package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// JobStatus represents the status of a fine-tuning job
type JobStatus string

const (
	JobStatusQueued    JobStatus = "QUEUED"
	JobStatusStarted   JobStatus = "STARTED"
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusFailed    JobStatus = "FAILED"
	JobStatusSuccess   JobStatus = "SUCCESS"
	JobStatusCancelled JobStatus = "CANCELLED"
	JobStatusTimedOut  JobStatus = "TIMED_OUT"
)

// FineTuneableModelType represents the type of model that can be fine-tuned
type FineTuneableModelType string

const (
	FineTuneableModelTypeFineTuning FineTuneableModelType = "FT"
	FineTuneableModelTypeClassifier FineTuneableModelType = "CLASSIFIER"
)

// TrainingFile represents a training file with optional weight
type TrainingFile struct {
	FileID string   `json:"file_id"`
	Weight *float64 `json:"weight,omitempty"`
}

// Hyperparameters represents hyperparameters for fine-tuning
type Hyperparameters struct {
	TrainingSteps  *int     `json:"training_steps,omitempty"`
	LearningRate   *float64 `json:"learning_rate,omitempty"`
	WeightDecay    *float64 `json:"weight_decay,omitempty"`
	WarmupFraction *float64 `json:"warmup_fraction,omitempty"`
	Epochs         *float64 `json:"epochs,omitempty"`
	FimRatio       *float64 `json:"fim_ratio,omitempty"`
}

// WandbIntegration represents Weights & Biases integration
type WandbIntegration struct {
	Type    string  `json:"type"` // "wandb"
	Project string  `json:"project"`
	Name    *string `json:"name,omitempty"`
	APIKey  *string `json:"api_key,omitempty"`
}

// JobMetadata represents metadata for a fine-tuning job
type JobMetadata struct {
	ExpectedDurationSeconds *int     `json:"expected_duration_seconds,omitempty"`
	Cost                    *float64 `json:"cost,omitempty"`
	CostCurrency            *string  `json:"cost_currency,omitempty"`
	TrainTokensPerStep      *int     `json:"train_tokens_per_step,omitempty"`
	DataTokens              *int     `json:"data_tokens,omitempty"`
	EstimatedStartTime      *int64   `json:"estimated_start_time,omitempty"`
}

// JobOut represents a fine-tuning job
type JobOut struct {
	ID                          string                 `json:"id"`
	Hyperparameters             Hyperparameters        `json:"hyperparameters"`
	FineTunedModel              *string                `json:"fine_tuned_model,omitempty"`
	Model                       string                 `json:"model"`
	Status                      JobStatus              `json:"status"`
	JobType                     *FineTuneableModelType `json:"job_type,omitempty"`
	CreatedAt                   int64                  `json:"created_at"`
	ModifiedAt                  int64                  `json:"modified_at"`
	TrainingFiles               []string               `json:"training_files"`
	ValidationFiles             []string               `json:"validation_files,omitempty"`
	Object                      string                 `json:"object"`
	Integrations                []interface{}          `json:"integrations,omitempty"`
	TrainedTokens               *int                   `json:"trained_tokens,omitempty"`
	Suffix                      *string                `json:"suffix,omitempty"`
	Metadata                    *JobMetadata           `json:"metadata,omitempty"`
	InvalidSampleSkipPercentage *float64               `json:"invalid_sample_skip_percentage,omitempty"`
	AutoStart                   *bool                  `json:"auto_start,omitempty"`
}

// JobsOut represents a list of fine-tuning jobs
type JobsOut struct {
	Data   []JobOut `json:"data"`
	Object string   `json:"object"`
	Total  int      `json:"total"`
}

// CreateFineTuningJobRequest represents the request to create a fine-tuning job
type CreateFineTuningJobRequest struct {
	Model                       string                 `json:"model"`
	TrainingFiles               []TrainingFile         `json:"training_files,omitempty"`
	ValidationFiles             []string               `json:"validation_files,omitempty"`
	Hyperparameters             Hyperparameters        `json:"hyperparameters"`
	Suffix                      *string                `json:"suffix,omitempty"`
	Integrations                []interface{}          `json:"integrations,omitempty"`
	AutoStart                   *bool                  `json:"auto_start,omitempty"`
	InvalidSampleSkipPercentage *float64               `json:"invalid_sample_skip_percentage,omitempty"`
	JobType                     *FineTuneableModelType `json:"job_type,omitempty"`
}

// ListFineTuningJobsParams represents parameters for listing fine-tuning jobs
type ListFineTuningJobsParams struct {
	Page          *int       `json:"page,omitempty"`
	PageSize      *int       `json:"page_size,omitempty"`
	Model         *string    `json:"model,omitempty"`
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	CreatedByMe   *bool      `json:"created_by_me,omitempty"`
	Status        *JobStatus `json:"status,omitempty"`
	WandbProject  *string    `json:"wandb_project,omitempty"`
	WandbName     *string    `json:"wandb_name,omitempty"`
	Suffix        *string    `json:"suffix,omitempty"`
}

// CreateFineTuningJob creates a new fine-tuning job
func (c *MistralClient) CreateFineTuningJob(req *CreateFineTuningJobRequest) (*JobOut, error) {
	response, err := c.request(http.MethodPost, map[string]interface{}{
		"model":                          req.Model,
		"training_files":                 req.TrainingFiles,
		"validation_files":               req.ValidationFiles,
		"hyperparameters":                req.Hyperparameters,
		"suffix":                         req.Suffix,
		"integrations":                   req.Integrations,
		"auto_start":                     req.AutoStart,
		"invalid_sample_skip_percentage": req.InvalidSampleSkipPercentage,
		"job_type":                       req.JobType,
	}, "v1/fine_tuning/jobs", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var jobOut JobOut
	err = mapToStruct(respData, &jobOut)
	if err != nil {
		return nil, err
	}

	return &jobOut, nil
}

// ListFineTuningJobs gets a list of fine-tuning jobs
func (c *MistralClient) ListFineTuningJobs(params *ListFineTuningJobsParams) (*JobsOut, error) {
	if params == nil {
		params = &ListFineTuningJobsParams{}
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
	if params.CreatedAfter != nil {
		queryParams.Add("created_after", params.CreatedAfter.Format(time.RFC3339))
	}
	if params.CreatedBefore != nil {
		queryParams.Add("created_before", params.CreatedBefore.Format(time.RFC3339))
	}
	if params.CreatedByMe != nil {
		queryParams.Add("created_by_me", fmt.Sprintf("%t", *params.CreatedByMe))
	}
	if params.Status != nil {
		queryParams.Add("status", string(*params.Status))
	}
	if params.WandbProject != nil {
		queryParams.Add("wandb_project", *params.WandbProject)
	}
	if params.WandbName != nil {
		queryParams.Add("wandb_name", *params.WandbName)
	}
	if params.Suffix != nil {
		queryParams.Add("suffix", *params.Suffix)
	}

	path := "v1/fine_tuning/jobs"
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

	var jobsOut JobsOut
	err = mapToStruct(respData, &jobsOut)
	if err != nil {
		return nil, err
	}

	return &jobsOut, nil
}

// GetFineTuningJob gets details of a specific fine-tuning job
func (c *MistralClient) GetFineTuningJob(jobID string) (*JobOut, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/fine_tuning/jobs/%s", jobID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var jobOut JobOut
	err = mapToStruct(respData, &jobOut)
	if err != nil {
		return nil, err
	}

	return &jobOut, nil
}

// CancelFineTuningJob cancels a fine-tuning job
func (c *MistralClient) CancelFineTuningJob(jobID string) (*JobOut, error) {
	response, err := c.request(http.MethodPost, nil, fmt.Sprintf("v1/fine_tuning/jobs/%s/cancel", jobID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var jobOut JobOut
	err = mapToStruct(respData, &jobOut)
	if err != nil {
		return nil, err
	}

	return &jobOut, nil
}

// StartFineTuningJob starts a fine-tuning job
func (c *MistralClient) StartFineTuningJob(jobID string) (*JobOut, error) {
	response, err := c.request(http.MethodPost, nil, fmt.Sprintf("v1/fine_tuning/jobs/%s/start", jobID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var jobOut JobOut
	err = mapToStruct(respData, &jobOut)
	if err != nil {
		return nil, err
	}

	return &jobOut, nil
}
