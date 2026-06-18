package sdk

import (
	"fmt"
	"net/http"
	"net/url"
)

// ShareEnum represents the access level granted to an entity.
type ShareEnum string

const (
	ShareEnumOwner  ShareEnum = "Owner"
	ShareEnumEditor ShareEnum = "Editor"
	ShareEnumViewer ShareEnum = "Viewer"
)

// EntityType represents a share target entity type.
type EntityType string

const (
	EntityTypeUser         EntityType = "User"
	EntityTypeWorkspace    EntityType = "Workspace"
	EntityTypeOrganization EntityType = "Organization"
)

// LibraryShare represents access control for a library share entry.
type LibraryShare struct {
	Level         ShareEnum  `json:"level"`
	ShareWithUUID string     `json:"share_with_uuid"`
	ShareWithType EntityType `json:"share_with_type"`
	OrgID         *string    `json:"org_id,omitempty"`

	// Backward-compatible fields used by previous SDK versions.
	UserID      string   `json:"user_id,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// LibraryAccess is kept as an alias for backward compatibility.
type LibraryAccess = LibraryShare

// AccessListResponse represents a list of library shares.
type AccessListResponse struct {
	Object string         `json:"object"`
	Data   []LibraryShare `json:"data"`
}

// UpdateAccessRequest represents a request to update/create a library share.
type UpdateAccessRequest struct {
	Level         ShareEnum  `json:"level"`
	ShareWithUUID string     `json:"share_with_uuid"`
	ShareWithType EntityType `json:"share_with_type"`
	OrgID         *string    `json:"org_id,omitempty"`

	// Backward-compatible fields used by previous SDK versions.
	UserID      string   `json:"user_id,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// DeleteAccessResponse represents the response from deleting a library share.
type DeleteAccessResponse struct {
	Level         ShareEnum  `json:"level"`
	ShareWithUUID string     `json:"share_with_uuid"`
	ShareWithType EntityType `json:"share_with_type"`
	OrgID         *string    `json:"org_id,omitempty"`

	// Backward-compatible fields used by previous SDK versions.
	UserID  string `json:"user_id,omitempty"`
	Deleted bool   `json:"deleted,omitempty"`
}

// ListLibraryAccesses lists all shares for a library.
func (c *MistralClient) ListLibraryAccesses(libraryID string) (*AccessListResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/share", libraryID), false, nil)
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

// UpdateOrCreateLibraryAccess updates or creates a library share entry.
func (c *MistralClient) UpdateOrCreateLibraryAccess(libraryID string, req *UpdateAccessRequest) (*LibraryShare, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"level":           req.Level,
		"share_with_uuid": req.ShareWithUUID,
		"share_with_type": req.ShareWithType,
	}
	if req.UserID != "" && req.ShareWithUUID == "" {
		reqMap["share_with_uuid"] = req.UserID
	}
	if req.Level == "" && len(req.Permissions) > 0 {
		reqMap["level"] = ShareEnumViewer
	}
	if req.ShareWithType == "" {
		reqMap["share_with_type"] = EntityTypeUser
	}
	if req.OrgID != nil {
		reqMap["org_id"] = *req.OrgID
	}

	response, err := c.request(http.MethodPut, reqMap, fmt.Sprintf("v1/libraries/%s/share", libraryID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var access LibraryShare
	err = mapToStruct(respData, &access)
	if err != nil {
		return nil, err
	}

	// Fill backward-compatible fields when only share fields are present.
	if access.UserID == "" {
		access.UserID = access.ShareWithUUID
	}
	if len(access.Permissions) == 0 && access.Level != "" {
		switch access.Level {
		case ShareEnumOwner, ShareEnumEditor:
			access.Permissions = []string{"read", "write"}
		default:
			access.Permissions = []string{"read"}
		}
	}

	return &access, nil
}

// DeleteLibraryShare removes a share entry from a library.
func (c *MistralClient) DeleteLibraryShare(libraryID, shareWithUUID string, shareWithType EntityType, orgID *string) (*DeleteAccessResponse, error) {
	query := url.Values{}
	query.Add("share_with_uuid", shareWithUUID)
	query.Add("share_with_type", string(shareWithType))
	if orgID != nil {
		query.Add("org_id", *orgID)
	}

	path := fmt.Sprintf("v1/libraries/%s/share?%s", libraryID, query.Encode())
	response, err := c.request(http.MethodDelete, nil, path, false, nil)
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

	var deleteResponse DeleteAccessResponse
	err = mapToStruct(respData, &deleteResponse)
	if err != nil {
		return nil, err
	}

	// Fill backward-compatible fields.
	if deleteResponse.UserID == "" {
		deleteResponse.UserID = deleteResponse.ShareWithUUID
	}

	return &deleteResponse, nil
}

// DeleteLibraryAccess removes a user share entry from a library.
// Kept for backward compatibility.
func (c *MistralClient) DeleteLibraryAccess(libraryID, userID string) (*DeleteAccessResponse, error) {
	return c.DeleteLibraryShare(libraryID, userID, EntityTypeUser, nil)
}
