package sdk

import (
	"fmt"
	"io"
	"net/http"
)

// AgentCompletionRequest represents a request for agent completion
type AgentCompletionRequest struct {
	AgentID           string                `json:"agent_id"`
	Messages          []ChatMessage         `json:"messages"`
	MaxTokens         *int                  `json:"max_tokens,omitempty"`
	Stream            *bool                 `json:"stream,omitempty"`
	Stop              interface{}           `json:"stop,omitempty"` // string or []string
	RandomSeed        *int                  `json:"random_seed,omitempty"`
	Metadata          map[string]any        `json:"metadata,omitempty"`
	ResponseFormat    *ResponseFormatSpec   `json:"response_format,omitempty"`
	Tools             []Tool                `json:"tools,omitempty"`
	ToolChoice        interface{}           `json:"tool_choice,omitempty"` // string or ToolChoice object
	PresencePenalty   *float64              `json:"presence_penalty,omitempty"`
	FrequencyPenalty  *float64              `json:"frequency_penalty,omitempty"`
	N                 *int                  `json:"n,omitempty"`
	Prediction        *Prediction           `json:"prediction,omitempty"`
	ParallelToolCalls *bool                 `json:"parallel_tool_calls,omitempty"`
	PromptMode        *MistralPromptMode    `json:"prompt_mode,omitempty"`
	ReasoningEffort   *ReasoningEffort      `json:"reasoning_effort,omitempty"`
	Guardrails        []GuardrailConfig     `json:"guardrails,omitempty"`
	PromptCacheKey    *string               `json:"prompt_cache_key,omitempty"`
	ServiceTier       *RequestedServiceTier `json:"service_tier,omitempty"`
}

// ResponseFormatSpec specifies the response format
type ResponseFormatSpec struct {
	Type ResponseFormat `json:"type"`
}

// AgentComplete performs an agent completion request
//
// Parameters:
//   - agentID: The ID of the agent to use
//   - messages: The conversation messages
//   - params: Optional parameters for the completion
//
// Returns a ChatCompletionResponse (same structure as regular chat)
func (c *MistralClient) AgentComplete(agentID string, messages []ChatMessage, params *AgentCompletionRequest) (*ChatCompletionResponse, error) {
	if params == nil {
		params = &AgentCompletionRequest{}
	}
	if err := validateCompletionTools(params.Tools); err != nil {
		return nil, err
	}

	// Set required fields
	params.AgentID = agentID
	params.Messages = messages

	// Ensure stream is false for non-streaming
	streamFalse := false
	params.Stream = &streamFalse

	// Convert params to map for request
	reqMap := map[string]interface{}{
		"agent_id": params.AgentID,
		"messages": params.Messages,
		"stream":   params.Stream,
	}

	// Add optional parameters
	if params.MaxTokens != nil {
		reqMap["max_tokens"] = params.MaxTokens
	}
	if params.Stop != nil {
		reqMap["stop"] = params.Stop
	}
	if params.RandomSeed != nil {
		reqMap["random_seed"] = params.RandomSeed
	}
	if params.Metadata != nil {
		reqMap["metadata"] = params.Metadata
	}
	if params.ResponseFormat != nil {
		reqMap["response_format"] = params.ResponseFormat
	}
	if len(params.Tools) > 0 {
		reqMap["tools"] = params.Tools
	}
	if params.ToolChoice != nil {
		reqMap["tool_choice"] = params.ToolChoice
	}
	if params.PresencePenalty != nil {
		reqMap["presence_penalty"] = params.PresencePenalty
	}
	if params.FrequencyPenalty != nil {
		reqMap["frequency_penalty"] = params.FrequencyPenalty
	}
	if params.N != nil {
		reqMap["n"] = params.N
	}
	if params.Prediction != nil {
		reqMap["prediction"] = params.Prediction
	}
	if params.ParallelToolCalls != nil {
		reqMap["parallel_tool_calls"] = params.ParallelToolCalls
	}
	if params.PromptMode != nil {
		reqMap["prompt_mode"] = params.PromptMode
	}
	if params.ReasoningEffort != nil {
		reqMap["reasoning_effort"] = params.ReasoningEffort
	}
	if params.Guardrails != nil {
		reqMap["guardrails"] = params.Guardrails
	}
	if params.PromptCacheKey != nil {
		reqMap["prompt_cache_key"] = params.PromptCacheKey
	}
	if params.ServiceTier != nil {
		reqMap["service_tier"] = params.ServiceTier
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/agents/completions", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var chatResponse ChatCompletionResponse
	err = mapToStruct(respData, &chatResponse)
	if err != nil {
		return nil, err
	}

	return &chatResponse, nil
}

// AgentCompleteStream performs a streaming agent completion request
//
// Parameters:
//   - agentID: The ID of the agent to use
//   - messages: The conversation messages
//   - params: Optional parameters for the completion
//
// Returns a channel of ChatCompletionStreamResponse
func (c *MistralClient) AgentCompleteStream(agentID string, messages []ChatMessage, params *AgentCompletionRequest) (<-chan ChatCompletionStreamResponse, error) {
	if params == nil {
		params = &AgentCompletionRequest{}
	}
	if err := validateCompletionTools(params.Tools); err != nil {
		return nil, err
	}

	// Set required fields
	params.AgentID = agentID
	params.Messages = messages

	// Ensure stream is true for streaming
	streamTrue := true
	params.Stream = &streamTrue

	// Convert params to map for request
	reqMap := map[string]interface{}{
		"agent_id": params.AgentID,
		"messages": params.Messages,
		"stream":   params.Stream,
	}

	// Add optional parameters
	if params.MaxTokens != nil {
		reqMap["max_tokens"] = params.MaxTokens
	}
	if params.Stop != nil {
		reqMap["stop"] = params.Stop
	}
	if params.RandomSeed != nil {
		reqMap["random_seed"] = params.RandomSeed
	}
	if params.Metadata != nil {
		reqMap["metadata"] = params.Metadata
	}
	if params.ResponseFormat != nil {
		reqMap["response_format"] = params.ResponseFormat
	}
	if len(params.Tools) > 0 {
		reqMap["tools"] = params.Tools
	}
	if params.ToolChoice != nil {
		reqMap["tool_choice"] = params.ToolChoice
	}
	if params.PresencePenalty != nil {
		reqMap["presence_penalty"] = params.PresencePenalty
	}
	if params.FrequencyPenalty != nil {
		reqMap["frequency_penalty"] = params.FrequencyPenalty
	}
	if params.N != nil {
		reqMap["n"] = params.N
	}
	if params.Prediction != nil {
		reqMap["prediction"] = params.Prediction
	}
	if params.ParallelToolCalls != nil {
		reqMap["parallel_tool_calls"] = params.ParallelToolCalls
	}
	if params.PromptMode != nil {
		reqMap["prompt_mode"] = params.PromptMode
	}
	if params.ReasoningEffort != nil {
		reqMap["reasoning_effort"] = params.ReasoningEffort
	}
	if params.Guardrails != nil {
		reqMap["guardrails"] = params.Guardrails
	}
	if params.PromptCacheKey != nil {
		reqMap["prompt_cache_key"] = params.PromptCacheKey
	}
	if params.ServiceTier != nil {
		reqMap["service_tier"] = params.ServiceTier
	}

	// Create response channel
	var responseChan <-chan ChatCompletionStreamResponse

	response, err := c.request(http.MethodPost, reqMap, "v1/agents/completions", true, nil)
	if err != nil {
		return nil, err
	}

	respBody, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	// Start streaming in a goroutine (same pattern as ChatStream)
	responseChan = streamJSON(c.requestContext(), respBody, func(err error) ChatCompletionStreamResponse { return ChatCompletionStreamResponse{Error: err} })

	return responseChan, nil
}

// NewAgentCompletionRequest creates a new AgentCompletionRequest with default values
func NewAgentCompletionRequest() *AgentCompletionRequest {
	return &AgentCompletionRequest{}
}
