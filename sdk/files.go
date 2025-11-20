package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
)

// FilePurpose represents the purpose of a file upload
type FilePurpose string

const (
	FilePurposeFineTune   FilePurpose = "fine-tune"
	FilePurposeBatch      FilePurpose = "batch"
	FilePurposeAssistants FilePurpose = "assistants"
)

// SampleType represents the type of sample in a file
type SampleType string

const (
	SampleTypePretrain SampleType = "pretrain"
	SampleTypeInstruct SampleType = "instruct"
)

// Source represents the source of a file
type Source string

const (
	SourceUpload     Source = "upload"
	SourceRepository Source = "repository"
)

// FileSchema represents a file object
type FileSchema struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	Bytes      int64       `json:"bytes"`
	CreatedAt  int64       `json:"created_at"`
	Filename   string      `json:"filename"`
	Purpose    FilePurpose `json:"purpose"`
	SampleType *SampleType `json:"sample_type,omitempty"`
	NumLines   *int        `json:"num_lines,omitempty"`
	Source     *Source     `json:"source,omitempty"`
	Deleted    *bool       `json:"deleted,omitempty"`
	Signature  *string     `json:"signature,omitempty"`
}

// UploadFileOut represents the response from uploading a file
type UploadFileOut struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	Bytes      int64       `json:"bytes"`
	CreatedAt  int64       `json:"created_at"`
	Filename   string      `json:"filename"`
	Purpose    FilePurpose `json:"purpose"`
	SampleType *SampleType `json:"sample_type,omitempty"`
	NumLines   *int        `json:"num_lines,omitempty"`
	Source     *Source     `json:"source,omitempty"`
	Signature  *string     `json:"signature,omitempty"`
}

// ListFilesOut represents the response from listing files
type ListFilesOut struct {
	Data   []FileSchema `json:"data"`
	Object string       `json:"object"`
	Total  int          `json:"total"`
}

// RetrieveFileOut represents the response from retrieving a file
type RetrieveFileOut struct {
	ID         string      `json:"id"`
	Object     string      `json:"object"`
	Bytes      int64       `json:"bytes"`
	CreatedAt  int64       `json:"created_at"`
	Filename   string      `json:"filename"`
	Purpose    FilePurpose `json:"purpose"`
	SampleType *SampleType `json:"sample_type,omitempty"`
	NumLines   *int        `json:"num_lines,omitempty"`
	Source     *Source     `json:"source,omitempty"`
	Deleted    *bool       `json:"deleted,omitempty"`
	Signature  *string     `json:"signature,omitempty"`
}

// DeleteFileOut represents the response from deleting a file
type DeleteFileOut struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// FileSignedURL represents a signed URL for file access
type FileSignedURL struct {
	URL string `json:"url"`
}

// ListFilesParams represents parameters for listing files
type ListFilesParams struct {
	Page       *int         `json:"page,omitempty"`
	PageSize   *int         `json:"page_size,omitempty"`
	SampleType []SampleType `json:"sample_type,omitempty"`
	Source     []Source     `json:"source,omitempty"`
	Search     *string      `json:"search,omitempty"`
	Purpose    *FilePurpose `json:"purpose,omitempty"`
}

// UploadFile uploads a file that can be used across various endpoints.
// The size of individual files can be a maximum of 512 MB.
// The Fine-tuning API only supports .jsonl files.
//
// Parameters:
//   - file: The file content as io.Reader
//   - filename: The name of the file
//   - purpose: The intended purpose of the file (fine-tune, batch, etc.)
func (c *MistralClient) UploadFile(file io.Reader, filename string, purpose FilePurpose) (*UploadFileOut, error) {
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

	// Add purpose if specified
	if purpose != "" {
		if err := writer.WriteField("purpose", string(purpose)); err != nil {
			return nil, fmt.Errorf("failed to write purpose field: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, c.endpoint+"/v1/files", body)
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
		var result UploadFileOut
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}

		return &result, nil
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

// ListFiles returns a list of files that belong to the user's organization.
func (c *MistralClient) ListFiles(params *ListFilesParams) (*ListFilesOut, error) {
	if params == nil {
		params = &ListFilesParams{}
	}

	// Build query parameters
	queryParams := url.Values{}
	if params.Page != nil {
		queryParams.Add("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%d", *params.PageSize))
	}
	if params.Search != nil {
		queryParams.Add("search", *params.Search)
	}
	if params.Purpose != nil {
		queryParams.Add("purpose", string(*params.Purpose))
	}
	for _, st := range params.SampleType {
		queryParams.Add("sample_type", string(st))
	}
	for _, src := range params.Source {
		queryParams.Add("source", string(src))
	}

	urlStr := c.endpoint + "/v1/files"
	if len(queryParams) > 0 {
		urlStr += "?" + queryParams.Encode()
	}

	response, err := c.request(http.MethodGet, nil, "v1/files?"+queryParams.Encode(), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listFilesOut ListFilesOut
	err = mapToStruct(respData, &listFilesOut)
	if err != nil {
		return nil, err
	}

	return &listFilesOut, nil
}

// RetrieveFile retrieves information about a specific file.
func (c *MistralClient) RetrieveFile(fileID string) (*RetrieveFileOut, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/files/%s", fileID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var retrieveFileOut RetrieveFileOut
	err = mapToStruct(respData, &retrieveFileOut)
	if err != nil {
		return nil, err
	}

	return &retrieveFileOut, nil
}

// DeleteFile deletes a file.
func (c *MistralClient) DeleteFile(fileID string) (*DeleteFileOut, error) {
	response, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/files/%s", fileID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var deleteFileOut DeleteFileOut
	err = mapToStruct(respData, &deleteFileOut)
	if err != nil {
		return nil, err
	}

	return &deleteFileOut, nil
}

// DownloadFile downloads the content of a file.
// Returns the file content as a byte slice.
func (c *MistralClient) DownloadFile(fileID string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.endpoint+fmt.Sprintf("/v1/files/%s/content", fileID), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/octet-stream")

	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, NewMistralConnectionError(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, NewMistralAPIError(string(body), resp.StatusCode, resp.Header)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	return content, nil
}

// GetSignedURL retrieves a signed URL for accessing a file.
// The URL will be valid for the specified number of hours (default: 24h).
func (c *MistralClient) GetSignedURL(fileID string, expiryHours *int) (*FileSignedURL, error) {
	queryParams := url.Values{}
	if expiryHours != nil {
		queryParams.Add("expiry", fmt.Sprintf("%d", *expiryHours))
	}

	path := fmt.Sprintf("v1/files/%s/url", fileID)
	if len(queryParams) > 0 {
		path += "?" + queryParams.Encode()
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var fileSignedURL FileSignedURL
	err = mapToStruct(respData, &fileSignedURL)
	if err != nil {
		return nil, err
	}

	return &fileSignedURL, nil
}
