# 28-state-manager.md

# State Manager Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a State Manager?

A **State Manager** is the subsystem responsible for storing, updating, synchronizing, and distributing the current state of the AI Coding Agent.

It ensures every subsystem has a consistent view of the application's current status, including active sessions, workflows, conversations, tools, models, plugins, permissions, and workspace information.

Instead of each module maintaining its own independent state, the State Manager acts as the **single source of truth** for runtime state.

---

# Why State Manager?

Without State Manager

```
Module A

↓

Own State

Module B

↓

Different State

Module C

↓

Another State
```

Problems

- Inconsistent data
- Synchronization issues
- Race conditions
- Difficult debugging
- Duplicate state

---

With State Manager

```
Application

↓

State Manager

↓

Shared State

↓

All Components
```

---

# Goals

A production State Manager should provide

- Centralized state
- Reactive updates
- State synchronization
- State persistence
- State validation
- Undo/redo support
- Snapshot management
- Event notifications
- Thread-safe updates
- State recovery

---

# High-Level Architecture

```
              Application

                    │

                    ▼

             State Manager

                    │

      ┌─────────────┼─────────────┐

      ▼             ▼             ▼

 Store        Observer      Persistence

      ▼             ▼             ▼

 Snapshot     Events      Recovery

      └─────────────┼─────────────┘

                    ▼

             Application
```

---

# Folder Structure

```
src/

state/

    StateManager.ts

    StateStore.ts

    StateSnapshot.ts

    StatePersistence.ts

    StateObserver.ts

    StateValidator.ts

    StateSerializer.ts

    StateRecovery.ts

    StateHistory.ts

    StateEvents.ts

    StateMetrics.ts
```

---

# Core Components

## State Manager

Central controller.

Responsibilities

- Store state
- Update state
- Notify observers
- Recover state

---

## State Store

Stores

```
Current State

Metadata

Version

Timestamp
```

---

## State Snapshot

Captures a complete copy of the current state.

Useful for

- Undo
- Rollback
- Recovery
- Debugging

---

## State Persistence

Stores state on disk for restoration after restart.

---

## State Observer

Notifies interested components when state changes.

---

## State Validator

Checks

- Data integrity
- Schema compliance
- State consistency

---

## State Serializer

Converts

```
Runtime State

↓

Serialized Data

↓

Storage
```

and restores it when needed.

---

## State Recovery

Restores state after

- Crash
- Restart
- Failure
- Rollback

---

## State History

Maintains historical state changes.

Supports

- Undo
- Redo
- Auditing

---

# State Lifecycle

```
Create

↓

Update

↓

Validate

↓

Notify

↓

Persist

↓

Recover
```

---

# State Object

Contains

```
Session

Conversation

Workspace

Workflow

Models

Tools

Plugins

Permissions

Settings

Metadata
```

---

# State Types

Examples

```
Application State

Session State

Conversation State

Workspace State

Workflow State

Tool State

Provider State

Plugin State

Theme State
```

---

# State Updates

```
Component

↓

State Manager

↓

Validation

↓

Store

↓

Notify
```

---

# Reactive Updates

```
State Changed

↓

Observers

↓

UI Updated
```

No manual refresh required.

---

# Snapshot Flow

```
Current State

↓

Snapshot

↓

Stored

↓

Restore Later
```

---

# Undo / Redo

```
State A

↓

State B

↓

State C

↓

Undo

↓

State B
```

---

# Persistence Flow

```
Runtime State

↓

Serializer

↓

Disk

↓

Restart

↓

Restore
```

---

# Event Bus Integration

Common events

```
state:create

state:update

state:restore

state:snapshot

state:error

state:reset
```

---

# Session Manager Integration

Stores

```
Session Status

Session Variables

Temporary Data
```

---

# Conversation Manager Integration

Maintains

```
Messages

Context

Summaries

Branches
```

---

# Workflow Engine Integration

Tracks

```
Running Workflows

Task Progress

Execution State
```

---

# Workspace Indexer Integration

Stores

```
Indexed Workspace

Open Files

Selected Project
```

---

# Search Engine Integration

Maintains

```
Recent Searches

Filters

Cached Results
```

---

# Model Router Integration

Tracks

```
Current Provider

Selected Model

Routing Status
```

---

# Plugin Integration

Plugins may store

- Configuration
- Runtime state
- Temporary data
- UI preferences

---

# Skills Integration

Skills may persist

- Templates
- Knowledge state
- Progress
- Custom variables

---

# Cache Strategy

Cache

```
Frequently Accessed State

Snapshots

History

Metadata
```

to improve performance.

---

# Error Handling

```
Invalid State

↓

Reject Update

↓

Restore Previous Snapshot

↓

Continue
```

Never allow corrupted state to propagate.

---

# Security

Always

- Validate updates
- Encrypt persisted state if required
- Restrict access
- Audit critical changes
- Protect sensitive runtime data

Never

- Store secrets in plain text
- Accept invalid state
- Skip validation
- Lose recovery checkpoints

---

# Performance Optimizations

Use

- Incremental updates
- Immutable state patterns
- Snapshot compression
- Lazy persistence
- Batched notifications

Avoid

- Full state rewrites
- Duplicate state objects
- Excessive serialization

---

# Best Practices

Always

- Centralize runtime state
- Validate every update
- Keep snapshots
- Notify observers
- Persist critical state
- Support recovery

Never

- Maintain duplicate state stores
- Ignore update conflicts
- Skip persistence for important data
- Mix unrelated state domains

---

# Common Mistakes

Bad

```
Each Module

↓

Own Runtime State
```

Inconsistent behavior.

---

Good

```
All Modules

↓

State Manager

↓

Shared Runtime State
```

Reliable and synchronized.

---

# Testing Checklist

- State creation
- State updates
- Validation
- Observer notifications
- Snapshots
- Persistence
- Recovery
- Undo/Redo
- Cache
- Error handling

---

# Advantages

- Consistent runtime data
- Easier debugging
- Reliable recovery
- Better synchronization
- Modular architecture
- Improved scalability

---

# Disadvantages

- Additional memory usage
- Snapshot storage
- Synchronization complexity
- Serialization overhead

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Workspace
- Enterprise AI Platforms

---

# Complete State Flow

```
Application

↓

State Manager

↓

Validator

↓

State Store

↓

Observers

↓

Persistence

↓

Snapshots

↓

Recovery

↓

Application Components
```

---

# Summary

The **State Manager** is the centralized runtime state layer responsible for maintaining, validating, synchronizing, persisting, and recovering application state across all components of an AI Coding Agent.

A production-grade State Manager should include:

- State Manager
- State Store
- State Snapshot
- State Persistence
- State Observer
- State Validator
- State Serializer
- State Recovery
- State History
- Event Bus Integration

By providing a single source of truth for runtime information, along with reactive updates, snapshots, persistence, and recovery mechanisms, the State Manager enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to operate consistently, reliably, and efficiently throughout complex development workflows.