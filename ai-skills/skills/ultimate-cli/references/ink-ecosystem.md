# Ink Ecosystem — Complete A-to-Z Reference

## 1. Ink Core (vadimdemedes/ink)
**GitHub**: https://github.com/vadimdemedes/ink | **npm**: `ink` | **License**: MIT
**npm weekly downloads**: ~4.2M | **Stars**: 35K+ | **Author**: Vadim Demedes
**Website**: https://term.ink/

### 1.1 Installation
```bash
npm install ink react
# TypeScript types
npm install --save-dev @types/react
```

### 1.2 render() — Entry Point
```typescript
import { render, type Instance } from 'ink';

const instance: Instance = render(
  <App />,
  options?: RenderOptions
);
```

#### RenderOptions
| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `stdout` | NodeJS.WriteStream | `process.stdout` | Custom stdout stream |
| `stdin` | NodeJS.ReadStream | `process.stdin` | Custom stdin stream |
| `stderr` | NodeJS.WriteStream | `process.stderr` | Custom stderr stream |
| `exitOnCtrlC` | boolean | `true` | Exit app on Ctrl+C |
| `patchConsole` | boolean | `true` | Patch console.log to rerender |
| `debug` | boolean | `false` | Debug mode |

#### Instance Methods
| Method | Return | Description |
|--------|--------|-------------|
| `waitUntilExit()` | `Promise<void>` | Wait until component unmounts |
| `clear()` | `void` | Clear rendered output |
| `unmount()` | `void` | Unmount component |
| `rerender()` | `void` | Rerender root component |

### 1.3 Built-in Components

#### `<Box>` — Flexbox Container
```typescript
import { Box } from 'ink';

<Box
  // FlexDirection
  flexDirection?: 'row' | 'column'        // default: 'row'

  // Alignment
  alignItems?: 'flex-start' | 'center' | 'flex-end'
  justifyContent?: 'flex-start' | 'flex-end' | 'center' | 'space-between' | 'space-around'

  // Flex
  flexGrow?: number
  flexShrink?: number
  flexBasis?: number | string

  // Dimensions
  width?: number | string                   // e.g., 10, '50%'
  height?: number | string
  minWidth?: number | string
  minHeight?: number | string

  // Padding (number or string like '1 2')
  padding?: number
  paddingTop?: number
  paddingBottom?: number
  paddingLeft?: number
  paddingRight?: number
  paddingX?: number                          // left + right
  paddingY?: number                          // top + bottom

  // Margin
  margin?: number
  marginTop?: number
  marginBottom?: number
  marginLeft?: number
  marginRight?: number
  marginX?: number
  marginY?: number

  // Border
  borderStyle?: 'single' | 'double' | 'round' | 'bold'
                | 'singleDouble' | 'doubleSingle' | 'classic'
  borderColor?: string                       // Chalk color

  // Display
  display?: 'flex' | 'none'

  // Position
  position?: 'absolute' | 'relative'
  top?: number | string
  left?: number | string
  bottom?: number | string
  right?: number | string

  // Overflow
  overflow?: 'visible' | 'hidden'

  // Gaps (Ink 5+)
  gap?: number
  columnGap?: number
  rowGap?: number
> {children}</Box>
```

#### `<Text>` — Text Output
```typescript
import { Text } from 'ink';

<Text
  // Colors (chalk-based)
  color?: string                    // 'red', '#ff0000', 'rgb(255,0,0)'
  backgroundColor?: string

  // Styles
  bold?: boolean
  italic?: boolean
  underline?: boolean
  strikethrough?: boolean
  dim?: boolean
  inverse?: boolean

  // Text wrapping
  wrap?: 'wrap' | 'truncate' | 'truncate-start' | 'truncate-middle' | 'truncate-end'
>{children}</Text>
```

#### `<Newline>` — Line Break
```typescript
import { Newline } from 'ink';

<Newline count={1} />    // number of newlines (default: 1)
```

#### `<Spacer>` — Fill Remaining Space
```typescript
import { Spacer } from 'ink';

// Takes up all available space in flex container
<Box>
  <Text>Left</Text>
  <Spacer />
  <Text>Right</Text>
</Box>
```

#### `<Static>` — Persist Children After Rerender
```typescript
import { Static } from 'ink';

<Static items={itemsArray}>
  {(item: ItemType) => <Text key={item.id}>{item.text}</Text>}
</Static>
// items: Array — items to render (each stays mounted)
```

#### `<Transform>` — Transform Child Output
```typescript
import { Transform } from 'ink';

<Transform transform={(children: string) => transformText(children)}>
  <Text>This will be transformed</Text>
</Transform>
```

### 1.4 Built-in Hooks

#### `useApp()` — App Control
```typescript
import { useApp } from 'ink';

const { exit, exitError } = useApp();
exit();          // Exit with code 0
exitError();     // Exit with code 1
```

#### `useInput()` — Keyboard Input
```typescript
import { useInput } from 'ink';

type KeyObject = {
  upArrow?: boolean;
  downArrow?: boolean;
  leftArrow?: boolean;
  rightArrow?: boolean;
  return?: boolean;      // Enter
  escape?: boolean;
  ctrl?: boolean;
  shift?: boolean;
  tab?: boolean;
  backspace?: boolean;
  delete?: boolean;
  meta?: boolean;
  pageUp?: boolean;
  pageDown?: boolean;
};

useInput(
  (input: string, key: KeyObject) => {
    // input: the character pressed
    // key: modifier keys object
    if (key.escape) useApp().exit();
    if (key.return) handleSubmit();
    if (key.ctrl && input === 'c') useApp().exit();
    if (key.upArrow) moveUp();
    if (key.downArrow) moveDown();
  },
  { isActive?: boolean }  // Only process input when active
);
```

#### `useStdin()` — Raw Stdin
```typescript
import { useStdin } from 'ink';

const { stdin, isRawModeSupported, setRawMode } = useStdin();
// stdin: NodeJS.ReadStream
// isRawModeSupported: boolean
// setRawMode: (enabled: boolean) => void
```

#### `useStdout()` — Stdout Access
```typescript
import { useStdout } from 'ink';

const { stdout, write } = useStdout();
// stdout: NodeJS.WriteStream
// write: (data: string) => void — write directly (bypass Ink)
```

#### `useStderr()` — Stderr Access
```typescript
import { useStderr } from 'ink';

const { stderr, write } = useStderr();
```

#### `useFocus()` — Focus Management
```typescript
import { useFocus } from 'ink';

const { isFocused } = useFocus({
  id?: string,       // Unique identifier
  autoFocus?: boolean // Auto-focus on mount
});
```

#### `useFocusManager()` — Focus Controller
```typescript
import { useFocusManager } from 'ink';

const { enableFocus, disableFocus, focusNext, focusPrevious, setFocus } = useFocusManager();

focusNext();       // Move to next focusable
focusPrevious();   // Move to previous
setFocus(id);      // Focus specific component
enableFocus();     // Enable focus handling
disableFocus();    // Disable focus handling
```

#### `useMeasure()` — Component Dimensions
```typescript
import { useMeasure } from 'ink';

const { ref } = useMeasure();
// ref: callback ref — attach to Box
// Returns dimension information of the measured element

<div ref={ref}>...</div>
```

### 1.5 Types
```typescript
import { type Instance, type MeasureResult } from 'ink';

type MeasureResult = {
  width: number;
  height: number;
};
```

### 1.6 Ink Testing Library
```bash
npm install --save-dev ink-testing-library
```

```typescript
import { render } from 'ink-testing-library';

const { lastFrame, frames, rerender, stdin } = render(<MyComponent />);

// lastFrame: string | undefined — current rendered output
// frames: string[] — all frames captured
// rerender: (tree: ReactNode) => void — rerender with new tree
// stdin: { write: (data: string) => void } — simulate input

// Simulate keystrokes:
stdin.write('hello');
stdin.write('\r');     // Enter key
stdin.write('\x03');   // Ctrl+C
```

### 1.7 Ink Version History

| Version | React | Key Changes |
|---------|-------|-------------|
| 1.x | 16.x | Initial release |
| 2.x | 16.x | Better hooks, Yoga upgrade |
| 3.x | 16.x | TypeScript rewrite |
| 4.x | 17.x | Current LTS |
| 5.x | 18.x | Static/Memo changes |
| 6.x | 18+/19 | Latest, better focus management |

---

## 2. ink-text-input
**GitHub**: https://github.com/vadimdemedes/ink-text-input | **Stars**: 183
**npm**: `ink-text-input` | **Weekly**: ~1.4M | **License**: MIT

### Complete Props
```typescript
import TextInput from 'ink-text-input';

<TextInput
  value: string                              // required — current input value
  onChange: (value: string) => void          // required — called on each change
  onSubmit?: (value: string) => void         // called on Enter
  placeholder?: string                       // placeholder text when empty
  focus?: boolean                            // default: true — whether focused
  mask?: string                              // mask character (e.g., '*' for password)
  showCursor?: boolean                       // default: true
  highlightPastedText?: boolean              // highlight pasted content
/>
```

### UncontrolledTextInput
```typescript
import { UncontrolledTextInput } from 'ink-text-input';

<UncontrolledTextInput
  onSubmit?: (value: string) => void
  placeholder?: string
  focus?: boolean
  mask?: string
/>
```

---

## 3. ink-spinner
**GitHub**: https://github.com/vadimdemedes/ink-spinner
**npm**: `ink-spinner` | **Weekly**: ~1.6M | **License**: MIT

### Props
```typescript
import Spinner from 'ink-spinner';

<Spinner type?: string />  // Spinner animation type from cli-spinners
```

### All Spinner Types (from cli-spinners)
```
dots, dots2, dots3, dots4, dots5, dots6, dots7, dots8, dots9, dots10, dots11, dots12,
line, line2, pipe, simpleDots, simpleDots2, star, star2, flip, hamburger, growVertical,
growHorizontal, balloon, balloon2, noise, noise2, nope, arc, circle, squareDots,
toggle, toggle2, toggle3, toggle4, toggle5, toggle6, toggle7, toggle8, toggle9,
toggle10, toggle11, toggle12, toggle13, arrow, arrow2, layer, layer2, betaWave,
aesthetic, clock, earth, moon, runner, pong, shark, monkey, meter, heartbeat,
dots13, dots14, christmas, grenade, point, pointer, smiley, squish, dashed,
bouncingBar, bouncingBall, triangle, binary, camera, circleHalves, circleQuarters,
triangleHalves, triPrism, diamond, sand, dog, cat, dragon
```

### Usage Pattern
```tsx
<Text color="green">
  <Spinner type="dots" /> Loading...
</Text>
```

---

## 4. ink-select-input
**GitHub**: https://github.com/vadimdemedes/ink-select-input
**npm**: `ink-select-input` | **License**: MIT

### Types
```typescript
type Item = {
  label: string;
  value: any;
  key?: string | number;    // keyboard shortcut
};

type Separator = {
  separator: boolean;        // renders divider line
};
```

### Props
```typescript
import SelectInput from 'ink-select-input';

<SelectInput
  items: (Item | Separator)[]         // required — list of items
  isFocused?: boolean                  // default: true
  initialIndex?: number               // default: 0
  limit?: number                      // max visible items (scrolls)
  indicatorComponent?: React.ComponentType<{isSelected?: boolean}> // custom indicator
  itemComponent?: React.ComponentType<{isSelected?: boolean; label: string}> // custom item
  onSelect?: (item: Item) => void      // called on Enter
  onHighlight?: (item: Item) => void   // called on cursor move
/>
```

### Keyboard Controls
| Key | Action |
|-----|--------|
| Up / k | Move up |
| Down / j | Move down |
| Enter | Select |
| 0-9 | Jump to item by index |

---

## 5. ink-multi-select
**GitHub**: https://github.com/karaggeorge/ink-multi-select
**npm**: `ink-multi-select` | **Weekly**: ~16.2K | **License**: MIT

### Types
```typescript
type Item = {
  label: string;
  value: any;
  key?: string | number;
  checked?: boolean;          // default selection
};
```

### Props
```typescript
import MultiSelect from 'ink-multi-select';

<MultiSelect
  items: Item[]                                  // required
  selected?: any[]                                // controlled mode values
  onSelect?: (item: Item) => void                 // toggle item
  onSubmit?: (selected: Item[]) => void           // called on Enter
  focus?: boolean                                 // default: true
  limit?: number                                  // max visible
  indicatorComponent?: React.ComponentType
  checkboxComponent?: React.ComponentType
  itemComponent?: React.ComponentType
/>
```

### Controlled vs Uncontrolled
- **Controlled**: Pass `selected` prop → parent manages state
- **Uncontrolled**: Omit `selected` → internal state

---

## 6. ink-confirm-input
**GitHub**: https://github.com/kevva/ink-confirm-input
**npm**: `ink-confirm-input` | **Weekly**: ~28K | **License**: MIT
⚠️ **Deprecated for Ink 5+**

### Props
```typescript
import ConfirmInput from 'ink-confirm-input';

<ConfirmInput
  isChecked?: boolean               // default: true
  value?: string                    // current input
  onChange?: (value: string) => void
  onSubmit?: (value: string) => void
  placeholder?: string
  focus?: boolean
/>
```

### Modern Replacement (@inkjs/ui)
```tsx
import { ConfirmInput } from '@inkjs/ui';

<ConfirmInput
  defaultValue?: boolean
  onChange?: (value: boolean) => void
/>
```

---

## 7. ink-progress-bar
**GitHub**: https://github.com/brigand/ink-progress-bar
**npm**: `ink-progress-bar` | **License**: MIT

### Props
```typescript
import ProgressBar from 'ink-progress-bar';

<ProgressBar
  percent: number      // 0.0 to 1.0
  left?: number        // left offset
  right?: number       // right offset
  character?: string   // fill character (default: '█')
  pad?: boolean        // padding
/>
```

### @inkjs/ui Version
```tsx
import { ProgressBar } from '@inkjs/ui';

<ProgressBar value={75} />  // 0 to 100
```

---

## 8. ink-table
**GitHub**: https://github.com/maticzav/ink-table
**npm**: `ink-table` | **License**: MIT

### Props
```typescript
import Table from 'ink-table';

<Table
  data: Record<string, any>[]              // required — array of row objects
  columns?: string[]                        // column order/selection
  cell?: React.ComponentType<{value: any; column: string; row: number}>  // custom cell
  header?: React.ComponentType<{value: string}>                          // custom header
  padding?: number                          // cell padding
/>
```

---

## 9. ink-markdown
**GitHub**: https://github.com/cameronhunter/ink-markdown
**npm**: `ink-markdown` | **Weekly**: ~6.3K | **License**: MIT

### Props
```typescript
import Markdown from 'ink-markdown';

<Markdown
  children: string           // markdown string to render
  dim?: boolean              // dim text
  inline?: boolean           // inline mode
  syntaxTheme?: object       // custom syntax highlighting theme
/>
```

### ESM Version (@inkkit/ink-markdown)
```bash
npm install @inkkit/ink-markdown
```

---

## 10. ink-link
**GitHub**: https://github.com/sindresorhus/ink-link | **Stars**: 139
**npm**: `ink-link` | **Weekly**: ~180K | **License**: MIT | **Author**: Sindre Sorhus

### Props
```typescript
import Link from 'ink-link';

<Link
  url: string                    // required — URL to open
  fallback?: boolean | ((children: string, url: string) => ReactNode)  // fallback render
>{children}</Link>
```

### Behavior
- Uses **OSC 8** hyperlink escape sequences
- Works in: iTerm2, Hyper, VS Code terminal, Kitty, Terminal.app, WezTerm, Windows Terminal
- **Fallback**: If `fallback=true`, shows `text (url)` in unsupported terminals

---

## 11. ink-picture (Successor to ink-image)
**GitHub**: https://github.com/endernoke/ink-picture
**npm**: `ink-picture` | **Weekly**: ~68K | **License**: MIT

### Props
```typescript
import Picture, { TerminalInfoProvider } from 'ink-picture';

// MUST wrap in TerminalInfoProvider
<TerminalInfoProvider>
  <Picture
    src: string                          // required — path or URL to image
    width?: number                       // max width in columns
    height?: number                      // max height in rows
    alt?: string                         // fallback text
    protocol?: 'auto' | 'ascii' | 'braille' | 'halfBlock' | 'sixel' | 'iterm2' | 'kitty'
    backgroundColor?: string             // (Ink 5+)
  />
</TerminalInfoProvider>
```

### Protocol Detection Order (auto mode)
1. Kitty (icat escape codes)
2. iTerm2 (inline images protocol)
3. Sixel (graphics protocol)
4. Half-block (unicode half blocks)
5. Braille (braille patterns)
6. ASCII (ASCII art fallback)

### Dependencies
- Uses `sharp` for image processing
- Automatically detects terminal capabilities

---

## 12. ink-big-text
**GitHub**: https://github.com/sindresorhus/ink-big-text | **Stars**: 139
**npm**: `ink-big-text` | **Weekly**: ~73K | **License**: MIT | **Author**: Sindre Sorhus

### Props
```typescript
import BigText from 'ink-big-text';

<BigText
  text: string                      // required — text to display

  // cfonts options:
  font?: string                     // 'block' | 'shade' | 'simple' | '3d' | 'chrome'
                                    // 'huge' | 'slick' | 'grid' | 'pallet' | 'simpleBlock'
                                    // 'simple3d' | 'console' | 'tiny'

  colors?: string[]                 // array of chalk colors
  background?: 'transparent' | string  // background color
  space?: boolean                   // add/remove space between chars
  align?: 'left' | 'center' | 'right'

  // Additional cfonts options:
  gradient?: string[]               // gradient colors
  transitionGradient?: boolean      // smooth gradient
  env?: 'node' | 'browser'
/>
```

### All Fonts (from cfonts)
```
block, shade, simple, 3d, chrome, huge, slick, grid, pallet,
simpleBlock, simple3d, console, tiny
```

### Common Pattern (with ink-gradient)
```tsx
<Gradient name="rainbow">
  <BigText text="HELLO" font="block" />
</Gradient>
```

---

## 13. ink-gradient
**GitHub**: https://github.com/sindresorhus/ink-gradient
**npm**: `ink-gradient` | **Weekly**: ~321K | **License**: MIT | **Author**: Sindre Sorhus

### Props
```typescript
import Gradient from 'ink-gradient';

// Named gradient
<Gradient name="rainbow">
  <Text>Colored text</Text>
</Gradient>

// Custom colors
<Gradient colors={['#FF6B6B', '#4ECDC4', '#45B7D1']}>
  <BigText text="CUSTOM" />
</Gradient>
```

### Built-in Gradients
| Name | Description |
|------|-------------|
| `rainbow` | Full rainbow spectrum |
| `pastel` | Soft pastel colors |
| `teen` | Bright youth colors |
| `mind` | Mind/psychedelic |
| `morning` | Soft morning hues |
| `vice` | Vice city neon |
| `passion` | Red/orange passion |

---

## 14. ink-divider
**GitHub**: https://github.com/JureSotosek/ink-divider
**npm**: `ink-divider` | **Weekly**: ~4.5K | **License**: MIT

### Props
```typescript
import Divider from 'ink-divider';

<Divider
  title?: string                    // centered title text
  dividerColor?: string             // color of the divider line
  titleColor?: string               // color of the title
  padding?: number                  // padding top/bottom
  width?: number                    // width (default: auto)
/>
```

### Output Example
```
─────────────── Section Title ───────────────
```

---

## 15. ink-tab
**GitHub**: https://github.com/jdeniau/ink-tab
**npm**: `ink-tab` | **Weekly**: ~12.6K | **License**: MIT

### Props
```typescript
import { Tabs, Tab } from 'ink-tab';

<Tabs
  onChange: (name: string, tab: TabComponent) => void  // required
  flexBasis?: number                                     // tab width
  isFocused?: boolean                                    // default: true
  keyMap?: Record<string, string>                       // custom key bindings
>
  <Tab name="tab1">Tab 1</Tab>
  <Tab name="tab2">Tab 2</Tab>
</Tabs>

// Tab props:
// name: string — required, unique identifier
// children: ReactNode — tab display content
```

### Key Bindings (default)
| Key | Action |
|-----|--------|
| Left / Shift+Tab | Previous tab |
| Right / Tab | Next tab |

---

## 16. ink-ascii
**GitHub**: https://github.com/hexrcs/ink-ascii
**npm**: `ink-ascii` | **License**: MIT

### Props
```typescript
import Ascii from 'ink-ascii';

<Ascii
  text: string                           // required — text to convert
  font?: string                          // figlet font name
  horizontalLayout?: 'default' | 'full' | 'fitted' | 'controlled smushing'
  verticalLayout?: 'default' | 'full' | 'fitted' | 'controlled smushing'
  width?: number                         // max width
  whitespaceBreak?: boolean              // break on whitespace
/>
```

### All Figlet Fonts (100+)
```
1Row, 3-D, 3D Diagonal, 3D-ASCII, 5 Line Oblique, AMC 3 Line, AMC 3 Liv1,
AMC AAA01, AMC Neko, AMC Razor, AMC Razor2, AMC Slash, AMC Slider,
AMC Thin, AMC Tubes, AMC Untitled, ANSI Regular, ANSI Shadow, ASCII New Roman,
Acrobatic, Alligator, Alligator2, Alpha, Alphabet, Arrows, Avatar, B1FF,
Banner, Banner3-D, Banner3, Banner4, Barbwire, Basic, Bear, Bell, Benjamin,
Big Chief, Big Money-ne, Big Money-nw, Big Money-se, Big Money-sw, Big,
Bigfig, Binary, Block, Blocks, Bloody, Bolger, Braced, Bright, Broadway,
Bubble, Bulbhead, Caligraphy, Caligraphy2, Caligphy, Calvin S, Cards,
Catwalk, Chiseled, Chunky, Coinstak, Cola, Cold, Colossal, Comic, Computer,
Contessa, Contrast, Cosmike, Crawford, Crawford2, Crazy, Cricket, Cursive,
Cyberlarge, Cybermedium, Cybersmall, Cygnet, DS-Digital, Dancing Font,
Decimal, Def Leppard, Delta Corps, Demon, Diamond, Digital, Doom, Double,
Dr Pepper, Efti Chess, Efti Font, Efti Italic, Efti Piti, Efti Robot,
Efti Wall, Efti Water, Electronic, Elite, Epic, Etched, Fender, Filter,
Fire Font, Fire Font-kz, Flipped, Floral, Flower Power, Fly, Four Tops,
Fraktur, Funny, Galahad, Ghost, Ghoulish, Gidy, Gimel, Gloom, Goofy,
Gothic, Graceful, Gradient, Graffiti, Greek, Green Bevel, H2O2, Happy,
Hash, Heavy, Helvet, Herox, Hiero, High Noon, Knight, Konto slant, Kontors,
Krit, L4Me, LA, Latin, Lean, Letters, Lexible, LI, Lil Devil, Line Blocks,
Linux, Lockergnome, Madrid, Marquee, Max Four, Maxi, Mayhem, Maze,
Merlin, Messy, Mini, Miniwi, Mirjam, Mirror, Mnemonic, Modular, Morse,
Morton, Mosco, Moscow, Mshebrew210, Mushroom, NS-Times, NScript, NV Script,
Nancyj-Fancy, Nancyj-Improved, Nancyj-Underlined, Nancyj, Neep, Neg, Neo,
Neon, New Rep, Nipples, Notched, Notes, Nuclear, O8, OCR-A, OCR-B,
Oblivion, Ogre, Old Banner, Olympiad, Orange, Outrun, Oz, PBM, PC, PCb,
Panther, Pawp, Peaks, Pearl, Peppers, Phantom, Pism, Plancy2, Poast, Poison,
Pony, Poop, Portal, Posse, Pound, Pound-cw, Presbyterian, Pseudo, Punk,
Puppy, Pure, Pyramids, R2-D2, Radical, Rainbow, Rammstein, Random, Rectangles,
Red Phoenix, Relief, Relieve, Rescue, Retro, Rev, Reverse, Rich, Rounded,
Rowenta, Ruby, Rune, Runyc, S Blood, SL Script, S T Forks, Saint Moritz,
Sans, Sans Serif, Scat, Sblood, Script, Serif, Serifcap, Shadow, Shangri-La,
Sharp, SheWas, Short, Slant, Slant Relief, Slide, Small, Small Caps, Small Isometric,
Small Keyboard, Small Poison, Small Script, Small Shadow, Small Slant, Small Tengwar,
Soft, Soft Serve, Sokol, Soviet, Speed, Spider, Spiral, Spliff, Square,
Stacey, Stampate, Stampatello, Standard, Star Strips, Star Wars, Stellar,
Stforek, Sticky, Stop, Stora, Straight, Strong, Sub-Zero, Swan, Sweet,
Synth, Sys, TN, TRASH, Tafel, Tarty, Tengwar, Term, Test, The Edge,
Thick, Thin, Thorned, Three Point, Ticks, Tickslant, Tiles, Tim, Tired,
Tombstone, Train, Trek, Tsal, Tubular, Twisted, Two Point, USA Flag,
UTMF, UTOPIA, Ultra, Uncle, Underdog, Uni, University, Varsity, Wavy,
Weird, Wet Letter, Whimsy, Wolf, XBrite, XCalibur, Xatc, Xding, Xheights,
Yigyan, Yosif
```

---

## 17. ink-quicksearch-input
**GitHub**: https://github.com/Eximchain/ink-quicksearch-input
**npm**: `ink-quicksearch-input` | **Ink 4**: `@inkkit/ink-quicksearch-input` | **License**: MIT

### Types
```typescript
type Item = {
  label: string;
  value: any;
};
```

### Props
```typescript
import QuickSearchInput from 'ink-quicksearch-input';

<QuickSearchInput
  items: Item[]                           // required — items to search
  onSelect: (item: Item) => void          // required — called on selection
  focus?: boolean                         // default: true
  caseSensitive?: boolean                 // default: false
  limit?: number                          // max visible items
  forceMatchingQuery?: boolean            // force input to match existing item
  clearQueryChars?: string               // chars to clear query (default: ESC)
  initialSelectionIndex?: number          // default: 0
  label?: string                          // input label
  placeholder?: string                    // input placeholder
  indicatorComponent?: React.ComponentType
  itemComponent?: React.ComponentType
/>
```

### Keyboard Controls
| Key | Action |
|-----|--------|
| Type | Filter items (real-time search) |
| Up/Down | Navigate items |
| Enter | Select highlighted |
| ESC | Clear search query |

---

## 18. @inkjs/ui — Official UI Kit
**GitHub**: https://github.com/vadimdemedes/ink-ui | **Stars**: 2K
**npm**: `@inkjs/ui` | **Weekly**: ~257K | **License**: MIT | **Latest**: v2.0.0

### Installation
```bash
npm install @inkjs/ui
```

### TextInput
```tsx
import { TextInput } from '@inkjs/ui';

<TextInput
  placeholder?: string
  defaultValue?: string
  onSubmit?: (value: string) => void
  onChange?: (value: string) => void
  validate?: (value: string) => string | undefined  // return error message or undefined
  isDisabled?: boolean
/>
```

### Select
```tsx
import { Select } from '@inkjs/ui';

<Select
  options: { label: string; value: any }[]
  onChange: (value: any) => void
  defaultValue?: any
  isDisabled?: boolean
/>
```

### MultiSelect
```tsx
import { MultiSelect } from '@inkjs/ui';

<MultiSelect
  options: { label: string; value: any; checked?: boolean }[]
  onChange: (values: any[]) => void
  defaultValue?: any[]
  isDisabled?: boolean
/>
```

### ConfirmInput
```tsx
import { ConfirmInput } from '@inkjs/ui';

<ConfirmInput
  defaultValue?: boolean
  onChange?: (value: boolean) => void
  isDisabled?: boolean
/>
```

### Spinner
```tsx
import { Spinner } from '@inkjs/ui';

<Spinner
  type?: string   // cli-spinners type
  label?: string  // label next to spinner
/>
```

### ProgressBar
```tsx
import { ProgressBar } from '@inkjs/ui';

<ProgressBar
  value: number       // 0 to 100
  type?: 'bar' | 'percent'  // display style
/>
```

### StatusMessage
```tsx
import { StatusMessage } from '@inkjs/ui';

<StatusMessage
  variant: 'info' | 'success' | 'warning' | 'error' | 'muted'
>{message}</StatusMessage>
```

### Key — Display Keyboard Shortcut
```tsx
import { Key } from '@inkjs/ui';

<Key>Ctrl+C</Key>   // Renders highlighted key display
```

### KeyMap — Keyboard Legend
```tsx
import { KeyMap } from '@inkjs/ui';

<KeyMap>
  <KeyMap.Group>
    <KeyMap.Key shortcut="q" /> — Quit
    <KeyMap.Key shortcut="↑↓" /> — Navigate
  </KeyMap.Group>
</KeyMap>
```

---

## 19. Ink Web (ink-web.dev)
**Website**: https://www.ink-web.dev
**GitHub**: https://github.com/cjroth/ink-web

### Concept
Run Ink CLI apps in the browser using XTerm.js. Uses shadcn-style registry for components.

### Setup
```bash
npx shadcn@latest add https://raw.githubusercontent.com/cjroth/ink-web/main/packages/ink-ui/registry/[component].json
```

### CSS Requirements
```css
import 'ink-web/css';
import '@xterm/xterm/css/xterm.css';
```

### Components Available (Ink Web)
| Component | Description |
|-----------|-------------|
| `InkXterm` | Terminal emulator wrapper |
| `Table` | Browser-compatible table |
| `Ascii` | Browser-compatible ASCII art |
| `Spinner` | Browser-compatible spinner |
| `Key` | Browser-compatible key display |

### Architecture
At build time, aliases `ink` → `ink-web`, enabling the same Ink components to render in both terminal and browser environments.
