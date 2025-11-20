# Cloudflare 400 Error Fix

## Problem

The SDK was returning HTTP 400 errors when used with Cloudflare-protected endpoints (including the official Mistral API). The error manifested as:

```
(HTTP Error 400) <html>
<head><title>400 Bad Request</title></head>
<body>
<center><h1>400 Bad Request</h1></center>
<hr><center>cloudflare</center>
</body>
</html>
```

## Root Cause

The SDK was not sending a `User-Agent` header with HTTP requests. Many APIs behind Cloudflare require a User-Agent header, and Cloudflare will reject requests without one with a 400 Bad Request error.

### Original Code (client.go)

```go
req.Header.Set("Authorization", "Bearer "+c.apiKey)
req.Header.Set("Content-Type", "application/json")
// Missing User-Agent header!
```

## Solution

Added the `User-Agent` header to all HTTP requests:

```go
req.Header.Set("Authorization", "Bearer "+c.apiKey)
req.Header.Set("Content-Type", "application/json")
req.Header.Set("User-Agent", UserAgent) // UserAgent = "mistral-go/2.0.0"
```

## Changes Made

1. **client.go**: Added `User-Agent` header to the `request()` method (line 90)
2. **integration_test.go**: 
   - Added User-Agent verification to `TestChatWithMock` (lines 44-51)
   - Added dedicated `TestUserAgentHeaderWithMock` test (lines 446-469)
3. **CHANGELOG.md**: Documented the fix

## Testing

All mock tests pass, verifying that:
- The User-Agent header is present in all requests
- The header value is correct: `mistral-go/2.0.0`
- The fix doesn't break existing functionality

```bash
$ go test ./sdk -run "Mock"
ok      github.com/ZaguanLabs/mistral-go/v2/sdk 1.509s
```

## Impact

This fix resolves the issue for:
- Direct SDK usage with Mistral API
- SDK usage through proxies/gateways (like Zaguán)
- Any deployment behind Cloudflare or similar CDN/security services

## Version

This fix will be included in the next release (2.0.1 or 2.1.0).
