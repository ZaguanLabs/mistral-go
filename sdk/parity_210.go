package sdk

import "encoding/json"

type ShareRelation string

const (
	ShareRelationUnspecified ShareRelation = "share_relation_unspecified"
	ShareRelationReader      ShareRelation = "reader"
	ShareRelationWriter      ShareRelation = "writer"
)

type DeploymentK8sBackendSpec struct {
	Type       string  `json:"type"`
	Entrypoint *string `json:"entrypoint,omitempty"`
	WorkingDir *string `json:"working_dir,omitempty"`
}
type DeploymentKoyebBackendSpec struct {
	Type           string  `json:"type"`
	BuildDirectory *string `json:"build_directory,omitempty"`
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
}

func (v DeploymentK8sBackendSpec) MarshalJSON() ([]byte, error) {
	type wire DeploymentK8sBackendSpec
	v.Type = "kubernetes"
	return json.Marshal(wire(v))
}
func (v DeploymentKoyebBackendSpec) MarshalJSON() ([]byte, error) {
	type wire DeploymentKoyebBackendSpec
	v.Type = "koyeb"
	return json.Marshal(wire(v))
}

// UsageInfoDollarDefs is the dollar-usage variant. Unlike UsageInfo it has no service_tier.
type UsageInfoDollarDefs struct {
	PromptTokens            int            `json:"prompt_tokens"`
	TotalTokens             int            `json:"total_tokens"`
	CompletionTokens        *int           `json:"completion_tokens,omitempty"`
	PromptAudioSeconds      *int           `json:"prompt_audio_seconds,omitempty"`
	RequestCount            *int           `json:"request_count,omitempty"`
	PromptTokensDetails     map[string]any `json:"prompt_tokens_details,omitempty"`
	CompletionTokensDetails map[string]any `json:"completion_tokens_details,omitempty"`
	PromptTokenDetails      map[string]any `json:"prompt_token_details,omitempty"`
	NumCachedTokens         *int           `json:"num_cached_tokens,omitempty"`
}
