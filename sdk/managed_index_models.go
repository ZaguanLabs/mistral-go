// Models for the managed-index API in official Python SDK 2.10.1.
package sdk

import (
	"encoding/json"
	"time"
)

type ManagedIndexStatus string

const (
	ManagedIndexStatusProvisioning ManagedIndexStatus = "provisioning"
	ManagedIndexStatusReady        ManagedIndexStatus = "ready"
	ManagedIndexStatusFailed       ManagedIndexStatus = "failed"
	ManagedIndexStatusDeleting     ManagedIndexStatus = "deleting"
)

type DistanceMetric string

const (
	DistanceMetricCosine       DistanceMetric = "cosine"
	DistanceMetricInnerProduct DistanceMetric = "inner_product"
	DistanceMetricL2           DistanceMetric = "l2"
)

type VectorDType string

const (
	VectorDTypeFloat16 VectorDType = "float16"
	VectorDTypeFloat32 VectorDType = "float32"
)

type AcceptedDocumentResponse struct {
	Position int `json:"position"`

	DocumentID string `json:"document_id"`

	ChunkCount int `json:"chunk_count"`

	Status string `json:"status"`
}

func (v AcceptedDocumentResponse) MarshalJSON() ([]byte, error) {
	type wire AcceptedDocumentResponse
	v.Status = "accepted"
	return json.Marshal(wire(v))
}

type And struct {
	Matches []any `json:"matches"`

	Type string `json:"type"`
}

func (v And) MarshalJSON() ([]byte, error) {
	type wire And
	v.Type = "and"
	return json.Marshal(wire(v))
}

type BoolArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v BoolArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire BoolArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "bool[]"
	return json.Marshal(wire(v))
}

type BoolFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v BoolFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire BoolFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "bool"
	return json.Marshal(wire(v))
}

type CreateManagedIndexRequest struct {
	Name string `json:"name"`

	Config ManagedIndexConfig `json:"config"`

	Schema *ManagedIndexFields `json:"schema,omitempty"`
}

type CustomEmbeddingModel struct {
	Name string `json:"name"`

	Dimensions int `json:"dimensions"`

	Type string `json:"type"`

	DType *VectorDType `json:"dtype,omitempty"`

	DistanceMetric *DistanceMetric `json:"distance_metric,omitempty"`
}

func (v CustomEmbeddingModel) MarshalJSON() ([]byte, error) {
	type wire CustomEmbeddingModel
	v.Type = "custom"
	return json.Marshal(wire(v))
}

type DeleteDocumentsRequest struct {
	DocumentIDs []string `json:"document_ids"`
}

type DeleteDocumentsResponse struct {
	DeletedDocumentIDs []string `json:"deleted_document_ids"`

	Missing []string `json:"missing"`
}

type DeleteManagedIndexResponse struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Status ManagedIndexStatus `json:"status"`
}

type DenseVectorFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v DenseVectorFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire DenseVectorFieldDefinition
	if v.Indexed == nil {
		value := true
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "dense_vector"
	return json.Marshal(wire(v))
}

type DocumentFieldErrorResponse struct {
	Location []any `json:"location"`

	Type string `json:"type"`

	Message string `json:"message"`
}

type Equal struct {
	Field string `json:"field"`

	Value any `json:"value"`

	Type string `json:"type"`
}

func (v Equal) MarshalJSON() ([]byte, error) {
	type wire Equal
	v.Type = "equal"
	return json.Marshal(wire(v))
}

type FloatArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v FloatArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire FloatArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "float[]"
	return json.Marshal(wire(v))
}

type FloatFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v FloatFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire FloatFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "float"
	return json.Marshal(wire(v))
}

type In struct {
	Field string `json:"field"`

	Values []any `json:"values"`

	Type string `json:"type"`
}

func (v In) MarshalJSON() ([]byte, error) {
	type wire In
	v.Type = "in"
	return json.Marshal(wire(v))
}

type IngestDocumentsRequest struct {
	Documents []map[string]any `json:"documents"`
}

type IngestDocumentsResponse struct {
	Accepted int `json:"accepted"`

	Rejected int `json:"rejected"`

	Results []any `json:"results"`
}

type IntArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v IntArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire IntArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "int[]"
	return json.Marshal(wire(v))
}

type IntFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v IntFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire IntFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "int"
	return json.Marshal(wire(v))
}

type JSONFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v JSONFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire JSONFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "json"
	return json.Marshal(wire(v))
}

type KeywordArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v KeywordArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire KeywordArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "keyword[]"
	return json.Marshal(wire(v))
}

type KeywordFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v KeywordFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire KeywordFieldDefinition
	if v.Indexed == nil {
		value := true
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "keyword"
	return json.Marshal(wire(v))
}

type KeywordRetriever struct {
	Query string `json:"query"`

	TopK *int `json:"top_k,omitempty"`

	Type string `json:"type"`

	Filter any `json:"filter,omitempty"`

	Field *string `json:"field,omitempty"`
}

func (v KeywordRetriever) MarshalJSON() ([]byte, error) {
	type wire KeywordRetriever
	if v.TopK == nil {
		value := 20
		v.TopK = &value
	}
	v.Type = "keyword"
	return json.Marshal(wire(v))
}

type ListManagedIndexesResponse struct {
	Data []ManagedIndexResponse `json:"data"`

	NextPageToken *string `json:"next_page_token,omitempty"`
}

type LongArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v LongArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire LongArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "long[]"
	return json.Marshal(wire(v))
}

type LongFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v LongFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire LongFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "long"
	return json.Marshal(wire(v))
}

type ManagedIndexConfig struct {
	Embedding any `json:"embedding"`
}

type ManagedIndexFields struct {
	DocumentFields map[string]any `json:"document_fields,omitempty"`

	ChunkFields map[string]any `json:"chunk_fields,omitempty"`
}

type ManagedIndexResponse struct {
	CreatorID string `json:"creator_id"`
	ID        string `json:"id"`

	Name string `json:"name"`

	Status ManagedIndexStatus `json:"status"`

	Schema ManagedIndexFields `json:"schema"`

	Config ManagedIndexConfig `json:"config"`

	StatusMessage *string `json:"status_message,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

type MistralEmbeddingModel struct {
	Name string `json:"name"`

	Dimensions int `json:"dimensions"`

	Type string `json:"type"`

	DType *VectorDType `json:"dtype,omitempty"`

	DistanceMetric *DistanceMetric `json:"distance_metric,omitempty"`
}

func (v MistralEmbeddingModel) MarshalJSON() ([]byte, error) {
	type wire MistralEmbeddingModel
	v.Type = "mistral"
	return json.Marshal(wire(v))
}

type NearestNeighbourRetriever struct {
	MaxCandidates *int `json:"max_candidates,omitempty"`
	TopK          *int `json:"top_k,omitempty"`

	Type string `json:"type"`

	Query *string `json:"query,omitempty"`

	QueryEmbedding []float64 `json:"query_embedding,omitempty"`

	Field *string `json:"field,omitempty"`

	Filter any `json:"filter,omitempty"`
}

func (v NearestNeighbourRetriever) MarshalJSON() ([]byte, error) {
	type wire NearestNeighbourRetriever
	if v.TopK == nil {
		value := 20
		v.TopK = &value
	}
	v.Type = "nearest_neighbour"
	return json.Marshal(wire(v))
}

type Not struct {
	Match any `json:"match"`

	Type string `json:"type"`
}

func (v Not) MarshalJSON() ([]byte, error) {
	type wire Not
	v.Type = "not"
	return json.Marshal(wire(v))
}

type Or struct {
	Matches []any `json:"matches"`

	Type string `json:"type"`
}

func (v Or) MarshalJSON() ([]byte, error) {
	type wire Or
	v.Type = "or"
	return json.Marshal(wire(v))
}

type Range struct {
	Field string `json:"field"`

	Type string `json:"type"`

	Gt any `json:"gt,omitempty"`

	Gte any `json:"gte,omitempty"`

	Lt any `json:"lt,omitempty"`

	Lte any `json:"lte,omitempty"`
}

func (v Range) MarshalJSON() ([]byte, error) {
	type wire Range
	v.Type = "range"
	return json.Marshal(wire(v))
}

type RejectedDocumentResponse struct {
	Position int `json:"position"`

	Errors []DocumentFieldErrorResponse `json:"errors"`

	Status string `json:"status"`
}

func (v RejectedDocumentResponse) MarshalJSON() ([]byte, error) {
	type wire RejectedDocumentResponse
	v.Status = "rejected"
	return json.Marshal(wire(v))
}

type SearchChunkResponse struct {
	ID string `json:"id"`

	SourceID string `json:"source_id"`

	Locator string `json:"locator"`

	StartOffset *int `json:"start_offset"`

	EndOffset *int `json:"end_offset"`

	ChunkType string `json:"chunk_type"`

	Content string `json:"content"`

	ParentRef *string `json:"parent_ref,omitempty"`

	Metadata map[string]any `json:"metadata,omitempty"`

	Extra map[string]json.RawMessage `json:"-"`
}

type SearchHitResponse struct {
	Chunk SearchChunkResponse `json:"chunk"`

	Score float64 `json:"score"`

	IndexName string `json:"index_name"`

	Distance *float64 `json:"distance,omitempty"`
}

type SearchRequest struct {
	Retriever any `json:"retriever"`
}

type SearchResponse struct {
	Hits []SearchHitResponse `json:"hits"`
}

type TextArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v TextArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire TextArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "text[]"
	return json.Marshal(wire(v))
}

type TextFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v TextFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire TextFieldDefinition
	if v.Indexed == nil {
		value := true
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "text"
	return json.Marshal(wire(v))
}

type TimestampArrayFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v TimestampArrayFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire TimestampArrayFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "timestamp[]"
	return json.Marshal(wire(v))
}

type TimestampFieldDefinition struct {
	Indexed *bool `json:"indexed,omitempty"`

	Required *bool `json:"required,omitempty"`

	System *bool `json:"system,omitempty"`

	Default any `json:"default,omitempty"`

	Type string `json:"type"`
}

func (v TimestampFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire TimestampFieldDefinition
	if v.Indexed == nil {
		value := false
		v.Indexed = &value
	}
	if v.Required == nil {
		value := false
		v.Required = &value
	}
	if v.System == nil {
		value := false
		v.System = &value
	}
	v.Type = "timestamp"
	return json.Marshal(wire(v))
}

type UpdateManagedIndexRequest struct {
	Schema ManagedIndexFields `json:"schema"`
}

// UnmarshalJSON retains extension fields returned by managed-index search.
func (v *SearchChunkResponse) UnmarshalJSON(data []byte) error {
	type wire SearchChunkResponse
	var w wire
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	for _, key := range []string{"id", "source_id", "locator", "start_offset", "end_offset", "chunk_type", "content", "parent_ref", "metadata"} {
		delete(fields, key)
	}
	*v = SearchChunkResponse(w)
	v.Extra = fields
	return nil
}
func (v SearchChunkResponse) MarshalJSON() ([]byte, error) {
	type wire SearchChunkResponse
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.Extra {
		if _, exists := fields[key]; !exists {
			fields[key] = value
		}
	}
	return json.Marshal(fields)
}
