# Configuration System

## Overview

The configuration system manages all user settings, preferences, and customization options. It uses a layered model where multiple configuration sources are merged into a single effective configuration.

## Configuration Layers

Configuration is resolved from multiple layers (highest priority first):

### Layer 1: Runtime (CLI Flags)
- Command-line arguments passed when starting the agent
- Highest priority — overrides all other layers
- Example: `--model gpt-4`, `--theme tokyo-night`, `--dir /path/to/project`

### Layer 2: Environment Variables
- Environment variables prefixed with the agent's identifier
- Example: `AGENT_MODEL=gpt-4`, `AGENT_THEME=tokyo-night`
- Override file-based config but not CLI flags

### Layer 3: Session-Level
- Settings specific to the current session
- Changed via slash commands or inline settings
- Persist for the session duration (can be saved as defaults)

### Layer 4: Workspace-Level
- Settings specific to the current workspace
- Stored in a dotfile within the project
- Version-controllable (shared with the team)
- Override user and default settings

### Layer 5: User-Level
- The user's personal configuration
- Stored in the user's home directory
- Applies to all workspaces (unless overridden)
- Created on first run (with sensible defaults)

### Layer 6: Default (Built-In)
- Hardcoded default values
- Used when no other layer specifies a value
- Never modified by the user (shipped with the agent)

## Configuration Merging

- Each layer is a key-value map
- Layers are merged in priority order (higher priority overwrites lower)
- String, number, boolean, and array values are overwritten
- Object/map values are deep-merged (nested keys are merged individually)
- Arrays are either replaced entirely or appended (configurable per key)

## Configuration Schema

Every configuration key has:
- **Key path** — dot-notation path (e.g., `ui.theme.name`)
- **Type** — string, number, boolean, array, object
- **Default value** — from the default layer
- **Allowed values** — enum where applicable
- **Validation** — constraints (min/max, pattern, required)
- **Description** — human-readable explanation
- **Category** — grouping for UI display
- **Mutable** — can be changed at runtime or requires restart

### Configuration Categories
- **UI** — theme, layout mode, compact mode, borders, animations
- **Input** — paste threshold, long-press timing, char batch delay
- **AI** — default model, default provider, temperature, max tokens
- **Tools** — default timeout, output truncation limit, retry count
- **Permissions** — auto-approve policies, permission tiers
- **Storage** — data directory, backup interval, retention policy
- **Network** — timeout, retry count, proxy settings
- **Notifications** — enabled events, notification types, quiet hours
- **Behavior** — auto-save interval, compaction threshold, verbose mode

## Configuration Commands

The user interacts with configuration via:
- `/settings` — open the settings dialog (visual editor)
- `/settings get <key>` — get a specific value
- `/settings set <key> <value>` — set a specific value
- `/settings list [category]` — list all settings or by category
- `/settings reset <key>` — reset a key to its default value
- `/settings reset --all` — reset all to defaults
- `/settings export` — export current configuration
- `/settings import <file>` — import configuration

## Configuration File Format

User and workspace configuration files use a human-readable format (YAML, TOML, or JSON):
```yaml
ui:
  theme: tokyo-night
  compact: false

ai:
  default_model: gpt-4o
  temperature: 0.7

permissions:
  auto_approve:
    - tool: read_file
    - tool: glob
```

The format supports:
- Comments
- Multiline values
- Environment variable interpolation
- File path references (relative paths resolved relative to config file)
- Nested structure (no flat keys)

## Configuration Validation

- Config is validated on load
- Invalid values are reported with clear messages (path, expected type, actual value)
- Invalid values are replaced with defaults (graceful degradation)
- Unknown keys are warned but not rejected (forward compatibility)
- Validation errors don't prevent startup — the agent reports and continues

## Configuration Migration

Between versions:
- New keys are added with their defaults
- Deprecated keys trigger warnings with migration instructions
- Removed keys are silently ignored
- Automatic migration converts old format to new if possible
- Migration is logged for reference

## Key Design Decisions

- Layered configuration gives fine-grained control without complexity
- CLI flags are highest priority for temporary overrides
- Workspace config enables team-shared settings
- Deep merge preserves unrelated settings when overriding
- Validation is non-fatal — the agent runs with minimal defaults if config is broken
- Config export/import supports backup and sync
- Category grouping makes settings discoverable in the UI
- Runtime-mutable settings change instantly — no restart needed
