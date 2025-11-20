package sdk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListDocuments(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := DocumentListResponse{
			Object: "list",
			Data: []Document{
				{ID: "doc-1", Name: "Document 1", Status: "processed"},
				{ID: "doc-2", Name: "Document 2", Status: "processing"},
			},
			Total: 2,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ListDocuments("lib-123", 0)
	if err != nil {
		t.Fatalf("ListDocuments failed: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 documents, got %d", len(resp.Data))
	}
}

func TestUploadDocument(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := DocumentUploadResponse{
			ID:     "doc-123",
			Object: "document",
			Status: "uploaded",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	file := bytes.NewReader([]byte("test file content"))

	resp, err := client.UploadDocument("lib-123", file, "test.txt")
	if err != nil {
		t.Fatalf("UploadDocument failed: %v", err)
	}

	if resp.ID != "doc-123" {
		t.Errorf("Expected doc-123, got %s", resp.ID)
	}
}

func TestGetDocument(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Document{
			ID:     "doc-123",
			Name:   "Test Document",
			Status: "processed",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.GetDocument("lib-123", "doc-123")
	if err != nil {
		t.Fatalf("GetDocument failed: %v", err)
	}

	if resp.ID != "doc-123" {
		t.Errorf("Expected doc-123, got %s", resp.ID)
	}
}

func TestUpdateDocument(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Document{
			ID:   "doc-123",
			Name: "Updated Document",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &UpdateDocumentRequest{
		Name: StringPtr("Updated Document"),
	}

	resp, err := client.UpdateDocument("lib-123", "doc-123", req)
	if err != nil {
		t.Fatalf("UpdateDocument failed: %v", err)
	}

	if resp.Name != "Updated Document" {
		t.Errorf("Expected Updated Document, got %s", resp.Name)
	}
}

func TestUpdateDocumentNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.UpdateDocument("lib-123", "doc-123", nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}

func TestDeleteDocument(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := DeleteDocumentResponse{
			ID:      "doc-123",
			Object:  "document",
			Deleted: true,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.DeleteDocument("lib-123", "doc-123")
	if err != nil {
		t.Fatalf("DeleteDocument failed: %v", err)
	}

	if !resp.Deleted {
		t.Error("Expected document to be deleted")
	}
}

func TestGetDocumentStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/status") {
			t.Errorf("Expected /status in path, got %s", r.URL.Path)
		}

		response := DocumentStatusResponse{
			ID:     "doc-123",
			Status: "processed",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.GetDocumentStatus("lib-123", "doc-123")
	if err != nil {
		t.Fatalf("GetDocumentStatus failed: %v", err)
	}

	if resp.Status != "processed" {
		t.Errorf("Expected processed, got %s", resp.Status)
	}
}
