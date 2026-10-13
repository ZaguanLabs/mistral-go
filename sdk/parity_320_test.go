package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPython320EvaluationContracts(t *testing.T) {
	cases := []struct {
		method, path string
		call         func(*MistralClient) error
	}{
		{"GET", "/v1/observability/projects?page=1&page_size=50", func(c *MistralClient) error { _, e := c.ListEvaluationProjects(nil); return e }},
		{"GET", "/v1/observability/projects/a%2Fb", func(c *MistralClient) error { _, e := c.GetEvaluationProject("a/b"); return e }},
		{"GET", "/v1/observability/projects/p/evaluations_v2?page=1&page_size=50", func(c *MistralClient) error { _, e := c.ListEvaluationsV2("p", nil); return e }},
		{"GET", "/v1/observability/evaluations_v2/e", func(c *MistralClient) error { _, e := c.GetEvaluationV2("e"); return e }},
		{"POST", "/v1/observability/evaluations_v2/e/runs/search?page=1&page_size=50", func(c *MistralClient) error { _, e := c.SearchEvaluationRunsV2("e", nil, nil); return e }},
		{"POST", "/v1/observability/evaluation-runs_v2/search?page=1&page_size=50", func(c *MistralClient) error { _, e := c.SearchAllEvaluationRunsV2(nil, nil); return e }},
		{"GET", "/v1/observability/evaluation-runs_v2/r", func(c *MistralClient) error { _, e := c.GetEvaluationRunV2("r"); return e }},
		{"GET", "/v1/observability/evaluation-runs_v2/r/statistics", func(c *MistralClient) error { _, e := c.GetEvaluationRunStatisticsV2("r"); return e }},
		{"POST", "/v1/observability/evaluation-runs_v2/r/statistics/compute", func(c *MistralClient) error { _, e := c.ComputeEvaluationRunStatistics("r", nil); return e }},
		{"GET", "/v1/observability/evaluation-runs_v2/r/records/x", func(c *MistralClient) error { _, e := c.GetEvaluationRunRecordV2("r", "x"); return e }},
		{"POST", "/v1/observability/evaluation-runs_v2/r/records/search?page=1&page_size=50", func(c *MistralClient) error { _, e := c.SearchEvaluationRunRecordsV2("r", nil, nil); return e }},
		{"GET", "/v1/observability/optimizations?page=1&page_size=50&tags=one&tags=two", func(c *MistralClient) error {
			_, e := c.ListOptimizations(&ListOptimizationsParams{Tags: []string{"one", "two"}})
			return e
		}},
		{"GET", "/v1/observability/optimizations/o", func(c *MistralClient) error { _, e := c.GetOptimization("o"); return e }},
		{"GET", "/v1/observability/optimizations/o/trials?page=1&page_size=50", func(c *MistralClient) error { _, e := c.ListOptimizationTrials("o", nil); return e }},
		{"GET", "/v1/whoami", func(c *MistralClient) error { _, e := c.Whoami(); return e }},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.RequestURI() != tc.path {
					t.Errorf("%s %s want %s %s", r.Method, r.URL.RequestURI(), tc.method, tc.path)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{}`))
			}))
			defer server.Close()
			c := NewMistralClient("test", server.URL, 1, 0)
			if err := tc.call(c); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestPython320ChatAndModelContracts(t *testing.T) {
	no := false
	zero := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		if body["logprobs"] != false || body["top_k"] != float64(0) || body["prompt_logprobs"] != false {
			t.Error(body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"choices":[{"index":0,"finish_reason":"stop","logprobs":{"content":[{"token":"A","token_id":65,"bytes":[65],"logprob":-0.5}]},"prompt_logprobs":{"content":[]}}]}`))
	}))
	defer server.Close()
	c := NewMistralClient("test", server.URL, 1, 0)
	resp, err := c.Chat("model", []ChatMessage{}, &ChatRequestParams{Logprobs: &no, TopK: &zero, PromptLogprobs: &no})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Choices[0].Logprobs.Content[0].TokenID != 65 || resp.Choices[0].PromptLogprobs == nil {
		t.Fatal(resp)
	}
	for _, value := range []any{ProbabilisticSampling{Percentage: 25}, AttributeHashSampling{Attribute: "id", Percentage: 50}, ResourceChunk{Uri: "test", Content: map[string]any{"type": "text", "text": "hi"}}, WhoamiServiceAccountPrincipal{ServiceAccountID: "sa"}} {
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), `"type"`) && !strings.Contains(string(data), `"principal_type"`) {
			t.Fatal(string(data))
		}
	}
	var runs SearchAllEvaluationRunsResponse
	if err := json.Unmarshal([]byte(`{"runs":{"count":1,"results":[{"id":"run","passed":false,"optimization":{"optimization_id":"opt","name":"test"}}]}}`), &runs); err != nil {
		t.Fatal(err)
	}
	if runs.Runs.Results[0].Passed == nil || *runs.Runs.Results[0].Passed || runs.Runs.Results[0].Optimization.OptimizationID != "opt" {
		t.Fatal(runs)
	}
}

func TestPython320PartialPipelineAndSpeech(t *testing.T) {
	no := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		switch r.URL.Path {
		case "/v1/observability/pipelines/id":
			if len(body) != 1 || body["enabled"] != false {
				t.Error(body)
			}
		case "/v1/audio/speech":
			if body["phoneme_hints"].(map[string]any)["word"].([]any)[0] != "w" {
				t.Error(body)
			}
		case "/v1/observability/pipeline-results/search":
			if r.URL.Query().Get("page_size") != "50" || body["pipeline_config_slugs"].([]any)[0] != "pipe" {
				t.Error(body)
			}
		default:
			t.Error(r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	}))
	defer server.Close()
	c := NewMistralClient("test", server.URL, 1, 0)
	if _, err := c.UpdatePipeline("id", &UpdatePipelineRequest{Enabled: &no}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.Speech(&SpeechRequest{Input: "word", PhonemeHints: map[string][]IPAPhoneme{"word": {"w"}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.SearchPipelineResults(&SearchPipelineResultsParams{PipelineConfigSlugs: []string{"pipe"}}); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ContentSearchFilterGroup{Filters: []ContentSearchFilter{}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"filters":[]`) {
		t.Fatal(string(data))
	}
}
