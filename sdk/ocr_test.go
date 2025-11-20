package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProcessOCR(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := OCRResponse{
			ID:     "ocr-123",
			Object: "ocr",
			Model:  "pixtral-12b-2409",
			Pages: []OCRPageObject{
				{
					PageNumber: 0,
					Text:       "Extracted text from page 1",
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	doc := OCRDocument{
		URL: StringPtr("https://example.com/doc.pdf"),
	}

	resp, err := client.ProcessOCR("pixtral-12b-2409", doc, nil)
	if err != nil {
		t.Fatalf("ProcessOCR failed: %v", err)
	}

	if resp.ID != "ocr-123" {
		t.Errorf("Expected ocr-123, got %s", resp.ID)
	}

	if len(resp.Pages) != 1 {
		t.Errorf("Expected 1 page, got %d", len(resp.Pages))
	}
}

func TestProcessOCRFromURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OCRResponse{
			ID:    "ocr-123",
			Model: "pixtral-12b-2409",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ProcessOCRFromURL("pixtral-12b-2409", "https://example.com/doc.pdf", nil)
	if err != nil {
		t.Fatalf("ProcessOCRFromURL failed: %v", err)
	}

	if resp.ID != "ocr-123" {
		t.Errorf("Expected ocr-123, got %s", resp.ID)
	}
}

func TestProcessOCRFromBase64(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OCRResponse{
			ID:    "ocr-123",
			Model: "pixtral-12b-2409",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ProcessOCRFromBase64("pixtral-12b-2409", "base64data", nil)
	if err != nil {
		t.Fatalf("ProcessOCRFromBase64 failed: %v", err)
	}

	if resp.ID != "ocr-123" {
		t.Errorf("Expected ocr-123, got %s", resp.ID)
	}
}

func TestProcessOCRFromFileID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OCRResponse{
			ID:    "ocr-123",
			Model: "pixtral-12b-2409",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.ProcessOCRFromFileID("pixtral-12b-2409", "file-123", nil)
	if err != nil {
		t.Fatalf("ProcessOCRFromFileID failed: %v", err)
	}

	if resp.ID != "ocr-123" {
		t.Errorf("Expected ocr-123, got %s", resp.ID)
	}
}

func TestProcessOCRWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := OCRResponse{
			ID:    "ocr-123",
			Model: "pixtral-12b-2409",
			Pages: []OCRPageObject{
				{PageNumber: 0, Text: "Page 1"},
				{PageNumber: 1, Text: "Page 2"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	doc := OCRDocument{
		URL: StringPtr("https://example.com/doc.pdf"),
	}

	params := &OCRRequest{
		Pages:              []int{0, 1},
		IncludeImageBase64: BoolPtr(true),
		ImageLimit:         IntPtr(5),
	}

	resp, err := client.ProcessOCR("pixtral-12b-2409", doc, params)
	if err != nil {
		t.Fatalf("ProcessOCR with params failed: %v", err)
	}

	if len(resp.Pages) != 2 {
		t.Errorf("Expected 2 pages, got %d", len(resp.Pages))
	}
}
