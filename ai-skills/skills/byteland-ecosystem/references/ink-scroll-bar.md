# ink-scroll-bar — Reference

> **NPM**: `ink-scroll-bar`
> **GitHub**: https://github.com/ByteLandTechnology/ink-scroll-bar
> **License**: MIT

A customizable, high-precision vertical scroll bar component for Ink CLI applications with two distinct rendering modes.

---

## Two Rendering Modes

### Border Mode
Seamlessly integrates with container borders — replaces one side of the border with the scroll bar. Supports all Ink border styles.

```
┌─────────────────────┬─┐
│ Content             │█│  ← ScrollBar replaces
│ Content             │ │     right border
│ Content             │█│
└─────────────────────┴─┘
```

### Inset Mode
Renders inside the content area, like a floating scroll bar. Supports `autoHide`.

```
┌─────────────────────┐
│ Content         ░█░ │  ← Inset scrollbar
│ Content         ░░  │     inside container
│ Content         ░█░ │
└─────────────────────┘
```

---

## Components

### `<ScrollBar>` — Standalone

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `contentHeight` | `number` | **required** | Total height of scrollable content |
| `viewportHeight` | `number` | **required** | Height of visible area |
| `scrollOffset` | `number` | **required** | Current scroll position (0 to max) |
| `placement` | `'left' \| 'right' \| 'inset'` | `'right'` | Rendering mode/position |
| `style` | `ScrollBarStyle` | `'single'` / `'block'` | Visual style (see below) |
| `color` | `string` | `undefined` | Color of scroll bar characters |
| `dimColor` | `boolean` | `false` | Whether to dim the scroll bar |
| `autoHide` | `boolean` | `false` | Hide when content fits (Inset mode only) |
| `thumbChar` | `string` | auto | Custom thumb character (Inset only) |
| `trackChar` | `string` | auto | Custom track character (Inset only) |

### `<ScrollBarBox>` — Wrapper

Inherits all Ink Box props, plus:

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `contentHeight` | `number` | **required** | Total content height |
| `viewportHeight` | `number` | **required** | Viewport height |
| `scrollOffset` | `number` | **required** | Current scroll position |
| `scrollBarPosition` | `'left' \| 'right'` | `'right'` | Which border to replace |
| `scrollBarAutoHide` | `boolean` | `false` | Hide thumb if content fits |

---

## Styles

### Border Mode Styles (match Ink borders)
`single` (┌┐), `double` (╔╗), `round` (╭╮), `bold` (┏┓), `singleDouble`, `doubleSingle`, `classic` (+--+), `arrow`

### Inset Mode Styles
| Style | Thumb | Track |
|-------|-------|-------|
| `block` | █ | ░ |
| `line` | │ | (blank) |
| `thick` | ┃ | ╏ |
| `dots` | ● | · |

---

## Critical Patterns

### Border Mode — Remove the matching border from your container

```tsx
// ✅ CORRECT
<Box flexDirection="row">
  <Box borderStyle="single" borderRight={false}>
    <Content />
  </Box>
  <ScrollBar placement="right" style="single"
    contentHeight={100} viewportHeight={20} scrollOffset={offset} />
</Box>

// ❌ WRONG — double border where scrollbar sits
<Box borderStyle="single">
  <Content />
</Box>
<ScrollBar placement="right" />  {/* clashes with border */}
```

### Inset Mode — autoHide for clean UX when content fits

```tsx
<Box borderStyle="round" padding={1}>
  <Box flexDirection="row">
    <Content />
    <ScrollBar placement="inset" style="block" color="cyan"
      contentHeight={totalItems}
      viewportHeight={10}
      scrollOffset={offset}
      autoHide   // ← completely hides when content ≤ viewport
    />
  </Box>
</Box>
```

### ScrollBarBox — Easiest Integration

```tsx
<ScrollBarBox
  height={12} width={40}
  borderStyle="single"
  scrollBarPosition="right"
  contentHeight={totalItems}
  viewportHeight={10}
  scrollOffset={offset}
>
  <Content />
</ScrollBarBox>
```
