package sdk

import (
	"testing"
)

// Test helper message functions
func TestSystemMessage(t *testing.T) {
	msg := SystemMessage("system prompt")
	if msg.Role != RoleSystem {
		t.Errorf("Expected role %s, got %s", RoleSystem, msg.Role)
	}
	if msg.Content != "system prompt" {
		t.Error("Content not set correctly")
	}
}

func TestUserMessage(t *testing.T) {
	msg := UserMessage("user input")
	if msg.Role != RoleUser {
		t.Errorf("Expected role %s, got %s", RoleUser, msg.Role)
	}
	if msg.Content != "user input" {
		t.Error("Content not set correctly")
	}
}

func TestAssistantMessage(t *testing.T) {
	msg := AssistantMessage("assistant response")
	if msg.Role != RoleAssistant {
		t.Errorf("Expected role %s, got %s", RoleAssistant, msg.Role)
	}
	if msg.Content != "assistant response" {
		t.Error("Content not set correctly")
	}
}

func TestToolMessage(t *testing.T) {
	msg := ToolMessage("call-123", "tool result")
	if msg.Role != RoleTool {
		t.Errorf("Expected role %s, got %s", RoleTool, msg.Role)
	}
	if msg.Content != "tool result" {
		t.Error("Content not set correctly")
	}
	if msg.ToolCallID != "call-123" {
		t.Error("ToolCallID not set correctly")
	}
}

func TestMistralPromptModePtr(t *testing.T) {
	mode := MistralPromptModePtr(PromptModeReasoning)
	if mode == nil {
		t.Fatal("Should return non-nil pointer")
	}
	if *mode != PromptModeReasoning {
		t.Error("Value not set correctly")
	}
}

func TestMistralErrorError(t *testing.T) {
	err := &MistralError{Message: "test error"}
	if err.Error() != "test error" {
		t.Error("Error() should return message")
	}
}

// Test additional coverage for mapToStruct
func TestMapToStructEdgeCases(t *testing.T) {
	// Test with empty map
	var result struct {
		Field string `json:"field"`
	}
	err := mapToStruct(map[string]interface{}{}, &result)
	if err != nil {
		t.Errorf("Should handle empty map: %v", err)
	}

	// Test with nil values
	err = mapToStruct(map[string]interface{}{"field": nil}, &result)
	if err != nil {
		t.Errorf("Should handle nil values: %v", err)
	}
}

// Test NewChatRequestParams
func TestNewChatRequestParamsDefaults(t *testing.T) {
	params := NewChatRequestParams()
	if params == nil {
		t.Fatal("Should return non-nil params")
	}
	// All fields should be nil by default
	if params.Temperature != nil {
		t.Error("Temperature should be nil by default")
	}
	if params.MaxTokens != nil {
		t.Error("MaxTokens should be nil by default")
	}
}

// Test client constructors with various inputs
func TestNewCodestralClientDefaultEndpoint(t *testing.T) {
	client := NewCodestralClientDefault("test-key")
	if client.endpoint != CodestralEndpoint {
		t.Errorf("Expected Codestral endpoint, got %s", client.endpoint)
	}
}

// Test constants are defined
func TestRoleConstants(t *testing.T) {
	roles := []string{RoleSystem, RoleUser, RoleAssistant, RoleTool}
	for _, role := range roles {
		if role == "" {
			t.Error("Role constant is empty")
		}
	}
}

func TestFinishReasonConstants(t *testing.T) {
	reasons := []FinishReason{
		FinishReasonStop,
		FinishReasonLength,
		FinishReasonModelLength,
		FinishReasonError,
		FinishReasonToolCalls,
	}
	for _, reason := range reasons {
		if reason == "" {
			t.Error("FinishReason constant is empty")
		}
	}
}

func TestResponseFormatConstants(t *testing.T) {
	if ResponseFormatText == "" {
		t.Error("ResponseFormatText constant is empty")
	}
	if ResponseFormatJsonObject == "" {
		t.Error("ResponseFormatJsonObject constant is empty")
	}
}

func TestToolTypeConstant(t *testing.T) {
	if ToolTypeFunction == "" {
		t.Error("ToolTypeFunction constant is empty")
	}
}

func TestPromptModeConstants(t *testing.T) {
	modes := []MistralPromptMode{PromptModeReasoning}
	for _, mode := range modes {
		if mode == "" {
			t.Error("PromptMode constant is empty")
		}
	}
}
