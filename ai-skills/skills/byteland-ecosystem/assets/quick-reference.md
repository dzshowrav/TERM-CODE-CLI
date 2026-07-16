# ByteLand Ecosystem — Quick Reference Cheatsheet

## Installation Commands

```bash
# Ink Ecosystem
npm install ink-scroll-view          # Scroll container
npm install ink-scroll-list          # Selectable list
npm install ink-scroll-bar           # Scroll bar
npm install ink-multiline-input      # Multi-line text input
npm install ink-canvas               # Ink in browser

# Tinky Framework
npm install tinky                    # Core framework
npm install tinky-image              # Terminal image rendering

# Layout Engine
npm install taffy-layout             # Taffy WASM layout

# API Client
npm install jules-api-node           # Jules API client
```

## Ink — ScrollView Quick Start

```tsx
const ref = useRef<ScrollViewRef>(null);
useInput((_, key) => {
  if (key.upArrow) ref.current?.scrollBy(-1);
  if (key.downArrow) ref.current?.scrollBy(1);
  if (key.pageUp) ref.current?.scrollBy(-(ref.current?.getViewportHeight() || 1));
  if (key.pageDown) ref.current?.scrollBy(ref.current?.getViewportHeight() || 1);
});
// Terminal resize:
useEffect(() => {
  const h = () => ref.current?.remeasure();
  process.stdout.on("resize", h);
  return () => process.stdout.off("resize", h);
}, []);
<ScrollView ref={ref}>{/* children with keys */}</ScrollView>
```

## Ink — ScrollList Quick Start

```tsx
const [sel, setSel] = useState(0);
const items = ["A", "B", "C"];
useInput((_, key) => {
  if (key.upArrow) setSel(p => Math.max(p - 1, 0));
  if (key.downArrow) setSel(p => Math.min(p + 1, items.length - 1));
});
<ScrollList selectedIndex={sel}>
  {items.map((item, i) => (
    <Text key={i} color={i === sel ? "green" : "white"}>
      {i === sel ? "> " : "  "}{item}
    </Text>
  ))}
</ScrollList>
```

## Ink — MultilineInput Quick Start

```tsx
const [val, setVal] = useState("");
<MultilineInput
  value={val} onChange={setVal}
  rows={3} maxRows={10}
  placeholder="Type here..."
  highlightStyle={{ backgroundColor: "blue" }}
  onSubmit={(v) => console.log("Submitted:", v)}
  keyBindings={{ submit: (key) => key.ctrl && key.return }}
/>
```

## Ink Canvas Quick Start

```tsx
// Vite: add inkCanvasPolyfills() to vite.config.ts
// Container MUST have explicit dimensions
<InkCanvas focused style={{ width: "100%", height: "400px" }}>
  <MyApp />
</InkCanvas>
```

## Tinky — Component Reference

```tsx
// Box — Flexbox + Grid layout
<Box flexDirection="row" gap={2} />
<Box display="grid" gridTemplateColumns="1fr 2fr" />

// Text — Styled text
<Text color="green" bold italic underline>...</Text>

// Hooks
useInput((input, key) => {});
useApp();        // { exit }
useFocus();      // { isFocused }
useFocusManager(); // { focusNext, focusPrevious }

// render
render(<App />, { exitOnCtrlC: true });
const { unmount, waitUntilExit } = render(<App />);
```

## Taffy Layout Quick Start

```ts
await loadTaffy();
const tree = new TaffyTree();
const style = new Style();
style.display = Display.Flex;
style.flexDirection = FlexDirection.Column;
style.size = { width: 300, height: 200 };
const node = tree.newWithChildren(style, children);
tree.computeLayout(node, { width: 300, height: 200 });
const layout = tree.getLayout(node);
// layout.width, layout.height, layout.x, layout.y
```

## Scroll Bar Quick Start

```tsx
// ScrollBarBox (easiest)
<ScrollBarBox height={12} width={40} borderStyle="single"
  contentHeight={100} viewportHeight={10} scrollOffset={offset}>
  <Content />
</ScrollBarBox>

// Standalone ScrollBar (border mode)
<ScrollBar placement="right" style="single"
  contentHeight={100} viewportHeight={20} scrollOffset={offset} />

// Inset mode with autoHide
<ScrollBar placement="inset" style="block" color="cyan"
  contentHeight={100} viewportHeight={20} scrollOffset={offset} autoHide />
```

## Image Rendering

```tsx
<Image src="./logo.png" width={40} resizeMode="contain" />
// Auto-detects: kitty → iterm → sixel → halfblock → braille → ascii
// Override: <Image renderer="halfblock" />
```

## Headless Ghidra Quick Start

```
"Use the headless-ghidra skill to analyze ./sample-target.
Start at P0 intake and stop after each phase gate."
```

## Spec Forge Quick Start

```
"Use $spec-forge to turn this request into an approved spec."
```
