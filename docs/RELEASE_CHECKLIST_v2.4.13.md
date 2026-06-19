# v2.4.13 Release Checklist

## Completed

### Code & Implementation

- [x] Python SDK v2.4.13 parity updates implemented
- [x] REST endpoint surface updated for new generated deployment log resources
- [x] Go 1.20+ module compatibility maintained
- [x] Module path remains v2
- [x] Version updated to v2.4.13 in `sdk/version.go`

### New Features / Parity Additions

- [x] Deployment log search API
- [x] Deployment log SSE stream API
- [x] Workflow run `root_execution_id` filter
- [x] Workflow run `include_internal` filter
- [x] Search index schema file JSON response helper

### Core Reliability

- [x] Search index schema file helper updated for v2.4.13 JSON `content` responses
- [x] Existing byte-returning schema file helper preserved for source compatibility
- [x] Deployment log stream uses the existing generic SSE parser

### Testing

- [x] Mock-server parity tests added for new v2.4.13 endpoints
- [x] Stream test coverage added for deployment log SSE endpoint
- [x] Search index schema file test updated for JSON response shape
- [x] Compile-only package validation passing
- [x] Focused parity tests passing

### Documentation

- [x] README version banner updated to v2.4.13 / Python SDK v2.4.13
- [x] README compatibility section updated for deployment logs, workflow run filters, and schema file response handling
- [x] CHANGELOG includes v2.4.13 release entry
- [x] Release notes prepared for v2.4.13
- [x] Release checklist prepared for v2.4.13
- [x] Current release checklist pointer updated to v2.4.13

### Repository Hygiene

- [x] Bundled Python SDK source remains under `docs/client-python/`
- [x] No credentials or secrets added
- [x] No new third-party Go dependencies added

## Release Summary

### Version

- Previous: v2.4.12
- Current: v2.4.13
- Python SDK Compatibility Target: v2.4.13

### Final Verification

```bash
go test ./sdk -run 'TestNewParity|TestListSearchIndexesEndpoint|TestVersion|Test.*Parity' -count=1
go test ./sdk -run '^$'
grep 'Version = ' sdk/version.go
```

### Known Validation Notes

- The full test suite contains live/API-dependent tests that require valid API credentials and stable model output.
- Release validation should rely on mock-server tests and compile checks unless running in a fully configured live-test environment.

## Release Steps

1. Commit:
   ```bash
   git add -A
   git commit -m "Release v2.4.13 - Python SDK v2.4.13 parity updates"
   ```

2. Tag:
   ```bash
   git tag -a v2.4.13 -m "Release v2.4.13 - Python SDK v2.4.13 parity updates"
   ```

3. Push:
   ```bash
   git push origin main
   git push origin v2.4.13
   ```

4. GitHub Release:
   - Create release from tag `v2.4.13`
   - Use `docs/RELEASE_NOTES_v2.4.13.md` as release notes

---

## Checklist Complete - Ready to Release v2.4.13
