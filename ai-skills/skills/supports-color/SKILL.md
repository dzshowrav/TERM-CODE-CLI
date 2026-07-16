---
name: supports-color
description: Detect whether a terminal supports color (basic/256/truecolor) using the `supports-color` library. Also detects `--color`/`--no-color` flags and `FORCE_COLOR` env var.
---

# supports-color

Detect terminal color support level. Returns `false` or an object with `.level` (1=16 colors, 2=256, 3=truecolor) and `.hasBasic`/`.has256`/`.has16m` flags.

## Install

```sh
npm install supports-color
```

## Usage

```typescript
import supportsColor from 'supports-color';

if (supportsColor.stdout) {
  console.log('Terminal supports color');
}
if (supportsColor.stdout.has256) {
  console.log('Terminal supports 256 colors');
}
if (supportsColor.stderr.has16m) {
  console.log('Terminal stderr supports truecolor');
}
```

### Custom stream

```typescript
import { createSupportsColor } from 'supports-color';

const result = createSupportsColor(process.stdout);
if (result) { /* has level/flag properties */ }
```

Options: `{ sniffFlags: false }` to disable `process.argv` checking.

## Info

- Obeys `--color` / `--no-color` / `--color=256` / `--color=16m` CLI flags
- `FORCE_COLOR=0|1|2|3` env var overrides all detection

## Target Processes

- cli-output-formatting
- terminal-capability-detection
