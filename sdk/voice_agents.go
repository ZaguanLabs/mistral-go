package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// VoiceAgentTransport is the native media/RPC boundary. It carries PCM audio,
// agent events, function invocations and acknowledged session updates.
type VoiceAgentTransport interface {
	Start(context.Context) error
	WaitForAgent(context.Context) error
	SendAudio(context.Context, []byte) error
	SetInputEnabled(context.Context, bool) error
	UpdateSession(context.Context, []byte) ([]byte, error)
	AddFunctionHandlers([]*VoiceFunctionHandler)
	Close() error
}

type VoiceAudioInterface interface {
	Preflight(context.Context) error
	Start(context.Context) error
	InputAudio(context.Context) ([]byte, error)
	OutputAudio(context.Context, []byte) error
	Stop() error
}

type VoiceEvent struct {
	Name string
	Data any
}
type VoiceTransportFactory func(*AudioSessionCreateResponse, func(VoiceEvent), []*VoiceFunctionHandler) (VoiceAgentTransport, error)
type VoiceAgentSessionOptions struct {
	TransportFactory VoiceTransportFactory
	Audio            VoiceAudioInterface
	FunctionHandlers []*VoiceFunctionHandler
}

type VoiceFunctionHandler struct {
	Name        string
	Description string
	Parameters  map[string]any
	Filler      *StaticToolFiller
	Timeout     time.Duration
	Function    func(context.Context, map[string]any) (any, error)
}

func (h *VoiceFunctionHandler) MarshalJSON() ([]byte, error) {
	if h == nil || h.Name == "" || h.Function == nil {
		return nil, fmt.Errorf("function handler requires name and callback")
	}
	function := map[string]any{"name": h.Name, "description": h.Description, "parameters": h.Parameters}
	timeout := h.Timeout
	if timeout == 0 {
		timeout = 15 * time.Second
	}
	if timeout.Milliseconds() < 1 {
		return nil, fmt.Errorf("function timeout must be positive")
	}
	value := map[string]any{"type": "function", "function": function, "timeout_ms": timeout.Milliseconds()}
	if h.Filler != nil {
		value["filler"] = h.Filler
	}
	return json.Marshal(value)
}
func (h *VoiceFunctionHandler) Invoke(ctx context.Context, arguments json.RawMessage) (any, error) {
	if h == nil || h.Function == nil {
		return nil, fmt.Errorf("function handler is unavailable")
	}
	var input map[string]any
	if len(arguments) > 0 && string(arguments) != "null" {
		if err := json.Unmarshal(arguments, &input); err != nil {
			return nil, fmt.Errorf("function arguments must be an object: %w", err)
		}
	}
	if input == nil {
		input = map[string]any{}
	}
	timeout := h.Timeout
	if timeout == 0 {
		timeout = 15 * time.Second
	}
	if timeout < 0 {
		return nil, fmt.Errorf("function timeout must be positive")
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	result, err := h.Function(ctx, input)
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return result, err
}

type VoiceAgentSession struct {
	client          *MistralClient
	response        *AudioSessionCreateResponse
	options         VoiceAgentSessionOptions
	mu              sync.Mutex
	startMu         sync.Mutex
	transport       VoiceAgentTransport
	handlers        map[string]map[uint64]func(VoiceEvent)
	nextHandler     uint64
	transcript      []VoiceAgentTranscriptEvent
	transcriptIndex map[string]int
	started         bool
	closed          bool
	closeOnce       sync.Once
	done            chan struct{}
	closeErr        error
	fatalErr        error
	cancel          context.CancelFunc
}

// InitVoiceAgentSession creates the remote session after checking local facilities.
// A native transport and optional audio device are supplied by the application.
func (c *MistralClient) InitVoiceAgentSession(ctx context.Context, request *AudioSessionCreateRequest, options VoiceAgentSessionOptions) (*VoiceAgentSession, error) {
	if request == nil || options.TransportFactory == nil {
		return nil, fmt.Errorf("request and voice transport factory are required")
	}
	if request.AgentVersion != nil && (request.AgentID == nil || *request.AgentID == "" || *request.AgentVersion < 1) {
		return nil, fmt.Errorf("positive agent_version requires agent_id")
	}
	if options.Audio != nil {
		if err := options.Audio.Preflight(ctx); err != nil {
			return nil, err
		}
	}
	copy := *request
	if request.Tools != nil {
		copy.Tools = append(make([]any, 0, len(request.Tools)), request.Tools...)
	}
	byName := map[string]*VoiceFunctionHandler{}
	for _, handler := range options.FunctionHandlers {
		if handler == nil || handler.Name == "" || handler.Function == nil {
			return nil, fmt.Errorf("invalid function handler")
		}
		if prior, ok := byName[handler.Name]; ok && prior != handler {
			return nil, fmt.Errorf("conflicting function handlers for %q", handler.Name)
		}
		byName[handler.Name] = handler
	}
	for _, tool := range copy.Tools {
		if handler, ok := tool.(*VoiceFunctionHandler); ok {
			if handler == nil || handler.Name == "" || handler.Function == nil {
				return nil, fmt.Errorf("invalid inline function handler")
			}
			if prior, exists := byName[handler.Name]; exists && prior != handler {
				return nil, fmt.Errorf("conflicting function handlers for %q", handler.Name)
			}
			byName[handler.Name] = handler
		}
	}
	options.FunctionHandlers = nil
	for _, handler := range byName {
		options.FunctionHandlers = append(options.FunctionHandlers, handler)
	}
	response, err := c.WithContext(ctx).CreateAudioSession(&copy)
	if err != nil {
		return nil, err
	}
	return &VoiceAgentSession{client: c, response: response, options: options, handlers: map[string]map[uint64]func(VoiceEvent){}, transcriptIndex: map[string]int{}, done: make(chan struct{})}, nil
}

func (s *VoiceAgentSession) ID() string                          { return s.response.ID }
func (s *VoiceAgentSession) Session() AudioSessionCreateResponse { return *s.response }
func (s *VoiceAgentSession) Transcript() []VoiceAgentTranscriptEvent {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]VoiceAgentTranscriptEvent(nil), s.transcript...)
}
func (s *VoiceAgentSession) On(name string, handler func(VoiceEvent)) func() {
	s.mu.Lock()
	s.nextHandler++
	id := s.nextHandler
	if s.handlers[name] == nil {
		s.handlers[name] = map[uint64]func(VoiceEvent){}
	}
	s.handlers[name][id] = handler
	s.mu.Unlock()
	return func() { s.mu.Lock(); delete(s.handlers[name], id); s.mu.Unlock() }
}
func (s *VoiceAgentSession) emit(event VoiceEvent) {
	s.mu.Lock()
	if event.Name == "error" {
		if failure, ok := event.Data.(VoiceAgentErrorEvent); ok && failure.Fatal && s.fatalErr == nil {
			s.fatalErr = errors.New(failure.Message)
		}
	}
	if event.Name == "transcript" {
		if transcript, ok := event.Data.(VoiceAgentTranscriptEvent); ok {
			if i, exists := s.transcriptIndex[transcript.ID]; exists {
				s.transcript[i] = transcript
			} else {
				s.transcriptIndex[transcript.ID] = len(s.transcript)
				s.transcript = append(s.transcript, transcript)
			}
		}
	}
	ids := make([]uint64, 0, len(s.handlers[event.Name]))
	for id := range s.handlers[event.Name] {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	handlers := make([]func(VoiceEvent), 0, len(ids))
	for _, id := range ids {
		handlers = append(handlers, s.handlers[event.Name][id])
	}
	s.mu.Unlock()
	for _, handler := range handlers {
		if handler != nil {
			func() { defer func() { _ = recover() }(); handler(event) }()
		}
	}
	if event.Name == "error" {
		if failure, ok := event.Data.(VoiceAgentErrorEvent); ok && failure.Fatal {
			go s.Close(context.Background(), "error")
		}
	}
	if event.Name == "audio" && s.options.Audio != nil {
		if audio, ok := event.Data.([]byte); ok {
			if err := s.options.Audio.OutputAudio(context.Background(), audio); err != nil {
				go s.Close(context.Background(), "audio_output_error")
			}
		}
	}
}
func (s *VoiceAgentSession) Connect(ctx context.Context) error {
	s.startMu.Lock()
	defer s.startMu.Unlock()
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return fmt.Errorf("cannot connect a closed voice session")
	}
	if s.started {
		s.mu.Unlock()
		return nil
	}
	s.mu.Unlock()
	transport, err := s.options.TransportFactory(s.response, s.emit, s.options.FunctionHandlers)
	if err == nil && transport == nil {
		err = fmt.Errorf("voice transport factory returned nil")
	}
	if err != nil {
		_ = s.Close(context.Background(), "start_failed")
		return err
	}
	s.mu.Lock()
	s.transport = transport
	s.mu.Unlock()
	if err = transport.Start(ctx); err != nil {
		_ = s.Close(context.Background(), "start_failed")
		return err
	}
	if s.options.Audio != nil {
		if err = s.options.Audio.Start(ctx); err != nil {
			_ = s.Close(context.Background(), "start_failed")
			return err
		}
	}
	pumpCtx, cancel := context.WithCancel(context.Background())
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		cancel()
		_ = transport.Close()
		return fmt.Errorf("voice session closed during connect")
	}
	s.cancel = cancel
	s.started = true
	s.mu.Unlock()
	if s.options.Audio != nil {
		go func() {
			for {
				audio, err := s.options.Audio.InputAudio(pumpCtx)
				if err != nil {
					if pumpCtx.Err() == nil {
						_ = s.Close(context.Background(), "audio_input_error")
					}
					return
				}
				if err := transport.SendAudio(pumpCtx, audio); err != nil {
					_ = s.Close(context.Background(), "audio_input_error")
					return
				}
			}
		}()
	}
	s.emit(VoiceEvent{Name: "connected", Data: map[string]any{"session_id": s.ID()}})
	return nil
}
func (s *VoiceAgentSession) requireTransport() (VoiceAgentTransport, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.closed || !s.started || s.transport == nil {
		return nil, fmt.Errorf("voice session is not connected")
	}
	return s.transport, nil
}
func (s *VoiceAgentSession) SendAudio(ctx context.Context, audio []byte) error {
	t, err := s.requireTransport()
	if err != nil {
		return err
	}
	return t.SendAudio(ctx, audio)
}
func (s *VoiceAgentSession) SetInputEnabled(ctx context.Context, enabled bool) error {
	t, err := s.requireTransport()
	if err != nil {
		return err
	}
	return t.SetInputEnabled(ctx, enabled)
}
func (s *VoiceAgentSession) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.done:
		s.mu.Lock()
		fatal := s.fatalErr
		s.mu.Unlock()
		if fatal != nil {
			return fatal
		}
		return s.closeErr
	}
}

type VoiceSessionUpdate struct {
	Instructions *string          `json:"instructions,omitempty"`
	Tools        []any            `json:"tools,omitempty"`
	Turn         *AudioTurnConfig `json:"turn,omitempty"`
}

func (s *VoiceAgentSession) Update(ctx context.Context, update VoiceSessionUpdate) (*VoiceAgentSessionUpdatedEvent, error) {
	t, err := s.requireTransport()
	if err != nil {
		return nil, err
	}
	payload := map[string]any{}
	if update.Instructions != nil {
		payload["instructions"] = *update.Instructions
	}
	if update.Tools != nil {
		payload["tools"] = update.Tools
	}
	if update.Turn != nil {
		payload["turn"] = update.Turn
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	response, err := t.UpdateSession(ctx, data)
	if err != nil {
		return nil, err
	}
	var event VoiceAgentSessionUpdatedEvent
	if err = json.Unmarshal(response, &event); err != nil {
		return nil, err
	}
	if event.Type != nil && *event.Type != "session.updated" {
		return nil, fmt.Errorf("invalid session update response type %q", *event.Type)
	}
	confirmed := map[string]bool{}
	for _, tool := range event.Session.Tools {
		data, err := json.Marshal(tool)
		if err != nil {
			continue
		}
		var wire struct {
			Type     string `json:"type"`
			Function struct {
				Name string `json:"name"`
			} `json:"function"`
		}
		if json.Unmarshal(data, &wire) == nil && wire.Type == "function" {
			confirmed[wire.Function.Name] = true
		}
	}
	var handlers []*VoiceFunctionHandler
	for _, tool := range update.Tools {
		if h, ok := tool.(*VoiceFunctionHandler); ok && h != nil && confirmed[h.Name] {
			handlers = append(handlers, h)
		}
	}
	t.AddFunctionHandlers(handlers)
	return &event, nil
}
func (s *VoiceAgentSession) Close(ctx context.Context, reason string) error {
	didClose := false
	s.closeOnce.Do(func() {
		didClose = true
		s.mu.Lock()
		s.closed = true
		transport := s.transport
		cancel := s.cancel
		s.mu.Unlock()
		if cancel != nil {
			cancel()
		}
		if transport != nil {
			s.closeErr = errors.Join(s.closeErr, transport.Close())
		}
		if s.options.Audio != nil {
			s.closeErr = errors.Join(s.closeErr, s.options.Audio.Stop())
		}
		s.closeErr = errors.Join(s.closeErr, s.client.WithContext(ctx).CloseAudioSession(s.ID()))
		close(s.done)
	})
	if didClose {
		s.emit(VoiceEvent{Name: "disconnected", Data: map[string]any{"session_id": s.ID(), "reason": reason}})
	}
	return s.closeErr
}

type VoiceWorkerToken struct {
	WSURL     string     `json:"ws_url"`
	AgentName string     `json:"agent_name"`
	Token     string     `json:"token"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

func requireSecureVoiceURL(raw, secure, insecure string) error {
	u, err := url.Parse(raw)
	if err != nil || u.Hostname() == "" {
		return fmt.Errorf("voice endpoint must be an absolute %s URL", secure)
	}
	if u.Scheme == secure {
		return nil
	}
	host := u.Hostname()
	ip := net.ParseIP(host)
	if u.Scheme == insecure && (strings.EqualFold(host, "localhost") || ip != nil && ip.IsLoopback()) {
		return nil
	}
	return fmt.Errorf("voice credentials require %s except on loopback", secure)
}
func (c *MistralClient) FetchVoiceWorkerToken(ctx context.Context) (*VoiceWorkerToken, error) {
	if err := requireSecureVoiceURL(c.endpoint, "https", "http"); err != nil {
		return nil, err
	}
	result, err := requestTyped[VoiceWorkerToken](c.WithContext(ctx), "POST", map[string]any{}, "v1/audio/workers/registration-token")
	if err != nil {
		return nil, err
	}
	if err := requireSecureVoiceURL(result.WSURL, "wss", "ws"); err != nil {
		return nil, err
	}
	return result, nil
}

// WaitForAgent waits until the native transport observes the remote agent.
func (s *VoiceAgentSession) WaitForAgent(ctx context.Context) error {
	transport, err := s.requireTransport()
	if err != nil {
		return err
	}
	return transport.WaitForAgent(ctx)
}
