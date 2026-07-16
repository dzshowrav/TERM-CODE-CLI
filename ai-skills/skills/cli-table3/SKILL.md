# cli-table3

Render unicode-aided tables in the command line from Node.js scripts. API-compatible with `cli-table` / `cli-table2`.

## Install

```bash
npm install cli-table3
```

## Usage

```js
const Table = require('cli-table3');
```

## Table Types

### Horizontal

```js
const table = new Table({
  head: ['TH 1 label', 'TH 2 label'],
  colWidths: [100, 200]
});
table.push(
  ['First value', 'Second value'],
  ['First value', 'Second value']
);
console.log(table.toString());
```

`table` is an Array — use `push`, `unshift`, `splice`, etc.

### Vertical

```js
const table = new Table();
table.push(
  { 'Some key': 'Some value' },
  { 'Another key': 'Another value' }
);
console.log(table.toString());
```

### Cross

Requires `head` with empty string as first header. Rows as `{ "Header": ["Row", "Values"] }`.

```js
const table = new Table({
  head: ["", "Top Header 1", "Top Header 2"]
});
table.push(
  { 'Left Header 1': ['Value Row 1 Col 1', 'Value Row 1 Col 2'] },
  { 'Left Header 2': ['Value Row 2 Col 1', 'Value Row 2 Col 2'] }
);
console.log(table.toString());
```

## Options

### `chars` — Custom border characters

```js
const table = new Table({
  chars: {
    'top': '═', 'top-mid': '╤', 'top-left': '╔', 'top-right': '╗',
    'bottom': '═', 'bottom-mid': '╧', 'bottom-left': '╚', 'bottom-right': '╝',
    'left': '║', 'left-mid': '╟', 'mid': '─', 'mid-mid': '┼',
    'right': '║', 'right-mid': '╢', 'middle': '│'
  }
});
```

Set `mid`, `left-mid`, `mid-mid`, `right-mid` to `''` to hide separator rows:

```js
const table = new Table({
  chars: { mid: '', 'left-mid': '', 'mid-mid': '', 'right-mid': '' }
});
```

Compact table with no decorations:

```js
const table = new Table({
  chars: { top: '', 'top-mid': '', 'top-left': '', 'top-right': '',
           bottom: '', 'bottom-mid': '', 'bottom-left': '', 'bottom-right': '',
           left: '', 'left-mid': '', mid: '', 'mid-mid': '',
           right: '', 'right-mid': '', middle: ' ' },
  style: { 'padding-left': 0, 'padding-right': 0 }
});
```

### `style` — Colors and formatting

Styles applied via `ansis` (optional dependency). Supports base 16 named colors, `hex(#CODE)` for truecolor, `bgHex(#CODE)` for background.

```js
const table = new Table({
  head: ['Name', 'Age'],
  style: {
    head: ['green', 'bold'],           // or ['hex(#FFA500)', 'italic']
    border: ['hex(#FFD700)'],
    'padding-left': 1,
    'padding-right': 1
  }
});
```

Colorize individual cells with `ansis`:

```js
const ansi = require('ansis');
table.push([ansi.green('Walter'), ansi.red('50')]);
```

### `colWidths` — Array of column widths (in chars)

### `wordWrap` — Boolean (default: true). Wrap on word boundaries

### `wrapOnWordBoundary` — Boolean (default: true). Set `false` to wrap chars (requires `wordWrap: true`)

### `colAligns` — Array: `'left'`, `'center'`, `'right'` per column

### `rowAligns` — Array: `'top'`, `'center'`, `'bottom'` per row

### `rowHeights` — Array of row heights

### `span` — Cell spanning (rows/cols). See [advanced-usage](https://github.com/cli-table/cli-table3/blob/master/advanced-usage.md)

### `debug` — Integer debug level. Messages at `table.messages[]`

## Features

- Cells span columns/rows
- Custom styles per cell (chars, colors, padding)
- Vertical alignment: top, center, bottom
- Word wrapping (`wordWrap: true`, `wrapOnWordBoundary`)
- Hyperlinked cells via `{ content, href }` objects
- ANSI-safe truncation
- Truecolor support via `hex(#CODE)` / `bgHex(#CODE)`

## Cell Object Format

```js
{ content: 'Text', href: 'http://example.com', colSpan: 2, rowSpan: 1 }
```

Pass objects in row arrays for per-cell control over content, hyperlinks, and spanning.
