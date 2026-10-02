package sdk

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestConnectorGatewayHTTPBoundary(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		if r.Header.Get("Authorization") != "Bearer key" || r.Header.Get("X-Credentials-Name") != "cred" {
			t.Errorf("headers %v", r.Header)
		}
		if strings.HasSuffix(r.URL.Path, "/redirect") {
			w.Header().Set("Location", "https://example.invalid/")
			w.WriteHeader(302)
			return
		}
		w.Write([]byte("ok"))
	}))
	defer server.Close()
	client, err := NewMistralClient("key", server.URL, 1, time.Second).ConnectorHTTPClient("a/b", StringPtr("cred"))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"/item?x=1", "item", "redirect"} {
		req, err := client.NewRequest(context.Background(), "GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		response, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		response.Body.Close()
		if path == "redirect" && response.StatusCode != 302 {
			t.Fatal("redirect followed")
		}
	}
	before := calls.Load()
	for _, path := range []string{"../other", "%2e%2e/other", "./item", "https://example.invalid/", "/../x", server.URL + "/v1/connectors-gateway/other/http/x"} {
		req, err := client.NewRequest(context.Background(), "GET", path, nil)
		if err == nil {
			_, err = client.Do(req)
		}
		if err == nil {
			t.Errorf("escaped connector: %s", path)
		}
	}
	// Direct requests through the exposed client retain the same boundary.
	req, _ := http.NewRequest("GET", server.URL+"/elsewhere", nil)
	if _, err = client.Do(req); err == nil {
		t.Fatal("direct request escaped")
	}
	if calls.Load() != before {
		t.Fatal("rejected request reached network")
	}
}

func TestConnectorGatewayProxyStatusErrors(t *testing.T) {
	cases := []struct {
		header  string
		gateway bool
		detail  string
	}{
		{`connectors-gateway; error=connection_refused; status-code=502; details="upstream down"`, true, "upstream down"},
		{`cache; error=timeout, connectors-gateway; error=denied; details="a, b; c\"d"`, true, `a, b; c"d`},
		{`origin; error=timeout`, false, ""},
		{`"connectors-gateway"; error=denied`, false, ""},
		{`malformed ???, connectors-gateway; error=denied`, false, ""},
	}
	for _, tc := range cases {
		t.Run(tc.header, func(t *testing.T) {
			resp := &http.Response{StatusCode: 502, Header: http.Header{"Proxy-Status": []string{tc.header}}, Body: io.NopCloser(strings.NewReader("body"))}
			err := gatewayResponseError(resp)
			var gateway *ConnectorsGatewayError
			if errors.As(err, &gateway) != tc.gateway {
				t.Fatalf("error %v", err)
			}
			if tc.gateway && (gateway.Details != tc.detail || string(gateway.Body) != "body") {
				t.Fatalf("%+v", gateway)
			}
		})
	}
}

func TestConnectorGatewayMCPSession(t *testing.T) {
	backend := mcp.NewServer(&mcp.Implementation{Name: "fixture", Version: "1"}, nil)
	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server { return backend }, nil)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/connectors-gateway/conn/mcp" || r.Header.Get("Authorization") != "Bearer key" || r.Header.Get("X-Credentials-Name") != "cred" {
			t.Errorf("request %s %v", r.URL.Path, r.Header)
		}
		handler.ServeHTTP(w, r)
	}))
	defer server.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	session, err := NewMistralClient("key", server.URL, 1, time.Second).ConnectorMCPClient(ctx, "conn", StringPtr("cred"))
	if err != nil {
		t.Fatal(err)
	}
	if err := session.Ping(ctx, nil); err != nil {
		t.Fatal(err)
	}
	if err := session.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestConnectorGatewayExplicitAuthorization(t *testing.T) {
	t.Setenv("MISTRAL_API_KEY", "environment")
	t.Setenv(ServiceAccountTokenPathEnv, t.TempDir()+"/missing-token")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer environment" {
			t.Errorf("caller authorization was replaced: %q", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	client, err := NewMistralClient("", server.URL, 1, time.Second).ConnectorHTTPClient("conn", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	req, err := client.NewRequest(context.Background(), http.MethodGet, "item", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer environment")
	response, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	response.Body.Close()
}
