# v2.10.1 Release Checklist

- [x] Official PyPI sdist downloaded and SHA-256 verified.
- [x] `python -m build` completed and resulting sdist unpacked.
- [x] Python 2.10.0-to-2.10.1 source delta reviewed.
- [x] Managed-index endpoints and models added; changed request shapes ported.
- [x] Endpoint, pagination, schema, filter, trace, and migration regression coverage.
- [x] `zsh -f scripts/test-offline.zsh` passes with race detector.
- [x] `go build ./...`, `go vet ./...`, and `git diff --check` pass.
- [x] Version, README, changelog, and release notes updated.
