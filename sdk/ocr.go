package sdk

import (
	"fmt"
	"net/http"
)

// OCRDocument represents a document for OCR processing
type OCRDocument struct {
	URL    *string `json:"url,omitempty"`     // URL of the document
	Base64 *string `json:"base64,omitempty"`  // Base64-encoded document
	FileID *string `json:"file_id,omitempty"` // ID of uploaded file
}

// OCRRequest represents a request for OCR processing
type OCRRequest struct {
	Model                    *string         `json:"model,omitempty"`
	ID                       *string         `json:"id,omitempty"`
	Document                 OCRDocument     `json:"document"`
	Pages                    []int           `json:"pages,omitempty"`
	IncludeImageBase64       *bool           `json:"include_image_base64,omitempty"`
	ImageLimit               *int            `json:"image_limit,omitempty"`
	ImageMinSize             *int            `json:"image_min_size,omitempty"`
	BboxAnnotationFormat     *ResponseFormat `json:"bbox_annotation_format,omitempty"`
	DocumentAnnotationFormat *ResponseFormat `json:"document_annotation_format,omitempty"`
}

// OCRPageDimensions represents the dimensions of a page
type OCRPageDimensions struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// OCRImageObject represents an extracted image from the document
type OCRImageObject struct {
	ImageURL    *string   `json:"image_url,omitempty"`
	ImageBase64 *string   `json:"image_base64,omitempty"`
	BBox        []float64 `json:"bbox,omitempty"` // [x, y, width, height]
}

// OCRPageObject represents a processed page
type OCRPageObject struct {
	PageNumber int                `json:"page_number"`
	Dimensions *OCRPageDimensions `json:"dimensions,omitempty"`
	Text       string             `json:"text"`
	Images     []OCRImageObject   `json:"images,omitempty"`
}

// OCRUsageInfo represents usage information for OCR
type OCRUsageInfo struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

// OCRResponse represents the response from OCR processing
type OCRResponse struct {
	ID     string          `json:"id"`
	Object string          `json:"object"`
	Model  string          `json:"model"`
	Pages  []OCRPageObject `json:"pages"`
	Usage  *OCRUsageInfo   `json:"usage,omitempty"`
}

// ProcessOCR processes a document with OCR
//
// Parameters:
//   - model: The model to use for OCR (e.g., "pixtral-12b-2409")
//   - document: The document to process (URL, base64, or file ID)
//   - params: Optional parameters for OCR processing
//
// Returns OCR results with extracted text and images
func (c *MistralClient) ProcessOCR(model string, document OCRDocument, params *OCRRequest) (*OCRResponse, error) {
	if params == nil {
		params = &OCRRequest{}
	}

	// Set required fields
	params.Model = &model
	params.Document = document

	reqMap := map[string]interface{}{
		"model":    params.Model,
		"document": params.Document,
	}

	// Add optional parameters
	if params.ID != nil {
		reqMap["id"] = params.ID
	}
	if len(params.Pages) > 0 {
		reqMap["pages"] = params.Pages
	}
	if params.IncludeImageBase64 != nil {
		reqMap["include_image_base64"] = params.IncludeImageBase64
	}
	if params.ImageLimit != nil {
		reqMap["image_limit"] = params.ImageLimit
	}
	if params.ImageMinSize != nil {
		reqMap["image_min_size"] = params.ImageMinSize
	}
	if params.BboxAnnotationFormat != nil {
		reqMap["bbox_annotation_format"] = map[string]interface{}{
			"type": *params.BboxAnnotationFormat,
		}
	}
	if params.DocumentAnnotationFormat != nil {
		reqMap["document_annotation_format"] = map[string]interface{}{
			"type": *params.DocumentAnnotationFormat,
		}
	}

	response, err := c.request(http.MethodPost, reqMap, "v1/ocr", false, nil)
	if err != nil {
		return nil, err
	}

	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}

	var ocrResponse OCRResponse
	err = mapToStruct(respData, &ocrResponse)
	if err != nil {
		return nil, err
	}

	return &ocrResponse, nil
}

// ProcessOCRFromURL is a convenience method for processing a document from a URL
func (c *MistralClient) ProcessOCRFromURL(model string, url string, params *OCRRequest) (*OCRResponse, error) {
	document := OCRDocument{
		URL: &url,
	}
	return c.ProcessOCR(model, document, params)
}

// ProcessOCRFromBase64 is a convenience method for processing a base64-encoded document
func (c *MistralClient) ProcessOCRFromBase64(model string, base64Data string, params *OCRRequest) (*OCRResponse, error) {
	document := OCRDocument{
		Base64: &base64Data,
	}
	return c.ProcessOCR(model, document, params)
}

// ProcessOCRFromFileID is a convenience method for processing an uploaded file
func (c *MistralClient) ProcessOCRFromFileID(model string, fileID string, params *OCRRequest) (*OCRResponse, error) {
	document := OCRDocument{
		FileID: &fileID,
	}
	return c.ProcessOCR(model, document, params)
}
