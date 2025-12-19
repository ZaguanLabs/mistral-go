# Release Notes - Mistral Go SDK v2.1.0

**Release Date**: December 19, 2025  
**Python SDK Compatibility**: v1.10.0

## Overview

This release updates the Mistral Go SDK to maintain 100% feature parity with the Mistral Python SDK v1.10.0. All new features and enhancements from the Python SDK have been successfully ported to Go.

## What's New

### 1. RequestSource Type
Added support for filtering agents by their creation source.

```go
type RequestSource string

const (
    RequestSourceAPI            RequestSource = "api"
    RequestSourcePlayground     RequestSource = "playground"
    RequestSourceAgentBuilderV1 RequestSource = "agent_builder_v1"
)
```

### 2. Enhanced OCR with Table Extraction

#### OCRTableObject
New type for extracting tables from documents:

```go
type OCRTableObject struct {
    ID      string         `json:"id"`
    Content string         `json:"content"`
    Format  OCRTableFormat `json:"format"` // "markdown" or "html"
}
```

#### Enhanced OCRPageObject
Added new fields for better document structure extraction:
- `Tables` - List of extracted tables
- `Hyperlinks` - List of all hyperlinks in the page
- `Header` - Page header (optional)
- `Footer` - Page footer (optional)

**Example:**
```go
response, err := client.ProcessOCRFromURL(
    "pixtral-12b-2409",
    "https://example.com/document.pdf",
    nil,
)

for _, page := range response.Pages {
    fmt.Printf("Page %d:\n", page.PageNumber)
    
    // Access tables
    for _, table := range page.Tables {
        fmt.Printf("  Table %s (%s):\n%s\n", 
            table.ID, table.Format, table.Content)
    }
    
    // Access hyperlinks
    fmt.Printf("  Hyperlinks: %d\n", len(page.Hyperlinks))
    
    // Access header/footer
    if page.Header != nil {
        fmt.Printf("  Header: %s\n", *page.Header)
    }
}
```

### 3. Agent Metadata Support

Added `Metadata` field to `AgentCompletionRequest` for custom tracking:

```go
params := &AgentCompletionRequest{
    Metadata: map[string]any{
        "user_id":     "user-123",
        "session_id":  "session-456",
        "environment": "production",
    },
}

response, err := client.AgentComplete("agent-id", messages, params)
```

### 4. Delete Operations

#### Delete Agent
```go
err := client.DeleteMistralAgent("agent-id")
if err != nil {
    log.Fatalf("Failed to delete agent: %v", err)
}
```

#### Delete Conversation
```go
err := client.DeleteConversation("conversation-id")
if err != nil {
    log.Fatalf("Failed to delete conversation: %v", err)
}
```

## Breaking Changes

**None** - This release is fully backward compatible with v2.0.1.

## Migration Guide

### From v2.0.1 to v2.1.0

No code changes required. All new features are additive.

#### Optional: Use New Features

**OCR with Tables:**
```go
// Your existing OCR code continues to work
response, err := client.ProcessOCRFromURL(model, url, nil)

// New: Access table data
for _, page := range response.Pages {
    for _, table := range page.Tables {
        // Process table content
    }
}
```

**Agent Metadata:**
```go
// Your existing agent code continues to work
response, err := client.AgentComplete(agentID, messages, nil)

// New: Add metadata
params := &AgentCompletionRequest{
    Metadata: map[string]any{"tracking": "enabled"},
}
response, err := client.AgentComplete(agentID, messages, params)
```

**Cleanup Operations:**
```go
// New: Delete agents when done
err := client.DeleteMistralAgent("agent-id")

// New: Delete conversations when done
err := client.DeleteConversation("conversation-id")
```

## Testing

All new features include comprehensive tests:

```bash
# Run all tests
go test ./sdk/...

# Run specific v2.1.0 tests
go test ./sdk/ -run TestRequestSourceConstants
go test ./sdk/ -run TestOCRTableObject
go test ./sdk/ -run TestOCRPageObjectWithTables
```

## Documentation

- **Comprehensive Guide**: `docs/PYTHON_SDK_v1.10.0_UPDATES.md`
- **Changelog**: `CHANGELOG.md`
- **API Reference**: All new types and methods include inline documentation

## Compatibility

| Component | Version | Status |
|-----------|---------|--------|
| Go SDK | v2.1.0 | ✅ Current |
| Python SDK | v1.10.0 | ✅ Compatible |
| Go Version | 1.21+ | ✅ Required |
| Feature Parity | 100% | ✅ Maintained |

## Files Changed

### Modified
- `sdk/types.go` - Added RequestSource type
- `sdk/ocr.go` - Enhanced OCR types with table support
- `sdk/agents.go` - Added metadata parameter
- `sdk/mistral_agents.go` - Added DeleteMistralAgent method
- `sdk/conversations.go` - Added DeleteConversation method
- `sdk/version.go` - Updated to v2.1.0
- `CHANGELOG.md` - Added v2.1.0 release notes

### Added
- `sdk/agents_v1_10_test.go` - Tests for v2.1.0 features
- `docs/PYTHON_SDK_v1.10.0_UPDATES.md` - Comprehensive update guide
- `docs/RELEASE_NOTES_v2.1.0.md` - This file

## Known Limitations

**OpenTelemetry Tracing**: The Python SDK v1.10.0 added optional OpenTelemetry tracing support. This feature has not been ported to the Go SDK in this release. Go developers can implement their own tracing using standard Go observability tools if needed.

## Performance

- Zero external dependencies maintained
- No performance regressions
- All operations remain efficient and type-safe

## Support

- **Issues**: https://github.com/ZaguanLabs/mistral-go/issues
- **Documentation**: https://github.com/ZaguanLabs/mistral-go
- **Python SDK**: https://github.com/mistralai/client-python

## Next Steps

1. Update your dependency:
   ```bash
   go get github.com/ZaguanLabs/mistral-go/v2@v2.1.0
   ```

2. Review the new features in `docs/PYTHON_SDK_v1.10.0_UPDATES.md`

3. Optionally integrate new capabilities:
   - Use table extraction in OCR workflows
   - Add metadata tracking to agent completions
   - Implement cleanup with delete operations

## Acknowledgments

This release maintains our commitment to 100% feature parity with the official Mistral Python SDK while providing the performance and type safety benefits of Go.

---

**Full Changelog**: https://github.com/ZaguanLabs/mistral-go/compare/v2.0.1...v2.1.0
