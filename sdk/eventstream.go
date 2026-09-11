package sdk

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
)

// ServerEvent retains the SSE envelope, including an explicitly empty data field.
type ServerEvent struct {
	Event string
	Data  string
	ID    string
	Retry *int
}

// EventStream is a pull-based stream. Close it when stopping before EOF.
type EventStream struct {
	finished   bool
	closeOnce  sync.Once
	closeError error
	body       io.ReadCloser
	reader     *bufio.Reader
	first      bool
	id         string
}

func NewEventStream(body io.ReadCloser) *EventStream {
	return &EventStream{body: body, reader: bufio.NewReader(body), first: true}
}
func (s *EventStream) Close() error {
	s.closeOnce.Do(func() { s.closeError = s.body.Close() })
	return s.closeError
}
func (s *EventStream) line() (string, error) {
	var b strings.Builder
	for {
		ch, err := s.reader.ReadByte()
		if err != nil {
			return b.String(), err
		}
		if ch == '\n' {
			return b.String(), nil
		}
		if ch == '\r' {
			if next, _ := s.reader.Peek(1); len(next) > 0 && next[0] == '\n' {
				_, _ = s.reader.ReadByte()
			}
			return b.String(), nil
		}
		b.WriteByte(ch)
	}
}
func (s *EventStream) Next() (ServerEvent, error) {
	if s.finished {
		return ServerEvent{}, io.EOF
	}
	event := ServerEvent{}
	parts := []string{}
	for {
		line, err := s.line()
		if err != nil {
			s.finished = true
		}
		if s.first {
			line = strings.TrimPrefix(line, "\ufeff")
			s.first = false
		}
		if line != "" && !strings.HasPrefix(line, ":") {
			field, value, found := strings.Cut(line, ":")
			if !found {
				value = ""
			}
			value = strings.TrimPrefix(value, " ")
			switch field {
			case "event":
				event.Event = value
			case "data":
				parts = append(parts, value)
			case "id":
				if !strings.ContainsRune(value, 0) {
					s.id = value
				}
			case "retry":
				if n, e := strconv.Atoi(value); e == nil && n >= 0 {
					event.Retry = &n
				}
			}
		}
		if err != nil && err != io.EOF {
			s.Close()
			return ServerEvent{}, err
		}
		if line == "" || err == io.EOF {
			if len(parts) > 0 {
				event.Data = strings.Join(parts, "\n")
				event.ID = s.id
				if event.Data == "[DONE]" {
					s.finished = true
					s.Close()
					return ServerEvent{}, io.EOF
				}
				if err == io.EOF {
					s.Close()
				}
				return event, nil
			}
			if err == io.EOF {
				s.Close()
				return ServerEvent{}, io.EOF
			}
			event = ServerEvent{}
		}
	}
}
func streamValues[T any](ctx context.Context, body io.ReadCloser, decode func(ServerEvent) (T, error), failure func(error) T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		stream := NewEventStream(body)
		defer stream.Close()
		done := make(chan struct{})
		defer close(done)
		go func() {
			select {
			case <-ctx.Done():
				stream.Close()
			case <-done:
			}
		}()
		send := func(value T) bool {
			select {
			case out <- value:
				return true
			case <-ctx.Done():
				return false
			}
		}
		for {
			event, err := stream.Next()
			if err == io.EOF {
				return
			}
			if err != nil {
				if ctx.Err() == nil {
					send(failure(err))
				}
				return
			}
			if event.Event == "error" {
				payload := map[string]any{}
				_ = json.Unmarshal([]byte(event.Data), &payload)
				reason, _ := payload["reason"].(string)
				if reason == "" {
					reason = "stream_error"
				}
				message, _ := payload["error"].(string)
				if message == "" {
					message = event.Data
				}
				send(failure(&StreamDisconnectedError{Reason: reason, ErrorMessage: message}))
				return
			}
			value, err := decode(event)
			if err != nil {
				if !send(failure(fmt.Errorf("error decoding stream event: %w", err))) {
					return
				}
				continue
			}
			if !send(value) {
				return
			}
		}
	}()
	return out
}
func streamJSON[T any](ctx context.Context, body io.ReadCloser, failure func(error) T) <-chan T {
	return streamValues(ctx, body, func(event ServerEvent) (T, error) {
		var value T
		err := json.Unmarshal([]byte(event.Data), &value)
		return value, err
	}, failure)
}
func parseGenericStreamContext(ctx context.Context, body io.ReadCloser) <-chan StreamEvent {
	return streamValues(ctx, body, func(event ServerEvent) (StreamEvent, error) {
		var payload map[string]any
		err := json.Unmarshal([]byte(event.Data), &payload)
		name, _ := payload["type"].(string)
		if name == "" {
			name = event.Event
		}
		return StreamEvent{Type: name, Data: payload}, err
	}, func(err error) StreamEvent { return StreamEvent{Error: err} })
}
