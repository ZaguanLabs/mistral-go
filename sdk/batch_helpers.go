package sdk

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// BatchInput owns JSONL bytes. Raw input is preserved verbatim for upload.
type BatchInput struct{ Raw []byte }
type BatchInputFile struct {
	FileID string `json:"file_id"`
}

func NewBatchInput[T any](payload map[string]T) (BatchInput, error) {
	keys := make([]string, 0, len(payload))
	for k := range payload {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer
	for _, k := range keys {
		if err := json.NewEncoder(&b).Encode(struct {
			CustomID string `json:"custom_id"`
			Body     T      `json:"body"`
		}{k, payload[k]}); err != nil {
			return BatchInput{}, err
		}
	}
	return BatchInput{Raw: bytes.TrimSuffix(b.Bytes(), []byte("\n"))}, nil
}
func (b BatchInput) Len() int {
	n := 0
	for _, line := range bytes.Split(b.Raw, []byte("\n")) {
		if len(bytes.TrimSpace(line)) > 0 {
			n++
		}
	}
	return n
}
func (b BatchInput) ToJSONLBytes() []byte { return append([]byte(nil), b.Raw...) }

// BatchInputRequest preserves missing, null, and empty custom IDs on inline inputs.
type BatchInputRequest struct {
	CustomID *string        `json:"custom_id"`
	Body     map[string]any `json:"body"`
}

func (b BatchInput) ToRequests() ([]BatchInputRequest, error) {
	result := []BatchInputRequest{}
	for _, line := range bytes.Split(b.Raw, []byte("\n")) {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var r BatchInputRequest
		if err := json.Unmarshal(line, &r); err != nil {
			return nil, err
		}
		if r.Body == nil {
			return nil, fmt.Errorf("batch request is missing body")
		}
		result = append(result, r)
	}
	return result, nil
}

type BatchRequestError struct {
	CustomID   *string
	StatusCode *int
	ErrorData  json.RawMessage
}

func (e *BatchRequestError) Error() string {
	return fmt.Sprintf("batch request failed: %s", e.ErrorData)
}

type BatchResponseError struct {
	CustomID   *string
	StatusCode *int
	Body       json.RawMessage
	Err        error
}

func (e *BatchResponseError) Error() string {
	return fmt.Sprintf("batch response could not be decoded: %v", e.Err)
}
func (e *BatchResponseError) Unwrap() error { return e.Err }

type BatchError struct{ Failures map[string]error }

func (e *BatchError) Error() string {
	return fmt.Sprintf("batch output has no successful result for %d requested IDs", len(e.Failures))
}

type BatchItem[T any] struct {
	CustomID *string
	Response *T
	Err      error
}
type BatchOutput[T any] struct {
	Responses      map[string]T
	RequestErrors  []*BatchRequestError
	ResponseErrors []*BatchResponseError
}

func parseBatchItem[T any](raw []byte) (BatchItem[T], error) {
	var line struct {
		CustomID *string         `json:"custom_id"`
		Error    json.RawMessage `json:"error"`
		Response struct {
			StatusCode *int            `json:"status_code"`
			Body       json.RawMessage `json:"body"`
		} `json:"response"`
	}
	if err := json.Unmarshal(raw, &line); err != nil {
		return BatchItem[T]{}, err
	}
	item := BatchItem[T]{CustomID: line.CustomID}
	if (len(line.Error) > 0 && string(line.Error) != "null") || (line.Response.StatusCode != nil && *line.Response.StatusCode >= 400) {
		item.Err = &BatchRequestError{line.CustomID, line.Response.StatusCode, line.Error}
		return item, nil
	}
	var response T
	err := json.Unmarshal(line.Response.Body, &response)
	if bytes.Equal(bytes.TrimSpace(line.Response.Body), []byte("null")) {
		err = fmt.Errorf("response body is null")
	}
	if err == nil && line.CustomID == nil {
		err = fmt.Errorf("response missing custom_id")
	}
	if err != nil {
		item.Err = &BatchResponseError{line.CustomID, line.Response.StatusCode, line.Response.Body, err}
		return item, nil
	}
	item.Response = &response
	return item, nil
}
func ParseBatchOutput[T any](raw []byte) (*BatchOutput[T], error) {
	out := &BatchOutput[T]{Responses: map[string]T{}}
	err := readBatchLines(bytes.NewReader(raw), func(line []byte) error {
		item, err := parseBatchItem[T](line)
		if err != nil {
			return err
		}
		switch e := item.Err.(type) {
		case *BatchRequestError:
			out.RequestErrors = append(out.RequestErrors, e)
		case *BatchResponseError:
			out.ResponseErrors = append(out.ResponseErrors, e)
		default:
			id := *item.CustomID
			if _, exists := out.Responses[id]; exists {
				return fmt.Errorf("duplicate custom_id %q", id)
			}
			out.Responses[id] = *item.Response
		}
		return nil
	})
	return out, err
}

// readBatchLines has no scanner token limit, so large embeddings remain readable.
func readBatchLines(r io.Reader, visit func([]byte) error) error {
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadBytes('\n')
		if len(bytes.TrimSpace(line)) > 0 {
			if e := visit(line); e != nil {
				return e
			}
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
func (o *BatchOutput[T]) Results() (map[string]BatchItem[T], error) {
	result := map[string]BatchItem[T]{}
	for id, r := range o.Responses {
		id, r := id, r
		result[id] = BatchItem[T]{CustomID: &id, Response: &r}
	}
	add := func(id *string, e error) error {
		if id == nil {
			return nil
		}
		if _, ok := result[*id]; ok {
			return fmt.Errorf("duplicate custom_id %q", *id)
		}
		result[*id] = BatchItem[T]{CustomID: id, Err: e}
		return nil
	}
	for _, e := range o.RequestErrors {
		if err := add(e.CustomID, e); err != nil {
			return nil, err
		}
	}
	for _, e := range o.ResponseErrors {
		if err := add(e.CustomID, e); err != nil {
			return nil, err
		}
	}
	return result, nil
}
func (o *BatchOutput[T]) Ordered(ids []string) ([]T, error) {
	items, err := o.Results()
	if err != nil {
		return nil, err
	}
	failures := map[string]error{}
	result := make([]T, 0, len(ids))
	for _, id := range ids {
		item, ok := items[id]
		if !ok || item.Err != nil {
			failures[id] = item.Err
			continue
		}
		result = append(result, *item.Response)
	}
	if len(failures) > 0 {
		return nil, &BatchError{Failures: failures}
	}
	return result, nil
}
func IsBatchRunning(status BatchJobStatus) bool {
	return status == BatchJobStatusQueued || status == BatchJobStatusRunning || status == BatchJobStatusCancellationRequested
}
func IsBatchTerminal(status BatchJobStatus) bool { return !IsBatchRunning(status) }

type BatchJobHandle struct {
	Job BatchJobOut `json:"job"`
}
type BatchResult[T any] struct {
	Job         BatchJobOut
	Output      *BatchOutput[T]
	ErrorOutput *BatchOutput[T]
}

func (r *BatchResult[T]) Results() (map[string]BatchItem[T], error) {
	result, err := r.Output.Results()
	if err != nil {
		return nil, err
	}
	if r.ErrorOutput != nil {
		other, err := r.ErrorOutput.Results()
		if err != nil {
			return nil, err
		}
		for id, item := range other {
			if _, exists := result[id]; exists {
				return nil, fmt.Errorf("duplicate custom_id %q", id)
			}
			result[id] = item
		}
	}
	return result, nil
}
func (r *BatchResult[T]) ByID() (map[string]T, error) {
	items, err := r.Results()
	if err != nil {
		return nil, err
	}
	out := map[string]T{}
	for id, item := range items {
		if item.Err == nil {
			out[id] = *item.Response
		}
	}
	return out, nil
}
func (r *BatchResult[T]) Errors() []error {
	var out []error
	for _, o := range []*BatchOutput[T]{r.Output, r.ErrorOutput} {
		if o == nil {
			continue
		}
		for _, e := range o.RequestErrors {
			out = append(out, e)
		}
		for _, e := range o.ResponseErrors {
			out = append(out, e)
		}
	}
	return out
}
func (r *BatchResult[T]) Ordered(ids []string) ([]T, error) {
	items, err := r.Results()
	if err != nil {
		return nil, err
	}
	out := make([]T, 0, len(ids))
	failures := map[string]error{}
	for _, id := range ids {
		item, ok := items[id]
		if !ok || item.Err != nil {
			failures[id] = item.Err
		} else {
			out = append(out, *item.Response)
		}
	}
	if len(failures) > 0 {
		return nil, &BatchError{Failures: failures}
	}
	return out, nil
}

// BatchClient binds a batch endpoint to a response type. Every network operation
// accepts a context; goroutines provide the Python helper's asynchronous use case.
// HTTPHeaders may be configured before use, but must not be changed concurrently.
type BatchClient[T any] struct {
	Client      *MistralClient
	Endpoint    BatchEndpoint
	HTTPHeaders http.Header
}

func NewBatchClient[T any](client *MistralClient, endpoint BatchEndpoint) *BatchClient[T] {
	return &BatchClient[T]{Client: client, Endpoint: endpoint}
}

type BatchCreateOptions struct {
	Model        *string
	Metadata     map[string]any
	TimeoutHours int
}
type BatchRunOptions struct {
	BatchCreateOptions
	PollInterval        time.Duration
	MaxFailedJobRetries int
	InlineThreshold     *int
}

func (b *BatchClient[T]) do(ctx context.Context, method, path, contentType string, body []byte) (*http.Response, error) {
	if b.Client == nil {
		return nil, fmt.Errorf("batch client is nil")
	}
	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(b.Client.endpoint, "/")+"/"+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+b.Client.apiKey)
	req.Header.Set("User-Agent", UserAgent)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	for k, values := range b.HTTPHeaders {
		req.Header[k] = append([]string(nil), values...)
	}
	resp, err := b.Client.WithContext(ctx).doRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, NewMistralAPIError(string(raw), resp.StatusCode, resp.Header)
	}
	return resp, nil
}

func batchDelay(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
func (b *BatchClient[T]) json(ctx context.Context, method, path string, body any, result any) error {
	var raw []byte
	var err error
	if body != nil {
		raw, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	resp, err := b.do(ctx, method, path, "application/json", raw)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result == nil {
		_, err = io.Copy(io.Discard, resp.Body)
		return err
	}
	return json.NewDecoder(resp.Body).Decode(result)
}
func (b *BatchClient[T]) Upload(ctx context.Context, input BatchInput, filename string) (BatchInputFile, error) {
	if filename == "" {
		filename = "batch-input.jsonl"
	}
	var raw bytes.Buffer
	w := multipart.NewWriter(&raw)
	part, err := w.CreateFormFile("file", filename)
	if err != nil {
		return BatchInputFile{}, err
	}
	if _, err = part.Write(input.Raw); err != nil {
		return BatchInputFile{}, err
	}
	if err = w.WriteField("purpose", "batch"); err != nil {
		return BatchInputFile{}, err
	}
	if err = w.Close(); err != nil {
		return BatchInputFile{}, err
	}
	resp, err := b.do(ctx, http.MethodPost, "v1/files", w.FormDataContentType(), raw.Bytes())
	if err != nil {
		return BatchInputFile{}, err
	}
	defer resp.Body.Close()
	var file UploadFileOut
	err = json.NewDecoder(resp.Body).Decode(&file)
	return BatchInputFile{FileID: file.ID}, err
}
func (b *BatchClient[T]) Download(ctx context.Context, id string) ([]byte, error) {
	resp, err := b.do(ctx, http.MethodGet, "v1/files/"+url.PathEscape(id)+"/content", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
func (b *BatchClient[T]) SignedURL(ctx context.Context, id string, expiryHours int) (string, error) {
	if expiryHours == 0 {
		expiryHours = 24
	}
	var out FileSignedURL
	err := b.json(ctx, http.MethodGet, fmt.Sprintf("v1/files/%s/url?expiry=%d", url.PathEscape(id), expiryHours), nil, &out)
	return out.URL, err
}
func (b *BatchClient[T]) Delete(ctx context.Context, id string) error {
	return b.json(ctx, http.MethodDelete, "v1/files/"+url.PathEscape(id), nil, nil)
}
func (b *BatchClient[T]) Create(ctx context.Context, input any, opts *BatchCreateOptions) (*BatchJobHandle, error) {
	if opts == nil {
		opts = &BatchCreateOptions{}
	}
	hours := opts.TimeoutHours
	if hours == 0 {
		hours = 24
	}
	body := map[string]any{"endpoint": b.Endpoint, "timeout_hours": hours}
	if opts.Model != nil {
		body["model"] = *opts.Model
	}
	if opts.Metadata != nil {
		body["metadata"] = opts.Metadata
	}
	switch v := input.(type) {
	case BatchInput:
		requests, err := v.ToRequests()
		if err != nil {
			return nil, err
		}
		body["requests"] = requests
	case BatchInputFile:
		body["input_files"] = []string{v.FileID}
	default:
		return nil, fmt.Errorf("input must be BatchInput or BatchInputFile")
	}
	var job BatchJobOut
	if err := b.json(ctx, http.MethodPost, "v1/batch/jobs", body, &job); err != nil {
		return nil, err
	}
	return &BatchJobHandle{Job: job}, nil
}
func (b *BatchClient[T]) Get(ctx context.Context, id string) (*BatchJobHandle, error) {
	var job BatchJobOut
	if err := b.json(ctx, http.MethodGet, "v1/batch/jobs/"+url.PathEscape(id), nil, &job); err != nil {
		return nil, err
	}
	return &BatchJobHandle{Job: job}, nil
}
func (b *BatchClient[T]) Refresh(ctx context.Context, h *BatchJobHandle) (*BatchJobHandle, error) {
	if h == nil {
		return nil, fmt.Errorf("handle cannot be nil")
	}
	return b.Get(ctx, h.Job.ID)
}
func (b *BatchClient[T]) Cancel(ctx context.Context, id string) (*BatchJobHandle, error) {
	var job BatchJobOut
	if err := b.json(ctx, http.MethodPost, "v1/batch/jobs/"+url.PathEscape(id)+"/cancel", nil, &job); err != nil {
		return nil, err
	}
	return &BatchJobHandle{Job: job}, nil
}
func (b *BatchClient[T]) Wait(ctx context.Context, id string, pollInterval time.Duration) (*BatchJobHandle, error) {
	if pollInterval <= 0 {
		pollInterval = 30 * time.Second
	}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 24*time.Hour)
		defer cancel()
	}
	for {
		h, err := b.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		if IsBatchTerminal(h.Job.Status) {
			return h, nil
		}
		if err := batchDelay(ctx, pollInterval); err != nil {
			return nil, err
		}
	}
}
func (b *BatchClient[T]) terminal(ctx context.Context, h *BatchJobHandle, poll time.Duration) (*BatchJobHandle, error) {
	if h == nil {
		return nil, fmt.Errorf("handle cannot be nil")
	}
	if IsBatchTerminal(h.Job.Status) {
		return h, nil
	}
	return b.Wait(ctx, h.Job.ID, poll)
}
func (b *BatchClient[T]) Output(ctx context.Context, h *BatchJobHandle, poll time.Duration) (*BatchOutput[T], error) {
	h, err := b.terminal(ctx, h, poll)
	if err != nil {
		return nil, err
	}
	if h.Job.OutputFile == nil || *h.Job.OutputFile == "" {
		return ParseBatchOutput[T](nil)
	}
	raw, err := b.Download(ctx, *h.Job.OutputFile)
	if err != nil {
		return nil, err
	}
	return ParseBatchOutput[T](raw)
}
func (b *BatchClient[T]) ErrorOutput(ctx context.Context, h *BatchJobHandle, poll time.Duration) (*BatchOutput[T], error) {
	h, err := b.terminal(ctx, h, poll)
	if err != nil {
		return nil, err
	}
	if h.Job.ErrorFile == nil || *h.Job.ErrorFile == "" {
		return nil, nil
	}
	raw, err := b.Download(ctx, *h.Job.ErrorFile)
	if err != nil {
		return nil, err
	}
	return ParseBatchOutput[T](raw)
}

// StreamResults reads both result files one line at a time and never deletes them.
func (b *BatchClient[T]) StreamResults(ctx context.Context, h *BatchJobHandle, poll time.Duration, visit func(BatchItem[T]) error) error {
	if visit == nil {
		return fmt.Errorf("visitor cannot be nil")
	}
	h, err := b.terminal(ctx, h, poll)
	if err != nil {
		return err
	}
	for _, id := range []*string{h.Job.OutputFile, h.Job.ErrorFile} {
		if id == nil || *id == "" {
			continue
		}
		resp, err := b.do(ctx, http.MethodGet, "v1/files/"+url.PathEscape(*id)+"/content", "", nil)
		if err != nil {
			return err
		}
		err = readBatchLines(resp.Body, func(line []byte) error {
			item, err := parseBatchItem[T](line)
			if err != nil {
				return err
			}
			return visit(item)
		})
		resp.Body.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// Run owns only files it uploads and the jobs' result files. It materializes
// results before cleanup, and cancels an in-flight job if the caller aborts.
// MaxFailedJobRetries resubmits every request and can incur additional charges.
func (b *BatchClient[T]) Run(ctx context.Context, input any, opts *BatchRunOptions) (result *BatchResult[T], err error) {
	if opts == nil {
		opts = &BatchRunOptions{}
	}
	if opts.MaxFailedJobRetries < 0 {
		return nil, fmt.Errorf("retry count cannot be negative")
	}
	source := input
	switch v := input.(type) {
	case []byte:
		source = BatchInput{Raw: v}
	case string:
		source = BatchInput{Raw: []byte(v)}
	case map[string]any:
		source, err = NewBatchInput(v)
		if err != nil {
			return nil, err
		}
	}
	threshold := 1000
	if opts.InlineThreshold != nil {
		threshold = *opts.InlineThreshold
	}
	createdInput := ""
	if v, ok := source.(BatchInput); ok && v.Len() > threshold {
		var file BatchInputFile
		file, err = b.Upload(ctx, v, "")
		if err != nil {
			return nil, err
		}
		createdInput = file.FileID
		source = file
	}
	hours := opts.TimeoutHours
	if hours == 0 {
		hours = 24
	}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(hours)*time.Hour)
	defer cancel()
	var handle *BatchJobHandle
	var files []string
	defer func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		if err != nil {
			if handle != nil && IsBatchRunning(handle.Job.Status) {
				_, _ = b.Cancel(cleanupCtx, handle.Job.ID)
			}
			if createdInput != "" {
				_ = b.Delete(cleanupCtx, createdInput)
			}
			return
		}
		if createdInput != "" {
			files = append(files, createdInput)
		}
		for _, id := range files {
			_ = b.Delete(cleanupCtx, id)
		}
	}()
	for attempt := 0; attempt <= opts.MaxFailedJobRetries; attempt++ {
		deadline, _ := ctx.Deadline()
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return nil, context.DeadlineExceeded
		}
		create := opts.BatchCreateOptions
		create.TimeoutHours = int((remaining + time.Hour - 1) / time.Hour)
		handle, err = b.Create(ctx, source, &create)
		if err != nil {
			return nil, err
		}
		settled, waitErr := b.Wait(ctx, handle.Job.ID, opts.PollInterval)
		if waitErr != nil {
			return nil, waitErr
		}
		handle = settled
		for _, id := range []*string{handle.Job.OutputFile, handle.Job.ErrorFile} {
			if id != nil && *id != "" {
				files = append(files, *id)
			}
		}
		if handle.Job.FailedRequests == 0 {
			break
		}
	}
	output, err := b.Output(ctx, handle, opts.PollInterval)
	if err != nil {
		return nil, err
	}
	errors, err := b.ErrorOutput(ctx, handle, opts.PollInterval)
	if err != nil {
		return nil, err
	}
	return &BatchResult[T]{Job: handle.Job, Output: output, ErrorOutput: errors}, nil
}
