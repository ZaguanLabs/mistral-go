package sdk

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math"
	"net/http"
	"net/url"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/klauspost/compress/zstd"
	"github.com/vmihailenco/msgpack/v5"
)

// WorkflowBlobStore stores immutable payloads under SDK-generated content keys.
// PutIfAbsent must treat an existing object as success.
type WorkflowBlobStore interface {
	PutIfAbsent(context.Context, string, []byte) error
	Get(context.Context, string) ([]byte, error)
}
type WorkflowCompression struct {
	Level        int
	MinSizeBytes int
}
type WorkflowEncodingConfig struct {
	Namespace                   string
	EncryptionMode              string // "full" or "partial"
	MainKeyHex, SecondaryKeyHex string
	Compression                 *WorkflowCompression
	BlobStore                   WorkflowBlobStore
	OffloadMinSizeBytes         int
}
type workflowEncoding struct {
	namespace       string
	mode            string
	main, secondary cipher.AEAD
	hashKey         []byte
	compress        *WorkflowCompression
	store           WorkflowBlobStore
	offloadMin      int
}

func (c *MistralClient) WithWorkflowEncoding(ctx context.Context, config WorkflowEncodingConfig) (*MistralClient, error) {
	enc := &workflowEncoding{namespace: config.Namespace, mode: config.EncryptionMode, store: config.BlobStore, offloadMin: config.OffloadMinSizeBytes}
	if enc.offloadMin < 0 {
		return nil, fmt.Errorf("negative offloading threshold")
	}
	if config.Compression != nil {
		compression := *config.Compression
		if compression.Level == 0 {
			compression.Level = 3
		}
		if compression.MinSizeBytes < 0 || compression.Level < 1 || compression.Level > 22 {
			return nil, fmt.Errorf("invalid workflow compression settings")
		}
		enc.compress = &compression
	}
	if enc.mode != "" || config.MainKeyHex != "" || config.SecondaryKeyHex != "" {
		if enc.mode != "full" && enc.mode != "partial" {
			return nil, fmt.Errorf("workflow encryption mode must be full or partial")
		}
		key, err := hex.DecodeString(config.MainKeyHex)
		if err != nil {
			return nil, fmt.Errorf("invalid main key encoding")
		}
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		enc.main, err = cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}
		enc.hashKey = workflowMAC(key, []byte("schedule-input-hash"))
		if config.SecondaryKeyHex != "" {
			key, err = hex.DecodeString(config.SecondaryKeyHex)
			if err != nil {
				return nil, fmt.Errorf("invalid secondary key encoding")
			}
			block, err = aes.NewCipher(key)
			if err != nil {
				return nil, err
			}
			enc.secondary, err = cipher.NewGCM(block)
			if err != nil {
				return nil, err
			}
		}
	}
	if enc.namespace == "" {
		who, err := requestTyped[struct {
			Namespace string `json:"namespace"`
		}](c.WithContext(ctx), http.MethodGet, nil, "v1/workflows/workers/whoami")
		if err != nil {
			return nil, err
		}
		if who.Namespace == "" {
			return nil, fmt.Errorf("worker namespace missing")
		}
		enc.namespace = who.Namespace
	}
	result := *c
	result.workflowEncoding = enc
	return &result, nil
}
func workflowMAC(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
func (e *workflowEncoding) seal(data []byte) ([]byte, error) {
	nonce := make([]byte, e.main.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return e.main.Seal(nonce, nonce, data, nil), nil
}
func (e *workflowEncoding) open(data []byte) ([]byte, error) {
	if e.main == nil || len(data) < e.main.NonceSize() {
		return nil, fmt.Errorf("missing workflow key or encrypted data")
	}
	n := e.main.NonceSize()
	plain, err := e.main.Open(nil, data[:n], data[n:], nil)
	if err != nil && e.secondary != nil {
		return e.secondary.Open(nil, data[:n], data[n:], nil)
	}
	return plain, err
}
func (e *workflowEncoding) transformFields(value any, encrypt bool) (bool, error) {
	var changed bool
	var children []any
	switch v := value.(type) {
	case map[string]any:
		if v["field_type"] == "__encrypted_str__" {
			s, ok := v["data"].(string)
			if !ok {
				return false, fmt.Errorf("workflow encrypted field requires string data")
			}
			if encrypt {
				data, err := e.seal([]byte(s))
				if err != nil {
					return false, err
				}
				v["data"] = base64.StdEncoding.EncodeToString(data)
			} else {
				data, err := base64.StdEncoding.DecodeString(s)
				if err != nil {
					return false, err
				}
				data, err = e.open(data)
				if err != nil {
					return false, err
				}
				v["data"] = string(data)
			}
			return true, nil
		}
		for _, child := range v {
			children = append(children, child)
		}
	case []any:
		children = v
	}
	for _, child := range children {
		ok, err := e.transformFields(child, encrypt)
		if err != nil {
			return false, err
		}
		changed = changed || ok
	}
	return changed, nil
}

type workflowPackedPayload struct {
	Compression struct {
		Algorithm string `msgpack:"algorithm"`
		Level     int    `msgpack:"level"`
	} `msgpack:"compression"`
	Payload []byte `msgpack:"payload"`
}

func workflowJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(buf.Bytes(), []byte("\n")), nil
}
func workflowValue(data []byte) (any, error) {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	var v any
	err := d.Decode(&v)
	return v, err
}

// scheduleCanonicalJSON matches Python json.dumps(sort_keys=True, ensure_ascii=False,
// separators=(",", ":")) after the request has been decoded from the wire.
func scheduleCanonicalJSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	var emit func(any) error
	quoted := func(s string) {
		buf.WriteByte('"')
		for _, r := range s {
			switch r {
			case '"', '\\':
				buf.WriteByte('\\')
				buf.WriteRune(r)
			case '\n':
				buf.WriteString(`\n`)
			case '\r':
				buf.WriteString(`\r`)
			case '\t':
				buf.WriteString(`\t`)
			case '\b':
				buf.WriteString(`\b`)
			case '\f':
				buf.WriteString(`\f`)
			default:
				if r < 32 {
					fmt.Fprintf(&buf, `\u%04x`, r)
				} else {
					buf.WriteRune(r)
				}
			}
		}
		buf.WriteByte('"')
	}
	emit = func(value any) error {
		switch value := value.(type) {
		case nil:
			buf.WriteString("null")
		case bool:
			buf.WriteString(strconv.FormatBool(value))
		case string:
			quoted(value)
		case json.Number:
			literal := value.String()
			if strings.ContainsAny(literal, ".eE") {
				number, err := value.Float64()
				if err != nil {
					return err
				}
				if abs := math.Abs(number); abs != 0 && (abs < 1e-4 || abs >= 1e16) {
					literal = strconv.FormatFloat(number, 'e', -1, 64)
				} else {
					literal = strconv.FormatFloat(number, 'f', -1, 64)
					if !strings.ContainsRune(literal, '.') {
						literal += ".0"
					}
				}
			} else if literal == "-0" {
				literal = "0"
			}
			buf.WriteString(literal)
		case []any:
			buf.WriteByte('[')
			for i, item := range value {
				if i > 0 {
					buf.WriteByte(',')
				}
				if err := emit(item); err != nil {
					return err
				}
			}
			buf.WriteByte(']')
		case map[string]any:
			keys := make([]string, 0, len(value))
			for key := range value {
				keys = append(keys, key)
			}
			sort.Strings(keys)
			buf.WriteByte('{')
			for i, key := range keys {
				if i > 0 {
					buf.WriteByte(',')
				}
				quoted(key)
				buf.WriteByte(':')
				if err := emit(value[key]); err != nil {
					return err
				}
			}
			buf.WriteByte('}')
		default:
			return fmt.Errorf("invalid JSON value %T", value)
		}
		return nil
	}
	err := emit(v)
	return buf.Bytes(), err
}

func (e *workflowEncoding) encode(ctx context.Context, value any, id string, schedule, create bool) (map[string]any, error) {
	raw, err := workflowJSON(value)
	if err != nil {
		return nil, err
	}
	value, err = workflowValue(raw)
	if err != nil {
		return nil, err
	}
	// Round trip through maps so struct field order also uses sorted JSON keys.
	raw, err = workflowJSON(value)
	if err != nil {
		return nil, err
	}
	envelope := map[string]any{"empty": false}
	opts := []string{}
	if schedule && create {
		canonical, err := scheduleCanonicalJSON(value)
		if err != nil {
			return nil, err
		}
		sum := sha256.Sum256(canonical)
		hash := sum[:]
		if e.main != nil {
			hash = workflowMAC(e.hashKey, canonical)
		}
		envelope["input_hash"] = hex.EncodeToString(hash)
	}
	if e.mode == "partial" {
		changed, err := e.transformFields(value, true)
		if err != nil {
			return nil, err
		}
		if changed {
			raw, err = workflowJSON(value)
			if err != nil {
				return nil, err
			}
			opts = append(opts, "encrypted-partial")
		}
	}
	if cfg := e.compress; cfg != nil && len(raw) >= cfg.MinSizeBytes {
		enc, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.EncoderLevelFromZstd(cfg.Level)), zstd.WithEncoderConcurrency(1))
		if err != nil {
			return nil, err
		}
		compressed := enc.EncodeAll(raw, nil)
		enc.Close()
		if len(compressed) < len(raw) {
			packed := workflowPackedPayload{Payload: compressed}
			packed.Compression.Algorithm = "zstd"
			packed.Compression.Level = cfg.Level
			raw, err = msgpack.Marshal(packed)
			if err != nil {
				return nil, err
			}
			opts = append(opts, "compressed")
		}
	}
	if !schedule && e.store != nil && len(raw) >= e.offloadMin {
		key := fmt.Sprintf("temporal-payload/%s/%s/sha256:%x", url.QueryEscape(e.namespace), url.QueryEscape(id), sha256.Sum256(raw))
		key = strings.ReplaceAll(key, "+", "%20")
		if err = e.store.PutIfAbsent(ctx, key, raw); err != nil {
			return nil, err
		}
		raw, err = json.Marshal(map[string]string{"key": key})
		if err != nil {
			return nil, err
		}
		opts = append(opts, "offloaded")
	}
	if e.mode == "full" {
		raw, err = e.seal(raw)
		if err != nil {
			return nil, err
		}
		opts = append(opts, "encrypted")
	}
	envelope["b64payload"] = base64.StdEncoding.EncodeToString(raw)
	envelope["encoding_options"] = opts
	return envelope, nil
}
func (e *workflowEncoding) decode(ctx context.Context, value map[string]any) (any, error) {
	encoded, ok := value["b64payload"].(string)
	if !ok {
		return nil, fmt.Errorf("missing b64payload")
	}
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var opts []string
	if options, ok := value["encoding_options"]; ok {
		b, err := json.Marshal(options)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(b, &opts); err != nil {
			return nil, err
		}
	}
	for i := len(opts) - 1; i >= 0; i-- {
		switch opts[i] {
		case "encrypted":
			data, err = e.open(data)
		case "offloaded":
			var ref struct {
				Key string `json:"key"`
			}
			if err = json.Unmarshal(data, &ref); err != nil {
				return nil, err
			}
			if e.store == nil {
				return nil, fmt.Errorf("workflow blob store required")
			}
			data, err = e.store.Get(ctx, ref.Key)
		case "compressed":
			var pack workflowPackedPayload
			if err = msgpack.Unmarshal(data, &pack); err != nil {
				return nil, err
			}
			if pack.Compression.Algorithm != "zstd" {
				return nil, fmt.Errorf("unsupported workflow compression %q", pack.Compression.Algorithm)
			}
			dec, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
			if err != nil {
				return nil, err
			}
			data, err = dec.DecodeAll(pack.Payload, nil)
			dec.Close()
			if err != nil {
				return nil, err
			}
		case "encrypted-partial":
			v, err := workflowValue(data)
			if err != nil {
				return nil, err
			}
			if _, err = e.transformFields(v, false); err != nil {
				return nil, err
			}
			data, err = workflowJSON(v)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported workflow encoding %q", opts[i])
		}
		if err != nil {
			return nil, err
		}
	}
	var decoded any
	if json.Unmarshal(data, &decoded) != nil {
		return data, nil
	}
	return decoded, nil
}

// encodeRequest applies the workflow protocol to one request body owned by the caller.
func (e *workflowEncoding) encodeRequest(ctx context.Context, method, path string, body map[string]any) (map[string]any, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 || parts[0] != "v1" || parts[1] != "workflows" {
		return body, nil
	}
	schedule := parts[2] == "schedules" && (len(parts) == 3 && method == "POST" || len(parts) == 4 && method == "PATCH")
	execute := method == "POST" && parts[len(parts)-1] == "execute" && (len(parts) == 4 || len(parts) == 5 && parts[2] == "registrations")
	invocation := len(parts) == 5 && parts[2] == "executions" && method == "POST" && (parts[4] == "signals" || parts[4] == "queries" || parts[4] == "updates")
	if !schedule && !execute && !invocation {
		return body, nil
	}
	// Copy through JSON to retain custom serializers and isolate nested schedule maps.
	raw, err := workflowJSON(body)
	if err != nil {
		return nil, err
	}
	v, err := workflowValue(raw)
	if err != nil {
		return nil, err
	}
	result, ok := v.(map[string]any)
	if !ok {
		return body, nil
	}
	holder := result
	if schedule {
		var ok bool
		holder, ok = result["schedule"].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("workflow schedule object required")
		}
	}
	input := holder["input"]
	if input == nil {
		return result, nil
	}
	id, _ := result["execution_id"].(string)
	if id == "" && execute {
		seed := parts[len(parts)-2]
		secondary, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		first := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(seed))
		second := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(hex.EncodeToString(secondary[:])))
		id = hex.EncodeToString(first[:]) + hex.EncodeToString(second[:])
		result["execution_id"] = id
	} else if id == "" && invocation {
		id = parts[3]
	}
	encoded, err := e.encode(ctx, input, id, schedule, schedule && len(parts) == 3)
	if err != nil {
		return nil, err
	}
	if execute {
		result["encoded_input"] = encoded
		holder["input"] = nil
	} else {
		holder["input"] = encoded
	}
	return result, nil
}
func (e *workflowEncoding) decodeResponse(ctx context.Context, method, path string, result any) error {
	body, ok := result.(map[string]any)
	if !ok {
		return nil
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 || parts[0] != "v1" || parts[1] != "workflows" {
		return nil
	}
	if method == "GET" && parts[2] == "schedules" && len(parts) <= 4 {
		defs, ok := body["schedules"].([]any)
		if !ok {
			defs = []any{body}
		}
		for _, def := range defs {
			entry, _ := def.(map[string]any)
			value, _ := entry["input"].(map[string]any)
			opts, _ := value["encoding_options"].([]any)
			if len(opts) > 0 {
				if decoded, err := e.decode(ctx, value); err == nil {
					if _, binary := decoded.([]byte); !binary {
						entry["input"] = decoded
					}
				}
			}
		}
	} else if method == "GET" && path == "v1/workflows/events/list" {
		events, _ := body["events"].([]any)
		for _, evt := range events {
			event, _ := evt.(map[string]any)
			if err := e.decodeEvent(ctx, event); err != nil {
				return err
			}
		}
	} else if (method == "GET" && parts[2] == "executions" && len(parts) == 4) || (method == "POST" && (parts[len(parts)-1] == "execute" || parts[2] == "executions" && (parts[len(parts)-1] == "updates" || parts[len(parts)-1] == "queries"))) {
		if encoded, ok := body["result"].(map[string]any); ok && encoded["b64payload"] != nil {
			decoded, err := e.decode(ctx, encoded)
			if err != nil {
				return err
			}
			body["result"] = decoded
		}
	}
	return nil
}
func (e *workflowEncoding) decodeEvent(ctx context.Context, event map[string]any) error {
	attributes, _ := event["attributes"].(map[string]any)
	for _, field := range attributes {
		payload, _ := field.(map[string]any)
		options, _ := payload["encoding_options"].([]any)
		if len(options) == 0 || (payload["type"] != "json" && payload["type"] != "json_patch") {
			continue
		}
		if patches, ok := payload["value"].([]any); ok && slices.Contains(options, any("encrypted-partial")) {
			for _, item := range patches {
				patch, _ := item.(map[string]any)
				marker, _ := patch["value"].(map[string]any)
				if marker["type"] == "__encrypted__" {
					str, _ := marker["value"].(string)
					data, err := base64.StdEncoding.DecodeString(str)
					if err != nil {
						return err
					}
					data, err = e.open(data)
					if err != nil {
						return err
					}
					var value any
					if err = json.Unmarshal(data, &value); err != nil {
						return err
					}
					patch["value"] = value
				}
			}
		} else {
			value, err := e.decode(ctx, map[string]any{"b64payload": payload["value"], "encoding_options": options})
			if err != nil {
				return err
			}
			payload["value"] = value
		}
		payload["encoding_options"] = []string{}
	}
	return nil
}
