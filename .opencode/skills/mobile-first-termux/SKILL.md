---
name: mobile-first-termux
description: Mobile-first development targeting Android Termux environment with constraints on screen, memory, and CPU
license: MIT
compatibility: opencode
metadata:
  audience: developers
  platform: termux
---

## Termux Environment
- Shell: bash (default), zsh available via pkg
- Filesystem: Android app sandbox (no root, no `/usr`)
- Network: localhost networking, no privileged ports
- Display: 80x24 minimum terminal size, 256 colors minimum
- Arch: aarch64 (ARM64) primary, arm32 secondary
- Memory: 512MB-4GB RAM typical, 100-200MB per process limit

## Build Constraints
- `CGO_ENABLED=0` preferred for cross-compatibility
- Static binaries: `-ldflags="-s -w -extldflags=-static"`
- Platform detection: runtime.GOOS/GOARCH for conditional code
- Avoid mmap-heavy operations (Android kernel limits)

## UI Constraints
- Minimum terminal width: 80 columns
- Minimum terminal height: 24 rows
- No mouse dependency (optional enhancement)
- No multimedia or GUI dependencies
- Keyboard: basic ASCII input, no function keys F13-F24
- Color: 256 color palette, handle TERM=xterm-256color and TERM=dumb

## Performance Targets
- Startup time: under 500ms
- Render time: under 16ms per frame
- Memory: under 50MB idle, under 200MB peak
- Binary size: under 15MB uncompressed
- Battery: avoid polling loops, use event-driven patterns

## Testing on Termux
- `pkg install golang` for Go toolchain
- Cross-compile with `GOOS=android GOARCH=arm64 go build`
- Use `CGO_ENABLED=0` for pure Go binaries
- Test interactive apps with `script` command for input simulation
