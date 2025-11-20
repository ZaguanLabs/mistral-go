package sdk

import (
	"fmt"
	"net/http"
)

// ModelPermission represents the permissions of a model.
type ModelPermission struct {
	ID                 string `json:"id"`
	Object             string `json:"object"`
	Created            int    `json:"created"`
	AllowCreateEngine  bool   `json:"allow_create_engine"`
	AllowSampling      bool   `json:"allow_sampling"`
	AllowLogprobs      bool   `json:"allow_logprobs"`
	AllowSearchIndices bool   `json:"allow_search_indices"`
	AllowView          bool   `json:"allow_view"`
	AllowFineTuning    bool   `json:"allow_fine_tuning"`
	Organization       string `json:"organization"`
	Group              string `json:"group,omitempty"`
	IsBlocking         bool   `json:"is_blocking"`
}

// ModelCard represents a model card.
type ModelCard struct {
	ID         string            `json:"id"`
	Object     string            `json:"object"`
	Created    int               `json:"created"`
	OwnedBy    string            `json:"owned_by"`
	Root       string            `json:"root,omitempty"`
	Parent     string            `json:"parent,omitempty"`
	Permission []ModelPermission `json:"permission"`
}

// ModelList represents a list of models.
type ModelList struct {
	Object string      `json:"object"`
	Data   []ModelCard `json:"data"`
}

// DeleteModelResponse represents the response from deleting a model
type DeleteModelResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// FineTunedModel represents a fine-tuned model with full details
type FineTunedModel struct {
	ID          string  `json:"id"`
	Object      string  `json:"object"`
	Created     int64   `json:"created"`
	OwnedBy     string  `json:"owned_by"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Archived    bool    `json:"archived"`
	Job         *string `json:"job,omitempty"`
}

// ArchiveModelResponse represents the response from archiving a model
type ArchiveModelResponse struct {
	ID       string `json:"id"`
	Object   string `json:"object"`
	Archived bool   `json:"archived"`
}

// UnarchiveModelResponse represents the response from unarchiving a model
type UnarchiveModelResponse struct {
	ID       string `json:"id"`
	Object   string `json:"object"`
	Archived bool   `json:"archived"`
}

func (c *MistralClient) ListModels() (*ModelList, error) {
	response, err := c.request(http.MethodGet, nil, "v1/models", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var modelList ModelList
	err = mapToStruct(respData, &modelList)
	if err != nil {
		return nil, err
	}

	return &modelList, nil
}

// RetrieveModel retrieves information about a specific model
//
// Parameters:
//   - modelID: The ID of the model to retrieve
//
// Returns detailed information about the model
func (c *MistralClient) RetrieveModel(modelID string) (*ModelCard, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/models/%s", modelID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var model ModelCard
	err = mapToStruct(respData, &model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// DeleteModel deletes a fine-tuned model
//
// Parameters:
//   - modelID: The ID of the model to delete
//
// Returns confirmation of deletion
func (c *MistralClient) DeleteModel(modelID string) (*DeleteModelResponse, error) {
	response, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/models/%s", modelID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var deleteResponse DeleteModelResponse
	err = mapToStruct(respData, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}

// UpdateModelRequest represents a request to update a model
type UpdateModelRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateModel updates a fine-tuned model's name or description
//
// Parameters:
//   - modelID: The ID of the model to update
//   - req: The update request with name and/or description
//
// Returns the updated model information
func (c *MistralClient) UpdateModel(modelID string, req *UpdateModelRequest) (*FineTunedModel, error) {
	if req == nil {
		return nil, fmt.Errorf("update request cannot be nil")
	}

	reqMap := make(map[string]interface{})
	if req.Name != nil {
		reqMap["name"] = req.Name
	}
	if req.Description != nil {
		reqMap["description"] = req.Description
	}

	response, err := c.request(http.MethodPatch, reqMap, fmt.Sprintf("v1/fine_tuning/models/%s", modelID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var model FineTunedModel
	err = mapToStruct(respData, &model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// ArchiveModel archives a fine-tuned model
//
// Parameters:
//   - modelID: The ID of the model to archive
//
// Returns confirmation of archival
func (c *MistralClient) ArchiveModel(modelID string) (*ArchiveModelResponse, error) {
	response, err := c.request(http.MethodPost, nil, fmt.Sprintf("v1/fine_tuning/models/%s/archive", modelID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var archiveResponse ArchiveModelResponse
	err = mapToStruct(respData, &archiveResponse)
	if err != nil {
		return nil, err
	}

	return &archiveResponse, nil
}

// UnarchiveModel unarchives a fine-tuned model
//
// Parameters:
//   - modelID: The ID of the model to unarchive
//
// Returns confirmation of unarchival
func (c *MistralClient) UnarchiveModel(modelID string) (*UnarchiveModelResponse, error) {
	response, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/fine_tuning/models/%s/archive", modelID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var unarchiveResponse UnarchiveModelResponse
	err = mapToStruct(respData, &unarchiveResponse)
	if err != nil {
		return nil, err
	}

	return &unarchiveResponse, nil
}
