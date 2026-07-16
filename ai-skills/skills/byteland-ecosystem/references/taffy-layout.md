# taffy-layout — Reference

> **NPM**: `taffy-layout` | **Stars**: ⭐12 (most starred ByteLand project)
> **GitHub**: https://github.com/ByteLandTechnology/taffy-layout
> **Website**: https://taffylayout.com
> **License**: MIT | **Status**: v2.0.3 (12 releases, 67 commits)
> **Languages**: TypeScript 61.4%, Rust 38.6%

High-performance **WebAssembly bindings** for the [Taffy](https://github.com/DioxusLabs/taffy) layout engine, bringing CSS Flexbox and Grid layout algorithms to JavaScript with near-native performance.

---

## Architecture

```
Rust (Taffy engine) → WASM (wasm-bindgen) → TypeScript bindings → JS/TS app
```

---

## Quick Start

```ts
import { loadTaffy, TaffyTree, Style, Display, FlexDirection, AlignItems } from "taffy-layout";

// 1. Initialize WASM (must await once at app startup)
await loadTaffy();

// 2. Create a layout tree
const tree = new TaffyTree();

// 3. Define styles
const containerStyle = new Style();
containerStyle.display = Display.Flex;
containerStyle.flexDirection = FlexDirection.Column;
containerStyle.alignItems = AlignItems.Center;
containerStyle.size = { width: 300, height: 200 };
containerStyle.padding = { left: 10, right: 10, top: 10, bottom: 10 };

const childStyle = new Style();
childStyle.flexGrow = 1;
childStyle.width = "100%";
childStyle.height = "auto";

// 4. Create nodes
const child1 = tree.newLeaf(childStyle);
const child2 = tree.newLeaf(childStyle);
const container = tree.newWithChildren(containerStyle, [child1, child2]);

// 5. Compute layout
tree.computeLayout(container, { width: 300, height: 200 });

// 6. Read results
const containerLayout = tree.getLayout(container);
console.log(`Container: ${containerLayout.width}x${containerLayout.height}`);
```

---

## Core API

### `loadTaffy()` — Initialize WASM module

```ts
await loadTaffy();  // Must call before using any other API
```

### `TaffyTree` — Main class

| Method | Description |
|--------|-------------|
| `newLeaf(style)` | Create a leaf node |
| `newWithChildren(style, children)` | Create a node with children |
| `newLeafWithContext(style, context)` | Leaf with custom measurement context |
| `computeLayout(node, availableSpace)` | Compute layout (synchronous) |
| `computeLayoutWithMeasure(node, space, measureFn)` | Layout with custom measurement callback |
| `getLayout(node)` | Read computed layout result |
| `remove(node)` | Remove a node |

### `Style` — Layout configuration

**Size**: Can be set as object or individual properties:
```ts
style.size = { width: 300, height: 200 };          // object form
style.width = 300; style.height = 200;              // individual
style.size = { width: "50%", height: "auto" };      // percentages
```

**Padding**: Can be object or individual:
```ts
style.padding = { left: 10, right: 10, top: 10, bottom: 10 };
style.paddingLeft = 10;  // individual form also works
```

**Display**: `Display.Flex`, `Display.Grid`, `Display.None`, `Display.Block`

**Position**: `Position.Relative`, `Position.Absolute`

### `Layout` — Read-only result

```ts
const layout = tree.getLayout(nodeId);
layout.width;   // number
layout.height;  // number
layout.x;       // number (left offset)
layout.y;       // number (top offset)
```

---

## Advanced Features

### Custom Text Measurement

```ts
const tree = new TaffyTree();
const textNode = tree.newLeafWithContext(style, { text: "Hello!" });

tree.computeLayoutWithMeasure(
  rootNode,
  { width: 800, height: "max-content" },
  (known, available, node, context, style) => {
    if (context?.text) {
      const width = context.text.length * 8;
      const height = 20;
      return { width, height };
    }
    return { width: 0, height: 0 };
  }
);
```

### CSS Grid with Template Areas

```ts
const gridStyle = new Style();
gridStyle.display = Display.Grid;
gridStyle.gridTemplateAreas = [
  { name: "header", rowStart: 1, rowEnd: 2, columnStart: 1, columnEnd: 4 },
  { name: "sidebar", rowStart: 2, rowEnd: 4, columnStart: 1, columnEnd: 2 },
  { name: "main", rowStart: 2, rowEnd: 4, columnStart: 2, columnEnd: 4 },
  { name: "footer", rowStart: 4, rowEnd: 5, columnStart: 1, columnEnd: 4 },
];
```

### Absolute Positioning

```ts
const absoluteStyle = new Style();
absoluteStyle.position = Position.Absolute;
absoluteStyle.inset = { left: 10, top: 10, right: "auto", bottom: "auto" };
absoluteStyle.size = { width: 100, height: 50 };
```

### Block Layout with Replaced Elements

```ts
const imgStyle = new Style();
imgStyle.itemIsReplaced = true;
imgStyle.aspectRatio = 16 / 9;
imgStyle.size = { width: "100%", height: "auto" };
```

### Error Handling

```ts
try {
  const nodeId = tree.newLeaf(style);
} catch (e) {
  if (e instanceof TaffyError) {
    console.error("Layout error:", e.message);
  }
}
```

---

## Browser Support

Chrome 57+, Firefox 52+, Safari 11+, Edge 16+ (any modern browser with WASM support)

---

## Build from Source

```bash
git clone https://github.com/ByteLandTechnology/taffy-layout.git
cd taffy-layout
npm install
npm run build     # Build WASM module
npm test          # Run tests
```
