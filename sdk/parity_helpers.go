package sdk

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
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
		case *Order:
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
		case *Order:
			if typed != nil {
				query.Add(key, string(*typed))
			}
		case *DeploymentStatus:
			if typed != nil {
				query.Add(key, string(*typed))
			}
		case *WorkflowRunSortBy:
			if typed != nil {
				query.Add(key, string(*typed))
			}
		case *WorkflowExecutionStatus:
			if typed != nil {
				query.Add(key, string(*typed))
			}
		case *time.Time:
			if typed != nil {
				query.Add(key, typed.Format(time.RFC3339Nano))
			}
		case []string:
			for _, item := range typed {
				query.Add(key, item)
			}
		case []WorkflowExecutionStatus:
			for _, item := range typed {
				query.Add(key, string(item))
			}
		case string:
			query.Add(key, typed)
		case int:
			query.Add(key, strconv.Itoa(typed))
		case int64:
			query.Add(key, strconv.FormatInt(typed, 10))
		case bool:
			query.Add(key, strconv.FormatBool(typed))
		case Order:
			query.Add(key, string(typed))
		case DeploymentStatus:
			query.Add(key, string(typed))
		case WorkflowRunSortBy:
			query.Add(key, string(typed))
		case WorkflowExecutionStatus:
			query.Add(key, string(typed))
		case time.Time:
			query.Add(key, typed.Format(time.RFC3339Nano))
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
		var eventName string
		var dataLines []string
		emit := func() bool {
			data := strings.TrimSpace(strings.Join(dataLines, "\n"))
			defer func() { eventName = ""; dataLines = nil }()
			if eventName == "error" {
				payload := map[string]any{}
				_ = json.Unmarshal([]byte(data), &payload)
				message, _ := payload["error"].(string)
				if message == "" {
					message = data
				}
				reason, _ := payload["reason"].(string)
				if reason == "" {
					reason = "stream_error"
				}
				out <- StreamEvent{Error: &StreamDisconnectedError{Reason: reason, ErrorMessage: message}}
				return true
			}
			if data == "" {
				return false
			}
			if data == "[DONE]" {
				return true
			}
			var payload map[string]any
			if err := json.Unmarshal([]byte(data), &payload); err != nil {
				out <- StreamEvent{Error: fmt.Errorf("error decoding stream event: %w", err)}
				return false
			}
			event := StreamEvent{Data: payload}
			if eventType, ok := payload["type"].(string); ok {
				event.Type = eventType
			}
			out <- event
			return false
		}
		for {
			line, readErr := reader.ReadString('\n')
			if readErr != nil {
				if readErr == io.EOF {
					if strings.TrimSpace(line) != "" {
						if strings.HasPrefix(line, "data:") {
							dataLines = append(dataLines, strings.TrimSpace(strings.TrimPrefix(line, "data:")))
						}
					}
					emit()
					return
				}
				out <- StreamEvent{Error: fmt.Errorf("error reading stream response: %w", readErr)}
				return
			}
			line = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
			if line == "" {
				if emit() {
					return
				}
				continue
			}
			if strings.HasPrefix(line, ":") {
				continue
			}
			field, value, found := strings.Cut(line, ":")
			if !found {
				continue
			}
			value = strings.TrimPrefix(value, " ")
			switch field {
			case "event":
				eventName = value
			case "data":
				dataLines = append(dataLines, value)
			}
		}
	}()
	return out
}
