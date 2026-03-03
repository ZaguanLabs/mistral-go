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

// Document represents a document in a library
type Document struct {
	ID                          string         `json:"id"`
	Object                      string         `json:"object,omitempty"`
	LibraryID                   string         `json:"library_id"`
	Hash                        *string        `json:"hash"`
	MimeType                    *string        `json:"mime_type"`
	Extension                   *string        `json:"extension"`
	Size                        *int64         `json:"size"`
	Name                        string         `json:"name"`
	CreatedAt                   int64          `json:"created_at"`
	ProcessingStatus            string         `json:"processing_status"`
	UploadedByID                *string        `json:"uploaded_by_id"`
	UploadedByType              string         `json:"uploaded_by_type"`
	TokensProcessingTotal       int            `json:"tokens_processing_total"`
	Summary                     *string        `json:"summary,omitempty"`
	LastProcessedAt             *int64         `json:"last_processed_at,omitempty"`
	NumberOfPages               *int           `json:"number_of_pages,omitempty"`
	TokensProcessingMainContent *int           `json:"tokens_processing_main_content,omitempty"`
	TokensProcessingSummary     *int           `json:"tokens_processing_summary,omitempty"`
	URL                         *string        `json:"url,omitempty"`
	Attributes                  map[string]any `json:"attributes,omitempty"`
}

// DocumentListResponse represents a list of documents
type DocumentListResponse struct {
	Object string     `json:"object"`
	Data   []Document `json:"data"`
	Total  int        `json:"total"`
}

// DocumentUploadResponse represents the response from uploading a document
type DocumentUploadResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Status string `json:"status"`
}

// UpdateDocumentRequest represents a request to update a document
type UpdateDocumentRequest struct {
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Attributes  map[string]any `json:"attributes,omitempty"`
}

// DeleteDocumentResponse represents the response from deleting a document
type DeleteDocumentResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

// DocumentStatusResponse represents the status of a document
type DocumentStatusResponse struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Error  *string `json:"error,omitempty"`
}

// DocumentSignedURLResponse represents a signed URL string response.
type DocumentSignedURLResponse struct {
	URL string `json:"url"`
}

// DocumentTextContent represents extracted text content for a document.
type DocumentTextContent struct {
	Content string `json:"content"`
}

// ListDocumentsParams represents query parameters for listing documents.
type ListDocumentsParams struct {
	Search            *string `json:"search,omitempty"`
	PageSize          *int    `json:"page_size,omitempty"`
	Page              *int    `json:"page,omitempty"`
	FiltersAttributes *string `json:"filters_attributes,omitempty"`
	SortBy            *string `json:"sort_by,omitempty"`
	SortOrder         *string `json:"sort_order,omitempty"`
}

// ListDocuments lists all documents in a library
func (c *MistralClient) ListDocuments(libraryID string, page int) (*DocumentListResponse, error) {
	return c.ListDocumentsWithParams(libraryID, &ListDocumentsParams{Page: &page})
}

// ListDocumentsWithParams lists documents in a library with filters.
func (c *MistralClient) ListDocumentsWithParams(libraryID string, params *ListDocumentsParams) (*DocumentListResponse, error) {
	if params == nil {
		params = &ListDocumentsParams{}
	}

	query := url.Values{}
	if params.Search != nil {
		query.Add("search", *params.Search)
	}
	if params.PageSize != nil {
		query.Add("page_size", fmt.Sprintf("%d", *params.PageSize))
	}
	if params.Page != nil {
		query.Add("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.FiltersAttributes != nil {
		query.Add("filters_attributes", *params.FiltersAttributes)
	}
	if params.SortBy != nil {
		query.Add("sort_by", *params.SortBy)
	}
	if params.SortOrder != nil {
		query.Add("sort_order", *params.SortOrder)
	}

	path := fmt.Sprintf("v1/libraries/%s/documents", libraryID)
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	response, err := c.request(http.MethodGet, nil, path, false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var listResponse DocumentListResponse
	err = mapToStruct(respData, &listResponse)
	if err != nil {
		return nil, err
	}

	return &listResponse, nil
}

// UploadDocument uploads a document to a library
func (c *MistralClient) UploadDocument(libraryID string, file io.Reader, filename string) (*DocumentUploadResponse, error) {
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

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, c.endpoint+fmt.Sprintf("/v1/libraries/%s/documents", libraryID), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, NewMistralConnectionError(err.Error())
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		return nil, NewMistralAPIError(string(respBody), resp.StatusCode, resp.Header)
	}

	// Parse response
	var result DocumentUploadResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

// GetDocument retrieves a specific document
func (c *MistralClient) GetDocument(libraryID, documentID string) (*Document, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/documents/%s", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var document Document
	err = mapToStruct(respData, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// UpdateDocument updates a document
func (c *MistralClient) UpdateDocument(libraryID, documentID string, req *UpdateDocumentRequest) (*Document, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}

	reqMap := make(map[string]interface{})

	if req.Name != nil {
		reqMap["name"] = *req.Name
	}
	if req.Description != nil {
		reqMap["description"] = *req.Description
	}
	if req.Attributes != nil {
		reqMap["attributes"] = req.Attributes
	}

	response, err := c.request(http.MethodPatch, reqMap, fmt.Sprintf("v1/libraries/%s/documents/%s", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var document Document
	err = mapToStruct(respData, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// DeleteDocument deletes a document
func (c *MistralClient) DeleteDocument(libraryID, documentID string) (*DeleteDocumentResponse, error) {
	response, err := c.request(http.MethodDelete, nil, fmt.Sprintf("v1/libraries/%s/documents/%s", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var deleteResponse DeleteDocumentResponse
	err = mapToStruct(respData, &deleteResponse)
	if err != nil {
		return nil, err
	}

	return &deleteResponse, nil
}

// GetDocumentStatus retrieves the processing status of a document
func (c *MistralClient) GetDocumentStatus(libraryID, documentID string) (*DocumentStatusResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/documents/%s/status", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var statusResponse DocumentStatusResponse
	err = mapToStruct(respData, &statusResponse)
	if err != nil {
		return nil, err
	}

	return &statusResponse, nil
}

// GetDocumentTextContent retrieves extracted text content for a document.
func (c *MistralClient) GetDocumentTextContent(libraryID, documentID string) (*DocumentTextContent, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/documents/%s/text_content", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var textContent DocumentTextContent
	err = mapToStruct(respData, &textContent)
	if err != nil {
		return nil, err
	}

	return &textContent, nil
}

// GetDocumentSignedURL retrieves a signed URL for the document binary.
func (c *MistralClient) GetDocumentSignedURL(libraryID, documentID string) (*DocumentSignedURLResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/documents/%s/signed-url", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var signedURL DocumentSignedURLResponse
	err = mapToStruct(respData, &signedURL)
	if err != nil {
		return nil, err
	}

	return &signedURL, nil
}

// GetDocumentExtractedTextSignedURL retrieves a signed URL for OCR extracted text.
func (c *MistralClient) GetDocumentExtractedTextSignedURL(libraryID, documentID string) (*DocumentSignedURLResponse, error) {
	response, err := c.request(http.MethodGet, nil, fmt.Sprintf("v1/libraries/%s/documents/%s/extracted-text-signed-url", libraryID, documentID), false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var signedURL DocumentSignedURLResponse
	err = mapToStruct(respData, &signedURL)
	if err != nil {
		return nil, err
	}

	return &signedURL, nil
}

// ReprocessDocument requests document reprocessing.
func (c *MistralClient) ReprocessDocument(libraryID, documentID string) error {
	_, err := c.request(http.MethodPost, map[string]interface{}{}, fmt.Sprintf("v1/libraries/%s/documents/%s/reprocess", libraryID, documentID), false, nil)
	return err
}
