package sdk

import (
	"net/http"
	"strings"
	"testing"
)

func TestListModelsWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/models" {
			t.Errorf("Expected path /v1/models, got %s", r.URL.Path)
		}
		MockListModelsResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	models, err := client.ListModels()

	if err != nil {
		t.Fatalf("ListModels failed: %v", err)
	}

	if models == nil {
		t.Fatal("Models should not be nil")
	}

	if len(models.Data) != 2 {
		t.Errorf("Expected 2 models, got %d", len(models.Data))
	}

	if models.Data[0].ID != "mistral-small-latest" {
		t.Errorf("Expected first model ID 'mistral-small-latest', got '%s'", models.Data[0].ID)
	}
}

func TestChatWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("Expected path /v1/chat/completions, got %s", r.URL.Path)
		}

		body := ReadRequestBody(r)
		if !strings.Contains(body, "Hello") {
			t.Error("Request body should contain 'Hello'")
		}

		MockChatResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.Chat(
		"mistral-small-latest",
		[]ChatMessage{
			{Role: RoleUser, Content: "Hello"},
		},
		nil,
	)

	if err != nil {
		t.Fatalf("Chat failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	if len(response.Choices) == 0 {
		t.Fatal("Should have at least one choice")
	}

	if response.Choices[0].Message.Role != RoleAssistant {
		t.Errorf("Expected assistant role, got %s", response.Choices[0].Message.Role)
	}
}

func TestEmbeddingsWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/embeddings" {
			t.Errorf("Expected path /v1/embeddings, got %s", r.URL.Path)
		}
		MockEmbeddingsResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.Embeddings("mistral-embed", []string{"test"})

	if err != nil {
		t.Fatalf("Embeddings failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	if len(response.Data) == 0 {
		t.Fatal("Should have at least one embedding")
	}

	if len(response.Data[0].Embedding) != 3 {
		t.Errorf("Expected 3 dimensions, got %d", len(response.Data[0].Embedding))
	}
}

func TestUploadFileWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files" {
			t.Errorf("Expected path /v1/files, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		MockFileUploadResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	content := strings.NewReader(`{"test": "data"}`)
	response, err := client.UploadFile(content, "test.jsonl", FilePurposeFineTune)

	if err != nil {
		t.Fatalf("UploadFile failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	if response.ID != "file-123" {
		t.Errorf("Expected file ID 'file-123', got '%s'", response.ID)
	}
}

func TestListFilesWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/v1/files") {
			t.Errorf("Expected path /v1/files, got %s", r.URL.Path)
		}
		MockListFilesResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.ListFiles(nil)

	if err != nil {
		t.Fatalf("ListFiles failed: %v", err)
	}

	if response == nil {
		t.Fatal("Response should not be nil")
	}

	if response.Total != 1 {
		t.Errorf("Expected total 1, got %d", response.Total)
	}
}

func TestRetrieveFileWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/v1/files/file-123") {
			t.Errorf("Expected path to contain /v1/files/file-123, got %s", r.URL.Path)
		}

		MockJSONResponse(200, `{
			"id": "file-123",
			"object": "file",
			"bytes": 1024,
			"created_at": 1234567890,
			"filename": "test.jsonl",
			"purpose": "fine-tune"
		}`).Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.RetrieveFile("file-123")

	if err != nil {
		t.Fatalf("RetrieveFile failed: %v", err)
	}

	if response.ID != "file-123" {
		t.Errorf("Expected file ID 'file-123', got '%s'", response.ID)
	}
}

func TestDeleteFileWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		MockJSONResponse(200, `{
			"id": "file-123",
			"object": "file",
			"deleted": true
		}`).Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.DeleteFile("file-123")

	if err != nil {
		t.Fatalf("DeleteFile failed: %v", err)
	}

	if !response.Deleted {
		t.Error("File should be marked as deleted")
	}
}

func TestCreateFineTuningJobWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/fine_tuning/jobs" {
			t.Errorf("Expected path /v1/fine_tuning/jobs, got %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		MockFineTuningJobResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.CreateFineTuningJob(&CreateFineTuningJobRequest{
		Model: "open-mistral-7b",
		TrainingFiles: []TrainingFile{
			{FileID: "file-123"},
		},
		Hyperparameters: Hyperparameters{
			TrainingSteps: IntPtr(100),
		},
	})

	if err != nil {
		t.Fatalf("CreateFineTuningJob failed: %v", err)
	}

	if response.ID != "job-123" {
		t.Errorf("Expected job ID 'job-123', got '%s'", response.ID)
	}
}

func TestListFineTuningJobsWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		MockListFineTuningJobsResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.ListFineTuningJobs(nil)

	if err != nil {
		t.Fatalf("ListFineTuningJobs failed: %v", err)
	}

	if response.Total != 1 {
		t.Errorf("Expected total 1, got %d", response.Total)
	}
}

func TestGetFineTuningJobWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		MockFineTuningJobResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.GetFineTuningJob("job-123")

	if err != nil {
		t.Fatalf("GetFineTuningJob failed: %v", err)
	}

	if response.ID != "job-123" {
		t.Errorf("Expected job ID 'job-123', got '%s'", response.ID)
	}
}

func TestCancelFineTuningJobWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		MockFineTuningJobResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.CancelFineTuningJob("job-123")

	if err != nil {
		t.Fatalf("CancelFineTuningJob failed: %v", err)
	}

	if response.ID != "job-123" {
		t.Errorf("Expected job ID 'job-123', got '%s'", response.ID)
	}
}

func TestCreateBatchJobWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		MockBatchJobResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.CreateBatchJob(&CreateBatchJobRequest{
		InputFiles: []string{"file-123"},
		Endpoint:   BatchEndpointChat,
	})

	if err != nil {
		t.Fatalf("CreateBatchJob failed: %v", err)
	}

	if response.ID != "batch-123" {
		t.Errorf("Expected batch ID 'batch-123', got '%s'", response.ID)
	}
}

func TestListBatchJobsWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		MockListBatchJobsResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.ListBatchJobs(nil)

	if err != nil {
		t.Fatalf("ListBatchJobs failed: %v", err)
	}

	if response.Total != 1 {
		t.Errorf("Expected total 1, got %d", response.Total)
	}
}

func TestGetBatchJobWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		MockBatchJobResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.GetBatchJob("batch-123")

	if err != nil {
		t.Fatalf("GetBatchJob failed: %v", err)
	}

	if response.ID != "batch-123" {
		t.Errorf("Expected batch ID 'batch-123', got '%s'", response.ID)
	}
}

func TestFIMWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		MockFIMResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	response, err := client.FIM(&FIMRequestParams{
		Model:  "codestral-latest",
		Prompt: "def f(",
		Suffix: StringPtr("return a + b"),
	})

	if err != nil {
		t.Fatalf("FIM failed: %v", err)
	}

	if len(response.Choices) == 0 {
		t.Fatal("Should have at least one choice")
	}
}

func TestErrorHandlingWithMock(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		MockErrorResponse(400, "Bad Request").Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	_, err := client.ListModels()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Just verify we got an error - the exact type depends on internal implementation
	if err.Error() == "" {
		t.Error("Error message should not be empty")
	}
}

func TestRetryLogicWithMock(t *testing.T) {
	attemptCount := 0
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount < 3 {
			// Return retryable error
			w.WriteHeader(429)
			return
		}
		// Success on third attempt
		MockListModelsResponse().Write(w)
	})
	defer mock.Close()

	client := mock.GetClient()
	_, err := client.ListModels()

	if err != nil {
		t.Fatalf("Expected success after retries, got error: %v", err)
	}

	if attemptCount < 2 {
		t.Errorf("Expected at least 2 attempts, got %d", attemptCount)
	}
}
