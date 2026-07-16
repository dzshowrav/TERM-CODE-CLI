# Terminal UI Libraries — Complete A-to-Z Reference

---

## 1. Blessed (chjj/blessed)
**GitHub**: https://github.com/chjj/blessed | **Stars**: 11K+
**npm**: `blessed` | **License**: MIT | **Type**: Curses-like widget system

### 1.1 Installation
```bash
npm install blessed
npm install --save-dev @types/blessed  # TypeScript
```

### 1.2 Screen — Main Container
```typescript
import blessed from 'blessed';

const screen = blessed.screen({
  // Core
  program?: blessed.Program,
  smartCSR?: boolean,           // default: true — use CSR for smoother rendering
  title?: string,                // window title
  tabSize?: number,              // default: 4

  // Input
  input?: NodeJS.ReadStream,    // default: process.stdin
  output?: NodeJS.WriteStream,  // default: process.stdout
  log?: string,                  // log file path
  logPad?: number,               // log padding

  // Cursor
  cursor?: 'block' | 'underline' | 'line',
  cursorArtificial?: boolean,
  cursorBlink?: boolean,         // default: false

  // Colors
  forceUnicode?: boolean,       // force unicode (default: auto)
  useBCE?: boolean,             // use background color erase

  // Performance
  fastCSR?: boolean,            // fast CSR detection
  dockBorders?: boolean,         // dock borders between elements
  ignoreLocked?: string[],       // elements to ignore locked
  grabKeys?: boolean,            // grab all keyboard input
  sendFocus?: boolean,           // send focus in/out events
  buffer?: boolean,              // double buffering
  warnings?: boolean,            // show warnings

  // Debug
  terminal?: string,            // TERM override
  fullUnicode?: boolean,        // full unicode support
  dump?: any,                   // dump output
  debug?: boolean,              // debug mode
  features?: any,               // feature flags
});

// Screen Methods
screen.append(element);          // Add widget
screen.render();                 // Render screen
screen.destroy();                // Destroy
screen.key(['q', 'C-c'], () => process.exit(0));  // Key bindings
screen.exec(fn);                 // Execute in screen context
screen.readEditor(options, callback);  // Open editor
screen.setEffects(el, fg, bg, callback);  // Set effects
screen.insertLine(n, y, top, bottom);    // Insert line
screen.deleteLine(n, y, top, bottom);    // Delete line
screen.insertBottom(lines);              // Insert at bottom
screen.insertTop(lines);                 // Insert at top
screen.focusPush(element);     // Push focus
screen.focusPop();             // Pop focus
screen.focusSet(index);        // Set focus
screen.focusNext();            // Next widget
screen.focusPrevious();        // Previous widget
```

### 1.3 Widgets — Complete Reference

#### Box
```typescript
const box = blessed.box({
  // Position
  top: number | 'center' | 'center-1' | string,  // default: 0
  left: number | 'center' | string,               // default: 0
  width: number | 'half' | 'shrink' | string,      // default: 'auto'
  height: number | 'half' | 'shrink' | string,     // default: 'auto'

  // Content
  content: string,
  tags: boolean,                     // parse tags like {bold} {/bold}
  input: boolean,                    // accept input
  label: string | {text: string; side: 'left' | 'right'},

  // Border & Scroll
  border: 'line' | 'bg' | 'ascii' | 'dotted' | 'dashed' | {type: string; fg?: string; bg?: string},
  scrollable: boolean,
  alwaysScroll: boolean,
  scrollbar: {ch?: string; track?: {fg?: string; bg?: string} | boolean; style?: any},

  // Visual
  style: {
    fg?: string | number,
    bg?: string | number,
    bold?: boolean,
    underline?: boolean,
    blink?: boolean,
    inverse?: boolean,
    invisible?: boolean,
    transparent?: boolean,
    border?: {fg?: string; bg?: string},
    scrollbar?: {fg?: string; bg?: string},
    focus?: {fg?: string; bg?: string},
    hover?: {fg?: string; bg?: string},
    label?: {fg?: string; bg?: string},
  },

  // Children behavior
  children?: blessed.Widgets.BlessedElement[],
  shrink?: boolean,
  align?: 'left' | 'center' | 'right',
  valign?: 'top' | 'middle' | 'bottom',
  padding?: number | {top?: number; right?: number; bottom?: number; left?: number},
  wrap?: boolean,
  hidden?: boolean,
  transparent?: boolean,
  shadow?: boolean,
  clickable?: boolean,
  draggable?: boolean,
  keys?: boolean,
  vi?: boolean,
  mouse?: boolean,
  screen?: blessed.Widgets.Screen,
});
```

#### All Widgets
| Widget Class | npm/Usage | Description |
|-------------|-----------|-------------|
| **Screen** | `blessed.screen()` | Root container |
| **Box** | `blessed.box()` | Generic container |
| **Text** | `blessed.text()` | Read-only text |
| **Line** | `blessed.line()` | Horizontal/vertical line |
| **ScrollableText** | `blessed.scrollabletext()` | Scrollable text |
| **Input** | `blessed.input()` | Base input |
| **Textarea** | `blessed.textarea()` | Multi-line text input |
| **Textbox** | `blessed.textbox()` | Single-line text input |
| **Button** | `blessed.button()` | Clickable button |
| **Checkbox** | `blessed.checkbox()` | Checkbox |
| **RadioSet** | `blessed.radioset()` | Radio button group |
| **RadioButton** | `blessed.radiobutton()` | Radio button |
| **List** | `blessed.list()` | Selectable list |
| **Listbar** | `blessed.listbar()` | Horizontal action bar |
| **Table** | `blessed.table()` | Data table |
| **Form** | `blessed.form()` | Input form |
| **Prompt** | `blessed.prompt()` | Input prompt dialog |
| **Question** | `blessed.question()` | Question dialog |
| **Message** | `blessed.message()` | Notification |
| **Loading** | `blessed.loading()` | Loading indicator |
| **ProgressBar** | `blessed.progressbar()` | Progress bar |
| **FileManager** | `blessed.filemanager()` | File browser |
| **Terminal** | `blessed.terminal()` | Embedded terminal |
| **Image** | `blessed.image()` | ANSI image (w3m) |
| **ANSIImage** | `blessed.ansiimage()` | ANSI image |
| **OverlayImage** | `blessed.overlayimage()` | Overlay image |
| **BigText** | `blessed.bigtext()` | Large text |
| **Log** | `blessed.log()` | Scrollable log |
| **ListTable** | `blessed.listtable()` | Table in list |

### 1.4 Events
```typescript
box.on('click', (el, coords) => {});
box.on('keypress', (ch, key) => {});
box.on('focus', (el) => {});
box.on('blur', (el) => {});
box.on('element focus', (el) => {});
box.on('element blur', (el) => {});
box.on('element keypress', (el, ch, key) => {});
box.on('mouse', (mouseEvent) => {});
box.on('resize', (el) => {});
box.on('screen resize', (type, position) => {});
screen.on('key', (ch, key) => {});
screen.on('element keypress', (el, ch, key) => {});
```

### 1.5 Unblessed (Modern Successor)
**GitHub**: https://github.com/avoidwork/unblessed
- TypeScript rewrite with 98.5% test coverage
- Browser support via XTerm.js
- Optional React renderer
- 2355+ passing tests
- Works with shadow DOM, WebWorkers

---

## 2. Blessed-Contrib (yaronn/blessed-contrib)
**GitHub**: https://github.com/yaronn/blessed-contrib
**npm**: `blessed-contrib` | **License**: MIT

### 2.1 Installation
```bash
npm install blessed blessed-contrib
```

### 2.2 Grid Layout
```typescript
import blessed from 'blessed';
import contrib from 'blessed-contrib';

const screen = blessed.screen();
const grid = new contrib.grid({
  rows: 12,       // grid rows
  cols: 12,       // grid columns
  screen: screen,
});
```

### 2.3 All Widgets

#### Line Chart
```typescript
const line = grid.set(0, 0, 6, 6, contrib.line, {
  style: {
    line: 'yellow',
    text: 'green',
    baseline: 'black',
  },
  xLabelPadding: 3,
  xPadding: 5,
  label: 'Title',
  showLegend: true,
  legend: { width: 12 },
  wholeNumbersOnly: false,
  numYLabels: 5,
});

line.setData([
  { title: 'Series 1', x: ['Mon', 'Tue', 'Wed', 'Thu'], y: [10, 20, 15, 25] },
  { title: 'Series 2', x: ['Mon', 'Tue', 'Wed', 'Thu'], y: [5, 15, 10, 20] },
]);
```

#### Bar Chart
```typescript
const bar = grid.set(0, 6, 4, 6, contrib.bar, {
  label: 'Server Load',
  barWidth: 4,
  barSpacing: 6,
  xOffset: 2,
  maxHeight: 9,
  barBgColor: 'green',
});

bar.setData({
  titles: ['Server1', 'Server2', 'Server3'],
  data: [45, 78, 23],
});
```

#### Table
```typescript
const table = grid.set(0, 0, 6, 6, contrib.table, {
  keys: true,
  fg: 'white',
  selectedFg: 'white',
  selectedBg: 'blue',
  interactive: true,
  label: 'Active Processes',
  width: '30%',
  height: '30%',
  border: {type: 'line', fg: 'cyan'},
  columnSpacing: 3,
  columnWidth: [16, 12],
});

table.setData({
  headers: ['col1', 'col2'],
  data: [
    ['value1', 'value2'],
    ['value3', 'value4'],
  ],
});
```

#### Gauge
```typescript
const gauge = grid.set(0, 0, 4, 4, contrib.gauge, {
  label: 'Progress',
  percent: 55,
  style: {
    bar: { bg: 'cyan' },
    label: { fg: 'white', bold: true },
  },
});

// Update
gauge.setPercent(75);
// Or with color
gauge.setData({ percent: 85, barColor: 'green' });
```

#### Donut
```typescript
const donut = grid.set(0, 0, 6, 6, contrib.donut, {
  label: 'Memory Usage',
  radius: 8,
  arcWidth: 3,
  remainColor: 'black',
  yPadding: 2,
  data: [
    { percent: 60, label: 'Used', color: 'red' },
    { percent: 40, label: 'Free', color: 'green' },
  ],
});

// Add data
donut.setData([
  { percent: 75, label: 'Used', color: 'red' },
  { percent: 25, label: 'Free', color: 'green' },
]);
```

#### Sparkline
```typescript
const sparkline = grid.set(0, 6, 4, 6, contrib.sparkline, {
  label: 'Network Activity',
  tags: true,
  style: { fg: 'blue' },
});

sparkline.setData(['Series1', 'Series2'], [
  [10, 20, 30, 20, 15, 25, 35],
  [5, 15, 25, 15, 10, 20, 30],
]);
```

#### Log
```typescript
const log = grid.set(0, 6, 6, 6, contrib.log, {
  fg: 'green',
  selectedFg: 'green',
  label: 'Server Log',
});

log.log('Server started on port 3000');
```

#### Picture
```typescript
const picture = grid.set(0, 0, 12, 12, contrib.picture, {
  file: '/path/to/image.png',
  onError: (err) => console.error(err),
});
```

#### Map
```typescript
const map = grid.set(0, 0, 12, 12, contrib.map, {
  label: 'Geo Distribution',
});
// Requires w3m/img2txt for terminal image display
```

#### Tree
```typescript
const tree = grid.set(0, 0, 6, 6, contrib.tree, {
  fg: 'white',
  selectedFg: 'white',
  selectedBg: 'blue',
  template: { lines: true },
  extended: true,
});

tree.setData({
  extended: true,
  children: {
    'src': {
      extended: true,
      children: {
        'index.ts': {},
        'app.ts': {},
      },
    },
    'package.json': {},
  },
});

tree.on('select', (node) => {
  console.log('Selected:', node.name);
});
```

---

## 3. Chalk (chalk/chalk)
**GitHub**: https://github.com/chalk/chalk | **Stars**: 21K+
**npm**: `chalk` | **Weekly**: 25M+ | **License**: MIT | **Author**: Sindre Sorhus

### 3.1 Installation
```bash
npm install chalk
```

### 3.2 Complete API
```typescript
import chalk from 'chalk';

// === Foreground Colors ===
chalk.black('text');
chalk.red('text');
chalk.green('text');
chalk.yellow('text');
chalk.blue('text');
chalk.magenta('text');
chalk.cyan('text');
chalk.white('text');
chalk.gray('text');
chalk.grey('text');
chalk.blackBright('text');
chalk.redBright('text');
chalk.greenBright('text');
chalk.yellowBright('text');
chalk.blueBright('text');
chalk.magentaBright('text');
chalk.cyanBright('text');
chalk.whiteBright('text');

// === Background Colors ===
chalk.bgBlack('text');
chalk.bgRed('text');
chalk.bgGreen('text');
chalk.bgYellow('text');
chalk.bgBlue('text');
chalk.bgMagenta('text');
chalk.bgCyan('text');
chalk.bgWhite('text');
chalk.bgGray('text');
chalk.bgGrey('text');
chalk.bgBlackBright('text');
chalk.bgRedBright('text');
chalk.bgGreenBright('text');
chalk.bgYellowBright('text');
chalk.bgBlueBright('text');
chalk.bgMagentaBright('text');
chalk.bgCyanBright('text');
chalk.bgWhiteBright('text');

// === Text Styles ===
chalk.bold('text');
chalk.dim('text');
chalk.italic('text');
chalk.underline('text');
chalk.inverse('text');
chalk.hidden('text');
chalk.strikethrough('text');
chalk.visible('text');      // only visible when chalk.level > 0

// === 256 Colors ===
chalk.hex('#DEADED')('text');
chalk.rgb(15, 100, 204)('text');
chalk.ansi256(200)('text');
chalk.bgHex('#DEADED')('text');
chalk.bgRgb(15, 100, 204)('text');
chalk.bgAnsi256(200)('text');

// === Chaining ===
chalk.blue.bgRed.bold('text');
chalk.red.bold.underline('text');

// === Template Literal API ===
console.log(chalk`
  CPU: {red ${cpuPercent}%}
  RAM: {green ${ramUsed}/${ramTotal}}
  {bold Status:} {bgGreen OK}
`);

// === Level Detection ===
chalk.level;  // 0: none, 1: 16 colors, 2: 256 colors, 3: truecolor (16M)
chalk.supportsColor;      // boolean
chalk.supportsColor.level; // 0-3

// === Utility ===
chalk.stderr(text);           // write to stderr with colors
chalk.stderr.blue('error');   // stderr with color

// === Modifiers ===
chalk.reset('text');
```

### 3.3 Chalk Instance (Custom Level)
```typescript
const customChalk = new chalk.Instance({ level: 1 });  // force 16 colors only
```

---

## 4. Gradient-String (bokub/gradient-string)
**GitHub**: https://github.com/bokub/gradient-string
**npm**: `gradient-string` | **License**: MIT

### 4.1 Installation
```bash
npm install gradient-string
```

### 4.2 Complete API
```typescript
import gradient from 'gradient-string';

// === Named Gradients ===
gradient.rainbow('text');
gradient.pastel('text');
gradient.teen('text');
gradient.mind('text');
gradient.morning('text');
gradient.vice('text');
gradient.passion('text');

// === Custom Gradients ===
const custom = gradient(['#FF6B6B', '#4ECDC4', '#45B7D1']);
custom('gradient text');

// === With Colors ===
gradient(['cyan', 'pink'])('Hello world!');

// === Options ===
gradient(['red', 'blue'], {
  interpolation: 'rgb' | 'hsv',  // default: 'rgb'
  hsvSpin: 'short' | 'long',     // default: 'short' — HSV rotation direction
})(text);

// === Multi-line (vertically aligned) ===
gradient.rainbow.multiline('Line 1\nLine 2\nLine 3');

// === Method Signature ===
// gradient(colors: string[], options?: GradientOptions)(text: string): string
// gradient.multiline(text: string): string
```

---

## 5. CLI-Table3 (cli-table/cli-table3)
**GitHub**: https://github.com/cli-table/cli-table3
**npm**: `cli-table3` | **License**: MIT

### 5.1 Installation
```bash
npm install cli-table3
```

### 5.2 Complete API
```typescript
import Table from 'cli-table3';

// === Simple Table ===
const table = new Table({
  head: ['Name', 'Role', 'Status'],     // header row
  colWidths: [20, 15, 10],              // column widths
  colAligns: ['left', 'center', 'right'],  // alignment
  style: {
    'padding-left': 1,
    'padding-right': 1,
    head: ['white', 'bold'],    // header style (colors.js colors)
    border: ['grey'],
    compact: false,
  },
});
table.push(['Alice', 'Engineer', 'Active']);
table.push(['Bob', 'Designer', 'Away']);
console.log(table.toString());

// === Custom Characters ===
const styled = new Table({
  chars: {
    'top': '═', 'top-mid': '╤', 'top-left': '╔', 'top-right': '╗',
    'bottom': '═', 'bottom-mid': '╧', 'bottom-left': '╚', 'bottom-right': '╝',
    'left': '║', 'left-mid': '╟', 'mid': '─', 'mid-mid': '┼',
    'right': '║', 'right-mid': '╢', 'middle': '│',
  },
});

// === Char Options (defaults) ===
chars: {
  'top': '─', 'top-mid': '┬', 'top-left': '┌', 'top-right': '┐',
  'bottom': '─', 'bottom-mid': '┴', 'bottom-left': '└', 'bottom-right': '┘',
  'left': '│', 'left-mid': '├', 'mid': '─', 'mid-mid': '┼',
  'right': '│', 'right-mid': '┤', 'middle': '│',
}

// === Cell Spanning ===
const spanTable = new Table();
spanTable.push(
  { content: 'Cross-column header', colSpan: 2, hAlign: 'center' },
  ['Col 1', 'Col 2']
);

// === Vertical Alignment ===
const vAlignTable = new Table();
vAlignTable.push(
  { content: 'Top', vAlign: 'top' },
  { content: 'Center', vAlign: 'center' },
  { content: 'Bottom', vAlign: 'bottom' },
);

// === Word Wrap ===
const wrapTable = new Table({
  colWidths: [20],
  wordWrap: true,  // wrap on word boundaries, not chars
});

// === Row Spanning ===
const rowSpanTable = new Table();
rowSpanTable.push(
  { content: 'Row spanning', rowSpan: 2 },
  'Row 1 Col 2'
);
rowSpanTable.push('Row 2 Col 2');  // fills the spanned cell

// === Compact Mode ===
const compactTable = new Table({
  style: { compact: true, 'padding-left': 0, 'padding-right': 0 },
  chars: {
    'mid': '', 'left-mid': '', 'mid-mid': '', 'right-mid': '',
    'middle': ' ',
  },
});

// === Truncation ===
const truncTable = new Table({ colWidths: [10] });
truncTable.push(['Long text that will be truncated...']);

// === Complete Style Object ===
style: {
  'padding-left': number,     // default: 1
  'padding-right': number,    // default: 1
  head: string[],             // colors.js styles for header
  border: string[],           // colors.js styles for border
  compact: boolean,           // compact mode
}
```

### 5.3 Methods
| Method | Description |
|--------|-------------|
| `.push(row)` | Add row(s) to table |
| `.toString()` | Render to string |
| `.length` | Number of rows |
| `.options` | Current options |

### 5.4 Row Types
```typescript
// Array row
table.push(['col1', 'col2', 'col3']);

// Object row (with spanning)
table.push({
  content: 'Text',
  colSpan: 2,         // span N columns
  rowSpan: 2,         // span N rows
  hAlign: 'left' | 'center' | 'right',
  vAlign: 'top' | 'center' | 'bottom',
});
```
