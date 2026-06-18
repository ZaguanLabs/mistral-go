package sdk

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type ListWorkflowsParams struct {
	Status                   any               `json:"status,omitempty"`
	ActiveOnly               *bool             `json:"active_only,omitempty"`
	IncludeShared            *bool             `json:"include_shared,omitempty"`
	AvailableInChatAssistant *bool             `json:"available_in_chat_assistant,omitempty"`
	DeploymentName           []string          `json:"deployment_name,omitempty"`
	DeploymentStatus         *DeploymentStatus `json:"deployment_status,omitempty"`
	Archived                 *bool             `json:"archived,omitempty"`
	Tags                     []string          `json:"tags,omitempty"`
	SortBy                   *string           `json:"sort_by,omitempty"`
	Order                    *Order            `json:"order,omitempty"`
	Cursor                   *string           `json:"cursor,omitempty"`
	Limit                    *int              `json:"limit,omitempty"`
}

type ListWorkflowRegistrationsParams struct {
	WorkflowID               *string `json:"workflow_id,omitempty"`
	TaskQueue                *string `json:"task_queue,omitempty"`
	ActiveOnly               *bool   `json:"active_only,omitempty"`
	IncludeShared            *bool   `json:"include_shared,omitempty"`
	WorkflowSearch           *string `json:"workflow_search,omitempty"`
	Archived                 *bool   `json:"archived,omitempty"`
	WithWorkflow             *bool   `json:"with_workflow,omitempty"`
	AvailableInChatAssistant *bool   `json:"available_in_chat_assistant,omitempty"`
	Limit                    *int    `json:"limit,omitempty"`
	Cursor                   *string `json:"cursor,omitempty"`
}

type ExecuteWorkflowRequest struct {
	ExecutionID             *string        `json:"execution_id,omitempty"`
	Input                   any            `json:"input,omitempty"`
	WaitForResult           *bool          `json:"wait_for_result,omitempty"`
	TimeoutSeconds          *int           `json:"timeout_seconds,omitempty"`
	CustomTracingAttributes map[string]any `json:"custom_tracing_attributes,omitempty"`
	Extensions              map[string]any `json:"extensions,omitempty"`
	TaskQueue               *string        `json:"task_queue,omitempty"`
	DeploymentName          *string        `json:"deployment_name,omitempty"`
}

type UpdateWorkflowRequest struct {
	DisplayName              *string `json:"display_name,omitempty"`
	Description              *string `json:"description,omitempty"`
	AvailableInChatAssistant *bool   `json:"available_in_chat_assistant,omitempty"`
}

type GetWorkflowRegistrationParams struct {
	WithWorkflow  *bool `json:"with_workflow,omitempty"`
	IncludeShared *bool `json:"include_shared,omitempty"`
}

type ListWorkflowRunsParams struct {
	WorkflowIdentifier *string            `json:"workflow_identifier,omitempty"`
	Search             *string            `json:"search,omitempty"`
	Status             any                `json:"status,omitempty"`
	DeploymentName     *string            `json:"deployment_name,omitempty"`
	SortBy             *WorkflowRunSortBy `json:"sort_by,omitempty"`
	Order              *Order             `json:"order,omitempty"`
	StartTimeAfter     *time.Time         `json:"start_time_after,omitempty"`
	StartTimeBefore    *time.Time         `json:"start_time_before,omitempty"`
	EndTimeAfter       *time.Time         `json:"end_time_after,omitempty"`
	EndTimeBefore      *time.Time         `json:"end_time_before,omitempty"`
	UserID             *string            `json:"user_id,omitempty"`
	PageSize           *int               `json:"page_size,omitempty"`
	NextPageToken      *string            `json:"next_page_token,omitempty"`
}

type ListWorkflowEventsParams struct {
	WorkflowName  *string `json:"workflow_name,omitempty"`
	RunID         *string `json:"run_id,omitempty"`
	ExecutionID   *string `json:"execution_id,omitempty"`
	PageSize      *int    `json:"page_size,omitempty"`
	NextPageToken *string `json:"next_page_token,omitempty"`
}

type ListWorkflowSchedulesParams struct {
	WorkflowName  *string `json:"workflow_name,omitempty"`
	UserID        *string `json:"user_id,omitempty"`
	Status        *string `json:"status,omitempty"`
	PageSize      *int    `json:"page_size,omitempty"`
	NextPageToken *string `json:"next_page_token,omitempty"`
}

type ScheduleWorkflowRequest struct {
	Schedule               any     `json:"schedule,omitempty"`
	WorkflowRegistrationID *string `json:"workflow_registration_id,omitempty"`
	WorkflowVersionID      *string `json:"workflow_version_id,omitempty"`
	WorkflowIdentifier     *string `json:"workflow_identifier,omitempty"`
	WorkflowTaskQueue      *string `json:"workflow_task_queue,omitempty"`
	ScheduleID             *string `json:"schedule_id,omitempty"`
	DeploymentName         *string `json:"deployment_name,omitempty"`
}

type UpdateScheduleRequest struct {
	Schedule any `json:"schedule"`
}

type ScheduleNoteRequest struct {
	Note *string `json:"note,omitempty"`
}

type TriggerScheduleRequest struct {
	Overlap *string `json:"overlap,omitempty"`
}

type ExecuteWorkflowAndWaitParams struct {
	WorkflowIdentifier      string
	Input                   any
	ExecutionID             *string
	DeploymentName          *string
	CustomTracingAttributes map[string]any
	TaskQueue               *string
	PollingInterval         time.Duration
	MaxAttempts             *int
	UseAPISync              bool
	TimeoutSeconds          *int
}

func (c *MistralClient) GetWorkflows(params *ListWorkflowsParams) (APIResponse, error) {
	if params == nil {
		params = &ListWorkflowsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"status":                      params.Status,
		"active_only":                 params.ActiveOnly,
		"include_shared":              params.IncludeShared,
		"available_in_chat_assistant": params.AvailableInChatAssistant,
		"deployment_name":             params.DeploymentName,
		"deployment_status":           params.DeploymentStatus,
		"archived":                    params.Archived,
		"tags":                        params.Tags,
		"sort_by":                     params.SortBy,
		"order":                       params.Order,
		"cursor":                      params.Cursor,
		"limit":                       params.Limit,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery("v1/workflows", query))
}

func (c *MistralClient) GetWorkflowRegistrations(params *ListWorkflowRegistrationsParams) (APIResponse, error) {
	if params == nil {
		params = &ListWorkflowRegistrationsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"workflow_id":                 params.WorkflowID,
		"task_queue":                  params.TaskQueue,
		"active_only":                 params.ActiveOnly,
		"include_shared":              params.IncludeShared,
		"workflow_search":             params.WorkflowSearch,
		"archived":                    params.Archived,
		"with_workflow":               params.WithWorkflow,
		"available_in_chat_assistant": params.AvailableInChatAssistant,
		"limit":                       params.Limit,
		"cursor":                      params.Cursor,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery("v1/workflows/registrations", query))
}

func (c *MistralClient) ExecuteWorkflow(workflowIdentifier string, req *ExecuteWorkflowRequest) (APIResponse, error) {
	return c.executeWorkflowPath(fmt.Sprintf("v1/workflows/%s/execute", workflowIdentifier), req)
}

func (c *MistralClient) ExecuteWorkflowRegistration(workflowRegistrationID string, req *ExecuteWorkflowRequest) (APIResponse, error) {
	return c.executeWorkflowPath(fmt.Sprintf("v1/workflows/registrations/%s/execute", workflowRegistrationID), req)
}

func (c *MistralClient) ExecuteWorkflowAndWait(params *ExecuteWorkflowAndWaitParams) (any, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	req := &ExecuteWorkflowRequest{
		ExecutionID:             params.ExecutionID,
		Input:                   params.Input,
		TimeoutSeconds:          params.TimeoutSeconds,
		CustomTracingAttributes: params.CustomTracingAttributes,
		TaskQueue:               params.TaskQueue,
		DeploymentName:          params.DeploymentName,
	}
	if params.UseAPISync {
		wait := true
		req.WaitForResult = &wait
		response, err := c.ExecuteWorkflow(params.WorkflowIdentifier, req)
		if err != nil {
			return nil, err
		}
		return response["result"], nil
	}
	response, err := c.ExecuteWorkflow(params.WorkflowIdentifier, req)
	if err != nil {
		return nil, err
	}
	executionID, ok := response["execution_id"].(string)
	if !ok || executionID == "" {
		return nil, fmt.Errorf("workflow execution response missing execution_id")
	}
	finalExecution, err := c.WaitForWorkflowCompletion(executionID, params.PollingInterval, params.MaxAttempts)
	if err != nil {
		return nil, err
	}
	return finalExecution["result"], nil
}

func (c *MistralClient) WaitForWorkflowCompletion(executionID string, pollingInterval time.Duration, maxAttempts *int) (APIResponse, error) {
	if pollingInterval == 0 {
		pollingInterval = 5 * time.Second
	}
	attempts := 0
	for {
		response, err := c.GetWorkflowExecution(executionID)
		if err != nil {
			return nil, err
		}
		status, _ := response["status"].(string)
		if status != "RUNNING" {
			if status == "COMPLETED" {
				return response, nil
			}
			return nil, fmt.Errorf("workflow failed with status: %s", status)
		}
		attempts++
		if maxAttempts != nil && attempts >= *maxAttempts {
			return nil, fmt.Errorf("workflow is still running after %d polling attempts", *maxAttempts)
		}
		time.Sleep(pollingInterval)
	}
}

func (c *MistralClient) GetWorkflow(workflowIdentifier string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/%s", workflowIdentifier))
}

func (c *MistralClient) UpdateWorkflow(workflowIdentifier string, req *UpdateWorkflowRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"display_name":                req.DisplayName,
		"description":                 req.Description,
		"available_in_chat_assistant": req.AvailableInChatAssistant,
	})
	return c.requestMap(http.MethodPut, body, fmt.Sprintf("v1/workflows/%s", workflowIdentifier))
}

func (c *MistralClient) GetWorkflowRegistration(workflowRegistrationID string, params *GetWorkflowRegistrationParams) (APIResponse, error) {
	if params == nil {
		params = &GetWorkflowRegistrationParams{}
	}
	query := queryWithOptionalValues(map[string]any{"with_workflow": params.WithWorkflow, "include_shared": params.IncludeShared})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/workflows/registrations/%s", workflowRegistrationID), query))
}

func (c *MistralClient) ArchiveWorkflow(workflowIdentifier string) (APIResponse, error) {
	return c.requestMap(http.MethodPut, nil, fmt.Sprintf("v1/workflows/%s/archive", workflowIdentifier))
}

func (c *MistralClient) UnarchiveWorkflow(workflowIdentifier string) (APIResponse, error) {
	return c.requestMap(http.MethodPut, nil, fmt.Sprintf("v1/workflows/%s/unarchive", workflowIdentifier))
}

func (c *MistralClient) BulkArchiveWorkflows(workflowIDs []string) (APIResponse, error) {
	body := map[string]interface{}{"workflow_ids": workflowIDs}
	return c.requestMap(http.MethodPut, body, "v1/workflows/archive")
}

func (c *MistralClient) BulkUnarchiveWorkflows(workflowIDs []string) (APIResponse, error) {
	body := map[string]interface{}{"workflow_ids": workflowIDs}
	return c.requestMap(http.MethodPut, body, "v1/workflows/unarchive")
}

func (c *MistralClient) ListWorkflowDeployments() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/workflows/deployments")
}

func (c *MistralClient) GetWorkflowDeployment(name string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/deployments/%s", name))
}

func (c *MistralClient) GetWorkflowMetrics(workflowName string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/%s/metrics", workflowName))
}

func (c *MistralClient) ListWorkflowRuns(params *ListWorkflowRunsParams) (APIResponse, error) {
	if params == nil {
		params = &ListWorkflowRunsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"workflow_identifier": params.WorkflowIdentifier,
		"search":              params.Search,
		"status":              params.Status,
		"deployment_name":     params.DeploymentName,
		"sort_by":             params.SortBy,
		"order":               params.Order,
		"start_time_after":    params.StartTimeAfter,
		"start_time_before":   params.StartTimeBefore,
		"end_time_after":      params.EndTimeAfter,
		"end_time_before":     params.EndTimeBefore,
		"user_id":             params.UserID,
		"page_size":           params.PageSize,
		"next_page_token":     params.NextPageToken,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery("v1/workflows/runs", query))
}

func (c *MistralClient) GetWorkflowRun(runID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/runs/%s", runID))
}

func (c *MistralClient) GetWorkflowRunHistory(runID string, decodePayloads *bool) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"decode_payloads": decodePayloads})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/workflows/runs/%s/history", runID), query))
}

func (c *MistralClient) GetWorkflowStreamEvents(params *ListWorkflowEventsParams) (<-chan StreamEvent, error) {
	if params == nil {
		params = &ListWorkflowEventsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"workflow_name":   params.WorkflowName,
		"run_id":          params.RunID,
		"execution_id":    params.ExecutionID,
		"page_size":       params.PageSize,
		"next_page_token": params.NextPageToken,
	})
	response, err := c.request(http.MethodGet, nil, appendQuery("v1/workflows/events/stream", query), true, nil)
	if err != nil {
		return nil, err
	}
	body, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	return parseGenericStream(body), nil
}

func (c *MistralClient) GetWorkflowEvents(params *ListWorkflowEventsParams) (APIResponse, error) {
	if params == nil {
		params = &ListWorkflowEventsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"workflow_name":   params.WorkflowName,
		"run_id":          params.RunID,
		"execution_id":    params.ExecutionID,
		"page_size":       params.PageSize,
		"next_page_token": params.NextPageToken,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery("v1/workflows/events/list", query))
}

func (c *MistralClient) GetWorkflowSchedules(params ...*ListWorkflowSchedulesParams) (APIResponse, error) {
	requestParams := &ListWorkflowSchedulesParams{}
	if len(params) > 0 && params[0] != nil {
		requestParams = params[0]
	}
	query := queryWithOptionalValues(map[string]any{
		"workflow_name":   requestParams.WorkflowName,
		"user_id":         requestParams.UserID,
		"status":          requestParams.Status,
		"page_size":       requestParams.PageSize,
		"next_page_token": requestParams.NextPageToken,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery("v1/workflows/schedules", query))
}

func (c *MistralClient) ScheduleWorkflow(req *ScheduleWorkflowRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"schedule":                 req.Schedule,
		"workflow_registration_id": req.WorkflowRegistrationID,
		"workflow_version_id":      req.WorkflowVersionID,
		"workflow_identifier":      req.WorkflowIdentifier,
		"workflow_task_queue":      req.WorkflowTaskQueue,
		"schedule_id":              req.ScheduleID,
		"deployment_name":          req.DeploymentName,
	})
	return c.requestMap(http.MethodPost, body, "v1/workflows/schedules")
}

func (c *MistralClient) UnscheduleWorkflow(scheduleID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/workflows/schedules/%s", scheduleID))
}

func (c *MistralClient) GetWorkflowSchedule(scheduleID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/workflows/schedules/%s", scheduleID))
}

func (c *MistralClient) UpdateWorkflowSchedule(scheduleID string, req *UpdateScheduleRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"schedule": req.Schedule})
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v1/workflows/schedules/%s", scheduleID))
}

func (c *MistralClient) PauseSchedule(scheduleID string, req *ScheduleNoteRequest) (APIResponse, error) {
	body := map[string]interface{}{}
	if req != nil {
		body = optionalRequestMap(map[string]any{"note": req.Note})
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/schedules/%s/pause", scheduleID))
}

func (c *MistralClient) ResumeSchedule(scheduleID string, req *ScheduleNoteRequest) (APIResponse, error) {
	body := map[string]interface{}{}
	if req != nil {
		body = optionalRequestMap(map[string]any{"note": req.Note})
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/schedules/%s/resume", scheduleID))
}

func (c *MistralClient) TriggerSchedule(scheduleID string, req *TriggerScheduleRequest) (APIResponse, error) {
	body := map[string]interface{}{}
	if req != nil {
		body = optionalRequestMap(map[string]any{"overlap": req.Overlap})
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/workflows/schedules/%s/trigger", scheduleID))
}

func (c *MistralClient) executeWorkflowPath(path string, req *ExecuteWorkflowRequest) (APIResponse, error) {
	if req == nil {
		req = &ExecuteWorkflowRequest{}
	}
	body := optionalRequestMap(map[string]any{
		"execution_id":              req.ExecutionID,
		"input":                     req.Input,
		"wait_for_result":           req.WaitForResult,
		"timeout_seconds":           req.TimeoutSeconds,
		"custom_tracing_attributes": req.CustomTracingAttributes,
		"extensions":                req.Extensions,
		"task_queue":                req.TaskQueue,
		"deployment_name":           req.DeploymentName,
	})
	return c.requestMap(http.MethodPost, body, path)
}
