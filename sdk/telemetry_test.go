package sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	collect "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/protobuf/proto"
)

func Test390TelemetryDedicatedExport(t *testing.T) {
	ctx := context.Background()
	exported := make(chan []byte, 1)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/custom" || r.Header.Get("Authorization") != "Bearer rotated" {
			t.Errorf("incorrect export destination or authorization")
		}
		data, _ := io.ReadAll(r.Body)
		exported <- data
		w.Header().Set("Content-Type", "application/x-protobuf")
	}))
	defer server.Close()
	path := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(path, []byte("initial"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(ServiceAccountTokenPathEnv, path)
	t.Setenv("MISTRAL_API_KEY", "")
	t.Setenv("MISTRAL_OTLP_TRACES_ENDPOINT", server.URL+"/custom")
	client, err := NewMistralClientDefault("").WithTelemetry(ctx, TelemetryConfig{Mode: "dedicated", ExporterHTTPClient: server.Client()})
	if err != nil {
		t.Fatal(err)
	}
	tracer, err := client.TelemetryTracer("")
	if err != nil {
		t.Fatal(err)
	}
	_, span := tracer.Start(ctx, "alice@example.com")
	span.SetAttributes(attribute.String("description", "Bearer very-private-token"))
	span.End()
	if err = os.WriteFile(path, []byte("rotated"), 0600); err != nil {
		t.Fatal(err)
	}
	if err = client.ShutdownTelemetry(ctx); err != nil {
		t.Fatal(err)
	}
	var payload collect.ExportTraceServiceRequest
	data := <-exported
	if err = proto.Unmarshal(data, &payload); err != nil {
		t.Fatal(err)
	}
	spans := payload.ResourceSpans[0].ScopeSpans[0].Spans
	if len(spans) != 1 || spans[0].Name != "[REDACTED]" || strings.Contains(string(data), "very-private-token") {
		t.Fatal("sensitive trace content was exported")
	}
}
func Test390TelemetryCustomProvider(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	defer provider.Shutdown(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"id":"id","model":"embed","usage":{"prompt_tokens":4}}`)
	}))
	defer server.Close()
	client, err := NewMistralClient("key", server.URL, 1, 0).WithTelemetry(context.Background(), TelemetryConfig{Provider: provider})
	if err != nil {
		t.Fatal(err)
	}
	if _, err = client.CreateEmbeddings(EmbeddingRequest{Model: "embed", Input: "hi"}); err != nil {
		t.Fatal(err)
	}
	ended := recorder.Ended()
	if len(ended) != 1 || ended[0].Name() != "embeddings embed" {
		t.Fatalf("ended spans %v", ended)
	}
	attrs := attribute.NewSet(ended[0].Attributes()...)
	v, _ := attrs.Value("gen_ai.usage.input_tokens")
	if v.AsInt64() != 4 {
		t.Error("response usage missing")
	}
	t.Setenv("MISTRAL_SDK_TELEMETRY", "bad")
	if _, err := NewMistralClient("key", server.URL, 1, 0).GetBudget("dashboard"); err != nil {
		t.Fatal("optional telemetry broke request", err)
	}
}
