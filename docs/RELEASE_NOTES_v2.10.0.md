# Mistral Go SDK v2.10.0

Updates the Go SDK for the official Python SDK v2.10.0, compared against the previous v2.9.4 baseline.

## API changes

- Adds all 17 new upstream operations: eight service-account operations, five evaluation pipeline configuration operations, two organization connector sharing operations, and scoped credential creation/update.
- Adds HTTP/MCP connector request and response variants, protocol defaults, OAuth client-credentials authentication, and optional OAuth authorization endpoints.
- Adds workflow backend specifications, creator/location filters, execution search-key requests, prompt/skill workspace sharing relations, model billing names, and dataset record source selection.
- Preserves null transcription timestamps and required nullable deployment response fields.
- Keeps removed upstream methods as deprecated Go compatibility methods.

## Runtime and batch helpers

- Typed JSONL input/output, job snapshots, upload/create/get/refresh/cancel/wait, streamed results, downloads, signed URLs, and full batch execution.
- Ordered results report all missing/failed IDs; duplicate IDs are rejected when combining outputs.
- Batch execution preserves caller-owned input files, materializes outputs before cleanup, preserves output files if download fails, and cancels running jobs on abort. Whole-job retry is opt-in and may incur additional charges.
- Configurable retry status codes and jitter, `retry-after-ms` precedence, replayable retry bodies, and response closure before retry.
- Cancellable clients and streams, explicit `EventStream.Close`, multiline SSE data, CR/LF framing, and final events at EOF.
- Base64 conversion for files and readers, restoring seekable stream positions.

## Migration

`TranscriptionSegment.Start` and `End` are now `*float64`. `ExtendedOAuthServerMetadata.AuthorizationEndpoint` is now `*string`. These preserve the nullable upstream fields; dereference only after checking for nil.

`CreateConnector` and `UpdateConnector` accept typed MCP/HTTP requests or JSON objects. An omitted protocol defaults to MCP. Existing request structs remain usable.

Use `client.WithContext(ctx)` and cancel the context when abandoning a channel stream. Use `WithRetryConfig` to obtain an independently configured client.

## Verification

The official archive and PyPI source distribution are stored locally under `docs/client-python/`. The Python package was built with `python -m build`, and its generated source archive was unpacked. The published source SHA-256 was verified; all 1,208 Python files match the release checkout, including its Azure/GCP package directories.

Validation: `zsh -f scripts/test-offline.zsh`, `go build ./...`, `go vet ./...`, and `git diff --check`.

The local suite uses mock HTTP/WebSocket servers and Go's race detector. Live API integration tests were not run. These checks verify the implemented update; they do not independently certify every historical SDK behavior against the live service.
