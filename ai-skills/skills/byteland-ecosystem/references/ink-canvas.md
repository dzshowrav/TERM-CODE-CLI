# ink-canvas — Reference

> **NPM**: `ink-canvas` | **Stars**: ⭐2 | **Forks**: 1
> **GitHub**: https://github.com/ByteLandTechnology/ink-canvas
> **Docs**: https://ink-canvas.byteland.app
> **License**: MIT
> **Peer deps**: `ink >=6`, `react >=19`, `@xterm/xterm`, `@xterm/addon-fit`

A library for rendering Ink applications in the browser using Xterm.js. Bridges Node.js-based CLI UIs with web-based terminal emulators.

---

## Architecture

```
Browser DOM
  └─ InkCanvas component (React)
       ├─ Xterm.js Terminal (renders ANSI/text)
       ├─ TerminalWritableStream (stdout) — converts LF→CRLF, cursor ops
       ├─ TerminalReadableStream (stdin) — captures keyboard → Ink
       └─ Process Shim — mocks process.env, process.nextTick, etc.
            └─ Ink instance (your app)
```

### 1. Process Shim (`shims/process.ts`)
Browser-compatible mock of Node.js `process`:
- `process.env` — mocked env vars with sensible defaults
- `process.stdout/stderr` — minimal stream mocks with TTY properties
- `process.stdin` — input stream mock
- `process.nextTick` — implemented via `setTimeout(0)`
- `platform`, `version`, `argv`, `cwd()` — all mocked

### 2. Custom Streams (`utils/streams.ts`)
- **TerminalWritableStream** — receives ANSI codes from Ink, converts LF→CRLF, provides cursor manipulation methods
- **TerminalReadableStream** — captures Xterm.js `onData` events, buffers input, supports raw mode

### 3. Canvas Component — auto-sizes to match stdout dimensions, listens for resize

### 4. React Lifecycle Management — Xterm.js init/cleanup, stream creation, Ink lifecycle

---

## API

### `<InkCanvas>` Component

```tsx
import { InkCanvas } from "ink-canvas";
```

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `children` | `ReactNode` | — | The Ink application to render |
| `focused` | `boolean` | `false` | Whether terminal captures keyboard input |
| `cols` | `number` | `undefined` | Fixed columns (omit = fit container) |
| `rows` | `number` | `undefined` | Fixed rows (omit = fit container) |
| `terminalOptions` | `ITerminalOptions` | `{}` | Xterm.js Terminal configuration |
| `onResize` | `(dims: {cols, rows}) => void` | — | Fired when terminal dimensions change |
| `...divProps` | `HTMLAttributes<HTMLDivElement>` | — | Passed to container div |

### `InkCanvasHandle` (via ref)

| Property | Type | Description |
|----------|------|-------------|
| `terminal` | `Terminal \| null` | The Xterm.js Terminal instance |
| `dimensions` | `{cols, rows} \| null` | Current terminal columns and rows |
| `instance` | `Instance \| null` | The Ink instance returned by `render()` |

### Terminal Options (common)

```tsx
<InkCanvas
  terminalOptions={{
    fontSize: 16,
    fontFamily: "JetBrains Mono, Fira Code, monospace",
    cursorStyle: "bar",      // 'block' | 'underline' | 'bar'
    cursorBlink: true,
    theme: {
      background: "#1a1b26",
      foreground: "#a9b1d6",
      cursor: "#c0caf5",
      selectionBackground: "#33467c",
      black: "#15161e",
      red: "#f7768e",
      green: "#9ece6a",
      yellow: "#e0af68",
      blue: "#7aa2f7",
      magenta: "#bb9af7",
      cyan: "#7dcfff",
      white: "#a9b1d6",
    },
    scrollback: 1000,
    allowProposedApi: true,
  }}
>
  <MyApp />
</InkCanvas>
```

---

## Build Integration

### Vite — `inkCanvasPolyfills()` plugin

```ts
// vite.config.ts
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { inkCanvasPolyfills } from "ink-canvas/plugin";

export default defineConfig({
  plugins: [react(), inkCanvasPolyfills()],
});
```

### Webpack / Next.js — `InkCanvasWebpackPlugin`

```js
// next.config.mjs or webpack.config.js
import { InkCanvasWebpackPlugin } from "ink-canvas/plugin";

const nextConfig = {
  transpilePackages: ["ink-canvas"],
  webpack: (config, { isServer }) => {
    if (!isServer) {
      config.plugins.push(new InkCanvasWebpackPlugin());
    }
    return config;
  },
};
```

### Manual Polyfill

```bash
npm install vite-plugin-node-polyfills
```

```ts
// vite.config.ts (manual)
import { nodePolyfills } from "vite-plugin-node-polyfills";
export default defineConfig({
  plugins: [
    nodePolyfills({
      exclude: ["process"],
      globals: { Buffer: true, global: true },
      protocolImports: true,
    }),
  ],
  resolve: {
    alias: { "node:process": "ink-canvas/shims/process" },
  },
});
```

---

## Critical Patterns

### [CRITICAL] Container MUST have explicit dimensions

```tsx
// ✅ CORRECT
<InkCanvas style={{ width: "100%", height: "400px" }}>
  <MyApp />
</InkCanvas>

// ❌ WRONG — no dimensions = empty terminal
<InkCanvas>
  <MyApp />
</InkCanvas>
```

### [CRITICAL] focused MUST be true for keyboard input

```tsx
// ✅ CORRECT
<InkCanvas focused={true}>
  <InteractiveApp />
</InkCanvas>

// ❌ WRONG — keyboard wont work
<InkCanvas>
  <InteractiveApp />
</InkCanvas>
```

### Handle initial layout with onResize

```tsx
const [ready, setReady] = useState(false);
<InkCanvas onResize={() => setReady(true)}>
  {ready && <MyApp />}
</InkCanvas>
```

### TypeScript config must include DOM lib

```json
{
  "compilerOptions": {
    "lib": ["ES2020", "DOM", "DOM.Iterable"]
  }
}
```
