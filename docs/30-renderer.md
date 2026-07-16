# 30-renderer.md

# Renderer Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Renderer?

A **Renderer** is the subsystem responsible for transforming internal application data into a visual representation that users can see and interact with.

In an AI CLI Coding Agent, the Renderer converts messages, terminal output, code blocks, markdown, tool status, progress bars, tables, trees, and UI components into a structured Terminal User Interface (TUI).

It is the final presentation layer between the AI Agent and the user.

---

# Why Renderer?

Without Renderer

```
Application

↓

Raw Data

↓

User
```

Problems

- Poor readability
- Inconsistent layout
- Difficult navigation
- No visual hierarchy
- Weak user experience

---

With Renderer

```
Application

↓

Renderer

↓

Layout Engine

↓

TUI Components

↓

Terminal Screen
```

---

# Goals

A production Renderer should provide

- Fast rendering
- Incremental updates
- Layout management
- Component rendering
- Markdown rendering
- Syntax highlighting
- Theme support
- Responsive terminal layout
- Animation support
- Accessibility

---

# High-Level Architecture

```
             Application

                   │

                   ▼

               Renderer

                   │

      ┌────────────┼────────────┐

      ▼            ▼            ▼

 Layout      Components      Theme

      ▼            ▼            ▼

 Diffing     Styling      Terminal

      └────────────┼────────────┘

                   ▼

               Display
```

---

# Folder Structure

```
src/

renderer/

    Renderer.ts

    RenderPipeline.ts

    LayoutEngine.ts

    ComponentRenderer.ts

    MarkdownRenderer.ts

    CodeRenderer.ts

    DiffRenderer.ts

    ThemeRenderer.ts

    ScreenBuffer.ts

    RendererEvents.ts

    RendererMetrics.ts

    RendererValidator.ts
```

---

# Core Components

## Renderer

Central controller.

Responsibilities

- Receive render requests
- Build layouts
- Update screen
- Coordinate rendering pipeline

---

## Render Pipeline

Processes rendering in stages

```
Input

↓

Transform

↓

Layout

↓

Style

↓

Render

↓

Display
```

---

## Layout Engine

Calculates

- Width
- Height
- Position
- Alignment
- Spacing

for every component.

---

## Component Renderer

Renders

```
Panels

Lists

Tables

Trees

Progress Bars

Status Lines

Input Boxes
```

---

## Markdown Renderer

Converts

```
Markdown

↓

Formatted Terminal Output
```

Supports

- Headers
- Lists
- Tables
- Code blocks
- Quotes
- Links

---

## Code Renderer

Displays

```
Source Code

↓

Syntax Highlighted Code
```

Supports multiple programming languages.

---

## Diff Renderer

Displays

```
Added Lines

Removed Lines

Modified Lines
```

Useful for code reviews and patches.

---

## Theme Renderer

Applies

```
Colors

Borders

Icons

Typography
```

according to the active theme.

---

## Screen Buffer

Stores the current rendered frame before writing it to the terminal.

Supports efficient updates.

---

## Renderer Validator

Checks

- Component structure
- Layout validity
- Rendering consistency

---

# Rendering Lifecycle

```
Application State

↓

Renderer

↓

Layout

↓

Theme

↓

Buffer

↓

Terminal Display
```

---

# Render Object

Contains

```
Component

Content

Style

Position

Metadata
```

---

# Supported Components

```
Text

Markdown

Code

Tables

Lists

Trees

Panels

Forms

Progress Bars

Notifications

Status Bars
```

---

# Incremental Rendering

Instead of redrawing the entire screen

```
Old Frame

↓

Diff

↓

Changed Regions

↓

Update Only
```

Improves performance.

---

# Screen Buffer Flow

```
Components

↓

Screen Buffer

↓

Diff

↓

Terminal
```

---

# Layout Types

Examples

```
Vertical

Horizontal

Grid

Split View

Floating Panel
```

---

# Markdown Rendering Flow

```
Markdown

↓

Parser

↓

Renderer

↓

Terminal Output
```

---

# Code Rendering Flow

```
Source Code

↓

Lexer

↓

Syntax Highlighter

↓

Terminal
```

---

# Diff Rendering Flow

```
Old Version

↓

Compare

↓

Colored Diff

↓

Display
```

---

# Theme Integration

Theme controls

```
Foreground

Background

Borders

Highlights

Icons

Spacing
```

---

# Animation Support

Examples

```
Progress Bars

Loading Indicators

Typing Effects

Spinner

Status Updates
```

Animations should be lightweight.

---

# Event Bus Integration

Common events

```
render:start

render:update

render:complete

render:error

screen:refresh
```

---

# State Manager Integration

Renderer reacts automatically when application state changes.

---

# Conversation Manager Integration

Displays

```
Messages

Streaming Responses

Tool Results
```

---

# Stream Handler Integration

Streams partial model output directly into the renderer.

---

# Terminal Engine Integration

Displays

```
stdout

stderr

Interactive Sessions
```

---

# Theme Engine Integration

Renderer requests styling information from the Theme Engine.

---

# Plugin Integration

Plugins may register

- Custom components
- Custom renderers
- New layouts
- Specialized visualizations

---

# Skills Integration

Skills may contribute

- Render templates
- Specialized panels
- Interactive widgets
- Documentation viewers

---

# Cache Strategy

Cache

```
Parsed Markdown

Syntax Tokens

Rendered Components

Theme Data
```

to reduce rendering overhead.

---

# Error Handling

```
Rendering Failure

↓

Fallback Renderer

↓

Plain Text Output

↓

Continue
```

The interface should remain usable even if advanced rendering fails.

---

# Security

Always

- Sanitize rendered content
- Escape control characters
- Validate markdown
- Protect terminal integrity

Never

- Execute embedded content
- Render unsafe escape sequences
- Trust unvalidated plugins

---

# Performance Optimizations

Use

- Incremental rendering
- Double buffering
- Layout caching
- Component reuse
- Background parsing

Avoid

- Full screen redraws
- Re-rendering unchanged components
- Blocking the UI thread

---

# Best Practices

Always

- Separate layout from rendering
- Cache parsed content
- Minimize redraws
- Keep components modular
- Respect terminal size changes
- Support accessibility

Never

- Hardcode layouts
- Ignore terminal resizing
- Mix rendering with business logic
- Block user interaction

---

# Common Mistakes

Bad

```
Application

↓

Print Text

↓

Terminal
```

No structure or visual organization.

---

Good

```
Application

↓

Renderer

↓

Layout Engine

↓

Theme Engine

↓

Screen Buffer

↓

Terminal
```

Fast, structured, and visually consistent.

---

# Testing Checklist

- Layout rendering
- Markdown rendering
- Code rendering
- Diff rendering
- Theme support
- Incremental updates
- Screen resize handling
- Performance
- Accessibility
- Error recovery

---

# Advantages

- Better user experience
- Faster screen updates
- Consistent UI
- Rich terminal components
- Modular rendering pipeline
- Scalable presentation layer

---

# Disadvantages

- Rendering complexity
- Theme maintenance
- Terminal compatibility differences
- Buffer synchronization overhead

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

# Complete Rendering Flow

```
Application State

↓

State Manager

↓

Conversation Manager

↓

Stream Handler

↓

Renderer

↓

Render Pipeline

↓

Layout Engine

↓

Component Renderer

↓

Markdown Renderer

↓

Theme Renderer

↓

Screen Buffer

↓

Diff Renderer

↓

Terminal Display

↓

User
```

---

# Summary

The **Renderer** is the presentation layer responsible for converting application state and structured data into an efficient, interactive, and visually consistent Terminal User Interface.

A production-grade Renderer should include:

- Renderer
- Render Pipeline
- Layout Engine
- Component Renderer
- Markdown Renderer
- Code Renderer
- Diff Renderer
- Theme Renderer
- Screen Buffer
- Renderer Validator
- Event Bus Integration

By combining incremental rendering, layout management, syntax highlighting, markdown support, theme integration, and efficient screen updates, the Renderer enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to deliver responsive, readable, and feature-rich terminal experiences.