# Mistral Go SDK

**Version 2.0.1** - Critical Bug Fix Release 🔧

[![Zero Dependencies](https://img.shields.io/badge/dependencies-0-brightgreen.svg)](https://github.com/ZaguanLabs/mistral-go)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue.svg)](https://golang.org/)
[![Standard Library Only](https://img.shields.io/badge/stdlib-only-success.svg)](https://pkg.go.dev/std)

A Go SDK for the Mistral AI API., designed to provide developers with powerful tools to integrate advanced AI capabilities into their applications. This SDK has been significantly enhanced to match the official Mistral Python SDK v1.9.11.

**✨ Zero Dependencies** - Built entirely with Go's standard library for maximum compatibility, security, and performance.

## Features

- **Chat Completions**: Full-featured chat completions with all parameters from the official API
- **Chat Completions Streaming**: Real-time streaming responses with proper SSE handling
- **Embeddings**: Generate vector embeddings for semantic search and ML applications
- **Fill-in-the-Middle (FIM)**: Code completion with context-aware suggestions
- **Models API**: List and retrieve available models
- **Tool/Function Calling**: Support for tools with parallel execution
- **Advanced Parameters**: Presence/frequency penalties, multiple completions, reasoning mode, and more
- **Files API**: Upload, list, retrieve, delete, and download files for fine-tuning and batch processing
- **Fine-tuning API**: Create and manage fine-tuning jobs for custom model training
- **Batch API**: Efficient bulk processing with batch jobs for cost-effective large-scale operations
- **Agents API**: Agentic workflows with tool use and streaming support
- **Classifiers API**: Content moderation and classification for safety and compliance
- **OCR API**: Document processing and text extraction from images
- **Audio/Transcriptions API**: Speech-to-text transcription with timestamp support

## Version 2.0.1 - Critical Bug Fix

🔧 **Fixed Cloudflare 400 Errors**: Added `User-Agent` header to all HTTP requests. This resolves issues where Cloudflare-protected endpoints (including the official Mistral API) would reject requests with 400 Bad Request errors.

This fix is critical for production deployments and resolves issues when using the SDK:
- Directly with Mistral API
- Through proxies or gateways (e.g., Zaguán)
- Behind any Cloudflare or similar CDN/security service

See [docs/CLOUDFLARE_400_FIX.md](docs/CLOUDFLARE_400_FIX.md) for technical details.

## Recent Major Improvements (v2.0.0)

This SDK has been significantly enhanced based on deep analysis of the official Mistral Python SDK:

✅ **All chat completion parameters** - presence_penalty, frequency_penalty, n, prediction, parallel_tool_calls, prompt_mode, stop sequences, and more
✅ **Pointer-based optional parameters** - Proper nil handling matching Python SDK's OptionalNullable pattern
✅ **Helper functions** - Easy message creation and pointer utilities
✅ **Extended type system** - New types for predictions, prompt modes, and embeddings
✅ **Improved API consistency** - Parameters only sent when explicitly set

See [docs/IMPROVEMENTS.md](docs/IMPROVEMENTS.md) for detailed changes.

## Getting Started

To begin using the Mistral Go Client in your project, ensure you have Go installed on your system. This client library is compatible with Go 1.20 and higher.

### Installation

To install the Mistral Go Client, run the following command:

```bash
go get github.com/ZaguanLabs/mistral-go/v2/sdk
```

### Basic Usage

```go
package main

import (
	"log"
	"github.com/ZaguanLabs/mistral-go/v2/sdk"
)

func main() {
	// Initialize client (loads from MISTRAL_API_KEY env var if empty)
	client := sdk.NewMistralClientDefault("")

	// Simple chat completion
	response, err := client.Chat(
		"mistral-small-latest",
		[]sdk.ChatMessage{
			sdk.UserMessage("Hello, how are you?"),
		},
		nil, // Use nil for default parameters
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.Choices[0].Message.Content)
}
```

### Listing Available Models

Instead of using hardcoded model IDs, you can fetch the current list of available models dynamically:

```go
client := sdk.NewMistralClientDefault("")

// Get all available models
models, err := client.ListModels()
if err != nil {
	log.Fatal(err)
}

// Print all available models
for _, model := range models.Data {
	log.Printf("Model: %s (owned by: %s)\n", model.ID, model.OwnedBy)
}

// Use a model from the list
if len(models.Data) > 0 {
	modelID := models.Data[0].ID
	response, err := client.Chat(
		modelID,
		[]sdk.ChatMessage{
			sdk.UserMessage("Hello!"),
		},
		nil,
	)
	// ...
}
```

**Note:** Model IDs change frequently as Mistral releases new versions. The SDK does not include hardcoded model constants. Always use `ListModels()` to get the current available models, or refer to the [Mistral AI documentation](https://docs.mistral.ai/getting-started/models/) for the latest model IDs.

### Advanced Usage with Parameters

```go
// Create parameters with optional settings
params := sdk.NewChatRequestParams()
params.Temperature = sdk.Float64Ptr(0.7)
params.MaxTokens = sdk.IntPtr(500)
params.PresencePenalty = sdk.Float64Ptr(0.1)
params.FrequencyPenalty = sdk.Float64Ptr(0.1)
params.Stop = []string{"END", "STOP"}

response, err := client.Chat(
	"mistral-small-latest",
	[]sdk.ChatMessage{
		sdk.SystemMessage("You are a helpful assistant."),
		sdk.UserMessage("Write a short poem about Go."),
	},
	params,
)
```

### Streaming Responses

```go
params := sdk.NewChatRequestParams()
params.Temperature = sdk.Float64Ptr(0.7)
params.MaxTokens = sdk.IntPtr(100)

stream, err := client.ChatStream(
	"mistral-small-latest",
	[]sdk.ChatMessage{
		sdk.UserMessage("Tell me a story"),
	},
	params,
)
if err != nil {
	log.Fatal(err)
}

for chunk := range stream {
	if chunk.Error != nil {
		log.Fatal(chunk.Error)
	}
	if len(chunk.Choices) > 0 {
		log.Print(chunk.Choices[0].Delta.Content)
	}
}
```

### Tool/Function Calling

```go
params := sdk.NewChatRequestParams()
params.Tools = []sdk.Tool{
	{
		Type: sdk.ToolTypeFunction,
		Function: sdk.Function{
			Name:        "get_weather",
			Description: "Get the current weather for a location",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"location": map[string]interface{}{
						"type":        "string",
						"description": "City and state, e.g. San Francisco, CA",
					},
				},
				"required": []string{"location"},
			},
		},
	},
}
params.ToolChoice = sdk.ToolChoiceAuto
params.ParallelToolCalls = sdk.BoolPtr(true)

response, err := client.Chat(
	"mistral-small-latest",
	[]sdk.ChatMessage{
		sdk.UserMessage("What's the weather in Paris?"),
	},
	params,
)

// Check for tool calls in response
if len(response.Choices[0].Message.ToolCalls) > 0 {
	toolCall := response.Choices[0].Message.ToolCalls[0]
	log.Printf("Function: %s, Args: %s", 
		toolCall.Function.Name, 
		toolCall.Function.Arguments)
}
```

### Embeddings

```go
response, err := client.Embeddings(
	"mistral-embed",
	[]string{
		"Embed this sentence.",
		"As well as this one.",
	},
)
if err != nil {
	log.Fatal(err)
}

for _, embedding := range response.Data {
	log.Printf("Embedding %d: %d dimensions", 
		embedding.Index, 
		len(embedding.Embedding))
}
```

### Multiple Completions

```go
params := sdk.NewChatRequestParams()
params.N = sdk.IntPtr(3) // Get 3 different completions
params.Temperature = sdk.Float64Ptr(0.8)

response, err := client.Chat(
	"mistral-small-latest",
	[]sdk.ChatMessage{
		sdk.UserMessage("Suggest a name for my startup"),
	},
	params,
)

// response.Choices will contain 3 different suggestions
for i, choice := range response.Choices {
	log.Printf("Suggestion %d: %s", i+1, choice.Message.Content)
}
```

### Reasoning Mode

```go
params := sdk.NewChatRequestParams()
params.PromptMode = sdk.MistralPromptModePtr(sdk.PromptModeReasoning)
params.Temperature = sdk.Float64Ptr(0.3)

response, err := client.Chat(
	"mistral-large-latest",
	[]sdk.ChatMessage{
		sdk.UserMessage("Solve this logic puzzle: ..."),
	},
	params,
)
```

### Files API

```go
// Upload a file
file, err := os.Open("training_data.jsonl")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

uploadResp, err := client.UploadFile(file, "training_data.jsonl", sdk.FilePurposeFineTune)
if err != nil {
	log.Fatal(err)
}
log.Printf("Uploaded file ID: %s\n", uploadResp.ID)

// List files
files, err := client.ListFiles(&sdk.ListFilesParams{
	Purpose: &sdk.FilePurposeFineTune,
})
if err != nil {
	log.Fatal(err)
}

// Download a file
content, err := client.DownloadFile(uploadResp.ID)
if err != nil {
	log.Fatal(err)
}
```

### Fine-tuning API

```go
// Create a fine-tuning job
job, err := client.CreateFineTuningJob(&sdk.CreateFineTuningJobRequest{
	Model: "open-mistral-7b",
	TrainingFiles: []sdk.TrainingFile{
		{FileID: "file-abc123"},
	},
	Hyperparameters: sdk.Hyperparameters{
		TrainingSteps: sdk.IntPtr(10),
		LearningRate:  sdk.Float64Ptr(0.0001),
	},
})
if err != nil {
	log.Fatal(err)
}

// List fine-tuning jobs
jobs, err := client.ListFineTuningJobs(&sdk.ListFineTuningJobsParams{
	Status: &sdk.JobStatusRunning,
})

// Get job status
jobStatus, err := client.GetFineTuningJob(job.ID)

// Cancel a job
cancelled, err := client.CancelFineTuningJob(job.ID)
```

### Batch API

```go
// Create a batch job
batchJob, err := client.CreateBatchJob(&sdk.CreateBatchJobRequest{
	InputFiles: []string{"file-abc123"},
	Endpoint:   sdk.BatchEndpointChat,
	Model:      sdk.StringPtr("mistral-small-latest"),
})
if err != nil {
	log.Fatal(err)
}

// List batch jobs
jobs, err := client.ListBatchJobs(&sdk.ListBatchJobsParams{
	Status: []sdk.BatchJobStatus{sdk.BatchJobStatusRunning},
})

// Get batch job status
status, err := client.GetBatchJob(batchJob.ID)

// Cancel a batch job
cancelled, err := client.CancelBatchJob(batchJob.ID)
```

### Agents API

```go
// Agent completion (non-streaming)
response, err := client.AgentComplete(
	"agent-abc123",
	[]sdk.ChatMessage{
		sdk.UserMessage("What's the weather in Paris?"),
	},
	&sdk.AgentCompletionRequest{
		MaxTokens: sdk.IntPtr(500),
		Tools: []sdk.Tool{
			{
				Type: sdk.ToolTypeFunction,
				Function: sdk.Function{
					Name:        "get_weather",
					Description: "Get weather for a location",
					Parameters: map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"location": map[string]interface{}{
								"type": "string",
							},
						},
					},
				},
			},
		},
	},
)

// Agent completion with streaming
stream, err := client.AgentCompleteStream(
	"agent-abc123",
	[]sdk.ChatMessage{
		sdk.UserMessage("Tell me a story"),
	},
	nil,
)

for response := range stream {
	if response.Error != nil {
		log.Printf("Error: %v", response.Error)
		break
	}
	for _, choice := range response.Choices {
		fmt.Print(choice.Delta.Content)
	}
}
```

### Classifiers API

```go
// Content moderation
response, err := client.ModerateText(
	"mistral-moderation-latest",
	[]string{
		"This is a test message",
		"Another message to moderate",
	},
)

if err != nil {
	log.Fatal(err)
}

for i, result := range response.Results {
	fmt.Printf("Input %d moderation:\n", i)
	for _, category := range result.Categories {
		fmt.Printf("  %s: %.2f\n", category.CategoryName, category.Score)
	}
}
```

### OCR API

```go
// Process document from URL
response, err := client.ProcessOCRFromURL(
	"pixtral-12b-2409",
	"https://example.com/document.pdf",
	&sdk.OCRRequest{
		Pages:              []int{0, 1, 2}, // Process first 3 pages
		IncludeImageBase64: sdk.BoolPtr(true),
		ImageLimit:         sdk.IntPtr(10),
	},
)

if err != nil {
	log.Fatal(err)
}

// Process results
for _, page := range response.Pages {
	fmt.Printf("Page %d:\n%s\n", page.PageNumber, page.Text)
	fmt.Printf("Found %d images\n", len(page.Images))
}

// Process from base64-encoded document
response, err := client.ProcessOCRFromBase64(
	"pixtral-12b-2409",
	base64EncodedDocument,
	nil,
)

// Process from uploaded file
response, err := client.ProcessOCRFromFileID(
	"pixtral-12b-2409",
	"file-abc123",
	nil,
)
```

### Audio/Transcriptions API

```go
// Transcribe audio file
file, err := os.Open("audio.mp3")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

response, err := client.Transcribe(
	"whisper-large-v3",
	file,
	"audio.mp3",
	&sdk.TranscriptionRequest{
		Language: sdk.StringPtr("en"),
		TimestampGranularities: []sdk.TimestampGranularity{
			sdk.TimestampGranularityWord,
			sdk.TimestampGranularitySegment,
		},
	},
)

if err != nil {
	log.Fatal(err)
}

fmt.Printf("Transcription: %s\n", response.Text)

// Process word-level timestamps
for _, word := range response.Words {
	fmt.Printf("[%.2f-%.2f] %s\n", word.Start, word.End, word.Word)
}

// Transcribe from URL
response, err := client.TranscribeFromURL(
	"whisper-large-v3",
	"https://example.com/audio.mp3",
	nil,
)

// Transcribe from uploaded file
response, err := client.TranscribeFromFileID(
	"whisper-large-v3",
	"file-abc123",
	nil,
)
```

### Models API

```go
// List all models
models, err := client.ListModels()
for _, model := range models.Data {
	fmt.Printf("%s: %s\n", model.ID, model.OwnedBy)
}

// Retrieve specific model details
model, err := client.RetrieveModel("mistral-small-latest")
fmt.Printf("Model: %s (created: %d)\n", model.ID, model.Created)

// Update a fine-tuned model
updated, err := client.UpdateModel(
	"ft:open-mistral-7b:abc123",
	&sdk.UpdateModelRequest{
		Name:        sdk.StringPtr("My Custom Model"),
		Description: sdk.StringPtr("Fine-tuned for customer support"),
	},
)

// Archive a fine-tuned model
archived, err := client.ArchiveModel("ft:open-mistral-7b:abc123")
fmt.Printf("Model archived: %v\n", archived.Archived)

// Unarchive a model
unarchived, err := client.UnarchiveModel("ft:open-mistral-7b:abc123")

// Delete a fine-tuned model
deleted, err := client.DeleteModel("ft:open-mistral-7b:abc123")
fmt.Printf("Model deleted: %v\n", deleted.Deleted)
```

### Enhanced Embeddings API

```go
// Simple embeddings
embeddings, err := client.Embeddings(
	"mistral-embed",
	[]string{"Hello world", "Goodbye world"},
)

// Advanced embeddings with parameters
embeddings, err := client.EmbeddingsWithParams(
	"mistral-embed",
	[]string{"Text to embed"},
	&sdk.EmbeddingRequest{
		EncodingFormat:  &sdk.EncodingFormatFloat,
		OutputDimension: sdk.IntPtr(512),
		OutputDtype:     &sdk.EmbeddingDtypeFloat32,
	},
)

for _, emb := range embeddings.Data {
	fmt.Printf("Embedding %d: %d dimensions\n", emb.Index, len(emb.Embedding))
}
```

### Enhanced FIM API

```go
// Simple FIM completion
response, err := client.FIM(&sdk.FIMRequestParams{
	Model:  "codestral-latest",
	Prompt: "def fibonacci(n):\n    ",
	Suffix: sdk.StringPtr("\n    return result"),
})

// FIM with streaming
stream, err := client.FIMStream(&sdk.FIMRequestParams{
	Model:       "codestral-latest",
	Prompt:      "def add(a, b):\n    ",
	Suffix:      sdk.StringPtr("\n    return result"),
	MaxTokens:   sdk.IntPtr(100),
	Temperature: sdk.Float64Ptr(0.1),
})

for chunk := range stream {
	if chunk.Error != nil {
		log.Printf("Error: %v", chunk.Error)
		break
	}
	for _, choice := range chunk.Choices {
		fmt.Print(choice.Delta.Content)
	}
}
```

## Documentation

### API Documentation
For detailed documentation on the Mistral AI API and the available endpoints, please refer to the [Mistral AI API Documentation](https://docs.mistral.ai).

### SDK Documentation
Comprehensive guides and documentation are available in the [`docs/`](docs/) directory:

- **[Feature Comparison](docs/FEATURE_COMPARISON.md)** - Complete API parity analysis with Python SDK
- **[Beta Features Guide](docs/BETA_FEATURES_IMPLEMENTATION.md)** - Conversations, Libraries, Documents, Accesses, Mistral Agents
- **[Improvements](docs/IMPROVEMENTS.md)** - Zero dependencies achievement and enhancements
- **[OCR & Audio APIs](docs/OCR_AUDIO_IMPLEMENTATION.md)** - Document processing and transcription
- **[Models API](docs/MODELS_API_COMPLETION.md)** - Complete model management
- **[Embeddings & FIM](docs/EMBEDDINGS_FIM_COMPLETION.md)** - Enhanced embeddings and code completion
- **[Security Audit](docs/AUDIT_REPORT.md)** - Comprehensive security analysis (A- grade, 0 vulnerabilities)
- **[Test Coverage](docs/TEST_COVERAGE_REPORT.md)** - Testing analysis (65.2% coverage)

### Version Information
```go
import "github.com/ZaguanLabs/mistral-go/v2/sdk"

info := sdk.GetVersionInfo()
fmt.Printf("SDK: %s v%s\n", info.SDKName, info.Version)
fmt.Printf("Feature Parity: %s\n", info.FeatureParity) // "100%"
```

## Contributing

Contributions are welcome! If you would like to contribute to the project, please fork the repository and submit a pull request with your changes.

## License

The Mistral Go Client is open-sourced software licensed under the [MIT license](LICENSE).

## Acknowledgments

**Huge thanks to [Gage Technologies](https://github.com/Gage-Technologies) for creating the initial version of this SDK!** 🙏

This project builds upon their excellent foundation from the [original repository](https://github.com/Gage-Technologies/mistral-go). We're grateful for their pioneering work in bringing Mistral AI capabilities to the Go ecosystem.

## Support

If you encounter any issues or require assistance, please file an issue on the GitHub repository issue tracker.
