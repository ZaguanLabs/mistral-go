# v2.4.12 Release Checklist

## Completed

### Code & Implementation

- [x] Python SDK v2.4.12 parity updates implemented
- [x] REST endpoint surface updated for new generated resources
- [x] Go 1.20+ module compatibility maintained
- [x] Module path remains v2
- [x] Version updated to v2.4.12 in `sdk/version.go`

### New Features / Parity Additions

- [x] Observability log search, field list, and field option APIs
- [x] Observability span search, span evaluation search, latest span evaluation search, fields, and field options
- [x] Observability trace search, trace fetch, trace spans, fields, field options, and span-by-ID APIs
- [x] RAG search index APIs migrated to `v1/rag/indexes`
- [x] RAG search index detail, metrics update, summary-field, schema detail, schema summary, and schema file operations
- [x] Workflow execution log search and SSE stream operations
- [x] Workflow schedule trigger operation
- [x] Workflow list filters for status, deployment, tags, sort field, and order
- [x] Workflow run filters for deployment, sort field, order, start time, and end time
- [x] Connector auth URL `github_installation_link` option

### Core Reliability

- [x] RFC3339Nano timestamp query serialization added for new datetime filters
- [x] Typed order/status/sort enum query serialization added
- [x] Workflow status list query serialization added
- [x] Library delete and access delete helpers tolerate empty `204 No Content` responses
- [x] Existing connector auth URL call sites remain source-compatible via optional variadic parameter

### Testing

- [x] Mock-server parity tests added for new v2.4.12 endpoints
- [x] Stream test coverage added for workflow execution log SSE endpoint
- [x] Binary response test coverage added for search index schema files
- [x] Compile-only package validation passing
- [x] Focused parity tests passing
- [x] Full live/API-dependent suite reviewed; requires credentials and deterministic model outputs

### Documentation

- [x] Release notes prepared for v2.4.12
- [x] Release checklist prepared for v2.4.12
- [x] Current release checklist pointer updated to v2.4.12

### Repository Hygiene

- [x] Bundled Python SDK source remains under `docs/client-python/`
- [x] No credentials or secrets added
- [x] No new third-party Go dependencies added

## Release Summary

### Version

- Previous: v2.4.9
- Current: v2.4.12
- Python SDK Compatibility Target: v2.4.12

### Final Verification

```bash
go test ./sdk -run 'TestNewParity|TestListSearchIndexesEndpoint|TestVersion|Test.*Parity' -count=1
go test ./sdk -run '^$'
grep 'Version = ' sdk/version.go
```

### Known Validation Notes

- The full test suite contains live/API-dependent tests that require valid API credentials and stable model output.
- In this workspace, `go test ./...` still fails in live chat, embeddings, and FIM tests for those reasons.
- Release validation should rely on mock-server tests and compile checks unless running in a fully configured live-test environment.

## Release Steps

1. Commit:
   ```bash
   git add -A
   git commit -m "Release v2.4.12 - Python SDK v2.4.12 parity updates"
   ```

2. Tag:
   ```bash
   git tag -a v2.4.12 -m "Release v2.4.12 - Python SDK v2.4.12 parity updates"
   ```

3. Push:
   ```bash
   git push origin main
   git push origin v2.4.12
   ```

4. GitHub Release:
   - Create release from tag `v2.4.12`
   - Use `docs/RELEASE_NOTES_v2.4.12.md` as release notes

---

## Checklist Complete - Ready to Release v2.4.12
