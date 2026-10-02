package sdk

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"nhooyr.io/websocket"
)

func TestServiceAccountAuthPrecedenceAndRotation(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	t.Setenv(ServiceAccountTokenPathEnv, path)
	t.Setenv("MISTRAL_API_KEY", "env-key")
	write := func(value string) {
		t.Helper()
		if err := os.WriteFile(path, []byte(value), 0600); err != nil {
			t.Fatal(err)
		}
	}
	var calls atomic.Int32
	want := "Bearer sa-one"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		if r.Header.Get("Authorization") != want {
			t.Errorf("Authorization %q want %q", r.Header.Get("Authorization"), want)
		}
		w.Write([]byte(`{"data":[]}`))
	}))
	defer server.Close()
	write(" sa-one\n")
	client := NewMistralClient("", server.URL, 1, time.Second)
	if _, err := client.ListModels(); err != nil {
		t.Fatal(err)
	}
	write("Bearer sa-two")
	want = "Bearer sa-two"
	if _, err := client.ListModels(); err != nil {
		t.Fatal(err)
	}
	want = "Bearer explicit"
	if _, err := NewMistralClient("explicit", server.URL, 1, time.Second).ListModels(); err != nil {
		t.Fatal(err)
	}
	write("")
	before := calls.Load()
	_, err := client.ListModels()
	var tokenErr *ServiceAccountTokenError
	if !errors.As(err, &tokenErr) || calls.Load() != before {
		t.Fatalf("empty token did not fail before HTTP: %v", err)
	}
	want = "Bearer env-key"
	if _, err := client.WithHeaders(http.Header{"authorization": []string{"Bearer env-key"}}).ListModels(); err != nil {
		t.Fatal(err)
	}
	t.Setenv(ServiceAccountTokenPathEnv, "")
	if _, err := client.ListModels(); err != nil {
		t.Fatal(err)
	}
	t.Setenv(ServiceAccountTokenPathEnv, path+"-missing")
	_, err = client.ListModels()
	if !errors.As(err, &tokenErr) {
		t.Fatal(err)
	}
}

func TestServiceAccountRealtimeAuth(t *testing.T) {
	path := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(path, []byte("sa-token"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(ServiceAccountTokenPathEnv, path)
	t.Setenv("MISTRAL_API_KEY", "env-key")
	for _, override := range []string{"", "Bearer env-key"} {
		t.Run(override, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				want := "Bearer sa-token"
				if override != "" {
					want = override
				}
				if r.Header.Get("Authorization") != want {
					t.Errorf("got %q want %q", r.Header.Get("Authorization"), want)
				}
				conn, err := websocket.Accept(w, r, nil)
				if err != nil {
					return
				}
				defer conn.CloseNow()
				conn.Write(r.Context(), websocket.MessageText, []byte(`{"type":"session.created","session":{"request_id":"req","model":"model","audio_format":{"encoding":"pcm_s16le","sample_rate":16000}}}`))
				conn.Read(r.Context())
			}))
			defer server.Close()
			params := &RealtimeTranscriptionConnectParams{}
			if override != "" {
				params.Headers = http.Header{"Authorization": []string{override}}
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			conn, err := NewMistralClient("", server.URL, 1, time.Second).RealtimeTranscriptionConnect(ctx, "model", params)
			if err != nil {
				t.Fatal(err)
			}
			conn.Close(websocket.StatusNormalClosure, "")
		})
	}
}
