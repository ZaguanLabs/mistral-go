package sdk

import (
	"net/http"
	"net/url"
)

// EvaluationPageParams uses upstream pagination defaults when fields are omitted.
type EvaluationPageParams struct {
	Page, PageSize *int
	Q              *string
}

func evaluationPageQuery(p *EvaluationPageParams) url.Values {
	page, size := 1, 50
	var q *string
	if p != nil {
		if p.Page != nil {
			page = *p.Page
		}
		if p.PageSize != nil {
			size = *p.PageSize
		}
		q = p.Q
	}
	values, _ := url.ParseQuery(queryWithOptionalValues(map[string]any{"page": page, "page_size": size, "q": q}))
	return values
}
func (c *MistralClient) ListEvaluationProjects(p *EvaluationPageParams) (*Projects, error) {
	return requestTyped[Projects](c, http.MethodGet, nil, appendQuery("v1/observability/projects", evaluationPageQuery(p).Encode()))
}
func (c *MistralClient) GetEvaluationProject(slug string) (*Project, error) {
	return requestTyped[Project](c, http.MethodGet, nil, "v1/observability/projects/"+url.PathEscape(slug))
}
func (c *MistralClient) ListEvaluationsV2(projectSlug string, p *EvaluationPageParams) (*EvaluationsV2, error) {
	return requestTyped[EvaluationsV2](c, http.MethodGet, nil, appendQuery("v1/observability/projects/"+url.PathEscape(projectSlug)+"/evaluations_v2", evaluationPageQuery(p).Encode()))
}
func (c *MistralClient) GetEvaluationV2(slug string) (*EvaluationV2, error) {
	return requestTyped[EvaluationV2](c, http.MethodGet, nil, "v1/observability/evaluations_v2/"+url.PathEscape(slug))
}
func (c *MistralClient) SearchEvaluationRunsV2(slug string, filter *EvaluationRunFilterInput, p *EvaluationPageParams) (*ListEvaluationRunsResponse, error) {
	return requestTyped[ListEvaluationRunsResponse](c, http.MethodPost, optionalEvaluationBody(filter), appendQuery("v1/observability/evaluations_v2/"+url.PathEscape(slug)+"/runs/search", evaluationPageQuery(p).Encode()))
}
func (c *MistralClient) SearchAllEvaluationRunsV2(filter *EvaluationRunFilterInput, p *EvaluationPageParams) (*SearchAllEvaluationRunsResponse, error) {
	return requestTyped[SearchAllEvaluationRunsResponse](c, http.MethodPost, optionalEvaluationBody(filter), appendQuery("v1/observability/evaluation-runs_v2/search", evaluationPageQuery(p).Encode()))
}
func optionalEvaluationBody[T any](v *T) any {
	if v == nil {
		return nil
	}
	return v
}
func (c *MistralClient) GetEvaluationRunV2(id string) (*EvaluationRunV2Response, error) {
	return requestTyped[EvaluationRunV2Response](c, http.MethodGet, nil, "v1/observability/evaluation-runs_v2/"+url.PathEscape(id))
}
func (c *MistralClient) GetEvaluationRunStatisticsV2(id string) (*RunStatisticsResponse, error) {
	return requestTyped[RunStatisticsResponse](c, http.MethodGet, nil, "v1/observability/evaluation-runs_v2/"+url.PathEscape(id)+"/statistics")
}
func (c *MistralClient) ComputeEvaluationRunStatistics(id string, filter *EvaluationRunRecordFilterInput) (*RunStatisticsResponse, error) {
	return requestTyped[RunStatisticsResponse](c, http.MethodPost, optionalEvaluationBody(filter), "v1/observability/evaluation-runs_v2/"+url.PathEscape(id)+"/statistics/compute")
}
func (c *MistralClient) GetEvaluationRunRecordV2(runID, recordID string) (*EvaluationRunRecordV2Response, error) {
	return requestTyped[EvaluationRunRecordV2Response](c, http.MethodGet, nil, "v1/observability/evaluation-runs_v2/"+url.PathEscape(runID)+"/records/"+url.PathEscape(recordID))
}
func (c *MistralClient) SearchEvaluationRunRecordsV2(id string, filter *EvaluationRunRecordFilterInput, p *EvaluationPageParams) (*ListEvaluationRunRecordsResponse, error) {
	return requestTyped[ListEvaluationRunRecordsResponse](c, http.MethodPost, optionalEvaluationBody(filter), appendQuery("v1/observability/evaluation-runs_v2/"+url.PathEscape(id)+"/records/search", evaluationPageQuery(p).Encode()))
}

type ListOptimizationsParams struct {
	EvaluationPageParams
	Project, Evaluation, Status, Algorithm *string
	Tags                                   []string
}

func (c *MistralClient) ListOptimizations(p *ListOptimizationsParams) (*ListOptimizationsResponse, error) {
	if p == nil {
		p = &ListOptimizationsParams{}
	}
	q := evaluationPageQuery(&p.EvaluationPageParams)
	extra, _ := url.ParseQuery(queryWithOptionalValues(map[string]any{"project": p.Project, "evaluation": p.Evaluation, "status": p.Status, "algorithm": p.Algorithm}))
	for k, v := range extra {
		q[k] = v
	}
	for _, tag := range p.Tags {
		q.Add("tags", tag)
	}
	return requestTyped[ListOptimizationsResponse](c, http.MethodGet, nil, appendQuery("v1/observability/optimizations", q.Encode()))
}
func (c *MistralClient) GetOptimization(id string) (*OptimizationDetail, error) {
	return requestTyped[OptimizationDetail](c, http.MethodGet, nil, "v1/observability/optimizations/"+url.PathEscape(id))
}
func (c *MistralClient) ListOptimizationTrials(id string, p *EvaluationPageParams) (*ListOptimizationTrialsResponse, error) {
	return requestTyped[ListOptimizationTrialsResponse](c, http.MethodGet, nil, appendQuery("v1/observability/optimizations/"+url.PathEscape(id)+"/trials", evaluationPageQuery(p).Encode()))
}

type SearchPipelineResultsParams struct {
	ObservabilitySearchParams
	PipelineConfigSlugs []string
}

func (c *MistralClient) SearchPipelineResults(p *SearchPipelineResultsParams) (*GetPipelineResults, error) {
	if p == nil {
		p = &SearchPipelineResultsParams{}
	}
	params := p.ObservabilitySearchParams
	if params.PageSize == nil {
		size := 50
		params.PageSize = &size
	}
	q := observabilitySearchQuery(&params)
	return requestTyped[GetPipelineResults](c, http.MethodPost, &PipelineResultsRequest{PipelineConfigSlugs: p.PipelineConfigSlugs, SearchExpression: p.SearchExpression}, appendQuery("v1/observability/pipeline-results/search", q))
}

// Whoami returns the current user or service-account principal and workspace.
func (c *MistralClient) Whoami() (*WhoamiResponse, error) {
	return requestTyped[WhoamiResponse](c, http.MethodGet, nil, "v1/whoami")
}
