# Release Notes - v2.0.1

**Release Date**: November 20, 2025

## 🔧 Critical Bug Fix Release

This is a critical patch release that fixes a production-blocking issue where the SDK would fail with Cloudflare 400 errors.

### What's Fixed

#### Critical: Cloudflare 400 Error Resolution

**Problem**: The SDK was missing the `User-Agent` header in HTTP requests, causing Cloudflare-protected endpoints to reject requests with `400 Bad Request` errors.

**Impact**: This affected:
- Direct usage with Mistral API
- Usage through proxies/gateways (e.g., Zaguán)
- Any deployment behind Cloudflare or similar CDN/security services

**Solution**: Added `User-Agent: mistral-go/2.0.1` header to all HTTP requests.

### Technical Details

- **File Changed**: `sdk/client.go` (line 90)
- **Test Coverage**: Added comprehensive tests to prevent regression
  - `TestChatWithMock` - Verifies User-Agent in chat requests
  - `TestUserAgentHeaderWithMock` - Dedicated User-Agent validation
- **Documentation**: Added `docs/CLOUDFLARE_400_FIX.md` with detailed analysis

### Upgrade Instructions

Update your dependency:

```bash
go get github.com/ZaguanLabs/mistral-go/v2/sdk@v2.0.1
```

Or update your `go.mod`:

```go
require github.com/ZaguanLabs/mistral-go/v2 v2.0.1
```

Then run:

```bash
go mod tidy
```

### Breaking Changes

None. This is a backward-compatible bug fix.

### Verification

All tests pass:
```bash
$ go test ./sdk -run "Mock"
ok      github.com/ZaguanLabs/mistral-go/v2/sdk 1.509s
```

### Recommendation

**Immediate upgrade recommended** for all users experiencing 400 errors or deploying to production environments.

### Links

- [Full Changelog](CHANGELOG.md)
- [Technical Details](docs/CLOUDFLARE_400_FIX.md)
- [GitHub Repository](https://github.com/ZaguanLabs/mistral-go)

---

## Previous Release (v2.0.0)

For information about the v2.0.0 major release with 100% feature parity, see the [full CHANGELOG](CHANGELOG.md).
