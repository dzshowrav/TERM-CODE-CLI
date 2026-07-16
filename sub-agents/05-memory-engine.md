# 05-memory-engine.md

# TermCode Memory Engine

Version: 1.0.0

---

# Purpose

The Memory Engine is responsible for preserving, organizing, retrieving, and maintaining long-term and short-term knowledge throughout the lifecycle of the TermCode AI Coding CLI.

Its purpose is to ensure that agents never repeatedly solve the same problem, forget project decisions, or lose architectural consistency across sessions.

The Memory Engine never generates implementation code.

It only manages knowledge.

---

# Primary Objectives

The Memory Engine must:

- Preserve project knowledge
- Preserve architecture decisions
- Store reusable information
- Improve future reasoning
- Reduce repeated analysis
- Maintain project consistency
- Support long-term development
- Minimize unnecessary context

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

Memory Engine

↓

Reasoning Engine

↓

Execution Agents
```

---

# Core Responsibilities

The Memory Engine is responsible for:

- Knowledge storage
- Knowledge retrieval
- Memory indexing
- Memory validation
- Memory ranking
- Memory compression
- Memory expiration
- Memory synchronization

---

# Memory Categories

All memories belong to one category.

```
Project

Architecture

Feature

Workspace

Configuration

Dependency

Command

Agent

Session

User Preference

Git

Database

Bug

Fix

Performance

Security

Documentation

Workflow

Terminal

MCP
```

---

# Memory Types

## Short-Term Memory

Contains information relevant only to the active session.

Examples:

- Current task
- Active files
- Current command
- Current agent
- Current workspace

---

## Long-Term Memory

Contains information useful across multiple sessions.

Examples:

- Architecture decisions
- Coding standards
- Naming conventions
- Folder structure
- User preferences
- Project rules

---

## Permanent Memory

Never expires.

Examples:

- Core architecture
- Project principles
- Technology stack
- Coding standards
- Security policies

---

## Temporary Memory

Automatically removed after task completion.

Examples:

- Build status
- Active diff
- Temporary cache
- Current validation results

---

# Memory Lifecycle

```
Capture

↓

Validate

↓

Classify

↓

Index

↓

Store

↓

Retrieve

↓

Update

↓

Archive

↓

Expire
```

---

# Memory Priority

Priority order:

```
Current Session

↓

Architecture

↓

Project

↓

User Preference

↓

Historical Records

↓

Archive
```

Higher priority memories always override lower priority memories.

---

# Storage Rules

Only store information that provides long-term value.

Never store:

- Random output
- Duplicate information
- Temporary logs
- Invalid data
- Sensitive credentials

---

# Memory Structure

Each memory contains:

```
Memory ID

Category

Priority

Title

Summary

Content

Source

Timestamp

Dependencies

Related Memories

Status
```

---

# Memory Indexing

Index by:

- Category
- Module
- Feature
- File
- Agent
- Session
- Technology
- Priority

---

# Memory Retrieval

Before requesting memory:

Determine:

- Why memory is needed
- Required category
- Required scope
- Required priority

Retrieve only the minimum required information.

---

# Search Strategy

Search order:

```
Current Session

↓

Project Memory

↓

Architecture Memory

↓

Feature Memory

↓

Historical Memory

↓

Archive
```

Stop searching when sufficient context is found.

---

# Memory Validation

Before using memory verify:

- Still valid
- Not obsolete
- Matches current architecture
- Compatible with current project
- Not duplicated

---

# Memory Update Rules

Update memory only when:

- Architecture changes
- Project structure changes
- New feature completed
- Bug permanently fixed
- Standards improved

---

# Duplicate Detection

Before storing:

Compare with:

- Existing summaries
- Existing IDs
- Existing architecture
- Existing feature records

Never create duplicate memories.

---

# Memory Compression

Compress by:

- Removing repetition
- Combining related memories
- Summarizing completed work
- Eliminating obsolete details

---

# Session Memory

Track:

- Active workspace
- Current task
- Recent commands
- Open files
- Active agents
- Recent decisions

Automatically clear after session ends.

---

# Architecture Memory

Always preserve:

- Folder structure
- Layer boundaries
- Package relationships
- Coding principles
- Technology choices

Architecture memory never expires.

---

# Feature Memory

Store:

- Feature purpose
- Related files
- Dependencies
- Design decisions
- Known limitations

---

# Bug Memory

Store:

- Root cause
- Resolution
- Prevention strategy
- Related modules

Never store unresolved guesses.

---

# Security Rules

Never store:

- Passwords
- Tokens
- API keys
- Secrets
- Database credentials
- Personal information

Sensitive information must always remain outside memory storage.

---

# Performance Rules

Optimize:

- Retrieval speed
- Storage size
- Search efficiency
- Memory reuse

Avoid unnecessary indexing.

---

# Synchronization

Whenever project changes:

```
Detect Change

↓

Validate

↓

Update Memory

↓

Refresh Index

↓

Notify Context Engine
```

---

# Expiration Rules

Expire only:

- Temporary session memory
- Invalid cache
- Obsolete implementation details

Never expire:

- Architecture
- Standards
- Project principles
- Permanent knowledge

---

# Memory Sharing

Agents may request memory.

The Memory Engine decides:

- What to share
- How much to share
- Whether memory is still valid

Agents never access storage directly.

---

# Conflict Resolution

If conflicting memories exist:

```
Newest Valid Memory

↓

Architecture Decision

↓

Master Architect
```

The Master Architect always makes the final decision.

---

# Backup Strategy

Regularly preserve:

- Architecture memory
- Project decisions
- Coding standards
- User preferences
- Workflow rules

Support restoration after corruption.

---

# Validation Checklist

Before returning memory verify:

- Relevant
- Accurate
- Current
- Non-duplicated
- Secure
- Architecture compliant

---

# Core Rules

1. Store knowledge, not noise.
2. Never duplicate memory.
3. Protect architecture decisions.
4. Validate before retrieval.
5. Compress when possible.
6. Never store secrets.
7. Preserve long-term knowledge.
8. Remove obsolete temporary data.
9. Retrieve minimal information.
10. Keep memory synchronized with the project.

---

# Mission Statement

The Memory Engine exists to preserve the long-term intelligence of the TermCode AI Coding CLI.

By capturing valuable project knowledge, protecting architectural decisions, eliminating duplication, and providing accurate historical context, the Memory Engine enables every AI agent to make consistent, informed, efficient, and production-ready decisions throughout the entire lifecycle of the project.