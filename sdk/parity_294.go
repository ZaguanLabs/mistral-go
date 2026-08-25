package sdk

// MessageTokens reports per-message tokenizer usage.
type MessageTokens struct {
	Role           string `json:"role"`
	TotalTokens    *int   `json:"total_tokens,omitempty"`
	SettingsTokens *int   `json:"settings_tokens,omitempty"`
	Truncated      bool   `json:"truncated,omitempty"`
	UsageCount     int    `json:"usage_count,omitempty"`
}

type RegisterDeploymentRequestVespaField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type RegisterDeploymentRequestVespaIndex struct {
	Name                string                                `json:"name"`
	Fields              []RegisterDeploymentRequestVespaField `json:"fields"`
	SD                  string                                `json:"sd"`
	EmbeddingDimensions *int                                  `json:"embedding_dimensions,omitempty"`
}

type TempoTraceAttributeDoubleValue struct {
	DoubleValue float64 `json:"doubleValue"`
}

type TempoTraceAttributeArrayElement struct {
	StringValue *string  `json:"stringValue,omitempty"`
	IntValue    *string  `json:"intValue,omitempty"`
	DoubleValue *float64 `json:"doubleValue,omitempty"`
	BoolValue   *bool    `json:"boolValue,omitempty"`
}
