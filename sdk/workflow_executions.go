package sdk

import (
	"fmt"
	"io"
	"net/http"
)

type WorkflowSignalRequest struct {
	Name  string `json:"name"`
	Input any    `json:"input,omitempty"`
}

type WorkflowQueryRequest struct {
	Name  string `json:"name"`
	Input any    `json:"input,omitempty"`
}

type WorkflowUpdateRequest struct {
	Name  string `json:"name"`
	Input any    `json:"input,omitempty"`
}

type BatchWorkflowExecutionRequest struct {
	ExecutionIDs []string `json:"execution_ids"`
}

type ResetWorkflowRequest struct {
	EventID        string  `json:"event_id"`
	Reason         *string `json:"reason,omitempty"`
	ExcludeSignals *bool   `json:"exclude_signals,omitempty"`
	ExcludeUpdates *bool   `json:"exclude_updates,omitempty"`
}

type WorkflowTraceEventsParams struct {
	MergeSameIDEvents     *bool `json:"merge_same_id_events,omitempty"`
	IncludeInternalEvents *bool `json:"include_internal_events,omitempty"`
}

type WorkflowExecutionStreamParams struct {
	EventSource *string `json:"event_source,omitempty"`
	LastEventID *string `json:"last_event_id,omitempty"`
}

func (c *MistralClient) GetWorkflowExecution(executionID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/executions/%s", executionID))
}

func (c *MistralClient) GetWorkflowExecutionHistory(executionID string, decodePayloads *bool) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"decode_payloads": decodePayloads})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/workflows/executions/%s/history", executionID), query))
}

func (c *MistralClient) SignalWorkflowExecution(executionID string, req *WorkflowSignalRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"name": req.Name, "input": req.Input})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/executions/%s/signals", executionID))
}

func (c *MistralClient) QueryWorkflowExecution(executionID string, req *WorkflowQueryRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"name": req.Name, "input": req.Input})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/executions/%s/queries", executionID))
}

func (c *MistralClient) TerminateWorkflowExecution(executionID string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, nil, fmt.Sprintf("v1/workflows/executions/%s/terminate", executionID))
}

func (c *MistralClient) BatchTerminateWorkflowExecutions(executionIDs []string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"execution_ids": executionIDs}, "v1/workflows/executions/terminate")
}

func (c *MistralClient) CancelWorkflowExecution(executionID string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, nil, fmt.Sprintf("v1/workflows/executions/%s/cancel", executionID))
}

func (c *MistralClient) BatchCancelWorkflowExecutions(executionIDs []string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"execution_ids": executionIDs}, "v1/workflows/executions/cancel")
}

func (c *MistralClient) ResetWorkflow(executionID string, req *ResetWorkflowRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"event_id":        req.EventID,
		"reason":          req.Reason,
		"exclude_signals": req.ExcludeSignals,
		"exclude_updates": req.ExcludeUpdates,
	})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/executions/%s/reset", executionID))
}

func (c *MistralClient) UpdateWorkflowExecution(executionID string, req *WorkflowUpdateRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"name": req.Name, "input": req.Input})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/executions/%s/updates", executionID))
}

func (c *MistralClient) GetWorkflowExecutionTraceOTEL(executionID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/executions/%s/trace/otel", executionID))
}

func (c *MistralClient) GetWorkflowExecutionTraceSummary(executionID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/executions/%s/trace/summary", executionID))
}

func (c *MistralClient) GetWorkflowExecutionTraceEvents(executionID string, params *WorkflowTraceEventsParams) (APIResponse, error) {
	if params == nil {
		params = &WorkflowTraceEventsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"merge_same_id_events":    params.MergeSameIDEvents,
		"include_internal_events": params.IncludeInternalEvents,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/workflows/executions/%s/trace/events", executionID), query))
}

func (c *MistralClient) StreamWorkflowExecution(executionID string, params *WorkflowExecutionStreamParams) (<-chan StreamEvent, error) {
	if params == nil {
		params = &WorkflowExecutionStreamParams{}
	}
	query := queryWithOptionalValues(map[string]any{"event_source": params.EventSource, "last_event_id": params.LastEventID})
	response, err := c.request(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/workflows/executions/%s/stream", executionID), query), true, nil)
	if err != nil {
		return nil, err
	}
	body, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	return parseGenericStream(body), nil
}
