# Release Notes - v2.4.12

Mistral Go SDK v2.4.12 updates the SDK for compatibility with the official Mistral Python SDK v2.4.12.

## Highlights

- Python SDK v2.4.12 parity updates
- New observability logs, spans, span evaluations, and traces APIs
- Updated RAG search index APIs for the new `v1/rag/indexes` surface
- Workflow execution log search and SSE streaming
- Expanded workflow list/run filters and schedule trigger support
- Connector auth URL support for GitHub installation links

## Added

### Observability

- `SearchLogs()`
- `ListLogFields()`
- `FetchLogFieldOptions()`
- `SearchSpans()`
- `SearchSpanEvaluations()`
- `SearchLatestSpanEvaluations()`
- `ListSpanFields()`
- `ListSpanEvaluationFields()`
- `FetchSpanFieldOptions()`
- `FetchSpanEvaluationFieldOptions()`
- `SearchTraces()`
- `ListTraceFields()`
- `GetTraceByID()`
- `GetTraceSpans()`
- `FetchTraceFieldOptions()`
- `GetSpanByID()`

### RAG

- `GetSearchIndexSummaries()`
- `UnregisterSearchIndex()`
- `UpdateSearchIndexMetrics()`
- `GetSearchIndexDetail()`
- `SetSearchIndexSummary()`
- `GetSearchIndexSchemaDetail()`
- `SetSearchIndexSchemaSummary()`
- `GetSearchIndexSchemaFile()`

### Workflows

- `GetWorkflowExecutionLogs()`
- `StreamWorkflowExecutionLogs()`
- `TriggerSchedule()`
- Workflow list filters for status, deployment name, deployment status, tags, sort field, and sort order
- Workflow run filters for deployment name, sort field, sort order, start time, and end time

## Changed

- SDK version updated to `2.4.12`
- User-Agent updated to `mistral-go/2.4.12`
- `ListSearchIndexes()` now uses the Python SDK v2.4.12 summary route, `v1/rag/indexes/summary`
- `RegisterSearchIndex()` now uses `v1/rag/indexes`
- Query serialization now supports typed order/status/sort enums and RFC3339Nano timestamps
- `GetConnectorAuthURL()` accepts optional `github_installation_link`
- Library and access deletion helpers now tolerate empty `204 No Content` responses

## Validation

Validated with:

```bash
gofmt -w sdk/connectors.go sdk/accesses.go sdk/libraries.go sdk/parity_helpers.go sdk/types.go sdk/workflows.go sdk/workflow_executions.go sdk/observability_signals.go sdk/rag.go sdk/version.go sdk/parity_methods_test.go
go test ./sdk -run 'TestNewParity|TestListSearchIndexesEndpoint|TestVersion|Test.*Parity' -count=1
go test ./sdk -run '^$'
```

The full test suite includes live/API-dependent tests. In this workspace, `go test ./...` still reaches live chat, embeddings, and FIM tests; those require valid credentials and deterministic model outputs.
