# Example: Complete ScrollView App

```tsx
import React, { useRef, useEffect, useState } from "react";
import { render, Text, Box, useInput, useStdout } from "ink";
import { ScrollView, ScrollViewRef } from "ink-scroll-view";
import { ScrollBar } from "ink-scroll-bar";

const ITEMS = Array.from({ length: 100 }, (_, i) => ({
  id: i,
  text: `Item ${i + 1}`,
  desc: `Description for item ${i + 1}`,
}));

const ScrollableBrowser = () => {
  const scrollRef = useRef<ScrollViewRef>(null);
  const { stdout } = useStdout();
  const [scrollOffset, setScrollOffset] = useState(0);
  const [viewportHeight, setViewportHeight] = useState(0);
  const [contentHeight, setContentHeight] = useState(0);

  // Handle terminal resize
  useEffect(() => {
    const handle = () => scrollRef.current?.remeasure();
    stdout?.on("resize", handle);
    return () => stdout?.off("resize", handle);
  }, [stdout]);

  // Global keyboard navigation
  useInput((input, key) => {
    if (key.upArrow) scrollRef.current?.scrollBy(-1);
    if (key.downArrow) scrollRef.current?.scrollBy(1);
    if (key.pageUp) {
      const h = scrollRef.current?.getViewportHeight() || 1;
      scrollRef.current?.scrollBy(-h);
    }
    if (key.pageDown) {
      const h = scrollRef.current?.getViewportHeight() || 1;
      scrollRef.current?.scrollBy(h);
    }
    if (key.home) scrollRef.current?.scrollToTop();
    if (key.end) scrollRef.current?.scrollToBottom();
    if (input === "j") scrollRef.current?.scrollBy(1);
    if (input === "k") scrollRef.current?.scrollBy(-1);
    if (input === "q") process.exit(0);
  });

  return (
    <Box flexDirection="row" height={20}>
      {/* Scrollable content area */}
      <Box flexDirection="column" flexGrow={1} borderStyle="single">
        <Box height={18} flexDirection="column">
          <ScrollView
            ref={scrollRef}
            onScroll={(offset) => setScrollOffset(offset)}
            onViewportSizeChange={(size) => setViewportHeight(size.height)}
            onContentHeightChange={(h) => setContentHeight(h)}
          >
            {ITEMS.map((item) => (
              <Box key={item.id} flexDirection="column" marginBottom={1}>
                <Text bold color="cyan">{item.text}</Text>
                <Text dimColor>{item.desc}</Text>
              </Box>
            ))}
          </ScrollView>
        </Box>
        {/* Status bar */}
        <Box borderStyle="single" borderTop={true}>
          <Text dimColor>
            Line {Math.round(scrollOffset + 1)}-{Math.round(scrollOffset + viewportHeight)}
            {' '}/ {Math.round(contentHeight)} | ↑↓ pgup/pgdn home/end j/k
          </Text>
        </Box>
      </Box>

      {/* Scroll bar */}
      <ScrollBar
        placement="inset"
        style="block"
        color="cyan"
        contentHeight={contentHeight}
        viewportHeight={viewportHeight}
        scrollOffset={scrollOffset}
        autoHide
      />
    </Box>
  );
};

render(<ScrollableBrowser />);
```
