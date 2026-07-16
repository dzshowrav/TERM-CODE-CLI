---
name: boxen
description: Create styled boxes in the terminal using the `boxen` library. Supports multiple border styles (single, double, round, bold, arrow, classic, none), titles, padding, margin, colors, and fullscreen mode.
---

# boxen

Create boxes in the terminal with configurable borders, padding, margin, titles, and colors.

## Install

```sh
npm install boxen
```

## Usage

```typescript
import boxen from 'boxen';

console.log(boxen('unicorn', { padding: 1 }));
// ┌─────────────┐
// │             │
// │  unicorn    │
// │             │
// └─────────────┘

console.log(boxen('unicorn', { padding: 1, margin: 1, borderStyle: 'double' }));
// ╔═════════════╗
// ║             ║
// ║  unicorn    ║
// ║             ║
// ╚═════════════╝

console.log(boxen('unicorns love rainbows', { title: 'magical', titleAlignment: 'center' }));
// ┌────── magical ───────┐
// │unicorns love rainbows│
// └──────────────────────┘
```

## API

### boxen(text, options?)

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `borderColor` | string | — | 'black', 'red', 'green', 'yellow', 'blue', 'magenta', 'cyan', 'white', 'gray', or hex `'#ff0000'` |
| `borderStyle` | string\|object | `'single'` | `'single'`, `'double'`, `'round'`, `'bold'`, `'singleDouble'`, `'doubleSingle'`, `'classic'`, `'arrow'`, `'none'`, or custom object |
| `dimBorder` | boolean | false | Reduce border opacity |
| `title` | string | — | Title displayed in top border |
| `titleAlignment` | string | `'left'` | `'left'`, `'center'`, `'right'` |
| `width` | number | — | Fixed width |
| `height` | number | — | Fixed height (crops overflow) |
| `fullscreen` | boolean\|function | false | Fit available space; pass `(width, height) => [w, h]` |
| `padding` | number\|object | 0 | Space between text and border |
| `margin` | number\|object | 0 | Space around the box |

## Target Processes

- cli-output-formatting
- terminal-ui-components
