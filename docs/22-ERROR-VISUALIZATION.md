# Error & Notification System

## Overview

The error and notification system handles how the agent communicates problems, warnings, and informational messages to the user. It covers error display, silent failures, notifications, and logging.

## Error Types

### User-Facing Errors
Errors the user needs to know about and act on:
- **Tool errors** — file not found, permission denied, command failed
- **AI errors** — provider unreachable, rate limited, invalid response
- **Configuration errors** — invalid config, missing required fields
- **Storage errors** — corrupt data, write failure, migration failure

### Non-User-Facing Errors
Errors the system handles internally:
- **Transient failures** — network timeouts (auto-retried)
- **Expected failures** — file already exists (handled by tool logic)
- **Degraded operations** — optional feature unavailable (agent continues)
- **Background errors** — MCP server disconnected (auto-reconnect attempted)

## Error Display

### Inline Error Messages
- Displayed in the conversation as system messages
- Include: error type, description, suggested action
- Styled with the error color from the theme
- Include error codes for reference

### Error Cards
For tool execution errors:
- Compact card showing tool name and error
- Expandable to show full error details
- Retry button (re-run the tool)
- Dismiss option

### Toast Notifications
For transient information:
- Appear briefly at the top or bottom of the screen
- Auto-dismiss after a configurable timeout
- Types: success (green), warning (yellow), error (red), info (blue)
- Stacks of toasts are shown in order

### Fatal Error Screen
For unrecoverable errors:
- Full-screen error display
- Error code and description
- Suggested recovery steps
- Option to restart or exit
- Crash log reference

## Silent Failures

Some failures are handled silently (without user disruption):
- Auto-retry of network requests (up to 3 attempts)
- Graceful degradation (optional feature disabled, core functionality continues)
- Background task failure (MCP reconnection, index rebuild)
- Minor validation warnings (format recommendation, deprecation notice)

Silent failures are always logged. The user can review them with `/logs` or `/debug`.

## Notification System

### Types
- **Toast** — brief, auto-dismissing popup at screen boundary
- **Status bar alert** — persistent indicator in the status bar
- **Terminal bell** — audible alert (configurable)
- **Desktop notification** — system-level notification (configurable)
- **Title update** — terminal title change (configurable)

### Notification Events
- Tool execution complete (with result or error)
- Background task finished
- Update available
- Provider disconnected/reconnected
- Session auto-saved
- Permission request pending
- Long-running operation progress

### Notification Configuration
The user can configure:
- Which events trigger notifications
- Which notification type for each event
- Timeout duration for toasts
- Quiet hours (no non-critical notifications)

## Logging

### Log Levels
- **ERROR** — unrecoverable errors, unexpected failures
- **WARN** — non-critical issues, deprecated usage
- **INFO** — major state changes, lifecycle events
- **DEBUG** — detailed operation information
- **TRACE** — verbose debugging (tool call arguments, AI response raw data)

### Log Storage
- Logs are stored in a circular buffer (fixed size, oldest entries removed)
- Logs persist across sessions (in the storage engine)
- Logs are never automatically deleted (user configurable)
- Sensitive data (API keys, file contents) is redacted from logs

### Log Viewing
- `/logs` command shows recent logs
- Filter by level, time range, component
- Export logs for bug reports
- Real-time log tail mode

## Error Recovery Suggestions

Every error message includes a suggested action:
- "File not found" → "Create the file? Check the path? Search for the file?"
- "Provider unreachable" → "Check network connection? Try a different provider? View provider status?"
- "Invalid configuration" → "Open settings to fix? Reset to defaults?"

Suggestions are executable — the user can press a key to execute the suggestion directly.

## Key Design Decisions

- Errors are contextual — the user sees what went wrong and why
- Non-user-facing errors are logged but not displayed
- Toast notifications are non-blocking — they don't interrupt the workflow
- Fatal errors are rare but have a clear recovery path
- Silent failures are always auditable
- Logging is privacy-aware (sensitive data redacted)
- Error messages include actionable suggestions, not just descriptions
- Notification configuration puts the user in control of interruptions
