package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestModerate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := ModerationResponse{
			ID:    "mod-123",
			Model: "mistral-moderation-latest",
			Results: []ModerationResult{
				{
					Categories: []ModerationCategory{
						{CategoryName: "sexual", Score: 0.01},
						{CategoryName: "hate", Score: 0.02},
						{CategoryName: "violence", Score: 0.01},
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	inputs := []ClassificationInput{"Test input"}
	resp, err := client.Moderate("mistral-moderation-latest", inputs)
	if err != nil {
		t.Fatalf("Moderate failed: %v", err)
	}

	if resp.ID != "mod-123" {
		t.Errorf("Expected mod-123, got %s", resp.ID)
	}

	if len(resp.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(resp.Results))
	}
}

func TestModerateText(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ModerationResponse{
			ID:    "mod-456",
			Model: "mistral-moderation-latest",
			Results: []ModerationResult{
				{
					Categories: []ModerationCategory{
						{CategoryName: "sexual", Score: 0.95},
						{CategoryName: "violence", Score: 0.05},
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ModerateText("mistral-moderation-latest", []string{"Test text"})
	if err != nil {
		t.Fatalf("ModerateText failed: %v", err)
	}

	if len(resp.Results) == 0 {
		t.Fatal("Expected at least one result")
	}

	if len(resp.Results[0].Categories) != 2 {
		t.Errorf("Expected 2 categories, got %d", len(resp.Results[0].Categories))
	}
}

func TestModerateWithRawScores(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ModerationResponse{
			ID:    "mod-123",
			Model: "mistral-moderation-latest",
			Results: []ModerationResult{
				{
					Categories: []ModerationCategory{
						{CategoryName: "sexual", Score: 0.95},
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	inputs := []ClassificationInput{"Inappropriate content"}
	resp, err := client.Moderate("mistral-moderation-latest", inputs)
	if err != nil {
		t.Fatalf("Moderate failed: %v", err)
	}

	if len(resp.Results) == 0 {
		t.Fatal("Expected at least one result")
	}
	if len(resp.Results[0].Categories) == 0 {
		t.Fatal("Expected at least one category")
	}

	if resp.Results[0].Categories[0].CategoryName != "sexual" {
		t.Errorf("Expected sexual category, got %s", resp.Results[0].Categories[0].CategoryName)
	}

	if resp.Results[0].Categories[0].Score != 0.95 {
		t.Errorf("Expected score 0.95, got %f", resp.Results[0].Categories[0].Score)
	}
}

func TestModerateChat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/chat/moderations" {
			t.Errorf("Expected /v1/chat/moderations, got %s", r.URL.Path)
		}

		response := ModerationResponse{
			ID:    "mod-chat-123",
			Model: "mistral-moderation-latest",
			Results: []ModerationResult{{
				Categories: []ModerationCategory{{CategoryName: "violence", Score: 0.02}},
			}},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)
	inputs := []ChatMessage{{Role: RoleUser, Content: "hello"}}

	resp, err := client.ModerateChat("mistral-moderation-latest", inputs)
	if err != nil {
		t.Fatalf("ModerateChat failed: %v", err)
	}
	if resp.ID != "mod-chat-123" {
		t.Errorf("Expected mod-chat-123, got %s", resp.ID)
	}
}

func TestClassify(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/classifications" {
			t.Errorf("Expected /v1/classifications, got %s", r.URL.Path)
		}

		response := ClassificationResponse{
			ID:    "cls-123",
			Model: "mistral-moderation-latest",
			Results: []ModerationResult{{
				Categories: []ModerationCategory{{CategoryName: "safety", Score: 0.9}},
			}},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)
	resp, err := client.Classify("mistral-moderation-latest", []ClassificationInput{"sample"})
	if err != nil {
		t.Fatalf("Classify failed: %v", err)
	}
	if resp.ID != "cls-123" {
		t.Errorf("Expected cls-123, got %s", resp.ID)
	}
}

func TestClassifyChat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/chat/classifications" {
			t.Errorf("Expected /v1/chat/classifications, got %s", r.URL.Path)
		}

		response := ClassificationResponse{
			ID:    "cls-chat-123",
			Model: "mistral-moderation-latest",
			Results: []ModerationResult{{
				Categories: []ModerationCategory{{CategoryName: "spam", Score: 0.11}},
			}},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)
	inputs := []ChatClassificationInput{{
		Messages: []ChatMessage{{Role: RoleUser, Content: "hello"}},
	}}
	resp, err := client.ClassifyChat("mistral-moderation-latest", inputs)
	if err != nil {
		t.Fatalf("ClassifyChat failed: %v", err)
	}
	if resp.ID != "cls-chat-123" {
		t.Errorf("Expected cls-chat-123, got %s", resp.ID)
	}
}
