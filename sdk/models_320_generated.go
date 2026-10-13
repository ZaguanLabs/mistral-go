// Code generated from official Mistral Python SDK 3.2.0 models. DO NOT EDIT.
package sdk

import (
	"encoding/json"
	"time"
)

type AuthenticationType = string
type ClassificationStatisticSpecKind = string
type Average = string
type ContentSearchOperator = string
type EvaluationResultSpec = any
type Sort = any
type EvaluatorStatistics = any
type Statistic = any
type EvaluatorV2Direction = string
type FilterCombinator = string
type GoalResultItemDirection = string
type GoalStatus = string
type IPAPhoneme = string
type MetadataFilterValue = any
type NumericScoreMetric = string
type Percentile = any
type PipelineConfigSampling = any
type RecordSortField = string
type ResourceChunkContent = any
type Metadata = any
type ScoreFilterValue = any
type ScoreOperator = string
type SimpleNumericStatisticSpecKind = string
type SortDirection = string
type SystemParamFilterValue = any
type SystemParamOperator = string
type Params = any
type VoiceKind = string
type Principal = any
type AttributeHashSampling struct {
	Attribute  string  `json:"attribute"`
	Percentage float64 `json:"percentage"`
	Type       string  `json:"type,omitempty"`
}

func (v AttributeHashSampling) MarshalJSON() ([]byte, error) {
	type wire AttributeHashSampling
	if v.Type == "" {
		v.Type = "attribute_hash"
	}
	return json.Marshal(wire(v))
}

type CategoricalStatistics struct {
	Frequencies     map[string]int         `json:"frequencies"`
	Mode            []string               `json:"mode"`
	Count           int                    `json:"count"`
	Type            *string                `json:"type,omitempty"`
	Values          []ScalarStatisticValue `json:"values,omitempty"`
	ExcludedCount   *int                   `json:"excluded_count,omitempty"`
	ConfusionMatrix *ConfusionMatrix       `json:"confusion_matrix,omitempty"`
}

func (v CategoricalStatistics) MarshalJSON() ([]byte, error) {
	type wire CategoricalStatistics
	if v.Type == nil {
		x := "categorical"
		v.Type = &x
	}
	if v.ExcludedCount == nil {
		x := 0
		v.ExcludedCount = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Values != nil {
		data, err := json.Marshal(v.Values)
		if err != nil {
			return nil, err
		}
		fields["values"] = data
	}
	return json.Marshal(fields)
}

type ChatCompletionResponseChoiceLogprobs struct {
	Content []ChatCompletionResponseTokenLogprobsTopK `json:"content"`
}
type ChatCompletionResponseTokenLogprobs struct {
	Token   string  `json:"token"`
	Logprob float64 `json:"logprob"`
	Bytes   []int   `json:"bytes"`
	TokenID int     `json:"token_id"`
}
type ChatCompletionResponseTokenLogprobsTopK struct {
	Token       string                                `json:"token"`
	Logprob     float64                               `json:"logprob"`
	Bytes       []int                                 `json:"bytes"`
	TokenID     int                                   `json:"token_id"`
	TopLogprobs []ChatCompletionResponseTokenLogprobs `json:"top_logprobs,omitempty"`
}

func (v ChatCompletionResponseTokenLogprobsTopK) MarshalJSON() ([]byte, error) {
	type wire ChatCompletionResponseTokenLogprobsTopK
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.TopLogprobs != nil {
		data, err := json.Marshal(v.TopLogprobs)
		if err != nil {
			return nil, err
		}
		fields["top_logprobs"] = data
	}
	return json.Marshal(fields)
}

type ClassificationStatisticSpec struct {
	Kind          ClassificationStatisticSpecKind `json:"kind"`
	PositiveLabel *string                         `json:"positive_label,omitempty"`
	Average       *Average                        `json:"average,omitempty"`
	ExpectedKey   *string                         `json:"expected_key,omitempty"`
	Labels        []string                        `json:"labels,omitempty"`
	Goal          *GoalSpec                       `json:"goal,omitempty"`
}

func (v ClassificationStatisticSpec) MarshalJSON() ([]byte, error) {
	type wire ClassificationStatisticSpec
	if v.ExpectedKey == nil {
		x := "expected"
		v.ExpectedKey = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Labels != nil {
		data, err := json.Marshal(v.Labels)
		if err != nil {
			return nil, err
		}
		fields["labels"] = data
	}
	return json.Marshal(fields)
}

type ConfusionMatrix struct {
	Labels []string `json:"labels"`
	Matrix [][]int  `json:"matrix"`
}
type ContentSearchFilter struct {
	Value    string                 `json:"value"`
	Path     *string                `json:"path,omitempty"`
	Operator *ContentSearchOperator `json:"operator,omitempty"`
}
type ContentSearchFilterGroup struct {
	Combinator *FilterCombinator     `json:"combinator,omitempty"`
	Filters    []ContentSearchFilter `json:"filters,omitempty"`
}

func (v ContentSearchFilterGroup) MarshalJSON() ([]byte, error) {
	type wire ContentSearchFilterGroup
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Filters != nil {
		data, err := json.Marshal(v.Filters)
		if err != nil {
			return nil, err
		}
		fields["filters"] = data
	}
	return json.Marshal(fields)
}

type DetectionDefinitionV1 struct {
	TargetAttributes []string           `json:"target_attributes"`
	Patterns         []DetectionPattern `json:"patterns"`
}
type EvaluationBooleanResultSpec struct {
	Type string `json:"type,omitempty"`
}

func (v EvaluationBooleanResultSpec) MarshalJSON() ([]byte, error) {
	type wire EvaluationBooleanResultSpec
	if v.Type == "" {
		v.Type = "boolean"
	}
	return json.Marshal(wire(v))
}

type EvaluationCategoricalResultSpec struct {
	Labels []string `json:"labels"`
	Type   string   `json:"type,omitempty"`
}

func (v EvaluationCategoricalResultSpec) MarshalJSON() ([]byte, error) {
	type wire EvaluationCategoricalResultSpec
	if v.Type == "" {
		v.Type = "categorical"
	}
	return json.Marshal(wire(v))
}

type EvaluationNumericResultSpec struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Type string  `json:"type,omitempty"`
}

func (v EvaluationNumericResultSpec) MarshalJSON() ([]byte, error) {
	type wire EvaluationNumericResultSpec
	if v.Type == "" {
		v.Type = "numeric"
	}
	return json.Marshal(wire(v))
}

type UnknownEvaluationResultSpec struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownEvaluationResultSpec) MarshalJSON() ([]byte, error) {
	type wire UnknownEvaluationResultSpec
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type EvaluationRunFilterInput struct {
	RunIds        []string               `json:"run_ids,omitempty"`
	EvaluationIds []string               `json:"evaluation_ids,omitempty"`
	CreatorIds    []string               `json:"creator_ids,omitempty"`
	ProjectIds    []string               `json:"project_ids,omitempty"`
	Ownership     *RunOwnershipFilter    `json:"ownership,omitempty"`
	Tags          *TagFilter             `json:"tags,omitempty"`
	ScoreFilters  *ScoreFilterGroup      `json:"score_filters,omitempty"`
	GoalFilters   *GoalStatusFilterGroup `json:"goal_filters,omitempty"`
	System        *SystemFilter          `json:"system,omitempty"`
	Metadata      *MetadataFilterGroup   `json:"metadata,omitempty"`
}

func (v EvaluationRunFilterInput) MarshalJSON() ([]byte, error) {
	type wire EvaluationRunFilterInput
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.RunIds != nil {
		data, err := json.Marshal(v.RunIds)
		if err != nil {
			return nil, err
		}
		fields["run_ids"] = data
	}
	if v.EvaluationIds != nil {
		data, err := json.Marshal(v.EvaluationIds)
		if err != nil {
			return nil, err
		}
		fields["evaluation_ids"] = data
	}
	if v.CreatorIds != nil {
		data, err := json.Marshal(v.CreatorIds)
		if err != nil {
			return nil, err
		}
		fields["creator_ids"] = data
	}
	if v.ProjectIds != nil {
		data, err := json.Marshal(v.ProjectIds)
		if err != nil {
			return nil, err
		}
		fields["project_ids"] = data
	}
	return json.Marshal(fields)
}

type EvaluationRunGenerationV2 struct {
	ID          string                            `json:"id"`
	CreatedAt   time.Time                         `json:"created_at"`
	UpdatedAt   time.Time                         `json:"updated_at"`
	DeletedAt   *time.Time                        `json:"deleted_at"`
	Output      any                               `json:"output"`
	Metadata    map[string]any                    `json:"metadata_"`
	Status      string                            `json:"status"`
	Error       *string                           `json:"error,omitempty"`
	Scores      map[string][]EvaluationRunScoreV2 `json:"scores,omitempty"`
	Statistics  map[string]EvaluatorStatistics    `json:"statistics,omitempty"`
	GoalResults map[string]GoalResult             `json:"goal_results,omitempty"`
}

func (v EvaluationRunGenerationV2) MarshalJSON() ([]byte, error) {
	type wire EvaluationRunGenerationV2
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Scores != nil {
		data, err := json.Marshal(v.Scores)
		if err != nil {
			return nil, err
		}
		fields["scores"] = data
	}
	if v.Statistics != nil {
		data, err := json.Marshal(v.Statistics)
		if err != nil {
			return nil, err
		}
		fields["statistics"] = data
	}
	if v.GoalResults != nil {
		data, err := json.Marshal(v.GoalResults)
		if err != nil {
			return nil, err
		}
		fields["goal_results"] = data
	}
	return json.Marshal(fields)
}

type EvaluationRunOutputRecordV2 struct {
	ID          string                         `json:"id"`
	CreatedAt   time.Time                      `json:"created_at"`
	UpdatedAt   time.Time                      `json:"updated_at"`
	DeletedAt   *time.Time                     `json:"deleted_at"`
	Metadata    map[string]any                 `json:"metadata_"`
	RecordInfo  map[string]any                 `json:"record_info,omitempty"`
	Generations []EvaluationRunGenerationV2    `json:"generations,omitempty"`
	Statistics  map[string]EvaluatorStatistics `json:"statistics,omitempty"`
	GoalResults map[string]GoalResult          `json:"goal_results,omitempty"`
}

func (v EvaluationRunOutputRecordV2) MarshalJSON() ([]byte, error) {
	type wire EvaluationRunOutputRecordV2
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.RecordInfo != nil {
		data, err := json.Marshal(v.RecordInfo)
		if err != nil {
			return nil, err
		}
		fields["record_info"] = data
	}
	if v.Generations != nil {
		data, err := json.Marshal(v.Generations)
		if err != nil {
			return nil, err
		}
		fields["generations"] = data
	}
	if v.Statistics != nil {
		data, err := json.Marshal(v.Statistics)
		if err != nil {
			return nil, err
		}
		fields["statistics"] = data
	}
	if v.GoalResults != nil {
		data, err := json.Marshal(v.GoalResults)
		if err != nil {
			return nil, err
		}
		fields["goal_results"] = data
	}
	return json.Marshal(fields)
}

type EvaluationRunRecordFilterInput struct {
	InputFilters  *ContentSearchFilterGroup `json:"input_filters,omitempty"`
	OutputFilters *ContentSearchFilterGroup `json:"output_filters,omitempty"`
	ScoreFilters  *ScoreFilterGroup         `json:"score_filters,omitempty"`
	GoalFilters   *GoalStatusFilterGroup    `json:"goal_filters,omitempty"`
	Sort          *Sort                     `json:"sort,omitempty"`
}
type EvaluationRunRecordV2Response struct {
	ID              string                       `json:"id"`
	CreatedAt       time.Time                    `json:"created_at"`
	UpdatedAt       time.Time                    `json:"updated_at"`
	DeletedAt       *time.Time                   `json:"deleted_at"`
	Input           any                          `json:"input"`
	DatasetRecordID *string                      `json:"dataset_record_id,omitempty"`
	OutputRecord    *EvaluationRunOutputRecordV2 `json:"output_record,omitempty"`
}
type EvaluationRunScoreV2 struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at"`
	Value     any            `json:"value"`
	Metadata  map[string]any `json:"metadata_"`
	Status    string         `json:"status"`
	Rationale *string        `json:"rationale,omitempty"`
	Error     *string        `json:"error,omitempty"`
}
type EvaluationRunV2Response struct {
	ID             string                          `json:"id"`
	CreatedAt      time.Time                       `json:"created_at"`
	UpdatedAt      time.Time                       `json:"updated_at"`
	DeletedAt      *time.Time                      `json:"deleted_at"`
	CreatorID      string                          `json:"creator_id"`
	EvaluationID   string                          `json:"evaluation_id"`
	Name           string                          `json:"name"`
	Description    *string                         `json:"description"`
	Tags           []string                        `json:"tags"`
	Metadata       map[string]any                  `json:"metadata_"`
	NumGenerations int                             `json:"num_generations"`
	Evaluators     []EvaluatorV2                   `json:"evaluators"`
	RunInfo        map[string]any                  `json:"run_info,omitempty"`
	Status         *string                         `json:"status,omitempty"`
	DatasetID      *string                         `json:"dataset_id,omitempty"`
	RunEvaluators  []EvaluatorV2                   `json:"run_evaluators,omitempty"`
	System         *SystemV2                       `json:"system,omitempty"`
	Statistics     map[string]EvaluatorStatistics  `json:"statistics,omitempty"`
	GoalResults    map[string]GoalResult           `json:"goal_results,omitempty"`
	Passed         *bool                           `json:"passed,omitempty"`
	RunScores      map[string]EvaluationRunScoreV2 `json:"run_scores,omitempty"`
	Optimization   *RunOptimizationRef             `json:"optimization,omitempty"`
}

func (v EvaluationRunV2Response) MarshalJSON() ([]byte, error) {
	type wire EvaluationRunV2Response
	if v.Status == nil {
		x := "completed"
		v.Status = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.RunInfo != nil {
		data, err := json.Marshal(v.RunInfo)
		if err != nil {
			return nil, err
		}
		fields["run_info"] = data
	}
	if v.RunEvaluators != nil {
		data, err := json.Marshal(v.RunEvaluators)
		if err != nil {
			return nil, err
		}
		fields["run_evaluators"] = data
	}
	if v.Statistics != nil {
		data, err := json.Marshal(v.Statistics)
		if err != nil {
			return nil, err
		}
		fields["statistics"] = data
	}
	if v.GoalResults != nil {
		data, err := json.Marshal(v.GoalResults)
		if err != nil {
			return nil, err
		}
		fields["goal_results"] = data
	}
	if v.RunScores != nil {
		data, err := json.Marshal(v.RunScores)
		if err != nil {
			return nil, err
		}
		fields["run_scores"] = data
	}
	return json.Marshal(fields)
}

type EvaluationRunV2WithRefsResponse struct {
	ID             string                          `json:"id"`
	CreatedAt      time.Time                       `json:"created_at"`
	UpdatedAt      time.Time                       `json:"updated_at"`
	DeletedAt      *time.Time                      `json:"deleted_at"`
	CreatorID      string                          `json:"creator_id"`
	EvaluationID   string                          `json:"evaluation_id"`
	Name           string                          `json:"name"`
	Description    *string                         `json:"description"`
	Tags           []string                        `json:"tags"`
	Metadata       map[string]any                  `json:"metadata_"`
	NumGenerations int                             `json:"num_generations"`
	Evaluators     []EvaluatorV2                   `json:"evaluators"`
	RunInfo        map[string]any                  `json:"run_info,omitempty"`
	Status         *string                         `json:"status,omitempty"`
	DatasetID      *string                         `json:"dataset_id,omitempty"`
	RunEvaluators  []EvaluatorV2                   `json:"run_evaluators,omitempty"`
	System         *SystemV2                       `json:"system,omitempty"`
	Statistics     map[string]EvaluatorStatistics  `json:"statistics,omitempty"`
	GoalResults    map[string]GoalResult           `json:"goal_results,omitempty"`
	Passed         *bool                           `json:"passed,omitempty"`
	RunScores      map[string]EvaluationRunScoreV2 `json:"run_scores,omitempty"`
	Optimization   *RunOptimizationRef             `json:"optimization,omitempty"`
	Evaluation     *RunEvaluationSummary           `json:"evaluation,omitempty"`
	Project        *RunProjectSummary              `json:"project,omitempty"`
}

func (v EvaluationRunV2WithRefsResponse) MarshalJSON() ([]byte, error) {
	type wire EvaluationRunV2WithRefsResponse
	if v.Status == nil {
		x := "completed"
		v.Status = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.RunInfo != nil {
		data, err := json.Marshal(v.RunInfo)
		if err != nil {
			return nil, err
		}
		fields["run_info"] = data
	}
	if v.RunEvaluators != nil {
		data, err := json.Marshal(v.RunEvaluators)
		if err != nil {
			return nil, err
		}
		fields["run_evaluators"] = data
	}
	if v.Statistics != nil {
		data, err := json.Marshal(v.Statistics)
		if err != nil {
			return nil, err
		}
		fields["statistics"] = data
	}
	if v.GoalResults != nil {
		data, err := json.Marshal(v.GoalResults)
		if err != nil {
			return nil, err
		}
		fields["goal_results"] = data
	}
	if v.RunScores != nil {
		data, err := json.Marshal(v.RunScores)
		if err != nil {
			return nil, err
		}
		fields["run_scores"] = data
	}
	return json.Marshal(fields)
}

type EvaluationsV2 struct {
	Evaluations PaginatedResultEvaluationV2 `json:"evaluations"`
}
type EvaluationV2 struct {
	ID          string         `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at"`
	WorkspaceID string         `json:"workspace_id"`
	CreatorID   string         `json:"creator_id"`
	ProjectID   string         `json:"project_id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description *string        `json:"description"`
	Tags        []string       `json:"tags"`
	Metadata    map[string]any `json:"metadata_"`
	IsFavorite  *bool          `json:"is_favorite,omitempty"`
}

func (v EvaluationV2) MarshalJSON() ([]byte, error) {
	type wire EvaluationV2
	if v.IsFavorite == nil {
		x := false
		v.IsFavorite = &x
	}
	return json.Marshal(wire(v))
}

type UnknownStatistic struct {
	Kind      string `json:"kind,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownStatistic) MarshalJSON() ([]byte, error) {
	type wire UnknownStatistic
	if v.Kind == "" {
		v.Kind = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type EvaluatorV2 struct {
	ID            string                `json:"id"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	DeletedAt     *time.Time            `json:"deleted_at"`
	Name          string                `json:"name"`
	Description   *string               `json:"description"`
	Tags          []string              `json:"tags"`
	NumScores     int                   `json:"num_scores"`
	Goal          *GoalSpec             `json:"goal,omitempty"`
	AggregateGoal *GoalSpec             `json:"aggregate_goal,omitempty"`
	Statistics    []Statistic           `json:"statistics,omitempty"`
	Weight        *float64              `json:"weight,omitempty"`
	Direction     *EvaluatorV2Direction `json:"direction,omitempty"`
	MinValue      *float64              `json:"min_value,omitempty"`
	MaxValue      *float64              `json:"max_value,omitempty"`
}

func (v EvaluatorV2) MarshalJSON() ([]byte, error) {
	type wire EvaluatorV2
	if v.Weight == nil {
		x := float64(1)
		v.Weight = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Statistics != nil {
		data, err := json.Marshal(v.Statistics)
		if err != nil {
			return nil, err
		}
		fields["statistics"] = data
	}
	return json.Marshal(fields)
}

type ExportDefinitionV1 struct {
	Destination OTLPDestination `json:"destination"`
}
type FeedResultGetPipelineResult struct {
	Results []PipelineResult `json:"results,omitempty"`
	Next    *string          `json:"next,omitempty"`
	Cursor  *string          `json:"cursor,omitempty"`
}

func (v FeedResultGetPipelineResult) MarshalJSON() ([]byte, error) {
	type wire FeedResultGetPipelineResult
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type FieldSortInput struct {
	Field     RecordSortField `json:"field"`
	Kind      string          `json:"kind,omitempty"`
	Direction *SortDirection  `json:"direction,omitempty"`
}

func (v FieldSortInput) MarshalJSON() ([]byte, error) {
	type wire FieldSortInput
	if v.Kind == "" {
		v.Kind = "field"
	}
	return json.Marshal(wire(v))
}

type GetPipelineResults struct {
	PipelineResults FeedResultGetPipelineResult `json:"pipeline_results"`
}
type GoalResult struct {
	Passed  *bool            `json:"passed"`
	Results []GoalResultItem `json:"results"`
}
type GoalResultItem struct {
	Metric    string                   `json:"metric"`
	Operator  string                   `json:"operator"`
	Actual    float64                  `json:"actual"`
	Passed    bool                     `json:"passed"`
	Target    *float64                 `json:"target,omitempty"`
	TargetMin *float64                 `json:"target_min,omitempty"`
	TargetMax *float64                 `json:"target_max,omitempty"`
	Direction *GoalResultItemDirection `json:"direction,omitempty"`
}
type GoalSpec struct {
	Operator *string  `json:"operator,omitempty"`
	Value    any      `json:"value,omitempty"`
	Min      *float64 `json:"min,omitempty"`
	Max      *float64 `json:"max,omitempty"`
	Values   []any    `json:"values,omitempty"`
	Metric   *string  `json:"metric,omitempty"`
}

func (v GoalSpec) MarshalJSON() ([]byte, error) {
	type wire GoalSpec
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Values != nil {
		data, err := json.Marshal(v.Values)
		if err != nil {
			return nil, err
		}
		fields["values"] = data
	}
	return json.Marshal(fields)
}

type GoalStatusFilter struct {
	Values        []GoalStatus `json:"values"`
	EvaluatorName *string      `json:"evaluator_name,omitempty"`
}
type GoalStatusFilterGroup struct {
	Combinator *FilterCombinator  `json:"combinator,omitempty"`
	Filters    []GoalStatusFilter `json:"filters,omitempty"`
}

func (v GoalStatusFilterGroup) MarshalJSON() ([]byte, error) {
	type wire GoalStatusFilterGroup
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Filters != nil {
		data, err := json.Marshal(v.Filters)
		if err != nil {
			return nil, err
		}
		fields["filters"] = data
	}
	return json.Marshal(fields)
}

type JudgeDefinitionV1 struct {
	Model       string               `json:"model"`
	Prompt      string               `json:"prompt"`
	ResultSpec  EvaluationResultSpec `json:"result_spec"`
	Temperature *float64             `json:"temperature,omitempty"`
	TopP        *float64             `json:"top_p,omitempty"`
	MaxTokens   *int                 `json:"max_tokens,omitempty"`
	RandomSeed  *int                 `json:"random_seed,omitempty"`
	Mapping     map[string]string    `json:"mapping,omitempty"`
	UseCache    *bool                `json:"use_cache,omitempty"`
}

func (v JudgeDefinitionV1) MarshalJSON() ([]byte, error) {
	type wire JudgeDefinitionV1
	if v.UseCache == nil {
		x := true
		v.UseCache = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Mapping != nil {
		data, err := json.Marshal(v.Mapping)
		if err != nil {
			return nil, err
		}
		fields["mapping"] = data
	}
	return json.Marshal(fields)
}

type ListEvaluationRunRecordsResponse struct {
	Records PaginatedResultEvaluationRunRecordV2Response `json:"records"`
}
type ListEvaluationRunsResponse struct {
	Runs PaginatedResultEvaluationRunV2Response `json:"runs"`
}
type ListOptimizationsResponse struct {
	Optimizations PaginatedResultOptimization `json:"optimizations"`
}
type ListOptimizationTrialsResponse struct {
	Trials PaginatedResultOptimizationTrialWithObservations `json:"trials"`
}
type LlmJudgeReference struct {
	Slug     string                  `json:"slug"`
	Mapping  map[string]string       `json:"mapping,omitempty"`
	Sampling *PipelineConfigSampling `json:"sampling,omitempty"`
}

func (v LlmJudgeReference) MarshalJSON() ([]byte, error) {
	type wire LlmJudgeReference
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Mapping != nil {
		data, err := json.Marshal(v.Mapping)
		if err != nil {
			return nil, err
		}
		fields["mapping"] = data
	}
	return json.Marshal(fields)
}

type MetadataFilter struct {
	Key      string              `json:"key"`
	Operator SystemParamOperator `json:"operator"`
	Value    MetadataFilterValue `json:"value"`
}
type MetadataFilterGroup struct {
	Combinator *FilterCombinator `json:"combinator,omitempty"`
	Filters    []MetadataFilter  `json:"filters,omitempty"`
}

func (v MetadataFilterGroup) MarshalJSON() ([]byte, error) {
	type wire MetadataFilterGroup
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Filters != nil {
		data, err := json.Marshal(v.Filters)
		if err != nil {
			return nil, err
		}
		fields["filters"] = data
	}
	return json.Marshal(fields)
}

type ModerationDefinitionV1 struct {
	Model            *string  `json:"model,omitempty"`
	TargetAttributes []string `json:"target_attributes,omitempty"`
}

func (v ModerationDefinitionV1) MarshalJSON() ([]byte, error) {
	type wire ModerationDefinitionV1
	if v.Model == nil {
		x := "mistral-moderation-latest"
		v.Model = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.TargetAttributes != nil {
		data, err := json.Marshal(v.TargetAttributes)
		if err != nil {
			return nil, err
		}
		fields["target_attributes"] = data
	}
	return json.Marshal(fields)
}

type NumericStatistics struct {
	Type        *string                `json:"type,omitempty"`
	Values      []ScalarStatisticValue `json:"values,omitempty"`
	SampleCount *int                   `json:"sample_count,omitempty"`
	Avg         *float64               `json:"avg,omitempty"`
	Min         *float64               `json:"min,omitempty"`
	Max         *float64               `json:"max,omitempty"`
	Std         *float64               `json:"std,omitempty"`
	Count       *int                   `json:"count,omitempty"`
}

func (v NumericStatistics) MarshalJSON() ([]byte, error) {
	type wire NumericStatistics
	if v.Type == nil {
		x := "numeric"
		v.Type = &x
	}
	if v.SampleCount == nil {
		x := 0
		v.SampleCount = &x
	}
	if v.Count == nil {
		x := 0
		v.Count = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Values != nil {
		data, err := json.Marshal(v.Values)
		if err != nil {
			return nil, err
		}
		fields["values"] = data
	}
	return json.Marshal(fields)
}

type Optimization struct {
	ID                  string         `json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           *time.Time     `json:"deleted_at"`
	CreatorID           string         `json:"creator_id"`
	ProjectID           string         `json:"project_id"`
	Name                string         `json:"name"`
	Algorithm           string         `json:"algorithm"`
	Status              string         `json:"status"`
	Source              string         `json:"source"`
	EvaluationID        *string        `json:"evaluation_id,omitempty"`
	Description         *string        `json:"description,omitempty"`
	Steer               *string        `json:"steer,omitempty"`
	Tags                []string       `json:"tags,omitempty"`
	WorkflowExecutionID *string        `json:"workflow_execution_id,omitempty"`
	WorkflowRunID       *string        `json:"workflow_run_id,omitempty"`
	ObjectiveSpec       map[string]any `json:"objective_spec,omitempty"`
	Outcome             map[string]any `json:"outcome,omitempty"`
}

func (v Optimization) MarshalJSON() ([]byte, error) {
	type wire Optimization
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Tags != nil {
		data, err := json.Marshal(v.Tags)
		if err != nil {
			return nil, err
		}
		fields["tags"] = data
	}
	if v.ObjectiveSpec != nil {
		data, err := json.Marshal(v.ObjectiveSpec)
		if err != nil {
			return nil, err
		}
		fields["objective_spec"] = data
	}
	if v.Outcome != nil {
		data, err := json.Marshal(v.Outcome)
		if err != nil {
			return nil, err
		}
		fields["outcome"] = data
	}
	return json.Marshal(fields)
}

type OptimizationDetail struct {
	ID                  string                 `json:"id"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	DeletedAt           *time.Time             `json:"deleted_at"`
	CreatorID           string                 `json:"creator_id"`
	ProjectID           string                 `json:"project_id"`
	Name                string                 `json:"name"`
	Algorithm           string                 `json:"algorithm"`
	Status              string                 `json:"status"`
	Source              string                 `json:"source"`
	EvaluationID        *string                `json:"evaluation_id,omitempty"`
	Description         *string                `json:"description,omitempty"`
	Steer               *string                `json:"steer,omitempty"`
	Tags                []string               `json:"tags,omitempty"`
	WorkflowExecutionID *string                `json:"workflow_execution_id,omitempty"`
	WorkflowRunID       *string                `json:"workflow_run_id,omitempty"`
	ObjectiveSpec       map[string]any         `json:"objective_spec,omitempty"`
	Outcome             map[string]any         `json:"outcome,omitempty"`
	Evaluation          *OptimizationReference `json:"evaluation,omitempty"`
	Project             *OptimizationReference `json:"project,omitempty"`
}

func (v OptimizationDetail) MarshalJSON() ([]byte, error) {
	type wire OptimizationDetail
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Tags != nil {
		data, err := json.Marshal(v.Tags)
		if err != nil {
			return nil, err
		}
		fields["tags"] = data
	}
	if v.ObjectiveSpec != nil {
		data, err := json.Marshal(v.ObjectiveSpec)
		if err != nil {
			return nil, err
		}
		fields["objective_spec"] = data
	}
	if v.Outcome != nil {
		data, err := json.Marshal(v.Outcome)
		if err != nil {
			return nil, err
		}
		fields["outcome"] = data
	}
	return json.Marshal(fields)
}

type OptimizationReference struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
type OptimizationTrialObservation struct {
	Run             EvaluationRunV2Response `json:"run"`
	ObservationType *string                 `json:"observation_type,omitempty"`
	Score           *float64                `json:"score,omitempty"`
	Metadata        map[string]any          `json:"metadata,omitempty"`
}

func (v OptimizationTrialObservation) MarshalJSON() ([]byte, error) {
	type wire OptimizationTrialObservation
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Metadata != nil {
		data, err := json.Marshal(v.Metadata)
		if err != nil {
			return nil, err
		}
		fields["metadata"] = data
	}
	return json.Marshal(fields)
}

type OptimizationTrialWithObservations struct {
	ID                string                         `json:"id"`
	CreatedAt         time.Time                      `json:"created_at"`
	UpdatedAt         time.Time                      `json:"updated_at"`
	DeletedAt         *time.Time                     `json:"deleted_at"`
	OptimizationID    string                         `json:"optimization_id"`
	AlgorithmTrialKey string                         `json:"algorithm_trial_key"`
	Status            string                         `json:"status"`
	Metadata          map[string]any                 `json:"metadata_"`
	ParentTrialID     *string                        `json:"parent_trial_id,omitempty"`
	Hypothesis        *string                        `json:"hypothesis,omitempty"`
	Resolution        *string                        `json:"resolution,omitempty"`
	Observations      []OptimizationTrialObservation `json:"observations,omitempty"`
}

func (v OptimizationTrialWithObservations) MarshalJSON() ([]byte, error) {
	type wire OptimizationTrialWithObservations
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Observations != nil {
		data, err := json.Marshal(v.Observations)
		if err != nil {
			return nil, err
		}
		fields["observations"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultEvaluationRunRecordV2Response struct {
	Count    int                             `json:"count"`
	Results  []EvaluationRunRecordV2Response `json:"results,omitempty"`
	Next     *string                         `json:"next,omitempty"`
	Previous *string                         `json:"previous,omitempty"`
}

func (v PaginatedResultEvaluationRunRecordV2Response) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultEvaluationRunRecordV2Response
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultEvaluationRunV2Response struct {
	Count    int                       `json:"count"`
	Results  []EvaluationRunV2Response `json:"results,omitempty"`
	Next     *string                   `json:"next,omitempty"`
	Previous *string                   `json:"previous,omitempty"`
}

func (v PaginatedResultEvaluationRunV2Response) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultEvaluationRunV2Response
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultEvaluationRunV2WithRefsResponse struct {
	Count    int                               `json:"count"`
	Results  []EvaluationRunV2WithRefsResponse `json:"results,omitempty"`
	Next     *string                           `json:"next,omitempty"`
	Previous *string                           `json:"previous,omitempty"`
}

func (v PaginatedResultEvaluationRunV2WithRefsResponse) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultEvaluationRunV2WithRefsResponse
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultEvaluationV2 struct {
	Count    int            `json:"count"`
	Results  []EvaluationV2 `json:"results,omitempty"`
	Next     *string        `json:"next,omitempty"`
	Previous *string        `json:"previous,omitempty"`
}

func (v PaginatedResultEvaluationV2) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultEvaluationV2
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultOptimization struct {
	Count    int            `json:"count"`
	Results  []Optimization `json:"results,omitempty"`
	Next     *string        `json:"next,omitempty"`
	Previous *string        `json:"previous,omitempty"`
}

func (v PaginatedResultOptimization) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultOptimization
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultOptimizationTrialWithObservations struct {
	Count    int                                 `json:"count"`
	Results  []OptimizationTrialWithObservations `json:"results,omitempty"`
	Next     *string                             `json:"next,omitempty"`
	Previous *string                             `json:"previous,omitempty"`
}

func (v PaginatedResultOptimizationTrialWithObservations) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultOptimizationTrialWithObservations
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PaginatedResultProject struct {
	Count    int       `json:"count"`
	Results  []Project `json:"results,omitempty"`
	Next     *string   `json:"next,omitempty"`
	Previous *string   `json:"previous,omitempty"`
}

func (v PaginatedResultProject) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultProject
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PercentileStatisticSpec struct {
	Percentile Percentile `json:"percentile"`
	Kind       string     `json:"kind,omitempty"`
	Goal       *GoalSpec  `json:"goal,omitempty"`
}

func (v PercentileStatisticSpec) MarshalJSON() ([]byte, error) {
	type wire PercentileStatisticSpec
	if v.Kind == "" {
		v.Kind = "percentile"
	}
	return json.Marshal(wire(v))
}

type UnknownPipelineConfigSampling struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownPipelineConfigSampling) MarshalJSON() ([]byte, error) {
	type wire UnknownPipelineConfigSampling
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type PipelineResult struct {
	TraceID           string             `json:"trace_id"`
	SpanID            string             `json:"span_id"`
	SpanName          string             `json:"span_name"`
	DurationNS        int                `json:"duration_ns"`
	EvaluatedAt       time.Time          `json:"evaluated_at"`
	JudgeScores       map[string]float64 `json:"judge_scores"`
	JudgeLabels       map[string]string  `json:"judge_labels"`
	JudgeExplanations map[string]string  `json:"judge_explanations"`
}
type PipelineResultsRequest struct {
	PipelineConfigSlugs []string `json:"pipeline_config_slugs"`
	SearchExpression    *string  `json:"search_expression,omitempty"`
}
type ProbabilisticSampling struct {
	Percentage float64 `json:"percentage"`
	Type       string  `json:"type,omitempty"`
}

func (v ProbabilisticSampling) MarshalJSON() ([]byte, error) {
	type wire ProbabilisticSampling
	if v.Type == "" {
		v.Type = "probabilistic"
	}
	return json.Marshal(wire(v))
}

type Project struct {
	ID          string                     `json:"id"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	DeletedAt   *time.Time                 `json:"deleted_at"`
	WorkspaceID string                     `json:"workspace_id"`
	CreatorID   string                     `json:"creator_id"`
	Name        string                     `json:"name"`
	Slug        string                     `json:"slug"`
	Description *string                    `json:"description"`
	IsFavorite  *bool                      `json:"is_favorite,omitempty"`
	Evaluations []ProjectEvaluationSummary `json:"evaluations,omitempty"`
}

func (v Project) MarshalJSON() ([]byte, error) {
	type wire Project
	if v.IsFavorite == nil {
		x := false
		v.IsFavorite = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Evaluations != nil {
		data, err := json.Marshal(v.Evaluations)
		if err != nil {
			return nil, err
		}
		fields["evaluations"] = data
	}
	return json.Marshal(fields)
}

type ProjectEvaluationSummary struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type Projects struct {
	Projects PaginatedResultProject `json:"projects"`
}
type UnknownResourceChunkContent struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownResourceChunkContent) MarshalJSON() ([]byte, error) {
	type wire UnknownResourceChunkContent
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type ResourceChunk struct {
	Uri     string               `json:"uri"`
	Content ResourceChunkContent `json:"content"`
	Type    string               `json:"type,omitempty"`
}

func (v ResourceChunk) MarshalJSON() ([]byte, error) {
	type wire ResourceChunk
	if v.Type == "" {
		v.Type = "resource"
	}
	return json.Marshal(wire(v))
}

type ResourceLinkChunk struct {
	Uri      string              `json:"uri"`
	Type     string              `json:"type,omitempty"`
	Metadata map[string]Metadata `json:"metadata,omitempty"`
}

func (v ResourceLinkChunk) MarshalJSON() ([]byte, error) {
	type wire ResourceLinkChunk
	if v.Type == "" {
		v.Type = "resource_link"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Metadata != nil {
		data, err := json.Marshal(v.Metadata)
		if err != nil {
			return nil, err
		}
		fields["metadata"] = data
	}
	return json.Marshal(fields)
}

type RunEvaluationSummary struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type RunOptimizationRef struct {
	OptimizationID string `json:"optimization_id"`
	Name           string `json:"name"`
}
type RunOwnershipFilter struct {
	Combinator  *FilterCombinator `json:"combinator,omitempty"`
	CreatorIds  []string          `json:"creator_ids,omitempty"`
	ProjectIds  []string          `json:"project_ids,omitempty"`
	StarredByMe *bool             `json:"starred_by_me,omitempty"`
}

func (v RunOwnershipFilter) MarshalJSON() ([]byte, error) {
	type wire RunOwnershipFilter
	if v.StarredByMe == nil {
		x := false
		v.StarredByMe = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.CreatorIds != nil {
		data, err := json.Marshal(v.CreatorIds)
		if err != nil {
			return nil, err
		}
		fields["creator_ids"] = data
	}
	if v.ProjectIds != nil {
		data, err := json.Marshal(v.ProjectIds)
		if err != nil {
			return nil, err
		}
		fields["project_ids"] = data
	}
	return json.Marshal(fields)
}

type RunProjectSummary struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type RunStatisticsResponse struct {
	RunID       string                         `json:"run_id"`
	Statistics  map[string]EvaluatorStatistics `json:"statistics"`
	GoalResults map[string]GoalResult          `json:"goal_results,omitempty"`
	Passed      *bool                          `json:"passed,omitempty"`
}

func (v RunStatisticsResponse) MarshalJSON() ([]byte, error) {
	type wire RunStatisticsResponse
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.GoalResults != nil {
		data, err := json.Marshal(v.GoalResults)
		if err != nil {
			return nil, err
		}
		fields["goal_results"] = data
	}
	return json.Marshal(fields)
}

type ScalarStatisticValue struct {
	ID    string   `json:"id"`
	Type  *string  `json:"type,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

func (v ScalarStatisticValue) MarshalJSON() ([]byte, error) {
	type wire ScalarStatisticValue
	if v.Type == nil {
		x := "scalar"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ScoreFilter struct {
	EvaluatorName string           `json:"evaluator_name"`
	Metric        string           `json:"metric"`
	Operator      ScoreOperator    `json:"operator"`
	Value         ScoreFilterValue `json:"value"`
}
type ScoreFilterGroup struct {
	Combinator *FilterCombinator `json:"combinator,omitempty"`
	Filters    []ScoreFilter     `json:"filters,omitempty"`
}

func (v ScoreFilterGroup) MarshalJSON() ([]byte, error) {
	type wire ScoreFilterGroup
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Filters != nil {
		data, err := json.Marshal(v.Filters)
		if err != nil {
			return nil, err
		}
		fields["filters"] = data
	}
	return json.Marshal(fields)
}

type ScoreSortInput struct {
	EvaluatorName string              `json:"evaluator_name"`
	Kind          string              `json:"kind,omitempty"`
	Metric        *NumericScoreMetric `json:"metric,omitempty"`
	Direction     *SortDirection      `json:"direction,omitempty"`
}

func (v ScoreSortInput) MarshalJSON() ([]byte, error) {
	type wire ScoreSortInput
	if v.Kind == "" {
		v.Kind = "score"
	}
	return json.Marshal(wire(v))
}

type SearchAllEvaluationRunsResponse struct {
	Runs PaginatedResultEvaluationRunV2WithRefsResponse `json:"runs"`
}
type SimpleNumericStatisticSpec struct {
	Kind SimpleNumericStatisticSpecKind `json:"kind"`
	Goal *GoalSpec                      `json:"goal,omitempty"`
}
type SystemFilter struct {
	Name   *SystemNameFilterGroup  `json:"name,omitempty"`
	Params *SystemParamFilterGroup `json:"params,omitempty"`
}
type SystemNameFilter struct {
	Value    string                 `json:"value"`
	Operator *ContentSearchOperator `json:"operator,omitempty"`
}
type SystemNameFilterGroup struct {
	Combinator *FilterCombinator  `json:"combinator,omitempty"`
	Filters    []SystemNameFilter `json:"filters,omitempty"`
}

func (v SystemNameFilterGroup) MarshalJSON() ([]byte, error) {
	type wire SystemNameFilterGroup
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Filters != nil {
		data, err := json.Marshal(v.Filters)
		if err != nil {
			return nil, err
		}
		fields["filters"] = data
	}
	return json.Marshal(fields)
}

type SystemParamFilter struct {
	Key      string                 `json:"key"`
	Operator SystemParamOperator    `json:"operator"`
	Value    SystemParamFilterValue `json:"value"`
}
type SystemParamFilterGroup struct {
	Combinator *FilterCombinator   `json:"combinator,omitempty"`
	Filters    []SystemParamFilter `json:"filters,omitempty"`
}

func (v SystemParamFilterGroup) MarshalJSON() ([]byte, error) {
	type wire SystemParamFilterGroup
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Filters != nil {
		data, err := json.Marshal(v.Filters)
		if err != nil {
			return nil, err
		}
		fields["filters"] = data
	}
	return json.Marshal(fields)
}

type SystemV2 struct {
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt *time.Time        `json:"deleted_at"`
	Name      string            `json:"name"`
	Params    map[string]Params `json:"params"`
}
type TagFilter struct {
	Combinator *FilterCombinator `json:"combinator,omitempty"`
	Values     []string          `json:"values,omitempty"`
}

func (v TagFilter) MarshalJSON() ([]byte, error) {
	type wire TagFilter
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Values != nil {
		data, err := json.Marshal(v.Values)
		if err != nil {
			return nil, err
		}
		fields["values"] = data
	}
	return json.Marshal(fields)
}

type UnsupportedStatistics struct {
	Count int     `json:"count"`
	Type  *string `json:"type,omitempty"`
}

func (v UnsupportedStatistics) MarshalJSON() ([]byte, error) {
	type wire UnsupportedStatistics
	if v.Type == nil {
		x := "unsupported"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type WhoamiOrganization struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	OrgTier  *string `json:"org_tier,omitempty"`
	Internal *bool   `json:"internal,omitempty"`
	Kind     *string `json:"kind,omitempty"`
}

func (v WhoamiOrganization) MarshalJSON() ([]byte, error) {
	type wire WhoamiOrganization
	if v.Internal == nil {
		x := false
		v.Internal = &x
	}
	if v.Kind == nil {
		x := "S"
		v.Kind = &x
	}
	return json.Marshal(wire(v))
}

type UnknownPrincipal struct {
	PrincipalType string `json:"principal_type,omitempty"`
	Raw           any    `json:"raw"`
	IsUnknown     bool   `json:"is_unknown,omitempty"`
}

func (v UnknownPrincipal) MarshalJSON() ([]byte, error) {
	type wire UnknownPrincipal
	if v.PrincipalType == "" {
		v.PrincipalType = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type WhoamiResponse struct {
	Principal    Principal           `json:"principal"`
	Organization *WhoamiOrganization `json:"organization,omitempty"`
	Workspace    any                 `json:"workspace,omitempty"`
}
type WhoamiServiceAccountPrincipal struct {
	ServiceAccountID string  `json:"service_account_id"`
	PrincipalType    string  `json:"principal_type,omitempty"`
	Name             *string `json:"name,omitempty"`
}

func (v WhoamiServiceAccountPrincipal) MarshalJSON() ([]byte, error) {
	type wire WhoamiServiceAccountPrincipal
	if v.PrincipalType == "" {
		v.PrincipalType = "service_account"
	}
	return json.Marshal(wire(v))
}

type WhoamiUserPrincipal struct {
	ID            string  `json:"id"`
	PrincipalType string  `json:"principal_type,omitempty"`
	Email         *string `json:"email,omitempty"`
	FirstName     *string `json:"first_name,omitempty"`
	LastName      *string `json:"last_name,omitempty"`
}

func (v WhoamiUserPrincipal) MarshalJSON() ([]byte, error) {
	type wire WhoamiUserPrincipal
	if v.PrincipalType == "" {
		v.PrincipalType = "user"
	}
	return json.Marshal(wire(v))
}
