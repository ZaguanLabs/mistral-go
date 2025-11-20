package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListLibraries(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		response := LibraryListResponse{
			Object: "list",
			Data: []Library{
				{ID: "lib-1", Name: "Library 1"},
				{ID: "lib-2", Name: "Library 2"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ListLibraries()
	if err != nil {
		t.Fatalf("ListLibraries failed: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 libraries, got %d", len(resp.Data))
	}
}

func TestCreateLibrary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := Library{
			ID:     "lib-123",
			Name:   "Test Library",
			Object: "library",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &CreateLibraryRequest{
		Name:        "Test Library",
		Description: StringPtr("A test library"),
	}

	resp, err := client.CreateLibrary(req)
	if err != nil {
		t.Fatalf("CreateLibrary failed: %v", err)
	}

	if resp.ID != "lib-123" {
		t.Errorf("Expected lib-123, got %s", resp.ID)
	}
}

func TestCreateLibraryNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.CreateLibrary(nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}

func TestGetLibrary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Library{
			ID:     "lib-123",
			Name:   "Test Library",
			Object: "library",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.GetLibrary("lib-123")
	if err != nil {
		t.Fatalf("GetLibrary failed: %v", err)
	}

	if resp.ID != "lib-123" {
		t.Errorf("Expected lib-123, got %s", resp.ID)
	}
}

func TestUpdateLibrary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH, got %s", r.Method)
		}

		response := Library{
			ID:   "lib-123",
			Name: "Updated Library",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &UpdateLibraryRequest{
		Name: StringPtr("Updated Library"),
	}

	resp, err := client.UpdateLibrary("lib-123", req)
	if err != nil {
		t.Fatalf("UpdateLibrary failed: %v", err)
	}

	if resp.Name != "Updated Library" {
		t.Errorf("Expected Updated Library, got %s", resp.Name)
	}
}

func TestUpdateLibraryNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.UpdateLibrary("lib-123", nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}

func TestDeleteLibrary(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}

		response := DeleteLibraryResponse{
			ID:      "lib-123",
			Object:  "library",
			Deleted: true,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.DeleteLibrary("lib-123")
	if err != nil {
		t.Fatalf("DeleteLibrary failed: %v", err)
	}

	if !resp.Deleted {
		t.Error("Expected library to be deleted")
	}
}
