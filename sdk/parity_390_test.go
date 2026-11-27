package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// Exercise public client calls with real URL construction and response decoding.
func Test390Routes(t *testing.T) {
	id := "one/two"
	cases := []struct {
		method, path, reply string
		call                func(*MistralClient) error
	}{
		{"GET", "/v1/billing/budget", "{}", func(c *MistralClient) error { _, err := c.GetBudget("dashboard-key"); return err }},
		{"GET", "/v1/billing/subscription", `{"plan_name":"team"}`, func(c *MistralClient) error {
			r, err := c.GetSubscription("dashboard-key")
			if err == nil && r.PlanName != "team" {
				t.Errorf("subscription: %+v", r)
			}
			return err
		}},
		{"GET", "/v1/webhooks", `[{"id":"w"}]`, func(c *MistralClient) error {
			r, e := c.ListWebhooks()
			if e == nil && (len(r) != 1 || r[0].ID != "w") {
				t.Errorf("webhooks: %+v", r)
			}
			return e
		}},
		{"POST", "/v1/webhooks", "{}", func(c *MistralClient) error { _, e := c.CreateWebhook(CreateWebhook{}); return e }},
		{"GET", "/v1/webhooks/one%2Ftwo", "{}", func(c *MistralClient) error { _, e := c.GetWebhook(id); return e }},
		{"DELETE", "/v1/webhooks/one%2Ftwo", "", func(c *MistralClient) error { return c.DeleteWebhook(id) }},
		{"GET", "/v1/webhooks/one%2Ftwo/signing-secret", `{"signing_secret":"test"}`, func(c *MistralClient) error {
			r, e := c.GetWebhookSigningSecret(id)
			if e == nil && r.SigningSecret != "test" {
				t.Error("missing secret")
			}
			return e
		}},
		{"POST", "/v1/triggers", "{}", func(c *MistralClient) error { _, e := c.CreateTrigger(CreateTrigger{}); return e }},
		{"GET", "/v1/triggers?primitive_id=x", "[]", func(c *MistralClient) error {
			x := "x"
			_, e := c.ListTriggers(&ListTriggersParams{PrimitiveID: &x})
			return e
		}},
		{"GET", "/v1/triggers/one%2Ftwo", "{}", func(c *MistralClient) error { _, e := c.GetTrigger(id); return e }},
		{"DELETE", "/v1/triggers/one%2Ftwo", "", func(c *MistralClient) error { return c.DeleteTrigger(id) }},
		{"GET", "/v1/triggers/one%2Ftwo/events?limit=10", "{}", func(c *MistralClient) error { _, e := c.ListTriggerEvents(id, nil); return e }},
		{"POST", "/v1/observability/scorers", "{}", func(c *MistralClient) error { _, e := c.CreateScorer(CreateScorerRequest{}); return e }},
		{"GET", "/v1/observability/scorers?created_by_me=false&ootb=false&page=1&page_size=50", "{}", func(c *MistralClient) error { _, e := c.ListScorers(nil); return e }},
		{"GET", "/v1/observability/scorers/one%2Ftwo", "{}", func(c *MistralClient) error { _, e := c.GetScorer(id); return e }},
		{"PATCH", "/v1/observability/scorers/one%2Ftwo", "{}", func(c *MistralClient) error { _, e := c.UpdateScorer(id, EditScorerRequest{}); return e }},
		{"POST", "/v1/observability/datasets/one%2Ftwo/records/bulk", "{}", func(c *MistralClient) error {
			_, e := c.CreateDatasetRecords(id, []CreateDatasetRecordsItem{})
			return e
		}},
		{"GET", "/v1/observability/pipelines/one%2Ftwo/results/fields", "{}", func(c *MistralClient) error { _, e := c.ListPipelineResultFields(id); return e }},
		{"POST", "/v1/workflows/deployments/archives", "{}", func(c *MistralClient) error { _, e := c.CreateDeploymentArchive(123); return e }},
		{"GET", "/v1/workflows/deployments/one%2Ftwo/archive", "{}", func(c *MistralClient) error { _, e := c.GetDeploymentArchive(id); return e }},
		{"POST", "/v1/workflows/one%2Ftwo/tags", "", func(c *MistralClient) error { return c.AddWorkflowTags(id, []string{}) }},
		{"PATCH", "/v1/audio/sessions/one%2Ftwo", "{}", func(c *MistralClient) error {
			_, e := c.UpdateAudioSession(id, AudioSessionUpdateRequest{Tools: []any{}})
			return e
		}},
		{"POST", "/v1/moderations", `{"results":[{"categories":{"safe":true},"category_scores":{"safe":0.9}}]}`, func(c *MistralClient) error {
			r, e := c.ModerateWithParams(ModerationRequest{Model: "model", Inputs: "hello", Queries: map[string]string{"safe": "check"}})
			if e == nil && (!r.Results[0].Thresholds["safe"] || r.Results[0].Categories[0].Score != 0.9) {
				t.Error("moderation response lost data")
			}
			return e
		}},
	}
	for _, tc := range cases {
		t.Run(tc.method+" "+tc.path, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.RequestURI() != tc.path {
					t.Errorf("got %s %s", r.Method, r.URL.RequestURI())
				}
				if strings.Contains(tc.path, "/billing/") {
					if r.Header.Get("X-Api-Key") != "dashboard-key" || r.Header.Get("Authorization") != "" {
						t.Error("billing used wrong credentials")
					}
				}
				if tc.path == "/v1/moderations" || tc.path == "/v1/audio/sessions/one%2Ftwo" {
					b, _ := io.ReadAll(r.Body)
					if !strings.Contains(string(b), `"input":"hello"`) && !strings.Contains(string(b), `"tools":[]`) {
						t.Errorf("body %s", b)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, tc.reply)
			}))
			defer server.Close()
			if err := tc.call(NewMistralClient("sdk-key", server.URL, 1, time.Second)); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func Test390RealtimeSpeech(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/audio/speech/realtime" || r.URL.Query().Get("model") != "vox" || r.Header.Get("Authorization") != "Bearer key" {
			t.Error("bad speech handshake")
		}
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.CloseNow()
		for _, msg := range []string{`{"type":"server.notice"}`, `{"type":"session.created","session":{"request_id":"r","model":"vox"}}`} {
			if err := conn.Write(ctx, websocket.MessageText, []byte(msg)); err != nil {
				return
			}
		}
		for _, want := range []string{"session.update", "input_text.append", "input_text.end"} {
			var m map[string]any
			if err = wsjson.Read(ctx, conn, &m); err != nil {
				t.Error(err)
				return
			}
			if m["type"] != want {
				t.Errorf("wanted %s got %v", want, m)
			}
		}
		for _, msg := range []string{`{"type":"speech.audio.delta","audio_data":"YWJj"}`, `{"type":"speech.audio.done","usage":{"total_tokens":6}}`} {
			if err := conn.Write(ctx, websocket.MessageText, []byte(msg)); err != nil {
				return
			}
		}
		conn.Close(websocket.StatusNormalClosure, "")
	}))
	defer server.Close()
	text := make(chan string, 1)
	text <- "hello"
	close(text)
	voice := "en"
	stream, err := NewMistralClient("key", server.URL, 1, time.Second).RealtimeSpeakStream(ctx, "vox", text, &RealtimeSpeechConnectParams{Session: &RealtimeSpeechSessionUpdatePayload{Voice: &voice}})
	if err != nil {
		t.Fatal(err)
	}
	var types []string
	for ev := range stream {
		if ev.Err != nil {
			t.Fatal(ev.Err)
		}
		types = append(types, ev.Type)
		if ev.AudioData != "" {
			b, e := ev.AudioBytes()
			if e != nil || string(b) != "abc" {
				t.Fatal("bad audio")
			}
		}
	}
	if !reflect.DeepEqual(types, []string{"server.notice", "session.created", "speech.audio.delta", "speech.audio.done"}) {
		t.Fatalf("got %v", types)
	}
}
func Test390WorkflowScheduleWire(t *testing.T) {
	input := map[string]any{"x": false, "a": []string{"hello", "España"}}
	var stored map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			var body struct {
				Schedule struct {
					Input map[string]any `json:"input"`
				} `json:"schedule"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			stored = body.Schedule.Input
			if stored["input_hash"] != "d2f7c4413cc632031c014772996e0571d6008ad3b4aaa1fa403c7f6a30f8a053" {
				t.Errorf("wrong hash %v", stored)
			}
			if !reflect.DeepEqual(stored["encoding_options"], []any{"encrypted"}) {
				t.Error("schedule must stay inline")
			}
			fmt.Fprint(w, "{}")
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"schedules": []any{map[string]any{"input": stored}}})
	}))
	defer server.Close()
	ctx := context.Background()
	client, err := NewMistralClient("x", server.URL, 1, time.Second).WithWorkflowEncoding(ctx, WorkflowEncodingConfig{Namespace: "ns", EncryptionMode: "full", MainKeyHex: strings.Repeat("58", 32)})
	if err != nil {
		t.Fatal(err)
	}
	if _, err = client.ScheduleWorkflow(&ScheduleWorkflowRequest{Schedule: map[string]any{"input": input}}); err != nil {
		t.Fatal(err)
	}
	response, err := client.GetWorkflowSchedules()
	if err != nil {
		t.Fatal(err)
	}
	actual := response["schedules"].([]any)[0].(map[string]any)["input"]
	a, _ := json.Marshal(actual)
	b, _ := json.Marshal(input)
	if string(a) != string(b) {
		t.Fatalf("decoded %s", a)
	}
}

type memoryWorkflowStore map[string][]byte

func (s memoryWorkflowStore) PutIfAbsent(_ context.Context, key string, data []byte) error {
	if _, ok := s[key]; !ok {
		s[key] = append([]byte(nil), data...)
	}
	return nil
}
func (s memoryWorkflowStore) Get(_ context.Context, key string) ([]byte, error) {
	data, ok := s[key]
	if !ok {
		return nil, fmt.Errorf("missing %s", key)
	}
	return data, nil
}

func Test390WorkflowTransforms(t *testing.T) {
	ctx := context.Background()
	store := memoryWorkflowStore{}
	client, err := NewMistralClientDefault("key").WithWorkflowEncoding(ctx, WorkflowEncodingConfig{Namespace: "my ns", EncryptionMode: "partial", MainKeyHex: strings.Repeat("66", 32), BlobStore: store, Compression: &WorkflowCompression{Level: 3}})
	if err != nil {
		t.Fatal(err)
	}
	input := map[string]any{"message": strings.Repeat("repeated", 500), "private": map[string]any{"field_type": "__encrypted_str__", "data": "secret"}}
	envelope, err := client.workflowEncoding.encode(ctx, input, "id/with/slash", false, false)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(envelope["encoding_options"], []string{"encrypted-partial", "compressed", "offloaded"}) || len(store) != 1 {
		t.Fatalf("workflow transforms %+v", envelope)
	}
	for key := range store {
		if !strings.HasPrefix(key, "temporal-payload/my%20ns/id%2Fwith%2Fslash/sha256:") {
			t.Fatal(key)
		}
	}
	other, err := NewMistralClientDefault("key").WithWorkflowEncoding(ctx, WorkflowEncodingConfig{Namespace: "my ns", EncryptionMode: "full", MainKeyHex: strings.Repeat("77", 32), SecondaryKeyHex: strings.Repeat("66", 32), BlobStore: store})
	if err != nil {
		t.Fatal(err)
	}
	actual, err := other.workflowEncoding.decode(ctx, envelope)
	if err != nil || !reflect.DeepEqual(input, actual) {
		t.Fatalf("workflow roundtrip: %v", err)
	}
	envelope["encoding_options"] = []string{"bogus"}
	if _, err = other.workflowEncoding.decode(ctx, envelope); err == nil {
		t.Fatal("unknown encoding accepted")
	}
}
func Test390GCSCustomAPIRoot(t *testing.T) {
	var data []byte
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if r.URL.Query().Get("ifGenerationMatch") != "0" || r.URL.Query().Get("name") != "prefix/k/1" {
				t.Errorf("query %v", r.URL.Query())
			}
			if data != nil {
				w.WriteHeader(412)
				return
			}
			data, _ = io.ReadAll(r.Body)
			return
		}
		if r.URL.EscapedPath() != "/custom/storage/v1/b/bucket/o/prefix%2Fk%2F1" || r.URL.Query().Get("alt") != "media" {
			t.Error(r.URL)
		}
		w.Write(data)
	}))
	defer server.Close()
	ctx := context.Background()
	store, err := NewGCSWorkflowBlobStore(ctx, "bucket", "prefix", server.URL+"/custom/", server.Client())
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"first", "second"} {
		if err := store.PutIfAbsent(ctx, "k/1", []byte(s)); err != nil {
			t.Fatal(err)
		}
	}
	b, err := store.Get(ctx, "k/1")
	if err != nil || string(b) != "first" {
		t.Fatalf("stored %q %v", b, err)
	}
}
func Test390ExplicitNullsAndConversationEntries(t *testing.T) {
	req := AudioSessionUpdateRequest{Tools: []any{}, ExplicitNulls: []string{"instructions"}}
	b, err := json.Marshal(req)
	if err != nil || string(b) != `{"instructions":null,"tools":[]}` {
		t.Fatalf("patch %s: %v", b, err)
	}
	req.ExplicitNulls = []string{"misspelled"}
	if _, err = json.Marshal(req); err == nil {
		t.Fatal("unknown patch field accepted")
	}
	entries, err := ConversationInputs("hello", FunctionResultEntry{ToolCallID: "call", Result: "ok"}, map[string]any{"type": "function.call", "tool_call_id": "call", "name": "price", "arguments": map[string]any{"currency": "EUR"}})
	if err != nil {
		t.Fatal(err)
	}
	b, err = json.Marshal(entries)
	if err != nil || !strings.Contains(string(b), `"tool_call_id":"call"`) || !strings.Contains(string(b), `"role":"user"`) {
		t.Fatalf("entries %s %v", b, err)
	}
}

func Test390PythonWorkflowFixtures(t *testing.T) {
	// Python cryptography AESGCM (nonce prefix), msgpack.packb and Zstandard frames.
	key := "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
	c, err := NewMistralClientDefault("key").WithWorkflowEncoding(context.Background(), WorkflowEncodingConfig{Namespace: "ns", EncryptionMode: "full", MainKeyHex: key})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ mode, b64, k, want string }{
		{"encrypted", "AAECAwQFBgcICQoLPCC7fraWo3zoY62px4wKBOW/4lDQGSZcaB6R7XIHIs8/Idi/VgxFT1gkAJg37C8b", "message", "verified by Python"},
		{"compressed", "gqtjb21wcmVzc2lvboKpYWxnb3JpdGhtpHpzdGSlbGV2ZWwDp3BheWxvYWTEKSi1L/0EWOUAAKh7InRleHQiOiJyZXBlYXQgbWUgIn0BALe3KmEYSxH2", "text", strings.Repeat("repeat me ", 100)},
	}
	for _, test := range cases {
		value, err := c.workflowEncoding.decode(context.Background(), map[string]any{"b64payload": test.b64, "encoding_options": []string{test.mode}})
		if err != nil {
			t.Fatal(err)
		}
		if value.(map[string]any)[test.k] != test.want {
			t.Error("Python fixture mismatch", test.mode)
		}
	}
	v, err := workflowValue([]byte(`{"z":"\u2028","x":1e-7,"n":1.0,"amp":"<&>"}`))
	if err != nil {
		t.Fatal(err)
	}
	data, err := scheduleCanonicalJSON(v)
	if err != nil || string(data) != "{\"amp\":\"<&>\",\"n\":1.0,\"x\":1e-07,\"z\":\"\u2028\"}" {
		t.Errorf("canonical JSON %q %v", data, err)
	}
}
