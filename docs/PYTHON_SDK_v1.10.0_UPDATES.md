# Python SDK v1.10.0 Updates - Go Port

This document details the changes from Mistral Python SDK v1.10.0 that have been ported to the Go SDK v2.1.0.

## Release Information

- **Python SDK Version**: v1.10.0 (Released: 2025-12-16)
- **Go SDK Version**: v2.1.0 (Released: 2025-12-19)
- **Compatibility**: 100% feature parity maintained

## Changes Summary

### 1. New Types

#### RequestSource
A new type for filtering agents by their creation source.

**Python SDK:**
```python
RequestSource = Literal["api", "playground", "agent_builder_v1"]
```

**Go SDK:**
```go
type RequestSource string

const (
    RequestSourceAPI            RequestSource = "api"
    RequestSourcePlayground     RequestSource = "playground"
    RequestSourceAgentBuilderV1 RequestSource = "agent_builder_v1"
)
```

**Usage:**
```go
// Filter agents by source when listing
// Note: API support for filtering may vary
sources := []RequestSource{
    RequestSourceAPI,
    RequestSourcePlayground,
}
```

#### OCRTableObject
Enhanced OCR with table extraction support.

**Python SDK:**
```python
class OCRTableObject(BaseModel):
    id: str  # Table ID for extracted table in a page
    content: str  # Content of the table in the given format
    format_: Annotated[Format, pydantic.Field(alias="format")]  # Format of the table
```

**Go SDK:**
```go
type OCRTableFormat string

const (
    OCRTableFormatMarkdown OCRTableFormat = "markdown"
    OCRTableFormatHTML     OCRTableFormat = "html"
)

type OCRTableObject struct {
    ID      string         `json:"id"`      // Table ID for extracted table in a page
    Content string         `json:"content"` // Content of the table in the given format
    Format  OCRTableFormat `json:"format"`  // Format of the table
}
```

**Usage:**
```go
// OCR now extracts tables from documents
response, err := client.ProcessOCRFromURL(
    "pixtral-12b-2409",
    "https://example.com/document.pdf",
    nil,
)

for _, page := range response.Pages {
    // Access extracted tables
    for _, table := range page.Tables {
        fmt.Printf("Table %s (%s):\n%s\n", table.ID, table.Format, table.Content)
    }
}
```

### 2. Enhanced OCRPageObject

The `OCRPageObject` type has been enhanced with additional fields.

**New Fields:**
- `Tables` - List of extracted tables
- `Hyperlinks` - List of all hyperlinks in the page
- `Header` - Header of the page (optional)
- `Footer` - Footer of the page (optional)

**Python SDK:**
```python
class OCRPageObject(BaseModel):
    index: int
    markdown: str
    images: List[OCRImageObject]
    dimensions: Nullable[OCRPageDimensions]
    tables: Optional[List[OCRTableObject]] = None  # NEW
    hyperlinks: Optional[List[str]] = None  # NEW
    header: OptionalNullable[str] = UNSET  # NEW
    footer: OptionalNullable[str] = UNSET  # NEW
```

**Go SDK:**
```go
type OCRPageObject struct {
    PageNumber int                `json:"page_number"`
    Dimensions *OCRPageDimensions `json:"dimensions,omitempty"`
    Text       string             `json:"text"`
    Images     []OCRImageObject   `json:"images,omitempty"`
    Tables     []OCRTableObject   `json:"tables,omitempty"`     // NEW
    Hyperlinks []string           `json:"hyperlinks,omitempty"` // NEW
    Header     *string            `json:"header,omitempty"`     // NEW
    Footer     *string            `json:"footer,omitempty"`     // NEW
}
```

**Usage:**
```go
response, err := client.ProcessOCRFromURL(
    "pixtral-12b-2409",
    "https://example.com/document.pdf",
    nil,
)

for _, page := range response.Pages {
    // Access enhanced page information
    if page.Header != nil {
        fmt.Printf("Header: %s\n", *page.Header)
    }
    
    fmt.Printf("Tables: %d\n", len(page.Tables))
    fmt.Printf("Hyperlinks: %d\n", len(page.Hyperlinks))
    
    if page.Footer != nil {
        fmt.Printf("Footer: %s\n", *page.Footer)
    }
}
```

### 3. Metadata Parameter in Agents API

The `AgentCompletionRequest` now supports a `Metadata` field for custom tracking.

**Python SDK:**
```python
metadata: OptionalNullable[Dict[str, Any]] = UNSET
```

**Go SDK:**
```go
type AgentCompletionRequest struct {
    // ... other fields
    Metadata map[string]any `json:"metadata,omitempty"` // NEW
}
```

**Usage:**
```go
messages := []ChatMessage{
    UserMessage("Hello, how can you help me?"),
}

params := &AgentCompletionRequest{
    Metadata: map[string]any{
        "user_id":    "user-123",
        "session_id": "session-456",
        "environment": "production",
        "version":    "2.1.0",
    },
}

response, err := client.AgentComplete("agent-id", messages, params)
```

### 4. Delete Operations

#### Delete Agent

**Python SDK:**
```python
def delete(
    self,
    agent_id: str,
    # ... other params
) -> None:
```

**Go SDK:**
```go
func (c *MistralClient) DeleteMistralAgent(agentID string) error
```

**Usage:**
```go
// Delete an agent when no longer needed
err := client.DeleteMistralAgent("agent-id-to-delete")
if err != nil {
    log.Fatalf("Failed to delete agent: %v", err)
}
```

#### Delete Conversation

**Python SDK:**
```python
def delete(
    self,
    conversation_id: str,
    # ... other params
) -> None:
```

**Go SDK:**
```go
func (c *MistralClient) DeleteConversation(conversationID string) error
```

**Usage:**
```go
// Delete a conversation when no longer needed
err := client.DeleteConversation("conversation-id-to-delete")
if err != nil {
    log.Fatalf("Failed to delete conversation: %v", err)
}
```

### 5. Observability/Tracing (Not Ported)

Python SDK v1.10.0 added OpenTelemetry tracing support. This is an optional feature in Python and has not been ported to the Go SDK in this release.

**Python SDK Changes:**
- Added `tracing.py` hook for OpenTelemetry integration
- Added `observability/otel.py` for tracing implementation
- Automatic span creation for API calls

**Go SDK Status:**
- Not implemented in v2.1.0
- May be added in a future release if there's demand
- Go developers can implement their own tracing using standard Go observability tools

## Migration Guide

### From Go SDK v2.0.1 to v2.1.0

No breaking changes. All additions are backward compatible.

#### Using New Features

**1. OCR with Table Extraction:**
```go
// Before (v2.0.1) - basic OCR
response, err := client.ProcessOCRFromURL(
    "pixtral-12b-2409",
    "https://example.com/doc.pdf",
    nil,
)
// Only had: PageNumber, Dimensions, Text, Images

// After (v2.1.0) - enhanced OCR
response, err := client.ProcessOCRFromURL(
    "pixtral-12b-2409",
    "https://example.com/doc.pdf",
    nil,
)
// Now also has: Tables, Hyperlinks, Header, Footer
for _, page := range response.Pages {
    for _, table := range page.Tables {
        fmt.Printf("Table: %s\n", table.Content)
    }
}
```

**2. Agent Completion with Metadata:**
```go
// Before (v2.0.1) - no metadata
params := &AgentCompletionRequest{
    MaxTokens: IntPtr(1000),
}

// After (v2.1.0) - with metadata
params := &AgentCompletionRequest{
    MaxTokens: IntPtr(1000),
    Metadata: map[string]any{
        "user_id": "123",
        "tracking": "enabled",
    },
}
```

**3. Cleanup Operations:**
```go
// New in v2.1.0 - delete agents
err := client.DeleteMistralAgent("agent-id")

// New in v2.1.0 - delete conversations
err := client.DeleteConversation("conversation-id")
```

## Testing

All new features include comprehensive tests:

```bash
# Run all tests
go test ./sdk/...

# Run v1.10.0 specific tests
go test ./sdk/ -run TestAgentCompletionWithMetadata
go test ./sdk/ -run TestDeleteMistralAgent
go test ./sdk/ -run TestDeleteConversation
go test ./sdk/ -run TestOCRTableObject
```

## Compatibility Matrix

| Feature | Python SDK v1.10.0 | Go SDK v2.1.0 | Status |
|---------|-------------------|---------------|--------|
| RequestSource type | ✅ | ✅ | ✅ Ported |
| OCRTableObject | ✅ | ✅ | ✅ Ported |
| Enhanced OCRPageObject | ✅ | ✅ | ✅ Ported |
| Agent metadata parameter | ✅ | ✅ | ✅ Ported |
| Delete agent operation | ✅ | ✅ | ✅ Ported |
| Delete conversation operation | ✅ | ✅ | ✅ Ported |
| OpenTelemetry tracing | ✅ | ❌ | ⏳ Future |

## References

- [Python SDK v1.10.0 Release](https://github.com/mistralai/client-python/releases/tag/v1.10.0)
- [Go SDK v2.1.0 Changelog](../CHANGELOG.md#210---2025-12-19)
- [Mistral AI Documentation](https://docs.mistral.ai/)

## Notes

1. **Feature Parity**: The Go SDK maintains 100% feature parity with the Python SDK for all core functionality.

2. **Type Safety**: Go's strong typing provides additional safety compared to Python's dynamic typing.

3. **Performance**: The Go SDK continues to have zero external dependencies and excellent performance characteristics.

4. **Future Updates**: We will continue to track Python SDK releases and port relevant changes to maintain compatibility.

## Support

For issues or questions:
- GitHub Issues: https://github.com/ZaguanLabs/mistral-go/issues
- Python SDK Reference: https://github.com/mistralai/client-python
