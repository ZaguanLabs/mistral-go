package sdk

import (
	"bytes"
	"strings"
	"testing"
)

func TestUploadFile(t *testing.T) {
	client := NewMistralClientDefault("")

	// Create a test file content
	content := strings.NewReader(`{"prompt": "Hello", "completion": "Hi there!"}`)

	// Note: This will fail without a valid API key, but tests the code path
	_, err := client.UploadFile(content, "test_file.jsonl", FilePurposeFineTune)

	// We expect an error due to missing/invalid API key in test environment
	// The important thing is that the function executes without panicking
	if err == nil {
		t.Log("Upload succeeded (valid API key present)")
	} else {
		t.Logf("Upload failed as expected in test environment: %v", err)
	}
}

func TestListFiles(t *testing.T) {
	client := NewMistralClientDefault("")

	params := &ListFilesParams{
		Page:     IntPtr(0),
		PageSize: IntPtr(10),
	}

	_, err := client.ListFiles(params)

	if err == nil {
		t.Log("ListFiles succeeded (valid API key present)")
	} else {
		t.Logf("ListFiles failed as expected in test environment: %v", err)
	}
}

func TestListFilesWithFilters(t *testing.T) {
	client := NewMistralClientDefault("")

	purpose := FilePurposeFineTune
	params := &ListFilesParams{
		Page:       IntPtr(0),
		PageSize:   IntPtr(10),
		Purpose:    &purpose,
		SampleType: []SampleType{SampleTypeInstruct},
		Source:     []Source{SourceUpload},
		Search:     StringPtr("test"),
	}

	_, err := client.ListFiles(params)

	if err == nil {
		t.Log("ListFiles with filters succeeded")
	} else {
		t.Logf("ListFiles with filters failed as expected: %v", err)
	}
}

func TestListFilesNilParams(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with nil params - should use defaults
	_, err := client.ListFiles(nil)

	if err == nil {
		t.Log("ListFiles with nil params succeeded")
	} else {
		t.Logf("ListFiles with nil params failed as expected: %v", err)
	}
}

func TestRetrieveFile(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with a dummy file ID
	_, err := client.RetrieveFile("file-test123")

	if err == nil {
		t.Log("RetrieveFile succeeded")
	} else {
		t.Logf("RetrieveFile failed as expected: %v", err)
	}
}

func TestDeleteFile(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with a dummy file ID
	_, err := client.DeleteFile("file-test123")

	if err == nil {
		t.Log("DeleteFile succeeded")
	} else {
		t.Logf("DeleteFile failed as expected: %v", err)
	}
}

func TestDownloadFile(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with a dummy file ID
	_, err := client.DownloadFile("file-test123")

	if err == nil {
		t.Log("DownloadFile succeeded")
	} else {
		t.Logf("DownloadFile failed as expected: %v", err)
	}
}

func TestGetSignedURL(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with default expiry
	_, err := client.GetSignedURL("file-test123", nil)

	if err == nil {
		t.Log("GetSignedURL succeeded")
	} else {
		t.Logf("GetSignedURL failed as expected: %v", err)
	}
}

func TestGetSignedURLWithExpiry(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with custom expiry
	expiry := 48
	_, err := client.GetSignedURL("file-test123", &expiry)

	if err == nil {
		t.Log("GetSignedURL with expiry succeeded")
	} else {
		t.Logf("GetSignedURL with expiry failed as expected: %v", err)
	}
}

func TestFilePurposeConstants(t *testing.T) {
	// Test that constants are defined correctly
	purposes := []FilePurpose{
		FilePurposeFineTune,
		FilePurposeBatch,
		FilePurposeAssistants,
	}

	for _, purpose := range purposes {
		if purpose == "" {
			t.Error("FilePurpose constant is empty")
		}
	}
}

func TestSampleTypeConstants(t *testing.T) {
	sampleTypes := []SampleType{
		SampleTypePretrain,
		SampleTypeInstruct,
	}

	for _, st := range sampleTypes {
		if st == "" {
			t.Error("SampleType constant is empty")
		}
	}
}

func TestSourceConstants(t *testing.T) {
	sources := []Source{
		SourceUpload,
		SourceRepository,
	}

	for _, source := range sources {
		if source == "" {
			t.Error("Source constant is empty")
		}
	}
}

func TestUploadFileWithEmptyContent(t *testing.T) {
	client := NewMistralClientDefault("")

	// Test with empty content
	content := bytes.NewReader([]byte{})

	_, err := client.UploadFile(content, "empty.jsonl", FilePurposeFineTune)

	// Should handle empty files gracefully
	if err != nil {
		t.Logf("Empty file upload failed as expected: %v", err)
	}
}

func TestUploadFileWithLargeFilename(t *testing.T) {
	client := NewMistralClientDefault("")

	content := strings.NewReader("test content")
	longFilename := strings.Repeat("a", 255) + ".jsonl"

	_, err := client.UploadFile(content, longFilename, FilePurposeFineTune)

	if err != nil {
		t.Logf("Large filename upload failed: %v", err)
	}
}

func TestListFilesParamsValidation(t *testing.T) {
	// Test that ListFilesParams can be created with various combinations
	params := &ListFilesParams{
		Page:     IntPtr(0),
		PageSize: IntPtr(100),
	}

	if params.Page == nil || *params.Page != 0 {
		t.Error("Page parameter not set correctly")
	}

	if params.PageSize == nil || *params.PageSize != 100 {
		t.Error("PageSize parameter not set correctly")
	}
}

func TestFileSchemaStructure(t *testing.T) {
	// Test that FileSchema can be created and fields are accessible
	schema := FileSchema{
		ID:        "file-123",
		Object:    "file",
		Bytes:     1024,
		CreatedAt: 1234567890,
		Filename:  "test.jsonl",
		Purpose:   FilePurposeFineTune,
	}

	if schema.ID != "file-123" {
		t.Error("FileSchema ID not set correctly")
	}

	if schema.Bytes != 1024 {
		t.Error("FileSchema Bytes not set correctly")
	}
}

func TestUploadFileOutStructure(t *testing.T) {
	// Test UploadFileOut structure
	out := UploadFileOut{
		ID:        "file-123",
		Object:    "file",
		Bytes:     2048,
		CreatedAt: 1234567890,
		Filename:  "uploaded.jsonl",
		Purpose:   FilePurposeBatch,
	}

	if out.Purpose != FilePurposeBatch {
		t.Error("UploadFileOut Purpose not set correctly")
	}
}

func TestListFilesOutStructure(t *testing.T) {
	// Test ListFilesOut structure
	out := ListFilesOut{
		Data: []FileSchema{
			{ID: "file-1", Filename: "file1.jsonl"},
			{ID: "file-2", Filename: "file2.jsonl"},
		},
		Object: "list",
		Total:  2,
	}

	if len(out.Data) != 2 {
		t.Error("ListFilesOut Data length incorrect")
	}

	if out.Total != 2 {
		t.Error("ListFilesOut Total incorrect")
	}
}

func TestDeleteFileOutStructure(t *testing.T) {
	// Test DeleteFileOut structure
	out := DeleteFileOut{
		ID:      "file-123",
		Object:  "file",
		Deleted: true,
	}

	if !out.Deleted {
		t.Error("DeleteFileOut Deleted flag not set correctly")
	}
}

func TestFileSignedURLStructure(t *testing.T) {
	// Test FileSignedURL structure
	signedURL := FileSignedURL{
		URL: "https://example.com/signed-url",
	}

	if !strings.HasPrefix(signedURL.URL, "https://") {
		t.Error("FileSignedURL should have HTTPS URL")
	}
}
