# Command Palette

## Overview

The command palette provides quick access to the agent's commands and features without leaving the input line. It encompasses slash commands, fuzzy command search, and special syntactic shortcuts.

## Slash Commands

Slash commands start with `/` and are typed directly in the input. They are processed before the message is sent to the AI.

### Command Structure
```
/command [subcommand] [--flag] [value] [arg1 arg2 ...]
```

### Built-In Commands
- **`/provider`** — list, add, edit, delete, test AI providers
- **`/model`** — list, set, switch AI models
- **`/session`** — create, list, switch, rename, export sessions
- **`/agents`** — list, load, switch agent profiles
- **`/skills`** — list, load, unload skills
- **`/tools`** — list, enable, disable, inspect tools
- **`/mcp`** — list, connect, disconnect MCP servers
- **`/history`** — view, search, export conversation history
- **`/settings`** — open settings, change specific config values
- **`/themes`** — list, preview, switch themes
- **`/permissions`** — view, modify permission policies
- **`/context`** — view context usage, trigger compaction
- **`/backup`** — backup providers, models, sessions, config
- **`/restore`** — restore from a backup file
- **`/update`** — check for updates, apply update, view changelog
- **`/uninstall`** — clean removal of the agent (with confirmation)
- **`/help`** — list commands, get help on a specific command
- **`/clear`** — clear the current display
- **`/export`** — export current session
- **`/import`** — import a session
- **`/stats`** — show usage statistics
- **`/about`** — version, license, credits

### Custom (User-Defined) Commands
Users can define custom slash commands:
- Map to a sequence of tool calls
- Map to a shell script
- Map to a series of AI prompts
- Defined in configuration, shared across sessions
- Can have subcommands, flags, and arguments

### Command Resolution
1. First, check built-in commands
2. Then, check user-defined commands
3. Then, check skill-provided commands
4. Then, check MCP-provided commands
5. If not found, show "unknown command" with suggestions

## Fuzzy Command Search

When the user types `/` without a known command, or types `/help` with a partial query:
- Commands are searched by name, aliases, and description
- Results are scored by fuzzy match relevance
- Top matches are shown in a dropdown
- User can tab through results or continue typing to narrow
- Shows command syntax and brief description for each match

## Shortcuts & Syntactic Sugar

### Inline Bash (`!`)
- A line starting with `!` is treated as a direct shell command
- `!npm run test` executes `npm run test` in the shell
- Output is displayed inline, not sent to the AI
- Supports all shell syntax (pipes, redirects, variables)

### File Reference (`@`)
- `@filename` references a file (expanded to full path)
- `@./relative/path` references relative to workspace
- `@dir/` references an entire directory
- @-mentions are resolved and replaced with file name chips

### Quick Actions
- `:q` — quick quit (with confirmation if unsaved)
- `:w` — save session
- `:wq` — save and quit
- `:e` — refresh UI

## Command History
- Previously used commands are navigable with arrow keys
- Command history is per-session and global
- History search with Ctrl+R (or equivalent)
- Frequently used commands may be suggested

## Key Design Decisions
- Slash commands are processed BEFORE the AI receives the message
- All commands have consistent syntax (no special parsing rules per command)
- Custom commands extend the system without modifying it
- Fuzzy search makes discovery natural — the user doesn't need to memorize commands
- Inline execution (!) is a direct UX path, not an AI-mediated action
- @-mentions bridge the gap between natural language and file paths
