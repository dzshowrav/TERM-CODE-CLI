# 06-typography.md

# Typography
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete typography system for the Mobile AI CLI.

Typography is the primary communication layer of the interface.

Since the application runs inside a terminal environment, text must always prioritize readability, consistency, and information density over visual decoration.

Every visible character must follow this specification.

---

# Design Goals

Typography must be

- Mobile First
- Terminal Native
- Highly Readable
- Monospaced
- Consistent
- Accessible
- Performance Friendly

---

# Supported Platform

Primary

- Android
- Termux

Rendering

- Terminal
- ANSI Compatible
- Unicode Compatible

---

# Typography Philosophy

Text is the interface.

Users spend nearly all of their time reading.

Typography therefore has the highest visual priority.

---

# Typography Hierarchy

```
Heading

↓

Section

↓

Body

↓

Code

↓

Metadata

↓

Hint
```

---

# Font Family

Preferred

```
Monospace
```

Recommended fonts

- JetBrains Mono
- Cascadia Mono
- Fira Code
- IBM Plex Mono
- DejaVu Sans Mono
- Noto Sans Mono
- Hack
- Source Code Pro

---

# Font Rules

Always

- Monospaced
- Unicode Compatible
- Terminal Friendly

Never

- Serif
- Decorative Fonts
- Script Fonts
- Handwritten Fonts

---

# Font Weight

Supported

```
Regular

Medium

Bold
```

Avoid excessive weight variation.

---

# Font Size

Terminal controls the physical font size.

The UI defines only logical hierarchy.

Never assume a fixed pixel size.

---

# Line Height

Recommended

```
1 Line
```

Additional spacing should be created using layout,

not typography.

---

# Letter Spacing

```
Normal
```

Never modify character spacing.

---

# Word Spacing

Use

```
Single Space
```

Never insert multiple spaces for alignment.

Use layout instead.

---

# Text Alignment

Default

```
Left
```

Allowed

```
Center
```

Only for

- Splash
- Empty State
- Loading

Never right-align conversation text.

---

# Text Direction

Supported

```
Left to Right
```

Unicode text should render correctly.

---

# Heading

Purpose

Major section titles.

Examples

- Settings
- Workspace
- Help

Use bold when supported.

---

# Section Title

Purpose

Group related content.

Must remain visually distinct.

---

# Body Text

Purpose

Conversation

Documentation

Descriptions

Primary readable content.

---

# Code Text

Purpose

Source code

Commands

Terminal output

Always monospaced.

Never wrap individual tokens.

---

# Inline Code

Purpose

Commands inside normal text.

Must remain visually distinguishable.

---

# Metadata

Purpose

Secondary information.

Examples

- Timestamp
- Token count
- Model
- Path

Should use muted color.

---

# Hint Text

Purpose

Placeholder

Guidance

Suggestions

Must never compete with primary content.

---

# Placeholder

Used inside

Input

Search

Command Palette

Always use muted styling.

---

# Conversation Typography

User

Normal weight.

Assistant

Normal weight.

System

Muted.

Tool Output

Monospaced.

---

# Code Blocks

Use

Monospaced

Preserve

Indentation

Spacing

Blank lines

Never reflow code.

---

# Tables

Columns

Must align perfectly.

Padding

Consistent.

Never rely on proportional fonts.

---

# Tree View

Indentation

Fixed width.

Example

```
Workspace

    src

        components

        pages

    package.json
```

---

# Markdown

Typography follows markdown semantics.

Examples

Heading

Body

List

Quote

Code

Table

---

# Quotes

Purpose

Highlight referenced content.

Maintain clear visual separation.

---

# Links

Must be distinguishable.

Never rely on underline alone.

---

# Terminal Commands

Always display exactly as typed.

Never auto-correct.

Never reformat.

---

# File Paths

Display exactly.

Preserve

Case

Separators

Spacing

---

# Numbers

Keep aligned.

Avoid proportional formatting.

---

# Truncation

When necessary

```
End Truncation
```

Example

```
very_long_file...
```

Never truncate from the beginning.

---

# Wrapping

Conversation

Wrap

Code

Do not wrap individual tokens.

Tables

Horizontal scrolling if required.

---

# Cursor

Always clearly visible.

Cursor must never overlap text.

---

# Selection

Selected text remains readable.

Selection should preserve syntax highlighting whenever possible.

---

# Input Typography

Input uses

Monospaced font.

Supports

Multi-line

Unicode

Paste

Selection

Cursor

---

# Status Bar Typography

Compact.

Readable.

Metadata oriented.

Avoid unnecessary labels.

---

# Dialog Typography

Title

Bold.

Body

Regular.

Action

Clear.

---

# Search Results

Match

Highlighted.

Path

Muted.

Filename

Primary.

---

# Accessibility

Support

Large fonts.

Screen readers.

High contrast.

Reduced motion.

Never communicate meaning through typography alone.

---

# Performance

Typography rendering should

- Cache measurements
- Avoid unnecessary reflow
- Preserve terminal alignment

---

# Responsive Rules

Small screens

Wrap paragraphs.

Never wrap

Commands

Paths

Identifiers

Unless unavoidable.

---

# Restrictions

Never use

- Decorative fonts
- Italic paragraphs
- Excessive bold text
- Random capitalization
- Variable-width fonts
- Justified alignment

---

# Visual Priority

```
Heading

↓

Assistant

↓

User

↓

Tool Output

↓

Metadata

↓

Hint
```

---

# Example Conversation

```
Assistant

Project initialized successfully.

Run the following command.

npm install

The workspace is now ready.
```

---

# Example Code Block

```
function hello() {
    console.log("Hello");
}
```

Formatting must remain unchanged.

---

# Typography Checklist

Every screen must

- Use monospaced text
- Preserve alignment
- Preserve code formatting
- Support Unicode
- Support wrapping
- Support accessibility
- Maintain readability
- Avoid decorative styles

---

# Core Rules

1. Monospaced fonts only.
2. Typography communicates information.
3. Preserve exact code formatting.
4. Never modify character spacing.
5. Never replace terminal formatting.
6. Wrap conversation, not code.
7. Maintain consistent hierarchy.
8. Accessibility always takes priority.
9. Readability is more important than density.
10. Typography must remain terminal-native.

---

# Summary

The Typography System defines a clean, monospaced, terminal-first text architecture for the Mobile AI CLI.

It ensures consistent rendering of conversations, commands, code, markdown, tables, and metadata while maintaining maximum readability, accessibility, and compatibility with Android Termux and mobile-first interaction patterns.