# v2.9.4 Release Checklist

- [x] Official Python SDK v2.9.4 source downloaded, built, and unpacked.
- [x] Python v2.9.3-to-v2.9.4 model delta reviewed.
- [x] New request and response fields represented in Go.
- [x] SDK version, changelog, and README updated.
- [x] Focused local parity tests pass.
- [x] `go build ./...`
- [x] `go vet ./...`
- [x] `git diff --check`

The repository's legacy integration tests require a valid live Mistral API key and are not part of the local release gate.
