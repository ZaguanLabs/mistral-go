package sdk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ConnectorsGatewayError identifies gateway-authored failures. An upstream
// service's HTTP error remains an ordinary response for the caller to inspect.
type ConnectorsGatewayError struct {
	Response        *http.Response
	Body            []byte
	ProxyStatus     string
	ProxyError      string
	ProxyStatusCode *int
	Details         string
}

func (e *ConnectorsGatewayError) Error() string {
	detail := e.Details
	if detail == "" {
		detail = e.ProxyError
	}
	return fmt.Sprintf("Connectors Gateway request failed with status %d: %s", e.Response.StatusCode, detail)
}

var proxyMemberRE = regexp.MustCompile(`^\s*([A-Za-z*][!#$%&'*+\-.^_` + "`" + `|~0-9A-Za-z:/]*|"(?:[\x20\x21\x23-\x5b\x5d-\x7e]|\\["\\])*")`)
var proxyParamRE = regexp.MustCompile(`^\s*;\s*([a-z*][a-z0-9_\-.*]*)(?:=(?:"((?:[\x20\x21\x23-\x5b\x5d-\x7e]|\\["\\])*)"|([^\s;,"\x00-\x1f]+)))?`)

func gatewayResponseError(response *http.Response) error {
	if response.StatusCode < 400 {
		return nil
	}
	for _, header := range response.Header.Values("Proxy-Status") {
		rest := header
		for {
			m := proxyMemberRE.FindStringSubmatch(rest)
			if m == nil {
				break
			}
			name := m[1]
			raw := strings.TrimLeft(m[0], " \t")
			rest = rest[len(m[0]):]
			params := map[string]string{}
			for {
				p := proxyParamRE.FindStringSubmatch(rest)
				if p == nil {
					break
				}
				value := "?1"
				if strings.Contains(p[0], "=") {
					value = p[3]
					if strings.Contains(p[0], "=\"") {
						value = strings.NewReplacer(`\"`, `"`, `\\`, `\`).Replace(p[2])
					}
				}
				params[p[1]] = value
				raw += p[0]
				rest = rest[len(p[0]):]
			}
			proxyError, hasError := params["error"]
			if name == "connectors-gateway" && hasError {
				body, err := io.ReadAll(response.Body)
				response.Body.Close()
				if err != nil {
					return err
				}
				var code *int
				if s := params["status-code"]; s != "" && strings.Trim(s, "0123456789") == "" {
					if n, err := strconv.Atoi(s); err == nil {
						code = &n
					}
				}
				return &ConnectorsGatewayError{Response: response, Body: body, ProxyStatus: raw, ProxyError: proxyError, ProxyStatusCode: code, Details: params["details"]}
			}
			rest = strings.TrimLeft(rest, " \t")
			if !strings.HasPrefix(rest, ",") {
				break
			}
			rest = rest[1:]
		}
	}
	return nil
}

type connectorGatewayTransport struct {
	base        *url.URL
	client      *MistralClient
	credentials *string
	transport   http.RoundTripper
}

func sameGatewayOrigin(a, b *url.URL) bool {
	port := func(u *url.URL) string {
		if p := u.Port(); p != "" {
			return p
		}
		if u.Scheme == "https" {
			return "443"
		}
		if u.Scheme == "http" {
			return "80"
		}
		return ""
	}
	return strings.EqualFold(a.Scheme, b.Scheme) && strings.EqualFold(a.Hostname(), b.Hostname()) && port(a) == port(b)
}
func allowedGatewayURL(request, base *url.URL) bool {
	if request == nil || request.User != nil || !sameGatewayOrigin(request, base) {
		return false
	}
	prefix := strings.TrimRight(base.EscapedPath(), "/")
	path := request.EscapedPath()
	if path != prefix && !strings.HasPrefix(path, prefix+"/") {
		return false
	}
	relative := strings.TrimPrefix(path, prefix)
	for _, part := range strings.Split(relative, "/") {
		decoded, err := url.PathUnescape(part)
		if err != nil || decoded == "." || decoded == ".." {
			return false
		}
	}
	return true
}
func (t *connectorGatewayTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if !allowedGatewayURL(req.URL, t.base) {
		if req.Body != nil {
			req.Body.Close()
		}
		return nil, fmt.Errorf("connector requests must target the configured gateway connector")
	}
	copy := req.Clone(req.Context())
	if copy.Header.Get("User-Agent") == "" {
		copy.Header.Set("User-Agent", UserAgent)
	}
	if t.credentials != nil {
		copy.Header.Set("X-Credentials-Name", *t.credentials)
	}
	// A caller's Authorization header wins even when its value happens to
	// equal the environment key captured when the SDK was constructed.
	callerAuth := false
	for key := range copy.Header {
		if strings.EqualFold(key, "Authorization") {
			callerAuth = true
		}
	}
	if !callerAuth {
		if err := t.client.authorizeRequest(copy); err != nil {
			if copy.Body != nil {
				copy.Body.Close()
			}
			return nil, err
		}
	}
	resp, err := t.transport.RoundTrip(copy)
	if err != nil {
		return nil, err
	}
	if err = gatewayResponseError(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (c *MistralClient) gatewayHTTPClient(connector, suffix string, credentialsName *string) (*http.Client, *url.URL, error) {
	raw := strings.TrimRight(c.endpoint, "/") + "/v1/connectors-gateway/" + url.PathEscape(connector) + "/" + suffix
	base, err := url.Parse(raw)
	if err != nil {
		return nil, nil, err
	}
	if connector == "" || connector == "." || connector == ".." || base.Host == "" || (base.Scheme != "http" && base.Scheme != "https") {
		return nil, nil, fmt.Errorf("invalid connector or endpoint")
	}
	var credentials *string
	if credentialsName != nil {
		v := *credentialsName
		credentials = &v
	}
	timeout := c.timeout
	if timeout == 0 {
		timeout = 5 * time.Minute
	}
	client := &http.Client{Timeout: timeout, Transport: &connectorGatewayTransport{base: base, client: c, credentials: credentials, transport: http.DefaultTransport}, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	return client, base, nil
}

// ConnectorHTTPClient scopes arbitrary HTTP calls to one configured connector.
type ConnectorHTTPClient struct {
	HTTPClient *http.Client
	base       *url.URL
}

func (c *MistralClient) ConnectorHTTPClient(connector string, credentialsName *string) (*ConnectorHTTPClient, error) {
	client, base, err := c.gatewayHTTPClient(connector, "http", credentialsName)
	if err != nil {
		return nil, err
	}
	return &ConnectorHTTPClient{client, base}, nil
}
func (c *ConnectorHTTPClient) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	parsed, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	if !parsed.IsAbs() && parsed.Host == "" {
		path = strings.TrimRight(c.base.String(), "/") + "/" + strings.TrimLeft(path, "/")
	}
	request, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}
	if !allowedGatewayURL(request.URL, c.base) {
		return nil, fmt.Errorf("connector requests must target the configured gateway connector")
	}
	return request, nil
}
func (c *ConnectorHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}
func (c *ConnectorHTTPClient) Close() { c.HTTPClient.CloseIdleConnections() }

// ConnectorMCPClient initializes an official MCP Go client session through the
// gateway. Call Close on the returned session when finished.
func (c *MistralClient) ConnectorMCPClient(ctx context.Context, connector string, credentialsName *string) (*mcp.ClientSession, error) {
	client, base, err := c.gatewayHTTPClient(connector, "mcp", credentialsName)
	if err != nil {
		return nil, err
	}
	session, err := mcp.NewClient(&mcp.Implementation{Name: SDKName, Version: Version}, nil).Connect(ctx, &mcp.StreamableClientTransport{Endpoint: base.String(), HTTPClient: client}, nil)
	if err != nil {
		var gateway *ConnectorsGatewayError
		if errors.As(err, &gateway) {
			return nil, gateway
		}
		return nil, err
	}
	return session, nil
}

func (t *connectorGatewayTransport) CloseIdleConnections() {
	if closer, ok := t.transport.(interface{ CloseIdleConnections() }); ok {
		closer.CloseIdleConnections()
	}
}
