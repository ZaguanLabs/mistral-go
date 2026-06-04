package sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"nhooyr.io/websocket"
)

func TestRealtimeTranscriptionConnectAndAudioMessages(t *testing.T) {
	messages := make(chan map[string]any, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/audio/transcriptions/realtime" {
			t.Errorf("expected realtime path, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("model") != "voxtral-mini-realtime" {
			t.Errorf("expected model query, got %s", r.URL.RawQuery)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("missing Authorization header")
		}
		if r.Header.Get("User-Agent") != UserAgent {
			t.Errorf("expected User-Agent %s, got %s", UserAgent, r.Header.Get("User-Agent"))
		}

		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			t.Errorf("accept websocket: %v", err)
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		ctx := r.Context()
		err = conn.Write(ctx, websocket.MessageText, []byte(`{"type":"session.created","session":{"request_id":"req-1","model":"voxtral-mini-realtime","audio_format":{"encoding":"pcm_s16le","sample_rate":16000}}}`))
		if err != nil {
			t.Errorf("write handshake: %v", err)
			return
		}

		for i := 0; i < 4; i++ {
			_, data, err := conn.Read(ctx)
			if err != nil {
				t.Errorf("read client message: %v", err)
				return
			}
			var msg map[string]any
			if err := json.Unmarshal(data, &msg); err != nil {
				t.Errorf("decode client message: %v", err)
				return
			}
			messages <- msg
		}
	}))
	defer server.Close()

	client := NewMistralClient("test-api-key", server.URL, DefaultMaxRetries, DefaultTimeout)
	delay := 120
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := client.RealtimeTranscriptionConnect(ctx, "voxtral-mini-realtime", &RealtimeTranscriptionConnectParams{
		AudioFormat:            &AudioFormat{Encoding: AudioEncodingPCMS16LE, SampleRate: 16000},
		TargetStreamingDelayMS: &delay,
	})
	if err != nil {
		t.Fatalf("RealtimeTranscriptionConnect failed: %v", err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	if conn.RequestID() != "req-1" {
		t.Fatalf("expected request id req-1, got %s", conn.RequestID())
	}
	if err := conn.SendAudio(ctx, []byte("pcm")); err != nil {
		t.Fatalf("SendAudio failed: %v", err)
	}
	if err := conn.FlushAudio(ctx); err != nil {
		t.Fatalf("FlushAudio failed: %v", err)
	}
	if err := conn.EndAudio(ctx); err != nil {
		t.Fatalf("EndAudio failed: %v", err)
	}

	update := <-messages
	if update["type"] != "session.update" {
		t.Fatalf("expected session.update, got %#v", update)
	}
	session, ok := update["session"].(map[string]any)
	if !ok {
		t.Fatalf("expected session payload, got %#v", update["session"])
	}
	if session["target_streaming_delay_ms"].(float64) != float64(delay) {
		t.Fatalf("expected delay %d, got %#v", delay, session["target_streaming_delay_ms"])
	}

	appendMsg := <-messages
	if appendMsg["type"] != "input_audio.append" {
		t.Fatalf("expected input_audio.append, got %#v", appendMsg)
	}
	decoded, err := base64.StdEncoding.DecodeString(appendMsg["audio"].(string))
	if err != nil {
		t.Fatalf("decode audio: %v", err)
	}
	if string(decoded) != "pcm" {
		t.Fatalf("expected pcm audio, got %q", string(decoded))
	}
	if flush := <-messages; flush["type"] != "input_audio.flush" {
		t.Fatalf("expected input_audio.flush, got %#v", flush)
	}
	if end := <-messages; end["type"] != "input_audio.end" {
		t.Fatalf("expected input_audio.end, got %#v", end)
	}
}

func TestRealtimeTranscribeStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			t.Errorf("accept websocket: %v", err)
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		ctx := r.Context()
		if err := conn.Write(ctx, websocket.MessageText, []byte(`{"type":"session.created","session":{"request_id":"req-2","model":"voxtral","audio_format":{"encoding":"pcm_s16le","sample_rate":16000}}}`)); err != nil {
			t.Errorf("write handshake: %v", err)
			return
		}

		sawAppend := false
		for {
			_, data, err := conn.Read(ctx)
			if err != nil {
				t.Errorf("read client message: %v", err)
				return
			}
			var msg map[string]any
			if err := json.Unmarshal(data, &msg); err != nil {
				t.Errorf("decode client message: %v", err)
				return
			}
			if msg["type"] == "input_audio.append" {
				sawAppend = true
			}
			if msg["type"] == "input_audio.end" {
				break
			}
		}
		if !sawAppend {
			t.Errorf("expected at least one input_audio.append message")
		}
		err = conn.Write(ctx, websocket.MessageText, []byte(`{"type":"transcription.text.delta","text":"hello"}`))
		if err != nil {
			t.Errorf("write delta: %v", err)
			return
		}
		err = conn.Write(ctx, websocket.MessageText, []byte(`{"type":"transcription.done","model":"voxtral","text":"hello","language":"en","usage":{"prompt_tokens":1,"total_tokens":1}}`))
		if err != nil {
			t.Errorf("write done: %v", err)
			return
		}
	}))
	defer server.Close()

	client := NewMistralClient("test-api-key", server.URL, DefaultMaxRetries, DefaultTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events, err := client.RealtimeTranscribeStream(ctx, "voxtral", strings.NewReader("pcm-data"), nil)
	if err != nil {
		t.Fatalf("RealtimeTranscribeStream failed: %v", err)
	}

	var text string
	for event := range events {
		if event.Err != nil && event.Err != io.EOF {
			t.Fatalf("unexpected event error: %v", event.Err)
		}
		if event.Type == "transcription.text.delta" {
			text += event.Text
		}
		if event.Type == "transcription.done" {
			if event.Text != "hello" {
				t.Fatalf("expected final text hello, got %q", event.Text)
			}
			break
		}
	}
	if text != "hello" {
		t.Fatalf("expected delta text hello, got %q", text)
	}
}
