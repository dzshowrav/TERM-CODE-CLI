# ink-multiline-input — Expert Knowledge

> **Source**: https://github.com/ByteLandTechnology/ink-multiline-input.git
> **Package**: `ink-multiline-input` v0.1.0 | MIT | ESM
> **Stack**: React 19, Ink 6, TypeScript, tsup, Vitest
> **Author**: ByteLand Technology

A multi-line text input component library for **Ink** (React-based CLI framework) applications. Teaches powerful patterns for building terminal UI text editors.

---

## 1. Two-Component Architecture (Display/Logic Separation)

The library splits concerns into two components — a **crucial architectural pattern**:

### A. `MultilineInput` — Interactive Controller
- **Handles**: ALL keyboard events, cursor movement, text editing
- **Props**: `value`, `onChange`, `onSubmit?`, `keyBindings?`, `highlightPastedText?`, `useCustomInput?`, `focus?`, `showCursor?`
- **Internal State**: `cursorIndex` (number — where cursor is), `pasteLength` (number — for paste detection)
- **Renders**: Delegates to `ControlledMultilineInput` with computed cursorIndex + highlight

### B. `ControlledMultilineInput` — Pure Display
- **Handles**: Text rendering with cursor, highlighting, scrolling — **NO input logic**
- **Props**: `value`, `cursorIndex`, `rows?`, `maxRows?`, `highlightStyle?`, `textStyle?`, `placeholder?`, `mask?`, `showCursor?`, `focus?`, `tabSize` (4), `highlight?`, `refreshKey?`
- **Pattern**: Splits text at cursor into preCursor/postCursor segments, each typed for styling

---

## 2. Cursor Visibility & Scrolling Algorithm

**Goal**: Keep cursor visible in terminal viewport during typing/navigation.

**Implementation**:
1. **MeasureBox** utility measures rendered height via Ink's `measureElement()`
2. Used **twice**:
   - Full text → `contentHeight` (total scrollable height)
   - Pre-cursor text only → `markerHeight` (cursor line position, 1-indexed)
3. **Scroll calculation**:
   ```
   visibleRows = max(rows, min(maxRows, contentHeight))
   viewportStart = scrollOffset
   viewportEnd = scrollOffset + visibleRows
   
   if cursorLine <= viewportStart → scroll UP (cursor at top)
   if cursorLine > viewportEnd → scroll DOWN (cursor at bottom)
   if contentHeight < viewportEnd → scroll to bottom
   ```
4. **Visual mechanism**: `marginTop: -scrollOffset` inside `overflow: hidden` Box

---

## 3. Vertical Navigation with Column Preservation

When user presses **Up/Down** arrows:

1. **Find current line & column** from absolute `cursorIndex`:
   - Split value by `\n`, walk through accumulating lengths
   - `col = cursorIndex - lineStartPosition`

2. **Calculate target**:
   ```
   targetLine = lines[currentLineIndex + 1]  // for Down
   newCol = Math.min(col, targetLine.length)  // CLAMP to avoid overshoot
   newIndex = sum(prev line lengths + newlines) + newCol
   ```

3. **Boundary check**: If at first/last line, navigation is a no-op

---

## 4. Key Features Reference

| Feature | Prop | Details |
|---------|------|---------|
| Arrow navigation | `showCursor=true` | Up/Down preserves column, Left/Right moves by char |
| Submit | `onSubmit`, keyBindings | Default: Ctrl+Enter |
| Newline | keyBindings.newline | Default: Enter |
| Mask (password) | `mask="*"` | Replaces all non-`\n` chars |
| Placeholder | `placeholder="..."` | Shown when empty AND unfocused |
| Paste highlight | `highlightPastedText` | Highlights multi-char input with highlightStyle |
| Custom input hook | `useCustomInput` | Replace Ink's useInput (for demos, routing) |
| Rows control | `rows` / `maxRows` | Min/max visible lines |
| Tab expansion | `tabSize` (default: 4) | \t → spaces |
| Custom key bindings | `keyBindings` | `{ submit: (key) => ..., newline: (key) => ... }` |

**Key Bindings API**:
```tsx
keyBindings={{
  submit: (key) => key.ctrl && key.return,   // Ctrl+Enter
  newline: (key) => key.shift && key.return,  // Shift+Enter
}}
```

---

## 5. Paste Detection Pattern

```tsx
// In input handler:
if (input.length > 1) {
  nextPasteLength = input.length;  // Detected as paste
}
// ...insert text at cursor...
setCursorIndex(cursorIndex + input.length);
setPasteLength(nextPasteLength);

// Computed highlight:
const highlight = highlightPastedText && pasteLength > 1 
  ? { start: cursorIndex - pasteLength, end: cursorIndex }
  : undefined;
```

---

## 6. Input Event Flow (in MultilineInput)

```
1. submitKey(key)? → onSubmit(value), return
2. newlineKey(key)? → insert '\n' at cursor, return
3. tab / shift+tab / ctrl+c → ignore, return
4. upArrow → vertical nav (column preserved)
5. downArrow → vertical nav (column preserved)
6. leftArrow → cursorIndex-- (clamped to 0)
7. rightArrow → cursorIndex++ (clamped to value.length)
8. return → insert '\n' at cursor
9. backspace/delete → remove char before cursor
10. default → insert input at cursor, detect paste
```

---

## 7. Display Rendering Pattern (in ControlledMultilineInput)

```
Format:
1. normalizeLineEndings() → \r\n, \r → \n
2. mask? → replace all /[^\n]/g with mask char
3. expandTabs() → \t → tabSize spaces
4. Split at cursorIndex:

   preCursor = [ {value: "text before line", type: "normal"},
                 {value: "cursor line text", type: "highlight"},
                 {value: " ", type: "cursor"} ]
   postCursor = [ {value: "rest of cursor line", type: "highlight"},
                  {value: "lines after cursor", type: "normal"} ]

5. Style mapping:
   - "normal" → textStyle
   - "highlight" → highlightStyle ?? textStyle
   - "cursor" → inverse: true + (highlightStyle ?? textStyle)
   - "placeholder" → textStyle + dimColor: true
```

---

## 8. MeasureBox Pattern (Generic Reusable)

```tsx
// A Box that reports its measured height
const MeasureBox = ({ children, onHeightChange }) => {
  const ref = useRef(null);
  const lastHeightRef = useRef(undefined);

  useEffect(() => {
    if (ref.current) {
      const { height } = measureElement(ref.current);
      if (lastHeightRef.current !== height) {
        lastHeightRef.current = height;
        onHeightChange?.(height);
      }
    }
  });  // No deps — runs on every render to catch content changes

  return <Box ref={ref}>{children}</Box>;
};
```

---

## 9. Testing Patterns (Vitest + ink-testing-library)

### Mock MeasureBox for test environment:
```tsx
vi.mock("../src/MeasureBox", () => ({
  MeasureBox: ({ children, onHeightChange }) => {
    const extractText = (node) => {
      if (!node) return "";
      if (typeof node === "string") return node;
      if (Array.isArray(node)) return node.map(extractText).join("");
      if (node.props?.children) return extractText(node.props.children);
      return "";
    };
    React.useEffect(() => {
      const lines = extractText(children).split("\n").length;
      onHeightChange?.(lines);
    }, [children, onHeightChange]);
    return <>{children}</>;
  },
}));
```

### Controlled wrapper for interactive tests:
```tsx
const Wrapped = () => {
  const [val, setVal] = React.useState("");
  return <MultilineInput value={val} onChange={setVal} />;
};
```

### Logic testing via useCustomInput:
```tsx
let capture: any;
render(
  <MultilineInput
    value="test"
    onChange={onChange}
    useCustomInput={(handler) => { capture = handler; }}
  />
);
// Simulate keys:
capture("", { return: true, ctrl: true });  // Ctrl+Enter → submit
capture("X", {});                            // Type 'X'
capture("", { backspace: true });            // Backspace
capture("", { upArrow: true });              // Arrow Up
```

---

## 10. Scrolling Demo Architecture Pattern

The auto-demo (`demo/auto-demo.tsx`) demonstrates a powerful pattern for automating Ink apps:

1. Capture input handler via `useCustomInput` → store in ref
2. Use `setTimeout` chains with programmatic handler calls
3. Include layout components (`DemoLayout`, `SplitDemoLayout`) for professional presentation
4. Run with: `npx tsx demo/auto-demo.tsx <category>`

---

## 11. Utility Functions

```typescript
// Expand tabs to spaces
expandTabs(text: string, tabSize: number): string
  → text.replace(/\t/g, ' '.repeat(tabSize))

// Normalize any line ending to \n
normalizeLineEndings(text: string): string
  → text.replace(/\r\n/g, '\n').replace(/\r/g, '\n')
  → null/undefined → ""
```

---

## 12. Project Patterns to Reuse

1. **Display/Logic component separation** — testable, reusable presentation
2. **Height measurement trick** — Ink's measureElement for scroll calculations
3. **Negative margin scrolling** in overflow:hidden containers
4. **Custom input hook injection** for testing/demos
5. **Paste detection via input.length > 1**
6. **Column preservation on vertical navigation** with Math.min clamping
7. **Typed text segments** for flexible styling
8. **Controlled component + wrapper pattern** for testing

---

## 13. Build & Tooling

- **Build**: tsup → ESM output to `dist/`
- **TypeScript**: @sindresorhus/tsconfig base, `jsx: "react-jsx"`, `target: "es2022"`
- **Git hooks**: Husky + commitlint (conventional commits) + lint-staged (Prettier)
- **Testing**: Vitest + ink-testing-library
- **Docs**: Typedoc with typedoc-plugin-markdown
- **peerDependencies**: `ink: ">=6"`, `react: ">=19"`
