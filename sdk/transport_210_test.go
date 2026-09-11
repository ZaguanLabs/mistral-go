package sdk

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestPython210RetryControls(t *testing.T) {
	var attempts atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		if string(raw) != `{"value":"replayed"}` {
			t.Errorf("retry body %s", raw)
		}
		if attempts.Add(1) == 1 {
			w.Header().Set("retry-after-ms", "0")
			w.WriteHeader(409)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{}`)
	}))
	defer server.Close()
	zero := time.Duration(0)
	base := NewMistralClient("test", server.URL, 3, time.Second)
	client, err := base.WithRetryConfig(RetryConfig{Strategy: "backoff", Backoff: BackoffStrategy{Jitter: &zero}, StatusCodesOverride: []string{"409"}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.requestMap("POST", map[string]any{"value": "replayed"}, "v1/test"); err != nil {
		t.Fatal(err)
	}
	if attempts.Load() != 2 {
		t.Errorf("attempts=%d", attempts.Load())
	}
	if base.retryConfig != nil {
		t.Error("configuration mutated base client")
	}
	attempts.Store(0)
	disabled, _ := client.WithRetryConfig(RetryConfig{Strategy: "none"})
	if _, err := disabled.requestMap("POST", map[string]any{"value": "replayed"}, "v1/test"); err == nil {
		t.Fatal("expected HTTP error")
	}
	if attempts.Load() != 1 {
		t.Errorf("disabled retry attempts=%d", attempts.Load())
	}
	for _, tc := range []struct {
		headers http.Header
		want    time.Duration
	}{
		{http.Header{"Retry-After-Ms": {"1.5"}, "Retry-After": {"10"}}, 1500 * time.Microsecond},
		{http.Header{"Retry-After": {"2"}}, 2 * time.Second},
	} {
		got, ok := retryAfter(tc.headers, time.Now())
		if !ok || got != tc.want {
			t.Errorf("delay %v %v", got, ok)
		}
	}
	negative := -time.Millisecond
	if _, err := base.WithRetryConfig(RetryConfig{Strategy: "backoff", Backoff: BackoffStrategy{Jitter: &negative}}); err == nil {
		t.Error("negative jitter accepted")
	}
}

func TestPython210RetryCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("retry-after-ms", "60000")
		w.WriteHeader(503)
	}))
	defer server.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	c := NewMistralClient("test", server.URL, 3, time.Second).WithContext(ctx)
	_, err := c.requestMap("GET", nil, "v1/models")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("cancellation: %v", err)
	}
}

type trackingReadCloser struct {
	io.Reader
	closed atomic.Bool
}

func (r *trackingReadCloser) Close() error { r.closed.Store(true); return nil }

type oneByteReader struct{ r io.Reader }

func (r oneByteReader) Read(b []byte) (int, error) { return r.r.Read(b[:1]) }
func TestPython210EventStreamFraming(t *testing.T) {
	raw := "\ufeffid: first\rdata: {\rdata: \"value\":1}\r\rdata:\n\ndata: {\"value\":2}"
	body := &trackingReadCloser{Reader: oneByteReader{strings.NewReader(raw)}}
	stream := NewEventStream(body)
	first, err := stream.Next()
	if err != nil || first.ID != "first" || first.Data != "{\n\"value\":1}" {
		t.Fatalf("first %+v %v", first, err)
	}
	empty, err := stream.Next()
	if err != nil || empty.Data != "" {
		t.Fatalf("empty data %+v %v", empty, err)
	}
	last, err := stream.Next()
	if err != nil || last.Data != `{"value":2}` || !body.closed.Load() {
		t.Fatalf("EOF event %+v %v closed=%v", last, err, body.closed.Load())
	}
	if _, err := stream.Next(); err != io.EOF {
		t.Errorf("after EOF %v", err)
	}
}
func TestPython210StreamCancellationAndErrors(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	body := &trackingReadCloser{Reader: bytes.NewBufferString("data: {\"value\":1}\n\ndata: {\"value\":2}\n\n")}
	stream := streamJSON(ctx, body, func(err error) batchTestBody { return batchTestBody{Value: -1} })
	<-stream
	cancel()
	for range stream {
	}
	if !body.closed.Load() {
		t.Error("abandoned stream not closed")
	}
	body = &trackingReadCloser{Reader: strings.NewReader("event: error\ndata: {\"reason\":\"timeout\",\"error\":\"expired\"}\n\n")}
	events := parseGenericStream(body)
	event := <-events
	var disconnected *StreamDisconnectedError
	if !errors.As(event.Error, &disconnected) || disconnected.Reason != "timeout" {
		t.Fatalf("stream error %+v", event)
	}
	for range events {
	}
	if !body.closed.Load() {
		t.Error("error stream not closed")
	}
}

func TestPython210Base64InputRestoresPosition(t *testing.T) {
	reader := strings.NewReader("prefix payload")
	if _, err := reader.Seek(7, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	encoded, err := EncodeBase64FileInput(reader)
	if err != nil || encoded != "cGF5bG9hZA==" {
		t.Fatalf("base64 %q %v", encoded, err)
	}
	pos, _ := reader.Seek(0, io.SeekCurrent)
	if pos != 7 {
		t.Errorf("reader position %d", pos)
	}
	if encoded, err := EncodeBase64FileInput("already encoded"); err != nil || encoded != "already encoded" {
		t.Fatal("string input changed")
	}
}

func TestPython210EscapedPathAndMultipartRetry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.URL.EscapedPath() != "/v1/service-accounts/a%2Fb" {
				t.Errorf("escaped path %s", r.URL.EscapedPath())
			}
			io.WriteString(w, `{"id":"a/b"}`)
			return
		}
		attempts++
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			t.Error(err)
			return
		}
		f, _, err := r.FormFile("file")
		if err != nil {
			t.Error(err)
			return
		}
		defer f.Close()
		raw, _ := io.ReadAll(f)
		if string(raw) != "payload" {
			t.Errorf("upload retry body %s", raw)
		}
		if attempts == 1 {
			w.Header().Set("retry-after-ms", "0")
			w.WriteHeader(503)
			return
		}
		io.WriteString(w, `{"id":"file"}`)
	}))
	defer server.Close()
	c := NewMistralClient("test", server.URL, 2, time.Second)
	if _, err := c.GetServiceAccount("a/b"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.UploadFile(strings.NewReader("payload"), "test.jsonl", FilePurposeBatch); err != nil {
		t.Fatal(err)
	}
	if attempts != 2 {
		t.Errorf("attempts %d", attempts)
	}
}
