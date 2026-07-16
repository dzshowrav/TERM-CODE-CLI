---
name: ansi-escapes
description: Use the `ansi-escapes` library (sindresorhus) to generate ANSI escape codes for cursor positioning, screen clearing, text erasing, terminal mode switching, image display, clickable links, iTerm2 annotations, and synchronized output.
---

# ansi-escapes

Use `ansi-escapes` when you need to manipulate the terminal via ANSI escape codes: cursor movement, screen/line erasing, scroll, beep, hyperlinks, image display, alternative screen, and synchronized output.

## Install

```sh
npm install ansi-escapes
```

## Usage

```typescript
import ansiEscapes from 'ansi-escapes';

// Move cursor two rows up and to the left
process.stdout.write(ansiEscapes.cursorUp(2) + ansiEscapes.cursorLeft);
```

Named imports:

```typescript
import { cursorUp, cursorLeft } from 'ansi-escapes';
```

Also works with Xterm.js in the browser:

```typescript
import ansiEscapes from 'ansi-escapes';
import { Terminal } from 'xterm';

const terminal = new Terminal({});
terminal.write(ansiEscapes.cursorUp(2) + ansiEscapes.cursorLeft);
```

## API

### Cursor Positioning

| Export | Type | Description |
|--------|------|-------------|
| `cursorTo(x, y?)` | `(number, number?) => string` | Absolute cursor position. `x0 y0` = top-left |
| `cursorMove(x, y?)` | `(number, number?) => string` | Relative cursor position |
| `cursorUp(count?)` | `(number?) => string` | Move up N rows (default 1) |
| `cursorDown(count?)` | `(number?) => string` | Move down N rows (default 1) |
| `cursorForward(count?)` | `(number?) => string` | Move forward N columns (default 1) |
| `cursorBackward(count?)` | `(number?) => string` | Move backward N columns (default 1) |
| `cursorLeft` | `string` | Move cursor to left side |
| `cursorSavePosition` | `string` | Save cursor position |
| `cursorRestorePosition` | `string` | Restore saved cursor position |
| `cursorGetPosition` | `string` | Get cursor position (device status report) |
| `cursorNextLine` | `string` | Move to next line |
| `cursorPrevLine` | `string` | Move to previous line |
| `cursorHide` | `string` | Hide cursor |
| `cursorShow` | `string` | Show cursor |

### Erasing

| Export | Type | Description |
|--------|------|-------------|
| `eraseLines(count)` | `(number) => string` | Erase N rows from current position up |
| `eraseEndLine` | `string` | Erase to end of line |
| `eraseStartLine` | `string` | Erase to start of line |
| `eraseLine` | `string` | Erase entire current line |
| `eraseDown` | `string` | Erase screen from current line down |
| `eraseUp` | `string` | Erase screen from current line up |
| `eraseScreen` | `string` | Erase entire screen, cursor to top-left |

### Screen & Scroll

| Export | Type | Description |
|--------|------|-------------|
| `scrollUp` | `string` | Scroll display up one line |
| `scrollDown` | `string` | Scroll display down one line |
| `clearViewport` | `string` | Clear visible screen only (safe, no scrollback or state change) |
| `clearScreen` | `string` | Full clear (uses RIS — may clear scrollback and reset modes in some terminals) |
| `clearTerminal` | `string` | Clear everything including scrollback buffer |

### Terminal Mode Switching

| Export | Type | Description |
|--------|------|-------------|
| `enterAlternativeScreen` | `string` | Switch to alternative screen buffer |
| `exitAlternativeScreen` | `string` | Exit alternative screen buffer |

### Synchronized Output (reduce flicker)

| Export | Type | Description |
|--------|------|-------------|
| `beginSynchronizedOutput` | `string` | Begin atomic render group |
| `endSynchronizedOutput` | `string` | End atomic render group |
| `synchronizedOutput(text)` | `(string) => string` | Wrap text in sync output sequences |

### Utilities

| Export | Type | Description |
|--------|------|-------------|
| `beep` | `string` | Output beeping sound |
| `link(text, url)` | `(string, string) => string` | Create clickable hyperlink |
| `image(filePath, options?)` | `(string, ImageOptions?) => string` | Display image in terminal |
| `setCwd(path?)` | `(string?) => string` | Set CWD for iTerm2 + ConEmu |

### iTerm2

| Export | Type | Description |
|--------|------|-------------|
| `iTerm.setCwd(path?)` | `(string?) => string` | Inform iTerm2 of CWD |
| `iTerm.annotation(msg, opts?)` | `(string, AnnotationOptions?) => string` | Create iTerm2 annotation |

### ConEmu

| Export | Type | Description |
|--------|------|-------------|
| `ConEmu.setCwd(path?)` | `(string?) => string` | Inform ConEmu of CWD |

## ImageOptions

```typescript
type ImageOptions = {
  width?: 'auto' | number | string;   // N=chars, Npx=pixels, N%=percent
  height?: 'auto' | number | string;
  preserveAspectRatio?: boolean;        // default true
};
```

## AnnotationOptions

```typescript
type AnnotationOptions = {
  length?: number;    // columns to annotate (default: rest of line)
  x?: number;         // starting X (default: cursor position)
  y?: number;         // starting Y (default: cursor position)
  isHidden?: boolean; // hidden until "Show Annotations" cmd
};
```

## Related

- `ansi-styles` — ANSI escape codes for styling strings (colors, bold, italic, etc.)
- `supports-hyperlinks` — detect if terminal supports clickable links
- `term-img` — higher-level image display module

## Dependencies

```json
{
  "dependencies": {
    "ansi-escapes": "^7.0.0",
    "environment": "^1.0.0"
  }
}
```

## Target Processes

- cli-output-formatting
- terminal-cursor-control
- screen-management
- tui-rendering
