package sdk

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestV293NewEndpoints(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		query map[string]string
		body  map[string]any
		call  func(*MistralClient) error
	}{
		{
			name: "CreateRealtimeSession", path: "/v1/client/sessions",
			body: map[string]any{"model": "voxtral-mini-latest", "purpose": "realtime"},
			call: func(c *MistralClient) error {
				_, err := c.CreateRealtimeSession(&CreateRealtimeSessionRequest{Model: "voxtral-mini-latest"})
				return err
			},
		},
		{
			name: "ListOrganizations", path: "/v1/users/me/organizations",
			query: map[string]string{"offset": "0", "limit": "100"},
			call: func(c *MistralClient) error {
				_, err := c.ListOrganizations(nil)
				return err
			},
		},
		{
			name: "ListWorkspaces", path: "/v1/users/me/workspaces",
			query: map[string]string{"organization_id": "org", "offset": "0", "limit": "100"},
			call: func(c *MistralClient) error {
				org := "org"
				_, err := c.ListWorkspaces(&ListWorkspacesParams{OrganizationID: &org})
				return err
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tc.path {
					t.Errorf("expected path %s, got %s", tc.path, r.URL.Path)
				}
				for key, want := range tc.query {
					if got := r.URL.Query().Get(key); got != want {
						t.Errorf("query %s: expected %q, got %q", key, want, got)
					}
				}
				if tc.body != nil {
					var got map[string]any
					if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
						t.Fatalf("decode request: %v", err)
					}
					for key, want := range tc.body {
						if got[key] != want {
							t.Errorf("body %s: expected %#v, got %#v", key, want, got[key])
						}
					}
				}
				response := `{}`
				if tc.name == "CreateRealtimeSession" {
					response = `{"purpose":"realtime","expires_at":"2026-08-17T12:00:00Z","client_secret":{"value":"secret","expires_at":"2026-08-17T12:00:00Z"},"object":"client.session"}`
				} else if tc.name == "ListOrganizations" {
					response = `{"organizations":[]}`
				} else if tc.name == "ListWorkspaces" {
					response = `{"workspaces":[]}`
				}
				MockJSONResponse(http.StatusOK, response).Write(w)
			})
			defer mock.Close()

			if err := tc.call(mock.GetClient()); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestV293RequestAdditions(t *testing.T) {
	tier := RequestedServiceTierStandardOnly
	workflowName := "workflow"
	secret := true

	tests := []struct {
		name   string
		method string
		path   string
		assert func(*testing.T, *http.Request, map[string]any)
		call   func(*MistralClient) error
	}{
		{
			name: "ChatServiceTier", method: http.MethodPost, path: "/v1/chat/completions",
			assert: func(t *testing.T, _ *http.Request, body map[string]any) {
				if body["service_tier"] != "standard_only" {
					t.Fatalf("service_tier = %#v", body["service_tier"])
				}
			},
			call: func(c *MistralClient) error {
				_, err := c.Chat("model", []ChatMessage{{Role: RoleUser, Content: "hi"}}, &ChatRequestParams{ServiceTier: &tier})
				return err
			},
		},
		{
			name: "AgentServiceTier", method: http.MethodPost, path: "/v1/agents/completions",
			assert: func(t *testing.T, _ *http.Request, body map[string]any) {
				if body["service_tier"] != "standard_only" {
					t.Fatalf("service_tier = %#v", body["service_tier"])
				}
			},
			call: func(c *MistralClient) error {
				_, err := c.AgentComplete("agent", []ChatMessage{{Role: RoleUser, Content: "hi"}}, &AgentCompletionRequest{ServiceTier: &tier})
				return err
			},
		},
		{
			name: "ConnectorGlobalHeaders", method: http.MethodPost, path: "/v1/connectors",
			assert: func(t *testing.T, _ *http.Request, body map[string]any) {
				headers, ok := body["global_headers"].(map[string]any)
				if !ok || headers["X-Global"] == nil {
					t.Fatalf("global_headers = %#v", body["global_headers"])
				}
			},
			call: func(c *MistralClient) error {
				_, err := c.CreateConnector(&ConnectorRequest{Name: "connector", GlobalHeaders: map[string]GlobalHeaderValue{"X-Global": {Value: "value", IsSecret: &secret}}})
				return err
			},
		},
		{
			name: "ConnectorAuthMethods", method: http.MethodPatch, path: "/v1/connectors/id",
			assert: func(t *testing.T, _ *http.Request, body map[string]any) {
				methods, ok := body["auth_methods"].([]any)
				if !ok || len(methods) != 1 {
					t.Fatalf("auth_methods = %#v", body["auth_methods"])
				}
			},
			call: func(c *MistralClient) error {
				_, err := c.UpdateConnector("id", &UpdateConnectorRequest{AuthMethods: []AuthenticationMethodCreateOrUpdateRequest{{MethodType: ConnectorAuthenticationOAuth2}}})
				return err
			},
		},
		{
			name: "IngestionTargetIndexes", method: http.MethodPut, path: "/v1/rag/ingestion_pipeline_configurations",
			assert: func(t *testing.T, _ *http.Request, body map[string]any) {
				targets, ok := body["target_indexes"].([]any)
				if !ok || len(targets) != 1 || targets[0].(map[string]any)["type"] != "vespa" {
					t.Fatalf("target_indexes = %#v", body["target_indexes"])
				}
			},
			call: func(c *MistralClient) error {
				_, err := c.RegisterIngestionPipelineConfiguration(&RegisterIngestionPipelineConfigurationRequest{Name: "pipeline", PipelineComposition: map[string]any{}, TargetIndexes: []IngestionPipelineTargetIndexRef{{Name: "index"}}})
				return err
			},
		},
		{
			name: "DeploymentWorkflowName", method: http.MethodGet, path: "/v1/workflows/deployments/deployment",
			assert: func(t *testing.T, r *http.Request, _ map[string]any) {
				if got := r.URL.Query().Get("workflow_name"); got != workflowName {
					t.Fatalf("workflow_name = %q", got)
				}
			},
			call: func(c *MistralClient) error {
				_, err := c.GetWorkflowDeployment("deployment", &workflowName)
				return err
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.Path != tc.path {
					t.Fatalf("expected %s %s, got %s %s", tc.method, tc.path, r.Method, r.URL.Path)
				}
				var body map[string]any
				if r.Body != nil && r.Method != http.MethodGet {
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Fatalf("decode request: %v", err)
					}
				}
				tc.assert(t, r, body)
				MockJSONResponse(http.StatusOK, `{}`).Write(w)
			})
			defer mock.Close()

			if err := tc.call(mock.GetClient()); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestV293ResponseModels(t *testing.T) {
	payload := `{"prompt_tokens":1,"total_tokens":2,"prompt_audio_seconds":3,"service_tier":"auto"}`
	var usage UsageInfo
	if err := json.Unmarshal([]byte(payload), &usage); err != nil {
		t.Fatal(err)
	}
	if usage.PromptAudioSeconds == nil || *usage.PromptAudioSeconds != 3 || usage.ServiceTier == nil || *usage.ServiceTier != "auto" {
		t.Fatalf("unexpected usage: %+v", usage)
	}

	eventPayload := `{"id":"event","continued_run_id":"next","first_execution_run_id":"first","schedule_id":"schedule"}`
	var event WorkflowEventResponse
	if err := json.Unmarshal([]byte(eventPayload), &event); err != nil {
		t.Fatal(err)
	}
	if event.ContinuedRunID == nil || event.FirstExecutionRunID == nil || event.ScheduleID == nil {
		t.Fatalf("missing workflow linkage fields: %+v", event)
	}
}
