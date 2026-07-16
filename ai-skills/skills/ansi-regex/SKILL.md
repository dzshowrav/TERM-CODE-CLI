---
name: ansi-regex
description: Regular expression for matching ANSI escape codes using the `ansi-regex` library. Supports standard and non-standard codes, with option to match only the first occurrence.
---

# ansi-regex

Regular expression for matching ANSI escape codes in strings.

## Install

```sh
npm install ansi-regex
```

## Usage

```typescript
import ansiRegex from 'ansi-regex';

ansiRegex().test('\u001B[4mcake\u001B[0m'); // true
'\u001B[4mcake\u001B[0m'.match(ansiRegex());
// ['\u001B[4m', '\u001B[0m']

// Match only first occurrence
'\u001B[4mcake\u001B[0m'.match(ansiRegex({ onlyFirst: true }));
// ['\u001B[4m']
```

## Target Processes

- text-processing
- regex-matching
- ansi-code-detection
