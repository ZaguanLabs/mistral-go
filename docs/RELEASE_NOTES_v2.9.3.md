# Release Notes - v2.9.3

Mistral Go SDK v2.9.3 updates the SDK for compatibility with the official Mistral Python SDK v2.9.3.

## Highlights

- Realtime client sessions via `CreateRealtimeSession`
- User organization and workspace list operations
- Chat and agent `service_tier` request controls and usage reporting
- Connector global headers and typed authentication-method updates
- Ingestion pipeline target-index references with Python's `vespa` default
- Workflow deployment scoping, event linkage, execution attempts, and Git commit metadata

## Compatibility

Existing connector header/auth-data fields and the original one-argument workflow deployment
lookup remain available for source compatibility. New typed fields produce the Python 2.9.3
wire shapes.

## Validation

Validation uses formatting, deterministic local mock-server parity tests, complete package
compilation, and `go vet`. The repository's legacy live API tests require valid Mistral
credentials and deterministic remote model output.
