# Mistral Go SDK Improvements

## Overview
This document details the comprehensive improvements made to the Mistral Go SDK based on deep analysis of the official Mistral Python SDK v1.9.11.

## ⭐ Zero Dependencies Achieved

**Major Achievement:** The SDK now has **ZERO external dependencies** - not even for tests!

**What changed:**
- Removed `github.com/stretchr/testify` (and its 3 indirect dependencies)
- Converted all test assertions to use Go's standard `testing` package
- Achieved 100% standard library usage across production and test code

**Benefits:**
- ✅ **Smaller binary size** - No external code to compile
- ✅ **Faster builds** - No dependency downloads
- ✅ **Better security** - Minimal attack surface, no supply chain risks
- ✅ **Maximum compatibility** - Works everywhere Go works
- ✅ **Easier maintenance** - No dependency updates to track
- ✅ **Production ready** - Enterprise-grade reliability

**go.mod:**
```go
module github.com/ZaguanLabs/mistral-go

go 1.20
```

That's it! No `require` section, no `go.sum` file. Pure Go standard library.

## Removed Hardcoded Model Constants

**Why this change?**
Model IDs change frequently as Mistral releases new versions. Hardcoded constants become outdated quickly and can mislead developers into using deprecated models.

**What changed:**
- Removed all hardcoded model constants (e.g., `ModelMistralSmallLatest`, `ModelMistralTiny`, etc.)
- Added comprehensive documentation in `types.go` explaining how to get models dynamically
- Updated all examples and tests to use string literals for model IDs

**How to get available models:**
```go
// Dynamically fetch current models
models, err := client.ListModels()
if err != nil {
    log.Fatal(err)
}

for _, model := range models.Data {
    fmt.Printf("Model: %s\n", model.ID)
}
```

**Common model ID patterns** (subject to change):
- `"mistral-large-latest"` - Latest large model
- `"mistral-small-latest"` - Latest small model
- `"codestral-latest"` - Latest code model
- `"mistral-embed"` - Embedding model

For production use:
1. Fetch models dynamically with `ListModels()`
2. Use specific versioned model IDs (e.g., `"mistral-large-2411"`)
3. Define your own constants based on your requirements

## Project Restructuring

The project has been reorganized for better maintainability:

**New Structure:**
- All Go source files moved to `sdk/` directory
- Package renamed from `mistral` to `sdk`
- Cleaner root directory with only documentation and configuration files
- Import path changed to `github.com/ZaguanLabs/mistral-go/sdk`

**Migration:**
```go
// Before
import "github.com/ZaguanLabs/mistral-go"
client := mistral.NewMistralClientDefault("")

// After
import "github.com/ZaguanLabs/mistral-go/sdk"
client := sdk.NewMistralClientDefault("")
```

See [STRUCTURE.md](STRUCTURE.md) for detailed project organization.

## Major Improvements Implemented

### 1. Enhanced Chat Completion Parameters
**Added missing parameters to match Python SDK:**
- `PresencePenalty` - Penalizes repetition of words/phrases
- `FrequencyPenalty` - Penalizes based on frequency in generated text
- `N` - Number of completions to return for each request
- `Prediction` - Speculative decoding prediction support
- `ParallelToolCalls` - Enable parallel tool calls
- `PromptMode` - Toggle between reasoning mode and no system prompt
- `MinTokens` - Minimum tokens to generate (for FIM)
- `Stop` - Stop sequences support (array of strings)

**Changed parameter structure:**
- All optional parameters now use pointers (`*int`, `*float64`, `*bool`) for proper nil handling
- This matches the Python SDK's `OptionalNullable` pattern
- Allows distinguishing between "not set" and "set to zero/false"

### 2. Type System Enhancements

**New message types and helpers:**
- `SystemMessage(content)` - Helper to create system messages
- `UserMessage(content)` - Helper to create user messages  
- `AssistantMessage(content)` - Helper to create assistant messages
- `ToolMessage(toolCallID, content)` - Helper to create tool messages

**Extended ChatMessage struct:**
- Added `ToolCallID` field for tool role messages
- Added `Name` field for function/tool messages
- Made `Content` optional with `omitempty`

**New types:**
- `Prediction` - For speculative decoding
- `MistralPromptMode` - For prompt mode control
- `EncodingFormat` - For embedding encoding format
- `EmbeddingDtype` - For embedding data type
- Additional `FinishReason` constants: `FinishReasonToolCalls`, `FinishReasonModelLength`

**Pointer helper functions:**
- `Float64Ptr(v)` - Create pointer to float64
- `IntPtr(v)` - Create pointer to int
- `BoolPtr(v)` - Create pointer to bool
- `StringPtr(v)` - Create pointer to string
- `MistralPromptModePtr(v)` - Create pointer to MistralPromptMode

### 3. Model Constants
**Added new model constants:**
- `ModelMistralEmbed` - For embeddings
- `ModelCodestral2405` - Specific Codestral version

### 4. API Consistency

**Parameter handling:**
- Changed from required default values to optional pointer-based parameters
- Removed `DefaultChatRequestParams` in favor of `NewChatRequestParams()`
- All parameters now properly omitted when not set (using `omitempty` tags)

**Request building:**
- Updated `Chat()` and `ChatStream()` methods to only include set parameters in requests
- Proper nil checking before adding parameters to request payload
- Matches Python SDK's behavior of not sending unset optional parameters

## Comparison with Python SDK

### Python SDK Structure (v1.9.11)
```python
class Mistral:
    models: Models
    beta: Beta  
    files: Files
    fine_tuning: FineTuning
    batch: Batch
    chat: Chat
    fim: Fim
    agents: Agents
    embeddings: Embeddings
    classifiers: Classifiers
    ocr: Ocr
    audio: Audio
```

### Current Go SDK Coverage
✅ **Implemented:**
- Chat completions (complete)
- Chat streaming (complete)
- Embeddings (basic)
- FIM completions (basic)
- Models listing (basic)

❌ **Not Yet Implemented (Future Work):**
- Files API (upload, list, retrieve, delete, download)
- Agents API (complete, stream)
- Batch API (jobs management)
- Fine-tuning API (jobs, models)
- Classifiers API (moderate, classify)
- OCR API (process)
- Audio/Transcriptions API
- Beta features (conversations, libraries, documents)
- Context manager pattern for resource management
- Retry configuration with backoff strategy
- Custom HTTP headers per request
- Server URL customization per request
- Timeout configuration per request

## Breaking Changes

### 1. Import Path
**Before:**
```go
import "github.com/ZaguanLabs/mistral-go"
```

**After:**
```go
import "github.com/ZaguanLabs/mistral-go/sdk"
```

### 2. Package Name
**Before:**
```go
client := mistral.NewMistralClientDefault("")
```

**After:**
```go
client := sdk.NewMistralClientDefault("")
```

### 3. Parameter Structure
**Before:**
```go
params := DefaultChatRequestParams
params.Temperature = 0.7
params.MaxTokens = 100
```

**After:**
```go
params := sdk.NewChatRequestParams()
params.Temperature = sdk.Float64Ptr(0.7)
params.MaxTokens = sdk.IntPtr(100)
```

### Why These Changes?
1. **Proper nil handling** - Can distinguish between "not set" and "set to zero"
2. **API compatibility** - Matches Mistral's API expectations
3. **Flexibility** - Allows using model defaults when parameters aren't specified
4. **Python SDK alignment** - Matches the `OptionalNullable` pattern
5. **Better organization** - Cleaner project structure

## Usage Examples

### Basic Chat with New Parameters
```go
client := sdk.NewMistralClientDefault("")

params := sdk.NewChatRequestParams()
params.Temperature = sdk.Float64Ptr(0.7)
params.MaxTokens = sdk.IntPtr(500)
params.PresencePenalty = sdk.Float64Ptr(0.1)
params.FrequencyPenalty = sdk.Float64Ptr(0.1)
params.Stop = []string{"END", "STOP"}

response, err := client.Chat(
    sdk.ModelMistralSmallLatest,
    []sdk.ChatMessage{
        sdk.UserMessage("Hello, how are you?"),
    },
    params,
)
```

### Using Reasoning Mode
```go
params := sdk.NewChatRequestParams()
params.PromptMode = sdk.MistralPromptModePtr(sdk.PromptModeReasoning)
params.Temperature = sdk.Float64Ptr(0.3)

response, err := client.Chat(model, messages, params)
```

### Multiple Completions
```go
params := sdk.NewChatRequestParams()
params.N = sdk.IntPtr(3) // Get 3 different completions
params.Temperature = sdk.Float64Ptr(0.8)

response, err := client.Chat(model, messages, params)
// response.Choices will contain 3 different completions
```

### Tool Calling with Parallel Execution
```go
params := sdk.NewChatRequestParams()
params.Tools = []sdk.Tool{/* your tools */}
params.ToolChoice = sdk.ToolChoiceAuto
params.ParallelToolCalls = sdk.BoolPtr(true)

response, err := client.Chat(model, messages, params)
```

## Testing Updates

All tests have been updated to use the new parameter structure:
- Fixed parameter assignments to use pointer helpers
- Updated all test cases to pass pointers instead of values
- Tests now properly demonstrate the new API usage

## Future Enhancements

### High Priority
1. **Files API** - Essential for fine-tuning workflows
2. **Fine-tuning API** - Model customization support
3. **Batch API** - Efficient bulk processing
4. **Retry logic** - Automatic retry with exponential backoff
5. **Context/Options pattern** - Per-request configuration

### Medium Priority
1. **Agents API** - Agentic workflows
2. **Classifiers API** - Content moderation
3. **Audio API** - Transcription support
4. **Better error types** - Structured error responses
5. **Streaming improvements** - Better error handling in streams

### Low Priority
1. **OCR API** - Document processing
2. **Beta features** - Conversations, libraries
3. **Provider SDKs** - Azure, GCP integration
4. **Async patterns** - Go context support
5. **Middleware/hooks** - Request/response interception

## Conclusion

These improvements significantly enhance the Go SDK's feature parity with the official Python SDK while maintaining idiomatic Go patterns. The SDK now supports all major chat completion parameters and provides a more flexible, Python SDK-aligned API.

The restructuring into the `sdk/` package provides better organization and sets the foundation for future growth as more APIs are added.

The pointer-based optional parameters may require migration effort for existing users, but provide much better API compatibility and flexibility going forward.
