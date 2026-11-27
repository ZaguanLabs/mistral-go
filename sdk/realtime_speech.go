package sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type RealtimeSpeechConnectParams struct {
	Session *RealtimeSpeechSessionUpdatePayload
	Headers http.Header
}
type RealtimeSpeechEvent struct {
	Type      string                     `json:"type"`
	Session   *RealtimeSpeechSession     `json:"session,omitempty"`
	AudioData string                     `json:"audio_data,omitempty"`
	Usage     *UsageInfo                 `json:"usage,omitempty"`
	Error     *RealtimeSpeechErrorDetail `json:"error,omitempty"`
	Raw       json.RawMessage            `json:"-"`
	Err       error                      `json:"-"`
}

func (e RealtimeSpeechEvent) AudioBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(e.AudioData)
}

type RealtimeSpeechConnection struct {
	conn    *websocket.Conn
	mu      sync.Mutex
	initial []RealtimeSpeechEvent
	session RealtimeSpeechSession
	closed  bool
}

// RealtimeSpeechConnect waits for session.created and retains handshake events for ReadEvent.
func (c *MistralClient) RealtimeSpeechConnect(ctx context.Context, model string, params *RealtimeSpeechConnectParams) (*RealtimeSpeechConnection, error) {
	if c.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.timeout)
		defer cancel()
	}
	if params == nil {
		params = &RealtimeSpeechConnectParams{}
	}
	conn, err := c.dialRealtime(ctx, model, "v1/audio/speech/realtime", params.Headers)
	if err != nil {
		return nil, err
	}
	speech := &RealtimeSpeechConnection{conn: conn}
	for {
		event, err := speech.read(ctx)
		if err != nil {
			conn.CloseNow()
			return nil, err
		}
		speech.initial = append(speech.initial, event)
		if event.Type == "error" {
			conn.CloseNow()
			return nil, fmt.Errorf("realtime speech handshake error: %s", event.Raw)
		}
		if event.Type == "session.created" && event.Session != nil {
			break
		}
	}
	if params.Session != nil {
		if err = speech.UpdateSession(ctx, *params.Session); err != nil {
			conn.CloseNow()
			return nil, err
		}
	}
	return speech, nil
}
func (r *RealtimeSpeechConnection) Session() RealtimeSpeechSession {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.session
}
func (r *RealtimeSpeechConnection) RequestID() string { return r.Session().RequestID }
func (r *RealtimeSpeechConnection) IsClosed() bool    { r.mu.Lock(); defer r.mu.Unlock(); return r.closed }
func (r *RealtimeSpeechConnection) SendText(ctx context.Context, text string) error {
	return wsjson.Write(ctx, r.conn, RealtimeSpeechInputTextAppend{Text: text})
}
func (r *RealtimeSpeechConnection) Flush(ctx context.Context, text *string) error {
	return wsjson.Write(ctx, r.conn, RealtimeSpeechInputTextFlush{Text: text})
}
func (r *RealtimeSpeechConnection) End(ctx context.Context) error {
	if r.IsClosed() {
		return nil
	}
	return wsjson.Write(ctx, r.conn, RealtimeSpeechInputTextEnd{})
}
func (r *RealtimeSpeechConnection) UpdateSession(ctx context.Context, p RealtimeSpeechSessionUpdatePayload) error {
	if p.Voice == nil && p.RefAudio == nil && p.StreamingMode == nil && p.ResponseFormat == nil && p.Expressiveness == nil && p.PhonemeHints == nil {
		return fmt.Errorf("at least one speech session field is required")
	}
	return wsjson.Write(ctx, r.conn, RealtimeSpeechSessionUpdateMessage{Session: p})
}
func (r *RealtimeSpeechConnection) Close(status websocket.StatusCode, reason string) error {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil
	}
	r.closed = true
	r.mu.Unlock()
	return r.conn.Close(status, reason)
}
func (r *RealtimeSpeechConnection) read(ctx context.Context) (RealtimeSpeechEvent, error) {
	_, data, err := r.conn.Read(ctx)
	if err != nil {
		return RealtimeSpeechEvent{}, err
	}
	var event RealtimeSpeechEvent
	if err = json.Unmarshal(data, &event); err != nil {
		event = RealtimeSpeechEvent{Err: err}
	}
	event.Raw = data
	if (event.Type == "session.created" || event.Type == "session.updated") && event.Session != nil {
		r.mu.Lock()
		r.session = *event.Session
		r.mu.Unlock()
	}
	return event, nil
}
func (r *RealtimeSpeechConnection) ReadEvent(ctx context.Context) (RealtimeSpeechEvent, error) {
	r.mu.Lock()
	if len(r.initial) > 0 {
		event := r.initial[0]
		r.initial = r.initial[1:]
		r.mu.Unlock()
		return event, nil
	}
	r.mu.Unlock()
	return r.read(ctx)
}

// RealtimeSpeakStream sends each fragment as it arrives and ends input when text closes.
// Cancel ctx when abandoning the output stream.
func (c *MistralClient) RealtimeSpeakStream(ctx context.Context, model string, text <-chan string, params *RealtimeSpeechConnectParams) (<-chan RealtimeSpeechEvent, error) {
	conn, err := c.RealtimeSpeechConnect(ctx, model, params)
	if err != nil {
		return nil, err
	}
	streamCtx, cancel := context.WithCancel(ctx)
	sendError := make(chan error, 1)
	go func() {
		for {
			select {
			case <-streamCtx.Done():
				return
			case fragment, ok := <-text:
				var err error
				if ok {
					err = conn.SendText(streamCtx, fragment)
				} else {
					err = conn.End(streamCtx)
				}
				if err != nil {
					sendError <- err
					cancel()
					return
				}
				if !ok {
					return
				}
			}
		}
	}()
	events := make(chan RealtimeSpeechEvent)
	go func() {
		defer close(events)
		defer conn.Close(websocket.StatusNormalClosure, "")
		defer cancel()
		for {
			event, err := conn.ReadEvent(streamCtx)
			if err != nil {
				select {
				case err = <-sendError:
				default:
				}
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					return
				}
				event = RealtimeSpeechEvent{Err: err}
			}
			select {
			case events <- event:
			case <-ctx.Done():
				return
			}
			if err != nil || event.Type == "error" || event.Type == "speech.audio.done" {
				return
			}
		}
	}()
	return events, nil
}
