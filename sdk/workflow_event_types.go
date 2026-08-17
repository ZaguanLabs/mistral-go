package sdk

// WorkflowEventResponse is the common wire envelope used by workflow,
// activity-task, custom-task, and workflow-task events.
type WorkflowEventResponse struct {
	EventID              string         `json:"event_id"`
	EventTimestamp       int            `json:"event_timestamp"`
	RootWorkflowExecID   string         `json:"root_workflow_exec_id"`
	ParentWorkflowExecID *string        `json:"parent_workflow_exec_id"`
	ContinuedRunID       *string        `json:"continued_run_id"`
	FirstExecutionRunID  *string        `json:"first_execution_run_id"`
	ScheduleID           *string        `json:"schedule_id"`
	WorkflowExecID       string         `json:"workflow_exec_id"`
	WorkflowRunID        string         `json:"workflow_run_id"`
	WorkflowName         string         `json:"workflow_name"`
	Attributes           map[string]any `json:"attributes"`
	EventType            string         `json:"event_type"`
}

type ActivityTaskCompletedResponse = WorkflowEventResponse
type ActivityTaskFailedResponse = WorkflowEventResponse
type ActivityTaskRetryingResponse = WorkflowEventResponse
type ActivityTaskStartedResponse = WorkflowEventResponse
type CustomTaskCanceledResponse = WorkflowEventResponse
type CustomTaskCompletedResponse = WorkflowEventResponse
type CustomTaskFailedResponse = WorkflowEventResponse
type CustomTaskInProgressResponse = WorkflowEventResponse
type CustomTaskStartedResponse = WorkflowEventResponse
type CustomTaskTimedOutResponse = WorkflowEventResponse
type WorkflowExecutionCanceledResponse = WorkflowEventResponse
type WorkflowExecutionCompletedResponse = WorkflowEventResponse
type WorkflowExecutionContinuedAsNewResponse = WorkflowEventResponse
type WorkflowExecutionFailedResponse = WorkflowEventResponse
type WorkflowExecutionStartedResponse = WorkflowEventResponse
type WorkflowTaskFailedResponse = WorkflowEventResponse
type WorkflowTaskTimedOutResponse = WorkflowEventResponse

type WorkflowExecutionCanceledAttributes struct {
	TaskID  string  `json:"task_id"`
	Reason  *string `json:"reason,omitempty"`
	Attempt *int    `json:"attempt,omitempty"`
}

type WorkflowExecutionCompletedAttributesResponse struct {
	TaskID  string `json:"task_id"`
	Result  any    `json:"result"`
	Attempt *int   `json:"attempt,omitempty"`
}

type WorkflowExecutionFailedAttributes struct {
	TaskID  string `json:"task_id"`
	Failure any    `json:"failure"`
	Attempt *int   `json:"attempt,omitempty"`
}

type WorkflowExecutionStartedAttributesResponse struct {
	TaskID       string  `json:"task_id"`
	WorkflowName string  `json:"workflow_name"`
	Input        any     `json:"input"`
	DisplayName  *string `json:"display_name,omitempty"`
	Attempt      *int    `json:"attempt,omitempty"`
}
