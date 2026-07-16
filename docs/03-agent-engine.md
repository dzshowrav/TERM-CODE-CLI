# 03-agent-engine.md

# Agent Engine Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is an Agent Engine?

The **Agent Engine** is the brain of an AI Coding CLI.

It receives a user request, understands the goal, builds context, selects tools, communicates with the LLM, executes actions, evaluates results, and returns the final response.

The Agent Engine does **not** directly control the UI, plugins, or file system. Instead, it coordinates them through the Event Bus.

---

# Responsibilities

The Agent Engine is responsible for:

- Understanding user intent
- Building prompt context
- Selecting models
- Planning tasks
- Calling tools
- Managing MCP servers
- Executing skills
- Reading and writing files
- Streaming responses
- Handling retries
- Maintaining conversation memory
- Producing structured outputs

---

# High-Level Architecture

```
                User
                  │
                  ▼
          Command Parser
                  │
                  ▼
           Agent Engine
                  │
    ┌─────────────┼─────────────┐
    ▼             ▼             ▼
Context      Planner       Memory
Builder
    │             │             │
    └──────┬──────┴──────┬──────┘
           ▼             ▼
      Tool Manager   Model Router
           │             │
           ▼             ▼
      MCP Servers      AI Model
           │             │
           └──────┬──────┘
                  ▼
          Response Stream
                  │
                  ▼
             Event Bus
                  │
                  ▼
                 TUI
```

---

# Goals

A production Agent Engine should provide

- Multi-step reasoning
- Tool usage
- Streaming
- Context awareness
- Session memory
- Retry handling
- Plugin compatibility
- Skill compatibility
- Multiple model support
- Safe execution

---

# Folder Structure

```
src/

agent/

    Agent.ts

    AgentEngine.ts

    AgentSession.ts

    Planner.ts

    ContextBuilder.ts

    PromptBuilder.ts

    Conversation.ts

    Memory.ts

    TokenManager.ts

    ResponseParser.ts

    StreamHandler.ts

    RetryManager.ts

    ModelRouter.ts

    ToolManager.ts

    MCPManager.ts

    SkillRunner.ts

    ExecutionManager.ts

    ResultEvaluator.ts

    Safety.ts

    Metrics.ts
```

---

# Core Components

## Agent

Represents one AI assistant.

Responsibilities

- Accept requests
- Coordinate execution
- Maintain state

---

## Agent Engine

The central controller.

Responsibilities

- Execute workflow
- Manage lifecycle
- Coordinate components

---

## Planner

Breaks one large task into smaller tasks.

Example

```
Create Blog

↓

Create Folder

↓

Generate Files

↓

Install Packages

↓

Run Tests
```

---

## Context Builder

Collects information from

- Conversation
- Workspace
- Open files
- Git
- Settings
- Skills
- MCP
- Plugins

Produces one unified context.

---

## Prompt Builder

Creates the final prompt sent to the LLM.

Combines

```
System Prompt

+

Conversation

+

Workspace

+

Tools

+

User Input
```

---

## Model Router

Chooses which model to use.

Example

```
Simple Question

↓

Small Model

----------------

Complex Refactor

↓

Large Model
```

---

## Tool Manager

Maintains all available tools.

Examples

```
Read File

Write File

Search

Terminal

Git

Browser

MCP Tool
```

---

## MCP Manager

Communicates with external MCP servers.

Responsibilities

- Connect
- Disconnect
- Discover tools
- Execute tools
- Handle errors

---

## Skill Runner

Loads reusable skills.

Example

```
Laravel Skill

↓

Prompt Template

↓

Execution Rules

↓

Output
```

---

## Memory

Stores

```
Conversation

Workspace Facts

Recent Actions

Agent Notes
```

Memory improves future responses.

---

## Conversation Manager

Stores

```
User Messages

Assistant Messages

Tool Results

System Messages
```

---

## Token Manager

Tracks

```
Prompt Tokens

Completion Tokens

Remaining Context

Cost
```

---

## Stream Handler

Receives streamed tokens.

Example

```
H

He

Hel

Hell

Hello
```

Each token updates the UI.

---

## Retry Manager

Handles

```
Timeout

↓

Retry

↓

Fallback Model

↓

Error
```

---

## Execution Manager

Controls execution order.

Example

```
Plan

↓

Tool

↓

Result

↓

Next Step

↓

Finish
```

---

## Result Evaluator

Checks whether the task succeeded.

Example

```
Compile

↓

Passed

↓

Done
```

or

```
Compile

↓

Failed

↓

Retry
```

---

# Agent Lifecycle

```
Receive Request

↓

Load Session

↓

Build Context

↓

Plan

↓

Select Tools

↓

Call LLM

↓

Execute Tools

↓

Evaluate

↓

Respond

↓

Save Session
```

---

# Request Flow

```
User

↓

Command Parser

↓

Agent

↓

Planner

↓

Context Builder

↓

Prompt Builder

↓

Model

↓

Response

↓

TUI
```

---

# Tool Execution Flow

```
Model Requests Tool

↓

Tool Manager

↓

Permission Check

↓

Execute Tool

↓

Result

↓

Return To Model
```

---

# Streaming Flow

```
Model

↓

Token

↓

Stream Handler

↓

Event Bus

↓

TUI

↓

User
```

---

# Context Sources

The Agent may collect context from

- Conversation history
- Workspace files
- Git repository
- Configuration
- Installed plugins
- Skills
- MCP servers
- Environment variables
- Active terminal

---

# State Machine

```
Idle

↓

Planning

↓

Thinking

↓

Calling Tool

↓

Streaming

↓

Completed
```

Possible error state

```
Failed
```

---

# Event Bus Integration

Example events

```
agent:start

agent:thinking

agent:tool

agent:stream

agent:complete

agent:error
```

Other modules subscribe to these events.

---

# File Watcher Integration

```
File Changed

↓

Context Invalidated

↓

Reload Context

↓

Continue
```

---

# Theme Integration

Agent never changes colors directly.

Instead

```
Emit Event

↓

Renderer Updates Theme
```

---

# Plugin Integration

Plugins may

- Add tools
- Modify prompts
- Listen to events
- Add context
- Validate output

---

# Skills Integration

Skills may

Provide

- Prompt templates
- Execution rules
- Coding standards
- Project conventions

---

# Memory Strategy

Short-Term Memory

Stores

```
Current Session
```

Long-Term Memory

Stores

```
Project Facts

Preferences

Learned Information
```

---

# Error Handling

```
LLM Error

↓

Retry

↓

Fallback

↓

Log

↓

Notify User
```

---

# Security

Before executing tools

Always verify

- Permissions
- Allowed paths
- Dangerous commands
- Plugin trust
- MCP trust

---

# Performance Optimizations

Use

- Streaming
- Prompt caching
- Context compression
- Parallel tool execution
- Incremental memory loading
- Lazy plugin loading

Avoid

- Reloading everything
- Blocking UI
- Duplicate tool calls

---

# Best Practices

Always

- Keep engine modular
- Separate planning from execution
- Stream responses
- Log important events
- Cache reusable context
- Use event-driven communication
- Validate tool outputs

Never

- Mix rendering with logic
- Directly print to terminal
- Hardcode prompts
- Ignore tool failures
- Block execution unnecessarily

---

# Common Mistakes

Bad

```
Agent

↓

Print UI

↓

Read File

↓

Call Model

↓

Write File
```

Everything tightly coupled.

---

Good

```
Agent

↓

Planner

↓

Tool Manager

↓

Model Router

↓

Event Bus

↓

TUI
```

Each component has one responsibility.

---

# Testing Checklist

- Request processing
- Planning
- Tool execution
- Streaming
- Retry logic
- Memory loading
- Context building
- Plugin integration
- Skill execution
- MCP communication
- Error recovery

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Gemini CLI
- GitHub Copilot CLI
- Cursor Agent
- Continue.dev
- Enterprise AI Agents

---

# Summary

The **Agent Engine** is the central intelligence layer of an AI CLI Coding Agent.

It coordinates planning, context management, model interaction, tool execution, memory, streaming, and evaluation while communicating with the rest of the system through the Event Bus.

A production-grade Agent Engine should be:

- Modular
- Event-driven
- Tool-aware
- Context-aware
- Streaming-first
- Secure
- Extensible
- Model-independent
- Plugin-friendly
- Highly testable

A well-designed Agent Engine allows OpenCode- or Antigravity-style applications to scale from simple chat interactions to complex autonomous coding workflows while keeping each subsystem cleanly separated.