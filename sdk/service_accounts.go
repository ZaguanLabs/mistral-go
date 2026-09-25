package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ServiceAccount is a workspace-scoped identity for automated clients.
type ServiceAccount struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	CustomerID     string     `json:"customer_id"`
	OrganizationID string     `json:"organization_id"`
	WorkspaceID    string     `json:"workspace_id"`
	Description    *string    `json:"description"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
type CreateServiceAccountRequest struct {
	RoleIDs     []string `json:"role_ids,omitempty"`
	Name        string   `json:"name"`
	WorkspaceID string   `json:"workspace_id"`
	Description *string  `json:"description,omitempty"`
}
type UpdateServiceAccountRequest struct {
	// A nil Description clears the description, matching an explicit Python None.
	Description *string `json:"description"`
}
type ListServiceAccountsParams struct {
	WorkspaceID    string
	Offset         int
	Limit          int
	IncludeDeleted bool
}
type ServiceAccountRole struct {
	RoleID string `json:"role_id"`
}
type AssignableServiceAccountRole struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsCustomRole bool   `json:"is_custom_role"`
}
type ListServiceAccountsResponse struct {
	Items []ServiceAccount `json:"items"`
}
type ListServiceAccountRolesResponse struct {
	Items []ServiceAccountRole `json:"items"`
}
type ListAssignableServiceAccountRolesResponse struct {
	Items []AssignableServiceAccountRole `json:"items"`
}
type SetServiceAccountRolesRequest struct {
	RoleIDs []string `json:"role_ids"`
}

func (c *MistralClient) CreateServiceAccount(req *CreateServiceAccountRequest) (*ServiceAccount, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[ServiceAccount](c, http.MethodPost, req, "v1/service-accounts")
}
func (c *MistralClient) ListServiceAccounts(params *ListServiceAccountsParams) (*ListServiceAccountsResponse, error) {
	if params == nil {
		params = &ListServiceAccountsParams{Limit: 100}
	}
	var workspace *string
	if params.WorkspaceID != "" {
		workspace = &params.WorkspaceID
	}
	query := queryWithOptionalValues(map[string]any{"workspace_id": workspace, "offset": params.Offset, "limit": params.Limit, "include_deleted": params.IncludeDeleted})
	return requestTyped[ListServiceAccountsResponse](c, http.MethodGet, nil, appendQuery("v1/service-accounts", query))
}
func (c *MistralClient) GetServiceAccount(id string) (*ServiceAccount, error) {
	return requestTyped[ServiceAccount](c, http.MethodGet, nil, "v1/service-accounts/"+url.PathEscape(id))
}
func (c *MistralClient) UpdateServiceAccount(id string, req *UpdateServiceAccountRequest) (*ServiceAccount, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[ServiceAccount](c, http.MethodPatch, req, "v1/service-accounts/"+url.PathEscape(id))
}
func (c *MistralClient) DeleteServiceAccount(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/service-accounts/"+url.PathEscape(id), false, nil)
	return err
}
func (c *MistralClient) ListAssignableServiceAccountRoles(workspaceID string) (*ListAssignableServiceAccountRolesResponse, error) {
	return requestTyped[ListAssignableServiceAccountRolesResponse](c, http.MethodGet, nil, appendQuery("v1/service-accounts/assignable-roles", queryWithOptionalValues(map[string]any{"workspace_id": workspaceID})))
}
func (c *MistralClient) ListServiceAccountRoles(id string) (*ListServiceAccountRolesResponse, error) {
	return requestTyped[ListServiceAccountRolesResponse](c, http.MethodGet, nil, "v1/service-accounts/"+url.PathEscape(id)+"/roles")
}
func (c *MistralClient) SetServiceAccountRoles(id string, roleIDs []string) (*ListServiceAccountRolesResponse, error) {
	if roleIDs == nil {
		roleIDs = []string{}
	}
	return requestTyped[ListServiceAccountRolesResponse](c, http.MethodPut, &SetServiceAccountRolesRequest{RoleIDs: roleIDs}, "v1/service-accounts/"+url.PathEscape(id)+"/roles")
}

func requestTyped[T any](c *MistralClient, method string, value any, path string) (*T, error) {
	var body map[string]interface{}
	var err error
	if value != nil {
		body, err = structToMap(value)
		if err != nil {
			return nil, err
		}
	}
	response, err := c.requestMap(method, body, path)
	if err != nil {
		return nil, err
	}
	var result T
	if err := mapToStruct(map[string]interface{}(response), &result); err != nil {
		return nil, err
	}
	return &result, nil
}
