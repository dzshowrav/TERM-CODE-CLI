# ink-multiline-input

Learned from: https://github.com/ByteLandTechnology/ink-multiline-input.git

A React component for multi-line text input in [Ink](https://github.com/vadimdemedes/ink) terminal applications.

## Quick Reference

- **Components**: `MultilineInput` (interactive) + `ControlledMultilineInput` (display)
- **Props**: value, onChange, onSubmit, rows/maxRows, highlightStyle, textStyle, placeholder, mask, tabSize, keyBindings, focus, showCursor, useCustomInput
- **Key Bindings**: keyBindings.submit (default Ctrl+Enter), keyBindings.newline (default Enter)
- **Scrolling**: MeasureBox + negative margin offset auto-scroll
- **Navigation**: Column-preserving vertical arrow keys
- **Testing**: Mock MeasureBox, use useCustomInput for logic tests
- **Build**: tsup → ESM, peer deps: ink>=6, react>=19

Full skill: `skill(name="ink-multiline-input")`
