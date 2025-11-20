package sdk

import (
	"fmt"
	"net/http"
)

// ConversationInput represents input for a conversation
type ConversationInput struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// ConversationStartRequest represents a request to start a conversation
type ConversationStartRequest struct {
	Inputs         []ConversationInput    `json:"inputs"`
	Instructions   *string                `json:"instructions,omitempty"`
	Tools          []Tool                 `json:"tools,omitempty"`
	CompletionArgs map[string]interface{} `json:"completion_args,omitempty"`
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
}

// ConversationResponse represents a conversation response
type ConversationResponse struct {
	ConversationID string               `json:"conversation_id"`
	Object         string               `json:"object"`
	Created        int64                `json:"created"`
	Status         string               `json:"status"`
	Outputs        []ConversationOutput `json:"outputs,omitempty"`
}

// ConversationOutput represents output from a conversation
type ConversationOutput struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// ConversationListResponse represents a list of conversations
type ConversationListResponse struct {
	Object string                 `json:"object"`
	Data   []ConversationResponse `json:"data"`
	Total  int                    `json:"total"`
}

// ConversationHistoryResponse represents conversation history
type ConversationHistoryResponse struct {
	ConversationID string              `json:"conversation_id"`
	Entries        []ConversationEntry `json:"entries"`
}

// ConversationEntry represents a single entry in conversation history
type ConversationEntry struct {
	Type      string      `json:"type"`
	Content   interface{} `json:"content"`
	Timestamp int64       `json:"timestamp"`
}

// StartConversation starts a new conversation
func (c *MistralClient) StartConversation(req *ConversationStartRequest) (*ConversationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"inputs": req.Inputs,
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
	if req.Name != nil {
		reqMap["name"] = *req.Name
	}
	if req.Description != nil {
		reqMap["description"] = *req.Description
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/conversations", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var convResponse ConversationResponse
	err = mapToStruct(respData, &convResponse)
	if err != nil {
		return nil, err
	}

	return &convResponse, nil
}

// ListConversations lists all conversations
func (c *MistralClient) ListConversations(page int) (*ConversationListResponse, error) {
	path := fmt.Sprintf("v1/conversations?page=%d", page)

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listResponse ConversationListResponse
	err = mapToStruct(respData, &listResponse)
	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}

// GetConversation retrieves a specific conversation
func (c *MistralClient) GetConversation(conversationID string) (*ConversationResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/conversations/%s", conversationID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var convResponse ConversationResponse
	err = mapToStruct(respData, &convResponse)
	if err != nil {
		return nil, err
	}

	return &convResponse, nil
}

// AppendToConversation appends inputs to an existing conversation
func (c *MistralClient) AppendToConversation(conversationID string, inputs []ConversationInput) (*ConversationResponse, error) {
	reqMap := map[string]interface{}{
		"inputs": inputs,
	}

	response, err := c.request(http.MethodPost, reqMap, fmt.Sprintf("v1/conversations/%s", conversationID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var convResponse ConversationResponse
	err = mapToStruct(respData, &convResponse)
	if err != nil {
		return nil, err
	}

	return &convResponse, nil
}

// GetConversationHistory retrieves the history of a conversation
func (c *MistralClient) GetConversationHistory(conversationID string) (*ConversationHistoryResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/conversations/%s/history", conversationID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var historyResponse ConversationHistoryResponse
	err = mapToStruct(respData, &historyResponse)
	if err != nil {
		return nil, err
	}

	return &historyResponse, nil
}

// RestartConversation restarts a conversation from a specific point
func (c *MistralClient) RestartConversation(conversationID string, inputs []ConversationInput) (*ConversationResponse, error) {
	reqMap := map[string]interface{}{
		"inputs": inputs,
	}

	response, err := c.request(http.MethodPost, reqMap, fmt.Sprintf("v1/conversations/%s/restart", conversationID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var convResponse ConversationResponse
	err = mapToStruct(respData, &convResponse)
	if err != nil {
		return nil, err
	}

	return &convResponse, nil
}
