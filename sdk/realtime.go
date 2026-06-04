package sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"nhooyr.io/websocket"
)

const realtimeTranscriptionPath = "v1/audio/transcriptions/realtime"

type AudioEncoding string

const (
	AudioEncodingPCMS16LE AudioEncoding = "pcm_s16le"
)

type AudioFormat struct {
	Encoding   AudioEncoding `json:"encoding"`
	SampleRate int           `json:"sample_rate"`
}

type RealtimeTranscriptionConnectParams struct {
	AudioFormat            *AudioFormat `json:"audio_format,omitempty"`
	TargetStreamingDelayMS *int         `json:"target_streaming_delay_ms,omitempty"`
	Headers                http.Header  `json:"-"`
}

type RealtimeTranscriptionSession struct {
	RequestID              string      `json:"request_id"`
	Model                  string      `json:"model"`
	AudioFormat            AudioFormat `json:"audio_format"`
	TargetStreamingDelayMS *int        `json:"target_streaming_delay_ms,omitempty"`
}

type RealtimeTranscriptionErrorDetail struct {
	Message any     `json:"message,omitempty"`
	Type    string  `json:"type,omitempty"`
	Code    *string `json:"code,omitempty"`
	Param   *string `json:"param,omitempty"`
}

type RealtimeTranscriptionEvent struct {
	Type          string                            `json:"type,omitempty"`
	Session       *RealtimeTranscriptionSession     `json:"session,omitempty"`
	Error         *RealtimeTranscriptionErrorDetail `json:"error,omitempty"`
	Text          string                            `json:"text,omitempty"`
	Model         string                            `json:"model,omitempty"`
	Language      *string                           `json:"language,omitempty"`
	AudioLanguage string                            `json:"audio_language,omitempty"`
	Start         *float64                          `json:"start,omitempty"`
	End           *float64                          `json:"end,omitempty"`
	SpeakerID     *string                           `json:"speaker_id,omitempty"`
	Usage         *UsageInfo                        `json:"usage,omitempty"`
	Segments      []TranscriptionSegment            `json:"segments,omitempty"`
	Raw           map[string]any                    `json:"-"`
	Err           error                             `json:"-"`
}

type RealtimeConnection struct {
	conn    *websocket.Conn
	session RealtimeTranscriptionSession
	initial []RealtimeTranscriptionEvent
	closed  bool
	mu      sync.Mutex
}

func (c *MistralClient) RealtimeTranscriptionConnect(ctx context.Context, model string, params *RealtimeTranscriptionConnectParams) (*RealtimeConnection, error) {
	wsURL, err := c.realtimeTranscriptionURL(model)
	if err != nil {
		return nil, err
	}

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.apiKey)
	headers.Set("User-Agent", UserAgent)
	if params != nil {
		for key, values := range params.Headers {
			for _, value := range values {
				headers.Add(key, value)
			}
		}
	}

	conn, _, err := websocket.Dial(ctx, wsURL, &websocket.DialOptions{
		HTTPHeader: headers,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect realtime transcription websocket: %w", err)
	}

	session, initial, err := recvRealtimeHandshake(ctx, conn)
	if err != nil {
		_ = conn.Close(websocket.StatusInternalError, "handshake failed")
		return nil, err
	}

	realtime := &RealtimeConnection{
		conn:    conn,
		session: session,
		initial: initial,
	}

	if params != nil && (params.AudioFormat != nil || params.TargetStreamingDelayMS != nil) {
		if err := realtime.UpdateSession(ctx, params.AudioFormat, params.TargetStreamingDelayMS); err != nil {
			_ = realtime.Close(websocket.StatusInternalError, "session update failed")
			return nil, err
		}
	}

	return realtime, nil
}

func (c *MistralClient) RealtimeTranscribeStream(ctx context.Context, model string, audio io.Reader, params *RealtimeTranscriptionConnectParams) (<-chan RealtimeTranscriptionEvent, error) {
	conn, err := c.RealtimeTranscriptionConnect(ctx, model, params)
	if err != nil {
		return nil, err
	}

	out := make(chan RealtimeTranscriptionEvent)
	go func() {
		defer close(out)
		defer conn.Close(websocket.StatusNormalClosure, "")

		sendErr := make(chan error, 1)
		go func() {
			buf := make([]byte, 4096)
			for {
				n, readErr := audio.Read(buf)
				if n > 0 {
					if err := conn.SendAudio(ctx, buf[:n]); err != nil {
						sendErr <- err
						return
					}
				}
				if readErr == io.EOF {
					if err := conn.FlushAudio(ctx); err != nil {
						sendErr <- err
						return
					}
					sendErr <- conn.EndAudio(ctx)
					return
				}
				if readErr != nil {
					sendErr <- readErr
					return
				}
			}
		}()

		events := conn.Events(ctx)
		for {
			select {
			case err := <-sendErr:
				if err != nil {
					out <- RealtimeTranscriptionEvent{Type: "error", Err: err}
					return
				}
				sendErr = nil
			case event, ok := <-events:
				if !ok {
					return
				}
				out <- event
				if event.Error != nil || event.Type == "transcription.done" {
					return
				}
			case <-ctx.Done():
				out <- RealtimeTranscriptionEvent{Type: "error", Err: ctx.Err()}
				return
			}
		}
	}()

	return out, nil
}

func (r *RealtimeConnection) Session() RealtimeTranscriptionSession {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.session
}

func (r *RealtimeConnection) RequestID() string {
	return r.Session().RequestID
}

func (r *RealtimeConnection) SendAudio(ctx context.Context, audio []byte) error {
	if err := r.ensureOpen(); err != nil {
		return err
	}
	message := map[string]any{
		"type":  "input_audio.append",
		"audio": base64.StdEncoding.EncodeToString(audio),
	}
	return r.writeJSON(ctx, message)
}

func (r *RealtimeConnection) FlushAudio(ctx context.Context) error {
	if err := r.ensureOpen(); err != nil {
		return err
	}
	return r.writeJSON(ctx, map[string]any{"type": "input_audio.flush"})
}

func (r *RealtimeConnection) EndAudio(ctx context.Context) error {
	if err := r.ensureOpen(); err != nil {
		return nil
	}
	return r.writeJSON(ctx, map[string]any{"type": "input_audio.end"})
}

func (r *RealtimeConnection) UpdateSession(ctx context.Context, audioFormat *AudioFormat, targetStreamingDelayMS *int) error {
	if audioFormat == nil && targetStreamingDelayMS == nil {
		return fmt.Errorf("at least one session field must be provided")
	}
	if err := r.ensureOpen(); err != nil {
		return err
	}
	session := optionalRequestMap(map[string]any{
		"audio_format":              audioFormat,
		"target_streaming_delay_ms": targetStreamingDelayMS,
	})
	return r.writeJSON(ctx, map[string]any{
		"type":    "session.update",
		"session": session,
	})
}

func (r *RealtimeConnection) Events(ctx context.Context) <-chan RealtimeTranscriptionEvent {
	out := make(chan RealtimeTranscriptionEvent)
	go func() {
		defer close(out)
		for _, event := range r.initial {
			r.applySessionUpdate(event)
			out <- event
		}
		r.initial = nil

		for {
			event, err := r.ReadEvent(ctx)
			if err != nil {
				if !r.isClosed() && ctx.Err() == nil {
					out <- RealtimeTranscriptionEvent{Type: "error", Err: err}
				}
				return
			}
			out <- event
		}
	}()
	return out
}

func (r *RealtimeConnection) ReadEvent(ctx context.Context) (RealtimeTranscriptionEvent, error) {
	_, payload, err := r.conn.Read(ctx)
	if err != nil {
		return RealtimeTranscriptionEvent{}, err
	}
	event := parseRealtimeEvent(payload)
	r.applySessionUpdate(event)
	return event, nil
}

func (r *RealtimeConnection) Close(status websocket.StatusCode, reason string) error {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil
	}
	r.closed = true
	r.mu.Unlock()
	return r.conn.Close(status, reason)
}

func (r *RealtimeConnection) writeJSON(ctx context.Context, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.conn.Write(ctx, websocket.MessageText, data)
}

func (r *RealtimeConnection) ensureOpen() error {
	if r.isClosed() {
		return fmt.Errorf("realtime connection is closed")
	}
	return nil
}

func (r *RealtimeConnection) isClosed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closed
}

func (r *RealtimeConnection) applySessionUpdate(event RealtimeTranscriptionEvent) {
	if event.Session == nil {
		return
	}
	r.mu.Lock()
	r.session = *event.Session
	r.mu.Unlock()
}

func (c *MistralClient) realtimeTranscriptionURL(model string) (string, error) {
	base, err := url.Parse(c.endpoint)
	if err != nil {
		return "", err
	}
	switch base.Scheme {
	case "https":
		base.Scheme = "wss"
	case "http":
		base.Scheme = "ws"
	case "ws", "wss":
	default:
		return "", fmt.Errorf("unsupported endpoint scheme %q", base.Scheme)
	}
	base.Path = strings.TrimRight(base.Path, "/") + "/" + realtimeTranscriptionPath
	query := base.Query()
	query.Set("model", model)
	base.RawQuery = query.Encode()
	return base.String(), nil
}

func recvRealtimeHandshake(ctx context.Context, conn *websocket.Conn) (RealtimeTranscriptionSession, []RealtimeTranscriptionEvent, error) {
	var initial []RealtimeTranscriptionEvent
	for {
		_, payload, err := conn.Read(ctx)
		if err != nil {
			return RealtimeTranscriptionSession{}, initial, fmt.Errorf("failed to receive realtime handshake: %w", err)
		}
		event := parseRealtimeEvent(payload)
		initial = append(initial, event)
		if event.Error != nil {
			return RealtimeTranscriptionSession{}, initial, realtimeErrorMessage(event)
		}
		if event.Type == "session.created" && event.Session != nil {
			return *event.Session, initial, nil
		}
	}
}

func parseRealtimeEvent(payload []byte) RealtimeTranscriptionEvent {
	var raw map[string]any
	if err := json.Unmarshal(payload, &raw); err != nil {
		return RealtimeTranscriptionEvent{Err: fmt.Errorf("invalid realtime event JSON: %w", err)}
	}
	var event RealtimeTranscriptionEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return RealtimeTranscriptionEvent{Raw: raw, Err: fmt.Errorf("invalid realtime event payload: %w", err)}
	}
	event.Raw = raw
	return event
}

func realtimeErrorMessage(event RealtimeTranscriptionEvent) error {
	if event.Error == nil {
		return fmt.Errorf("realtime transcription error")
	}
	switch msg := event.Error.Message.(type) {
	case string:
		if msg != "" {
			return fmt.Errorf("realtime transcription error: %s", msg)
		}
	case map[string]any:
		if detail, ok := msg["detail"].(string); ok && detail != "" {
			return fmt.Errorf("realtime transcription error: %s", detail)
		}
	}
	return fmt.Errorf("realtime transcription error")
}
