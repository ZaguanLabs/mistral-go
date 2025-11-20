package sdk

import (
"fmt"
"net/http"
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
Model        string                 `json:"model"`
Name         *string                `json:"name,omitempty"`
Description  *string                `json:"description,omitempty"`
Instructions *string                `json:"instructions,omitempty"`
Tools        []Tool                 `json:"tools,omitempty"`
Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// UpdateMistralAgentRequest represents a request to update an agent
type UpdateMistralAgentRequest struct {
Model        *string                `json:"model,omitempty"`
Name         *string                `json:"name,omitempty"`
Description  *string                `json:"description,omitempty"`
Instructions *string                `json:"instructions,omitempty"`
Tools        []Tool                 `json:"tools,omitempty"`
Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// MistralAgentListResponse represents a list of agents
type MistralAgentListResponse struct {
Object string         `json:"object"`
Data   []MistralAgent `json:"data"`
Total  int            `json:"total"`
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
if req.Metadata != nil {
reqMap["metadata"] = req.Metadata
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
path := fmt.Sprintf("v1/agents?page=%d", page)

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
response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/agents/%s", agentID), false, nil)
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
if req.Metadata != nil {
reqMap["metadata"] = req.Metadata
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
