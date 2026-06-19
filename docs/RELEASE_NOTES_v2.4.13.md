# Release Notes - v2.4.13

Mistral Go SDK v2.4.13 updates the SDK for compatibility with the official Mistral Python SDK v2.4.13.

## Highlights

- Python SDK v2.4.13 parity updates
- Deployment log search and SSE streaming
- Workflow run filters for execution trees and internal workflow runs
- Updated RAG search index schema-file response handling
- README, CHANGELOG, release notes, and release checklist updated for v2.4.13

## Added

### Deployments

- `GetDeploymentLogs()`
- `GetWorkflowDeploymentLogs()`
- `StreamDeploymentLogs()`
- `StreamWorkflowDeploymentLogs()`

### Workflows

- `ListWorkflowRunsParams.RootExecutionID`
- `ListWorkflowRunsParams.IncludeInternal`

### RAG

- `GetSearchIndexSchemaFileResponse()`

## Changed

- SDK version updated to `2.4.13`
- User-Agent updated to `mistral-go/2.4.13`
- `GetSearchIndexSchemaFile()` now reads the Python SDK v2.4.13 JSON `content` response while preserving the existing `[]byte` return shape
- README compatibility notes updated to Python SDK v2.4.13
- CHANGELOG includes a v2.4.13 release entry

## Validation

Validated with:

```bash
gofmt -w sdk/workflows.go sdk/rag.go sdk/version.go sdk/parity_methods_test.go
go test ./sdk -run 'TestNewParity|TestListSearchIndexesEndpoint|TestVersion|Test.*Parity' -count=1
go test ./sdk -run '^$'
```

The full test suite includes live/API-dependent tests that require valid credentials and deterministic model outputs.
