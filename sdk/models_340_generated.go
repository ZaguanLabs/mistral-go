// Code generated from official Mistral Python SDK 3.4.0 models. DO NOT EDIT.
package sdk

import (
	"encoding/json"
	"time"
)

type AgentsCompletionRequestStop = any
type AgentsCompletionRequestMessage = any
type AgentsCompletionRequestTool = any
type AgentsCompletionRequestToolChoice = any
type AgentsCompletionStreamRequestStop = any
type AgentsCompletionStreamRequestMessage = any
type AgentsCompletionStreamRequestTool = any
type AgentsCompletionStreamRequestToolChoice = any
type Audience = string
type AudioAgentConfigurationRequestTool = any
type AudioAgentConfigurationResponseConfigTool = any
type AudioSessionCreateRequestTool = any
type AudioSessionResponseStatus = string
type Transcript = any
type AudioToolTranscriptTurnStatus = string
type ChatCompletionRequestStop = any
type ChatCompletionRequestMessage = any
type ChatCompletionRequestTool = any
type ChatCompletionRequestToolChoice = any
type ChatCompletionStreamRequestStop = any
type ChatCompletionStreamRequestMessage = any
type ChatCompletionStreamRequestTool = any
type ChatCompletionStreamRequestToolChoice = any
type ConnectorToolCallResponseContent = any
type CountResponseInputTokensRequestInput1 = any
type CountResponseInputTokensRequestInput2 = any
type CountResponseInputTokensRequestTool = any
type CountResponseInputTokensRequestToolChoiceEnum = string
type CountResponseInputTokensRequestToolChoiceUnion = any
type CountResponseInputTokensRequestEffort = string
type CountResponseInputTokensRequestFormat = any
type CountResponseInputTokensRequestVerbosity = string
type CreateConnectorRequestAuthMethod = any
type CreateHTTPConnectorRequestAuthMethod = any
type CreateResponseRequestInput1 = any
type CreateResponseRequestInput2 = any
type CreateResponseRequestTool = any
type CreateResponseRequestToolChoiceEnum = string
type CreateResponseRequestToolChoiceUnion = any
type CreateResponseRequestEffort = string
type CreateResponseRequestFormat = any
type CreateResponseRequestVerbosity = string
type Code = string
type Reason = string
type CreateResponseResponseOutput = any
type CreateResponseResponseToolChoiceEnum = string
type CreateResponseResponseToolChoiceUnion = any
type CreateResponseResponseTool = any
type PromptCacheRetention = string
type ResponsesSupportedReasoningEffort = string
type ServiceTier = string
type CreateResponseResponseStatus = string
type ResponsesTextConfigFormat = any
type ResponsesTextConfigVerbosity = string
type Truncation = string
type CustomConnectorAuthorization = any
type Resource = any
type Operator = string
type KeywordRetrieverFilter = any
type Resources = any
type Tools = any
type Prompts = any
type Theme = string
type MessageInputEntryRole = string
type MessageInputEntryContent = any
type FormatEnum = string
type ResponsesContentPartAddedEventPart = any
type ResponsesContentPartDoneEventPart = any
type ResponsesCustomToolAllowedCaller = string
type Syntax = string
type ResponsesCustomToolFormat = any
type ResponsesEasyInputMessageContent1 = any
type ResponsesEasyInputMessageContent2 = any
type ResponsesEasyInputMessageRole = string
type ResponsesEasyInputMessagePhase = string
type ResponsesFunctionCallOutputOutput1 = any
type ResponsesFunctionCallOutputOutput2 = any
type ResponsesFunctionCallOutputCaller = any
type ResponsesFunctionCallOutputStatus = string
type ResponsesFunctionToolAllowedCaller = string
type ResponsesFunctionToolCallCaller = any
type ResponsesFunctionToolCallStatus = string
type ResponsesInputFileDetail = string
type ResponsesInputFileContentDetail = string
type ResponsesInputImageDetail = string
type ResponsesInputImageContentDetail = string
type ResponsesMessageContent = any
type ResponsesMessageRole = string
type ResponsesMessageStatus = string
type ResponsesNamespaceToolTool = any
type ResponsesOutputItemAddedEventItem = any
type ResponsesOutputItemDoneEventItem = any
type ResponsesOutputMessageContent = any
type ResponsesOutputMessageStatus = string
type ResponsesOutputMessagePhase = string
type Annotation = any
type ResponsesReasoningItemStatus = string
type ResponsesStreamFrameData = any
type ResponsesToolFunctionAllowedCaller = string
type ConversationHistoryUnion = any
type SummaryStatus = string
type SummaryStatusDollarDefs = string
type TaskSupport = string
type Destination = any
type TurnDetectionEndReason = string
type UpdateSummaryRequestSummaryStatus = string
type VoiceAgentCustomConnectorAuthorization = any
type VoiceAgentErrorEventSource = string
type VoiceAgentEvent = any
type VoiceAgentFunctionToolType = string
type OldState = string
type State = string
type VoiceAgentToolCallEndedEventStatus = string
type VoiceAgentTranscriptEventRole = string
type AgentsCompletionRequest struct {
	Messages          []AgentsCompletionRequestMessage   `json:"messages"`
	AgentID           string                             `json:"agent_id"`
	MaxTokens         *int                               `json:"max_tokens,omitempty"`
	Stream            *bool                              `json:"stream,omitempty"`
	Stop              *AgentsCompletionRequestStop       `json:"stop,omitempty"`
	RandomSeed        *int                               `json:"random_seed,omitempty"`
	Metadata          map[string]any                     `json:"metadata,omitempty"`
	ResponseFormat    *ResponseFormat                    `json:"response_format,omitempty"`
	Tools             []AgentsCompletionRequestTool      `json:"tools,omitempty"`
	ToolChoice        *AgentsCompletionRequestToolChoice `json:"tool_choice,omitempty"`
	PresencePenalty   *float64                           `json:"presence_penalty,omitempty"`
	FrequencyPenalty  *float64                           `json:"frequency_penalty,omitempty"`
	MinTokens         *int                               `json:"min_tokens,omitempty"`
	RepetitionPenalty *float64                           `json:"repetition_penalty,omitempty"`
	TopK              *int                               `json:"top_k,omitempty"`
	Logprobs          *bool                              `json:"logprobs,omitempty"`
	TopLogprobs       *int                               `json:"top_logprobs,omitempty"`
	PromptLogprobs    *bool                              `json:"prompt_logprobs,omitempty"`
	TopPromptLogprobs *int                               `json:"top_prompt_logprobs,omitempty"`
	N                 *int                               `json:"n,omitempty"`
	Prediction        *Prediction                        `json:"prediction,omitempty"`
	ParallelToolCalls *bool                              `json:"parallel_tool_calls,omitempty"`
	ReasoningEffort   *ReasoningEffort                   `json:"reasoning_effort,omitempty"`
	PromptMode        *MistralPromptMode                 `json:"prompt_mode,omitempty"`
	Guardrails        []GuardrailConfig                  `json:"guardrails,omitempty"`
	PromptCacheKey    *string                            `json:"prompt_cache_key,omitempty"`
	ServiceTier       *RequestedServiceTier              `json:"service_tier,omitempty"`
}

func (v AgentsCompletionRequest) MarshalJSON() ([]byte, error) {
	type wire AgentsCompletionRequest
	if v.Stream == nil {
		x := false
		v.Stream = &x
	}
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
	if v.Tools != nil {
		data, err := json.Marshal(v.Tools)
		if err != nil {
			return nil, err
		}
		fields["tools"] = data
	}
	if v.Guardrails != nil {
		data, err := json.Marshal(v.Guardrails)
		if err != nil {
			return nil, err
		}
		fields["guardrails"] = data
	}
	return json.Marshal(fields)
}

type AgentsCompletionStreamRequest struct {
	Messages          []AgentsCompletionStreamRequestMessage   `json:"messages"`
	AgentID           string                                   `json:"agent_id"`
	MaxTokens         *int                                     `json:"max_tokens,omitempty"`
	Stream            *bool                                    `json:"stream,omitempty"`
	Stop              *AgentsCompletionStreamRequestStop       `json:"stop,omitempty"`
	RandomSeed        *int                                     `json:"random_seed,omitempty"`
	Metadata          map[string]any                           `json:"metadata,omitempty"`
	ResponseFormat    *ResponseFormat                          `json:"response_format,omitempty"`
	Tools             []AgentsCompletionStreamRequestTool      `json:"tools,omitempty"`
	ToolChoice        *AgentsCompletionStreamRequestToolChoice `json:"tool_choice,omitempty"`
	PresencePenalty   *float64                                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty  *float64                                 `json:"frequency_penalty,omitempty"`
	MinTokens         *int                                     `json:"min_tokens,omitempty"`
	RepetitionPenalty *float64                                 `json:"repetition_penalty,omitempty"`
	TopK              *int                                     `json:"top_k,omitempty"`
	Logprobs          *bool                                    `json:"logprobs,omitempty"`
	TopLogprobs       *int                                     `json:"top_logprobs,omitempty"`
	PromptLogprobs    *bool                                    `json:"prompt_logprobs,omitempty"`
	TopPromptLogprobs *int                                     `json:"top_prompt_logprobs,omitempty"`
	N                 *int                                     `json:"n,omitempty"`
	Prediction        *Prediction                              `json:"prediction,omitempty"`
	ParallelToolCalls *bool                                    `json:"parallel_tool_calls,omitempty"`
	ReasoningEffort   *ReasoningEffort                         `json:"reasoning_effort,omitempty"`
	PromptMode        *MistralPromptMode                       `json:"prompt_mode,omitempty"`
	Guardrails        []GuardrailConfig                        `json:"guardrails,omitempty"`
	PromptCacheKey    *string                                  `json:"prompt_cache_key,omitempty"`
	ServiceTier       *RequestedServiceTier                    `json:"service_tier,omitempty"`
}

func (v AgentsCompletionStreamRequest) MarshalJSON() ([]byte, error) {
	type wire AgentsCompletionStreamRequest
	if v.Stream == nil {
		x := true
		v.Stream = &x
	}
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
	if v.Tools != nil {
		data, err := json.Marshal(v.Tools)
		if err != nil {
			return nil, err
		}
		fields["tools"] = data
	}
	if v.Guardrails != nil {
		data, err := json.Marshal(v.Guardrails)
		if err != nil {
			return nil, err
		}
		fields["guardrails"] = data
	}
	return json.Marshal(fields)
}

type Annotations struct {
	Audience     []Audience `json:"audience,omitempty"`
	Priority     *float64   `json:"priority,omitempty"`
	LastModified *string    `json:"lastModified,omitempty"`
}

func (v Annotations) MarshalJSON() ([]byte, error) {
	type wire Annotations
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Audience != nil {
		data, err := json.Marshal(v.Audience)
		if err != nil {
			return nil, err
		}
		fields["audience"] = data
	}
	return json.Marshal(fields)
}

type AudioAgentConfigurationRequest struct {
	Llm            *AudioLLMConfig                      `json:"llm,omitempty"`
	Tts            *AudioTTSConfig                      `json:"tts,omitempty"`
	Stt            *AudioSTTConfig                      `json:"stt,omitempty"`
	Turn           *AudioTurnConfig                     `json:"turn,omitempty"`
	Instructions   *string                              `json:"instructions,omitempty"`
	Tools          []AudioAgentConfigurationRequestTool `json:"tools,omitempty"`
	WelcomeMessage *string                              `json:"welcome_message,omitempty"`
}

func (v AudioAgentConfigurationRequest) MarshalJSON() ([]byte, error) {
	type wire AudioAgentConfigurationRequest
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
	return json.Marshal(fields)
}

type UnknownAudioAgentConfigurationResponseConfigTool struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownAudioAgentConfigurationResponseConfigTool) MarshalJSON() ([]byte, error) {
	type wire UnknownAudioAgentConfigurationResponseConfigTool
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type AudioAgentConfigurationResponseConfig struct {
	Llm            AudioLLMConfig                              `json:"llm"`
	Tts            AudioTTSConfig                              `json:"tts"`
	Stt            AudioSTTConfig                              `json:"stt"`
	Turn           AudioTurnConfig                             `json:"turn"`
	Instructions   *string                                     `json:"instructions,omitempty"`
	Tools          []AudioAgentConfigurationResponseConfigTool `json:"tools,omitempty"`
	WelcomeMessage *string                                     `json:"welcome_message,omitempty"`
}

func (v AudioAgentConfigurationResponseConfig) MarshalJSON() ([]byte, error) {
	type wire AudioAgentConfigurationResponseConfig
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
	return json.Marshal(fields)
}

type AudioAgentCreateRequest struct {
	Name          string                         `json:"name"`
	Configuration AudioAgentConfigurationRequest `json:"configuration"`
	Description   *string                        `json:"description,omitempty"`
}
type AudioAgentListPage struct {
	Data          []AudioAgentResponse `json:"data"`
	NextPageToken *string              `json:"next_page_token,omitempty"`
}
type AudioAgentResponse struct {
	ID            string                                 `json:"id"`
	OwnerID       string                                 `json:"owner_id"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
	Name          *string                                `json:"name,omitempty"`
	Description   *string                                `json:"description,omitempty"`
	Version       *int                                   `json:"version,omitempty"`
	Configuration *AudioAgentConfigurationResponseConfig `json:"configuration,omitempty"`
}
type AudioAgentSessionListPage struct {
	Data          []AudioAgentSessionSummary `json:"data"`
	NextPageToken *string                    `json:"next_page_token,omitempty"`
}
type AudioAgentSessionSummary struct {
	ID            string     `json:"id"`
	AgentID       string     `json:"agent_id"`
	StartedAt     *time.Time `json:"started_at,omitempty"`
	Duration      *int       `json:"duration,omitempty"`
	PhoneNumberID *string    `json:"phone_number_id,omitempty"`
}
type AudioAgentUpdateRequest struct {
	Name          string                         `json:"name"`
	Configuration AudioAgentConfigurationRequest `json:"configuration"`
	Description   *string                        `json:"description,omitempty"`
}
type AudioAssistantTranscriptTurn struct {
	Timestamp      time.Time                      `json:"timestamp"`
	Latency        AudioLatencyBreakdownAssistant `json:"latency"`
	Content        *string                        `json:"content,omitempty"`
	DurationMs     *float64                       `json:"duration_ms,omitempty"`
	Role           string                         `json:"role,omitempty"`
	IsInterruption *bool                          `json:"is_interruption,omitempty"`
	ToolCalls      []AudioToolCall                `json:"tool_calls,omitempty"`
}

func (v AudioAssistantTranscriptTurn) MarshalJSON() ([]byte, error) {
	type wire AudioAssistantTranscriptTurn
	if v.Role == "" {
		v.Role = "assistant"
	}
	if v.IsInterruption == nil {
		x := false
		v.IsInterruption = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.ToolCalls != nil {
		data, err := json.Marshal(v.ToolCalls)
		if err != nil {
			return nil, err
		}
		fields["tool_calls"] = data
	}
	return json.Marshal(fields)
}

type AudioContent struct {
	Data        string         `json:"data"`
	MimeType    string         `json:"mimeType"`
	Type        string         `json:"type,omitempty"`
	Annotations *Annotations   `json:"annotations,omitempty"`
	Meta        map[string]any `json:"_meta,omitempty"`
}

func (v AudioContent) MarshalJSON() ([]byte, error) {
	type wire AudioContent
	if v.Type == "" {
		v.Type = "audio"
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
	return json.Marshal(fields)
}

type AudioInactivityConfig struct {
	TimeoutMs    int      `json:"timeout_ms"`
	MaxReminders int      `json:"max_reminders"`
	Messages     []string `json:"messages,omitempty"`
}

func (v AudioInactivityConfig) MarshalJSON() ([]byte, error) {
	type wire AudioInactivityConfig
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Messages != nil {
		data, err := json.Marshal(v.Messages)
		if err != nil {
			return nil, err
		}
		fields["messages"] = data
	}
	return json.Marshal(fields)
}

type AudioLatencyBreakdownAssistant struct {
	LlmTtftMs     *float64 `json:"llm_ttft_ms,omitempty"`
	LlmDurationMs *float64 `json:"llm_duration_ms,omitempty"`
	TtsTtfbMs     *float64 `json:"tts_ttfb_ms,omitempty"`
	TtsDurationMs *float64 `json:"tts_duration_ms,omitempty"`
	E2eLatencyMs  *float64 `json:"e2e_latency_ms,omitempty"`
}
type AudioLatencyBreakdownUser struct {
	TranscriptionDelayMs *float64 `json:"transcription_delay_ms,omitempty"`
	EndOfTurnDelayMs     *float64 `json:"end_of_turn_delay_ms,omitempty"`
}
type AudioLLMConfig struct {
	Model *string `json:"model,omitempty"`
}
type AudioSessionCreateRequest struct {
	Llm            *AudioLLMConfig                 `json:"llm,omitempty"`
	Tts            *AudioTTSConfig                 `json:"tts,omitempty"`
	Stt            *AudioSTTConfig                 `json:"stt,omitempty"`
	Turn           *AudioTurnConfig                `json:"turn,omitempty"`
	Instructions   *string                         `json:"instructions,omitempty"`
	Tools          []AudioSessionCreateRequestTool `json:"tools,omitempty"`
	WelcomeMessage *string                         `json:"welcome_message,omitempty"`
	AgentID        *string                         `json:"agent_id,omitempty"`
	AgentVersion   *int                            `json:"agent_version,omitempty"`
}

func (v AudioSessionCreateRequest) MarshalJSON() ([]byte, error) {
	type wire AudioSessionCreateRequest
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
	return json.Marshal(fields)
}

type AudioSessionCreateResponse struct {
	ID               string                   `json:"id"`
	ParticipantToken string                   `json:"participant_token"`
	ServerURL        string                   `json:"server_url"`
	CreatedAt        time.Time                `json:"created_at"`
	ExpiresAt        time.Time                `json:"expires_at"`
	FunctionTools    []VoiceAgentFunctionTool `json:"function_tools,omitempty"`
}

func (v AudioSessionCreateResponse) MarshalJSON() ([]byte, error) {
	type wire AudioSessionCreateResponse
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.FunctionTools != nil {
		data, err := json.Marshal(v.FunctionTools)
		if err != nil {
			return nil, err
		}
		fields["function_tools"] = data
	}
	return json.Marshal(fields)
}

type AudioSessionPhoneCall struct {
	PhoneNumberID string                   `json:"phone_number_id"`
	PhoneNumber   *AudioSessionPhoneNumber `json:"phone_number,omitempty"`
}
type AudioSessionPhoneNumber struct {
	ID         string  `json:"id"`
	Identifier string  `json:"identifier"`
	Name       *string `json:"name,omitempty"`
}
type UnknownTranscript struct {
	Role      string `json:"role,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownTranscript) MarshalJSON() ([]byte, error) {
	type wire UnknownTranscript
	if v.Role == "" {
		v.Role = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type AudioSessionResponse struct {
	ID              string                     `json:"id"`
	Status          AudioSessionResponseStatus `json:"status"`
	NumParticipants int                        `json:"num_participants"`
	CreatedAt       time.Time                  `json:"created_at"`
	PhoneCall       *AudioSessionPhoneCall     `json:"phone_call,omitempty"`
	Transcripts     []Transcript               `json:"transcripts,omitempty"`
}

func (v AudioSessionResponse) MarshalJSON() ([]byte, error) {
	type wire AudioSessionResponse
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Transcripts != nil {
		data, err := json.Marshal(v.Transcripts)
		if err != nil {
			return nil, err
		}
		fields["transcripts"] = data
	}
	return json.Marshal(fields)
}

type AudioSpeechCapabilities struct {
	Modes           []string `json:"modes,omitempty"`
	Instruction     *bool    `json:"instruction,omitempty"`
	ResponseFormats []string `json:"response_formats,omitempty"`
}

func (v AudioSpeechCapabilities) MarshalJSON() ([]byte, error) {
	type wire AudioSpeechCapabilities
	if v.Instruction == nil {
		x := false
		v.Instruction = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Modes != nil {
		data, err := json.Marshal(v.Modes)
		if err != nil {
			return nil, err
		}
		fields["modes"] = data
	}
	if v.ResponseFormats != nil {
		data, err := json.Marshal(v.ResponseFormats)
		if err != nil {
			return nil, err
		}
		fields["response_formats"] = data
	}
	return json.Marshal(fields)
}

type AudioSTTConfig struct {
	Model                  *string                   `json:"model,omitempty"`
	Language               *string                   `json:"language,omitempty"`
	ContextBias            []string                  `json:"context_bias,omitempty"`
	TargetStreamingDelayMs *int                      `json:"target_streaming_delay_ms,omitempty"`
	TurnDetection          *AudioTurnDetectionConfig `json:"turn_detection,omitempty"`
}

func (v AudioSTTConfig) MarshalJSON() ([]byte, error) {
	type wire AudioSTTConfig
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.ContextBias != nil {
		data, err := json.Marshal(v.ContextBias)
		if err != nil {
			return nil, err
		}
		fields["context_bias"] = data
	}
	return json.Marshal(fields)
}

type AudioToolCall struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Arguments *string `json:"arguments,omitempty"`
}
type AudioToolTranscriptTurn struct {
	ToolCallID string                        `json:"tool_call_id"`
	Name       string                        `json:"name"`
	Status     AudioToolTranscriptTurnStatus `json:"status"`
	Timestamp  time.Time                     `json:"timestamp"`
	Role       string                        `json:"role,omitempty"`
	Result     any                           `json:"result,omitempty"`
	DurationMs *float64                      `json:"duration_ms,omitempty"`
}

func (v AudioToolTranscriptTurn) MarshalJSON() ([]byte, error) {
	type wire AudioToolTranscriptTurn
	if v.Role == "" {
		v.Role = "tool"
	}
	return json.Marshal(wire(v))
}

type AudioTranscriptionCapabilities struct {
	Modes                  []string              `json:"modes,omitempty"`
	Diarization            *bool                 `json:"diarization,omitempty"`
	ContextBias            *bool                 `json:"context_bias,omitempty"`
	TurnDetection          *bool                 `json:"turn_detection,omitempty"`
	TimestampGranularities []string              `json:"timestamp_granularities,omitempty"`
	Languages              []string              `json:"languages,omitempty"`
	StreamingDelaysMs      []StreamingDelayRange `json:"streaming_delays_ms,omitempty"`
}

func (v AudioTranscriptionCapabilities) MarshalJSON() ([]byte, error) {
	type wire AudioTranscriptionCapabilities
	if v.Diarization == nil {
		x := false
		v.Diarization = &x
	}
	if v.ContextBias == nil {
		x := false
		v.ContextBias = &x
	}
	if v.TurnDetection == nil {
		x := false
		v.TurnDetection = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Modes != nil {
		data, err := json.Marshal(v.Modes)
		if err != nil {
			return nil, err
		}
		fields["modes"] = data
	}
	if v.TimestampGranularities != nil {
		data, err := json.Marshal(v.TimestampGranularities)
		if err != nil {
			return nil, err
		}
		fields["timestamp_granularities"] = data
	}
	if v.Languages != nil {
		data, err := json.Marshal(v.Languages)
		if err != nil {
			return nil, err
		}
		fields["languages"] = data
	}
	if v.StreamingDelaysMs != nil {
		data, err := json.Marshal(v.StreamingDelaysMs)
		if err != nil {
			return nil, err
		}
		fields["streaming_delays_ms"] = data
	}
	return json.Marshal(fields)
}

type AudioTranscriptionRequest struct {
	Model                  string                 `json:"model"`
	File                   any                    `json:"file,omitempty"`
	FileURL                *string                `json:"file_url,omitempty"`
	FileID                 *string                `json:"file_id,omitempty"`
	Language               *string                `json:"language,omitempty"`
	Temperature            *float64               `json:"temperature,omitempty"`
	Stream                 *bool                  `json:"stream,omitempty"`
	Diarize                *bool                  `json:"diarize,omitempty"`
	ContextBias            []string               `json:"context_bias,omitempty"`
	SpeakerReferences      map[string]string      `json:"speaker_references,omitempty"`
	TimestampGranularities []TimestampGranularity `json:"timestamp_granularities,omitempty"`
}

func (v AudioTranscriptionRequest) MarshalJSON() ([]byte, error) {
	type wire AudioTranscriptionRequest
	if v.Stream == nil {
		x := false
		v.Stream = &x
	}
	if v.Diarize == nil {
		x := false
		v.Diarize = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.ContextBias != nil {
		data, err := json.Marshal(v.ContextBias)
		if err != nil {
			return nil, err
		}
		fields["context_bias"] = data
	}
	if v.SpeakerReferences != nil {
		data, err := json.Marshal(v.SpeakerReferences)
		if err != nil {
			return nil, err
		}
		fields["speaker_references"] = data
	}
	if v.TimestampGranularities != nil {
		data, err := json.Marshal(v.TimestampGranularities)
		if err != nil {
			return nil, err
		}
		fields["timestamp_granularities"] = data
	}
	return json.Marshal(fields)
}

type AudioTranscriptionRequestStream struct {
	Model                  string                 `json:"model"`
	File                   any                    `json:"file,omitempty"`
	FileURL                *string                `json:"file_url,omitempty"`
	FileID                 *string                `json:"file_id,omitempty"`
	Language               *string                `json:"language,omitempty"`
	Temperature            *float64               `json:"temperature,omitempty"`
	Stream                 *bool                  `json:"stream,omitempty"`
	Diarize                *bool                  `json:"diarize,omitempty"`
	ContextBias            []string               `json:"context_bias,omitempty"`
	SpeakerReferences      map[string]string      `json:"speaker_references,omitempty"`
	TimestampGranularities []TimestampGranularity `json:"timestamp_granularities,omitempty"`
}

func (v AudioTranscriptionRequestStream) MarshalJSON() ([]byte, error) {
	type wire AudioTranscriptionRequestStream
	if v.Stream == nil {
		x := true
		v.Stream = &x
	}
	if v.Diarize == nil {
		x := false
		v.Diarize = &x
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.ContextBias != nil {
		data, err := json.Marshal(v.ContextBias)
		if err != nil {
			return nil, err
		}
		fields["context_bias"] = data
	}
	if v.SpeakerReferences != nil {
		data, err := json.Marshal(v.SpeakerReferences)
		if err != nil {
			return nil, err
		}
		fields["speaker_references"] = data
	}
	if v.TimestampGranularities != nil {
		data, err := json.Marshal(v.TimestampGranularities)
		if err != nil {
			return nil, err
		}
		fields["timestamp_granularities"] = data
	}
	return json.Marshal(fields)
}

type AudioTTSConfig struct {
	Model          *string  `json:"model,omitempty"`
	VoiceID        *string  `json:"voice_id,omitempty"`
	Expressiveness *float64 `json:"expressiveness,omitempty"`
}
type AudioTurnConfig struct {
	AllowInterruptions *bool                  `json:"allow_interruptions,omitempty"`
	Inactivity         *AudioInactivityConfig `json:"inactivity,omitempty"`
}
type AudioTurnDetectionConfig struct {
	StartThreshold    *float64 `json:"start_threshold,omitempty"`
	EagerEndThreshold *float64 `json:"eager_end_threshold,omitempty"`
	EndThreshold      *float64 `json:"end_threshold,omitempty"`
	MinSilenceMs      *int     `json:"min_silence_ms,omitempty"`
	EndTimeoutMs      *int     `json:"end_timeout_ms,omitempty"`
}
type AudioUserTranscriptTurn struct {
	Timestamp  time.Time                 `json:"timestamp"`
	Latency    AudioLatencyBreakdownUser `json:"latency"`
	Content    *string                   `json:"content,omitempty"`
	DurationMs *float64                  `json:"duration_ms,omitempty"`
	Role       string                    `json:"role,omitempty"`
}

func (v AudioUserTranscriptTurn) MarshalJSON() ([]byte, error) {
	type wire AudioUserTranscriptTurn
	if v.Role == "" {
		v.Role = "user"
	}
	return json.Marshal(wire(v))
}

type BaseModelCard struct {
	ID                             string                          `json:"id"`
	Capabilities                   ModelCapabilities               `json:"capabilities"`
	Object                         *string                         `json:"object,omitempty"`
	Created                        *int                            `json:"created,omitempty"`
	OwnedBy                        *string                         `json:"owned_by,omitempty"`
	Name                           *string                         `json:"name,omitempty"`
	Description                    *string                         `json:"description,omitempty"`
	MaxContextLength               *int                            `json:"max_context_length,omitempty"`
	Aliases                        []string                        `json:"aliases,omitempty"`
	Deprecation                    *time.Time                      `json:"deprecation,omitempty"`
	DeprecationReplacementModel    *string                         `json:"deprecation_replacement_model,omitempty"`
	DefaultModelTemperature        *float64                        `json:"default_model_temperature,omitempty"`
	Internal                       *bool                           `json:"internal,omitempty"`
	BillingModelName               *string                         `json:"billing_model_name,omitempty"`
	Type                           string                          `json:"type,omitempty"`
	AudioTranscriptionCapabilities *AudioTranscriptionCapabilities `json:"audio_transcription_capabilities,omitempty"`
	AudioSpeechCapabilities        *AudioSpeechCapabilities        `json:"audio_speech_capabilities,omitempty"`
}

func (v BaseModelCard) MarshalJSON() ([]byte, error) {
	type wire BaseModelCard
	if v.Object == nil {
		x := "model"
		v.Object = &x
	}
	if v.OwnedBy == nil {
		x := "mistralai"
		v.OwnedBy = &x
	}
	if v.MaxContextLength == nil {
		x := 32768
		v.MaxContextLength = &x
	}
	if v.Internal == nil {
		x := false
		v.Internal = &x
	}
	if v.Type == "" {
		v.Type = "base"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Aliases != nil {
		data, err := json.Marshal(v.Aliases)
		if err != nil {
			return nil, err
		}
		fields["aliases"] = data
	}
	return json.Marshal(fields)
}

type BlobResourceContents struct {
	Uri      string         `json:"uri"`
	Blob     string         `json:"blob"`
	MimeType *string        `json:"mimeType,omitempty"`
	Meta     map[string]any `json:"_meta,omitempty"`
}

func (v BlobResourceContents) MarshalJSON() ([]byte, error) {
	type wire BlobResourceContents
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
	return json.Marshal(fields)
}

type ChatCompletionRequest struct {
	Model             string                           `json:"model"`
	Messages          []ChatCompletionRequestMessage   `json:"messages"`
	Temperature       *float64                         `json:"temperature,omitempty"`
	TopP              *float64                         `json:"top_p,omitempty"`
	MaxTokens         *int                             `json:"max_tokens,omitempty"`
	Stream            *bool                            `json:"stream,omitempty"`
	Stop              *ChatCompletionRequestStop       `json:"stop,omitempty"`
	RandomSeed        *int                             `json:"random_seed,omitempty"`
	Metadata          map[string]any                   `json:"metadata,omitempty"`
	ResponseFormat    *ResponseFormat                  `json:"response_format,omitempty"`
	Tools             []ChatCompletionRequestTool      `json:"tools,omitempty"`
	ToolChoice        *ChatCompletionRequestToolChoice `json:"tool_choice,omitempty"`
	PresencePenalty   *float64                         `json:"presence_penalty,omitempty"`
	FrequencyPenalty  *float64                         `json:"frequency_penalty,omitempty"`
	MinTokens         *int                             `json:"min_tokens,omitempty"`
	RepetitionPenalty *float64                         `json:"repetition_penalty,omitempty"`
	TopK              *int                             `json:"top_k,omitempty"`
	Logprobs          *bool                            `json:"logprobs,omitempty"`
	TopLogprobs       *int                             `json:"top_logprobs,omitempty"`
	PromptLogprobs    *bool                            `json:"prompt_logprobs,omitempty"`
	TopPromptLogprobs *int                             `json:"top_prompt_logprobs,omitempty"`
	N                 *int                             `json:"n,omitempty"`
	Prediction        *Prediction                      `json:"prediction,omitempty"`
	ParallelToolCalls *bool                            `json:"parallel_tool_calls,omitempty"`
	ReasoningEffort   *ReasoningEffort                 `json:"reasoning_effort,omitempty"`
	PromptMode        *MistralPromptMode               `json:"prompt_mode,omitempty"`
	Guardrails        []GuardrailConfig                `json:"guardrails,omitempty"`
	PromptCacheKey    *string                          `json:"prompt_cache_key,omitempty"`
	ServiceTier       *RequestedServiceTier            `json:"service_tier,omitempty"`
	SafePrompt        *bool                            `json:"safe_prompt,omitempty"`
}

func (v ChatCompletionRequest) MarshalJSON() ([]byte, error) {
	type wire ChatCompletionRequest
	if v.Stream == nil {
		x := false
		v.Stream = &x
	}
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
	if v.Tools != nil {
		data, err := json.Marshal(v.Tools)
		if err != nil {
			return nil, err
		}
		fields["tools"] = data
	}
	if v.Guardrails != nil {
		data, err := json.Marshal(v.Guardrails)
		if err != nil {
			return nil, err
		}
		fields["guardrails"] = data
	}
	return json.Marshal(fields)
}

type ChatCompletionStreamRequest struct {
	Model             string                                 `json:"model"`
	Messages          []ChatCompletionStreamRequestMessage   `json:"messages"`
	Temperature       *float64                               `json:"temperature,omitempty"`
	TopP              *float64                               `json:"top_p,omitempty"`
	MaxTokens         *int                                   `json:"max_tokens,omitempty"`
	Stream            *bool                                  `json:"stream,omitempty"`
	Stop              *ChatCompletionStreamRequestStop       `json:"stop,omitempty"`
	RandomSeed        *int                                   `json:"random_seed,omitempty"`
	Metadata          map[string]any                         `json:"metadata,omitempty"`
	ResponseFormat    *ResponseFormat                        `json:"response_format,omitempty"`
	Tools             []ChatCompletionStreamRequestTool      `json:"tools,omitempty"`
	ToolChoice        *ChatCompletionStreamRequestToolChoice `json:"tool_choice,omitempty"`
	PresencePenalty   *float64                               `json:"presence_penalty,omitempty"`
	FrequencyPenalty  *float64                               `json:"frequency_penalty,omitempty"`
	MinTokens         *int                                   `json:"min_tokens,omitempty"`
	RepetitionPenalty *float64                               `json:"repetition_penalty,omitempty"`
	TopK              *int                                   `json:"top_k,omitempty"`
	Logprobs          *bool                                  `json:"logprobs,omitempty"`
	TopLogprobs       *int                                   `json:"top_logprobs,omitempty"`
	PromptLogprobs    *bool                                  `json:"prompt_logprobs,omitempty"`
	TopPromptLogprobs *int                                   `json:"top_prompt_logprobs,omitempty"`
	N                 *int                                   `json:"n,omitempty"`
	Prediction        *Prediction                            `json:"prediction,omitempty"`
	ParallelToolCalls *bool                                  `json:"parallel_tool_calls,omitempty"`
	ReasoningEffort   *ReasoningEffort                       `json:"reasoning_effort,omitempty"`
	PromptMode        *MistralPromptMode                     `json:"prompt_mode,omitempty"`
	Guardrails        []GuardrailConfig                      `json:"guardrails,omitempty"`
	PromptCacheKey    *string                                `json:"prompt_cache_key,omitempty"`
	ServiceTier       *RequestedServiceTier                  `json:"service_tier,omitempty"`
	SafePrompt        *bool                                  `json:"safe_prompt,omitempty"`
}

func (v ChatCompletionStreamRequest) MarshalJSON() ([]byte, error) {
	type wire ChatCompletionStreamRequest
	if v.Stream == nil {
		x := true
		v.Stream = &x
	}
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
	if v.Tools != nil {
		data, err := json.Marshal(v.Tools)
		if err != nil {
			return nil, err
		}
		fields["tools"] = data
	}
	if v.Guardrails != nil {
		data, err := json.Marshal(v.Guardrails)
		if err != nil {
			return nil, err
		}
		fields["guardrails"] = data
	}
	return json.Marshal(fields)
}

type ClientCapabilities struct {
	Experimental map[string]map[string]any `json:"experimental,omitempty"`
	Sampling     *SamplingCapability       `json:"sampling,omitempty"`
	Elicitation  *ElicitationCapability    `json:"elicitation,omitempty"`
	Roots        *RootsCapability          `json:"roots,omitempty"`
	Extensions   map[string]map[string]any `json:"extensions,omitempty"`
	Tasks        *ClientTasksCapability    `json:"tasks,omitempty"`
}

func (v ClientCapabilities) MarshalJSON() ([]byte, error) {
	type wire ClientCapabilities
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Experimental != nil {
		data, err := json.Marshal(v.Experimental)
		if err != nil {
			return nil, err
		}
		fields["experimental"] = data
	}
	if v.Extensions != nil {
		data, err := json.Marshal(v.Extensions)
		if err != nil {
			return nil, err
		}
		fields["extensions"] = data
	}
	return json.Marshal(fields)
}

type ClientTasksCapability struct {
	List     *TasksListCapability           `json:"list,omitempty"`
	Cancel   *TasksCancelCapability         `json:"cancel,omitempty"`
	Requests *ClientTasksRequestsCapability `json:"requests,omitempty"`
}
type ClientTasksRequestsCapability struct {
	Sampling    *TasksSamplingCapability    `json:"sampling,omitempty"`
	Elicitation *TasksElicitationCapability `json:"elicitation,omitempty"`
}
type CompletionsCapability struct {
}
type CompletionTokensDetails struct {
	ReasoningTokens *int `json:"reasoning_tokens,omitempty"`
}
type UnknownConnectorToolCallResponseContent struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownConnectorToolCallResponseContent) MarshalJSON() ([]byte, error) {
	type wire UnknownConnectorToolCallResponseContent
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type ConnectorToolCallResponse struct {
	AdditionalProperties map[string]any                     `json:"-"`
	Content              []ConnectorToolCallResponseContent `json:"content"`
	Metadata             any                                `json:"metadata,omitempty"`
}

func (v ConnectorToolCallResponse) MarshalJSON() ([]byte, error) {
	type wire ConnectorToolCallResponse
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
		case "content", "metadata":
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
func (v *ConnectorToolCallResponse) UnmarshalJSON(data []byte) error {
	type wire ConnectorToolCallResponse
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content")
	delete(fields, "metadata")
	*v = ConnectorToolCallResponse(decoded)
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

type CountResponseInputTokensRequestResponsesToolChoiceFunction struct {
	AdditionalProperties map[string]any `json:"-"`
	Name                 string         `json:"name"`
	Type                 string         `json:"type,omitempty"`
}

func (v CountResponseInputTokensRequestResponsesToolChoiceFunction) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensRequestResponsesToolChoiceFunction
	if v.Type == "" {
		v.Type = "function"
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
		case "name", "type":
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
func (v *CountResponseInputTokensRequestResponsesToolChoiceFunction) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensRequestResponsesToolChoiceFunction
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "type")
	*v = CountResponseInputTokensRequestResponsesToolChoiceFunction(decoded)
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

type CountResponseInputTokensRequestResponsesSupportedReasoning struct {
	Effort *CountResponseInputTokensRequestEffort `json:"effort,omitempty"`
}
type CountResponseInputTokensRequestResponsesFormatJSONObject struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v CountResponseInputTokensRequestResponsesFormatJSONObject) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensRequestResponsesFormatJSONObject
	if v.Type == "" {
		v.Type = "json_object"
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
func (v *CountResponseInputTokensRequestResponsesFormatJSONObject) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensRequestResponsesFormatJSONObject
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = CountResponseInputTokensRequestResponsesFormatJSONObject(decoded)
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

type CountResponseInputTokensRequestResponsesFormatTextJSONSchemaConfig struct {
	AdditionalProperties map[string]any `json:"-"`
	Name                 string         `json:"name"`
	Schema               map[string]any `json:"schema"`
	Type                 string         `json:"type,omitempty"`
	Description          *string        `json:"description,omitempty"`
	Strict               *bool          `json:"strict,omitempty"`
}

func (v CountResponseInputTokensRequestResponsesFormatTextJSONSchemaConfig) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensRequestResponsesFormatTextJSONSchemaConfig
	if v.Type == "" {
		v.Type = "json_schema"
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
		case "name", "schema", "type", "description", "strict":
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
func (v *CountResponseInputTokensRequestResponsesFormatTextJSONSchemaConfig) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensRequestResponsesFormatTextJSONSchemaConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "schema")
	delete(fields, "type")
	delete(fields, "description")
	delete(fields, "strict")
	*v = CountResponseInputTokensRequestResponsesFormatTextJSONSchemaConfig(decoded)
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

type CountResponseInputTokensRequestResponsesFormatText struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v CountResponseInputTokensRequestResponsesFormatText) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensRequestResponsesFormatText
	if v.Type == "" {
		v.Type = "text"
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
func (v *CountResponseInputTokensRequestResponsesFormatText) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensRequestResponsesFormatText
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = CountResponseInputTokensRequestResponsesFormatText(decoded)
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

type CountResponseInputTokensRequestResponsesTextConfig struct {
	AdditionalProperties map[string]any                            `json:"-"`
	Format               *CountResponseInputTokensRequestFormat    `json:"format,omitempty"`
	Verbosity            *CountResponseInputTokensRequestVerbosity `json:"verbosity,omitempty"`
}

func (v CountResponseInputTokensRequestResponsesTextConfig) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensRequestResponsesTextConfig
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
		case "format", "verbosity":
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
func (v *CountResponseInputTokensRequestResponsesTextConfig) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensRequestResponsesTextConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "format")
	delete(fields, "verbosity")
	*v = CountResponseInputTokensRequestResponsesTextConfig(decoded)
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

type CountResponseInputTokensRequest struct {
	AdditionalProperties map[string]any                                              `json:"-"`
	Model                string                                                      `json:"model"`
	Input                CountResponseInputTokensRequestInput2                       `json:"input"`
	Instructions         *string                                                     `json:"instructions,omitempty"`
	ParallelToolCalls    *bool                                                       `json:"parallel_tool_calls,omitempty"`
	PreviousResponseID   any                                                         `json:"previous_response_id,omitempty"`
	Truncation           *string                                                     `json:"truncation,omitempty"`
	Tools                []CountResponseInputTokensRequestTool                       `json:"tools,omitempty"`
	ToolChoice           *CountResponseInputTokensRequestToolChoiceUnion             `json:"tool_choice,omitempty"`
	Reasoning            *CountResponseInputTokensRequestResponsesSupportedReasoning `json:"reasoning,omitempty"`
	Conversation         any                                                         `json:"conversation,omitempty"`
	Prompt               any                                                         `json:"prompt,omitempty"`
	Text                 *CountResponseInputTokensRequestResponsesTextConfig         `json:"text,omitempty"`
}

func (v CountResponseInputTokensRequest) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensRequest
	if v.ParallelToolCalls == nil {
		x := true
		v.ParallelToolCalls = &x
	}
	if v.Truncation == nil {
		x := "disabled"
		v.Truncation = &x
	}
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
	for key, value := range v.AdditionalProperties {
		switch key {
		case "model", "input", "instructions", "parallel_tool_calls", "previous_response_id", "truncation", "tools", "tool_choice", "reasoning", "conversation", "prompt", "text":
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
func (v *CountResponseInputTokensRequest) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensRequest
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "model")
	delete(fields, "input")
	delete(fields, "instructions")
	delete(fields, "parallel_tool_calls")
	delete(fields, "previous_response_id")
	delete(fields, "truncation")
	delete(fields, "tools")
	delete(fields, "tool_choice")
	delete(fields, "reasoning")
	delete(fields, "conversation")
	delete(fields, "prompt")
	delete(fields, "text")
	*v = CountResponseInputTokensRequest(decoded)
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

type CountResponseInputTokensResponse struct {
	AdditionalProperties map[string]any `json:"-"`
	InputTokens          int            `json:"input_tokens"`
	Object               string         `json:"object,omitempty"`
}

func (v CountResponseInputTokensResponse) MarshalJSON() ([]byte, error) {
	type wire CountResponseInputTokensResponse
	if v.Object == "" {
		v.Object = "response.input_tokens"
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
		case "input_tokens", "object":
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
func (v *CountResponseInputTokensResponse) UnmarshalJSON(data []byte) error {
	type wire CountResponseInputTokensResponse
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "input_tokens")
	delete(fields, "object")
	*v = CountResponseInputTokensResponse(decoded)
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

type CreateResponseRequestResponsesToolChoiceFunction struct {
	AdditionalProperties map[string]any `json:"-"`
	Name                 string         `json:"name"`
	Type                 string         `json:"type,omitempty"`
}

func (v CreateResponseRequestResponsesToolChoiceFunction) MarshalJSON() ([]byte, error) {
	type wire CreateResponseRequestResponsesToolChoiceFunction
	if v.Type == "" {
		v.Type = "function"
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
		case "name", "type":
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
func (v *CreateResponseRequestResponsesToolChoiceFunction) UnmarshalJSON(data []byte) error {
	type wire CreateResponseRequestResponsesToolChoiceFunction
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "type")
	*v = CreateResponseRequestResponsesToolChoiceFunction(decoded)
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

type CreateResponseRequestResponsesSupportedReasoning struct {
	Effort *CreateResponseRequestEffort `json:"effort,omitempty"`
}
type CreateResponseRequestResponsesFormatJSONObject struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v CreateResponseRequestResponsesFormatJSONObject) MarshalJSON() ([]byte, error) {
	type wire CreateResponseRequestResponsesFormatJSONObject
	if v.Type == "" {
		v.Type = "json_object"
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
func (v *CreateResponseRequestResponsesFormatJSONObject) UnmarshalJSON(data []byte) error {
	type wire CreateResponseRequestResponsesFormatJSONObject
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = CreateResponseRequestResponsesFormatJSONObject(decoded)
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

type CreateResponseRequestResponsesFormatTextJSONSchemaConfig struct {
	AdditionalProperties map[string]any `json:"-"`
	Name                 string         `json:"name"`
	Schema               map[string]any `json:"schema"`
	Type                 string         `json:"type,omitempty"`
	Description          *string        `json:"description,omitempty"`
	Strict               *bool          `json:"strict,omitempty"`
}

func (v CreateResponseRequestResponsesFormatTextJSONSchemaConfig) MarshalJSON() ([]byte, error) {
	type wire CreateResponseRequestResponsesFormatTextJSONSchemaConfig
	if v.Type == "" {
		v.Type = "json_schema"
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
		case "name", "schema", "type", "description", "strict":
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
func (v *CreateResponseRequestResponsesFormatTextJSONSchemaConfig) UnmarshalJSON(data []byte) error {
	type wire CreateResponseRequestResponsesFormatTextJSONSchemaConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "schema")
	delete(fields, "type")
	delete(fields, "description")
	delete(fields, "strict")
	*v = CreateResponseRequestResponsesFormatTextJSONSchemaConfig(decoded)
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

type CreateResponseRequestResponsesFormatText struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v CreateResponseRequestResponsesFormatText) MarshalJSON() ([]byte, error) {
	type wire CreateResponseRequestResponsesFormatText
	if v.Type == "" {
		v.Type = "text"
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
func (v *CreateResponseRequestResponsesFormatText) UnmarshalJSON(data []byte) error {
	type wire CreateResponseRequestResponsesFormatText
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = CreateResponseRequestResponsesFormatText(decoded)
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

type CreateResponseRequestResponsesTextConfig struct {
	AdditionalProperties map[string]any                  `json:"-"`
	Format               *CreateResponseRequestFormat    `json:"format,omitempty"`
	Verbosity            *CreateResponseRequestVerbosity `json:"verbosity,omitempty"`
}

func (v CreateResponseRequestResponsesTextConfig) MarshalJSON() ([]byte, error) {
	type wire CreateResponseRequestResponsesTextConfig
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
		case "format", "verbosity":
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
func (v *CreateResponseRequestResponsesTextConfig) UnmarshalJSON(data []byte) error {
	type wire CreateResponseRequestResponsesTextConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "format")
	delete(fields, "verbosity")
	*v = CreateResponseRequestResponsesTextConfig(decoded)
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

type CreateResponseRequest struct {
	AdditionalProperties map[string]any                                    `json:"-"`
	Model                string                                            `json:"model"`
	Input                CreateResponseRequestInput2                       `json:"input"`
	Instructions         *string                                           `json:"instructions,omitempty"`
	ParallelToolCalls    *bool                                             `json:"parallel_tool_calls,omitempty"`
	PreviousResponseID   any                                               `json:"previous_response_id,omitempty"`
	Truncation           *string                                           `json:"truncation,omitempty"`
	Tools                []CreateResponseRequestTool                       `json:"tools,omitempty"`
	ToolChoice           *CreateResponseRequestToolChoiceUnion             `json:"tool_choice,omitempty"`
	Reasoning            *CreateResponseRequestResponsesSupportedReasoning `json:"reasoning,omitempty"`
	Conversation         any                                               `json:"conversation,omitempty"`
	Prompt               any                                               `json:"prompt,omitempty"`
	Text                 *CreateResponseRequestResponsesTextConfig         `json:"text,omitempty"`
	Background           *bool                                             `json:"background,omitempty"`
	Stream               *bool                                             `json:"stream,omitempty"`
	Store                *bool                                             `json:"store,omitempty"`
	MaxOutputTokens      *int                                              `json:"max_output_tokens,omitempty"`
	Temperature          *float64                                          `json:"temperature,omitempty"`
	TopP                 *float64                                          `json:"top_p,omitempty"`
	Metadata             map[string]string                                 `json:"metadata,omitempty"`
	PromptCacheKey       *string                                           `json:"prompt_cache_key,omitempty"`
	MaxToolCalls         *int                                              `json:"max_tool_calls,omitempty"`
	Moderation           any                                               `json:"moderation,omitempty"`
	TopLogprobs          *int                                              `json:"top_logprobs,omitempty"`
}

func (v CreateResponseRequest) MarshalJSON() ([]byte, error) {
	type wire CreateResponseRequest
	if v.ParallelToolCalls == nil {
		x := true
		v.ParallelToolCalls = &x
	}
	if v.Truncation == nil {
		x := "disabled"
		v.Truncation = &x
	}
	if v.Background == nil {
		x := false
		v.Background = &x
	}
	if v.Stream == nil {
		x := false
		v.Stream = &x
	}
	if v.Store == nil {
		x := false
		v.Store = &x
	}
	if v.TopLogprobs == nil {
		x := 0
		v.TopLogprobs = &x
	}
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
	if v.Metadata != nil {
		data, err := json.Marshal(v.Metadata)
		if err != nil {
			return nil, err
		}
		fields["metadata"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "model", "input", "instructions", "parallel_tool_calls", "previous_response_id", "truncation", "tools", "tool_choice", "reasoning", "conversation", "prompt", "text", "background", "stream", "store", "max_output_tokens", "temperature", "top_p", "metadata", "prompt_cache_key", "max_tool_calls", "moderation", "top_logprobs":
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
func (v *CreateResponseRequest) UnmarshalJSON(data []byte) error {
	type wire CreateResponseRequest
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "model")
	delete(fields, "input")
	delete(fields, "instructions")
	delete(fields, "parallel_tool_calls")
	delete(fields, "previous_response_id")
	delete(fields, "truncation")
	delete(fields, "tools")
	delete(fields, "tool_choice")
	delete(fields, "reasoning")
	delete(fields, "conversation")
	delete(fields, "prompt")
	delete(fields, "text")
	delete(fields, "background")
	delete(fields, "stream")
	delete(fields, "store")
	delete(fields, "max_output_tokens")
	delete(fields, "temperature")
	delete(fields, "top_p")
	delete(fields, "metadata")
	delete(fields, "prompt_cache_key")
	delete(fields, "max_tool_calls")
	delete(fields, "moderation")
	delete(fields, "top_logprobs")
	*v = CreateResponseRequest(decoded)
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

type ResponsesError struct {
	AdditionalProperties map[string]any `json:"-"`
	Code                 Code           `json:"code"`
	Message              string         `json:"message"`
}

func (v ResponsesError) MarshalJSON() ([]byte, error) {
	type wire ResponsesError
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
		case "code", "message":
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
func (v *ResponsesError) UnmarshalJSON(data []byte) error {
	type wire ResponsesError
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "code")
	delete(fields, "message")
	*v = ResponsesError(decoded)
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

type ResponsesIncompleteDetails struct {
	AdditionalProperties map[string]any `json:"-"`
	Reason               *Reason        `json:"reason,omitempty"`
}

func (v ResponsesIncompleteDetails) MarshalJSON() ([]byte, error) {
	type wire ResponsesIncompleteDetails
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
		case "reason":
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
func (v *ResponsesIncompleteDetails) UnmarshalJSON(data []byte) error {
	type wire ResponsesIncompleteDetails
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "reason")
	*v = ResponsesIncompleteDetails(decoded)
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

type CreateResponseResponseResponsesSupportedReasoning struct {
	Effort *ResponsesSupportedReasoningEffort `json:"effort,omitempty"`
}
type ResponsesTextConfigResponsesFormatJSONObject struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesTextConfigResponsesFormatJSONObject) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextConfigResponsesFormatJSONObject
	if v.Type == "" {
		v.Type = "json_object"
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
func (v *ResponsesTextConfigResponsesFormatJSONObject) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextConfigResponsesFormatJSONObject
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = ResponsesTextConfigResponsesFormatJSONObject(decoded)
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

type ResponsesTextConfigResponsesFormatTextJSONSchemaConfig struct {
	AdditionalProperties map[string]any `json:"-"`
	Name                 string         `json:"name"`
	Schema               map[string]any `json:"schema"`
	Type                 string         `json:"type,omitempty"`
	Description          *string        `json:"description,omitempty"`
	Strict               *bool          `json:"strict,omitempty"`
}

func (v ResponsesTextConfigResponsesFormatTextJSONSchemaConfig) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextConfigResponsesFormatTextJSONSchemaConfig
	if v.Type == "" {
		v.Type = "json_schema"
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
		case "name", "schema", "type", "description", "strict":
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
func (v *ResponsesTextConfigResponsesFormatTextJSONSchemaConfig) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextConfigResponsesFormatTextJSONSchemaConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "schema")
	delete(fields, "type")
	delete(fields, "description")
	delete(fields, "strict")
	*v = ResponsesTextConfigResponsesFormatTextJSONSchemaConfig(decoded)
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

type ResponsesTextConfigResponsesFormatText struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesTextConfigResponsesFormatText) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextConfigResponsesFormatText
	if v.Type == "" {
		v.Type = "text"
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
func (v *ResponsesTextConfigResponsesFormatText) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextConfigResponsesFormatText
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = ResponsesTextConfigResponsesFormatText(decoded)
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

type CreateResponseResponseResponsesTextConfig struct {
	AdditionalProperties map[string]any                `json:"-"`
	Format               *ResponsesTextConfigFormat    `json:"format,omitempty"`
	Verbosity            *ResponsesTextConfigVerbosity `json:"verbosity,omitempty"`
}

func (v CreateResponseResponseResponsesTextConfig) MarshalJSON() ([]byte, error) {
	type wire CreateResponseResponseResponsesTextConfig
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
		case "format", "verbosity":
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
func (v *CreateResponseResponseResponsesTextConfig) UnmarshalJSON(data []byte) error {
	type wire CreateResponseResponseResponsesTextConfig
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "format")
	delete(fields, "verbosity")
	*v = CreateResponseResponseResponsesTextConfig(decoded)
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

type ResponsesUsage struct {
	AdditionalProperties map[string]any               `json:"-"`
	InputTokens          int                          `json:"input_tokens"`
	InputTokensDetails   ResponsesInputTokensDetails  `json:"input_tokens_details"`
	OutputTokens         int                          `json:"output_tokens"`
	OutputTokensDetails  ResponsesOutputTokensDetails `json:"output_tokens_details"`
	TotalTokens          int                          `json:"total_tokens"`
}

func (v ResponsesUsage) MarshalJSON() ([]byte, error) {
	type wire ResponsesUsage
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
		case "input_tokens", "input_tokens_details", "output_tokens", "output_tokens_details", "total_tokens":
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
func (v *ResponsesUsage) UnmarshalJSON(data []byte) error {
	type wire ResponsesUsage
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "input_tokens")
	delete(fields, "input_tokens_details")
	delete(fields, "output_tokens")
	delete(fields, "output_tokens_details")
	delete(fields, "total_tokens")
	*v = ResponsesUsage(decoded)
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

type CreateResponseResponse struct {
	AdditionalProperties map[string]any                                     `json:"-"`
	ID                   string                                             `json:"id"`
	CreatedAt            float64                                            `json:"created_at"`
	Model                string                                             `json:"model"`
	Output               []CreateResponseResponseOutput                     `json:"output"`
	ParallelToolCalls    bool                                               `json:"parallel_tool_calls"`
	ToolChoice           CreateResponseResponseToolChoiceUnion              `json:"tool_choice"`
	Tools                []CreateResponseResponseTool                       `json:"tools"`
	Error                *ResponsesError                                    `json:"error,omitempty"`
	IncompleteDetails    *ResponsesIncompleteDetails                        `json:"incomplete_details,omitempty"`
	Instructions         *string                                            `json:"instructions,omitempty"`
	Metadata             map[string]string                                  `json:"metadata,omitempty"`
	Object               string                                             `json:"object,omitempty"`
	Temperature          *float64                                           `json:"temperature,omitempty"`
	TopP                 *float64                                           `json:"top_p,omitempty"`
	Background           *bool                                              `json:"background,omitempty"`
	CompletedAt          *float64                                           `json:"completed_at,omitempty"`
	MaxOutputTokens      *int                                               `json:"max_output_tokens,omitempty"`
	MaxToolCalls         *int                                               `json:"max_tool_calls,omitempty"`
	PreviousResponseID   *string                                            `json:"previous_response_id,omitempty"`
	PromptCacheKey       *string                                            `json:"prompt_cache_key,omitempty"`
	PromptCacheRetention *PromptCacheRetention                              `json:"prompt_cache_retention,omitempty"`
	Reasoning            *CreateResponseResponseResponsesSupportedReasoning `json:"reasoning,omitempty"`
	SafetyIdentifier     *string                                            `json:"safety_identifier,omitempty"`
	ServiceTier          *ServiceTier                                       `json:"service_tier,omitempty"`
	Status               *CreateResponseResponseStatus                      `json:"status,omitempty"`
	Text                 *CreateResponseResponseResponsesTextConfig         `json:"text,omitempty"`
	TopLogprobs          *int                                               `json:"top_logprobs,omitempty"`
	Truncation           *Truncation                                        `json:"truncation,omitempty"`
	Usage                *ResponsesUsage                                    `json:"usage,omitempty"`
	User                 *string                                            `json:"user,omitempty"`
}

func (v CreateResponseResponse) MarshalJSON() ([]byte, error) {
	type wire CreateResponseResponse
	if v.Object == "" {
		v.Object = "response"
	}
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
	for key, value := range v.AdditionalProperties {
		switch key {
		case "id", "created_at", "model", "output", "parallel_tool_calls", "tool_choice", "tools", "error", "incomplete_details", "instructions", "metadata", "object", "temperature", "top_p", "background", "completed_at", "max_output_tokens", "max_tool_calls", "previous_response_id", "prompt_cache_key", "prompt_cache_retention", "reasoning", "safety_identifier", "service_tier", "status", "text", "top_logprobs", "truncation", "usage", "user":
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
func (v *CreateResponseResponse) UnmarshalJSON(data []byte) error {
	type wire CreateResponseResponse
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "id")
	delete(fields, "created_at")
	delete(fields, "model")
	delete(fields, "output")
	delete(fields, "parallel_tool_calls")
	delete(fields, "tool_choice")
	delete(fields, "tools")
	delete(fields, "error")
	delete(fields, "incomplete_details")
	delete(fields, "instructions")
	delete(fields, "metadata")
	delete(fields, "object")
	delete(fields, "temperature")
	delete(fields, "top_p")
	delete(fields, "background")
	delete(fields, "completed_at")
	delete(fields, "max_output_tokens")
	delete(fields, "max_tool_calls")
	delete(fields, "previous_response_id")
	delete(fields, "prompt_cache_key")
	delete(fields, "prompt_cache_retention")
	delete(fields, "reasoning")
	delete(fields, "safety_identifier")
	delete(fields, "service_tier")
	delete(fields, "status")
	delete(fields, "text")
	delete(fields, "top_logprobs")
	delete(fields, "truncation")
	delete(fields, "usage")
	delete(fields, "user")
	*v = CreateResponseResponse(decoded)
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

type CredentialsResponse struct {
	Credentials                       []any                `json:"credentials"`
	ConnectorPresetCredentialsForAuth []AuthenticationType `json:"connector_preset_credentials_for_auth,omitempty"`
}

func (v CredentialsResponse) MarshalJSON() ([]byte, error) {
	type wire CredentialsResponse
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.ConnectorPresetCredentialsForAuth != nil {
		data, err := json.Marshal(v.ConnectorPresetCredentialsForAuth)
		if err != nil {
			return nil, err
		}
		fields["connector_preset_credentials_for_auth"] = data
	}
	return json.Marshal(fields)
}

type UnknownCustomConnectorAuthorization struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownCustomConnectorAuthorization) MarshalJSON() ([]byte, error) {
	type wire UnknownCustomConnectorAuthorization
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type CustomConnector struct {
	ConnectorID       string                        `json:"connector_id"`
	Type              string                        `json:"type,omitempty"`
	Authorization     *CustomConnectorAuthorization `json:"authorization,omitempty"`
	ToolConfiguration any                           `json:"tool_configuration,omitempty"`
}

func (v CustomConnector) MarshalJSON() ([]byte, error) {
	type wire CustomConnector
	if v.Type == "" {
		v.Type = "connector"
	}
	return json.Marshal(wire(v))
}

type ElicitationCapability struct {
	Form *FormElicitationCapability `json:"form,omitempty"`
	URL  *URLElicitationCapability  `json:"url,omitempty"`
}
type EmbeddedResource struct {
	Resource    Resource       `json:"resource"`
	Type        string         `json:"type,omitempty"`
	Annotations *Annotations   `json:"annotations,omitempty"`
	Meta        map[string]any `json:"_meta,omitempty"`
}

func (v EmbeddedResource) MarshalJSON() ([]byte, error) {
	type wire EmbeddedResource
	if v.Type == "" {
		v.Type = "resource"
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
	return json.Marshal(fields)
}

type FormElicitationCapability struct {
}
type FunctionRequest struct {
	Name        string         `json:"name"`
	Parameters  map[string]any `json:"parameters"`
	Description *string        `json:"description,omitempty"`
	Strict      *bool          `json:"strict,omitempty"`
}

func (v FunctionRequest) MarshalJSON() ([]byte, error) {
	type wire FunctionRequest
	if v.Description == nil {
		x := ""
		v.Description = &x
	}
	if v.Strict == nil {
		x := false
		v.Strict = &x
	}
	return json.Marshal(wire(v))
}

type GetSummaryResponseSummary struct {
	Content    string        `json:"content"`
	Status     SummaryStatus `json:"status"`
	Translated bool          `json:"translated"`
}
type HangUpTool struct {
	Type   string            `json:"type,omitempty"`
	Filler *StaticToolFiller `json:"filler,omitempty"`
}

func (v HangUpTool) MarshalJSON() ([]byte, error) {
	type wire HangUpTool
	if v.Type == "" {
		v.Type = "hang_up"
	}
	return json.Marshal(wire(v))
}

type ImageContent struct {
	Data        string         `json:"data"`
	MimeType    string         `json:"mimeType"`
	Type        string         `json:"type,omitempty"`
	Annotations *Annotations   `json:"annotations,omitempty"`
	Meta        map[string]any `json:"_meta,omitempty"`
}

func (v ImageContent) MarshalJSON() ([]byte, error) {
	type wire ImageContent
	if v.Type == "" {
		v.Type = "image"
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
	return json.Marshal(fields)
}

type LoggingCapability struct {
}
type MCPServerCard struct {
	AdditionalProperties map[string]any      `json:"-"`
	Name                 string              `json:"name"`
	Version              string              `json:"version"`
	DollarSchema         *string             `json:"$schema,omitempty"`
	Capabilities         *ServerCapabilities `json:"capabilities,omitempty"`
	Title                *string             `json:"title,omitempty"`
	Description          *string             `json:"description,omitempty"`
	WebsiteURL           *string             `json:"websiteUrl,omitempty"`
	Repository           any                 `json:"repository,omitempty"`
	Icons                []MCPServerIcon     `json:"icons,omitempty"`
	Remotes              []any               `json:"remotes,omitempty"`
	Requires             *ClientCapabilities `json:"requires,omitempty"`
	Resources            *Resources          `json:"resources,omitempty"`
	Tools                *Tools              `json:"tools,omitempty"`
	Prompts              *Prompts            `json:"prompts,omitempty"`
	Meta                 any                 `json:"_meta,omitempty"`
}

func (v MCPServerCard) MarshalJSON() ([]byte, error) {
	type wire MCPServerCard
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Icons != nil {
		data, err := json.Marshal(v.Icons)
		if err != nil {
			return nil, err
		}
		fields["icons"] = data
	}
	if v.Remotes != nil {
		data, err := json.Marshal(v.Remotes)
		if err != nil {
			return nil, err
		}
		fields["remotes"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "name", "version", "$schema", "capabilities", "title", "description", "websiteUrl", "repository", "icons", "remotes", "requires", "resources", "tools", "prompts", "_meta":
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
func (v *MCPServerCard) UnmarshalJSON(data []byte) error {
	type wire MCPServerCard
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "version")
	delete(fields, "$schema")
	delete(fields, "capabilities")
	delete(fields, "title")
	delete(fields, "description")
	delete(fields, "websiteUrl")
	delete(fields, "repository")
	delete(fields, "icons")
	delete(fields, "remotes")
	delete(fields, "requires")
	delete(fields, "resources")
	delete(fields, "tools")
	delete(fields, "prompts")
	delete(fields, "_meta")
	*v = MCPServerCard(decoded)
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

type MCPServerCardTool struct {
	AdditionalProperties map[string]any   `json:"-"`
	Name                 string           `json:"name"`
	InputSchema          map[string]any   `json:"inputSchema"`
	Title                *string          `json:"title,omitempty"`
	Description          *string          `json:"description,omitempty"`
	Execution            *ToolExecution   `json:"execution,omitempty"`
	OutputSchema         map[string]any   `json:"outputSchema,omitempty"`
	Icons                []MCPServerIcon  `json:"icons,omitempty"`
	Annotations          *ToolAnnotations `json:"annotations,omitempty"`
	Meta                 any              `json:"_meta,omitempty"`
}

func (v MCPServerCardTool) MarshalJSON() ([]byte, error) {
	type wire MCPServerCardTool
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.OutputSchema != nil {
		data, err := json.Marshal(v.OutputSchema)
		if err != nil {
			return nil, err
		}
		fields["outputSchema"] = data
	}
	if v.Icons != nil {
		data, err := json.Marshal(v.Icons)
		if err != nil {
			return nil, err
		}
		fields["icons"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "name", "inputSchema", "title", "description", "execution", "outputSchema", "icons", "annotations", "_meta":
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
func (v *MCPServerCardTool) UnmarshalJSON(data []byte) error {
	type wire MCPServerCardTool
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "inputSchema")
	delete(fields, "title")
	delete(fields, "description")
	delete(fields, "execution")
	delete(fields, "outputSchema")
	delete(fields, "icons")
	delete(fields, "annotations")
	delete(fields, "_meta")
	*v = MCPServerCardTool(decoded)
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

type MCPServerIcon struct {
	Src      string   `json:"src"`
	MimeType *string  `json:"mimeType,omitempty"`
	Sizes    []string `json:"sizes,omitempty"`
	Theme    *Theme   `json:"theme,omitempty"`
}

func (v MCPServerIcon) MarshalJSON() ([]byte, error) {
	type wire MCPServerIcon
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Sizes != nil {
		data, err := json.Marshal(v.Sizes)
		if err != nil {
			return nil, err
		}
		fields["sizes"] = data
	}
	return json.Marshal(fields)
}

type MCPTool struct {
	AdditionalProperties map[string]any   `json:"-"`
	Name                 string           `json:"name"`
	InputSchema          map[string]any   `json:"inputSchema"`
	Title                *string          `json:"title,omitempty"`
	Description          *string          `json:"description,omitempty"`
	Execution            *ToolExecution   `json:"execution,omitempty"`
	OutputSchema         map[string]any   `json:"outputSchema,omitempty"`
	Icons                []MCPServerIcon  `json:"icons,omitempty"`
	Annotations          *ToolAnnotations `json:"annotations,omitempty"`
	Meta                 any              `json:"_meta,omitempty"`
}

func (v MCPTool) MarshalJSON() ([]byte, error) {
	type wire MCPTool
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.OutputSchema != nil {
		data, err := json.Marshal(v.OutputSchema)
		if err != nil {
			return nil, err
		}
		fields["outputSchema"] = data
	}
	if v.Icons != nil {
		data, err := json.Marshal(v.Icons)
		if err != nil {
			return nil, err
		}
		fields["icons"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "name", "inputSchema", "title", "description", "execution", "outputSchema", "icons", "annotations", "_meta":
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
func (v *MCPTool) UnmarshalJSON(data []byte) error {
	type wire MCPTool
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "inputSchema")
	delete(fields, "title")
	delete(fields, "description")
	delete(fields, "execution")
	delete(fields, "outputSchema")
	delete(fields, "icons")
	delete(fields, "annotations")
	delete(fields, "_meta")
	*v = MCPTool(decoded)
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

type MessageInputEntry struct {
	Role        MessageInputEntryRole    `json:"role"`
	Content     MessageInputEntryContent `json:"content"`
	Object      *string                  `json:"object,omitempty"`
	Type        *string                  `json:"type,omitempty"`
	CreatedAt   *time.Time               `json:"created_at,omitempty"`
	CompletedAt *time.Time               `json:"completed_at,omitempty"`
	ID          *string                  `json:"id,omitempty"`
	Prefix      *bool                    `json:"prefix,omitempty"`
}

func (v MessageInputEntry) MarshalJSON() ([]byte, error) {
	type wire MessageInputEntry
	if v.Object == nil {
		x := "entry"
		v.Object = &x
	}
	if v.Type == nil {
		x := "message.input"
		v.Type = &x
	}
	if v.Prefix == nil {
		x := false
		v.Prefix = &x
	}
	return json.Marshal(wire(v))
}

type PhoneDestination struct {
	Number string `json:"number"`
	Type   string `json:"type,omitempty"`
}

func (v PhoneDestination) MarshalJSON() ([]byte, error) {
	type wire PhoneDestination
	if v.Type == "" {
		v.Type = "phone"
	}
	return json.Marshal(wire(v))
}

type PromptArgument struct {
	Name        string  `json:"name"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Required    *bool   `json:"required,omitempty"`
}
type PromptsCapability struct {
	ListChanged *bool `json:"listChanged,omitempty"`
}
type RawAudio struct {
	Data string `json:"data"`
}
type RealtimeTranscriptionSessionUpdatePayload struct {
	AudioFormat            *AudioFormat         `json:"audio_format,omitempty"`
	TargetStreamingDelayMs *int                 `json:"target_streaming_delay_ms,omitempty"`
	Diarize                *bool                `json:"diarize,omitempty"`
	ContextBias            []string             `json:"context_bias,omitempty"`
	Language               *string              `json:"language,omitempty"`
	SpeakerReferences      map[string]string    `json:"speaker_references,omitempty"`
	TurnDetection          *TurnDetectionConfig `json:"turn_detection,omitempty"`
}

func (v RealtimeTranscriptionSessionUpdatePayload) MarshalJSON() ([]byte, error) {
	type wire RealtimeTranscriptionSessionUpdatePayload
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.ContextBias != nil {
		data, err := json.Marshal(v.ContextBias)
		if err != nil {
			return nil, err
		}
		fields["context_bias"] = data
	}
	if v.SpeakerReferences != nil {
		data, err := json.Marshal(v.SpeakerReferences)
		if err != nil {
			return nil, err
		}
		fields["speaker_references"] = data
	}
	return json.Marshal(fields)
}

type ResourceLink struct {
	Name        string          `json:"name"`
	Uri         string          `json:"uri"`
	Title       *string         `json:"title,omitempty"`
	Description *string         `json:"description,omitempty"`
	MimeType    *string         `json:"mimeType,omitempty"`
	Size        *int            `json:"size,omitempty"`
	Icons       []MCPServerIcon `json:"icons,omitempty"`
	Annotations *Annotations    `json:"annotations,omitempty"`
	Meta        map[string]any  `json:"_meta,omitempty"`
	Type        string          `json:"type,omitempty"`
}

func (v ResourceLink) MarshalJSON() ([]byte, error) {
	type wire ResourceLink
	if v.Type == "" {
		v.Type = "resource_link"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Icons != nil {
		data, err := json.Marshal(v.Icons)
		if err != nil {
			return nil, err
		}
		fields["icons"] = data
	}
	if v.Meta != nil {
		data, err := json.Marshal(v.Meta)
		if err != nil {
			return nil, err
		}
		fields["_meta"] = data
	}
	return json.Marshal(fields)
}

type ResourcesCapability struct {
	Subscribe   *bool `json:"subscribe,omitempty"`
	ListChanged *bool `json:"listChanged,omitempty"`
}
type ResponsesAnnotationContainerFileCitation struct {
	AdditionalProperties map[string]any `json:"-"`
	ContainerID          string         `json:"container_id"`
	EndIndex             int            `json:"end_index"`
	FileID               string         `json:"file_id"`
	Filename             string         `json:"filename"`
	StartIndex           int            `json:"start_index"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesAnnotationContainerFileCitation) MarshalJSON() ([]byte, error) {
	type wire ResponsesAnnotationContainerFileCitation
	if v.Type == "" {
		v.Type = "container_file_citation"
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
		case "container_id", "end_index", "file_id", "filename", "start_index", "type":
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
func (v *ResponsesAnnotationContainerFileCitation) UnmarshalJSON(data []byte) error {
	type wire ResponsesAnnotationContainerFileCitation
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "container_id")
	delete(fields, "end_index")
	delete(fields, "file_id")
	delete(fields, "filename")
	delete(fields, "start_index")
	delete(fields, "type")
	*v = ResponsesAnnotationContainerFileCitation(decoded)
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

type ResponsesAnnotationFileCitation struct {
	AdditionalProperties map[string]any `json:"-"`
	FileID               string         `json:"file_id"`
	Filename             string         `json:"filename"`
	Index                int            `json:"index"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesAnnotationFileCitation) MarshalJSON() ([]byte, error) {
	type wire ResponsesAnnotationFileCitation
	if v.Type == "" {
		v.Type = "file_citation"
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
		case "file_id", "filename", "index", "type":
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
func (v *ResponsesAnnotationFileCitation) UnmarshalJSON(data []byte) error {
	type wire ResponsesAnnotationFileCitation
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "file_id")
	delete(fields, "filename")
	delete(fields, "index")
	delete(fields, "type")
	*v = ResponsesAnnotationFileCitation(decoded)
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

type ResponsesAnnotationFilePath struct {
	AdditionalProperties map[string]any `json:"-"`
	FileID               string         `json:"file_id"`
	Index                int            `json:"index"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesAnnotationFilePath) MarshalJSON() ([]byte, error) {
	type wire ResponsesAnnotationFilePath
	if v.Type == "" {
		v.Type = "file_path"
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
		case "file_id", "index", "type":
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
func (v *ResponsesAnnotationFilePath) UnmarshalJSON(data []byte) error {
	type wire ResponsesAnnotationFilePath
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "file_id")
	delete(fields, "index")
	delete(fields, "type")
	*v = ResponsesAnnotationFilePath(decoded)
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

type ResponsesAnnotationURLCitation struct {
	AdditionalProperties map[string]any `json:"-"`
	EndIndex             int            `json:"end_index"`
	StartIndex           int            `json:"start_index"`
	Title                string         `json:"title"`
	URL                  string         `json:"url"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesAnnotationURLCitation) MarshalJSON() ([]byte, error) {
	type wire ResponsesAnnotationURLCitation
	if v.Type == "" {
		v.Type = "url_citation"
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
		case "end_index", "start_index", "title", "url", "type":
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
func (v *ResponsesAnnotationURLCitation) UnmarshalJSON(data []byte) error {
	type wire ResponsesAnnotationURLCitation
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "end_index")
	delete(fields, "start_index")
	delete(fields, "title")
	delete(fields, "url")
	delete(fields, "type")
	*v = ResponsesAnnotationURLCitation(decoded)
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

type ResponsesCompletedEvent struct {
	AdditionalProperties map[string]any         `json:"-"`
	Response             CreateResponseResponse `json:"response"`
	SequenceNumber       int                    `json:"sequence_number"`
	Type                 string                 `json:"type,omitempty"`
}

func (v ResponsesCompletedEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesCompletedEvent
	if v.Type == "" {
		v.Type = "response.completed"
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
		case "response", "sequence_number", "type":
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
func (v *ResponsesCompletedEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesCompletedEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "response")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesCompletedEvent(decoded)
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

type ResponsesContent struct {
	AdditionalProperties map[string]any `json:"-"`
	Text                 string         `json:"text"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesContent) MarshalJSON() ([]byte, error) {
	type wire ResponsesContent
	if v.Type == "" {
		v.Type = "reasoning_text"
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
func (v *ResponsesContent) UnmarshalJSON(data []byte) error {
	type wire ResponsesContent
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
	*v = ResponsesContent(decoded)
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

type ResponsesContentPartAddedEvent struct {
	AdditionalProperties map[string]any                     `json:"-"`
	ContentIndex         int                                `json:"content_index"`
	ItemID               string                             `json:"item_id"`
	OutputIndex          int                                `json:"output_index"`
	Part                 ResponsesContentPartAddedEventPart `json:"part"`
	SequenceNumber       int                                `json:"sequence_number"`
	Type                 string                             `json:"type,omitempty"`
}

func (v ResponsesContentPartAddedEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesContentPartAddedEvent
	if v.Type == "" {
		v.Type = "response.content_part.added"
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
		case "content_index", "item_id", "output_index", "part", "sequence_number", "type":
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
func (v *ResponsesContentPartAddedEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesContentPartAddedEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content_index")
	delete(fields, "item_id")
	delete(fields, "output_index")
	delete(fields, "part")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesContentPartAddedEvent(decoded)
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

type ResponsesContentPartDoneEvent struct {
	AdditionalProperties map[string]any                    `json:"-"`
	ContentIndex         int                               `json:"content_index"`
	ItemID               string                            `json:"item_id"`
	OutputIndex          int                               `json:"output_index"`
	Part                 ResponsesContentPartDoneEventPart `json:"part"`
	SequenceNumber       int                               `json:"sequence_number"`
	Type                 string                            `json:"type,omitempty"`
}

func (v ResponsesContentPartDoneEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesContentPartDoneEvent
	if v.Type == "" {
		v.Type = "response.content_part.done"
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
		case "content_index", "item_id", "output_index", "part", "sequence_number", "type":
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
func (v *ResponsesContentPartDoneEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesContentPartDoneEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content_index")
	delete(fields, "item_id")
	delete(fields, "output_index")
	delete(fields, "part")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesContentPartDoneEvent(decoded)
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

type ResponsesCreatedEvent struct {
	AdditionalProperties map[string]any         `json:"-"`
	Response             CreateResponseResponse `json:"response"`
	SequenceNumber       int                    `json:"sequence_number"`
	Type                 string                 `json:"type,omitempty"`
}

func (v ResponsesCreatedEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesCreatedEvent
	if v.Type == "" {
		v.Type = "response.created"
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
		case "response", "sequence_number", "type":
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
func (v *ResponsesCreatedEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesCreatedEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "response")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesCreatedEvent(decoded)
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

type ResponsesGrammar struct {
	AdditionalProperties map[string]any `json:"-"`
	Definition           string         `json:"definition"`
	Syntax               Syntax         `json:"syntax"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesGrammar) MarshalJSON() ([]byte, error) {
	type wire ResponsesGrammar
	if v.Type == "" {
		v.Type = "grammar"
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
		case "definition", "syntax", "type":
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
func (v *ResponsesGrammar) UnmarshalJSON(data []byte) error {
	type wire ResponsesGrammar
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "definition")
	delete(fields, "syntax")
	delete(fields, "type")
	*v = ResponsesGrammar(decoded)
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

type ResponsesText struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesText) MarshalJSON() ([]byte, error) {
	type wire ResponsesText
	if v.Type == "" {
		v.Type = "text"
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
func (v *ResponsesText) UnmarshalJSON(data []byte) error {
	type wire ResponsesText
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = ResponsesText(decoded)
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

type ResponsesCustomTool struct {
	AdditionalProperties map[string]any                     `json:"-"`
	Name                 string                             `json:"name"`
	Type                 string                             `json:"type,omitempty"`
	AllowedCallers       []ResponsesCustomToolAllowedCaller `json:"allowed_callers,omitempty"`
	DeferLoading         *bool                              `json:"defer_loading,omitempty"`
	Description          *string                            `json:"description,omitempty"`
	Format               *ResponsesCustomToolFormat         `json:"format,omitempty"`
}

func (v ResponsesCustomTool) MarshalJSON() ([]byte, error) {
	type wire ResponsesCustomTool
	if v.Type == "" {
		v.Type = "custom"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.AllowedCallers != nil {
		data, err := json.Marshal(v.AllowedCallers)
		if err != nil {
			return nil, err
		}
		fields["allowed_callers"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "name", "type", "allowed_callers", "defer_loading", "description", "format":
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
func (v *ResponsesCustomTool) UnmarshalJSON(data []byte) error {
	type wire ResponsesCustomTool
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "type")
	delete(fields, "allowed_callers")
	delete(fields, "defer_loading")
	delete(fields, "description")
	delete(fields, "format")
	*v = ResponsesCustomTool(decoded)
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

type ResponsesEasyInputMessage struct {
	AdditionalProperties map[string]any                    `json:"-"`
	Content              ResponsesEasyInputMessageContent2 `json:"content"`
	Role                 ResponsesEasyInputMessageRole     `json:"role"`
	Phase                *ResponsesEasyInputMessagePhase   `json:"phase,omitempty"`
	Type                 *string                           `json:"type,omitempty"`
}

func (v ResponsesEasyInputMessage) MarshalJSON() ([]byte, error) {
	type wire ResponsesEasyInputMessage
	if v.Type == nil {
		x := "message"
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
		case "content", "role", "phase", "type":
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
func (v *ResponsesEasyInputMessage) UnmarshalJSON(data []byte) error {
	type wire ResponsesEasyInputMessage
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content")
	delete(fields, "role")
	delete(fields, "phase")
	delete(fields, "type")
	*v = ResponsesEasyInputMessage(decoded)
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

type ResponsesErrorEvent struct {
	AdditionalProperties map[string]any `json:"-"`
	Message              string         `json:"message"`
	SequenceNumber       int            `json:"sequence_number"`
	Code                 *string        `json:"code,omitempty"`
	Param                *string        `json:"param,omitempty"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesErrorEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesErrorEvent
	if v.Type == "" {
		v.Type = "error"
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
		case "message", "sequence_number", "code", "param", "type":
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
func (v *ResponsesErrorEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesErrorEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "message")
	delete(fields, "sequence_number")
	delete(fields, "code")
	delete(fields, "param")
	delete(fields, "type")
	*v = ResponsesErrorEvent(decoded)
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

type ResponsesFailedEvent struct {
	AdditionalProperties map[string]any         `json:"-"`
	Response             CreateResponseResponse `json:"response"`
	SequenceNumber       int                    `json:"sequence_number"`
	Type                 string                 `json:"type,omitempty"`
}

func (v ResponsesFailedEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesFailedEvent
	if v.Type == "" {
		v.Type = "response.failed"
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
		case "response", "sequence_number", "type":
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
func (v *ResponsesFailedEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesFailedEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "response")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesFailedEvent(decoded)
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

type ResponsesFunctionCallArgumentsDeltaEvent struct {
	AdditionalProperties map[string]any `json:"-"`
	Delta                string         `json:"delta"`
	ItemID               string         `json:"item_id"`
	OutputIndex          int            `json:"output_index"`
	SequenceNumber       int            `json:"sequence_number"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesFunctionCallArgumentsDeltaEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionCallArgumentsDeltaEvent
	if v.Type == "" {
		v.Type = "response.function_call_arguments.delta"
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
		case "delta", "item_id", "output_index", "sequence_number", "type":
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
func (v *ResponsesFunctionCallArgumentsDeltaEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionCallArgumentsDeltaEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "delta")
	delete(fields, "item_id")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesFunctionCallArgumentsDeltaEvent(decoded)
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

type ResponsesFunctionCallArgumentsDoneEvent struct {
	AdditionalProperties map[string]any `json:"-"`
	Arguments            string         `json:"arguments"`
	ItemID               string         `json:"item_id"`
	Name                 string         `json:"name"`
	OutputIndex          int            `json:"output_index"`
	SequenceNumber       int            `json:"sequence_number"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesFunctionCallArgumentsDoneEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionCallArgumentsDoneEvent
	if v.Type == "" {
		v.Type = "response.function_call_arguments.done"
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
		case "arguments", "item_id", "name", "output_index", "sequence_number", "type":
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
func (v *ResponsesFunctionCallArgumentsDoneEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionCallArgumentsDoneEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "arguments")
	delete(fields, "item_id")
	delete(fields, "name")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesFunctionCallArgumentsDoneEvent(decoded)
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

type ResponsesFunctionCallOutput struct {
	AdditionalProperties map[string]any                     `json:"-"`
	CallID               string                             `json:"call_id"`
	Output               ResponsesFunctionCallOutputOutput2 `json:"output"`
	Type                 string                             `json:"type,omitempty"`
	ID                   *string                            `json:"id,omitempty"`
	Caller               *ResponsesFunctionCallOutputCaller `json:"caller,omitempty"`
	Status               *ResponsesFunctionCallOutputStatus `json:"status,omitempty"`
}

func (v ResponsesFunctionCallOutput) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionCallOutput
	if v.Type == "" {
		v.Type = "function_call_output"
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
		case "call_id", "output", "type", "id", "caller", "status":
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
func (v *ResponsesFunctionCallOutput) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionCallOutput
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "call_id")
	delete(fields, "output")
	delete(fields, "type")
	delete(fields, "id")
	delete(fields, "caller")
	delete(fields, "status")
	*v = ResponsesFunctionCallOutput(decoded)
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

type ResponsesFunctionCallOutputCallerDirect struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesFunctionCallOutputCallerDirect) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionCallOutputCallerDirect
	if v.Type == "" {
		v.Type = "direct"
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
func (v *ResponsesFunctionCallOutputCallerDirect) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionCallOutputCallerDirect
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = ResponsesFunctionCallOutputCallerDirect(decoded)
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

type ResponsesFunctionCallOutputCallerProgram struct {
	AdditionalProperties map[string]any `json:"-"`
	CallerID             string         `json:"caller_id"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesFunctionCallOutputCallerProgram) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionCallOutputCallerProgram
	if v.Type == "" {
		v.Type = "program"
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
		case "caller_id", "type":
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
func (v *ResponsesFunctionCallOutputCallerProgram) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionCallOutputCallerProgram
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "caller_id")
	delete(fields, "type")
	*v = ResponsesFunctionCallOutputCallerProgram(decoded)
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

type ResponsesFunctionTool struct {
	AdditionalProperties map[string]any                       `json:"-"`
	Name                 string                               `json:"name"`
	Parameters           map[string]any                       `json:"parameters,omitempty"`
	Strict               *bool                                `json:"strict,omitempty"`
	Type                 string                               `json:"type,omitempty"`
	AllowedCallers       []ResponsesFunctionToolAllowedCaller `json:"allowed_callers,omitempty"`
	DeferLoading         *bool                                `json:"defer_loading,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	OutputSchema         map[string]any                       `json:"output_schema,omitempty"`
}

func (v ResponsesFunctionTool) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionTool
	if v.Type == "" {
		v.Type = "function"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Parameters != nil {
		data, err := json.Marshal(v.Parameters)
		if err != nil {
			return nil, err
		}
		fields["parameters"] = data
	}
	if v.AllowedCallers != nil {
		data, err := json.Marshal(v.AllowedCallers)
		if err != nil {
			return nil, err
		}
		fields["allowed_callers"] = data
	}
	if v.OutputSchema != nil {
		data, err := json.Marshal(v.OutputSchema)
		if err != nil {
			return nil, err
		}
		fields["output_schema"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "name", "parameters", "strict", "type", "allowed_callers", "defer_loading", "description", "output_schema":
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
func (v *ResponsesFunctionTool) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionTool
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "parameters")
	delete(fields, "strict")
	delete(fields, "type")
	delete(fields, "allowed_callers")
	delete(fields, "defer_loading")
	delete(fields, "description")
	delete(fields, "output_schema")
	*v = ResponsesFunctionTool(decoded)
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

type ResponsesCallerProgram struct {
	AdditionalProperties map[string]any `json:"-"`
	CallerID             string         `json:"caller_id"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesCallerProgram) MarshalJSON() ([]byte, error) {
	type wire ResponsesCallerProgram
	if v.Type == "" {
		v.Type = "program"
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
		case "caller_id", "type":
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
func (v *ResponsesCallerProgram) UnmarshalJSON(data []byte) error {
	type wire ResponsesCallerProgram
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "caller_id")
	delete(fields, "type")
	*v = ResponsesCallerProgram(decoded)
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

type ResponsesCallerDirect struct {
	AdditionalProperties map[string]any `json:"-"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesCallerDirect) MarshalJSON() ([]byte, error) {
	type wire ResponsesCallerDirect
	if v.Type == "" {
		v.Type = "direct"
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
func (v *ResponsesCallerDirect) UnmarshalJSON(data []byte) error {
	type wire ResponsesCallerDirect
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	*v = ResponsesCallerDirect(decoded)
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

type ResponsesFunctionToolCall struct {
	AdditionalProperties map[string]any                   `json:"-"`
	Arguments            string                           `json:"arguments"`
	CallID               string                           `json:"call_id"`
	Name                 string                           `json:"name"`
	Type                 string                           `json:"type,omitempty"`
	ID                   *string                          `json:"id,omitempty"`
	Caller               *ResponsesFunctionToolCallCaller `json:"caller,omitempty"`
	Namespace            *string                          `json:"namespace,omitempty"`
	Status               *ResponsesFunctionToolCallStatus `json:"status,omitempty"`
}

func (v ResponsesFunctionToolCall) MarshalJSON() ([]byte, error) {
	type wire ResponsesFunctionToolCall
	if v.Type == "" {
		v.Type = "function_call"
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
		case "arguments", "call_id", "name", "type", "id", "caller", "namespace", "status":
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
func (v *ResponsesFunctionToolCall) UnmarshalJSON(data []byte) error {
	type wire ResponsesFunctionToolCall
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "arguments")
	delete(fields, "call_id")
	delete(fields, "name")
	delete(fields, "type")
	delete(fields, "id")
	delete(fields, "caller")
	delete(fields, "namespace")
	delete(fields, "status")
	*v = ResponsesFunctionToolCall(decoded)
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

type ResponsesIncompleteEvent struct {
	AdditionalProperties map[string]any         `json:"-"`
	Response             CreateResponseResponse `json:"response"`
	SequenceNumber       int                    `json:"sequence_number"`
	Type                 string                 `json:"type,omitempty"`
}

func (v ResponsesIncompleteEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesIncompleteEvent
	if v.Type == "" {
		v.Type = "response.incomplete"
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
		case "response", "sequence_number", "type":
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
func (v *ResponsesIncompleteEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesIncompleteEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "response")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesIncompleteEvent(decoded)
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

type ResponsesInProgressEvent struct {
	AdditionalProperties map[string]any         `json:"-"`
	Response             CreateResponseResponse `json:"response"`
	SequenceNumber       int                    `json:"sequence_number"`
	Type                 string                 `json:"type,omitempty"`
}

func (v ResponsesInProgressEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesInProgressEvent
	if v.Type == "" {
		v.Type = "response.in_progress"
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
		case "response", "sequence_number", "type":
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
func (v *ResponsesInProgressEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesInProgressEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "response")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesInProgressEvent(decoded)
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

type ResponsesInputFileResponsesPromptCacheBreakpoint struct {
	AdditionalProperties map[string]any `json:"-"`
	Mode                 string         `json:"mode,omitempty"`
}

func (v ResponsesInputFileResponsesPromptCacheBreakpoint) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputFileResponsesPromptCacheBreakpoint
	if v.Mode == "" {
		v.Mode = "explicit"
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
		case "mode":
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
func (v *ResponsesInputFileResponsesPromptCacheBreakpoint) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputFileResponsesPromptCacheBreakpoint
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "mode")
	*v = ResponsesInputFileResponsesPromptCacheBreakpoint(decoded)
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

type ResponsesInputFile struct {
	AdditionalProperties  map[string]any                                    `json:"-"`
	Type                  string                                            `json:"type,omitempty"`
	Detail                *ResponsesInputFileDetail                         `json:"detail,omitempty"`
	FileData              *string                                           `json:"file_data,omitempty"`
	FileID                *string                                           `json:"file_id,omitempty"`
	FileURL               *string                                           `json:"file_url,omitempty"`
	Filename              *string                                           `json:"filename,omitempty"`
	PromptCacheBreakpoint *ResponsesInputFileResponsesPromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

func (v ResponsesInputFile) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputFile
	if v.Type == "" {
		v.Type = "input_file"
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
		case "type", "detail", "file_data", "file_id", "file_url", "filename", "prompt_cache_breakpoint":
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
func (v *ResponsesInputFile) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputFile
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	delete(fields, "detail")
	delete(fields, "file_data")
	delete(fields, "file_id")
	delete(fields, "file_url")
	delete(fields, "filename")
	delete(fields, "prompt_cache_breakpoint")
	*v = ResponsesInputFile(decoded)
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

type ResponsesInputFileContentResponsesPromptCacheBreakpoint struct {
	AdditionalProperties map[string]any `json:"-"`
	Mode                 string         `json:"mode,omitempty"`
}

func (v ResponsesInputFileContentResponsesPromptCacheBreakpoint) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputFileContentResponsesPromptCacheBreakpoint
	if v.Mode == "" {
		v.Mode = "explicit"
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
		case "mode":
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
func (v *ResponsesInputFileContentResponsesPromptCacheBreakpoint) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputFileContentResponsesPromptCacheBreakpoint
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "mode")
	*v = ResponsesInputFileContentResponsesPromptCacheBreakpoint(decoded)
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

type ResponsesInputFileContent struct {
	AdditionalProperties  map[string]any                                           `json:"-"`
	Type                  string                                                   `json:"type,omitempty"`
	Detail                *ResponsesInputFileContentDetail                         `json:"detail,omitempty"`
	FileData              *string                                                  `json:"file_data,omitempty"`
	FileID                *string                                                  `json:"file_id,omitempty"`
	FileURL               *string                                                  `json:"file_url,omitempty"`
	Filename              *string                                                  `json:"filename,omitempty"`
	PromptCacheBreakpoint *ResponsesInputFileContentResponsesPromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

func (v ResponsesInputFileContent) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputFileContent
	if v.Type == "" {
		v.Type = "input_file"
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
		case "type", "detail", "file_data", "file_id", "file_url", "filename", "prompt_cache_breakpoint":
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
func (v *ResponsesInputFileContent) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputFileContent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	delete(fields, "detail")
	delete(fields, "file_data")
	delete(fields, "file_id")
	delete(fields, "file_url")
	delete(fields, "filename")
	delete(fields, "prompt_cache_breakpoint")
	*v = ResponsesInputFileContent(decoded)
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

type ResponsesInputImageResponsesPromptCacheBreakpoint struct {
	AdditionalProperties map[string]any `json:"-"`
	Mode                 string         `json:"mode,omitempty"`
}

func (v ResponsesInputImageResponsesPromptCacheBreakpoint) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputImageResponsesPromptCacheBreakpoint
	if v.Mode == "" {
		v.Mode = "explicit"
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
		case "mode":
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
func (v *ResponsesInputImageResponsesPromptCacheBreakpoint) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputImageResponsesPromptCacheBreakpoint
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "mode")
	*v = ResponsesInputImageResponsesPromptCacheBreakpoint(decoded)
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

type ResponsesInputImage struct {
	AdditionalProperties  map[string]any                                     `json:"-"`
	Detail                ResponsesInputImageDetail                          `json:"detail"`
	Type                  string                                             `json:"type,omitempty"`
	FileID                *string                                            `json:"file_id,omitempty"`
	ImageURL              *string                                            `json:"image_url,omitempty"`
	PromptCacheBreakpoint *ResponsesInputImageResponsesPromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

func (v ResponsesInputImage) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputImage
	if v.Type == "" {
		v.Type = "input_image"
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
		case "detail", "type", "file_id", "image_url", "prompt_cache_breakpoint":
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
func (v *ResponsesInputImage) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputImage
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "detail")
	delete(fields, "type")
	delete(fields, "file_id")
	delete(fields, "image_url")
	delete(fields, "prompt_cache_breakpoint")
	*v = ResponsesInputImage(decoded)
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

type ResponsesInputImageContentResponsesPromptCacheBreakpoint struct {
	AdditionalProperties map[string]any `json:"-"`
	Mode                 string         `json:"mode,omitempty"`
}

func (v ResponsesInputImageContentResponsesPromptCacheBreakpoint) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputImageContentResponsesPromptCacheBreakpoint
	if v.Mode == "" {
		v.Mode = "explicit"
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
		case "mode":
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
func (v *ResponsesInputImageContentResponsesPromptCacheBreakpoint) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputImageContentResponsesPromptCacheBreakpoint
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "mode")
	*v = ResponsesInputImageContentResponsesPromptCacheBreakpoint(decoded)
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

type ResponsesInputImageContent struct {
	AdditionalProperties  map[string]any                                            `json:"-"`
	Type                  string                                                    `json:"type,omitempty"`
	Detail                *ResponsesInputImageContentDetail                         `json:"detail,omitempty"`
	FileID                *string                                                   `json:"file_id,omitempty"`
	ImageURL              *string                                                   `json:"image_url,omitempty"`
	PromptCacheBreakpoint *ResponsesInputImageContentResponsesPromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

func (v ResponsesInputImageContent) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputImageContent
	if v.Type == "" {
		v.Type = "input_image"
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
		case "type", "detail", "file_id", "image_url", "prompt_cache_breakpoint":
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
func (v *ResponsesInputImageContent) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputImageContent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "type")
	delete(fields, "detail")
	delete(fields, "file_id")
	delete(fields, "image_url")
	delete(fields, "prompt_cache_breakpoint")
	*v = ResponsesInputImageContent(decoded)
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

type ResponsesInputTextResponsesPromptCacheBreakpoint struct {
	AdditionalProperties map[string]any `json:"-"`
	Mode                 string         `json:"mode,omitempty"`
}

func (v ResponsesInputTextResponsesPromptCacheBreakpoint) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputTextResponsesPromptCacheBreakpoint
	if v.Mode == "" {
		v.Mode = "explicit"
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
		case "mode":
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
func (v *ResponsesInputTextResponsesPromptCacheBreakpoint) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputTextResponsesPromptCacheBreakpoint
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "mode")
	*v = ResponsesInputTextResponsesPromptCacheBreakpoint(decoded)
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

type ResponsesInputText struct {
	AdditionalProperties  map[string]any                                    `json:"-"`
	Text                  string                                            `json:"text"`
	Type                  string                                            `json:"type,omitempty"`
	PromptCacheBreakpoint *ResponsesInputTextResponsesPromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

func (v ResponsesInputText) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputText
	if v.Type == "" {
		v.Type = "input_text"
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
		case "text", "type", "prompt_cache_breakpoint":
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
func (v *ResponsesInputText) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputText
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
	delete(fields, "prompt_cache_breakpoint")
	*v = ResponsesInputText(decoded)
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

type ResponsesInputTextContentResponsesPromptCacheBreakpoint struct {
	AdditionalProperties map[string]any `json:"-"`
	Mode                 string         `json:"mode,omitempty"`
}

func (v ResponsesInputTextContentResponsesPromptCacheBreakpoint) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputTextContentResponsesPromptCacheBreakpoint
	if v.Mode == "" {
		v.Mode = "explicit"
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
		case "mode":
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
func (v *ResponsesInputTextContentResponsesPromptCacheBreakpoint) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputTextContentResponsesPromptCacheBreakpoint
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "mode")
	*v = ResponsesInputTextContentResponsesPromptCacheBreakpoint(decoded)
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

type ResponsesInputTextContent struct {
	AdditionalProperties  map[string]any                                           `json:"-"`
	Text                  string                                                   `json:"text"`
	Type                  string                                                   `json:"type,omitempty"`
	PromptCacheBreakpoint *ResponsesInputTextContentResponsesPromptCacheBreakpoint `json:"prompt_cache_breakpoint,omitempty"`
}

func (v ResponsesInputTextContent) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputTextContent
	if v.Type == "" {
		v.Type = "input_text"
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
		case "text", "type", "prompt_cache_breakpoint":
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
func (v *ResponsesInputTextContent) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputTextContent
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
	delete(fields, "prompt_cache_breakpoint")
	*v = ResponsesInputTextContent(decoded)
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

type ResponsesInputTokensDetails struct {
	AdditionalProperties map[string]any `json:"-"`
	CacheWriteTokens     int            `json:"cache_write_tokens"`
	CachedTokens         int            `json:"cached_tokens"`
}

func (v ResponsesInputTokensDetails) MarshalJSON() ([]byte, error) {
	type wire ResponsesInputTokensDetails
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
		case "cache_write_tokens", "cached_tokens":
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
func (v *ResponsesInputTokensDetails) UnmarshalJSON(data []byte) error {
	type wire ResponsesInputTokensDetails
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "cache_write_tokens")
	delete(fields, "cached_tokens")
	*v = ResponsesInputTokensDetails(decoded)
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

type ResponsesMessage struct {
	AdditionalProperties map[string]any            `json:"-"`
	Content              []ResponsesMessageContent `json:"content"`
	Role                 ResponsesMessageRole      `json:"role"`
	Status               *ResponsesMessageStatus   `json:"status,omitempty"`
	Type                 *string                   `json:"type,omitempty"`
}

func (v ResponsesMessage) MarshalJSON() ([]byte, error) {
	type wire ResponsesMessage
	if v.Type == nil {
		x := "message"
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
		case "content", "role", "status", "type":
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
func (v *ResponsesMessage) UnmarshalJSON(data []byte) error {
	type wire ResponsesMessage
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content")
	delete(fields, "role")
	delete(fields, "status")
	delete(fields, "type")
	*v = ResponsesMessage(decoded)
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

type ResponsesNamespaceTool struct {
	AdditionalProperties map[string]any               `json:"-"`
	Description          string                       `json:"description"`
	Name                 string                       `json:"name"`
	Tools                []ResponsesNamespaceToolTool `json:"tools"`
	Type                 string                       `json:"type,omitempty"`
}

func (v ResponsesNamespaceTool) MarshalJSON() ([]byte, error) {
	type wire ResponsesNamespaceTool
	if v.Type == "" {
		v.Type = "namespace"
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
		case "description", "name", "tools", "type":
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
func (v *ResponsesNamespaceTool) UnmarshalJSON(data []byte) error {
	type wire ResponsesNamespaceTool
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "description")
	delete(fields, "name")
	delete(fields, "tools")
	delete(fields, "type")
	*v = ResponsesNamespaceTool(decoded)
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

type ResponsesOutputItemAddedEvent struct {
	AdditionalProperties map[string]any                    `json:"-"`
	Item                 ResponsesOutputItemAddedEventItem `json:"item"`
	OutputIndex          int                               `json:"output_index"`
	SequenceNumber       int                               `json:"sequence_number"`
	Type                 string                            `json:"type,omitempty"`
}

func (v ResponsesOutputItemAddedEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputItemAddedEvent
	if v.Type == "" {
		v.Type = "response.output_item.added"
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
		case "item", "output_index", "sequence_number", "type":
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
func (v *ResponsesOutputItemAddedEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputItemAddedEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "item")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesOutputItemAddedEvent(decoded)
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

type ResponsesOutputItemDoneEvent struct {
	AdditionalProperties map[string]any                   `json:"-"`
	Item                 ResponsesOutputItemDoneEventItem `json:"item"`
	OutputIndex          int                              `json:"output_index"`
	SequenceNumber       int                              `json:"sequence_number"`
	Type                 string                           `json:"type,omitempty"`
}

func (v ResponsesOutputItemDoneEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputItemDoneEvent
	if v.Type == "" {
		v.Type = "response.output_item.done"
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
		case "item", "output_index", "sequence_number", "type":
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
func (v *ResponsesOutputItemDoneEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputItemDoneEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "item")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesOutputItemDoneEvent(decoded)
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

type ResponsesOutputMessage struct {
	AdditionalProperties map[string]any                  `json:"-"`
	ID                   string                          `json:"id"`
	Content              []ResponsesOutputMessageContent `json:"content"`
	Status               ResponsesOutputMessageStatus    `json:"status"`
	Role                 string                          `json:"role,omitempty"`
	Type                 string                          `json:"type,omitempty"`
	Phase                *ResponsesOutputMessagePhase    `json:"phase,omitempty"`
}

func (v ResponsesOutputMessage) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputMessage
	if v.Role == "" {
		v.Role = "assistant"
	}
	if v.Type == "" {
		v.Type = "message"
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
		case "id", "content", "status", "role", "type", "phase":
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
func (v *ResponsesOutputMessage) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputMessage
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "id")
	delete(fields, "content")
	delete(fields, "status")
	delete(fields, "role")
	delete(fields, "type")
	delete(fields, "phase")
	*v = ResponsesOutputMessage(decoded)
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

type ResponsesOutputRefusal struct {
	AdditionalProperties map[string]any `json:"-"`
	Refusal              string         `json:"refusal"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesOutputRefusal) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputRefusal
	if v.Type == "" {
		v.Type = "refusal"
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
		case "refusal", "type":
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
func (v *ResponsesOutputRefusal) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputRefusal
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "refusal")
	delete(fields, "type")
	*v = ResponsesOutputRefusal(decoded)
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

type ResponsesOutputText struct {
	AdditionalProperties map[string]any               `json:"-"`
	Annotations          []Annotation                 `json:"annotations"`
	Text                 string                       `json:"text"`
	Type                 string                       `json:"type,omitempty"`
	Logprobs             []ResponsesOutputTextLogprob `json:"logprobs,omitempty"`
}

func (v ResponsesOutputText) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputText
	if v.Type == "" {
		v.Type = "output_text"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Logprobs != nil {
		data, err := json.Marshal(v.Logprobs)
		if err != nil {
			return nil, err
		}
		fields["logprobs"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "annotations", "text", "type", "logprobs":
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
func (v *ResponsesOutputText) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputText
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "annotations")
	delete(fields, "text")
	delete(fields, "type")
	delete(fields, "logprobs")
	*v = ResponsesOutputText(decoded)
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

type ResponsesOutputTextLogprob struct {
	AdditionalProperties map[string]any                         `json:"-"`
	Token                string                                 `json:"token"`
	Bytes                []int                                  `json:"bytes"`
	Logprob              float64                                `json:"logprob"`
	TopLogprobs          []ResponsesOutputTextLogprobTopLogprob `json:"top_logprobs"`
}

func (v ResponsesOutputTextLogprob) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputTextLogprob
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
		case "token", "bytes", "logprob", "top_logprobs":
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
func (v *ResponsesOutputTextLogprob) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputTextLogprob
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "token")
	delete(fields, "bytes")
	delete(fields, "logprob")
	delete(fields, "top_logprobs")
	*v = ResponsesOutputTextLogprob(decoded)
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

type ResponsesOutputTextLogprobTopLogprob struct {
	AdditionalProperties map[string]any `json:"-"`
	Token                string         `json:"token"`
	Bytes                []int          `json:"bytes"`
	Logprob              float64        `json:"logprob"`
}

func (v ResponsesOutputTextLogprobTopLogprob) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputTextLogprobTopLogprob
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
		case "token", "bytes", "logprob":
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
func (v *ResponsesOutputTextLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputTextLogprobTopLogprob
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "token")
	delete(fields, "bytes")
	delete(fields, "logprob")
	*v = ResponsesOutputTextLogprobTopLogprob(decoded)
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

type ResponsesOutputTokensDetails struct {
	AdditionalProperties map[string]any `json:"-"`
	ReasoningTokens      int            `json:"reasoning_tokens"`
}

func (v ResponsesOutputTokensDetails) MarshalJSON() ([]byte, error) {
	type wire ResponsesOutputTokensDetails
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
		case "reasoning_tokens":
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
func (v *ResponsesOutputTokensDetails) UnmarshalJSON(data []byte) error {
	type wire ResponsesOutputTokensDetails
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "reasoning_tokens")
	*v = ResponsesOutputTokensDetails(decoded)
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

type ResponsesPartReasoningText struct {
	AdditionalProperties map[string]any `json:"-"`
	Text                 string         `json:"text"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesPartReasoningText) MarshalJSON() ([]byte, error) {
	type wire ResponsesPartReasoningText
	if v.Type == "" {
		v.Type = "reasoning_text"
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
func (v *ResponsesPartReasoningText) UnmarshalJSON(data []byte) error {
	type wire ResponsesPartReasoningText
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
	*v = ResponsesPartReasoningText(decoded)
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

type ResponsesReasoningItem struct {
	AdditionalProperties map[string]any                `json:"-"`
	ID                   string                        `json:"id"`
	Summary              []ResponsesSummary            `json:"summary"`
	Type                 string                        `json:"type,omitempty"`
	Content              []ResponsesContent            `json:"content,omitempty"`
	EncryptedContent     *string                       `json:"encrypted_content,omitempty"`
	Status               *ResponsesReasoningItemStatus `json:"status,omitempty"`
}

func (v ResponsesReasoningItem) MarshalJSON() ([]byte, error) {
	type wire ResponsesReasoningItem
	if v.Type == "" {
		v.Type = "reasoning"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Content != nil {
		data, err := json.Marshal(v.Content)
		if err != nil {
			return nil, err
		}
		fields["content"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "id", "summary", "type", "content", "encrypted_content", "status":
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
func (v *ResponsesReasoningItem) UnmarshalJSON(data []byte) error {
	type wire ResponsesReasoningItem
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "id")
	delete(fields, "summary")
	delete(fields, "type")
	delete(fields, "content")
	delete(fields, "encrypted_content")
	delete(fields, "status")
	*v = ResponsesReasoningItem(decoded)
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

type ResponsesReasoningTextDeltaEvent struct {
	AdditionalProperties map[string]any `json:"-"`
	ContentIndex         int            `json:"content_index"`
	Delta                string         `json:"delta"`
	ItemID               string         `json:"item_id"`
	OutputIndex          int            `json:"output_index"`
	SequenceNumber       int            `json:"sequence_number"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesReasoningTextDeltaEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesReasoningTextDeltaEvent
	if v.Type == "" {
		v.Type = "response.reasoning_text.delta"
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
		case "content_index", "delta", "item_id", "output_index", "sequence_number", "type":
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
func (v *ResponsesReasoningTextDeltaEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesReasoningTextDeltaEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content_index")
	delete(fields, "delta")
	delete(fields, "item_id")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesReasoningTextDeltaEvent(decoded)
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

type ResponsesReasoningTextDoneEvent struct {
	AdditionalProperties map[string]any `json:"-"`
	ContentIndex         int            `json:"content_index"`
	ItemID               string         `json:"item_id"`
	OutputIndex          int            `json:"output_index"`
	SequenceNumber       int            `json:"sequence_number"`
	Text                 string         `json:"text"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesReasoningTextDoneEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesReasoningTextDoneEvent
	if v.Type == "" {
		v.Type = "response.reasoning_text.done"
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
		case "content_index", "item_id", "output_index", "sequence_number", "text", "type":
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
func (v *ResponsesReasoningTextDoneEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesReasoningTextDoneEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content_index")
	delete(fields, "item_id")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "text")
	delete(fields, "type")
	*v = ResponsesReasoningTextDoneEvent(decoded)
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

type ResponsesStreamFrame struct {
	Event string                   `json:"event"`
	Data  ResponsesStreamFrameData `json:"data"`
}
type ResponsesSummary struct {
	AdditionalProperties map[string]any `json:"-"`
	Text                 string         `json:"text"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesSummary) MarshalJSON() ([]byte, error) {
	type wire ResponsesSummary
	if v.Type == "" {
		v.Type = "summary_text"
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
func (v *ResponsesSummary) UnmarshalJSON(data []byte) error {
	type wire ResponsesSummary
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
	*v = ResponsesSummary(decoded)
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

type ResponsesTextDeltaEvent struct {
	AdditionalProperties map[string]any                   `json:"-"`
	ContentIndex         int                              `json:"content_index"`
	Delta                string                           `json:"delta"`
	ItemID               string                           `json:"item_id"`
	Logprobs             []ResponsesTextDeltaEventLogprob `json:"logprobs"`
	OutputIndex          int                              `json:"output_index"`
	SequenceNumber       int                              `json:"sequence_number"`
	Type                 string                           `json:"type,omitempty"`
}

func (v ResponsesTextDeltaEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextDeltaEvent
	if v.Type == "" {
		v.Type = "response.output_text.delta"
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
		case "content_index", "delta", "item_id", "logprobs", "output_index", "sequence_number", "type":
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
func (v *ResponsesTextDeltaEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextDeltaEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content_index")
	delete(fields, "delta")
	delete(fields, "item_id")
	delete(fields, "logprobs")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "type")
	*v = ResponsesTextDeltaEvent(decoded)
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

type ResponsesTextDeltaEventLogprob struct {
	AdditionalProperties map[string]any                             `json:"-"`
	Token                string                                     `json:"token"`
	Logprob              float64                                    `json:"logprob"`
	TopLogprobs          []ResponsesTextDeltaEventLogprobTopLogprob `json:"top_logprobs,omitempty"`
}

func (v ResponsesTextDeltaEventLogprob) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextDeltaEventLogprob
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.TopLogprobs != nil {
		data, err := json.Marshal(v.TopLogprobs)
		if err != nil {
			return nil, err
		}
		fields["top_logprobs"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "token", "logprob", "top_logprobs":
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
func (v *ResponsesTextDeltaEventLogprob) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextDeltaEventLogprob
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "token")
	delete(fields, "logprob")
	delete(fields, "top_logprobs")
	*v = ResponsesTextDeltaEventLogprob(decoded)
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

type ResponsesTextDeltaEventLogprobTopLogprob struct {
	AdditionalProperties map[string]any `json:"-"`
	Token                *string        `json:"token,omitempty"`
	Logprob              *float64       `json:"logprob,omitempty"`
}

func (v ResponsesTextDeltaEventLogprobTopLogprob) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextDeltaEventLogprobTopLogprob
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
		case "token", "logprob":
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
func (v *ResponsesTextDeltaEventLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextDeltaEventLogprobTopLogprob
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "token")
	delete(fields, "logprob")
	*v = ResponsesTextDeltaEventLogprobTopLogprob(decoded)
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

type ResponsesTextDoneEvent struct {
	AdditionalProperties map[string]any                  `json:"-"`
	ContentIndex         int                             `json:"content_index"`
	ItemID               string                          `json:"item_id"`
	Logprobs             []ResponsesTextDoneEventLogprob `json:"logprobs"`
	OutputIndex          int                             `json:"output_index"`
	SequenceNumber       int                             `json:"sequence_number"`
	Text                 string                          `json:"text"`
	Type                 string                          `json:"type,omitempty"`
}

func (v ResponsesTextDoneEvent) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextDoneEvent
	if v.Type == "" {
		v.Type = "response.output_text.done"
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
		case "content_index", "item_id", "logprobs", "output_index", "sequence_number", "text", "type":
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
func (v *ResponsesTextDoneEvent) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextDoneEvent
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "content_index")
	delete(fields, "item_id")
	delete(fields, "logprobs")
	delete(fields, "output_index")
	delete(fields, "sequence_number")
	delete(fields, "text")
	delete(fields, "type")
	*v = ResponsesTextDoneEvent(decoded)
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

type ResponsesTextDoneEventLogprob struct {
	AdditionalProperties map[string]any                            `json:"-"`
	Token                string                                    `json:"token"`
	Logprob              float64                                   `json:"logprob"`
	TopLogprobs          []ResponsesTextDoneEventLogprobTopLogprob `json:"top_logprobs,omitempty"`
}

func (v ResponsesTextDoneEventLogprob) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextDoneEventLogprob
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.TopLogprobs != nil {
		data, err := json.Marshal(v.TopLogprobs)
		if err != nil {
			return nil, err
		}
		fields["top_logprobs"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "token", "logprob", "top_logprobs":
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
func (v *ResponsesTextDoneEventLogprob) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextDoneEventLogprob
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "token")
	delete(fields, "logprob")
	delete(fields, "top_logprobs")
	*v = ResponsesTextDoneEventLogprob(decoded)
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

type ResponsesTextDoneEventLogprobTopLogprob struct {
	AdditionalProperties map[string]any `json:"-"`
	Token                *string        `json:"token,omitempty"`
	Logprob              *float64       `json:"logprob,omitempty"`
}

func (v ResponsesTextDoneEventLogprobTopLogprob) MarshalJSON() ([]byte, error) {
	type wire ResponsesTextDoneEventLogprobTopLogprob
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
		case "token", "logprob":
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
func (v *ResponsesTextDoneEventLogprobTopLogprob) UnmarshalJSON(data []byte) error {
	type wire ResponsesTextDoneEventLogprobTopLogprob
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "token")
	delete(fields, "logprob")
	*v = ResponsesTextDoneEventLogprobTopLogprob(decoded)
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

type ResponsesToolChoiceFunction struct {
	AdditionalProperties map[string]any `json:"-"`
	Name                 string         `json:"name"`
	Type                 string         `json:"type,omitempty"`
}

func (v ResponsesToolChoiceFunction) MarshalJSON() ([]byte, error) {
	type wire ResponsesToolChoiceFunction
	if v.Type == "" {
		v.Type = "function"
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
		case "name", "type":
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
func (v *ResponsesToolChoiceFunction) UnmarshalJSON(data []byte) error {
	type wire ResponsesToolChoiceFunction
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "type")
	*v = ResponsesToolChoiceFunction(decoded)
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

type ResponsesToolFunction struct {
	AdditionalProperties map[string]any                       `json:"-"`
	Name                 string                               `json:"name"`
	Type                 string                               `json:"type,omitempty"`
	AllowedCallers       []ResponsesToolFunctionAllowedCaller `json:"allowed_callers,omitempty"`
	DeferLoading         *bool                                `json:"defer_loading,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	OutputSchema         map[string]any                       `json:"output_schema,omitempty"`
	Parameters           any                                  `json:"parameters,omitempty"`
	Strict               *bool                                `json:"strict,omitempty"`
}

func (v ResponsesToolFunction) MarshalJSON() ([]byte, error) {
	type wire ResponsesToolFunction
	if v.Type == "" {
		v.Type = "function"
	}
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.AllowedCallers != nil {
		data, err := json.Marshal(v.AllowedCallers)
		if err != nil {
			return nil, err
		}
		fields["allowed_callers"] = data
	}
	if v.OutputSchema != nil {
		data, err := json.Marshal(v.OutputSchema)
		if err != nil {
			return nil, err
		}
		fields["output_schema"] = data
	}
	for key, value := range v.AdditionalProperties {
		switch key {
		case "name", "type", "allowed_callers", "defer_loading", "description", "output_schema", "parameters", "strict":
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
func (v *ResponsesToolFunction) UnmarshalJSON(data []byte) error {
	type wire ResponsesToolFunction
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	delete(fields, "name")
	delete(fields, "type")
	delete(fields, "allowed_callers")
	delete(fields, "defer_loading")
	delete(fields, "description")
	delete(fields, "output_schema")
	delete(fields, "parameters")
	delete(fields, "strict")
	*v = ResponsesToolFunction(decoded)
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

type RootsCapability struct {
	ListChanged *bool `json:"listChanged,omitempty"`
}
type SamplingCapability struct {
	Context *SamplingContextCapability `json:"context,omitempty"`
	Tools   *SamplingToolsCapability   `json:"tools,omitempty"`
}
type SamplingContextCapability struct {
}
type SamplingToolsCapability struct {
}
type ServerCapabilities struct {
	Experimental map[string]map[string]any `json:"experimental,omitempty"`
	Logging      *LoggingCapability        `json:"logging,omitempty"`
	Prompts      *PromptsCapability        `json:"prompts,omitempty"`
	Resources    *ResourcesCapability      `json:"resources,omitempty"`
	Tools        *ToolsCapability          `json:"tools,omitempty"`
	Completions  *CompletionsCapability    `json:"completions,omitempty"`
	Extensions   map[string]map[string]any `json:"extensions,omitempty"`
	Tasks        *ServerTasksCapability    `json:"tasks,omitempty"`
}

func (v ServerCapabilities) MarshalJSON() ([]byte, error) {
	type wire ServerCapabilities
	data, err := json.Marshal(wire(v))
	if err != nil {
		return nil, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	if v.Experimental != nil {
		data, err := json.Marshal(v.Experimental)
		if err != nil {
			return nil, err
		}
		fields["experimental"] = data
	}
	if v.Extensions != nil {
		data, err := json.Marshal(v.Extensions)
		if err != nil {
			return nil, err
		}
		fields["extensions"] = data
	}
	return json.Marshal(fields)
}

type ServerTasksCapability struct {
	List     *TasksListCapability           `json:"list,omitempty"`
	Cancel   *TasksCancelCapability         `json:"cancel,omitempty"`
	Requests *ServerTasksRequestsCapability `json:"requests,omitempty"`
}
type ServerTasksRequestsCapability struct {
	Tools *TasksToolsCapability `json:"tools,omitempty"`
}
type SharingDelete struct {
	ShareWithUuid string     `json:"share_with_uuid"`
	ShareWithType EntityType `json:"share_with_type"`
	OrgID         *string    `json:"org_id,omitempty"`
}
type SharingRequest struct {
	Level         ShareEnum  `json:"level"`
	ShareWithUuid string     `json:"share_with_uuid"`
	ShareWithType EntityType `json:"share_with_type"`
	OrgID         *string    `json:"org_id,omitempty"`
}
type SipDestination struct {
	Uri  string `json:"uri"`
	Type string `json:"type,omitempty"`
}

func (v SipDestination) MarshalJSON() ([]byte, error) {
	type wire SipDestination
	if v.Type == "" {
		v.Type = "sip"
	}
	return json.Marshal(wire(v))
}

type SpeechAssistantTurn struct {
	Text  string   `json:"text"`
	Audio RawAudio `json:"audio"`
	Role  string   `json:"role,omitempty"`
}

func (v SpeechAssistantTurn) MarshalJSON() ([]byte, error) {
	type wire SpeechAssistantTurn
	if v.Role == "" {
		v.Role = "assistant"
	}
	return json.Marshal(wire(v))
}

type SpeechUserTurn struct {
	Audio RawAudio `json:"audio"`
	Role  string   `json:"role,omitempty"`
}

func (v SpeechUserTurn) MarshalJSON() ([]byte, error) {
	type wire SpeechUserTurn
	if v.Role == "" {
		v.Role = "user"
	}
	return json.Marshal(wire(v))
}

type StaticToolFiller struct {
	Text string  `json:"text"`
	Type *string `json:"type,omitempty"`
}

func (v StaticToolFiller) MarshalJSON() ([]byte, error) {
	type wire StaticToolFiller
	if v.Type == nil {
		x := "static"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type StreamingDelayRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}
type SummaryChunk struct {
	Content string `json:"content"`
}
type SummaryError struct {
	Error string `json:"error"`
}
type SummaryMetadata struct {
	Status     SummaryStatusDollarDefs `json:"status"`
	Translated bool                    `json:"translated"`
}
type TasksCallCapability struct {
}
type TasksCancelCapability struct {
}
type TasksCreateElicitationCapability struct {
}
type TasksCreateMessageCapability struct {
}
type TasksElicitationCapability struct {
	Create *TasksCreateElicitationCapability `json:"create,omitempty"`
}
type TasksListCapability struct {
}
type TasksSamplingCapability struct {
	CreateMessage *TasksCreateMessageCapability `json:"createMessage,omitempty"`
}
type TasksToolsCapability struct {
	Call *TasksCallCapability `json:"call,omitempty"`
}
type TextContent struct {
	Text        string         `json:"text"`
	Type        string         `json:"type,omitempty"`
	Annotations *Annotations   `json:"annotations,omitempty"`
	Meta        map[string]any `json:"_meta,omitempty"`
}

func (v TextContent) MarshalJSON() ([]byte, error) {
	type wire TextContent
	if v.Type == "" {
		v.Type = "text"
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
	return json.Marshal(fields)
}

type TextResourceContents struct {
	Uri      string         `json:"uri"`
	Text     string         `json:"text"`
	MimeType *string        `json:"mimeType,omitempty"`
	Meta     map[string]any `json:"_meta,omitempty"`
}

func (v TextResourceContents) MarshalJSON() ([]byte, error) {
	type wire TextResourceContents
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
	return json.Marshal(fields)
}

type ToolAnnotations struct {
	Title           *string `json:"title,omitempty"`
	ReadOnlyHint    *bool   `json:"readOnlyHint,omitempty"`
	DestructiveHint *bool   `json:"destructiveHint,omitempty"`
	IdempotentHint  *bool   `json:"idempotentHint,omitempty"`
	OpenWorldHint   *bool   `json:"openWorldHint,omitempty"`
}
type ToolChoice struct {
	Function any     `json:"function"`
	Type     *string `json:"type,omitempty"`
}

func (v ToolChoice) MarshalJSON() ([]byte, error) {
	type wire ToolChoice
	if v.Type == nil {
		x := "function"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type ToolExecution struct {
	TaskSupport *TaskSupport `json:"taskSupport,omitempty"`
}
type ToolsCapability struct {
	ListChanged *bool `json:"listChanged,omitempty"`
}
type TransferCallTool struct {
	Destinations []TransferTarget  `json:"destinations"`
	Type         string            `json:"type,omitempty"`
	Filler       *StaticToolFiller `json:"filler,omitempty"`
}

func (v TransferCallTool) MarshalJSON() ([]byte, error) {
	type wire TransferCallTool
	if v.Type == "" {
		v.Type = "transfer_call"
	}
	return json.Marshal(wire(v))
}

type UnknownDestination struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownDestination) MarshalJSON() ([]byte, error) {
	type wire UnknownDestination
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type TransferTarget struct {
	ID          string      `json:"id"`
	Condition   string      `json:"condition"`
	Destination Destination `json:"destination"`
}
type TurnDetectionConfig struct {
	StartThreshold    *float64 `json:"start_threshold,omitempty"`
	EagerEndThreshold *float64 `json:"eager_end_threshold,omitempty"`
	EndThreshold      *float64 `json:"end_threshold,omitempty"`
	MinSilenceMs      *int     `json:"min_silence_ms,omitempty"`
	EndTimeoutMs      *int     `json:"end_timeout_ms,omitempty"`
}

func (v TurnDetectionConfig) MarshalJSON() ([]byte, error) {
	type wire TurnDetectionConfig
	if v.StartThreshold == nil {
		x := float64(0.5)
		v.StartThreshold = &x
	}
	if v.EagerEndThreshold == nil {
		x := float64(0.5)
		v.EagerEndThreshold = &x
	}
	if v.EndThreshold == nil {
		x := float64(0.33)
		v.EndThreshold = &x
	}
	if v.MinSilenceMs == nil {
		x := 500
		v.MinSilenceMs = &x
	}
	return json.Marshal(wire(v))
}

type TurnDetectionEagerEndEvent struct {
	TimeS float64 `json:"time_s"`
	Type  *string `json:"type,omitempty"`
}

func (v TurnDetectionEagerEndEvent) MarshalJSON() ([]byte, error) {
	type wire TurnDetectionEagerEndEvent
	if v.Type == nil {
		x := "turn.eager_end"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type TurnDetectionEndEvent struct {
	TimeS     float64                `json:"time_s"`
	EndReason TurnDetectionEndReason `json:"end_reason"`
	Type      *string                `json:"type,omitempty"`
}

func (v TurnDetectionEndEvent) MarshalJSON() ([]byte, error) {
	type wire TurnDetectionEndEvent
	if v.Type == nil {
		x := "turn.end"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type TurnDetectionResumeEvent struct {
	TimeS float64 `json:"time_s"`
	Type  *string `json:"type,omitempty"`
}

func (v TurnDetectionResumeEvent) MarshalJSON() ([]byte, error) {
	type wire TurnDetectionResumeEvent
	if v.Type == nil {
		x := "turn.resume"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type TurnDetectionStartEvent struct {
	TimeS float64 `json:"time_s"`
	Type  *string `json:"type,omitempty"`
}

func (v TurnDetectionStartEvent) MarshalJSON() ([]byte, error) {
	type wire TurnDetectionStartEvent
	if v.Type == nil {
		x := "turn.start"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type UpdateSummaryRequestSummary struct {
	Content    string                            `json:"content"`
	Status     UpdateSummaryRequestSummaryStatus `json:"status"`
	Translated bool                              `json:"translated"`
}
type URLElicitationCapability struct {
}
type UnknownVoiceAgentCustomConnectorAuthorization struct {
	Type      string `json:"type,omitempty"`
	Raw       any    `json:"raw"`
	IsUnknown bool   `json:"is_unknown,omitempty"`
}

func (v UnknownVoiceAgentCustomConnectorAuthorization) MarshalJSON() ([]byte, error) {
	type wire UnknownVoiceAgentCustomConnectorAuthorization
	if v.Type == "" {
		v.Type = "UNKNOWN"
	}
	return json.Marshal(wire(v))
}

type VoiceAgentCustomConnector struct {
	ConnectorID       string                                  `json:"connector_id"`
	Type              string                                  `json:"type,omitempty"`
	Authorization     *VoiceAgentCustomConnectorAuthorization `json:"authorization,omitempty"`
	ToolConfiguration any                                     `json:"tool_configuration,omitempty"`
	Filler            *StaticToolFiller                       `json:"filler,omitempty"`
}

func (v VoiceAgentCustomConnector) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentCustomConnector
	if v.Type == "" {
		v.Type = "connector"
	}
	return json.Marshal(wire(v))
}

type VoiceAgentDocumentLibraryTool struct {
	LibraryIds        []string          `json:"library_ids"`
	ToolConfiguration any               `json:"tool_configuration,omitempty"`
	Type              string            `json:"type,omitempty"`
	Filler            *StaticToolFiller `json:"filler,omitempty"`
}

func (v VoiceAgentDocumentLibraryTool) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentDocumentLibraryTool
	if v.Type == "" {
		v.Type = "document_library"
	}
	return json.Marshal(wire(v))
}

type VoiceAgentErrorEvent struct {
	Source  VoiceAgentErrorEventSource `json:"source"`
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Fatal   bool                       `json:"fatal"`
	Type    *string                    `json:"type,omitempty"`
	Payload map[string]any             `json:"payload,omitempty"`
}

func (v VoiceAgentErrorEvent) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentErrorEvent
	if v.Type == nil {
		x := "error"
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
	if v.Payload != nil {
		data, err := json.Marshal(v.Payload)
		if err != nil {
			return nil, err
		}
		fields["payload"] = data
	}
	return json.Marshal(fields)
}

type VoiceAgentFunctionTool struct {
	Function  Function                    `json:"function"`
	Type      *VoiceAgentFunctionToolType `json:"type,omitempty"`
	Filler    *StaticToolFiller           `json:"filler,omitempty"`
	TimeoutMs *int                        `json:"timeout_ms,omitempty"`
}

func (v VoiceAgentFunctionTool) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentFunctionTool
	if v.Type == nil {
		x := "function"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type VoiceAgentFunctionToolRequest struct {
	Function  FunctionRequest   `json:"function"`
	Type      string            `json:"type,omitempty"`
	Filler    *StaticToolFiller `json:"filler,omitempty"`
	TimeoutMs *int              `json:"timeout_ms,omitempty"`
}

func (v VoiceAgentFunctionToolRequest) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentFunctionToolRequest
	if v.Type == "" {
		v.Type = "function"
	}
	return json.Marshal(wire(v))
}

type VoiceAgentSessionUpdatedEvent struct {
	Session AudioAgentConfigurationResponseConfig `json:"session"`
	Ts      float64                               `json:"ts"`
	Type    *string                               `json:"type,omitempty"`
}

func (v VoiceAgentSessionUpdatedEvent) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentSessionUpdatedEvent
	if v.Type == nil {
		x := "session.updated"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type VoiceAgentStateChangedEvent struct {
	OldState OldState `json:"old_state"`
	State    State    `json:"state"`
	Ts       float64  `json:"ts"`
	Type     *string  `json:"type,omitempty"`
}

func (v VoiceAgentStateChangedEvent) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentStateChangedEvent
	if v.Type == nil {
		x := "agent.state.changed"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type VoiceAgentToolCallEndedEvent struct {
	ID        string                             `json:"id"`
	Name      string                             `json:"name"`
	Ts        float64                            `json:"ts"`
	Arguments string                             `json:"arguments"`
	Status    VoiceAgentToolCallEndedEventStatus `json:"status"`
	Type      *string                            `json:"type,omitempty"`
	Phase     *string                            `json:"phase,omitempty"`
	Result    *string                            `json:"result,omitempty"`
}

func (v VoiceAgentToolCallEndedEvent) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentToolCallEndedEvent
	if v.Type == nil {
		x := "tool_call.ended"
		v.Type = &x
	}
	if v.Phase == nil {
		x := "ended"
		v.Phase = &x
	}
	return json.Marshal(wire(v))
}

type VoiceAgentToolCallStartedEvent struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Ts    float64 `json:"ts"`
	Type  *string `json:"type,omitempty"`
	Phase *string `json:"phase,omitempty"`
}

func (v VoiceAgentToolCallStartedEvent) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentToolCallStartedEvent
	if v.Type == nil {
		x := "tool_call.started"
		v.Type = &x
	}
	if v.Phase == nil {
		x := "started"
		v.Phase = &x
	}
	return json.Marshal(wire(v))
}

type VoiceAgentTranscriptEvent struct {
	ID        string                        `json:"id"`
	Role      VoiceAgentTranscriptEventRole `json:"role"`
	Content   string                        `json:"content"`
	Final     bool                          `json:"final"`
	Ts        float64                       `json:"ts"`
	Type      *string                       `json:"type,omitempty"`
	StartTime *float64                      `json:"start_time,omitempty"`
	EndTime   *float64                      `json:"end_time,omitempty"`
	Language  *string                       `json:"language,omitempty"`
}

func (v VoiceAgentTranscriptEvent) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentTranscriptEvent
	if v.Type == nil {
		x := "transcription.segment"
		v.Type = &x
	}
	return json.Marshal(wire(v))
}

type VoiceAgentWebSearchTool struct {
	ToolConfiguration any               `json:"tool_configuration,omitempty"`
	Type              string            `json:"type,omitempty"`
	Filler            *StaticToolFiller `json:"filler,omitempty"`
}

func (v VoiceAgentWebSearchTool) MarshalJSON() ([]byte, error) {
	type wire VoiceAgentWebSearchTool
	if v.Type == "" {
		v.Type = "web_search"
	}
	return json.Marshal(wire(v))
}

type VoiceUpdateRequest struct {
	Name            *string  `json:"name,omitempty"`
	Languages       []string `json:"languages,omitempty"`
	Gender          *string  `json:"gender,omitempty"`
	Age             *int     `json:"age,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	Description     *string  `json:"description,omitempty"`
	RetentionNotice *int     `json:"retention_notice,omitempty"`
}

func (v VoiceUpdateRequest) MarshalJSON() ([]byte, error) {
	type wire VoiceUpdateRequest
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
	return json.Marshal(fields)
}
