# 33-loading-states.md

# Loading States
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Loading State System used throughout the Mobile AI CLI.

The Loading State System provides clear feedback when the application is performing background operations, waiting for responses, processing data, loading resources, or executing tools.

The system must keep users informed without blocking interaction.

---

# Design Goals

The Loading State System must be

- Mobile First
- Terminal Native
- Lightweight
- Informative
- Non-Blocking
- Streaming Compatible
- Battery Friendly
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

Loading states are communication, not decoration.

Users should understand:

- What is loading
- Why it is loading
- Whether progress is happening
- What they can do meanwhile

---

# Loading Architecture

```
Operation Started

↓

Loading State Created

↓

Progress Updates

↓

Operation Completed

↓

Success / Error State
```

---

# Loading Categories

Supported

- AI Generation Loading
- Tool Execution Loading
- File Loading
- Workspace Loading
- Model Loading
- Network Loading
- Application Loading
- Background Loading

---

# Loading State Types

Supported states

```
Idle

Queued

Preparing

Loading

Processing

Streaming

Completing

Completed

Failed

Cancelled
```

---

# Idle State

No operation is running.

Example

```text
Ready
```

---

# Queued State

Operation waiting for execution.

Example

```text
Waiting...
```

---

# Preparing State

Initial setup is happening.

Examples

```text
Preparing Context...

Loading Workspace...
```

---

# Loading State

Resource is being loaded.

Examples

```text
Loading Model...
```

---

# Processing State

Active computation.

Examples

```text
Analyzing Files...
```

---

# Streaming State

Continuous output generation.

Example

```text
Generating Response...
```

---

# Completing State

Final steps.

Example

```text
Finalizing...
```

---

# Completed State

Operation finished successfully.

Example

```text
Completed
```

---

# Failed State

Operation failed.

Example

```text
Failed

Connection Error
```

---

# Cancelled State

User stopped the operation.

Example

```text
Cancelled
```

---

# Display Locations

Loading states may appear in:

- Conversation Area
- Status Bar
- Progress UI
- Tool Execution View
- Dialog System

---

# Chat Loading

Used during AI response generation.

Example

```text
AI is thinking...

⠋
```

---

# Streaming Loading

During token generation.

Example

```text
Generating response...

The architecture is...
```

---

# Tool Loading

Example

```text
Running Tool

filesystem.search
```

---

# File Loading

Example

```text
Reading File

src/App.tsx
```

---

# Workspace Loading

Example

```text
Indexing Workspace

125 / 500 Files
```

---

# Model Loading

Example

```text
Loading GPT-5

Initializing...
```

---

# Network Loading

Example

```text
Connecting Provider...
```

---

# Application Startup Loading

Example

```text
Starting Mobile AI CLI

Loading Configuration
```

---

# Loading Indicator Types

Supported

- Spinner
- Progress Bar
- Status Text
- Skeleton Output
- Step Indicator

---

# Spinner

Used for unknown duration operations.

Example

```text
Processing

⠋
```

---

# Spinner Rules

Spinner must

- Use low CPU
- Remain readable
- Stop after completion

---

# Progress Bar

Used for measurable operations.

Example

```text
████████░░ 80%
```

---

# Status Text

Always recommended.

Example

```text
Scanning project files
```

---

# Step Indicator

Used for multi-step operations.

Example

```text
Step 2/5

Installing Dependencies
```

---

# Skeleton Output

Optional.

Used when structure is known before content.

Example

```text
Loading Messages...
```

---

# Multiple Loading Operations

Supported.

Example

```text
AI Response

+

File Indexing

+

MCP Sync
```

---

# Priority System

Priority order

```
Critical

↓

User Requested

↓

Foreground

↓

Background
```

---

# Foreground Loading

Visible immediately.

Examples

- AI Response
- Command Execution

---

# Background Loading

Shown minimally.

Examples

- Cache Update
- Index Refresh

---

# Cancel Action

Long operations should support cancellation.

Example

```text
Indexing Workspace

Cancel
```

---

# Auto Timeout

Optional.

Long-running operations may show timeout information.

Example

```text
Taking longer than expected
```

---

# Keyboard Behavior

Loading states never block typing.

Users may continue writing messages.

---

# Safe Area

Loading indicators must respect

- Command Input
- Status Bar
- Keyboard
- Display boundaries

---

# Streaming Integration

Loading transitions smoothly into streaming.

Example

```
Thinking...

↓

Generating...

↓

Response Text
```

---

# Performance

Loading updates should

- Render incrementally
- Avoid full redraw
- Minimize CPU usage

---

# Battery Optimization

Reduce animation frequency when

- Battery saver enabled
- Device performance is limited

---

# Accessibility

Support

- Text alternatives
- High Contrast
- Reduced Motion

Never rely only on animation.

---

# Error Handling

If loading fails:

Display clear message.

Example

```text
Loading Failed

Unable to load model.
```

---

# Security

Never display sensitive loading information.

Do not expose:

- API tokens
- Internal secrets
- Private paths

---

# Restrictions

Never

- Show fake progress
- Freeze the interface
- Hide operation status
- Loop forever without information
- Block user input

---

# Example AI Flow

```text
User Prompt

↓

Thinking...

↓

Generating Response...

↓

Response Complete
```

---

# Example File Flow

```text
Opening File

↓

Reading Content

↓

File Loaded
```

---

# Example Tool Flow

```text
Running Command

↓

Processing Output

↓

Completed
```

---

# Loading Checklist

Every Loading System must

- Show current state
- Support multiple operations
- Support cancellation
- Support streaming
- Avoid blocking
- Respect safe areas
- Protect privacy
- Optimize performance
- Work on Android Termux
- Remain terminal-native

---

# Core Rules

1. Loading must always communicate progress.
2. Unknown duration uses indicators.
3. Known duration uses progress.
4. User interaction remains available.
5. Streaming transitions smoothly.
6. Never display fake progress.
7. Cancel long operations when possible.
8. Protect sensitive information.
9. Keep animations lightweight.
10. Optimize for mobile CLI usage.

---

# Summary

The Loading State System provides transparent feedback during every background and foreground operation inside the Mobile AI CLI. By combining status messages, progress indicators, streaming transitions, and non-blocking behavior, it creates a responsive terminal-native experience optimized for Android Termux.