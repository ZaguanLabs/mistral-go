package sdk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ConversationInput represents input for a conversation
type ConversationInput struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// ConversationStartRequest represents a request to start a conversation
type ConversationStartRequest struct {
	Inputs           []ConversationInput    `json:"inputs"`
	Stream           *bool                  `json:"stream,omitempty"`
	Store            *bool                  `json:"store,omitempty"`
	HandoffExecution *string                `json:"handoff_execution,omitempty"`
	Instructions     *string                `json:"instructions,omitempty"`
	Tools            []Tool                 `json:"tools,omitempty"`
	CompletionArgs   map[string]interface{} `json:"completion_args,omitempty"`
	Name             *string                `json:"name,omitempty"`
	Description      *string                `json:"description,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	AgentID          *string                `json:"agent_id,omitempty"`
	AgentVersion     interface{}            `json:"agent_version,omitempty"`
	Model            *string                `json:"model,omitempty"`
}

// ConversationAppendRequest represents request params for append.
type ConversationAppendRequest struct {
	Inputs           []ConversationInput    `json:"inputs"`
	Stream           *bool                  `json:"stream,omitempty"`
	Store            *bool                  `json:"store,omitempty"`
	HandoffExecution *string                `json:"handoff_execution,omitempty"`
	CompletionArgs   map[string]interface{} `json:"completion_args,omitempty"`
}

// ConversationRestartRequest represents request params for restart.
type ConversationRestartRequest struct {
	Inputs           []ConversationInput    `json:"inputs"`
	FromEntryID      *string                `json:"from_entry_id,omitempty"`
	Stream           *bool                  `json:"stream,omitempty"`
	Store            *bool                  `json:"store,omitempty"`
	HandoffExecution *string                `json:"handoff_execution,omitempty"`
	CompletionArgs   map[string]interface{} `json:"completion_args,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	AgentVersion     interface{}            `json:"agent_version,omitempty"`
}

// ListConversationsParams represents list filters.
type ListConversationsParams struct {
	Page     *int                   `json:"page,omitempty"`
	PageSize *int                   `json:"page_size,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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

// ConversationMessagesResponse represents messages in a conversation.
type ConversationMessagesResponse struct {
	ConversationID string        `json:"conversation_id,omitempty"`
	Messages       []ChatMessage `json:"messages,omitempty"`
	Data           []ChatMessage `json:"data,omitempty"`
}

// ConversationStreamEvent represents one SSE conversation event.
type ConversationStreamEvent struct {
	Type  string                 `json:"type,omitempty"`
	Data  map[string]interface{} `json:"data,omitempty"`
	Error error                  `json:"-"`
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
	if req.Stream != nil {
		reqMap["stream"] = *req.Stream
	}
	if req.Store != nil {
		reqMap["store"] = *req.Store
	}
	if req.HandoffExecution != nil {
		reqMap["handoff_execution"] = *req.HandoffExecution
	}
	if req.Metadata != nil {
		reqMap["metadata"] = req.Metadata
	}
	if req.AgentID != nil {
		reqMap["agent_id"] = *req.AgentID
	}
	if req.AgentVersion != nil {
		reqMap["agent_version"] = req.AgentVersion
	}
	if req.Model != nil {
		reqMap["model"] = *req.Model
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

// StartConversationStream starts a conversation and returns SSE events.
func (c *MistralClient) StartConversationStream(req *ConversationStartRequest) (<-chan ConversationStreamEvent, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	streamTrue := true
	req.Stream = &streamTrue

	reqMap := map[string]interface{}{"inputs": req.Inputs, "stream": true}
	if req.Store != nil {
		reqMap["store"] = *req.Store
	}
	if req.HandoffExecution != nil {
		reqMap["handoff_execution"] = *req.HandoffExecution
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
	if req.Metadata != nil {
		reqMap["metadata"] = req.Metadata
	}
	if req.AgentID != nil {
		reqMap["agent_id"] = *req.AgentID
	}
	if req.AgentVersion != nil {
		reqMap["agent_version"] = req.AgentVersion
	}
	if req.Model != nil {
		reqMap["model"] = *req.Model
	}

	return c.conversationStream("v1/conversations", reqMap)
}

// AppendToConversationStream appends to a conversation and returns SSE events.
func (c *MistralClient) AppendToConversationStream(conversationID string, req *ConversationAppendRequest) (<-chan ConversationStreamEvent, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	reqMap := map[string]interface{}{"inputs": req.Inputs, "stream": true}
	if req.Store != nil {
		reqMap["store"] = *req.Store
	}
	if req.HandoffExecution != nil {
		reqMap["handoff_execution"] = *req.HandoffExecution
	}
	if req.CompletionArgs != nil {
		reqMap["completion_args"] = req.CompletionArgs
	}

	return c.conversationStream(fmt.Sprintf("v1/conversations/%s", conversationID), reqMap)
}

// RestartConversationStream restarts a conversation and returns SSE events.
func (c *MistralClient) RestartConversationStream(conversationID string, req *ConversationRestartRequest) (<-chan ConversationStreamEvent, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	reqMap := map[string]interface{}{"inputs": req.Inputs, "stream": true}
	if req.FromEntryID != nil {
		reqMap["from_entry_id"] = *req.FromEntryID
	}
	if req.Store != nil {
		reqMap["store"] = *req.Store
	}
	if req.HandoffExecution != nil {
		reqMap["handoff_execution"] = *req.HandoffExecution
	}
	if req.CompletionArgs != nil {
		reqMap["completion_args"] = req.CompletionArgs
	}
	if req.Metadata != nil {
		reqMap["metadata"] = req.Metadata
	}
	if req.AgentVersion != nil {
		reqMap["agent_version"] = req.AgentVersion
	}

	return c.conversationStream(fmt.Sprintf("v1/conversations/%s/restart", conversationID), reqMap)
}

func (c *MistralClient) conversationStream(path string, reqMap map[string]interface{}) (<-chan ConversationStreamEvent, error) {
	response, err := c.request(http.MethodPost, reqMap, path, true, nil)
	if err != nil {
		return nil, err
	}

	body, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	out := make(chan ConversationStreamEvent)
	go func() {
		defer close(out)
		defer body.Close()

		reader := bufio.NewReader(body)
		for {
			line, readErr := reader.ReadBytes('\n')
			if readErr == io.EOF {
				break
			}
			if readErr != nil {
				out <- ConversationStreamEvent{Error: fmt.Errorf("error reading stream response: %w", readErr)}
				return
			}

			if bytes.Equal(line, []byte("\n")) {
				continue
			}
			if !bytes.HasPrefix(line, []byte("data: ")) {
				continue
			}

			jsonLine := bytes.TrimSpace(bytes.TrimPrefix(line, []byte("data: ")))
			if bytes.Equal(jsonLine, []byte("[DONE]")) {
				break
			}

			var payload map[string]interface{}
			if err := json.Unmarshal(jsonLine, &payload); err != nil {
				out <- ConversationStreamEvent{Error: fmt.Errorf("error decoding stream event: %w", err)}
				continue
			}

			event := ConversationStreamEvent{Data: payload}
			if t, ok := payload["type"].(string); ok {
				event.Type = t
			}
			out <- event
		}
	}()

	return out, nil
}

// ListConversations lists all conversations
func (c *MistralClient) ListConversations(page int) (*ConversationListResponse, error) {
	return c.ListConversationsWithParams(&ListConversationsParams{Page: &page})
}

// ListConversationsWithParams lists all conversations with filters.
func (c *MistralClient) ListConversationsWithParams(params *ListConversationsParams) (*ConversationListResponse, error) {
	if params == nil {
		params = &ListConversationsParams{}
	}

	query := url.Values{}
	if params.Page != nil {
		query.Add("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.PageSize != nil {
		query.Add("page_size", fmt.Sprintf("%d", *params.PageSize))
	}
	for k, v := range params.Metadata {
		query.Add("metadata", fmt.Sprintf("%s:%v", k, v))
	}

	path := "v1/conversations"
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
	return c.AppendToConversationWithParams(conversationID, &ConversationAppendRequest{Inputs: inputs})
}

// AppendToConversationWithParams appends inputs with optional params.
func (c *MistralClient) AppendToConversationWithParams(conversationID string, req *ConversationAppendRequest) (*ConversationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"inputs": req.Inputs,
	}
	if req.Stream != nil {
		reqMap["stream"] = *req.Stream
	}
	if req.Store != nil {
		reqMap["store"] = *req.Store
	}
	if req.HandoffExecution != nil {
		reqMap["handoff_execution"] = *req.HandoffExecution
	}
	if req.CompletionArgs != nil {
		reqMap["completion_args"] = req.CompletionArgs
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

// GetConversationMessages retrieves only message entries from a conversation.
func (c *MistralClient) GetConversationMessages(conversationID string) (*ConversationMessagesResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/conversations/%s/messages", conversationID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var messagesResponse ConversationMessagesResponse
	err = mapToStruct(respData, &messagesResponse)
	if err != nil {
		return nil, err
	}

	return &messagesResponse, nil
}

// RestartConversation restarts a conversation from a specific point
func (c *MistralClient) RestartConversation(conversationID string, inputs []ConversationInput) (*ConversationResponse, error) {
	return c.RestartConversationWithParams(conversationID, &ConversationRestartRequest{Inputs: inputs})
}

// RestartConversationWithParams restarts a conversation with optional params.
func (c *MistralClient) RestartConversationWithParams(conversationID string, req *ConversationRestartRequest) (*ConversationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := map[string]interface{}{
		"inputs": req.Inputs,
	}
	if req.FromEntryID != nil {
		reqMap["from_entry_id"] = *req.FromEntryID
	}
	if req.Stream != nil {
		reqMap["stream"] = *req.Stream
	}
	if req.Store != nil {
		reqMap["store"] = *req.Store
	}
	if req.HandoffExecution != nil {
		reqMap["handoff_execution"] = *req.HandoffExecution
	}
	if req.CompletionArgs != nil {
		reqMap["completion_args"] = req.CompletionArgs
	}
	if req.Metadata != nil {
		reqMap["metadata"] = req.Metadata
	}
	if req.AgentVersion != nil {
		reqMap["agent_version"] = req.AgentVersion
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

// DeleteConversation deletes a conversation
//
// Parameters:
//   - conversationID: The ID of the conversation to delete
//
// Returns an error if the deletion fails
func (c *MistralClient) DeleteConversation(conversationID string) error {
	_, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/conversations/%s", conversationID), false, nil)
	return err
}
