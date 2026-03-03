package sdk

import (
	"fmt"
	"net/http"
	"net/url"
)

// MistralAgent represents a Mistral agent
type MistralAgent struct {
	ID           string                 `json:"id"`
	Object       string                 `json:"object"`
	Model        string                 `json:"model"`
	Name         *string                `json:"name,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Instructions *string                `json:"instructions,omitempty"`
	Tools        []Tool                 `json:"tools,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Created      int64                  `json:"created"`
	Updated      int64                  `json:"updated"`
}

// CreateMistralAgentRequest represents a request to create an agent
type CreateMistralAgentRequest struct {
	Model          string                 `json:"model"`
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Instructions   *string                `json:"instructions,omitempty"`
	Tools          []Tool                 `json:"tools,omitempty"`
	CompletionArgs map[string]any         `json:"completion_args,omitempty"`
	Handoffs       []string               `json:"handoffs,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	VersionMessage *string                `json:"version_message,omitempty"`
}

// UpdateMistralAgentRequest represents a request to update an agent
type UpdateMistralAgentRequest struct {
	Model          *string                `json:"model,omitempty"`
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Instructions   *string                `json:"instructions,omitempty"`
	Tools          []Tool                 `json:"tools,omitempty"`
	CompletionArgs map[string]any         `json:"completion_args,omitempty"`
	Handoffs       []string               `json:"handoffs,omitempty"`
	DeploymentChat *bool                  `json:"deployment_chat,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	VersionMessage *string                `json:"version_message,omitempty"`
}

// MistralAgentListResponse represents a list of agents
type MistralAgentListResponse struct {
	Object string         `json:"object"`
	Data   []MistralAgent `json:"data"`
	Total  int            `json:"total"`
}

// ListMistralAgentsParams represents list filters for beta agents.
type ListMistralAgentsParams struct {
	Page           *int            `json:"page,omitempty"`
	PageSize       *int            `json:"page_size,omitempty"`
	DeploymentChat *bool           `json:"deployment_chat,omitempty"`
	Sources        []RequestSource `json:"sources,omitempty"`
	Name           *string         `json:"name,omitempty"`
	Search         *string         `json:"search,omitempty"`
	ID             *string         `json:"id,omitempty"`
	Metadata       map[string]any  `json:"metadata,omitempty"`
}

// AgentAliasResponse represents an alias attached to an agent version.
type AgentAliasResponse struct {
	Alias   string `json:"alias"`
	Version int    `json:"version"`
}

// CreateMistralAgent creates a new Mistral agent
func (c *MistralClient) CreateMistralAgent(req *CreateMistralAgentRequest) (*MistralAgent, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"model": req.Model,
	}

	if req.Name != nil {
		reqMap["name"] = *req.Name
	}
	if req.Description != nil {
		reqMap["description"] = *req.Description
	}
	if req.Instructions != nil {
		reqMap["instructions"] = *req.Instructions
	}
	if len(req.Tools) > 0 {
		reqMap["tools"] = req.Tools
	}
	if req.CompletionArgs != nil {
		reqMap["completion_args"] = req.CompletionArgs
	}
	if len(req.Handoffs) > 0 {
		reqMap["handoffs"] = req.Handoffs
	}
	if req.Metadata != nil {
		reqMap["metadata"] = req.Metadata
	}
	if req.VersionMessage != nil {
		reqMap["version_message"] = *req.VersionMessage
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/agents", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var agent MistralAgent
	err = mapToStruct(respData, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// ListMistralAgents lists all Mistral agents
func (c *MistralClient) ListMistralAgents(page int) (*MistralAgentListResponse, error) {
	return c.ListMistralAgentsWithParams(&ListMistralAgentsParams{Page: &page})
}

// ListMistralAgentsWithParams lists all Mistral agents with filters.
func (c *MistralClient) ListMistralAgentsWithParams(params *ListMistralAgentsParams) (*MistralAgentListResponse, error) {
	if params == nil {
		params = &ListMistralAgentsParams{}
	}

	query := url.Values{}
	if params.Page != nil {
		query.Add("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.PageSize != nil {
		query.Add("page_size", fmt.Sprintf("%d", *params.PageSize))
	}
	if params.DeploymentChat != nil {
		query.Add("deployment_chat", fmt.Sprintf("%t", *params.DeploymentChat))
	}
	for _, source := range params.Sources {
		query.Add("sources", string(source))
	}
	if params.Name != nil {
		query.Add("name", *params.Name)
	}
	if params.Search != nil {
		query.Add("search", *params.Search)
	}
	if params.ID != nil {
		query.Add("id", *params.ID)
	}

	path := "v1/agents"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listResponse MistralAgentListResponse
	err = mapToStruct(respData, &listResponse)
	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}

// GetMistralAgent retrieves a specific Mistral agent
func (c *MistralClient) GetMistralAgent(agentID string) (*MistralAgent, error) {
	return c.GetMistralAgentWithVersion(agentID, nil)
}

// GetMistralAgentWithVersion retrieves a specific Mistral agent and optional version/alias.
func (c *MistralClient) GetMistralAgentWithVersion(agentID string, agentVersion *string) (*MistralAgent, error) {
	path := fmt.Sprintf("v1/agents/%s", agentID)
	if agentVersion != nil {
		path += "?" + url.Values{"agent_version": []string{*agentVersion}}.Encode()
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var agent MistralAgent
	err = mapToStruct(respData, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// UpdateMistralAgent updates a Mistral agent
func (c *MistralClient) UpdateMistralAgent(agentID string, req *UpdateMistralAgentRequest) (*MistralAgent, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := make(map[string]interface{})

	if req.Model != nil {
		reqMap["model"] = *req.Model
	}
	if req.Name != nil {
		reqMap["name"] = *req.Name
	}
	if req.Description != nil {
		reqMap["description"] = *req.Description
	}
	if req.Instructions != nil {
		reqMap["instructions"] = *req.Instructions
	}
	if len(req.Tools) > 0 {
		reqMap["tools"] = req.Tools
	}
	if req.CompletionArgs != nil {
		reqMap["completion_args"] = req.CompletionArgs
	}
	if len(req.Handoffs) > 0 {
		reqMap["handoffs"] = req.Handoffs
	}
	if req.DeploymentChat != nil {
		reqMap["deployment_chat"] = *req.DeploymentChat
	}
	if req.Metadata != nil {
		reqMap["metadata"] = req.Metadata
	}
	if req.VersionMessage != nil {
		reqMap["version_message"] = *req.VersionMessage
	}

	response, err := c.request(http.MethodPatch, reqMap, fmt.Sprintf("v1/agents/%s", agentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var agent MistralAgent
	err = mapToStruct(respData, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// DeleteMistralAgent deletes a Mistral agent
//
// Parameters:
//   - agentID: The ID of the agent to delete
//
// Returns an error if the deletion fails
func (c *MistralClient) DeleteMistralAgent(agentID string) error {
	_, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/agents/%s", agentID), false, nil)
	return err
}

// UpdateMistralAgentVersion switches the active version for an agent.
func (c *MistralClient) UpdateMistralAgentVersion(agentID string, version int) (*MistralAgent, error) {
	response, err := c.request(http.MethodPatch, map[string]interface{}{"version": version}, fmt.Sprintf("v1/agents/%s/version", agentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var agent MistralAgent
	err = mapToStruct(respData, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// ListMistralAgentVersions lists available versions for an agent.
func (c *MistralClient) ListMistralAgentVersions(agentID string, page, pageSize *int) (*MistralAgentListResponse, error) {
	query := url.Values{}
	if page != nil {
		query.Add("page", fmt.Sprintf("%d", *page))
	}
	if pageSize != nil {
		query.Add("page_size", fmt.Sprintf("%d", *pageSize))
	}

	path := fmt.Sprintf("v1/agents/%s/versions", agentID)
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listResponse MistralAgentListResponse
	err = mapToStruct(respData, &listResponse)
	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}

// GetMistralAgentVersion retrieves a specific version for an agent.
func (c *MistralClient) GetMistralAgentVersion(agentID, version string) (*MistralAgent, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/agents/%s/versions/%s", agentID, version), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var agent MistralAgent
	err = mapToStruct(respData, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

// CreateOrUpdateMistralAgentAlias creates/updates a version alias.
func (c *MistralClient) CreateOrUpdateMistralAgentAlias(agentID, alias string, version int) (*AgentAliasResponse, error) {
	response, err := c.request(http.MethodPut, map[string]interface{}{"alias": alias, "version": version}, fmt.Sprintf("v1/agents/%s/aliases", agentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var aliasResp AgentAliasResponse
	err = mapToStruct(respData, &aliasResp)
	if err != nil {
		return nil, err
	}

	return &aliasResp, nil
}

// ListMistralAgentAliases lists aliases for an agent.
func (c *MistralClient) ListMistralAgentAliases(agentID string) ([]AgentAliasResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/agents/%s/aliases", agentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	rawData, ok := respData["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid alias list response format")
	}

	aliases := make([]AgentAliasResponse, 0, len(rawData))
	for _, item := range rawData {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		var aliasResp AgentAliasResponse
		if err := mapToStruct(itemMap, &aliasResp); err == nil {
			aliases = append(aliases, aliasResp)
		}
	}

	return aliases, nil
}

// DeleteMistralAgentAlias deletes an alias for an agent.
func (c *MistralClient) DeleteMistralAgentAlias(agentID, alias string) error {
	path := fmt.Sprintf("v1/agents/%s/aliases?%s", agentID, url.Values{"alias": []string{alias}}.Encode())
	_, err := c.request(http.MethodDelete, nil, path, false, nil)
	return err
}
