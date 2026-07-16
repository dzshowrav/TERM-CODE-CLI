# 06-reasoning-engine.md

# TermCode Reasoning Engine

Version: 1.0.0

---

# Purpose

The Reasoning Engine is responsible for transforming context into intelligent decisions.

It receives validated context from the Context Engine, relevant knowledge from the Memory Engine, and execution plans from the Task Planner.

The Reasoning Engine never performs direct implementation.

Its responsibility is to determine **what should be done**, **why it should be done**, **how it should be done**, and **which agent should perform it**.

---

# Primary Objectives

The Reasoning Engine must:

- Understand intent
- Analyze requirements
- Evaluate architecture
- Compare alternatives
- Select the optimal strategy
- Minimize implementation risk
- Improve code quality
- Protect long-term maintainability

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

↓

Review Engineer
```

---

# Responsibilities

The Reasoning Engine is responsible for:

- Requirement analysis
- Logical reasoning
- Architectural reasoning
- Dependency reasoning
- Tool selection
- Agent selection
- Risk analysis
- Strategy generation
- Decision validation

---

# Reasoning Pipeline

Every request follows this process.

```
Receive Context

↓

Understand Intent

↓

Analyze Requirements

↓

Evaluate Constraints

↓

Generate Strategies

↓

Compare Alternatives

↓

Select Best Strategy

↓

Validate Decision

↓

Delegate Task
```

---

# Reasoning Principles

Always prioritize:

- Correctness
- Simplicity
- Maintainability
- Performance
- Security
- Scalability
- Readability

Never optimize prematurely.

---

# Intent Analysis

Identify:

- User objective
- Technical objective
- Business objective
- Architectural impact
- User experience impact

---

# Requirement Classification

Each request belongs to one or more categories.

```
Feature

Bug Fix

Optimization

Refactoring

Testing

Documentation

Security

Performance

Configuration

Research
```

---

# Context Evaluation

Analyze:

- Current workspace
- Active module
- Related files
- Dependencies
- Architecture
- Existing implementation
- Current session

Never reason without verified context.

---

# Memory Evaluation

Retrieve only relevant knowledge.

Examples:

- Previous implementation
- Similar feature
- Known bug
- Architecture decision
- Coding standard

Ignore unrelated memories.

---

# Constraint Analysis

Identify constraints before making decisions.

Examples:

- Mobile-first
- Termux compatibility
- Offline-first
- Go best practices
- Bubble Tea architecture
- Existing APIs
- Existing dependencies

---

# Alternative Generation

Always generate multiple possible solutions.

Example:

```
Option A

Reuse existing component

Option B

Extend existing module

Option C

Create new module
```

Never assume the first solution is the best.

---

# Decision Matrix

Evaluate each alternative based on:

- Complexity
- Performance
- Maintainability
- Security
- Testability
- Reusability
- User experience

Choose the highest overall value.

---

# Architecture Reasoning

Verify:

- Layer boundaries
- Package relationships
- Dependency direction
- Interface contracts
- Module ownership

Never violate architecture.

---

# Dependency Reasoning

Before introducing anything new:

Check:

- Existing package
- Existing utility
- Existing component
- Existing interface

Reuse before creating.

---

# Tool Selection

Use the smallest appropriate tool.

Priority:

```
Go Standard Library

↓

Existing Utility

↓

Existing Dependency

↓

New Dependency
```

Avoid unnecessary external libraries.

---

# Agent Selection

Choose the most suitable execution agent.

Examples:

```
Go Engineer

Business logic

Bubble Tea Engineer

UI

Database Engineer

Storage

Git Engineer

Repository operations

Documentation Engineer

Documentation
```

Never assign work to the wrong agent.

---

# Risk Assessment

Evaluate:

- Build failures
- Breaking changes
- Performance impact
- Security impact
- User impact
- Regression risk

High-risk tasks require additional review.

---

# Validation Rules

Every decision must satisfy:

- Architecture
- Coding standards
- Performance
- Security
- Maintainability
- Project rules

---

# Optimization Strategy

Prefer:

- Existing implementation
- Smaller changes
- Reusable modules
- Lower complexity

Avoid:

- Duplicate logic
- Unnecessary abstraction
- Deep nesting
- Large refactors without reason

---

# Error Reasoning

When failures occur:

```
Detect

↓

Analyze

↓

Find Root Cause

↓

Generate Solutions

↓

Select Safest Fix

↓

Validate

↓

Continue
```

Never patch symptoms without understanding the cause.

---

# Conflict Resolution

If multiple valid solutions exist:

Prioritize:

```
Maintainability

↓

Architecture

↓

Performance

↓

Development Speed
```

---

# Performance Reasoning

Analyze:

- CPU usage
- Memory usage
- Rendering cost
- Disk access
- Network usage
- Token usage

Optimize only where necessary.

---

# Security Reasoning

Always verify:

- Input validation
- Permission boundaries
- Sensitive data handling
- File access
- Command execution
- Database safety

Never weaken security for convenience.

---

# User Experience Reasoning

Always consider:

- Mobile devices
- Small screens
- Touch interaction
- Keyboard visibility
- Terminal responsiveness
- Accessibility

---

# Long-Term Reasoning

Every decision should support:

- Future features
- Easier maintenance
- Better testing
- Cleaner architecture
- Reduced technical debt

---

# Failure Handling

If reasoning cannot produce a safe decision:

```
Stop

↓

Identify Missing Information

↓

Request Clarification

↓

Reevaluate
```

Never guess.

---

# Decision Output

The engine produces:

```
Reasoning Summary

Chosen Strategy

Alternative Strategies

Risk Analysis

Agent Assignment

Execution Recommendation

Validation Checklist
```

---

# Validation Checklist

Before approving a decision verify:

- Goal understood
- Context complete
- Architecture preserved
- Dependencies validated
- Risks acceptable
- Agent selected
- Strategy optimized
- Security maintained

---

# Core Rules

1. Think before acting.
2. Never reason without context.
3. Always compare alternatives.
4. Protect architecture.
5. Prefer reuse over creation.
6. Minimize complexity.
7. Validate every decision.
8. Never guess missing information.
9. Optimize for long-term quality.
10. Delegate to the correct agent.

---

# Mission Statement

The Reasoning Engine exists to ensure that every implementation inside TermCode begins with sound engineering judgment rather than immediate execution.

By combining context, memory, architectural knowledge, and structured analysis, it produces reliable, maintainable, secure, and scalable decisions that guide every AI agent toward production-ready outcomes while preserving the long-term integrity of the project.