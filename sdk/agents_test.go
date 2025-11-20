package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAgentCompletion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := ChatCompletionResponse{
			ID:      "agent-completion-123",
			Object:  "chat.completion",
			Model:   "mistral-large-latest",
			Created: 1234567890,
			Choices: []ChatCompletionResponseChoice{
				{
					Index:        0,
					Message:      ChatMessage{Role: "assistant", Content: "Agent response"},
					FinishReason: FinishReasonStop,
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	messages := []ChatMessage{
		{Role: "user", Content: "Hello agent"},
	}

	resp, err := client.AgentComplete("agent-123", messages, nil)
	if err != nil {
		t.Fatalf("AgentComplete failed: %v", err)
	}

	if resp.ID != "agent-completion-123" {
		t.Errorf("Expected agent-completion-123, got %s", resp.ID)
	}

	if len(resp.Choices) != 1 {
		t.Errorf("Expected 1 choice, got %d", len(resp.Choices))
	}
}

func TestAgentCompleteWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ChatCompletionResponse{
			ID: "agent-completion-456",
			Choices: []ChatCompletionResponseChoice{
				{Message: ChatMessage{Content: "Response with params"}},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	messages := []ChatMessage{{Role: "user", Content: "Test"}}
	params := &AgentCompletionRequest{
		MaxTokens: IntPtr(100),
	}

	resp, err := client.AgentComplete("agent-123", messages, params)
	if err != nil {
		t.Fatalf("AgentComplete with params failed: %v", err)
	}

	if resp.ID != "agent-completion-456" {
		t.Errorf("Expected agent-completion-456, got %s", resp.ID)
	}
}

func TestAgentCompleteNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.AgentComplete("agent-123", nil, nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}
