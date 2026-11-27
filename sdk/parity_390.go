package sdk

import (
	"fmt"
	"net/http"
	"net/url"
)

// Billing endpoints authenticate with an organization dashboard-context key.
func (c *MistralClient) billingClient(key string) (*MistralClient, error) {
	if key == "" {
		return nil, fmt.Errorf("dashboard context API key is required")
	}
	headers := c.requestHeaders.Clone()
	if headers == nil {
		headers = http.Header{}
	}
	headers.Set("X-Api-Key", key)
	// Suppress bearer authentication for operations using their own scheme.
	headers.Del("Authorization")
	headers["Authorization"] = nil
	return c.WithHeaders(headers), nil
}
func (c *MistralClient) GetBudget(dashboardContextKey string) (*GetBudgetResponse, error) {
	client, err := c.billingClient(dashboardContextKey)
	if err != nil {
		return nil, err
	}
	return requestTyped[GetBudgetResponse](client, http.MethodGet, nil, "v1/billing/budget")
}
func (c *MistralClient) GetSubscription(dashboardContextKey string) (*GetSubscriptionResponse, error) {
	client, err := c.billingClient(dashboardContextKey)
	if err != nil {
		return nil, err
	}
	return requestTyped[GetSubscriptionResponse](client, http.MethodGet, nil, "v1/billing/subscription")
}

type ListScorersParams struct {
	CreatedByMe bool
	OOTB        bool
	Query       *string
	PageSize    *int
	Page        *int
}

func (c *MistralClient) ListScorers(params *ListScorersParams) (*ListScorersResponse, error) {
	if params == nil {
		params = &ListScorersParams{}
	}
	page, size := 1, 50
	if params.Page != nil {
		page = *params.Page
	}
	if params.PageSize != nil {
		size = *params.PageSize
	}
	query := queryWithOptionalValues(map[string]any{"page": page, "page_size": size, "q": params.Query, "created_by_me": params.CreatedByMe, "ootb": params.OOTB})
	return requestTyped[ListScorersResponse](c, http.MethodGet, nil, appendQuery("v1/observability/scorers", query))
}
func (c *MistralClient) CreateScorer(request CreateScorerRequest) (*Scorer, error) {
	return requestTyped[Scorer](c, http.MethodPost, request, "v1/observability/scorers")
}
func (c *MistralClient) GetScorer(slug string) (*Scorer, error) {
	return requestTyped[Scorer](c, http.MethodGet, nil, "v1/observability/scorers/"+url.PathEscape(slug))
}
func (c *MistralClient) UpdateScorer(slug string, request EditScorerRequest) (*Scorer, error) {
	return requestTyped[Scorer](c, http.MethodPatch, request, "v1/observability/scorers/"+url.PathEscape(slug))
}

func (c *MistralClient) CreateTrigger(request CreateTrigger) (*Trigger, error) {
	return requestTyped[Trigger](c, http.MethodPost, request, "v1/triggers")
}

type ListTriggersParams struct {
	Handler       *TriggerHandler
	PrimitiveID   *string
	EventSourceID *string
}

func (c *MistralClient) ListTriggers(params *ListTriggersParams) ([]Trigger, error) {
	if params == nil {
		params = &ListTriggersParams{}
	}
	query := queryWithOptionalValues(map[string]any{"handler": params.Handler, "primitive_id": params.PrimitiveID, "event_source_id": params.EventSourceID})
	return requestList[Trigger](c, appendQuery("v1/triggers", query))
}
func (c *MistralClient) GetTrigger(id string) (*Trigger, error) {
	return requestTyped[Trigger](c, http.MethodGet, nil, "v1/triggers/"+url.PathEscape(id))
}
func (c *MistralClient) DeleteTrigger(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/triggers/"+url.PathEscape(id), false, nil)
	return err
}
func (c *MistralClient) ListTriggerEvents(id string, limit *int) (*TriggerEventList, error) {
	size := 10
	if limit != nil {
		size = *limit
	}
	return requestTyped[TriggerEventList](c, http.MethodGet, nil, appendQuery("v1/triggers/"+url.PathEscape(id)+"/events", queryWithOptionalValues(map[string]any{"limit": size})))
}
func (c *MistralClient) ListWebhooks() ([]Webhook, error) {
	return requestList[Webhook](c, "v1/webhooks")
}
func requestList[T any](client *MistralClient, path string) ([]T, error) {
	response, err := requestTyped[[]T](client, http.MethodGet, nil, path)
	if err != nil {
		return nil, err
	}
	return *response, nil
}
func (c *MistralClient) CreateWebhook(request CreateWebhook) (*Webhook, error) {
	return requestTyped[Webhook](c, http.MethodPost, request, "v1/webhooks")
}
func (c *MistralClient) GetWebhook(id string) (*Webhook, error) {
	return requestTyped[Webhook](c, http.MethodGet, nil, "v1/webhooks/"+url.PathEscape(id))
}
func (c *MistralClient) DeleteWebhook(id string) error {
	_, err := c.request(http.MethodDelete, nil, "v1/webhooks/"+url.PathEscape(id), false, nil)
	return err
}
func (c *MistralClient) GetWebhookSigningSecret(id string) (*WebhookSigningSecret, error) {
	return requestTyped[WebhookSigningSecret](c, http.MethodGet, nil, "v1/webhooks/"+url.PathEscape(id)+"/signing-secret")
}
func (c *MistralClient) CreateDeploymentArchive(size int) (*DeploymentArchiveUploadResponse, error) {
	return requestTyped[DeploymentArchiveUploadResponse](c, http.MethodPost, CreateDeploymentArchiveRequest{Size: size}, "v1/workflows/deployments/archives")
}
func (c *MistralClient) GetDeploymentArchive(name string) (*DeploymentArchiveResponse, error) {
	return requestTyped[DeploymentArchiveResponse](c, http.MethodGet, nil, "v1/workflows/deployments/"+url.PathEscape(name)+"/archive")
}
func (c *MistralClient) AddWorkflowTags(id string, tags []string) error {
	_, err := c.request(http.MethodPost, map[string]any{"tags": tags}, "v1/workflows/"+url.PathEscape(id)+"/tags", false, nil)
	return err
}
func (c *MistralClient) CreateDatasetRecords(id string, records []CreateDatasetRecordsItem) (*CreateDatasetRecordsResponse, error) {
	return requestTyped[CreateDatasetRecordsResponse](c, http.MethodPost, CreateDatasetRecordsRequest{Records: records}, "v1/observability/datasets/"+url.PathEscape(id)+"/records/bulk")
}
func (c *MistralClient) ListPipelineResultFields(id string) (*GetPipelineResultFields, error) {
	return requestTyped[GetPipelineResultFields](c, http.MethodGet, nil, "v1/observability/pipelines/"+url.PathEscape(id)+"/results/fields")
}
func (c *MistralClient) UpdateAudioSession(id string, request AudioSessionUpdateRequest) (*AudioSessionUpdateResponse, error) {
	return requestTyped[AudioSessionUpdateResponse](c, http.MethodPatch, request, "v1/audio/sessions/"+url.PathEscape(id))
}

func (c *MistralClient) ModerateWithParams(request ModerationRequest) (*ModerationResponse, error) {
	return requestTyped[ModerationResponse](c, http.MethodPost, request, "v1/moderations")
}
