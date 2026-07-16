# ByteLand Technology Limited — Open Source Knowledge

> Researched: 2026-07-13 from https://byteland.app/opensource and GitHub
> Category: project-config / third-party-knowledge

---

## Company Overview

| Attribute | Value |
|-----------|-------|
| **Name** | ByteLand Technology Limited |
| **Founded** | August 2024 |
| **Location** | Hong Kong — ROOM 25, 15/F, BLOCK E, LUK HOP INDUSTRIAL BUILDING, 8 LUK HOP STREET, SAN PO KONG |
| **Website** | https://byteland.app |
| **GitHub** | https://github.com/ByteLandTechnology |
| **Email** | github@byteland.app |
| **Products** | Cross-platform Mobile apps, Progressive Web Apps (PWAs) |
| **Focus** | CLI tools, Ink/React terminal UI, open-source infrastructure |
| **Verification** | Verified GitHub organization controlling byteland.app domain |

ByteLand builds practical, high-quality apps for daily life (weather, transit, coding-on-the-go). Their open-source initiative is called **OpenLand**.

---

## Ink Ecosystem Projects (CLI Terminal UI)

ByteLand has built a comprehensive suite of components for [Ink](https://github.com/vadimdemedes/ink) — React for CLI applications. These projects form a hierarchical ecosystem:

### 1. ink-scroll-view
- **NPM**: `ink-scroll-view`
- **GitHub**: https://github.com/ByteLandTechnology/ink-scroll-view
- **Stars**: 10 | **Forks**: 2
- **License**: MIT
- **Docs**: https://ink-scroll-view.byteland.app
- **Status**: Active, stable

**What it is**: A robust, performance-optimized ScrollView component for Ink CLI applications. It's a layout primitive that handles content larger than the visible terminal viewport.

**Key Features**:
- **Optimistic Updates** — immediate state updates for smoother interaction
- **Efficient Re-rendering** — renders all children but manages visibility via `overflow` and offsets
- **Auto-Measurement** — automatically measures child heights using virtually rendered DOM
- **Dynamic Content** — supports adding, removing, expanding/collapsing items
- **Layout Stability** — maintains scroll position when content changes

**Architecture**: Renders all children in a container, shifts content using `marginTop`, parent box with `overflow="hidden"` = viewport.

**API Highlights**:
- `<ScrollView>` — core component, extends Ink BoxProps
- `<ControlledScrollView>` — for advanced cases (sync multiple views, animate transitions)
- Ref methods: `scrollTo()`, `scrollBy()`, `scrollToTop()`, `scrollToBottom()`, `getScrollOffset()`, `getContentHeight()`, `getViewportHeight()`, `getBottomOffset()`, `getItemHeight(index)`, `getItemPosition(index)`, `remeasure()`, `remeasureItem(index)`
- Callbacks: `onScroll`, `onViewportSizeChange`, `onContentHeightChange`, `onItemHeightChange`

**Critical Patterns**:
1. DOES NOT capture user input automatically — control programmatically via refs + Ink's `useInput`
2. MUST call `remeasure()` on terminal `resize` event
3. Children MUST have unique `key` props
4. Use `remeasureItem(index)` for expand/collapse instead of full update

---

### 2. ink-scroll-list
- **NPM**: `ink-scroll-list`
- **GitHub**: https://github.com/ByteLandTechnology/ink-scroll-list
- **Stars**: 8 | **Forks**: 1
- **License**: MIT
- **Docs**: https://ink-scroll-list.byteland.app
- **Status**: Active, v0.4.0

**What it is**: High-level ScrollList component built on top of ink-scroll-view with focus management and item selection.

**Key Features**:
- Controlled selection via `selectedIndex` prop
- Auto-scrolling to ensure selected item is visible
- Flexible alignment modes for how selected item aligns in viewport
- Optimized for selection position tracking

**API Highlights**:
- Props: `selectedIndex` (controlled), `scrollAlignment` (`'auto' | 'top' | 'bottom' | 'center'`)
- Extends all `ScrollViewProps` from ink-scroll-view
- Ref methods extend `ScrollViewRef` (scrolling methods are constrained to keep selection visible)

**Scroll Alignment Modes**:
| Mode | Behavior | Best For |
|------|----------|----------|
| `'auto'` | Minimal scrolling to bring item into view | Keyboard navigation |
| `'top'` | Always aligns selected item to viewport top | — |
| `'bottom'` | Always aligns selected item to viewport bottom | — |
| `'center'` | Always centers selected item in viewport | Search/spotlight UX |

**v0.4.0 Breaking Change**: Became a **fully controlled component**. Removed `onSelectionChange`, `selectNext/Previous`, `scrollToItem` etc. Parent now manages `selectedIndex` directly via state.

**Important**: Does NOT handle input internally — parent must use `useInput` and update `selectedIndex`.

---

### 3. ink-scroll-bar
- **NPM**: `ink-scroll-bar`
- **GitHub**: https://github.com/ByteLandTechnology/ink-scroll-bar
- **License**: MIT
- **Status**: Active

**What it is**: Customizable, high-precision vertical scroll bar component for Ink CLI applications.

**Two Rendering Modes**:
| Mode | Description |
|------|-------------|
| **Border Mode** | Seamlessly integrates with container borders (replaces one side) |
| **Inset Mode** | Renders inside content area, supports `autoHide` when content fits viewport |

**Border Styles** (matching Ink): `single`, `double`, `round`, `bold`, `singleDouble`, `doubleSingle`, `classic`, `arrow`
**Inset Styles**: `block` (█/░), `line` (│), `thick` (┃/╏), `dots` (●/·)

**Components**:
- `<ScrollBar>` — standalone, flexible positioning
- `<ScrollBarBox>` — wrapper for instant bordered containers with scroll bars

---

### 4. ink-multiline-input
- **NPM**: `ink-multiline-input`
- **GitHub**: https://github.com/ByteLandTechnology/ink-multiline-input
- **Stars**: 7 | **Forks**: 1
- **License**: MIT
- **Status**: Active

**What it is**: A robust multi-line text input component for Ink applications.

**Features**: Vertical scrolling, cursor navigation (arrow keys), editing (Backspace/Delete), customizable active line styling (`highlightStyle`), mask support for passwords, flexible rows (`rows`/`maxRows`).

**Key Props**:
- `value`, `onChange`, `onSubmit`, `rows`, `maxRows`, `highlightStyle`, `textStyle`, `placeholder`, `mask`, `showCursor`, `focus`, `tabSize`, `keyBindings`, `highlightPastedText`

**Custom Key Bindings**:
```js
keyBindings={{
  submit: (key) => key.ctrl && key.return,  // Ctrl+Enter to submit
  newline: (key) => key.return,              // Enter for newline
}}
```

**Advanced**: `ControlledMultilineInput` for controlling cursor position externally.

---

### 5. ink-canvas
- **NPM**: `ink-canvas`
- **GitHub**: https://github.com/ByteLandTechnology/ink-canvas
- **Stars**: 2 | **Forks**: 1
- **License**: MIT
- **Docs**: https://ink-canvas.byteland.app
- **Status**: Active

**What it is**: A library for rendering Ink applications in the browser using Xterm.js. Bridges Node.js-based CLI UIs with web-based terminal emulators.

**Architecture**:
1. **Process Shim** (`shims/process.ts`) — browser-compatible mock of Node.js `process` (env, stdout/stdin/stderr, nextTick)
2. **Custom Streams** (`utils/streams.ts`) — `TerminalWritableStream` (stdout), `TerminalReadableStream` (stdin) optimized for browser
3. **Canvas Component** — auto-sizing wrapper that listens for resize events
4. **React Lifecycle Management** — Xterm.js init/cleanup, stream connection, Ink lifecycle

**Props**:
- `children`, `focused`, `cols`, `rows`, `terminalOptions` (Xterm.js config), `onResize`

**Handle Properties**: `terminal` (Xterm.js Terminal), `dimensions` (cols/rows), `instance` (Ink render instance)

**Build Integration**:
- **Vite**: Use `inkCanvasPolyfills()` plugin
- **Webpack/Next.js**: Use `InkCanvasWebpackPlugin` (`new InkCanvasWebpackPlugin()` in webpack config)

**Troubleshooting**: Container MUST have explicit dimensions. `focused` must be `true` for keyboard input. Ensure `tsconfig.json` includes `"lib": ["ES2020", "DOM", "DOM.Iterable"]`.

---

## Tinky Framework

### tinky
- **NPM**: `tinky`
- **GitHub**: https://github.com/ByteLandTechnology/tinky
- **Stars**: 1 | **Forks**: 1
- **License**: MIT
- **Status**: Active, v1.9.0 (17 releases, 61 commits)

**What it is**: A modern React-based framework for building beautiful and interactive CLIs. Uses the Taffy layout engine for CSS Flexbox/Grid layout support in the terminal. Think of it as "Ink re-imagined with Taffy."

**Key Features**:
- React Components (Box, Text, Static, Transform, Newline, Spacer)
- **CSS Flexbox & Grid Layout** powered by Taffy (not Yoga/Ink's yoga-layout)
- Keyboard input hooks (`useInput`)
- Focus management (`useFocus`, `useFocusManager`) with Tab/Shift+Tab navigation
- Borders & backgrounds, accessibility (ARIA attributes)
- Hot reloading with React DevTools support
- TypeScript-first with comprehensive type definitions

**Components**:
- `<Box>` — fundamental building block like `<div>` in browser; supports Flexbox AND Grid layout
- `<Text>` — styled text with colors, bold, italic, hex/ANSI colors
- `<Static>` — renders static content that won't be updated (perfect for logs)
- `<Transform>` — transforms output of children (e.g., `toUpperCase()`)
- `<Newline>` / `<Spacer>` — spacing helpers

**Hooks**: `useInput`, `useApp` (with `exit()`), `useFocus`, `useFocusManager`, `useStdin`, `useStdout`, `useStderr`

**Incremental Rendering**: Two strategies — `"run"` (diffs terminal cells, writes minimal runs) and `"line"` (diffs lines, rewrites changed lines). Auto-falls back in debug/screen-reader/CI envs.

**Testing**: Uses Bun for tests.

---

### tinky-image
- **NPM**: `tinky-image`
- **GitHub**: https://github.com/ByteLandTechnology/tinky-image
- **Stars**: 1
- **License**: MIT
- **Status**: Active, v1.0.0

**What it is**: Terminal image rendering for the Tinky CLI framework.

**Image Backends** (auto-detected, overridable):
1. `kitty` — terminal graphics protocol (Kitty-compatible)
2. `iterm` — inline image protocol (iTerm2-compatible)
3. `sixel` — bitmap graphics for Sixel-supporting terminals
4. `halfblock` — full-color using Unicode half-block `▄` + ANSI colors
5. `braille` — monochrome high-density using Braille patterns
6. `ascii` — plain character-art fallback

**Props**: `src` (file path, URL, Uint8Array, Blob), `width`/`height` (in terminal cells or %), `renderer` (backend override), `resizeMode` (`contain`/`cover`/`fill`/`none`)

**Dual Entry**: `tinky-image` for Node runtime, `tinky-image/browser` for bundled apps.

---

## Layout Infrastructure

### taffy-layout
- **NPM**: `taffy-layout`
- **GitHub**: https://github.com/ByteLandTechnology/taffy-layout
- **Stars**: 12 (most starred ByteLand project) | **Forks**: 0
- **License**: MIT
- **Website**: https://taffylayout.com
- **Status**: Active, v2.0.3 (12 releases, 67 commits)
- **Languages**: TypeScript 61.4%, Rust 38.6%

**What it is**: High-performance WebAssembly bindings for the [Taffy](https://github.com/DioxusLabs/taffy) layout engine (Rust). Brings CSS Flexbox and Grid layout algorithms to JavaScript with near-native performance.

**Core API**:
- `loadTaffy()` — initialize WebAssembly module (must await)
- `TaffyTree` — main class for managing layout trees
- `Style` — configuration object for layout properties (display, flexDirection, alignItems, gap, padding, etc.)
- `Layout` — read-only computed layout result (width, height, x, y)

**Key Features**: Full Flexbox + CSS Grid, custom text measurement callbacks, TypeScript ready, tree-based API, CSS-like property names

**Advanced**: Custom text measurement with `newLeafWithContext()` and `computeLayoutWithMeasure()`; Grid template areas, named grid lines, absolute positioning, percentage sizing, block layout with replaced elements (`itemIsReplaced`, `aspectRatio`)

**Browser Support**: Chrome 57+, Firefox 52+, Safari 11+, Edge 16+

---

## Agent Skill Families

### headless-ghidra
- **GitHub**: https://github.com/ByteLandTechnology/headless-ghidra
- **Stars**: 6
- **License**: MIT
- **Status**: Active, v1.8.0 (18 releases, 59 commits)
- **Languages**: Rust 55%, Java 28.8%, JS 8.8%, Python 5.7%, Shell 1.7%

**What it is**: A Ghidra reverse-engineering skill family for AI agents. Provides reproducible, evidence-backed workflows with audit-ready Markdown outputs.

**Architecture**: 5-phase pipeline (P0-P4) plus single-function analysis:
| Phase | Purpose |
|-------|---------|
| P0 — Intake | Confirm target, init workspace, set scope |
| P1 — Baseline | Import into Ghidra, export baseline artifacts, record runtime |
| P2 — Evidence | Identify third-party code and evidence sources |
| P3 — Discovery | Enrich names, signatures, types, constants, strings |
| P4 — Batch Decompile | Apply metadata and decompile selected functions |

Also includes `ghidra-agent-cli` (Rust CLI), the orchestration layer, and a single-function deep analysis skill.

**Install**: Via Codex `$skill-installer` or `npx skills add` command.

---

### spec-forge
- **GitHub**: https://github.com/ByteLandTechnology/spec-forge
- **Stars**: 0 (new)
- **License**: MIT
- **Status**: v1.0.0 (3 commits)
- **Languages**: Rust 83.3%, JavaScript 15.6%, Shell 1.1%

**What it is**: A YAML-first workflow for turning rough requests into approved implementation specs and auditable implementation reports. Combines agent-guided stage progression with a Rust CLI.

**Workflow Stages** (P0-P5):
| Stage | Purpose |
|-------|---------|
| P0 — Intake | Frame request, stakeholders, scope, constraints |
| P1 — Architecture | Lock solution outline, journey/component indexes |
| P2 — Journeys | Refine in-scope journeys in reviewable batches |
| P3 — Components | Refine components into implementation contracts |
| P4 — Readiness | Consolidate into final implementation spec |
| P5 — Implement | Record delivery status, validations, blockers |

**Workspace**: Lives under `.spec-forge/` with `registry.yaml`, per-spec subdirs for pipeline-state, framing, architecture, journeys, components, synthesis, gates.

---

## Design System

### byteland-design
- **GitHub**: https://github.com/ByteLandTechnology/byteland-design
- **Stars**: 0 (new)
- **License**: MIT
- **Status**: New (1 commit)

**What it is**: ByteLand's design-language Skill suite — 5 Skills defining brand visual, web, icon, motion, and video direction.

**Skill Hierarchy**:
```
byteland-visual-style (foundation: tokens, color, surface, anti-patterns)
├── byteland-web-style
├── byteland-icon-style
├── byteland-motion-style
└── byteland-video-style
```

**Testing**: Node built-in test runner (`node --test`), 3 suites: style rules, WCAG AA contrast checks, eval suite alignment.

---

## Jules API

### jules-api-node
- **NPM**: `jules-api-node`
- **GitHub**: https://github.com/ByteLandTechnology/jules-api-node
- **License**: MIT
- **Status**: Active

**What it is**: The unofficial TypeScript client for the Jules API. Auto-generated from the Jules OpenAPI specification using `@hey-api/openapi-ts`.

**Methods**: `approvePlan`, `listSessions`, `createSession`, `getSession`, `sendMessage`, `getActivity`, `listActivities`, `getSource`, `listSources`

**Auth**: Pass `X-Goog-Api-Key` header with API key.

---

## Other Repositories

- **jellyfin-agent-cli** (Rust, 1 star) — CLI for Jellyfin media server agent interactions
- **ink-taffy** (TypeScript, 1 star) — Fork of vadimdemedes/ink with Taffy layout integration
- **conventional-commit-skill** (Python) — Agent skill for conventional commits
- **taffy-layout-docs** (TypeScript, 1 star) — Documentation for taffy-layout

---

## Blog Articles (ByteLog)

| Date | Title | Topics |
|------|-------|--------|
| Dec 18, 2025 | React Puck: The Visual Editor for Developers | React, Visual Editor, Next.js, Headless CMS, Open Source |
| Dec 14, 2025 | AI Agent Development with Mastra | AI, TypeScript |
| Dec 6, 2025 | Introducing ink-scroll-view | Ink, CLI, React, TypeScript |
| Dec 5, 2025 | Hello, ByteLand! | Announcement |

---

## Key Technical Patterns & Architecture

1. **Terminal UI Ecosystem**: ByteLand has built the most comprehensive Ink component ecosystem including scrolling, lists, scroll bars, multi-line input, and browser rendering — all MIT-licensed.

2. **Dual Layout Engine Strategy**:
   - **Ink ecosystem** uses Yoga-like layout (traditional Ink)
   - **Tinky** uses Taffy layout engine (WebAssembly-powered CSS Flexbox + Grid)

3. **Agent Skills Architecture**: ByteLand heavily develops "skill families" — hierarchical collections of agent instructions for AI coding assistants (Codex, Claude Code). Skills follow P0-P4/P5 phased pipeline patterns with workspace persistence.

4. **WASM Bridge Pattern**: `taffy-layout` demonstrates ByteLand's pattern of wrapping Rust WASM modules in TypeScript for high-performance JS packages.

5. **i18n Support**: Many repos have READMEs in English, Simplified Chinese, and Japanese.

6. **Monorepo Style**: Skill families use monorepo structures with independent sub-packages (e.g., headless-ghidra has 8 sibling skills, spec-forge has 7 stage skills).

---

## Key Facts for Recall

- ByteLand = Hong Kong company, founded Aug 2024
- OpenLand = their open-source initiative
- Primary expertise: Ink CLI components, Tinky framework, Taffy layout WASM bindings, AI agent skills, Ghidra RE automation
- All projects MIT licensed
- Strong TypeScript focus, some Rust for WASM and CLI tooling
- taffy-layout is their most starred project (12 stars)
- Notable: ink-scroll-view (10 stars), ink-scroll-list (8 stars), ink-multiline-input (7 stars), headless-ghidra (6 stars)
