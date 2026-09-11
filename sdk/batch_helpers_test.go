package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

type batchTestBody struct {
	Value int `json:"value"`
}

func TestBatchJSONLAndOrdering(t *testing.T) {
	input, err := NewBatchInput(map[string]batchTestBody{"b": {2}, "a": {1}})
	if err != nil {
		t.Fatal(err)
	}
	requests, err := input.ToRequests()
	if err != nil || len(requests) != 2 || (requests[0].CustomID == nil || *requests[0].CustomID != "a") || input.Len() != 2 {
		t.Fatalf("input %v %v", requests, err)
	}
	raw := `{"custom_id":"b","response":{"status_code":200,"body":{"value":2}}}
{"custom_id":"a","response":{"status_code":200,"body":{"value":1}}}
{"custom_id":"bad","response":{"status_code":200,"body":{"value":"wrong type"}}}
{"custom_id":"failed","response":{"status_code":429},"error":{"message":"rate limited"}}
{"response":{"status_code":400},"error":"no id"}
{"response":{"status_code":200,"body":{"value":3}}}`
	output, err := ParseBatchOutput[batchTestBody]([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	if len(output.Responses) != 2 || len(output.RequestErrors) != 2 || len(output.ResponseErrors) != 2 {
		t.Fatalf("output %+v", output)
	}
	ordered, err := output.Ordered([]string{"a", "b", "a"})
	if err != nil || !reflect.DeepEqual(ordered, []batchTestBody{{1}, {2}, {1}}) {
		t.Fatalf("ordered %+v %v", ordered, err)
	}
	_, err = output.Ordered([]string{"missing", "bad", "failed"})
	var batchErr *BatchError
	if !errors.As(err, &batchErr) || len(batchErr.Failures) != 3 {
		t.Fatalf("failures: %v", err)
	}
	duplicate := `{"custom_id":"a","response":{"body":{"value":1}}}`
	if _, err := ParseBatchOutput[batchTestBody]([]byte(duplicate + "\n" + duplicate)); err == nil {
		t.Fatal("duplicate successes accepted")
	}
	result := &BatchResult[batchTestBody]{Output: output, ErrorOutput: &BatchOutput[batchTestBody]{Responses: map[string]batchTestBody{"a": {7}}}}
	if _, err := result.ByID(); err == nil {
		t.Fatal("duplicate across output files accepted")
	}
	for _, status := range []BatchJobStatus{BatchJobStatusQueued, BatchJobStatusRunning, BatchJobStatusCancellationRequested} {
		if !IsBatchRunning(status) {
			t.Errorf("status %s", status)
		}
	}
	if !IsBatchTerminal("FUTURE_STATUS") {
		t.Fatal("unknown statuses must stop polling")
	}
}

func TestBatchRunLifecycle(t *testing.T) {
	for _, preuploaded := range []bool{false, true} {
		t.Run(fmt.Sprint(preuploaded), func(t *testing.T) {
			var created, polled, uploaded int
			deleted := map[string]bool{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("X-Batch-Test") != "yes" {
					t.Error("custom headers missing")
				}
				w.Header().Set("Content-Type", "application/json")
				switch {
				case r.Method == "POST" && r.URL.Path == "/v1/files":
					uploaded++
					if err := r.ParseMultipartForm(1 << 20); err != nil {
						t.Error(err)
					}
					if r.FormValue("purpose") != "batch" {
						t.Error("purpose missing")
					}
					fmt.Fprint(w, `{"id":"owned-input"}`)
				case r.Method == "POST" && r.URL.Path == "/v1/batch/jobs":
					created++
					var body map[string]any
					_ = json.NewDecoder(r.Body).Decode(&body)
					if body["endpoint"] != "/v1/embeddings" || body["input_files"] == nil {
						t.Errorf("body %v", body)
					}
					fmt.Fprintf(w, `{"id":"job-%d","status":"QUEUED"}`, created)
				case r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/v1/batch/jobs/"):
					polled++
					if polled == 1 {
						fmt.Fprint(w, `{"id":"job-1","status":"CANCELLATION_REQUESTED"}`)
						return
					}
					failed := 0
					if created == 1 {
						failed = 1
					}
					fmt.Fprintf(w, `{"id":"job-%d","status":"SUCCESS","failed_requests":%d,"succeeded_requests":1,"output_file":"output-%d","error_file":"error-%d"}`, created, failed, created, created)
				case r.Method == "GET" && strings.HasSuffix(r.URL.Path, "/content"):
					if strings.Contains(r.URL.Path, "output-") {
						fmt.Fprint(w, `{"custom_id":"a","response":{"status_code":200,"body":{"value":9}}}`)
					}
				case r.Method == "DELETE":
					deleted[r.URL.Path] = true
					w.WriteHeader(204)
				default:
					t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
					w.WriteHeader(400)
				}
			}))
			defer server.Close()
			client := NewBatchClient[batchTestBody](NewMistralClient("test", server.URL, 1, time.Second), BatchEndpointEmbeddings)
			client.HTTPHeaders = http.Header{"X-Batch-Test": []string{"yes"}}
			input, _ := NewBatchInput(map[string]batchTestBody{"a": {1}})
			var source any = input
			if preuploaded {
				source = BatchInputFile{FileID: "caller-input"}
			}
			threshold := 0
			result, err := client.Run(context.Background(), source, &BatchRunOptions{InlineThreshold: &threshold, PollInterval: time.Millisecond, MaxFailedJobRetries: 1})
			if err != nil {
				t.Fatal(err)
			}
			rows, err := result.Ordered([]string{"a"})
			if err != nil || rows[0].Value != 9 {
				t.Fatalf("results %v %v", rows, err)
			}
			if created != 2 || polled != 3 {
				t.Fatalf("created %d polled %d", created, polled)
			}
			if deleted["/v1/files/caller-input"] {
				t.Error("caller input deleted")
			}
			if !preuploaded && (!deleted["/v1/files/owned-input"] || uploaded != 1) {
				t.Errorf("uploaded %d deleted %v", uploaded, deleted)
			}
			for _, id := range []string{"output-1", "error-1", "output-2", "error-2"} {
				if !deleted["/v1/files/"+id] {
					t.Errorf("not cleaned: %s", id)
				}
			}
		})
	}
}

func TestBatchRunAbortCancelsAndPreservesResults(t *testing.T) {
	for _, downloadFailure := range []bool{false, true} {
		t.Run(fmt.Sprint(downloadFailure), func(t *testing.T) {
			cancelled := false
			deleted := []string{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch {
				case r.Method == "POST" && r.URL.Path == "/v1/batch/jobs":
					fmt.Fprint(w, `{"id":"job","status":"QUEUED"}`)
				case r.Method == "POST" && strings.HasSuffix(r.URL.Path, "/cancel"):
					cancelled = true
					fmt.Fprint(w, `{"id":"job","status":"CANCELLATION_REQUESTED"}`)
				case r.Method == "GET" && r.URL.Path == "/v1/batch/jobs/job":
					if downloadFailure {
						fmt.Fprint(w, `{"id":"job","status":"SUCCESS","output_file":"recoverable"}`)
					} else {
						fmt.Fprint(w, `{"id":"job","status":"RUNNING"}`)
					}
				case r.Method == "GET":
					http.Error(w, "download failed", 500)
				case r.Method == "DELETE":
					deleted = append(deleted, r.URL.Path)
					w.WriteHeader(204)
				default:
					t.Errorf("unexpected %s %s", r.Method, r.URL)
					w.WriteHeader(400)
				}
			}))
			defer server.Close()
			client := NewBatchClient[batchTestBody](NewMistralClient("test", server.URL, 1, time.Second), BatchEndpointEmbeddings)
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()
			_, err := client.Run(ctx, BatchInputFile{FileID: "caller-input"}, &BatchRunOptions{PollInterval: time.Millisecond})
			if err == nil {
				t.Fatal("expected failure")
			}
			if cancelled == downloadFailure {
				t.Errorf("cancelled %v", cancelled)
			}
			if len(deleted) != 0 {
				t.Errorf("files lost on failure: %v", deleted)
			}
		})
	}
}

func TestBatchStreamingLargeLinesAndEarlyStop(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != "GET" {
			t.Errorf("stream must only read: %s", r.Method)
		}
		fmt.Fprintf(w, "{\"custom_id\":\"large\",\"response\":{\"body\":{\"value\":1,\"padding\":%q}}}\n", strings.Repeat("x", 100000))
	}))
	defer server.Close()
	client := NewBatchClient[batchTestBody](NewMistralClient("test", server.URL, 1, time.Second), BatchEndpointEmbeddings)
	output, errorFile := "out", "err"
	h := &BatchJobHandle{Job: BatchJobOut{Status: BatchJobStatusSuccess, OutputFile: &output, ErrorFile: &errorFile}}
	stop := errors.New("stop")
	err := client.StreamResults(context.Background(), h, 0, func(item BatchItem[batchTestBody]) error {
		if item.Err != nil || item.Response.Value != 1 {
			t.Errorf("item %v", item)
		}
		return stop
	})
	if !errors.Is(err, stop) || calls != 1 {
		t.Fatalf("stop: calls=%d err=%v", calls, err)
	}
}

func TestBatchInputEmptyAndMissingIDs(t *testing.T) {
	input := BatchInput{Raw: []byte("{\"custom_id\":\"\",\"body\":{}}\n{\"body\":{}}")}
	requests, err := input.ToRequests()
	if err != nil {
		t.Fatal(err)
	}
	if requests[0].CustomID == nil || *requests[0].CustomID != "" || requests[1].CustomID != nil {
		t.Fatalf("IDs %+v", requests)
	}
	raw, err := json.Marshal(requests)
	if err != nil || !strings.Contains(string(raw), `"custom_id":""`) {
		t.Fatalf("empty ID lost: %s %v", raw, err)
	}
}
