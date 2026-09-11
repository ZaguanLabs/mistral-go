package sdk

import (
	"fmt"
	"net/http"
)

type ConnectorRequest struct {
	AuthMethods             any                          `json:"auth_methods,omitempty"`
	Name                    string                       `json:"name,omitempty"`
	Description             *string                      `json:"description,omitempty"`
	Server                  any                          `json:"server,omitempty"`
	Protocol                *string                      `json:"protocol,omitempty"`
	Title                   *string                      `json:"title,omitempty"`
	IconURL                 *string                      `json:"icon_url,omitempty"`
	Visibility              any                          `json:"visibility,omitempty"`
	Headers                 map[string]any               `json:"headers,omitempty"`
	GlobalHeaders           map[string]GlobalHeaderValue `json:"global_headers,omitempty"`
	AuthData                map[string]any               `json:"auth_data,omitempty"`
	OAuth2ServerMetadata    map[string]any               `json:"oauth2_server_metadata,omitempty"`
	OAuth2ServerMetadataURL *string                      `json:"oauth2_server_metadata_url,omitempty"`
	SystemPrompt            *string                      `json:"system_prompt,omitempty"`
}

type UpdateConnectorRequest struct {
	Title             *string        `json:"title,omitempty"`
	Name              *string        `json:"name,omitempty"`
	Description       *string        `json:"description,omitempty"`
	IconURL           *string        `json:"icon_url,omitempty"`
	SystemPrompt      *string        `json:"system_prompt,omitempty"`
	ConnectionConfig  map[string]any `json:"connection_config,omitempty"`
	ConnectionSecrets map[string]any `json:"connection_secrets,omitempty"`
	Protocol          *string        `json:"protocol,omitempty"`
	Server            any            `json:"server,omitempty"`
	Headers           map[string]any `json:"headers,omitempty"`
	AuthData          map[string]any `json:"auth_data,omitempty"`
	AuthMethods       any            `json:"auth_methods,omitempty"`
}

type ConnectorAuthenticationMethodType string

const (
	ConnectorAuthenticationOAuth2    ConnectorAuthenticationMethodType = "oauth2"
	ConnectorAuthenticationBearer    ConnectorAuthenticationMethodType = "bearer"
	ConnectorAuthenticationNone      ConnectorAuthenticationMethodType = "none"
	ConnectorAuthenticationGitHubApp ConnectorAuthenticationMethodType = "github_app"
	ConnectorAuthenticationSlackApp  ConnectorAuthenticationMethodType = "slack_app"
	ConnectorAuthenticationWebhook   ConnectorAuthenticationMethodType = "webhook"
)

type AuthDirection string

const (
	AuthDirectionInbound  AuthDirection = "inbound"
	AuthDirectionOutbound AuthDirection = "outbound"
)

type OAuthMetadataSource string

const (
	OAuthMetadataSourceAutodiscovery OAuthMetadataSource = "autodiscovery"
	OAuthMetadataSourceProvided      OAuthMetadataSource = "provided"
)

type GlobalHeaderValue struct {
	Value    string `json:"value"`
	IsSecret *bool  `json:"is_secret,omitempty"`
}

type ConnectorAuthenticationHeader struct {
	Name       string `json:"name"`
	IsRequired *bool  `json:"is_required,omitempty"`
	IsSecret   *bool  `json:"is_secret,omitempty"`
}

type OAuth2MetadataSecrets struct {
	ClientID              *string `json:"client_id,omitempty"`
	ClientSecret          *string `json:"client_secret,omitempty"`
	ClientIDIssuedAt      *int    `json:"client_id_issued_at,omitempty"`
	ClientSecretExpiresAt *int    `json:"client_secret_expires_at,omitempty"`
}

type ExtendedOAuthServerMetadata struct {
	Issuer                                             string               `json:"issuer"`
	AuthorizationEndpoint                              *string              `json:"authorization_endpoint,omitempty"`
	TokenEndpoint                                      string               `json:"token_endpoint"`
	RegistrationEndpoint                               *string              `json:"registration_endpoint,omitempty"`
	ScopesSupported                                    []string             `json:"scopes_supported,omitempty"`
	ResponseTypesSupported                             []string             `json:"response_types_supported,omitempty"`
	ResponseModesSupported                             []string             `json:"response_modes_supported,omitempty"`
	GrantTypesSupported                                []string             `json:"grant_types_supported,omitempty"`
	TokenEndpointAuthMethodsSupported                  []string             `json:"token_endpoint_auth_methods_supported,omitempty"`
	TokenEndpointAuthSigningAlgValuesSupported         []string             `json:"token_endpoint_auth_signing_alg_values_supported,omitempty"`
	ServiceDocumentation                               *string              `json:"service_documentation,omitempty"`
	UILocalesSupported                                 []string             `json:"ui_locales_supported,omitempty"`
	OPPolicyURI                                        *string              `json:"op_policy_uri,omitempty"`
	OPTOSURI                                           *string              `json:"op_tos_uri,omitempty"`
	RevocationEndpoint                                 *string              `json:"revocation_endpoint,omitempty"`
	RevocationEndpointAuthMethodsSupported             []string             `json:"revocation_endpoint_auth_methods_supported,omitempty"`
	RevocationEndpointAuthSigningAlgValuesSupported    []string             `json:"revocation_endpoint_auth_signing_alg_values_supported,omitempty"`
	IntrospectionEndpoint                              *string              `json:"introspection_endpoint,omitempty"`
	IntrospectionEndpointAuthMethodsSupported          []string             `json:"introspection_endpoint_auth_methods_supported,omitempty"`
	IntrospectionEndpointAuthSigningAlgValuesSupported []string             `json:"introspection_endpoint_auth_signing_alg_values_supported,omitempty"`
	CodeChallengeMethodsSupported                      []string             `json:"code_challenge_methods_supported,omitempty"`
	ClientIDMetadataDocumentSupported                  *bool                `json:"client_id_metadata_document_supported,omitempty"`
	XSource                                            *OAuthMetadataSource `json:"x_source,omitempty"`
	XResourceURL                                       *string              `json:"x_resource_url,omitempty"`
	XScope                                             *string              `json:"x_scope,omitempty"`
}

type AuthenticationMethodCreateOrUpdateRequest struct {
	MethodType            ConnectorAuthenticationMethodType `json:"method_type"`
	AuthDirection         *AuthDirection                    `json:"auth_direction,omitempty"`
	Headers               []ConnectorAuthenticationHeader   `json:"headers,omitempty"`
	GlobalHeaders         map[string]GlobalHeaderValue      `json:"global_headers,omitempty"`
	OAuth2MetadataSecrets *OAuth2MetadataSecrets            `json:"oauth2_metadata_secrets,omitempty"`
	OAuth2ServerMetadata  *ExtendedOAuthServerMetadata      `json:"oauth2_server_metadata,omitempty"`
}

type PublicAuthenticationMethod struct {
	GrantType             *OAuth2GrantType                  `json:"grant_type,omitempty"`
	MethodType            ConnectorAuthenticationMethodType `json:"method_type"`
	HasDefaultCredentials bool                              `json:"has_default_credentials"`
	Headers               []ConnectorAuthenticationHeader   `json:"headers,omitempty"`
	GlobalHeaders         map[string]GlobalHeaderValue      `json:"global_headers,omitempty"`
	OAuth2ServerMetadata  *ExtendedOAuthServerMetadata      `json:"oauth2_server_metadata,omitempty"`
}

type ListConnectorsParams struct {
	QueryFilters map[string]any `json:"query_filters,omitempty"`
	Cursor       *string        `json:"cursor,omitempty"`
	PageSize     *int           `json:"page_size,omitempty"`
}

type ConnectorCredentialsRequest struct {
	Name        string  `json:"name,omitempty"`
	Title       *string `json:"title,omitempty"`
	IsDefault   *bool   `json:"is_default,omitempty"`
	Credentials any     `json:"credentials,omitempty"`
}

type ConnectorAuthURLParams struct {
	AppReturnURL           *string
	MethodType             *string
	CredentialsName        *string
	CredentialsTitle       *string
	GitHubInstallationLink *bool
}

type ListConnectorCredentialsParams struct {
	AuthType     *string `json:"auth_type,omitempty"`
	FetchDefault *bool   `json:"fetch_default,omitempty"`
}

type ListConnectorToolsParams struct {
	Page            *int    `json:"page,omitempty"`
	PageSize        *int    `json:"page_size,omitempty"`
	Refresh         *bool   `json:"refresh,omitempty"`
	Pretty          *bool   `json:"pretty,omitempty"`
	CredentialsName *string `json:"credentials_name,omitempty"`
}

type ToolExecutionConfiguration struct {
	RequiresConfirmation any      `json:"requires_confirmation,omitempty"`
	SkipConfirmation     any      `json:"skip_confirmation,omitempty"`
	Include              []string `json:"include,omitempty"`
	Exclude              []string `json:"exclude,omitempty"`
}

// CreateConnector accepts an MCP or HTTP request, or an equivalent JSON object.
func (c *MistralClient) CreateConnector(req any) (APIResponse, error) {
	body, err := connectorRequestBody(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPost, body, "v1/connectors")
}

func connectorRequestBody(req any) (map[string]interface{}, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body, err := structToMap(req)
	if err != nil {
		return nil, err
	}
	if body == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if _, ok := body["protocol"]; !ok {
		body["protocol"] = "mcp"
	}
	if body["protocol"] != "mcp" && body["protocol"] != "http" {
		return nil, fmt.Errorf("connector protocol must be mcp or http")
	}
	return body, nil
}

func (c *MistralClient) ListConnectors(params *ListConnectorsParams) (APIResponse, error) {
	if params == nil {
		params = &ListConnectorsParams{}
	}
	query := queryWithOptionalValues(map[string]any{"cursor": params.Cursor, "page_size": params.PageSize})
	body := optionalRequestMap(map[string]any{"query_filters": params.QueryFilters})
	if len(body) == 0 {
		return c.requestMap(http.MethodGet, nil, appendQuery("v1/connectors", query))
	}
	return c.requestMap(http.MethodGet, body, appendQuery("v1/connectors", query))
}

func (c *MistralClient) GetConnectorAuthURL(connectorIDOrName string, appReturnURL *string, credentialsName *string, githubInstallationLink ...*bool) (APIResponse, error) {
	var githubLink *bool
	if len(githubInstallationLink) > 0 {
		githubLink = githubInstallationLink[0]
	}
	query := queryWithOptionalValues(map[string]any{
		"app_return_url":           appReturnURL,
		"credentials_name":         credentialsName,
		"github_installation_link": githubLink,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/connectors/%s/auth_url", connectorIDOrName), query))
}

func (c *MistralClient) GetConnectorAuthURLWithParams(connectorIDOrName string, params *ConnectorAuthURLParams) (APIResponse, error) {
	if params == nil {
		params = &ConnectorAuthURLParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"app_return_url":           params.AppReturnURL,
		"method_type":              params.MethodType,
		"credentials_name":         params.CredentialsName,
		"credentials_title":        params.CredentialsTitle,
		"github_installation_link": params.GitHubInstallationLink,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/connectors/%s/auth_url", connectorIDOrName), query))
}

func (c *MistralClient) CallConnectorTool(connectorIDOrName, toolName string, credentialsName *string, arguments map[string]any) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"credentials_name": credentialsName})
	body := optionalRequestMap(map[string]any{"arguments": arguments})
	return c.requestMap(http.MethodPost, body, appendQuery(fmt.Sprintf("v1/connectors/%s/tools/%s/call", connectorIDOrName, toolName), query))
}

func (c *MistralClient) ListConnectorTools(connectorIDOrName string, params *ListConnectorToolsParams) (APIResponse, error) {
	if params == nil {
		params = &ListConnectorToolsParams{}
	}
	query := queryWithOptionalValues(map[string]any{
		"page":             params.Page,
		"page_size":        params.PageSize,
		"refresh":          params.Refresh,
		"pretty":           params.Pretty,
		"credentials_name": params.CredentialsName,
	})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/connectors/%s/tools", connectorIDOrName), query))
}

func (c *MistralClient) GetConnectorAuthenticationMethods(connectorIDOrName string) (APIResponse, error) {
	return c.requestMap(http.MethodGet, nil, fmt.Sprintf("v1/connectors/%s/authentication_methods", connectorIDOrName))
}

func (c *MistralClient) ActivateConnectorForOrganization(connectorID string, config *ToolExecutionConfiguration) (APIResponse, error) {
	return c.activateConnector(connectorID, "organization", config)
}

func (c *MistralClient) DeactivateConnectorForOrganization(connectorID string) (APIResponse, error) {
	return c.deactivateConnector(connectorID, "organization")
}

func (c *MistralClient) ActivateConnectorForWorkspace(connectorID string, config *ToolExecutionConfiguration) (APIResponse, error) {
	return c.activateConnector(connectorID, "workspace", config)
}

func (c *MistralClient) DeactivateConnectorForWorkspace(connectorID string) (APIResponse, error) {
	return c.deactivateConnector(connectorID, "workspace")
}

func (c *MistralClient) ActivateConnectorForUser(connectorID string, config *ToolExecutionConfiguration) (APIResponse, error) {
	return c.activateConnector(connectorID, "user", config)
}

func (c *MistralClient) DeactivateConnectorForUser(connectorID string) (APIResponse, error) {
	return c.deactivateConnector(connectorID, "user")
}

func (c *MistralClient) ActivateConnectorForConsumer(connectorID, consumerScope string, config *ToolExecutionConfiguration) (APIResponse, error) {
	return c.activateConnector(connectorID, consumerScope, config)
}

func (c *MistralClient) DeactivateConnectorForConsumer(connectorID, consumerScope string) (APIResponse, error) {
	return c.deactivateConnector(connectorID, consumerScope)
}

func (c *MistralClient) ShareConnector(connectorID string, requestBody map[string]any) (APIResponse, error) {
	return c.requestMap(http.MethodPut, requestBody, fmt.Sprintf("v1/connectors/%s/share", connectorID))
}

func (c *MistralClient) UnshareConnector(connectorID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/connectors/%s/share", connectorID))
}

func (c *MistralClient) DeleteAllUserConnectorCredentials(connectorIDOrName string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/connectors/%s/user/credentials", connectorIDOrName))
}

func (c *MistralClient) ListOrganizationConnectorCredentials(connectorIDOrName string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	return c.listConnectorCredentials(connectorIDOrName, "organization", params)
}

// Deprecated: removed from the official Python SDK in v2.10.0.
func (c *MistralClient) CreateOrUpdateOrganizationConnectorCredentials(connectorIDOrName string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.createOrUpdateConnectorCredentials(connectorIDOrName, "organization", req)
}

func (c *MistralClient) ListWorkspaceConnectorCredentials(connectorIDOrName string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	return c.listConnectorCredentials(connectorIDOrName, "workspace", params)
}

// Deprecated: removed from the official Python SDK in v2.10.0.
func (c *MistralClient) CreateOrUpdateWorkspaceConnectorCredentials(connectorIDOrName string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.createOrUpdateConnectorCredentials(connectorIDOrName, "workspace", req)
}

func (c *MistralClient) ListUserConnectorCredentials(connectorIDOrName string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	return c.listConnectorCredentials(connectorIDOrName, "user", params)
}

// Deprecated: removed from the official Python SDK in v2.10.0.
func (c *MistralClient) CreateOrUpdateUserConnectorCredentials(connectorIDOrName string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.createOrUpdateConnectorCredentials(connectorIDOrName, "user", req)
}

func (c *MistralClient) DeleteOrganizationConnectorCredentials(connectorIDOrName, credentialsName string) (APIResponse, error) {
	return c.deleteConnectorCredentials(connectorIDOrName, "organization", credentialsName)
}

func (c *MistralClient) DeleteWorkspaceConnectorCredentials(connectorIDOrName, credentialsName string) (APIResponse, error) {
	return c.deleteConnectorCredentials(connectorIDOrName, "workspace", credentialsName)
}

func (c *MistralClient) DeleteUserConnectorCredentials(connectorIDOrName, credentialsName string) (APIResponse, error) {
	return c.deleteConnectorCredentials(connectorIDOrName, "user", credentialsName)
}

func (c *MistralClient) GetConnector(connectorIDOrName string, fetchCustomerData *bool, fetchConnectionSecrets *bool) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"fetch_customer_data": fetchCustomerData, "fetch_connection_secrets": fetchConnectionSecrets})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/connectors/%s", connectorIDOrName), query))
}

func (c *MistralClient) UpdateConnector(connectorID string, req any) (APIResponse, error) {
	body, err := connectorRequestBody(req)
	if err != nil {
		return nil, err
	}
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v1/connectors/%s", connectorID))
}

func (c *MistralClient) DeleteConnector(connectorID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/connectors/%s", connectorID))
}

func (c *MistralClient) activateConnector(connectorID, scope string, config *ToolExecutionConfiguration) (APIResponse, error) {
	var body map[string]interface{}
	if config != nil {
		body = optionalRequestMap(map[string]any{
			"requires_confirmation": config.RequiresConfirmation,
			"skip_confirmation":     config.SkipConfirmation,
			"include":               config.Include,
			"exclude":               config.Exclude,
		})
	}
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/connectors/%s/%s/activate", connectorID, scope))
}

func (c *MistralClient) deactivateConnector(connectorID, scope string) (APIResponse, error) {
	return c.requestMap(http.MethodPost, nil, fmt.Sprintf("v1/connectors/%s/%s/deactivate", connectorID, scope))
}

func (c *MistralClient) listConnectorCredentials(connectorIDOrName, scope string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	if params == nil {
		params = &ListConnectorCredentialsParams{}
	}
	query := queryWithOptionalValues(map[string]any{"auth_type": params.AuthType, "fetch_default": params.FetchDefault})
	return c.requestMap(http.MethodGet, nil, appendQuery(fmt.Sprintf("v1/connectors/%s/%s/credentials", connectorIDOrName, scope), query))
}

func (c *MistralClient) createOrUpdateConnectorCredentials(connectorIDOrName, scope string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{"name": req.Name, "title": req.Title, "is_default": req.IsDefault, "credentials": req.Credentials})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/connectors/%s/%s/credentials", connectorIDOrName, scope))
}

func (c *MistralClient) deleteConnectorCredentials(connectorIDOrName, scope, credentialsName string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/connectors/%s/%s/credentials/%s", connectorIDOrName, scope, credentialsName))
}
