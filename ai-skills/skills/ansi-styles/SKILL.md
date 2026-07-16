---
name: ansi-styles
description: Use the `ansi-styles` library (chalk/sindresorhus) to generate ANSI escape codes for terminal text styling — colors, background colors, and modifiers (bold, dim, italic, underline, etc.) with support for 16, 256, and 16 million truecolor.
---

# ansi-styles

Use `ansi-styles` when you need low-level ANSI escape codes for styling terminal strings: foreground/background colors (16/256/truecolor) and text modifiers. For a higher-level API, use `chalk` instead.

## Install

```sh
npm install ansi-styles
```

## Usage

```typescript
import styles from 'ansi-styles';

console.log(`${styles.green.open}Hello world!${styles.green.close}`);

// Color conversion between 256/truecolor
console.log(`${styles.color.ansi(styles.rgbToAnsi(199, 20, 250))}Hello World${styles.color.close}`);
console.log(`${styles.color.ansi256(styles.rgbToAnsi256(199, 20, 250))}Hello World${styles.color.close}`);
console.log(`${styles.color.ansi16m(...styles.hexToRgb('#abcdef'))}Hello World${styles.color.close}`);
```

## API

### `open` and `close`

Every style has `open` and `close` string properties.

### Arrays

- `modifierNames` — all modifier names (e.g., `'bold'`, `'italic'`)
- `foregroundColorNames` — all foreground color names
- `backgroundColorNames` — all background color names
- `colorNames` — combined foreground + background

Use for validation:

```typescript
import { modifierNames } from 'ansi-styles';
modifierNames.includes('bold'); // true
```

### Groups (non-enumerable)

- `styles.modifier`
- `styles.color`
- `styles.bgColor`

### Raw codes

```typescript
styles.codes.get(36); // 39 (Map<openCode, closeCode>)
```

## Styles

### Modifiers

| Name | Notes |
|------|-------|
| `reset` | |
| `bold` | |
| `dim` | |
| `italic` | Not widely supported |
| `underline` | |
| `overline` | VTE, GNOME terminal, mintty, Git Bash |
| `inverse` | |
| `hidden` | |
| `strikethrough` | Not widely supported |

### Colors / Backgrounds (`bg` prefix)

16 base colors + 8 bright variants:

- `black` / `red` / `green` / `yellow` / `blue` / `magenta` / `cyan` / `white`
- `blackBright` (alias: `gray`, `grey`) + bright variants
- `bgBlack` / `bgRed` / ... / `bgWhite` + `bgBlackBright` (alias: `bgGray`, `bgGrey`) + bright variants

## Color Conversion

| Function | Input | Output |
|----------|-------|--------|
| `rgbToAnsi(r, g, b)` | 0-255 each | 16-color code |
| `rgbToAnsi256(r, g, b)` | 0-255 each | 256-color code |
| `hexToRgb(hex)` | `'#rgb'` or `'#rrggbb'` | `[r, g, b]` |
| `hexToAnsi(hex)` | hex string | 16-color code |
| `hexToAnsi256(hex)` | hex string | 256-color code |

Usage:

```typescript
styles.color.ansi(styles.rgbToAnsi(100, 200, 15));        // 16-color fg
styles.bgColor.ansi(styles.hexToAnsi('#C0FFEE'));          // 16-color bg
styles.color.ansi256(styles.rgbToAnsi256(100, 200, 15));   // 256-color fg
styles.color.ansi16m(100, 200, 15);                         // truecolor fg
styles.bgColor.ansi16m(...styles.hexToRgb('#C0FFEE'));      // truecolor bg
```

## Dependencies

Zero dependencies.

## Related

- `ansi-escapes` — ANSI escape codes for cursor/screen manipulation
- `chalk` — higher-level terminal styling (this library's consumer)
- `supports-color` — detect terminal color support

## Target Processes

- cli-output-formatting
- terminal-text-styling
- ansi-color-management
