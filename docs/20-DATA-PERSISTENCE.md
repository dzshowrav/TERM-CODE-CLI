# Data Persistence

## Overview

The data persistence system handles all long-term storage of the agent's data — sessions, configuration, providers, themes, and metadata. It provides backup, restore, and migration capabilities.

## Stored Data

### Sessions
- Full message history (user and AI messages)
- Tool call records (invocations, results, timings)
- Session metadata (title, created/updated timestamps, model used)
- Branch/child relationships
- Scroll position, cursor state (for resumption)

### Providers & Models
- Provider definitions (name, URL, auth type, custom headers)
- Model definitions (name, provider ref, capabilities, pricing)
- Authentication data (API keys — stored securely)
- Priority and routing rules

### Configuration
- All configuration layers (global, user, workspace, session)
- User preferences and settings
- Permission policies
- Custom slash commands
- Custom agent profiles

### Themes
- Built-in theme definitions
- User-created themes
- Active theme selection per-session
- Terminal color mappings

### Skills & Knowledge
- Skill activation state per-session
- Custom skill definitions
- Custom knowledge entries
- Skill preferences and configuration

### Metadata
- Session statistics (token usage, cost, duration)
- Audit logs (permission decisions, tool executions)
- Installation and update history

## Storage Engine

### Design
- Data is stored in a structured, queryable format
- The storage layer is abstracted — the schema works with any backend
- All operations are transactional (atomic commits)
- Concurrent access is handled with locking

### Schema
- Tables/collections are organized by data domain
- Each record has: ID, type, data blob, timestamps, version
- Cross-references use unique IDs
- Schema version is tracked for migrations

### Constraints
- Maximum storage size is configurable
- Old data can be auto-purged by retention policy
- Individual records are bounded in size
- Storage failures are non-fatal — the agent degrades gracefully

## Backup & Restore

### Backup
The agent can create full or partial backups:
- **Full backup** — all data (sessions, providers, config, themes)
- **Partial backup** — selected data domains
- **Snapshot backup** — point-in-time capture

Backup format:
- Single portable file (compressed)
- Human-readable content (for inspection without the agent)
- Encrypted (optional, for sensitive data)
- Self-contained (includes schema version for compatibility)

Backup operations:
- `create` — generate a backup file
- `list` — show available backups with timestamps and sizes
- `info` — inspect backup contents without restoring

### Restore
Restoring from a backup:
- **Full restore** — replaces all current data
- **Selective restore** — restore specific domains
- **Preview** — show what would be restored without applying
- **Merge** — merge backup data with current data (with conflict resolution)

Conflict resolution during merge:
- Keep current (skip backup version)
- Use backup (overwrite current)
- Keep both (with renamed duplicate)
- Review each conflict manually

### Backup Automation
- Scheduled backups (configurable interval)
- Pre-update automatic backup
- On-demand backup before destructive operations

## Migration

When the data schema changes between versions:
1. The current schema version is checked on startup
2. If the stored version doesn't match the expected version, migration runs
3. Migration scripts are versioned and run in sequence
4. Each migration is transactional (rollback on failure)
5. A pre-migration backup is automatically created
6. Migration progress is displayed
7. On failure, the previous version's data is preserved

## Data Integrity

- All writes are validated before commit
- Checksums verify data integrity on read
- Corrupted records are isolated (not loaded) and reported
- Repair tools rebuild corrupted data from available backups

## Key Design Decisions

- All data is stored locally (no cloud dependency for core operations)
- Backup format is portable and human-readable
- Migration is automatic but careful — never destructive
- The storage schema is versioned and forward-compatible
- Storage is abstracted — the schema works with various backends
- Backup before destructive operations is the default
- Data integrity checks happen at read time, not just write time
