# 16-code-block.md

# Code Block
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Code Block system used throughout the Mobile AI CLI.

Code Blocks provide a consistent way to display, stream, copy, scroll, and review source code, configuration files, shell commands, logs, and terminal output while preserving exact formatting.

Every code-related component must follow this specification.

---

# Design Goals

The Code Block system must be

- Mobile First
- Terminal Native
- Monospaced
- Readable
- Copy Friendly
- Streaming Optimized
- Performance Friendly
- Accessible

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Code is first-class content.

Formatting is part of the code.

The renderer must preserve every visible character exactly.

---

# Display Location

Code Blocks appear only inside

- Conversation
- Tool Output
- File Preview
- Markdown Renderer

They never open a separate screen.

---

# Layout

```text
┌────────────────────────────────────┐
│ JavaScript                         │
├────────────────────────────────────┤
│ function hello() {                 │
│     console.log("Hello");          │
│ }                                  │
└────────────────────────────────────┘
```

---

# Structure

Each Code Block contains

- Language
- Source Code
- Optional Metadata

---

# Language Label

Displayed above the code.

Examples

```text
JavaScript
```

```text
TypeScript
```

```text
JSON
```

```text
Shell
```

---

# Supported Languages

Examples

- JavaScript
- TypeScript
- JSX
- TSX
- HTML
- CSS
- SCSS
- JSON
- YAML
- TOML
- XML
- Markdown
- Bash
- Shell
- Python
- PHP
- Go
- Rust
- Java
- Kotlin
- Swift
- SQL
- Dockerfile
- Git Diff
- Plain Text

---

# Font

Always

```
Monospaced
```

Never use proportional fonts.

---

# Formatting

Preserve

- Spaces
- Tabs
- Blank Lines
- Indentation
- Line Breaks
- Unicode Characters

Never automatically reformat code.

---

# Syntax Highlighting

Supported.

Highlighting is language-aware.

If language detection fails

Use plain text.

---

# Horizontal Scrolling

Supported.

Long lines remain on a single line.

Never wrap code automatically.

---

# Vertical Scrolling

Large code blocks are vertically scrollable.

Conversation scrolling remains independent.

---

# Line Numbers

Optional.

When enabled

Displayed on the left.

Example

```text
1  function hello() {
2      console.log("Hello");
3  }
```

---

# Selection

Supported.

Users may select

- Single line
- Multiple lines
- Entire block

Selection preserves formatting.

---

# Copy

Supported.

Copied content must match the original exactly.

---

# Streaming

Streaming code appears incrementally.

Previously rendered lines remain unchanged.

---

# Incremental Rendering

Only newly received lines are rendered.

Never redraw the entire block.

---

# Diff Rendering

Supported.

Example

```text
Added

Removed

Modified
```

Visual differences must remain readable.

---

# File Metadata

Optional.

Display

```text
src/App.tsx
```

or

```text
package.json
```

Above the language label.

---

# Shell Commands

Display exactly as entered.

Example

```text
npm install
```

Never modify shell syntax.

---

# Logs

Display

- Standard Output
- Standard Error
- Build Logs

Formatting remains unchanged.

---

# JSON

Preserve indentation.

Maintain key ordering.

---

# Markdown Source

Render as code when inside a Code Block.

Do not interpret markdown syntax.

---

# Search Highlighting

Matched text may be highlighted.

Code layout must remain unchanged.

---

# Folding

Optional.

Large Code Blocks may collapse.

Collapsed example

```text
Show Code
```

Expanded example

Displays the full content.

---

# Performance

Only visible lines should be rendered.

Large files should use virtual scrolling.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Syntax highlighting must not be the only source of meaning.

---

# Keyboard Behavior

Keyboard opening

Does not modify code formatting.

Horizontal scrolling remains available.

---

# Safe Area

Code Blocks remain inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Error Handling

If rendering fails

Display

```text
Unable to display code.
```

Never corrupt the original content.

---

# Security

Never execute displayed code automatically.

Never modify copied content.

Never hide dangerous commands.

---

# Restrictions

Never

- Wrap long code lines
- Remove indentation
- Convert tabs automatically
- Reorder lines
- Change file formatting
- Hide syntax errors
- Alter whitespace

---

# Example JavaScript

```text
function hello() {
    console.log("Hello World");
}
```

---

# Example JSON

```text
{
    "name": "mobile-cli",
    "version": "1.0.0"
}
```

---

# Example Shell

```text
npm create vite@latest
```

---

# Example File Preview

```text
package.json

JSON

{
    "name": "mobile-cli"
}
```

---

# Code Block Checklist

Every Code Block must

- Preserve formatting
- Preserve indentation
- Preserve blank lines
- Support syntax highlighting
- Support horizontal scrolling
- Support vertical scrolling
- Support copying
- Support streaming
- Respect safe areas
- Remain terminal-native

---

# Core Rules

1. Code formatting is never modified.
2. Use monospaced fonts only.
3. Never wrap long lines automatically.
4. Preserve whitespace exactly.
5. Support incremental streaming.
6. Syntax highlighting is optional but recommended.
7. Copy operations preserve original formatting.
8. Large files use virtual scrolling.
9. Code Blocks remain inside the Conversation Area.
10. Optimize rendering for Android Termux.

---

# Summary

The Code Block system provides a reliable, terminal-native method for displaying source code, commands, logs, configuration files, and structured text within the Mobile AI CLI. By preserving formatting, supporting incremental streaming, enabling efficient scrolling, and maintaining exact copy fidelity, it delivers a professional coding experience optimized for Android Termux and mobile-first AI-assisted development.