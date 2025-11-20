package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListLibraryAccesses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		response := AccessListResponse{
			Object: "list",
			Data: []LibraryAccess{
				{UserID: "user-1", Permissions: []string{"read"}},
				{UserID: "user-2", Permissions: []string{"read", "write"}},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ListLibraryAccesses("lib-123")
	if err != nil {
		t.Fatalf("ListLibraryAccesses failed: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("Expected 2 accesses, got %d", len(resp.Data))
	}
}

func TestUpdateOrCreateLibraryAccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT, got %s", r.Method)
		}

		response := LibraryAccess{
			UserID:      "user-123",
			Permissions: []string{"read", "write"},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	req := &UpdateAccessRequest{
		UserID:      "user-123",
		Permissions: []string{"read", "write"},
	}

	resp, err := client.UpdateOrCreateLibraryAccess("lib-123", req)
	if err != nil {
		t.Fatalf("UpdateOrCreateLibraryAccess failed: %v", err)
	}

	if resp.UserID != "user-123" {
		t.Errorf("Expected user-123, got %s", resp.UserID)
	}

	if len(resp.Permissions) != 2 {
		t.Errorf("Expected 2 permissions, got %d", len(resp.Permissions))
	}
}

func TestUpdateOrCreateLibraryAccessNilRequest(t *testing.T) {
	client := NewMistralClientDefault("")

	_, err := client.UpdateOrCreateLibraryAccess("lib-123", nil)
	if err == nil {
		t.Error("Expected error for nil request")
	}
}

func TestDeleteLibraryAccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE, got %s", r.Method)
		}

		response := DeleteAccessResponse{
			UserID:  "user-123",
			Deleted: true,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.DeleteLibraryAccess("lib-123", "user-123")
	if err != nil {
		t.Fatalf("DeleteLibraryAccess failed: %v", err)
	}

	if !resp.Deleted {
		t.Error("Expected access to be deleted")
	}

	if resp.UserID != "user-123" {
		t.Errorf("Expected user-123, got %s", resp.UserID)
	}
}
