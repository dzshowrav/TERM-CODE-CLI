# TERM CODE (tc) - Complete Architecture Plan

## Version 1.0
## AI Coding Agent CLI for Android Termux

---

# TABLE OF CONTENTS

1. [Executive Summary](#1-executive-summary)
2. [Project Identity & Philosophy](#2-project-identity--philosophy)
3. [Tech Stack & Dependencies](#3-tech-stack--dependencies)
4. [Project Structure](#4-project-structure)
5. [Clean Architecture Layers](#5-clean-architecture-layers)
6. [Core Systems Architecture](#6-core-systems-architecture)
7. [AI Model & Provider Management System](#7-ai-model--provider-management-system)
8. [Model Router Architecture](#8-model-router-architecture)
9. [LLM Provider Manager](#9-llm-provider-manager)
10. [Stream Handler & Token Management](#10-stream-handler--token-management)
11. [TUI Shell Architecture](#11-tui-shell-architecture)
12. [UI-UX Mobile-First Design](#12-ui-ux-mobile-first-design)
13. [Key Screens & Components](#13-key-screens--components)
14. [State Management](#14-state-management)
15. [Database Schema](#15-database-schema)
16. [Event Bus System](#16-event-bus-system)
17. [Configuration System](#17-configuration-system)
18. [Security & Permissions](#18-security--permissions)
19. [Development Workflow](#19-development-workflow)
20. [Implementation Roadmap](#20-implementation-roadmap)

---

# 1. EXECUTIVE SUMMARY

## What is TERM CODE?

TERM CODE (short form: `tc`) is a 100% Termux-compatible, mobile-first, mobile-friendly AI coding agent CLI application. It combines the best UX patterns from **opencode CLI** into a single, unified terminal interface that works entirely on Android Termux.

## Core Purpose

Provide a complete AI-assisted coding experience directly in the terminal:
- Chat with AI models (OpenAI, Anthropic, OpenRouter, local models, etc.)
- Read, write, edit, and search files
- Execute shell commands with permission control
- Manage sessions, providers, models, and tools
- Full Git integration
- MCP (Model Context Protocol) support
- Plugin system for extensibility

## Key Differentiators

- **100% Mobile First** - Designed for portrait mode, thumb reach, and on-screen keyboard
- **No Sidebars** - All navigation via command input, slash commands, and modal overlays
- **Zero-Config Start** - Just run `tc` and everything is already configured; no commands needed to start
- **Home Screen First** - When no conversation exists, show config overview (provider, model, agent, workspace) so user knows current state before typing
- **Model Header on Chat** - Active model name shown at top of conversation area when chat is active
- **Status Bar Below Input** - Status bar is at the very bottom, below the input area
- **AI Provider Agnostic** - Any OpenAI-compatible provider works
- **Local Model Support** - Ollama, Hugging Face, and custom local models
- **Offline First** - Core features work without internet
- **No Emoji** - Text-only output using terminal-compatible icons

---

# 2. PROJECT IDENTITY & PHILOSOPHY

## Name & Branding

- **Full Name**: TERM CODE
- **Short Name**: tc
- **Binary Name**: `tc`
- **Config Directory**: `~/.config/tc/`
- **Data Directory**: `~/.local/share/tc/`
- **Database**: `~/.local/share/tc/tc.db`

## Design Philosophy

```
Mobile First     - Everything designed for a phone first
Terminal Native  - Feels like a real CLI, not a GUI
AI First         - AI conversation is primary content
Content First    - Information over decoration
Keyboard Safe    - Keyboard never hides important content
Bottom Focused   - Primary actions near the bottom
Touch Friendly   - Minimum 44px touch targets
Minimal          - Only what is necessary
Fast             - Under 500ms startup, 60fps rendering
Offline First    - Core features work without network
```

## Core Rules

1. No sidebars - ever
2. No emoji - use terminal icons only (Nerd Font, Codicons, Powerline)
3. Status bar always below input area
4. Input always visible (keyboard open or closed)
5. 80x24 minimum terminal support
6. 256 color support minimum
7. Portrait orientation primary
8. Single-column layout only
9. CGO_ENABLED=0 for static builds
10. Full file output, never partial diffs

---

# 3. TECH STACK & DEPENDENCIES

## Go Runtime

- **Go**: 1.26.3+
- **Build**: `CGO_ENABLED=0 go build -ldflags="-s -w -extldflags=-static" ./cmd/tc`
- **Platform**: `linux/arm64` (primary), `linux/amd64` (secondary)

## Core TUI Framework

| Package | Purpose | Version |
|---------|---------|---------|
| `charm.land/bubbletea/v2` | TUI framework (Elm Architecture) | v2.0.8 |
| `github.com/charmbracelet/bubbles` | Reusable TUI components (textinput, viewport, spinner, etc.) | v1.0.0 |
| `github.com/charmbracelet/lipgloss` | Terminal styling (colors, layout, borders) | v1.1.1 |
| `github.com/charmbracelet/glamour` | Markdown rendering in terminal | v1.0.0 |
| `github.com/charmbracelet/harmonica` | Smooth animations for window size | v0.2.0 |
| `github.com/charmbracelet/huh` | Form prompts | v1.0.0 |
| `github.com/charmbracelet/log` | Structured logging (charm style) | v1.0.0 |
| `github.com/muesli/termenv` | Terminal color detection and profiles | v0.16.0 |
| `github.com/muesli/reflow` | Text wrapping, word wrapping, ANSI-aware | v0.3.0 |
| `github.com/muesli/ansi` | ANSI sequence utilities | - |
| `github.com/mattn/go-runewidth` | Character width detection | v0.0.24 |
| `github.com/rivo/uniseg` | Unicode text segmentation | v0.4.7 |
| `github.com/mattn/go-isatty` | TTY detection | v0.0.23 |

## CLI Framework

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/spf13/cobra` | CLI commands & subcommands | v1.10.2 |
| `github.com/spf13/pflag` | POSIX-style flag parsing | v1.0.10 |
| `github.com/spf13/viper` | Configuration management | v1.21.0 |
| `github.com/knadh/koanf` | Alternative config with providers | v1.5.0 |

## AI / MCP

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/modelcontextprotocol/go-sdk` | MCP protocol implementation | v1.6.1 |
| `github.com/go-resty/resty/v2` | HTTP client for API calls | v2.17.2 |
| `github.com/gorilla/websocket` | WebSocket for streaming | v1.5.3 |

## Database

| Package | Purpose | Version |
|---------|---------|---------|
| `modernc.org/sqlite` | SQLite driver (pure Go, no CGO) | v1.54.0 |
| `github.com/jackc/pgx/v5` | PostgreSQL driver (optional) | v5.10.0 |
| `github.com/redis/go-redis/v9` | Redis client (optional caching) | v9.21.0 |

## File System & Git

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/fsnotify/fsnotify` | File system notifications | v1.10.1 |
| `github.com/go-git/go-git/v5` | Pure Go Git operations | v5.19.1 |
| `github.com/shirou/gopsutil/v4` | System metrics (CPU, memory) | v4.26.6 |

## Code Analysis

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/tree-sitter/go-tree-sitter` | Incremental parsing (syntax trees) | v0.25.0 |
| `github.com/yuin/goldmark` | Markdown parser | v1.8.4 |
| `github.com/alecthomas/chroma/v2` | Syntax highlighting | v2.20.0 |

## Search & Fuzzy

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/lithammer/fuzzysearch` | Fuzzy string matching | v1.1.8 |
| `github.com/sahilm/fuzzy` | Fuzzy matching (filenames) | v0.1.3 |

## Diff

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/hexops/gotextdiff` | Text diffs (used by gopls) | v1.0.3 |
| `github.com/sergi/go-diff` | Diff match patch | v1.4.0 |

## Table / Progress Rendering

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/jedib0t/go-pretty/v6` | Table rendering | v6.8.2 |
| `github.com/olekukonko/tablewriter` | ASCII table writer | v1.1.4 |
| `github.com/vbauerster/mpb/v8` | Multi-progress bars | v8.12.1 |
| `github.com/schollz/progressbar/v3` | Single progress bar | v3.19.1 |
| `github.com/briandowns/spinner` | Terminal spinner | v1.23.2 |

## JSON / Config

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/tidwall/gjson` | Fast JSON path queries | v1.19.0 |
| `github.com/tidwall/sjson` | Fast JSON setting | v1.2.5 |
| `github.com/pelletier/go-toml/v2` | TOML parsing | v2.4.3 |
| `gopkg.in/yaml.v3` | YAML parsing | v3.0.1 |
| `github.com/joho/godotenv` | .env file loading | v1.5.1 |

## Concurrency

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/avast/retry-go/v4` | Retry with backoff | v4.7.0 |
| `github.com/panjf2000/ants/v2` | Goroutine pool | v2.12.1 |
| `github.com/patrickmn/go-cache` | In-memory cache | v2.1.0 |

## Validation

| Package | Purpose | Version |
|---------|---------|---------|
| `github.com/go-playground/validator/v10` | Struct validation | v10.30.3 |
| `github.com/google/uuid` | UUID generation | v1.6.0 |

## External Tools (executed as subprocess)

| Tool | Purpose | Installation |
|------|---------|-------------|
| `ripgrep` (rg) | Ultra-fast code search | `pkg install ripgrep` |
| `fd` | Fast file find | `pkg install fd` |
| `ollama` | Local model serving | `pkg install ollama` |

---

# 4. PROJECT STRUCTURE

```
termcode/
├── cmd/
│   └── tc/
│       └── main.go                    # Application entry point
│
├── internal/
│   ├── domain/                        # Enterprise business rules
│   │   ├── provider/
│   │   │   ├── provider.go            # Provider entity & value objects
│   │   │   └── errors.go              # Provider-specific errors
│   │   ├── model/
│   │   │   ├── model.go               # Model entity & value objects
│   │   │   ├── capabilities.go        # Capabilities value object
│   │   │   └── errors.go              # Model-specific errors
│   │   ├── session/
│   │   │   ├── session.go             # Session entity
│   │   │   ├── message.go             # Message value object
│   │   │   └── errors.go              # Session-specific errors
│   │   ├── conversation/
│   │   │   └── conversation.go        # Conversation aggregate
│   │   ├── tool/
│   │   │   ├── tool.go                # Tool entity
│   │   │   └── result.go              # Tool result value object
│   │   ├── workspace/
│   │   │   └── workspace.go           # Workspace entity
│   │   ├── agent/
│   │   │   └── agent.go               # Agent entity
│   │   ├── skill/
│   │   │   └── skill.go               # Skill entity
│   │   ├── token/
│   │   │   └── token.go               # Token tracking value objects
│   │   ├── mcp/
│   │   │   └── server.go              # MCP server entity
│   │   ├── plugin/
│   │   │   └── plugin.go              # Plugin entity
│   │   ├── theme/
│   │   │   └── theme.go               # Theme entity
│   │   ├── config/
│   │   │   └── config.go              # Configuration aggregates
│   │   └── permission/
│   │       └── permission.go          # Permission policy entity
│   │
│   ├── application/                   # Use case / application services
│   │   ├── provider/
│   │   │   ├── provider_service.go    # Provider CRUD use cases
│   │   │   ├── provider_test.go       # Provider connection test
│   │   │   └── provider_discovery.go  # Auto-discover models from provider
│   │   ├── model/
│   │   │   ├── model_service.go       # Model CRUD use cases
│   │   │   ├── model_selector.go      # Model selection logic
│   │   │   └── local_model.go         # Local model management
│   │   ├── session/
│   │   │   ├── session_service.go     # Session management
│   │   │   └── session_export.go      # Session import/export
│   │   ├── conversation/
│   │   │   ├── chat_service.go        # Chat/request use case
│   │   │   └── context_builder.go     # Context composition
│   │   ├── tool/
│   │   │   ├── tool_service.go        # Tool execution orchestration
│   │   │   └── permission_service.go  # Permission checks
│   │   ├── workspace/
│   │   │   └── workspace_service.go   # Workspace management
│   │   ├── router/
│   │   │   ├── model_router.go        # Model routing logic
│   │   │   ├── capability_matcher.go  # Capability matching
│   │   │   ├── cost_engine.go         # Cost calculation
│   │   │   └── failover_manager.go    # Provider failover
│   │   ├── stream/
│   │   │   ├── stream_handler.go      # Stream processing
│   │   │   ├── token_tracker.go       # Token usage tracking
│   │   │   └── stream_buffer.go       # Stream buffering
│   │   ├── llm/
│   │   │   ├── provider_manager.go    # Unified provider interface
│   │   │   ├── request_builder.go     # OpenAI-compatible request building
│   │   │   ├── response_parser.go     # Response normalization
│   │   │   └── auth_manager.go        # Authentication handling
│   │   ├── mcp/
│   │   │   └── mcp_service.go         # MCP server management
│   │   ├── file/
│   │   │   ├── read_service.go        # File reading
│   │   │   ├── write_service.go       # File writing
│   │   │   ├── edit_service.go        # File editing
│   │   │   ├── search_service.go      # File search (rg/fd)
│   │   │   └── diff_service.go        # Diff computation
│   │   ├── git/
│   │   │   └── git_service.go         # Git operations
│   │   ├── plugin/
│   │   │   └── plugin_service.go      # Plugin loading/unloading
│   │   ├── skill/
│   │   │   └── skill_service.go       # Skill management
│   │   ├── agent/
│   │   │   └── agent_service.go       # Agent management
│   │   ├── cache/
│   │   │   └── cache_service.go       # Caching logic
│   │   └── backup/
│   │       └── backup_service.go      # Backup & restore
│   │
│   ├── adapters/                      # Interface adapters
│   │   ├── cli/
│   │   │   ├── root.go                # Root cobra command
│   │   │   ├── serve.go               # `tc serve` - start TUI
│   │   │   ├── provider_cmd.go        # `tc provider` subcommands
│   │   │   ├── model_cmd.go           # `tc model` subcommands
│   │   │   ├── session_cmd.go         # `tc session` subcommands
│   │   │   ├── config_cmd.go          # `tc config` subcommands
│   │   │   ├── backup_cmd.go          # `tc backup` subcommands
│   │   │   └── version_cmd.go         # `tc version`
│   │   │
│   │   ├── tui/
│   │   │   ├── app.go                 # Main TUI application model
│   │   │   ├── app_update.go          # Main update loop
│   │   │   ├── app_view.go            # Main view renderer
│   │   │   │
│   │   │   ├── screens/
│   │   │   │   ├── chat_screen.go          # Primary chat screen
│   │   │   │   ├── provider_list_screen.go  # Provider list
│   │   │   │   ├── provider_add_screen.go   # Add provider form
│   │   │   │   ├── model_list_screen.go     # Model list (/all models)
│   │   │   │   ├── model_add_screen.go      # Add model form
│   │   │   │   ├── model_selector.go        # Model selector overlay
│   │   │   │   ├── session_screen.go        # Session manager
│   │   │   │   ├── settings_screen.go       # Settings
│   │   │   │   ├── help_screen.go           # Help
│   │   │   │   └── tool_execution_screen.go # Tool execution view
│   │   │   │
│   │   │   ├── components/
│   │   │   │   ├── status_bar.go            # Bottom status bar
│   │   │   │   ├── command_input.go          # Input area
│   │   │   │   ├── message_list.go           # Message list (viewport)
│   │   │   │   ├── message_item.go           # Single message renderer
│   │   │   │   ├── code_block.go             # Syntax highlighted code
│   │   │   │   ├── thinking_indicator.go     # AI thinking state
│   │   │   │   ├── streaming_view.go         # Streaming output
│   │   │   │   ├── tool_card.go              # Tool execution card
│   │   │   │   ├── token_gauge.go            # Token usage bar
│   │   │   │   ├── progress_bar.go           # Progress bar component
│   │   │   │   ├── search_input.go            # Search/filter input
│   │   │   │   ├── filter_list.go            # Filterable list
│   │   │   │   ├── help_text.go              # Help text display
│   │   │   │   ├── empty_state.go            # Empty state display
│   │   │   │   ├── error_state.go            # Error state display
│   │   │   │   └── loading_state.go          # Loading state display
│   │   │   │
│   │   │   ├── dialogs/
│   │   │   │   ├── confirm_dialog.go         # Confirmation dialog
│   │   │   │   ├── prompt_dialog.go          # Text input dialog
│   │   │   │   ├── select_dialog.go          # Selection dialog
│   │   │   │   ├── form_dialog.go            # Multi-field form
│   │   │   │   └── notification.go           # Toast notification
│   │   │   │
│   │   │   ├── styles/
│   │   │   │   ├── theme.go                  # Theme definitions
│   │   │   │   ├── colors.go                 # Color palette
│   │   │   │   ├── spacing.go                # Spacing tokens
│   │   │   │   ├── typography.go             # Text styles
│   │   │   │   └── borders.go                # Border styles
│   │   │   │
│   │   │   └── eventbus.go              # UI event bus integration
│   │   │
│   │   ├── handlers/                    # HTTP handlers (future API)
│   │   │   └── ...
│   │   │
│   │   └── presenters/
│   │       ├── message_presenter.go     # Format messages for display
│   │       ├── provider_presenter.go    # Format provider info
│   │       ├── model_presenter.go       # Format model info
│   │       └── token_presenter.go       # Format token stats
│   │
│   └── infrastructure/                 # Frameworks, drivers, external tools
│       ├── database/
│       │   ├── sqlite/
│       │   │   ├── connection.go        # SQLite connection & WAL mode
│       │   │   ├── migrations.go        # Schema migrations
│       │   │   ├── provider_repo.go     # Provider repository
│       │   │   ├── model_repo.go        # Model repository
│       │   │   ├── session_repo.go      # Session repository
│       │   │   ├── message_repo.go      # Message repository
│       │   │   ├── config_repo.go       # Config repository
│       │   │   ├── agent_repo.go        # Agent repository
│       │   │   ├── skill_repo.go        # Skill repository
│       │   │   ├── tool_repo.go         # Tool repository
│       │   │   ├── mcp_repo.go          # MCP server repository
│       │   │   ├── plugin_repo.go       # Plugin repository
│       │   │   ├── theme_repo.go        # Theme repository
│       │   │   ├── permissions_repo.go  # Permissions repository
│       │   │   ├── cache_repo.go        # Cache repository
│       │   │   └── statistics_repo.go   # Statistics repository
│       │   │
│       │   └── redis/
│       │       └── cache.go             # Redis cache implementation
│       │
│       ├── llm/
│       │   ├── openai_adapter.go        # OpenAI-compatible API adapter
│       │   ├── ollama_adapter.go        # Ollama local adapter
│       │   ├── huggingface_adapter.go   # Hugging Face adapter
│       │   ├── streaming.go             # SSE stream reader
│       │   └── tokenizer.go             # Token counting (tiktoken)
│       │
│       ├── executor/
│       │   ├── shell_executor.go        # Shell command execution
│       │   ├── file_executor.go         # File operation executor
│       │   ├── git_executor.go          # Git operation executor
│       │   ├── search_executor.go       # Search executor (rg/fd)
│       │   └── mcp_executor.go          # MCP tool executor
│       │
│       ├── filesystem/
│       │   ├── reader.go                # File reader with encoding detection
│       │   ├── writer.go                # File writer
│       │   ├── editor.go                # File editor (line-based)
│       │   ├── searcher.go              # ripgrep integration (rg)
│       │   ├── finder.go                # fd integration
│       │   ├── watcher.go               # fsnotify watcher
│       │   └── workspace.go             # Workspace scanner
│       │
│       ├── git/
│       │   └── go_git.go               # go-git wrapper
│       │
│       ├── mcp/
│       │   ├── client.go                # MCP client connection
│       │   ├── server.go                # MCP server management
│       │   └── transport.go             # stdio/SSE transport
│       │
│       ├── plugin/
│       │   └── loader.go                # Plugin loading
│       │
│       ├── cache/
│       │   ├── memory.go                # In-memory cache (go-cache)
│       │   └── redis.go                 # Redis cache
│       │
│       ├── config/
│       │   └── viper.go                 # Viper configuration provider
│       │
│       └── logging/
│           └── charm_log.go             # Charm Log setup
│
├── pkg/                               # Shared library code
│   ├── apitypes/
│   │   ├── chat.go                     # Chat completion types
│   │   ├── models.go                   # Models list types
│   │   ├── streaming.go                # SSE event types
│   │   └── errors.go                   # API error types
│   │
│   ├── constants/
│   │   ├── defaults.go                 # Default values
│   │   ├── capabilities.go             # Capability constants
│   │   ├── categories.go               # Model categories
│   │   └── limits.go                   # Size limits
│   │
│   └── helpers/
│       ├── crypto.go                   # API key encryption
│       ├── format.go                   # Text formatting
│       ├── truncate.go                 # String truncation
│       ├── validate.go                 # Validation helpers
│       └── platform.go                 # Platform detection
│
├── docs/                              # Documentation
│   └── ...
│
├── go.mod
├── go.sum
├── ARCHITECTURE.md                     # This file
├── AGENTS.md                           # Multi-agent spec
└── opencode.json                        # opencode configuration
```

---

# 5. CLEAN ARCHITECTURE LAYERS

## Layer Hierarchy

```
┌─────────────────────────────────────────┐
│            UI Layer (TUI)               │
│  screens, components, dialogs, styles   │
├─────────────────────────────────────────┤
│         Command Layer (CLI)             │
│  cobra commands, flag parsing           │
├─────────────────────────────────────────┤
│       Application Layer (Use Cases)     │
│  services, orchestrators, builders      │
├─────────────────────────────────────────┤
│         Domain Layer (Entities)         │
│  provider, model, session, tool, etc.   │
├─────────────────────────────────────────┤
│    Adapters Layer (Interface Adapters)  │
│  presenters, CLI handlers, TUI adapters │
├─────────────────────────────────────────┤
│   Infrastructure Layer (Frameworks)     │
│  database, LLM clients, executor, MCP   │
└─────────────────────────────────────────┘
```

## Dependency Rule

- **Domain** has ZERO external dependencies
- **Application** depends ONLY on domain interfaces
- **Adapters** depend on application (via interfaces)
- **Infrastructure** implements adapter interfaces

---

# 6. CORE SYSTEMS ARCHITECTURE

## System Interaction Diagram

```
┌─────────────┐     ┌──────────────┐     ┌─────────────┐
│   TUI App   │────▶│  Event Bus   │◀────│ CLI Commands│
└──────┬──────┘     └──────┬───────┘     └─────────────┘
       │                    │
       ▼                    ▼
┌──────────────────────────────────────────────┐
│           Application Services                │
│  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ │
│  │ Chat   │ │ Router │ │ Tool   │ │Session │ │
│  │ Service│ │Manager │ │Service │ │Service │ │
│  └───┬────┘ └───┬────┘ └───┬────┘ └───┬────┘ │
│  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ │
│  │Provider│ │ Model  │ │ MCP    │ │ Plugin │ │
│  │Service │ │Service │ │Service │ │Service │ │
│  └───┬────┘ └───┬────┘ └───┬────┘ └───┬────┘ │
└──────┼──────────┼──────────┼──────────┼───────┘
       │          │          │          │
       ▼          ▼          ▼          ▼
┌──────────────────────────────────────────────┐
│            Domain Entities                     │
│  provider, model, session, tool, workspace,   │
│  agent, skill, mcp, plugin, theme, token      │
└──────────────────────────────────────────────┘
       │          │          │          │
       ▼          ▼          ▼          ▼
┌──────────────────────────────────────────────┐
│          Infrastructure Layer                  │
│  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ │
│  │ SQLite │ │ LLM    │ │ MCP    │ │  Git   │ │
│  │  Repos │ │Adapters│ │ Client │ │Executor│ │
│  └────────┘ └────────┘ └────────┘ └────────┘ │
│  ┌────────┐ ┌────────┐ ┌────────┐ ┌────────┐ │
│  │ Shell  │ │ File   │ │ Search │ │ Cache  │ │
│  │Exec.   │ │ Ops    │ │ (rg/fd)│ │        │ │
│  └────────┘ └────────┘ └────────┘ └────────┘ │
└──────────────────────────────────────────────┘
```

---

# 7. AI MODEL & PROVIDER MANAGEMENT SYSTEM

This is the highest-priority subsystem, implementing the complete provider, model, local model, and token management system.

## 7.1 Provider System

### Provider Entity (Domain)

```go
type Provider struct {
    ID          string    `json:"id" validate:"required,uuid"`
    Name        string    `json:"name" validate:"required,max=50"`
    BaseURL     string    `json:"base_url" validate:"required,url"`
    APIKey      string    `json:"api_key,omitempty"` // encrypted at rest
    Description string    `json:"description,omitempty" max:"200"`
    Status      ProviderStatus
    Latency     time.Duration
    Priority    int       `json:"priority" default:"0"`
    IsDefault   bool      `json:"is_default"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type ProviderStatus string
const (
    ProviderStatusConnected    ProviderStatus = "connected"
    ProviderStatusConnecting   ProviderStatus = "connecting"
    ProviderStatusDisconnected ProviderStatus = "disconnected"
    ProviderStatusOffline      ProviderStatus = "offline"
    ProviderStatusAuthFailed   ProviderStatus = "auth_failed"
    ProviderStatusTimeout      ProviderStatus = "timeout"
    ProviderStatusUnknown      ProviderStatus = "unknown"
)
```

### Provider Repository Interface

```go
type ProviderRepository interface {
    Create(ctx context.Context, p *domain.Provider) error
    GetByID(ctx context.Context, id string) (*domain.Provider, error)
    GetByName(ctx context.Context, name string) (*domain.Provider, error)
    List(ctx context.Context) ([]*domain.Provider, error)
    Update(ctx context.Context, p *domain.Provider) error
    Delete(ctx context.Context, id string) error
    SetDefault(ctx context.Context, id string) error
    GetDefault(ctx context.Context) (*domain.Provider, error)
}
```

### Provider Service (Application)

```go
type ProviderService struct {
    repo    ProviderRepository
    encrypt CryptoService // for API key encryption
    http    HTTPClient
    logger  *slog.Logger
}

func (s *ProviderService) Create(ctx context.Context, name, baseURL, apiKey, desc string) (*domain.Provider, error)
func (s *ProviderService) Update(ctx context.Context, id, name, baseURL, apiKey, desc string) (*domain.Provider, error)
func (s *ProviderService) Delete(ctx context.Context, id string) error
func (s *ProviderService) TestConnection(ctx context.Context, id string) (*ConnectionTestResult, error)
func (s *ProviderService) List(ctx context.Context) ([]*domain.Provider, error)
func (s *ProviderService) GetByID(ctx context.Context, id string) (*domain.Provider, error)
func (s *ProviderService) ValidateProviderConfig(name, baseURL string) error
func (s *ProviderService) EncryptAPIKey(key string) (string, error)
func (s *ProviderService) DecryptAPIKey(encrypted string) (string, error)
```

### API Key Security

- Keys encrypted at rest using AES-256-GCM
- Key encryption key derived from machine-specific seed
- Keys NEVER displayed in UI, logs, exports, or error messages
- Keys masked in provider details view: `sk-****5678`
- Environment variable override: `TC_PROVIDER_<NAME>_API_KEY`

## 7.2 Model System

### Model Entity (Domain)

```go
type Model struct {
    ID             string            `json:"id" validate:"required,uuid"`
    ProviderID     string            `json:"provider_id" validate:"required"`
    ModelID        string            `json:"model_id" validate:"required,max=100"` // API model ID
    DisplayName    string            `json:"display_name" validate:"max=100"`
    Description    string            `json:"description,omitempty" max:"500"`
    Category       ModelCategory     `json:"category"`
    Capabilities   ModelCapabilities `json:"capabilities"`
    MaxContext     int               `json:"max_context" default:"4096"`
    MaxOutput      int               `json:"max_output" default:"4096"`
    PricingInput   float64           `json:"pricing_input"`  // per 1K tokens
    PricingOutput  float64           `json:"pricing_output"` // per 1K tokens
    IsLocal        bool              `json:"is_local"`
    IsFavorite     bool              `json:"is_favorite"`
    Enabled        bool              `json:"enabled" default:"true"`
    CreatedAt      time.Time         `json:"created_at"`
    UpdatedAt      time.Time         `json:"updated_at"`
}

type ModelCategory string
const (
    ModelCategoryGeneral     ModelCategory = "general"
    ModelCategoryCoding      ModelCategory = "coding"
    ModelCategoryReasoning   ModelCategory = "reasoning"
    ModelCategoryVision      ModelCategory = "vision"
    ModelCategoryEmbedding   ModelCategory = "embedding"
    ModelCategoryAudio       ModelCategory = "audio"
    ModelCategoryExperimental ModelCategory = "experimental"
    ModelCategoryCustom      ModelCategory = "custom"
)

type ModelCapabilities struct {
    Streaming         bool `json:"streaming"`
    ToolCalling       bool `json:"tool_calling"`
    Reasoning         bool `json:"reasoning"`
    Vision            bool `json:"vision"`
    Embeddings        bool `json:"embeddings"`
    JSONMode          bool `json:"json_mode"`
    FunctionCalling   bool `json:"function_calling"`
    SystemPrompt      bool `json:"system_prompt"`
}
```

### Model Repository Interface

```go
type ModelRepository interface {
    Create(ctx context.Context, m *domain.Model) error
    GetByID(ctx context.Context, id string) (*domain.Model, error)
    List(ctx context.Context) ([]*domain.Model, error)
    ListByProvider(ctx context.Context, providerID string) ([]*domain.Model, error)
    ListByCategory(ctx context.Context, category domain.ModelCategory) ([]*domain.Model, error)
    ListLocal(ctx context.Context) ([]*domain.Model, error)
    ListFavorites(ctx context.Context) ([]*domain.Model, error)
    Search(ctx context.Context, query string) ([]*domain.Model, error)
    Update(ctx context.Context, m *domain.Model) error
    Delete(ctx context.Context, id string) error
    SetFavorite(ctx context.Context, id string, favorite bool) error
    GetActive(ctx context.Context) (*domain.Model, error)
    SetActive(ctx context.Context, id string) error
    GetRecent(ctx context.Context, limit int) ([]*domain.Model, error)
}
```

## 7.3 Local Model Support

### Local Model Discovery

```go
type LocalModelService struct {
    ollamaClient   *OllamaAdapter
    huggingFaceDir string // ~/.cache/huggingface/
    configDir      string // ~/.config/tc/models/local/
}

// Methods
func (s *LocalModelService) DiscoverOllamaModels(ctx context.Context) ([]*domain.Model, error)
func (s *LocalModelService) DiscoverHuggingFaceModels(ctx context.Context) ([]*domain.Model, error)
func (s *LocalModelService) AddLocalManually(name, modelID, providerType string) (*domain.Model, error)
func (s *LocalModelService) ValidateLocalEnvironment(model *domain.Model) (*CompatibilityWarning, error)
func (s *LocalModelService) RemoveLocal(modelID string) error
```

### Compatibility Warnings

```go
type CompatibilityWarning struct {
    ModelName       string
    Warnings        []string
    Unsupporteds    []string
    MissingDeps     []string
    InsufficientRAM bool
    WrongArch       bool // not aarch64
    Severity        WarningSeverity
}

type WarningSeverity string
const (
    SeverityInfo     WarningSeverity = "info"
    SeverityWarning  WarningSeverity = "warning"
    SeverityBlocking WarningSeverity = "blocking" // cannot run
)
```

### Local Model Providers

| Provider | Discovery | API | Auth |
|----------|-----------|-----|------|
| Ollama | `ollama list` via API | OpenAI-compatible | None (local) |
| Hugging Face | Scan `.cache/huggingface/` | Custom TGI | None (local) |
| Manual | User-specified config | OpenAI-compatible | None (local) |
| LM Studio | Local API scan | OpenAI-compatible | None (local) |
| vLLM | Local API scan | OpenAI-compatible | None (local) |

## 7.4 Runtime Router

### Router Flow

```
User Request
    │
    ▼
Model Selected
    │
    ▼
Provider Resolver
    │
    ├── Get Provider by ID
    ├── Get API Key (decrypted)
    ├── Get Base URL
    └── Get Model ID
    │
    ▼
Runtime Router
    │
    ├── Determine: streaming or non-streaming
    ├── Build OpenAI-compatible request
    │   {
    │     "model": "<model_id>",
    │     "messages": [...],
    │     "stream": true/false,
    │     "temperature": 0.7,
    │     "max_tokens": <max_output>
    │   }
    │
    ├── Add Authorization: Bearer <api_key>
    ├── Add Custom Headers (if any)
    │
    ▼
LLM Provider Adapter (OpenAI-Compatible)
    │
    ├── POST /v1/chat/completions
    ├── Handle SSE stream (if streaming)
    ├── Parse response chunks
    └── Return normalized response
    │
    ▼
Stream Handler / Token Tracker
    │
    ▼
Conversation Manager → Renderer
```

## 7.5 Token Management System

### Token Tracker

```go
type TokenUsage struct {
    InputTokens      int     `json:"input_tokens"`
    OutputTokens     int     `json:"output_tokens"`
    TotalTokens      int     `json:"total_tokens"`
    ContextWindow    int     `json:"context_window"`
    RemainingContext int     `json:"remaining_context"`
    ContextPercent   float64 `json:"context_percent"` // 0.0 - 1.0
    EstimatedCost    float64 `json:"estimated_cost"`
    RequestCount     int     `json:"request_count"`
}

type SessionTokenTotals struct {
    TotalInputTokens  int     `json:"total_input"`
    TotalOutputTokens int     `json:"total_output"`
    TotalCost         float64 `json:"total_cost"`
    RequestCount      int     `json:"request_count"`
    SessionDuration   time.Duration
}

type TokenTrackerService struct {
    current SessionTokenUsage
    session SessionTokenTotals
}

// Methods
func (s *TokenTrackerService) TrackRequest(usage *TokenUsage)
func (s *TokenTrackerService) GetCurrentUsage() *TokenUsage
func (s *TokenTrackerService) GetSessionTotals() *SessionTokenTotals
func (s *TokenTrackerService) GetContextWarning() *ContextWarning
func (s *TokenTrackerService) ResetSession()
```

### Context Warning Thresholds

| Level | Threshold | Action |
|-------|-----------|--------|
| Normal | < 70% | No action |
| Warning | 70-85% | Show warning in status bar |
| Critical | 85-95% | Suggest compaction |
| Full | > 95% | Block new requests until compaction |

### Cost Calculation

```go
func CalculateCost(inputTokens, outputTokens int, pricingInput, pricingOutput float64) float64 {
    inputCost := (float64(inputTokens) / 1000.0) * pricingInput
    outputCost := (float64(outputTokens) / 1000.0) * pricingOutput
    return inputCost + outputCost
}
```

### Token Counter (Infrastructure)

```go
type TokenCounter interface {
    CountTokens(text string) (int, error)
    CountMessages(messages []ChatMessage) (int, error)
}

// Implementation uses tiktoken-go for OpenAI models
// Falls back to approximate counting (len/4) for unknown models
```

---

# 8. MODEL ROUTER ARCHITECTURE

## Components

```
┌──────────────────────────────────────────┐
│            Model Router                    │
│                                          │
│  ┌─────────────┐  ┌──────────────────┐   │
│  │ Capability  │  │   Cost Engine     │   │
│  │  Matcher    │  │  (budget-aware)   │   │
│  └──────┬──────┘  └───────┬──────────┘   │
│         │                 │              │
│  ┌──────▼──────────────────▼──────────┐  │
│  │      Routing Decision Engine       │  │
│  │  - Task type analysis             │  │
│  │  - Context length check           │  │
│  │  - Tool compatibility check       │  │
│  │  - Streaming support check        │  │
│  └──────┬──────────────────┬──────────┘  │
│         │                 │              │
│  ┌──────▼──────┐  ┌───────▼──────────┐  │
│  │ Failover    │  │  Provider        │  │
│  │  Manager    │  │  Resolver        │  │
│  └─────────────┘  └──────────────────┘  │
└──────────────────────────────────────────┘
```

## Routing Factors

| Factor | Source | Priority |
|--------|--------|----------|
| Task type | Request analysis | High |
| Required tools | Request + skill config | High |
| Context length | Current conversation | High |
| User preference | Session config | Medium |
| Model health | Health monitor | High |
| Cost | Pricing data | Medium |
| Latency | Historical data | Low |
| Provider priority | Provider config | Low |

---

# 9. LLM PROVIDER MANAGER

## Architecture

```
                    Model Router
                         │
                         ▼
              LLM Provider Manager
                         │
           ┌─────────────┼──────────────┐
           ▼             ▼              ▼
    Authentication   Request Builder  Response Parser
           │             │              │
           ▼             ▼              ▼
         ┌─────────────────────────────────┐
         │      Provider Adapter           │
         │  (OpenAI-compatible interface)   │
         └──────────────┬──────────────────┘
                        │
         ┌──────────────┼──────────────┐
         ▼              ▼              ▼
    OpenAI API     Ollama API     Custom API
```

## Unified Request

```go
type UnifiedRequest struct {
    Model      string           `json:"model"`
    Messages   []ChatMessage    `json:"messages"`
    Tools      []ToolDefinition `json:"tools,omitempty"`
    Stream     bool             `json:"stream"`
    Temperature float64         `json:"temperature,omitempty"`
    MaxTokens  int              `json:"max_tokens,omitempty"`
    Stop       []string         `json:"stop,omitempty"`
    Metadata   map[string]any   `json:"metadata,omitempty"`
}
```

## Unified Response

```go
type UnifiedResponse struct {
    Content      string            `json:"content"`
    ToolCalls    []ToolCall        `json:"tool_calls,omitempty"`
    Usage        *TokenUsage       `json:"usage"`
    FinishReason string            `json:"finish_reason"`
    Model        string            `json:"model"`
    Provider     string            `json:"provider"`
    Latency      time.Duration     `json:"latency"`
}
```

## Supported Providers

| Provider | Adapter | Auth Method | Models Endpoint |
|----------|---------|-------------|-----------------|
| OpenAI | openai_adapter | Bearer token | GET /v1/models |
| OpenRouter | openai_adapter | Bearer token | GET /v1/models |
| Groq | openai_adapter | Bearer token | GET /v1/models |
| DeepSeek | openai_adapter | Bearer token | GET /v1/models |
| Together AI | openai_adapter | Bearer token | GET /v1/models |
| Azure OpenAI | openai_adapter | API Key header | GET /openai/models |
| LiteLLM | openai_adapter | Bearer token | GET /v1/models |
| Ollama | ollama_adapter | None | GET /v1/models |
| LM Studio | openai_adapter | None | GET /v1/models |
| vLLM | openai_adapter | Bearer token | GET /v1/models |
| Hugging Face TGI | hf_adapter | Bearer token | GET /v1/models |
| Custom | openai_adapter | Configurable | Configurable |

---

# 10. STREAM HANDLER & TOKEN MANAGEMENT

## Stream Handler Architecture

```
LLM Provider
    │
    ▼ (SSE Chunks)
┌──────────────────┐
│  Stream Parser    │  Parse SSE events, extract tokens
└────────┬─────────┘
         ▼
┌──────────────────┐
│  Stream Buffer    │  Buffer for smooth rendering
└────────┬─────────┘
         ▼
┌──────────────────┐
│ Stream Aggregator │  Combine chunks into coherent response
└────────┬─────────┘
         ▼
┌──────────────────┐
│ Token Tracker     │  Count tokens, update usage stats
└────────┬─────────┘
         ▼
┌──────────────────┐
│ Stream Dispatcher │  Send to conversation manager + renderer
└────────┬─────────┘
         ▼
  Conversation Manager
         ▼
      Renderer
```

## Stream States

```
Idle → Connecting → Streaming → Completed
                         ↓
                      Paused → Resumed
                         ↓
                    Cancelled / Error
```

## Token Tracking Pipeline

```
Stream Chunk Received
    │
    ▼
Extract content delta
    │
    ▼
Count tokens (approximate: len/4 or tiktoken)
    │
    ▼
Update current usage:
  - Add delta to output_tokens
  - Recalculate total_tokens
  - Recalculate remaining_context
  - Recalculate context_percent
  - Calculate estimated_cost
    │
    ▼
Emit token update event
    │
    ▼
Status bar updates token gauge
```

---

# 11. TUI SHELL ARCHITECTURE

## Application Structure (Bubble Tea)

The application has a **dual-screen** model: a Home Screen when idle, and a Chat Screen when a conversation is active. Both share the same input area and status bar.

```go
type AppModel struct {
    width  int
    height int
    ready  bool

    // Screens
    chatScreen     *ChatScreen
    homeScreen     *HomeScreen

    // Shared
    commandInput   *CommandInput
    statusBar      *StatusBar
    dialogStack    *DialogStack

    // State
    state          *AppState
    eventBus       *EventBus

    // Services (injected)
    chatService      *application.ChatService
    providerService  *application.ProviderService
    modelService     *application.ModelService
    sessionService   *application.SessionService
    toolService      *application.ToolService
    workspaceService *application.WorkspaceService
    tokenTracker     *application.TokenTrackerService
    router           *application.ModelRouter
    streamHandler    *application.StreamHandler

    tea.Model
}
```

## Application Lifecycle

```
tc serve
    │
    ▼
Load configuration (Viper)
    │
    ▼
Open database (SQLite)
    │
    ▼
Load default provider, model, agent, workspace
    │
    ▼
Apply theme
    │
    ▼
Display HOME SCREEN
    │
    ▼
User types prompt + Enter
    │
    ▼
Create session (if none active)
    │
    ▼
Switch to CHAT SCREEN
    │
    ▼
Send message, receive response, streaming...
    │
    ▼
Conversation continues on CHAT SCREEN
    │
    ▼
User clears session (/clear) or session ends
    │
    ▼
Return to HOME SCREEN
```

## Screen Types

```go
type ScreenType int
const (
    ScreenHome       ScreenType = iota // Welcome/config screen - no conversation
    ScreenChat                          // Active conversation
    ScreenProviderList
    ScreenProviderAdd
    ScreenModelList
    ScreenModelAdd
    ScreenModelSelect
    ScreenSessionManager
    ScreenSettings
    ScreenHelp
)
```

## Dialog Stack

```go
type Dialog interface {
    Init() tea.Cmd
    Update(msg tea.Msg) (Dialog, tea.Cmd)
    View() string
    Focused() bool
    SetFocused(bool)
}

type DialogStack struct {
    dialogs []Dialog
}

func (s *DialogStack) Push(d Dialog)
func (s *DialogStack) Pop() Dialog
func (s *DialogStack) Peek() Dialog
func (s *DialogStack) IsEmpty() bool
func (s *DialogStack) View() string
```

## Event Bus

```go
type EventBus struct {
    listeners map[string][]EventHandler
    mu        sync.RWMutex
}

type Event struct {
    Type      string
    Data      any
    Timestamp time.Time
}

const (
    EventAppStarted       = "app.started"
    EventModelChanged     = "model.changed"
    EventProviderChanged  = "provider.changed"
    EventSessionCreated   = "session.created"
    EventSessionDeleted   = "session.deleted"
    EventMessageSent      = "message.sent"
    EventMessageReceived  = "message.received"
    EventStreamStarted    = "stream.started"
    EventStreamChunk      = "stream.chunk"
    EventStreamComplete   = "stream.complete"
    EventTokenUpdate      = "token.update"
    EventToolStarted      = "tool.started"
    EventToolComplete     = "tool.complete"
    EventToolFailed       = "tool.failed"
    EventWorkspaceChanged = "workspace.changed"
    EventThemeChanged     = "theme.changed"
    EventConfigChanged    = "config.changed"
    EventAttention        = "attention.required"
    EventNotification     = "notification.show"
    EventScreenChanged    = "screen.changed" // home <-> chat
)
```

---

# 12. UI-UX MOBILE-FIRST DESIGN

## Dual-Mode Screen System

The application has exactly two primary visual modes:

1. **HOME SCREEN** - Shown when there is NO active conversation
2. **CHAT SCREEN** - Shown when a conversation is active

Both modes share the same **Command Input** and **Status Bar** at the bottom. Only the main content area changes.

---

## 12.1 HOME SCREEN

### Purpose

- Show active configuration at a glance (provider, model, agent, workspace)
- Indicate the app is ready and waiting for input
- Provide quick access to configuration via `/` commands

### When Displayed

- On fresh start (no saved sessions)
- After clearing conversation (`/clear`)
- After deleting the last session
- When user explicitly navigates to home

### Layout

```
┌──────────────────────────────────────┐
│          TERM CODE CLI               │
│      Universal Coding Agent          │
│──────────────────────────────────────│
│  Provider : OpenCode Zen             │
│  Model    : DeepSeek V4              │
│  Agent    : General                  │
│  Workspace: ~/projects               │
│──────────────────────────────────────│
│                                      │
│                                      │
│                                      │
│  Press / for commands                │
│                                      │
├──────────────────────────────────────┤
│ > [command input placeholder]        │
├──────────────────────────────────────┤
│ ◐ DeepSeek │ General │ v1.0.0        │
└──────────────────────────────────────┘
```

### Home Screen Elements

| Element | Description | Source |
|---------|-------------|--------|
| Title | `TERM CODE CLI` centered | Hardcoded |
| Subtitle | `Universal Coding Agent` centered | Hardcoded |
| Provider | Current provider name | Active provider config |
| Model | Current model display name | Active model config |
| Agent | Current agent name | Active agent config |
| Workspace | Current workspace path | Workspace config |
| Prompt hint | `> ` with placeholder text | Static |
| Help text | `Press / for commands` | Static |
| Command Input | Always visible at bottom | Shared component |
| Status Bar | Always visible at bottom | Shared component |

### Home Screen States

| State | Content |
|-------|---------|
| Default | Shows current config with all fields populated |
| No Provider | Shows "No provider configured. Run /add provider" |
| No Model | Shows "No model configured. Run /add model" |
| No Workspace | Shows "No workspace selected. Run /workspace" |
| First Run | Shows setup instructions and quick start guide |

---

## 12.2 CHAT SCREEN

### Purpose

- Display conversation history (user + AI messages)
- Show active model name as header
- Render AI responses with streaming, markdown, code highlighting
- Show tool execution cards interleaved
- Keep input and status bar fixed at bottom

### When Displayed

- User sends first message from Home Screen
- User resumes an existing session
- User selects a session from session manager

### Layout

```
┌──────────────────────────────────────┐
│                                      │
│  User                                │
│  > Create a REST API in Go          │
│                                      │
│  Assistant                           │
│  I will help you create a Go         │
│  REST API...                         │
│                                      │
│  ✓ reading go.mod                    │
│  ✓ writing main.go                   │
│  ✓ completed                         │
│                                      │
├──────────────────────────────────────┤
│ > [typing area]                      │
├──────────────────────────────────────┤
│ Zen │ DeepSeek │ General │ 2.8K      │
└──────────────────────────────────────┘
```

### Layout with Keyboard Open

```
┌──────────────────────────────────────┐
│  User                                │
│  > Create a REST API...              │
│                                      │
│  Assistant (last visible messages)   │
│                                      │
├──────────────────────────────────────┤
│ > [type here...]                     │
├──────────────────────────────────────┤
│ Zen │ DeepSeek │ General │ 2.8K      │
├──────────────────────────────────────┤
│       Android Keyboard               │
└──────────────────────────────────────┘
```

### Chat Screen Elements

| Element | Description | Behavior |
|---------|-------------|----------|
| Messages | Conversation history | Scrollable viewport; newest at bottom |
| User Message | `> ` prefix | Left aligned, no bubble |
| Assistant Message | Model prefix or none | Left aligned, streaming supported |
| Tool Card | `[Tool Name]` with status | Shows during execution |
| Streaming Cursor | `▌` at end of text | Visible only during generation |
| Command Input | Always bottom | Shared component |
| Status Bar | Always below input | Shared component |

### Chat Screen States

| State | Content |
|-------|---------|
| Empty | "No conversation. Type a prompt below." |
| Loading | "Loading session..." |
| Conversation | Normal message history |
| Streaming | Real-time token display with `▌` cursor |
| Tool Execution | Tool cards interspersed |
| Error | Error message with retry option |

---

## 12.3 Screen Transition: Home -> Chat

```
HOME SCREEN
    │
    │ User types prompt + Enter
    ▼
Create new session (auto-title)
    │
    │
    ▼
Switch to CHAT SCREEN
    │
    │ Conversation area replaces home content
    │ Input + Status Bar remain unchanged
    │
    ▼
CHAT SCREEN (streaming response)

    │ User types /clear
    ▼
Clear session
    │
    ▼
HOME SCREEN (config overview)

```

---

## 12.4 Shared Components

Both Home and Chat screens share:

```
┌──────────────────────────────────────┐
│  [MAIN CONTENT - changes per screen] │
├──────────────────────────────────────┤
│  > Command Input (always visible)    │  ← Shared
├──────────────────────────────────────┤
│  Status Bar (always visible)         │  ← Shared
└──────────────────────────────────────┘
```

---

## 12.5 Master Layout Summary

```
HOME SCREEN (no conversation)
┌──────────────────────────────────────┐
│          TERM CODE CLI               │
│      Universal Coding Agent          │
│                                      │
│  Provider : OpenAI                   │
│  Model    : GPT-5                    │
│  Agent    : General                  │
│  Workspace: ~/my-project             │
│                                      │
│  > Type a message...                 │
│  Press / for commands                │
├──────────────────────────────────────┤
│ > [input area]                       │
├──────────────────────────────────────┤
│ GPT-5 │ my-project │ Ready │ 2.8K   │
└──────────────────────────────────────┘

CHAT SCREEN (active conversation)
┌──────────────────────────────────────┐
│  User                                │
│  > Hello                             │
│                                      │
│  Assistant                           │
│  Hello! How can I help...           │
├──────────────────────────────────────┤
│ > [input area]                       │
├──────────────────────────────────────┤
│ GPT-5 │ my-project │ Ready │ 2.8K   │
└──────────────────────────────────────┘
```

## Key Layout Rules

1. **Home Screen has title** - `TERM CODE CLI` + subtitle at top
2. **No headers on Chat Screen, dialogs, palettes, or popups** - content starts from top
3. **Status Bar is ALWAYS below Input Area** - Never above
4. **Input Area is ALWAYS visible** - Keyboard open or closed
5. **Home Screen**: Title + config overview when idle
6. **Chat Screen**: Conversation with messages starting from top
7. **Transition**: Home -> Chat is instant on first message
8. **Both screens share** input area and status bar
9. **No sidebars** - All navigation via `/` commands or modal overlays
10. **Single column** - No split views or multi-column layouts
11. **Portrait only** - Landscape not supported

## Touch Zones

```
┌──────────────────────────────────────┐
│                                      │
│         Hard to Reach                │
│         (Top of content)             │
│                                      │
├──────────────────────────────────────┤
│                                      │
│         Easy to Reach                │
│         (Scroll zone)                │
│                                      │
├──────────────────────────────────────┤
│   > Easy to Reach (Input)            │  ← Thumb zone
├──────────────────────────────────────┤
│   Model │ Status │ Tokens            │  ← Thumb zone
└──────────────────────────────────────┘
```

---

# 13. KEY SCREENS & COMPONENTS

## 13.1 Home Screen (Idle/Welcome)

### Purpose

Shown when there is no active conversation. Displays the app identity and current configuration so the user knows the state before typing.

### Wireframe

```
┌──────────────────────────────────────┐
│                                      │
│          TERM CODE CLI               │
│      Universal Coding Agent          │
│                                      │
│                                      │
│  Provider : OpenCode Zen             │
│  Model    : DeepSeek V4              │
│  Agent    : General                  │
│  Workspace: ~/projects/termcode     │
│                                      │
│                                      │
│  > Type a message or / for commands  │
│                                      │
├──────────────────────────────────────┤
│ >                                    │
├──────────────────────────────────────┤
│ Zen │ DeepSeek │ General │ Ready     │
└──────────────────────────────────────┘
```

### States

| State | Display |
|-------|---------|
| Default | All config fields populated |
| No Provider | `Provider : Not configured` + `/add provider` hint |
| No Model | `Model : Not configured` + `/add model` hint |
| No Agent | `Agent : General (default)` (always has default) |
| No Workspace | `Workspace: None` + `/workspace` hint |
| First Run | Welcome message + quick setup instructions |

### Transitions

- User types message + Enter -> **Chat Screen** (auto-creates session)
- User types `/add provider` -> overlays dialog
- User types `/add model` -> overlays dialog
- User types `/model` -> opens model selector dialog
- User types `/provider` -> opens provider selector dialog

---

## 13.2 Chat Screen (Active Conversation)

### Purpose

Display conversation history with AI. Interleaves messages, tool executions, and streaming responses.

### Wireframe

```
┌──────────────────────────────────────┐
│                                      │
│  User                                │
│  > Create a Go REST API              │
│                                      │
│  Assistant                           │
│  I will help you create a Go         │
│  REST API with the following         │
│  structure...                        │
│                                      │
│  ✓ reading go.mod                    │
│  ✓ writing main.go                   │
│  ✓ completed                         │
│                                      │
│  Assistant                           │
│  Here is the complete code...        │
│  ```go                              │
│  func main() {                      │
│      // ...                         │
│  }                                  │
│  ```                                │
│                                      │
├──────────────────────────────────────┤
│ >                                    │
├──────────────────────────────────────┤
│ Zen │ DeepSeek │ General │ 2.8K      │
└──────────────────────────────────────┘
```

### States

- **Empty**: "No conversation. Type a prompt below."
- **Loading**: "Loading session..."
- **Conversation**: Normal message history
- **Streaming**: Real-time token display with cursor `▌`
- **Tool Execution**: Tool cards interspersed between messages
- **Error**: Error message with retry option

### Message Types

```
User
  > How do I create a React component?

Assistant
  I will help you create a React component...

Tool
  [Reading src/components/...]
  [Writing src/components/Button.tsx]
  [Completed]

System
  Model changed to GPT-5
  Session saved
```

### Message Rendering

- No chat bubbles - terminal-style messages
- Left-aligned
- User prefix: `> `
- Assistant prefix: none (or model name)
- Tool prefix: `[Tool Name]`
- System prefix: muted color
- Code blocks: syntax highlighted
- Streaming: cursor `▌` visible during generation

## 13.2 Command Input

```
┌──────────────────────────────────────┐
│ > Type your prompt here...           │
└──────────────────────────────────────┘
```

### Features

- Single/multi-line (up to 8 lines)
- Slash command autocomplete (/model, /provider, etc.)
- File path autocomplete
- Command history (up/down arrows)
- Draft persistence
- Prompt prefix `>`
- Placeholder text when empty

## 13.3 Status Bar

```
┌──────────────────────────────────────┐
│ ◐ DeepSeek │ General │ v1.0.0        │
└──────────────────────────────────────┘
```

### Sections (left to right)

| Section | Description | Priority |
|---------|-------------|----------|
| Spinner | Animated loading indicator when AI is working | High |
| Model | Active model name (truncated) | High |
| Agent | Current agent name | Medium |
| Version | App version (e.g. v1.0.0) | Low |

### Spinner Animation

Progress bar animation that runs continuously while AI is working:

```
Frame 1: [⬝⬝⬝⬝⬝⬝]   (all dots)
Frame 2: [■⬝⬝⬝⬝⬝]
Frame 3: [■■⬝⬝⬝⬝]
Frame 4: [■■■⬝⬝⬝]
Frame 5: [■■■■⬝⬝]
Frame 6: [■■■■■⬝]
Frame 7: [■■■■■■]
Frame 8: [■■■■■⬝]
Frame 9: [■■■■⬝⬝]
Frame 10: [■■■⬝⬝⬝]
Frame 11: [■■⬝⬝⬝⬝]
Frame 12: [■⬝⬝⬝⬝⬝]
```

- Visible only when AI is working (streaming, thinking, tool execution)
- Hidden when idle/ready
- Updates every 150ms
- Uses `⬝` (U+2B1D) for empty and `■` (U+25A0) for filled

### Dynamic Width Behavior

- Small (< 30 cols): `[■■■⬝⬝⬝] DeepSeek`
- Medium (30-50 cols): `[■■■⬝⬝⬝] DeepSeek │ General`
- Large (> 50 cols): `[■■■⬝⬝⬝] DeepSeek │ General │ v1.0.0`

## 13.4 Model Selector (Overlay)

```
┌──────────────────────────────────────┐
│ Select Model                         │
├──────────────────────────────────────┤
│ Search models...                     │
├──────────────────────────────────────┤
│ ✓ GPT-5                              │
│   OpenAI · Coding + Reasoning        │
│                                      │
│   Claude Opus                        │
│   Anthropic · Long Context           │
│                                      │
│ ── Local Models ──                   │
│   Llama 3 (Ollama)                   │
│   Qwen 2.5 (Ollama)                  │
│                                      │
│ ── Favorites ──                      │
│   ★ DeepSeek V4                      │
└──────────────────────────────────────┘
```

## 13.5 Provider List Screen

```
┌──────────────────────────────────────┐
│ Provider Management                  │
├──────────────────────────────────────┤
│ ● OpenCode Zen              312ms    │
│   https://api.opencode.ai/v1         │
│                                      │
│ ○ OpenAI                    Disabled │
│                                      │
│ ● Ollama (Local)             18ms    │
│   http://localhost:11434/v1          │
├──────────────────────────────────────┤
│ [/add] Add Provider                  │
└──────────────────────────────────────┘
```

## 13.6 Add Provider Screen (Form)

```
┌──────────────────────────────────────┐
│ Add Provider                         │
├──────────────────────────────────────┤
│ Provider Name                        │
│ [____________________________]       │
│                                      │
│ Base URL                             │
│ [https://api.openai.com/v1_____]     │
│                                      │
│ API Key                              │
│ [****************************]       │
│                                      │
│ Description (optional)               │
│ [____________________________]       │
├──────────────────────────────────────┤
│ [Save & Test]       [Cancel]         │
└──────────────────────────────────────┘
```

## 13.7 Model List Screen (/all models)

```
┌──────────────────────────────────────┐
│ All Models                           │
├──────────────────────────────────────┤
│ Search models...                     │
├──────────────────────────────────────┤
│ ★ GPT-5                              │
│   OpenAI · Coding · 128K             │
│                                      │
│   Claude Opus                        │
│   Anthropic · 200K                   │
│                                      │
│   DeepSeek V4                        │
│   OpenCode Zen · Reasoning · 128K    │
│                                      │
│ ── Local ──                          │
│   ★ Llama 3 70B                      │
│   Ollama · 128K                      │
├──────────────────────────────────────┤
│ [/add] Add Model                     │
└──────────────────────────────────────┘
```

---

---

## 13.3 Screen State Machine

```
                    ┌─────────────┐
                    │  App Start  │
                    └──────┬──────┘
                           │
                           ▼
                    ┌─────────────┐
              ┌────▶│  HOME       │◀────────┐
              │     │  SCREEN     │         │
              │     └──────┬──────┘         │
              │            │                │
              │      ┌─────┴─────┐          │
              │      │           │          │
              │      ▼           ▼          │
              │  type msg    /command       │
              │      │           │          │
              │      ▼           ▼          │
              │  create      dialog/        │
              │  session     overlay        │
              │      │           │          │
              │      ▼           │          │
              │  ┌─────────┐    │          │
              │  │  CHAT   │    │          │
              │  │  SCREEN │    │          │
              │  └────┬────┘    │          │
              │       │         │          │
              │  ┌────┴────┐    │          │
              │  │         │    │          │
              │  ▼         ▼    │          │
              │ /clear   /cmd   │          │
              │  │         │    │          │
              └──┘         └────┘          │
                                           │
                    ┌─────────────┐         │
                    │  QUIT       │─────────┘
                    └─────────────┘
```

# 14. STATE MANAGEMENT

## Global State Tree

```go
type AppState struct {
    // Application
    App AppStateData

    // Workspace
    Workspace WorkspaceState

    // AI Configuration
    Provider ProviderState
    Model    ModelState

    // Session
    Session SessionState
    Conversation ConversationState

    // Tools
    Tools ToolExecutionState

    // Streaming
    Streaming StreamingState

    // UI
    UI UIState
}

type AppStateData struct {
    Version    string
    StartTime  time.Time
    ConfigPath string
}

type ProviderState struct {
    Current       *domain.Provider
    List          []*domain.Provider
    Health        map[string]ProviderHealth
}

type ModelState struct {
    Current       *domain.Model
    List          []*domain.Model
    Recent        []*domain.Model
    Favorites     []*domain.Model
}

type SessionState struct {
    Current       *domain.Session
    List          []*domain.SessionSummary
}

type ConversationState struct {
    Messages      []*domain.Message
    InputDraft    string
    ScrollPos     int
}

type StreamingState struct {
    IsStreaming   bool
    CurrentText   string
    Tokens        TokenUsage
    StartTime     time.Time
}

type ToolExecutionState struct {
    ActiveTools   []ToolExecution
    History       []ToolExecution
}

type UIState struct {
    Screen         ScreenType
    DialogStack    []DialogType
    LayoutMode     LayoutMode
    Theme          string
    Focus          FocusTarget
}
```

## State Update Pattern

```go
// Every state update follows this pattern:
func (m *AppModel) handleEvent(evt Event) {
    // 1. Compute new state
    newState := m.state.Clone()

    // 2. Apply changes
    switch evt.Type {
    case EventModelChanged:
        newState.Model.Current = evt.Data.(*domain.Model)
    case EventStreamChunk:
        newState.Streaming.CurrentText += evt.Data.(string)
        newState.Streaming.Tokens.OutputTokens++
    }

    // 3. Update UI
    m.state = newState
    m.updateUI()
}
```

---

# 15. DATABASE SCHEMA

## Tables

```sql
-- Providers
CREATE TABLE providers (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    base_url    TEXT NOT NULL,
    api_key     TEXT,  -- encrypted at rest
    description TEXT DEFAULT '',
    status      TEXT DEFAULT 'unknown',
    latency     INTEGER DEFAULT 0,  -- milliseconds
    priority    INTEGER DEFAULT 0,
    is_default  INTEGER DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Models
CREATE TABLE models (
    id                  TEXT PRIMARY KEY,
    provider_id         TEXT NOT NULL REFERENCES providers(id) ON DELETE CASCADE,
    model_id            TEXT NOT NULL,
    display_name        TEXT DEFAULT '',
    description         TEXT DEFAULT '',
    category            TEXT DEFAULT 'general',
    supports_streaming  INTEGER DEFAULT 1,
    supports_tools      INTEGER DEFAULT 0,
    supports_reasoning  INTEGER DEFAULT 0,
    supports_vision     INTEGER DEFAULT 0,
    supports_embeddings INTEGER DEFAULT 0,
    supports_json_mode  INTEGER DEFAULT 0,
    max_context         INTEGER DEFAULT 4096,
    max_output          INTEGER DEFAULT 4096,
    pricing_input       REAL DEFAULT 0.0,
    pricing_output      REAL DEFAULT 0.0,
    is_local            INTEGER DEFAULT 0,
    is_favorite         INTEGER DEFAULT 0,
    enabled             INTEGER DEFAULT 1,
    created_at          TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at          TEXT NOT NULL DEFAULT (datetime('now')),
    UNIQUE(provider_id, model_id)
);

-- Sessions
CREATE TABLE sessions (
    id          TEXT PRIMARY KEY,
    title       TEXT DEFAULT 'New Session',
    provider_id TEXT REFERENCES providers(id),
    model_id    TEXT REFERENCES models(id),
    agent_id    TEXT REFERENCES agents(id),
    workspace_id TEXT REFERENCES workspaces(id),
    summary     TEXT DEFAULT '',
    is_favorite INTEGER DEFAULT 0,
    is_archived INTEGER DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Messages
CREATE TABLE messages (
    id          TEXT PRIMARY KEY,
    session_id  TEXT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    role        TEXT NOT NULL CHECK(role IN ('user','assistant','system','tool')),
    content     TEXT NOT NULL DEFAULT '',
    tool_calls  TEXT,  -- JSON array
    token_input INTEGER DEFAULT 0,
    token_output INTEGER DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Agents
CREATE TABLE agents (
    id              TEXT PRIMARY KEY,
    name            TEXT NOT NULL UNIQUE,
    description     TEXT DEFAULT '',
    system_prompt   TEXT DEFAULT '',
    temperature     REAL DEFAULT 0.7,
    reasoning_level TEXT DEFAULT 'normal',
    allowed_tools   TEXT DEFAULT '[]',  -- JSON array
    default_skills  TEXT DEFAULT '[]',  -- JSON array
    enabled         INTEGER DEFAULT 1,
    is_builtin      INTEGER DEFAULT 0,
    created_at      TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at      TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Skills
CREATE TABLE skills (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    version     TEXT DEFAULT '1.0',
    category    TEXT DEFAULT 'general',
    description TEXT DEFAULT '',
    priority    INTEGER DEFAULT 0,
    path        TEXT DEFAULT '',
    enabled     INTEGER DEFAULT 1,
    is_builtin  INTEGER DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Tools
CREATE TABLE tools (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    category    TEXT DEFAULT 'general',
    description TEXT DEFAULT '',
    schema      TEXT DEFAULT '{}',  -- JSON schema
    plugin_id   TEXT,
    enabled     INTEGER DEFAULT 1,
    version     TEXT DEFAULT '1.0'
);

-- Tool Logs
CREATE TABLE tool_logs (
    id          TEXT PRIMARY KEY,
    session_id  TEXT REFERENCES sessions(id),
    tool_name   TEXT NOT NULL,
    arguments   TEXT DEFAULT '{}',
    result      TEXT DEFAULT '',
    status      TEXT DEFAULT 'completed',
    duration_ms INTEGER DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Workspaces
CREATE TABLE workspaces (
    id              TEXT PRIMARY KEY,
    name            TEXT NOT NULL,
    path            TEXT NOT NULL UNIQUE,
    language        TEXT DEFAULT '',
    framework       TEXT DEFAULT '',
    package_manager TEXT DEFAULT '',
    git_branch      TEXT DEFAULT '',
    last_scan       TEXT,
    created_at      TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Settings (key-value)
CREATE TABLE settings (
    key         TEXT PRIMARY KEY,
    value       TEXT NOT NULL,
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- MCP Servers
CREATE TABLE mcp_servers (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    transport   TEXT NOT NULL CHECK(transport IN ('stdio','sse','websocket')),
    command     TEXT DEFAULT '',
    args        TEXT DEFAULT '[]',  -- JSON array
    url         TEXT DEFAULT '',
    env         TEXT DEFAULT '{}',  -- JSON object
    status      TEXT DEFAULT 'disconnected',
    enabled     INTEGER DEFAULT 1,
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Plugins
CREATE TABLE plugins (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    version     TEXT DEFAULT '1.0',
    author      TEXT DEFAULT '',
    description TEXT DEFAULT '',
    enabled     INTEGER DEFAULT 1,
    installed_at TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Themes
CREATE TABLE themes (
    id          TEXT PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    author      TEXT DEFAULT '',
    version     TEXT DEFAULT '1.0',
    palette     TEXT NOT NULL DEFAULT '{}',  -- JSON
    is_dark     INTEGER DEFAULT 1,
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Permissions
CREATE TABLE permissions (
    tool_name   TEXT PRIMARY KEY,
    permission  TEXT NOT NULL CHECK(permission IN ('always_allow','allow_once','ask','deny')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Favorites
CREATE TABLE favorites (
    id          TEXT PRIMARY KEY,
    type        TEXT NOT NULL CHECK(type IN ('model','session','command','skill','agent','file')),
    target_id   TEXT NOT NULL,
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    UNIQUE(type, target_id)
);

-- History (recently used)
CREATE TABLE history (
    id          TEXT PRIMARY KEY,
    type        TEXT NOT NULL CHECK(type IN ('command','model','provider','session','file')),
    target_id   TEXT NOT NULL,
    target_name TEXT DEFAULT '',
    used_at     TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Statistics
CREATE TABLE statistics (
    id          TEXT PRIMARY KEY,
    metric      TEXT NOT NULL,
    value       REAL NOT NULL,
    recorded_at TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Cache
CREATE TABLE cache (
    key         TEXT PRIMARY KEY,
    value       TEXT NOT NULL,
    expires_at  TEXT,
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Bookmarks
CREATE TABLE bookmarks (
    id          TEXT PRIMARY KEY,
    session_id  TEXT REFERENCES sessions(id) ON DELETE CASCADE,
    message_id  TEXT REFERENCES messages(id) ON DELETE SET NULL,
    label       TEXT DEFAULT '',
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);

-- Attachments
CREATE TABLE attachments (
    id          TEXT PRIMARY KEY,
    session_id  TEXT REFERENCES sessions(id) ON DELETE CASCADE,
    message_id  TEXT REFERENCES messages(id) ON DELETE SET NULL,
    filename    TEXT NOT NULL,
    path        TEXT NOT NULL,
    mime_type   TEXT DEFAULT 'text/plain',
    size_bytes  INTEGER DEFAULT 0,
    created_at  TEXT NOT NULL DEFAULT (datetime('now'))
);
```

## Indexes

```sql
CREATE INDEX idx_models_provider ON models(provider_id);
CREATE INDEX idx_models_category ON models(category);
CREATE INDEX idx_models_favorite ON models(is_favorite);
CREATE INDEX idx_models_local ON models(is_local);
CREATE INDEX idx_messages_session ON messages(session_id);
CREATE INDEX idx_sessions_updated ON sessions(updated_at);
CREATE INDEX idx_tool_logs_session ON tool_logs(session_id);
CREATE INDEX idx_tool_logs_tool ON tool_logs(tool_name);
CREATE INDEX idx_history_type ON history(type);
CREATE INDEX idx_history_used ON history(used_at);
CREATE INDEX idx_statistics_metric ON statistics(metric);
CREATE INDEX idx_cache_expires ON cache(expires_at);
CREATE INDEX idx_favorites_type ON favorites(type);
```

## Migrations

```go
type Migration struct {
    Version int
    Name    string
    Up      string // SQL to apply
    Down    string // SQL to rollback
}

var migrations = []Migration{
    {1, "initial_schema", up1, down1},
    {2, "add_local_models", up2, down2},
    {3, "add_token_tracking", up3, down3},
    // ...
}
```

---

# 16. EVENT BUS SYSTEM

## Event Definitions

```go
// App Events
const (
    EventAppStarted     = "app.started"
    EventAppShutdown    = "app.shutdown"
)

// Provider Events
const (
    EventProviderCreated  = "provider.created"
    EventProviderUpdated  = "provider.updated"
    EventProviderDeleted  = "provider.deleted"
    EventProviderSelected = "provider.selected"
    EventProviderTested   = "provider.tested"
    EventProviderHealth   = "provider.health"
)

// Model Events
const (
    EventModelCreated  = "model.created"
    EventModelUpdated  = "model.updated"
    EventModelDeleted  = "model.deleted"
    EventModelSelected = "model.selected"
    EventLocalDiscovered = "model.local_discovered"
)

// Session Events
const (
    EventSessionCreated   = "session.created"
    EventSessionDeleted   = "session.deleted"
    EventSessionUpdated   = "session.updated"
    EventSessionSwitched  = "session.switched"
)

// Conversation Events
const (
    EventMessageSent     = "message.sent"
    EventMessageReceived = "message.received"
)

// Stream Events
const (
    EventStreamStart    = "stream.start"
    EventStreamChunk    = "stream.chunk"
    EventStreamEnd      = "stream.end"
    EventStreamError    = "stream.error"
    EventStreamCancel   = "stream.cancel"
)

// Tool Events
const (
    EventToolStart   = "tool.start"
    EventToolEnd     = "tool.end"
    EventToolError   = "tool.error"
    EventToolPermission = "tool.permission"
)

// Workspace Events
const (
    EventWorkspaceChanged = "workspace.changed"
    EventWorkspaceScanned = "workspace.scanned"
)

// Token Events
const (
    EventTokenUpdate    = "token.update"
    EventTokenWarning   = "token.warning"
    EventTokenCritical  = "token.critical"
)

// UI Events
const (
    EventThemeChanged   = "theme.changed"
    EventLayoutChanged  = "layout.changed"
    EventNotification   = "notification.show"
    EventAttention      = "attention.required"
)
```

## Event Bus Implementation

```go
type EventHandler func(event Event)

type EventBus struct {
    handlers map[string][]EventHandler
    mu       sync.RWMutex
    logger   *slog.Logger
}

func NewEventBus(logger *slog.Logger) *EventBus
func (eb *EventBus) Subscribe(eventType string, handler EventHandler) func() // returns unsubscribe
func (eb *EventBus) Publish(event Event)
func (eb *EventBus) PublishAsync(event Event) // non-blocking
```

---

# 17. CONFIGURATION SYSTEM

## Configuration Layers (Viper)

```
1. Defaults (hardcoded)
2. Global Config (~/.config/tc/config.yaml)
3. Workspace Config (.tc/config.yaml in project)
4. Environment Variables (TC_*)
5. Runtime Overrides (/settings commands)
6. CLI Flags (--model, --provider)
```

## Config Structure

```go
type Config struct {
    AI struct {
        DefaultProvider string  `mapstructure:"default_provider"`
        DefaultModel    string  `mapstructure:"default_model"`
        Temperature     float64 `mapstructure:"temperature"`
        MaxTokens       int     `mapstructure:"max_tokens"`
        Streaming       bool    `mapstructure:"streaming"`
        StreamSpeed     string  `mapstructure:"stream_speed"` // realtime, fast, instant
    } `mapstructure:"ai"`

    UI struct {
        Theme           string `mapstructure:"theme"`
        CompactMode     bool   `mapstructure:"compact_mode"`
        ShowTokenGauge  bool   `mapstructure:"show_token_gauge"`
        StatusBarItems  string `mapstructure:"status_bar_items"`
        AnimationSpeed  string `mapstructure:"animation_speed"`
    } `mapstructure:"ui"`

    Editor struct {
        TabSize         int    `mapstructure:"tab_size"`
        WordWrap        bool   `mapstructure:"word_wrap"`
        LineNumbers     bool   `mapstructure:"line_numbers"`
        SyntaxHighlight bool   `mapstructure:"syntax_highlight"`
    } `mapstructure:"editor"`

    Session struct {
        AutoSave        bool   `mapstructure:"auto_save"`
        AutoSaveInterval int   `mapstructure:"auto_save_interval"`
        MaxHistory      int    `mapstructure:"max_history"`
        DefaultTitle    string `mapstructure:"default_title"`
    } `mapstructure:"session"`

    Storage struct {
        DatabasePath string `mapstructure:"database_path"`
        MaxBackups   int    `mapstructure:"max_backups"`
        AutoBackup   bool   `mapstructure:"auto_backup"`
    } `mapstructure:"storage"`

    Security struct {
        EncryptKeys     bool   `mapstructure:"encrypt_keys"`
        DefaultPermission string `mapstructure:"default_permission"` // ask, allow, deny
        MaskAPIKeys     bool   `mapstructure:"mask_api_keys"`
    } `mapstructure:"security"`

    Network struct {
        Timeout         int    `mapstructure:"timeout"` // seconds
        RetryCount      int    `mapstructure:"retry_count"`
        ProxyURL        string `mapstructure:"proxy_url"`
    } `mapstructure:"network"`
}
```

---

# 18. SECURITY & PERMISSIONS

## Permission Tiers

| Tier | Description | Default |
|------|-------------|---------|
| Always Allow | No prompt needed | read, search, list |
| Allow Once | Prompt on first use, then allow | write, edit |
| Ask | Always prompt | shell, delete, install |
| Deny | Never allow | configurable |

## API Key Security

- Encrypted at rest using AES-256-GCM
- Encryption key derived from machine ID + salt
- Keys never appear in logs, exports, or error messages
- Key display: `sk-****5678` (last 4 chars only)
- Environment variable override: `TC_PROVIDER_<NAME>_API_KEY`

## Permission Store

```go
type PermissionEntry struct {
    ToolName   string
    Permission PermissionLevel
    UpdatedAt  time.Time
}

type PermissionLevel string
const (
    PermissionAlwaysAllow PermissionLevel = "always_allow"
    PermissionAllowOnce   PermissionLevel = "allow_once"
    PermissionAsk         PermissionLevel = "ask"
    PermissionDeny        PermissionLevel = "deny"
)
```

---

# 19. DEVELOPMENT WORKFLOW

## Build Commands

```bash
# Development
go build -o tc ./cmd/tc/

# Static build (Termux)
CGO_ENABLED=0 go build -ldflags="-s -w -extldflags=-static" -o tc ./cmd/tc/

# Run
./tc serve

# Test
go test ./... -v -count=1

# Lint
gofumpt -l -w .
goimports -local -w .
go vet ./...
```

## Commit Workflow

1. Format: `gofumpt -l -w . && goimports -local -w .`
2. Lint: `go vet ./...`
3. Test: `go test ./... -v -count=1`
4. Build: `go build ./...`

---

# 20. IMPLEMENTATION ROADMAP

## Phase 1: Foundation (Weeks 1-2)

- [ ] Set up Go module structure
- [ ] Implement database schema + migrations (SQLite)
- [ ] Create domain entities (provider, model, session, etc.)
- [ ] Implement provider repository (SQLite)
- [ ] Implement model repository (SQLite)
- [ ] Implement provider service (CRUD + test connection)
- [ ] Implement model service (CRUD + validation)
- [ ] Set up Viper configuration system
- [ ] Set up Charm Log logging
- [ ] Implement API key encryption

## Phase 2: Core TUI Shell (Weeks 3-4)

- [ ] Create main Bubble Tea application model (dual-screen: home + chat)
- [ ] Implement status bar component (shared between screens)
- [ ] Implement command input component (shared between screens)
- [ ] Implement HOME SCREEN (config overview, welcome, ready state)
- [ ] Implement CHAT SCREEN (model header + viewport + messages)
- [ ] Implement home screen states (default, no-provider, no-model, first-run)
- [ ] Implement screen state machine (home <-> chat transitions)
- [ ] Implement message list + item rendering
- [ ] Implement layout system (responsive, keyboard-aware)
- [ ] Set up theme/color system (Lip Gloss)
- [ ] Implement dialog stack
- [ ] Implement event bus

## Phase 3: AI Communication (Weeks 5-6)

- [ ] Implement OpenAI-compatible HTTP client (resty)
- [ ] Implement streaming SSE parser
- [ ] Implement stream handler + buffer
- [ ] Implement token tracker
- [ ] Implement model router (capability match + cost engine)
- [ ] Implement provider manager + adapters
- [ ] Implement failover manager
- [ ] Implement request builder + response parser

## Phase 4: Provider & Model UI (Weeks 7-8)

- [ ] Implement provider list screen
- [ ] Implement add provider form
- [ ] Implement model list screen (/all models)
- [ ] Implement add model form
- [ ] Implement model selector overlay
- [ ] Implement provider test connection
- [ ] Implement model auto-discovery from provider
- [ ] Implement local model management

## Phase 5: Tools & Execution (Weeks 9-10)

- [ ] Implement file reader/writer/editor
- [ ] Implement search service (ripgrep + fd)
- [ ] Implement diff computation
- [ ] Implement shell executor with permission
- [ ] Implement git service (go-git)
- [ ] Implement tool card rendering
- [ ] Implement permission system
- [ ] Implement workspace scanner

## Phase 6: Advanced Features (Weeks 11-12)

- [ ] Implement MCP client + server management
- [ ] Implement session management (CRUD + export/import)
- [ ] Implement context management (compaction)
- [ ] Implement backup/restore
- [ ] Implement plugin loader
- [ ] Implement skill system
- [ ] Implement agent system
- [ ] Implement file watcher (fsnotify)
- [ ] Implement tree-sitter code analysis

## Phase 7: Polish & Optimization (Weeks 13-14)

- [ ] Performance optimization (memory, startup, rendering)
- [ ] Mobile testing on Termux (various Android versions)
- [ ] Error handling improvement
- [ ] Loading/empty/error states for all screens
- [ ] Keyboard behavior optimization
- [ ] Theme customization
- [ ] Accessibility (high contrast, large fonts)
- [ ] Documentation

---

# APPENDIX A: FILE REFERENCE

## All Docs Used for This Architecture

| File | Purpose |
|------|---------|
| `docs/00-OVERVIEW.md` | Complete concept map |
| `docs/01-IDENTITY-PHILOSOPHY.md` | Identity & philosophy |
| `docs/02-SESSION-LIFECYCLE.md` | Session lifecycle |
| `docs/03-AI-MODEL-PROVIDER.md` | AI model & provider spec |
| `docs/03-ARCHITECTURE.md` | Software architecture spec |
| `docs/04-AI-COMMUNICATION.md` | AI communication flow |
| `docs/05-CONTEXT-MANAGEMENT.md` | Context management |
| `docs/05-PROVIDERS.md` | Provider management |
| `docs/06-MODELS.md` | Model management |
| `docs/06-TOOL-ARCHITECTURE.md` | Tool architecture |
| `docs/07-FILE-OPERATIONS.md` | File operations |
| `docs/08-SHELL-EXECUTION.md` | Shell execution |
| `docs/09-PERMISSION-SECURITY.md` | Permission & security |
| `docs/10-CHAT-ENGINE.md` | Chat engine |
| `docs/11-COMMAND-PALETTE.md` | Command palette |
| `docs/12-DIALOG-SYSTEM.md` | Dialog system |
| `docs/13-DATABASE.MD` | Database schema |
| `docs/13-DISPLAY-RENDERING.md` | Display rendering |
| `docs/14-UI-SHELL.md` | UI shell layout |
| `docs/15-MODEL-ROUTER.md` | Model router |
| `docs/15-THEMING.md` | Theming |
| `docs/16-AGENT-ARCHITECTURE.md` | Agent architecture |
| `docs/17-LLM-PROVIDER-MANAGER.md` | LLM provider manager |
| `docs/17-SKILL-KNOWLEDGE.md` | Skills & knowledge |
| `docs/17-STREAMING.md` | Streaming engine |
| `docs/18-MCP-EXTENSIBILITY.md` | MCP extensibility |
| `docs/18-STREAM-HANDLER.md` | Stream handler |
| `docs/19-STATE-MANAGEMENT.md` | State management |
| `docs/19-ACTION-ENGINE.MD` | Action engine |
| `docs/20-DATA-PERSISTENCE.md` | Data persistence |
| `docs/UI-UX/*.md` | Complete UI-UX specs (40 files) |
| `ai-skills/skills/*/SKILL.md` | All installed skills |
| `sub-agents/*.md` | Agent specifications |

## AI SKILLS Referenced

All skills from `/data/data/com.termux/files/home/termcode/ai-skills/` were analyzed for patterns, best practices, and implementation strategies applicable to this architecture.

---

# APPENDIX B: COMMANDS REFERENCE

## Slash Commands

| Command | Description |
|---------|-------------|
| `/help` | Show help |
| `/model` | Switch model (opens model selector) |
| `/provider` | View/switch provider |
| `/add model` | Add new model |
| `/add provider` | Add new provider |
| `/all models` | List all models |
| `/providers` | List all providers |
| `/session` | Session management |
| `/sessions` | List all sessions |
| `/clear` | Clear conversation |
| `/settings` | Open settings |
| `/theme` | Change theme |
| `/workspace` | Workspace info |
| `/search` | Search workspace |
| `/git` | Git operations |
| `/mcp` | MCP server management |
| `/backup` | Database backup |
| `/restore` | Database restore |
| `/export` | Export session |
| `/import` | Import session |
| `/retry` | Retry last AI response |
| `/cancel` | Cancel current operation |
| `/quit` | Quit application |

## CLI Commands

| Command | Description |
|---------|-------------|
| `tc serve` | Start TUI application |
| `tc provider list` | List providers |
| `tc provider add` | Add provider |
| `tc provider remove` | Remove provider |
| `tc model list` | List models |
| `tc model add` | Add model |
| `tc model remove` | Remove model |
| `tc session list` | List sessions |
| `tc session export` | Export session |
| `tc session import` | Import session |
| `tc config get` | Get config value |
| `tc config set` | Set config value |
| `tc backup` | Create backup |
| `tc restore` | Restore backup |
| `tc version` | Show version |

---

# APPENDIX C: UI COMMANDS

## `/add provider` Screen Flow

```
User types: /add provider

→ Opens Add Provider form
→ User fills: Name, Base URL, API Key, Description
→ User taps "Save & Test"
→ System validates inputs
→ System sends GET /models to test connection
→ On success: save provider, show "Provider connected" notification
→ On failure: show error with details, allow retry
```

## `/add model` Screen Flow

```
User types: /add model

→ Opens Add Model form
→ User fills: Model ID, Display Name, Provider (dropdown), Description
→ User configures Capabilities (toggles)
→ User sets Context Window, Max Output
→ User taps "Save"
→ System validates inputs
→ System checks if model ID exists on provider (GET /models)
→ On success: save model
→ Auto-activate model (optional)
```

## `/all models` Screen Flow

```
User types: /all models

→ Opens Model List screen
→ Shows all models grouped by provider
→ Search bar at top
→ Each model shows: name, provider, category, context size
→ Active model marked with ✓
→ Favorites marked with ★
→ Tap model to view details
→ Tap ✓ to activate
→ Long press for actions: Edit, Delete, Duplicate, Favorite
```

## `/provider switch` Flow

```
User types: /provider

→ Shows provider selection list
→ Each provider shows: name, status (●/○), latency
→ Search support
→ Tap to select
→ System updates default provider
→ Status bar updates
```

---

*End of Architecture Plan - Version 1.0*
*This document serves as the complete architectural reference for the TERM CODE (tc) project.*
