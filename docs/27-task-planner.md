# 27-task-planner.md

# Task Planner Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Task Planner?

A **Task Planner** is the reasoning and planning subsystem responsible for analyzing a user's request, breaking it into manageable tasks, determining execution order, estimating dependencies, assigning resources, and generating an optimized execution plan.

Instead of immediately calling tools or models, the AI Agent first creates a structured plan.

The Task Planner acts as the brain responsible for **"thinking before acting."**

---

# Why Task Planner?

Without Task Planner

```
User Request

↓

Agent

↓

Random Tool Calls

↓

Response
```

Problems

- Poor reasoning
- Duplicate work
- Wrong execution order
- Tool misuse
- Difficult recovery

---

With Task Planner

```
User Request

↓

Task Planner

↓

Execution Plan

↓

Workflow Engine

↓

Response
```

---

# Goals

A production Task Planner should provide

- Task decomposition
- Dependency analysis
- Execution planning
- Tool selection
- Resource estimation
- Parallel task detection
- Priority management
- Risk analysis
- Plan optimization
- Progress tracking

---

# High-Level Architecture

```
             User Request

                   │

                   ▼

            Task Planner

                   │

      ┌────────────┼────────────┐

      ▼            ▼            ▼

 Analyzer    Dependency    Optimizer

      ▼            ▼            ▼

 Plan       Priorities    Resources

      └────────────┼────────────┘

                   ▼

           Workflow Engine
```

---

# Folder Structure

```
src/

planner/

    TaskPlanner.ts

    TaskAnalyzer.ts

    TaskDecomposer.ts

    DependencyAnalyzer.ts

    PlanOptimizer.ts

    ResourceEstimator.ts

    PriorityManager.ts

    TaskScheduler.ts

    PlanValidator.ts

    PlanStore.ts

    PlannerEvents.ts

    PlannerMetrics.ts
```

---

# Core Components

## Task Planner

Central controller.

Responsibilities

- Analyze requests
- Create plans
- Optimize execution
- Deliver execution plan

---

## Task Analyzer

Understands

```
User Goal

↓

Task Requirements
```

Extracts intent, objectives, and constraints.

---

## Task Decomposer

Breaks large requests into smaller tasks.

Example

```
Build Website

↓

Setup Project

↓

Create Pages

↓

Add Components

↓

Test

↓

Deploy
```

---

## Dependency Analyzer

Determines

```
Task B

↓

Needs Task A
```

before execution.

---

## Plan Optimizer

Improves

- Execution order
- Parallelism
- Resource usage
- Completion time

---

## Resource Estimator

Estimates

```
Models

Tools

Memory

Workspace

Time
```

required for each task.

---

## Priority Manager

Assigns

```
High

Medium

Low

Background
```

priorities.

---

## Task Scheduler

Creates executable task queues.

---

## Plan Validator

Checks

- Completeness
- Dependency consistency
- Permission requirements
- Workflow compatibility

---

## Plan Store

Stores

- Current plan
- Plan history
- Execution metadata

---

# Planning Lifecycle

```
Receive Request

↓

Analyze

↓

Decompose

↓

Analyze Dependencies

↓

Optimize

↓

Validate

↓

Return Plan
```

---

# Task Object

Contains

```
Task ID

Description

Priority

Dependencies

Assigned Tools

Estimated Cost

Estimated Time

Status
```

---

# Plan Object

Contains

```
Plan ID

Tasks

Execution Order

Parallel Groups

Metadata

Status
```

---

# Task States

```
Created

↓

Planned

↓

Queued

↓

Running

↓

Completed

↓

Failed

↓

Cancelled
```

---

# Sequential Planning

```
Task A

↓

Task B

↓

Task C
```

Each task depends on the previous one.

---

# Parallel Planning

```
Task A

↓

┌────────┬────────┐

▼        ▼        ▼

TaskB   TaskC   TaskD

↓

Merge

↓

Continue
```

Independent tasks execute simultaneously.

---

# Dependency Graph

Example

```
Initialize Project

↓

Install Packages

↓

Generate Files

↓

Run Tests

↓

Deploy
```

---

# Tool Selection

Planner determines

```
Filesystem

Search

Git

Terminal

MCP

Browser

Database
```

required for each task.

---

# Risk Analysis

Identify

- Dangerous commands
- Permission requirements
- Network operations
- Long-running tasks

---

# Execution Optimization

Reduce

- Duplicate work
- Token usage
- Tool invocations
- Execution time

---

# Event Bus Integration

Common events

```
planner:start

planner:analyze

planner:plan

planner:update

planner:complete

planner:error
```

---

# Workflow Engine Integration

```
Task Plan

↓

Workflow Engine

↓

Execution
```

---

# Sub-Agent System Integration

Planner may assign tasks to specialized agents.

Example

```
Planner

↓

Coding Agent

Testing Agent

Documentation Agent
```

---

# Permission Engine Integration

Planner checks required permissions before execution.

---

# Search Engine Integration

Search may be planned before coding begins.

---

# Workspace Indexer Integration

Planner retrieves project structure before generating tasks.

---

# Model Router Integration

Planner may choose different models for

- Planning
- Coding
- Reviewing
- Documentation

---

# Conversation Manager Integration

The execution plan may be summarized and presented to the user.

---

# Plugin Integration

Plugins may contribute

- Planning rules
- Custom task types
- Optimizers
- Validators

---

# Skills Integration

Skills may provide reusable planning templates for

- Project creation
- Refactoring
- Debugging
- Testing
- Deployment

---

# Cache Strategy

Cache

```
Execution Plans

Dependency Graphs

Task Templates

Analysis Results
```

to avoid repeated planning.

---

# Error Handling

```
Invalid Plan

↓

Reanalyze

↓

Regenerate

↓

Validate Again
```

Never execute an invalid plan.

---

# Security

Always

- Validate task dependencies
- Respect permissions
- Estimate risks
- Verify execution order

Never

- Skip planning
- Ignore dangerous operations
- Execute incomplete plans
- Bypass validation

---

# Performance Optimizations

Use

- Incremental planning
- Cached templates
- Parallel analysis
- Lightweight dependency graphs
- Lazy optimization

Avoid

- Planning identical tasks repeatedly
- Overcomplicated plans
- Blocking execution while planning unrelated tasks

---

# Best Practices

Always

- Plan before execution
- Keep tasks small
- Validate plans
- Optimize execution
- Track progress
- Reuse templates

Never

- Execute without planning
- Ignore dependencies
- Duplicate work
- Create oversized tasks

---

# Common Mistakes

Bad

```
User Request

↓

Execute Immediately
```

No reasoning.

---

Good

```
User Request

↓

Task Planner

↓

Optimized Plan

↓

Workflow Engine

↓

Execution
```

Structured and reliable.

---

# Testing Checklist

- Task analysis
- Task decomposition
- Dependency analysis
- Priority assignment
- Tool selection
- Parallel planning
- Plan validation
- Cache
- Optimization
- Error recovery

---

# Advantages

- Better reasoning
- Predictable execution
- Lower token usage
- Reduced errors
- Improved scalability
- Easier debugging

---

# Disadvantages

- Planning overhead
- Dependency management complexity
- Template maintenance
- Optimization cost

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

# Complete Planning Flow

```
User Request

↓

Task Analyzer

↓

Task Decomposer

↓

Dependency Analyzer

↓

Priority Manager

↓

Resource Estimator

↓

Plan Optimizer

↓

Plan Validator

↓

Workflow Engine

↓

Sub-Agent System

↓

Tool Manager

↓

Execution

↓

Conversation Manager

↓

User
```

---

# Summary

The **Task Planner** is the reasoning and orchestration preparation layer responsible for analyzing requests, decomposing work, optimizing execution, and generating structured execution plans for an AI Coding Agent.

A production-grade Task Planner should include:

- Task Planner
- Task Analyzer
- Task Decomposer
- Dependency Analyzer
- Plan Optimizer
- Resource Estimator
- Priority Manager
- Task Scheduler
- Plan Validator
- Plan Store
- Event Bus Integration

By transforming high-level user requests into validated, dependency-aware, optimized execution plans, the Task Planner enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to execute complex development tasks with greater efficiency, reliability, and scalability.