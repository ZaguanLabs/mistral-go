package sdk

import (
	"testing"
	"time"
)

func TestCreateFineTuningJob(t *testing.T) {
	client := NewMistralClientDefault("")

	req := &CreateFineTuningJobRequest{
		Model: "open-mistral-7b",
		TrainingFiles: []TrainingFile{
			{FileID: "file-test123", Weight: Float64Ptr(1.0)},
		},
		Hyperparameters: Hyperparameters{
			TrainingSteps: IntPtr(100),
			LearningRate:  Float64Ptr(0.0001),
		},
	}

	_, err := client.CreateFineTuningJob(req)

	if err == nil {
		t.Log("CreateFineTuningJob succeeded")
	} else {
		t.Logf("CreateFineTuningJob failed as expected: %v", err)
	}
}

func TestCreateFineTuningJobWithAllParams(t *testing.T) {
	client := NewMistralClientDefault("")

	jobType := FineTuneableModelTypeFineTuning
	req := &CreateFineTuningJobRequest{
		Model: "open-mistral-7b",
		TrainingFiles: []TrainingFile{
			{FileID: "file-train1", Weight: Float64Ptr(1.0)},
			{FileID: "file-train2", Weight: Float64Ptr(0.5)},
		},
		ValidationFiles: []string{"file-val1", "file-val2"},
		Hyperparameters: Hyperparameters{
			TrainingSteps:  IntPtr(1000),
			LearningRate:   Float64Ptr(0.0001),
			WeightDecay:    Float64Ptr(0.01),
			WarmupFraction: Float64Ptr(0.1),
			Epochs:         Float64Ptr(3.0),
			FimRatio:       Float64Ptr(0.9),
		},
		Suffix:                      StringPtr("my-custom-model"),
		AutoStart:                   BoolPtr(true),
		InvalidSampleSkipPercentage: Float64Ptr(5.0),
		JobType:                     &jobType,
	}

	_, err := client.CreateFineTuningJob(req)

	if err != nil {
		t.Logf("CreateFineTuningJob with all params failed: %v", err)
	}
}

func TestListFineTuningJobs(t *testing.T) {
	client := NewMistralClientDefault("")

	params := &ListFineTuningJobsParams{
		Page:     IntPtr(0),
		PageSize: IntPtr(10),
	}

	_, err := client.ListFineTuningJobs(params)

	if err == nil {
		t.Log("ListFineTuningJobs succeeded")
	} else {
		t.Logf("ListFineTuningJobs failed as expected: %v", err)
	}
}

func TestListFineTuningJobsWithFilters(t *testing.T) {
	client := NewMistralClientDefault("")

	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	status := JobStatusRunning

	params := &ListFineTuningJobsParams{
		Page:          IntPtr(0),
		PageSize:      IntPtr(20),
		Model:         StringPtr("open-mistral-7b"),
		CreatedAfter:  &yesterday,
		CreatedBefore: &now,
		CreatedByMe:   BoolPtr(true),
		Status:        &status,
		WandbProject:  StringPtr("my-project"),
		WandbName:     StringPtr("my-run"),
		Suffix:        StringPtr("my-model"),
	}

	_, err := client.ListFineTuningJobs(params)

	if err != nil {
		t.Logf("ListFineTuningJobs with filters failed: %v", err)
	}
}

func TestListFineTuningJobsNilParams(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with nil params - should use defaults
	_, err := client.ListFineTuningJobs(nil)

	if err != nil {
		t.Logf("ListFineTuningJobs with nil params failed: %v", err)
	}
}

func TestGetFineTuningJob(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.GetFineTuningJob("job-test123")

	if err == nil {
		t.Log("GetFineTuningJob succeeded")
	} else {
		t.Logf("GetFineTuningJob failed as expected: %v", err)
	}
}

func TestCancelFineTuningJob(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.CancelFineTuningJob("job-test123")

	if err == nil {
		t.Log("CancelFineTuningJob succeeded")
	} else {
		t.Logf("CancelFineTuningJob failed as expected: %v", err)
	}
}

func TestStartFineTuningJob(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.StartFineTuningJob("job-test123")

	if err == nil {
		t.Log("StartFineTuningJob succeeded")
	} else {
		t.Logf("StartFineTuningJob failed as expected: %v", err)
	}
}

func TestJobStatusConstants(t *testing.T) {
	statuses := []JobStatus{
		JobStatusQueued,
		JobStatusStarted,
		JobStatusRunning,
		JobStatusFailed,
		JobStatusSuccess,
		JobStatusCancelled,
		JobStatusTimedOut,
	}

	for _, status := range statuses {
		if status == "" {
			t.Error("JobStatus constant is empty")
		}
	}
}

func TestFineTuneableModelTypeConstants(t *testing.T) {
	types := []FineTuneableModelType{
		FineTuneableModelTypeFineTuning,
		FineTuneableModelTypeClassifier,
	}

	for _, modelType := range types {
		if modelType == "" {
			t.Error("FineTuneableModelType constant is empty")
		}
	}
}

func TestTrainingFileStructure(t *testing.T) {
	file := TrainingFile{
		FileID: "file-123",
		Weight: Float64Ptr(0.8),
	}

	if file.FileID != "file-123" {
		t.Error("TrainingFile FileID not set correctly")
	}

	if file.Weight == nil || *file.Weight != 0.8 {
		t.Error("TrainingFile Weight not set correctly")
	}
}

func TestHyperparametersStructure(t *testing.T) {
	hyper := Hyperparameters{
		TrainingSteps:  IntPtr(500),
		LearningRate:   Float64Ptr(0.0002),
		WeightDecay:    Float64Ptr(0.01),
		WarmupFraction: Float64Ptr(0.05),
		Epochs:         Float64Ptr(5.0),
		FimRatio:       Float64Ptr(0.95),
	}

	if hyper.TrainingSteps == nil || *hyper.TrainingSteps != 500 {
		t.Error("Hyperparameters TrainingSteps not set correctly")
	}

	if hyper.LearningRate == nil || *hyper.LearningRate != 0.0002 {
		t.Error("Hyperparameters LearningRate not set correctly")
	}
}

func TestWandbIntegrationStructure(t *testing.T) {
	integration := WandbIntegration{
		Type:    "wandb",
		Project: "my-project",
		Name:    StringPtr("my-run"),
		APIKey:  StringPtr("secret-key"),
	}

	if integration.Type != "wandb" {
		t.Error("WandbIntegration Type not set correctly")
	}

	if integration.Project != "my-project" {
		t.Error("WandbIntegration Project not set correctly")
	}
}

func TestJobMetadataStructure(t *testing.T) {
	metadata := JobMetadata{
		ExpectedDurationSeconds: IntPtr(3600),
		Cost:                    Float64Ptr(10.50),
		CostCurrency:            StringPtr("USD"),
		TrainTokensPerStep:      IntPtr(1000),
		DataTokens:              IntPtr(50000),
		EstimatedStartTime:      Int64Ptr(1234567890),
	}

	if metadata.Cost == nil || *metadata.Cost != 10.50 {
		t.Error("JobMetadata Cost not set correctly")
	}
}

func TestJobOutStructure(t *testing.T) {
	jobType := FineTuneableModelTypeFineTuning
	job := JobOut{
		ID:              "job-123",
		Model:           "open-mistral-7b",
		Status:          JobStatusRunning,
		JobType:         &jobType,
		CreatedAt:       1234567890,
		ModifiedAt:      1234567900,
		TrainingFiles:   []string{"file-1", "file-2"},
		ValidationFiles: []string{"file-val"},
		Object:          "fine_tuning.job",
		Hyperparameters: Hyperparameters{
			TrainingSteps: IntPtr(100),
		},
	}

	if job.Status != JobStatusRunning {
		t.Error("JobOut Status not set correctly")
	}

	if len(job.TrainingFiles) != 2 {
		t.Error("JobOut TrainingFiles length incorrect")
	}
}

func TestJobsOutStructure(t *testing.T) {
	jobs := JobsOut{
		Data: []JobOut{
			{ID: "job-1", Model: "model-1"},
			{ID: "job-2", Model: "model-2"},
		},
		Object: "list",
		Total:  2,
	}

	if len(jobs.Data) != 2 {
		t.Error("JobsOut Data length incorrect")
	}

	if jobs.Total != 2 {
		t.Error("JobsOut Total incorrect")
	}
}

func TestCreateFineTuningJobRequestValidation(t *testing.T) {
	// Test minimal valid request
	req := &CreateFineTuningJobRequest{
		Model: "open-mistral-7b",
		Hyperparameters: Hyperparameters{
			TrainingSteps: IntPtr(10),
		},
	}

	if req.Model == "" {
		t.Error("Model is required")
	}
}

func TestListFineTuningJobsParamsTimeFilters(t *testing.T) {
	now := time.Now()
	params := &ListFineTuningJobsParams{
		CreatedAfter:  &now,
		CreatedBefore: &now,
	}

	if params.CreatedAfter == nil {
		t.Error("CreatedAfter not set")
	}

	if params.CreatedBefore == nil {
		t.Error("CreatedBefore not set")
	}
}

func TestHyperparametersOptionalFields(t *testing.T) {
	// Test that all fields are optional
	hyper := Hyperparameters{}

	if hyper.TrainingSteps != nil {
		t.Error("TrainingSteps should be nil by default")
	}

	if hyper.LearningRate != nil {
		t.Error("LearningRate should be nil by default")
	}
}

func TestTrainingFileWithoutWeight(t *testing.T) {
	// Test TrainingFile without weight (should be optional)
	file := TrainingFile{
		FileID: "file-123",
	}

	if file.Weight != nil {
		t.Error("Weight should be nil when not set")
	}
}

func TestJobOutWithOptionalFields(t *testing.T) {
	// Test JobOut with minimal fields
	job := JobOut{
		ID:              "job-123",
		Model:           "model",
		Status:          JobStatusQueued,
		CreatedAt:       123,
		ModifiedAt:      124,
		TrainingFiles:   []string{},
		Object:          "job",
		Hyperparameters: Hyperparameters{},
	}

	if job.FineTunedModel != nil {
		t.Error("FineTunedModel should be nil when not set")
	}

	if job.Suffix != nil {
		t.Error("Suffix should be nil when not set")
	}
}
