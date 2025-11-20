package sdk

import (
	"fmt"
	"net/http"
)

// LibraryAccess represents access control for a library
type LibraryAccess struct {
	UserID      string   `json:"user_id"`
	Permissions []string `json:"permissions"`
}

// AccessListResponse represents a list of library accesses
type AccessListResponse struct {
	Object string          `json:"object"`
	Data   []LibraryAccess `json:"data"`
}

// UpdateAccessRequest represents a request to update library access
type UpdateAccessRequest struct {
	UserID      string   `json:"user_id"`
	Permissions []string `json:"permissions"`
}

// DeleteAccessResponse represents the response from deleting access
type DeleteAccessResponse struct {
	UserID  string `json:"user_id"`
	Deleted bool   `json:"deleted"`
}

// ListLibraryAccesses lists all accesses for a library
func (c *MistralClient) ListLibraryAccesses(libraryID string) (*AccessListResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/accesses", libraryID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listResponse AccessListResponse
	err = mapToStruct(respData, &listResponse)
	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}

// UpdateOrCreateLibraryAccess updates or creates library access for a user
func (c *MistralClient) UpdateOrCreateLibraryAccess(libraryID string, req *UpdateAccessRequest) (*LibraryAccess, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"user_id":     req.UserID,
		"permissions": req.Permissions,
	}

	response, err := c.request(http.MethodPut, reqMap, fmt.Sprintf("v1/libraries/%s/accesses", libraryID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var access LibraryAccess
	err = mapToStruct(respData, &access)
	if err != nil {
		return nil, err
	}

	return &access, nil
}

// DeleteLibraryAccess removes access for a user from a library
func (c *MistralClient) DeleteLibraryAccess(libraryID, userID string) (*DeleteAccessResponse, error) {
	response, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/libraries/%s/accesses/%s", libraryID, userID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var deleteResponse DeleteAccessResponse
	err = mapToStruct(respData, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}
