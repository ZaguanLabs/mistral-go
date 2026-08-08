package sdk

import (
	"fmt"
	"net/http"
	"time"
)

type ObservabilitySearchParams struct {
	From             *time.Time `json:"from,omitempty"`
	To               *time.Time `json:"to,omitempty"`
	PageSize         *int       `json:"page_size,omitempty"`
	Cursor           *string    `json:"cursor,omitempty"`
	SearchExpression *string    `json:"search_expression,omitempty"`
	Order            *Order     `json:"order,omitempty"`
}

type ObservabilityFieldOptionsParams struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type ObservabilityAggregationParams struct {
	Metric           any
	From             *time.Time
	To               *time.Time
	Dimensions       []string
	TimeDimension    any
	SearchExpression *string
	OrderBy          []map[string]any
	Limit            *int
}

func (c *MistralClient) SearchLogs(params *ObservabilitySearchParams) (APIResponse, error) {
	if params == nil {
		params = &ObservabilitySearchParams{}
	}
	query := observabilitySearchQuery(params)
	body := optionalRequestMap(map[string]any{
		"search_expression": params.SearchExpression,
		"order":             params.Order,
	})
	return c.requestMap(http.MethodPost, body, appendQuery("v1/observability/logs/search", query))
}

func (c *MistralClient) ListLogFields() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/observability/logs/fields")
}

func (c *MistralClient) FetchLogFieldOptions(fieldName string, params *ObservabilityFieldOptionsParams) (APIResponse, error) {
	return c.observabilityFieldOptions(fmt.Sprintf("v1/observability/logs/fields/%s/options", fieldName), params)
}

func (c *MistralClient) SearchSpans(params *ObservabilitySearchParams) (APIResponse, error) {
	return c.searchObservabilitySignals("v1/observability/spans/search", params)
}

func (c *MistralClient) AggregateSpans(params *ObservabilityAggregationParams) (APIResponse, error) {
	return c.aggregateObservabilitySignals("v1/observability/spans/aggregate", params)
}

func (c *MistralClient) SearchSpanEvaluations(params *ObservabilitySearchParams) (APIResponse, error) {
	return c.searchObservabilitySignals("v1/observability/spans/evaluations/search", params)
}

func (c *MistralClient) SearchLatestSpanEvaluations(params *ObservabilitySearchParams) (APIResponse, error) {
	return c.searchObservabilitySignals("v1/observability/spans/evaluations/search/latest", params)
}

func (c *MistralClient) ListSpanFields() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/observability/spans/fields")
}

func (c *MistralClient) ListSpanEvaluationFields() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/observability/spans/evaluations/fields")
}

func (c *MistralClient) FetchSpanFieldOptions(fieldName string, params *ObservabilityFieldOptionsParams) (APIResponse, error) {
	return c.observabilityFieldOptions(fmt.Sprintf("v1/observability/spans/fields/%s/options", fieldName), params)
}

func (c *MistralClient) FetchSpanEvaluationFieldOptions(fieldName string, params *ObservabilityFieldOptionsParams) (APIResponse, error) {
	return c.observabilityFieldOptions(fmt.Sprintf("v1/observability/spans/evaluations/fields/%s/options", fieldName), params)
}

func (c *MistralClient) SearchTraces(params *ObservabilitySearchParams) (APIResponse, error) {
	return c.searchObservabilitySignals("v1/observability/traces/search", params)
}

func (c *MistralClient) AggregateTraces(params *ObservabilityAggregationParams) (APIResponse, error) {
	return c.aggregateObservabilitySignals("v1/observability/traces/aggregate", params)
}

func (c *MistralClient) ListTraceFields() (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, "v1/observability/traces/fields")
}

func (c *MistralClient) GetTraceByID(traceID string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/observability/traces/%s", traceID))
}

func (c *MistralClient) GetTraceSpans(traceID string, params *ObservabilitySearchParams) (APIResponse, error) {
	if params == nil {
		params = &ObservabilitySearchParams{}
	}
	query := observabilitySearchQuery(params)
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/observability/traces/%s/spans", traceID), query))
}

func (c *MistralClient) FetchTraceFieldOptions(fieldName string, params *ObservabilityFieldOptionsParams) (APIResponse, error) {
	return c.observabilityFieldOptions(fmt.Sprintf("v1/observability/traces/fields/%s/options", fieldName), params)
}

func (c *MistralClient) GetSpanByID(traceID, spanID string) (APIResponse, error) {
	return c.GetSpanByIDWithParams(traceID, spanID, nil)
}

func (c *MistralClient) GetSpanByIDWithParams(traceID, spanID string, params *ObservabilityFieldOptionsParams) (APIResponse, error) {
	if params == nil {
		params = &ObservabilityFieldOptionsParams{}
	}
	query := queryWithOptionalValues(map[string]any{"from": params.From, "to": params.To})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/observability/traces/%s/spans/%s", traceID, spanID), query))
}

func (c *MistralClient) searchObservabilitySignals(path string, params *ObservabilitySearchParams) (APIResponse, error) {
	if params == nil {
		params = &ObservabilitySearchParams{}
	}
	query := observabilitySearchQuery(params)
	body := optionalRequestMap(map[string]any{"search_expression": params.SearchExpression})
	return c.requestMap(http.MethodPost, body, appendQuery(path, query))
}

func (c *MistralClient) observabilityFieldOptions(path string, params *ObservabilityFieldOptionsParams) (APIResponse, error) {
	if params == nil {
		params = &ObservabilityFieldOptionsParams{}
	}
	query := queryWithOptionalValues(map[string]any{"from": params.From, "to": params.To})
	return c.requestMap(http.MethodGet, nil, appendQuery(path, query))
}

func (c *MistralClient) aggregateObservabilitySignals(path string, params *ObservabilityAggregationParams) (APIResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	query := queryWithOptionalValues(map[string]any{"from": params.From, "to": params.To})
	body := optionalRequestMap(map[string]any{
		"metric": params.Metric, "dimensions": params.Dimensions,
		"time_dimension": params.TimeDimension, "search_expression": params.SearchExpression,
		"order_by": params.OrderBy, "limit": params.Limit,
	})
	return c.requestMap(http.MethodPost, body, appendQuery(path, query))
}

func observabilitySearchQuery(params *ObservabilitySearchParams) string {
	return queryWithOptionalValues(map[string]any{
		"from":      params.From,
		"to":        params.To,
		"page_size": params.PageSize,
		"cursor":    params.Cursor,
	})
}
