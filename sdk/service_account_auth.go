package sdk

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
)

const ServiceAccountTokenPathEnv = "MISTRAL_SA_TOKEN_PATH"

type ServiceAccountTokenError struct {
	Path string
	Err  error
}

func (e *ServiceAccountTokenError) Error() string {
	return fmt.Sprintf("cannot read service-account token from %q: %v", e.Path, e.Err)
}
func (e *ServiceAccountTokenError) Unwrap() error { return e.Err }

// ReadServiceAccountToken reads on every call so projected token rotation works.
// An unset path returns an empty token. A configured unreadable/empty file fails.
func ReadServiceAccountToken() (string, error) {
	path := os.Getenv(ServiceAccountTokenPathEnv)
	if path == "" {
		return "", nil
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", &ServiceAccountTokenError{path, err}
	}
	token := strings.TrimSpace(string(raw))
	if token == "" || !utf8.Valid(raw) {
		return "", &ServiceAccountTokenError{path, fmt.Errorf("empty or invalid UTF-8 token")}
	}
	return token, nil
}
func bearerHeader(token string) string {
	if strings.HasPrefix(strings.ToLower(token), "bearer ") {
		return token
	}
	return "Bearer " + token
}

// WithHeaders supplies per-request overrides on an independent client copy.
func (c *MistralClient) WithHeaders(headers http.Header) *MistralClient {
	copy := *c
	copy.requestHeaders = headers.Clone()
	return &copy
}
func (c *MistralClient) authorizeRequest(req *http.Request) error {
	if req.Header == nil {
		req.Header = make(http.Header)
	}
	for key, values := range c.requestHeaders {
		req.Header.Del(key)
		for _, v := range values {
			req.Header.Add(key, v)
		}
	}
	for key := range c.requestHeaders {
		if strings.EqualFold(key, "Authorization") {
			return nil
		}
	}
	existing := req.Header.Get("Authorization")
	if existing != "" && existing != "Bearer "+c.apiKey && existing != bearerHeader(c.apiKey) {
		return nil
	}
	if c.explicitAPIKey {
		req.Header.Set("Authorization", bearerHeader(c.apiKey))
		return nil
	}
	token, err := ReadServiceAccountToken()
	if err != nil {
		return err
	}
	if token == "" {
		token = os.Getenv("MISTRAL_API_KEY")
	}
	if token != "" {
		req.Header.Set("Authorization", bearerHeader(token))
	}
	return nil
}
