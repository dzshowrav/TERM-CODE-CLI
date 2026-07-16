# 04-context-engine.md

# TermCode Context Engine

Version: 1.0.0

---

# Purpose

The Context Engine is responsible for collecting, organizing, prioritizing, maintaining, and distributing all contextual information required by the AI agents inside TermCode.

It ensures that every decision is made using the most relevant project information while minimizing unnecessary token usage.

The Context Engine never generates implementation code.

Its only responsibility is managing context.

---

# Primary Objectives

The Context Engine must:

- Understand the current project
- Understand the current workspace
- Understand the current task
- Understand project history
- Understand architecture
- Minimize context size
- Maximize context quality
- Eliminate duplicate information
- Keep context synchronized

---

# Position in Architecture

```
User

↓

Master Architect

↓

Task Planner

↓

Context Engine

↓

Reasoning Engine

↓

Execution Agents
```

---

# Core Responsibilities

The Context Engine is responsible for:

- Workspace discovery
- Project indexing
- Context extraction
- Context prioritization
- Context compression
- Context validation
- Context synchronization
- Context distribution

---

# Context Sources

The engine collects context from:

- User requests
- Project files
- Configuration files
- Documentation
- Memory Engine
- Git history
- Previous sessions
- MCP responses
- Agent outputs
- Runtime state

---

# Context Categories

Every context belongs to one category.

```
Project

Architecture

Workspace

Task

Session

User

Configuration

Source Code

Database

Git

Documentation

Terminal

MCP

Memory

Runtime
```

---

# Context Hierarchy

Priority order:

```
Current User Request

↓

Current Task

↓

Current File

↓

Current Module

↓

Project Architecture

↓

Workspace

↓

Session

↓

Memory

↓

Documentation

↓

Historical Context
```

Higher priority always overrides lower priority.

---

# Context Lifecycle

```
Collect

↓

Validate

↓

Filter

↓

Prioritize

↓

Compress

↓

Distribute

↓

Monitor

↓

Update

↓

Expire
```

---

# Collection Rules

Always collect:

- Active workspace
- Current task
- Active module
- Relevant files
- Dependencies
- Architecture rules

Never collect unrelated information.

---

# Context Window

The active context should contain only:

- Relevant files
- Related modules
- Required documentation
- Required configuration
- Current session information

Avoid unnecessary expansion.

---

# Context Priority Levels

Level 1

Critical

Examples:

- Current task
- Current file
- Active module

---

Level 2

High

Examples:

- Related interfaces
- Related packages
- Dependencies

---

Level 3

Medium

Examples:

- Documentation
- Previous commits
- Configuration

---

Level 4

Low

Examples:

- Historical notes
- Archived sessions

---

# Context Compression

Before sending context:

Remove:

- Duplicate content
- Unused files
- Irrelevant documentation
- Obsolete state
- Repeated messages

---

# Context Expansion

Expand only when:

- Required dependency exists
- User requests additional work
- Architecture requires it
- Missing references detected

Never expand automatically without reason.

---

# Context Validation

Before distribution verify:

- File exists
- Module exists
- Reference valid
- Dependency valid
- Session active

---

# Context Synchronization

Whenever files change:

```
Detect Change

↓

Update Index

↓

Refresh Context

↓

Notify Agents
```

---

# Context Cache

Frequently used context should be cached.

Examples:

- Project structure
- Configuration
- Commands
- Active modules
- Symbols

---

# Cache Rules

Cache only:

- Stable information
- Frequently accessed information

Never cache:

- Secrets
- Passwords
- Tokens
- Temporary runtime state

---

# Workspace Analysis

Identify:

- Root directory
- Project type
- Programming language
- Build system
- Package manager
- Dependencies
- Configuration files

---

# Project Discovery

Automatically detect:

- Go project
- Node project
- Rust project
- Python project
- Mixed workspace

Load only relevant context.

---

# Symbol Indexing

Index:

- Packages
- Structs
- Interfaces
- Functions
- Methods
- Constants
- Variables
- Commands

---

# Dependency Mapping

Track:

```
Package

↓

Imports

↓

Interfaces

↓

Services

↓

Commands
```

Never lose dependency relationships.

---

# Session Context

Track:

- Active task
- Open files
- Recent commands
- Current branch
- Current workspace
- Active agents

---

# Memory Integration

Request memory only when:

- Historical information needed
- Previous implementation required
- User references past work

---

# Git Context

Collect:

- Current branch
- Modified files
- Recent commits
- Relevant history

Ignore unrelated commits.

---

# Documentation Context

Load only:

- Relevant specifications
- Related markdown
- API documentation
- Architecture documents

---

# Runtime Context

Monitor:

- Running processes
- Active terminal
- Background tasks
- Build status

---

# MCP Context

Collect only required responses.

Examples:

Filesystem

Git

Database

Search

Browser

Memory

Never request unnecessary MCP data.

---

# Security Rules

Never include:

- API Keys
- Passwords
- Secrets
- Tokens
- Private credentials

Sensitive information must remain isolated.

---

# Performance Rules

Minimize:

- Token usage
- Context size
- Disk reads
- File scanning
- Memory allocation

Reuse cached context whenever possible.

---

# Failure Handling

If context becomes invalid:

```
Invalidate Cache

↓

Reload

↓

Validate

↓

Redistribute
```

If recovery fails:

Escalate to Master Architect.

---

# Context Distribution

Each agent receives only the minimum context required.

Example:

Go Engineer

Receives:

- Go files
- Related interfaces
- Relevant packages

UI Engineer

Receives:

- Bubble Tea models
- UI components
- Theme files

Documentation Engineer

Receives:

- Markdown
- Specifications
- Changelog

---

# Expiration Rules

Remove context when:

- Task completed
- Session closed
- Workspace changed
- File deleted
- Cache invalid

Never retain stale context.

---

# Validation Checklist

Before distributing context verify:

- Relevant
- Complete
- Current
- Minimal
- Valid
- Secure
- Consistent

---

# Core Rules

1. Collect only relevant information.
2. Minimize token usage.
3. Prioritize current task.
4. Remove duplicate context.
5. Never expose secrets.
6. Cache stable information.
7. Refresh after file changes.
8. Distribute least required context.
9. Validate before use.
10. Keep context synchronized.

---

# Mission Statement

The Context Engine exists to ensure that every AI agent inside TermCode operates with accurate, relevant, minimal, and up-to-date information.

By intelligently managing project knowledge, reducing unnecessary context, and maintaining synchronization across the entire system, the Context Engine enables fast, reliable, architecture-aware, and production-ready decision making while preserving efficiency and scalability.