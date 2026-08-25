package sdk

import (
	"fmt"
	"net/http"
	"time"
)

type ClientSessionPurpose string

const ClientSessionPurposeRealtime ClientSessionPurpose = "realtime"

type CreateRealtimeSessionRequest struct {
	Model      string               `json:"model"`
	Purpose    ClientSessionPurpose `json:"purpose,omitempty"`
	TTLSeconds *int                 `json:"ttl_seconds,omitempty"`
}

type ClientSecret struct {
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
}

type CreateRealtimeSessionResponse struct {
	Purpose      ClientSessionPurpose `json:"purpose"`
	ExpiresAt    time.Time            `json:"expires_at"`
	ClientSecret ClientSecret         `json:"client_secret"`
	Object       string               `json:"object,omitempty"`
}

// CreateRealtimeSession creates a short-lived client session for the realtime API.
func (c *MistralClient) CreateRealtimeSession(req *CreateRealtimeSessionRequest) (*CreateRealtimeSessionResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.Model == "" {
		return nil, fmt.Errorf("model is required")
	}
	purpose := req.Purpose
	if purpose == "" {
		purpose = ClientSessionPurposeRealtime
	}
	body := map[string]interface{}{"model": req.Model, "purpose": purpose}
	if req.TTLSeconds != nil {
		body["ttl_seconds"] = *req.TTLSeconds
	}
	response, err := c.request(http.MethodPost, body, "v1/client/sessions", false, nil)
	if err != nil {
		return nil, err
	}
	data, ok := response.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response type: %T", response)
	}
	var out CreateRealtimeSessionResponse
	if err := mapToStruct(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
