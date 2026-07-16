# Tinky Framework — Reference

> **NPM**: `tinky` | **Stars**: ⭐1 | **Forks**: 1
> **GitHub**: https://github.com/ByteLandTechnology/tinky
> **License**: MIT | **Status**: v1.9.0 (17 releases, 61 commits)
> **Tests**: Bun | **Language**: TypeScript 100%
> **Website**: https://byteland.app (project page)

**React for CLIs, re-imagined with the Taffy layout engine.** A modern React-based framework for building beautiful and interactive command-line interfaces. Supports CSS Flexbox AND CSS Grid layout (unlike Ink which only supports Yoga/Flexbox).

---

## Key Difference from Ink

| Feature | Tinky | Ink |
|---------|-------|-----|
| **Layout Engine** | Taffy (Flexbox + Grid) | Yoga (Flexbox only) |
| **CSS Grid** | ✅ Full support | ❌ |
| **Focus Management** | Built-in (`useFocus`) | Manual |
| **Accessibility** | ARIA attributes built-in | Manual |
| **Hot Reloading** | React DevTools support | Limited |
| **Incremental Rendering** | Run-level and line-level | Full-frame |

---

## Components

### `<Box>` — Fundamental building block

Like `<div>` in browser. Supports Flexbox AND Grid layouts.

```tsx
// Flexbox layout
<Box flexDirection="row" gap={2}>
  <Text>Left</Text>
  <Text>Right</Text>
</Box>

// Grid layout
<Box display="grid"
  gridTemplateColumns="1fr 2fr 1fr"
  gap={1}>
  <Text>Col 1</Text>
  <Text>Col 2</Text>
  <Text>Col 3</Text>
</Box>

// With borders
<Box borderStyle="round" borderColor="cyan" padding={1}>
  <Text>Styled Box</Text>
</Box>
```

**Flexbox Properties**: `flexDirection`, `justifyContent`, `alignItems`, `flexWrap`, `flexGrow`, `flexShrink`, `gap`

**Grid Properties**: `display="grid"`, `gridTemplateColumns`, `gridTemplateRows`, `columnGap`, `rowGap`, `justifyItems`, `alignItems`

**Border Styles**: `single` (┌┐), `double` (╔╗), `round` (╭╮), `bold` (┏┓), `classic` (+--+)

### `<Text>` — Styled text

```tsx
<Text color="blue">Blue text</Text>
<Text backgroundColor="red" color="white">Highlighted</Text>
<Text bold italic underline>Styled text</Text>
<Text color="#ff6600">Hex colors work too!</Text>
<Text color="rgb(255, 102, 0)">RGB</Text>
<Text color="ansi256:208">ANSI 256 colors</Text>
```

### `<Static>` — Non-updating content (logs, history)

```tsx
const logs = ["Log 1", "Log 2", "Log 3"];
<Static items={logs}>
  {(log, index) => <Text key={index}>{log}</Text>}
</Static>
```

### `<Transform>` — Transform output

```tsx
<Transform transform={(output) => output.toUpperCase()}>
  <Text>hello</Text>
</Transform>
{/* Renders: HELLO */}
```

### `<Newline>` / `<Spacer>` — Spacing

```tsx
<Newline count={2} />           // 2 blank lines
<Spacer />                      // Flexible space in flex containers
```

---

## Hooks

### `useInput` — Keyboard handling

```tsx
import { useInput, useApp } from "tinky";

function MyComponent() {
  const { exit } = useApp();
  useInput((input, key) => {
    if (key.escape) exit();
    if (key.upArrow) { /* handle */ }
    if (input === "q") exit();
  });
  return <Text>Press 'q' to quit</Text>;
}
```

### `useApp` — App lifecycle

```tsx
const { exit } = useApp();
exit(new Error("Something went wrong"));  // exit with error
exit();                                     // exit normally
```

### `useFocus` / `useFocusManager` — Focus management

```tsx
function FocusableItem({ label }: { label: string }) {
  const { isFocused } = useFocus();
  return (
    <Box borderStyle={isFocused ? "bold" : "single"}>
      <Text color={isFocused ? "green" : "white"}>{label}</Text>
    </Box>
  );
}
```

### `useStdin`, `useStdout`, `useStderr` — Stream access

```tsx
const { write } = useStdout();
useEffect(() => { write("Hello from stdout!\n"); }, []);
```

---

## `render()` Function

```tsx
import { render } from "tinky";

const { unmount, waitUntilExit, rerender, clear } = render(<App />, {
  stdout: process.stdout,
  stdin: process.stdin,
  stderr: process.stderr,
  exitOnCtrlC: true,
  patchConsole: true,
});

await waitUntilExit();
```

### Incremental Rendering

```tsx
// Run-level (default) — diffs terminal cells, writes minimal runs
render(<App />, { incrementalRendering: true });

// Line-level — diffs lines, rewrites changed lines
render(<App />, { incrementalRendering: { strategy: "line" } });

// Disable
render(<App />, { incrementalRendering: { enabled: false } });
```

Auto-falls back to non-run paths in `debug`, screen-reader, and CI environments.

### `measureElement(ref)` — Measure rendered dimensions

```tsx
import { measureElement } from "tinky";

function MyComponent() {
  const ref = useRef(null);
  useEffect(() => {
    if (ref.current) {
      const { width, height } = measureElement(ref.current);
      console.log(`Size: ${width}x${height}`);
    }
  }, []);
  return <Box ref={ref}>Content</Box>;
}
```

---

## Testing (Bun)

```bash
bun test                    # Run test suite
bun run perf:render         # Benchmark incremental rendering
bun run perf:gate           # CI performance threshold check
```

---

## Complete Example

```tsx
import { render, Box, Text, useInput, useApp } from "tinky";

function Counter() {
  const [count, setCount] = useState(0);
  const { exit } = useApp();

  useInput((input, key) => {
    if (input === "+" || key.upArrow) setCount(c => c + 1);
    if (input === "-" || key.downArrow) setCount(c => c - 1);
    if (input === "q") exit();
  });

  return (
    <Box flexDirection="column" padding={1} alignItems="center">
      <Text bold>Counter: {count}</Text>
      <Text dimColor>Press +/- to change, q to quit</Text>
    </Box>
  );
}

render(<Counter />);
```
