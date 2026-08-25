package sdk

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPython294SchemaParity(t *testing.T) {
	ttl := 600
	b, err := json.Marshal(CreateRealtimeSessionRequest{Model: "voxtral", TTLSeconds: &ttl})
	if err != nil || !json.Valid(b) || string(b) != `{"model":"voxtral","ttl_seconds":600}` {
		t.Fatalf("realtime request = %s, %v", b, err)
	}
	double := 1.5
	element, err := json.Marshal(TempoTraceAttributeArrayElement{DoubleValue: &double})
	if err != nil || string(element) != `{"doubleValue":1.5}` {
		t.Fatalf("tempo element = %s, %v", element, err)
	}
	now := time.Now().UTC().Truncate(time.Second)
	prompt, err := json.Marshal(Prompt{CreatedBy: stringPointer("user"), VersionCreatedAt: &now})
	if err != nil || len(prompt) == 0 {
		t.Fatalf("prompt = %s, %v", prompt, err)
	}
}

func stringPointer(value string) *string { return &value }
