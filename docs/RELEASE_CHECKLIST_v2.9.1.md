# v2.9.1 Release Checklist

## Code and parity

- [x] Official Python SDK v2.9.1 source downloaded, built, and unpacked under `docs/client-python/`.
- [x] Python SDK v2.4.13-to-v2.9.1 public API delta reviewed.
- [x] v2 prompt and skill registries implemented.
- [x] User identity and paginated beta-agent APIs implemented.
- [x] Workflow deployment, execution trace, and expanded filtering APIs implemented.
- [x] Observability aggregation and deployment-oriented RAG APIs implemented.
- [x] Connector, OCR, library, speech, and model parity fields implemented.
- [x] SSE stream errors represented by `StreamDisconnectedError`.
- [x] Version and User-Agent updated to v2.9.1.

## Validation

- [x] Go formatting passes.
- [x] Local mock-server parity tests pass.
- [x] All packages compile.
- [x] `go vet ./...` passes.
- [x] `git diff --check` passes.

## Documentation

- [x] README compatibility banner and summary updated.
- [x] CHANGELOG updated.
- [x] Release notes prepared.
- [x] Official source and build artifacts retained under the ignored documentation reference tree.

## Release

- [x] Changes committed and pushed to `main`.
- [x] Annotated `v2.9.1` tag pushed.
- [x] GitHub release published from `docs/RELEASE_NOTES_v2.9.1.md`.
