# cosmiconfig

Hierarchical configuration loading for Node.js. Searches up the directory tree for config in JSON, YAML, JS, TS, and more.

## Install

```bash
npm install cosmiconfig
```

v9.x — ESM. Requires Node ≥18.

## Core API

```js
import { cosmiconfig, cosmiconfigSync } from 'cosmiconfig';
```

### `cosmiconfig(moduleName, options?)`

Async explorer. Walks up directories from `searchFrom` looking for config.

```js
const explorer = cosmiconfig('myapp');
const result = await explorer.search();
// { config: {...}, filepath: '/path/to/.myapprc.json', isEmpty: false }
```

### `cosmiconfigSync(moduleName, options?)`

Synchronous version.

```js
const explorer = cosmiconfigSync('myapp');
const result = explorer.search();
```

## Methods

### `explorer.search(searchFrom?)`

Walk up directory tree from `searchFrom` (default: `process.cwd()`), checking each `searchPlace` in each directory. Stops at first non-empty match.

```js
const result = await explorer.search('/home/user/projects/my-app/src');
```

Returns `CosmiconfigResult | null`:

```ts
interface CosmiconfigResult {
  config: unknown;       // Parsed config content
  filepath: string;      // Absolute path to found config file
  isEmpty?: boolean;     // True if file was empty
}
```

### `explorer.load(filepath)`

Load a specific file directly (no directory walk).

```js
const result = await explorer.load('/path/to/.myapprc.json');
```

### Cache clearing

```js
explorer.clearSearchCache();  // Clear search location cache
explorer.clearLoadCache();    // Clear file content cache
explorer.clearCaches();       // Clear both
```

## Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `packageProp` | `string \| string[]` | `moduleName` | Property key in `package.json`/`package.yaml` |
| `searchPlaces` | `string[]` | auto-generated | File paths to check per directory |
| `searchStrategy` | `'none' \| 'project' \| 'global'` | depends on `stopDir` | Directory traversal strategy |
| `stopDir` | `string` | home dir | Where upward search stops (`'global'` strategy) |
| `cache` | `boolean` | `true` | Cache search and load results |
| `ignoreEmptySearchPlaces` | `boolean` | `true` | Skip empty files, continue searching |
| `loaders` | `Loaders` | `defaultLoaders` | Custom file format parsers |
| `transform` | `Transform` | identity | Transform loaded config |
| `mergeImportArrays` | `boolean` | `true` | Merge arrays from `$import` |
| `mergeSearchPlaces` | `boolean` | `true` | Merge with default search places |

### searchStrategy

- `'none'` — scan only the exact `searchFrom` directory
- `'project'` — walk up until a `package.json` / `.git` / `stopDir` is found
- `'global'` — walk up from `searchFrom` to `stopDir` (default: home directory)

## Default Search Places

For module name `'myapp'`:

```
package.json
.myapprc
.myapprc.json
.myapprc.yaml
.myapprc.yml
.myapprc.js
.myapprc.ts         (sync: excluded)
.myapprc.cjs
.myapprc.mjs        (async only)
.config/myapprc
.config/myapprc.json
.config/myapprc.yaml
.config/myapprc.yml
.config/myapprc.js
.config/myapprc.ts  (sync: excluded)
.config/myapprc.cjs
.config/myapprc.mjs (async only)
myapp.config.js
myapp.config.ts     (sync: excluded)
myapp.config.cjs
myapp.config.mjs    (async only)
```

## Default Loaders

```js
import { defaultLoaders, defaultLoadersSync } from 'cosmiconfig';
```

| Ext | Async | Sync |
|-----|-------|------|
| `.json` | `JSON.parse` | `JSON.parse` |
| `.yaml` / `.yml` | `yaml` | `yaml` |
| `.js` | ESM `import()` | `require()` |
| `.cjs` | `require()` | `require()` |
| `.mjs` | ESM `import()` | ❌ not supported |
| `.ts` | `jiti` or `tsx` | `jiti` or `tsx` |
| `noExt` (extensionless) | ESM `import()` | JS `require()` |

## Custom Loaders

```js
cosmiconfig('myapp', {
  loaders: {
    '.toml': (filepath, content) => parseToml(content),
    '.json5': json5Loader,
  },
});
```

Loader signature: `(filepath: string, content: string) => unknown`

For sync: `(filepath: string, content: string) => unknown` (must not return a Promise).

## Transform

```js
cosmiconfig('myapp', {
  transform: async (result) => {
    if (!result || result.isEmpty) return result;
    return {
      ...result,
      config: await validateSchema(result.config),
    };
  },
});
```

Transform receives `CosmiconfigResult | null`, must return the same shape. Returning `null` makes it as if no config was found.

## packageProp

Nested property access with array notation:

```js
cosmiconfig('myapp', { packageProp: 'myapp' });
// Looks in: package.json -> { myapp: ... }

cosmiconfig('myapp', { packageProp: ['config', 'myapp'] });
// Looks in: package.json -> { config: { myapp: ... } }
```

## $import Directive

Config files can import from other files:

```json
{
  "$import": "./base-config.json",
  "extends": "./shared-settings.json"
}
```

- Relative paths resolve from the importing file's directory
- Arrays merge by default (set `mergeImportArrays: false` to replace)
- Works in JSON, YAML, and JS configs

## Example: Setup Pattern

```js
import { cosmiconfig, defaultLoaders } from 'cosmiconfig';

const explorer = cosmiconfig('myapp', {
  searchPlaces: [
    'package.json',
    '.myapprc',
    '.myapprc.json',
    'myapp.config.js',
    'myapp.config.ts',
  ],
  searchStrategy: 'project',
  stopDir: process.cwd(),
  loaders: {
    ...defaultLoaders,
  },
  transform: async (result) => {
    if (!result || result.isEmpty) return result;
    return { ...result, config: Object.freeze(result.config) };
  },
  cache: true,
  ignoreEmptySearchPlaces: true,
  mergeImportArrays: true,
  mergeSearchPlaces: true,
});

const { config, filepath } = (await explorer.search()) ?? {
  config: getDefaults(),
  filepath: null,
};
```

## Example: Multiple Configs

```js
const [db, app] = await Promise.all([
  cosmiconfig('database').search(),
  cosmiconfig('myapp').search(),
]);
```

## Notes

- Cosmiconfig handles `package.yaml` in addition to `package.json`
- JS/TS configs can export an object (`module.exports = {}` or `export default {}`)
- Use `explorer.load()` for explicit paths, `explorer.search()` for auto-discovery
- Cache is per-explorer instance; create new instances for independent searches
- For TypeScript loading at runtime, cosmiconfig internally uses `jiti` (async) or `jiti/register` (sync); external `cosmiconfig-typescript-loader` package is also available
