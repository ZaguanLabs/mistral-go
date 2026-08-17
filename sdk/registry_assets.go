package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegistryListParams struct {
	PageSize      *int
	PageToken     *string
	Alias         *string
	Fields        []string
	SortField     *string
	SortDirection *string
	SortBy        *string
}

type RegistryGetParams struct {
	Version *int
	Alias   *string
	Fields  []string
}

type PromptVariable struct {
	Name *string `json:"name,omitempty"`
}

type PromptDefinition struct {
	Content   string           `json:"content"`
	Variables []PromptVariable `json:"variables,omitempty"`
}

type CreatePromptRequest struct {
	Name         string           `json:"name"`
	Definition   PromptDefinition `json:"definition"`
	Title        *string          `json:"title,omitempty"`
	Description  *string          `json:"description,omitempty"`
	Notes        *string          `json:"notes,omitempty"`
	SharingScope *string          `json:"sharing_scope,omitempty"`
	Aliases      []string         `json:"aliases,omitempty"`
}

type CreatePromptVersionRequest struct {
	Definition PromptDefinition `json:"definition"`
	Notes      *string          `json:"notes,omitempty"`
	Aliases    []string         `json:"aliases,omitempty"`
}

type UpdatePromptMetadataRequest struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	SharingScope *string `json:"sharing_scope,omitempty"`
}

type UpdateRegistryVersionMetadataRequest struct {
	Notes   *string  `json:"notes,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
}

type SkillAssetContent struct {
	TextContent  *string `json:"textContent,omitempty"`
	RawContent   *string `json:"rawContent,omitempty"`
	IsExecutable *bool   `json:"isExecutable,omitempty"`
}

type SkillDefinition struct {
	Description *string                      `json:"description,omitempty"`
	Body        *string                      `json:"body,omitempty"`
	Assets      map[string]SkillAssetContent `json:"assets,omitempty"`
}

type CreateSkillRequest struct {
	Name         string          `json:"name"`
	Definition   SkillDefinition `json:"definition"`
	Notes        *string         `json:"notes,omitempty"`
	SharingScope *string         `json:"sharing_scope,omitempty"`
	Aliases      []string        `json:"aliases,omitempty"`
}

type CreateSkillVersionRequest struct {
	Definition SkillDefinition `json:"definition"`
	Notes      *string         `json:"notes,omitempty"`
	Aliases    []string        `json:"aliases,omitempty"`
}

func (c *MistralClient) ListPrompts(params *RegistryListParams) (APIResponse, error) {
	return c.listRegistryAssets("v2/prompts", params)
}

func (c *MistralClient) CreatePrompt(req *CreatePromptRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPost, body, "v2/prompts")
}

func (c *MistralClient) GetPrompt(promptID string, params *RegistryGetParams) (APIResponse, error) {
	return c.getRegistryAsset(fmt.Sprintf("v2/prompts/%s", promptID), params)
}

func (c *MistralClient) DeletePrompt(promptID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v2/prompts/%s", promptID))
}

func (c *MistralClient) UpdatePromptMetadata(promptID string, req *UpdatePromptMetadataRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v2/prompts/%s", promptID))
}

func (c *MistralClient) ListPromptVersions(promptID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v2/prompts/%s/versions", promptID))
}

func (c *MistralClient) CreatePromptVersion(promptID string, req *CreatePromptVersionRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v2/prompts/%s/versions", promptID))
}

func (c *MistralClient) GetPromptVersion(promptID string, version int, fields []string) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"fields": fields})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v2/prompts/%s/versions/%d", promptID, version), query))
}

func (c *MistralClient) UpdatePromptVersionMetadata(promptID string, version int, req *UpdateRegistryVersionMetadataRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v2/prompts/%s/versions/%d", promptID, version))
}

func (c *MistralClient) ListSkills(params *RegistryListParams) (APIResponse, error) {
	return c.listRegistryAssets("v2/skills", params)
}

func (c *MistralClient) CreateSkill(req *CreateSkillRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPost, body, "v2/skills")
}

func (c *MistralClient) GetSkill(skillID string, params *RegistryGetParams) (APIResponse, error) {
	return c.getRegistryAsset(fmt.Sprintf("v2/skills/%s", skillID), params)
}

func (c *MistralClient) DeleteSkill(skillID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v2/skills/%s", skillID))
}

func (c *MistralClient) UpdateSkillMetadata(skillID string, sharingScope *string) (APIResponse, error) {
	return c.requestMap(http.MethodPatch, optionalRequestMap(map[string]any{"sharing_scope": sharingScope}), fmt.Sprintf("v2/skills/%s", skillID))
}

func (c *MistralClient) ListSkillVersions(skillID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v2/skills/%s/versions", skillID))
}

func (c *MistralClient) CreateSkillVersion(skillID string, req *CreateSkillVersionRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v2/skills/%s/versions", skillID))
}

func (c *MistralClient) GetSkillVersion(skillID string, version int, fields []string) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"fields": fields})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v2/skills/%s/versions/%d", skillID, version), query))
}

func (c *MistralClient) UpdateSkillVersionMetadata(skillID string, version int, req *UpdateRegistryVersionMetadataRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v2/skills/%s/versions/%d", skillID, version))
}

func (c *MistralClient) GetUserIdentity() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/users/me")
}

type ListOrganizationsParams struct {
	Offset *int
	Limit  *int
}

type UserOrganization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListOrganizationsResponse struct {
	Organizations []UserOrganization `json:"organizations"`
}

type ListWorkspacesParams struct {
	OrganizationID *string
	Offset         *int
	Limit          *int
}

type UserWorkspace struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
}

type ListWorkspacesResponse struct {
	Workspaces []UserWorkspace `json:"workspaces"`
}

func (c *MistralClient) ListOrganizations(params *ListOrganizationsParams) (*ListOrganizationsResponse, error) {
	if params == nil {
		params = &ListOrganizationsParams{}
	}
	offset, limit := 0, 100
	if params.Offset != nil {
		offset = *params.Offset
	}
	if params.Limit != nil {
		limit = *params.Limit
	}
	query := queryWithOptionalValues(map[string]any{"offset": offset, "limit": limit})
	response, err := c.request(http.MethodGet, nil, appendQuery("v1/users/me/organizations", query), false, nil)
	if err != nil {
		return nil, err
	}
	data, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	var out ListOrganizationsResponse
	if err := mapToStruct(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MistralClient) ListWorkspaces(params *ListWorkspacesParams) (*ListWorkspacesResponse, error) {
	if params == nil {
		params = &ListWorkspacesParams{}
	}
	offset, limit := 0, 100
	if params.Offset != nil {
		offset = *params.Offset
	}
	if params.Limit != nil {
		limit = *params.Limit
	}
	query := queryWithOptionalValues(map[string]any{
		"organization_id": params.OrganizationID,
		"offset":          offset,
		"limit":           limit,
	})
	response, err := c.request(http.MethodGet, nil, appendQuery("v1/users/me/workspaces", query), false, nil)
	if err != nil {
		return nil, err
	}
	data, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	var out ListWorkspacesResponse
	if err := mapToStruct(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MistralClient) listRegistryAssets(path string, params *RegistryListParams) (APIResponse, error) {
	if params == nil {
		params = &RegistryListParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"page_size": params.PageSize, "page_token": params.PageToken, "alias": params.Alias,
		"fields": params.Fields, "sort.field": params.SortField,
		"sort.direction": params.SortDirection, "sort_by": params.SortBy,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery(path, query))
}

func (c *MistralClient) getRegistryAsset(path string, params *RegistryGetParams) (APIResponse, error) {
	if params == nil {
		params = &RegistryGetParams{}
	}
	query := queryWithOptionalValues(map[string]any{"version": params.Version, "alias": params.Alias, "fields": params.Fields})
	return c.requestMap(http.MethodGet, nil, appendQuery(path, query))
}

func structToMap(value any) (map[string]interface{}, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("convert request: %w", err)
	}
	return result, nil
}
