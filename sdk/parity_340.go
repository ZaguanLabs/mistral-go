package sdk

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type AudioAgentListParams struct {
	PageSize  *int
	PageToken *string
	Search    *string
}

func audioAgentQuery(params *AudioAgentListParams) string {
	if params == nil {
		params = &AudioAgentListParams{}
	}
	size := 10
	if params.PageSize != nil {
		size = *params.PageSize
	}
	return queryWithOptionalValues(map[string]any{"page_size": size, "page_token": params.PageToken, "search": params.Search})
}

func (c *MistralClient) ListAudioAgents(params *AudioAgentListParams) (*AudioAgentListPage, error) {
	return requestTyped[AudioAgentListPage](c, http.MethodGet, nil, appendQuery("v1/audio/agents", audioAgentQuery(params)))
}
func (c *MistralClient) CreateAudioAgent(request *AudioAgentCreateRequest) (*AudioAgentResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	return requestTyped[AudioAgentResponse](c, http.MethodPost, request, "v1/audio/agents")
}
func (c *MistralClient) GetAudioAgent(id string, version *int) (*AudioAgentResponse, error) {
	return requestTyped[AudioAgentResponse](c, http.MethodGet, nil, appendQuery("v1/audio/agents/"+url.PathEscape(id), queryWithOptionalValues(map[string]any{"version": version})))
}
func (c *MistralClient) UpdateAudioAgent(id string, request *AudioAgentUpdateRequest) (*AudioAgentResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	return requestTyped[AudioAgentResponse](c, http.MethodPut, request, "v1/audio/agents/"+url.PathEscape(id))
}
func (c *MistralClient) DeleteAudioAgent(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/audio/agents/"+url.PathEscape(id), false, nil)
	return err
}
func (c *MistralClient) ListAudioAgentVersions(id string, params *AudioAgentListParams) (*AudioAgentListPage, error) {
	return requestTyped[AudioAgentListPage](c, http.MethodGet, nil, appendQuery("v1/audio/agents/"+url.PathEscape(id)+"/versions", audioAgentQuery(params)))
}
func (c *MistralClient) ListAudioAgentSessions(id string, params *AudioAgentListParams) (*AudioAgentSessionListPage, error) {
	return requestTyped[AudioAgentSessionListPage](c, http.MethodGet, nil, appendQuery("v1/audio/agents/"+url.PathEscape(id)+"/sessions", audioAgentQuery(params)))
}
func (c *MistralClient) CreateAudioSession(request *AudioSessionCreateRequest) (*AudioSessionCreateResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	return requestTyped[AudioSessionCreateResponse](c, http.MethodPost, request, "v1/audio/sessions")
}
func (c *MistralClient) GetAudioSession(id string) (*AudioSessionResponse, error) {
	return requestTyped[AudioSessionResponse](c, http.MethodGet, nil, "v1/audio/sessions/"+url.PathEscape(id))
}
func (c *MistralClient) CloseAudioSession(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/audio/sessions/"+url.PathEscape(id), false, nil)
	return err
}

func (c *MistralClient) CreateResponse(request *CreateResponseRequest) (*CreateResponseResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	copy := *request
	stream := false
	copy.Stream = &stream
	return requestTyped[CreateResponseResponse](c, http.MethodPost, &copy, "v1/responses")
}
func (c *MistralClient) CountResponseInputTokens(request *CountResponseInputTokensRequest) (*CountResponseInputTokensResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	return requestTyped[CountResponseInputTokensResponse](c, http.MethodPost, request, "v1/responses/input_tokens")
}

// CreateResponseStream returns a closable SSE stream with complete event envelopes.
func (c *MistralClient) CreateResponseStream(request *CreateResponseRequest) (*EventStream, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	copy := *request
	stream := true
	copy.Stream = &stream
	body, err := structToMap(&copy)
	if err != nil {
		return nil, err
	}
	response, err := c.withAccept("text/event-stream").request(http.MethodPost, body, "v1/responses", true, nil)
	if err != nil {
		return nil, err
	}
	reader, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("unexpected stream response %T", response)
	}
	return NewEventStream(reader), nil
}

func (c *MistralClient) GetManagedIndexSummary(name, language string) (*GetSummaryResponseSummary, error) {
	return requestTyped[GetSummaryResponseSummary](c, http.MethodGet, nil, managedIndexPath(name)+"/summary_field/"+url.PathEscape(language))
}
func (c *MistralClient) SetManagedIndexSummary(name, language string, request *UpdateSummaryRequestSummary) (any, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	body, err := structToMap(request)
	if err != nil {
		return nil, err
	}
	return c.request(http.MethodPut, body, managedIndexPath(name)+"/summary_field/"+url.PathEscape(language), false, nil)
}

// JSONLStream reads one JSON value at a time and preserves unknown event fields.
type JSONLStream struct {
	body    io.ReadCloser
	decoder *json.Decoder
}

func (s *JSONLStream) Next() (map[string]any, error) {
	var value map[string]any
	if err := s.decoder.Decode(&value); err != nil {
		s.Close()
		return nil, err
	}
	return value, nil
}
func (s *JSONLStream) Close() error { return s.body.Close() }
func (c *MistralClient) GenerateManagedIndexSummary(name, language string, force *bool) (*JSONLStream, error) {
	enabled := false
	if force != nil {
		enabled = *force
	}
	path := appendQuery(managedIndexPath(name)+"/summary_field/"+url.PathEscape(language), queryWithOptionalValues(map[string]any{"force": enabled}))
	response, err := c.withAccept("application/x-ndjson").request(http.MethodPost, nil, path, true, nil)
	if err != nil {
		return nil, err
	}
	reader, ok := response.(io.ReadCloser)
	if !ok {
		return nil, fmt.Errorf("unexpected stream response %T", response)
	}
	return &JSONLStream{body: reader, decoder: json.NewDecoder(bufio.NewReader(reader))}, nil
}

func (c *MistralClient) withAccept(accept string) *MistralClient {
	headers := c.requestHeaders.Clone()
	if headers == nil {
		headers = http.Header{}
	}
	headers.Set("Accept", accept)
	return c.WithHeaders(headers)
}
