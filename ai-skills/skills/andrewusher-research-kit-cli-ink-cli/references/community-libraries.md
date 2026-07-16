# Ink Community Libraries

Essential third-party libraries that extend Ink's capabilities for AI agent terminal interfaces.

## ink-chart

Render charts and graphs in the terminal.

- **npm**: `npm install ink-chart`
- **License**: MIT
- **Size**: 3.7 kB
- **Dependencies**: ervy, prop-types

### Usage

```tsx
import React from 'react';
import {render} from 'ink';
import Chart from 'ink-chart';

const data = [
  { key: 'Tokens Used', value: 45230 },
  { key: 'API Calls',   value: 1280 },
  { key: 'Cost ($)',    value: 0.89 },
];

render(<Chart data={data} type="bar" />);
```

### API

#### `<Chart />`

| Prop      | Type     | Description                                              |
| --------- | -------- | -------------------------------------------------------- |
| `data`    | `array`  | Array of `{ key, value }` objects to display             |
| `type`    | `string` | Chart type: `'bar'`, `'pie'`, and others from ervy       |
| `options` | `object` | Any options supported by [ervy](https://github.com/chunqiuyiu/ervy) |

### AI Use Cases

- Token count and cost tracking dashboards
- API call volume monitoring
- Processing speed sparklines
- Real-time metrics for agent operations

---

## ink-gradient

Apply ANSI color gradients to terminal text.

- **npm**: `npm install ink-gradient`
- **License**: MIT
- **Size**: 11.2 kB
- **Author**: Sindre Sorhus
- **Dependencies**: gradient-string, strip-ansi

### Usage

```tsx
import React from 'react';
import {render, Text} from 'ink';
import Gradient from 'ink-gradient';
import BigText from 'ink-big-text';

render(
  <Gradient name="rainbow">
    <BigText text="AI AGENT READY"/>
  </Gradient>
);
```

### API

#### `<Gradient>`

| Prop       | Type                  | Description                                                |
| ---------- | --------------------- | ---------------------------------------------------------- |
| `children` | `string \| Component` | Content to colorize. Multiple `<Text>` treated as separate |
| `name`     | `string`              | Built-in gradient name (mutually exclusive with `colors`)  |
| `colors`   | `string[] \| object[]`| Custom color array (mutually exclusive with `name`)        |

#### Built-in Gradient Names

`rainbow`, `pastel`, `teen`, `mind`, `morning`, `vice`, `passion`, `fruit`, `instagram`, `atlas`, `retro`, `summer`, `candy`, `mine`

#### Custom Colors

```tsx
<Gradient colors={['#00FFAA', '#FF00AA', '#FFAA00']}>
  <Text>Branded Gradient Title</Text>
</Gradient>
```

### AI Use Cases

- Emphasized agent titles and section headers
- Visual separation between AI output and user commands
- Branded CLI splash screens
- Status indicators with color-coded gradients

---

## ink-picture

Display images in the terminal with automatic protocol detection and graceful fallbacks. Supports Sixel, Kitty, iTerm2, half-block, braille, and ASCII art.

- **npm**: `npm install ink-picture`
- **License**: MIT
- **Size**: 116.5 kB
- **Author**: endernoke
- **Dependencies**: chalk, jimp, sixel, node-fetch, supports-color

### Usage

```tsx
import React from 'react';
import {Box, render} from 'ink';
import Image, {InkPictureProvider} from 'ink-picture';

function App() {
  return (
    <InkPictureProvider>
      <Image
        src="https://picsum.photos/200/200"
        width={20}
        height={10}
        alt="Example image"
      />
    </InkPictureProvider>
  );
}

render(<App />);
```

### API

#### `<InkPictureProvider />`

Required wrapper at the app root. Detects terminal capabilities.

| Prop                    | Type       | Default | Description                                    |
| ----------------------- | ---------- | ------- | ---------------------------------------------- |
| `terminalInfo`          | `object`   | —       | Override detected terminal capabilities        |
| `config`                | `object`   | —       | Library-wide configuration                     |
| `config.pollIntervalMs` | `number`   | `16`    | Layout polling interval (ms)                   |
| `config.cacheSize`      | `number`   | `10`    | Max cached images (0 = disable)                |
| `onTerminalInfoDetection` | `function` | —     | Callback after terminal detection completes    |

#### `<Image />`

| Prop         | Type                              | Default    | Description                               |
| ------------ | --------------------------------- | ---------- | ----------------------------------------- |
| `src`        | `string \| ArrayBuffer \| Buffer` | —          | Image source (URL, path, or buffer)       |
| `width`      | `number \| string`                | `"100%"`   | Width in terminal cells or percentage     |
| `height`     | `number \| string`                | `"100%"`   | Height in terminal cells or percentage    |
| `alt`        | `string`                          | —          | Alt text for screen readers and loading   |
| `objectFit`  | `"fill" \| "contain" \| "cover"`  | `"fill"`   | Image resize behavior                     |
| `protocol`   | `string \| object`                | auto       | Force a specific rendering protocol       |
| `getVisibility` | `function`                     | —          | Custom visibility detection callback      |

### Rendering Protocols

| Protocol    | Resolution   | Requirements              | Priority |
| ----------- | ------------ | ------------------------- | -------- |
| `kitty`     | Full         | Kitty-compatible terminal | 1st      |
| `iterm2`    | Full         | iTerm2                    | 2nd      |
| `sixel`     | Full         | Sixel support             | 3rd      |
| `halfBlock` | 1×2 per cell | Unicode + color           | 4th      |
| `braille`   | 2×4 per cell | Unicode support           | 5th      |
| `ascii`     | 1×1 per cell | Universal fallback        | 6th      |

### Individual Protocol Components

```tsx
import {
  AsciiImage,
  BrailleImage,
  HalfBlockImage,
  SixelImage,
  KittyImage,
  ITerm2Image,
} from 'ink-picture';
```

### Hooks

```tsx
import {
  useInkPictureConfig,
  useTerminalInfo,
  useImageCache,
} from 'ink-picture';
```

### AI Use Cases

- Display AI-generated images from vision models
- Show multi-modal AI results (image analysis)
- Render screenshots and diagrams in CLI
- ASCII art fallback for universal compatibility

---

## Combined Example: AI Agent Dashboard

```tsx
import React from 'react';
import {render, Box, Text} from 'ink';
import Gradient from 'ink-gradient';
import Chart from 'ink-chart';
import Image, {InkPictureProvider} from 'ink-picture';

function AIDashboard() {
  const usage = [
    { key: 'GPT-4',   value: 15230 },
    { key: 'Claude',  value: 8900 },
    { key: 'Gemini',  value: 4200 },
  ];

  return (
    <InkPictureProvider>
      <Box flexDirection="column" padding={1}>
        <Gradient name="mind">
          <Text>🤖 AI AGENT TERMINAL</Text>
        </Gradient>

        <Text>── Token Usage ──</Text>
        <Chart data={usage} type="bar" />

        <Text>── Vision Result ──</Text>
        <Image
          src="output/generated-image.png"
          width={30} height={15}
          alt="AI-generated image"
        />
      </Box>
    </InkPictureProvider>
  );
}

render(<AIDashboard />);
```

## Quick Install

```bash
npm install ink ink-chart ink-gradient ink-picture
```

For TypeScript:

```bash
npm install --save-dev @types/react
```
