# 19-tree-renderer.md

# Tree Renderer
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Tree Rendering System used throughout the Mobile AI CLI.

The Tree Renderer displays hierarchical data such as project folders, file systems, dependency graphs, workspace indexes, Git trees, symbol outlines, and AI-generated structures.

The renderer must preserve hierarchy, indentation, and readability while remaining optimized for Android Termux.

---

# Design Goals

The Tree Renderer must be

- Mobile First
- Terminal Native
- Readable
- Hierarchy Focused
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

Hierarchy must always be immediately understandable.

Indentation is the primary visual cue.

The renderer must preserve structure rather than maximize information density.

---

# Rendering Pipeline

```
Tree Data

↓

Hierarchy Analysis

↓

Indentation Calculation

↓

Node Layout

↓

Terminal Renderer

↓

Display
```

---

# Supported Sources

Tree structures may originate from

- Workspace Files
- Project Explorer
- Git Trees
- Symbol Outline
- AST
- Dependency Graph
- Package Structure
- Search Results
- AI Responses

---

# Tree Structure

Each tree contains

- Root Node
- Parent Nodes
- Child Nodes
- Leaf Nodes

---

# Basic Example

```text
workspace
├── src
│   ├── components
│   ├── pages
│   └── utils
├── public
├── package.json
└── README.md
```

---

# Root Node

Always displayed first.

Only one root node exists unless multiple trees are intentionally rendered.

---

# Parent Nodes

May contain one or more child nodes.

---

# Leaf Nodes

Nodes without children.

---

# Indentation

Default indentation

```
4 Spaces
```

Indentation must remain consistent throughout the tree.

---

# Branch Characters

Preferred

```text
├──

└──

│
```

Fallback

```text
+

-

|
```

when Unicode is unavailable.

---

# Node Alignment

Every sibling node aligns vertically.

Tree connectors remain continuous.

---

# Node Types

Supported

- Folder
- File
- Symbol
- Package
- Module
- Function
- Class
- Namespace
- Custom Node

---

# Expandable Nodes

Supported.

Collapsed

```text
src
```

Expanded

```text
src
├── components
├── pages
└── utils
```

---

# Expansion State

Supported

```
Expanded

Collapsed
```

The renderer preserves expansion state during updates.

---

# Lazy Loading

Large trees may load child nodes on demand.

---

# Streaming Trees

Nodes appear incrementally.

Example

```text
workspace
├── src
```

↓

```text
workspace
├── src
│   ├── components
│   └── pages
```

---

# Incremental Rendering

Only changed nodes are rendered.

Avoid full tree redraws.

---

# Long Names

Long node names may be truncated.

Example

```text
very-long-component...
```

Full names remain accessible through selection.

---

# Empty Folder

Display

```text
empty-folder
└── (empty)
```

---

# Hidden Files

Optional.

Example

```text
.gitignore
```

Visibility depends on user settings.

---

# Symbol Tree

Example

```text
App.tsx
├── App
├── Header
├── Footer
└── Button
```

---

# Dependency Tree

Example

```text
project
├── React
├── TypeScript
└── Vite
```

---

# Search Tree

Grouped search results

```text
src
├── App.tsx
└── main.ts

docs
└── README.md
```

---

# Selection

Supported.

Users may select

- Single Node
- Multiple Nodes
- Entire Branch

---

# Copy

Supported.

Copied content preserves indentation.

---

# Search Highlighting

Matched nodes may be highlighted.

Tree structure remains unchanged.

---

# Horizontal Scrolling

Supported.

Long trees never wrap automatically.

---

# Vertical Scrolling

Supported.

Large trees scroll efficiently.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Hierarchy must remain understandable without color alone.

---

# Safe Area

Trees remain inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Keyboard Behavior

Keyboard opening

Reduces visible height.

Tree layout remains unchanged.

---

# Performance

Large trees

Use virtual rendering.

Render only visible nodes.

---

# Error Handling

Invalid hierarchy

Render best-effort structure.

Never crash the renderer.

---

# Security

Tree nodes are rendered as text only.

Never execute node content.

---

# Restrictions

Never

- Break indentation
- Reorder nodes automatically
- Wrap long node names
- Remove branch connectors
- Collapse nodes unexpectedly
- Hide hierarchy information

---

# Example Workspace Tree

```text
workspace
├── src
│   ├── components
│   ├── hooks
│   ├── pages
│   └── utils
├── public
├── package.json
├── tsconfig.json
└── README.md
```

---

# Example Symbol Tree

```text
main.ts
├── initialize
├── createApp
├── registerCommands
└── start
```

---

# Example Dependency Tree

```text
dependencies
├── React
├── Ink
├── Effect
├── Zod
└── TypeScript
```

---

# Tree Renderer Checklist

Every tree must

- Preserve hierarchy
- Preserve indentation
- Support expandable nodes
- Support incremental rendering
- Support horizontal scrolling
- Support vertical scrolling
- Support selection
- Support copying
- Respect safe areas
- Remain terminal-native

---

# Core Rules

1. Hierarchy is the primary visual structure.
2. Indentation remains consistent.
3. Tree connectors remain aligned.
4. Long names never wrap automatically.
5. Support expandable nodes.
6. Streaming updates incrementally.
7. Preserve expansion state.
8. Render efficiently for large trees.
9. Trees remain inside the Conversation Area.
10. Optimize for Android Termux.

---

# Summary

The Tree Renderer provides a structured, terminal-native visualization of hierarchical information within the Mobile AI CLI. By preserving indentation, branch alignment, expandable nodes, and incremental rendering, it enables users to explore workspaces, project structures, dependency graphs, and symbol outlines efficiently while maintaining a consistent mobile-first experience on Android Termux.