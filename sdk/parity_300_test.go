package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPython300Endpoints(t *testing.T) {
	zero := 0
	disabled := false
	order := Order("asc")
	status := ManagedIndexStatus("ready")
	logType := DeploymentLogBuild
	cases := []struct {
		name, method, path string
		query              map[string][]string
		body               map[string]any
		response           string
		call               func(*MistralClient) error
	}{
		{name: "create pipeline", method: "POST", path: "/v1/observability/pipelines", body: map[string]any{"name": "eval", "selectors": []any{}, "definitions": []any{map[string]any{"slug": "judge"}}, "enabled": true}, call: func(c *MistralClient) error {
			_, e := c.CreatePipeline(&CreatePipelineRequest{Name: "eval", Selectors: []PipelineConfigSelector{}, Definitions: []JudgeDefinition{{Slug: "judge"}}})
			return e
		}},
		{name: "list pipelines", method: "GET", path: "/v1/observability/pipelines", query: map[string][]string{"page": {"1"}, "page_size": {"50"}, "enabled": {"false"}}, call: func(c *MistralClient) error {
			_, e := c.ListPipelines(&ListPipelinesParams{Enabled: &disabled})
			return e
		}},
		{name: "get pipeline", method: "GET", path: "/v1/observability/pipelines/a%2Fb", call: func(c *MistralClient) error { _, e := c.GetPipeline("a/b"); return e }},
		{name: "update pipeline", method: "PUT", path: "/v1/observability/pipelines/id", body: map[string]any{"name": "eval", "selectors": []any{}, "definitions": []any{}, "enabled": false}, call: func(c *MistralClient) error {
			_, e := c.UpdatePipeline("id", &UpdatePipelineRequest{Name: "eval", Selectors: []PipelineConfigSelector{}, Definitions: []JudgeDefinition{}})
			return e
		}},
		{name: "delete pipeline", method: "DELETE", path: "/v1/observability/pipelines/id", response: "204", call: func(c *MistralClient) error { return c.DeletePipeline("id") }},
		{name: "navigate", method: "POST", path: "/v1/rag/managed_indexes/index/navigate", body: map[string]any{"source_id": "src", "start_offset": float64(0), "end_offset": float64(4), "direction": "next", "top_k": float64(1), "content_type": "content"}, call: func(c *MistralClient) error {
			_, e := c.NavigateManagedIndex("index", &NavigateRequest{SourceID: "src", StartOffset: 0, EndOffset: 4, Direction: NavigationNext})
			return e
		}},
		{name: "read", method: "POST", path: "/v1/rag/managed_indexes/index/read", body: map[string]any{"source_id": "src", "start_offset": float64(0), "top_k": float64(20), "content_type": "content"}, call: func(c *MistralClient) error {
			_, e := c.ReadManagedIndex("index", &ReadRequest{SourceID: "src", StartOffset: &zero})
			return e
		}},
		{name: "grep", method: "POST", path: "/v1/rag/managed_indexes/index/grep", body: map[string]any{"source_id": "src", "pattern": "text", "top_k": float64(5), "content_type": "content"}, call: func(c *MistralClient) error {
			_, e := c.GrepManagedIndex("index", &GrepRequest{SourceID: "src", Pattern: "text"})
			return e
		}},
		{name: "chunk", method: "GET", path: "/v1/rag/managed_indexes/a%2Fb/chunks/c%2Fd", response: `{"id":"chunk","extension":42}`, call: func(c *MistralClient) error {
			v, e := c.GetManagedIndexChunk("a/b", "c/d")
			if e == nil && string(v.Extra["extension"]) != "42" {
				t.Error(v)
			}
			return e
		}},
		{name: "indexes", method: "GET", path: "/v1/rag/managed_indexes", query: map[string][]string{"page_size": {"20"}, "name": {"index"}, "status": {"ready"}, "creator_id": {"creator"}}, call: func(c *MistralClient) error {
			_, e := c.ListManagedIndexes(&ListManagedIndexesParams{Name: StringPtr("index"), Status: &status, CreatorID: StringPtr("creator")})
			return e
		}},
		{name: "credentials", method: "DELETE", path: "/v1/connectors/a%2Fb/workspace/credentials/c%2Fd", call: func(c *MistralClient) error {
			_, e := c.DeleteConnectorCredentials("a/b", "workspace", "c/d")
			return e
		}},
		{name: "connector", method: "GET", path: "/v1/connectors/id", query: map[string][]string{"fetch_user_data": {"false"}}, call: func(c *MistralClient) error { _, e := c.GetConnector("id", &disabled); return e }},
		{name: "tools", method: "GET", path: "/v1/connectors/id/tools", query: map[string][]string{"page_size": {"0"}}, call: func(c *MistralClient) error {
			_, e := c.ListConnectorTools("id", &ListConnectorToolsParams{Page: &zero, PageSize: &zero})
			return e
		}},
		{name: "import spans", method: "POST", path: "/v1/observability/datasets/id/imports/from-spans", body: map[string]any{"span_references": []any{map[string]any{"trace_id": "trace", "span_id": "span"}}, "mapping_contract": map[string]any{"version": float64(1), "mappings": []any{}}}, response: `{"result":{"requested_record_count":2,"imported_record_count":1,"skipped_record_count":1}}`, call: func(c *MistralClient) error {
			v, e := c.ImportDatasetFromSpans("id", &ImportDatasetFromSpansRequest{SpanReferences: []TelemetrySpanReference{{TraceID: "trace", SpanID: "span"}}, MappingContract: SpanDatasetMappingContract{Mappings: []SpanDatasetMapping{}}})
			if e == nil && (v.Result == nil || v.Result.ImportedRecordCount != 1) {
				t.Error(v)
			}
			return e
		}},
		{name: "aggregate evaluations", method: "POST", path: "/v1/observability/spans/evaluations/aggregate", call: func(c *MistralClient) error {
			_, e := c.AggregateSpanEvaluations(&ObservabilityAggregationParams{Metric: map[string]any{"type": "count"}})
			return e
		}},
		{name: "unharden", method: "POST", path: "/v1/workflows/deployments/id/unharden", query: map[string][]string{"workspace_id": {"ws"}}, response: "204", call: func(c *MistralClient) error { return c.UnhardenWorkflowDeployment("id", StringPtr("ws")) }},
		{name: "deployments", method: "GET", path: "/v1/workflows/deployments", query: map[string][]string{"owner": {"me"}}, call: func(c *MistralClient) error {
			_, e := c.ListWorkflowDeploymentsWithParams(&ListWorkflowDeploymentsParams{Owner: StringPtr("me")})
			return e
		}},
		{name: "deployment logs", method: "GET", path: "/v1/workflows/deployments/id/logs", query: map[string][]string{"log_type": {"build"}}, call: func(c *MistralClient) error {
			_, e := c.GetDeploymentLogs("id", &DeploymentLogsParams{LogType: &logType})
			return e
		}},
		{name: "service accounts", method: "GET", path: "/v1/service-accounts", query: map[string][]string{"offset": {"0"}, "limit": {"10"}, "include_deleted": {"false"}, "q": {"bot"}, "order": {"asc"}}, call: func(c *MistralClient) error {
			_, e := c.ListServiceAccounts(&ListServiceAccountsParams{Limit: 10, Q: StringPtr("bot"), Order: &order})
			return e
		}},
		{name: "voice search", method: "GET", path: "/v2/audio/voices", query: map[string][]string{"page_size": {"10"}, "type": {"all"}, "gender": {"neutral", "female"}, "language": {"en"}, "page_token": {"next"}}, response: `{"data":[{"id":"voice"}],"next_page_token":"more"}`, call: func(c *MistralClient) error {
			v, e := c.SearchVoices(&SearchVoicesParams{Gender: []VoiceGender{VoiceGenderNeutral, VoiceGenderFemale}, Language: []string{"en"}, PageToken: StringPtr("next")})
			if e == nil && (v.NextPageToken == nil || *v.NextPageToken != "more") {
				t.Error(v)
			}
			return e
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.EscapedPath() != tc.path {
					t.Errorf("%s %s", r.Method, r.URL.EscapedPath())
				}
				if len(r.URL.Query()) != len(tc.query) {
					t.Errorf("query %v want %v", r.URL.Query(), tc.query)
				}
				for key, want := range tc.query {
					if !reflect.DeepEqual(r.URL.Query()[key], want) {
						t.Errorf("query %s: %v want %v", key, r.URL.Query()[key], want)
					}
				}
				if tc.body != nil {
					var body map[string]any
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Error(err)
					}
					if !reflect.DeepEqual(body, tc.body) {
						t.Errorf("body %v want %v", body, tc.body)
					}
				}
				if tc.response == "204" {
					w.WriteHeader(204)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				response := tc.response
				if response == "" {
					response = "{}"
				}
				w.Write([]byte(response))
			}))
			defer server.Close()
			if err := tc.call(NewMistralClient("key", server.URL, 1, 0)); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPython300ModelsAndStructuredContent(t *testing.T) {
	retriever := RRFRetriever{Retrievers: []any{KeywordRetriever{Query: "hello"}, NearestNeighbourRetriever{Query: StringPtr("hello"), MaxCandidates: IntPtr(50)}}, Weights: []float64{1, 2}}
	raw, err := json.Marshal(SearchRequest{Retriever: retriever})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"type":"rrf"`, `"rank_constant":60`, `"top_k":20`, `"max_candidates":50`} {
		if !strings.Contains(string(raw), want) {
			t.Errorf("missing %s: %s", want, raw)
		}
	}
	raw, err = json.Marshal(DeploymentWorkerSpec{BackendSpec: DeploymentMistralCloudBackendSpec{Secrets: []DeploymentSecretBinding{{EnvVarName: "TOKEN", Reference: "secret"}}}, Entrypoint: StringPtr("removed")})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "entrypoint") || !strings.Contains(string(raw), `"type":"mistral_cloud"`) {
		t.Fatal(string(raw))
	}
	if _, err = json.Marshal(CreatePipelineConfigRequest{Definitions: []PipelineConfigDefinition{JudgeDefinition{Slug: "one"}, JudgeDefinition{Slug: "two"}}}); err == nil {
		t.Fatal("ambiguous legacy definition list accepted")
	}
	for _, content := range []string{`"{\"value\":7}"`, `[{"type":"text","text":"{\"value\":"},{"type":"image_url","image_url":"ignore"},{"type":"text","text":"7}"}]`} {
		var response ChatCompletionResponse
		if err := json.Unmarshal([]byte(`{"choices":[{"message":{"role":"assistant","content":`+content+`}}]}`), &response); err != nil {
			t.Fatal(err)
		}
		parsed, err := ParseChatCompletion[struct {
			Value int `json:"value"`
		}](&response)
		if err != nil || parsed.Choices[0].Parsed == nil || parsed.Choices[0].Parsed.Value != 7 {
			t.Fatalf("%+v %v", parsed, err)
		}
	}
	for _, tool := range []ToolType{ToolTypeWebSearch, ToolTypeWebSearchPrem, ToolTypeCodeInterpreter} {
		if err := validateCompletionTools([]any{BuiltInTool{Type: tool}}); err == nil {
			t.Errorf("removed tool accepted: %s", tool)
		}
	}
}
