package sdk

import (
	"os"
	"testing"
	"time"
)

func TestNewMistralClient(t *testing.T) {
	client := NewMistralClient("test-key", Endpoint, 3, 60*time.Second)

	if client == nil {
		t.Fatal("NewMistralClient returned nil")
	}

	if client.apiKey != "test-key" {
		t.Error("API key not set correctly")
	}

	if client.endpoint != Endpoint {
		t.Error("Endpoint not set correctly")
	}

	if client.maxRetries != 3 {
		t.Error("MaxRetries not set correctly")
	}

	if client.timeout != 60*time.Second {
		t.Error("Timeout not set correctly")
	}
}

func TestNewMistralClientDefaults(t *testing.T) {
	client := NewMistralClient("", "", 0, 0)

	if client.endpoint != Endpoint {
		t.Error("Should use default endpoint")
	}

	if client.maxRetries != DefaultMaxRetries {
		t.Error("Should use default max retries")
	}

	if client.timeout != DefaultTimeout {
		t.Error("Should use default timeout")
	}
}

func TestNewMistralClientFromEnv(t *testing.T) {
	// Set environment variable
	os.Setenv("MISTRAL_API_KEY", "env-test-key")
	defer os.Unsetenv("MISTRAL_API_KEY")

	client := NewMistralClient("", "", 0, 0)

	if client.apiKey != "env-test-key" {
		t.Error("Should load API key from environment")
	}
}

func TestNewMistralClientDefault(t *testing.T) {
	client := NewMistralClientDefault("test-key")

	if client == nil {
		t.Fatal("NewMistralClientDefault returned nil")
	}

	if client.apiKey != "test-key" {
		t.Error("API key not set correctly")
	}

	if client.endpoint != Endpoint {
		t.Error("Should use default endpoint")
	}
}

func TestNewCodestralClientDefault(t *testing.T) {
	client := NewCodestralClientDefault("test-key")

	if client == nil {
		t.Fatal("NewCodestralClientDefault returned nil")
	}

	if client.endpoint != CodestralEndpoint {
		t.Error("Should use Codestral endpoint")
	}
}

func TestEndpointConstants(t *testing.T) {
	if Endpoint == "" {
		t.Error("Endpoint constant is empty")
	}

	if CodestralEndpoint == "" {
		t.Error("CodestralEndpoint constant is empty")
	}

	if Endpoint == CodestralEndpoint {
		t.Error("Endpoint and CodestralEndpoint should be different")
	}
}

func TestDefaultConstants(t *testing.T) {
	if DefaultMaxRetries <= 0 {
		t.Error("DefaultMaxRetries should be positive")
	}

	if DefaultTimeout <= 0 {
		t.Error("DefaultTimeout should be positive")
	}
}

func TestRetryStatusCodes(t *testing.T) {
	expectedCodes := []int{429, 500, 502, 503, 504}

	for _, code := range expectedCodes {
		if !retryStatusCodes[code] {
			t.Errorf("Status code %d should be in retry list", code)
		}
	}

	// Test that non-retry codes are not in the map
	nonRetryCodes := []int{200, 400, 401, 403, 404}
	for _, code := range nonRetryCodes {
		if retryStatusCodes[code] {
			t.Errorf("Status code %d should not be in retry list", code)
		}
	}
}

func TestMistralClientStructure(t *testing.T) {
	client := &MistralClient{
		apiKey:     "test",
		endpoint:   "https://test.com",
		maxRetries: 5,
		timeout:    120 * time.Second,
	}

	if client.apiKey != "test" {
		t.Error("apiKey field not accessible")
	}

	if client.endpoint != "https://test.com" {
		t.Error("endpoint field not accessible")
	}
}

func TestClientWithCustomTimeout(t *testing.T) {
	customTimeout := 30 * time.Second
	client := NewMistralClient("test-key", Endpoint, DefaultMaxRetries, customTimeout)

	if client.timeout != customTimeout {
		t.Error("Custom timeout not set correctly")
	}
}

func TestClientWithCustomRetries(t *testing.T) {
	customRetries := 10
	client := NewMistralClient("test-key", Endpoint, customRetries, DefaultTimeout)

	if client.maxRetries != customRetries {
		t.Error("Custom max retries not set correctly")
	}
}

func TestClientWithCustomEndpoint(t *testing.T) {
	customEndpoint := "https://custom.mistral.ai"
	client := NewMistralClient("test-key", customEndpoint, DefaultMaxRetries, DefaultTimeout)

	if client.endpoint != customEndpoint {
		t.Error("Custom endpoint not set correctly")
	}
}

func TestNewMistralClientDefaultWithEmptyKey(t *testing.T) {
	// Clear environment variable
	os.Unsetenv("MISTRAL_API_KEY")

	client := NewMistralClientDefault("")

	if client == nil {
		t.Fatal("Should create client even with empty key")
	}

	// API key will be empty, which will cause API calls to fail
	// but the client should still be created
	if client.apiKey != "" {
		t.Error("API key should be empty when not provided")
	}
}

func TestClientAPIKeyPrecedence(t *testing.T) {
	// Set environment variable
	os.Setenv("MISTRAL_API_KEY", "env-key")
	defer os.Unsetenv("MISTRAL_API_KEY")

	// Explicit key should take precedence
	client := NewMistralClient("explicit-key", "", 0, 0)

	if client.apiKey != "explicit-key" {
		t.Error("Explicit API key should take precedence over environment variable")
	}
}

func TestClientTimeoutValues(t *testing.T) {
	testCases := []struct {
		name     string
		timeout  time.Duration
		expected time.Duration
	}{
		{"Zero timeout uses default", 0, DefaultTimeout},
		{"Custom timeout", 45 * time.Second, 45 * time.Second},
		{"Very short timeout", 1 * time.Second, 1 * time.Second},
		{"Very long timeout", 300 * time.Second, 300 * time.Second},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewMistralClient("test", "", 0, tc.timeout)
			if client.timeout != tc.expected {
				t.Errorf("Expected timeout %v, got %v", tc.expected, client.timeout)
			}
		})
	}
}

func TestClientRetryValues(t *testing.T) {
	testCases := []struct {
		name     string
		retries  int
		expected int
	}{
		{"Zero retries uses default", 0, DefaultMaxRetries},
		{"Custom retries", 3, 3},
		{"Single retry", 1, 1},
		{"Many retries", 20, 20},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewMistralClient("test", "", tc.retries, 0)
			if client.maxRetries != tc.expected {
				t.Errorf("Expected %d retries, got %d", tc.expected, client.maxRetries)
			}
		})
	}
}
