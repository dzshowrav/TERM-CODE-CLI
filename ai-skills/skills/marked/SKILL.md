# Marked

Fast, low-level Markdown compiler in JavaScript/TypeScript. Parses Markdown without caching or blocking. Works in browsers, Node.js, Deno, Bun, and via CLI.

## Installation

```sh
npm install marked
```

## Quick Start

```js
import { marked } from 'marked'

const html = marked.parse('# Hello *world*')
console.log(html) // '<h1>Hello <em>world</em></h1>'
```

## Options

Set via `marked.use()` (global) or per-call in `marked.parse(str, options)`.

| Option    | Type      | Default | Description                                        |
|-----------|-----------|---------|----------------------------------------------------|
| `async`   | `boolean` | `false` | Enable async `walkTokens`; `parse` returns a Promise |
| `breaks`  | `boolean` | `false` | Add `<br>` on single line breaks (requires `gfm`)    |
| `gfm`     | `boolean` | `true`  | GitHub Flavored Markdown                            |
| `pedantic`| `boolean` | `false` | Conform to original `markdown.pl`                   |
| `silent`  | `boolean` | `false` | Return errors as strings instead of throwing        |

```js
marked.use({ gfm: true, breaks: false })
```

## Core API

```js
// Full document (wraps in <p> tags for blocks)
marked.parse(markdownString, options?)

// Inline only (no <p> wrapper)
marked.parseInline(markdownString, options?)

// Direct lexer access (get token AST)
const tokens = marked.lexer(markdownString, options?)
marked.parser(tokens, options?)

// New instances (no shared state)
const lexer = new marked.Lexer(options)
const tokens = lexer.lex(markdownString)

const parser = new marked.Parser(options)
const html = parser.parse(tokens)
```

## Custom Renderer

Override default HTML generation for specific token types:

```js
const renderer = {
  heading(text, level) {
    return `<h${level} class="my-heading">${text}</h${level}>\n`
  }
}

marked.use({ renderer })

marked.parse('# Hello')
// '<h1 class="my-heading">Hello</h1>'
```

## Custom Tokenizer

Override how markdown source is tokenized for specific types:

```js
const tokenizer = {
  codespan(src) {
    const match = /^\$+([^$\n]+?)\$+/.exec(src)
    if (match) {
      return {
        type: 'codespan',
        raw: match[0],
        text: match[1].trim()
      }
    }
    return false // fall back to default
  }
}

marked.use({ tokenizer })
```

## Custom Extensions

Add entirely new syntax. An extension has `name`, `level` (`'block'` or `'inline'`), optional `start(src)`, `tokenizer(src, tokens)`, and `renderer(token)`.

```js
const emojiExtension = {
  name: 'emoji',
  level: 'inline',
  start(src) { return src.indexOf(':') },
  tokenizer(src) {
    const match = /^:([a-z0-9_]+):/.exec(src)
    if (match) {
      return { type: 'emoji', raw: match[0], emoji: match[1] }
    }
  },
  renderer(token) {
    const map = { smile: '😊', laugh: '😂', thumbsup: '👍' }
    return map[token.emoji] || token.raw
  }
}

marked.use({ extensions: [emojiExtension] })
```

Within renderers, access `this.parser.parse(tokens)` for block children and `this.parser.parseInline(tokens)` for inline children. Within tokenizers, access `this.lexer.blockTokens(src)`, `this.lexer.inline(src)`, and `this.lexer.inlineTokens(src)`.

## walkTokens

Traverse and modify every token before rendering:

```js
marked.use({
  walkTokens(token) {
    if (token.type === 'heading') {
      token.depth += 1
    }
  }
})
```

For async operations (e.g., fetching remote content), enable `async: true`:

```js
marked.use({
  async: true,
  async walkTokens(token) {
    if (token.type === 'link') {
      const res = await fetch(token.href)
      token.title = res.ok ? 'valid' : 'invalid'
    }
  }
})

const html = await marked.parse(markdown)
```

## Syntax Highlighting

Use the `marked-highlight` package:

```sh
npm install marked-highlight highlight.js
```

```js
import { marked } from 'marked'
import { markedHighlight } from 'marked-highlight'
import hljs from 'highlight.js'

marked.use(markedHighlight({
  langPrefix: 'hljs language-',
  highlight(code, lang) {
    const language = hljs.getLanguage(lang) ? lang : 'plaintext'
    return hljs.highlight(code, { language }).value
  }
}))
```

## CLI

```sh
cat in.md | marked > out.html
echo "hello *world*" | marked
marked -o out.html -i in.md --gfm
```

Config file: `~/.marked.json`, `~/.marked.js`, or `~/.marked/index.js`.

## Security

Marked does **not** sanitize output HTML. Always sanitize user-provided Markdown:

```js
import DOMPurify from 'dompurify'
const html = DOMPurify.sanitize(marked.parse(userInput))
```

## Hooks (Advanced)

```js
marked.use({
  hooks: {
    preprocess(markdown) { return markdown },
    processAllTokens(tokens) { return tokens },
    postprocess(html) { return html }
  }
})
```

## Token Types

Block: `space`, `code`, `heading`, `table`, `hr`, `blockquote`, `list`, `list_item`, `paragraph`, `html`, `def`, `escape`
Inline: `escape`, `html`, `link`, `image`, `strong`, `em`, `codespan`, `br`, `del`, `text`
