# Process Lifecycle

## Overview

The process lifecycle governs how the agent starts, runs, and shuts down. It handles initialization, signal management, crash recovery, self-update, and self-uninstall.

## Startup Sequence

When the agent process starts:

### Phase 1: Environment Detection
1. Detect operating system and terminal capabilities
2. Detect terminal dimensions and color support
3. Read environment variables (paths, config locations, auth tokens)
4. Check for required dependencies and external tools
5. Detect workspace (current directory, git state)

### Phase 2: Configuration Loading
1. Load default configuration
2. Load global user configuration
3. Load workspace-specific configuration
4. Merge configuration layers
5. Validate configuration (report errors without crashing)

### Phase 3: Storage Initialization
1. Initialize storage engine
2. Run schema migrations (if any)
3. Load session index
4. Restore last active session (if configured)
5. Verify data integrity

### Phase 4: Resource Loading
1. Load theme
2. Initialize skill system
3. Connect to MCP servers
4. Register built-in tools
5. Initialize AI provider connections

### Phase 5: UI Initialization
1. Compute layout based on terminal size
2. Render initial display (status bar, welcome message, input area)
3. Set up input handling
4. Set up signal handlers
5. Set up auto-save timer

### Phase 6: Ready
1. Display ready indicator
2. Wait for user input
3. Start idle timer

Startup is progressive — if any phase fails, the agent falls back to a minimal operational state and reports the error.

## Runtime Loop

After startup, the agent runs an event loop:
1. Wait for input
2. Process input (keystroke, paste, command, AI message)
3. Execute AI communication or tools
4. Update state
5. Re-render UI
6. Go back to step 1

The loop is non-blocking — input is always accepted, even during AI response generation.

## Shutdown Sequence

### Phase 1: Graceful Initiation
Shutdown can be triggered by:
- User command (/exit, Ctrl+D, Ctrl+C)
- System signal (SIGTERM, SIGINT)
- Error condition (irrecoverable error)
- Self-update (shutdown before restart)

### Phase 2: Save State
1. Save current session immediately
2. Save session index
3. Save configuration (if modified)
4. Store cursor position and scroll state

### Phase 3: Cleanup
1. Abort any in-flight AI requests
2. Kill any running tools/shell commands (with grace period)
3. Disconnect MCP servers
4. Close storage engine
5. Release resources

### Phase 4: Terminal Cleanup
1. Restore terminal settings (cursor, raw mode, echo)
2. Clear or preserve screen content (configurable)
3. Show goodbye message (optional)
4. Exit with appropriate status code

## Signal Handling

### SIGINT (Ctrl+C)
- First SIGINT: abort current operation (AI response, tool execution)
- Second SIGINT within 1 second: show exit confirmation
- Third SIGINT: force exit immediately

### SIGTERM
- Start graceful shutdown sequence
- If shutdown doesn't complete within 5 seconds, force exit

### SIGHUP
- Terminal closed unexpectedly
- Save session immediately
- Exit with save confirmation message

### SIGPIPE
- Output pipe closed
- Continue running (don't crash)

### Triple SIGTERM Threshold
If 3 rapid termination signals arrive within a short window:
- Skip graceful shutdown
- Save critical state only (session, config)
- Force exit immediately
- On next startup, detect abnormal termination and offer recovery

## Crash Recovery

### Crash Detection
On next startup after abnormal exit:
1. Detect crash (no clean shutdown record)
2. Scan for partial session data
3. Check for backup availability
4. Report crash to user with details

### Recovery Options
1. **Restore last session** — load from auto-save
2. **Restore from backup** — load from latest backup
3. **Start fresh** — begin new session, archive corrupted data
4. **Diagnose** — show crash log for user inspection

### Crash Prevention
- State is saved frequently (auto-save every N seconds)
- Large operations checkpoint their progress
- Memory usage is bounded and monitored
- Unhandled errors are caught at the top level

## Self-Update

### Update Flow
1. User runs `/update` command (or `latest`, `check`)
2. System checks current version
3. System checks latest version (from configured source — git, package manager)
4. If update is available:
   - **Check mode**: display version info and changelog
   - **Latest mode**: download and apply update
5. Download update
6. Verify integrity
7. Auto-backup current installation
8. Apply update
9. Restart the agent (user sees brief restart, session resumes)

### Rollback
- If update fails, auto-rollback to previous version
- User can manually rollback to a previous version
- Previous version's backup is preserved for N days

## Self-Uninstall

### Uninstall Flow
1. User runs `/uninstall`
2. Show warning: what will be deleted (session data, config, themes)
3. Show what will be preserved (if anything)
4. Offer backup option before proceeding
5. Require explicit confirmation (type "yes" or equivalent)
6. Create final backup
7. Remove all agent files
8. Clean up configuration directories
9. Show completion message
10. Exit

## Key Design Decisions

- Startup is progressive — the agent is usable even if some components fail
- SIGINT has a progressive response — first abort, then confirm, then force
- Triple SIGTERM protects against signal storms
- Crash recovery is automatic but transparent — the user knows what happened
- Self-update is non-destructive — backup is always created first
- Self-uninstall is careful — backup is offered, confirmation is explicit
- All cleanup operations have timeouts — the agent won't hang on shutdown
- Auto-save ensures minimal data loss in any failure scenario
