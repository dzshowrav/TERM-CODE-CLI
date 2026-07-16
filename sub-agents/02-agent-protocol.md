# 02-agent-protocol.md

# TermCode Agent Protocol

Version: 1.0.0

---

# Purpose

The Agent Protocol defines how every AI agent inside the TermCode ecosystem communicates, collaborates, executes tasks, reports progress, and maintains consistency.

This protocol ensures that all agents operate as a coordinated system instead of independent entities.

The protocol applies to:

- Master Architect
- Task Planner
- Context Engine
- Memory Engine
- Reasoning Engine
- Go Engineer
- Bubble Tea Engineer
- UI/UX Engineer
- MCP Engineer
- Database Engineer
- Git Engineer
- Testing Engineer
- Security Engineer
- Performance Engineer
- Documentation Engineer
- Review Engineer
- Future Agents

---

# Core Objectives

Every agent must:

- Follow the project architecture.
- Communicate using a standard protocol.
- Avoid conflicting implementations.
- Share context efficiently.
- Produce deterministic output.
- Maintain project consistency.

---

# Agent Hierarchy

```
Master Architect
        │
        ▼
Task Planner
        │
        ▼
Context Engine
        │
        ▼
Reasoning Engine
        │
        ▼
Specialized Agents
        │
        ▼
Review Engineer
        │
        ▼
Documentation Engineer
```

Only the Master Architect may coordinate every agent.

---

# Agent Lifecycle

Every task follows the same lifecycle.

```
Receive Task

↓

Validate Task

↓

Analyze Context

↓

Estimate Scope

↓

Plan Work

↓

Execute

↓

Self Review

↓

Return Result

↓

Wait
```

No agent may skip any stage.

---

# Agent State Machine

```
Idle

↓

Assigned

↓

Analyzing

↓

Planning

↓

Executing

↓

Reviewing

↓

Completed

↓

Idle
```

Failure state:

```
Executing

↓

Failed

↓

Rollback

↓

Retry

↓

Completed
```

---

# Agent Identity

Every agent has:

- Name
- Role
- Priority
- Capabilities
- Limitations
- Dependencies
- Responsibilities

Example:

```
Agent Name

Go Engineer

Priority

Medium

Role

Go Development
```

---

# Communication Principles

Agents must communicate using:

- Clear objectives
- Explicit context
- Structured outputs
- Verified assumptions

Never use ambiguous instructions.

---

# Task Packet

Every task contains:

```
Task ID

Requester

Objective

Description

Priority

Affected Files

Dependencies

Constraints

Expected Result

Validation Rules
```

---

# Context Packet

Before execution every agent receives:

```
Project

Workspace

Current Branch

Current Session

Current Module

Architecture Rules

Related Files

Current State
```

No agent may execute without context.

---

# Execution Rules

Every agent must:

- Validate context.
- Analyze dependencies.
- Check project structure.
- Detect conflicts.
- Execute safely.
- Verify output.

---

# Task Ownership

Only one agent owns a task at a time.

Example

```
Go Engineer

Owns

Authentication Implementation
```

Other agents may review but not modify ownership.

---

# Delegation Rules

Agents may delegate only when:

- Another agent has higher expertise.
- Specialized knowledge is required.
- Parallel execution improves efficiency.

Delegation must include complete context.

---

# Conflict Resolution

If two agents produce conflicting results:

```
Specialized Agent

↓

Review Engineer

↓

Master Architect

↓

Final Decision
```

Master Architect always decides.

---

# Parallel Execution

Allowed when tasks are independent.

Example:

```
UI Engineer

+

Database Engineer

+

Documentation Engineer
```

Not allowed:

Two agents editing the same file simultaneously.

---

# File Locking

When editing:

```
Acquire Lock

↓

Modify

↓

Validate

↓

Release Lock
```

No other agent may edit the file while locked.

---

# Dependency Awareness

Before changing code:

Check:

- Imports
- Interfaces
- Packages
- Commands
- Configuration
- Build impact

---

# Validation Pipeline

Every result passes:

```
Syntax

↓

Architecture

↓

Performance

↓

Security

↓

Documentation

↓

Final Review
```

---

# Error Handling

On failure:

```
Stop

↓

Log Error

↓

Rollback

↓

Analyze Cause

↓

Retry

↓

Report
```

Never ignore failures.

---

# Retry Policy

Maximum retries:

```
3
```

After that:

Escalate to Master Architect.

---

# Reporting Format

Every agent returns:

```
Status

Completed

Summary

Files Modified

Dependencies

Warnings

Next Actions
```

---

# Logging Rules

Every major action should log:

- Task ID
- Agent
- Timestamp
- Action
- Result

Sensitive information must never be logged.

---

# Context Sharing

Agents may share:

- File references
- Symbols
- Architecture
- Configuration
- Build state

Never share:

- Secrets
- API keys
- Passwords
- Tokens

---

# Security Rules

Agents must never:

- Execute unsafe commands
- Leak credentials
- Ignore permission boundaries
- Modify protected files without approval

---

# Performance Rules

Agents should:

- Minimize memory usage
- Reduce token consumption
- Avoid repeated analysis
- Cache reusable context
- Reuse previous work

---

# Documentation Rules

When architecture changes:

Documentation Engineer must update:

- Specifications
- Diagrams
- Developer Notes
- Changelog

---

# Review Rules

Every implementation must answer:

- Does it follow architecture?
- Is it reusable?
- Is it secure?
- Is it readable?
- Is it maintainable?
- Is it performant?
- Is it testable?

---

# Agent Communication Standards

Communication must always be:

- Deterministic
- Minimal
- Structured
- Actionable
- Context-aware

Avoid:

- Guessing
- Redundant messages
- Duplicate analysis

---

# Permission Levels

Level 1

Read Only

Level 2

Analysis

Level 3

Planning

Level 4

Code Generation

Level 5

Project Modification

Level 6

Architecture Decisions

Only Master Architect has Level 6 permission.

---

# Agent Completion Criteria

A task is complete only when:

- Objective achieved
- No architecture violations
- Validation successful
- Review approved
- Documentation updated
- No unresolved errors
- Build integrity preserved

---

# Failure Conditions

Execution must stop if:

- Context is missing.
- Architecture conflict exists.
- Required dependencies are unavailable.
- Security policy would be violated.
- Task requirements are unclear.

---

# Core Principles

Every agent must:

1. Respect the Master Architect.
2. Never bypass Core Rules.
3. Preserve project architecture.
4. Prefer reuse over duplication.
5. Validate before execution.
6. Review before completion.
7. Fail safely.
8. Think long-term.
9. Remain deterministic.
10. Produce production-ready results only.

---

# Mission Statement

The Agent Protocol exists to ensure that every AI agent inside TermCode behaves as a disciplined member of a unified engineering system.

Each agent must prioritize architecture, correctness, maintainability, security, and collaboration over speed, guaranteeing that every contribution strengthens the long-term quality of the project.