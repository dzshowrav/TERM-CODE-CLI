# 14-prompt-builder.md

# Prompt Builder Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Prompt Builder?

A **Prompt Builder** is the subsystem responsible for transforming structured application context into the final prompt that is sent to the Large Language Model (LLM).

The Prompt Builder does **not** generate AI responses.

Instead, it constructs a well-organized, token-efficient, and model-compatible prompt from multiple context sources.

---

# Why Prompt Builder?

Without Prompt Builder

```
User Prompt

↓

LLM
```

Problems

- Poor prompt quality
- Missing instructions
- No context organization
- High token usage
- Inconsistent responses

---

With Prompt Builder

```
Context

↓

Prompt Builder

↓

Optimized Prompt

↓

LLM
```

The model receives a structured instruction package.

---

# Goals

A production Prompt Builder should provide

- Structured prompt generation
- Model-specific formatting
- Token optimization
- Instruction hierarchy
- Context merging
- Prompt templates
- Variable substitution
- Safety instructions
- Tool definitions
- Prompt validation

---

# High-Level Architecture

```
             Context Builder

                   │

                   ▼

            Prompt Builder

                   │

      ┌────────────┼────────────┐

      ▼            ▼            ▼

 Templates    Instructions   Variables

      ▼            ▼            ▼

 Tokenizer    Optimizer     Validator

      └────────────┼────────────┘

                   ▼

            Final Prompt

                   ▼

            Model Router
```

---

# Folder Structure

```
src/

prompt/

    PromptBuilder.ts

    PromptTemplate.ts

    PromptComposer.ts

    PromptOptimizer.ts

    PromptValidator.ts

    PromptVariables.ts

    PromptFormatter.ts

    PromptTokenizer.ts

    PromptMetrics.ts

    PromptEvents.ts

    PromptCache.ts

    PromptPolicies.ts
```

---

# Core Components

## Prompt Builder

Central controller.

Responsibilities

- Build prompt
- Merge sections
- Validate output
- Return final prompt

---

## Prompt Composer

Combines

- System prompt
- Context
- User message
- Tool definitions
- Policies

---

## Prompt Template

Provides reusable layouts.

Examples

```
Coding

Debugging

Review

Planning

Documentation
```

---

## Prompt Optimizer

Reduces

- Duplicate text
- Redundant instructions
- Unused context

---

## Prompt Formatter

Formats prompts according to model requirements.

Examples

```
Markdown

XML

JSON

Plain Text
```

---

## Prompt Variables

Supports placeholders.

Examples

```
{{workspace}}

{{language}}

{{framework}}

{{model}}

{{current_file}}
```

---

## Prompt Validator

Checks

- Token limit
- Required sections
- Missing variables
- Invalid formatting

---

## Prompt Cache

Stores recently generated prompts.

Benefits

- Faster rebuilds
- Reduced processing

---

# Prompt Lifecycle

```
Receive Context

↓

Select Template

↓

Merge Sections

↓

Resolve Variables

↓

Optimize

↓

Validate

↓

Return Prompt
```

---

# Prompt Structure

Typical prompt layout

```
System Instructions

↓

Policies

↓

Workspace Context

↓

Memory

↓

Conversation

↓

Files

↓

Tool Definitions

↓

User Request
```

---

# Prompt Sections

## System Prompt

Defines

- Agent identity
- Rules
- Behavior
- Constraints

---

## Context Section

Contains

- Workspace
- Memory
- Session
- Files

---

## Tool Section

Defines

- Available tools
- Parameters
- Permissions

---

## User Section

Contains the latest user request.

---

# Variable Resolution

Example

Template

```
Current Framework:

{{framework}}
```

Resolved

```
Current Framework:

Laravel
```

---

# Instruction Hierarchy

Priority

```
System

↓

Developer

↓

Policies

↓

Skills

↓

Plugins

↓

User
```

Higher-priority instructions override lower-priority ones.

---

# Prompt Templates

Examples

```
General Chat

Coding

Refactoring

Bug Fixing

Architecture

Documentation

Testing

Deployment
```

---

# Token Optimization

Strategies

- Remove duplicates
- Compress history
- Summarize long context
- Prioritize relevant sections

---

# Prompt Validation

Check

- Size
- Completeness
- Variable resolution
- Required instructions
- Formatting consistency

---

# Model Compatibility

Adapt prompts for

```
OpenAI

Anthropic

Google Gemini

OpenRouter

Local LLMs
```

Each model may require different formatting.

---

# Event Bus Integration

Common events

```
prompt:start

prompt:build

prompt:optimize

prompt:validate

prompt:complete

prompt:error
```

---

# Context Builder Integration

```
Context Builder

↓

Prompt Builder
```

The Prompt Builder never gathers context directly.

---

# Model Router Integration

```
Prompt Builder

↓

Model Router

↓

LLM
```

---

# Tool Manager Integration

Inject

```
Tool Schemas

Capabilities

Usage Rules
```

into the prompt when needed.

---

# Skills Integration

Skills may contribute

```
Templates

Instructions

Examples
```

before prompt generation.

---

# Plugin Integration

Plugins may extend

- Prompt templates
- Variables
- Formatting rules
- Validation

---

# Session Integration

Use

```
Conversation

Workspace

Preferences
```

when building prompts.

---

# Cache Strategy

Cache prompts using

```
Context Hash

+

Template

+

Model
```

to avoid unnecessary rebuilding.

---

# Error Handling

```
Missing Variable

↓

Default Value

↓

Continue
```

or

```
Invalid Prompt

↓

Validation Error

↓

Rebuild
```

---

# Security

Always

- Remove secrets
- Validate variables
- Escape unsafe content
- Respect workspace boundaries

Never

- Inject private credentials
- Include unrelated context
- Skip validation

---

# Performance Optimizations

Use

- Prompt cache
- Incremental updates
- Lazy formatting
- Parallel variable resolution
- Efficient token counting

Avoid

- Rebuilding identical prompts
- Duplicating context
- Unnecessary formatting passes

---

# Best Practices

Always

- Use templates
- Separate context from formatting
- Optimize tokens
- Validate prompts
- Support multiple models

Never

- Hardcode prompts
- Mix prompt logic with context collection
- Ignore model limits
- Duplicate instructions

---

# Common Mistakes

Bad

```
Concatenate Strings

↓

LLM
```

Hard to maintain.

---

Good

```
Template

↓

Composer

↓

Optimizer

↓

Validator

↓

LLM
```

Structured and reusable.

---

# Testing Checklist

- Template selection
- Variable resolution
- Prompt optimization
- Validation
- Token limits
- Model compatibility
- Tool injection
- Plugin extensions
- Cache
- Error handling

---

# Advantages

- Consistent prompts
- Better AI quality
- Lower token usage
- Easier maintenance
- Multi-model support
- Extensible architecture

---

# Disadvantages

- Template management
- Additional processing
- Variable maintenance
- Model-specific complexity

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Agents
- Enterprise AI platforms

---

# Complete Prompt Flow

```
Context Builder

↓

Prompt Template

↓

Composer

↓

Variable Resolver

↓

Optimizer

↓

Validator

↓

Final Prompt

↓

Model Router

↓

LLM
```

---

# Summary

The **Prompt Builder** converts structured application context into an optimized, model-ready prompt.

A production-grade Prompt Builder should include:

- Prompt Builder
- Composer
- Template System
- Variable Resolver
- Formatter
- Optimizer
- Validator
- Cache
- Event Bus Integration
- Multi-model compatibility

By separating prompt construction from context collection and model execution, the Prompt Builder creates consistent, efficient, and maintainable prompts that maximize AI performance while minimizing token usage, forming a critical layer in modern AI Coding Agents such as OpenCode and Antigravity CLI.