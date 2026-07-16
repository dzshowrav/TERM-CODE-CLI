# 06-plugin-loader.md

# Plugin Loader Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Plugin?

A **Plugin** is an installable software module that extends the capabilities of an AI CLI without modifying the core application.

Unlike a Skill, which mainly provides knowledge and instructions, a Plugin can execute code, register services, add commands, contribute UI components, expose tools, subscribe to events, and integrate with external systems.

---

# Skill vs Plugin

| Skill | Plugin |
|--------|---------|
| Knowledge | Executable Code |
| Prompt Instructions | Runtime Features |
| Markdown/Templates | JavaScript/TypeScript Modules |
| Loaded into Context | Loaded into Runtime |
| Teaches the Agent | Extends the Application |

---

# Why Plugin Loader?

Without a Plugin Loader

```
Core Application

↓

Every Feature Built-In

↓

Large Codebase

↓

Hard To Maintain
```

---

With a Plugin Loader

```
Core

↓

Plugin Loader

↓

Installed Plugins

↓

New Features
```

The application remains small while gaining unlimited extensibility.

---

# Goals

A production Plugin Loader should provide

- Dynamic loading
- Safe isolation
- Version management
- Dependency resolution
- Lifecycle management
- Event integration
- Tool registration
- Command registration
- Hot reload (optional)
- Permission control

---

# High-Level Architecture

```
                  Startup
                     │
                     ▼
              Plugin Manager
                     │
          ┌──────────┼──────────┐
          ▼          ▼          ▼
      Registry     Loader    Validator
          │          │          │
          └──────┬───┴──────┬───┘
                 ▼          ▼
          Runtime Manager  Event Bus
                 │
                 ▼
         Loaded Plugins
                 │
    ┌────────────┼────────────┐
    ▼            ▼            ▼
 Commands      Tools       UI Hooks
```

---

# Folder Structure

```
plugins/

    github/

        manifest.json

        index.ts

        commands/

        tools/

        hooks/

        assets/

    docker/

    figma/

src/

plugins/

    PluginManager.ts

    PluginLoader.ts

    PluginRegistry.ts

    PluginRuntime.ts

    PluginManifest.ts

    PluginValidator.ts

    PluginSandbox.ts

    PluginPermissions.ts

    PluginHooks.ts

    PluginEvents.ts

    PluginMetrics.ts

    PluginCache.ts

    PluginInstaller.ts

    PluginUpdater.ts
```

---

# Core Components

## Plugin Manager

Central controller.

Responsibilities

- Discover plugins
- Load plugins
- Enable plugins
- Disable plugins
- Update plugins
- Remove plugins

---

## Plugin Registry

Stores

```
Installed Plugins

Versions

Status

Permissions

Dependencies

Metadata
```

---

## Plugin Loader

Loads plugin code into memory.

Responsibilities

- Read manifest
- Import module
- Validate exports
- Initialize runtime

---

## Plugin Runtime

Executes plugin lifecycle.

Responsibilities

- Initialize
- Activate
- Execute
- Deactivate
- Dispose

---

## Plugin Validator

Checks

- Manifest
- API version
- Required files
- Dependencies
- Digital signature (optional)

---

## Plugin Sandbox

Provides runtime isolation.

Possible restrictions

- Filesystem
- Network
- Environment
- Terminal
- Process execution

---

## Plugin Permissions

Controls access to

```
Filesystem

Network

Git

Terminal

Clipboard

Environment

MCP

Workspace
```

---

# Plugin Lifecycle

```
Install

↓

Validate

↓

Register

↓

Load

↓

Initialize

↓

Activate

↓

Running

↓

Deactivate

↓

Unload
```

---

# Startup Flow

```
Application Starts

↓

Plugin Manager

↓

Discover Plugins

↓

Validate

↓

Load

↓

Initialize

↓

Ready
```

---

# Plugin Structure

```
plugin/

    manifest.json

    index.ts

    commands/

    tools/

    hooks/

    views/

    assets/

    README.md
```

---

# Manifest

Contains

```
Name

ID

Version

Description

Author

Dependencies

Permissions

Keywords

API Version
```

---

# Entry Module

Example responsibilities

- Register commands
- Register tools
- Subscribe events
- Initialize services
- Cleanup resources

---

# Command Registration

Plugin may provide

```
/deploy

/test

/review

/github
```

Commands become available immediately after activation.

---

# Tool Registration

Plugins may expose

```
deploy()

publish()

search()

review()

sync()
```

The Tool Manager registers them.

---

# Hook System

Hooks allow plugins to extend application behavior.

Examples

```
Before Prompt

After Prompt

Before Tool

After Tool

Before Render

After Render

Startup

Shutdown
```

---

# Event Bus Integration

Common events

```
plugin:install

plugin:load

plugin:activate

plugin:disable

plugin:update

plugin:error
```

Plugins can also subscribe to

```
agent:start

agent:stream

theme:update

mcp:connect

session:save
```

---

# Agent Integration

```
Agent

↓

Plugin Hook

↓

Modify Prompt

↓

Continue Execution
```

---

# Skills Integration

Plugins may

- Install skills
- Remove skills
- Register skill sources
- Update skill repositories

---

# MCP Integration

Plugins may

- Register MCP servers
- Add transports
- Add authentication
- Register external tools

---

# Dependency Resolution

Example

```
Plugin A

↓

Requires

↓

Plugin B

↓

Load B First
```

---

# Version Compatibility

Plugin declares

```
Minimum API

Maximum API

Supported Version
```

Loader validates compatibility.

---

# Configuration

Plugins may store

```
Settings

Tokens

Endpoints

Feature Flags
```

Configuration should be isolated per plugin.

---

# Hot Reload (Optional)

Development flow

```
Source Changed

↓

Unload Plugin

↓

Reload

↓

Reactivate
```

Useful during development.

---

# Error Handling

```
Plugin Crash

↓

Catch

↓

Disable Plugin

↓

Log Error

↓

Continue Application
```

One plugin must never crash the core application.

---

# Security

Always

- Validate manifests
- Restrict permissions
- Isolate execution
- Verify compatibility
- Log privileged actions

Never

- Execute unknown code automatically
- Grant unrestricted filesystem access
- Ignore version mismatches
- Share internal state directly

---

# Performance Optimizations

Use

- Lazy loading
- Plugin cache
- Background initialization
- Dependency graph
- Incremental updates
- Deferred activation

Avoid

- Loading every plugin immediately
- Blocking startup
- Duplicate initialization

---

# Best Practices

Always

- Keep plugins independent
- Expose stable APIs
- Version every release
- Request minimum permissions
- Clean resources on unload
- Document configuration
- Emit lifecycle events

Never

- Modify core application directly
- Depend on internal private APIs
- Leak memory
- Block the main thread

---

# Common Mistakes

Bad

```
Plugin

↓

Directly Modify Core

↓

Application Breaks
```

---

Good

```
Plugin

↓

Public API

↓

Plugin Loader

↓

Core
```

Everything flows through stable interfaces.

---

# Testing Checklist

- Installation
- Validation
- Activation
- Deactivation
- Update
- Dependency resolution
- Permission enforcement
- Event handling
- Command registration
- Tool registration
- Error recovery

---

# Example Plugin Categories

Examples

- GitHub
- GitLab
- Docker
- Kubernetes
- Figma
- Jira
- Slack
- Notion
- AWS
- Azure
- Google Cloud
- PostgreSQL
- Redis
- Browser Automation
- CI/CD
- Code Review
- Deployment
- Monitoring

---

# Advantages

- Highly extensible
- Smaller core
- Easier maintenance
- Community ecosystem
- Feature isolation
- Runtime customization
- Faster development
- Reusable integrations
- Better scalability

---

# Disadvantages

- Requires API stability
- Dependency management
- Security considerations
- Version compatibility challenges

---

# Used In

- OpenCode
- VS Code
- JetBrains IDEs
- Obsidian
- Neovim
- Eclipse
- IntelliJ Platform
- Cursor
- Enterprise AI platforms

---

# Summary

The **Plugin Loader** is the runtime extension system that allows an AI CLI Coding Agent to grow without increasing the complexity of its core.

A production-grade Plugin Loader should include:

- Plugin Manager
- Registry
- Loader
- Runtime
- Validator
- Sandbox
- Permission System
- Dependency Resolver
- Lifecycle Manager
- Event Bus Integration
- Tool Registration
- Command Registration

By keeping plugins isolated, versioned, permission-aware, and event-driven, the application becomes scalable, maintainable, and capable of supporting a rich ecosystem of third-party extensions similar to OpenCode, VS Code, and other modern developer platforms.