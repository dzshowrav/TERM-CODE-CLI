---
name: wrap-ansi
description: Wordwrap a string containing ANSI escape codes to a specified column width using the `wrap-ansi` library. Preserves ANSI styling across line breaks.
---

# wrap-ansi

Wordwrap a string with ANSI escape codes to a given column width. Preserves all styling (colors, bold, etc.) across wrapped lines.

## Install

```sh
npm install wrap-ansi
```

## Usage

```typescript
import chalk from 'chalk';
import wrapAnsi from 'wrap-ansi';

const input = 'The quick brown ' + chalk.red('fox jumped over ') + 'the lazy ' + chalk.green('dog and then ran away with the unicorn.');
console.log(wrapAnsi(input, 20));
```

## API

### wrapAnsi(string, columns, options?)

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `hard` | boolean | false | Hard wrap at column width (default: soft wrap) |
| `wordWrap` | boolean | true | Split words at spaces; if false, fills each column |
| `trim` | boolean | true | Remove whitespace on all lines |

## Related

- `slice-ansi` — slice a string with ANSI codes
- `cli-truncate` — truncate string to terminal width

## Target Processes

- cli-output-formatting
- text-wrapping
