package sdk

import (
	"os"
	"testing"
)

func TestAgentCompletionWithMetadata(t *testing.T) {
	apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
		t.Skip("MISTRAL_API_KEY not set")
	}
	client := NewMistralClient(apiKey, Endpoint, DefaultMaxRetries, DefaultTimeout)

	messages := []ChatMessage{
		UserMessage("Hello, what can you help me with?"),
	}

	params := &AgentCompletionRequest{
		Metadata: map[string]any{
			"user_id":    "test-user-123",
			"session_id": "session-456",
			"custom_tag": "v1.10.0-test",
		},
	}

	// Note: This test requires a valid agent ID
	// In a real test, you would use a test agent or mock the API
	t.Skip("Skipping test that requires valid agent ID")

	response, err := client.AgentComplete("test-agent-id", messages, params)
	if err != nil {
		t.Fatalf("AgentComplete with metadata failed: %v", err)
	}

	if response == nil {
		t.Fatal("Expected non-nil response")
	}
}

func TestDeleteMistralAgent(t *testing.T) {
	apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
		t.Skip("MISTRAL_API_KEY not set")
	}
	client := NewMistralClient(apiKey, Endpoint, DefaultMaxRetries, DefaultTimeout)

	// Note: This test requires a valid agent ID
	// In a real test, you would create a test agent first
	t.Skip("Skipping test that requires valid agent ID")

	err := client.DeleteMistralAgent("test-agent-id")
	if err != nil {
		t.Fatalf("DeleteMistralAgent failed: %v", err)
	}
}

func TestDeleteConversation(t *testing.T) {
	apiKey := os.Getenv("MISTRAL_API_KEY")
	if apiKey == "" {
		t.Skip("MISTRAL_API_KEY not set")
	}
	client := NewMistralClient(apiKey, Endpoint, DefaultMaxRetries, DefaultTimeout)

	// Note: This test requires a valid conversation ID
	// In a real test, you would create a test conversation first
	t.Skip("Skipping test that requires valid conversation ID")

	err := client.DeleteConversation("test-conversation-id")
	if err != nil {
		t.Fatalf("DeleteConversation failed: %v", err)
	}
}

func TestRequestSourceConstants(t *testing.T) {
	// Test that RequestSource constants are defined correctly
	sources := []RequestSource{
		RequestSourceAPI,
		RequestSourcePlayground,
		RequestSourceAgentBuilderV1,
	}

	expectedValues := []string{"api", "playground", "agent_builder_v1"}

	for i, source := range sources {
		if string(source) != expectedValues[i] {
			t.Errorf("Expected RequestSource %s, got %s", expectedValues[i], source)
		}
	}
}

func TestOCRTableFormatConstants(t *testing.T) {
	// Test that OCRTableFormat constants are defined correctly
	formats := []OCRTableFormat{
		OCRTableFormatMarkdown,
		OCRTableFormatHTML,
	}

	expectedValues := []string{"markdown", "html"}

	for i, format := range formats {
		if string(format) != expectedValues[i] {
			t.Errorf("Expected OCRTableFormat %s, got %s", expectedValues[i], format)
		}
	}
}

func TestOCRTableObject(t *testing.T) {
	// Test OCRTableObject structure
	table := OCRTableObject{
		ID:      "table-1",
		Content: "| Header 1 | Header 2 |\n|----------|----------|\n| Cell 1   | Cell 2   |",
		Format:  OCRTableFormatMarkdown,
	}

	if table.ID != "table-1" {
		t.Errorf("Expected table ID 'table-1', got %s", table.ID)
	}

	if table.Format != OCRTableFormatMarkdown {
		t.Errorf("Expected format markdown, got %s", table.Format)
	}
}

func TestOCRPageObjectWithTables(t *testing.T) {
	// Test OCRPageObject with enhanced fields
	header := "Document Header"
	footer := "Page 1 of 10"

	page := OCRPageObject{
		PageNumber: 1,
		Text:       "Sample text content",
		Tables: []OCRTableObject{
			{
				ID:      "table-1",
				Content: "| A | B |\n|---|---|\n| 1 | 2 |",
				Format:  OCRTableFormatMarkdown,
			},
		},
		Hyperlinks: []string{
			"https://example.com",
			"https://mistral.ai",
		},
		Header: &header,
		Footer: &footer,
	}

	if len(page.Tables) != 1 {
		t.Errorf("Expected 1 table, got %d", len(page.Tables))
	}

	if len(page.Hyperlinks) != 2 {
		t.Errorf("Expected 2 hyperlinks, got %d", len(page.Hyperlinks))
	}

	if page.Header == nil || *page.Header != "Document Header" {
		t.Error("Expected header to be set")
	}

	if page.Footer == nil || *page.Footer != "Page 1 of 10" {
		t.Error("Expected footer to be set")
	}
}
