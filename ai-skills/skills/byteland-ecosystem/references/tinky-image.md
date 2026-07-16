# tinky-image — Reference

> **NPM**: `tinky-image` | **Stars**: ⭐1
> **GitHub**: https://github.com/ByteLandTechnology/tinky-image
> **License**: MIT | **Status**: v1.0.0
> **Language**: TypeScript 100%

Terminal image rendering for the Tinky CLI framework.

---

## Entry Points

```tsx
// Node runtime (default)
import { Image } from "tinky-image";
<Image src="./logo.png" width={40} alt="logo" />;

// Browser (when bundled with sharp-web)
import { Image } from "tinky-image/browser";
```

---

## Image Backends (auto-detected, overridable)

| Backend | Protocol | Best For |
|---------|----------|----------|
| `kitty` | Terminal graphics protocol | Kitty-compatible terminals |
| `iterm` | Inline image protocol | iTerm2-compatible terminals |
| `sixel` | Bitmap graphics | Sixel-supporting terminals |
| `halfblock` | Unicode half-block `▄` + ANSI | Full-color, any terminal |
| `braille` | Unicode Braille patterns | Monochrome, high-density |
| `ascii` | Plain character art | Maximum compatibility |

Override with `renderer` prop:
```tsx
<Image src="./logo.png" width={40} renderer="halfblock" />
```

---

## Props

| Prop | Type | Default | Description |
|------|------|---------|-------------|
| `src` | `string \| Uint8Array \| Blob` | — | Image source |
| `width` | `number \| string` | — | Width in terminal cells or `"50%"` |
| `height` | `number \| string` | auto | Height in terminal cells or `"50%"` |
| `renderer` | `Renderer` | auto | Backend override |
| `resizeMode` | `'contain' \| 'cover' \| 'fill' \| 'none'` | `'contain'` | How to fit image in box |
| `alt` | `string` | — | Alt text for accessibility |

### Resize Modes

| Mode | Behavior |
|------|----------|
| `contain` (default) | Scale to fit entirely inside box, preserving aspect ratio |
| `cover` | Scale to fill box, preserving aspect ratio (may crop) |
| `fill` | Stretch to exact width×height, ignoring aspect ratio |
| `none` | Use natural cell size, ignore width/height |

Every mode except `none` clamps to terminal bounds.

---

## Source Types

**Node**: File paths, absolute URLs, data URLs, Uint8Array, Blob
**Browser**: Relative URLs, absolute URLs, blob: URLs, data URLs, Uint8Array, Blob

---

## Complete Example

```tsx
import React from "react";
import { render, Box, Text } from "tinky";
import { Image } from "tinky-image";

const App = () => (
  <Box flexDirection="column" padding={1} alignItems="center">
    <Text bold>Terminal Image Demo</Text>
    <Image src="./logo.png" width={40} resizeMode="contain" />
    <Text dimColor>Rendered with auto-detected backend</Text>
  </Box>
);

render(<App />);
```
