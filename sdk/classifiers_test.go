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
