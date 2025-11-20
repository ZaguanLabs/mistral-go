# Mistral Go SDK Examples

This directory contains comprehensive examples demonstrating all major features of the Mistral Go SDK v2.0.0.

## Prerequisites

Set your Mistral API key as an environment variable:

```bash
export MISTRAL_API_KEY="your-api-key-here"
```

For Codestral examples:
```bash
export CODESTRAL_API_KEY="your-codestral-key-here"
```

## Running Examples

Each example can be run independently:

```bash
# Chat completions
go run examples/chat/main.go

# Tool/Function calling
go run examples/tools/main.go

# Embeddings
go run examples/embeddings/main.go

# Fill-in-the-Middle (FIM) code completion
go run examples/fim/main.go

# Agents
go run examples/agents/main.go

# Content moderation
go run examples/moderation/main.go

# Models API
go run examples/models/main.go

# Basic example (all features)
go run examples/main.go
```

## Examples Overview

### 1. **chat/** - Chat Completions
- Simple chat completion
- Chat with parameters (temperature, max_tokens)
- Streaming responses
- System messages

### 2. **tools/** - Tool/Function Calling
- Define custom tools/functions
- Tool calling with auto mode
- Processing tool responses
- Multi-turn conversations with tools

### 3. **embeddings/** - Text Embeddings
- Generate embeddings for semantic search
- Custom embedding dimensions
- Batch embedding generation
- Different encoding formats

### 4. **fim/** - Fill-in-the-Middle
- Code completion with context
- Streaming FIM completions
- Custom parameters (temperature, stop sequences)
- Multi-language code completion

### 5. **agents/** - Mistral Agents
- Create custom agents
- List and manage agents
- Agent completions
- Agent streaming

### 6. **moderation/** - Content Moderation
- Text classification
- Safety scoring
- Batch moderation
- Category-based filtering

### 7. **models/** - Models API
- List available models
- Get model details
- Model capabilities
- Context length information

### 8. **main.go** - All-in-One Example
- Demonstrates multiple APIs
- Basic usage patterns
- Quick start reference

## API Coverage

All examples use the v2 import path:
```go
import "github.com/ZaguanLabs/mistral-go/v2/sdk"
```

## Additional Resources

- **[Main README](../README.md)** - Complete SDK documentation
- **[API Documentation](https://docs.mistral.ai)** - Official Mistral AI API docs
- **[Feature Comparison](../docs/FEATURE_COMPARISON.md)** - 100% parity with Python SDK
- **[Beta Features](../docs/BETA_FEATURES_IMPLEMENTATION.md)** - Conversations, Libraries, Documents

## Notes

- All examples include error handling
- Examples use environment variables for API keys
- Some examples require specific API access (e.g., Codestral, Agents)
- Streaming examples demonstrate proper channel handling
- Tool examples show complete request/response cycles

## Support

For issues or questions:
- Check the [main README](../README.md)
- Review the [documentation](../docs/)
- File an issue on GitHub
