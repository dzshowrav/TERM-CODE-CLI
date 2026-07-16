# 19-permission-engine.md

# Permission Engine Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Permission Engine?

A **Permission Engine** is the security subsystem responsible for determining **whether an AI Agent is allowed to perform a requested action**.

Every operation that may affect the system, user data, workspace, network, or external services must pass through the Permission Engine before execution.

The Permission Engine is the central authorization layer between the Agent and every executable capability.

---

# Why Permission Engine?

Without Permission Engine

```
Agent

↓

Filesystem

↓

Delete Files
```

Problems

- Dangerous operations
- Unauthorized access
- Data loss
- Security risks
- No auditing

---

With Permission Engine

```
Agent

↓

Permission Engine

↓

Allow?

↓

Tool Execution

↓

Result
```

Every action is validated before execution.

---

# Goals

A production Permission Engine should provide

- Central authorization
- Policy-based permissions
- Workspace isolation
- User approval workflow
- Tool restrictions
- Filesystem protection
- Network restrictions
- Command validation
- Audit logging
- Plugin security

---

# High-Level Architecture

```
              Agent Engine

                    │

                    ▼

          Permission Engine

                    │

      ┌─────────────┼─────────────┐

      ▼             ▼             ▼

 Policies      Validators     Approvals

      ▼             ▼             ▼

 Audit Log     Rule Engine   Security

      └─────────────┼─────────────┘

                    ▼

             Tool Execution
```

---

# Folder Structure

```
src/

permission/

    PermissionEngine.ts

    PermissionPolicy.ts

    PermissionValidator.ts

    PermissionRule.ts

    PermissionResolver.ts

    ApprovalManager.ts

    WorkspaceGuard.ts

    FilesystemGuard.ts

    NetworkGuard.ts

    CommandGuard.ts

    PluginGuard.ts

    PermissionEvents.ts

    PermissionMetrics.ts

    AuditLogger.ts
```

---

# Core Components

## Permission Engine

Central controller.

Responsibilities

- Check permissions
- Evaluate policies
- Request approvals
- Return authorization result

---

## Permission Policy

Defines rules for

- Users
- Tools
- Workspaces
- Plugins
- Commands

---

## Permission Validator

Checks

```
Action

↓

Policy

↓

Allowed?
```

---

## Rule Engine

Evaluates

- Allow rules
- Deny rules
- Conditional rules
- Priority rules

---

## Approval Manager

Handles

```
Sensitive Action

↓

Ask User

↓

Approve

↓

Continue
```

or

```
Reject

↓

Stop Execution
```

---

## Workspace Guard

Protects

- Current workspace
- Parent directories
- External projects

---

## Filesystem Guard

Controls

```
Read

Write

Delete

Rename

Move

Execute
```

---

## Network Guard

Controls

- HTTP requests
- HTTPS requests
- Downloads
- Uploads
- External APIs

---

## Command Guard

Validates terminal commands before execution.

Examples

Allowed

```
ls

pwd

git status
```

Restricted

```
rm -rf /

shutdown

format
```

---

## Plugin Guard

Ensures plugins

- Cannot bypass permissions
- Cannot access restricted resources
- Cannot execute unauthorized actions

---

## Audit Logger

Records

```
Who

What

When

Result

Reason
```

---

# Permission Lifecycle

```
Request

↓

Identify Action

↓

Load Policies

↓

Evaluate Rules

↓

Approval Needed?

↓

Allow / Deny

↓

Log Result
```

---

# Permission Object

Contains

```
Action

Tool

Workspace

User

Plugin

Risk Level

Timestamp

Metadata
```

---

# Permission Levels

Examples

```
Read Only

↓

Write

↓

Execute

↓

Administrator
```

---

# Risk Levels

```
Low

Medium

High

Critical
```

High-risk actions may always require user confirmation.

---

# Policy Types

Examples

```
Filesystem

Network

Terminal

Git

Database

Plugins

MCP

Environment Variables
```

---

# Rule Priority

```
Explicit Deny

↓

Explicit Allow

↓

Conditional Rules

↓

Default Policy
```

---

# Approval Flow

```
Delete File

↓

Permission Engine

↓

User Approval

↓

Approved

↓

Continue
```

---

# Filesystem Protection

Restrict

```
System Directories

Private Files

Configuration

Secrets
```

Allow only approved workspace access.

---

# Network Protection

Policies may allow

```
Specific Domains

↓

Allowed
```

while blocking unknown destinations.

---

# Terminal Protection

Before execution

```
Command

↓

Validator

↓

Safe?

↓

Execute
```

---

# Plugin Security

Plugins receive only

```
Granted Permissions
```

They cannot elevate privileges themselves.

---

# MCP Integration

Each MCP server declares

```
Capabilities

↓

Permission Engine

↓

Authorization
```

before tools are executed.

---

# Event Bus Integration

Common events

```
permission:request

permission:allow

permission:deny

permission:approval

permission:error
```

---

# Tool Manager Integration

```
Tool Request

↓

Permission Engine

↓

Authorized

↓

Tool Execution
```

No tool executes without authorization.

---

# Session Integration

Session stores

```
Temporary Approvals

Permission Cache

Security Context
```

---

# Conversation Integration

User approvals become part of the conversation history.

---

# Plugin Integration

Plugins register

- Required permissions
- Optional permissions
- Capability declarations

---

# Cache Strategy

Cache

```
Permission Decisions

Policy Metadata

Approval Tokens
```

for repeated operations within a session.

---

# Error Handling

```
Unknown Permission

↓

Deny

↓

Log

↓

Notify
```

Default to the safest behavior.

---

# Security Principles

Always

- Deny by default
- Validate every action
- Require approval for risky operations
- Log all privileged actions
- Isolate workspaces
- Protect secrets

Never

- Trust user input blindly
- Allow unrestricted terminal access
- Skip authorization
- Allow privilege escalation
- Expose confidential data

---

# Performance Optimizations

Use

- Permission cache
- Policy indexing
- Rule compilation
- Incremental validation
- Lightweight authorization checks

Avoid

- Reloading policies every request
- Recalculating identical permissions
- Blocking unrelated operations

---

# Best Practices

Always

- Centralize authorization
- Separate policy from execution
- Validate before execution
- Audit every privileged action
- Support user approvals
- Keep rules explicit

Never

- Embed permissions inside tools
- Duplicate security logic
- Ignore failed validations
- Allow hidden privilege escalation

---

# Common Mistakes

Bad

```
Agent

↓

Tool

↓

Filesystem
```

Tool executes immediately.

---

Good

```
Agent

↓

Permission Engine

↓

Authorization

↓

Tool

↓

Filesystem
```

Secure and auditable.

---

# Testing Checklist

- Read permissions
- Write permissions
- Delete protection
- Network restrictions
- Terminal validation
- Plugin permissions
- Approval workflow
- Workspace isolation
- Audit logging
- Cache behavior
- Error recovery

---

# Advantages

- Strong security
- Central authorization
- Reduced risk
- Plugin isolation
- Better auditing
- Safer automation

---

# Disadvantages

- Additional validation overhead
- Policy maintenance
- Approval workflow complexity
- Rule management

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Agents
- Enterprise AI Platforms

---

# Complete Permission Flow

```
User Request

↓

Agent Engine

↓

Tool Manager

↓

Permission Engine

↓

Policy Evaluation

↓

Approval Manager

↓

Authorization

↓

Tool Execution

↓

Audit Logger

↓

Conversation Manager

↓

Renderer

↓

User
```

---

# Summary

The **Permission Engine** is the centralized security and authorization layer responsible for deciding whether an AI Agent may execute a requested action.

A production-grade Permission Engine should include:

- Permission Engine
- Policy Manager
- Rule Engine
- Permission Validator
- Approval Manager
- Workspace Guard
- Filesystem Guard
- Network Guard
- Command Guard
- Plugin Guard
- Audit Logger
- Permission Cache
- Event Bus Integration

By separating authorization from execution, the Permission Engine provides secure, policy-driven control over tools, files, networks, plugins, terminals, and MCP servers, ensuring that AI Coding Agents such as OpenCode, Antigravity CLI, and enterprise AI platforms operate safely, transparently, and with full accountability.