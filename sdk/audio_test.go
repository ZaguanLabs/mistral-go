package sdk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranscribe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		response := TranscriptionResponse{
			Text:     "This is a test transcription",
			Language: StringPtr("en"),
			Duration: Float64Ptr(5.5),
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	audioData := bytes.NewReader([]byte("fake audio data"))

	resp, err := client.Transcribe("whisper-large-v3", audioData, "test.mp3", nil)
	if err != nil {
		t.Fatalf("Transcribe failed: %v", err)
	}

	if resp.Text != "This is a test transcription" {
		t.Errorf("Expected 'This is a test transcription', got %s", resp.Text)
	}

	if resp.Language == nil || *resp.Language != "en" {
		t.Error("Expected language 'en'")
	}
}

func TestTranscribeFromURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := TranscriptionResponse{
			Text:     "Transcription from URL",
			Language: StringPtr("en"),
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.TranscribeFromURL("whisper-large-v3", "https://example.com/audio.mp3", nil)
	if err != nil {
		t.Fatalf("TranscribeFromURL failed: %v", err)
	}

	if resp.Text != "Transcription from URL" {
		t.Errorf("Expected 'Transcription from URL', got %s", resp.Text)
	}
}

func TestTranscribeFromFileID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := TranscriptionResponse{
			Text:     "Transcription from file ID",
			Language: StringPtr("en"),
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	resp, err := client.TranscribeFromFileID("whisper-large-v3", "file-123", nil)
	if err != nil {
		t.Fatalf("TranscribeFromFileID failed: %v", err)
	}

	if resp.Text != "Transcription from file ID" {
		t.Errorf("Expected 'Transcription from file ID', got %s", resp.Text)
	}
}

func TestTranscribeWithParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := TranscriptionResponse{
			Text:     "Detailed transcription",
			Language: StringPtr("fr"),
			Words: []TranscriptionWord{
				{Word: "Bonjour", Start: 0.0, End: 0.5},
				{Word: "monde", Start: 0.6, End: 1.0},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := NewMistralClient("test-key", server.URL, 1, DefaultTimeout)

	audioData := bytes.NewReader([]byte("fake audio data"))

	params := &TranscriptionRequest{
		Language:               StringPtr("fr"),
		Temperature:            Float64Ptr(0.0),
		TimestampGranularities: []TimestampGranularity{TimestampGranularityWord},
	}

	resp, err := client.Transcribe("whisper-large-v3", audioData, "test.mp3", params)
	if err != nil {
		t.Fatalf("Transcribe with params failed: %v", err)
	}

	if resp.Language == nil || *resp.Language != "fr" {
		t.Errorf("Expected language 'fr', got %s", *resp.Language)
	}

	if len(resp.Words) != 2 {
		t.Errorf("Expected 2 words, got %d", len(resp.Words))
	}
}
