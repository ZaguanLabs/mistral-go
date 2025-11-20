package sdk

import (
	"testing"
)

func TestEmbeddings(t *testing.T) {
	client := NewMistralClientDefault("")
	res, err := client.Embeddings("mistral-embed", []string{"Embed this sentence.", "As well as this one."})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res == nil {
		t.Fatal("res should not be nil")
	}

	if len(res.Data) != 2 {
		t.Errorf("expected %v data items, got %v", 2, len(res.Data))
	}
	if len(res.Data[0].Embedding) != 1024 {
		t.Errorf("expected embedding length %v, got %v", 1024, len(res.Data[0].Embedding))
	}
}
