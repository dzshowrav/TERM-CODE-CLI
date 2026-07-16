---
name: has-ansi
description: Check if a string has ANSI escape codes using the `has-ansi` library. Returns boolean for whether terminal formatting codes are present.
---

# has-ansi

Check if a string contains ANSI escape codes.

## Install

```sh
npm install has-ansi
```

## Usage

```typescript
import hasAnsi from 'has-ansi';

hasAnsi('\u001B[4mUnicorn\u001B[0m'); // true
hasAnsi('cake'); // false
```

## Related

- `strip-ansi` — strip ANSI codes
- `ansi-regex` — regex for matching ANSI codes

## Target Processes

- text-validation
- terminal-output-sanitization
