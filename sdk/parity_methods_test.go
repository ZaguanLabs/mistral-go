package sdk

import (
	"io"
	"net/http"
	"testing"
)

func TestNewParityEndpoints(t *testing.T) {
	str := "value"
	boolVal := true
	intVal := 10

	tests := []struct {
		name   string
		method string
		path   string
		call   func(*MistralClient) error
	}{
		{"Speech", http.MethodPost, "/v1/audio/speech", func(c *MistralClient) error { _, err := c.Speech(&SpeechRequest{Input: "hello"}); return err }},
		{"ListVoices", http.MethodGet, "/v1/audio/voices", func(c *MistralClient) error { _, err := c.ListVoices(&ListVoicesParams{Limit: &intVal}); return err }},
		{"CreateVoice", http.MethodPost, "/v1/audio/voices", func(c *MistralClient) error {
			_, err := c.CreateVoice(&VoiceRequest{Name: "voice", SampleAudio: "base64"})
			return err
		}},
		{"GetVoice", http.MethodGet, "/v1/audio/voices/voice-1", func(c *MistralClient) error { _, err := c.GetVoice("voice-1"); return err }},
		{"UpdateVoice", http.MethodPatch, "/v1/audio/voices/voice-1", func(c *MistralClient) error {
			_, err := c.UpdateVoice("voice-1", &UpdateVoiceRequest{Name: &str})
			return err
		}},
		{"DeleteVoice", http.MethodDelete, "/v1/audio/voices/voice-1", func(c *MistralClient) error { _, err := c.DeleteVoice("voice-1"); return err }},
		{"CreateConnector", http.MethodPost, "/v1/connectors", func(c *MistralClient) error { _, err := c.CreateConnector(&ConnectorRequest{Name: "conn"}); return err }},
		{"ListConnectors", http.MethodGet, "/v1/connectors", func(c *MistralClient) error {
			_, err := c.ListConnectors(&ListConnectorsParams{Cursor: &str})
			return err
		}},
		{"GetConnectorAuthURL", http.MethodGet, "/v1/connectors/conn/auth_url", func(c *MistralClient) error { _, err := c.GetConnectorAuthURL("conn", &str, nil); return err }},
		{"CallConnectorTool", http.MethodPost, "/v1/connectors/conn/tools/tool/call", func(c *MistralClient) error {
			_, err := c.CallConnectorTool("conn", "tool", nil, map[string]any{"x": 1})
			return err
		}},
		{"ListConnectorTools", http.MethodGet, "/v1/connectors/conn/tools", func(c *MistralClient) error {
			_, err := c.ListConnectorTools("conn", &ListConnectorToolsParams{Refresh: &boolVal})
			return err
		}},
		{"GetConnectorAuthenticationMethods", http.MethodGet, "/v1/connectors/conn/authentication_methods", func(c *MistralClient) error { _, err := c.GetConnectorAuthenticationMethods("conn"); return err }},
		{"ListOrganizationConnectorCredentials", http.MethodGet, "/v1/connectors/conn/organization/credentials", func(c *MistralClient) error {
			_, err := c.ListOrganizationConnectorCredentials("conn", nil)
			return err
		}},
		{"CreateOrUpdateWorkspaceConnectorCredentials", http.MethodPost, "/v1/connectors/conn/workspace/credentials", func(c *MistralClient) error {
			_, err := c.CreateOrUpdateWorkspaceConnectorCredentials("conn", &ConnectorCredentialsRequest{Name: "cred"})
			return err
		}},
		{"DeleteUserConnectorCredentials", http.MethodDelete, "/v1/connectors/conn/user/credentials/cred", func(c *MistralClient) error { _, err := c.DeleteUserConnectorCredentials("conn", "cred"); return err }},
		{"GetConnector", http.MethodGet, "/v1/connectors/conn", func(c *MistralClient) error { _, err := c.GetConnector("conn", nil, nil); return err }},
		{"UpdateConnector", http.MethodPatch, "/v1/connectors/conn", func(c *MistralClient) error {
			_, err := c.UpdateConnector("conn", &UpdateConnectorRequest{Name: &str})
			return err
		}},
		{"DeleteConnector", http.MethodDelete, "/v1/connectors/conn", func(c *MistralClient) error { _, err := c.DeleteConnector("conn"); return err }},
		{"ActivateConnectorForOrganization", http.MethodPost, "/v1/connectors/conn/organization/activate", func(c *MistralClient) error {
			_, err := c.ActivateConnectorForOrganization("conn", &ToolExecutionConfiguration{Include: []string{"tool"}})
			return err
		}},
		{"DeactivateConnectorForOrganization", http.MethodPost, "/v1/connectors/conn/organization/deactivate", func(c *MistralClient) error {
			_, err := c.DeactivateConnectorForOrganization("conn")
			return err
		}},
		{"ActivateConnectorForWorkspace", http.MethodPost, "/v1/connectors/conn/workspace/activate", func(c *MistralClient) error {
			_, err := c.ActivateConnectorForWorkspace("conn", nil)
			return err
		}},
		{"DeactivateConnectorForWorkspace", http.MethodPost, "/v1/connectors/conn/workspace/deactivate", func(c *MistralClient) error {
			_, err := c.DeactivateConnectorForWorkspace("conn")
			return err
		}},
		{"ActivateConnectorForUser", http.MethodPost, "/v1/connectors/conn/user/activate", func(c *MistralClient) error {
			_, err := c.ActivateConnectorForUser("conn", nil)
			return err
		}},
		{"DeactivateConnectorForUser", http.MethodPost, "/v1/connectors/conn/user/deactivate", func(c *MistralClient) error {
			_, err := c.DeactivateConnectorForUser("conn")
			return err
		}},
		{"ListIngestionPipelineConfigurations", http.MethodGet, "/v1/rag/ingestion_pipeline_configurations", func(c *MistralClient) error { _, err := c.ListIngestionPipelineConfigurations(); return err }},
		{"RegisterIngestionPipelineConfiguration", http.MethodPut, "/v1/rag/ingestion_pipeline_configurations", func(c *MistralClient) error {
			_, err := c.RegisterIngestionPipelineConfiguration(&RegisterIngestionPipelineConfigurationRequest{Name: "pipe", PipelineComposition: map[string]any{}})
			return err
		}},
		{"UpdateIngestionPipelineRunInfo", http.MethodPut, "/v1/rag/ingestion_pipeline_configurations/id/run_info", func(c *MistralClient) error {
			_, err := c.UpdateIngestionPipelineRunInfo("id", &UpdateIngestionPipelineRunInfoRequest{ChunksCount: &intVal})
			return err
		}},
		{"RegisterSearchIndex", http.MethodPut, "/v1/rag/search_index", func(c *MistralClient) error {
			status := SearchIndexStatusOffline
			_, err := c.RegisterSearchIndex(&RegisterSearchIndexRequest{Name: "idx", Index: map[string]any{"type": "vespa"}, Status: &status})
			return err
		}},
		{"GetWorkflows", http.MethodGet, "/v1/workflows", func(c *MistralClient) error {
			_, err := c.GetWorkflows(&ListWorkflowsParams{Limit: &intVal})
			return err
		}},
		{"GetWorkflowRegistrations", http.MethodGet, "/v1/workflows/registrations", func(c *MistralClient) error {
			_, err := c.GetWorkflowRegistrations(&ListWorkflowRegistrationsParams{ActiveOnly: &boolVal})
			return err
		}},
		{"ExecuteWorkflow", http.MethodPost, "/v1/workflows/wf/execute", func(c *MistralClient) error {
			_, err := c.ExecuteWorkflow("wf", &ExecuteWorkflowRequest{Input: map[string]any{}})
			return err
		}},
		{"ExecuteWorkflowRegistration", http.MethodPost, "/v1/workflows/registrations/reg/execute", func(c *MistralClient) error { _, err := c.ExecuteWorkflowRegistration("reg", nil); return err }},
		{"GetWorkflow", http.MethodGet, "/v1/workflows/wf", func(c *MistralClient) error { _, err := c.GetWorkflow("wf"); return err }},
		{"UpdateWorkflow", http.MethodPut, "/v1/workflows/wf", func(c *MistralClient) error {
			_, err := c.UpdateWorkflow("wf", &UpdateWorkflowRequest{DisplayName: &str})
			return err
		}},
		{"GetWorkflowRegistration", http.MethodGet, "/v1/workflows/registrations/reg", func(c *MistralClient) error { _, err := c.GetWorkflowRegistration("reg", nil); return err }},
		{"ArchiveWorkflow", http.MethodPut, "/v1/workflows/wf/archive", func(c *MistralClient) error { _, err := c.ArchiveWorkflow("wf"); return err }},
		{"UnarchiveWorkflow", http.MethodPut, "/v1/workflows/wf/unarchive", func(c *MistralClient) error { _, err := c.UnarchiveWorkflow("wf"); return err }},
		{"BulkArchiveWorkflows", http.MethodPut, "/v1/workflows/archive", func(c *MistralClient) error { _, err := c.BulkArchiveWorkflows([]string{"wf"}); return err }},
		{"BulkUnarchiveWorkflows", http.MethodPut, "/v1/workflows/unarchive", func(c *MistralClient) error { _, err := c.BulkUnarchiveWorkflows([]string{"wf"}); return err }},
		{"ListWorkflowDeployments", http.MethodGet, "/v1/workflows/deployments", func(c *MistralClient) error { _, err := c.ListWorkflowDeployments(); return err }},
		{"GetWorkflowDeployment", http.MethodGet, "/v1/workflows/deployments/dep", func(c *MistralClient) error { _, err := c.GetWorkflowDeployment("dep"); return err }},
		{"GetWorkflowMetrics", http.MethodGet, "/v1/workflows/wf/metrics", func(c *MistralClient) error { _, err := c.GetWorkflowMetrics("wf"); return err }},
		{"ListWorkflowRuns", http.MethodGet, "/v1/workflows/runs", func(c *MistralClient) error {
			_, err := c.ListWorkflowRuns(&ListWorkflowRunsParams{PageSize: &intVal})
			return err
		}},
		{"GetWorkflowRun", http.MethodGet, "/v1/workflows/runs/run", func(c *MistralClient) error { _, err := c.GetWorkflowRun("run"); return err }},
		{"GetWorkflowRunHistory", http.MethodGet, "/v1/workflows/runs/run/history", func(c *MistralClient) error { _, err := c.GetWorkflowRunHistory("run", &boolVal); return err }},
		{"GetWorkflowEvents", http.MethodGet, "/v1/workflows/events/list", func(c *MistralClient) error {
			_, err := c.GetWorkflowEvents(&ListWorkflowEventsParams{PageSize: &intVal})
			return err
		}},
		{"GetWorkflowSchedules", http.MethodGet, "/v1/workflows/schedules", func(c *MistralClient) error { _, err := c.GetWorkflowSchedules(); return err }},
		{"GetWorkflowSchedulesWithParams", http.MethodGet, "/v1/workflows/schedules", func(c *MistralClient) error {
			_, err := c.GetWorkflowSchedules(&ListWorkflowSchedulesParams{WorkflowName: &str, PageSize: &intVal})
			return err
		}},
		{"ScheduleWorkflow", http.MethodPost, "/v1/workflows/schedules", func(c *MistralClient) error {
			_, err := c.ScheduleWorkflow(&ScheduleWorkflowRequest{ScheduleID: &str})
			return err
		}},
		{"GetWorkflowSchedule", http.MethodGet, "/v1/workflows/schedules/sched", func(c *MistralClient) error { _, err := c.GetWorkflowSchedule("sched"); return err }},
		{"UpdateWorkflowSchedule", http.MethodPatch, "/v1/workflows/schedules/sched", func(c *MistralClient) error {
			_, err := c.UpdateWorkflowSchedule("sched", &UpdateScheduleRequest{Schedule: map[string]any{"input": map[string]any{}}})
			return err
		}},
		{"UnscheduleWorkflow", http.MethodDelete, "/v1/workflows/schedules/sched", func(c *MistralClient) error { _, err := c.UnscheduleWorkflow("sched"); return err }},
		{"PauseSchedule", http.MethodPost, "/v1/workflows/schedules/sched/pause", func(c *MistralClient) error {
			_, err := c.PauseSchedule("sched", &ScheduleNoteRequest{Note: &str})
			return err
		}},
		{"ResumeSchedule", http.MethodPost, "/v1/workflows/schedules/sched/resume", func(c *MistralClient) error { _, err := c.ResumeSchedule("sched", nil); return err }},
		{"GetWorkflowExecution", http.MethodGet, "/v1/workflows/executions/exec", func(c *MistralClient) error { _, err := c.GetWorkflowExecution("exec"); return err }},
		{"GetWorkflowExecutionHistory", http.MethodGet, "/v1/workflows/executions/exec/history", func(c *MistralClient) error { _, err := c.GetWorkflowExecutionHistory("exec", nil); return err }},
		{"SignalWorkflowExecution", http.MethodPost, "/v1/workflows/executions/exec/signals", func(c *MistralClient) error {
			_, err := c.SignalWorkflowExecution("exec", &WorkflowSignalRequest{Name: "sig"})
			return err
		}},
		{"QueryWorkflowExecution", http.MethodPost, "/v1/workflows/executions/exec/queries", func(c *MistralClient) error {
			_, err := c.QueryWorkflowExecution("exec", &WorkflowQueryRequest{Name: "query"})
			return err
		}},
		{"TerminateWorkflowExecution", http.MethodPost, "/v1/workflows/executions/exec/terminate", func(c *MistralClient) error { _, err := c.TerminateWorkflowExecution("exec"); return err }},
		{"BatchTerminateWorkflowExecutions", http.MethodPost, "/v1/workflows/executions/terminate", func(c *MistralClient) error {
			_, err := c.BatchTerminateWorkflowExecutions([]string{"exec"})
			return err
		}},
		{"CancelWorkflowExecution", http.MethodPost, "/v1/workflows/executions/exec/cancel", func(c *MistralClient) error { _, err := c.CancelWorkflowExecution("exec"); return err }},
		{"BatchCancelWorkflowExecutions", http.MethodPost, "/v1/workflows/executions/cancel", func(c *MistralClient) error { _, err := c.BatchCancelWorkflowExecutions([]string{"exec"}); return err }},
		{"ResetWorkflow", http.MethodPost, "/v1/workflows/executions/exec/reset", func(c *MistralClient) error {
			_, err := c.ResetWorkflow("exec", &ResetWorkflowRequest{EventID: "event"})
			return err
		}},
		{"UpdateWorkflowExecution", http.MethodPost, "/v1/workflows/executions/exec/updates", func(c *MistralClient) error {
			_, err := c.UpdateWorkflowExecution("exec", &WorkflowUpdateRequest{Name: "upd"})
			return err
		}},
		{"GetWorkflowExecutionTraceOTEL", http.MethodGet, "/v1/workflows/executions/exec/trace/otel", func(c *MistralClient) error { _, err := c.GetWorkflowExecutionTraceOTEL("exec"); return err }},
		{"GetWorkflowExecutionTraceSummary", http.MethodGet, "/v1/workflows/executions/exec/trace/summary", func(c *MistralClient) error { _, err := c.GetWorkflowExecutionTraceSummary("exec"); return err }},
		{"GetWorkflowExecutionTraceEvents", http.MethodGet, "/v1/workflows/executions/exec/trace/events", func(c *MistralClient) error {
			_, err := c.GetWorkflowExecutionTraceEvents("exec", &WorkflowTraceEventsParams{MergeSameIDEvents: &boolVal})
			return err
		}},
		{"CreateCampaign", http.MethodPost, "/v1/observability/campaigns", func(c *MistralClient) error {
			_, err := c.CreateCampaign(&CreateCampaignRequest{Name: "camp"})
			return err
		}},
		{"ListCampaigns", http.MethodGet, "/v1/observability/campaigns", func(c *MistralClient) error {
			_, err := c.ListCampaigns(&ListObservabilityParams{PageSize: &intVal})
			return err
		}},
		{"FetchCampaign", http.MethodGet, "/v1/observability/campaigns/camp", func(c *MistralClient) error { _, err := c.FetchCampaign("camp"); return err }},
		{"FetchCampaignStatus", http.MethodGet, "/v1/observability/campaigns/camp/status", func(c *MistralClient) error { _, err := c.FetchCampaignStatus("camp"); return err }},
		{"ListCampaignEvents", http.MethodGet, "/v1/observability/campaigns/camp/selected-events", func(c *MistralClient) error { _, err := c.ListCampaignEvents("camp", &intVal, nil); return err }},
		{"CreateDataset", http.MethodPost, "/v1/observability/datasets", func(c *MistralClient) error { _, err := c.CreateDataset(&CreateDatasetRequest{Name: "ds"}); return err }},
		{"ListDatasets", http.MethodGet, "/v1/observability/datasets", func(c *MistralClient) error { _, err := c.ListDatasets(nil); return err }},
		{"FetchDataset", http.MethodGet, "/v1/observability/datasets/ds", func(c *MistralClient) error { _, err := c.FetchDataset("ds"); return err }},
		{"UpdateDataset", http.MethodPatch, "/v1/observability/datasets/ds", func(c *MistralClient) error {
			_, err := c.UpdateDataset("ds", &UpdateDatasetRequest{Name: &str})
			return err
		}},
		{"ListDatasetRecords", http.MethodGet, "/v1/observability/datasets/ds/records", func(c *MistralClient) error { _, err := c.ListDatasetRecords("ds", nil, nil); return err }},
		{"CreateDatasetRecord", http.MethodPost, "/v1/observability/datasets/ds/records", func(c *MistralClient) error {
			_, err := c.CreateDatasetRecord("ds", &DatasetRecordRequest{Payload: map[string]any{}})
			return err
		}},
		{"ImportDatasetFromCampaign", http.MethodPost, "/v1/observability/datasets/ds/imports/from-campaign", func(c *MistralClient) error { _, err := c.ImportDatasetFromCampaign("ds", "camp"); return err }},
		{"ExportDatasetToJSONL", http.MethodGet, "/v1/observability/datasets/ds/exports/to-jsonl", func(c *MistralClient) error { _, err := c.ExportDatasetToJSONL("ds"); return err }},
		{"FetchDatasetTask", http.MethodGet, "/v1/observability/datasets/ds/tasks/task", func(c *MistralClient) error { _, err := c.FetchDatasetTask("ds", "task"); return err }},
		{"ListChatCompletionFields", http.MethodGet, "/v1/observability/chat-completion-fields", func(c *MistralClient) error { _, err := c.ListChatCompletionFields(); return err }},
		{"FetchChatCompletionFieldOptions", http.MethodGet, "/v1/observability/chat-completion-fields/field/options", func(c *MistralClient) error { _, err := c.FetchChatCompletionFieldOptions("field", nil); return err }},
		{"SearchChatCompletionEvents", http.MethodPost, "/v1/observability/chat-completion-events/search", func(c *MistralClient) error {
			_, err := c.SearchChatCompletionEvents(&SearchChatCompletionEventsRequest{SearchParams: map[string]any{}})
			return err
		}},
		{"FetchChatCompletionEvent", http.MethodGet, "/v1/observability/chat-completion-events/event", func(c *MistralClient) error { _, err := c.FetchChatCompletionEvent("event"); return err }},
		{"JudgeChatCompletionEvent", http.MethodPost, "/v1/observability/chat-completion-events/event/live-judging", func(c *MistralClient) error {
			_, err := c.JudgeChatCompletionEvent("event", map[string]any{})
			return err
		}},
		{"CreateJudge", http.MethodPost, "/v1/observability/judges", func(c *MistralClient) error { _, err := c.CreateJudge(&JudgeRequest{Name: "judge"}); return err }},
		{"ListJudges", http.MethodGet, "/v1/observability/judges", func(c *MistralClient) error { _, err := c.ListJudges(&ListJudgesParams{Q: &str}); return err }},
		{"FetchJudge", http.MethodGet, "/v1/observability/judges/judge", func(c *MistralClient) error { _, err := c.FetchJudge("judge"); return err }},
		{"JudgeConversation", http.MethodPost, "/v1/observability/judges/judge/live-judging", func(c *MistralClient) error {
			_, err := c.JudgeConversation("judge", &JudgeConversationRequest{Messages: []any{}})
			return err
		}},
		{"FetchDatasetRecord", http.MethodGet, "/v1/observability/dataset-records/rec", func(c *MistralClient) error { _, err := c.FetchDatasetRecord("rec"); return err }},
		{"BulkDeleteDatasetRecords", http.MethodPost, "/v1/observability/dataset-records/bulk-delete", func(c *MistralClient) error { _, err := c.BulkDeleteDatasetRecords([]string{"rec"}); return err }},
		{"UpdateDatasetRecordPayload", http.MethodPut, "/v1/observability/dataset-records/rec/payload", func(c *MistralClient) error {
			_, err := c.UpdateDatasetRecordPayload("rec", map[string]any{})
			return err
		}},
		{"DeleteBatchJob", http.MethodDelete, "/v1/batch/jobs/job", func(c *MistralClient) error {
			_, err := c.DeleteBatchJob("job")
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

func TestNewParityBinaryAndStreamEndpoints(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/audio/voices/voice/sample":
			w.Header().Set("Content-Type", "audio/wav")
			_, _ = w.Write([]byte("wav"))
		case "/v1/audio/speech", "/v1/workflows/events/stream", "/v1/workflows/executions/exec/stream":
			w.Header().Set("Content-Type", "text/event-stream")
			_, _ = io.WriteString(w, "data: {\"type\":\"ok\"}\n\n")
			_, _ = io.WriteString(w, "data: [DONE]\n\n")
		default:
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
	})
	defer mock.Close()
	client := mock.GetClient()

	if data, err := client.GetVoiceSampleAudio("voice"); err != nil || string(data) != "wav" {
		t.Fatalf("unexpected voice sample result: %q %v", string(data), err)
	}
	if events, err := client.SpeechStream(&SpeechRequest{Input: "hello"}); err != nil {
		t.Fatalf("unexpected speech stream error: %v", err)
	} else {
		for range events {
			break
		}
	}
	if events, err := client.GetWorkflowStreamEvents(nil); err != nil {
		t.Fatalf("unexpected workflow events stream error: %v", err)
	} else {
		for range events {
			break
		}
	}
	if events, err := client.StreamWorkflowExecution("exec", nil); err != nil {
		t.Fatalf("unexpected workflow execution stream error: %v", err)
	} else {
		for range events {
			break
		}
	}
}

func TestListSearchIndexesEndpoint(t *testing.T) {
	mock := NewMockHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected method %s, got %s", http.MethodGet, r.Method)
		}
		if r.URL.Path != "/v1/rag/search_index" {
			t.Errorf("expected path /v1/rag/search_index, got %s", r.URL.Path)
		}
		MockJSONResponse(http.StatusOK, `[{"id":"idx","name":"Index","creator_id":"user","document_count":1,"status":"online","created_at":"2026-01-01T00:00:00Z","modified_at":"2026-01-01T00:00:00Z","index":{"type":"vespa"}}]`).Write(w)
	})
	defer mock.Close()

	indexes, err := mock.GetClient().ListSearchIndexes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(indexes) != 1 || indexes[0].ID != "idx" {
		t.Fatalf("unexpected indexes: %+v", indexes)
	}
}
