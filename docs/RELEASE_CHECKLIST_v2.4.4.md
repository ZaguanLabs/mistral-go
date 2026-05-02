# v2.4.4 Release Checklist

## ✅ Completed

### Code & Implementation
- [x] Python SDK v2.4.4 parity updates implemented
- [x] 100% endpoint surface parity maintained for newly introduced generated resources
- [x] Zero external dependencies maintained
- [x] Go 1.20+ module compatibility maintained
- [x] Module path remains v2
- [x] Version updated to v2.4.4 in `sdk/version.go`

### New Features / Parity Additions
- [x] Audio Speech API: `Speech()` and `SpeechStream()`
- [x] Voices API: list/create/get/update/delete + sample audio download
- [x] Connectors API: connector CRUD, auth URL, auth methods, tools, tool calls, and credentials management
- [x] RAG API: ingestion pipeline configuration list/register/update-run-info
- [x] Workflows API: list, registrations, execute, update, archive, unarchive
- [x] Workflow helpers: `ExecuteWorkflowAndWait()` and `WaitForWorkflowCompletion()`
- [x] Workflow executions API: get/history, signal/query/update, reset, cancel/terminate, batch operations, trace, stream
- [x] Workflow runs, events, schedules, deployments, and metrics APIs
- [x] Observability API: campaigns, datasets, dataset records, imports/exports/tasks
- [x] Observability API: chat completion fields, chat completion events, judges, and live judging

### Core Reliability
- [x] Shared helper layer added for generated-style JSON endpoints
- [x] Generic SSE stream parser added for new streaming resources
- [x] Binary response helper added for voice sample audio
- [x] Flexible `APIResponse` support added for fast-moving generated response schemas
- [x] User-Agent updated to v2.4.4 in tests

### Testing
- [x] New mock-server parity tests added for the new resource groups
- [x] New binary and stream endpoint tests added
- [x] Compile-only SDK validation passing
- [x] Focused new parity tests passing
- [x] `go build ./sdk/...` successful
- [x] Full live/API-dependent suite reviewed; requires credentials and deterministic model outputs

### Documentation
- [x] README version banner updated to v2.4.4 / Python SDK v2.4.4
- [x] README feature list updated for new APIs
- [x] CHANGELOG includes v2.4.4 release entry
- [x] Release notes prepared for v2.4.4
- [x] Release checklist prepared for v2.4.4

### Repository Hygiene
- [x] Bundled Python SDK source ignored via `docs/client-python-2.*`
- [x] Python SDK tarballs ignored via `*.tar.gz`
- [x] No external dependencies added to `go.mod`
- [x] No credentials or secrets added

## 📋 Release Summary

### Version
- **Previous**: v2.2.0
- **Current**: v2.4.4
- **Python SDK Compatibility Target**: v2.4.4

### Final Verification
```bash
# Build check
go build ./sdk/...

# Compile-only package validation
go test ./sdk/... -run '^$'

# Focused parity regression checks
go test ./sdk/... -run 'TestNewParity'

# Version check
grep 'Version = ' sdk/version.go
```

### Known Validation Notes
- The full test suite contains live/API-dependent tests that require valid API credentials and stable model output.
- Release validation should rely on mock-server tests and compile checks unless running in a fully configured live-test environment.

## 📋 Release Steps

1. **Commit**
   ```bash
   git add -A
   git commit -m "Release v2.4.4 - Python SDK v2.4.4 parity updates"
   ```

2. **Tag**
   ```bash
   git tag -a v2.4.4 -m "Release v2.4.4 - Python SDK v2.4.4 parity updates"
   ```

3. **Push**
   ```bash
   git push origin main
   git push origin v2.4.4
   ```

4. **GitHub Release**
   - Create release from tag `v2.4.4`
   - Use `docs/RELEASE_NOTES_v2.4.4.md` as release notes

---

## ✅ Checklist Complete - Ready to Release v2.4.4
