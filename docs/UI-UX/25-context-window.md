# 25-context-window.md

# Context Window
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Context Window system used throughout the Mobile AI CLI.

The Context Window manages how conversations, files, tool outputs, workspace information, memory, and AI instructions are collected, displayed, optimized, and controlled before being sent to the language model.

The system must provide users with visibility and control while maintaining performance on Android Termux.

---

# Design Goals

The Context Window system must be

- Mobile First
- Terminal Native
- Transparent
- Efficient
- Token Optimized
- Memory Aware
- Privacy Focused
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

The AI should receive the most relevant information, not unlimited information.

Context management balances

```
Accuracy

+

Performance

+

Token Efficiency

+

User Control
```

---

# Context Window Concept

The Context Window is the active information space available to the AI model.

It may contain

- System Instructions
- User Messages
- Assistant Responses
- Tool Results
- Files
- Workspace Data
- Memory
- Metadata

---

# Context Pipeline

```
Input Sources

↓

Context Collector

↓

Context Analyzer

↓

Context Optimizer

↓

Prompt Builder

↓

LLM Request
```

---

# Context Sources

Supported sources

- Conversation History
- Current Prompt
- Selected Files
- Workspace Index
- Tool Results
- MCP Data
- User Preferences
- Project Metadata

---

# Context Priority

Priority order

```
1. Current User Request

2. System Instructions

3. Active Files

4. Recent Conversation

5. Tool Results

6. Workspace Information

7. Previous History
```

---

# Context Layers

The system uses multiple layers.

---

# Layer 1: System Context

Contains

- Application rules
- Security policies
- AI behavior settings

Highest priority.

---

# Layer 2: User Context

Contains

- Current request
- User preferences
- Selected options

---

# Layer 3: Conversation Context

Contains

- Previous messages
- Previous responses
- Current session data

---

# Layer 4: Workspace Context

Contains

- Project files
- Folder structure
- Dependencies
- Configuration

---

# Layer 5: Tool Context

Contains

- Command output
- API responses
- Search results

---

# Layer 6: Temporary Context

Contains short-lived information.

Examples

- Current operation
- Temporary variables
- Active selections

---

# Context Size

The system monitors

- Token usage
- Character size
- File size
- Memory usage

---

# Token Budget

Each request has a token allocation.

Example

```
System

↓

User

↓

Files

↓

Tools

↓

Response Space
```

---

# Context Compression

When context becomes large

The system may compress

- Old messages
- Repeated information
- Large outputs

---

# Compression Strategy

Priority

Keep

- Important decisions
- User requirements
- Active files

Compress

- Old explanations
- Repeated outputs
- Unused history

---

# Conversation Memory

Supported.

Stores useful information between sessions.

Memory must be controlled by user settings.

---

# File Context

Users may attach files manually.

Supported

- Single File
- Multiple Files
- Folder Selection

---

# Automatic File Context

Optional.

The system may automatically include relevant files from workspace search.

---

# Workspace Awareness

The Context Window may include

- Project structure
- Framework information
- Dependencies
- Configuration

---

# Tool Result Handling

Large tool outputs should not consume the entire context.

Use

- Summaries
- References
- Selected sections

---

# MCP Context

MCP responses are included only when relevant.

Unused MCP data must not enter the context.

---

# Context Preview

Users may view current context.

Example

```text
Context

Files

src/App.tsx

package.json


Messages

12


Tokens

8400 / 32000
```

---

# Context Indicator

Status Bar may display

```
Context 45%
```

---

# Context Warning

When context becomes large

Display

```text
Context limit approaching.
```

---

# Context Overflow

When maximum size is reached

Options

- Compress
- Remove old messages
- Start new session

---

# Context Search

Users may search within active context.

---

# Context Pinning

Users may pin important information.

Pinned data remains higher priority.

---

# Context Removal

Users may remove

- Files
- Messages
- Tool outputs

---

# Privacy

Context must protect

- Private files
- Credentials
- Secrets
- Tokens

---

# Security Rules

Never automatically include

- Password files
- Environment secrets
- Private keys
- Sensitive credentials

---

# Performance

Context processing should be incremental.

Avoid rebuilding the entire context unnecessarily.

---

# Streaming Integration

Context preparation occurs before streaming begins.

During streaming

Current context remains stable.

---

# Tool Integration

Before executing tools

The system evaluates required context.

Only necessary information is provided.

---

# Error Handling

Context failure examples

```
Unable to load file context.

Context size exceeded.

Invalid workspace data.
```

---

# Safe Area

Context information appears only inside approved UI areas.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Keyboard Behavior

Opening keyboard does not reset context state.

Draft messages preserve context.

---

# Restrictions

Never

- Send unnecessary files
- Expose private information
- Hide context usage
- Ignore user selection
- Fill context with irrelevant data
- Block user interaction

---

# Example Context View

```text
Context Window

System

Active


Files

3 Selected


Messages

18


Tools

2 Results


Tokens

12000 / 32000
```

---

# Context Optimization Example

Before

```text
1000 lines of logs
```

After

```text
Error summary

5 relevant lines
```

---

# Context Checklist

Every Context Window system must

- Manage token usage
- Prioritize relevant data
- Protect privacy
- Support compression
- Support file context
- Support tool context
- Support memory
- Provide visibility
- Optimize performance
- Work on Android Termux

---

# Core Rules

1. Current user intent has highest priority.
2. Only relevant information enters the context.
3. Large data must be compressed.
4. Sensitive information is protected.
5. Users control important context.
6. Context usage remains transparent.
7. Tool outputs are optimized.
8. Workspace data is included carefully.
9. Token limits are monitored.
10. Optimize for mobile environments.

---

# Summary

The Context Window system manages the information flow between the user, workspace, tools, and AI model. By prioritizing relevant data, optimizing tokens, protecting privacy, and providing user visibility, it creates an efficient AI coding environment optimized for Android Termux and mobile-first development workflows.