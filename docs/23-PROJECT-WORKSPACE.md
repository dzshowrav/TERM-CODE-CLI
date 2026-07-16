# Project & Workspace Awareness

## Overview

The workspace awareness system detects and understands the user's project context. It identifies project type, frameworks, dependencies, git state, and file structure to provide contextually relevant assistance.

## Project Auto-Detection

When the agent starts in a directory, it analyzes:

### Project Type Detection
Detected by the presence of specific files:
- **Node.js** — package.json, node_modules
- **Python** — pyproject.toml, setup.py, requirements.txt, Pipfile
- **Rust** — Cargo.toml
- **Go** — go.mod
- **Java** — pom.xml, build.gradle
- **Ruby** — Gemfile
- **Elixir** — mix.exs
- **Docker** — Dockerfile, docker-compose.yml
- **Generic** — Makefile, justfile, Taskfile.yml

### Framework Detection
Detected by dependencies or configuration files:
- **Frontend** — React, Vue, Angular, Svelte, Astro, Next.js, Nuxt
- **Backend** — Express, FastAPI, Django, Rails, Spring, Gin
- **Tooling** — Vite, Webpack, esbuild, Rollup
- **Testing** — Jest, Vitest, Playwright, pytest, Mocha

### Language Detection
Primary language is determined by:
- File extensions in the project
- Configuration files (package.json → JS/TS, Cargo.toml → Rust)
- Mixed projects identify multiple languages

## Git Integration

### Git State Detection
On startup and on each message:
- Current branch name
- Uncommitted changes (modified, staged, untracked)
- Ahead/behind remote
- Merge/rebase in progress
- Recent commit history

### Git Awareness
The agent understands:
- What branch the user is on
- Whether there are uncommitted changes
- Recent commit messages and authors
- Diff between working tree and HEAD

### Git Operations
The agent can (with permission):
- Create commits (with user-provided messages)
- Create and switch branches
- Stage/unstage files
- View diffs and history
- Create and review PRs (via service integration)
- Tag releases

### Safe Git Practices
- "Commit" is always suggested, never forced
- Destructive operations (rebase, force push) require explicit confirmation
- Git operations that would lose data are flagged
- The agent can propose commit messages based on changes

## LSP Integration

### Language Server Protocol
The agent can connect to Language Servers for:
- **Code intelligence** — go to definition, find references, hover info
- **Diagnostics** — errors, warnings, hints in real-time
- **Completions** — context-aware code completions
- **Code actions** — quick fixes, refactoring suggestions
- **Symbol search** — find by name across the project

### LSP Connection
- LSP servers are started per-language as needed
- Connection is lazy (started when first needed)
- Multiple LSP servers can run simultaneously
- Server output is parsed and presented as structured data

### LSP Benefits
- The agent can understand code semantics, not just text
- Diagnostics catch errors before the AI generates code
- Find references helps understand code usage patterns
- Go to definition enables deep code exploration

## File Structure Awareness

The agent maintains awareness of:
- **Project root** — determined by VCS root or config file
- **Source directories** — common source paths (src/, lib/, app/)
- **Configuration files** — project configuration and tool configs
- **Entry points** — main files, index files, server entry
- **Test directories** — test files, test configuration
- **Build output** — dist/, build/, out/
- **Excluded directories** — node_modules, .git, .venv, __pycache__

## File Watcher

The file watcher monitors the workspace for external changes:

### Watch Events
- **File created** — new file added to workspace
- **File modified** — external edit (user edited outside the agent)
- **File deleted** — file removed
- **Directory created/removed** — structure change

### Watch Responses
Depending on configuration and context:
- **Auto-reload** — if the file is on screen, refresh the display
- **Notification** — notify the user of the change
- **Re-check** — re-detect project type or workspace state
- **Re-read** — if a file is in context, update the AI context

### Watch Scope
- Watches the workspace directory (recursive)
- Respects .gitignore patterns
- Excludes binary and large files
- Configurable include/exclude patterns

## Workspace-Specific Configuration

Each workspace can have its own configuration:
- Preferred agent profile
- Auto-loaded skills
- Custom tools and commands
- Permission policies
- Model/provider preferences
- Environment variables
- Ignored file patterns

Workspace config is stored in a dotfile in the project root. It's version-controllable.

## Key Design Decisions

- Project auto-detection is heuristic-based — it's fast and doesn't require a full scan
- Framework detection enables framework-specific assistance
- Git integration is read-heavy by default — the agent observes more than it modifies
- LSP integration provides semantic understanding, not just text matching
- File watching is non-intrusive — changes are noted, not announced
- Workspace configuration keeps project-specific settings with the project
- The agent prioritizes the user's workspace config over global defaults
- Auto-detection failures are non-fatal — the agent works without project context
