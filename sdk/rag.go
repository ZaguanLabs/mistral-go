package sdk

import (
	"fmt"
	"net/http"
)

type RegisterIngestionPipelineConfigurationRequest struct {
	Name                string `json:"name"`
	PipelineComposition any    `json:"pipeline_composition"`
}

type UpdateIngestionPipelineRunInfoRequest struct {
	ExecutionTime *float64 `json:"execution_time,omitempty"`
	ChunksCount   *int     `json:"chunks_count,omitempty"`
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
