package sdk

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockHTTPServer creates a test HTTP server for mocking API responses
type MockHTTPServer struct {
	Server   *httptest.Server
	Requests []*http.Request
	t        *testing.T
}

// NewMockHTTPServer creates a new mock HTTP server
func NewMockHTTPServer(t *testing.T, handler http.HandlerFunc) *MockHTTPServer {
	mock := &MockHTTPServer{
		Requests: make([]*http.Request, 0),
		t:        t,
	}

	mock.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Store request for inspection
		mock.Requests = append(mock.Requests, r)

		// Call the handler
		handler(w, r)
	}))

	return mock
}

// Close closes the mock server
func (m *MockHTTPServer) Close() {
	m.Server.Close()
}

// GetClient returns a MistralClient configured to use the mock server
func (m *MockHTTPServer) GetClient() *MistralClient {
	return NewMistralClient("test-api-key", m.Server.URL, DefaultMaxRetries, DefaultTimeout)
}

// MockResponse is a helper to create mock HTTP responses
type MockResponse struct {
	StatusCode int
	Body       string
	Headers    map[string]string
}

// Write writes the mock response to the ResponseWriter
func (mr *MockResponse) Write(w http.ResponseWriter) {
	for key, value := range mr.Headers {
		w.Header().Set(key, value)
	}
	w.WriteHeader(mr.StatusCode)
	_, _ = io.WriteString(w, mr.Body) // Ignore error in test helper
}

// MockJSONResponse creates a mock JSON response
func MockJSONResponse(statusCode int, body string) *MockResponse {
	return &MockResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// MockErrorResponse creates a mock error response
func MockErrorResponse(statusCode int, message string) *MockResponse {
	return MockJSONResponse(statusCode, `{"error": "`+message+`"}`)
}

// MockListModelsResponse creates a mock response for ListModels
func MockListModelsResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"object": "list",
		"data": [
			{
				"id": "mistral-small-latest",
				"object": "model",
				"created": 1234567890,
				"owned_by": "mistral",
				"permission": []
			},
			{
				"id": "mistral-large-latest",
				"object": "model",
				"created": 1234567891,
				"owned_by": "mistral",
				"permission": []
			}
		]
	}`)
}

// MockChatResponse creates a mock response for Chat
func MockChatResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"id": "chat-123",
		"object": "chat.completion",
		"created": 1234567890,
		"model": "mistral-small-latest",
		"choices": [
			{
				"index": 0,
				"message": {
					"role": "assistant",
					"content": "Hello! How can I help you today?"
				},
				"finish_reason": "stop"
			}
		],
		"usage": {
			"prompt_tokens": 10,
			"completion_tokens": 20,
			"total_tokens": 30
		}
	}`)
}

// MockEmbeddingsResponse creates a mock response for Embeddings
func MockEmbeddingsResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"object": "list",
		"data": [
			{
				"object": "embedding",
				"embedding": [0.1, 0.2, 0.3],
				"index": 0
			}
		],
		"model": "mistral-embed",
		"usage": {
			"prompt_tokens": 5,
			"total_tokens": 5
		}
	}`)
}

// MockFileUploadResponse creates a mock response for file upload
func MockFileUploadResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"id": "file-123",
		"object": "file",
		"bytes": 1024,
		"created_at": 1234567890,
		"filename": "test.jsonl",
		"purpose": "fine-tune"
	}`)
}

// MockListFilesResponse creates a mock response for listing files
func MockListFilesResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"object": "list",
		"data": [
			{
				"id": "file-123",
				"object": "file",
				"bytes": 1024,
				"created_at": 1234567890,
				"filename": "test.jsonl",
				"purpose": "fine-tune"
			}
		],
		"total": 1
	}`)
}

// MockFineTuningJobResponse creates a mock response for fine-tuning job
func MockFineTuningJobResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"id": "job-123",
		"model": "open-mistral-7b",
		"status": "QUEUED",
		"job_type": "FT",
		"created_at": 1234567890,
		"modified_at": 1234567890,
		"training_files": ["file-123"],
		"validation_files": [],
		"object": "fine_tuning.job",
		"hyperparameters": {
			"training_steps": 100
		}
	}`)
}

// MockListFineTuningJobsResponse creates a mock response for listing fine-tuning jobs
func MockListFineTuningJobsResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"object": "list",
		"data": [
			{
				"id": "job-123",
				"model": "open-mistral-7b",
				"status": "QUEUED",
				"created_at": 1234567890,
				"modified_at": 1234567890,
				"training_files": ["file-123"],
				"validation_files": [],
				"object": "fine_tuning.job",
				"hyperparameters": {}
			}
		],
		"total": 1
	}`)
}

// MockBatchJobResponse creates a mock response for batch job
func MockBatchJobResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"id": "batch-123",
		"object": "batch",
		"endpoint": "/v1/chat/completions",
		"input_files": ["file-123"],
		"created_at": 1234567890,
		"status": "QUEUED"
	}`)
}

// MockListBatchJobsResponse creates a mock response for listing batch jobs
func MockListBatchJobsResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"object": "list",
		"data": [
			{
				"id": "batch-123",
				"object": "batch",
				"endpoint": "/v1/chat/completions",
				"input_files": ["file-123"],
				"created_at": 1234567890,
				"status": "QUEUED"
			}
		],
		"total": 1
	}`)
}

// MockFIMResponse creates a mock response for FIM
func MockFIMResponse() *MockResponse {
	return MockJSONResponse(200, `{
		"id": "fim-123",
		"object": "fim.completion",
		"created": 1234567890,
		"model": "codestral-latest",
		"choices": [
			{
				"index": 0,
				"message": {
					"role": "assistant",
					"content": "a, b):"
				},
				"finish_reason": "stop"
			}
		]
	}`)
}

// ReadRequestBody reads and returns the request body as a string
func ReadRequestBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}
	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	return string(body)
}
