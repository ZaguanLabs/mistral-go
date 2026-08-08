package sdk

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestV291ParityEndpoints(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		call   func(*MistralClient) error
	}{
		{"CreatePrompt", http.MethodPost, "/v2/prompts", func(c *MistralClient) error {
			_, err := c.CreatePrompt(&CreatePromptRequest{Name: "prompt", Definition: PromptDefinition{Content: "hello"}})
			return err
		}},
		{"CreateSkill", http.MethodPost, "/v2/skills", func(c *MistralClient) error {
			_, err := c.CreateSkill(&CreateSkillRequest{Name: "skill", Definition: SkillDefinition{}})
			return err
		}},
		{"GetUserIdentity", http.MethodGet, "/v1/users/me", func(c *MistralClient) error {
			_, err := c.GetUserIdentity()
			return err
		}},
		{"ActivateConnectorForConsumer", http.MethodPost, "/v1/connectors/connector/team/activate", func(c *MistralClient) error {
			_, err := c.ActivateConnectorForConsumer("connector", "team", nil)
			return err
		}},
		{"ShareConnector", http.MethodPut, "/v1/connectors/connector/share", func(c *MistralClient) error {
			_, err := c.ShareConnector("connector", map[string]any{"scope": "workspace"})
			return err
		}},
		{"CreateWorkflowDeployment", http.MethodPost, "/v1/workflows/deployments", func(c *MistralClient) error {
			_, err := c.CreateWorkflowDeployment(&CreateWorkflowDeploymentRequest{Name: "deployment", Spec: DeploymentWorkerSpec{GitHubURL: "https://example.com/repo"}})
			return err
		}},
		{"StartWorkflowDeployment", http.MethodPost, "/v1/workflows/deployments/deployment/start", func(c *MistralClient) error {
			_, err := c.StartWorkflowDeployment("deployment")
			return err
		}},
		{"GetWorkflowExecutionTraceInfo", http.MethodGet, "/v1/workflows/executions/execution/trace/info", func(c *MistralClient) error {
			_, err := c.GetWorkflowExecutionTraceInfo("execution")
			return err
		}},
		{"AggregateSpans", http.MethodPost, "/v1/observability/spans/aggregate", func(c *MistralClient) error {
			_, err := c.AggregateSpans(&ObservabilityAggregationParams{Metric: map[string]any{"count": true}})
			return err
		}},
		{"AggregateTraces", http.MethodPost, "/v1/observability/traces/aggregate", func(c *MistralClient) error {
			_, err := c.AggregateTraces(&ObservabilityAggregationParams{Metric: map[string]any{"count": true}})
			return err
		}},
		{"RegisterRAGDeployment", http.MethodPut, "/v1/rag/deployments", func(c *MistralClient) error {
			_, err := c.RegisterRAGDeployment(&RegisterRAGDeploymentRequest{Name: "rag", Deployment: map[string]any{}})
			return err
		}},
		{"UpdateRAGDeploymentMetrics", http.MethodPut, "/v1/rag/deployments/deployment/metrics", func(c *MistralClient) error {
			_, err := c.UpdateRAGDeploymentMetrics("deployment", &UpdateRAGDeploymentMetricsRequest{Status: "online"})
			return err
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method {
					t.Errorf("expected method %s, got %s", tc.method, r.Method)
				}
				if r.URL.Path != tc.path {
					t.Errorf("expected path %s, got %s", tc.path, r.URL.Path)
				}
				MockJSONResponse(http.StatusOK, `{}`).Write(w)
			})
			defer mock.Close()

			if err := tc.call(mock.GetClient()); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestV291StreamDisconnectedError(t *testing.T) {
	body := io.NopCloser(strings.NewReader("event: error\ndata: {\"reason\":\"server_shutdown\",\"error\":\"try again\"}\n\n"))
	event, ok := <-parseGenericStream(body)
	if !ok {
		t.Fatal("expected a stream event")
	}

	var disconnected *StreamDisconnectedError
	if !errors.As(event.Error, &disconnected) {
		t.Fatalf("expected StreamDisconnectedError, got %T", event.Error)
	}
	if disconnected.Reason != "server_shutdown" || disconnected.ErrorMessage != "try again" {
		t.Fatalf("unexpected stream error: %+v", disconnected)
	}
}
