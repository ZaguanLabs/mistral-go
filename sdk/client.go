package sdk

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	Endpoint          = "https://api.mistral.ai"
	CodestralEndpoint = "https://codestral.mistral.ai"
	DefaultMaxRetries = 5
	DefaultTimeout    = 120 * time.Second
)

var retryStatusCodes = map[int]bool{
	429: true,
	500: true,
	502: true,
	503: true,
	504: true,
}

type MistralClient struct {
	retryConfig *RetryConfig
	ctx         context.Context
	apiKey      string
	endpoint    string
	maxRetries  int
	timeout     time.Duration
}

func NewMistralClient(apiKey string, endpoint string, maxRetries int, timeout time.Duration) *MistralClient {
	if apiKey == "" {
		apiKey = os.Getenv("MISTRAL_API_KEY")
	}
	if endpoint == "" {
		endpoint = Endpoint
	}
	if maxRetries == 0 {
		maxRetries = DefaultMaxRetries
	}
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return &MistralClient{
		apiKey:     apiKey,
		endpoint:   endpoint,
		maxRetries: maxRetries,
		timeout:    timeout,
	}
}

// NewMistralClientDefault creates a new Mistral API client with the default endpoint and the given API key. Defaults to using MISTRAL_API_KEY from the environment.
func NewMistralClientDefault(apiKey string) *MistralClient {
	if apiKey == "" {
		apiKey = os.Getenv("MISTRAL_API_KEY")
	}

	return NewMistralClient(apiKey, Endpoint, DefaultMaxRetries, DefaultTimeout)
}

// NewCodestralClientDefault creates a new Codestral API client with the default endpoint and the given API key. Defaults to using CODESTRAL_API_KEY from the environment.
func NewCodestralClientDefault(apiKey string) *MistralClient {
	if apiKey == "" {
		apiKey = os.Getenv("CODESTRAL_API_KEY")
	}

	return NewMistralClient(apiKey, CodestralEndpoint, DefaultMaxRetries, DefaultTimeout)
}

func (c *MistralClient) request(method string, jsonData map[string]interface{}, path string, stream bool, params map[string]string) (interface{}, error) {
	uri, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	if pathURL, parseErr := url.Parse(path); parseErr == nil {
		if pathURL.Path != "" {
			uri.Path = pathURL.Path
			uri.RawPath = pathURL.RawPath
		}
		uri.RawQuery = pathURL.RawQuery
	} else {
		uri.Path = path
	}

	// The execute API carries trace context in both the body and header.
	traceparent := ""
	if method == http.MethodPost && strings.HasPrefix(strings.TrimPrefix(uri.Path, "/"), "v1/workflows/") && strings.HasSuffix(uri.Path, "/execute") {
		body := make(map[string]interface{}, len(jsonData)+1)
		for key, value := range jsonData {
			body[key] = value
		}
		traceparent, _ = body["traceparent"].(string)
		if traceparent == "" {
			var trace [24]byte
			if _, err := rand.Read(trace[:]); err != nil {
				return nil, err
			}
			traceparent = fmt.Sprintf("00-%x-%x-01", trace[:16], trace[16:])
			body["traceparent"] = traceparent
		}
		jsonData = body
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", UserAgent)
	if traceparent != "" {
		req.Header.Set("traceparent", traceparent)
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		responseBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("(HTTP Error %d) %s", resp.StatusCode, string(responseBytes))
	}

	if stream {
		return resp.Body, nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(bytes.TrimSpace(body)) == 0 {
		return map[string]interface{}{}, nil
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
