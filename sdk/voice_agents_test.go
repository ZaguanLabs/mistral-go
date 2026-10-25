package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type testVoiceTransport struct {
	starts, closes atomic.Int32
	update         []byte
	emit           func(VoiceEvent)
}

func (t *testVoiceTransport) Start(context.Context) error                 { t.starts.Add(1); return nil }
func (t *testVoiceTransport) WaitForAgent(context.Context) error          { return nil }
func (t *testVoiceTransport) SendAudio(context.Context, []byte) error     { return nil }
func (t *testVoiceTransport) SetInputEnabled(context.Context, bool) error { return nil }
func (t *testVoiceTransport) UpdateSession(_ context.Context, b []byte) ([]byte, error) {
	t.update = append([]byte(nil), b...)
	return []byte(`{"type":"session.updated","session":{"tools":[]}}`), nil
}
func (t *testVoiceTransport) AddFunctionHandlers([]*VoiceFunctionHandler) {}
func (t *testVoiceTransport) Close() error                                { t.closes.Add(1); return nil }

func TestPython340VoiceLifecycle(t *testing.T) {
	var creates, deletes atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			creates.Add(1)
			w.Write([]byte(`{"id":"session","participant_token":"secret","server_url":"wss://media.test"}`))
		case "DELETE":
			deletes.Add(1)
			w.WriteHeader(204)
		default:
			t.Errorf("unexpected %s", r.Method)
		}
	}))
	defer server.Close()
	client := NewMistralClient("test", server.URL, 1, 0)
	if _, err := client.InitVoiceAgentSession(context.Background(), &AudioSessionCreateRequest{}, VoiceAgentSessionOptions{}); err == nil || creates.Load() != 0 {
		t.Fatal("missing transport must fail before allocation")
	}
	transport := &testVoiceTransport{}
	session, err := client.InitVoiceAgentSession(context.Background(), &AudioSessionCreateRequest{}, VoiceAgentSessionOptions{TransportFactory: func(_ *AudioSessionCreateResponse, emit func(VoiceEvent), _ []*VoiceFunctionHandler) (VoiceAgentTransport, error) {
		transport.emit = emit
		return transport, nil
	}})
	if err != nil {
		t.Fatal(err)
	}
	var disconnected atomic.Int32
	session.On("disconnected", func(VoiceEvent) {
		disconnected.Add(1)
		if err := session.Close(context.Background(), "again"); err != nil {
			t.Error(err)
		}
	})
	for i := 0; i < 2; i++ {
		if err := session.Connect(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
	if transport.starts.Load() != 1 {
		t.Fatal("connect not idempotent")
	}
	transport.emit(VoiceEvent{Name: "transcript", Data: VoiceAgentTranscriptEvent{ID: "turn", Content: "hel"}})
	transport.emit(VoiceEvent{Name: "transcript", Data: VoiceAgentTranscriptEvent{ID: "turn", Content: "hello", Final: true}})
	if got := session.Transcript(); len(got) != 1 || got[0].Content != "hello" {
		t.Fatal(got)
	}
	if _, err := session.Update(context.Background(), VoiceSessionUpdate{Tools: []any{}}); err != nil {
		t.Fatal(err)
	}
	if string(transport.update) != `{"tools":[]}` {
		t.Fatal(string(transport.update))
	}
	for i := 0; i < 2; i++ {
		if err := session.Close(context.Background(), "normal"); err != nil {
			t.Fatal(err)
		}
	}
	if err := session.Wait(context.Background()); err != nil {
		t.Fatal(err)
	}
	if transport.closes.Load() != 1 || deletes.Load() != 1 || disconnected.Load() != 1 {
		t.Fatal("close not idempotent")
	}
	if session.Connect(context.Background()) == nil {
		t.Fatal("closed session reconnected")
	}
}
func TestPython340VoiceHandlersAndSecurity(t *testing.T) {
	h := &VoiceFunctionHandler{Name: "echo", Function: func(ctx context.Context, input map[string]any) (any, error) { return input, nil }}
	data, err := json.Marshal(h)
	if err != nil || !strings.Contains(string(data), `"timeout_ms":15000`) {
		t.Fatal(string(data), err)
	}
	if _, err := h.Invoke(context.Background(), json.RawMessage(`[]`)); err == nil {
		t.Fatal("array accepted")
	}
	h.Timeout = time.Millisecond
	h.Function = func(ctx context.Context, _ map[string]any) (any, error) { <-ctx.Done(); return nil, nil }
	if _, err := h.Invoke(context.Background(), nil); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
	for _, raw := range []string{"https://api.test", "http://127.0.0.1:8000", "http://localhost:8000"} {
		if err := requireSecureVoiceURL(raw, "https", "http"); err != nil {
			t.Fatal(err)
		}
	}
	for _, raw := range []string{"http://api.test", "http://localhost.evil.test", "//api.test"} {
		if requireSecureVoiceURL(raw, "https", "http") == nil {
			t.Fatal(raw)
		}
	}
}
