// Code generated from official Mistral Python SDK 3.9.0 models. DO NOT EDIT.
package sdk

import (
	"encoding/json"
	"fmt"
	"time"
)

type AudioSessionUpdateRequestTool = any
type AudioSessionUpdateResponseStatus = string
type AuthStatus = string
type BackendSpec = any
type BaseTaskStatus = string
type BuiltInConnectors = string
type ConnectorErrorCode = string
type ConsumerType = string
type CreateDatasetRecordsItemSource = string
type CredentialsStatusErrorReason = string
type DatasetRecordSource = string
type DisplayFormat = string
type EmbeddingInput = any
type EmbeddingRequestInputs = any
type EncodedPayloadOptions = string
type Fallback = StaticToolFiller
type FieldDefinition = any
type FunctionCallEntryArguments = any
type FunctionCallEntryConfirmationStatus = string
type GetSpanStatusCode = string
type GetTraceStatusCode = string
type HTTPStatus = any
type HangUpToolFiller = any
type ImageDetail = string
type ImageURLUnion = any
type InputEntries = any
type MessageOutputContentChunks = any
type MessageOutputEntryContent = any
type MetricAggregation = string
type MetricDirectionSpecDirection = string
type MetricValue = any
type ModerationRequestInputs = any
type ObservabilityErrorCode = string
type OpenAPISpecSource = any
type OtelEnumFieldDefinitionSupportedOperator = string
type OtelFieldDefinitionSupportedOperator = string
type OtelFieldDefinitionType = string
type OtelFloatFieldDefinitionSupportedOperator = string
type RealtimeSpeechErrorDetailMessage = any
type RealtimeSpeechStreamingMode = string
type ReferenceID = any
type ResourceVisibility = string
type ResponsesErrorReason = string
type ResponsesIncompleteDetailsReason = string
type ResponsesInputFileContentTTL = string
type ResponsesInputFileTTL = string
type ResponsesInputImageContentTTL = string
type ResponsesInputImageTTL = string
type ResponsesInputTextContentTTL = string
type ResponsesInputTextTTL = string
type ScheduleOverlapPolicy = any
type ScoreMetricValue = any
type ScorerKind = string
type SignalInvocationBodyCallContext = any
type SignalInvocationBodyInput = any
type Thinking = any
type TokenEndpointAuthMethod = string
type ToolExecutionEntryName = any
type ToolFileChunkTool = any
type ToolReferenceChunkTool = any
type TranscriptionStreamEventTypes = string
type TranscriptionStreamEventsData = any
type TransferCallToolFiller = any
type TriggerActionType = string
type TriggerHandler = string
type TriggerStatusValue = string
type TriggersOwnerType = string
type UpdateInvocationBodyCallContext = any
type UpdateInvocationBodyInput = any
type VoiceAgentCustomConnectorFiller = any
type VoiceAgentDocumentLibraryToolFiller = any
type VoiceAgentFunctionToolFiller = any
type VoiceAgentFunctionToolRequestFiller = any
type VoiceAgentWebSearchToolFiller = any
type WebhookProvider = string
type WebhookSignatureEncoding = string
type WebhookSignatureHashAlgorithm = string
type WorkflowExecutionRequestCallContext = any
type AgentHandoffEntry struct {
	PreviousAgentID   string     `json:"previous_agent_id"`
	PreviousAgentName string     `json:"previous_agent_name"`
	NextAgentID       string     `json:"next_agent_id"`
	NextAgentName     string     `json:"next_agent_name"`
	Object            *string    `json:"object,omitempty"`
	Type              *string    `json:"type,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	CompletedAt       *time.Time `json:"completed_at,omitempty"`
	ID                *string    `json:"id,omitempty"`
}

func (v AgentHandoffEntry) MarshalJSON() ([]byte, error) {
	type wire AgentHandoffEntry
	if v.Object == nil {
		x := string("entry")
		v.Object = &x
	}
	if v.Type == nil {
		x := string("agent.handoff")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type AmazonConnectContact struct {
	ContextID    string  `json:"context_id"`
	Type         string  `json:"type,omitempty"`
	InstanceArn  *string `json:"instance_arn,omitempty"`
	InstanceID   *string `json:"instance_id,omitempty"`
	ContactArn   *string `json:"contact_arn,omitempty"`
	ContactID    *string `json:"contact_id,omitempty"`
	CallerNumber *string `json:"caller_number,omitempty"`
	CalledNumber *string `json:"called_number,omitempty"`
	Channel      *string `json:"channel,omitempty"`
}

func (v AmazonConnectContact) MarshalJSON() ([]byte, error) {
	type wire AmazonConnectContact
	if v.Type == "" {
		v.Type = "amazon_connect"
	}
	return json.Marshal(wire(v))
}

type AudioChunk struct {
	InputAudio string  `json:"input_audio"`
	Type       *string `json:"type,omitempty"`
}

func (v AudioChunk) MarshalJSON() ([]byte, error) {
	type wire AudioChunk
	if v.Type == nil {
		x := string("input_audio")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type AudioSessionTurnUpdate struct {
	AllowInterruptions *bool    `json:"allow_interruptions,omitempty"`
	ExplicitNulls      []string `json:"-"`
}

func (v AudioSessionTurnUpdate) MarshalJSON() ([]byte, error) {
	type wire AudioSessionTurnUpdate
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "allow_interruptions":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type AudioSessionUpdateRequest struct {
	Instructions  *string                         `json:"instructions,omitempty"`
	Tools         []AudioSessionUpdateRequestTool `json:"tools,omitempty"`
	Turn          *AudioSessionTurnUpdate         `json:"turn,omitempty"`
	ExplicitNulls []string                        `json:"-"`
}

func (v AudioSessionUpdateRequest) MarshalJSON() ([]byte, error) {
	type wire AudioSessionUpdateRequest
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Tools != nil {
		data, err := json.Marshal(v.Tools)
		if err != nil {
			return nil, err
		}
		fields["tools"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "instructions", "tools", "turn":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type AudioSessionUpdateResponse struct {
	ID              string                           `json:"id"`
	Status          AudioSessionUpdateResponseStatus `json:"status"`
	NumParticipants int                              `json:"num_participants"`
	CreatedAt       time.Time                        `json:"created_at"`
	PhoneCall       *AudioSessionPhoneCall           `json:"phone_call,omitempty"`
	BridgeMetadata  *AmazonConnectContact            `json:"bridge_metadata,omitempty"`
}
type AuthenticationConfiguration struct {
	Name               string             `json:"name"`
	AuthenticationType AuthenticationType `json:"authentication_type"`
	Scope              ConsumerType       `json:"scope"`
	Title              *string            `json:"title,omitempty"`
	Status             *CredentialsStatus `json:"status,omitempty"`
	IsDefault          *bool              `json:"is_default,omitempty"`
	CreatorID          *string            `json:"creator_id,omitempty"`
}

func (v AuthenticationConfiguration) MarshalJSON() ([]byte, error) {
	type wire AuthenticationConfiguration
	if v.IsDefault == nil {
		x := bool(false)
		v.IsDefault = &x
	}
	return json.Marshal(wire(v))
}

type BudgetDetail struct {
	UsagePercentage float64   `json:"usage_percentage"`
	InitialBudget   float64   `json:"initial_budget"`
	ResetAt         time.Time `json:"reset_at"`
	Currency        *string   `json:"currency,omitempty"`
	PaygEnabled     *bool     `json:"payg_enabled,omitempty"`
}

func (v BudgetDetail) MarshalJSON() ([]byte, error) {
	type wire BudgetDetail
	if v.PaygEnabled == nil {
		x := bool(false)
		v.PaygEnabled = &x
	}
	return json.Marshal(wire(v))
}

type ConnectionPreference struct {
	ToolConfiguration ToolExecutionConfiguration `json:"tool_configuration"`
	ConsumerType      *ConsumerType              `json:"consumer_type,omitempty"`
	HttpProxyEnabled  *bool                      `json:"http_proxy_enabled,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	IsDefault         *bool                      `json:"is_default,omitempty"`
}

func (v ConnectionPreference) MarshalJSON() ([]byte, error) {
	type wire ConnectionPreference
	if v.Name == nil {
		x := string("default")
		v.Name = &x
	}
	return json.Marshal(wire(v))
}

type ConnectorOpenAPISpec struct {
	Source *OpenAPISpecSource `json:"source"`
}
type ConnectorTool struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	ExecutionConfig *ExecutionConfig     `json:"execution_config"`
	Visibility      ResourceVisibility   `json:"visibility"`
	CreatedAt       time.Time            `json:"created_at"`
	ModifiedAt      time.Time            `json:"modified_at"`
	SystemPrompt    *string              `json:"system_prompt,omitempty"`
	Locale          *ConnectorToolLocale `json:"locale,omitempty"`
	Jsonschema      map[string]any       `json:"jsonschema,omitempty"`
	Active          *bool                `json:"active,omitempty"`
}

func (v ConnectorTool) MarshalJSON() ([]byte, error) {
	type wire ConnectorTool
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Jsonschema != nil {
		data, err := json.Marshal(v.Jsonschema)
		if err != nil {
			return nil, err
		}
		fields["jsonschema"] = data
	}
	return json.Marshal(fields)
}

type ConnectorToolLocale struct {
	Name          map[string]string `json:"name"`
	Description   map[string]string `json:"description"`
	UsageSentence map[string]string `json:"usage_sentence"`
}
type ConnectorToolResultMetadata struct {
	AdditionalProperties map[string]any `json:"-"`
	IsError              *bool          `json:"isError,omitempty"`
	StructuredContent    any            `json:"structuredContent,omitempty"`
	Meta                 map[string]any `json:"_meta,omitempty"`
}

func (v ConnectorToolResultMetadata) MarshalJSON() ([]byte, error) {
	type wire ConnectorToolResultMetadata
	if v.IsError == nil {
		x := bool(false)
		v.IsError = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Meta != nil {
		data, err := json.Marshal(v.Meta)
		if err != nil {
			return nil, err
		}
		fields["_meta"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "isError", "structuredContent", "_meta":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *ConnectorToolResultMetadata) UnmarshalJSON(data []byte) error {
	type wire ConnectorToolResultMetadata
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "isError")
	delete(fields, "structuredContent")
	delete(fields, "_meta")
	*v = ConnectorToolResultMetadata(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type CreateDatasetRecordsItem struct {
	Payload       map[string]any                  `json:"payload"`
	Properties    map[string]any                  `json:"properties,omitempty"`
	Source        *CreateDatasetRecordsItemSource `json:"source,omitempty"`
	ID            *string                         `json:"id,omitempty"`
	ExplicitNulls []string                        `json:"-"`
}

func (v CreateDatasetRecordsItem) MarshalJSON() ([]byte, error) {
	type wire CreateDatasetRecordsItem
	if v.Source == nil {
		x := CreateDatasetRecordsItemSource("DIRECT_INPUT")
		v.Source = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Properties != nil {
		data, err := json.Marshal(v.Properties)
		if err != nil {
			return nil, err
		}
		fields["properties"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "id":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type CreateDatasetRecordsRequest struct {
	Records []CreateDatasetRecordsItem `json:"records"`
}
type CreateDatasetRecordsResponse struct {
	Records []DatasetRecord `json:"records"`
}
type CreateDeploymentArchiveRequest struct {
	Size int `json:"size"`
}
type CreateIngestionPipelineConfigurationRequest struct {
	Name                string            `json:"name"`
	PipelineComposition map[string]string `json:"pipeline_composition,omitempty"`
	ExplicitNulls       []string          `json:"-"`
}

func (v CreateIngestionPipelineConfigurationRequest) MarshalJSON() ([]byte, error) {
	type wire CreateIngestionPipelineConfigurationRequest
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.PipelineComposition != nil {
		data, err := json.Marshal(v.PipelineComposition)
		if err != nil {
			return nil, err
		}
		fields["pipeline_composition"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "pipeline_composition":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type CreateScorerRequest struct {
	Name          string         `json:"name"`
	Kind          ScorerKind     `json:"kind"`
	Config        LlmJudgeConfig `json:"config"`
	Description   *string        `json:"description,omitempty"`
	ExplicitNulls []string       `json:"-"`
}

func (v CreateScorerRequest) MarshalJSON() ([]byte, error) {
	type wire CreateScorerRequest
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "description":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type CreateTrigger struct {
	EventSourceID string              `json:"event_source_id"`
	Handler       TriggerHandler      `json:"handler"`
	Action        CreateTriggerAction `json:"action"`
	OwnerType     TriggersOwnerType   `json:"owner_type"`
	Title         string              `json:"title"`
	Filters       *string             `json:"filters,omitempty"`
	ExplicitNulls []string            `json:"-"`
}

func (v CreateTrigger) MarshalJSON() ([]byte, error) {
	type wire CreateTrigger
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "filters":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type CreateTriggerAction struct {
	PrimitiveID string            `json:"primitive_id"`
	Type        TriggerActionType `json:"type"`
}
type CreateWebhook struct {
	Provider                WebhookProvider                `json:"provider"`
	Name                    string                         `json:"name"`
	OwnerType               *TriggersOwnerType             `json:"owner_type,omitempty"`
	Description             *string                        `json:"description,omitempty"`
	Secret                  *string                        `json:"secret,omitempty"`
	EncryptionKey           *string                        `json:"encryption_key,omitempty"`
	URLVerificationToken    *string                        `json:"url_verification_token,omitempty"`
	ExpiresAt               *time.Time                     `json:"expires_at,omitempty"`
	SignatureHashAlgorithm  *WebhookSignatureHashAlgorithm `json:"signature_hash_algorithm,omitempty"`
	SignatureHeaderName     *string                        `json:"signature_header_name,omitempty"`
	ValidityPeriodSecond    *int                           `json:"validity_period_second,omitempty"`
	SignatureExtractorRegex *string                        `json:"signature_extractor_regex,omitempty"`
	TimestampExtractorRegex *string                        `json:"timestamp_extractor_regex,omitempty"`
	TimestampPattern        *string                        `json:"timestamp_pattern,omitempty"`
	SignaturePayloadPattern *string                        `json:"signature_payload_pattern,omitempty"`
	SignatureEncoding       *WebhookSignatureEncoding      `json:"signature_encoding,omitempty"`
	SuccessStatusCode       *int                           `json:"success_status_code,omitempty"`
	EventTypePattern        *string                        `json:"event_type_pattern,omitempty"`
	EventIDPattern          *string                        `json:"event_id_pattern,omitempty"`
	ExplicitNulls           []string                       `json:"-"`
}

func (v CreateWebhook) MarshalJSON() ([]byte, error) {
	type wire CreateWebhook
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "description", "secret", "encryption_key", "url_verification_token", "expires_at", "signature_hash_algorithm", "signature_header_name", "validity_period_second", "signature_extractor_regex", "timestamp_extractor_regex", "timestamp_pattern", "signature_payload_pattern", "signature_encoding", "success_status_code", "event_type_pattern", "event_id_pattern":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type CredentialsStatus struct {
	StatusType    AuthStatus                    `json:"status_type"`
	LastCheckedAt *time.Time                    `json:"last_checked_at,omitempty"`
	ErrorHttpCode *HTTPStatus                   `json:"error_http_code,omitempty"`
	ErrorMessage  *CredentialsStatusErrorReason `json:"error_message,omitempty"`
}
type Dataset struct {
	ID          string         `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	OwnerID     string         `json:"owner_id"`
	WorkspaceID string         `json:"workspace_id"`
	OwnerType   *PrincipalType `json:"owner_type,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

func (v Dataset) MarshalJSON() ([]byte, error) {
	type wire Dataset
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Metadata != nil {
		data, err := json.Marshal(v.Metadata)
		if err != nil {
			return nil, err
		}
		fields["metadata"] = data
	}
	return json.Marshal(fields)
}

type DatasetPreview struct {
	ID          string         `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	OwnerID     string         `json:"owner_id"`
	WorkspaceID string         `json:"workspace_id"`
	OwnerType   *PrincipalType `json:"owner_type,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

func (v DatasetPreview) MarshalJSON() ([]byte, error) {
	type wire DatasetPreview
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Metadata != nil {
		data, err := json.Marshal(v.Metadata)
		if err != nil {
			return nil, err
		}
		fields["metadata"] = data
	}
	return json.Marshal(fields)
}

type DatasetRecord struct {
	ID         string              `json:"id"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	DeletedAt  *time.Time          `json:"deleted_at"`
	DatasetID  string              `json:"dataset_id"`
	Payload    map[string]any      `json:"payload"`
	Properties map[string]any      `json:"properties"`
	Source     DatasetRecordSource `json:"source"`
}
type DeploymentArchiveResponse struct {
	ArchiveID   string  `json:"archive_id"`
	DownloadURL *string `json:"download_url,omitempty"`
}
type DeploymentArchiveUploadResponse struct {
	ArchiveID     string            `json:"archive_id"`
	UploadURL     string            `json:"upload_url"`
	UploadHeaders map[string]string `json:"upload_headers"`
}
type DeploymentBuildState struct {
	Phase      *string    `json:"phase,omitempty"`
	CommitSha  *string    `json:"commit_sha,omitempty"`
	Image      *string    `json:"image,omitempty"`
	Message    *string    `json:"message,omitempty"`
	StartedAt  *time.Time `json:"started_at,omitempty"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}
type DeploymentObservedState struct {
	Phase              *string               `json:"phase,omitempty"`
	AvailableReplicas  *int                  `json:"available_replicas,omitempty"`
	ReadyReplicas      *int                  `json:"ready_replicas,omitempty"`
	Endpoint           *string               `json:"endpoint,omitempty"`
	Message            *string               `json:"message,omitempty"`
	LastSeen           *time.Time            `json:"last_seen,omitempty"`
	DeployedRevision   *string               `json:"deployed_revision,omitempty"`
	Generation         *int                  `json:"generation,omitempty"`
	LastPhaseChangedAt *time.Time            `json:"last_phase_changed_at,omitempty"`
	BuildState         *DeploymentBuildState `json:"build_state,omitempty"`
}

func (v DeploymentObservedState) MarshalJSON() ([]byte, error) {
	type wire DeploymentObservedState
	if v.AvailableReplicas == nil {
		x := int(0)
		v.AvailableReplicas = &x
	}
	if v.ReadyReplicas == nil {
		x := int(0)
		v.ReadyReplicas = &x
	}
	return json.Marshal(wire(v))
}

type DeploymentWorkerSpecInput struct {
	GithubURL      *string                            `json:"github_url,omitempty"`
	RepositoryURL  *string                            `json:"repository_url,omitempty"`
	Revision       *string                            `json:"revision,omitempty"`
	ArchiveID      *string                            `json:"archive_id,omitempty"`
	BackendSpec    *DeploymentMistralCloudBackendSpec `json:"backend_spec,omitempty"`
	RequestedRoles []string                           `json:"requested_roles,omitempty"`
}

func (v DeploymentWorkerSpecInput) MarshalJSON() ([]byte, error) {
	type wire DeploymentWorkerSpecInput
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.RequestedRoles != nil {
		data, err := json.Marshal(v.RequestedRoles)
		if err != nil {
			return nil, err
		}
		fields["requested_roles"] = data
	}
	return json.Marshal(fields)
}

type DisplaySpec struct {
	Format   *DisplayFormat `json:"format,omitempty"`
	Unit     *string        `json:"unit,omitempty"`
	Decimals *int           `json:"decimals,omitempty"`
	Metric   *string        `json:"metric,omitempty"`
}
type DocumentURLChunk struct {
	DocumentURL  string  `json:"document_url"`
	Type         *string `json:"type,omitempty"`
	DocumentName *string `json:"document_name,omitempty"`
}

func (v DocumentURLChunk) MarshalJSON() ([]byte, error) {
	type wire DocumentURLChunk
	if v.Type == nil {
		x := string("document_url")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type EditScorerRequest struct {
	Name          *string         `json:"name,omitempty"`
	Description   *string         `json:"description,omitempty"`
	Config        *LlmJudgeConfig `json:"config,omitempty"`
	ExplicitNulls []string        `json:"-"`
}

func (v EditScorerRequest) MarshalJSON() ([]byte, error) {
	type wire EditScorerRequest
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "name", "description", "config":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type EnumConstraints struct {
	Options []string `json:"options,omitempty"`
}

func (v EnumConstraints) MarshalJSON() ([]byte, error) {
	type wire EnumConstraints
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Options != nil {
		data, err := json.Marshal(v.Options)
		if err != nil {
			return nil, err
		}
		fields["options"] = data
	}
	return json.Marshal(fields)
}

type EvaluationRunRunScoreV2 struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at"`
	Value     any            `json:"value"`
	Metadata  map[string]any `json:"metadata_"`
	Status    string         `json:"status"`
	Rationale *string        `json:"rationale,omitempty"`
	Error     *string        `json:"error,omitempty"`
	Metrics   []ScoreMetric  `json:"metrics,omitempty"`
}

func (v EvaluationRunRunScoreV2) MarshalJSON() ([]byte, error) {
	type wire EvaluationRunRunScoreV2
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Metrics != nil {
		data, err := json.Marshal(v.Metrics)
		if err != nil {
			return nil, err
		}
		fields["metrics"] = data
	}
	return json.Marshal(fields)
}

type ExecutionConfig struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type"`
}

func (v ExecutionConfig) MarshalJSON() ([]byte, error) {
	type wire ExecutionConfig
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "type":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *ExecutionConfig) UnmarshalJSON(data []byte) error {
	type wire ExecutionConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = ExecutionConfig(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type FloatConstraints struct {
	Min *float64 `json:"min,omitempty"`
	Max *float64 `json:"max,omitempty"`
}
type FunctionCallEntry struct {
	ToolCallID         string                               `json:"tool_call_id"`
	Name               string                               `json:"name"`
	Arguments          FunctionCallEntryArguments           `json:"arguments"`
	Object             *string                              `json:"object,omitempty"`
	Type               *string                              `json:"type,omitempty"`
	CreatedAt          *time.Time                           `json:"created_at,omitempty"`
	CompletedAt        *time.Time                           `json:"completed_at,omitempty"`
	AgentID            *string                              `json:"agent_id,omitempty"`
	Model              *string                              `json:"model,omitempty"`
	ID                 *string                              `json:"id,omitempty"`
	ConfirmationStatus *FunctionCallEntryConfirmationStatus `json:"confirmation_status,omitempty"`
}

func (v FunctionCallEntry) MarshalJSON() ([]byte, error) {
	type wire FunctionCallEntry
	if v.Object == nil {
		x := string("entry")
		v.Object = &x
	}
	if v.Type == nil {
		x := string("function.call")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type FunctionResultEntry struct {
	ToolCallID  string     `json:"tool_call_id"`
	Result      string     `json:"result"`
	Object      *string    `json:"object,omitempty"`
	Type        *string    `json:"type,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	ID          *string    `json:"id,omitempty"`
}

func (v FunctionResultEntry) MarshalJSON() ([]byte, error) {
	type wire FunctionResultEntry
	if v.Object == nil {
		x := string("entry")
		v.Object = &x
	}
	if v.Type == nil {
		x := string("function.result")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type GeneratedToolFiller struct {
	Type         string    `json:"type,omitempty"`
	Instructions *string   `json:"instructions,omitempty"`
	Fallback     *Fallback `json:"fallback,omitempty"`
}

func (v GeneratedToolFiller) MarshalJSON() ([]byte, error) {
	type wire GeneratedToolFiller
	if v.Type == "" {
		v.Type = "generated"
	}
	return json.Marshal(wire(v))
}

type GetBudgetResponse struct {
	VibeBudget BudgetDetail `json:"vibe_budget"`
	APIBudget  BudgetDetail `json:"api_budget"`
}
type GetLog struct {
	CustomerID         string            `json:"customer_id"`
	OrganizationID     string            `json:"organization_id"`
	WorkspaceID        string            `json:"workspace_id"`
	UserID             string            `json:"user_id"`
	Timestamp          time.Time         `json:"timestamp"`
	TraceID            string            `json:"trace_id"`
	SpanID             string            `json:"span_id"`
	TraceFlags         int               `json:"trace_flags"`
	SeverityText       string            `json:"severity_text"`
	SeverityNumber     int               `json:"severity_number"`
	ServiceName        string            `json:"service_name"`
	Body               string            `json:"body"`
	EventName          string            `json:"event_name"`
	ResourceSchemaURL  string            `json:"resource_schema_url"`
	ResourceAttributes map[string]string `json:"resource_attributes"`
	ScopeSchemaURL     string            `json:"scope_schema_url"`
	ScopeName          string            `json:"scope_name"`
	ScopeVersion       string            `json:"scope_version"`
	ScopeAttributes    map[string]string `json:"scope_attributes"`
	LogAttributes      map[string]string `json:"log_attributes"`
	ServiceAccountID   *string           `json:"service_account_id,omitempty"`
}

func (v GetLog) MarshalJSON() ([]byte, error) {
	type wire GetLog
	if v.ServiceAccountID == nil {
		x := string("")
		v.ServiceAccountID = &x
	}
	return json.Marshal(wire(v))
}

type GetPipelineResultFields struct {
	FieldDefinitions []FieldDefinition `json:"field_definitions"`
}
type GetSpan struct {
	CustomerID                    string            `json:"customer_id"`
	OrganizationID                string            `json:"organization_id"`
	WorkspaceID                   string            `json:"workspace_id"`
	UserID                        string            `json:"user_id"`
	TraceID                       string            `json:"trace_id"`
	SpanID                        string            `json:"span_id"`
	ParentSpanID                  string            `json:"parent_span_id"`
	TraceState                    string            `json:"trace_state"`
	StartTime                     time.Time         `json:"start_time"`
	EndTime                       time.Time         `json:"end_time"`
	DurationNS                    int               `json:"duration_ns"`
	SpanName                      string            `json:"span_name"`
	SpanKind                      string            `json:"span_kind"`
	ServiceName                   string            `json:"service_name"`
	StatusCode                    GetSpanStatusCode `json:"status_code"`
	StatusMessage                 string            `json:"status_message"`
	ErrorType                     string            `json:"error_type"`
	OperationName                 string            `json:"operation_name"`
	ProviderName                  string            `json:"provider_name"`
	RequestModel                  string            `json:"request_model"`
	ResponseModel                 string            `json:"response_model"`
	ResponseID                    string            `json:"response_id"`
	OutputType                    string            `json:"output_type"`
	ConversationID                string            `json:"conversation_id"`
	DataSourceID                  string            `json:"data_source_id"`
	AgentID                       string            `json:"agent_id"`
	AgentName                     string            `json:"agent_name"`
	AgentVersion                  string            `json:"agent_version"`
	AgentDescription              string            `json:"agent_description"`
	WorkflowName                  string            `json:"workflow_name"`
	PromptName                    string            `json:"prompt_name"`
	ToolName                      string            `json:"tool_name"`
	ToolType                      string            `json:"tool_type"`
	ToolCallID                    string            `json:"tool_call_id"`
	InputMessages                 string            `json:"input_messages"`
	OutputMessages                string            `json:"output_messages"`
	SystemInstructions            string            `json:"system_instructions"`
	ToolDefinitions               string            `json:"tool_definitions"`
	ToolCallArguments             string            `json:"tool_call_arguments"`
	ToolCallResult                string            `json:"tool_call_result"`
	RequestChoiceCount            int               `json:"request_choice_count"`
	RequestMaxTokens              int               `json:"request_max_tokens"`
	RequestTemperature            *float64          `json:"request_temperature"`
	RequestTopP                   *float64          `json:"request_top_p"`
	RequestTopK                   *float64          `json:"request_top_k"`
	RequestPresencePenalty        *float64          `json:"request_presence_penalty"`
	RequestFrequencyPenalty       *float64          `json:"request_frequency_penalty"`
	RequestSeed                   int               `json:"request_seed"`
	RequestStopSequences          []string          `json:"request_stop_sequences"`
	RequestEncodingFormats        []string          `json:"request_encoding_formats"`
	ResponseFinishReasons         []string          `json:"response_finish_reasons"`
	UsageInputTokens              int               `json:"usage_input_tokens"`
	UsageOutputTokens             int               `json:"usage_output_tokens"`
	UsageCacheReadInputTokens     int               `json:"usage_cache_read_input_tokens"`
	UsageCacheCreationInputTokens int               `json:"usage_cache_creation_input_tokens"`
	ResourceAttributes            map[string]string `json:"resource_attributes"`
	SpanAttributes                map[string]string `json:"span_attributes"`
	ScopeName                     string            `json:"scope_name"`
	ScopeVersion                  string            `json:"scope_version"`
	ServiceAccountID              *string           `json:"service_account_id,omitempty"`
}

func (v GetSpan) MarshalJSON() ([]byte, error) {
	type wire GetSpan
	if v.ServiceAccountID == nil {
		x := string("")
		v.ServiceAccountID = &x
	}
	return json.Marshal(wire(v))
}

type GetSpanEvaluation struct {
	CustomerID       string            `json:"customer_id"`
	OrganizationID   string            `json:"organization_id"`
	WorkspaceID      string            `json:"workspace_id"`
	UserID           string            `json:"user_id"`
	TraceID          string            `json:"trace_id"`
	SpanID           string            `json:"span_id"`
	ResponseID       string            `json:"response_id"`
	ConversationID   string            `json:"conversation_id"`
	Timestamp        time.Time         `json:"timestamp"`
	EvaluationName   string            `json:"evaluation_name"`
	ScoreValue       float64           `json:"score_value"`
	ScoreLabel       string            `json:"score_label"`
	Explanation      string            `json:"explanation"`
	Metadata         map[string]string `json:"metadata"`
	ServiceAccountID *string           `json:"service_account_id,omitempty"`
}

func (v GetSpanEvaluation) MarshalJSON() ([]byte, error) {
	type wire GetSpanEvaluation
	if v.ServiceAccountID == nil {
		x := string("")
		v.ServiceAccountID = &x
	}
	return json.Marshal(wire(v))
}

type GetSubscriptionResponse struct {
	PlanName        string `json:"plan_name"`
	PaygEnabled     bool   `json:"payg_enabled"`
	VibePaygEnabled bool   `json:"vibe_payg_enabled"`
	Enterprise      bool   `json:"enterprise"`
}
type GetTrace struct {
	CustomerID                 string             `json:"customer_id"`
	OrganizationID             string             `json:"organization_id"`
	WorkspaceID                string             `json:"workspace_id"`
	UserID                     string             `json:"user_id"`
	TraceID                    string             `json:"trace_id"`
	RootSpanID                 string             `json:"root_span_id"`
	RootSpanName               string             `json:"root_span_name"`
	StartTime                  time.Time          `json:"start_time"`
	EndTime                    time.Time          `json:"end_time"`
	DurationNS                 int                `json:"duration_ns"`
	ServiceName                string             `json:"service_name"`
	Environment                string             `json:"environment"`
	ConversationID             string             `json:"conversation_id"`
	WorkflowName               string             `json:"workflow_name"`
	AgentID                    string             `json:"agent_id"`
	AgentName                  string             `json:"agent_name"`
	StatusCode                 GetTraceStatusCode `json:"status_code"`
	ErrorCount                 int                `json:"error_count"`
	SpanCount                  int                `json:"span_count"`
	GenAiSpanCount             int                `json:"gen_ai_span_count"`
	LlmCallCount               int                `json:"llm_call_count"`
	ToolCallCount              int                `json:"tool_call_count"`
	RetrievalCount             int                `json:"retrieval_count"`
	EvaluationCount            int                `json:"evaluation_count"`
	InputTokens                int                `json:"input_tokens"`
	OutputTokens               int                `json:"output_tokens"`
	CacheReadInputTokens       int                `json:"cache_read_input_tokens"`
	CacheCreationInputTokens   int                `json:"cache_creation_input_tokens"`
	ModelsUsed                 []string           `json:"models_used"`
	ToolsUsed                  []string           `json:"tools_used"`
	FirstTurnLastInputMessage  string             `json:"first_turn_last_input_message"`
	FirstTurnLastOutputMessage string             `json:"first_turn_last_output_message"`
	LastTurnLastInputMessage   string             `json:"last_turn_last_input_message"`
	LastTurnLastOutputMessage  string             `json:"last_turn_last_output_message"`
	ServiceAccountID           *string            `json:"service_account_id,omitempty"`
}

func (v GetTrace) MarshalJSON() ([]byte, error) {
	type wire GetTrace
	if v.ServiceAccountID == nil {
		x := string("")
		v.ServiceAccountID = &x
	}
	return json.Marshal(wire(v))
}

type ImageURL struct {
	URL    string       `json:"url"`
	Detail *ImageDetail `json:"detail,omitempty"`
}
type ImageURLChunk struct {
	ImageURL ImageURLUnion `json:"image_url"`
	Type     *string       `json:"type,omitempty"`
}

func (v ImageURLChunk) MarshalJSON() ([]byte, error) {
	type wire ImageURLChunk
	if v.Type == nil {
		x := string("image_url")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ListScorersResponse struct {
	Scorers PaginatedResultScorerSummary `json:"scorers"`
}
type LlmJudgeConfig struct {
	Model          string               `json:"model"`
	PromptTemplate string               `json:"prompt_template"`
	OutputSpec     EvaluationResultSpec `json:"output_spec"`
	Params         *LlmJudgeParams      `json:"params,omitempty"`
}
type LlmJudgeParams struct {
	Temperature *float64 `json:"temperature,omitempty"`
	TopP        *float64 `json:"top_p,omitempty"`
	MaxTokens   *int     `json:"max_tokens,omitempty"`
	RandomSeed  *int     `json:"random_seed,omitempty"`
}
type MessageOutputEntry struct {
	Content     MessageOutputEntryContent `json:"content"`
	Object      *string                   `json:"object,omitempty"`
	Type        *string                   `json:"type,omitempty"`
	CreatedAt   *time.Time                `json:"created_at,omitempty"`
	CompletedAt *time.Time                `json:"completed_at,omitempty"`
	AgentID     *string                   `json:"agent_id,omitempty"`
	Model       *string                   `json:"model,omitempty"`
	ID          *string                   `json:"id,omitempty"`
	Role        *string                   `json:"role,omitempty"`
}

func (v MessageOutputEntry) MarshalJSON() ([]byte, error) {
	type wire MessageOutputEntry
	if v.Object == nil {
		x := string("entry")
		v.Object = &x
	}
	if v.Type == nil {
		x := string("message.output")
		v.Type = &x
	}
	if v.Role == nil {
		x := string("assistant")
		v.Role = &x
	}
	return json.Marshal(wire(v))
}

type MetricDirectionSpec struct {
	Metric    string                       `json:"metric"`
	Direction MetricDirectionSpecDirection `json:"direction"`
}
type MissingGoalItem struct {
	Metric    string   `json:"metric"`
	Operator  string   `json:"operator"`
	Message   string   `json:"message"`
	Target    *float64 `json:"target,omitempty"`
	TargetMin *float64 `json:"target_min,omitempty"`
	TargetMax *float64 `json:"target_max,omitempty"`
}
type ModerationRequest struct {
	Model         string                  `json:"model"`
	Inputs        ModerationRequestInputs `json:"input"`
	Metadata      map[string]any          `json:"metadata,omitempty"`
	Queries       map[string]string       `json:"queries,omitempty"`
	ExplicitNulls []string                `json:"-"`
}

func (v ModerationRequest) MarshalJSON() ([]byte, error) {
	type wire ModerationRequest
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Metadata != nil {
		data, err := json.Marshal(v.Metadata)
		if err != nil {
			return nil, err
		}
		fields["metadata"] = data
	}
	if v.Queries != nil {
		data, err := json.Marshal(v.Queries)
		if err != nil {
			return nil, err
		}
		fields["queries"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "metadata", "queries":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type NetworkEncodedInput struct {
	B64payload      string                  `json:"b64payload"`
	EncodingOptions []EncodedPayloadOptions `json:"encoding_options,omitempty"`
	Empty           *bool                   `json:"empty,omitempty"`
}

func (v NetworkEncodedInput) MarshalJSON() ([]byte, error) {
	type wire NetworkEncodedInput
	if v.Empty == nil {
		x := bool(false)
		v.Empty = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.EncodingOptions != nil {
		data, err := json.Marshal(v.EncodingOptions)
		if err != nil {
			return nil, err
		}
		fields["encoding_options"] = data
	}
	return json.Marshal(fields)
}

type OpenAPIFileSource struct {
	Filename string `json:"filename"`
	Type     string `json:"type,omitempty"`
}

func (v OpenAPIFileSource) MarshalJSON() ([]byte, error) {
	type wire OpenAPIFileSource
	if v.Type == "" {
		v.Type = "file"
	}
	return json.Marshal(wire(v))
}

type OpenAPIURLSource struct {
	URL  string `json:"url"`
	Type string `json:"type,omitempty"`
}

func (v OpenAPIURLSource) MarshalJSON() ([]byte, error) {
	type wire OpenAPIURLSource
	if v.Type == "" {
		v.Type = "url"
	}
	return json.Marshal(wire(v))
}

type OtelEnumFieldDefinition struct {
	Name                  string                                     `json:"name"`
	Label                 string                                     `json:"label"`
	SupportedOperators    []OtelEnumFieldDefinitionSupportedOperator `json:"supported_operators"`
	SupportedAggregations []MetricAggregation                        `json:"supported_aggregations"`
	Type                  string                                     `json:"type,omitempty"`
	Group                 *string                                    `json:"group,omitempty"`
	SourceAttributeKey    *string                                    `json:"source_attribute_key,omitempty"`
	Constraints           *EnumConstraints                           `json:"constraints,omitempty"`
}

func (v OtelEnumFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire OtelEnumFieldDefinition
	if v.Type == "" {
		v.Type = "ENUM"
	}
	return json.Marshal(wire(v))
}

type OtelFloatFieldDefinition struct {
	Name                  string                                      `json:"name"`
	Label                 string                                      `json:"label"`
	SupportedOperators    []OtelFloatFieldDefinitionSupportedOperator `json:"supported_operators"`
	SupportedAggregations []MetricAggregation                         `json:"supported_aggregations"`
	Type                  string                                      `json:"type,omitempty"`
	Group                 *string                                     `json:"group,omitempty"`
	SourceAttributeKey    *string                                     `json:"source_attribute_key,omitempty"`
	Constraints           *FloatConstraints                           `json:"constraints,omitempty"`
}

func (v OtelFloatFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire OtelFloatFieldDefinition
	if v.Type == "" {
		v.Type = "FLOAT"
	}
	return json.Marshal(wire(v))
}

type PaginatedResultScorerSummary struct {
	Count    int             `json:"count"`
	Results  []ScorerSummary `json:"results,omitempty"`
	Next     *string         `json:"next,omitempty"`
	Previous *string         `json:"previous,omitempty"`
}

func (v PaginatedResultScorerSummary) MarshalJSON() ([]byte, error) {
	type wire PaginatedResultScorerSummary
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Results != nil {
		data, err := json.Marshal(v.Results)
		if err != nil {
			return nil, err
		}
		fields["results"] = data
	}
	return json.Marshal(fields)
}

type PromptTokensDetails struct {
	CachedTokens *int `json:"cached_tokens,omitempty"`
	AudioTokens  *int `json:"audio_tokens,omitempty"`
}

func (v PromptTokensDetails) MarshalJSON() ([]byte, error) {
	type wire PromptTokensDetails
	if v.CachedTokens == nil {
		x := int(0)
		v.CachedTokens = &x
	}
	if v.AudioTokens == nil {
		x := int(0)
		v.AudioTokens = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechError struct {
	Error RealtimeSpeechErrorDetail `json:"error"`
	Type  *string                   `json:"type,omitempty"`
}

func (v RealtimeSpeechError) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechError
	if v.Type == nil {
		x := string("error")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechErrorDetail struct {
	Message RealtimeSpeechErrorDetailMessage `json:"message"`
	Code    int                              `json:"code"`
}
type RealtimeSpeechInputTextAppend struct {
	Text string  `json:"text"`
	Type *string `json:"type,omitempty"`
}

func (v RealtimeSpeechInputTextAppend) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechInputTextAppend
	if v.Type == nil {
		x := string("input_text.append")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechInputTextEnd struct {
	Type *string `json:"type,omitempty"`
}

func (v RealtimeSpeechInputTextEnd) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechInputTextEnd
	if v.Type == nil {
		x := string("input_text.end")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechInputTextFlush struct {
	Type *string `json:"type,omitempty"`
	Text *string `json:"text,omitempty"`
}

func (v RealtimeSpeechInputTextFlush) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechInputTextFlush
	if v.Type == nil {
		x := string("input_text.flush")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechSession struct {
	RequestID      string                      `json:"request_id"`
	Model          string                      `json:"model"`
	IsRefAudio     bool                        `json:"is_ref_audio"`
	StreamingMode  RealtimeSpeechStreamingMode `json:"streaming_mode"`
	ResponseFormat SpeechOutputFormat          `json:"response_format"`
	Expressiveness float64                     `json:"expressiveness"`
	Voice          *string                     `json:"voice,omitempty"`
	PhonemeHints   map[string][]IPAPhoneme     `json:"phoneme_hints,omitempty"`
}

func (v RealtimeSpeechSession) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechSession
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.PhonemeHints != nil {
		data, err := json.Marshal(v.PhonemeHints)
		if err != nil {
			return nil, err
		}
		fields["phoneme_hints"] = data
	}
	return json.Marshal(fields)
}

type RealtimeSpeechSessionCreated struct {
	Session RealtimeSpeechSession `json:"session"`
	Type    *string               `json:"type,omitempty"`
}

func (v RealtimeSpeechSessionCreated) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechSessionCreated
	if v.Type == nil {
		x := string("session.created")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechSessionUpdateMessage struct {
	Session RealtimeSpeechSessionUpdatePayload `json:"session"`
	Type    *string                            `json:"type,omitempty"`
}

func (v RealtimeSpeechSessionUpdateMessage) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechSessionUpdateMessage
	if v.Type == nil {
		x := string("session.update")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type RealtimeSpeechSessionUpdatePayload struct {
	Voice          *string                      `json:"voice,omitempty"`
	RefAudio       *string                      `json:"ref_audio,omitempty"`
	StreamingMode  *RealtimeSpeechStreamingMode `json:"streaming_mode,omitempty"`
	ResponseFormat *SpeechOutputFormat          `json:"response_format,omitempty"`
	Expressiveness *float64                     `json:"expressiveness,omitempty"`
	PhonemeHints   map[string][]IPAPhoneme      `json:"phoneme_hints,omitempty"`
	ExplicitNulls  []string                     `json:"-"`
}

func (v RealtimeSpeechSessionUpdatePayload) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechSessionUpdatePayload
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.PhonemeHints != nil {
		data, err := json.Marshal(v.PhonemeHints)
		if err != nil {
			return nil, err
		}
		fields["phoneme_hints"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "voice", "ref_audio", "streaming_mode", "response_format", "expressiveness", "phoneme_hints":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type RealtimeSpeechSessionUpdated struct {
	Session RealtimeSpeechSession `json:"session"`
	Type    *string               `json:"type,omitempty"`
}

func (v RealtimeSpeechSessionUpdated) MarshalJSON() ([]byte, error) {
	type wire RealtimeSpeechSessionUpdated
	if v.Type == nil {
		x := string("session.updated")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ReferenceChunk struct {
	ReferenceIds []ReferenceID `json:"reference_ids"`
	Type         *string       `json:"type,omitempty"`
}

func (v ReferenceChunk) MarshalJSON() ([]byte, error) {
	type wire ReferenceChunk
	if v.Type == nil {
		x := string("reference")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ResponsesMisalignment struct {
	AdditionalProperties map[string]any              `json:"-"`
	DetailedExplanation  *string                     `json:"detailed_explanation,omitempty"`
	ErrorType            *string                     `json:"error_type,omitempty"`
	ReviewTarget         *string                     `json:"review_target,omitempty"`
	Steer                *ResponsesMisalignmentSteer `json:"steer,omitempty"`
}

func (v ResponsesMisalignment) MarshalJSON() ([]byte, error) {
	type wire ResponsesMisalignment
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "detailed_explanation", "error_type", "review_target", "steer":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *ResponsesMisalignment) UnmarshalJSON(data []byte) error {
	type wire ResponsesMisalignment
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "detailed_explanation")
	delete(fields, "error_type")
	delete(fields, "review_target")
	delete(fields, "steer")
	*v = ResponsesMisalignment(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type ResponsesMisalignmentSteer struct {
	AdditionalProperties map[string]any `json:"-"`
	Message              string         `json:"message"`
}

func (v ResponsesMisalignmentSteer) MarshalJSON() ([]byte, error) {
	type wire ResponsesMisalignmentSteer
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "message":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *ResponsesMisalignmentSteer) UnmarshalJSON(data []byte) error {
	type wire ResponsesMisalignmentSteer
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "message")
	*v = ResponsesMisalignmentSteer(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type ScheduleCalendar struct {
	Second     []ScheduleRange `json:"second,omitempty"`
	Minute     []ScheduleRange `json:"minute,omitempty"`
	Hour       []ScheduleRange `json:"hour,omitempty"`
	DayOfMonth []ScheduleRange `json:"day_of_month,omitempty"`
	Month      []ScheduleRange `json:"month,omitempty"`
	Year       []ScheduleRange `json:"year,omitempty"`
	DayOfWeek  []ScheduleRange `json:"day_of_week,omitempty"`
	Comment    *string         `json:"comment,omitempty"`
}

func (v ScheduleCalendar) MarshalJSON() ([]byte, error) {
	type wire ScheduleCalendar
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Second != nil {
		data, err := json.Marshal(v.Second)
		if err != nil {
			return nil, err
		}
		fields["second"] = data
	}
	if v.Minute != nil {
		data, err := json.Marshal(v.Minute)
		if err != nil {
			return nil, err
		}
		fields["minute"] = data
	}
	if v.Hour != nil {
		data, err := json.Marshal(v.Hour)
		if err != nil {
			return nil, err
		}
		fields["hour"] = data
	}
	if v.DayOfMonth != nil {
		data, err := json.Marshal(v.DayOfMonth)
		if err != nil {
			return nil, err
		}
		fields["day_of_month"] = data
	}
	if v.Month != nil {
		data, err := json.Marshal(v.Month)
		if err != nil {
			return nil, err
		}
		fields["month"] = data
	}
	if v.Year != nil {
		data, err := json.Marshal(v.Year)
		if err != nil {
			return nil, err
		}
		fields["year"] = data
	}
	if v.DayOfWeek != nil {
		data, err := json.Marshal(v.DayOfWeek)
		if err != nil {
			return nil, err
		}
		fields["day_of_week"] = data
	}
	return json.Marshal(fields)
}

type ScheduleDefinitionOutput struct {
	Input               any                       `json:"input"`
	ScheduleID          string                    `json:"schedule_id"`
	WorkflowName        string                    `json:"workflow_name"`
	Paused              bool                      `json:"paused"`
	Calendars           []ScheduleCalendar        `json:"calendars,omitempty"`
	Intervals           []ScheduleInterval        `json:"intervals,omitempty"`
	CronExpressions     []string                  `json:"cron_expressions,omitempty"`
	Skip                []ScheduleCalendar        `json:"skip,omitempty"`
	StartAt             *time.Time                `json:"start_at,omitempty"`
	EndAt               *time.Time                `json:"end_at,omitempty"`
	Jitter              *string                   `json:"jitter,omitempty"`
	TimeZoneName        *string                   `json:"time_zone_name,omitempty"`
	Policy              *SchedulePolicy           `json:"policy,omitempty"`
	RemainingExecutions *int                      `json:"remaining_executions,omitempty"`
	DeploymentName      *string                   `json:"deployment_name,omitempty"`
	UserID              *string                   `json:"user_id,omitempty"`
	Note                *string                   `json:"note,omitempty"`
	FutureExecutions    []ScheduleFutureExecution `json:"future_executions,omitempty"`
	RecentExecutions    []ScheduleRecentExecution `json:"recent_executions,omitempty"`
}

func (v ScheduleDefinitionOutput) MarshalJSON() ([]byte, error) {
	type wire ScheduleDefinitionOutput
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Calendars != nil {
		data, err := json.Marshal(v.Calendars)
		if err != nil {
			return nil, err
		}
		fields["calendars"] = data
	}
	if v.Intervals != nil {
		data, err := json.Marshal(v.Intervals)
		if err != nil {
			return nil, err
		}
		fields["intervals"] = data
	}
	if v.CronExpressions != nil {
		data, err := json.Marshal(v.CronExpressions)
		if err != nil {
			return nil, err
		}
		fields["cron_expressions"] = data
	}
	if v.Skip != nil {
		data, err := json.Marshal(v.Skip)
		if err != nil {
			return nil, err
		}
		fields["skip"] = data
	}
	if v.FutureExecutions != nil {
		data, err := json.Marshal(v.FutureExecutions)
		if err != nil {
			return nil, err
		}
		fields["future_executions"] = data
	}
	if v.RecentExecutions != nil {
		data, err := json.Marshal(v.RecentExecutions)
		if err != nil {
			return nil, err
		}
		fields["recent_executions"] = data
	}
	return json.Marshal(fields)
}

type ScheduleFutureExecution struct {
	ScheduledAt time.Time `json:"scheduled_at"`
}
type ScheduleInterval struct {
	Every  string  `json:"every"`
	Offset *string `json:"offset,omitempty"`
}
type SchedulePolicy struct {
	CatchupWindowSeconds *int                   `json:"catchup_window_seconds,omitempty"`
	Overlap              *ScheduleOverlapPolicy `json:"overlap,omitempty"`
	PauseOnFailure       *bool                  `json:"pause_on_failure,omitempty"`
}

func (v SchedulePolicy) MarshalJSON() ([]byte, error) {
	type wire SchedulePolicy
	if v.CatchupWindowSeconds == nil {
		x := int(31536000)
		v.CatchupWindowSeconds = &x
	}
	if v.PauseOnFailure == nil {
		x := bool(false)
		v.PauseOnFailure = &x
	}
	return json.Marshal(wire(v))
}

type ScheduleRange struct {
	Start int  `json:"start"`
	End   *int `json:"end,omitempty"`
	Step  *int `json:"step,omitempty"`
}

func (v ScheduleRange) MarshalJSON() ([]byte, error) {
	type wire ScheduleRange
	if v.End == nil {
		x := int(0)
		v.End = &x
	}
	if v.Step == nil {
		x := int(0)
		v.Step = &x
	}
	return json.Marshal(wire(v))
}

type ScheduleRecentExecution struct {
	ScheduledAt time.Time `json:"scheduled_at"`
	StartedAt   time.Time `json:"started_at"`
	ExecutionID string    `json:"execution_id"`
}
type ScoreMetric struct {
	Name  string           `json:"name"`
	Value ScoreMetricValue `json:"value"`
}
type Scorer struct {
	Slug          string        `json:"slug"`
	Name          string        `json:"name"`
	Description   *string       `json:"description"`
	Kind          ScorerKind    `json:"kind"`
	LatestVersion ScorerVersion `json:"latest_version"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
type ScorerSummary struct {
	Slug        string         `json:"slug"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	Kind        ScorerKind     `json:"kind"`
	IsOotb      bool           `json:"is_ootb"`
	CreatorID   string         `json:"creator_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	CreatorType *PrincipalType `json:"creator_type,omitempty"`
}
type ScorerVersion struct {
	Version     int            `json:"version"`
	Config      LlmJudgeConfig `json:"config"`
	Variables   []string       `json:"variables"`
	CreatorID   string         `json:"creator_id"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatorType *PrincipalType `json:"creator_type,omitempty"`
}
type SignalInvocationBody struct {
	Name        string                                     `json:"name"`
	Input       *SignalInvocationBodyInput                 `json:"input,omitempty"`
	CallContext map[string]SignalInvocationBodyCallContext `json:"call_context,omitempty"`
}

func (v SignalInvocationBody) MarshalJSON() ([]byte, error) {
	type wire SignalInvocationBody
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.CallContext != nil {
		data, err := json.Marshal(v.CallContext)
		if err != nil {
			return nil, err
		}
		fields["call_context"] = data
	}
	return json.Marshal(fields)
}

type SignalInvocationBodyNetworkEncodedInput struct {
	AdditionalProperties map[string]any          `json:"-"`
	B64payload           string                  `json:"b64payload"`
	EncodingOptions      []EncodedPayloadOptions `json:"encoding_options,omitempty"`
	Empty                *bool                   `json:"empty,omitempty"`
}

func (v SignalInvocationBodyNetworkEncodedInput) MarshalJSON() ([]byte, error) {
	type wire SignalInvocationBodyNetworkEncodedInput
	if v.Empty == nil {
		x := bool(false)
		v.Empty = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.EncodingOptions != nil {
		data, err := json.Marshal(v.EncodingOptions)
		if err != nil {
			return nil, err
		}
		fields["encoding_options"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "b64payload", "encoding_options", "empty":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *SignalInvocationBodyNetworkEncodedInput) UnmarshalJSON(data []byte) error {
	type wire SignalInvocationBodyNetworkEncodedInput
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "b64payload")
	delete(fields, "encoding_options")
	delete(fields, "empty")
	*v = SignalInvocationBodyNetworkEncodedInput(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type SpeechStreamDone struct {
	Usage UsageInfo `json:"usage"`
	Type  string    `json:"type,omitempty"`
}

func (v SpeechStreamDone) MarshalJSON() ([]byte, error) {
	type wire SpeechStreamDone
	if v.Type == "" {
		v.Type = "speech.audio.done"
	}
	return json.Marshal(wire(v))
}

type TextChunk struct {
	Text string  `json:"text"`
	Type *string `json:"type,omitempty"`
}

func (v TextChunk) MarshalJSON() ([]byte, error) {
	type wire TextChunk
	if v.Type == nil {
		x := string("text")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ThinkChunk struct {
	Thinking  []Thinking `json:"thinking"`
	Type      *string    `json:"type,omitempty"`
	Signature *string    `json:"signature,omitempty"`
	Closed    *bool      `json:"closed,omitempty"`
}

func (v ThinkChunk) MarshalJSON() ([]byte, error) {
	type wire ThinkChunk
	if v.Type == nil {
		x := string("thinking")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ToolConfiguration struct {
	Exclude              []string `json:"exclude,omitempty"`
	Include              []string `json:"include,omitempty"`
	RequiresConfirmation []string `json:"requires_confirmation,omitempty"`
}

func (v ToolConfiguration) MarshalJSON() ([]byte, error) {
	type wire ToolConfiguration
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Exclude != nil {
		data, err := json.Marshal(v.Exclude)
		if err != nil {
			return nil, err
		}
		fields["exclude"] = data
	}
	if v.Include != nil {
		data, err := json.Marshal(v.Include)
		if err != nil {
			return nil, err
		}
		fields["include"] = data
	}
	if v.RequiresConfirmation != nil {
		data, err := json.Marshal(v.RequiresConfirmation)
		if err != nil {
			return nil, err
		}
		fields["requires_confirmation"] = data
	}
	return json.Marshal(fields)
}

type ToolExecutionEntry struct {
	Name        ToolExecutionEntryName `json:"name"`
	Arguments   string                 `json:"arguments"`
	Object      *string                `json:"object,omitempty"`
	Type        *string                `json:"type,omitempty"`
	CreatedAt   *time.Time             `json:"created_at,omitempty"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
	AgentID     *string                `json:"agent_id,omitempty"`
	Model       *string                `json:"model,omitempty"`
	ID          *string                `json:"id,omitempty"`
	Info        map[string]any         `json:"info,omitempty"`
}

func (v ToolExecutionEntry) MarshalJSON() ([]byte, error) {
	type wire ToolExecutionEntry
	if v.Object == nil {
		x := string("entry")
		v.Object = &x
	}
	if v.Type == nil {
		x := string("tool.execution")
		v.Type = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Info != nil {
		data, err := json.Marshal(v.Info)
		if err != nil {
			return nil, err
		}
		fields["info"] = data
	}
	return json.Marshal(fields)
}

type ToolFileChunk struct {
	Tool     ToolFileChunkTool `json:"tool"`
	FileID   string            `json:"file_id"`
	Type     *string           `json:"type,omitempty"`
	FileName *string           `json:"file_name,omitempty"`
	FileType *string           `json:"file_type,omitempty"`
}

func (v ToolFileChunk) MarshalJSON() ([]byte, error) {
	type wire ToolFileChunk
	if v.Type == nil {
		x := string("tool_file")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ToolReferenceChunk struct {
	Tool        ToolReferenceChunkTool `json:"tool"`
	Title       string                 `json:"title"`
	Type        *string                `json:"type,omitempty"`
	URL         *string                `json:"url,omitempty"`
	Favicon     *string                `json:"favicon,omitempty"`
	Description *string                `json:"description,omitempty"`
}

func (v ToolReferenceChunk) MarshalJSON() ([]byte, error) {
	type wire ToolReferenceChunk
	if v.Type == nil {
		x := string("tool_reference")
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type TranscriptionSegmentChunk struct {
	AdditionalProperties map[string]any `json:"-"`
	Text                 string         `json:"text"`
	Start                *float64       `json:"start"`
	End                  *float64       `json:"end"`
	Type                 *string        `json:"type,omitempty"`
	Score                *float64       `json:"score,omitempty"`
	SpeakerID            *string        `json:"speaker_id,omitempty"`
}

func (v TranscriptionSegmentChunk) MarshalJSON() ([]byte, error) {
	type wire TranscriptionSegmentChunk
	if v.Type == nil {
		x := string("transcription_segment")
		v.Type = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "text", "start", "end", "type", "score", "speaker_id":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *TranscriptionSegmentChunk) UnmarshalJSON(data []byte) error {
	type wire TranscriptionSegmentChunk
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "text")
	delete(fields, "start")
	delete(fields, "end")
	delete(fields, "type")
	delete(fields, "score")
	delete(fields, "speaker_id")
	*v = TranscriptionSegmentChunk(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type TranscriptionStreamDone struct {
	AdditionalProperties map[string]any              `json:"-"`
	Model                string                      `json:"model"`
	Text                 string                      `json:"text"`
	Usage                UsageInfo                   `json:"usage"`
	Language             *string                     `json:"language"`
	Segments             []TranscriptionSegmentChunk `json:"segments,omitempty"`
	Type                 string                      `json:"type,omitempty"`
}

func (v TranscriptionStreamDone) MarshalJSON() ([]byte, error) {
	type wire TranscriptionStreamDone
	if v.Type == "" {
		v.Type = "transcription.done"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Segments != nil {
		data, err := json.Marshal(v.Segments)
		if err != nil {
			return nil, err
		}
		fields["segments"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "model", "text", "usage", "language", "segments", "type":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *TranscriptionStreamDone) UnmarshalJSON(data []byte) error {
	type wire TranscriptionStreamDone
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "model")
	delete(fields, "text")
	delete(fields, "usage")
	delete(fields, "language")
	delete(fields, "segments")
	delete(fields, "type")
	*v = TranscriptionStreamDone(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type TranscriptionStreamLanguage struct {
	AdditionalProperties map[string]any `json:"-"`
	AudioLanguage        string         `json:"audio_language"`
	Type                 string         `json:"type,omitempty"`
}

func (v TranscriptionStreamLanguage) MarshalJSON() ([]byte, error) {
	type wire TranscriptionStreamLanguage
	if v.Type == "" {
		v.Type = "transcription.language"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "audio_language", "type":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *TranscriptionStreamLanguage) UnmarshalJSON(data []byte) error {
	type wire TranscriptionStreamLanguage
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "audio_language")
	delete(fields, "type")
	*v = TranscriptionStreamLanguage(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type TranscriptionStreamSegmentDelta struct {
	AdditionalProperties map[string]any `json:"-"`
	Text                 string         `json:"text"`
	Start                *float64       `json:"start"`
	End                  *float64       `json:"end"`
	Type                 string         `json:"type,omitempty"`
	SpeakerID            *string        `json:"speaker_id,omitempty"`
}

func (v TranscriptionStreamSegmentDelta) MarshalJSON() ([]byte, error) {
	type wire TranscriptionStreamSegmentDelta
	if v.Type == "" {
		v.Type = "transcription.segment"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "text", "start", "end", "type", "speaker_id":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *TranscriptionStreamSegmentDelta) UnmarshalJSON(data []byte) error {
	type wire TranscriptionStreamSegmentDelta
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "text")
	delete(fields, "start")
	delete(fields, "end")
	delete(fields, "type")
	delete(fields, "speaker_id")
	*v = TranscriptionStreamSegmentDelta(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type TranscriptionStreamTextDelta struct {
	AdditionalProperties map[string]any `json:"-"`
	Text                 string         `json:"text"`
	Type                 string         `json:"type,omitempty"`
}

func (v TranscriptionStreamTextDelta) MarshalJSON() ([]byte, error) {
	type wire TranscriptionStreamTextDelta
	if v.Type == "" {
		v.Type = "transcription.text.delta"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "text", "type":
			continue
		}
		data, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
func (v *TranscriptionStreamTextDelta) UnmarshalJSON(data []byte) error {
	type wire TranscriptionStreamTextDelta
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "text")
	delete(fields, "type")
	*v = TranscriptionStreamTextDelta(decoded)
	if len(fields) > 0 {
		v.AdditionalProperties = map[string]any{}
		for key, raw := range fields {
			var value any
			if err := json.Unmarshal(raw, &value); err != nil {
				return err
			}
			v.AdditionalProperties[key] = value
		}
	}
	return nil
}

type Trigger struct {
	ID             string            `json:"id"`
	CreatorID      string            `json:"creator_id"`
	WorkspaceID    string            `json:"workspace_id"`
	OrganizationID string            `json:"organization_id"`
	CustomerID     string            `json:"customer_id"`
	OwnerType      TriggersOwnerType `json:"owner_type"`
	EventSource    Webhook           `json:"event_source"`
	Handler        TriggerHandler    `json:"handler"`
	Action         TriggerAction     `json:"action"`
	Filters        *string           `json:"filters"`
	Title          string            `json:"title"`
}
type TriggerAction struct {
	PrimitiveID string            `json:"primitive_id"`
	Type        TriggerActionType `json:"type"`
}
type TriggerEvent struct {
	EventID     string             `json:"event_id"`
	TriggerID   string             `json:"trigger_id"`
	SourceID    string             `json:"source_id"`
	Status      TriggerStatusValue `json:"status"`
	ReceivedAt  time.Time          `json:"received_at"`
	SubmittedAt *time.Time         `json:"submitted_at"`
	TriggeredAt *time.Time         `json:"triggered_at"`
	Detail      *string            `json:"detail"`
	URL         *string            `json:"url"`
}
type TriggerEventList struct {
	Data []TriggerEvent `json:"data"`
}
type UnknownBackendSpec struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownBackendSpec) MarshalJSON() ([]byte, error) {
	type wire UnknownBackendSpec
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownFallback struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownFallback) MarshalJSON() ([]byte, error) {
	type wire UnknownFallback
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownFieldDefinition struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownFieldDefinition) MarshalJSON() ([]byte, error) {
	type wire UnknownFieldDefinition
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownHangUpToolFiller struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownHangUpToolFiller) MarshalJSON() ([]byte, error) {
	type wire UnknownHangUpToolFiller
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownOpenAPISpecSource struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownOpenAPISpecSource) MarshalJSON() ([]byte, error) {
	type wire UnknownOpenAPISpecSource
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownTranscriptionStreamEventsData struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownTranscriptionStreamEventsData) MarshalJSON() ([]byte, error) {
	type wire UnknownTranscriptionStreamEventsData
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownTransferCallToolFiller struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownTransferCallToolFiller) MarshalJSON() ([]byte, error) {
	type wire UnknownTransferCallToolFiller
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownVoiceAgentCustomConnectorFiller struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownVoiceAgentCustomConnectorFiller) MarshalJSON() ([]byte, error) {
	type wire UnknownVoiceAgentCustomConnectorFiller
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownVoiceAgentDocumentLibraryToolFiller struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownVoiceAgentDocumentLibraryToolFiller) MarshalJSON() ([]byte, error) {
	type wire UnknownVoiceAgentDocumentLibraryToolFiller
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownVoiceAgentFunctionToolFiller struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownVoiceAgentFunctionToolFiller) MarshalJSON() ([]byte, error) {
	type wire UnknownVoiceAgentFunctionToolFiller
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UnknownVoiceAgentWebSearchToolFiller struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownVoiceAgentWebSearchToolFiller) MarshalJSON() ([]byte, error) {
	type wire UnknownVoiceAgentWebSearchToolFiller
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type UpdateInvocationBody struct {
	Name        string                                     `json:"name"`
	Input       *UpdateInvocationBodyInput                 `json:"input,omitempty"`
	CallContext map[string]UpdateInvocationBodyCallContext `json:"call_context,omitempty"`
}

func (v UpdateInvocationBody) MarshalJSON() ([]byte, error) {
	type wire UpdateInvocationBody
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.CallContext != nil {
		data, err := json.Marshal(v.CallContext)
		if err != nil {
			return nil, err
		}
		fields["call_context"] = data
	}
	return json.Marshal(fields)
}

type UserIdentityWorkspace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type VoiceCreateRequest struct {
	Name            string       `json:"name"`
	Languages       []string     `json:"languages,omitempty"`
	Gender          *VoiceGender `json:"gender,omitempty"`
	Age             *int         `json:"age,omitempty"`
	Tags            []string     `json:"tags,omitempty"`
	Color           *string      `json:"color,omitempty"`
	Description     *string      `json:"description,omitempty"`
	RetentionNotice *int         `json:"retention_notice,omitempty"`
	SampleAudio     *string      `json:"sample_audio,omitempty"`
	SampleFilename  *string      `json:"sample_filename,omitempty"`
	ExplicitNulls   []string     `json:"-"`
}

func (v VoiceCreateRequest) MarshalJSON() ([]byte, error) {
	type wire VoiceCreateRequest
	if v.RetentionNotice == nil {
		x := int(30)
		v.RetentionNotice = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Languages != nil {
		data, err := json.Marshal(v.Languages)
		if err != nil {
			return nil, err
		}
		fields["languages"] = data
	}
	if v.Tags != nil {
		data, err := json.Marshal(v.Tags)
		if err != nil {
			return nil, err
		}
		fields["tags"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "gender", "age", "tags", "color", "description", "sample_audio", "sample_filename":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type Webhook struct {
	ID                      string                         `json:"id"`
	CreatorID               string                         `json:"creator_id"`
	OrganizationID          string                         `json:"organization_id"`
	WorkspaceID             string                         `json:"workspace_id"`
	OwnerType               TriggersOwnerType              `json:"owner_type"`
	Provider                WebhookProvider                `json:"provider"`
	Name                    string                         `json:"name"`
	URL                     string                         `json:"url"`
	HttpMethod              string                         `json:"http_method"`
	URLVerified             *bool                          `json:"url_verified"`
	SignatureHashAlgorithm  *WebhookSignatureHashAlgorithm `json:"signature_hash_algorithm"`
	SignatureHeaderName     *string                        `json:"signature_header_name"`
	ValidityPeriodSecond    *int                           `json:"validity_period_second"`
	SignatureExtractorRegex *string                        `json:"signature_extractor_regex"`
	TimestampExtractorRegex *string                        `json:"timestamp_extractor_regex"`
	SignaturePayloadPattern *string                        `json:"signature_payload_pattern"`
	SignatureEncoding       *WebhookSignatureEncoding      `json:"signature_encoding"`
	SuccessStatusCode       *int                           `json:"success_status_code"`
	ExpiresAt               *time.Time                     `json:"expires_at"`
	CreatedAt               time.Time                      `json:"created_at"`
	UpdatedAt               time.Time                      `json:"updated_at"`
	Description             *string                        `json:"description,omitempty"`
	URLVerificationToken    *string                        `json:"url_verification_token,omitempty"`
	EventTypePath           *string                        `json:"event_type_path,omitempty"`
	EventTypePattern        *string                        `json:"event_type_pattern,omitempty"`
	EventIDPattern          *string                        `json:"event_id_pattern,omitempty"`
	TimestampPattern        *string                        `json:"timestamp_pattern,omitempty"`
}
type WebhookSigningSecret struct {
	SigningSecret string `json:"signing_secret"`
}
type WorkflowExecutionRequest struct {
	ExecutionID             *string                                        `json:"execution_id,omitempty"`
	Input                   any                                            `json:"input,omitempty"`
	WaitForResult           *bool                                          `json:"wait_for_result,omitempty"`
	TimeoutSeconds          *float64                                       `json:"timeout_seconds,omitempty"`
	CustomTracingAttributes map[string]string                              `json:"custom_tracing_attributes,omitempty"`
	ForceNewTrace           *bool                                          `json:"force_new_trace,omitempty"`
	Traceparent             *string                                        `json:"traceparent,omitempty"`
	Extensions              map[string]any                                 `json:"extensions,omitempty"`
	CallContext             map[string]WorkflowExecutionRequestCallContext `json:"call_context,omitempty"`
	TaskQueue               *string                                        `json:"task_queue,omitempty"`
	DeploymentName          *string                                        `json:"deployment_name,omitempty"`
	ExplicitNulls           []string                                       `json:"-"`
}

func (v WorkflowExecutionRequest) MarshalJSON() ([]byte, error) {
	type wire WorkflowExecutionRequest
	if v.WaitForResult == nil {
		x := bool(false)
		v.WaitForResult = &x
	}
	if v.ForceNewTrace == nil {
		x := bool(false)
		v.ForceNewTrace = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.CustomTracingAttributes != nil {
		data, err := json.Marshal(v.CustomTracingAttributes)
		if err != nil {
			return nil, err
		}
		fields["custom_tracing_attributes"] = data
	}
	if v.Extensions != nil {
		data, err := json.Marshal(v.Extensions)
		if err != nil {
			return nil, err
		}
		fields["extensions"] = data
	}
	if v.CallContext != nil {
		data, err := json.Marshal(v.CallContext)
		if err != nil {
			return nil, err
		}
		fields["call_context"] = data
	}
	for _, key := range v.ExplicitNulls {
		switch key {
		case "execution_id", "input", "timeout_seconds", "custom_tracing_attributes", "traceparent", "extensions", "call_context", "task_queue", "deployment_name":
			fields[key] = json.RawMessage("null")
		default:
			return nil, fmt.Errorf("field %q does not accept explicit null", key)
		}
	}
	return json.Marshal(fields)
}

type WorkflowTagsAddRequest struct {
	Tags []string `json:"tags"`
}
type WorkflowsWorkerSpecUpdate struct {
	GithubURL      *string                            `json:"github_url,omitempty"`
	RepositoryURL  *string                            `json:"repository_url,omitempty"`
	Revision       *string                            `json:"revision,omitempty"`
	ArchiveID      *string                            `json:"archive_id,omitempty"`
	BackendSpec    *DeploymentMistralCloudBackendSpec `json:"backend_spec,omitempty"`
	RequestedRoles []string                           `json:"requested_roles,omitempty"`
}

func (v WorkflowsWorkerSpecUpdate) MarshalJSON() ([]byte, error) {
	type wire WorkflowsWorkerSpecUpdate
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.RequestedRoles != nil {
		data, err := json.Marshal(v.RequestedRoles)
		if err != nil {
			return nil, err
		}
		fields["requested_roles"] = data
	}
	return json.Marshal(fields)
}
