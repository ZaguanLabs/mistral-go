package sdk

import (
	"context"
	"regexp"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// redactingTraceExporter copies data at export so application processors retain their original spans.
type redactingTraceExporter struct {
	sdktrace.SpanExporter
	mask func(attribute.KeyValue) attribute.KeyValue
}
type redactedTraceSpan struct {
	sdktrace.ReadOnlySpan
	mask func(attribute.KeyValue) attribute.KeyValue
}

func (e *redactingTraceExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	result := make([]sdktrace.ReadOnlySpan, len(spans))
	for i, span := range spans {
		result[i] = &redactedTraceSpan{span, e.mask}
	}
	return e.SpanExporter.ExportSpans(ctx, result)
}
func (s *redactedTraceSpan) Name() string {
	return s.mask(attribute.String("span.name", s.ReadOnlySpan.Name())).Value.AsString()
}
func (s *redactedTraceSpan) attrs(input []attribute.KeyValue) []attribute.KeyValue {
	result := make([]attribute.KeyValue, 0, len(input))
	for _, a := range input {
		masked := s.mask(a)
		if masked.Valid() {
			result = append(result, masked)
		}
	}
	return result
}
func (s *redactedTraceSpan) Attributes() []attribute.KeyValue {
	return s.attrs(s.ReadOnlySpan.Attributes())
}
func (s *redactedTraceSpan) Events() []sdktrace.Event {
	events := append([]sdktrace.Event(nil), s.ReadOnlySpan.Events()...)
	for i := range events {
		events[i].Name = s.mask(attribute.String("event.name", events[i].Name)).Value.AsString()
		events[i].Attributes = s.attrs(events[i].Attributes)
	}
	return events
}
func (s *redactedTraceSpan) Links() []sdktrace.Link {
	links := append([]sdktrace.Link(nil), s.ReadOnlySpan.Links()...)
	for i := range links {
		links[i].Attributes = s.attrs(links[i].Attributes)
	}
	return links
}
func (s *redactedTraceSpan) Status() sdktrace.Status {
	status := s.ReadOnlySpan.Status()
	status.Description = s.mask(attribute.String("status.description", status.Description)).Value.AsString()
	return status
}
func (s *redactedTraceSpan) Resource() *resource.Resource {
	r := s.ReadOnlySpan.Resource()
	if r == nil {
		return nil
	}
	return resource.NewWithAttributes(r.SchemaURL(), s.attrs(r.Attributes())...)
}
func redactMistralAttribute(kv attribute.KeyValue) attribute.KeyValue {
	scrub := func(s string) string {
		for _, p := range mistralRedactionPatterns {
			s = p.ReplaceAllString(s, "[REDACTED]")
		}
		return s
	}
	switch kv.Value.Type() {
	case attribute.STRING:
		kv.Value = attribute.StringValue(scrub(kv.Value.AsString()))
	case attribute.STRINGSLICE:
		values := append([]string(nil), kv.Value.AsStringSlice()...)
		for i := range values {
			values[i] = scrub(values[i])
		}
		kv.Value = attribute.StringSliceValue(values)
	}
	return kv
}

// These are the default content redaction patterns shipped in Python SDK 3.9.0.
var mistralRedactionPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)bearer\s+[a-z0-9._\-]+`),
	regexp.MustCompile(`\bgh[pousr]_[A-Za-z0-9_]{20,}\b`),
	regexp.MustCompile(`\bxox[baprs]-[A-Za-z0-9-]{10,}\b`),
	regexp.MustCompile(`\bsk-[A-Za-z0-9]{20,}\b`),
	regexp.MustCompile(`\bAKIA[0-9A-Z]{16}\b`),
	regexp.MustCompile(`\bAIza[0-9A-Za-z_\-]{35}\b`),
	regexp.MustCompile(`\beyJ[A-Za-z0-9_\-]+\.[A-Za-z0-9_\-]+\.[A-Za-z0-9_\-]+\b`),
	regexp.MustCompile(`-----BEGIN [A-Z ]*PRIVATE KEY-----`),
	regexp.MustCompile(`\b[sr]k_(?:live|test)_[0-9A-Za-z]{10,}\b`),
	regexp.MustCompile(`\bsk-ant-[A-Za-z0-9\-_]{20,}\b`),
	regexp.MustCompile(`\bsk-proj-[A-Za-z0-9\-_]{20,}\b`),
	regexp.MustCompile(`\bhf_[A-Za-z0-9]{30,}\b`),
	regexp.MustCompile(`\bmstrl_[0-9A-Za-z]{32}_[0-9A-Za-z]{6}\b`),
	regexp.MustCompile(`\bgithub_pat_[A-Za-z0-9_]{22,}\b`),
	regexp.MustCompile(`\bglpat-[A-Za-z0-9\-=_]{20,22}\b`),
	regexp.MustCompile(`\bshp(?:at|ca|pa|ss)_[a-fA-F0-9]{32}\b`),
	regexp.MustCompile(`\bsq0(?:atp|csp|idp)-[0-9A-Za-z\-_]{22,43}\b`),
	regexp.MustCompile(`\bPMAK-[a-zA-Z0-9]{24,59}\b`),
	regexp.MustCompile(`\bphc_[a-zA-Z0-9_]{43}\b`),
	regexp.MustCompile(`\brubygems_[a-f0-9]{48}\b`),
	regexp.MustCompile(`\blin_api_[0-9A-Za-z]{40}\b`),
	regexp.MustCompile(`pypi-AgEIcHlwaS5vcmc[A-Za-z0-9\-_]{50,}`),
	regexp.MustCompile(`\bsecret_[A-Za-z0-9]{43}\b`),
	regexp.MustCompile(`[A-Za-z0-9]{14}\.atlasv1\.[A-Za-z0-9]{60,}`),
	regexp.MustCompile(`\bSG\.[A-Za-z0-9_\-]{22}\.[A-Za-z0-9_\-]{43}\b`),
	regexp.MustCompile(`\bpk_(?:live|test)_[0-9a-zA-Z]{24}\b`),
	regexp.MustCompile(`https://hooks\.slack\.com/services/[A-Za-z0-9/+]{40,}`),
	regexp.MustCompile(`https://discord(?:app)?\.com/api/webhooks/[0-9]{17,}/[A-Za-z0-9\-_]{60,}`),
	regexp.MustCompile(`https://hooks\.zapier\.com/hooks/catch/[A-Za-z0-9/]{16,}`),
	regexp.MustCompile(`\b[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}\b`),
	regexp.MustCompile(`\b(?:\d[ -]?){13,16}\b`),
	regexp.MustCompile(`\b(?:(?:25[0-5]|2[0-4]\d|1?\d?\d)\.){3}(?:25[0-5]|2[0-4]\d|1?\d?\d)\b`),
}
