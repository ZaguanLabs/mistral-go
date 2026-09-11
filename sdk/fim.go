package sdk

import (
	"fmt"
	"io"
	"net/http"
)

// FIMRequestParams represents the parameters for the FIM method of MistralClient.
type FIMRequestParams struct {
	Model       string         `json:"model"`
	Prompt      string         `json:"prompt"`
	Suffix      *string        `json:"suffix,omitempty"`
	MaxTokens   *int           `json:"max_tokens,omitempty"`
	Temperature *float64       `json:"temperature,omitempty"`
	TopP        *float64       `json:"top_p,omitempty"`
	MinTokens   *int           `json:"min_tokens,omitempty"`
	RandomSeed  *int           `json:"random_seed,omitempty"`
	Stop        []string       `json:"stop,omitempty"`
	Stream      *bool          `json:"stream,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"` // Custom metadata (added in v1.10.0)
}

// FIMCompletionResponse represents the response from the FIM completion endpoint.
type FIMCompletionResponse struct {
	ID      string                        `json:"id"`
	Object  string                        `json:"object"`
	Created int                           `json:"created"`
	Model   string                        `json:"model"`
	Choices []FIMCompletionResponseChoice `json:"choices"`
	Usage   UsageInfo                     `json:"usage"`
}

// FIMCompletionResponseChoice represents a choice in the FIM completion response.
type FIMCompletionResponseChoice struct {
	Index        int          `json:"index"`
	Message      ChatMessage  `json:"message"`
	FinishReason FinishReason `json:"finish_reason,omitempty"`
}

// FIMCompletionStreamResponse represents a streaming response from FIM
type FIMCompletionStreamResponse struct {
	ID      string                              `json:"id"`
	Object  string                              `json:"object"`
	Created int                                 `json:"created"`
	Model   string                              `json:"model"`
	Choices []FIMCompletionResponseChoiceStream `json:"choices"`
	Error   error                               `json:"-"`
}

// FIMCompletionResponseChoiceStream represents a streaming choice
type FIMCompletionResponseChoiceStream struct {
	Index        int          `json:"index"`
	Delta        DeltaMessage `json:"delta"`
	FinishReason FinishReason `json:"finish_reason,omitempty"`
}

// FIM sends a FIM request and returns the completion response (non-streaming).
func (c *MistralClient) FIM(params *FIMRequestParams) (*FIMCompletionResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	// Ensure stream is false
	streamFalse := false
	params.Stream = &streamFalse

	requestData := map[string]interface{}{
		"model":  params.Model,
		"prompt": params.Prompt,
		"stream": false,
	}

	// Add optional parameters
	if params.Suffix != nil {
		requestData["suffix"] = *params.Suffix
	}
	if params.MaxTokens != nil {
		requestData["max_tokens"] = *params.MaxTokens
	}
	if params.Temperature != nil {
		requestData["temperature"] = *params.Temperature
	}
	if params.TopP != nil {
		requestData["top_p"] = *params.TopP
	}
	if params.MinTokens != nil {
		requestData["min_tokens"] = *params.MinTokens
	}
	if params.RandomSeed != nil {
		requestData["random_seed"] = *params.RandomSeed
	}
	if len(params.Stop) > 0 {
		requestData["stop"] = params.Stop
	}

	response, err := c.request(http.MethodPost, requestData, "v1/fim/completions", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var fimResponse FIMCompletionResponse
	err = mapToStruct(respData, &fimResponse)
	if err != nil {
		return nil, err
	}

	return &fimResponse, nil
}

// FIMStream sends a streaming FIM request and returns a channel of responses
//
// Parameters:
//   - params: FIM request parameters (stream will be set to true automatically)
//
// Returns a channel that streams FIM completion chunks
func (c *MistralClient) FIMStream(params *FIMRequestParams) (<-chan FIMCompletionStreamResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	// Ensure stream is true
	streamTrue := true
	params.Stream = &streamTrue

	requestData := map[string]interface{}{
		"model":  params.Model,
		"prompt": params.Prompt,
		"stream": true,
	}

	// Add optional parameters
	if params.Suffix != nil {
		requestData["suffix"] = *params.Suffix
	}
	if params.MaxTokens != nil {
		requestData["max_tokens"] = *params.MaxTokens
	}
	if params.Temperature != nil {
		requestData["temperature"] = *params.Temperature
	}
	if params.TopP != nil {
		requestData["top_p"] = *params.TopP
	}
	if params.MinTokens != nil {
		requestData["min_tokens"] = *params.MinTokens
	}
	if params.RandomSeed != nil {
		requestData["random_seed"] = *params.RandomSeed
	}
	if len(params.Stop) > 0 {
		requestData["stop"] = params.Stop
	}

	// Create response channel
	var responseChan <-chan FIMCompletionStreamResponse

	response, err := c.request(http.MethodPost, requestData, "v1/fim/completions", true, nil)
	if err != nil {
		return nil, err
	}

	respBody, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	// Start streaming in a goroutine
	responseChan = streamJSON(c.requestContext(), respBody, func(err error) FIMCompletionStreamResponse { return FIMCompletionStreamResponse{Error: err} })

	return responseChan, nil
}
