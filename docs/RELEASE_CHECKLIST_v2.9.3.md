# v2.9.3 Release Checklist

## Code and parity

- [x] Official Python SDK v2.9.3 source downloaded, built, and unpacked under `docs/client-python/`.
- [x] Python SDK v2.9.1-to-v2.9.3 public API delta reviewed.
- [x] Realtime sessions and user organization/workspace operations implemented.
- [x] Service-tier, connector authentication, ingestion, and workflow additions implemented.
- [x] Python pagination and ingestion type defaults mirrored.
- [x] Version and User-Agent updated to v2.9.3.

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
- [x] Annotated `v2.9.3` tag pushed.
- [x] GitHub release published from `docs/RELEASE_NOTES_v2.9.3.md`.
