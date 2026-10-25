package sdk

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPython340Endpoints(t *testing.T) {
	cases := []struct {
		name, method, path, response string
		call                         func(*MistralClient) error
	}{
		{"agents list", "GET", "/v1/audio/agents", `{"data":[]}`, func(c *MistralClient) error { _, e := c.ListAudioAgents(nil); return e }},
		{"agents create", "POST", "/v1/audio/agents", `{"id":"agent"}`, func(c *MistralClient) error {
			_, e := c.CreateAudioAgent(&AudioAgentCreateRequest{Name: "test", Configuration: AudioAgentConfigurationRequest{}})
			return e
		}},
		{"agents get", "GET", "/v1/audio/agents/a%2Fb", `{"id":"a/b"}`, func(c *MistralClient) error { _, e := c.GetAudioAgent("a/b", IntPtr(1)); return e }},
		{"agents update", "PUT", "/v1/audio/agents/a", `{"id":"a"}`, func(c *MistralClient) error {
			_, e := c.UpdateAudioAgent("a", &AudioAgentUpdateRequest{Name: "test"})
			return e
		}},
		{"agents delete", "DELETE", "/v1/audio/agents/a", "", func(c *MistralClient) error { return c.DeleteAudioAgent("a") }},
		{"agent versions", "GET", "/v1/audio/agents/a/versions", `{"data":[]}`, func(c *MistralClient) error { _, e := c.ListAudioAgentVersions("a", nil); return e }},
		{"agent sessions", "GET", "/v1/audio/agents/a/sessions", `{"data":[]}`, func(c *MistralClient) error { _, e := c.ListAudioAgentSessions("a", nil); return e }},
		{"session create", "POST", "/v1/audio/sessions", `{"id":"s","participant_token":"test","server_url":"wss://example.test"}`, func(c *MistralClient) error {
			_, e := c.CreateAudioSession(&AudioSessionCreateRequest{Tools: []any{}})
			return e
		}},
		{"session get", "GET", "/v1/audio/sessions/s", `{"id":"s"}`, func(c *MistralClient) error { _, e := c.GetAudioSession("s"); return e }},
		{"session close", "DELETE", "/v1/audio/sessions/s", "", func(c *MistralClient) error { return c.CloseAudioSession("s") }},
		{"responses create", "POST", "/v1/responses", `{"id":"r","model":"m","output":[],"parallel_tool_calls":true,"tool_choice":"auto","tools":[]}`, func(c *MistralClient) error {
			_, e := c.CreateResponse(&CreateResponseRequest{Model: "m", Input: "hello", Tools: []any{}, AdditionalProperties: map[string]any{"extension": true}})
			return e
		}},
		{"responses count", "POST", "/v1/responses/input_tokens", `{"object":"response.input_tokens","input_tokens":3}`, func(c *MistralClient) error {
			r, e := c.CountResponseInputTokens(&CountResponseInputTokensRequest{Model: "m", Input: "hello"})
			if e == nil && r.InputTokens != 3 {
				t.Error(r)
			}
			return e
		}},
		{"summary get", "GET", "/v1/context/search-indexes/a%2Fb/summary_field/en", `{"content":"summary"}`, func(c *MistralClient) error { _, e := c.GetManagedIndexSummary("a/b", "en"); return e }},
		{"summary set", "PUT", "/v1/context/search-indexes/a/summary_field/en", `{}`, func(c *MistralClient) error {
			_, e := c.SetManagedIndexSummary("a", "en", &UpdateSummaryRequestSummary{Content: "summary", Status: "completed", Translated: false})
			return e
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.EscapedPath() != tc.path {
					t.Errorf("%s %s", r.Method, r.URL.EscapedPath())
				}
				if strings.HasSuffix(r.URL.Path, "agents") && r.Method == "GET" && r.URL.Query().Get("page_size") != "10" {
					t.Error(r.URL)
				}
				if tc.name == "responses create" {
					var body map[string]any
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Error(err)
					}
					if body["stream"] != false || body["parallel_tool_calls"] != true || body["extension"] != true {
						t.Error(body)
					}
					tools, ok := body["tools"].([]any)
					if !ok || len(tools) != 0 {
						t.Error(body)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				io.WriteString(w, tc.response)
			}))
			defer server.Close()
			if err := tc.call(NewMistralClient("test", server.URL, 1, 0)); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPython340Streaming(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/responses") {
			if r.Header.Get("Accept") != "text/event-stream" {
				t.Error(r.Header)
			}
			w.Header().Set("Content-Type", "text/event-stream")
			io.WriteString(w, "event: response.output_text.delta\ndata: {\"type\":\"response.output_text.delta\",\"delta\":\"hi\"}\n\ndata: [DONE]\n\n")
		} else {
			if r.Header.Get("Accept") != "application/x-ndjson" || r.URL.Query().Get("force") != "false" {
				t.Error(r.Header, r.URL)
			}
			w.Header().Set("Content-Type", "application/x-ndjson")
			io.WriteString(w, "{\"type\":\"chunk\",\"content\":\"hello\"}\n{\"type\":\"status\",\"status\":\"completed\"}\n")
		}
	}))
	defer server.Close()
	client := NewMistralClient("test", server.URL, 1, 0)
	request := &CreateResponseRequest{Model: "m", Input: "hello"}
	stream, err := client.CreateResponseStream(request)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()
	event, err := stream.Next()
	if err != nil || event.Event != "response.output_text.delta" {
		t.Fatal(event, err)
	}
	if request.Stream != nil {
		t.Fatal("request mutated")
	}
	summary, err := client.GenerateManagedIndexSummary("i", "en", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer summary.Close()
	for range 2 {
		if _, err := summary.Next(); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := summary.Next(); err != io.EOF {
		t.Fatal(err)
	}
}

func TestPython340ModelFields(t *testing.T) {
	var usage UsageInfo
	if err := json.Unmarshal([]byte(`{"completion_tokens_details":{"reasoning_tokens":0}}`), &usage); err != nil {
		t.Fatal(err)
	}
	if usage.CompletionTokensDetails == nil || usage.CompletionTokensDetails.ReasoningTokens == nil || *usage.CompletionTokensDetails.ReasoningTokens != 0 {
		t.Fatal(usage)
	}
	var session RealtimeTranscriptionSession
	if err := json.Unmarshal([]byte(`{"diarize":false,"context_bias":[],"language":"en","speaker_references":["a"],"turn_detection":{"type":"semantic_vad"}}`), &session); err != nil {
		t.Fatal(err)
	}
	if session.Diarize == nil || *session.Diarize || len(session.SpeakerReferences) != 1 || session.TurnDetection == nil {
		t.Fatal(session)
	}
}

func TestPython340ExtensionFieldsRoundTrip(t *testing.T) {
	var value ConnectorToolCallResponse
	if err := json.Unmarshal([]byte(`{"type":"connector_call","extension":{"answer":42}}`), &value); err != nil {
		t.Fatal(err)
	}
	if value.AdditionalProperties["extension"].(map[string]any)["answer"] != float64(42) {
		t.Fatal(value)
	}
	data, err := json.Marshal(value)
	if err != nil || !strings.Contains(string(data), `"extension":{"answer":42}`) {
		t.Fatal(string(data), err)
	}
	request := CreateResponseRequest{Model: "real", Input: "hello", AdditionalProperties: map[string]any{"model": "override", "extension": true}}
	data, err = json.Marshal(request)
	if err != nil || strings.Contains(string(data), "override") {
		t.Fatal(string(data), err)
	}
}
