# v2.10.0 Release Checklist

- [x] Official GitHub release archive and PyPI sdist downloaded.
- [x] PyPI SHA-256 verified and published Python source matched to the release checkout.
- [x] `python -m build` completed; generated sdist unpacked.
- [x] Python v2.9.4-to-v2.10.0 endpoint, model, batch-helper, retry, and stream changes reviewed.
- [x] New endpoints and changed request/response fields ported.
- [x] Local regression tests cover endpoints, variants, nulls, pagination, batch ownership/cancellation, retries, and streaming.
- [x] SDK version, README, changelog, and migration notes updated.
- [x] `zsh -f scripts/test-offline.zsh` (race detector enabled).
- [x] `go build ./...` and `go vet ./...`.
- [x] `git diff --check`.

Live credential-dependent integration tests are excluded from the local release gate. The test runner names each excluded legacy test explicitly, so newly added tests run by default.
