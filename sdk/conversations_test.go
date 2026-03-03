package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStartConversation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/conversations" {
			t.Errorf("Expected /v1/conversations, got %s", r.URL.Path)
		}

		response := ConversationResponse{
			ConversationID: "conv-123",
			Object:         "conversation",
			Created:        1234567890,
			Status:         "active",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &ConversationStartRequest{
		Inputs: []ConversationInput{
			{Type: "text", Content: "Hello"},
		},
		Instructions: StringPtr("Be helpful"),
	}

	resp, err := client.StartConversation(req)
	if err != nil {
		t.Fatalf("StartConversation failed: %v", err)
	}

	if resp.ConversationID != "conv-123" {
		t.Errorf("Expected conv-123, got %s", resp.ConversationID)
	}
}

func TestStartConversationNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.StartConversation(nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}

func TestListConversations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		response := ConversationListResponse{
			Object: "list",
			Data: []ConversationResponse{
				{ConversationID: "conv-1", Status: "active"},
				{ConversationID: "conv-2", Status: "completed"},
			},
			Total: 2,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ListConversations(0)
	if err != nil {
		t.Fatalf("ListConversations failed: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 conversations, got %d", len(resp.Data))
	}
}

func TestListConversationsWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/conversations" {
			t.Errorf("Expected /v1/conversations, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("page") != "2" {
			t.Errorf("Expected page=2, got %s", r.URL.Query().Get("page"))
		}
		if r.URL.Query().Get("page_size") != "25" {
			t.Errorf("Expected page_size=25, got %s", r.URL.Query().Get("page_size"))
		}

		response := ConversationListResponse{Object: "list", Data: []ConversationResponse{}, Total: 0}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)
	page := 2
	pageSize := 25
	_, err := client.ListConversationsWithParams(&ListConversationsParams{Page: &page, PageSize: &pageSize})
	if err != nil {
		t.Fatalf("ListConversationsWithParams failed: %v", err)
	}
}

func TestGetConversation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ConversationResponse{
			ConversationID: "conv-123",
			Object:         "conversation",
			Status:         "active",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.GetConversation("conv-123")
	if err != nil {
		t.Fatalf("GetConversation failed: %v", err)
	}

	if resp.ConversationID != "conv-123" {
		t.Errorf("Expected conv-123, got %s", resp.ConversationID)
	}
}

func TestAppendToConversation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ConversationResponse{
			ConversationID: "conv-123",
			Status:         "active",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	inputs := []ConversationInput{
		{Type: "text", Content: "Follow-up message"},
	}

	resp, err := client.AppendToConversation("conv-123", inputs)
	if err != nil {
		t.Fatalf("AppendToConversation failed: %v", err)
	}

	if resp.ConversationID != "conv-123" {
		t.Errorf("Expected conv-123, got %s", resp.ConversationID)
	}
}

func TestGetConversationHistory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ConversationHistoryResponse{
			ConversationID: "conv-123",
			Entries: []ConversationEntry{
				{Type: "user", Content: "Hello", Timestamp: 1234567890},
				{Type: "assistant", Content: "Hi there", Timestamp: 1234567891},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.GetConversationHistory("conv-123")
	if err != nil {
		t.Fatalf("GetConversationHistory failed: %v", err)
	}

	if len(resp.Entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(resp.Entries))
	}
}

func TestGetConversationMessages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/messages") {
			t.Errorf("Expected path ending with /messages, got %s", r.URL.Path)
		}

		response := ConversationMessagesResponse{
			ConversationID: "conv-123",
			Messages: []ChatMessage{
				{Role: RoleUser, Content: "Hello"},
				{Role: RoleAssistant, Content: "Hi"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)
	resp, err := client.GetConversationMessages("conv-123")
	if err != nil {
		t.Fatalf("GetConversationMessages failed: %v", err)
	}

	if len(resp.Messages) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(resp.Messages))
	}
}

func TestRestartConversation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := ConversationResponse{
			ConversationID: "conv-123",
			Status:         "restarted",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	inputs := []ConversationInput{
		{Type: "text", Content: "Restart message"},
	}

	resp, err := client.RestartConversation("conv-123", inputs)
	if err != nil {
		t.Fatalf("RestartConversation failed: %v", err)
	}

	if resp.Status != "restarted" {
		t.Errorf("Expected restarted, got %s", resp.Status)
	}
}
