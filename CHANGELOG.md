# Changelog

All notable changes to the Mistral Go SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.4.12] - 2026-06-18

### Added - Python SDK v2.4.12 Parity Updates

- Observability API:
  - `SearchLogs()`, `ListLogFields()`, and `FetchLogFieldOptions()`.
  - `SearchSpans()`, `SearchSpanEvaluations()`, and `SearchLatestSpanEvaluations()`.
  - Span field and span evaluation field list/options helpers.
  - `SearchTraces()`, `ListTraceFields()`, `GetTraceByID()`, `GetTraceSpans()`, `FetchTraceFieldOptions()`, and `GetSpanByID()`.
- RAG API:
  - Search index summary, unregister, metrics update, detail, summary-field, schema detail, schema summary, and schema file operations.
  - `RegisterSearchIndex()` migrated to the Python SDK v2.4.12 `v1/rag/indexes` route.
- Workflows API:
  - `GetWorkflowExecutionLogs()` for paginated workflow execution logs.
  - `StreamWorkflowExecutionLogs()` for SSE log streaming.
  - `TriggerSchedule()` for immediate schedule execution.
  - Workflow list filters for status, deployment name, deployment status, tags, sort field, and order.
  - Workflow run filters for deployment name, sort field, order, start time, and end time.
- Connectors API:
  - Optional `github_installation_link` query parameter for `GetConnectorAuthURL()`.

### Changed

- Version and User-Agent updated to `2.4.12`.
- README compatibility notes updated for Python SDK v2.4.12.
- Query serialization expanded for RFC3339Nano timestamps and typed order/status/sort enums.
- `ListSearchIndexes()` now uses `v1/rag/indexes/summary`.
- Library and access delete helpers now tolerate empty `204 No Content` responses.

### Tests

- Added mock-server parity coverage for the new 2.4.12 endpoints.
- Added stream coverage for workflow execution logs.
- Added binary response coverage for search index schema files.
- Verified focused parity tests and compile-only SDK validation.

## [2.4.9] - 2026-06-04

### Added - Python SDK v2.4.9 Parity Updates

- Batch API:
  - `DeleteBatchJob()` for `DELETE /v1/batch/jobs/{job_id}`.
- Chat API:
  - String-or-list `stop` values.
  - Full `response_format` objects for JSON schema mode.
  - Built-in tool payloads, specific tool choices, `reasoning_effort`, `guardrails`, and `prompt_cache_key`.
- Connectors API:
  - Organization, workspace, and user activate/deactivate operations.
  - OAuth2 server metadata fields and connector `protocol` request fields.
- RAG API:
  - Search index list and register operations.
- Workflows API:
  - Bulk archive/unarchive operations.
  - Schedule list filters, get schedule, and update schedule operations.
- Models API:
  - Updated model card and capability fields from Python SDK v2.4.9.
- Audio realtime API:
  - `RealtimeTranscriptionConnect()` for WebSocket session management.
  - `RealtimeTranscribeStream()` convenience helper for streaming audio from an `io.Reader`.
  - Session update, audio append, flush, end, close, and event-reading helpers.

### Changed

- Version and User-Agent updated to `2.4.9`.
- README compatibility notes updated for Python SDK v2.4.9.
- Added `nhooyr.io/websocket` as a pure-Go dependency to support official realtime API parity.

### Tests

- Added mock-server parity coverage for the new 2.4.9 endpoints and chat request fields.
- Verified compile-only test sweep with `go test ./... -run '^$'`.

## [2.4.4] - 2026-05-02

### Added - Python SDK v2.4.4 Parity Updates

- Audio/Speech and Voices API:
  - `Speech()` and `SpeechStream()`
  - Voice list/create/get/update/delete operations
  - Voice sample audio download
- Connectors API:
  - Connector create/list/get/update/delete operations
  - Connector authentication URL and authentication method helpers
  - Connector tool listing and tool call support
  - Organization, workspace, and user credentials management
- RAG API:
  - Ingestion pipeline configuration list/register/update-run-info operations
- Workflows API:
  - Workflow list, registration list/get, update, archive, unarchive
  - Workflow execution and registration execution
  - `ExecuteWorkflowAndWait()` and `WaitForWorkflowCompletion()`
  - Workflow deployments, metrics, runs, events, and schedules
- Workflow executions API:
  - Get/history, signal, query, update, reset
  - Cancel/terminate and batch cancel/terminate
  - Trace OTEL, trace summary, trace events, and execution stream support
- Observability API:
  - Campaigns, datasets, dataset records, imports/exports/tasks
  - Chat completion fields and chat completion events
  - Judges and live judging helpers

### Changed

- Version and User-Agent updated to `2.4.4`.
- README release banner and compatibility notes updated for Python SDK v2.4.4.
- Added broad flexible `APIResponse` support for generated endpoints with large/fast-moving response schemas.

### Tests

- Added mock-server parity coverage for new endpoint groups.
- Verified compile-only SDK test target.
- Full live/API-dependent suite still requires valid API credentials and deterministic model responses.

## [2.2.0] - 2026-03-03

### Added - Python SDK v1.12.4 Parity Updates

- Batch API improvements:
  - Inline request validation for `CreateBatchJob()` (`input_files` xor `requests`)
  - Support for `agent_id` and `order_by`
  - `GetBatchJob(jobID, inline ...bool)` for inline retrieval
- Classifiers API:
  - `ModerateChat()`
  - `Classify()`
  - `ClassifyChat()`
- Audio/Transcriptions API:
  - `TranscribeStream()` (SSE)
  - `diarize` and `context_bias` support
- Conversations API:
  - Stream operations (`StartConversationStream`, `AppendToConversationStream`, `RestartConversationStream`)
  - `GetConversationMessages()`
  - Extended request parameter support
- Beta Agents API enhancements:
  - Version management and alias CRUD operations
  - Extended list/get/update support
- Documents/Libraries API enhancements:
  - Documents: advanced list params, text/signed URL helpers, reprocess
  - Libraries: `chunk_size` support in create

### Changed

- Core request URL handling fixed to preserve query strings.
- Core response JSON decoding now supports non-object payloads.
- Files API list query support expanded (`include_total`, `mimetypes`) and direct HTTP calls include `User-Agent`.
- Library update uses `PUT` for parity.
- Access/share API aligned with beta share semantics while preserving backward-compatible wrappers.

### Tests

- Added targeted regression tests for:
  - Batch validation paths
  - Classifier chat/text classification endpoints
  - Conversation messages and params
  - Query param behavior (`inline`, `include_total`, `mimetypes`)

## [2.1.0] - 2025-12-19

### Added - Python SDK v1.10.0 Compatibility

#### New Types
- **`RequestSource`** - Type for filtering agents by source (API, playground, agent_builder_v1)
- **`OCRTableObject`** - Table extraction from documents with markdown/HTML format support
- **`ModelCapabilities`** - Model capability flags including new classification, moderation, and audio fields

#### New Operations
- **`DeleteMistralAgent()`** - Delete agents
- **`DeleteConversation()`** - Delete conversations

#### New Parameters
- **`Metadata`** field added to:
  - `ChatRequestParams` - Custom metadata for chat completions
  - `FIMRequestParams` - Custom metadata for FIM completions
  - `AgentCompletionRequest` - Custom metadata for agent completions
- **`IncludeTotal`** - Optional parameter in `ListFilesParams`

### Changed - Breaking Changes

#### Library API
- **Field Renames:**
  - `Created` → `CreatedAt`
  - `Updated` → `UpdatedAt`
- **New Required Fields:**
  - `OwnerID` (*string, nullable)
  - `OwnerType` (string)
  - `TotalSize` (int)
  - `NbDocuments` (int)
  - `ChunkSize` (*int, nullable)
- **New Optional Fields:**
  - `Emoji`, `GeneratedDescription`, `ExplicitUserMembersCount`, `ExplicitWorkspaceMembersCount`, `OrgSharingRole`, `GeneratedName`

#### Document API
- **Field Renames:**
  - `Created` → `CreatedAt`
  - `Status` → `ProcessingStatus`
- **Field Type Changes:**
  - `Size` → *int64 (now nullable)
- **New Required Fields:**
  - `Hash`, `MimeType`, `Extension` (all nullable strings)
  - `UploadedByID` (*string, nullable)
  - `UploadedByType` (string)
  - `TokensProcessingTotal` (int)
- **New Optional Fields:**
  - `Summary`, `LastProcessedAt`, `NumberOfPages`, `TokensProcessingMainContent`, `TokensProcessingSummary`, `URL`, `Attributes`
- **New Request Fields:**
  - `UpdateDocumentRequest.Attributes` (map[string]any)

#### Files API
- **Field Type Changes:**
  - `ListFilesOut.Total` → *int (now nullable/optional)

#### Model API
- **New Fields:**
  - `ModelCard.Capabilities` (*ModelCapabilities) - includes classification, moderation, audio flags

#### OCR API
- **Enhanced Fields:**
  - `OCRPageObject.Tables` ([]OCRTableObject) - extracted tables
  - `OCRPageObject.Hyperlinks` ([]string) - page hyperlinks
  - `OCRPageObject.Header` (*string) - page header
  - `OCRPageObject.Footer` (*string) - page footer

### Migration Guide

#### Library Type Updates
```go
// Before v2.1.0
library.Created  // int64
library.Updated  // int64

// After v2.1.0
library.CreatedAt  // int64
library.UpdatedAt  // int64
library.OwnerID    // *string (new, nullable)
library.OwnerType  // string (new, required)
```

#### Document Type Updates
```go
// Before v2.1.0
doc.Status  // string
doc.Size    // int64

// After v2.1.0
doc.ProcessingStatus  // string (renamed)
doc.Size              // *int64 (now nullable)
doc.UploadedByType    // string (new, required)
doc.Attributes        // map[string]any (new, optional)
```

#### Files API Updates
```go
// Before v2.1.0
response.Total  // int

// After v2.1.0
response.Total  // *int (now nullable)
if response.Total != nil {
    total := *response.Total
}
```

#### Metadata Parameters
```go
// Chat with metadata
params := &ChatRequestParams{
    Metadata: map[string]any{
        "user_id": "123",
        "session": "abc",
    },
}

// FIM with metadata
fimParams := &FIMRequestParams{
    Metadata: map[string]any{
        "tracking": "enabled",
    },
}
```

### Documentation
- Added `docs/BREAKING_CHANGES_v1.10.0.md` - Comprehensive breaking changes documentation
- Added `docs/V2.1.0_IMPLEMENTATION_PLAN.md` - Implementation plan and tracking
- Updated all affected types with inline documentation
- All tests updated to reflect new types

## [2.0.1] - 2025-11-20

### Fixed
- **Critical**: Added `User-Agent` header to all HTTP requests to prevent Cloudflare 400 errors
  - Cloudflare-protected endpoints require a User-Agent header or they reject requests with 400 Bad Request
  - All requests now include `User-Agent: mistral-go/2.0.1` header
  - Added comprehensive test coverage to prevent regression
  - This fix resolves issues when using the SDK directly or through proxies/gateways

## [2.0.0] - 2025-11-20

### 🎉 Major Release - 100% Feature Parity Achieved!

This is a **major milestone** release that brings the Go SDK to **100% feature parity** with the official Mistral Python SDK!

### Added

#### Core API Enhancements
- **Models API** - Complete implementation with all CRUD operations
  - `RetrieveModel()` - Get model details
  - `DeleteModel()` - Delete fine-tuned models
  - `UpdateModel()` - Update model metadata
  - `ArchiveModel()` / `UnarchiveModel()` - Model archival

- **Embeddings API** - Full parameter support
  - `EmbeddingsWithParams()` - Advanced embeddings control
  - `encoding_format` parameter (float, base64)
  - `output_dimension` parameter
  - `output_dtype` parameter (float32, ubinary)

- **FIM API** - Streaming support
  - `FIMStream()` - Streaming code completions
  - All parameters: `top_p`, `min_tokens`, `random_seed`, `stop`
  - Pointer-based optional parameters

#### New APIs - OCR and Audio
- **OCR API** (`ocr.go`) - Document processing
  - `ProcessOCR()` - Process documents with OCR
  - `ProcessOCRFromURL()` - Process from URL
  - `ProcessOCRFromBase64()` - Process from base64
  - `ProcessOCRFromFileID()` - Process from uploaded file
  - Page selection, image extraction, bounding boxes

- **Audio/Transcriptions API** (`audio.go`) - Speech-to-text
  - `Transcribe()` - Transcribe audio files
  - `TranscribeFromURL()` - Transcribe from URL
  - `TranscribeFromFileID()` - Transcribe from file ID
  - Word and segment-level timestamps
  - Language detection and specification

#### Beta Features - Complete Implementation
- **Conversations API** (`conversations.go`) - Multi-turn conversations
  - `StartConversation()` - Start new conversations
  - `ListConversations()` - List all conversations
  - `GetConversation()` - Get conversation details
  - `AppendToConversation()` - Continue conversations
  - `GetConversationHistory()` - Full history tracking
  - `RestartConversation()` - Restart from any point

- **Libraries API** (`libraries.go`) - RAG document libraries
  - `ListLibraries()` - List all libraries
  - `CreateLibrary()` - Create new library
  - `GetLibrary()` - Get library details
  - `UpdateLibrary()` - Update library metadata
  - `DeleteLibrary()` - Delete library

- **Documents API** (`documents.go`) - Document management
  - `ListDocuments()` - List documents in library
  - `UploadDocument()` - Upload documents (multipart)
  - `GetDocument()` - Get document details
  - `UpdateDocument()` - Update document metadata
  - `DeleteDocument()` - Delete document
  - `GetDocumentStatus()` - Check processing status

- **Accesses API** (`accesses.go`) - Access control
  - `ListLibraryAccesses()` - List all accesses
  - `UpdateOrCreateLibraryAccess()` - Grant/update access
  - `DeleteLibraryAccess()` - Revoke access
  - Multi-tenant support

- **Mistral Agents API** (`mistral_agents.go`) - Beta agents
  - `CreateMistralAgent()` - Create custom agents
  - `ListMistralAgents()` - List all agents
  - `GetMistralAgent()` - Get agent details
  - `UpdateMistralAgent()` - Update agent configuration

### Changed
- **BREAKING:** Module path updated to `github.com/ZaguanLabs/mistral-go/v2`
- **FIM API:** Parameters now use pointers for proper optional handling
- **OCR API:** `Document` type renamed to `OCRDocument` to avoid conflicts
- All APIs now follow consistent pointer-based optional parameter patterns

### Statistics
- **APIs Implemented:** 12/12 (100%)
- **Total Methods:** 100+
- **New Code:** ~10,000 lines across 5 new API modules
- **Feature Parity:** 100% with Python SDK
- **External Dependencies:** 0 (pure Go)

### Documentation
- Added `docs/BETA_FEATURES_IMPLEMENTATION.md` - Complete Beta features guide
- Added `docs/OCR_AUDIO_IMPLEMENTATION.md` - OCR and Audio API guide
- Added `docs/MODELS_API_COMPLETION.md` - Models API completion guide
- Added `docs/EMBEDDINGS_FIM_COMPLETION.md` - Enhanced APIs guide
- Added `docs/FEATURE_COMPARISON.md` - Now shows 100% parity
- Added `docs/IMPROVEMENTS.md` - Zero dependencies achievement
- Added `docs/AUDIT_REPORT.md` - Comprehensive security audit (A- grade, 0 vulnerabilities)
- Added `docs/TEST_COVERAGE_REPORT.md` - Test coverage analysis (65.2%)
- Updated `README.md` - Comprehensive examples for all APIs
- Added `CHANGELOG.md` - This file
- Added `sdk/version.go` - Centralized version management

### Migration Guide

#### From v1.1.0 to v2.0.0

**Import Path Change:**
```go
// Old (v1.1.0)
import "github.com/ZaguanLabs/mistral-go/sdk"

// New (v2.0.0)
import "github.com/ZaguanLabs/mistral-go/v2/sdk"
```

**FIM API Changes:**
```go
// Old (v1.1.0)
params := &sdk.FIMRequestParams{
    Model:       "codestral-latest",
    Prompt:      "code",
    Suffix:      "more code",  // string
    MaxTokens:   100,           // int
    Temperature: 0.1,           // float64
}

// New (v2.0.0)
params := &sdk.FIMRequestParams{
    Model:       "codestral-latest",
    Prompt:      "code",
    Suffix:      sdk.StringPtr("more code"),  // *string
    MaxTokens:   sdk.IntPtr(100),             // *int
    Temperature: sdk.Float64Ptr(0.1),         // *float64
}
```

**New APIs in v2.0.0:**
```go
// OCR API - Not available in v1.1.0
response, err := client.ProcessOCRFromURL(
    "pixtral-12b-2409",
    "https://example.com/doc.pdf",
    &sdk.OCRRequest{
        Pages: []int{0, 1, 2},
    },
)

// Audio API - Not available in v1.1.0
response, err := client.Transcribe(
    "whisper-large-v3",
    audioFile,
    "audio.mp3",
    nil,
)

// Beta Features - Not available in v1.1.0
conv, err := client.StartConversation(&sdk.ConversationStartRequest{
    Inputs: []sdk.ConversationInput{
        {Type: "text", Content: "Hello"},
    },
})
```

## [1.1.0] - Previous Release

Previous release with core functionality:
- Chat completions and streaming
- Embeddings (basic)
- FIM (basic)
- Models (basic - list only)
- Files API
- Fine-tuning API
- Batch API
- Agents API
- Classifiers API

## [1.0.0] - Initial Release

Initial release with basic functionality.

---

[2.2.0]: https://github.com/ZaguanLabs/mistral-go/compare/v2.1.0...v2.2.0
[2.1.0]: https://github.com/ZaguanLabs/mistral-go/compare/v2.0.1...v2.1.0
[2.0.1]: https://github.com/ZaguanLabs/mistral-go/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/ZaguanLabs/mistral-go/compare/v1.1.0...v2.0.0
[1.1.0]: https://github.com/ZaguanLabs/mistral-go/releases/tag/v1.1.0
[1.0.0]: https://github.com/ZaguanLabs/mistral-go/releases/tag/v1.0.0
