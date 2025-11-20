package sdk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ChatRequestParams represents the parameters for the Chat/ChatStream method of MistralClient.
type ChatRequestParams struct {
	// Sampling parameters
	Temperature *float64 `json:"temperature,omitempty"` // The temperature to use for sampling, between 0.0 and 0.7. Higher values make output more random.
	TopP        *float64 `json:"top_p,omitempty"`       // Nucleus sampling parameter. We recommend altering this or Temperature but not both.
	RandomSeed  *int     `json:"random_seed,omitempty"` // Seed for deterministic results

	// Token limits
	MaxTokens *int `json:"max_tokens,omitempty"` // Maximum tokens to generate
	MinTokens *int `json:"min_tokens,omitempty"` // Minimum tokens to generate (for FIM)

	// Stop sequences
	Stop []string `json:"stop,omitempty"` // Stop generation if these tokens are detected

	// Response format
	ResponseFormat ResponseFormat `json:"response_format,omitempty"` // Format for the response (text or json_object)

	// Tools and function calling
	Tools             []Tool `json:"tools,omitempty"`               // Available tools for the model
	ToolChoice        string `json:"tool_choice,omitempty"`         // How to select tools (auto, any, none, or specific tool)
	ParallelToolCalls *bool  `json:"parallel_tool_calls,omitempty"` // Whether to enable parallel tool calls

	// Penalties for repetition
	PresencePenalty  *float64 `json:"presence_penalty,omitempty"`  // Penalize repetition of words/phrases
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"` // Penalize based on frequency

	// Multiple completions
	N *int `json:"n,omitempty"` // Number of completions to return

	// Advanced features
	Prediction *Prediction        `json:"prediction,omitempty"`  // Speculative decoding prediction
	PromptMode *MistralPromptMode `json:"prompt_mode,omitempty"` // Prompt mode (e.g., "reasoning")
	SafePrompt *bool              `json:"safe_prompt,omitempty"` // Inject safety prompt
}

// NewChatRequestParams creates a new ChatRequestParams with sensible defaults
func NewChatRequestParams() *ChatRequestParams {
	return &ChatRequestParams{}
}

// ChatCompletionResponseChoice represents a choice in the chat completion response.
type ChatCompletionResponseChoice struct {
	Index        int          `json:"index"`
	Message      ChatMessage  `json:"message"`
	FinishReason FinishReason `json:"finish_reason,omitempty"`
}

// ChatCompletionResponseChoice represents a choice in the chat completion response.
type ChatCompletionResponseChoiceStream struct {
	Index        int          `json:"index"`
	Delta        DeltaMessage `json:"delta"`
	FinishReason FinishReason `json:"finish_reason,omitempty"`
}

// ChatCompletionResponse represents the response from the chat completion endpoint.
type ChatCompletionResponse struct {
	ID      string                         `json:"id"`
	Object  string                         `json:"object"`
	Created int                            `json:"created"`
	Model   string                         `json:"model"`
	Choices []ChatCompletionResponseChoice `json:"choices"`
	Usage   UsageInfo                      `json:"usage"`
}

// ChatCompletionStreamResponse represents the streamed response from the chat completion endpoint.
type ChatCompletionStreamResponse struct {
	ID      string                               `json:"id"`
	Model   string                               `json:"model"`
	Choices []ChatCompletionResponseChoiceStream `json:"choices"`
	Created int                                  `json:"created,omitempty"`
	Object  string                               `json:"object,omitempty"`
	Usage   UsageInfo                            `json:"usage,omitempty"`
	Error   error                                `json:"error,omitempty"`
}

// UsageInfo represents the usage information of a response.
type UsageInfo struct {
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
}

func (c *MistralClient) Chat(model string, messages []ChatMessage, params *ChatRequestParams) (*ChatCompletionResponse, error) {
	if params == nil {
		params = NewChatRequestParams()
	}

	requestData := map[string]interface{}{
		"model":    model,
		"messages": messages,
	}

	// Add optional parameters only if set
	if params.Temperature != nil {
		requestData["temperature"] = *params.Temperature
	}
	if params.TopP != nil {
		requestData["top_p"] = *params.TopP
	}
	if params.RandomSeed != nil {
		requestData["random_seed"] = *params.RandomSeed
	}
	if params.MaxTokens != nil {
		requestData["max_tokens"] = *params.MaxTokens
	}
	if params.MinTokens != nil {
		requestData["min_tokens"] = *params.MinTokens
	}
	if params.Stop != nil {
		requestData["stop"] = params.Stop
	}
	if params.SafePrompt != nil {
		requestData["safe_prompt"] = *params.SafePrompt
	}
	if params.PresencePenalty != nil {
		requestData["presence_penalty"] = *params.PresencePenalty
	}
	if params.FrequencyPenalty != nil {
		requestData["frequency_penalty"] = *params.FrequencyPenalty
	}
	if params.N != nil {
		requestData["n"] = *params.N
	}
	if params.Prediction != nil {
		requestData["prediction"] = params.Prediction
	}
	if params.PromptMode != nil {
		requestData["prompt_mode"] = *params.PromptMode
	}
	if params.ParallelToolCalls != nil {
		requestData["parallel_tool_calls"] = *params.ParallelToolCalls
	}
	if params.Tools != nil {
		requestData["tools"] = params.Tools
	}
	if params.ToolChoice != "" {
		requestData["tool_choice"] = params.ToolChoice
	}
	if params.ResponseFormat != "" {
		requestData["response_format"] = map[string]any{"type": params.ResponseFormat}
	}

	response, err := c.request(http.MethodPost, requestData, "v1/chat/completions", false, nil)
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

// ChatStream sends a chat message and returns a channel to receive streaming responses.
func (c *MistralClient) ChatStream(model string, messages []ChatMessage, params *ChatRequestParams) (<-chan ChatCompletionStreamResponse, error) {
	if params == nil {
		params = NewChatRequestParams()
	}

	responseChannel := make(chan ChatCompletionStreamResponse)

	requestData := map[string]interface{}{
		"model":    model,
		"messages": messages,
		"stream":   true,
	}

	// Add optional parameters only if set
	if params.Temperature != nil {
		requestData["temperature"] = *params.Temperature
	}
	if params.TopP != nil {
		requestData["top_p"] = *params.TopP
	}
	if params.RandomSeed != nil {
		requestData["random_seed"] = *params.RandomSeed
	}
	if params.MaxTokens != nil {
		requestData["max_tokens"] = *params.MaxTokens
	}
	if params.MinTokens != nil {
		requestData["min_tokens"] = *params.MinTokens
	}
	if params.Stop != nil {
		requestData["stop"] = params.Stop
	}
	if params.SafePrompt != nil {
		requestData["safe_prompt"] = *params.SafePrompt
	}
	if params.PresencePenalty != nil {
		requestData["presence_penalty"] = *params.PresencePenalty
	}
	if params.FrequencyPenalty != nil {
		requestData["frequency_penalty"] = *params.FrequencyPenalty
	}
	if params.N != nil {
		requestData["n"] = *params.N
	}
	if params.Prediction != nil {
		requestData["prediction"] = params.Prediction
	}
	if params.PromptMode != nil {
		requestData["prompt_mode"] = *params.PromptMode
	}
	if params.ParallelToolCalls != nil {
		requestData["parallel_tool_calls"] = *params.ParallelToolCalls
	}
	if params.Tools != nil {
		requestData["tools"] = params.Tools
	}
	if params.ToolChoice != "" {
		requestData["tool_choice"] = params.ToolChoice
	}
	if params.ResponseFormat != "" {
		requestData["response_format"] = map[string]any{"type": params.ResponseFormat}
	}

	response, err := c.request(http.MethodPost, requestData, "v1/chat/completions", true, nil)
	if err != nil {
		return nil, err
	}

	respBody, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	// Execute the HTTP request in a separate goroutine.
	go func() {
		defer close(responseChannel)
		defer respBody.Close()

		// Assuming ChatCompletionStreamResponse is already defined in your Go code.
		// Assuming responseChannel is a channel of ChatCompletionStreamResponse.

		// Create a buffered reader to read the stream line by line.
		reader := bufio.NewReader(respBody)

		for {
			// Read a line from the buffered reader.
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break // End of stream.
			} else if err != nil {
				responseChannel <- ChatCompletionStreamResponse{Error: fmt.Errorf("error reading stream response: %w", err)}
				return
			}

			// Skip empty lines.
			if bytes.Equal(line, []byte("\n")) {
				continue
			}

			// Check if the line starts with "data: ".
			if bytes.HasPrefix(line, []byte("data: ")) {
				// Trim the prefix and any leading or trailing whitespace.
				jsonLine := bytes.TrimSpace(bytes.TrimPrefix(line, []byte("data: ")))

				// Check for the special "[DONE]" message.
				if bytes.Equal(jsonLine, []byte("[DONE]")) {
					break
				}

				// Decode the JSON object from the line.
				var streamResponse ChatCompletionStreamResponse
				if err := json.Unmarshal(jsonLine, &streamResponse); err != nil {
					responseChannel <- ChatCompletionStreamResponse{Error: fmt.Errorf("error decoding stream response: %w", err)}
					continue
				}

				// Send the decoded response to the channel.
				responseChannel <- streamResponse
			}
		}
	}()

	// Return the response channel.
	return responseChannel, nil
}

// mapToStruct is a helper function to convert a map to a struct.
func mapToStruct(m map[string]interface{}, s interface{}) error {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, s)
}
