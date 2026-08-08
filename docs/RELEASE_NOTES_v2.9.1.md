# Release Notes - v2.9.1

Mistral Go SDK v2.9.1 updates the SDK for compatibility with the official Mistral Python SDK v2.9.1.

## Highlights

- v2 prompt and skill registries, with version and metadata operations
- User identity and paginated beta-agent resources
- Workflow deployment lifecycle, workers, execution trace information, and expanded filters
- Observability span and trace aggregation
- Deployment-oriented RAG registration and metrics APIs
- Expanded connector consumer, sharing, and credential operations
- OCR block output types and new OCR request controls
- Frame-aware SSE parsing with typed stream-disconnection errors

## Additional parity updates

- Library pagination tokens and filters
- Speech prompt-cache keys
- Model internal and unified-resource metadata
- Workflow force-new-trace, tags, and search fields

## Validation

The release is validated with formatting, local mock-server parity tests, complete package compilation, and `go vet`. The repository's legacy live API tests still require valid Mistral credentials and deterministic remote model output.
