# ink-scroll-list — Reference

> **NPM**: `ink-scroll-list` | **Stars**: ⭐8 | **Forks**: 1
> **GitHub**: https://github.com/ByteLandTechnology/ink-scroll-list
> **Docs**: https://ink-scroll-list.byteland.app
> **License**: MIT | **Status**: Active, v0.4.0
> **Peer deps**: `ink >=6`, `react >=19`, `ink-scroll-view`
> **Built on**: `ink-scroll-view` (inherits all its props + ref methods)

A high-level ScrollList component built on top of `ink-scroll-view` with focus management and item selection.

---

## [CRITICAL] v0.4.0 Breaking Change — Now Fully Controlled

**v0.4.0** converted ScrollList to a **fully controlled component**. The parent manages ALL selection state.

### Removed APIs (v0.3.x → v0.4.0 migration):

| Removed | Replacement |
|---------|-------------|
| `onSelectionChange` prop | Parent owns `selectedIndex` state directly |
| `selectNext()`, `selectPrevious()` | `setSelectedIndex(prev => Math.min(prev + 1, len-1))` |
| `selectFirst()`, `selectLast()` | `setSelectedIndex(0)` / `setSelectedIndex(len-1)` |
| `scrollToItem(index, mode)` | Set `selectedIndex` prop with `scrollAlignment` |
| `getSelectedIndex()` | Parent already knows from state |
| `getItemCount()` | Parent already knows from array |

### Migration Example

```tsx
// BEFORE (v0.3.x)
const listRef = useRef<ScrollListRef>(null);
const [selectedIndex, setSelectedIndex] = useState(0);
useInput((_, key) => {
  if (key.downArrow) {
    const newIndex = listRef.current?.selectNext() ?? 0;
    setSelectedIndex(newIndex);
  }
});
<ScrollList ref={listRef} selectedIndex={selectedIndex}
  onSelectionChange={setSelectedIndex} />

// AFTER (v0.4.0)
const listRef = useRef<ScrollListRef>(null);
const [selectedIndex, setSelectedIndex] = useState(0);
useInput((_, key) => {
  if (key.downArrow) {
    setSelectedIndex(prev => Math.min(prev + 1, items.length - 1));
  }
});
<ScrollList ref={listRef} selectedIndex={selectedIndex} />
```

---

## Props

Extends `ScrollViewProps` from `ink-scroll-view`.

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `selectedIndex` | `number` | **required** | Currently selected item index (controlled) |
| `scrollAlignment` | `'auto' \| 'top' \| 'bottom' \| 'center'` | `'auto'` | How selected item aligns in viewport |

### Scroll Alignment Modes

| Mode | Behavior | Best For |
|------|----------|----------|
| `'auto'` | Minimal scrolling to bring item into view | Keyboard navigation (default) |
| `'top'` | Always align selected item to viewport top | — |
| `'bottom'` | Always align selected item to viewport bottom | — |
| `'center'` | Always center selected item in viewport | Search/spotlight UX |

---

## Ref Methods

Extends `ScrollViewRef`. When `selectedIndex` is set, scrolling is **constrained** to keep selection visible.

| Method | Description |
|--------|-------------|
| `scrollTo(y)` | Scroll constrained to keep selection visible |
| `scrollBy(delta)` | Scroll by delta, constrained |
| `scrollToTop()` | Max scroll up keeping selection visible |
| `scrollToBottom()` | Max scroll down keeping selection visible |
| All `ScrollViewRef` methods | `getScrollOffset`, `getContentHeight`, `getViewportHeight`, `getBottomOffset`, `getItemHeight`, `getItemPosition`, `remeasure`, `remeasureItem` |

**Large Items**: If an item is larger than the viewport, scrolling is allowed within its bounds so users can see different parts.

---

## Critical Patterns

### Parent Handles Selection + Input

```tsx
const [selectedIndex, setSelectedIndex] = useState(0);
const items = Array.from({ length: 20 }, (_, i) => `Item ${i + 1}`);

useInput((input, key) => {
  if (key.upArrow) setSelectedIndex(p => Math.max(p - 1, 0));
  if (key.downArrow) setSelectedIndex(p => Math.min(p + 1, items.length - 1));
  if (input === "g") setSelectedIndex(0);        // jump to first
  if (input === "G") setSelectedIndex(items.length - 1); // jump to last
  if (key.return) console.log(`Selected: ${items[selectedIndex]}`);
});
```

### Dynamic Items — Adjust selectedIndex

```tsx
// Adding items at beginning
setSelectedIndex(prev => prev + addedCount);

// Removing items — clamp to valid range
setSelectedIndex(prev => Math.min(prev, newLength - 1));
```

### Complete Example

```tsx
import React, { useRef, useState, useEffect } from "react";
import { render, Text, Box, useInput } from "ink";
import { ScrollList, ScrollListRef } from "ink-scroll-list";

const App = () => {
  const listRef = useRef<ScrollListRef>(null);
  const [selectedIndex, setSelectedIndex] = useState(0);
  const items = Array.from({ length: 20 }, (_, i) => `Item ${i + 1}`);

  useEffect(() => {
    const h = () => listRef.current?.remeasure();
    process.stdout.on("resize", h);
    return () => process.stdout.off("resize", h);
  }, []);

  useInput((input, key) => {
    if (key.upArrow) setSelectedIndex(p => Math.max(p - 1, 0));
    if (key.downArrow) setSelectedIndex(p => Math.min(p + 1, items.length - 1));
    if (input === "g") setSelectedIndex(0);
    if (input === "G") setSelectedIndex(items.length - 1);
    if (key.return) console.log(`Selected: ${items[selectedIndex]}`);
  });

  return (
    <Box borderStyle="single" height={10}>
      <ScrollList ref={listRef} selectedIndex={selectedIndex}>
        {items.map((item, i) => (
          <Box key={i}>
            <Text color={i === selectedIndex ? "green" : "white"}>
              {i === selectedIndex ? "> " : "  "}{item}
            </Text>
          </Box>
        ))}
      </ScrollList>
    </Box>
  );
};

render(<App />);
```
