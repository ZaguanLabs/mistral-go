package sdk

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// BackoffStrategy uses durations rather than the Python SDK's millisecond integers.
type BackoffStrategy struct {
	InitialInterval time.Duration
	MaxInterval     time.Duration
	Exponent        float64
	MaxElapsedTime  time.Duration
	// Nil uses the Python default of up to one second of additive jitter.
	Jitter *time.Duration
}
type RetryConfig struct {
	Strategy              string // "none" or "backoff"
	Backoff               BackoffStrategy
	RetryConnectionErrors bool
	StatusCodesOverride   []string // Exact codes or classes such as "5XX".
}

// WithRetryConfig returns an independent client configuration.
func (c *MistralClient) WithRetryConfig(config RetryConfig) (*MistralClient, error) {
	if config.Strategy != "none" && config.Strategy != "backoff" {
		return nil, fmt.Errorf("retry strategy must be none or backoff")
	}
	if config.Backoff.Jitter != nil && *config.Backoff.Jitter < 0 {
		return nil, fmt.Errorf("jitter must be nonnegative")
	}
	if config.Backoff.InitialInterval < 0 || config.Backoff.MaxInterval < 0 || config.Backoff.MaxElapsedTime < 0 || config.Backoff.Exponent < 0 || math.IsNaN(config.Backoff.Exponent) || math.IsInf(config.Backoff.Exponent, 0) {
		return nil, fmt.Errorf("invalid backoff")
	}
	config.StatusCodesOverride = append([]string(nil), config.StatusCodesOverride...)
	for _, code := range config.StatusCodesOverride {
		if len(code) != 3 {
			return nil, fmt.Errorf("invalid retry status %q", code)
		}
		if code[1:] == "XX" && code[0] >= '1' && code[0] <= '5' {
			continue
		}
		n, err := strconv.Atoi(code)
		if err != nil || n < 100 || n > 599 {
			return nil, fmt.Errorf("invalid retry status %q", code)
		}
	}
	if config.Backoff.Jitter != nil {
		jitter := *config.Backoff.Jitter
		config.Backoff.Jitter = &jitter
	}
	copy := *c
	copy.retryConfig = &config
	return &copy, nil
}

// WithContext returns a client whose requests and streams stop when ctx is cancelled.
func (c *MistralClient) WithContext(ctx context.Context) *MistralClient {
	copy := *c
	copy.ctx = ctx
	return &copy
}
func (c *MistralClient) requestContext() context.Context {
	if c.ctx != nil {
		return c.ctx
	}
	return context.Background()
}
func retryAfter(header http.Header, now time.Time) (time.Duration, bool) {
	for _, key := range []string{"retry-after-ms", "Retry-After"} {
		value := header.Get(key)
		if value == "" {
			continue
		}
		if n, err := strconv.ParseFloat(value, 64); err == nil && n >= 0 && !math.IsInf(n, 0) && !math.IsNaN(n) {
			scale := float64(time.Second)
			if key == "retry-after-ms" {
				scale = float64(time.Millisecond)
			}
			if n*scale < float64(math.MaxInt64) {
				return time.Duration(n * scale), true
			}
		}
		if key == "Retry-After" {
			if date, err := http.ParseTime(value); err == nil {
				d := date.Sub(now)
				if d < 0 {
					d = 0
				}
				return d, true
			}
		}
	}
	return 0, false
}
func (c *MistralClient) shouldRetryStatus(code int) bool {
	if c.retryConfig == nil || len(c.retryConfig.StatusCodesOverride) == 0 {
		return retryStatusCodes[code]
	}
	for _, rule := range c.retryConfig.StatusCodesOverride {
		if rule == strconv.Itoa(code) || rule == fmt.Sprintf("%dXX", code/100) {
			return true
		}
	}
	return false
}
func (c *MistralClient) retryDelay(attempt int, header http.Header) time.Duration {
	if delay, ok := retryAfter(header, time.Now()); ok {
		return delay
	}
	if c.retryConfig == nil {
		return time.Duration(attempt+1) * 500 * time.Millisecond
	}
	b := c.retryConfig.Backoff
	initial, max, exponent := b.InitialInterval, b.MaxInterval, b.Exponent
	if max == 0 {
		max = 60 * time.Second
	}
	if exponent == 0 {
		exponent = 1.5
	}
	delay := float64(initial) * math.Pow(exponent, float64(attempt))
	jitter := time.Second
	if b.Jitter != nil {
		jitter = *b.Jitter
	}
	delay += rand.Float64() * float64(jitter)
	if delay > float64(max) {
		delay = float64(max)
	}
	return time.Duration(delay)
}
func (c *MistralClient) doRequest(req *http.Request) (*http.Response, error) {
	if c.ctx != nil {
		req = req.Clone(c.ctx)
	}
	attempts := c.maxRetries
	if attempts < 1 {
		attempts = 1
	}
	if c.retryConfig != nil && c.retryConfig.Strategy == "none" {
		attempts = 1
	}
	start := time.Now()
	client := &http.Client{Timeout: c.timeout}
	for attempt := 0; attempt < attempts; attempt++ {
		if attempt > 0 && req.Body != nil {
			if req.GetBody == nil {
				return nil, fmt.Errorf("request body cannot be replayed")
			}
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req = req.Clone(req.Context())
			req.Body = body
		}
		resp, err := client.Do(req)
		if req.Context().Err() != nil {
			if resp != nil {
				resp.Body.Close()
			}
			return nil, req.Context().Err()
		}
		retry := err == nil && c.shouldRetryStatus(resp.StatusCode)
		if err != nil {
			var networkError net.Error
			retry = errors.As(err, &networkError) && (c.retryConfig == nil || c.retryConfig.RetryConnectionErrors)
		}
		if !retry || attempt == attempts-1 {
			return resp, err
		}
		if req.Body != nil && req.GetBody == nil {
			return resp, err
		}
		header := http.Header{}
		if resp != nil {
			header = resp.Header
		}
		delay := c.retryDelay(attempt, header)
		if c.retryConfig != nil && c.retryConfig.Backoff.MaxElapsedTime > 0 && time.Since(start)+delay > c.retryConfig.Backoff.MaxElapsedTime {
			return resp, err
		}
		if resp != nil {
			resp.Body.Close()
		}
		if err := batchDelay(req.Context(), delay); err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("%s request retries exhausted", strings.ToLower(req.Method))
}
