# Project Structure

This document describes the organization of the Mistral Go SDK project.

## Directory Layout

```
mistral-go/
├── sdk/                    # Main SDK package
│   ├── chat.go            # Chat completion implementation
│   ├── chat_test.go       # Chat completion tests
│   ├── client.go          # Core client implementation
│   ├── embeddings.go      # Embeddings API
│   ├── embeddings_test.go # Embeddings tests
│   ├── errors.go          # Error types
│   ├── fim.go             # Fill-in-the-middle API
│   ├── fim_test.go        # FIM tests
│   ├── models.go          # Models API
│   └── types.go           # Type definitions and constants
├── examples/              # Usage examples
│   └── main.go           # Basic example
├── docs/                  # Documentation and reference materials
│   └── client-python-1.9.11/  # Python SDK reference
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── README.md              # Main documentation
├── IMPROVEMENTS.md        # Changelog and improvements
├── STRUCTURE.md           # This file
├── LICENSE                # MIT License
└── .gitignore            # Git ignore rules

## Package Structure

### `sdk` Package

The main SDK package contains all the core functionality:

- **Client Management** (`client.go`)
  - `MistralClient` - Main client struct
  - `NewMistralClient()` - Create client with custom config
  - `NewMistralClientDefault()` - Create client with defaults
  - `NewCodestralClientDefault()` - Create Codestral client

- **Chat Completions** (`chat.go`)
  - `Chat()` - Synchronous chat completions
  - `ChatStream()` - Streaming chat completions
  - `ChatRequestParams` - Request parameters
  - `ChatCompletionResponse` - Response types

- **Embeddings** (`embeddings.go`)
  - `Embeddings()` - Generate text embeddings
  - `EmbeddingResponse` - Response types

- **Fill-in-the-Middle** (`fim.go`)
  - `FIM()` - Code completion
  - `FIMRequestParams` - FIM parameters

- **Models** (`models.go`)
  - `ListModels()` - List available models
  - `ModelList` - Model information

- **Types** (`types.go`)
  - Model constants
  - Role constants
  - Type definitions
  - Helper functions

- **Errors** (`errors.go`)
  - `MistralError` - Base error type
  - `MistralAPIError` - API errors
  - `MistralConnectionError` - Connection errors

## Import Path

To use the SDK in your project:

```go
import "github.com/ZaguanLabs/mistral-go/sdk"
```

## Usage Example

```go
package main

import (
    "log"
    "github.com/ZaguanLabs/mistral-go/sdk"
)

func main() {
    client := sdk.NewMistralClientDefault("")
    
    response, err := client.Chat(
        sdk.ModelMistralSmallLatest,
        []sdk.ChatMessage{
            sdk.UserMessage("Hello!"),
        },
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println(response.Choices[0].Message.Content)
}
```

## Testing

Run tests with:

```bash
go test ./sdk/...
```

Build the SDK:

```bash
go build ./sdk/...
```

Build examples:

```bash
go build examples/main.go
```

## Future Organization

As the SDK grows, consider organizing into sub-packages:

```
sdk/
├── chat/          # Chat completions
├── embeddings/    # Embeddings
├── fim/           # Fill-in-the-middle
├── models/        # Models API
├── files/         # Files API (future)
├── agents/        # Agents API (future)
├── batch/         # Batch API (future)
└── internal/      # Internal utilities
```

This would allow for better separation of concerns and easier maintenance as new features are added.

## Migration from Old Structure

If you were using the old import path:

**Before:**
```go
import "github.com/ZaguanLabs/mistral-go"

client := mistral.NewMistralClientDefault("")
```

**After:**
```go
import "github.com/ZaguanLabs/mistral-go/sdk"

client := sdk.NewMistralClientDefault("")
```

All functionality remains the same, only the import path and package name have changed.
