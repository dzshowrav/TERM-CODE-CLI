# TermCode Project Rules

## Build & Test Commands
- Build: `go build ./...` (CGO_ENABLED=0 for static builds)
- Test: `go test ./...` or `go test -v -count=1 ./...` (no cache)
- Lint: `gofumpt -l -w . && goimports -local -w . && go vet ./...`
- Type check: `go build ./...` (Go compiler enforces types)
- Format: `gofumpt -l -w .` and `goimports -local -w .`

## Project Structure
- Go modules with `cmd/`, `internal/`, `pkg/` layout
- Clean Architecture layers: domain, application, adapters, infrastructure
- Bubble Tea apps follow Model-Update-View pattern

## Core Conventions
- Mobile First: design for 80x24 terminal, 256 colors, aarch64/arm64
- Termux Only: CGO_ENABLED=0, static binaries, no platform-specific deps
- No Mock Data: use real implementations, fakes, or test containers
- No Emoji: text-only output
- Full File Output: write complete files, never partial diffs
- Modular Code: small focused packages, clear interfaces
- Production Ready: graceful shutdown, structured logging, signal handling

## Code Standards
- Go best practices: idiomatic patterns, error wrapping, table-driven tests
- Bubble Tea Architecture: Elm-style model/update/view, cmd/message patterns
- Clean Architecture: dependency inversion, domain isolation

## Custom Agents (Legacy)
- @go-expert - General Go development
- @bubbletea-expert - General Bubble Tea TUI
- @lipgloss-expert - Terminal styling
- @mcp-expert - General MCP server development
- @treesitter-expert - Code parsing and analysis
- @postgres-expert - Database schema and queries
- @redis-expert - Caching and data structures
- @architect - System design and architecture review
- @security-auditor - General security analysis
- @perf-engineer - General performance optimization

## TermCode Multi-Agent System
Orchestration layer (planning/reasoning):
- @master-architect - Architecture decisions, agent coordination
- @task-planner - Request decomposition and execution planning
- @context-engine - Context collection and distribution
- @memory-engine - Long-term project knowledge
- @reasoning-engine - Structured decision making

Engineering agents (implementation):
- @go-engineer - Go business logic, services, domain models
- @bubbletea-engineer - Bubble Tea TUI architecture and screens
- @uiux-engineer - Mobile-first terminal UX and interaction design
- @terminal-engineer - ANSI, Unicode, keyboard, viewport
- @mcp-engineer - MCP integrations and tool execution
- @database-engineer - Schema, migrations, queries, persistence
- @git-engineer - Version control and repository workflow
- @testing-engineer - Quality assurance and automated testing
- @security-engineer - Vulnerability prevention and data protection
- @performance-engineer - Startup, memory, CPU, rendering optimization

Quality assurance layer:
- @documentation-engineer - Documentation and knowledge management
- @review-engineer - Final quality gate and code review
- @release-engineer - Release management and versioning

## Custom Commands
- /feature - Create a new feature
- /component - Create a Bubble Tea component
- /command - Create a custom command
- /mcp - Create an MCP server
- /screen - Create a Bubble Tea screen
- /test - Write tests
- /docs - Write documentation
