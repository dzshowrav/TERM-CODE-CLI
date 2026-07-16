# 05-skills-system.md

# Skills System Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Skill?

A **Skill** is a reusable, self-contained capability that teaches an AI agent how to perform a specific task.

A Skill is **not** the AI model itself.

Instead, it provides:

- Knowledge
- Instructions
- Rules
- Workflows
- Templates
- Best practices
- Tool requirements

The Agent dynamically loads and executes Skills when they are relevant to the user's request.

---

# Why Skills?

Without Skills

```
Agent

↓

Huge System Prompt

↓

Everything Mixed Together
```

Problems

- Difficult maintenance
- Large prompts
- Poor scalability
- Hard to update
- Context waste

---

With Skills

```
User Request

↓

Skill Resolver

↓

Relevant Skills

↓

Agent

↓

Response
```

Only the required knowledge is loaded.

---

# Goals

A production Skills System should provide

- Modular knowledge
- Dynamic loading
- Context-aware selection
- Versioning
- Dependencies
- Tool declarations
- Permission rules
- Fast loading
- Plugin compatibility
- Easy distribution

---

# High-Level Architecture

```
                 User
                  │
                  ▼
            Agent Engine
                  │
                  ▼
           Skill Manager
                  │
      ┌───────────┼────────────┐
      ▼           ▼            ▼
 Skill Registry Resolver    Loader
      │           │            │
      └──────┬────┴─────┬──────┘
             ▼          ▼
      Skill Runtime  Context Builder
             │
             ▼
         Prompt Builder
             │
             ▼
           AI Model
```

---

# Folder Structure

```
skills/

    laravel/

        skill.md

        manifest.json

        prompts/

        templates/

        examples/

        rules/

        workflows/

        resources/

    react/

    nodejs/

    docker/

    git/

src/

skills/

    SkillManager.ts

    SkillRegistry.ts

    SkillLoader.ts

    SkillResolver.ts

    SkillRuntime.ts

    SkillManifest.ts

    SkillValidator.ts

    SkillCache.ts

    SkillExecutor.ts

    SkillEvents.ts

    SkillMetrics.ts
```

---

# Core Components

## Skill Manager

Central controller.

Responsibilities

- Load skills
- Unload skills
- Cache skills
- Resolve dependencies
- Execute skills

---

## Skill Registry

Stores

```
Installed Skills

Versions

Categories

Metadata

Dependencies
```

---

## Skill Loader

Reads skills from disk.

Responsibilities

- Parse manifests
- Load markdown
- Load templates
- Validate structure

---

## Skill Resolver

Chooses which skills should be loaded.

Example

```
User

↓

"Create Laravel API"

↓

Resolver

↓

Laravel Skill

+

PHP Skill

+

REST API Skill
```

---

## Skill Runtime

Executes active skills.

Responsibilities

- Merge instructions
- Provide templates
- Apply rules
- Inject prompts

---

## Skill Cache

Stores frequently used skills in memory.

Benefits

- Faster startup
- Reduced disk access
- Better performance

---

## Skill Validator

Checks

- Manifest
- Version
- Dependencies
- Required files
- Schema

---

# Skill Lifecycle

```
Install

↓

Validate

↓

Register

↓

Resolve

↓

Load

↓

Execute

↓

Unload

↓

Cache
```

---

# Skill Resolution Flow

```
User Request

↓

Intent Detection

↓

Resolver

↓

Matching Skills

↓

Load

↓

Prompt Builder

↓

Model
```

---

# Typical Skill Structure

```
laravel/

    skill.md

    manifest.json

    prompts/

    templates/

    workflows/

    rules/

    examples/

    snippets/

    resources/
```

---

# Manifest

A manifest describes the skill.

Contains

```
Name

Version

Description

Author

Dependencies

Category

Tools

Permissions

Keywords
```

---

# skill.md

Contains

- Instructions
- Best practices
- Coding conventions
- Standards
- Examples
- Warnings

This becomes part of the prompt context.

---

# Templates

Reusable files.

Examples

```
Controller

Migration

Dockerfile

README

Config

Workflow
```

---

# Rules

Examples

```
Always use TypeScript

Never use any

Prefer async/await

Use dependency injection
```

---

# Workflows

Describe step-by-step execution.

Example

```
Create Project

↓

Install Packages

↓

Generate Files

↓

Run Tests

↓

Return Result
```

---

# Examples

Provide reference implementations.

Examples

```
CRUD

REST API

Authentication

Pagination
```

---

# Dependencies

One skill may require another.

Example

```
Laravel Skill

↓

Requires

↓

PHP Skill
```

Resolver loads both.

---

# Categories

Examples

```
Language

Framework

Database

Cloud

DevOps

Testing

Documentation

Security

Frontend

Backend
```

---

# Tool Requirements

A skill may request tools.

Example

```
Filesystem

Git

Terminal

Docker

Browser
```

The Agent checks availability before execution.

---

# Prompt Injection Flow

```
System Prompt

↓

Conversation

↓

Resolved Skills

↓

Templates

↓

Rules

↓

User Request

↓

Final Prompt
```

---

# Event Bus Integration

Common events

```
skill:install

skill:load

skill:resolve

skill:execute

skill:unload

skill:error
```

---

# Plugin Integration

Plugins may

- Install skills
- Update skills
- Remove skills
- Register categories
- Add templates

---

# MCP Integration

Skills may declare

```
Requires

↓

Git Tool

Filesystem Tool

Database Tool

Browser Tool
```

The Agent obtains them through MCP.

---

# Versioning

Example

```
1.0.0

1.1.0

2.0.0
```

Supports upgrades without breaking compatibility.

---

# Permission Model

Skill may request

```
Filesystem

Terminal

Git

Network

Environment
```

The Agent decides whether to grant access.

---

# Security

Always

- Verify source
- Validate manifest
- Check signatures (optional)
- Restrict permissions
- Prevent arbitrary execution

Never

- Execute unknown scripts automatically
- Trust unverified skills
- Ignore dependency conflicts

---

# Performance Optimizations

Use

- Lazy loading
- Memory cache
- Dependency cache
- Manifest indexing
- Incremental updates
- Background validation

Avoid

- Loading every skill at startup
- Parsing markdown repeatedly
- Duplicate dependency loading

---

# Best Practices

Always

- Keep skills focused
- Write reusable instructions
- Version every release
- Separate templates from rules
- Keep manifests small
- Use descriptive keywords
- Document dependencies

Never

- Put everything into one skill
- Duplicate templates
- Hardcode project paths
- Ignore compatibility

---

# Common Mistakes

Bad

```
One Skill

↓

Everything

Laravel

React

Python

Docker

Kubernetes

Git

AWS
```

Impossible to maintain.

---

Good

```
Laravel

React

Node

Docker

Git

Testing

Security
```

Each skill has one responsibility.

---

# Testing Checklist

- Install
- Validation
- Dependency resolution
- Loading
- Prompt injection
- Template loading
- Rule application
- Cache
- Version upgrade
- Removal
- Error recovery

---

# Example Skill Categories

Examples

- Laravel
- React
- Vue
- Angular
- Node.js
- Express
- PHP
- Python
- Go
- Rust
- Docker
- Kubernetes
- Git
- SQL
- MongoDB
- Redis
- GraphQL
- Tailwind CSS
- Flutter
- Android
- AWS
- Azure
- GCP

---

# Advantages

- Modular architecture
- Reusable knowledge
- Smaller prompts
- Faster execution
- Easier maintenance
- Better scalability
- Team sharing
- Plugin friendly
- Version controlled

---

# Disadvantages

- Dependency management required
- Version compatibility must be maintained
- Poorly designed skills can conflict
- Skill discovery adds complexity

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Continue.dev
- Cursor AI
- Enterprise AI Agents
- Internal developer assistants

---

# Summary

The **Skills System** enables an AI Coding Agent to acquire specialized capabilities without increasing the complexity of the core Agent Engine.

A production-grade Skills System should include:

- Skill Manager
- Registry
- Resolver
- Loader
- Runtime
- Manifest
- Templates
- Rules
- Workflows
- Versioning
- Permission model
- Event Bus integration

By keeping knowledge modular and loading only the skills needed for the current task, the Agent remains efficient, extensible, and easy to maintain while supporting a continuously growing ecosystem of reusable capabilities.