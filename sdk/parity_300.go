package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func enumString[T ~string](value *T) *string {
	if value == nil {
		return nil
	}
	s := string(*value)
	return &s
}

type DeploymentLogType string

const (
	DeploymentLogBuild   DeploymentLogType = "build"
	DeploymentLogRuntime DeploymentLogType = "runtime"
)

type PrincipalType string

const (
	PrincipalTypeAPIKey         PrincipalType = "api_key"
	PrincipalTypeServiceAccount PrincipalType = "service_account"
	PrincipalTypeUser           PrincipalType = "user"
)

type DeploymentSecretBinding struct {
	EnvVarName string `json:"env_var_name"`
	Reference  string `json:"reference"`
}
type DeploymentMistralCloudBackendSpec struct {
	Type           string                    `json:"type"`
	BuildDirectory *string                   `json:"build_directory,omitempty"`
	DockerfilePath *string                   `json:"dockerfile_path,omitempty"`
	Secrets        []DeploymentSecretBinding `json:"secrets,omitempty"`
}

func (v DeploymentMistralCloudBackendSpec) MarshalJSON() ([]byte, error) {
	type wire DeploymentMistralCloudBackendSpec
	if v.Type == "" {
		v.Type = "mistral_cloud"
	}
	return json.Marshal(wire(v))
}

type AuthData struct {
	ClientID     string   `json:"client_id"`
	ClientSecret *string  `json:"client_secret,omitempty"`
	ClientScopes []string `json:"client_scopes,omitempty"`
}

const BuiltInConnectorMistralMCP = "mistral_mcp"

type RRFRetriever struct {
	Retrievers   []any     `json:"retrievers"`
	TopK         *int      `json:"top_k,omitempty"`
	Type         string    `json:"type"`
	Weights      []float64 `json:"weights,omitempty"`
	RankConstant *int      `json:"rank_constant,omitempty"`
	Filter       any       `json:"filter,omitempty"`
}

func (v RRFRetriever) MarshalJSON() ([]byte, error) {
	type wire RRFRetriever
	v.Type = "rrf"
	if v.TopK == nil {
		n := 20
		v.TopK = &n
	}
	if v.RankConstant == nil {
		n := 60
		v.RankConstant = &n
	}
	return json.Marshal(wire(v))
}

type NavigationDirection string

const (
	NavigationNext     NavigationDirection = "next"
	NavigationPrevious NavigationDirection = "previous"
)

type GrepMode string

const (
	GrepPhrase GrepMode = "phrase"
	GrepTerm   GrepMode = "term"
)

type NavigateRequest struct {
	SourceID    string              `json:"source_id"`
	StartOffset int                 `json:"start_offset"`
	EndOffset   int                 `json:"end_offset"`
	Direction   NavigationDirection `json:"direction"`
	TopK        *int                `json:"top_k,omitempty"`
	ContentType *string             `json:"content_type,omitempty"`
}
type ReadRequest struct {
	SourceID    string  `json:"source_id"`
	StartOffset *int    `json:"start_offset,omitempty"`
	EndOffset   *int    `json:"end_offset,omitempty"`
	TopK        *int    `json:"top_k,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
}
type GrepRequest struct {
	SourceID    string    `json:"source_id"`
	Pattern     string    `json:"pattern"`
	Mode        *GrepMode `json:"mode,omitempty"`
	TopK        *int      `json:"top_k,omitempty"`
	ContentType *string   `json:"content_type,omitempty"`
}

func navigationDefaults(top **int, content **string, n int) {
	if *top == nil {
		*top = &n
	}
	if *content == nil {
		s := "content"
		*content = &s
	}
}
func (v NavigateRequest) MarshalJSON() ([]byte, error) {
	type wire NavigateRequest
	navigationDefaults(&v.TopK, &v.ContentType, 1)
	return json.Marshal(wire(v))
}
func (v ReadRequest) MarshalJSON() ([]byte, error) {
	type wire ReadRequest
	navigationDefaults(&v.TopK, &v.ContentType, 20)
	return json.Marshal(wire(v))
}
func (v GrepRequest) MarshalJSON() ([]byte, error) {
	type wire GrepRequest
	navigationDefaults(&v.TopK, &v.ContentType, 5)
	return json.Marshal(wire(v))
}

type NavigationChunkResponse = SearchChunkResponse
type NavigationResponse struct {
	Chunks []NavigationChunkResponse `json:"chunks"`
}

func (c *MistralClient) NavigateManagedIndex(name string, req *NavigateRequest) (*NavigationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[NavigationResponse](c, http.MethodPost, req, managedIndexPath(name)+"/navigate")
}
func (c *MistralClient) ReadManagedIndex(name string, req *ReadRequest) (*NavigationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[NavigationResponse](c, http.MethodPost, req, managedIndexPath(name)+"/read")
}
func (c *MistralClient) GrepManagedIndex(name string, req *GrepRequest) (*NavigationResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[NavigationResponse](c, http.MethodPost, req, managedIndexPath(name)+"/grep")
}
func (c *MistralClient) GetManagedIndexChunk(name, id string) (*NavigationChunkResponse, error) {
	return requestTyped[NavigationChunkResponse](c, http.MethodGet, nil, managedIndexPath(name)+"/chunks/"+url.PathEscape(id))
}

type CreatePipelineRequest struct {
	Name        string                   `json:"name"`
	Selectors   []PipelineConfigSelector `json:"selectors"`
	Definitions []JudgeDefinition        `json:"definitions"`
	Description *string                  `json:"description,omitempty"`
	Enabled     *bool                    `json:"enabled,omitempty"`
}
type UpdatePipelineRequest struct {
	Name        string                   `json:"name"`
	Selectors   []PipelineConfigSelector `json:"selectors"`
	Definitions []JudgeDefinition        `json:"definitions"`
	Enabled     bool                     `json:"enabled"`
	Description *string                  `json:"description,omitempty"`
}
type Pipeline struct {
	ID              string                   `json:"id"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
	WorkspaceID     string                   `json:"workspace_id"`
	Name            string                   `json:"name"`
	Slug            string                   `json:"slug"`
	Description     *string                  `json:"description"`
	Selectors       []PipelineConfigSelector `json:"selectors"`
	Enabled         bool                     `json:"enabled"`
	CreatorID       *string                  `json:"creator_id"`
	PipelineConfigs []PipelineConfig         `json:"pipeline_configs"`
}
type PaginatedResultPipeline struct {
	Count    int        `json:"count"`
	Results  []Pipeline `json:"results,omitempty"`
	Next     *string    `json:"next,omitempty"`
	Previous *string    `json:"previous,omitempty"`
}
type PipelinesResponse struct {
	Pipelines PaginatedResultPipeline `json:"pipelines"`
}
type ListPipelinesParams struct {
	Enabled  *bool
	PageSize *int
	Page     *int
	Q        *string
}

func (c *MistralClient) CreatePipeline(req *CreatePipelineRequest) (*Pipeline, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	copy := *req
	if copy.Enabled == nil {
		b := true
		copy.Enabled = &b
	}
	return requestTyped[Pipeline](c, http.MethodPost, &copy, "v1/observability/pipelines")
}
func (c *MistralClient) ListPipelines(params *ListPipelinesParams) (*PipelinesResponse, error) {
	if params == nil {
		params = &ListPipelinesParams{}
	}
	page, size := 1, 50
	if params.Page != nil {
		page = *params.Page
	}
	if params.PageSize != nil {
		size = *params.PageSize
	}
	return requestTyped[PipelinesResponse](c, http.MethodGet, nil, appendQuery("v1/observability/pipelines", queryWithOptionalValues(map[string]any{"enabled": params.Enabled, "page_size": size, "page": page, "q": params.Q})))
}
func (c *MistralClient) GetPipeline(id string) (*Pipeline, error) {
	return requestTyped[Pipeline](c, http.MethodGet, nil, "v1/observability/pipelines/"+url.PathEscape(id))
}
func (c *MistralClient) UpdatePipeline(id string, req *UpdatePipelineRequest) (*Pipeline, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[Pipeline](c, http.MethodPut, req, "v1/observability/pipelines/"+url.PathEscape(id))
}
func (c *MistralClient) DeletePipeline(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/observability/pipelines/"+url.PathEscape(id), false, nil)
	return err
}

type TelemetrySpanReference struct {
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
}
type SpanDatasetSourceNamespace string

const (
	SpanDatasetSpanAttributes     SpanDatasetSourceNamespace = "span_attributes"
	SpanDatasetResourceAttributes SpanDatasetSourceNamespace = "resource_attributes"
)

type SpanDatasetSourceReference struct {
	Namespace SpanDatasetSourceNamespace `json:"namespace"`
	Key       string                     `json:"key"`
}
type SpanDatasetMapping struct {
	TargetField string                     `json:"target_field"`
	SourceField SpanDatasetSourceReference `json:"source_field"`
}
type SpanDatasetMappingContract struct {
	Mappings []SpanDatasetMapping `json:"mappings"`
	Version  int                  `json:"version"`
}

func (v SpanDatasetMappingContract) MarshalJSON() ([]byte, error) {
	type wire SpanDatasetMappingContract
	v.Version = 1
	return json.Marshal(wire(v))
}

type ImportDatasetFromSpansRequest struct {
	SpanReferences  []TelemetrySpanReference   `json:"span_references"`
	MappingContract SpanDatasetMappingContract `json:"mapping_contract"`
}
type DatasetImportResult struct {
	RequestedRecordCount int `json:"requested_record_count"`
	ImportedRecordCount  int `json:"imported_record_count"`
	SkippedRecordCount   int `json:"skipped_record_count"`
}
type DatasetImportTask struct {
	ID          string               `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	DeletedAt   *time.Time           `json:"deleted_at"`
	CreatorID   string               `json:"creator_id"`
	DatasetID   string               `json:"dataset_id"`
	WorkspaceID string               `json:"workspace_id"`
	Status      string               `json:"status"`
	Progress    *int                 `json:"progress,omitempty"`
	Message     *string              `json:"message,omitempty"`
	Result      *DatasetImportResult `json:"result,omitempty"`
}

func (c *MistralClient) ImportDatasetFromSpans(id string, req *ImportDatasetFromSpansRequest) (*DatasetImportTask, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return requestTyped[DatasetImportTask](c, http.MethodPost, req, "v1/observability/datasets/"+url.PathEscape(id)+"/imports/from-spans")
}
func (c *MistralClient) AggregateSpanEvaluations(params *ObservabilityAggregationParams) (APIResponse, error) {
	return c.aggregateObservabilitySignals("v1/observability/spans/evaluations/aggregate", params)
}
func (c *MistralClient) UnhardenWorkflowDeployment(id string, workspaceID *string) error {
	_, err := c.request(http.MethodPost, nil, appendQuery("v1/workflows/deployments/"+url.PathEscape(id)+"/unharden", queryWithOptionalValues(map[string]any{"workspace_id": workspaceID})), false, nil)
	return err
}
func (c *MistralClient) DeleteConnectorCredentials(id, scope, name string) (APIResponse, error) {
	if scope != "user" && scope != "workspace" && scope != "organization" {
		return nil, fmt.Errorf("invalid consumer scope %q", scope)
	}
	return c.requestMap(http.MethodDelete, nil, "v1/connectors/"+url.PathEscape(id)+"/"+scope+"/credentials/"+url.PathEscape(name))
}

type VoiceGender string

const (
	VoiceGenderFemale  VoiceGender = "female"
	VoiceGenderMale    VoiceGender = "male"
	VoiceGenderNeutral VoiceGender = "neutral"
)

type VoiceListPage struct {
	Data          []Voice `json:"data"`
	NextPageToken *string `json:"next_page_token,omitempty"`
}
type SearchVoicesParams struct {
	PageSize  *int
	PageToken *string
	Type      *VoiceType
	Gender    []VoiceGender
	Language  []string
	Query     *string
}

func (c *MistralClient) SearchVoices(params *SearchVoicesParams) (*VoiceListPage, error) {
	if params == nil {
		params = &SearchVoicesParams{}
	}
	size := 10
	if params.PageSize != nil {
		size = *params.PageSize
	}
	kind := "all"
	if params.Type != nil {
		kind = string(*params.Type)
	}
	genders := make([]string, len(params.Gender))
	for i, v := range params.Gender {
		genders[i] = string(v)
	}
	query := queryWithOptionalValues(map[string]any{"page_size": size, "page_token": params.PageToken, "type": kind, "gender": genders, "language": params.Language, "query": params.Query})
	return requestTyped[VoiceListPage](c, http.MethodGet, nil, appendQuery("v2/audio/voices", query))
}

// OtelFieldDefinition retains the raw attribute key exposed by field discovery.
type OtelFieldDefinition struct {
	Name                  string   `json:"name"`
	Label                 string   `json:"label"`
	Type                  string   `json:"type"`
	SupportedOperators    []string `json:"supported_operators"`
	SupportedAggregations []string `json:"supported_aggregations"`
	Group                 *string  `json:"group,omitempty"`
	SourceAttributeKey    *string  `json:"source_attribute_key,omitempty"`
}
