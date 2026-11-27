package sdk

import (
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
	Model           string                 `json:"model"`
	Input           EmbeddingRequestInputs `json:"input"`
	Metadata        map[string]any         `json:"metadata,omitempty"`
	EncodingFormat  *EncodingFormat        `json:"encoding_format,omitempty"`
	OutputDimension *int                   `json:"output_dimension,omitempty"`
	OutputDtype     *EmbeddingDtype        `json:"output_dtype,omitempty"`
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
	var request EmbeddingRequest
	if params != nil {
		request = *params
	}
	request.Model = model
	request.Input = input
	return c.CreateEmbeddings(request)
}

// CreateEmbeddings accepts text, image, and audio inputs, individually or in batches.
func (c *MistralClient) CreateEmbeddings(request EmbeddingRequest) (*EmbeddingResponse, error) {
	return requestTyped[EmbeddingResponse](c, http.MethodPost, request, "v1/embeddings")
}
