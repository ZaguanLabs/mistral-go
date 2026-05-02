package sdk

import (
	"fmt"
	"io"
	"net/http"
)

type SpeechOutputFormat string

const (
	SpeechOutputFormatMP3  SpeechOutputFormat = "mp3"
	SpeechOutputFormatWAV  SpeechOutputFormat = "wav"
	SpeechOutputFormatPCM  SpeechOutputFormat = "pcm"
	SpeechOutputFormatFLAC SpeechOutputFormat = "flac"
)

type SpeechRequest struct {
	Input          string             `json:"input"`
	Model          *string            `json:"model,omitempty"`
	Metadata       map[string]any     `json:"metadata,omitempty"`
	Stream         *bool              `json:"stream,omitempty"`
	VoiceID        *string            `json:"voice_id,omitempty"`
	RefAudio       *string            `json:"ref_audio,omitempty"`
	ResponseFormat SpeechOutputFormat `json:"response_format,omitempty"`
}

type SpeechResponse struct {
	Audio    string         `json:"audio,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type VoiceType string

const (
	VoiceTypeAll    VoiceType = "all"
	VoiceTypeCustom VoiceType = "custom"
	VoiceTypePreset VoiceType = "preset"
)

type Voice struct {
	ID              string         `json:"id,omitempty"`
	Name            string         `json:"name,omitempty"`
	Slug            string         `json:"slug,omitempty"`
	Type            string         `json:"type,omitempty"`
	Languages       []string       `json:"languages,omitempty"`
	Gender          string         `json:"gender,omitempty"`
	Age             *int           `json:"age,omitempty"`
	Tags            []string       `json:"tags,omitempty"`
	Color           string         `json:"color,omitempty"`
	SampleFilename  string         `json:"sample_filename,omitempty"`
	RetentionNotice *int           `json:"retention_notice,omitempty"`
	CreatedAt       *int64         `json:"created_at,omitempty"`
	UpdatedAt       *int64         `json:"updated_at,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}

type VoiceListResponse struct {
	Object string  `json:"object,omitempty"`
	Data   []Voice `json:"data,omitempty"`
	Total  *int    `json:"total,omitempty"`
}

type VoiceRequest struct {
	Name            string   `json:"name"`
	SampleAudio     string   `json:"sample_audio,omitempty"`
	Slug            *string  `json:"slug,omitempty"`
	Languages       []string `json:"languages,omitempty"`
	Gender          *string  `json:"gender,omitempty"`
	Age             *int     `json:"age,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	Color           *string  `json:"color,omitempty"`
	RetentionNotice *int     `json:"retention_notice,omitempty"`
	SampleFilename  *string  `json:"sample_filename,omitempty"`
}

type UpdateVoiceRequest struct {
	Name            *string  `json:"name,omitempty"`
	Slug            *string  `json:"slug,omitempty"`
	Languages       []string `json:"languages,omitempty"`
	Gender          *string  `json:"gender,omitempty"`
	Age             *int     `json:"age,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	Color           *string  `json:"color,omitempty"`
	RetentionNotice *int     `json:"retention_notice,omitempty"`
}

type ListVoicesParams struct {
	Limit  *int       `json:"limit,omitempty"`
	Offset *int       `json:"offset,omitempty"`
	Type   *VoiceType `json:"type,omitempty"`
}

func (c *MistralClient) Speech(req *SpeechRequest) (*SpeechResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"input":           req.Input,
		"model":           req.Model,
		"metadata":        req.Metadata,
		"stream":          req.Stream,
		"voice_id":        req.VoiceID,
		"ref_audio":       req.RefAudio,
		"response_format": req.ResponseFormat,
	})
	response, err := c.requestMap(http.MethodPost, body, "v1/audio/speech")
	if err != nil {
		return nil, err
	}
	var speech SpeechResponse
	if err := mapToStruct(map[string]interface{}(response), &speech); err != nil {
		return nil, err
	}
	return &speech, nil
}

func (c *MistralClient) SpeechStream(req *SpeechRequest) (<-chan StreamEvent, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	stream := true
	req.Stream = &stream
	body := optionalRequestMap(map[string]any{
		"input":           req.Input,
		"model":           req.Model,
		"metadata":        req.Metadata,
		"stream":          req.Stream,
		"voice_id":        req.VoiceID,
		"ref_audio":       req.RefAudio,
		"response_format": req.ResponseFormat,
	})
	response, err := c.request(http.MethodPost, body, "v1/audio/speech", true, nil)
	if err != nil {
		return nil, err
	}
	bodyStream, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	return parseGenericStream(bodyStream), nil
}

func (c *MistralClient) ListVoices(params *ListVoicesParams) (*VoiceListResponse, error) {
	if params == nil {
		params = &ListVoicesParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"limit":  params.Limit,
		"offset": params.Offset,
	})
	if params.Type != nil {
		typeValue := string(*params.Type)
		typeQuery := queryWithOptionalValues(map[string]any{"type": &typeValue})
		if query == "" {
			query = typeQuery
		} else if typeQuery != "" {
			query += "&" + typeQuery
		}
	}
	response, err := c.requestMap(http.MethodGet, nil, appendQuery("v1/audio/voices", query))
	if err != nil {
		return nil, err
	}
	var voices VoiceListResponse
	if err := mapToStruct(map[string]interface{}(response), &voices); err != nil {
		return nil, err
	}
	return &voices, nil
}

func (c *MistralClient) CreateVoice(req *VoiceRequest) (*Voice, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"name":             req.Name,
		"sample_audio":     req.SampleAudio,
		"slug":             req.Slug,
		"languages":        req.Languages,
		"gender":           req.Gender,
		"age":              req.Age,
		"tags":             req.Tags,
		"color":            req.Color,
		"retention_notice": req.RetentionNotice,
		"sample_filename":  req.SampleFilename,
	})
	response, err := c.requestMap(http.MethodPost, body, "v1/audio/voices")
	if err != nil {
		return nil, err
	}
	var voice Voice
	if err := mapToStruct(map[string]interface{}(response), &voice); err != nil {
		return nil, err
	}
	return &voice, nil
}

func (c *MistralClient) DeleteVoice(voiceID string) (*Voice, error) {
	response, err := c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/audio/voices/%s", voiceID))
	if err != nil {
		return nil, err
	}
	var voice Voice
	if err := mapToStruct(map[string]interface{}(response), &voice); err != nil {
		return nil, err
	}
	return &voice, nil
}

func (c *MistralClient) UpdateVoice(voiceID string, req *UpdateVoiceRequest) (*Voice, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"name":             req.Name,
		"slug":             req.Slug,
		"languages":        req.Languages,
		"gender":           req.Gender,
		"age":              req.Age,
		"tags":             req.Tags,
		"color":            req.Color,
		"retention_notice": req.RetentionNotice,
	})
	response, err := c.requestMap(http.MethodPatch, body, fmt.Sprintf("v1/audio/voices/%s", voiceID))
	if err != nil {
		return nil, err
	}
	var voice Voice
	if err := mapToStruct(map[string]interface{}(response), &voice); err != nil {
		return nil, err
	}
	return &voice, nil
}

func (c *MistralClient) GetVoice(voiceID string) (*Voice, error) {
	response, err := c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/audio/voices/%s", voiceID))
	if err != nil {
		return nil, err
	}
	var voice Voice
	if err := mapToStruct(map[string]interface{}(response), &voice); err != nil {
		return nil, err
	}
	return &voice, nil
}

func (c *MistralClient) GetVoiceSampleAudio(voiceID string) ([]byte, error) {
	return c.requestBytes(http.MethodGet, fmt.Sprintf("v1/audio/voices/%s/sample", voiceID), "audio/wav")
}
