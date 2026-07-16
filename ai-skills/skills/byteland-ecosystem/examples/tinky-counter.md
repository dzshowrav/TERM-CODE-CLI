# Example: Tinky Counter with Grid Layout

```tsx
import React, { useState, useRef, useEffect } from "react";
import {
  render,
  Box,
  Text,
  useInput,
  useApp,
  useFocus,
  useFocusManager,
  measureElement,
} from "tinky";

// A focusable button component
function Button({
  label,
  onPress,
  color,
}: {
  label: string;
  onPress: () => void;
  color: string;
}) {
  const { isFocused } = useFocus();
  const ref = useRef(null);

  useInput((input, key) => {
    if (key.return && isFocused) onPress();
    if (input === " " && isFocused) onPress();
  });

  return (
    <Box
      ref={ref}
      borderStyle={isFocused ? "bold" : "single"}
      borderColor={isFocused ? "green" : color}
      paddingX={2}
      paddingY={1}
    >
      <Text color={isFocused ? "green" : color}>
        {isFocused ? "▶ " : "  "}{label}
      </Text>
    </Box>
  );
}

function CounterApp() {
  const [count, setCount] = useState(0);
  const { exit } = useApp();
  const { focusNext, focusPrevious } = useFocusManager();

  useInput((input, key) => {
    if (key.tab) focusNext();
    if (key.shift && key.tab) focusPrevious();
    if (input === "q") exit();
  });

  return (
    <Box
      display="grid"
      gridTemplateColumns="1fr 2fr 1fr"
      gridTemplateRows="auto auto auto"
      gap={1}
      padding={2}
      borderStyle="round"
      borderColor="cyan"
      width={60}
    >
      {/* Header: spans all columns */}
      <Box gridColumn={{ start: 1, end: 4 }} justifyContent="center">
        <Text bold color="yellow" fontSize="large">
          Tinky Counter
        </Text>
      </Box>

      {/* Left: Decrement button */}
      <Box justifyContent="center" alignItems="center">
        <Button
          label="- Decrement"
          color="red"
          onPress={() => setCount(c => c - 1)}
        />
      </Box>

      {/* Center: Count display */}
      <Box justifyContent="center" alignItems="center">
        <Box
          borderStyle="double"
          borderColor="green"
          paddingX={3}
          paddingY={1}
        >
          <Text bold color="green" fontSize="xlarge">
            {count}
          </Text>
        </Box>
      </Box>

      {/* Right: Increment button */}
      <Box justifyContent="center" alignItems="center">
        <Button
          label="+ Increment"
          color="green"
          onPress={() => setCount(c => c + 1)}
        />
      </Box>

      {/* Reset button: spans 3 columns */}
      <Box gridColumn={{ start: 1, end: 4 }} justifyContent="center">
        <Button
          label="Reset"
          color="magenta"
          onPress={() => setCount(0)}
        />
      </Box>

      {/* Footer: spans 3 columns */}
      <Box gridColumn={{ start: 1, end: 4 }} justifyContent="center">
        <Text dimColor>
          Tab/Shift+Tab to navigate • Enter/Space to press • q to quit
        </Text>
      </Box>
    </Box>
  );
}

render(<CounterApp />);
```
