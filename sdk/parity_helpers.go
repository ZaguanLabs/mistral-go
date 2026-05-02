package sdk

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type APIResponse map[string]any

type StreamEvent struct {
	Type  string         `json:"type,omitempty"`
	Data  map[string]any `json:"data,omitempty"`
	Error error          `json:"-"`
}

func optionalRequestMap(values map[string]any) map[string]interface{} {
	req := make(map[string]interface{})
	for key, value := range values {
		if value == nil {
			continue
		}
		switch typed := value.(type) {
		case *string:
			if typed != nil {
				req[key] = *typed
			}
		case *int:
			if typed != nil {
				req[key] = *typed
			}
		case *int64:
			if typed != nil {
				req[key] = *typed
			}
		case *float64:
			if typed != nil {
				req[key] = *typed
			}
		case *bool:
			if typed != nil {
				req[key] = *typed
			}
		default:
			req[key] = value
		}
	}
	return req
}

func queryWithOptionalValues(values map[string]any) string {
	query := url.Values{}
	for key, value := range values {
		if value == nil {
			continue
		}
		switch typed := value.(type) {
		case *string:
			if typed != nil {
				query.Add(key, *typed)
			}
		case *int:
			if typed != nil {
				query.Add(key, strconv.Itoa(*typed))
			}
		case *int64:
			if typed != nil {
				query.Add(key, strconv.FormatInt(*typed, 10))
			}
		case *bool:
			if typed != nil {
				query.Add(key, strconv.FormatBool(*typed))
			}
		case []string:
			for _, item := range typed {
				query.Add(key, item)
			}
		case string:
			query.Add(key, typed)
		case int:
			query.Add(key, strconv.Itoa(typed))
		case int64:
			query.Add(key, strconv.FormatInt(typed, 10))
		case bool:
			query.Add(key, strconv.FormatBool(typed))
		}
	}
	return query.Encode()
}

func appendQuery(path string, query string) string {
	if query == "" {
		return path
	}
	return path + "?" + query
}

func (c *MistralClient) requestMap(method string, body map[string]interface{}, path string) (APIResponse, error) {
	response, err := c.request(method, body, path, false, nil)
	if err != nil {
		return nil, err
	}
	respData, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	return APIResponse(respData), nil
}

func (c *MistralClient) requestBytes(method string, path string, accept string) ([]byte, error) {
	req, err := http.NewRequest(method, c.endpoint+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	req.Header.Set("User-Agent", UserAgent)
	client := &http.Client{Timeout: c.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, NewMistralConnectionError(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, NewMistralAPIError(string(body), resp.StatusCode, resp.Header)
	}
	return body, nil
}

func parseGenericStream(body io.ReadCloser) <-chan StreamEvent {
	out := make(chan StreamEvent)
	go func() {
		defer close(out)
		defer body.Close()
		reader := bufio.NewReader(body)
		for {
			line, readErr := reader.ReadBytes('\n')
			if readErr == io.EOF {
				break
			}
			if readErr != nil {
				out <- StreamEvent{Error: fmt.Errorf("error reading stream response: %w", readErr)}
				return
			}
			if bytes.Equal(line, []byte("\n")) || !bytes.HasPrefix(line, []byte("data: ")) {
				continue
			}
			jsonLine := bytes.TrimSpace(bytes.TrimPrefix(line, []byte("data: ")))
			if bytes.Equal(jsonLine, []byte("[DONE]")) {
				break
			}
			var payload map[string]any
			if err := json.Unmarshal(jsonLine, &payload); err != nil {
				out <- StreamEvent{Error: fmt.Errorf("error decoding stream event: %w", err)}
				continue
			}
			event := StreamEvent{Data: payload}
			if eventType, ok := payload["type"].(string); ok {
				event.Type = eventType
			}
			out <- event
		}
	}()
	return out
}
