---
name: terminal-image
description: Display images in the terminal using the `terminal-image` library. Supports PNG, JPEG, and animated GIFs. Uses native terminal graphics protocols (iTerm2, Kitty, WezTerm) with ANSI block character fallback.
---

# terminal-image

Display images in the terminal. Supports full resolution in iTerm2/Kitty/WezTerm, ANSI block character fallback in other terminals.

## Install

```sh
npm install terminal-image
```

## Usage

```typescript
import terminalImage from 'terminal-image';

// Display a file
console.log(await terminalImage.file('unicorn.jpg'));

// With sizing
console.log(await terminalImage.file('unicorn.jpg', {
  width: '50%',
  height: '50%',
}));

// Animated GIF
const stopAnimation = await terminalImage.gifFile('animation.gif');
// Call stopAnimation() to stop
```

## API

| Method | Description |
|--------|-------------|
| `terminalImage.buffer(imageBuffer, options?)` | Display from buffer |
| `terminalImage.file(filePath, options?)` | Display from file path |
| `terminalImage.gifBuffer(imageBuffer, options?)` | Display animated GIF from buffer |
| `terminalImage.gifFile(filePath, options?)` | Display animated GIF from file |

### Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `width` | string\|number | — | Percentage `'50%'` or columns (number) |
| `height` | string\|number | — | Percentage `'50%'` or rows (number) |
| `preserveAspectRatio` | boolean | true | Maintain aspect ratio |
| `preferNativeRender` | boolean | true | Prefer native terminal protocols |
| `maximumFrameRate` | number | 30 | Max FPS for GIF (ignored in iTerm) |

## Target Processes

- cli-image-display
- terminal-graphics
