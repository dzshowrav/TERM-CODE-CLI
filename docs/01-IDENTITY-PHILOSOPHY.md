# Identity & Philosophy

## Core Identity

A CLI coding agent is an AI-powered assistant that operates inside a terminal — it reads, writes, searches, and executes code on the user's behalf. It is:

- **A conversation partner** — the user talks to it in natural language
- **A tool operator** — it uses tools (file I/O, shell, search) to accomplish tasks
- **A context-aware worker** — it understands the project structure, git state, and workspace
- **A delegator** — it can spawn sub-agents for specialized work

## Design Philosophy

### Terminal-Native, Not Terminal-Ported

The agent is designed for the terminal from the ground up. It does not try to be a web app in a terminal emulator. Every interaction pattern — input, output, navigation — respects terminal constraints: limited colors, no mouse dependence, no graphics, text-only output, scrollable history.

### Multi-Surface Availability

The same agent core can be exposed through multiple interfaces:
- **Native CLI** — full-featured terminal experience
- **Web browser** — browser-based version with same capabilities
- **Editor extension** — embedded in VS Code, Neovim, or other editors
- **Headless mode** — automated, non-interactive execution

The core logic (AI communication, tool system, agent model) is shared across all surfaces. Only the presentation layer changes.

### Open Ecosystem

The agent is built on an open architecture:
- **Open protocol** — tool definitions follow a standard contract
- **Pluggable** — anyone can add tools, skills, agents, or providers
- **Transparent** — every decision, every tool call, every error is visible
- **Portable** — session data, configuration, and backups are plain-text, portable formats

### Human at the Center

- The user always has final say — every significant action requires approval (unless auto-approved by policy)
- The agent explains its reasoning — thinking stages are visible
- The agent can be interrupted, corrected, or redirected at any point
- The agent adapts to the user's workflow, not the other way around

### Fail Gracefully

- Partial failures are expected (network blips, tool timeouts, malformed input)
- Every failure produces a clear, actionable message
- The agent retries transient failures automatically
- The agent never silently drops work

## Core Tenets

1. **Trust through transparency** — show what you're doing, why, and how
2. **Speed through streaming** — show results as they arrive, not all at once
3. **Power through tools** — the agent is only as capable as its tools
4. **Safety through permissions** — every action is gated by user consent
5. **Memory through sessions** — every conversation is a persistent, resumable session
6. **Quality through specialization** — use the right agent for the right task
7. **Growth through skills** — the agent gets better as skills expand
