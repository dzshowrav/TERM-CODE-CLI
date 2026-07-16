---
name: strip-ansi
description: Strip ANSI escape codes from a string using the `strip-ansi` library. Removes color, cursor, and formatting escape sequences from terminal output strings.
---

# strip-ansi

Strip ANSI escape codes from a string. Node.js has a built-in `stripVTControlCharacters` but this package provides consistent cross-version behavior.

## Install

```sh
npm install strip-ansi
```

## Usage

```typescript
import stripAnsi from 'strip-ansi';

stripAnsi('\u001B[4mUnicorn\u001B[0m'); // 'Unicorn'
stripAnsi('\u001B]8;;https://github.com\u0007Click\u001B]8;;\u0007'); // 'Click'
```

## Related

- `has-ansi` — check if a string has ANSI codes
- `ansi-regex` — regex for matching ANSI codes
- `strip-ansi-cli` — CLI for this module

## Target Processes

- text-processing
- terminal-output-sanitization
