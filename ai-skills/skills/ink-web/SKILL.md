---
name: ink-web
description: "Complete reference for Ink Web (ink-web) — a shadcn-style component registry for browser-rendered terminal UIs. Covers all 11 official UI components: Ascii, Chat, Gradient, Link, Modal, MultiSelect, SelectInput, Spinner, StatusBar, TabBar, Table, and TextInput. Also covers setup, alias configuration, useTextInput hook, ACP-aligned Chat types, and best practices. Trigger: When building terminal interfaces in the browser, using Ink components via shadcn registry, rendering Ink apps in Next.js/Vite with Xterm.js, or needing clickable links, gradients, ASCII art, chat panels, tables, modals, select inputs, spinners, status bars, tab bars, or multi-select in terminal UIs."
license: "MIT"
metadata:
  version: "1.0"
  type: domain
  skills:
    dependencies: [byteland-ecosystem]
---

# Ink Web — Complete Reference

> **Author**: [Chris Roth](https://cjroth.com)
> **Repo**: [github.com/cjroth/ink-web](https://github.com/cjroth/ink-web)
> **Registry**: shadcn-compatible component registry
> **Stack**: React, Ink, ink-web (browser-compatible Ink), Xterm.js, shadcn CLI
> **License**: MIT

Ink Web provides a collection of React UI components built for use with [Ink](https://github.com/vadimdemedes/ink) that can run in the **browser** via Xterm.js. Components are installed via the [shadcn registry](https://ui.shadcn.com/docs/registry/registry-index).

---

## When to Use

Trigger this skill when:

- Building terminal interfaces that render in the **browser** (Next.js, Vite, etc.)
- Using **shadcn-style registry** to install Ink components
- Rendering Ink applications via **Xterm.js** in the browser
- Needing clickable **links**, **gradients**, **ASCII art**, **chat panels**, **tables**, **modals**, **select/multi-select inputs**, **spinners**, **status bars**, **tab bars**, or **text inputs** in terminal UIs
- Setting up **aliases** to redirect `ink` -> `ink-web` for browser compatibility

Don't use for:

- Native Node.js/terminal-only Ink apps (use `byteland-ecosystem` skill instead)
- General React web development (use `react-expert` or `frontend-ui-engineering`)
- General CLI/terminal apps without browser rendering

---

## Architecture & Core Concepts

### How It Works

```
┌─────────────────────────────────────────┐
│           Your App (Next.js/Vite)        │
├─────────────────────────────────────────┤
│  import { Box, Text } from 'ink'         │
│         v  (aliased at build time)       │
│  import { Box, Text } from 'ink-web'     │
│         v                                │
│  ink-web provides browser-compatible     │
│  shims + polyfills for Ink's Node APIs   │
│         v                                │
│  <InkTerminalBox> or <InkXterm> wrapper  │
│         v                                │
│  Renders in browser via Xterm.js          │
└─────────────────────────────────────────┘
```

### Key Principle: `ink` -> `ink-web` Alias

Components import from `ink`. Your bundler alias redirects to `ink-web` at build time, which includes all browser-compatible shims and polyfills. This means:
- Components work in both Node.js (real `ink`) and browser (`ink-web`)
- Full TypeScript support from `ink` types
- No changes needed to component source code

### Required Wrapper Components

- **`InkTerminalBox`** -- Simple terminal wrapper (lighter weight)
- **`InkXterm`** -- Full Xterm.js terminal with more features
- Both require `focus` prop and wrapping in `"use client"` + `dynamic()` import

---

## Setup & Installation

### 1. Configure Bundler Alias

**Next.js (Turbopack)** -- `next.config.mjs`:
```js
/** @type {import('next').NextConfig} */
const config = {
  turbopack: {
    resolveAlias: {
      ink: 'ink-web',
    },
  },
};
export default config;
```

**Next.js (Webpack)** -- `next.config.mjs`:
```js
const config = {
  webpack: (config) => {
    config.resolve.alias = {
      ...config.resolve.alias,
      ink: 'ink-web',
    };
    return config;
  },
};
export default config;
```

**Vite** -- `vite.config.ts`:
```ts
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      ink: 'ink-web',
    },
  },
});
```

### 2. Install Ink Types (Dev Dependency)
```bash
npm install -D ink
```

### 3. Initialize shadcn
```bash
npx shadcn@latest init
```

### 4. Install Components via Registry
```bash
npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/<component>.json
```

### 5. Basic Usage Pattern

```tsx
// app/page.tsx -- top-level
"use client";
import dynamic from "next/dynamic";
const Terminal = dynamic(() => import("./terminal"), { ssr: false });
export default function Home() { return <Terminal />; }

// app/terminal.tsx
"use client";
import { Box } from "ink";
import { InkXterm } from "ink-web";
import { Component } from '@/components/ui/component'
import "ink-web/css";
import "@xterm/xterm/css/xterm.css";

export default function Terminal() {
  return (
    <InkXterm focus>
      <Box flexDirection="column">
        <Component />
      </Box>
    </InkXterm>
  );
}
```

---

## Component Reference

### 1. Ascii -- ASCII Art Text

Render text as ASCII art using figlet fonts.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/ascii.json`
**Dep**: `npm install figlet`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `text` | `string` | required | The text to render |
| `font` | `string` | `'Standard'` | The figlet font |
| `horizontalLayout` | `string` | `'default'` | Horizontal layout mode |
| `verticalLayout` | `string` | `'default'` | Vertical layout mode |
| `color` | `string` | -- | Color of the ASCII art |

**Font Registration** (required -- at module load time, before component renders):
```tsx
import figlet from 'figlet'
import standard from 'figlet/importable-fonts/Standard.js'
figlet.parseFont('Standard', standard)
```

**Critical**: Do NOT use `figlet.preloadFonts()` in browser environments.

**Popular fonts**: `Standard`, `Big`, `Mini`, `Slant`, `Banner`, `Block`, `Bubble`, `Digital`, `Ivrit`, `Lean`, `Script`, `Shadow`, `Small`, `Speed`, `Star Wars`

**Examples**:
```tsx
<Ascii text="Hello" />
<Ascii text="World" color="cyan" />
<Ascii text="CLI App" font="Big" color="cyan" />
<Box flexDirection="column" gap={1}>
  <Ascii text="Doom" font="Doom" />
  <Ascii text="Ghost" font="Ghost" />
</Box>
```

---

### 2. Chat -- AI Chat Panel

Full chat panel with message list, streaming text, tool call display, and text input. ACP-aligned types.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/chat.json`

**ChatPanel Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `messages` | `ChatMessage[]` | required | Array of chat messages |
| `streamingText` | `string` | `""` | Text being streamed |
| `isLoading` | `boolean` | `false` | Show loading indicator |
| `activeToolCalls` | `ToolCallInfo[]` | `[]` | Active tool calls |
| `onSendMessage` | `(text) => void` | -- | User submits message |
| `onCancel` | `() => void` | -- | User cancels |
| `placeholder` | `string` | `"Type a message..."` | Input placeholder |
| `promptChar` | `string` | `"> "` | Prompt character |
| `promptColor` | `string` | `"green"` | Prompt color |
| `userColor` | `string` | `"green"` | User prefix color |
| `assistantColor` | `string` | `"blue"` | Assistant prefix color |
| `loadingText` | `string` | `"Thinking..."` | Loading text |

**Types (ACP-aligned)**:
```ts
interface ChatMessage {
  id: string
  role: 'user' | 'assistant'
  content: string
  toolCalls?: ToolCallInfo[]
}

interface ToolCallInfo {
  id: string
  title: string
  status: 'pending' | 'in_progress' | 'completed' | 'failed'
  result?: string
}
```

**Examples**:
```tsx
<ChatPanel messages={messages} onSendMessage={handleSend} />
<ChatPanel messages={messages} streamingText="Typing..." onSendMessage={handleSend} />
<ChatPanel messages={messages} activeToolCalls={[
  { id: '1', title: 'search', status: 'in_progress' },
]} onSendMessage={handleSend} />
<Modal title="AI Chat" onClose={() => setOpen(false)}>
  <ChatPanel messages={messages} onSendMessage={handleSend} />
</Modal>
```

---

### 3. Gradient -- Colored Text Gradients

Zero-dependency color gradients for terminal text.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/gradient.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `name` | `GradientName` | -- | Built-in gradient name |
| `colors` | `string[]` | -- | Custom hex colors |
| `children` | `ReactNode` | -- | Content for gradient |

**Built-in Gradients**: `rainbow`, `pastel`, `instagram`, `retro`, `cristal`, `teen`, `mind`, `morning`, `vice`, `passion`, `fruit`, `atlas`, `summer`

**Examples**:
```tsx
<Gradient name="rainbow"><Text>Rainbow</Text></Gradient>
<Gradient name="instagram"><Text bold>Instagram</Text></Gradient>
<Gradient colors={['#ff0000', '#00ff00', '#0000ff']}><Text>Custom</Text></Gradient>
<Gradient name="passion">
  <Box flexDirection="column"><Text>Line 1</Text><Text>Line 2</Text></Box>
</Gradient>
```

---

### 4. Link -- Clickable Hyperlinks

OSC 8 escape sequence hyperlinks for modern terminals.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/link.json`

**Props**: `children` (ReactNode), `url` (string -- required)

**Examples**:
```tsx
<Link url="https://github.com">GitHub</Link>
<Box><Text>Check </Text><Link url="https://github.com">GitHub</Link><Text> for more</Text></Box>
<Box flexDirection="column" gap={1}>
  <Link url="https://github.com">GitHub</Link>
  <Link url="https://twitter.com">Twitter</Link>
</Box>
```

**Supported in**: iTerm2, Windows Terminal, GNOME Terminal.

---

### 5. Modal -- Full-Screen Overlay

Full-screen modal with border, title, and Escape-key dismiss.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/modal.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `children` | `ReactNode` | -- | Content inside modal |
| `title` | `string` | -- | Optional title |
| `borderColor` | `string` | `"blue"` | Border color |
| `borderStyle` | `string` | `"round"` | Border style |
| `onClose` | `() => void` | -- | Called on Escape |

**Examples**:
```tsx
<Modal onClose={() => setOpen(false)}>
  <Box paddingX={1}><Text>Content</Text></Box>
</Modal>
<Modal title="Settings" onClose={() => setOpen(false)}>...</Modal>
<Modal title="Warning" borderColor="red" borderStyle="double" onClose={() => setOpen(false)}>
  <Box paddingX={1}><Text color="red">Error!</Text></Box>
</Modal>
```

---

### 6. MultiSelect -- Multiple Selection Input

Terminal multi-select with keyboard navigation.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/multi-select.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `items` | `MultiSelectItem<V>[]` | `[]` | Items with `label`, `value` |
| `selected` | `MultiSelectItem<V>[]` | -- | Controlled selection |
| `defaultSelected` | `MultiSelectItem<V>[]` | `[]` | Initial selection |
| `focus` | `boolean` | `true` | Listens to input |
| `initialIndex` | `number` | `0` | Initial highlight index |
| `limit` | `number` | -- | Scroll limit |
| `onSelect` | `(item) => void` | -- | Item selected |
| `onUnselect` | `(item) => void` | -- | Item unselected |
| `onSubmit` | `(items) => void` | -- | Enter pressed |
| `onHighlight` | `(item) => void` | -- | Highlight changed |

**Keyboard**: `up/k` -- up, `down/j` -- down, `Space` -- toggle, `Enter` -- submit

**Examples**:
```tsx
const items = [{ label: 'Apple', value: 'apple' }, { label: 'Banana', value: 'banana' }];
<MultiSelect items={items} onSubmit={(s) => console.log(s)} />
<MultiSelect items={items} defaultSelected={[items[1]]} onSubmit={handleSubmit} />
<MultiSelect items={items} limit={5} onSubmit={handleSubmit} />
```

---

### 7. SelectInput -- Single Selection Input

Terminal select with keyboard navigation.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/select-input.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `items` | `SelectInputItem<V>[]` | `[]` | Items with `label`, `value` |
| `isFocused` | `boolean` | `true` | Listens to input |
| `initialIndex` | `number` | `0` | Initial index |
| `limit` | `number` | -- | Scroll limit |
| `onSelect` | `(item) => void` | -- | Enter pressed |
| `onHighlight` | `(item) => void` | -- | Navigation |

**Keyboard**: `up/k` -- up, `down/j` -- down, `Enter` -- select, `1-9` -- instant select

**Multiple Select Inputs**: Use `isFocused` to route input:
```tsx
const [active, setActive] = useState('first');
<SelectInput items={first} isFocused={active==='first'} onSelect={(item) => { handle(item); setActive('second'); }} />
<SelectInput items={second} isFocused={active==='second'} onSelect={handle} />
```

---

### 8. Spinner -- Loading Animation

Animated loading spinner.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/spinner.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `text` | `string` | `"Loading"` | Text next to spinner |
| `color` | `string` | `"gray"` | Spinner/text color |
| `interval` | `number` | `100` | Animation speed (ms) |

**Examples**:
```tsx
<Spinner />
<Spinner text="Thinking" />
<Spinner text="Processing" color="cyan" />
<Spinner text="Loading" interval={200} />
```

---

### 9. StatusBar -- Keybinding Hints Bar

Bottom bar showing keyboard shortcuts.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/status-bar.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `items` | `StatusBarItem[]` | -- | Items with `key` + `label` |
| `extra` | `React.ReactNode` | -- | Content before items |

```ts
interface StatusBarItem {
  key: string    // inverse bold badge
  label: string  // dimmed description
}
```

**Examples**:
```tsx
<StatusBar items={[
  { key: "q", label: "quit" },
  { key: "Tab", label: "switch" },
]} />
<StatusBar items={[{ key: "q", label: "quit" }]}
  extra={<Text color="green">● Connected</Text>} />
```

---

### 10. TabBar -- Horizontal Tab Navigation

Horizontal tab bar with inverse highlight.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/tab-bar.json`

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `options` | `string[]` | -- | Tab labels |
| `selectedIndex` | `number` | -- | Selected tab index |
| `label` | `string` | -- | Label before tabs |
| `focused` | `boolean` | `true` | Styling |
| `activeColor` | `string` | `'cyan'` | Active tab color |

**Display-only** -- handle keyboard yourself:
```tsx
const [idx, setIdx] = useState(0);
useInput((_, key) => {
  if (key.leftArrow) setIdx(i => i > 0 ? i - 1 : tabs.length - 1);
  if (key.rightArrow) setIdx(i => i < tabs.length - 1 ? i + 1 : 0);
});
<TabBar options={tabs} selectedIndex={idx} />
```

---

### 11. Table -- Data Table with Borders

Table with box-drawing borders, alignment, per-cell styling, and footers.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/table.json`

**Simple mode**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `data` | `T[]` | required | Array of objects |
| `columns` | `(keyof T)[]` | all keys | Fields to show |
| `padding` | `number` | `1` | Cell padding |

**Advanced mode**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `columns` | `Column[]` | required | Column defs |
| `rows` | `Cell[][]` | required | Row data |
| `footerRows` | `Cell[][]` | `[]` | Footer rows |
| `padding` | `number` | `1` | Padding |

**Column**: `header` (string), `width` (number), `align` (left/right), `headerColor` (string)
**Cell**: `text` (string), `color` (string), `bold` (bool), `dimColor` (bool), `node` (ReactNode)

**Output**:
```
╭───────┬─────┬─────────╮
│ name  │ age │ city    │
├───────┼─────┼─────────┤
│ Alice │ 30  │ NYC     │
╰───────┴─────┴─────────╯
```

**Examples**:
```tsx
<Table data={[{ name: 'Alice', age: 30 }]} />
<Table data={data} columns={['name', 'email']} />
<Table columns={[{ header: 'Item' }, { header: 'Price', align: 'right' }]}
  rows={[[{ text: 'Widget' }, { text: '$10' }]]} />
<Table columns={[{ header: 'Status' }, { header: 'Count' }]}
  rows={[[{ text: 'Passing', color: 'green' }, { text: '42' }]]}
  footerRows={[[{ text: 'Total', bold: true }, { text: '42', bold: true }]]} />
```

---

### 12. TextInput -- Terminal Text Input

Terminal text input with keyboard handling, placeholder, prompt, and history.

**Install**: `npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/text-input.json`

**useTextInput Hook**:
```ts
const { value, setValue, history, clear, clearHistory } = useTextInput({
  onSubmit: (value) => console.log(value)
});
```

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `value` | `string` | -- | Controlled value |
| `onChange` | `(v) => void` | -- | Value changed |
| `onSubmit` | `(v) => void` | -- | Enter pressed |
| `placeholder` | `string` | `""` | Placeholder |
| `prompt` | `string` | `"> "` | Prompt string |
| `promptColor` | `string` | `"cyan"` | Prompt color |
| `cursorColor` | `string` | -- | Cursor color |
| `focus` | `boolean` | `true` | Focused |

**Examples**:
```tsx
<TextInput onSubmit={(v) => console.log(v)} />
<TextInput placeholder="Name..." onSubmit={(v) => console.log(v)} />
<TextInput prompt="$ " promptColor="green" onSubmit={(v) => console.log(v)} />
const [v, setV] = useState('');
<TextInput value={v} onChange={setV} onSubmit={(x) => { console.log(x); setV(''); }} />
```

---

## Best Practices

1. **Always use `dynamic(() => import(...), { ssr: false })`**
2. **Import CSS in order**: `"ink-web/css"` then `"@xterm/xterm/css/xterm.css"`
3. **Wrapper must have `focus` prop**: `<InkXterm focus>`
4. **Wrap in `"use client"`**
5. **Bundle aliases go in build config**, not tsconfig
6. **Figlet**: use `parseFont`, never `preloadFonts()`
7. **ChatPanel**: manage state externally
8. **TabBar**: display-only, pair with `useInput`

## Anti-Patterns

- Never import directly from `ink-web` -- use shadcn registry
- Never use `figlet.preloadFonts()` in browser
- Never forget the bundler alias (`ink` -> `ink-web`)
- Never render Ink components on the server
- Never skip CSS imports

## Registry URLs

| Component | URL |
|-----------|-----|
| Ascii | `.../registry/ascii.json` |
| Chat | `.../registry/chat.json` |
| Gradient | `.../registry/gradient.json` |
| Link | `.../registry/link.json` |
| Modal | `.../registry/modal.json` |
| MultiSelect | `.../registry/multi-select.json` |
| SelectInput | `.../registry/select-input.json` |
| Spinner | `.../registry/spinner.json` |
| StatusBar | `.../registry/status-bar.json` |
| TabBar | `.../registry/tab-bar.json` |
| Table | `.../registry/table.json` |
| TextInput | `.../registry/text-input.json` |

Base: `https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/`
