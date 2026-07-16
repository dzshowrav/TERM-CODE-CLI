---
name: string-width
description: Get the visual width of a string (columns required to display it) using the `string-width` library. Handles fullwidth characters, ANSI escape codes (stripped), and ambiguous-width characters.
---

# string-width

Get the visual width of a string in terminal columns. Fullwidth Unicode characters count as 2, ANSI codes are stripped.

## Install

```sh
npm install string-width
```

## Usage

```typescript
import stringWidth from 'string-width';

stringWidth('a');          // 1
stringWidth('古');         // 2
stringWidth('\u001B[1m古\u001B[22m'); // 2
```

## API

### stringWidth(string, options?)

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `ambiguousIsNarrow` | boolean | true | Count ambiguous-width chars as narrow (1) |
| `countAnsiEscapeCodes` | boolean | false | Whether to count ANSI codes in the width |

## Related

- `widest-line` — get the visual width of the widest line
- `get-east-asian-width` — determine East Asian Width of a character

## Target Processes

- cli-output-formatting
- text-measurement
