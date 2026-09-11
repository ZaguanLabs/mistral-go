package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type OAuth2GrantType string

const (
	OAuth2GrantAuthorizationCode OAuth2GrantType = "authorization_code"
	OAuth2GrantClientCredentials OAuth2GrantType = "client_credentials"
)

type OAuth2TokenEndpointAuthMethod string

const (
	OAuth2TokenEndpointAuthClientSecretPost  OAuth2TokenEndpointAuthMethod = "client_secret_post"
	OAuth2TokenEndpointAuthClientSecretBasic OAuth2TokenEndpointAuthMethod = "client_secret_basic"
)

type OAuth2ClientCredentialsInput struct {
	ClientID     string          `json:"client_id"`
	ClientSecret string          `json:"client_secret"`
	GrantType    OAuth2GrantType `json:"grant_type"`
}

func (v OAuth2ClientCredentialsInput) MarshalJSON() ([]byte, error) {
	type wire OAuth2ClientCredentialsInput
	if v.GrantType == "" {
		v.GrantType = OAuth2GrantClientCredentials
	}
	return json.Marshal(wire(v))
}

type ConnectionCredentialsInput struct {
	OAuth       *OAuth2ClientCredentialsInput `json:"oauth,omitempty"`
	Headers     map[string]string             `json:"headers,omitempty"`
	BearerToken *string                       `json:"bearer_token,omitempty"`
}

// ConnectorAuthMethod represents bearer, none, and both OAuth2 grant variants.
type ConnectorAuthMethod struct {
	MethodType              ConnectorAuthenticationMethodType `json:"method_type"`
	Headers                 []ConnectorAuthenticationHeader   `json:"headers,omitempty"`
	GrantType               OAuth2GrantType                   `json:"grant_type,omitempty"`
	OAuth2ServerMetadata    *ExtendedOAuthServerMetadata      `json:"oauth2_server_metadata,omitempty"`
	AuthData                map[string]any                    `json:"auth_data,omitempty"`
	TokenEndpointAuthMethod OAuth2TokenEndpointAuthMethod     `json:"token_endpoint_auth_method,omitempty"`
}
type BearerAuthMethod ConnectorAuthMethod
type NoneAuthMethod ConnectorAuthMethod
type OAuth2AuthorizationCodeAuthMethod ConnectorAuthMethod
type OAuth2ClientCredentialsAuthMethod ConnectorAuthMethod

func (v BearerAuthMethod) MarshalJSON() ([]byte, error) {
	w := ConnectorAuthMethod(v)
	w.MethodType = ConnectorAuthenticationBearer
	return json.Marshal(w)
}
func (v NoneAuthMethod) MarshalJSON() ([]byte, error) {
	w := ConnectorAuthMethod(v)
	w.MethodType = ConnectorAuthenticationNone
	return json.Marshal(w)
}
func (v OAuth2AuthorizationCodeAuthMethod) MarshalJSON() ([]byte, error) {
	w := ConnectorAuthMethod(v)
	w.MethodType = ConnectorAuthenticationOAuth2
	w.GrantType = OAuth2GrantAuthorizationCode
	return json.Marshal(w)
}
func (v OAuth2ClientCredentialsAuthMethod) MarshalJSON() ([]byte, error) {
	w := ConnectorAuthMethod(v)
	w.MethodType = ConnectorAuthenticationOAuth2
	w.GrantType = OAuth2GrantClientCredentials
	return json.Marshal(w)
}

type CreateConnectorRequest = ConnectorRequest
type CreateHTTPConnectorRequest ConnectorRequest
type ConnectorMCPPublicUpdate UpdateConnectorRequest
type UpdateHTTPConnectorRequest UpdateConnectorRequest

func (v CreateHTTPConnectorRequest) MarshalJSON() ([]byte, error) {
	w := ConnectorRequest(v)
	protocol := "http"
	w.Protocol = &protocol
	return json.Marshal(w)
}
func (v ConnectorMCPPublicUpdate) MarshalJSON() ([]byte, error) {
	w := UpdateConnectorRequest(v)
	protocol := "mcp"
	w.Protocol = &protocol
	return json.Marshal(w)
}
func (v UpdateHTTPConnectorRequest) MarshalJSON() ([]byte, error) {
	w := UpdateConnectorRequest(v)
	protocol := "http"
	w.Protocol = &protocol
	return json.Marshal(w)
}

// Connector retains both HTTP and MCP variant fields; Protocol is the discriminator.
// Unknown variants remain available through Raw.
type Connector struct {
	ID                    string                       `json:"id"`
	Name                  string                       `json:"name"`
	Description           string                       `json:"description"`
	CreatedAt             time.Time                    `json:"created_at"`
	ModifiedAt            time.Time                    `json:"modified_at"`
	OwnerType             string                       `json:"owner_type"`
	Visibility            string                       `json:"visibility"`
	PrivateToolExecution  bool                         `json:"private_tool_execution"`
	Protocol              string                       `json:"protocol"`
	Title                 *string                      `json:"title,omitempty"`
	Server                *string                      `json:"server,omitempty"`
	IconURL               *string                      `json:"icon_url,omitempty"`
	ServerCard            map[string]any               `json:"server_card,omitempty"`
	OwnerID               *string                      `json:"owner_id,omitempty"`
	CreatorID             *string                      `json:"creator_id,omitempty"`
	Locale                *ConnectorLocale             `json:"locale,omitempty"`
	SystemPrompt          *string                      `json:"system_prompt,omitempty"`
	SupportedAuthMethods  []PublicAuthenticationMethod `json:"supported_auth_methods,omitempty"`
	ConnectionPreferences []map[string]any             `json:"connection_preferences,omitempty"`
	ConnectionCredentials []map[string]any             `json:"connection_credentials,omitempty"`
	Active                *bool                        `json:"active,omitempty"`
	BlockedBy             []string                     `json:"blocked_by,omitempty"`
	Mistral               bool                         `json:"mistral,omitempty"`
	IsAuthenticated       *bool                        `json:"is_authenticated,omitempty"`
	Tools                 []map[string]any             `json:"tools,omitempty"`
	SystemPromptRoute     *string                      `json:"system_prompt_route,omitempty"`
	ConnectionConfig      map[string]any               `json:"connection_config,omitempty"`
	ExecutionEnv          map[string]any               `json:"execution_env,omitempty"`
	Raw                   json.RawMessage              `json:"-"`
}

func (v *Connector) UnmarshalJSON(data []byte) error {
	type wire Connector
	var w wire
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	*v = Connector(w)
	v.Raw = append(json.RawMessage(nil), data...)
	return nil
}

type HTTPConnector Connector
type MCPConnector Connector

func (v HTTPConnector) MarshalJSON() ([]byte, error) {
	w := Connector(v)
	w.Protocol = "http"
	return json.Marshal(w)
}
func (v MCPConnector) MarshalJSON() ([]byte, error) {
	w := Connector(v)
	w.Protocol = "mcp"
	return json.Marshal(w)
}

func (c *MistralClient) ShareConnectorToOrganization(id string) (APIResponse, error) {
	return c.requestMap(http.MethodPut, nil, "v1/connectors/"+url.PathEscape(id)+"/organization/share")
}
func (c *MistralClient) UnshareConnectorFromOrganization(id string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, "v1/connectors/"+url.PathEscape(id)+"/organization/share")
}
func (c *MistralClient) CreateConnectorCredentials(id, scope string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.writeConnectorCredentials(http.MethodPost, id, scope, req)
}
func (c *MistralClient) UpdateConnectorCredentials(id, scope string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.writeConnectorCredentials(http.MethodPatch, id, scope, req)
}
func (c *MistralClient) writeConnectorCredentials(method, id, scope string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if scope != "organization" && scope != "workspace" && scope != "user" {
		return nil, fmt.Errorf("invalid consumer scope %q", scope)
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(method, body, "v1/connectors/"+url.PathEscape(id)+"/"+scope+"/credentials")
}

type ConnectorLocale struct {
	Name          map[string]string `json:"name"`
	Description   map[string]string `json:"description"`
	UsageSentence map[string]string `json:"usage_sentence"`
}

func (v Connector) MarshalJSON() ([]byte, error) {
	if v.Protocol != "mcp" && v.Protocol != "http" && len(v.Raw) > 0 {
		return append([]byte(nil), v.Raw...), nil
	}
	type wire Connector
	return json.Marshal(wire(v))
}
