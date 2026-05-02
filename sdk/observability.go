package sdk

import (
	"fmt"
	"net/http"
)

type ListObservabilityParams struct {
	PageSize *int    `json:"page_size,omitempty"`
	Page     *int    `json:"page,omitempty"`
	Q        *string `json:"q,omitempty"`
}

type CreateCampaignRequest struct {
	SearchParams map[string]any `json:"search_params,omitempty"`
	JudgeID      string         `json:"judge_id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Description  *string        `json:"description,omitempty"`
	MaxNbEvents  *int           `json:"max_nb_events,omitempty"`
}

type CreateDatasetRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type UpdateDatasetRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type DatasetRecordRequest struct {
	Payload    any            `json:"payload,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
}

type JudgeRequest struct {
	Name         string  `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	ModelName    *string `json:"model_name,omitempty"`
	Output       any     `json:"output,omitempty"`
	Instructions *string `json:"instructions,omitempty"`
	Tools        any     `json:"tools,omitempty"`
}

type ListJudgesParams struct {
	TypeFilter  *string `json:"type_filter,omitempty"`
	ModelFilter *string `json:"model_filter,omitempty"`
	PageSize    *int    `json:"page_size,omitempty"`
	Page        *int    `json:"page,omitempty"`
	Q           *string `json:"q,omitempty"`
}

type JudgeConversationRequest struct {
	Messages   any            `json:"messages,omitempty"`
	Properties map[string]any `json:"properties,omitempty"`
}

type FieldOptionCountsRequest struct {
	FilterParams map[string]any `json:"filter_params,omitempty"`
}

type SearchChatCompletionEventsRequest struct {
	SearchParams map[string]any `json:"search_params"`
	PageSize     *int           `json:"-"`
	Cursor       *string        `json:"-"`
	ExtraFields  []string       `json:"extra_fields,omitempty"`
}

func (c *MistralClient) CreateCampaign(req *CreateCampaignRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"search_params": req.SearchParams,
		"judge_id":      req.JudgeID,
		"name":          req.Name,
		"description":   req.Description,
		"max_nb_events": req.MaxNbEvents,
	})
	return c.requestMap(http.MethodPost, body, "v1/observability/campaigns")
}

func (c *MistralClient) ListCampaigns(params *ListObservabilityParams) (APIResponse, error) {
	return c.listObservability("v1/observability/campaigns", params)
}

func (c *MistralClient) FetchCampaign(campaignID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/campaigns/%s", campaignID))
}

func (c *MistralClient) DeleteCampaign(campaignID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/observability/campaigns/%s", campaignID))
}

func (c *MistralClient) FetchCampaignStatus(campaignID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/campaigns/%s/status", campaignID))
}

func (c *MistralClient) ListCampaignEvents(campaignID string, pageSize, page *int) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"page_size": pageSize, "page": page})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/observability/campaigns/%s/selected-events", campaignID), query))
}

func (c *MistralClient) CreateDataset(req *CreateDatasetRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"name": req.Name, "description": req.Description})
	return c.requestMap(http.MethodPost, body, "v1/observability/datasets")
}

func (c *MistralClient) ListDatasets(params *ListObservabilityParams) (APIResponse, error) {
	return c.listObservability("v1/observability/datasets", params)
}

func (c *MistralClient) FetchDataset(datasetID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/datasets/%s", datasetID))
}

func (c *MistralClient) DeleteDataset(datasetID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/observability/datasets/%s", datasetID))
}

func (c *MistralClient) UpdateDataset(datasetID string, req *UpdateDatasetRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"name": req.Name, "description": req.Description})
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v1/observability/datasets/%s", datasetID))
}

func (c *MistralClient) ListDatasetRecords(datasetID string, pageSize, page *int) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"page_size": pageSize, "page": page})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/observability/datasets/%s/records", datasetID), query))
}

func (c *MistralClient) CreateDatasetRecord(datasetID string, req *DatasetRecordRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"payload": req.Payload, "properties": req.Properties})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/observability/datasets/%s/records", datasetID))
}

func (c *MistralClient) ImportDatasetFromCampaign(datasetID, campaignID string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"campaign_id": campaignID}, fmt.Sprintf("v1/observability/datasets/%s/imports/from-campaign", datasetID))
}

func (c *MistralClient) ImportDatasetFromExplorer(datasetID string, completionEventIDs []string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"completion_event_ids": completionEventIDs}, fmt.Sprintf("v1/observability/datasets/%s/imports/from-explorer", datasetID))
}

func (c *MistralClient) ImportDatasetFromFile(datasetID, fileID string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"file_id": fileID}, fmt.Sprintf("v1/observability/datasets/%s/imports/from-file", datasetID))
}

func (c *MistralClient) ImportDatasetFromPlayground(datasetID string, conversationIDs []string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"conversation_ids": conversationIDs}, fmt.Sprintf("v1/observability/datasets/%s/imports/from-playground", datasetID))
}

func (c *MistralClient) ImportDatasetFromDatasetRecords(datasetID string, datasetRecordIDs []string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"dataset_record_ids": datasetRecordIDs}, fmt.Sprintf("v1/observability/datasets/%s/imports/from-dataset", datasetID))
}

func (c *MistralClient) ExportDatasetToJSONL(datasetID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/datasets/%s/exports/to-jsonl", datasetID))
}

func (c *MistralClient) FetchDatasetTask(datasetID, taskID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/datasets/%s/tasks/%s", datasetID, taskID))
}

func (c *MistralClient) ListDatasetTasks(datasetID string, pageSize, page *int) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"page_size": pageSize, "page": page})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/observability/datasets/%s/tasks", datasetID), query))
}

func (c *MistralClient) ListChatCompletionFields() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/observability/chat-completion-fields")
}

func (c *MistralClient) FetchChatCompletionFieldOptions(fieldName string, operator *string) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"operator": operator})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/observability/chat-completion-fields/%s/options", fieldName), query))
}

func (c *MistralClient) FetchChatCompletionFieldOptionCounts(fieldName string, req *FieldOptionCountsRequest) (APIResponse, error) {
	body := map[string]interface{}{}
	if req != nil {
		body = optionalRequestMap(map[string]any{"filter_params": req.FilterParams})
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/observability/chat-completion-fields/%s/options-counts", fieldName))
}

func (c *MistralClient) SearchChatCompletionEvents(req *SearchChatCompletionEventsRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	query := queryWithOptionalValues(map[string]any{"page_size": req.PageSize, "cursor": req.Cursor})
	body := optionalRequestMap(map[string]any{"search_params": req.SearchParams, "extra_fields": req.ExtraFields})
	return c.requestMap(http.MethodPost, body, appendQuery("v1/observability/chat-completion-events/search", query))
}

func (c *MistralClient) SearchChatCompletionEventIDs(searchParams map[string]any, extraFields []string) (APIResponse, error) {
	body := optionalRequestMap(map[string]any{"search_params": searchParams, "extra_fields": extraFields})
	return c.requestMap(http.MethodPost, body, "v1/observability/chat-completion-events/search-ids")
}

func (c *MistralClient) FetchChatCompletionEvent(eventID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/chat-completion-events/%s", eventID))
}

func (c *MistralClient) FetchSimilarChatCompletionEvents(eventID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/chat-completion-events/%s/similar-events", eventID))
}

func (c *MistralClient) JudgeChatCompletionEvent(eventID string, judgeDefinition any) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"judge_definition": judgeDefinition}, fmt.Sprintf("v1/observability/chat-completion-events/%s/live-judging", eventID))
}

func (c *MistralClient) CreateJudge(req *JudgeRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return c.requestMap(http.MethodPost, judgeRequestMap(req), "v1/observability/judges")
}

func (c *MistralClient) ListJudges(params *ListJudgesParams) (APIResponse, error) {
	if params == nil {
		params = &ListJudgesParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"type_filter":  params.TypeFilter,
		"model_filter": params.ModelFilter,
		"page_size":    params.PageSize,
		"page":         params.Page,
		"q":            params.Q,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery("v1/observability/judges", query))
}

func (c *MistralClient) FetchJudge(judgeID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/judges/%s", judgeID))
}

func (c *MistralClient) DeleteJudge(judgeID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/observability/judges/%s", judgeID))
}

func (c *MistralClient) UpdateJudge(judgeID string, req *JudgeRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	return c.requestMap(http.MethodPut, judgeRequestMap(req), fmt.Sprintf("v1/observability/judges/%s", judgeID))
}

func (c *MistralClient) JudgeConversation(judgeID string, req *JudgeConversationRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"messages": req.Messages, "properties": req.Properties})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/observability/judges/%s/live-judging", judgeID))
}

func (c *MistralClient) FetchDatasetRecord(datasetRecordID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/dataset-records/%s", datasetRecordID))
}

func (c *MistralClient) DeleteDatasetRecord(datasetRecordID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/observability/dataset-records/%s", datasetRecordID))
}

func (c *MistralClient) BulkDeleteDatasetRecords(datasetRecordIDs []string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"dataset_record_ids": datasetRecordIDs}, "v1/observability/dataset-records/bulk-delete")
}

func (c *MistralClient) JudgeDatasetRecord(datasetRecordID string, judgeDefinition any) (APIResponse, error) {
	return c.requestMap(http.MethodPost, map[string]interface{}{"judge_definition": judgeDefinition}, fmt.Sprintf("v1/observability/dataset-records/%s/live-judging", datasetRecordID))
}

func (c *MistralClient) UpdateDatasetRecordPayload(datasetRecordID string, payload any) (APIResponse, error) {
	return c.requestMap(http.MethodPut, map[string]interface{}{"payload": payload}, fmt.Sprintf("v1/observability/dataset-records/%s/payload", datasetRecordID))
}

func (c *MistralClient) UpdateDatasetRecordProperties(datasetRecordID string, properties map[string]any) (APIResponse, error) {
	return c.requestMap(http.MethodPut, map[string]interface{}{"properties": properties}, fmt.Sprintf("v1/observability/dataset-records/%s/properties", datasetRecordID))
}

func (c *MistralClient) listObservability(path string, params *ListObservabilityParams) (APIResponse, error) {
	if params == nil {
		params = &ListObservabilityParams{}
	}
	query := queryWithOptionalValues(map[string]any{"page_size": params.PageSize, "page": params.Page, "q": params.Q})
	return c.requestMap(http.MethodGet, nil, appendQuery(path, query))
}

func judgeRequestMap(req *JudgeRequest) map[string]interface{} {
	return optionalRequestMap(map[string]any{
		"name":         req.Name,
		"description":  req.Description,
		"model_name":   req.ModelName,
		"output":       req.Output,
		"instructions": req.Instructions,
		"tools":        req.Tools,
	})
}
