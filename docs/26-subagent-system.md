# 26-subagent-system.md

# Sub-Agent System Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Sub-Agent System?

A **Sub-Agent System** is the subsystem responsible for creating, managing, coordinating, and supervising multiple specialized AI agents that work together to solve complex tasks.

Instead of one AI handling everything, the primary agent delegates specific responsibilities to specialized sub-agents, each with its own tools, permissions, context, and objectives.

The Sub-Agent System enables parallel reasoning, task specialization, and scalable automation.

---

# Why Sub-Agent System?

Without Sub-Agents

```
User Request

↓

Single Agent

↓

Everything

↓

Response
```

Problems

- Slow execution
- Large context
- Limited specialization
- Sequential processing
- Hard scalability

---

With Sub-Agent System

```
User Request

↓

Main Agent

↓

Task Planner

↓

Sub-Agents

↓

Results

↓

Final Response
```

---

# Goals

A production Sub-Agent System should provide

- Task decomposition
- Specialized agents
- Parallel execution
- Agent communication
- Context isolation
- Permission isolation
- Result aggregation
- Agent lifecycle management
- Resource optimization
- Failure recovery

---

# High-Level Architecture

```
              User Request

                    │

                    ▼

              Main Agent

                    │

                    ▼

           Sub-Agent Manager

                    │

     ┌──────────────┼──────────────┐

     ▼              ▼              ▼

 Coding Agent  Search Agent  Review Agent

     ▼              ▼              ▼

 Results       Results       Results

     └──────────────┼──────────────┘

                    ▼

          Result Aggregator

                    ▼

              Final Response
```

---

# Folder Structure

```
src/

subagent/

    SubAgentManager.ts

    AgentFactory.ts

    AgentRegistry.ts

    AgentScheduler.ts

    AgentExecutor.ts

    AgentContext.ts

    AgentCommunication.ts

    ResultAggregator.ts

    AgentLifecycle.ts

    AgentPolicy.ts

    SubAgentEvents.ts

    SubAgentMetrics.ts
```

---

# Core Components

## Sub-Agent Manager

Central controller.

Responsibilities

- Create agents
- Assign work
- Monitor execution
- Collect results

---

## Agent Factory

Creates new agents based on task requirements.

---

## Agent Registry

Stores

```
Agent ID

Capabilities

Status

Metadata
```

---

## Agent Scheduler

Schedules

- Sequential tasks
- Parallel tasks
- Priority tasks
- Background agents

---

## Agent Executor

Runs assigned tasks.

---

## Agent Context

Each sub-agent receives an isolated context containing

```
Task

Workspace

Files

Conversation

Permissions

Metadata
```

---

## Agent Communication

Allows agents to exchange

- Messages
- Intermediate results
- Status updates

without sharing unnecessary context.

---

## Result Aggregator

Combines outputs from multiple agents into a single coherent result.

---

## Agent Lifecycle

Tracks

```
Created

Running

Waiting

Completed

Failed

Destroyed
```

---

## Agent Policy

Defines

- Allowed tools
- Resource limits
- Maximum runtime
- Communication rules

---

# Agent Lifecycle

```
Create

↓

Initialize

↓

Assign Task

↓

Execute

↓

Return Result

↓

Destroy
```

---

# Agent Object

Contains

```
Agent ID

Role

Capabilities

Tools

Permissions

Status

Context
```

---

# Agent Types

Examples

```
Coding Agent

Review Agent

Testing Agent

Documentation Agent

Search Agent

Git Agent

Debug Agent

Deployment Agent

Planner Agent
```

---

# Task Delegation

```
Main Agent

↓

Split Task

↓

Assign

↓

Sub-Agents

↓

Results
```

---

# Parallel Execution

```
Task

↓

┌────────┬────────┬────────┐

▼        ▼        ▼

Agent1  Agent2  Agent3

↓

Aggregator

↓

Response
```

---

# Sequential Execution

```
Planner

↓

Coder

↓

Reviewer

↓

Tester

↓

Response
```

---

# Context Isolation

Each agent sees only the data required for its task.

Example

```
Coding Agent

↓

Code Only

Review Agent

↓

Generated Code

Search Agent

↓

Workspace Index
```

---

# Permission Isolation

Each agent has independent permissions.

Example

```
Search Agent

↓

Read Only

Coding Agent

↓

Read + Write

Deployment Agent

↓

Network + Execute
```

---

# Agent Communication Flow

```
Agent A

↓

Communication Bus

↓

Agent B
```

Messages are structured and validated.

---

# Result Aggregation

```
Code

↓

Tests

↓

Documentation

↓

Merge

↓

Final Output
```

---

# Event Bus Integration

Common events

```
agent:create

agent:start

agent:message

agent:complete

agent:fail

agent:destroy
```

---

# Workflow Engine Integration

The Workflow Engine may assign individual workflow steps to different sub-agents.

---

# Task Planner Integration

```
Task Planner

↓

Sub-Agent Manager

↓

Execution
```

---

# Conversation Manager Integration

Conversation history is shared selectively based on agent needs.

---

# Workspace Indexer Integration

Sub-agents retrieve only relevant workspace information.

---

# Search Engine Integration

Search agents perform retrieval tasks for other agents.

---

# Permission Engine Integration

Every sub-agent receives a restricted permission profile.

---

# Model Router Integration

Different sub-agents may use different language models depending on their specialization.

Example

```
Planner

↓

Reasoning Model

Coder

↓

Coding Model

Reviewer

↓

Fast Model
```

---

# Plugin Integration

Plugins may register

- New agent types
- Communication protocols
- Scheduling policies
- Aggregation strategies

---

# Skills Integration

Skills may define

- Specialized agents
- Agent templates
- Task decomposition rules
- Collaboration patterns

---

# Cache Strategy

Cache

```
Agent Templates

Capabilities

Intermediate Results

Execution Plans
```

to improve performance.

---

# Error Handling

```
Agent Failure

↓

Retry

↓

Replace Agent

↓

Continue Workflow
```

If recovery is impossible, notify the main agent.

---

# Security

Always

- Isolate contexts
- Restrict permissions
- Validate inter-agent messages
- Limit resource usage
- Audit agent actions

Never

- Share unnecessary context
- Allow unrestricted communication
- Grant excessive permissions
- Ignore failed agents

---

# Performance Optimizations

Use

- Parallel execution
- Lightweight agent contexts
- Shared caches
- Efficient scheduling
- Background workers

Avoid

- Creating unnecessary agents
- Duplicating work
- Oversized contexts
- Excessive communication

---

# Best Practices

Always

- Delegate specialized tasks
- Keep agent responsibilities focused
- Aggregate results centrally
- Monitor agent health
- Destroy idle agents

Never

- Use one agent for every task
- Share all context
- Ignore lifecycle management
- Bypass permission controls

---

# Common Mistakes

Bad

```
One Agent

↓

Everything

↓

Response
```

Slow and difficult to scale.

---

Good

```
Main Agent

↓

Sub-Agent Manager

↓

Specialized Agents

↓

Aggregator

↓

Final Response
```

Efficient and modular.

---

# Testing Checklist

- Agent creation
- Task delegation
- Parallel execution
- Sequential execution
- Context isolation
- Permission isolation
- Communication
- Aggregation
- Lifecycle management
- Failure recovery

---

# Advantages

- Better scalability
- Faster execution
- Specialized reasoning
- Parallel processing
- Improved maintainability
- Modular architecture

---

# Disadvantages

- Coordination complexity
- Resource management
- Communication overhead
- Scheduling complexity

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

# Complete Sub-Agent Flow

```
User Request

↓

Main Agent

↓

Task Planner

↓

Sub-Agent Manager

↓

Agent Scheduler

↓

Specialized Sub-Agents

↓

Workspace Index

↓

Search Engine

↓

Tool Manager

↓

Result Aggregator

↓

Conversation Manager

↓

Renderer

↓

User
```

---

# Summary

The **Sub-Agent System** is the distributed intelligence layer responsible for coordinating multiple specialized AI agents that collaborate to complete complex tasks efficiently.

A production-grade Sub-Agent System should include:

- Sub-Agent Manager
- Agent Factory
- Agent Registry
- Agent Scheduler
- Agent Executor
- Agent Context
- Agent Communication
- Result Aggregator
- Agent Lifecycle
- Agent Policy
- Event Bus Integration

By enabling task decomposition, context isolation, permission isolation, parallel execution, and intelligent result aggregation, the Sub-Agent System allows AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to achieve higher scalability, better specialization, and more efficient execution of sophisticated development workflows.