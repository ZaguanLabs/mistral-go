package sdk

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestManagedIndexesWireContract(t *testing.T) {
	cases := []struct {
		name, method, path, query, body, response string
		call                                      func(*MistralClient) error
	}{
		{"create", "POST", "/v1/rag/managed_indexes", "", `{"name":"index","config":{"embedding":{"name":"mistral-embed","dimensions":1024,"type":"mistral"}}}`, `{"id":"i","status":"provisioning"}`, func(c *MistralClient) error {
			r, e := c.CreateManagedIndex(&CreateManagedIndexRequest{Name: "index", Config: ManagedIndexConfig{Embedding: MistralEmbeddingModel{Name: "mistral-embed", Dimensions: 1024}}})
			if e == nil && r.Status != ManagedIndexStatusProvisioning {
				t.Fatal(r)
			}
			return e
		}},
		{"list", "GET", "/v1/rag/managed_indexes", "page_size=20", "", `{"data":[],"next_page_token":"next"}`, func(c *MistralClient) error {
			r, e := c.ListManagedIndexes(nil)
			if e == nil && (r.NextPageToken == nil || *r.NextPageToken != "next") {
				t.Fatal(r)
			}
			return e
		}},
		{"next", "GET", "/v1/rag/managed_indexes", "page_size=20&page_token=next", "", `{"data":[]}`, func(c *MistralClient) error {
			token := "next"
			_, e := c.ListManagedIndexes(&ListManagedIndexesParams{PageToken: &token})
			return e
		}},
		{"get", "GET", "/v1/rag/managed_indexes/a%2Fb", "", "", `{}`, func(c *MistralClient) error { _, e := c.GetManagedIndex("a/b"); return e }},
		{"update", "PUT", "/v1/rag/managed_indexes/index", "", `{"schema":{"document_fields":{"title":{"indexed":true,"required":false,"system":false,"type":"text"}}}}`, `{}`, func(c *MistralClient) error {
			_, e := c.UpdateManagedIndex("index", &UpdateManagedIndexRequest{Schema: ManagedIndexFields{DocumentFields: map[string]any{"title": TextFieldDefinition{}}}})
			return e
		}},
		{"delete", "DELETE", "/v1/rag/managed_indexes/index", "", "", `{"status":"deleting"}`, func(c *MistralClient) error { _, e := c.DeleteManagedIndex("index"); return e }},
		{"ingest", "POST", "/v1/rag/managed_indexes/index/documents", "", `{"documents":[{"id":"doc","content":"hello"}]}`, `{"accepted":1,"rejected":0,"results":[{"status":"accepted","document_id":"doc"}]}`, func(c *MistralClient) error {
			_, e := c.IngestManagedIndexDocuments("index", &IngestDocumentsRequest{Documents: []map[string]any{{"id": "doc", "content": "hello"}}})
			return e
		}},
		{"delete documents", "DELETE", "/v1/rag/managed_indexes/index/documents", "", `{"document_ids":["doc"]}`, `{"deleted_document_ids":["doc"],"missing":[]}`, func(c *MistralClient) error {
			_, e := c.DeleteManagedIndexDocuments("index", &DeleteDocumentsRequest{DocumentIDs: []string{"doc"}})
			return e
		}},
		{"search", "POST", "/v1/rag/managed_indexes/index/search", "", `{"retriever":{"query":"hello","top_k":20,"type":"keyword","filter":{"type":"equal","field":"lang","value":"en"}}}`, `{"hits":[{"chunk":{"id":"c","extra":"kept"},"score":0.5,"index_name":"index"}]}`, func(c *MistralClient) error {
			r, e := c.SearchManagedIndex("index", &SearchRequest{Retriever: KeywordRetriever{Query: "hello", Filter: Equal{Field: "lang", Value: "en"}}})
			if e == nil && string(r.Hits[0].Chunk.Extra["extra"]) != `"kept"` {
				t.Fatal(r)
			}
			return e
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.EscapedPath() != tc.path || r.URL.RawQuery != tc.query {
					t.Errorf("request %s %s", r.Method, r.URL)
				}
				if tc.body != "" {
					var got, want any
					if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
						t.Error(err)
					}
					json.Unmarshal([]byte(tc.body), &want)
					a, _ := json.Marshal(got)
					b, _ := json.Marshal(want)
					if string(a) != string(b) {
						t.Errorf("body %s want %s", a, b)
					}
				}
				MockJSONResponse(200, tc.response).Write(w)
			})
			defer server.Close()
			if err := tc.call(server.GetClient()); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPipelineDefinitionsAndOrganizationAccounts(t *testing.T) {
	server := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "pipeline-configs") {
			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if _, exists := body["definitions"]; exists {
				t.Error("obsolete definitions list")
			}
			if _, exists := body["slug"]; exists {
				t.Error("obsolete slug")
			}
			definition, ok := body["definition"].(map[string]any)
			if !ok || definition["slug"] != "judge" {
				t.Errorf("body: %#v", body)
			}
		} else if r.Method == "GET" {
			if _, exists := r.URL.Query()["workspace_id"]; exists {
				t.Error("organization listing sent workspace")
			}
		} else {
			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			if len(body["role_ids"].([]any)) != 1 {
				t.Error(body)
			}
		}
		MockJSONResponse(200, `{}`).Write(w)
	})
	defer server.Close()
	c := server.GetClient()
	if _, err := c.CreatePipelineConfig(&CreatePipelineConfigRequest{Name: "eval", Definitions: []PipelineConfigDefinition{JudgeDefinition{Slug: "judge", Mapping: map[string]string{"x": "y"}}}}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.ListServiceAccounts(&ListServiceAccountsParams{Limit: 10}); err != nil {
		t.Fatal(err)
	}
	if _, err := c.CreateServiceAccount(&CreateServiceAccountRequest{Name: "svc", RoleIDs: []string{"role"}}); err != nil {
		t.Fatal(err)
	}
}

func TestWorkflowTraceparentBodyAndHeader(t *testing.T) {
	for _, explicit := range []string{"", "00-aabbccddeeff00112233445566778899-0102030405060708-00"} {
		t.Run(explicit, func(t *testing.T) {
			server := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				var body map[string]any
				json.NewDecoder(r.Body).Decode(&body)
				trace, _ := body["traceparent"].(string)
				if trace != r.Header.Get("traceparent") || !regexp.MustCompile(`^00-[0-9a-f]{32}-[0-9a-f]{16}-0[01]$`).MatchString(trace) {
					t.Errorf("trace body=%s header=%s", trace, r.Header.Get("traceparent"))
				}
				if explicit != "" && trace != explicit {
					t.Error("explicit trace overwritten")
				}
				MockJSONResponse(200, `{}`).Write(w)
			})
			defer server.Close()
			req := &ExecuteWorkflowRequest{Input: map[string]any{"a": 1}}
			if explicit != "" {
				req.Traceparent = &explicit
			}
			if _, err := server.GetClient().ExecuteWorkflow("wf", req); err != nil {
				t.Fatal(err)
			}
		})
	}
}
