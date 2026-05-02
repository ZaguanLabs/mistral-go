package sdk

import (
	"fmt"
	"net/http"
)

type ConnectorRequest struct {
	Name         string         `json:"name,omitempty"`
	Description  *string        `json:"description,omitempty"`
	Server       any            `json:"server,omitempty"`
	Title        *string        `json:"title,omitempty"`
	IconURL      *string        `json:"icon_url,omitempty"`
	Visibility   any            `json:"visibility,omitempty"`
	Headers      map[string]any `json:"headers,omitempty"`
	AuthData     map[string]any `json:"auth_data,omitempty"`
	SystemPrompt *string        `json:"system_prompt,omitempty"`
}

type UpdateConnectorRequest struct {
	Title             *string        `json:"title,omitempty"`
	Name              *string        `json:"name,omitempty"`
	Description       *string        `json:"description,omitempty"`
	IconURL           *string        `json:"icon_url,omitempty"`
	SystemPrompt      *string        `json:"system_prompt,omitempty"`
	ConnectionConfig  map[string]any `json:"connection_config,omitempty"`
	ConnectionSecrets map[string]any `json:"connection_secrets,omitempty"`
	Server            any            `json:"server,omitempty"`
	Headers           map[string]any `json:"headers,omitempty"`
	AuthData          map[string]any `json:"auth_data,omitempty"`
}

type ListConnectorsParams struct {
	QueryFilters map[string]any `json:"query_filters,omitempty"`
	Cursor       *string        `json:"cursor,omitempty"`
	PageSize     *int           `json:"page_size,omitempty"`
}

type ConnectorCredentialsRequest struct {
	Name        string         `json:"name,omitempty"`
	IsDefault   *bool          `json:"is_default,omitempty"`
	Credentials map[string]any `json:"credentials,omitempty"`
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

func (c *MistralClient) CreateConnector(req *ConnectorRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"name":          req.Name,
		"description":   req.Description,
		"server":        req.Server,
		"title":         req.Title,
		"icon_url":      req.IconURL,
		"visibility":    req.Visibility,
		"headers":       req.Headers,
		"auth_data":     req.AuthData,
		"system_prompt": req.SystemPrompt,
	})
	return c.requestMap(http.MethodPost, body, "v1/connectors")
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

func (c *MistralClient) GetConnectorAuthURL(connectorIDOrName string, appReturnURL *string, credentialsName *string) (APIResponse, error) {
	query := queryWithOptionalValues(map[string]any{"app_return_url": appReturnURL, "credentials_name": credentialsName})
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

func (c *MistralClient) ListOrganizationConnectorCredentials(connectorIDOrName string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	return c.listConnectorCredentials(connectorIDOrName, "organization", params)
}

func (c *MistralClient) CreateOrUpdateOrganizationConnectorCredentials(connectorIDOrName string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.createOrUpdateConnectorCredentials(connectorIDOrName, "organization", req)
}

func (c *MistralClient) ListWorkspaceConnectorCredentials(connectorIDOrName string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	return c.listConnectorCredentials(connectorIDOrName, "workspace", params)
}

func (c *MistralClient) CreateOrUpdateWorkspaceConnectorCredentials(connectorIDOrName string, req *ConnectorCredentialsRequest) (APIResponse, error) {
	return c.createOrUpdateConnectorCredentials(connectorIDOrName, "workspace", req)
}

func (c *MistralClient) ListUserConnectorCredentials(connectorIDOrName string, params *ListConnectorCredentialsParams) (APIResponse, error) {
	return c.listConnectorCredentials(connectorIDOrName, "user", params)
}

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

func (c *MistralClient) UpdateConnector(connectorID string, req *UpdateConnectorRequest) (APIResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	body := optionalRequestMap(map[string]any{
		"title":              req.Title,
		"name":               req.Name,
		"description":        req.Description,
		"icon_url":           req.IconURL,
		"system_prompt":      req.SystemPrompt,
		"connection_config":  req.ConnectionConfig,
		"connection_secrets": req.ConnectionSecrets,
		"server":             req.Server,
		"headers":            req.Headers,
		"auth_data":          req.AuthData,
	})
	return c.requestMap(http.MethodPatch, body, fmt.Sprintf("v1/connectors/%s", connectorID))
}

func (c *MistralClient) DeleteConnector(connectorID string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/connectors/%s", connectorID))
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
	body := optionalRequestMap(map[string]any{"name": req.Name, "is_default": req.IsDefault, "credentials": req.Credentials})
	return c.requestMap(http.MethodPost, body, fmt.Sprintf("v1/connectors/%s/%s/credentials", connectorIDOrName, scope))
}

func (c *MistralClient) deleteConnectorCredentials(connectorIDOrName, scope, credentialsName string) (APIResponse, error) {
	return c.requestMap(http.MethodDelete, nil, fmt.Sprintf("v1/connectors/%s/%s/credentials/%s", connectorIDOrName, scope, credentialsName))
}
