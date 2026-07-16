# Session Lifecycle

## What Is a Session

A session is the fundamental unit of conversation. It represents a single thread of interaction between user and agent. Sessions are persistent — they survive restarts, crashes, and network interruptions.

## Session States

A session progresses through these states:
- **Dormant** — exists in storage, never activated
- **Active** — currently loaded and interactive
- **Paused** — user switched away, state preserved in memory
- **Archived** — saved to long-term storage, no longer in memory
- **Exported** — serialized for sharing or backup
- **Deleted** — removed from storage

## Session Operations

### Create
A new session starts with:
- A unique identifier (UUID or hash-based)
- Timestamp of creation
- Initial system prompt (core identity + workspace context)
- Empty message history
- Default configuration (inherited from global config + workspace config)

The user may optionally provide:
- A descriptive title or name
- An initial context file or reference
- A session template (predefined configuration)

### List
All sessions are listable with:
- Title (auto-generated from first message or user-named)
- Timestamp of last activity
- Preview of last message (truncated)
- Current state (active/paused/archived)
- Model/provider used

### Resume
Resuming a session:
- Loads full message history from storage
- Rebuilds the context window (applying compaction if needed)
- Reconnects to the AI provider if active
- Restores the scroll position and cursor state

### Branch / Child Sessions
A session can spawn children:
- **Child session** — inherits parent's history and context, diverges at branch point
- **Independent session** — starts fresh but references the parent
- **Fork** — creates an independent copy of the session at a checkpoint

Child sessions form a tree. The user can navigate up and down the tree. The tree structure is visible in the session listing.

### Export
Sessions export to a portable format containing:
- Full message history (both user and assistant messages)
- Tool call records (invocations, results, timings)
- Configuration at time of session
- Metadata (timestamps, model used, duration)
- Branch/child relationships

### Import
Imported sessions:
- Are validated for format correctness
- Can be merged into existing session tree
- Can be loaded as read-only (for reference)
- Preserve all metadata and tool call records

### Auto-Save
- The session auto-saves after every message exchange
- Auto-save also triggers on idle timeout (e.g., 30 seconds of no activity)
- Corrupted saves are detected and rolled back to last valid checkpoint

### Cleanup / Retention
- Old archived sessions may be auto-purged based on configurable retention policy
- The user is warned before bulk deletion
- Sessions can be individually protected from auto-cleanup

## Session Hierarchy

```
Global Defaults
  └─ Workspace Config
       └─ Session A (active)
            ├─ Session A.1 (child, paused)
            └─ Session A.2 (child, archived)
       └─ Session B (dormant)
```

## Key Design Decisions

- Sessions are the unit of context — loading a different session means switching context
- Session storage is portable across machines (sync via file sync or cloud)
- Session history is append-only — messages are never modified after storage
- Session IDs are content-addressable or UUID-based for uniqueness
