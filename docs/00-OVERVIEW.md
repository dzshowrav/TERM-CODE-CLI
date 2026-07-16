# CLI Coding Agent — Complete Concept Map

## What This Is

A purely conceptual design document for a CLI-native coding agent. Zero tech stack, zero programming language, zero framework assumptions. Every concept is described in universal terms so it can be built on any runtime (compiled, interpreted, or otherwise).

The only preserved translation from the current project is the **AI Model & Provider system** — every other concept is extracted, generalized, and mixed with inspirations from opencode CLI and other industry tools.

## File Map

### Core Foundation
| File | What It Describes |
|------|------------------|
| `01-IDENTITY-PHILOSOPHY` | What the agent IS — identity, principles, multi-surface availability, open ecosystem |
| `02-SESSION-LIFECYCLE` | Session creation, resumption, branching, export, cleanup — the user's conversation unit |

### AI & Communication
| File | What It Describes |
|------|------------------|
| `03-AI-MODEL-PROVIDER` | Provider & model registry, CRUD operations, authentication, priority routing |
| `04-AI-COMMUNICATION` | Streaming, SSE, thinking stages, tool call lifecycle, abort mechanics |
| `05-CONTEXT-MANAGEMENT` | System prompts, context window optimization, compaction, sliding window |

### Tool System
| File | What It Describes |
|------|------------------|
| `06-TOOL-ARCHITECTURE` | Tool definition, registration, execution pipeline, timeout, lifecycle hooks |
| `07-FILE-OPERATIONS` | Read, write, edit, search, glob, git snapshots, diff computation |
| `08-SHELL-EXECUTION` | Command execution, process management, output capture, inline bash |
| `09-PERMISSION-SECURITY` | Permission tiers, policy system, auto-approve, deny, approval prompts |

### Input & Navigation
| File | What It Describes |
|------|------------------|
| `10-INPUT-SYSTEM` | Key detection, paste detection (timing analysis), long-press, char batching, cursor model, @-mentions |
| `11-COMMAND-PALETTE` | Slash commands, fuzzy search, custom commands, inline execution (!) |
| `12-DIALOG-SYSTEM` | Modal dialogs, multi-purpose switchers, forms, selections, navigation |

### Display & UI
| File | What It Describes |
|------|------------------|
| `13-DISPLAY-RENDERING` | Markdown rendering, diff visualization, tool call cards, streaming output, syntax highlighting |
| `14-UI-SHELL` | Layout system, status bar, compact/mobile modes, attention system, screen regions |
| `15-THEMING` | Visual themes, terminal color integration, dynamic theme loading, dark/light |

### Knowledge & Intelligence
| File | What It Describes |
|------|------------------|
| `16-AGENT-ARCHITECTURE` | Multi-agent system, role hierarchy, sub-agent delegation, agent profiles |
| `17-SKILL-KNOWLEDGE` | Skill injection, knowledge bases, prompt templates, auto-detect behavior |
| `18-MCP-EXTENSIBILITY` | External service integration, plugin architecture, custom tool loading |

### Storage & State
| File | What It Describes |
|------|------------------|
| `19-STATE-MANAGEMENT` | Event-driven state, reactive updates, publish-subscribe, derived state |
| `20-DATA-PERSISTENCE` | Storage engine, schema design, backup/restore, migration strategy |

### System & Process
| File | What It Describes |
|------|------------------|
| `21-PROCESS-LIFECYCLE` | Startup sequence, shutdown phases, signal handling, crash recovery, self-update, self-uninstall |
| `22-ERROR-VISUALIZATION` | Error display, silent failures, notification system, log management |
| `23-PROJECT-WORKSPACE` | Project auto-detection, git integration, LSP, file watcher, branch awareness |
| `24-CONFIGURATION` | Settings hierarchy, layering, defaults, file-based config, runtime overrides |
| `25-SESSION-SHARING` | Session export/import, share links, collaboration model |

## Design Principles

1. **Zero lock-in** — Every concept works on any OS, terminal, runtime
2. **Atomic concepts** — Each file describes exactly one concern, no more
3. **Substrate-agnostic** — No tech stack implied anywhere
4. **Progressive complexity** — Simple defaults, deep customization when needed
5. **Failure transparency** — The user always knows what went wrong and why
6. **Offline-first** — Core features work without network; only AI calls need connectivity
7. **Permission by default deny** — Every tool execution needs explicit or policy-based approval
8. **Stateless rendering** — UI is a pure function of state; state is the single source of truth
