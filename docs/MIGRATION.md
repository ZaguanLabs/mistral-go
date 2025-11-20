# Repository Migration Guide

## New Repository Location

The Mistral Go SDK has moved to a new repository:

**Old:** `github.com/gage-technologies/mistral-go`  
**New:** `github.com/ZaguanLabs/mistral-go`

## What Changed

All import paths have been updated throughout the codebase:

### go.mod
```go
module github.com/ZaguanLabs/mistral-go

go 1.20
```

### Import Statements
**Before:**
```go
import "github.com/gage-technologies/mistral-go/sdk"
```

**After:**
```go
import "github.com/ZaguanLabs/mistral-go/sdk"
```

## Migration Steps for Users

If you're using this SDK in your project, update your imports:

### 1. Update go.mod
```bash
go mod edit -replace github.com/gage-technologies/mistral-go=github.com/ZaguanLabs/mistral-go
go get github.com/ZaguanLabs/mistral-go/sdk@latest
go mod tidy
```

### 2. Update Import Statements
Replace all occurrences in your code:
```bash
find . -type f -name "*.go" -exec sed -i 's|github.com/gage-technologies/mistral-go|github.com/ZaguanLabs/mistral-go|g' {} +
```

### 3. Verify
```bash
go build ./...
go test ./...
```

## Installation

New installation command:
```bash
go get github.com/ZaguanLabs/mistral-go/sdk
```

## No Breaking Changes

This is purely a repository move. All functionality remains identical:
- ✅ Same API
- ✅ Same package structure
- ✅ Same features
- ✅ Zero dependencies
- ✅ Same version compatibility

Only the import path has changed.

## Links

- **Repository:** https://github.com/ZaguanLabs/mistral-go
- **Documentation:** See [README.md](README.md)
- **Issues:** https://github.com/ZaguanLabs/mistral-go/issues
