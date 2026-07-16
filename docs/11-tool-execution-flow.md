# 11-tool-execution-flow.md

# Tool Execution Flow Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is Tool Execution?

A **Tool Execution Flow** is the complete lifecycle that begins when an AI model decides it needs an external capability and ends when the tool result is returned to the model.

Instead of relying only on its internal knowledge, the Agent can call tools such as:

- Filesystem
- Terminal
- Git
- Browser
- MCP Tools
- Database
- Search
- HTTP APIs
- Plugins

The Tool Execution Flow ensures that these tools are executed safely, efficiently, and consistently.

---

# Why Tool Execution Flow?

Without Tool Manager

```
LLM

↓

Calls Filesystem

↓

Calls Git

↓

Calls Browser

↓

Calls Database
```

Everything becomes tightly coupled.

---

With Tool Execution Flow

```
LLM

↓

Tool Manager

↓

Permission Engine

↓

Executor

↓

Tool

↓

Result

↓

LLM
```

Every tool follows the same execution pipeline.

---

# Goals

A production Tool Execution Flow should provide

- Safe execution
- Permission validation
- Tool discovery
- Input validation
- Output validation
- Streaming support
- Parallel execution
- Retry handling
- Metrics
- Event integration

---

# High-Level Architecture

```
                User

                  │

                  ▼

             Agent Engine

                  │

                  ▼

            Tool Manager

                  │

      ┌───────────┼─────────────┐

      ▼           ▼             ▼

 Permission   Registry     Scheduler

      │           │             │

      └──────┬────┴──────┬──────┘

             ▼           ▼

       Tool Executor  Result Parser

             │

             ▼

       Filesystem

       Terminal

       Git

       Browser

       MCP

       Plugins

             │

             ▼

         Tool Result

             │

             ▼

        Agent Engine
```

---

# Folder Structure

```
src/

tools/

    ToolManager.ts

    ToolRegistry.ts

    ToolExecutor.ts

    ToolScheduler.ts

    ToolQueue.ts

    ToolValidator.ts

    ToolPermissions.ts

    ToolContext.ts

    ToolRequest.ts

    ToolResponse.ts

    ToolEvents.ts

    ToolMetrics.ts

    ToolRetry.ts

    ToolTimeout.ts

    ToolCache.ts

    ToolLogger.ts

    ToolResultParser.ts
```

---

# Core Components

## Tool Manager

Central controller.

Responsibilities

- Register tools
- Discover tools
- Execute tools
- Remove tools
- Manage lifecycle

---

## Tool Registry

Stores

```
Available Tools

Capabilities

Categories

Permissions

Metadata

Version
```

---

## Tool Executor

Responsible for

```
Prepare

↓

Execute

↓

Collect Result

↓

Return
```

---

## Tool Scheduler

Controls

- Parallel execution
- Sequential execution
- Priority
- Resource limits

---

## Tool Queue

Stores pending requests.

```
Request

↓

Queue

↓

Executor
```

---

## Tool Validator

Checks

- Tool exists
- Arguments
- Schema
- Required fields
- Types

---

## Permission Engine

Verifies

```
Filesystem

Network

Terminal

Git

Database

Environment
```

before execution.

---

## Timeout Manager

Every tool receives

```
Start

↓

Execution

↓

Timeout

↓

Cancel
```

---

## Retry Manager

Handles

```
Failure

↓

Retry

↓

Fallback

↓

Abort
```

---

## Tool Cache

Stores

```
Frequently Used Results

Metadata

Schemas
```

Improves performance.

---

## Result Parser

Normalizes tool output.

Example

```
JSON

↓

Structured Result

↓

Agent
```

---

# Tool Lifecycle

```
Register

↓

Validate

↓

Authorize

↓

Execute

↓

Collect Result

↓

Parse

↓

Return

↓

Log
```

---

# Tool Request Object

Contains

```
Tool ID

Arguments

Session

Workspace

Permissions

Metadata

Timeout
```

---

# Tool Response Object

Contains

```
Status

Output

Errors

Execution Time

Metadata

Logs
```

---

# Execution Flow

```
Agent

↓

Tool Request

↓

Permission Check

↓

Validator

↓

Executor

↓

Tool

↓

Response

↓

Parser

↓

Agent
```

---

# Parallel Execution

Example

```
Read File

Git Status

Search

↓

Run Together
```

Reduces latency.

---

# Sequential Execution

Example

```
Read Config

↓

Generate File

↓

Write File

↓

Commit Git
```

Execution order matters.

---

# Event Bus Integration

Common events

```
tool:register

tool:start

tool:complete

tool:error

tool:timeout

tool:retry
```

---

# Agent Integration

```
LLM Requests Tool

↓

Tool Manager

↓

Execute

↓

Result

↓

Continue Reasoning
```

---

# Plugin Integration

Plugins may

- Register tools
- Override tools
- Extend metadata
- Add validators

---

# Skills Integration

Skills may declare

```
Required Tools

Filesystem

Git

Docker

Browser
```

The Tool Manager resolves them automatically.

---

# MCP Integration

External tools execute through MCP.

```
Tool Request

↓

MCP Client

↓

Server

↓

Tool

↓

Result
```

The Agent does not need to know where the tool is hosted.

---

# File Watcher Integration

Tool execution may trigger

```
Write File

↓

Watcher Event

↓

Cache Update

↓

UI Refresh
```

---

# Session Integration

Store

```
Recent Tool Calls

Execution History

Tool Results
```

Useful for recovery and debugging.

---

# Error Handling

```
Tool Failed

↓

Retry

↓

Fallback Tool

↓

Return Error

↓

Continue
```

The Agent should recover gracefully whenever possible.

---

# Timeout Strategy

Example

```
Start

↓

30 Seconds

↓

Timeout

↓

Cancel

↓

Report
```

Timeouts prevent hung tools.

---

# Validation Strategy

Before execution

Check

- Tool exists
- Permission granted
- Arguments valid
- Workspace allowed
- Session active

---

# Security

Always

- Validate input
- Sanitize paths
- Restrict filesystem access
- Limit terminal execution
- Log privileged actions
- Verify MCP server trust

Never

- Execute arbitrary commands blindly
- Ignore permissions
- Trust plugin input
- Skip validation

---

# Performance Optimizations

Use

- Tool cache
- Parallel execution
- Connection pooling
- Lazy initialization
- Request batching
- Streaming
- Incremental parsing

Avoid

- Blocking UI
- Reloading tool metadata
- Duplicate execution

---

# Best Practices

Always

- Separate discovery from execution
- Keep tools stateless where possible
- Use structured input/output
- Emit lifecycle events
- Measure execution time
- Handle retries centrally

Never

- Mix tool logic with Agent logic
- Hardcode tool locations
- Ignore failures
- Return inconsistent formats

---

# Common Mistakes

Bad

```
Agent

↓

Run Filesystem

↓

Run Git

↓

Run Browser
```

The Agent becomes tightly coupled.

---

Good

```
Agent

↓

Tool Manager

↓

Executor

↓

Tool

↓

Result
```

Every tool follows the same architecture.

---

# Testing Checklist

- Tool registration
- Validation
- Permission checks
- Parallel execution
- Sequential execution
- Retry logic
- Timeout handling
- Result parsing
- Event emission
- Error recovery
- MCP tools
- Plugin tools

---

# Example Tool Categories

Examples

- Filesystem
- Terminal
- Git
- GitHub
- Browser
- Search
- HTTP
- Docker
- Kubernetes
- PostgreSQL
- SQLite
- Redis
- Figma
- Slack
- Jira
- Notion
- AWS
- Azure
- GCP

---

# Advantages

- Standardized execution
- Modular architecture
- Secure tool access
- Easier debugging
- Plugin compatibility
- MCP compatibility
- Parallel processing
- Better scalability

---

# Disadvantages

- More infrastructure
- Scheduling complexity
- Permission management overhead
- Additional validation cost

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

# Complete Execution Flow

```
User Request

↓

Agent Engine

↓

Planner

↓

LLM Decides Tool Needed

↓

Tool Manager

↓

Permission Engine

↓

Validator

↓

Scheduler

↓

Tool Executor

↓

Filesystem / MCP / Plugin / Terminal

↓

Tool Result

↓

Result Parser

↓

Agent Continues Reasoning

↓

Final Response

↓

Renderer

↓

User
```

---

# Summary

The **Tool Execution Flow** is the standardized execution pipeline that allows an AI Coding Agent to safely interact with local and remote capabilities.

A production-grade implementation should include:

- Tool Manager
- Tool Registry
- Permission Engine
- Validator
- Scheduler
- Queue
- Executor
- Retry Manager
- Timeout Manager
- Result Parser
- Metrics
- Event Bus Integration

By separating discovery, validation, authorization, scheduling, execution, and result processing, the Tool Execution Flow becomes secure, extensible, and maintainable while supporting built-in tools, plugins, and MCP servers through a unified architecture comparable to OpenCode and Antigravity CLI.