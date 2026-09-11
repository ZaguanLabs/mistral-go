package sdk

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestPython210Endpoints(t *testing.T) {
	kind := PipelineKindJudge
	disabled := false
	cases := []struct {
		name, method, path, response string
		query                        map[string][]string
		body                         map[string]any
		call                         func(*MistralClient) error
	}{
		{name: "CreateServiceAccount", method: "POST", path: "/v1/service-accounts", body: map[string]any{"name": "bot", "workspace_id": "workspace"}, call: func(c *MistralClient) error {
			_, e := c.CreateServiceAccount(&CreateServiceAccountRequest{Name: "bot", WorkspaceID: "workspace"})
			return e
		}},
		{name: "ListServiceAccounts", method: "GET", path: "/v1/service-accounts", query: map[string][]string{"workspace_id": {"ws"}, "offset": {"0"}, "limit": {"20"}, "include_deleted": {"false"}}, response: `{"items":[{"id":"bot","description":null,"deleted_at":null}]}`, call: func(c *MistralClient) error {
			r, e := c.ListServiceAccounts(&ListServiceAccountsParams{WorkspaceID: "ws", Limit: 20})
			if e == nil && (len(r.Items) != 1 || r.Items[0].ID != "bot") {
				t.Errorf("response: %+v", r)
			}
			return e
		}},
		{name: "GetServiceAccount", method: "GET", path: "/v1/service-accounts/id", call: func(c *MistralClient) error { _, e := c.GetServiceAccount("id"); return e }},
		{name: "UpdateServiceAccount", method: "PATCH", path: "/v1/service-accounts/id", body: map[string]any{"description": nil}, call: func(c *MistralClient) error {
			_, e := c.UpdateServiceAccount("id", &UpdateServiceAccountRequest{})
			return e
		}},
		{name: "DeleteServiceAccount", method: "DELETE", path: "/v1/service-accounts/id", response: "204", call: func(c *MistralClient) error { return c.DeleteServiceAccount("id") }},
		{name: "ListAssignableRoles", method: "GET", path: "/v1/service-accounts/assignable-roles", query: map[string][]string{"workspace_id": {"ws"}}, call: func(c *MistralClient) error { _, e := c.ListAssignableServiceAccountRoles("ws"); return e }},
		{name: "ListRoles", method: "GET", path: "/v1/service-accounts/id/roles", call: func(c *MistralClient) error { _, e := c.ListServiceAccountRoles("id"); return e }},
		{name: "SetRoles", method: "PUT", path: "/v1/service-accounts/id/roles", body: map[string]any{"role_ids": []any{}}, call: func(c *MistralClient) error { _, e := c.SetServiceAccountRoles("id", nil); return e }},
		{name: "CreatePipeline", method: "POST", path: "/v1/observability/pipeline-configs", body: map[string]any{"pipeline_kind": "judge", "selectors": []any{}, "definition": map[string]any{"model": "m", "prompt": "p"}, "name": "eval", "enabled": true}, call: func(c *MistralClient) error {
			_, e := c.CreatePipelineConfig(&CreatePipelineConfigRequest{PipelineKind: kind, Selectors: []PipelineConfigSelector{}, Definition: JudgeDefinition{Model: "m", Prompt: "p"}, Name: "eval"})
			return e
		}},
		{name: "ListPipelines", method: "GET", path: "/v1/observability/pipeline-configs", query: map[string][]string{"page": {"1"}, "page_size": {"50"}, "pipeline_kind": {"judge"}, "enabled": {"false"}}, response: `{"pipeline_configs":{"count":1,"results":[{"id":"pipeline"}]}}`, call: func(c *MistralClient) error {
			r, e := c.ListPipelineConfigs(&ListPipelineConfigsParams{PipelineKind: &kind, Enabled: &disabled})
			if e == nil && r.PipelineConfigs.Count != 1 {
				t.Errorf("response: %+v", r)
			}
			return e
		}},
		{name: "GetPipeline", method: "GET", path: "/v1/observability/pipeline-configs/id", call: func(c *MistralClient) error { _, e := c.GetPipelineConfig("id"); return e }},
		{name: "UpdatePipeline", method: "PUT", path: "/v1/observability/pipeline-configs/id", body: map[string]any{"pipeline_kind": "judge", "selectors": []any{}, "definition": map[string]any{"model": "m", "prompt": "p"}, "name": "eval", "enabled": false}, call: func(c *MistralClient) error {
			_, e := c.UpdatePipelineConfig("id", &UpdatePipelineConfigRequest{PipelineKind: kind, Selectors: []PipelineConfigSelector{}, Definition: JudgeDefinition{Model: "m", Prompt: "p"}, Name: "eval"})
			return e
		}},
		{name: "DeletePipeline", method: "DELETE", path: "/v1/observability/pipeline-configs/id", response: "204", call: func(c *MistralClient) error { return c.DeletePipelineConfig("id") }},
		{name: "ShareConnector", method: "PUT", path: "/v1/connectors/id/organization/share", call: func(c *MistralClient) error { _, e := c.ShareConnectorToOrganization("id"); return e }},
		{name: "UnshareConnector", method: "DELETE", path: "/v1/connectors/id/organization/share", call: func(c *MistralClient) error { _, e := c.UnshareConnectorFromOrganization("id"); return e }},
		{name: "CreateCredentials", method: "POST", path: "/v1/connectors/id/workspace/credentials", body: map[string]any{"name": "cred", "credentials": map[string]any{"oauth": map[string]any{"client_id": "id", "client_secret": "test-secret", "grant_type": "client_credentials"}}}, call: func(c *MistralClient) error {
			_, e := c.CreateConnectorCredentials("id", "workspace", &ConnectorCredentialsRequest{Name: "cred", Credentials: ConnectionCredentialsInput{OAuth: &OAuth2ClientCredentialsInput{ClientID: "id", ClientSecret: "test-secret"}}})
			return e
		}},
		{name: "UpdateCredentials", method: "PATCH", path: "/v1/connectors/id/user/credentials", body: map[string]any{"name": "cred", "is_default": false}, call: func(c *MistralClient) error {
			_, e := c.UpdateConnectorCredentials("id", "user", &ConnectorCredentialsRequest{Name: "cred", IsDefault: &disabled})
			return e
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.Path != tc.path {
					t.Errorf("request %s %s", r.Method, r.URL)
				}
				if tc.query != nil && !reflect.DeepEqual(map[string][]string(r.URL.Query()), tc.query) {
					t.Errorf("query %v; want %v", r.URL.Query(), tc.query)
				}
				if tc.body != nil {
					var body map[string]any
					if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
						t.Error(e)
					}
					if !reflect.DeepEqual(body, tc.body) {
						t.Errorf("body %#v; want %#v", body, tc.body)
					}
				}
				if tc.response == "204" {
					w.WriteHeader(http.StatusNoContent)
					return
				}
				response := tc.response
				if response == "" {
					response = `{}`
				}
				MockJSONResponse(200, response).Write(w)
			})
			defer mock.Close()
			if err := tc.call(mock.GetClient()); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPython210ConnectorVariants(t *testing.T) {
	for _, protocol := range []string{"mcp", "http"} {
		t.Run(protocol, func(t *testing.T) {
			mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				var body map[string]any
				_ = json.NewDecoder(r.Body).Decode(&body)
				if body["protocol"] != protocol {
					t.Errorf("protocol %v", body["protocol"])
				}
				methods := body["auth_methods"].([]any)
				m := methods[0].(map[string]any)
				if m["grant_type"] != "client_credentials" || m["method_type"] != "oauth2" {
					t.Errorf("method %#v", m)
				}
				if _, exists := m["oauth2_server_metadata"].(map[string]any)["authorization_endpoint"]; exists {
					t.Error("optional authorization endpoint should be absent")
				}
				MockJSONResponse(200, `{}`).Write(w)
			})
			defer mock.Close()
			methods := []any{OAuth2ClientCredentialsAuthMethod{OAuth2ServerMetadata: &ExtendedOAuthServerMetadata{Issuer: "issuer", TokenEndpoint: "token"}, TokenEndpointAuthMethod: OAuth2TokenEndpointAuthClientSecretPost}}
			var create any = &ConnectorRequest{Name: "test", AuthMethods: methods}
			var update any = &UpdateConnectorRequest{AuthMethods: methods}
			if protocol == "http" {
				create = &CreateHTTPConnectorRequest{Name: "test", AuthMethods: methods}
				update = &UpdateHTTPConnectorRequest{AuthMethods: methods}
			}
			if _, e := mock.GetClient().CreateConnector(create); e != nil {
				t.Fatal(e)
			}
			if _, e := mock.GetClient().UpdateConnector("id", update); e != nil {
				t.Fatal(e)
			}
		})
	}
	var nilRequest *ConnectorRequest
	if _, err := connectorRequestBody(nilRequest); err == nil {
		t.Fatal("typed nil must fail")
	}
	var response Connector
	raw := `{"protocol":"future","name":"connector","extra":"preserved"}`
	if err := json.Unmarshal([]byte(raw), &response); err != nil || string(response.Raw) != raw {
		t.Fatalf("unknown connector: %+v %v", response, err)
	}
}
func TestPython210SchemaAndQuery(t *testing.T) {
	var segment TranscriptionSegment
	if e := json.Unmarshal([]byte(`{"start":null,"end":0}`), &segment); e != nil {
		t.Fatal(e)
	}
	if segment.Start != nil || segment.End == nil || *segment.End != 0 {
		t.Fatalf("nullable timestamps: %+v", segment)
	}
	raw, _ := json.Marshal(segment)
	if !strings.Contains(string(raw), `"start":null`) {
		t.Fatalf("round trip %s", raw)
	}
	relation := ShareRelationWriter
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if strings.Contains(r.URL.Path, "executions") && r.URL.Query().Get("include_search_keys") != "true" {
				t.Errorf("query: %s", r.URL)
			}
			if strings.Contains(r.URL.Path, "deployments") && (!reflect.DeepEqual(r.URL.Query()["location_types"], []string{"k8s", "managed"}) || r.URL.Query().Get("created_by") != "user") {
				t.Errorf("query: %s", r.URL)
			}
		} else {
			var body map[string]any
			_ = json.NewDecoder(r.Body).Decode(&body)
			if body["workspaceRelation"] != "writer" {
				t.Errorf("body: %v", body)
			}
		}
		MockJSONResponse(200, `{"search_keys":{"tenant":"test"}}`).Write(w)
	})
	defer mock.Close()
	c := mock.GetClient()
	if _, e := c.CreatePrompt(&CreatePromptRequest{WorkspaceRelation: &relation}); e != nil {
		t.Fatal(e)
	}
	if _, e := c.CreateSkill(&CreateSkillRequest{WorkspaceRelation: &relation}); e != nil {
		t.Fatal(e)
	}
	if _, e := c.UpdatePromptMetadata("id", &UpdatePromptMetadataRequest{WorkspaceRelation: &relation}); e != nil {
		t.Fatal(e)
	}
	if _, e := c.UpdateSkillMetadata("id", nil, relation); e != nil {
		t.Fatal(e)
	}
	r, e := c.GetWorkflowExecution("id", true)
	if e != nil || r["search_keys"] == nil {
		t.Fatalf("%v %v", r, e)
	}
	user := "user"
	if _, e := c.ListWorkflowDeploymentsWithParams(&ListWorkflowDeploymentsParams{CreatedBy: &user, LocationTypes: []string{"k8s", "managed"}}); e != nil {
		t.Fatal(e)
	}
}
