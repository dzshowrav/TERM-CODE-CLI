# Ultimate CLI — Complete AI CLI Agent Skill

## Metadata
- **Name**: ultimate-cli
- **Version**: 1.0.0
- **Trigger**: When building CLI applications with Ink/React, integrating AI agents, implementing terminal UIs, setting up vector search, handling file operations, parsing code, manipulating ASTs, managing git repositories, or building MCP servers.
- **Tags**: ink, react-cli, commander, yargs, blessed, chalk, mcp, langchain, llamaindex, chroma, qdrant, lancedb, tree-sitter, babel, ts-morph, jscodeshift, isomorphic-git, simple-git, prompts, inquirer, enquirer, fast-glob, globby, cli-ui
- **Platform**: Node.js, TypeScript, React 18+, Terminal
- **Dependencies**: ink, react, commander/yargs, chalk, fast-glob, tree-sitter, ts-morph, babel, isomorphic-git/simple-git, chromadb/qdrant/lancedb, @langchain/core, openai-agents-sdk

---

## 1. CORE INK ECOSYSTEM (React for CLI)

### 1.1 Ink — The Foundation

**GitHub**: `vadimdemedes/ink` — 33K+ stars, TypeScript  
**npm**: `ink`, ~2.1M weekly downloads, MIT License  
**Author**: Vadim Demedes  
**Website**: https://term.ink/  
**Last Updated**: 2025-11-19  
**Current Version**: v6.5.1

#### Architecture & Philosophy
Ink brings React's component model to the terminal. Instead of rendering to the DOM, Ink renders React components to terminal output using Yoga (Facebook's Flexbox layout engine for cross-platform layout).

**Core Principle**: Every Ink element is automatically a Flexbox container (`display: flex` in browser terms). All text MUST be wrapped in `<Text>` component.

#### Installation & Setup
```bash
npm install ink react react-reconciler
npm install --save-dev @types/react
```

**Scaffold a new Ink project:**
```bash
npx create-ink-app my-cli
```

**Manual setup for TypeScript + ESM:**
```json
// tsconfig.json
{
  "compilerOptions": {
    "jsx": "react-jsx",
    "module": "ESNext",
    "moduleResolution": "bundler",
    "target": "ES2022",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "outDir": "./dist",
    "rootDir": "./src",
    "declaration": true,
    "sourceMap": true
  }
}
```

**For JSX transpilation without bundler:**
```bash
npm install --save-dev @esbuild-kit/esm-loader
node --loader @esbuild-kit/esm-loader dist/index.js
```

#### Core API

##### `render()` — Entry Point
```typescript
import { render, type Instance } from 'ink';

const { waitUntilExit, clear, unmount, rerender } = render(<App />);

// waitUntilExit — Promise that resolves when app exits
// clear — clears output
// unmount — unmounts component
// rerender — updates root component
```

**Options:**
```typescript
render(<App />, {
  stdout: process.stdout,  // custom stdout
  stdin: process.stdin,    // custom stdin
  stderr: process.stderr,  // custom stderr
  exitOnCtrlC: true,       // Ctrl+C exits app (default: true)
  patchConsole: true,       // patch console.log (default: true)
  debug: false              // debug mode
});
```

##### Built-in Components

| Component | Purpose | Key Props |
|-----------|---------|-----------|
| `<Box>` | Flexbox container | `flexDirection`, `alignItems`, `justifyContent`, `padding`, `margin`, `width`, `height`, `borderStyle` |
| `<Text>` | Text output (ALL text must go here) | `color`, `backgroundColor`, `bold`, `italic`, `underline`, `strikethrough`, `dim`, `inverse`, `wrap` |
| `<Newline>` | Line break | `count` (default: 1) |
| `<Spacer>` | Fill remaining space in flex container | none |
| `<Static>` | Keep children mounted after initial render | `items` array |
| `<Transform>` | Transform rendered children | `transform(children: string): string` |

**Box border styles**: `single`, `double`, `round`, `bold`, `singleDouble`, `doubleSingle`, `classic`

**Text colors**: Named (256 colors) + hex (`#ffffff`) + `rgb(r,g,b)`

##### Built-in Hooks

**`useApp()`** — App lifecycle control
```typescript
import { useApp } from 'ink';

const { exit, exitError } = useApp();
exit();          // graceful exit
exitError();     // error exit
```

**`useInput()`** — Keyboard handling
```typescript
import { useInput } from 'ink';

useInput((input: string, key: KeyObject) => {
  // key: { upArrow, downArrow, leftArrow, rightArrow, return, escape, ctrl, shift, tab, backspace, delete, meta, pageUp, pageDown }
  if (input === 'q') useApp().exit();
  if (key.upArrow) moveSelection(-1);
  if (key.downArrow) moveSelection(1);
}, { isActive: true });  // isActive: only process when focused
```

**`useStdin()`** — Raw stdin access
```typescript
import { useStdin } from 'ink';
const { stdin, isRawModeSupported, setRawMode } = useStdin();
```

**`useStdout()` / `useStderr()`** — Output stream access
```typescript
import { useStdout } from 'ink';
const { stdout, write } = useStdout();
// write(data: string) — direct write to stdout
```

**`useFocus()` / `useFocusManager()`** — Focus management for multi-component apps
```typescript
import { useFocus, useFocusManager } from 'ink';

// In child component:
const { isFocused } = useFocus({ id: 'my-input', autoFocus: true });

// In parent controller:
const { enableFocus, disableFocus, focusNext, focusPrevious, setFocus } = useFocusManager();
```

**`useMeasure()`** — Component dimensions
```typescript
import { useMeasure } from 'ink';
const { ref } = useMeasure();
// Returns { width, height } of the measured element
```

#### Ink Testing
```bash
npm install --save-dev ink-testing-library
```

```typescript
import { render } from 'ink-testing-library';
const { lastFrame, frames, rerender, stdin } = render(<MyComponent />);
// lastFrame — last rendered frame as string
// frames — all rendered frames as string[]
// stdin.write(data) — simulate keyboard input
```

#### Version Migration Guide

| Ink Version | React Version | Breaking Changes |
|-------------|---------------|------------------|
| Ink 4.x | React 17/18 | Legacy API |
| Ink 5.x | React 18 | Static component changes, Memo changes |
| Ink 6.x | React 18/19 | Latest API, better hooks |

#### Ink 4 → Ink 6 Migration

1. `render()` now returns `{ waitUntilExit, clear, unmount, rerender }` instead of `{ waitUntilExit, clear, unmount }`
2. `useInput` callback signature may change
3. `useApp().exit()` replaces `process.exit()` patterns
4. `<Static>` uses `items` prop instead of children
5. `useStdin`, `useStdout`, `useStderr` return objects directly (not tuples)

---

#### Ink TL;DR (from Awesome React)
A custom React renderer that targets the terminal instead of the browser DOM. It brings the component model, hooks, and a complete Flexbox layout engine to the terminal, replacing messy `console.log` chains with declarative UI.

#### Why Ink?
- **Declarative UI**: Render state, Ink handles diffing and repainting
- **Flexbox Layout**: Yoga-based — `<Box flexDirection="column">` for complex layouts
- **Standard React Hooks**: useState, useEffect — your React knowledge transfers 100%
- **Input Handling**: `useInput` hook makes keyboard shortcuts trivial
- **Proven Stability**: Used by Gatsby CLI, Prisma CLI, Jest, Claude Code, Cursor AI

#### Ink vs Other CLI Libraries (Comparison)

| Library | Design Philosophy | Best For | Pain Points |
|---------|------------------|----------|-------------|
| **Ink** | Declarative UI — Full React renderer for terminal | Rich Applications: Installers, Dashboards, Wizards | Overhead — heavy for simple scripts |
| **Commander.js** | Argument Parsing — flags and args | Standard CLIs: tools like git/npm | No UI — no interactive interface |
| **Chalk** | Styling Only — colorizes stdout | Simple Output: making logs readable | Imperative — manual string management |
| **Inquirer.js** | Prompt Based — linear questions | Scaffolding: 5 questions in a row | Rigid — no dashboard-style UI |

#### Ink Pros
- **Superior DX**: Structured, readable, maintainable terminal UIs
- **Visual Layouts**: Flexbox for dashboards and side-by-side panels
- **Component Reusability**: `<Spinner />`, `<ProgressBar />`, `<SelectMenu />` across tools

#### Ink Cons
- **Runtime Overhead**: Requires Node.js + React (startup time + memory)
- **Terminal Quirks**: Older terminals may break on complex layouts
- **Overkill**: Use `prompts` or `inquirer` for single Yes/No questions

#### Verdict: When to Adopt Ink
**Ink** is the only viable choice for **Dashboard-style CLIs** — persisting on screen, real-time updates (download managers, server monitors), or complex navigation. For standard "fire and forget" tools, Commander + chalk is sufficient.

---

### 1.2 Ink Ecosystem Components

#### 1.2.1 `ink-text-input` — Text Input Field
**npm**: `ink-text-input` ~1.4M/week  
**GitHub**: `vadimdemedes/ink-text-input`

```tsx
import TextInput from 'ink-text-input';

function MyInput() {
  const [value, setValue] = useState('');
  return (
    <TextInput
      value={value}
      onChange={setValue}
      onSubmit={(submitted) => console.log('Submitted:', submitted)}
      placeholder="Type something..."
      focus={true}
      mask="*"          // password masking
      showCursor={true}
      highlightPastedText={true}
    />
  );
}
```

**Props**:
| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `value` | string | required | Current input value |
| `onChange` | (v: string) => void | required | Change handler |
| `onSubmit` | (v: string) => void | - | Enter pressed |
| `placeholder` | string | '' | Placeholder text |
| `focus` | boolean | true | Active/focused |
| `mask` | string | - | Mask character for passwords |
| `showCursor` | boolean | true | Show cursor |
| `highlightPastedText` | boolean | false | Highlight pasted text |

**Also exports**: `UncontrolledTextInput` — manages its own state internally

#### 1.2.2 `ink-spinner` — Loading Animation
**npm**: `ink-spinner` ~1.6M/week

```tsx
import Spinner from 'ink-spinner';

<Text color="green">
  <Spinner type="dots" /> Loading...
</Text>
```

**Props**: `type` — spinner style from `cli-spinners` (80+ animations: dots, line, pipe, simpleDots, globe, moon, monkey, etc.)

#### 1.2.3 `ink-select-input` — Single Select
**GitHub**: `vadimdemedes/ink-select-input`

```tsx
import SelectInput from 'ink-select-input';

const items = [
  { label: 'Option 1', value: 'opt1' },
  { label: 'Option 2', value: 'opt2' },
  { separator: true },  // renders separator
  { label: 'Option 3', value: 'opt3' },
];

<SelectInput
  items={items}
  isFocused={true}
  onSelect={(item) => console.log(item.value)}
  onHighlight={(item) => console.log('Highlighted:', item.label)}
  indicatorComponent={CustomIndicator}
  itemComponent={CustomItem}
  limit={5}  // max visible items (scrolls)
/>
```

**Key behavior**: Arrow up/down, j/k keys, number keys for index-based selection, Enter to confirm.

#### 1.2.4 `ink-multi-select` — Multi-Select Checkboxes
**GitHub**: `karaggeorge/ink-multi-select`  
**npm**: `ink-multi-select` ~16.2K/week

```tsx
import MultiSelect from 'ink-multi-select';

<MultiSelect
  items={[
    { label: 'Option A', value: 'a' },
    { label: 'Option B', value: 'b' },
  ]}
  selected={['a']}      // controlled mode
  onSelect={(item) => {}}
  onSubmit={(selected) => console.log(selected)}
  focus={true}
  indicatorComponent={CustomIndicator}
  checkboxComponent={CustomCheckbox}
  itemComponent={CustomItem}
  limit={10}
/>
```

**Modes**:
- **Controlled**: `selected` prop set — parent manages state
- **Uncontrolled**: `selected` prop undefined — internal state

#### 1.2.5 `ink-confirm-input` — Y/N Confirmation
**GitHub**: `kevva/ink-confirm-input`  
**npm**: `ink-confirm-input` ~28K/week  
**⚠️ Deprecated for Ink 5+** — use `@inkjs/ui` `ConfirmInput` instead.

```tsx
import ConfirmInput from 'ink-confirm-input';

<ConfirmInput
  isChecked={true}
  value="y"
  onChange={(val) => setValue(val)}
  onSubmit={(submitted) => {
    const confirmed = submitted.toLowerCase() === 'y';
    if (confirmed) proceed();
  }}
/>
```

**Modern replacement** (`@inkjs/ui`):
```tsx
import { ConfirmInput } from '@inkjs/ui';

<ConfirmInput
  defaultValue={true}
  onChange={(val) => setConfirmed(val)}
/>
```

#### 1.2.6 `ink-progress-bar` — Progress Indicator
**GitHub**: `brigand/ink-progress-bar`

```tsx
import ProgressBar from 'ink-progress-bar';

// Standalone
<ProgressBar percent={0.75} left={0} right={0} />

// Or @inkjs/ui version:
import { ProgressBar } from '@inkjs/ui';
<ProgressBar value={75} />  // 0-100
```

#### 1.2.7 `ink-table` — Tabular Data
**GitHub**: `maticzav/ink-table`

```tsx
import Table from 'ink-table';

const data = [
  { name: 'Alice', role: 'Engineer' },
  { name: 'Bob', role: 'Designer' },
];

<Table
  data={data}
  columns={['name', 'role']}    // column order
  cell={CellComponent}           // custom cell renderer
  header={HeaderComponent}       // custom header
  padding={1}
/>
```

#### 1.2.8 `ink-markdown` — Markdown Rendering
**GitHub**: `cameronhunter/ink-markdown`  
**npm**: `ink-markdown` ~6.3K/week  
**ESM**: `@inkkit/ink-markdown`

```tsx
import Markdown from 'ink-markdown';

<Markdown>
  {`
# Hello

This is **bold** and *italic* text.

- List item 1
- List item 2
  `}
</Markdown>
```

#### 1.2.9 `ink-link` — Clickable Hyperlinks
**GitHub**: `sindresorhus/ink-link`  
**npm**: `ink-link` ~180K/week

```tsx
import Link from 'ink-link';

<Link url="https://example.com" fallback={false}>
  Click here
</Link>
```

- **OSC 8 hyperlinks**: Works in iTerm2, Hyper, VS Code terminal, Kitty, etc.
- **Fallback**: Shows URL in parentheses if terminal unsupported

#### 1.2.10 `ink-picture` — Terminal Images
**GitHub**: `endernoke/ink-picture`  
**npm**: `ink-picture` ~68K/week  
**Replaces**: `ink-image` (deprecated)

```tsx
import Picture, { TerminalInfoProvider } from 'ink-picture';

const App = () => (
  <TerminalInfoProvider>
    <Picture
      src="/path/to/image.png"
      width={40}
      height={20}
      alt="Description"
      protocol="auto"       // auto | ascii | braille | halfBlock | sixel | iterm2 | kitty
    />
  </TerminalInfoProvider>
);
```

**Auto-detected protocols in order**: Kitty > iTerm2 > Sixel > Half-block > Braille > ASCII fallback

#### 1.2.11 `ink-big-text` — Large ASCII Text
**GitHub**: `sindresorhus/ink-big-text`  
**npm**: `ink-big-text` ~73K/week  
**Powered by**: `cfonts`

```tsx
import BigText from 'ink-big-text';

<BigText
  text="HELLO"
  font="block"              // block, shade, simple, 3d, chrome, etc.
  colors={['red', 'green']}  // multi-color
  background="transparent"
  space={true}
/>
```

#### 1.2.12 `ink-gradient` — Color Gradients
**GitHub**: `sindresorhus/ink-gradient`  
**npm**: `ink-gradient` ~321K/week

```tsx
import Gradient from 'ink-gradient';
import BigText from 'ink-big-text';

<Gradient name="rainbow">
  <BigText text="FLASHY" />
</Gradient>

<Gradient colors={['#FF6B6B', '#4ECDC4']}>
  <Text>Custom gradient text</Text>
</Gradient>
```

**Built-in gradients**: `rainbow`, `pastel`, `teen`, `mind`, `morning`, `vice`, `passion`

#### 1.2.13 `ink-divider` — Section Dividers
**GitHub**: `JureSotosek/ink-divider`  
**npm**: `ink-divider` ~4.5K/week

```tsx
import Divider from 'ink-divider';

<Divider title="Section Title" dividerColor="gray" titleColor="white" padding={1} />
// Output: ─────────── Section Title ───────────
```

#### 1.2.14 `ink-tab` — Tab Navigation
**GitHub**: `jdeniau/ink-tab`  
**npm**: `ink-tab` ~12.6K/week

```tsx
import { Tabs, Tab } from 'ink-tab';

<Tabs onChange={(activeTab) => setTab(activeTab)}>
  <Tab name="tab1">Tab 1</Tab>
  <Tab name="tab2">Tab 2</Tab>
  <Tab name="tab3">Tab 3</Tab>
</Tabs>
```

#### 1.2.15 `ink-ascii` — Figlet ASCII Art
**GitHub**: `hexrcs/ink-ascii`

```tsx
import Ascii from 'ink-ascii';

<Ascii
  text="AI CLI"
  font="Standard"           // figlet fonts
  horizontalLayout="default"
  verticalLayout="default"
/>
```

**Font sources**: figlet.js (patorjk) — 100+ figlet fonts

#### 1.2.16 `ink-quicksearch-input` — Fuzzy Search Input
**GitHub**: `Eximchain/ink-quicksearch-input`  
**Ink 4**: `@inkkit/ink-quicksearch-input`

```tsx
import QuickSearchInput from 'ink-quicksearch-input';

const items = [
  { label: 'JavaScript', value: 'js' },
  { label: 'TypeScript', value: 'ts' },
  { label: 'Python', value: 'py' },
];

<QuickSearchInput
  items={items}
  onSelect={(item) => console.log(item.value)}
  focus={true}
  caseSensitive={false}
  limit={5}
  forceMatchingQuery={false}
  clearQueryChars="\x1b"    // ESC to clear
  initialSelectionIndex={0}
  label="Search language: "
/>
```

**Behavior**: Type characters to filter items in real-time. Arrow keys to highlight. Enter to confirm.

#### 1.2.17 `@inkjs/ui` — Official UI Kit
**GitHub**: `vadimdemedes/ink-ui` — 2K+ stars  
**npm**: `@inkjs/ui` ~257K/week — v2.0.0 (May 2024)

```tsx
import {
  TextInput,
  Select,
  MultiSelect,
  ConfirmInput,
  Spinner,
  ProgressBar,
  StatusMessage,
  Key,
  KeyMap,
} from '@inkjs/ui';
```

**Components available**:
| Component | Description |
|-----------|-------------|
| `TextInput` | Text input with placeholder, mask, onSubmit |
| `Select` | Arrow-key select from options |
| `MultiSelect` | Checkbox multi-select |
| `ConfirmInput` | Y/n confirmation |
| `Spinner` | Animated loading spinner |
| `ProgressBar` | 0-100 progress bar |
| `StatusMessage` | Info/success/error status messages |
| `Key` | Single keyboard shortcut display |
| `KeyMap` | Keyboard shortcut legend |

**Modern API pattern**:
```tsx
import { Select } from '@inkjs/ui';

function App() {
  return (
    <Select
      options={[
        { label: 'Option A', value: 'a' },
        { label: 'Option B', value: 'b' },
      ]}
      onChange={(value) => console.log(value)}
    />
  );
}
```

### 1.3 Ink Web — Browser Runtime
**Website**: https://www.ink-web.dev  
**GitHub**: `cjroth/ink-web`

**Purpose**: Run Ink CLI apps in browser via XTerm.js (shadcn registry pattern)

**Installation**:
```bash
npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/[component].json
```

**Required imports**:
```css
import 'ink-web/css';
import '@xterm/xterm/css/xterm.css';
```

**Key concept**: At build time, `ink` is aliased to `ink-web`. Components like `InkXterm` provide the terminal emulator wrapper.

**Available browser components**: Table, Ascii, Spinner, Key, and more.

---

### 1.4 Best Practices for Ink CLI

#### Layout Pattern
```tsx
// Master-detail layout
<Box flexDirection="column" height="100%">
  {/* Header */}
  <Box borderStyle="single" padding={1}>
    <Text bold>My CLI App v1.0</Text>
  </Box>

  {/* Main content area — takes remaining space */}
  <Box flexGrow={1} flexDirection="column" padding={1}>
    <Text>Content here...</Text>
  </Box>

  {/* Footer/status bar */}
  <Box borderStyle="single" padding={1}>
    <Text dim>Status: Ready</Text>
  </Box>
</Box>
```

#### Error Boundary
```tsx
class ErrorBoundary extends React.Component<{children: ReactNode}, {error: Error | null}> {
  state = { error: null };

  static getDerivedStateFromError(error: Error) {
    return { error };
  }

  render() {
    if (this.state.error) {
      return (
        <Box flexDirection="column" padding={1}>
          <Text color="red" bold>Error:</Text>
          <Text>{this.state.error.message}</Text>
          <Text dim>{this.state.error.stack}</Text>
        </Box>
      );
    }
    return this.props.children;
  }
}
```

#### Full App Template
```tsx
#!/usr/bin/env node
import React, { useState, useCallback } from 'react';
import { render, Text, Box, useInput, useApp } from 'ink';
import TextInput from 'ink-text-input';
import SelectInput from 'ink-select-input';
import Spinner from 'ink-spinner';

function App() {
  const { exit } = useApp();
  const [screen, setScreen] = useState<'input' | 'loading' | 'result'>('input');
  const [query, setQuery] = useState('');
  const [result, setResult] = useState('');

  useInput((input, key) => {
    if (key.escape) exit();
    if (key.ctrl && input === 'c') exit();
  });

  const handleSubmit = useCallback(async (value: string) => {
    setQuery(value);
    setScreen('loading');
    // Simulate async work
    await new Promise(r => setTimeout(r, 1000));
    setResult(`You said: ${value}`);
    setScreen('result');
  }, []);

  return (
    <Box flexDirection="column" padding={1}>
      <Box marginBottom={1}>
        <Text bold color="blue">Ultimate CLI Demo</Text>
      </Box>

      {screen === 'input' && (
        <Box>
          <Text>Enter text: </Text>
          <TextInput
            value={query}
            onChange={setQuery}
            onSubmit={handleSubmit}
            placeholder="Type something and press Enter..."
          />
        </Box>
      )}

      {screen === 'loading' && (
        <Box>
          <Text color="yellow"><Spinner type="dots" /> </Text>
          <Text>Processing...</Text>
        </Box>
      )}

      {screen === 'result' && (
        <Box flexDirection="column">
          <Text color="green" bold>✓ Done!</Text>
          <Text>{result}</Text>
          <Box marginTop={1}>
            <SelectInput
              items={[
                { label: 'Try again', value: 'again' },
                { label: 'Exit', value: 'exit' },
              ]}
              onSelect={(item) => {
                if (item.value === 'again') setScreen('input');
                else exit();
              }}
            />
          </Box>
        </Box>
      )}
    </Box>
  );
}

const { waitUntilExit } = render(<App />);
await waitUntilExit;
```

---

## 2. CLI ARGUMENT PARSING

### 2.1 Commander.js
**GitHub**: `tj/commander.js` — 26K+ stars  
**npm**: `commander` ~50M/week — zero dependencies, TypeScript-first v8+

```typescript
import { Command } from 'commander';

const program = new Command();

program
  .name('my-cli')
  .description('CLI description')
  .version('1.0.0');

// Basic option
program
  .option('-d, --debug', 'Enable debug mode')
  .option('-o, --output <file>', 'Output file')
  .option('-p, --port <number>', 'Port number', parseInt, 3000);

// Command with arguments
program
  .command('serve [port]')
  .description('Start the server')
  .option('-v, --verbose', 'Verbose output')
  .action((port, options) => {
    console.log(`Serving on port ${port || 3000}`);
    if (options.verbose) console.log('Verbose mode on');
  });

// Required arguments
program
  .command('deploy <environment>')
  .argument('[path]', 'Path to deploy')
  .requiredOption('-k, --api-key <key>', 'API key (required)')
  .option('-f, --force', 'Force deploy')
  .action((environment, path, options) => {
    console.log(`Deploying to ${environment}`);
  });

program.parse();
```

**TypeScript usage** (v8+ auto-inferrence):
```typescript
const program = new Command();
program
  .option('--port <number>', 'port', parseInt)
  .action((opts) => {
    // opts.port is typed as number
  });
```

### 2.2 Yargs
**GitHub**: `yargs/yargs` — 11K+ stars  
**npm**: `yargs` ~30M/week — auto-help, validation, middleware, shell completions

```typescript
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

yargs(hideBin(process.argv))
  .scriptName('my-cli')
  .usage('$0 <cmd> [args]')
  .command(
    'deploy <environment>',
    'Deploy to environment',
    (yargs) => {
      yargs
        .positional('environment', {
          type: 'string',
          choices: ['staging', 'production'] as const,
          describe: 'Target environment',
        })
        .option('force', {
          alias: 'f',
          type: 'boolean',
          description: 'Force deploy',
        })
        .option('timeout', {
          type: 'number',
          default: 60,
          description: 'Timeout in seconds',
        });
    },
    (argv) => {
      console.log(`Deploying to ${argv.environment}`);
      if (argv.force) console.log('Force mode');
    }
  )
  .commandDir('./commands') // auto-load command modules from dir
  .demandCommand(1, 'Please specify a command')
  .recommendCommands()
  .completion()            // shell completion
  .middleware([(argv) => { /* pre-processing */ }])
  .help()
  .parse();
```

### 2.3 Comparison

| Feature | Commander | Yargs |
|---------|-----------|-------|
| Weekly downloads | ~50M | ~30M |
| Dependencies | 0 | 3 (small) |
| Auto help | Manual | Automatic |
| Subcommands | Good | Excellent |
| Shell completion | Manual | Built-in |
| Middleware | None | Built-in |
| Validation | Manual | Built-in (choices, required) |
| Bundle impact | Minimal | ~7.8MB node_modules |
| Best for | Published npm packages | Complex internal CLIs |

---

## 3. TERMINAL UI LIBRARIES (Non-Ink)

### 3.1 Blessed (Curses-Like Widget System)
**GitHub**: `chjj/blessed` — 11K+ stars  
**npm**: `blessed`

```typescript
import blessed from 'blessed';

const screen = blessed.screen({
  smartCSR: true,
  title: 'My App',
});

const box = blessed.box({
  top: 'center',
  left: 'center',
  width: '50%',
  height: '50%',
  content: 'Hello {bold}world{/bold}!',
  tags: true,
  border: { type: 'line' },
  style: {
    fg: 'white',
    bg: 'magenta',
    border: { fg: '#f0f0f0' },
  },
});

screen.append(box);
screen.render();

screen.key(['q', 'C-c'], () => process.exit(0));
```

**Successor**: `unblessed` — TypeScript rewrite, 98.5% test coverage, browser support, React renderer.

### 3.2 Blessed-Contrib (Dashboard Widgets)
**GitHub**: `yaronn/blessed-contrib`

```typescript
import contrib from 'blessed-contrib';

const screen = blessed.screen();
const grid = new contrib.grid({ rows: 12, cols: 12, screen });

const line = grid.set(0, 0, 6, 6, contrib.line, {
  style: { line: 'yellow', text: 'green', baseline: 'black' },
  label: 'Active Users',
});

line.setData([{ x: ['Mon','Tue','Wed'], y: [100, 200, 150] }]);
```

**Widgets**: line chart, bar chart, table, gauge, donut, tree, map, sparkline, log, picture.

### 3.3 Chalk (Terminal Styling)
**GitHub**: `chalk/chalk` — 25M+/week

```typescript
import chalk from 'chalk';

// Basic
console.log(chalk.blue('Hello world'));
console.log(chalk.red.bold('Error!'));

// Chaining & nesting
console.log(chalk.blue.bgRed.bold('Styled!'));

// Template literals
console.log(chalk`
  {bold Status:} {green Success}
  {bold Error:} {red ${errorMessage}}
`);

// 256 colors & Truecolor
chalk.hex('#DEADED').bold('Bold hex!');
chalk.rgb(15, 100, 204).inverse('Hello!');

// Level detection
chalk.level;        // 0: none, 1: 16, 2: 256, 3: 16M
chalk.supportsColor; // boolean
```

### 3.4 Gradient-String
**GitHub**: `bokub/gradient-string`

```typescript
import gradient from 'gradient-string';

// Built-in gradients
console.log(gradient.rainbow('Rainbow text'));
console.log(gradient.pastel('Pastel text'));

// Custom gradients
const custom = gradient(['#FF6B6B', '#4ECDC4', '#45B7D1']);
console.log(custom('Custom gradient'));

// Multi-line (vertically aligned)
console.log(gradient.rainbow.multiline('Line 1\nLine 2\nLine 3'));

// Advanced options
gradient(['red', 'blue'], { interpolation: 'hsv' })('HSV gradient');
gradient(['red', 'blue'], { interpolation: 'rgb', hsvSpin: 'long' })('Long HSV');
```

### 3.5 CLI-Table3
**npm**: `cli-table3`

```typescript
import Table from 'cli-table3';

// Basic table
const table = new Table({
  head: ['Name', 'Role', 'Status'],
  colWidths: [20, 15, 10],
});

table.push(['Alice', 'Engineer', 'Active']);
table.push(['Bob', 'Designer', 'Away']);

console.log(table.toString());

// Styled table
const styled = new Table({
  chars: {
    'top': '═', 'top-mid': '╤', 'top-left': '╔', 'top-right': '╗',
    'bottom': '═', 'bottom-mid': '╧', 'bottom-left': '╚', 'bottom-right': '╝',
    'left': '║', 'left-mid': '╟', 'mid': '─', 'mid-mid': '┼',
    'right': '║', 'right-mid': '╢', 'middle': '│',
  },
  style: { 'padding-left': 1, 'padding-right': 1 },
});

// Cell spanning
const spanTable = new Table();
spanTable.push(
  { content: 'Cross-column header', colSpan: 2, hAlign: 'center' },
  ['Col 1', 'Col 2']
);
```

---

## 4. INTERACTIVE PROMPT LIBRARIES

### 4.1 Inquirer
**GitHub**: `SBoudrias/Inquirer.js`  
**npm**: `@inquirer/prompts` ~28.5M/week

```typescript
import { input, select, checkbox, confirm, search, password } from '@inquirer/prompts';

const answer = await input({
  message: 'Enter your name',
  default: 'User',
  validate: (val) => val.length > 0 || 'Name required',
});

const color = await select({
  message: 'Pick a color',
  choices: [
    { name: 'Red', value: 'red', description: 'The color of passion' },
    { name: 'Blue', value: 'blue' },
    { value: 'other', name: 'Other' },
  ],
});

const toppings = await checkbox({
  message: 'Choose toppings',
  choices: [
    { name: 'Cheese', value: 'cheese', checked: true },
    { name: 'Pepperoni', value: 'pepperoni' },
  ],
});

const confirmed = await confirm({ message: 'Continue?' });
const pwd = await password({ message: 'Enter password', mask: true });
```

### 4.2 Enquirer
**GitHub**: `enquirer/enquirer`  
**npm**: `enquirer` ~26.7M/week

```typescript
import Enquirer from 'enquirer';
const enquirer = new Enquirer();

// Or use the prompt function
import { prompt } from 'enquirer';

const response = await prompt({
  type: 'autocomplete',
  name: 'framework',
  message: 'Choose a framework',
  choices: ['React', 'Vue', 'Svelte', 'Angular'],
  limit: 5,
});

// Custom prompt types: AutoComplete, BasicAuth, Confirm, Form, Input,
// Invisible, List, MultiSelect, Numeral, Password, Quiz, Survey,
// Scale, Select, Sort, Snippet, Toggle
```

**Built-in prompt types (17 total)**:
| Type | Description |
|------|-------------|
| AutoComplete | Filterable select with suggestions |
| BasicAuth | Username/password form |
| Confirm | Yes/no |
| Form | Multi-field form |
| Input | Text input |
| Invisible | Hidden input (no echo) |
| List | Comma-separated list |
| MultiSelect | Checkbox choices |
| Numeral | Number input |
| Password | Masked input |
| Quiz | Quiz with scoring |
| Scale | Rating scale |
| Select | Single selection |
| Snippet | Code template |
| Sort | Reorderable list |
| Survey | Multiple questions |
| Toggle | On/off toggle |

### 4.3 Prompts
**GitHub**: `terkelg/prompts`  
**npm**: `prompts` ~45.2M/week — lightweight, 2 dependencies

```typescript
import prompts from 'prompts';

const response = await prompts({
  type: 'text',
  name: 'name',
  message: 'What is your name?',
  initial: 'User',
  validate: (val: string) => val.length > 2 || 'Name too short',
});

// Multi-prompt form
const responses = await prompts([
  { type: 'text', name: 'username', message: 'Username' },
  { type: 'password', name: 'password', message: 'Password' },
  { type: 'select', name: 'role', message: 'Role',
    choices: [
      { title: 'Admin', value: 'admin' },
      { title: 'User', value: 'user' },
    ]
  },
]);

// Prompt types: text, password, invisible, number, confirm,
// list, toggle, select, multiselect, autocomplete, date
```

**Differences**:
| Library | Dependencies | API Style | Last Release | Best For |
|---------|-------------|-----------|-------------|----------|
| Inquirer | ~8 | Promise API | Active | Feature-complete forms |
| Enquirer | ~12 | Class + function | Active | Custom prompts |
| Prompts | 2 | Simple function | Active (2.4.2) | Lightweight usage |

---

## 5. FILE & SEARCH

### 5.1 Fast-Glob
**GitHub**: `mrmlnc/fast-glob`  
**npm**: `fast-glob`

```typescript
import fg from 'fast-glob';

// Async
const entries = await fg(['src/**/*.ts', '!src/**/*.test.ts']);
// => ['src/index.ts', 'src/utils/helper.ts', ...]

// Sync
const sync = fg.sync('**/*.json', { dot: true, cwd: './config' });

// Stream
const stream = fg.stream('**/*.ts');
for await (const entry of stream) {
  console.log(entry);
}

// Options
fg(['**/*'], {
  cwd: process.cwd(),
  dot: false,                // match dotfiles
  ignore: ['node_modules'],
  onlyFiles: true,
  onlyDirectories: false,
  absolute: false,
  caseSensitiveMatch: true,
  baseNameMatch: false,      // match basename against pattern
  braceExpansion: true,
  extglob: true,
  globstar: true,
  followSymbolicLinks: true,
  unique: true,
  markDirectories: false,
  objectMode: false,         // return Dirent objects
  stats: false,              // return fs.Stats
});
```

### 5.2 Globby
**GitHub**: `sindresorhus/globby`  
**npm**: `globby`

```typescript
import { globby, globbyStream } from 'globby';

// Advanced globbing with .gitignore support
const paths = await globby(['**/*.ts', '!**/*.spec.ts'], {
  gitignore: true,           // respect .gitignore
  ignore: ['node_modules'],
  cwd: '/project',
  dot: false,
  expandDirectories: true,   // automatically glob directories
});

// Negation-only patterns (auto-prepends **/*)
const configFiles = await globby(['!*.json', '!*.yaml'], { cwd: '/config' });

// Streaming
const stream = globbyStream('**/*');
for await (const path of stream) {
  processFile(path);
}

// URL cwd support
await globby('**/*', { cwd: new URL('file:///project') });
```

### 5.3 Ignore (.gitignore Parser)
**GitHub**: `kaelzhang/node-ignore`  
**npm**: `ignore` ~52.1M/week — follows gitignore spec 2.22.1

```typescript
import ignore from 'ignore';

const ig = ignore().add([
  'node_modules',
  '*.log',
  '!important.log',  // negate pattern
  '/dist',
  'build/',
]);

// Check files
ig.ignores('node_modules/package/index.js'); // true
ig.ignores('src/app.ts');                     // false

// Filter array
const files = [
  'node_modules/pkg/index.js',
  'src/index.ts',
  'dist/bundle.js',
  'important.log',
];
const filtered = ig.filter(files); // ['src/index.ts', 'important.log']

// Advanced: add with markers (e.g., line numbers from .gitignore file)
ig.add({ pattern: '*.test.ts', mark: '.gitignore:5' });

// Pattern support:
// * — any non-slash chars
// ? — single non-slash char
// ** — any number of directories
// [seq] — character range
// ! — negation
// \ — escape special chars
```

### 5.4 Tree-Sitter (Incremental Parser)
**GitHub**: `tree-sitter/tree-sitter` — 26K+ stars  
**Website**: https://tree-sitter.github.io  
**Written in**: Rust (core), C (runtime library)

#### Key Properties
1. **General**: Can parse any programming language
2. **Fast**: Operates in ~1ms per keystroke
3. **Robust**: Produces valid AST even with syntax errors
4. **Dependency-free**: Pure C11 runtime library

```typescript
import Parser from 'tree-sitter';
import JavaScript from 'tree-sitter-javascript';

const parser = new Parser();
parser.setLanguage(JavaScript);

const tree = parser.parse('function hello() { return 42; }');

// Navigate AST
const rootNode = tree.rootNode;
console.log(rootNode.type); // 'program'

// Query syntax
const query = new Parser.Query(JavaScript, '(function_declaration name: (identifier) @fn.name)');
const matches = query.matches(tree.rootNode);
for (const match of matches) {
  const nameNode = match.captures.find(c => c.name === 'fn.name').node;
  console.log(nameNode.text); // 'hello'
}

// Incremental parsing (edit a tree)
tree.edit({
  startIndex: 0,
  oldEndIndex: 0,
  newEndIndex: 5,
  startPosition: { row: 0, column: 0 },
  oldEndPosition: { row: 0, column: 0 },
  newEndPosition: { row: 0, column: 5 },
});
const newTree = parser.parse('const ', tree);
```

**Available language parsers**: 250+ languages (JavaScript, TypeScript, Python, Rust, Go, Java, C, C++, Ruby, PHP, Swift, Kotlin, HTML, CSS, JSON, YAML, Markdown, SQL, Bash, etc.)

**Used by**: GitHub code search, Neovim, Zed, Helix, Shopify, Discord

---

## 6. CODE EDITING & AST MANIPULATION

### 6.1 Babel — JavaScript/TypeScript Compiler
**GitHub**: `babel/babel` — 44K stars  
**Website**: https://babeljs.io

**Architecture**: input string → `@babel/parser` → AST → plugins (via `@babel/traverse`) → `@babel/generator` → output string

```typescript
import { parse } from '@babel/parser';
import traverse from '@babel/traverse';
import generate from '@babel/generator';
import * as t from '@babel/types';

// Parse
const ast = parse('const x = 42;', {
  sourceType: 'module',
  plugins: ['typescript', 'jsx'],
});

// Traverse & transform
traverse(ast, {
  VariableDeclaration(path) {
    if (path.node.kind === 'const') {
      path.node.kind = 'let';
    }
  },
  ArrowFunctionExpression(path) {
    // Replace arrow function with regular function
    const func = t.functionDeclaration(
      t.identifier('myFunc'),
      path.node.params,
      path.node.body,
      false, // generator
      path.node.async,
    );
    path.replaceWith(func);
  },
});

// Generate
const output = generate(ast, { retainLines: true }, 'const x = 42;');

// Code frame (error highlighting)
import { codeFrameColumns } from '@babel/code-frame';
const frame = codeFrameColumns('const x = 42;', {
  start: { line: 1, column: 7 },
}, { highlightCode: true, message: 'Error here' });
```

**Core packages**:
| Package | Purpose |
|---------|---------|
| `@babel/core` | Full compiler |
| `@babel/parser` | JS/TS/JSX parser |
| `@babel/traverse` | AST traversal |
| `@babel/generator` | AST to code |
| `@babel/types` | Type definitions & builders |
| `@babel/template` | Build AST from template strings |
| `@babel/code-frame` | Pretty error frames |
| `@babel/cli` | CLI tool |
| `@babel/register` | Require hook |
| `@babel/runtime` | Runtime helpers |

### 6.2 TS-Morph — TypeScript Compiler API Wrapper
**GitHub**: `dsherret/ts-morph`  
**Website**: https://ts-morph.com  
**npm**: `ts-morph`

```typescript
import { Project, SyntaxKind, Node, ScriptKind } from 'ts-morph';

// Create project
const project = new Project({
  tsConfigFilePath: 'tsconfig.json',
  // Or manually:
  compilerOptions: {
    target: ScriptTarget.ES2022,
    module: ModuleKind.ESNext,
    strict: true,
  },
});

// Add source files
const sourceFile = project.addSourceFileAtPath('src/index.ts');
// Or create in memory
const newFile = project.createSourceFile('src/new.ts', 'export const x = 42;');

// Navigation
sourceFile.getClasses().forEach(cls => {
  console.log(cls.getName());             // class name
  console.log(cls.getExtends()?.getText()); // extends clause
  cls.getProperties().forEach(prop => {
    console.log(prop.getName(), prop.getType().getText());
  });
  cls.getMethods().forEach(method => {
    method.getParameters().forEach(param => {
      console.log(param.getName(), param.getType().getText());
    });
  });
});

// Code manipulation
sourceFile.addImportDeclaration({
  moduleSpecifier: './utils',
  namedImports: ['helper'],
});

const func = sourceFile.addFunction({
  name: 'myFunc',
  parameters: [{ name: 'input', type: 'string' }],
  returnType: 'string',
  statements: [`return input.toUpperCase();`],
});

// Find and replace
const declarations = sourceFile.getVariableDeclarations(d => d.getName() === 'x');
declarations.forEach(d => d.rename('y'));

// Type checking
sourceFile.getFunctions().forEach(fn => {
  const returnType = fn.getReturnType();
  console.log(returnType.getText());    // string | null
  console.log(returnType.isNullable()); // true
});

// Save changes
project.saveSync();
// Or get text without saving:
const newContent = sourceFile.getFullText();
```

**Comparison**:
| Library | Use Case | Pros | Cons |
|---------|----------|------|------|
| Babel | Transform code | Broad support, many plugins | Loses formatting |
| TS-Morph | TypeScript manipulation | Full TS type awareness, preserves formatting | Heavier, TS-only |
| JSCodeshift | Codemods | Preserves style, batch mode | Recast-based, limited type info |

### 6.3 JSCodeshift — Codemod Toolkit
**GitHub**: `facebook/jscodeshift` — 10K stars  
**npm**: `jscodeshift`

```typescript
// Transform module format:
export default function transformer(fileInfo: FileInfo, api: API, options: Options) {
  const j = api.jscodeshift;

  return j(fileInfo.source)
    .find(j.Identifier)
    .filter(path => path.node.name === 'oldName')
    .forEach(path => {
      j(path).replaceWith(j.identifier('newName'));
    })
    .toSource();
}

// Collection API — jQuery-like traversal
j(source)
  .find(j.ArrowFunctionExpression)
  .filter(path => path.node.async)
  .forEach(path => {
    // Transform each match
  });

// Common operations
j(source)
  .findVariableDeclarators('foo')
  .renameTo('bar');

j(source)
  .find(j.ImportDeclaration)
  .filter(path => path.node.source.value === 'old-package')
  .forEach(path => {
    path.node.source.value = 'new-package';
  });

// Insert code
j(source)
  .find(j.Program)
  .forEach(path => {
    path.node.body.push(j.importDeclaration(
      [j.importDefaultSpecifier(j.identifier('React'))],
      j.literal('react')
    ));
  });

// Parser selection (export in module)
export const parser = 'tsx'; // or 'babel', 'flow', 'ts'
```

**How it works**: Uses `recast` under the hood — preserves original code style (whitespace, quotes, etc.) as much as possible.

**Running codemods**:
```bash
npx jscodeshift -t transform.js src/
npx jscodeshift -t transform.js --parser=tsx src/
```

---

## 7. GIT INTEGRATION

### 7.1 Isomorphic-Git (Pure JS Git)
**GitHub**: `isomorphic-git/isomorphic-git`  
**Website**: https://isomorphic-git.org  
**npm**: `isomorphic-git`

```typescript
import git from 'isomorphic-git';
import http from 'isomorphic-git/http/node';
import fs from 'fs';

const dir = '/path/to/repo';

// Clone
await git.clone({ fs, http, dir, url: 'https://github.com/user/repo.git' });

// Status
const status = await git.status({ fs, dir, filepath: 'README.md' });
// 'ignored' | 'unmodified' | '*modified' | '*deleted' | '*added' | 'absent' | 'modified' | 'deleted' | 'added'

// Add
await git.add({ fs, dir, filepath: 'README.md' });

// Commit
const sha = await git.commit({
  fs, dir,
  message: 'feat: add new feature',
  author: {
    name: 'Author',
    email: 'author@example.com',
  },
});

// Log
const commits = await git.log({ fs, dir, depth: 10 });
// Returns array of { oid, message, tree, parent, author, committer }

// Push
await git.push({
  fs, http, dir,
  remote: 'origin',
  ref: 'main',
  token: 'ghp_xxx',       // or onNote on Node 18+
});

// List branches
const branches = await git.listBranches({ fs, dir });
const tags = await git.listTags({ fs, dir });

// Diff
const diff = await git.diff({
  fs, dir, tree1: 'HEAD', tree2: 'WORKDIR', filepath: 'index.ts',
});

// Checkout
await git.checkout({ fs, dir, ref: 'main' });

// Raw: lower-level operations
await git.readBlob({ fs, dir, oid: 'abc123' });
await git.readTree({ fs, dir, oid: 'abc123' });
await git.resolveRef({ fs, dir, ref: 'HEAD' });
```

**Browser usage** (with LightningFS):
```typescript
import git from 'isomorphic-git';
import http from 'isomorphic-git/http/web';
import LightningFS from '@isomorphic-git/lightning-fs';

const fs = new LightningFS('myfs');
await git.clone({ fs, http, dir: '/repo', url, corsProxy: 'https://cors.isomorphic-git.org' });
```

### 7.2 Simple-Git (Native Git Wrapper)
**GitHub**: `simple-git-js/simple-git`  
**npm**: `simple-git`

```typescript
import { simpleGit, SimpleGit, CleanOptions } from 'simple-git';

const git: SimpleGit = simpleGit('/repo/path');

// Status
const status = await git.status();
console.log(status.modified, status.created, status.deleted);

// Log
const log = await git.log({ maxCount: 10 });
console.log(log.latest?.hash, log.latest?.message);

// Diff
const diff = await git.diffSummary();
const fileDiff = await git.diff(['--', 'src/index.ts']);

// Branch operations
await git.branch(['-m', 'master', 'main']);
await git.checkout('feature-branch');
await git.branchLocal();

// Add & commit
await git.add('.');
await git.commit('feat: update', {
  '--author': '"Name <email>"',
});

// Push & pull
await git.push('origin', 'main');
await git.pull('origin', 'main', { '--rebase': null });

// Tags
await git.addTag('v1.0.0');
const tags = await git.tags();

// Raw command (when no wrapper exists)
const raw = await git.raw(['log', '--oneline', '-5']);

// Checkout
await git.checkout('main');
await git.pull();
await git.checkoutLocalBranch('new-feature');

// Clean
await git.clean(CleanOptions.FORCE + CleanOptions.RECURSIVE);

// Advanced
const stash = await git.stash();
await git.bisect();
await git.submoduleUpdate();
```

**Comparison**:
| Library | Platform | Speed | Features | Dependency |
|---------|----------|-------|----------|------------|
| isomorphic-git | Node + browser | Moderate | Complete git, pure JS | None (native) |
| simple-git | Node only | Fast (uses git binary) | Very complete | git binary |

---

## 8. AI CLI FRAMEWORKS

### 8.1 LangChain.js — LLM Application Framework
**Website**: https://docs.langchain.com/oss/javascript/langchain/overview  
**GitHub**: `langchain-ai/langchainjs`  
**npm**: `@langchain/core`, `langchain`

```typescript
import { ChatOpenAI } from '@langchain/openai';
import { createAgent } from './agent.js';
import { tool } from '@langchain/core/tools';
import { MemorySaver } from '@langchain/langgraph';

// Define tools with Zod schemas
const searchTool = tool(
  async ({ query }: { query: string }) => {
    const result = await search(query);
    return JSON.stringify(result);
  },
  {
    name: 'web_search',
    description: 'Search the web',
    schema: z.object({ query: z.string() }),
  }
);

// Create agent
const agent = await createAgent({
  llm: new ChatOpenAI({ model: 'gpt-4o', temperature: 0 }),
  tools: [searchTool],
  systemPrompt: 'You are a helpful CLI assistant.',
  memory: new MemorySaver(),  // persistence
});

// Run agent
const result = await agent.invoke({ messages: [{ role: 'user', content: query }] });
console.log(result.messages[result.messages.length - 1].content);
```

**Key patterns**:
- **Tools**: Defined with Zod schemas, can include runtime context
- **Memory**: MemorySaver for conversation state between interactions
- **Tracing**: LangSmith observability integration
- **Deep Agents**: Built-in planning, file system, subagent capabilities

### 8.2 OpenAI Agents SDK
**GitHub (Python)**: `openai/openai-agents-python`  
**GitHub (TS)**: `openai/openai-agents-js`  
**Website**: https://openai.github.io/openai-agents-python/

```typescript
import { Agent, Runner, type AgentHooks } from 'openai-agents';

const agent = new Agent({
  name: 'Research Assistant',
  instructions: 'You are a research assistant.',
  model: 'gpt-4o',
  tools: [webSearchTool, fileReaderTool],
});

// Multi-agent handoffs
const editor = new Agent({
  name: 'Editor',
  instructions: 'Review and polish content.',
});

agent.addHandoff(editor);

// Run
const result = await Runner.run(agent, query);

// Guardrails
runner.withGuardrail({
  check: (message) => {
    if (containsPII(message)) throw new Error('PII detected');
  },
});
```

**Features**: Provider-agnostic (OpenAI + others), agent handoffs, guardrails, tracing.

### 8.3 LlamaIndex — Data Framework for LLMs
**Website**: https://www.llamaindex.ai  
**GitHub**: `run-llama/llama_index` — 50.8K stars

```typescript
import {
  VectorStoreIndex,
  Document,
  SimpleDirectoryReader,
  OpenAIEmbedding,
} from 'llamaindex';

// Ingestion
const documents = await new SimpleDirectoryReader()
  .loadData({ directoryPath: './data' });

const index = await VectorStoreIndex.fromDocuments(documents, {
  embedModel: new OpenAIEmbedding(),
});

// Query
const queryEngine = index.asQueryEngine();
const response = await queryEngine.query({
  query: 'What is this about?',
});

// Agentic RAG
const agent = await createAgent({
  tools: [index.asQueryTool()],
  llm: new OpenAI({ model: 'gpt-4o' }),
});
```

**Products**: LlamaParse (agentic OCR), LlamaAgents, LlamaCloud, Workflows.

---

## 9. MCP PROTOCOL (Model Context Protocol)

**Website**: https://modelcontextprotocol.io  
**GitHub**: `modelcontextprotocol/specification`

### Architecture
```
┌──────────────┐      JSON-RPC 2.0      ┌──────────────┐
│    Host      │ ◄─────────────────────► │   Server     │
│ (LLM App)    │    Stateful Session     │ (Capability) │
└──────┬───────┘                         └──────┬───────┘
       │                                         │
       │   ┌──────────────┐                      │
       └──►│   Client     │──────────────────────┘
           │  (Connector) │
           └──────────────┘
```

### Server Capabilities
| Capability | Description | Examples |
|------------|-------------|----------|
| **Resources** | Read-only data exposure | Files, DB schemas, API docs |
| **Prompts** | Templated interaction patterns | Workflows, multi-step tasks |
| **Tools** | Executable functions | Search, calculate, transform |

### Client Capabilities
| Capability | Description |
|------------|-------------|
| **Sampling** | Server-initiated generation requests |
| **Roots** | Filesystem boundary definition |
| **Elicitation** | Server-initiated user info requests |

### Transport
- **stdio**: Local process communication
- **HTTP + SSE**: Remote server communication
- **WebSocket**: Bidirectional streaming

### Implementation Pattern
```typescript
import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';

const server = new Server(
  { name: 'my-server', version: '1.0.0' },
  { capabilities: { tools: {}, resources: {}, prompts: {} } }
);

// Tool registration
server.setRequestHandler('tools/call', async (request) => {
  const { name, arguments: args } = request.params;
  switch (name) {
    case 'search':
      return { content: [{ type: 'text', text: await search(args.query) }] };
    case 'read_file':
      return { content: [{ type: 'resource', resource: { text: fs.readFileSync(args.path, 'utf-8') } }] };
    default:
      throw new Error(`Unknown tool: ${name}`);
  }
});

await server.connect(new StdioServerTransport());
```

---

## 10. VECTOR SEARCH & EMBEDDING DATABASES

### 10.1 Chroma — Open-Source AI Search
**Website**: https://www.trychroma.com  
**GitHub**: `chroma-core/chroma` — 26K+ stars, Apache 2.0  
**Downloads**: 11M+ per month

```typescript
import { ChromaClient } from 'chromadb';

const client = new ChromaClient({ path: 'http://localhost:8000' });

// Create collection
const collection = await client.createCollection({
  name: 'my_docs',
  metadata: { 'hnsw:space': 'cosine' },
});

// Add documents with embeddings
await collection.add({
  ids: ['doc1', 'doc2'],
  embeddings: [[0.1, 0.2, ...], [0.3, 0.4, ...]],  // or let Chroma auto-embed
  metadatas: [{ source: 'wiki' }, { source: 'blog' }],
  documents: ['Document text 1', 'Document text 2'],
});

// Search — hybrid (dense + sparse + full-text + metadata)
const results = await collection.query({
  queryEmbeddings: [[0.1, 0.2, ...]],
  nResults: 10,
  where: { source: 'wiki' },      // metadata filter
  whereDocument: { $contains: 'AI' },  // full-text filter
});
```

**Search types**: Dense vector (cosine/l2/ip), sparse vector (BM25/SPLADE), full-text, regex, metadata filtering.  
**Architecture**: Built on object storage, serverless, automatic query-aware data tiering.

### 10.2 Qdrant — High-Performance Vector Search
**GitHub**: `qdrant/qdrant` — Written in Rust  
**Website**: https://qdrant.tech

```typescript
import { QdrantClient } from '@qdrant/js-client-rest';

const client = new QdrantClient({ url: 'http://localhost:6333' });

// Create collection
await client.createCollection('my_collection', {
  vectors: { size: 384, distance: 'Cosine' },
  hnsw_config: { m: 16, ef_construct: 100 },
});

// Add points
await client.upsert('my_collection', {
  points: [
    {
      id: 1,
      vector: [0.1, 0.2, ...],
      payload: { title: 'Doc 1', category: 'tech' },
    },
  ],
});

// Search
const results = await client.search('my_collection', {
  vector: [0.1, 0.2, ...],
  limit: 10,
  filter: {
    must: [{ key: 'category', match: { value: 'tech' } }],
  },
});
```

**Features**: HNSW indexing, horizontal scaling, quantization, multi-tenancy, REST + gRPC APIs.

### 10.3 LanceDB — Embedded Lakehouse for AI
**Website**: https://www.lancedb.com  
**GitHub**: `lancedb/lance` — Apache 2.0  
**Written in**: Rust, with Python/JS/Rust SDKs

```typescript
import * as lancedb from '@lancedb/lancedb';

const db = await lancedb.connect('./data/lancedb');

// Create table
const table = await db.createTable('vectors', [
  { id: 1, vector: [0.1, 0.2, ...], text: 'Hello world' },
  { id: 2, vector: [0.3, 0.4, ...], text: 'AI is great' },
]);

// Create vector index
await table.createIndex({ column: 'vector', type: 'ivf_pq', num_partitions: 256 });

// Search (hybrid — vector + full-text + SQL)
const results = await table
  .search([0.1, 0.2, ...])
  .where('text IS NOT NULL')
  .limit(10)
  .toArray();
```

**Key features**:
- **Format**: Lance — columnar lakehouse format (100x faster random access than Parquet)
- **Search**: Dense vector (IVF-PQ), full-text (BM25), SQL analytics
- **Data versioning**: Zero-copy, ACID transactions, time travel, tags, branches
- **Ecosystem**: DuckDB, Pandas, Polars, PyTorch, Spark, LangChain, LlamaIndex

### Vector DB Comparison

| Feature | Chroma | Qdrant | LanceDB |
|---------|--------|--------|---------|
| Search types | Vector + Sparse + Full-text + Regex | Vector + Filter | Vector + BM25 + SQL |
| Architecture | Serverless (object storage) | Server (gRPC/REST) | Embedded/lakehouse |
| Index | HNSW | HNSW | IVF-PQ |
| Language | Python/JS | Rust | Rust |
| Versioning | No | No | Yes (zero-copy) |
| Best for | Quick prototype, serverless | Production search | ML + data lakehouse |

---

## 11. ARCHITECTURAL PATTERNS FOR AI CLI AGENTS

### 11.1 Full-Stack AI CLI Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                        CLI App Layer                             │
│  ┌─────────────────┐  ┌──────────────┐  ┌──────────────────┐   │
│  │  Argument Parser │  │  Ink UI      │  │  Interactive     │   │
│  │  (Commander/     │  │  (React      │  │  Prompts         │   │
│  │   yargs)         │  │   Terminal)  │  │  (Inquirer/etc)  │   │
│  └────────┬────────┘  └──────┬───────┘  └───────┬──────────┘   │
└───────────┼──────────────────┼──────────────────┼──────────────┘
            │                  │                  │
┌───────────┼──────────────────┼──────────────────┼──────────────┐
│           ▼                  ▼                  ▼              │
│                     Agent Orchestration Layer                    │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │  LangChain / OpenAI Agents SDK / LlamaIndex              │   │
│  │  • Tool definitions (Zod schemas)                       │   │
│  │  • Multi-agent handoffs                                 │   │
│  │  • Guardrails & validation                              │   │
│  │  • Memory & persistence                                 │   │
│  └──────────────────────────┬───────────────────────────────┘   │
└─────────────────────────────┼───────────────────────────────────┘
                              │
┌─────────────────────────────┼───────────────────────────────────┐
│                             ▼                                    │
│                      Capability Layer                             │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────────┐   │
│  │ File Ops │ │ Code     │ │ Git      │ │ Vector Search    │   │
│  │ (fast-   │ │ Analysis │ │ (isogit/ │ │ (Chroma/Qdrant/  │   │
│  │ glob/    │ │ (tree-   │ │ simple-  │ │ LanceDB)         │   │
│  │ globby)  │ │ sitter/  │ │ git)     │ │                  │   │
│  └──────────┘ │ babel/   │ └──────────┘ │                  │   │
│               │ ts-morph/│              │                  │   │
│               │ jscode-  │              │                  │   │
│               │ shift)   │              │                  │   │
│               └──────────┘              └──────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### 11.2 MCP-Based Agent Architecture

```
┌────────────────────┐
│   MCP Host (LLM)   │
│  (Your CLI Agent)  │
└────┬───────────┬───┘
     │           │
┌────▼───┐ ┌────▼────────────┐
│ MCP    │ │ MCP             │
│ Client │ │ File/Code/Tool  │
└────┬───┘ │ Servers         │
     │     └─────────────────┘
     │
┌────▼─────────────────────┐
│ MCP Transport (stdio/SSE)│
└──────────────────────────┘
```

### 11.3 Recommended Stack for Different Scenarios

| Scenario | CLI Parser | UI | AI Framework | Vector DB | File Ops | Code Analysis | Git |
|----------|-----------|-----|-------------|-----------|----------|--------------|-----|
| Simple CLI tool | Commander | Ink + @inkjs/ui | — | — | fast-glob | — | simple-git |
| Complex AI CLI | Yargs | Ink + custom components | LangChain.js | LanceDB | globby | tree-sitter | isomorphic-git |
| AI Agent Platform | Commander | Ink + ink-web | OpenAI Agents SDK | Chroma | globby + ignore | ts-morph | isomorphic-git |
| Code Analysis Tool | Yargs | Blessed | — | — | fast-glob | tree-sitter + babel | simple-git |
| RAG Pipeline CLI | Commander | Ink minimal | LlamaIndex | Qdrant | globby | — | simple-git |

---

## 12. NPM ECOSYSTEM DOWNLOAD COMPARISON

| Package | Weekly Downloads | Category |
|---------|-----------------|----------|
| chalk | 25M+ | Terminal styling |
| commander | ~50M | CLI parser |
| yargs | ~30M | CLI parser |
| prompts | ~45M | Interactive prompts |
| @inquirer/prompts | ~28.5M | Interactive prompts |
| enquirer | ~26.7M | Interactive prompts |
| ignore | ~52M | .gitignore matching |
| fast-glob | ~30M+ | File globbing |
| globby | ~20M+ | User-friendly globbing |
| simple-git | ~10M+ | Git commands |
| isomorphic-git | ~1M+ | Pure JS git |
| ink | ~4.2M | React for CLI |
| @inkjs/ui | ~257K | Ink UI kit |
| ink-text-input | ~1.4M | Ink text input |
| ink-spinner | ~1.6M | Ink spinner |
| ink-gradient | ~321K | Ink gradient |
| ink-link | ~180K | Ink links |
| ink-big-text | ~73K | Ink big text |
| ink-picture | ~68K | Terminal images |
| ink-quicksearch-input | ~16K | Search input |

---

## 13. AWESOME CLI APPS ECOSYSTEM

> Reference: `references/awesome-cli-apps.md` | Source: toolleeo/awesome-cli-apps-in-a-csv (2230 tools, 81 categories)
>
> **Dedicated A-to-Z references**: `references/modern-unix-tools.md` | `references/ai-cli-tools.md` | `references/git-tui-tools.md` | `references/tui-file-managers.md`

When building CLIs, you should also be aware of the CLI *application* ecosystem — tools users install and run. The Awesome CLI Apps list catalogs 2230 CLI/TUI tools. Key categories for CLI developers:

### Modern Replacements for Unix Classics
→ **Full A-to-Z reference**: `references/modern-unix-tools.md` (15 tools)

| Classic | Modern Replacement | Language | Why Better |
|---------|-------------------|----------|------------|
| `find` | `fd` | Rust | Simpler syntax, sensible defaults, faster |
| `grep` | `ripgrep` | Rust | Faster, .gitignore-aware, better defaults |
| `sed` | `sd` | Rust | Intuitive find/replace, no cryptic syntax |
| `cut` | `choose` | Rust | Human-friendly field selection |
| `ls` | `eza` / `lsd` | Rust | Icons, colors, tree view |
| `cd` | `zoxide` | Rust | Fuzzy + frecency-based directory jumping |
| `top` | `btm` (bottom) | Rust | Graphical, cross-platform |
| `ps` | `procs` | Rust | Colorful columns, Docker support |
| `diff` | `delta` | Rust | Syntax highlighting, side-by-side |
| `jq` | `jaq` / `gojq` | Rust/Go | Faster, memory-safe alternatives |

### Essential CLI Tools for Developers
| Category | Must-Know Tools | Reference File |
|----------|----------------|----------------|
| Git TUI | Lazygit, GitUI, tig | `references/git-tui-tools.md` |
| GitHub CLI | gh, gh-dash | `references/git-tui-tools.md` |
| File Manager | yazi, lf, ranger, nnn, broot | `references/tui-file-managers.md` |
| Text Search | ripgrep, ugrep, ast-grep | `references/modern-unix-tools.md` |
| Data Processing | jq, yq, dasel, xsv, visidata | — |
| HTTP Client | httpie, curlie | — |
| Markdown | glow | — |
| Disk Usage | dua-cli, dust, ncdu | `references/modern-unix-tools.md` |
| AI Tools | ollama, aichat, fabric, Mods! | `references/ai-cli-tools.md` |
| Security | gitleaks, trivy | — |

### Key Trends
1. **Rust dominates** modern CLI tooling (fd, rg, bat, delta, sd, bottom, yazi)
2. **Go is strong** for infrastructure (gh, lazygit, k9s, terraform)
3. **AI is explosive** — 75+ AI CLI tools across 3 categories
4. **TUIs are trending** — interactive UIs replacing simple CLIs

---

## 14. ULTIMATE CLI — QUICK REFERENCE CARD

### Ink Essentials
```tsx
render(<App />)           // Bootstrap
useInput((i, k) => {})    // Keyboard
useFocus()                // Focus management
<Box>                     // Flex container
<Text>                    // Text (MANDATORY for text)
```

### File Operations
```typescript
fg('**/*.ts')                       // Fast glob
globby(['**/*', '!**/*.test.ts'])   // Smart glob
ignore().add('node_modules')        // .gitignore filter
```

### Code Analysis
```typescript
treeSitter.parse(language, code)    // Incremental parse
tsMorph.project.addSourceFile()     // TS AST manipulation
jscodeshift(source).find(...)       // Codemod
```

### AI Integration
```typescript
createAgent({ llm, tools })         // LangChain
Runner.run(agent, query)            // OpenAI Agents
index.asQueryEngine()               // LlamaIndex
```

### Vector Search
```typescript
chroma.collection.add({...})        // Chroma
qdrant.upsert('col', {...})         // Qdrant
lancedb.table.search(vec)           // LanceDB
```

### Git
```typescript
git.clone({ fs, http, ... })        // isomorphic-git
simpleGit().status()                // simple-git
```

---

## 15. AWESOME TUIS — TUI APPLICATION ECOSYSTEM

### 13 Technology Categories (500+ TUIs)

| # | Category | Examples | Reference |
|---|----------|----------|-----------|
| 1 | **Dashboards** | gotop, btop, gtop, zenith, glances, ytop, fastfetch | `references/awesome-tuis-ecosystem.md` |
| 2 | **Development** | lazydocker, lazysql, Lite XL, micro, neovim, helix | `references/awesome-tuis-ecosystem.md` |
| 3 | **Docker & K8s** | lazydocker, docker-tui, k9s, kui, kubenav | `references/awesome-tuis-ecosystem.md` |
| 4 | **Editors** | neovim, helix, micro, vis, nvim, emacs (terminal), kakoune | `references/awesome-tuis-ecosystem.md` |
| 5 | **File Managers** | yazi, ranger, lf, nnn, broot, fzf (related) | `references/file-managers.md` |
| 6 | **Games** | nethack, brogue, vitetris, myman, adventure | `references/awesome-tuis-ecosystem.md` |
| 7 | **Libraries (Dev)** | Bubble Tea, Ratatui, Textual, Rich, Ink, FTXUI | `references/awesome-tuis-libraries.md` |
| 8 | **Messaging** | slack-term, weechat, irssi, discordchatex-cli | `references/awesome-tuis-ecosystem.md` |
| 9 | **Miscellaneous** | fzf, zoxide, cheat, tldr, lazygit | See below |
| 10 | **Multimedia** | cmus, mpv (TUI), ncmpcpp, spotify-tui | `references/awesome-tuis-ecosystem.md` |
| 11 | **Productivity** | taskwarrior, calcurse, newsboat, tuir, jrnl | `references/awesome-tuis-ecosystem.md` |
| 12 | **Screensavers** | cmatrix, pipes.sh, terminal-screensaver | `references/awesome-tuis-ecosystem.md` |
| 13 | **Web** | browsh, links, carbon-now-cli | `references/awesome-tuis-ecosystem.md` |

### TUI Library Framework Popularity

| Framework | Language | Ecosystem | Key TUIs |
|-----------|----------|-----------|----------|
| **Bubble Tea** | Go | 40+ TUIs | glow, charm, gum, slides, wish |
| **Ratatui** | Rust | 30+ TUIs | zellij, bottom, yazi, gping, broot |
| **Textual** | Python | 15+ TUIs | (growing fast) |
| **Rich** | Python | 10+ TUIs | (power tool for TUIs) |
| **Ink** | JavaScript | 10+ TUIs | (React-based) |
| **FTXUI** | C++ | 10+ TUIs | (fast C++ TUIs) |
| **ncurses** | C | 200+ TUIs | (foundational, used by vim/htop etc.) |

### TUI Dev Libraries Reference (65+ across 8 languages)

See full reference: `references/awesome-tuis-libraries.md`

| Language | Libraries | Highlights |
|----------|-----------|------------|
| **Python** | 13 (Textual, Rich, urwid, pytermgui, asciimatics, etc.) | Textual+Rich dominate modern Python TUIs |
| **C++** | 12 (FTXUI, range-v3, notcurses, Turbo Vision, etc.) | FTXUI leads; notcurses for GPU TUIs |
| **Rust** | 9 (Ratatui, cursive, termion, crossterm, tuirs, etc.) | Ratatui is the standard; cross-platform |
| **Go** | 7 (Bubble Tea, tview, termui, gocui, etc.) | Bubble Tea is Go's TUI champion |
| **C** | 6 (ncurses, notcurses, stfl, etc.) | Foundational — ncurses powers 200+ TUIs |
| **JavaScript** | 4 (Ink, blessed, neo-blessed, enquirer) | Ink = React for CLI; blessed for raw TUIs |
| **Java** | 4 (lanterna, jcurses, charva, text-io) | lanterna is most active |
| **.NET** | 6 (Terminal.Gui, gui.cs, Spectre.Console, etc.) | Terminal.Gui is the most active |
| **Other** | 6 (D, Nim, Zig, Dart, Swift, Common Lisp) | Dart has dart_ncurses for Flutter TUIs |

### Language Distribution
- **Python**: 13 libraries (20%) — most diverse ecosystem
- **C++**: 12 libraries (18%) — strong systems option
- **Rust**: 9 libraries (14%) — fastest growing
- **Go**: 7 libraries (11%) — excellent for distro TUIs
- **C**: 6 libraries (9%) — foundational
- **.NET**: 6 libraries (9%) — Windows-first
- **JavaScript**: 4 libraries (6%) — React-native-style Ink
- **Java**: 4 libraries (6%) — legacy strength
- **Other**: 6 libraries (9%) — niche but active

### Key TUI Development Insights
1. **Bubble Tea + Ratatui = dominant pair** — 70+ TUIs between them
2. **Rich ecosystem** — `gum` and `charm` tools make shell scripting gorgeous
3. **Ink bridges React devs** to CLI without learning new paradigms
4. **Cross-platform matters** — Ratatui, FTXUI, Textual all target Win/Mac/Linux
5. **TUIs are replacing basic CLIs** in dashboards, file managers, git tools
6. **ncurses is the legacy giant** — but its API shows its age

### TUI vs CLI Decision Guide
- **Use TUI when**: real-time updates (progress, logs), interactive navigation, visual hierarchy with colors/borders/panels, keyboard-driven workflows, multi-pane layouts
- **Use CLI when**: quick one-shot commands, pipe-friendly output, automation/scripts, CI/CD pipelines, remote/SSH over slow connections, accessibility (screen readers)
- **Hybrid approach**: CLI commands that produce output, `--watch`/`--interactive` flags for TUI mode

---

*This skill covers the complete CLI + AI agent ecosystem. For deep dives into specific packages, refer to the individual official documentation.*
