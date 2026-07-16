# ink-scroll-view — Reference

> **NPM**: `ink-scroll-view` | **Stars**: ⭐10 | **Forks**: 2
> **GitHub**: https://github.com/ByteLandTechnology/ink-scroll-view
> **Docs**: https://ink-scroll-view.byteland.app
> **License**: MIT | **Status**: Active, stable
> **Peer deps**: `ink >=6`, `react >=19`

A robust, performance-optimized ScrollView component for Ink CLI applications. A layout primitive that handles content larger than the visible terminal viewport.

---

## Architecture

The component renders ALL children but shifts content vertically using `marginTop`. The parent box with `overflow="hidden"` acts as the viewport. This avoids layout thrashing.

```
┌─────────────────────────┐
│ (hidden content)        │ ← Content above viewport
├─────────────────────────┤ ← scrollOffset (marginTop: -N)
│ ┌───────────────────┐   │
│ │ Visible Viewport  │   │ ← What user sees (overflow: hidden)
│ │                   │   │
│ └───────────────────┘   │
├─────────────────────────┤
│ (hidden content)        │ ← Content below viewport
└─────────────────────────┘
```

---

## Components

### `<ScrollView>` — Main component

Extends `BoxProps` from Ink. Accepts all Ink Box props.

**ScrollViewProps**:

| Prop | Type | Description |
|------|------|-------------|
| `children` | `ReactNode` | Child elements. **MUST have unique `key` props** |
| `onScroll` | `(offset: number) => void` | Called when scroll position changes |
| `onViewportSizeChange` | `(layout: {width, height}) => void` | Called when viewport dimensions change |
| `onContentHeightChange` | `(height: number) => void` | Called when total content height changes |
| `onItemHeightChange` | `(index, height, prevHeight) => void` | Called when an individual item's height changes |
| `debug` | `boolean` | If true, overflows content instead of hiding (debugging) |

**ScrollViewRef methods** (access via `ref.current`):

| Method | Description |
|--------|-------------|
| `scrollTo(offset)` | Scroll to absolute Y offset from top |
| `scrollBy(delta)` | Scroll relative (negative=up, positive=down) |
| `scrollToTop()` | Scroll to offset 0 |
| `scrollToBottom()` | Scroll to max offset (contentHeight - viewportHeight) |
| `getScrollOffset()` | Returns current scroll offset |
| `getContentHeight()` | Returns total content height |
| `getViewportHeight()` | Returns visible area height |
| `getBottomOffset()` | Returns scroll offset when scrolled to bottom |
| `getItemHeight(index)` | Returns measured height of item at index |
| `getItemPosition(index)` | Returns `{top, height}` of item at index |
| `remeasure()` | Re-check viewport dimensions. **Must call on terminal resize** |
| `remeasureItem(index)` | Force re-measure of specific child (for expand/collapse) |

### `<ControlledScrollView>` — Advanced

For cases needing full control over scroll state (sync multiple views, animate transitions). Accepts `scrollOffset` prop instead of managing it internally.

```tsx
const [offset, setOffset] = useState(0);
<ControlledScrollView scrollOffset={offset}>
  {children}
</ControlledScrollView>
```

---

## Features

| Feature | Details |
|---------|---------|
| **Optimistic Updates** | Immediate state updates for smoother interaction |
| **Efficient Re-rendering** | Manages visibility via `overflow` + offsets, no layout thrashing |
| **Auto-Measurement** | Automatically measures child heights using virtually rendered DOM |
| **Dynamic Content** | Supports adding, removing, expanding/collapsing items on the fly |
| **Layout Stability** | Maintains scroll position context when content changes |

---

## Critical Patterns

### [CRITICAL] Parent Handles All Input

ScrollView does NOT capture user input. Parent must use `useInput` and call ref methods:

```tsx
const scrollRef = useRef<ScrollViewRef>(null);

useInput((input, key) => {
  if (key.upArrow) scrollRef.current?.scrollBy(-1);
  if (key.downArrow) scrollRef.current?.scrollBy(1);
  if (key.pageUp) {
    const h = scrollRef.current?.getViewportHeight() || 1;
    scrollRef.current?.scrollBy(-h);
  }
  if (key.pageDown) {
    const h = scrollRef.current?.getViewportHeight() || 1;
    scrollRef.current?.scrollBy(h);
  }
});
```

### [CRITICAL] Terminal Resize → remeasure()

```tsx
useEffect(() => {
  const handleResize = () => scrollRef.current?.remeasure();
  const stdout = process.stdout;  // or from useStdout()
  stdout.on("resize", handleResize);
  return () => stdout.off("resize", handleResize);
}, [stdout]);
```

### [CRITICAL] Unique Keys on Children

```tsx
// ✅ CORRECT
<ScrollView>
  {items.map(item => <Text key={item.id}>{item.text}</Text>)}
</ScrollView>

// ❌ WRONG — breaks height tracking
<ScrollView>
  {items.map((item, i) => <Text key={i}>{item.text}</Text>)}
</ScrollView>
```

### Use remeasureItem for expand/collapse

```tsx
const handleExpand = (index: number) => {
  setItems(prev => prev.map((item, i) =>
    i === index ? { ...item, expanded: !item.expanded } : item
  ));
  // Signal item height changed
  scrollRef.current?.remeasureItem(index);
};
```

---

## Complete Example

```tsx
import React, { useRef, useEffect } from "react";
import { render, Text, Box, useInput, useStdout } from "ink";
import { ScrollView, ScrollViewRef } from "ink-scroll-view";

const App = () => {
  const scrollRef = useRef<ScrollViewRef>(null);
  const { stdout } = useStdout();

  // Terminal resize
  useEffect(() => {
    const handle = () => scrollRef.current?.remeasure();
    stdout?.on("resize", handle);
    return () => stdout?.off("resize", handle);
  }, [stdout]);

  // Input handling
  useInput((_, key) => {
    if (key.upArrow) scrollRef.current?.scrollBy(-1);
    if (key.downArrow) scrollRef.current?.scrollBy(1);
    if (key.pageUp) {
      const h = scrollRef.current?.getViewportHeight() || 1;
      scrollRef.current?.scrollBy(-h);
    }
    if (key.pageDown) {
      const h = scrollRef.current?.getViewportHeight() || 1;
      scrollRef.current?.scrollBy(h);
    }
  });

  return (
    <Box height={10} borderStyle="single" borderColor="green" flexDirection="column">
      <ScrollView ref={scrollRef}>
        {Array.from({ length: 50 }).map((_, i) => (
          <Text key={i}>
            Item {i + 1} - scrollable content line
          </Text>
        ))}
      </ScrollView>
    </Box>
  );
};

render(<App />);
```
