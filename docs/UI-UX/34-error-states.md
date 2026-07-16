# 34-error-states.md

# Error States
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Error State System used throughout the Mobile AI CLI.

The Error State System provides clear information when operations fail, explains what happened, prevents user confusion, and guides users toward possible solutions.

The system must handle failures gracefully without breaking the user workflow.

---

# Design Goals

The Error State System must be

- Mobile First
- Terminal Native
- Clear
- Actionable
- Secure
- Non-Blocking
- Developer Friendly
- Performance Optimized

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Errors are information, not dead ends.

Every error should answer:

- What happened?
- Why did it happen?
- What can the user do next?

---

# Error Architecture

```
Operation

↓

Failure Detected

↓

Error Classification

↓

Error Display

↓

Recovery Action
```

---

# Error Categories

Supported

- User Input Error
- Permission Error
- File Error
- Network Error
- Model Error
- Tool Error
- Configuration Error
- System Error
- Security Error

---

# Error Severity Levels

Supported

```
Info

Warning

Recoverable Error

Critical Error
```

---

# Info Error

Minor issue that does not stop workflow.

Example

```text
Using default configuration.
```

---

# Warning Error

Potential issue.

Example

```text
Large file may slow processing.
```

---

# Recoverable Error

Operation failed but recovery is possible.

Example

```text
Connection failed.

Retry
```

---

# Critical Error

Application-level failure.

Example

```text
Unable to start workspace.
```

---

# Error Display Location

Errors may appear in:

- Conversation Area
- Toast System
- Dialog System
- Status Bar
- Tool Execution View

---

# Inline Error

Used for local failures.

Example

```text
File not found.

src/App.tsx
```

---

# Toast Error

Used for temporary notifications.

Example

```text
Command failed.
```

---

# Dialog Error

Used for important failures requiring attention.

Example

```text
Operation Failed

Unable to connect provider.

Close
```

---

# Status Bar Error

Used for ongoing system problems.

Example

```text
MCP Offline
```

---

# Error Structure

Every error should contain

- Title
- Description
- Cause (optional)
- Solution (optional)
- Action

---

# Error Title

Must be short.

Good

```text
File Not Found
```

Bad

```text
Something went wrong
```

---

# Error Description

Explains the problem.

Example

```text
The selected file no longer exists.
```

---

# Cause Information

Optional.

Example

```text
The file was moved or deleted.
```

---

# Recovery Action

Examples

```
Retry

Open Settings

Change Path

Cancel
```

---

# User Input Errors

Examples

- Invalid command
- Empty required field
- Wrong format

---

# Example

```text
Invalid Command

Unknown command:
/abc

Try /help
```

---

# Permission Errors

Examples

- File access denied
- Tool permission missing

---

# Example

```text
Permission Required

Allow file access?
```

---

# File Errors

Examples

- Missing file
- Corrupted file
- Unsupported format

---

# Example

```text
Cannot Open File

Unsupported format.
```

---

# Network Errors

Examples

- No connection
- Timeout
- Provider unavailable

---

# Example

```text
Connection Failed

Retry
```

---

# Model Errors

Examples

- Model unavailable
- Invalid API configuration
- Context limit exceeded

---

# Example

```text
Model Unavailable

Select another model.
```

---

# Tool Errors

Examples

- Command failed
- MCP failure
- Execution blocked

---

# Example

```text
Tool Failed

Terminal execution denied.
```

---

# Configuration Errors

Examples

- Invalid settings
- Missing configuration
- Broken environment

---

# Example

```text
Configuration Error

API provider not configured.
```

---

# System Errors

Examples

- Memory issue
- Internal failure
- Unexpected crash

---

# Example

```text
System Error

Restart required.
```

---

# Security Errors

Examples

- Unauthorized access
- Invalid credentials

---

# Example

```text
Access Denied

Invalid credentials.
```

---

# Recovery Flow

```
Error

↓

Explanation

↓

Suggested Action

↓

Retry / Fix / Cancel
```

---

# Retry System

Supported.

Retry options

- Immediate Retry
- Retry After Fix
- Manual Retry

---

# Error Logging

Errors should be logged internally.

Log contains

- Error Type
- Timestamp
- Operation
- Stack Information

Sensitive data must be removed.

---

# Error Reporting

Optional.

Users may submit error reports.

Reports must exclude private data.

---

# Error History

Users may view recent errors.

Example

```text
Recent Errors

Network Timeout

File Access Denied
```

---

# Streaming Error Handling

If streaming fails

Display partial response.

Example

```text
Response interrupted.

Retry generation.
```

---

# Tool Error Handling

Tool failures must show

- Tool Name
- Failed Step
- Reason

---

# Keyboard Behavior

Errors never block typing.

Users may continue working.

---

# Safe Area

Errors respect

- Command Input
- Status Bar
- Keyboard
- Display boundaries

---

# Accessibility

Support

- High Contrast
- Screen Readers
- Text Descriptions

Never use only color to indicate errors.

---

# Performance

Error rendering should be instant.

Avoid expensive recovery operations automatically.

---

# Security Rules

Never expose

- API Keys
- Passwords
- Tokens
- Private Environment Data

---

# Error Prevention

The system should prevent errors through

- Validation
- Confirmation
- Permission Checks
- Clear Instructions

---

# Restrictions

Never

- Show meaningless errors
- Hide recovery options
- Crash after errors
- Lose user work
- Expose sensitive data
- Use technical stack traces as the only message

---

# Example Error Flow

```text
User

Run Command

↓

Error

Command Failed

↓

Reason

Permission denied

↓

Action

Allow Permission
```

---

# Example File Error

```text
Open File

↓

File Not Found

↓

Search Again
```

---

# Example Model Error

```text
Generate Response

↓

Context Limit Exceeded

↓

Compress Context
```

---

# Error Checklist

Every Error System must

- Explain failures
- Provide recovery options
- Protect sensitive information
- Support logging
- Preserve user workflow
- Handle streaming failures
- Support accessibility
- Remain terminal-native
- Work on Android Termux
- Stay user friendly

---

# Core Rules

1. Errors must explain the problem.
2. Recovery actions must be clear.
3. User work must never be lost.
4. Sensitive information must stay hidden.
5. Errors should not block normal workflow.
6. Partial results should be preserved.
7. Technical details are optional.
8. Validation should prevent failures.
9. Error messages must be readable.
10. Optimize for mobile CLI usage.

---

# Summary

The Error State System creates a reliable failure-handling experience inside the Mobile AI CLI. By providing clear explanations, recovery actions, secure reporting, and non-blocking behavior, it ensures that errors become useful guidance instead of workflow interruptions while maintaining Android Termux compatibility.