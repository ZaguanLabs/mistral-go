package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RegisterIngestionPipelineConfigurationRequest struct {
	Name                string `json:"name"`
	PipelineComposition any    `json:"pipeline_composition"`
}

type UpdateIngestionPipelineRunInfoRequest struct {
	ExecutionTime *float64 `json:"execution_time,omitempty"`
	ChunksCount   *int     `json:"chunks_count,omitempty"`
}

type SearchIndexStatus string

const (
	SearchIndexStatusOnline  SearchIndexStatus = "online"
	SearchIndexStatusOffline SearchIndexStatus = "offline"
)

type RegisterSearchIndexRequest struct {
	Name          string             `json:"name"`
	Index         any                `json:"index"`
	DocumentCount *int               `json:"document_count,omitempty"`
	Status        *SearchIndexStatus `json:"status,omitempty"`
}

type SearchIndexResponse struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	CreatorID     string            `json:"creator_id"`
	DocumentCount int               `json:"document_count"`
	Status        SearchIndexStatus `json:"status"`
	CreatedAt     time.Time         `json:"created_at"`
	ModifiedAt    time.Time         `json:"modified_at"`
	Index         map[string]any    `json:"index"`
}

func (c *MistralClient) ListIngestionPipelineConfigurations() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/rag/ingestion_pipeline_configurations")
}

func (c *MistralClient) RegisterIngestionPipelineConfiguration(req *RegisterIngestionPipelineConfigurationRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"name":                 req.Name,
		"pipeline_composition": req.PipelineComposition,
	})
	return c.requestMap(http.MethodPut, body, "v1/rag/ingestion_pipeline_configurations")
}

func (c *MistralClient) UpdateIngestionPipelineRunInfo(id string, req *UpdateIngestionPipelineRunInfoRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"execution_time": req.ExecutionTime,
		"chunks_count":   req.ChunksCount,
	})
	return c.requestMap(http.MethodPut, body, fmt.Sprintf("v1/rag/ingestion_pipeline_configurations/%s/run_info", id))
}

func (c *MistralClient) ListSearchIndexes() ([]SearchIndexResponse, error) {
	response, err := c.request(http.MethodGet, nil, "v1/rag/search_index", false, nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var indexes []SearchIndexResponse
	if err := json.Unmarshal(data, &indexes); err != nil {
		return nil, err
	}
	return indexes, nil
}

func (c *MistralClient) RegisterSearchIndex(req *RegisterSearchIndexRequest) (*SearchIndexResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"name":           req.Name,
		"index":          req.Index,
		"document_count": req.DocumentCount,
		"status":         req.Status,
	})
	response, err := c.request(http.MethodPut, body, "v1/rag/search_index", false, nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var index SearchIndexResponse
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, err
	}
	return &index, nil
}
