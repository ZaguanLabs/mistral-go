package sdk

import (
	"fmt"
	"net/http"
)

// EmbeddingObject represents an embedding object in the response.
type EmbeddingObject struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

// EmbeddingResponse represents the response from the embeddings endpoint.
type EmbeddingResponse struct {
	ID     string            `json:"id"`
	Object string            `json:"object"`
	Data   []EmbeddingObject `json:"data"`
	Model  string            `json:"model"`
	Usage  UsageInfo         `json:"usage"`
}

// EmbeddingRequest represents parameters for creating embeddings
type EmbeddingRequest struct {
	Model           string          `json:"model"`
	Input           []string        `json:"input"`
	EncodingFormat  *EncodingFormat `json:"encoding_format,omitempty"`
	OutputDimension *int            `json:"output_dimension,omitempty"`
	OutputDtype     *EmbeddingDtype `json:"output_dtype,omitempty"`
}

// Embeddings creates embeddings for the given input texts (simple version)
func (c *MistralClient) Embeddings(model string, input []string) (*EmbeddingResponse, error) {
	return c.EmbeddingsWithParams(model, input, nil)
}

// EmbeddingsWithParams creates embeddings with additional parameters
//
// Parameters:
//   - model: The model ID to use for embeddings
//   - input: The text inputs to embed
//   - params: Optional parameters (encoding_format, output_dimension, output_dtype)
//
// Returns embeddings for the input texts
func (c *MistralClient) EmbeddingsWithParams(model string, input []string, params *EmbeddingRequest) (*EmbeddingResponse, error) {
	if params == nil {
		params = &EmbeddingRequest{}
	}

	params.Model = model
	params.Input = input

	requestData := map[string]interface{}{
		"model": model,
		"input": input,
	}

	// Add optional parameters
	if params.EncodingFormat != nil {
		requestData["encoding_format"] = *params.EncodingFormat
	}
	if params.OutputDimension != nil {
		requestData["output_dimension"] = *params.OutputDimension
	}
	if params.OutputDtype != nil {
		requestData["output_dtype"] = *params.OutputDtype
	}

	response, err := c.request(http.MethodPost, requestData, "v1/embeddings", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var embeddingResponse EmbeddingResponse
	err = mapToStruct(respData, &embeddingResponse)
	if err != nil {
		return nil, err
	}

	return &embeddingResponse, nil
}
