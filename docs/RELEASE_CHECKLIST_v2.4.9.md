# v2.4.9 Release Checklist

## Completed

### Code & Implementation

- [x] Python SDK v2.4.9 parity updates implemented
- [x] REST endpoint surface updated for new generated resources
- [x] Realtime audio WebSocket API added for official SDK parity
- [x] Go 1.20+ module compatibility maintained
- [x] Module path remains v2
- [x] Version updated to v2.4.9 in `sdk/version.go`

### New Features / Parity Additions

- [x] Audio realtime transcription WebSocket connect/session/audio/event helpers
- [x] Batch job deletion
- [x] Chat request additions: reasoning effort, guardrails, prompt cache key, built-in tools, flexible tool choice, JSON schema response formats, and string stop values
- [x] Connector organization/workspace/user activate and deactivate operations
- [x] Connector OAuth2 metadata and protocol request fields
- [x] RAG search index list/register operations
- [x] Workflow bulk archive/unarchive operations
- [x] Workflow schedule filters, get schedule, and update schedule operations
- [x] Model card and capability fields updated

### Core Reliability

- [x] Empty `204 No Content` JSON responses handled without decode errors
- [x] Realtime events parsed tolerantly with raw payload preservation
- [x] Realtime audio messages use official base64 JSON event shapes
- [x] User-Agent updated to v2.4.9 in tests

### Testing

- [x] WebSocket mock tests added for realtime handshake, session update, audio append/flush/end, and stream helper
- [x] Mock-server parity tests added for new v2.4.9 endpoints
- [x] Chat request mock coverage added for new request fields
- [x] Compile-only package validation passing
- [x] Focused parity and realtime tests passing
- [x] `go build ./sdk/...` successful
- [x] Full live/API-dependent suite reviewed; requires credentials and deterministic model outputs

### Documentation

- [x] README version banner updated to v2.4.9 / Python SDK v2.4.9
- [x] README dependency wording updated for realtime WebSocket parity
- [x] README feature list updated for realtime audio and RAG search indexes
- [x] CHANGELOG includes v2.4.9 release entry
- [x] Release notes prepared for v2.4.9
- [x] Release checklist prepared for v2.4.9

### Repository Hygiene

- [x] Bundled Python SDK source ignored via `docs/client-python/`
- [x] Python SDK tarballs ignored via `*.tar.gz`
- [x] Only one pure-Go dependency added: `nhooyr.io/websocket`
- [x] No credentials or secrets added

## Release Summary

### Version

- Previous: v2.4.4
- Current: v2.4.9
- Python SDK Compatibility Target: v2.4.9

### Final Verification

```bash
go build ./sdk/...
go test ./... -run '^$'
go test ./sdk -run 'TestRealtime|TestChatWithMock|TestChatNewRequestFieldsWithMock|TestNoContentResponseWithMock|TestNewParity|TestListSearchIndexesEndpoint|TestUserAgentHeaderWithMock'
grep 'Version = ' sdk/version.go
```

### Known Validation Notes

- The full test suite contains live/API-dependent tests that require valid API credentials and stable model output.
- Release validation should rely on mock-server tests and compile checks unless running in a fully configured live-test environment.

## Release Steps

1. Commit:
   ```bash
   git add -A
   git commit -m "Release v2.4.9 - Python SDK v2.4.9 parity updates"
   ```

2. Tag:
   ```bash
   git tag -a v2.4.9 -m "Release v2.4.9 - Python SDK v2.4.9 parity updates"
   ```

3. Push:
   ```bash
   git push origin main
   git push origin v2.4.9
   ```

4. GitHub Release:
   - Create release from tag `v2.4.9`
   - Use `docs/RELEASE_NOTES_v2.4.9.md` as release notes

---

## Checklist Complete - Ready to Release v2.4.9
