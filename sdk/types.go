package sdk

// Model IDs
// Note: Model IDs change frequently as Mistral releases new versions.
// Instead of using hardcoded constants, use client.ListModels() to get
// the current list of available models dynamically.
//
// Example:
//   models, err := client.ListModels()
//   if err != nil {
//       log.Fatal(err)
//   }
//   for _, model := range models.Data {
//       fmt.Printf("Model: %s\n", model.ID)
//   }
//
// Common model ID patterns (subject to change):
//   - "mistral-large-latest" - Latest large model
//   - "mistral-small-latest" - Latest small model
//   - "codestral-latest" - Latest code model
//   - "mistral-embed" - Embedding model
//
// For production use, either:
//   1. Fetch models dynamically with ListModels()
//   2. Use specific versioned model IDs (e.g., "mistral-large-2411")
//   3. Define your own constants based on your requirements

const (
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleSystem    = "system"
	RoleTool      = "tool"
)

// FinishReason the reason that a chat message was finished
type FinishReason string

const (
	FinishReasonStop        FinishReason = "stop"
	FinishReasonLength      FinishReason = "length"
	FinishReasonError       FinishReason = "error"
	FinishReasonToolCalls   FinishReason = "tool_calls"
	FinishReasonModelLength FinishReason = "model_length"
)

// ResponseFormat the format that the response must adhere to
type ResponseFormat string

const (
	ResponseFormatText       ResponseFormat = "text"
	ResponseFormatJsonObject ResponseFormat = "json_object"
)

// ToolType type of tool defined for the llm
type ToolType string

const (
	ToolTypeFunction ToolType = "function"
)

const (
	ToolChoiceAny  = "any"
	ToolChoiceAuto = "auto"
	ToolChoiceNone = "none"
)

// Tool definition of a tool that the llm can call
type Tool struct {
	Type     ToolType `json:"type"`
	Function Function `json:"function"`
}

// Function definition of a function that the llm can call including its parameters
type Function struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}

// FunctionCall represents a request to call an external tool by the llm
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ToolCall represents the call to a tool by the llm
type ToolCall struct {
	Id       string       `json:"id"`
	Type     ToolType     `json:"type"`
	Function FunctionCall `json:"function"`
}

// DeltaMessage represents the delta between the prior state of the message and the new state of the message when streaming responses.
type DeltaMessage struct {
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls"`
}

// ChatMessage represents a single message in a chat.
type ChatMessage struct {
	Role       string     `json:"role"`
	Content    string     `json:"content,omitempty"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"` // For tool role messages
	Name       string     `json:"name,omitempty"`         // For function/tool messages
}

// SystemMessage creates a system message
func SystemMessage(content string) ChatMessage {
	return ChatMessage{Role: RoleSystem, Content: content}
}

// UserMessage creates a user message
func UserMessage(content string) ChatMessage {
	return ChatMessage{Role: RoleUser, Content: content}
}

// AssistantMessage creates an assistant message
func AssistantMessage(content string) ChatMessage {
	return ChatMessage{Role: RoleAssistant, Content: content}
}

// ToolMessage creates a tool message
func ToolMessage(toolCallID, content string) ChatMessage {
	return ChatMessage{Role: RoleTool, Content: content, ToolCallID: toolCallID}
}

// Prediction represents prediction parameters for speculative decoding
type Prediction struct {
	Type    string `json:"type"`    // "content" for content prediction
	Content string `json:"content"` // The predicted content
}

// MistralPromptMode represents the prompt mode
type MistralPromptMode string

const (
	PromptModeReasoning MistralPromptMode = "reasoning"
)

// EncodingFormat represents the encoding format for embeddings
type EncodingFormat string

const (
	EncodingFormatFloat EncodingFormat = "float"
)

// EmbeddingDtype represents the data type for embeddings
type EmbeddingDtype string

const (
	EmbeddingDtypeFloat32 EmbeddingDtype = "float32"
)

// Helper functions for creating pointers to common types

// Float64Ptr returns a pointer to a float64 value
func Float64Ptr(v float64) *float64 {
	return &v
}

// IntPtr returns a pointer to an int value
func IntPtr(v int) *int {
	return &v
}

// Int64Ptr returns a pointer to an int64 value
func Int64Ptr(v int64) *int64 {
	return &v
}

// BoolPtr returns a pointer to a bool value
func BoolPtr(v bool) *bool {
	return &v
}

// StringPtr returns a pointer to a string value
func StringPtr(v string) *string {
	return &v
}

// MistralPromptModePtr returns a pointer to a MistralPromptMode value
func MistralPromptModePtr(v MistralPromptMode) *MistralPromptMode {
	return &v
}
