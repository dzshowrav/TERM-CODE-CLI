# Session Sharing

## Overview

The session sharing system enables users to export, import, and share sessions with others. Sessions contain the full conversation history, tool call records, and configuration.

## Export Format

Exported sessions use a portable format:

### Format Structure
- **Metadata** — session title, timestamps, model used, agent profile, duration, token count
- **Messages** — ordered list of all messages (user + assistant), preserving roles and content
- **Tool calls** — each tool invocation with arguments, results, duration, status
- **Configuration** — session-level configuration at the time of export
- **Workspace context** — anonymized project info (project type, frameworks — no file contents)
- **Sensitive data** — API keys and credentials are never included (stripped on export)

### Format Properties
- Human-readable (viewable in any text editor)
- Diff-friendly (changes between exports are trackable)
- Compressed option (for sharing larger sessions)
- Self-contained (all references resolved)
- Signed (optional, for authenticity verification)

## Export Operations

### Full Export
- Exports the complete session including all messages, tool calls, and metadata
- File size warning for very large sessions
- Progress indicator during export

### Selective Export
- Export selected messages (by range or by selection)
- Export tool calls only (without conversation)
- Export configuration only
- Export without sensitive information

### Auto-Export
- Export on session close (configurable)
- Periodic export during long sessions
- Export before major operations (update, uninstall)

## Import Operations

### Full Import
- Imports a complete session as a new session or merged into an existing one
- Validates format and integrity before importing
- Reports any issues (missing fields, format errors)

### Merge Import
- Merges an imported session into the current session
- Messages are appended to the current conversation
- Configuration can be imported or ignored
- Duplicate detection (same session ID or content)

### Read-Only Import
- Opens an exported session in read-only mode
- The user can browse the conversation but not modify it
- Useful for reviewing shared sessions

## Share Mechanisms

### File-Based Sharing
- Export to a file, share the file (email, messaging, cloud storage)
- Recipient imports the file
- No server required

### URL-Based Sharing
- Session is uploaded to a sharing service (optional, opt-in)
- A URL is generated for access
- URL can be password-protected
- Optional expiration date
- The sharing service doesn't store sensitive data

### Clipboard Sharing
- Export to clipboard (for pasting into chat, email, etc.)
- Format is compact (suitable for pasting into conversations)
- Can include a summary instead of full content

## Privacy & Security

- All sensitive data (API keys, credentials, tokens) is stripped on export
- The user can review what will be shared before exporting
- Workspace file paths can be anonymized
- User can redact specific messages before export
- URL-based sharing is opt-in (never automatic)
- Exported files can be encrypted with a password

## Collaboration Model

Imported sessions enable collaboration:
- A teammate can see exactly what the agent did and why
- Tool call results are visible — no "trust me" moments
- Configuration is preserved — the same settings can be replicated
- The importer can continue the conversation from where it left off

## Key Design Decisions

- Export is always user-initiated — sessions are never shared automatically
- Sensitive data is stripped by default, not opt-in
- URL-based sharing is optional and doesn't require an account
- The format is human-readable for transparency
- Read-only import prevents accidental modification of shared sessions
- Merge import enables building on others' work
- File-based sharing works without any server infrastructure
- Session sharing is a collaboration feature, not a surveillance feature
