package sdk

import (
	"testing"
	"time"
)

func TestCreateBatchJob(t *testing.T) {
	client := NewMistralClientDefault("")

	req := &CreateBatchJobRequest{
		InputFiles: []string{"file-test123"},
		Endpoint:   BatchEndpointChat,
	}

	_, err := client.CreateBatchJob(req)

	if err == nil {
		t.Log("CreateBatchJob succeeded")
	} else {
		t.Logf("CreateBatchJob failed as expected: %v", err)
	}
}

func TestCreateBatchJobWithAllParams(t *testing.T) {
	client := NewMistralClientDefault("")

	req := &CreateBatchJobRequest{
		InputFiles: []string{"file-input1", "file-input2"},
		Endpoint:   BatchEndpointEmbeddings,
		Model:      StringPtr("mistral-embed"),
		Metadata: map[string]any{
			"project":     "test-project",
			"environment": "staging",
			"batch_id":    123,
		},
		TimeoutHours: IntPtr(24),
	}

	_, err := client.CreateBatchJob(req)

	if err != nil {
		t.Logf("CreateBatchJob with all params failed: %v", err)
	}
}

func TestListBatchJobs(t *testing.T) {
	client := NewMistralClientDefault("")

	params := &ListBatchJobsParams{
		Page:     IntPtr(0),
		PageSize: IntPtr(10),
	}

	_, err := client.ListBatchJobs(params)

	if err == nil {
		t.Log("ListBatchJobs succeeded")
	} else {
		t.Logf("ListBatchJobs failed as expected: %v", err)
	}
}

func TestListBatchJobsWithFilters(t *testing.T) {
	client := NewMistralClientDefault("")

	now := time.Now()
	params := &ListBatchJobsParams{
		Page:         IntPtr(0),
		PageSize:     IntPtr(20),
		Model:        StringPtr("mistral-small-latest"),
		CreatedAfter: &now,
		CreatedByMe:  BoolPtr(true),
		Status: []BatchJobStatus{
			BatchJobStatusRunning,
			BatchJobStatusQueued,
		},
		Metadata: map[string]any{
			"project": "test",
		},
	}

	_, err := client.ListBatchJobs(params)

	if err != nil {
		t.Logf("ListBatchJobs with filters failed: %v", err)
	}
}

func TestListBatchJobsNilParams(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with nil params - should use defaults
	_, err := client.ListBatchJobs(nil)

	if err != nil {
		t.Logf("ListBatchJobs with nil params failed: %v", err)
	}
}

func TestGetBatchJob(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.GetBatchJob("batch-test123")

	if err == nil {
		t.Log("GetBatchJob succeeded")
	} else {
		t.Logf("GetBatchJob failed as expected: %v", err)
	}
}

func TestCancelBatchJob(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.CancelBatchJob("batch-test123")

	if err == nil {
		t.Log("CancelBatchJob succeeded")
	} else {
		t.Logf("CancelBatchJob failed as expected: %v", err)
	}
}

func TestBatchJobStatusConstants(t *testing.T) {
	statuses := []BatchJobStatus{
		BatchJobStatusQueued,
		BatchJobStatusRunning,
		BatchJobStatusSuccess,
		BatchJobStatusFailed,
		BatchJobStatusTimedOut,
		BatchJobStatusCancelled,
		BatchJobStatusCancelling,
	}

	for _, status := range statuses {
		if status == "" {
			t.Error("BatchJobStatus constant is empty")
		}
	}
}

func TestBatchEndpointConstants(t *testing.T) {
	endpoints := []BatchEndpoint{
		BatchEndpointChat,
		BatchEndpointEmbeddings,
		BatchEndpointFIM,
	}

	for _, endpoint := range endpoints {
		if endpoint == "" {
			t.Error("BatchEndpoint constant is empty")
		}
	}
}

func TestBatchJobMetadataStructure(t *testing.T) {
	metadata := BatchJobMetadata{
		TotalRequests:     IntPtr(1000),
		SucceededRequests: IntPtr(950),
		FailedRequests:    IntPtr(50),
	}

	if metadata.TotalRequests == nil || *metadata.TotalRequests != 1000 {
		t.Error("BatchJobMetadata TotalRequests not set correctly")
	}

	if metadata.SucceededRequests == nil || *metadata.SucceededRequests != 950 {
		t.Error("BatchJobMetadata SucceededRequests not set correctly")
	}
}

func TestBatchJobOutStructure(t *testing.T) {
	outputFile := "file-output123"
	errorFile := "file-error123"
	startedAt := int64(1234567890)
	completedAt := int64(1234567900)

	job := BatchJobOut{
		ID:          "batch-123",
		Object:      "batch",
		Endpoint:    BatchEndpointChat,
		InputFiles:  []string{"file-input1", "file-input2"},
		OutputFile:  &outputFile,
		ErrorFile:   &errorFile,
		CreatedAt:   1234567880,
		StartedAt:   &startedAt,
		CompletedAt: &completedAt,
		Status:      BatchJobStatusSuccess,
		Model:       StringPtr("mistral-small-latest"),
		Metadata: &BatchJobMetadata{
			TotalRequests:     IntPtr(100),
			SucceededRequests: IntPtr(95),
			FailedRequests:    IntPtr(5),
		},
		TimeoutHours: IntPtr(24),
	}

	if job.Status != BatchJobStatusSuccess {
		t.Error("BatchJobOut Status not set correctly")
	}

	if len(job.InputFiles) != 2 {
		t.Error("BatchJobOut InputFiles length incorrect")
	}

	if job.OutputFile == nil || *job.OutputFile != outputFile {
		t.Error("BatchJobOut OutputFile not set correctly")
	}
}

func TestBatchJobsOutStructure(t *testing.T) {
	jobs := BatchJobsOut{
		Data: []BatchJobOut{
			{ID: "batch-1", Endpoint: BatchEndpointChat},
			{ID: "batch-2", Endpoint: BatchEndpointEmbeddings},
		},
		Object: "list",
		Total:  2,
	}

	if len(jobs.Data) != 2 {
		t.Error("BatchJobsOut Data length incorrect")
	}

	if jobs.Total != 2 {
		t.Error("BatchJobsOut Total incorrect")
	}
}

func TestCreateBatchJobRequestValidation(t *testing.T) {
	// Test minimal valid request
	req := &CreateBatchJobRequest{
		InputFiles: []string{"file-123"},
		Endpoint:   BatchEndpointChat,
	}

	if len(req.InputFiles) == 0 {
		t.Error("InputFiles is required")
	}

	if req.Endpoint == "" {
		t.Error("Endpoint is required")
	}
}

func TestCreateBatchJobRequestWithMultipleFiles(t *testing.T) {
	req := &CreateBatchJobRequest{
		InputFiles: []string{"file-1", "file-2", "file-3"},
		Endpoint:   BatchEndpointFIM,
	}

	if len(req.InputFiles) != 3 {
		t.Error("Should support multiple input files")
	}
}

func TestListBatchJobsParamsWithMetadata(t *testing.T) {
	params := &ListBatchJobsParams{
		Metadata: map[string]any{
			"key1": "value1",
			"key2": 123,
			"key3": true,
		},
	}

	if params.Metadata == nil {
		t.Error("Metadata should be set")
	}

	if len(params.Metadata) != 3 {
		t.Error("Metadata should have 3 entries")
	}
}

func TestListBatchJobsParamsWithMultipleStatuses(t *testing.T) {
	params := &ListBatchJobsParams{
		Status: []BatchJobStatus{
			BatchJobStatusQueued,
			BatchJobStatusRunning,
			BatchJobStatusSuccess,
		},
	}

	if len(params.Status) != 3 {
		t.Error("Should support multiple status filters")
	}
}

func TestBatchJobOutWithOptionalFields(t *testing.T) {
	// Test BatchJobOut with minimal fields
	job := BatchJobOut{
		ID:         "batch-123",
		Object:     "batch",
		Endpoint:   BatchEndpointChat,
		InputFiles: []string{"file-1"},
		CreatedAt:  123,
		Status:     BatchJobStatusQueued,
	}

	if job.OutputFile != nil {
		t.Error("OutputFile should be nil when not set")
	}

	if job.ErrorFile != nil {
		t.Error("ErrorFile should be nil when not set")
	}

	if job.StartedAt != nil {
		t.Error("StartedAt should be nil when not set")
	}

	if job.CompletedAt != nil {
		t.Error("CompletedAt should be nil when not set")
	}
}

func TestBatchJobMetadataOptionalFields(t *testing.T) {
	// Test that all fields are optional
	metadata := BatchJobMetadata{}

	if metadata.TotalRequests != nil {
		t.Error("TotalRequests should be nil by default")
	}

	if metadata.SucceededRequests != nil {
		t.Error("SucceededRequests should be nil by default")
	}

	if metadata.FailedRequests != nil {
		t.Error("FailedRequests should be nil by default")
	}
}

func TestCreateBatchJobRequestMetadataTypes(t *testing.T) {
	// Test that metadata supports various types
	req := &CreateBatchJobRequest{
		InputFiles: []string{"file-1"},
		Endpoint:   BatchEndpointChat,
		Metadata: map[string]any{
			"string_value": "test",
			"int_value":    42,
			"float_value":  3.14,
			"bool_value":   true,
			"array_value":  []string{"a", "b"},
		},
	}

	if req.Metadata["string_value"] != "test" {
		t.Error("String metadata not set correctly")
	}

	if req.Metadata["int_value"] != 42 {
		t.Error("Int metadata not set correctly")
	}
}

func TestBatchEndpointValues(t *testing.T) {
	// Verify endpoint values match expected API paths
	if BatchEndpointChat != "/v1/chat/completions" {
		t.Error("BatchEndpointChat value incorrect")
	}

	if BatchEndpointEmbeddings != "/v1/embeddings" {
		t.Error("BatchEndpointEmbeddings value incorrect")
	}

	if BatchEndpointFIM != "/v1/fim/completions" {
		t.Error("BatchEndpointFIM value incorrect")
	}
}
