# fast-glob

Fast and efficient glob library for Node.js. Synchronous, Promise, and Stream API.

## Install

```bash
npm install fast-glob
```

**v3.3.3** — latest. Requires Node.js ≥10.10 for modern mode.

## APIs

```js
import fg from 'fast-glob';
```

### Promise (async)

```js
const entries = await fg(['**/*.js', '!**/node_modules/**']);
```

### Synchronous

```js
const entries = fg.sync(['**/*.js', '!**/node_modules/**']);
```

### Stream

```js
const stream = fg.stream(['**/*.js', '!**/node_modules/**']);
stream.on('data', (entry) => console.log(entry));
```

## Pattern Syntax

| Pattern | Meaning |
|---------|---------|
| `*` | Any chars except `/`, excludes dotfiles |
| `**` | Zero or more directories (globstar) |
| `?` | Single char (except `/`) |
| `[abc]` | Any char in set |
| `[!abc]` | Any char NOT in set |
| `{a,b}` | Alternative patterns (brace expansion) |
| `!(pattern)` | Negate (extglob) |
| `?(pattern)` | Zero or one (extglob) |
| `+(pattern)` | One or more (extglob) |
| `@(pattern)` | Exactly one (extglob) |
| `!(pattern)` | Not matching (extglob) |

```js
fg.sync('src/**/*.{js,ts}');
fg.sync('*.+(json|md)');       // extglob
fg.sync(['**/*.js', '!**/node_modules/**']);   // negative patterns
fg.sync('**/*', { ignore: ['node_modules/**', '.git/**'] });
```

Always use forward-slashes in patterns.

## Helpers

```js
fg.isDynamicPattern('*.js');        // true — has wildcards
fg.isDynamicPattern('path/file.js'); // false — static path

fg.escapePath('!abc');             // \\!abc  — escape special chars
fg.convertPathToPattern('C:\\foo'); // C:/foo  — path to safe pattern
```

## Options

### Common

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `cwd` | string | `process.cwd()` | Working directory |
| `deep` | number | `Infinity` | Max recursion depth |
| `ignore` | string[] | `[]` | Patterns to exclude |
| `concurrency` | number | CPU cores | Max concurrent FS reads |
| `followSymbolicLinks` | boolean | `true` | Follow symlinks |
| `suppressErrors` | boolean | `false` | Suppress read errors |
| `throwErrorOnBrokenSymbolicLink` | boolean | `false` | Throw on broken link |
| `fs` | object | — | Custom FS implementation |

### Output Control

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `absolute` | boolean | `false` | Return absolute paths |
| `markDirectories` | boolean | `false` | Append `/` to dirs |
| `objectMode` | boolean | `false` | Return `{ path, dirent }` |
| `onlyDirectories` | boolean | `false` | Dirs only |
| `onlyFiles` | boolean | `true` | Files only |
| `stats` | boolean | `false` | Include `fs.Stats` |
| `unique` | boolean | `true` | Deduplicate entries |

### Matching Control

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `braceExpansion` | boolean | `true` | Enable `{a,b}` |
| `caseSensitiveMatch` | boolean | `true` | Case sensitive |
| `dot` | boolean | `false` | Match dotfiles |
| `extglob` | boolean | `true` | Enable `@()`, `!()`, etc. |
| `globstar` | boolean | `true` | Enable `**` |
| `baseNameMatch` | boolean | `false` | Match basename without dir |

## Examples

```js
// All JS files recursively, absolute paths
const files = await fg('**/*.js', { absolute: true });

// Dirs only, max depth 2
const dirs = fg.sync('**/*', { onlyDirectories: true, deep: 2 });

// Dotfiles + stats
const stats = await fg('**/.*', { dot: true, stats: true });

// Custom cwd
const src = await fg('**/*.ts', { cwd: './src' });

// Object mode with dirent
const objs = fg.sync('*', { objectMode: true });
objs.forEach(o => console.log(o.path, o.dirent.isFile()));

// Stream — large results
const stream = fg.stream('**/*', { onlyFiles: false });
stream.on('data', (path) => processPath(path));
```

## Tips

- **Negative patterns** use `!` prefix: `['**/*', '!**/node_modules/**']`
- **Exclude directories** with `!**/node_modules` or `!**/node_modules/**`
- **Windows**: always use forward-slash in patterns, or `convertPathToPattern()`
- **Environment variable** `UV_THREADPOOL_SIZE` can increase concurrency
- Use `fg.isDynamicPattern()` to check if a pattern needs globbing vs plain stat
