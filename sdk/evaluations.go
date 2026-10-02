package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type PipelineKind string

const (
	PipelineKindDetection  PipelineKind = "detection"
	PipelineKindModeration PipelineKind = "moderation"
	PipelineKindJudge      PipelineKind = "judge"
	PipelineKindExport     PipelineKind = "export"
)

type SourceKind string

const (
	SourceKindSpan   SourceKind = "span"
	SourceKindLog    SourceKind = "log"
	SourceKindMetric SourceKind = "metric"
)

type PipelineConfigScope string

const (
	PipelineConfigScopeWorkspace PipelineConfigScope = "workspace"
	PipelineConfigScopeShared    PipelineConfigScope = "shared"
)

type PipelineConfigSelector struct {
	SourceKind SourceKind `json:"source_kind"`
	Filter     *string    `json:"filter,omitempty"`
}
type DetectionPattern struct {
	Name       string `json:"name"`
	Regex      string `json:"regex"`
	Category   string `json:"category"`
	Confidence string `json:"confidence"`
}
type DetectionDefinition struct {
	TargetAttributes []string           `json:"target_attributes"`
	Patterns         []DetectionPattern `json:"patterns"`
}
type ModerationDefinition struct {
	Model            string   `json:"model,omitempty"`
	TargetAttributes []string `json:"target_attributes,omitempty"`
}
type JudgeDefinition struct {
	Slug    string            `json:"slug"`
	Mapping map[string]string `json:"mapping,omitempty"`
	// Deprecated: upstream judges now use Slug and Mapping.
	Model  string `json:"-"`
	Prompt string `json:"-"`
}
type PipelineConfigHeader struct {
	Value *string `json:"value,omitempty"`
}
type OTLPDestination struct {
	Protocol string                          `json:"protocol"`
	Endpoint string                          `json:"endpoint"`
	Insecure bool                            `json:"insecure"`
	Headers  map[string]PipelineConfigHeader `json:"headers,omitempty"`
}
type ExportDefinition struct {
	Destination OTLPDestination `json:"destination"`
}

// PipelineConfigDefinition accepts DetectionDefinition, ModerationDefinition,
// JudgeDefinition, ExportDefinition, or a map for a future definition variant.
type PipelineConfigDefinition = any

type CreatePipelineConfigRequest struct {
	PipelineKind PipelineKind             `json:"pipeline_kind"`
	Selectors    []PipelineConfigSelector `json:"selectors"`
	// Deprecated: set Definition; a single-element list is accepted for migration.
	Definitions []PipelineConfigDefinition `json:"definitions,omitempty"`
	Definition  PipelineConfigDefinition   `json:"definition"`
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	Slug        *string                    `json:"-"`
	Group       *string                    `json:"-"`
	Enabled     *bool                      `json:"enabled,omitempty"`
}
type UpdatePipelineConfigRequest struct {
	PipelineKind PipelineKind             `json:"pipeline_kind"`
	Selectors    []PipelineConfigSelector `json:"selectors"`
	// Deprecated: set Definition; a single-element list is accepted for migration.
	Definitions []PipelineConfigDefinition `json:"definitions,omitempty"`
	Definition  PipelineConfigDefinition   `json:"definition"`
	Name        string                     `json:"name"`
	Enabled     bool                       `json:"enabled"`
	Description *string                    `json:"description,omitempty"`
	Slug        *string                    `json:"-"`
	Group       *string                    `json:"-"`
}
type PipelineConfig struct {
	ID             string                     `json:"id"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	DeletedAt      *time.Time                 `json:"deleted_at"`
	Scope          PipelineConfigScope        `json:"scope"`
	WorkspaceID    string                     `json:"workspace_id"`
	Name           string                     `json:"name"`
	PipelineKind   PipelineKind               `json:"pipeline_kind"`
	Selectors      []PipelineConfigSelector   `json:"selectors"`
	Enabled        bool                       `json:"enabled"`
	DefinitionHash string                     `json:"-"`
	Definitions    []PipelineConfigDefinition `json:"definitions,omitempty"`
	Definition     PipelineConfigDefinition   `json:"definition"`
	Description    *string                    `json:"description,omitempty"`
	Slug           *string                    `json:"-"`
	Group          *string                    `json:"-"`
}
type PaginatedResultPipelineConfig struct {
	Count    int              `json:"count"`
	Results  []PipelineConfig `json:"results,omitempty"`
	Next     *string          `json:"next,omitempty"`
	Previous *string          `json:"previous,omitempty"`
}
type PipelineConfigsResponse struct {
	PipelineConfigs PaginatedResultPipelineConfig `json:"pipeline_configs"`
}
type ListPipelineConfigsParams struct {
	PipelineKind *PipelineKind
	Group        *string
	Enabled      *bool
	PageSize     *int
	Page         *int
	Q            *string
}

func (c *MistralClient) CreatePipelineConfig(req *CreatePipelineConfigRequest) (*PipelineConfig, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	copy := *req
	if copy.Enabled == nil {
		enabled := true
		copy.Enabled = &enabled
	}
	return requestTyped[PipelineConfig](c, http.MethodPost, &copy, "v1/observability/pipeline-configs")
}
func (c *MistralClient) ListPipelineConfigs(params *ListPipelineConfigsParams) (*PipelineConfigsResponse, error) {
	if params == nil {
		params = &ListPipelineConfigsParams{}
	}
	page, size := 1, 50
	if params.Page != nil {
		page = *params.Page
	}
	if params.PageSize != nil {
		size = *params.PageSize
	}
	values := map[string]any{"page": page, "page_size": size, "enabled": params.Enabled, "q": params.Q}
	if params.PipelineKind != nil {
		values["pipeline_kind"] = string(*params.PipelineKind)
	}
	return requestTyped[PipelineConfigsResponse](c, http.MethodGet, nil, appendQuery("v1/observability/pipeline-configs", queryWithOptionalValues(values)))
}
func (c *MistralClient) GetPipelineConfig(id string) (*PipelineConfig, error) {
	return requestTyped[PipelineConfig](c, http.MethodGet, nil, "v1/observability/pipeline-configs/"+url.PathEscape(id))
}
func (c *MistralClient) UpdatePipelineConfig(id string, req *UpdatePipelineConfigRequest) (*PipelineConfig, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[PipelineConfig](c, http.MethodPut, req, "v1/observability/pipeline-configs/"+url.PathEscape(id))
}
func (c *MistralClient) DeletePipelineConfig(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/observability/pipeline-configs/"+url.PathEscape(id), false, nil)
	return err
}

func (v ModerationDefinition) MarshalJSON() ([]byte, error) {
	type wire ModerationDefinition
	if v.Model == "" {
		v.Model = "mistral-moderation-latest"
	}
	return json.Marshal(wire(v))
}

func (v CreatePipelineConfigRequest) MarshalJSON() ([]byte, error) {
	type wire CreatePipelineConfigRequest
	if v.Definition == nil {
		if len(v.Definitions) != 1 {
			return nil, fmt.Errorf("pipeline config requires one definition; use CreatePipeline for multiple judges")
		}
		v.Definition = v.Definitions[0]
	}
	v.Definitions = nil
	return json.Marshal(wire(v))
}
func (v UpdatePipelineConfigRequest) MarshalJSON() ([]byte, error) {
	type wire UpdatePipelineConfigRequest
	if v.Definition == nil {
		if len(v.Definitions) != 1 {
			return nil, fmt.Errorf("pipeline config requires one definition; use UpdatePipeline for multiple judges")
		}
		v.Definition = v.Definitions[0]
	}
	v.Definitions = nil
	return json.Marshal(wire(v))
}
