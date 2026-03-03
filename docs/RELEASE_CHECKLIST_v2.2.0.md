# v2.2.0 Release Checklist

## ✅ Completed

### Code & Implementation
- [x] Python SDK v1.12.4 parity updates implemented
- [x] Zero external dependencies maintained
- [x] Go 1.21+ compatibility maintained
- [x] Module path remains v2
- [x] Version updated to v2.2.0 in `sdk/version.go`

### New Features / Parity Additions
- [x] Batch API: inline `requests` support validation (strict mode selection)
- [x] Batch API: `agent_id`, `order_by`, `inline` retrieval support
- [x] Classifiers API: `moderate_chat`, `classify`, `classify_chat`
- [x] Audio API: streaming transcription (`TranscribeStream`)
- [x] Audio API: `diarize`, `context_bias` parameters
- [x] Conversations API: stream endpoints (`start_stream`, `append_stream`, `restart_stream`)
- [x] Conversations API: `get_messages` + extended request params
- [x] Agents API: versions and aliases management support
- [x] Documents API: advanced listing + signed/extracted text URLs + reprocess
- [x] Libraries API: `chunk_size` create param
- [x] Files API: `include_total`, `mimetypes` filters and query handling fixes
- [x] Accesses API: beta share model support with backward compatibility wrappers

### Core Reliability
- [x] `client.request` query parsing fixed (preserve query params)
- [x] `client.request` JSON decoding supports non-object payloads
- [x] User-Agent updated to v2.2.0 in tests

### Testing
- [x] New/updated tests for batch validation paths
- [x] New/updated tests for classifiers chat/text classification endpoints
- [x] New/updated tests for conversation params + messages endpoint
- [x] New integration tests for query-param behavior (`inline`, `include_total`, `mimetypes`)
- [x] Targeted SDK test suites passing
- [x] `go build ./sdk/...` successful

### Documentation
- [x] README version banner updated to v2.2.0 / Python SDK v1.12.4
- [x] CHANGELOG includes v2.2.0 release entry
- [x] Release checklist prepared for v2.2.0

## 📋 Release Summary

### Version
- **Previous**: v2.1.0
- **Current**: v2.2.0
- **Python SDK Compatibility Target**: v1.12.4

### Final Verification
```bash
# Build check
go build ./sdk/...

# Focused regression checks
go test ./sdk -run 'Test(GetBatchJobWithInlineQueryWithMock|ListFilesWithQueryParamsWithMock|ListConversationsWithParams|GetConversationMessages|ModerateChat|Classify|ClassifyChat|CreateBatchJobRequestValidationStrict|CreateBatchJobRequestNoFieldsError)'

# Version check
grep 'Version = ' sdk/version.go
```

## 📋 Release Steps

1. **Commit**
   ```bash
   git add -A
   git commit -m "Release v2.2.0 - Python SDK v1.12.4 parity updates"
   ```

2. **Tag**
   ```bash
   git tag -a v2.2.0 -m "Release v2.2.0 - Python SDK v1.12.4 parity updates"
   ```

3. **Push**
   ```bash
   git push origin main
   git push origin v2.2.0
   ```

4. **GitHub Release**
   - Create release from tag `v2.2.0`
   - Use CHANGELOG v2.2.0 section as release notes

---

## ✅ Checklist Complete - Ready to Release v2.2.0
