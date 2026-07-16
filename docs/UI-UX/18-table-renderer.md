# 18-table-renderer.md

# Table Renderer
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Table Rendering System used throughout the Mobile AI CLI.

Tables are used to present structured information such as search results, command summaries, file lists, configuration values, API responses, benchmark results, and AI-generated comparisons.

The renderer must preserve alignment, readability, and terminal-native behavior across all supported Android Termux environments.

---

# Design Goals

The Table Renderer must be

- Mobile First
- Terminal Native
- Responsive
- Readable
- Streaming Friendly
- Keyboard Safe
- High Performance
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

A table should remain understandable regardless of screen width.

The renderer prioritizes

1. Readability
2. Alignment
3. Performance

before visual decoration.

---

# Rendering Pipeline

```
Table Data

↓

Column Analysis

↓

Width Calculation

↓

Layout Engine

↓

Terminal Renderer

↓

Display
```

---

# Supported Sources

Tables may originate from

- Markdown
- AI Responses
- Tool Output
- Database Queries
- CSV Preview
- JSON Conversion
- API Responses
- Git Information

---

# Table Structure

Each table contains

- Header
- Body
- Optional Footer

---

# Basic Example

```text
Name          Status

frontend      Ready

backend       Running

database      Connected
```

---

# Header

Always displayed first.

Headers use stronger visual emphasis.

---

# Body

Contains all data rows.

Rows maintain consistent spacing.

---

# Footer

Optional.

Example

```text
3 Rows
```

---

# Alignment

Supported

Left

Center

Right

Default

```
Left
```

---

# Column Width

Automatically calculated.

Based on

- Header length
- Longest value
- Available terminal width

---

# Dynamic Width

Columns resize automatically when

- Terminal width changes
- Orientation changes
- Font size changes

---

# Minimum Width

Each column must remain readable.

Very narrow columns should be truncated intelligently.

---

# Maximum Width

Long content should not force the table beyond practical limits.

---

# Overflow Handling

If total width exceeds the terminal

Enable

Horizontal Scrolling

Never wrap table cells automatically.

---

# Horizontal Scrolling

Supported.

Entire rows move together.

Column alignment never changes.

---

# Vertical Scrolling

Large tables scroll vertically.

Headers may remain visible if supported.

---

# Wrapping

Allowed

Paragraph text outside tables.

Not Allowed

Inside table cells.

---

# Truncation

Long text may be shortened.

Example

```text
Very Long File Name...

```

Full value remains available through selection or expansion.

---

# Empty Cells

Render as blank.

Maintain column alignment.

---

# Missing Data

Display

```text
—
```

or

```text
N/A
```

---

# Numeric Columns

Right aligned.

Example

```text
Files

120

45

3
```

---

# Text Columns

Left aligned.

---

# Boolean Values

Example

```text
Enabled

Disabled
```

---

# Status Columns

Examples

```text
Ready

Running

Failed

Offline

Pending
```

Status remains readable without relying solely on color.

---

# Sorting

Optional.

Current sort direction may be indicated.

---

# Filtering

Optional.

Filtered tables display only matching rows.

---

# Search Highlighting

Matched values may be highlighted.

Table layout must remain unchanged.

---

# Streaming Tables

Rows appear incrementally.

Example

```text
Loading...

Name      Status

frontend  Ready
```

Additional rows are appended.

---

# Incremental Rendering

Only newly added or modified rows are rendered.

Avoid full table redraws.

---

# Markdown Tables

Supported.

Rendered according to this specification.

---

# CSV Preview

CSV files render as structured tables.

Delimiter detection is automatic.

---

# JSON Conversion

Arrays of objects may render as tables.

Example

```json
[
  { "name": "App", "status": "Ready" }
]
```

↓

```text
Name      Status

App       Ready
```

---

# Code Integration

Tables inside Code Blocks are not interpreted.

Formatting remains unchanged.

---

# Selection

Supported.

Users may select

- Individual cells
- Rows
- Entire table

---

# Copy

Supported.

Copied text preserves table alignment where possible.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Headers remain semantically distinguishable.

---

# Safe Area

Tables remain inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Keyboard Behavior

Keyboard opening

Reduces available height.

Horizontal scrolling remains available.

---

# Performance

Large tables

Use virtual rendering.

Only visible rows should be drawn.

---

# Error Handling

Malformed table

Render best-effort layout.

Never crash the renderer.

---

# Security

Never execute table content.

Render all values as plain text unless explicitly formatted.

---

# Restrictions

Never

- Wrap cells automatically
- Break column alignment
- Hide headers
- Modify source data
- Rearrange rows unexpectedly
- Depend only on color for meaning

---

# Example Configuration Table

```text
Key              Value

Model            GPT-5

Workspace        mobile-cli

Theme            Dark

Streaming        Enabled
```

---

# Example File Table

```text
Name             Size

README.md        3 KB

package.json     1 KB

src              Folder
```

---

# Example Search Results

```text
File             Line

App.tsx          42

main.ts          18

README.md        91
```

---

# Table Renderer Checklist

Every table must

- Preserve alignment
- Support dynamic widths
- Support horizontal scrolling
- Support streaming
- Preserve headers
- Support selection
- Support copying
- Respect safe areas
- Render efficiently
- Remain terminal-native

---

# Core Rules

1. Tables preserve column alignment.
2. Cells never wrap automatically.
3. Horizontal scrolling is supported.
4. Column widths are calculated dynamically.
5. Streaming appends rows incrementally.
6. Headers remain visible and distinct.
7. Numeric values align to the right.
8. Rendering is optimized for large datasets.
9. Tables remain inside the Conversation Area.
10. Optimize for Android Termux.

---

# Summary

The Table Renderer provides a responsive, terminal-native system for displaying structured data within the Mobile AI CLI. By maintaining precise column alignment, supporting horizontal scrolling, enabling incremental streaming, and adapting dynamically to terminal size, it ensures that tables remain readable and efficient on Android Termux while preserving the integrity of the original data.