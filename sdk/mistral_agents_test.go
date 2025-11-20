package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMistralAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := MistralAgent{
			ID:     "agent-123",
			Object: "agent",
			Model:  "mistral-large-latest",
			Name:   StringPtr("Test Agent"),
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &CreateMistralAgentRequest{
		Model:        "mistral-large-latest",
		Name:         StringPtr("Test Agent"),
		Instructions: StringPtr("Be helpful"),
	}

	resp, err := client.CreateMistralAgent(req)
	if err != nil {
		t.Fatalf("CreateMistralAgent failed: %v", err)
	}

	if resp.ID != "agent-123" {
		t.Errorf("Expected agent-123, got %s", resp.ID)
	}
}

func TestCreateMistralAgentNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.CreateMistralAgent(nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}

func TestListMistralAgents(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		response := MistralAgentListResponse{
			Object: "list",
			Data: []MistralAgent{
				{ID: "agent-1", Model: "mistral-large-latest"},
				{ID: "agent-2", Model: "mistral-small-latest"},
			},
			Total: 2,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ListMistralAgents(0)
	if err != nil {
		t.Fatalf("ListMistralAgents failed: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 agents, got %d", len(resp.Data))
	}

	if resp.Total != 2 {
		t.Errorf("Expected total 2, got %d", resp.Total)
	}
}

func TestGetMistralAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := MistralAgent{
			ID:     "agent-123",
			Object: "agent",
			Model:  "mistral-large-latest",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.GetMistralAgent("agent-123")
	if err != nil {
		t.Fatalf("GetMistralAgent failed: %v", err)
	}

	if resp.ID != "agent-123" {
		t.Errorf("Expected agent-123, got %s", resp.ID)
	}
}

func TestUpdateMistralAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH, got %s", r.Method)
		}

		response := MistralAgent{
			ID:           "agent-123",
			Model:        "mistral-large-latest",
			Name:         StringPtr("Updated Agent"),
			Instructions: StringPtr("New instructions"),
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &UpdateMistralAgentRequest{
		Name:         StringPtr("Updated Agent"),
		Instructions: StringPtr("New instructions"),
	}

	resp, err := client.UpdateMistralAgent("agent-123", req)
	if err != nil {
		t.Fatalf("UpdateMistralAgent failed: %v", err)
	}

	if resp.Name == nil || *resp.Name != "Updated Agent" {
		t.Error("Expected name to be Updated Agent")
	}
}

func TestUpdateMistralAgentNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.UpdateMistralAgent("agent-123", nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}
