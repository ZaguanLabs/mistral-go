# Release Notes - v2.4.9

Mistral Go SDK v2.4.9 updates the SDK for compatibility with the official Mistral Python SDK v2.4.9.

## Highlights

- Python SDK v2.4.9 parity updates
- Realtime audio transcription over WebSocket
- New RAG search index APIs
- New workflow schedule and bulk archive helpers
- Expanded connector activation APIs
- Updated chat request fields for reasoning, guardrails, prompt caching, built-in tools, and JSON schema response formats

## Added

### Audio Realtime

- `RealtimeTranscriptionConnect()`
- `RealtimeTranscribeStream()`
- `RealtimeConnection.SendAudio()`
- `RealtimeConnection.FlushAudio()`
- `RealtimeConnection.EndAudio()`
- `RealtimeConnection.UpdateSession()`
- `RealtimeConnection.Events()`

Realtime audio uses `nhooyr.io/websocket`, a pure-Go WebSocket dependency.

### Batch

- `DeleteBatchJob()`

### Chat

- String or list `stop` values
- Full `response_format` objects for JSON schema mode
- Built-in tool payloads
- Specific tool choice payloads
- `reasoning_effort`
- `guardrails`
- `prompt_cache_key`

### Connectors

- Organization connector activate/deactivate
- Workspace connector activate/deactivate
- User connector activate/deactivate
- OAuth2 server metadata fields
- Connector protocol request fields

### RAG

- `ListSearchIndexes()`
- `RegisterSearchIndex()`

### Workflows

- `BulkArchiveWorkflows()`
- `BulkUnarchiveWorkflows()`
- Schedule list filters
- `GetWorkflowSchedule()`
- `UpdateWorkflowSchedule()`

## Changed

- SDK version updated to `2.4.9`
- User-Agent updated to `mistral-go/2.4.9`
- Model card and capability fields updated for Python SDK v2.4.9
- Shared JSON request helper now handles empty `204 No Content` responses
- README compatibility notes updated to Python SDK v2.4.9

## Validation

Validated with:

```bash
go build ./sdk/...
go test ./... -run '^$'
go test ./sdk -run 'TestRealtime|TestChatWithMock|TestChatNewRequestFieldsWithMock|TestNoContentResponseWithMock|TestNewParity|TestListSearchIndexesEndpoint|TestUserAgentHeaderWithMock'
```

The full test suite includes live/API-dependent tests that require valid credentials and deterministic model outputs.
