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
	Model                       *string                         `json:"model,omitempty"`
	ID                          *string                         `json:"id,omitempty"`
	Document                    OCRDocument                     `json:"document"`
	Pages                       any                             `json:"pages,omitempty"`
	IncludeImageBase64          *bool                           `json:"include_image_base64,omitempty"`
	ImageLimit                  *int                            `json:"image_limit,omitempty"`
	ImageMinSize                *int                            `json:"image_min_size,omitempty"`
	BboxAnnotationFormat        *ResponseFormat                 `json:"bbox_annotation_format,omitempty"`
	DocumentAnnotationFormat    *ResponseFormat                 `json:"document_annotation_format,omitempty"`
	DocumentAnnotationPrompt    *string                         `json:"document_annotation_prompt,omitempty"`
	TableFormat                 *OCRTableFormat                 `json:"table_format,omitempty"`
	ExtractHeader               *bool                           `json:"extract_header,omitempty"`
	ExtractFooter               *bool                           `json:"extract_footer,omitempty"`
	IncludeBlocks               *bool                           `json:"include_blocks,omitempty"`
	ConfidenceScoresGranularity *OCRConfidenceScoresGranularity `json:"confidence_scores_granularity,omitempty"`
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

// OCRTableFormat represents the format of extracted tables
type OCRTableFormat string

const (
	OCRTableFormatMarkdown OCRTableFormat = "markdown"
	OCRTableFormatHTML     OCRTableFormat = "html"
)

type OCRConfidenceScoresGranularity string

const (
	OCRConfidenceScoresWord  OCRConfidenceScoresGranularity = "word"
	OCRConfidenceScoresPage  OCRConfidenceScoresGranularity = "page"
	OCRConfidenceScoresBlock OCRConfidenceScoresGranularity = "block"
)

type OCRBlockConfidenceScores struct {
	AverageContentConfidenceScore *float64 `json:"average_content_confidence_score,omitempty"`
	MinimumContentConfidenceScore *float64 `json:"minimum_content_confidence_score,omitempty"`
	BlockTypeConfidenceScore      *float64 `json:"block_type_confidence_score,omitempty"`
}

type OCRBlock struct {
	TopLeftX         int                       `json:"top_left_x,omitempty"`
	TopLeftY         int                       `json:"top_left_y,omitempty"`
	BottomRightX     int                       `json:"bottom_right_x,omitempty"`
	BottomRightY     int                       `json:"bottom_right_y,omitempty"`
	Content          string                    `json:"content,omitempty"`
	ConfidenceScores *OCRBlockConfidenceScores `json:"confidence_scores,omitempty"`
	Type             string                    `json:"type"`
	ImageID          *string                   `json:"image_id,omitempty"`
	TableID          *string                   `json:"table_id,omitempty"`
	Raw              any                       `json:"raw,omitempty"`
	IsUnknown        *bool                     `json:"is_unknown,omitempty"`
}

// OCRTableObject represents an extracted table from the document
type OCRTableObject struct {
	ID      string         `json:"id"`      // Table ID for extracted table in a page
	Content string         `json:"content"` // Content of the table in the given format
	Format  OCRTableFormat `json:"format"`  // Format of the table
}

// OCRPageObject represents a processed page
type OCRPageObject struct {
	PageNumber int                `json:"page_number"`
	Dimensions *OCRPageDimensions `json:"dimensions,omitempty"`
	Text       string             `json:"text"`
	Images     []OCRImageObject   `json:"images,omitempty"`
	Tables     []OCRTableObject   `json:"tables,omitempty"`     // List of all extracted tables in the page
	Hyperlinks []string           `json:"hyperlinks,omitempty"` // List of all hyperlinks in the page
	Header     *string            `json:"header,omitempty"`     // Header of the page
	Footer     *string            `json:"footer,omitempty"`     // Footer of the page
	Blocks     []OCRBlock         `json:"blocks,omitempty"`
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
	if params.Pages != nil {
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
	if params.DocumentAnnotationPrompt != nil {
		reqMap["document_annotation_prompt"] = params.DocumentAnnotationPrompt
	}
	if params.TableFormat != nil {
		reqMap["table_format"] = params.TableFormat
	}
	if params.ExtractHeader != nil {
		reqMap["extract_header"] = params.ExtractHeader
	}
	if params.ExtractFooter != nil {
		reqMap["extract_footer"] = params.ExtractFooter
	}
	if params.IncludeBlocks != nil {
		reqMap["include_blocks"] = params.IncludeBlocks
	}
	if params.ConfidenceScoresGranularity != nil {
		reqMap["confidence_scores_granularity"] = params.ConfidenceScoresGranularity
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
