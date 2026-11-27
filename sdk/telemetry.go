package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// TelemetryConfig selects dedicated OTLP export, the global provider, or an application provider.
type TelemetryConfig struct {
	Mode               string
	Provider           trace.TracerProvider
	ExporterHTTPClient *http.Client
	DisableRedaction   bool
	Redact             func(attribute.KeyValue) attribute.KeyValue
}
type mistralTelemetry struct {
	once     sync.Once
	provider trace.TracerProvider
	owned    *sdktrace.TracerProvider
	err      error
}
type telemetryAuth struct {
	key            string
	serviceAccount bool
	base           http.RoundTripper
}

func (a telemetryAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	key := a.key
	if a.serviceAccount {
		var err error
		key, err = ReadServiceAccountToken()
		if err != nil {
			return nil, err
		}
		if key == "" {
			return nil, fmt.Errorf("service account token path is no longer configured")
		}
	}
	if key == "" {
		return nil, fmt.Errorf("telemetry credentials required")
	}
	req = req.Clone(req.Context())
	req.Header.Set("Authorization", bearerHeader(key))
	return a.base.RoundTrip(req)
}

var defaultOTelProviderType = reflect.TypeOf(otel.GetTracerProvider())

func (t *mistralTelemetry) configure(ctx context.Context, c *MistralClient, config TelemetryConfig, automatic bool) error {
	if config.Provider != nil {
		t.provider = config.Provider
		return nil
	}
	mode := strings.ToLower(strings.TrimSpace(config.Mode))
	switch mode {
	case "false", "":
		return nil
	case "global":
		t.provider = otel.GetTracerProvider()
		return nil
	case "dedicated":
		if automatic && reflect.TypeOf(otel.GetTracerProvider()) != defaultOTelProviderType {
			return nil
		}
	default:
		return fmt.Errorf("unrecognized Mistral telemetry provider %q", config.Mode)
	}
	endpoint := strings.TrimSpace(os.Getenv("MISTRAL_OTLP_TRACES_ENDPOINT"))
	if endpoint == "" {
		endpoint = strings.TrimRight(c.endpoint, "/") + "/telemetry/v1/traces"
	}
	u, err := url.Parse(endpoint)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" || u.User != nil {
		return fmt.Errorf("invalid OTLP trace endpoint")
	}
	serviceAccount := !c.explicitAPIKey && os.Getenv(ServiceAccountTokenPathEnv) != ""
	key := c.apiKey
	if !c.explicitAPIKey && !serviceAccount {
		key = os.Getenv("MISTRAL_API_KEY")
	}
	if !serviceAccount && key == "" {
		return fmt.Errorf("Mistral telemetry requires an API key or service account token")
	}
	client := http.Client{Timeout: 10 * time.Second}
	if config.ExporterHTTPClient != nil {
		client = *config.ExporterHTTPClient
	}
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	transport := client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	client.Transport = telemetryAuth{key: key, serviceAccount: serviceAccount, base: transport}
	exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(endpoint), otlptracehttp.WithHTTPClient(&client))
	if err != nil {
		return err
	}
	var exp sdktrace.SpanExporter = exporter
	if !config.DisableRedaction {
		mask := config.Redact
		if mask == nil {
			mask = redactMistralAttribute
		}
		exp = &redactingTraceExporter{SpanExporter: exp, mask: mask}
	}
	t.owned = sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp), sdktrace.WithResource(resource.NewSchemaless(attribute.String("service.name", "mistralai_sdk"))))
	t.provider = t.owned
	runtime.AddCleanup(t, func(provider *sdktrace.TracerProvider) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = provider.Shutdown(ctx)
	}, t.owned)
	return nil
}
func (c *MistralClient) WithTelemetry(ctx context.Context, config TelemetryConfig) (*MistralClient, error) {
	copy := *c
	t := &mistralTelemetry{}
	if config.Mode == "" && config.Provider == nil {
		config.Mode = "dedicated"
	}
	t.once.Do(func() { t.err = t.configure(ctx, c, config, false) })
	if t.err != nil {
		return nil, t.err
	}
	copy.telemetry = t
	return &copy, nil
}
func (c *MistralClient) resolveTelemetry(ctx context.Context) *mistralTelemetry {
	t := c.telemetry
	if t == nil {
		return nil
	}
	t.once.Do(func() {
		t.err = t.configure(ctx, c, TelemetryConfig{Mode: os.Getenv("MISTRAL_SDK_TELEMETRY")}, true)
		if t.err != nil {
			log.Printf("Mistral SDK telemetry disabled: %v", t.err)
		}
	})
	return t
}
func (c *MistralClient) TelemetryTracer(name string) (trace.Tracer, error) {
	t := c.resolveTelemetry(c.requestContext())
	if t == nil || t.provider == nil {
		return nil, fmt.Errorf("telemetry is disabled")
	}
	if name == "" {
		name = "mistralai_sdk_tracer"
	}
	return t.provider.Tracer(name), nil
}
func (c *MistralClient) ShutdownTelemetry(ctx context.Context) error {
	if t := c.resolveTelemetry(ctx); t != nil && t.owned != nil {
		return t.owned.Shutdown(ctx)
	}
	return nil
}
func (c *MistralClient) traceRequest(req *http.Request) (*http.Request, trace.Span) {
	t := c.resolveTelemetry(req.Context())
	if t == nil || t.provider == nil {
		return req, nil
	}
	operation := ""
	switch strings.Trim(req.URL.Path, "/") {
	case "v1/chat/completions":
		operation = "chat"
	case "v1/fim/completions":
		operation = "text_completion"
	case "v1/agents/completions":
		operation = "invoke_agent"
	case "v1/embeddings":
		operation = "embeddings"
	case "v1/ocr":
		operation = "generate_content"
	}
	name := req.Method
	attrs := []attribute.KeyValue{attribute.String("http.request.method", req.Method), attribute.String("server.address", req.URL.Hostname()), attribute.String("url.scheme", req.URL.Scheme)}
	if operation != "" {
		name = operation
		attrs = append(attrs, attribute.String("gen_ai.operation.name", operation), attribute.String("gen_ai.provider.name", "mistral_ai"))
		if req.GetBody != nil {
			reader, err := req.GetBody()
			if err == nil {
				var fields map[string]json.RawMessage
				if json.NewDecoder(io.LimitReader(reader, 1<<20)).Decode(&fields) == nil {
					var model string
					if json.Unmarshal(fields["model"], &model) == nil && model != "" {
						name += " " + model
						attrs = append(attrs, attribute.String("gen_ai.request.model", model))
					}
					for field, key := range map[string]string{"messages": "gen_ai.input.messages", "tools": "gen_ai.tool.definitions"} {
						if data := fields[field]; len(data) > 0 {
							attrs = append(attrs, attribute.String(key, string(data)))
						}
					}
				}
				reader.Close()
			}
		}
	}
	ctx, span := t.provider.Tracer("mistralai_sdk_tracer").Start(req.Context(), name, trace.WithSpanKind(trace.SpanKindClient), trace.WithAttributes(attrs...))
	return req.WithContext(ctx), span
}
func finishRequestSpan(span trace.Span, response *http.Response, err error) {
	if span == nil {
		return
	}
	if response != nil {
		span.SetAttributes(attribute.Int("http.response.status_code", response.StatusCode))
		if response.StatusCode >= 400 {
			span.SetStatus(codes.Error, response.Status)
		}
	}
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	span.End()
}

type tracedResponseBody struct {
	io.ReadCloser
	span     trace.Span
	response *http.Response
	mu       sync.Mutex
	capture  bool
	data     bytes.Buffer
	once     sync.Once
}

func (b *tracedResponseBody) Read(p []byte) (int, error) {
	n, err := b.ReadCloser.Read(p)
	b.mu.Lock()
	if b.capture && n > 0 {
		if b.data.Len()+n > 1<<20 {
			b.capture = false
			b.data.Reset()
		} else {
			b.data.Write(p[:n])
		}
	}
	b.mu.Unlock()
	if err == io.EOF {
		b.finish(nil)
	} else if err != nil {
		b.finish(err)
	}
	return n, err
}
func (b *tracedResponseBody) Close() error { err := b.ReadCloser.Close(); b.finish(err); return err }
func (b *tracedResponseBody) finish(err error) {
	b.once.Do(func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		if b.capture && b.span.IsRecording() {
			var result struct {
				Model string
				ID    string
				Usage struct {
					PromptTokens     *int `json:"prompt_tokens"`
					CompletionTokens *int `json:"completion_tokens"`
				}
				Choices json.RawMessage
			}
			if json.Unmarshal(b.data.Bytes(), &result) == nil {
				if result.Model != "" {
					b.span.SetAttributes(attribute.String("gen_ai.response.model", result.Model))
				}
				if result.ID != "" {
					b.span.SetAttributes(attribute.String("gen_ai.response.id", result.ID))
				}
				if result.Usage.PromptTokens != nil {
					b.span.SetAttributes(attribute.Int("gen_ai.usage.input_tokens", *result.Usage.PromptTokens))
				}
				if result.Usage.CompletionTokens != nil {
					b.span.SetAttributes(attribute.Int("gen_ai.usage.output_tokens", *result.Usage.CompletionTokens))
				}
				if len(result.Choices) > 0 {
					b.span.SetAttributes(attribute.String("gen_ai.output.messages", string(result.Choices)))
				}
			}
		}
		finishRequestSpan(b.span, b.response, err)
	})
}
