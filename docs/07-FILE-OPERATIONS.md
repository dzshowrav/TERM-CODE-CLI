# File Operations

## Overview

File operations are the most fundamental tools in the agent's toolkit. The agent reads files to understand code, writes files to make changes, edits files to apply modifications, and searches files to find relevant code.

## Core Operations

### Read
- Read a file by path with optional line range
- Returns full content or specified range
- Detects file encoding (UTF-8, UTF-16, Latin-1, binary detection)
- Large files warn before reading (>1MB)
- Supports reading from references (@-mentions of file paths)

### Write
- Write content to a file (full overwrite)
- Creates parent directories if they don't exist
- Validates write target is within workspace (security boundary)
- Confirms before overwriting existing files (unless auto-approved by policy)
- Supports binary files (images, compiled binaries) with size warnings

### Edit
- Apply surgical edits to existing files
- Patterns for editing:
  - **Line-based** — replace line range with new content
  - **Text-based** — find exact text and replace
  - **Regex-based** — find pattern and replace (with confirmation)
  - **Insert** — insert content at specific line
  - **Append** — append content to end of file
- Edits are validated before application:
  - If find pattern is not found, the tool reports failure with context
  - If multiple matches exist, the user or model disambiguates
- Edits can be previewed (show diff before applying)

### Delete
- Delete files or empty directories
- Confirms before deletion (mandatory, no auto-approve for delete)
- Recursive deletion with explicit confirmation
- Moves to trash/recycle bin if available (undoable)

### List Directory
- List directory contents
- Show file types, sizes, modification times
- Support filtering by glob pattern
- Support recursion with depth limit

### Glob
- Find files matching a glob pattern
- Case-insensitive option
- Exclude patterns (e.g., node_modules, .git)
- Return relative paths from workspace root
- Results are sorted and deduplicated

## Git Snapshots

Before any write or edit operation:
1. Optionally create a git snapshot (stash or temporary commit)
2. The snapshot is tagged with the tool call ID for identification
3. If the edit causes issues, the snapshot can be restored

The snapshot system:
- Does NOT commit to the user's branch
- Uses git stash, temporary branches, or worktrees
- Preserves the working tree state before the change
- Is cleaned up after a configurable number of snapshots

## Diff Computation

After every file modification, a diff is computed:
- Unified diff format (contextual, with line numbers)
- Diff against the previous snapshot (or HEAD if no snapshot)
- Word-level and line-level highlighting
- Language-aware diff for structured files
- Diff metadata: files changed, insertions, deletions

The diff is:
- Displayed to the user in a formatted view (see Display Rendering)
- Stored in the session for later reference
- Available for the model to reference in subsequent responses

## File Watcher

The agent can watch files for changes:
- Detect external modifications (user edited a file outside the agent)
- Trigger re-read of watched files on change
- Update displayed content if the file is currently on screen
- Used for: following logs, monitoring config files, watching build output

## Large File Handling

Files above a size threshold (configurable, default 1MB):
- Are not read in full — only the first N lines and last N lines
- The model is informed of the total size and that it was truncated
- The model can request specific line ranges
- Write/Edit operations on large files are checked for disk space

## Key Design Decisions

- Every file modification has an undo path (via git snapshots or trash)
- File operations validate against workspace boundaries (no escaping the project root)
- Edits use exact-match patterns to avoid ambiguity (AI models make mistakes on line numbers)
- Large files are handled incrementally, never loaded entirely into memory
- File operations are synchronous and blocking — the agent waits for completion before proceeding
- Binary file write is allowed but read returns metadata only (not content)
