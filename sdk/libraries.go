package sdk

import (
	"fmt"
	"net/http"
)

// Library represents a document library for RAG
type Library struct {
	ID                            string  `json:"id"`
	Object                        string  `json:"object,omitempty"`
	Name                          string  `json:"name"`
	CreatedAt                     int64   `json:"created_at"`
	UpdatedAt                     int64   `json:"updated_at"`
	OwnerID                       *string `json:"owner_id"`
	OwnerType                     string  `json:"owner_type"`
	TotalSize                     int     `json:"total_size"`
	NbDocuments                   int     `json:"nb_documents"`
	ChunkSize                     *int    `json:"chunk_size"`
	Emoji                         *string `json:"emoji,omitempty"`
	Description                   *string `json:"description,omitempty"`
	GeneratedDescription          *string `json:"generated_description,omitempty"`
	ExplicitUserMembersCount      *int    `json:"explicit_user_members_count,omitempty"`
	ExplicitWorkspaceMembersCount *int    `json:"explicit_workspace_members_count,omitempty"`
	OrgSharingRole                *string `json:"org_sharing_role,omitempty"`
	GeneratedName                 *string `json:"generated_name,omitempty"`
}

// LibraryListResponse represents a list of libraries
type LibraryListResponse struct {
	Object string    `json:"object"`
	Data   []Library `json:"data"`
}

// CreateLibraryRequest represents a request to create a library
type CreateLibraryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	ChunkSize   *int    `json:"chunk_size,omitempty"`
}

// UpdateLibraryRequest represents a request to update a library
type UpdateLibraryRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// DeleteLibraryResponse represents the response from deleting a library
type DeleteLibraryResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// ListLibraries lists all libraries
func (c *MistralClient) ListLibraries() (*LibraryListResponse, error) {
	response, err := c.request(http.MethodGet, nil, "v1/libraries", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listResponse LibraryListResponse
	err = mapToStruct(respData, &listResponse)
	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}

// CreateLibrary creates a new library
func (c *MistralClient) CreateLibrary(req *CreateLibraryRequest) (*Library, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"name": req.Name,
	}

	if req.Description != nil {
		reqMap["description"] = *req.Description
	}
	if req.ChunkSize != nil {
		reqMap["chunk_size"] = *req.ChunkSize
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/libraries", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var library Library
	err = mapToStruct(respData, &library)
	if err != nil {
		return nil, err
	}

	return &library, nil
}

// GetLibrary retrieves a specific library
func (c *MistralClient) GetLibrary(libraryID string) (*Library, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s", libraryID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var library Library
	err = mapToStruct(respData, &library)
	if err != nil {
		return nil, err
	}

	return &library, nil
}

// UpdateLibrary updates a library
func (c *MistralClient) UpdateLibrary(libraryID string, req *UpdateLibraryRequest) (*Library, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := make(map[string]interface{})

	if req.Name != nil {
		reqMap["name"] = *req.Name
	}
	if req.Description != nil {
		reqMap["description"] = *req.Description
	}

	response, err := c.request(http.MethodPut, reqMap, fmt.Sprintf("v1/libraries/%s", libraryID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var library Library
	err = mapToStruct(respData, &library)
	if err != nil {
		return nil, err
	}

	return &library, nil
}

// DeleteLibrary deletes a library
func (c *MistralClient) DeleteLibrary(libraryID string) (*DeleteLibraryResponse, error) {
	response, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/libraries/%s", libraryID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	if len(respData) == 0 {
		return nil, nil
	}

	var deleteResponse DeleteLibraryResponse
	err = mapToStruct(respData, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}
