# Changelog

All notable changes to the Mistral Go SDK will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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

[2.0.0]: https://github.com/ZaguanLabs/mistral-go/compare/v1.1.0...v2.0.0
[1.1.0]: https://github.com/ZaguanLabs/mistral-go/releases/tag/v1.1.0
[1.0.0]: https://github.com/ZaguanLabs/mistral-go/releases/tag/v1.0.0
