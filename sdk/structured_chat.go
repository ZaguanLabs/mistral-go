package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// UnmarshalJSON accepts plain text, null, or chunked assistant content. Text
// chunks are concatenated for Content; ContentChunks preserves the full wire data.
func (m *ChatMessage) UnmarshalJSON(data []byte) error {
	var raw struct {
		Role       string          `json:"role"`
		Content    json.RawMessage `json:"content"`
		ToolCalls  []ToolCall      `json:"tool_calls"`
		ToolCallID string          `json:"tool_call_id"`
		Name       string          `json:"name"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*m = ChatMessage{Role: raw.Role, ToolCalls: raw.ToolCalls, ToolCallID: raw.ToolCallID, Name: raw.Name}
	if len(raw.Content) == 0 || bytes.Equal(bytes.TrimSpace(raw.Content), []byte("null")) {
		return nil
	}
	if err := json.Unmarshal(raw.Content, &m.Content); err == nil {
		return nil
	}
	if err := json.Unmarshal(raw.Content, &m.ContentChunks); err != nil {
		return fmt.Errorf("invalid message content: %w", err)
	}
	var text strings.Builder
	for _, chunk := range m.ContentChunks {
		if chunk["type"] == "text" {
			if value, ok := chunk["text"].(string); ok {
				text.WriteString(value)
			}
		}
	}
	m.Content = text.String()
	return nil
}
func (m ChatMessage) MarshalJSON() ([]byte, error) {
	type wire ChatMessage
	if m.ContentChunks == nil {
		return json.Marshal(wire(m))
	}
	raw, err := json.Marshal(wire(m))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err = json.Unmarshal(raw, &fields); err != nil {
		return nil, err
	}
	fields["content"], err = json.Marshal(m.ContentChunks)
	if err != nil {
		return nil, err
	}
	return json.Marshal(fields)
}

type ParsedChatCompletionChoice[T any] struct {
	ChatCompletionResponseChoice
	Parsed *T
}
type ParsedChatCompletionResponse[T any] struct {
	Response *ChatCompletionResponse
	Choices  []ParsedChatCompletionChoice[T]
}

// ParseChatCompletion decodes structured JSON from string or text-chunk content.
// Empty and non-text-only content produce a nil Parsed value.
func ParseChatCompletion[T any](response *ChatCompletionResponse) (*ParsedChatCompletionResponse[T], error) {
	if response == nil {
		return nil, fmt.Errorf("response cannot be nil")
	}
	parsed := &ParsedChatCompletionResponse[T]{Response: response, Choices: make([]ParsedChatCompletionChoice[T], len(response.Choices))}
	for i, choice := range response.Choices {
		parsed.Choices[i].ChatCompletionResponseChoice = choice
		if choice.Message.Content != "" {
			var value T
			if err := json.Unmarshal([]byte(choice.Message.Content), &value); err != nil {
				return nil, err
			}
			parsed.Choices[i].Parsed = &value
		}
	}
	return parsed, nil
}

func validateCompletionTools(tools any) error {
	if tools == nil {
		return nil
	}
	data, err := json.Marshal(tools)
	if err != nil {
		return err
	}
	if string(data) == "null" {
		return nil
	}
	var list []struct {
		Type string `json:"type"`
	}
	if err = json.Unmarshal(data, &list); err != nil {
		return err
	}
	for _, tool := range list {
		switch tool.Type {
		case "function", "image_generation", "document_library", "connector":
		default:
			return fmt.Errorf("unsupported completion tool %q; use a connector for built-in search or code tools", tool.Type)
		}
	}
	return nil
}
