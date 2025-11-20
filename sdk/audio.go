package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

// TimestampGranularity represents the granularity of timestamps in transcription
type TimestampGranularity string

const (
	TimestampGranularityWord    TimestampGranularity = "word"
	TimestampGranularitySegment TimestampGranularity = "segment"
)

// TranscriptionRequest represents a request for audio transcription
type TranscriptionRequest struct {
	Model                  string                 `json:"model"`
	File                   io.Reader              `json:"-"` // File content (not serialized)
	Filename               string                 `json:"-"` // Filename (not serialized)
	FileURL                *string                `json:"file_url,omitempty"`
	FileID                 *string                `json:"file_id,omitempty"`
	Language               *string                `json:"language,omitempty"`
	Temperature            *float64               `json:"temperature,omitempty"`
	TimestampGranularities []TimestampGranularity `json:"timestamp_granularities,omitempty"`
}

// TranscriptionWord represents a word in the transcription with timestamp
type TranscriptionWord struct {
	Word  string  `json:"word"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

// TranscriptionSegment represents a segment in the transcription
type TranscriptionSegment struct {
	ID               int     `json:"id"`
	Seek             int     `json:"seek"`
	Start            float64 `json:"start"`
	End              float64 `json:"end"`
	Text             string  `json:"text"`
	Tokens           []int   `json:"tokens"`
	Temperature      float64 `json:"temperature"`
	AvgLogprob       float64 `json:"avg_logprob"`
	CompressionRatio float64 `json:"compression_ratio"`
	NoSpeechProb     float64 `json:"no_speech_prob"`
}

// TranscriptionResponse represents the response from audio transcription
type TranscriptionResponse struct {
	Text     string                 `json:"text"`
	Language *string                `json:"language,omitempty"`
	Duration *float64               `json:"duration,omitempty"`
	Words    []TranscriptionWord    `json:"words,omitempty"`
	Segments []TranscriptionSegment `json:"segments,omitempty"`
}

// Transcribe transcribes an audio file to text
//
// Parameters:
//   - model: The model to use for transcription (e.g., "whisper-large-v3")
//   - file: The audio file content as io.Reader
//   - filename: The name of the audio file
//   - params: Optional parameters for transcription
//
// Returns transcription text with optional timestamps
func (c *MistralClient) Transcribe(model string, file io.Reader, filename string, params *TranscriptionRequest) (*TranscriptionResponse, error) {
	if params == nil {
		params = &TranscriptionRequest{}
	}

	params.Model = model
	params.File = file
	params.Filename = filename

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	part, err := writer.CreateFormFile("file", filepath.Base(filename))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Add model
	if err := writer.WriteField("model", model); err != nil {
		return nil, fmt.Errorf("failed to write model field: %w", err)
	}

	// Add optional fields
	if params.Language != nil {
		if err := writer.WriteField("language", *params.Language); err != nil {
			return nil, fmt.Errorf("failed to write language field: %w", err)
		}
	}

	if params.Temperature != nil {
		if err := writer.WriteField("temperature", fmt.Sprintf("%f", *params.Temperature)); err != nil {
			return nil, fmt.Errorf("failed to write temperature field: %w", err)
		}
	}

	if len(params.TimestampGranularities) > 0 {
		for _, gran := range params.TimestampGranularities {
			if err := writer.WriteField("timestamp_granularities[]", string(gran)); err != nil {
				return nil, fmt.Errorf("failed to write timestamp_granularities field: %w", err)
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, c.endpoint+"/v1/audio/transcriptions", body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request with retry logic
	client := &http.Client{Timeout: c.timeout}
	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				continue
			}
			return nil, NewMistralConnectionError(err.Error())
		}
		defer resp.Body.Close()

		// Check if we should retry
		if retryStatusCodes[resp.StatusCode] && attempt < c.maxRetries {
			lastErr = fmt.Errorf("received retry status code: %d", resp.StatusCode)
			continue
		}

		// Read response body
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Check for errors
		if resp.StatusCode >= 400 {
			return nil, NewMistralAPIError(string(respBody), resp.StatusCode, resp.Header)
		}

		// Parse response
		var result TranscriptionResponse
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}

		return &result, nil
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

// TranscribeFromURL transcribes an audio file from a URL
func (c *MistralClient) TranscribeFromURL(model string, fileURL string, params *TranscriptionRequest) (*TranscriptionResponse, error) {
	if params == nil {
		params = &TranscriptionRequest{}
	}

	params.Model = model
	params.FileURL = &fileURL

	reqMap := map[string]interface{}{
		"model":    model,
		"file_url": fileURL,
	}

	// Add optional parameters
	if params.Language != nil {
		reqMap["language"] = params.Language
	}
	if params.Temperature != nil {
		reqMap["temperature"] = params.Temperature
	}
	if len(params.TimestampGranularities) > 0 {
		reqMap["timestamp_granularities"] = params.TimestampGranularities
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/audio/transcriptions", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var transcriptionResponse TranscriptionResponse
	err = mapToStruct(respData, &transcriptionResponse)
	if err != nil {
		return nil, err
	}

	return &transcriptionResponse, nil
}

// TranscribeFromFileID transcribes an audio file that was previously uploaded
func (c *MistralClient) TranscribeFromFileID(model string, fileID string, params *TranscriptionRequest) (*TranscriptionResponse, error) {
	if params == nil {
		params = &TranscriptionRequest{}
	}

	params.Model = model
	params.FileID = &fileID

	reqMap := map[string]interface{}{
		"model":   model,
		"file_id": fileID,
	}

	// Add optional parameters
	if params.Language != nil {
		reqMap["language"] = params.Language
	}
	if params.Temperature != nil {
		reqMap["temperature"] = params.Temperature
	}
	if len(params.TimestampGranularities) > 0 {
		reqMap["timestamp_granularities"] = params.TimestampGranularities
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/audio/transcriptions", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var transcriptionResponse TranscriptionResponse
	err = mapToStruct(respData, &transcriptionResponse)
	if err != nil {
		return nil, err
	}

	return &transcriptionResponse, nil
}
