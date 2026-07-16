---
name: terminal-link
description: Create clickable hyperlinks in the terminal using the `terminal-link` library. Falls back to plain text with URL for unsupported terminals. Check support with `terminalLink.isSupported`.
---

# terminal-link

Create clickable links in the terminal. Falls back gracefully for terminals that don't support hyperlinks.

## Install

```sh
npm install terminal-link
```

## Usage

```typescript
import terminalLink from 'terminal-link';

const link = terminalLink('My Website', 'https://sindresorhus.com');
console.log(link);
```

## API

### terminalLink(text, url, options?)

| Option | Type | Description |
|--------|------|-------------|
| `fallback` | function\|boolean | Override default fallback `(text, url) => string`. `false` returns text as-is. |

### terminalLink.isSupported

`boolean` — whether terminal's stdout supports hyperlinks.

### terminalLink.stderr(text, url, options?)

Create a link for stderr.

### terminalLink.stderr.isSupported

Same as stdout variant but for stderr.

## Related

- `supports-hyperlinks` — detect hyperlink support
- `ansi-escapes` — low-level link escape code

## Target Processes

- cli-output-formatting
- interactive-terminal
