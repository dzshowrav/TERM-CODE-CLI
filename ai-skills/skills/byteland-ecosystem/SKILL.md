---
name: byteland-ecosystem
description: "Master ByteLand Technology's open-source ecosystem — Ink CLI components, Tinky terminal framework, Taffy WASM layout engine, headless-ghidra RE skill family, spec-forge workflow, byteland-design system, and Jules API. Trigger: When building Ink terminal UIs, creating CLI apps with React, needing scroll/input/list components for Ink, rendering Ink in the browser, using Taffy layout engine in JS/TS, reverse-engineering with Ghidra via agents, creating spec-driven development workflows, or integrating with the Jules API."
license: "MIT"
metadata:
  version: "2.0"
  type: domain
  skills:
    dependencies: []
---

# ByteLand Ecosystem — Complete Reference

> **Organization**: [ByteLand Technology Limited](https://github.com/ByteLandTechnology) (Hong Kong, est. Aug 2024)
> **Open-Source Brand**: OpenLand
> **License**: All projects MIT
> **Stack**: TypeScript, Rust (WASM), React 19, Ink 6, Taffy layout engine
> **i18n**: English, 简体中文, 日本語 (all repos have multi-language READMEs)
> **Repo Count**: 51 public repositories

A comprehensive, family-level skill for every open-source project from ByteLand Technology. Covers the full Ink CLI component ecosystem, the Tinky framework, Taffy WASM layout bindings, agent skill families, design system, and Jules API client.

---

## When to Use

Trigger this skill when ANY of these conditions are met:

- **Ink CLI Components**: Building terminal UIs with Ink needing scroll views, scrollable lists, scroll bars, or multi-line text input
- **Ink → Browser**: Rendering Ink applications in the browser via Xterm.js
- **Tinky Framework**: Building CLIs with React + CSS Flexbox/Grid layout (Taffy-powered)
- **Terminal Images**: Rendering images in terminal emulators (Kitty, iTerm2, Sixel, etc.)
- **Taffy WASM Layout**: Using CSS Flexbox/Grid layout algorithms in JavaScript via WebAssembly
- **Agent RE Skills**: Automating Ghidra reverse-engineering pipelines with AI agents
- **Spec Workflows**: Creating YAML-first spec-driven development with agent-guided stage progression
- **Jules API**: Integrating with the Jules API from TypeScript/Node.js
- **ByteLand Design**: Applying ByteLand's design language (visual, web, icon, motion, video styles)

Don't use for:

- General React web development (use react-expert or frontend-ui-engineering)
- General Node.js backend development (use nodejs or backend-dev)
- General Rust/wasm development without Taffy integration
- Tasks unrelated to terminal UI, CLI frameworks, or ByteLand's specific projects

---

## Ecosystem Map

```
ByteLand Ecosystem Overview
├── INK ECOSYSTEM (React for CLIs — traditional Ink)
│   ├── ink-scroll-view (⭐10)  — Scroll container primitive
│   ├── ink-scroll-list (⭐8)   — Selectable list on scroll-view
│   ├── ink-scroll-bar          — Vertical scroll bar component
│   ├── ink-multiline-input (⭐7) — Multi-line text input
│   └── ink-canvas (⭐2)        — Ink apps in browser via Xterm.js
│
├── TINKY FRAMEWORK (React for CLIs — Taffy-powered)
│   ├── tinky (⭐1 / v1.9.0)    — Core framework (Flexbox + Grid)
│   └── tinky-image (⭐1)       — Terminal image rendering (6 backends)
│
├── LAYOUT INFRASTRUCTURE
│   └── taffy-layout (⭐12)     — WASM bindings for Taffy Rust engine
│
├── AGENT SKILL FAMILIES
│   ├── headless-ghidra (⭐6)   — Ghidra RE automation (5 phases)
│   ├── spec-forge              — YAML spec workflow (6 stages, Rust CLI)
│   └── byteland-design         — 5-skill brand design language suite
│
├── API CLIENTS
│   └── jules-api-node          — Unofficial Jules API TypeScript client
│
└── OTHER
    ├── ink-taffy               — Ink fork with Taffy layout
    ├── jellyfin-agent-cli      — Jellyfin media server CLI (Rust)
    ├── taffy-layout-docs       — Docs for taffy-layout
    └── conventional-commit-skill — Agent skill for conventional commits
```

---

## [CRITICAL] Core Principles Across All Projects

### 1. Controlled Component Pattern
ALL interactive ByteLand Ink/Tinky components follow the **controlled component pattern**. The parent owns state; the component renders based on props. **Never let the component manage state internally.**

```tsx
// ✅ CORRECT (ByteLand pattern)
const [selectedIndex, setSelectedIndex] = useState(0);
<ScrollList selectedIndex={selectedIndex}>
  {items.map((item, i) => <Text key={i}>{item}</Text>)}
</ScrollList>

// ❌ WRONG — components do NOT manage selection internally
<ScrollList>  {/* no selectedIndex — broken */}
```

### 2. Input Handling Is Parent's Responsibility
ByteLand scroll/list/input components **do not capture keyboard input**. The parent must use `useInput` (Ink) or the framework's input hooks.

```tsx
// ✅ CORRECT: Parent handles input, component renders
useInput((input, key) => {
  if (key.upArrow) scrollRef.current?.scrollBy(-1);
  if (key.downArrow) scrollRef.current?.scrollBy(1);
});
<ScrollView ref={scrollRef}>...</ScrollView>

// ❌ WRONG: Components do NOT handle input internally
```

### 3. Terminal Resize → Must Remeasure
Ink/Tinky components don't auto-detect terminal resizes. Always listen to `process.stdout.on("resize")` and call `remeasure()`.

```tsx
useEffect(() => {
  const handleResize = () => ref.current?.remeasure();
  process.stdout.on("resize", handleResize);
  return () => process.stdout.off("resize", handleResize);
}, []);
```

### 4. Unique Keys on Children
All ByteLand scroll components require stable, unique `key` props on children for accurate height tracking.

```tsx
// ✅ CORRECT
<ScrollView>
  {items.map((item, i) => <Text key={item.id}>{item.text}</Text>)}
</ScrollView>

// ❌ WRONG — no key or index-only key for dynamic lists
```

### 5. Overflow:Hidden + MarginTop Scrolling Pattern
The core scrolling mechanism across ALL ByteLand scroll components uses a negative margin approach:

```
┌─────────────────────────┐
│ (hidden content)        │ ← Content above viewport
├─────────────────────────┤ ← scrollOffset (marginTop: -N)
│ ┌───────────────────┐   │
│ │ Visible Viewport  │   │ ← overflow: hidden
│ │                   │   │
│ └───────────────────┘   │
├─────────────────────────┤
│ (hidden content)        │ ← Content below viewport
└─────────────────────────┘
```

---

## Decision Tree

```
User needs terminal UI component?
  → Using Ink (traditional)?
    → Need scrolling?       → See [ink-scroll-view](references/ink-scroll-view.md)
    → Need selectable list? → See [ink-scroll-list](references/ink-scroll-list.md)
    → Need scroll bar?      → See [ink-scroll-bar](references/ink-scroll-bar.md)
    → Need text input?      → See [ink-multiline-input](references/ink-multiline-input.md)
    → Need browser render?  → See [ink-canvas](references/ink-canvas.md)
  → Using Tinky (Taffy-powered)?
    → Need core framework?  → See [tinky](references/tinky.md)
    → Need image render?    → See [tinky-image](references/tinky-image.md)

User needs CSS layout in JS?
  → WebAssembly OK?         → See [taffy-layout](references/taffy-layout.md)

User needs agent skills?
  → Ghidra RE automation?   → See [headless-ghidra](references/headless-ghidra.md)
  → Spec-driven workflow?   → See [spec-forge](references/spec-forge.md)
  → Brand design language?  → See [byteland-design](references/byteland-design.md)

User needs Jules API client? → See [jules-api-node](references/jules-api-node.md)
```

---

## Workflow: Building a CLI App with ByteLand Components

### Phase 1: Framework Choice

1. **Choose framework**:
   - **Ink** (traditional) — use `ink`, `ink-scroll-view`, `ink-scroll-list`, `ink-scroll-bar`, `ink-multiline-input`
   - **Tinky** (Taffy-powered Flexbox/Grid) — use `tinky`, `tinky-image`
2. **Install**: `npm install <packages>` (peer deps: `react`, `ink` or `tinky`)
3. **Set up TypeScript**: target ES2022+, jsx react-jsx, include ES2020/DOM/DOM.Iterable lib

### Phase 2: Layout Structure

1. Create `Box` containers with explicit dimensions (`height` is critical for scrolling)
2. Add `ScrollView` with ref for scrollable content
3. Add `ScrollBar` (border or inset mode) for visual scroll position
4. Wire up keyboard input with `useInput`
5. Handle terminal resize with `process.stdout.on("resize") → remeasure()`

### Phase 3: Interactivity

1. **ScrollList**: Parent manages `selectedIndex`, updates on arrow keys
2. **MultilineInput**: Parent manages `value`, `onChange` for each edit
3. **Tinky useFocus**: For tabbable interactive elements

### Phase 4: Testing

1. Mock `MeasureBox` for scroll-related components
2. Use `useCustomInput` injection pattern for keyboard logic tests
3. Use controlled wrapper pattern for interactive component tests

### Phase 5: Browser Deployment (Optional)

1. Add `ink-canvas` with `inkCanvasPolyfills()` Vite plugin
2. Or use `tinky-web` for Tinky apps in browser
3. Container MUST have explicit width/height dimensions

---

## Conventions

- **All components are TypeScript-first** with full type definitions
- **All components are controlled** — parent owns state
- **No components capture input** — parent routes keyboard events
- **All scroll components use overflow:hidden + marginTop** scrolling
- **Children must have stable unique keys** for height tracking
- **Terminal resize must trigger remeasure** — always wire it up
- **All projects use MIT license**
- **Commit style**: conventional commits (enforced via commitlint + husky)
- **Build tooling**: tsup (ESM), Vitest, eslint flat config, TSUP/TypeScript
- **Docs**: Typedoc + typedoc-plugin-markdown; dedicated documentation websites on subdomains

---

## Critical Anti-Patterns

### ❌ NEVER: Skip key prop on scroll children
Without keys, item height tracking breaks on reorder/removal. Always provide stable keys.

### ❌ NEVER: Forget terminal resize handling
Ink/Tinky components do NOT auto-detect terminal resize. Without `remeasure()`, the viewport will be wrong after resize.

### ❌ NEVER: Let scroll-list manage selection
ScrollList is fully controlled. Parent must provide `selectedIndex` and update it. The component will NOT call any `onSelectionChange`.

### ❌ NEVER: Use scroll-view without parent input handling
ScrollView does NOT capture arrow keys, page up/down, or any input. The parent must call `scrollBy()` etc. via ref.

### ❌ NEVER: Try to use ink-canvas without explicit container dimensions
The Xterm.js terminal needs explicit container size (`width`/`height` or CSS class). Without it, the terminal renders empty/zero-size.

---

## Edge Cases

**Content smaller than viewport**: Scroll offset is 0, scroll bar hides (with `autoHide` in inset mode). Scrollview handles natively.

**Items larger than viewport**: Scrolling is constrained to keep at least part of the selected item visible. Users can scroll within large items.

**Dynamic item heights**: Use `remeasureItem(index)` for individual items (expand/collapse). Avoid full `remeasure()` which is more expensive.

**Terminal resize race condition**: Call `remeasure()` in a `useEffect` that depends on terminal dimensions. Use `setTimeout(0)` if needed.

**Paste detection**: `input.length > 1` = paste in `ink-multiline-input`. Highlights pasted text if `highlightPastedText` is set.

**Browser polyfills**: Ink uses Node.js APIs (`process`, `Buffer`). `ink-canvas` provides process shim + Buffer polyfill. Always use `inkCanvasPolyfills()` plugin.

**Grid + Flexbox in terminal**: Tinky supports both via Taffy. Use `display="grid"` or `flexDirection="row"` on Box. Grid supports template areas, named lines, auto-flow.

**WASM initialization**: `taffy-layout` requires `await loadTaffy()` before creating `TaffyTree`. Must run once at app startup.

---

## Checklist

### Before Using ByteLand Components
- [ ] Framework chosen (Ink or Tinky)
- [ ] Peer dependencies installed (`react`, `ink >=6` or `tinky`)
- [ ] TypeScript configured (ES2022+, react-jsx)
- [ ] Terminal resize handler planned
- [ ] Input handling architecture designed (parent routes keyboard)

### During Implementation
- [ ] All children have unique stable keys
- [ ] Scroll ref wired to keyboard input
- [ ] `remeasure()` called on terminal resize
- [ ] Controlled component pattern followed (state in parent)
- [ ] Container has explicit height for scroll to work

### For Browser Deployment
- [ ] `ink-canvas` / `tinky-web` added
- [ ] Polyfills configured (Vite plugin or Webpack plugin)
- [ ] Container has explicit dimensions
- [ ] `focused` prop set to `true`

---

## Resources

### Reference Files (per project deep-dive)
- [ink-scroll-view](references/ink-scroll-view.md) — Scroll container component
- [ink-scroll-list](references/ink-scroll-list.md) — Selectable list component
- [ink-scroll-bar](references/ink-scroll-bar.md) — Scroll bar component
- [ink-multiline-input](references/ink-multiline-input.md) — Multi-line text input
- [ink-canvas](references/ink-canvas.md) — Ink in browser via Xterm.js
- [tinky](references/tinky.md) — Tinky terminal framework
- [tinky-image](references/tinky-image.md) — Terminal image rendering
- [taffy-layout](references/taffy-layout.md) — Taffy WASM layout engine bindings
- [headless-ghidra](references/headless-ghidra.md) — Ghidra RE agent skill family
- [spec-forge](references/spec-forge.md) — YAML spec workflow
- [byteland-design](references/byteland-design.md) — Design language skill suite
- [jules-api-node](references/jules-api-node.md) — Jules API TypeScript client

### Examples
- [ink-scroll-view-app](examples/ink-scroll-view-app.md) — Full scrollview app example
- [tinky-counter](examples/tinky-counter.md) — Tinky counter app example
- [taffy-layout](examples/taffy-layout.md) — Taffy layout example

### Assets
- [ByteLand Quick Reference](assets/quick-reference.md) — Cheatsheet of all APIs
