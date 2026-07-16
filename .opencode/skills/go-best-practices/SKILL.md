---
name: go-best-practices
description: Go development conventions including idiomatic patterns, project layout, error handling, testing, and production readiness
license: MIT
compatibility: opencode
metadata:
  audience: developers
  language: go
---

## Project Layout
- `cmd/` - application entry points (thin main packages)
- `internal/` - private application code (not importable)
- `pkg/` - shared library code
- `config/` - configuration structs and defaults
- Use `go mod` for dependency management, pin dependencies

## Code Style
- Run `gofumpt -l -w .` and `goimports -local` before committing
- No global state, no `init()` side effects, no panics
- Error wrapping: `fmt.Errorf("context: %w", err)`
- Interface design: small (1-3 methods), focused, accept interfaces return structs
- Constructor idiom: `func New(deps ...) (*Type, error)` with dependency injection

## Concurrency
- Use `errgroup` for goroutine lifecycle management
- `sync.Once` for lazy initialization, `sync.Pool` for reusable buffers
- Channel ownership: producer closes, consumer reads
- Context propagation: first parameter, pass through all layers

## Testing
- Table-driven tests with sub-tests `t.Run`
- Use `testing/slogtest` for logger, `httptest` for HTTP, `iotest` for I/O
- Fakes over mocks: implement the interface with controlled behavior
- Integration tests with `//go:build integration` build tag
- Run `go vet ./...` and `staticcheck ./...` in CI

## Production Readiness
- Structured logging via `log/slog`
- `net/http/pprof` for profiling endpoints (conditional on build tag)
- Graceful shutdown: `signal.NotifyContext` + `Shutdown()`
- Metrics: `expvar` or Prometheus `promhttp`
- Configuration: env vars with `os.LookupEnv`, sensible defaults
