# Permission & Security

## Overview

The permission system governs what actions the agent can take and what user approval is required. It's a layered system that balances safety with efficiency.

## Permission Tiers

Every tool execution is classified into one of four tiers:

### Tier 1: Always Allow
- Actions that are purely informational and read-only
- Examples: read file, list directory, glob search, grep search
- Never requires approval
- User can also add custom tools to this tier

### Tier 2: Allow Once
- Actions that modify state but are well-understood
- Examples: write a new file, edit an existing file, run a non-destructive command
- Requires approval the first time in a session, then auto-approved for that session
- The approval can be "allow for this session" or "allow once"

### Tier 3: Always Ask
- Actions that modify state in potentially impactful ways
- Examples: delete files, install packages, modify configuration, run tests
- Requires approval every time
- The full command/tool call is displayed for review

### Tier 4: Require Confirmation
- Actions that could cause data loss or security impact
- Examples: recursive delete, format disk, modify git history, network operations
- Requires explicit approval with additional confirmation
- These actions show a warning about potential consequences

## Policy System

Policies define automatic behaviors for permission decisions:

### Policy Rules
- **Pattern-based** — match tool names, file paths, command patterns
- **Time-based** — auto-approve for a duration
- **Scope-based** — apply to specific sessions or workspace
- **Count-based** — auto-approve after manually approving N times

### Policy Examples
- "Auto-approve `npm install` if package.json is unchanged"
- "Never auto-approve anything containing `rm -rf`"
- "Auto-approve all edits to test files"
- "Require confirmation for any network-connected tool"
- "Auto-approve for the next 5 minutes"

### Policy Evaluation
1. Policies are evaluated in priority order
2. First matching policy wins
3. If no policy matches, use the tool's default tier
4. The user is shown which policy made the decision

## Approval Prompt

When approval is needed, the user sees:
- **Tool name** and description
- **Arguments** being passed (formatted for readability)
- **Context** — what triggered this tool call (the AI message that led to it)
- **Options** — allow, allow once, allow for session, allow always (for this tool), deny, deny permanently
- **Shortcuts** — keyboard shortcuts to quickly approve/deny
- **Timeout** — the prompt auto-denies if no response within a configurable period

## Auto Mode

When the user enables "auto mode":
- All Tier 1 and Tier 2 actions proceed without approval
- Tier 3 actions auto-approve after a 3-second warning (user can cancel)
- Tier 4 actions still require explicit confirmation
- Auto mode is visually indicated in the UI at all times
- The user can disable auto mode with a single keystroke at any time

## Silent Mode / Background

In silent mode:
- The agent runs without displaying intermediate tool calls
- Permission decisions follow policy only (no user prompts)
- Only errors and final results are displayed
- Used for automated/scripted operation

## Audit Trail

Every permission decision is logged:
- Tool name, arguments, timestamp
- Decision (approved, denied, auto-approved)
- Policy rule that made the decision (if applicable)
- User's selected option (allow once, allow session, allow always)
- Session ID for correlation

The audit trail is:
- Viewable per-session
- Exportable for review
- Privacy-preserving (no sensitive argument values)

## Key Design Decisions

- Default-deny for any state-changing action — the user must explicitly opt in
- Policy system allows the agent to be tuned from "baby sit" to "fully autonomous"
- Approval is a single keystroke — low friction for genuine usage
- Auto mode is visually unmistakable — the user always knows when it's on
- Audit trail provides accountability without being noisy
- Policies are evaluated before execution, never retroactively
- Sensitive operations (Tier 4) are never auto-approved, even in auto mode
