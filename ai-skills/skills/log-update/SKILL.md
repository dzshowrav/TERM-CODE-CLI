---
name: log-update
description: Log by overwriting the previous terminal output using the `log-update` library. Useful for progress bars, spinners, and real-time animations. Performs partial redraws to reduce flicker.
---

# log-update

Render in-place terminal output for progress bars, spinners, animations, etc. Uses partial redraws to minimize flicker.

## Install

```sh
npm install log-update
```

## Usage

```typescript
import logUpdate from 'log-update';

const frames = ['-', '\\', '|', '/'];
let index = 0;
setInterval(() => {
  const frame = frames[index = ++index % frames.length];
  logUpdate(` ♥♥ ${frame} unicorns ${frame} ♥♥ `);
}, 80);
```

## API

| Export | Description |
|--------|-------------|
| `logUpdate(text…)` | Log to stdout (overwrites previous) |
| `logUpdate.clear()` | Clear logged output |
| `logUpdate.done()` | Persist output (stops overwriting) |
| `logUpdate.persist(text…)` | Write output that stays in scrollback |
| `logUpdateStderr(text…)` | Log to stderr (overwrites) |
| `logUpdateStderr.clear()` | Clear stderr output |
| `logUpdateStderr.done()` | Persist stderr output |
| `logUpdateStderr.persist(text…)` | Write persistent stderr output |
| `createLogUpdate(stream, options?)` | Custom stream instance |

### Options

```typescript
type LogUpdateOptions = {
  showCursor?: boolean;  // default false
  defaultWidth?: number; // default 80
  defaultHeight?: number; // default 24
};
```

## Related

- `ora` — elegant terminal spinner (uses log-update)

## Target Processes

- cli-progress-rendering
- real-time-terminal-output
