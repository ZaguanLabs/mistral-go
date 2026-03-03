package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// BatchJobStatus represents the status of a batch job
type BatchJobStatus string

const (
	BatchJobStatusQueued     BatchJobStatus = "QUEUED"
	BatchJobStatusRunning    BatchJobStatus = "RUNNING"
	BatchJobStatusSuccess    BatchJobStatus = "SUCCESS"
	BatchJobStatusFailed     BatchJobStatus = "FAILED"
	BatchJobStatusTimedOut   BatchJobStatus = "TIMED_OUT"
	BatchJobStatusCancelled  BatchJobStatus = "CANCELLED"
	BatchJobStatusCancelling BatchJobStatus = "CANCELLING"
)

// BatchEndpoint represents the endpoint for batch processing
type BatchEndpoint string

const (
	BatchEndpointChat       BatchEndpoint = "/v1/chat/completions"
	BatchEndpointEmbeddings BatchEndpoint = "/v1/embeddings"
	BatchEndpointFIM        BatchEndpoint = "/v1/fim/completions"
	BatchEndpointModeration BatchEndpoint = "/v1/moderations"
)

// OrderBy represents sorting order for list endpoints.
type OrderBy string

const (
	OrderByCreatedAsc  OrderBy = "created"
	OrderByCreatedDesc OrderBy = "-created"
)

// BatchRequest represents an inline request to be executed in a batch job.
type BatchRequest struct {
	CustomID string         `json:"custom_id,omitempty"`
	Body     map[string]any `json:"body"`
}

// BatchJobMetadata represents metadata for a batch job
type BatchJobMetadata struct {
	TotalRequests     *int `json:"total_requests,omitempty"`
	SucceededRequests *int `json:"succeeded_requests,omitempty"`
	FailedRequests    *int `json:"failed_requests,omitempty"`
}

// BatchJobOut represents a batch job
type BatchJobOut struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	Endpoint     BatchEndpoint     `json:"endpoint"`
	InputFiles   []string          `json:"input_files"`
	OutputFile   *string           `json:"output_file,omitempty"`
	ErrorFile    *string           `json:"error_file,omitempty"`
	CreatedAt    int64             `json:"created_at"`
	StartedAt    *int64            `json:"started_at,omitempty"`
	CompletedAt  *int64            `json:"completed_at,omitempty"`
	Status       BatchJobStatus    `json:"status"`
	Model        *string           `json:"model,omitempty"`
	Metadata     *BatchJobMetadata `json:"metadata,omitempty"`
	TimeoutHours *int              `json:"timeout_hours,omitempty"`
}

// BatchJobsOut represents a list of batch jobs
type BatchJobsOut struct {
	Data   []BatchJobOut `json:"data"`
	Object string        `json:"object"`
	Total  int           `json:"total"`
}

// CreateBatchJobRequest represents the request to create a batch job
type CreateBatchJobRequest struct {
	InputFiles   []string       `json:"input_files"`
	Requests     []BatchRequest `json:"requests,omitempty"`
	Endpoint     BatchEndpoint  `json:"endpoint"`
	Model        *string        `json:"model,omitempty"`
	AgentID      *string        `json:"agent_id,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
	TimeoutHours *int           `json:"timeout_hours,omitempty"`
}

// ListBatchJobsParams represents parameters for listing batch jobs
type ListBatchJobsParams struct {
	Page         *int             `json:"page,omitempty"`
	PageSize     *int             `json:"page_size,omitempty"`
	Model        *string          `json:"model,omitempty"`
	AgentID      *string          `json:"agent_id,omitempty"`
	Metadata     map[string]any   `json:"metadata,omitempty"`
	CreatedAfter *time.Time       `json:"created_after,omitempty"`
	CreatedByMe  *bool            `json:"created_by_me,omitempty"`
	Status       []BatchJobStatus `json:"status,omitempty"`
	OrderBy      *OrderBy         `json:"order_by,omitempty"`
}

// CreateBatchJob creates a new batch job
func (c *MistralClient) CreateBatchJob(req *CreateBatchJobRequest) (*BatchJobOut, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	hasInputFiles := len(req.InputFiles) > 0
	hasRequests := len(req.Requests) > 0

	if hasInputFiles && hasRequests {
		return nil, fmt.Errorf("only one of input_files or requests should be provided, not both")
	}
	if !hasInputFiles && !hasRequests {
		return nil, fmt.Errorf("either input_files or requests must be provided")
	}

	payload := map[string]interface{}{
		"endpoint":      req.Endpoint,
		"model":         req.Model,
		"agent_id":      req.AgentID,
		"metadata":      req.Metadata,
		"timeout_hours": req.TimeoutHours,
	}
	if hasInputFiles {
		payload["input_files"] = req.InputFiles
	}
	if hasRequests {
		payload["requests"] = req.Requests
	}

	response, err := c.request(http.MethodPost, payload, "v1/batch/jobs", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var batchJobOut BatchJobOut
	err = mapToStruct(respData, &batchJobOut)
	if err != nil {
		return nil, err
	}

	return &batchJobOut, nil
}

// ListBatchJobs gets a list of batch jobs
func (c *MistralClient) ListBatchJobs(params *ListBatchJobsParams) (*BatchJobsOut, error) {
	if params == nil {
		params = &ListBatchJobsParams{}
	}

	// Build query parameters
	queryParams := url.Values{}
	if params.Page != nil {
		queryParams.Add("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%d", *params.PageSize))
	}
	if params.Model != nil {
		queryParams.Add("model", *params.Model)
	}
	if params.AgentID != nil {
		queryParams.Add("agent_id", *params.AgentID)
	}
	if params.CreatedAfter != nil {
		queryParams.Add("created_after", params.CreatedAfter.Format(time.RFC3339))
	}
	if params.CreatedByMe != nil {
		queryParams.Add("created_by_me", fmt.Sprintf("%t", *params.CreatedByMe))
	}
	for _, status := range params.Status {
		queryParams.Add("status", string(status))
	}
	if params.OrderBy != nil {
		queryParams.Add("order_by", string(*params.OrderBy))
	}

	path := "v1/batch/jobs"
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var batchJobsOut BatchJobsOut
	err = mapToStruct(respData, &batchJobsOut)
	if err != nil {
		return nil, err
	}

	return &batchJobsOut, nil
}

// GetBatchJob gets details of a specific batch job.
//
// Pass inline=true to request inline results in the response.
func (c *MistralClient) GetBatchJob(jobID string, inline ...bool) (*BatchJobOut, error) {
	path := fmt.Sprintf("v1/batch/jobs/%s", jobID)
	if len(inline) > 0 {
		path = fmt.Sprintf("%s?inline=%t", path, inline[0])
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var batchJobOut BatchJobOut
	err = mapToStruct(respData, &batchJobOut)
	if err != nil {
		return nil, err
	}

	return &batchJobOut, nil
}

// CancelBatchJob cancels a batch job
func (c *MistralClient) CancelBatchJob(jobID string) (*BatchJobOut, error) {
	response, err := c.request(http.MethodPost, nil, fmt.Sprintf("v1/batch/jobs/%s/cancel", jobID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var batchJobOut BatchJobOut
	err = mapToStruct(respData, &batchJobOut)
	if err != nil {
		return nil, err
	}

	return &batchJobOut, nil
}
