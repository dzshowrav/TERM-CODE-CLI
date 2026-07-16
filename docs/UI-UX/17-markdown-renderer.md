# 17-markdown-renderer.md

# Markdown Renderer
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Markdown Rendering System used throughout the Mobile AI CLI.

Every AI response, documentation page, README file, markdown preview, tool output, and generated content must be rendered consistently using this specification.

The renderer must preserve readability while remaining completely terminal-native.

---

# Design Goals

The Markdown Renderer must be

- Mobile First
- Terminal Native
- Fast
- Readable
- Streaming Friendly
- CommonMark Compatible
- GitHub-Flavored Markdown Compatible
- Performance Optimized

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Markdown is the default document format inside the application.

Rendering should improve readability without changing the author's intent.

The renderer must preserve document structure while adapting naturally to terminal constraints.

---

# Rendering Pipeline

```
Markdown Source

↓

Lexer

↓

Parser

↓

AST

↓

Renderer

↓

Terminal Output
```

---

# Supported Markdown

Supported

- CommonMark
- GitHub-Flavored Markdown (GFM)

---

# Rendering Flow

```
Input

↓

Parse

↓

AST

↓

Layout

↓

Render

↓

Display
```

---

# Block Elements

Supported

- Headings
- Paragraphs
- Lists
- Ordered Lists
- Task Lists
- Tables
- Block Quotes
- Code Blocks
- Horizontal Rules
- Images (placeholder)
- Links

---

# Inline Elements

Supported

- Bold
- Italic
- Bold Italic
- Inline Code
- Links
- Strike Through
- Escaped Characters

---

# Headings

Supported

```
# H1
## H2
### H3
#### H4
##### H5
###### H6
```

Rendering example

```text
Heading

--------

Subheading
```

Heading hierarchy must remain visually clear.

---

# Paragraphs

Paragraphs wrap naturally.

Paragraph spacing

```
1 Empty Line
```

---

# Line Wrapping

Normal text

Automatically wraps.

Code

Never wraps.

Tables

Never wrap cells automatically.

---

# Lists

Supported

Unordered

```text
- Item
- Item
```

Ordered

```text
1. Item
2. Item
```

Nested lists supported.

---

# Task Lists

Supported

```text
[ ] Todo

[x] Done
```

Interactive editing is optional.

---

# Block Quotes

Example

```text
│ This is a quote.
```

Nested block quotes supported.

---

# Horizontal Rules

Supported

Example

```text
────────────────────
```

---

# Inline Code

Rendered using monospaced formatting.

Example

```text
npm install
```

---

# Code Blocks

Rendered according to

```
16-code-block.md
```

Language labels supported.

Syntax highlighting recommended.

---

# Tables

Supported

Example

```text
Name      Status

App       Ready

API       Running
```

Columns remain aligned.

Horizontal scrolling allowed.

---

# Images

Terminal cannot display images directly.

Render placeholder

```text
[Image]

diagram.png
```

Optional image preview supported.

---

# Links

Display

```text
OpenAI

https://openai.com
```

Links remain selectable.

---

# Automatic URL Detection

Supported.

Plain URLs automatically become links.

---

# Escaping

Supported.

Markdown escape characters must render correctly.

---

# HTML

Raw HTML rendering

Disabled by default.

Unsafe HTML

Ignored.

---

# Emoji

Unicode emoji supported if terminal font supports them.

Renderer must not depend on emoji.

---

# Unicode

Fully supported.

Renderer must correctly display

- UTF-8
- CJK
- RTL
- Symbols

---

# Streaming Markdown

Markdown renders incrementally.

Incomplete markdown remains readable.

Example

```text
## Inst

↓

## Installation
```

---

# Streaming Tables

Rows appear progressively.

Columns remain aligned.

---

# Streaming Code

Code blocks render immediately.

Syntax highlighting updates as needed.

---

# Streaming Lists

Items appear individually.

Existing items remain stable.

---

# Search Highlighting

Matched text may be highlighted.

Layout must remain unchanged.

---

# Folding

Optional.

Large sections may collapse.

Example

```text
Show Section
```

---

# File Preview

Markdown files display directly.

Example

```text
README.md
```

Rendered immediately.

---

# Navigation

Internal headings may support quick navigation.

Optional.

---

# Scrolling

Vertical scrolling

Supported.

Horizontal scrolling

Available for

- Tables
- Code
- Long Links

---

# Typography

Use terminal font only.

No proportional fonts.

---

# Spacing

Paragraph

```
1 Empty Line
```

Heading

```
1 Empty Line Before

1 Empty Line After
```

Lists

Compact spacing.

---

# Safe Area

Markdown remains inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Keyboard Behavior

Opening the keyboard

Reduces visible content.

Rendering continues normally.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Semantic structure preserved.

---

# Performance

Large markdown files

Render progressively.

Virtual rendering recommended.

Only changed lines update.

---

# Error Handling

Invalid markdown

Render best-effort output.

Never crash the renderer.

---

# Security

Never execute

- HTML
- JavaScript
- Embedded scripts

Render only safe markdown.

---

# Restrictions

Never

- Reformat author content
- Remove whitespace unnecessarily
- Execute embedded HTML
- Wrap code automatically
- Break table alignment
- Hide unsupported syntax

---

# Example Document

```text
# Mobile AI CLI

A modern coding assistant.

## Features

- AI Chat
- Tool Execution
- Markdown Rendering

### Installation

npm install

────────────────────

> Mobile-first design.
```

---

# Example Task List

```text
[ ] Design UI

[x] Implement Renderer

[ ] Testing
```

---

# Example Table

```text
Feature      Status

Renderer     Ready

Streaming    Active

Search       Planned
```

---

# Markdown Renderer Checklist

Every Markdown Renderer must

- Support CommonMark
- Support GFM
- Preserve document structure
- Preserve spacing
- Preserve code formatting
- Support streaming
- Support tables
- Support task lists
- Respect safe areas
- Remain terminal-native

---

# Core Rules

1. Markdown is the default document format.
2. Preserve author intent.
3. Never modify code formatting.
4. Support incremental streaming.
5. Support CommonMark and GFM.
6. Render invalid markdown gracefully.
7. Never execute embedded HTML.
8. Keep rendering responsive.
9. Respect mobile safe areas.
10. Optimize for Android Termux.

---

# Summary

The Markdown Renderer provides a complete, terminal-native implementation of CommonMark and GitHub-Flavored Markdown for the Mobile AI CLI. It supports incremental streaming, structured rendering, code blocks, tables, task lists, and safe document display while preserving formatting and performance. Designed specifically for Android Termux, it ensures documentation, AI responses, and project files remain readable, responsive, and consistent across all mobile-first workflows.