# Release Notes - v2.4.4

Mistral Go SDK v2.4.4 updates the SDK for compatibility with the official Mistral Python SDK v2.4.4.

## Highlights

- **Python SDK v2.4.4 parity** for newly introduced generated resource groups
- **Zero external dependencies maintained**
- **New workflow orchestration APIs**
- **New observability APIs**
- **New connectors APIs**
- **New RAG ingestion pipeline configuration APIs**
- **New audio speech and voices APIs**

## Added

### Audio Speech and Voices

- `Speech()`
- `SpeechStream()`
- `ListVoices()`
- `CreateVoice()`
- `GetVoice()`
- `UpdateVoice()`
- `DeleteVoice()`
- `GetVoiceSampleAudio()`

### Connectors

- Connector create/list/get/update/delete
- Connector authentication URL helpers
- Connector authentication methods
- Connector tools and tool calls
- Organization, workspace, and user connector credentials management

### RAG

- `ListIngestionPipelineConfigurations()`
- `RegisterIngestionPipelineConfiguration()`
- `UpdateIngestionPipelineRunInfo()`

### Workflows

- Workflow listing and registration APIs
- Workflow execute/update/archive/unarchive APIs
- Workflow deployment, metrics, runs, events, and schedules APIs
- `ExecuteWorkflowAndWait()`
- `WaitForWorkflowCompletion()`

### Workflow Executions

- Execution get/history
- Signal/query/update/reset
- Cancel/terminate and batch cancel/terminate
- Trace OTEL, trace summary, trace events
- Execution streaming

### Observability

- Campaigns
- Datasets and dataset records
- Dataset imports, exports, and tasks
- Chat completion fields
- Chat completion events
- Judges and live judging

## Changed

- SDK version updated to `2.4.4`
- User-Agent updated to `mistral-go/2.4.4`
- README compatibility notes updated to Python SDK v2.4.4
- Added flexible `APIResponse` for large generated response schemas

## Validation

Validated with:

```bash
go build ./sdk/...
go test ./sdk/... -run '^$'
go test ./sdk/... -run 'TestNewParity'
```

The full test suite includes live/API-dependent tests that require valid credentials and deterministic model outputs.
