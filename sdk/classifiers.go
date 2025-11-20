package sdk

import (
	"fmt"
	"net/http"
)

// ClassificationInput represents input for classification/moderation
type ClassificationInput interface{}

// ClassificationRequest represents a request for content moderation
type ClassificationRequest struct {
	Model  string                `json:"model"`
	Inputs []ClassificationInput `json:"inputs"`
}

// ModerationCategory represents a moderation category result
type ModerationCategory struct {
	CategoryName string  `json:"category_name"`
	Score        float64 `json:"score"`
}

// ModerationResult represents the moderation result for a single input
type ModerationResult struct {
	Categories []ModerationCategory `json:"categories"`
}

// ModerationResponse represents the response from content moderation
type ModerationResponse struct {
	ID      string             `json:"id"`
	Model   string             `json:"model"`
	Results []ModerationResult `json:"results"`
}

// Moderate performs content moderation/classification
//
// Parameters:
//   - model: The model ID to use for moderation (e.g., "mistral-moderation-latest")
//   - inputs: The text inputs to moderate (can be strings or structured inputs)
//
// Returns moderation results with category scores
func (c *MistralClient) Moderate(model string, inputs []ClassificationInput) (*ModerationResponse, error) {
	reqMap := map[string]interface{}{
		"model":  model,
		"inputs": inputs,
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/moderations", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var moderationResponse ModerationResponse
	err = mapToStruct(respData, &moderationResponse)
	if err != nil {
		return nil, err
	}

	return &moderationResponse, nil
}

// ModerateText is a convenience function for moderating simple text inputs
//
// Parameters:
//   - model: The model ID to use for moderation
//   - texts: The text strings to moderate
//
// Returns moderation results
func (c *MistralClient) ModerateText(model string, texts []string) (*ModerationResponse, error) {
	inputs := make([]ClassificationInput, len(texts))
	for i, text := range texts {
		inputs[i] = text
	}

	return c.Moderate(model, inputs)
}
