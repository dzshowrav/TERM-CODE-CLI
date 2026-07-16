# cli-progress

Easy to use progress-bar for command-line/terminal applications in Node.js.

## Install

```bash
npm install cli-progress
```

## SingleBar Mode

```js
const cliProgress = require('cli-progress');
const bar = new cliProgress.SingleBar({}, cliProgress.Presets.shades_classic);
bar.start(200, 0);
bar.update(100);
bar.stop();
```

### Constructor

```js
new cliProgress.SingleBar(options[, preset])
```

### Methods

| Method | Description |
|--------|-------------|
| `start(totalValue, startValue[, payload])` | Start the bar |
| `update([currentValue[, payload]])` | Set progress value (null to skip, payload-only) |
| `increment([delta[, payload]])` | Increase by delta (default +1) |
| `setTotal(totalValue)` | Change total while active |
| `stop()` | Stop and go to next line |
| `updateETA()` | Force ETA recalculation without changing value |

## MultiBar Mode

```js
const multibar = new cliProgress.MultiBar({ clearOnComplete: false, hideCursor: true });
const b1 = multibar.create(200, 0);
const b2 = multibar.create(1000, 0);
b1.increment();
b2.update(20, { filename: "data.txt" });
multibar.stop();
```

### Methods

| Method | Description |
|--------|-------------|
| `create(total, start[, payload[, barOptions]])` | Add a new bar, returns SingleBar |
| `remove(bar)` | Remove a bar |
| `stop()` | Stop all bars |
| `log(msg)` | Output buffered text above bars (needs `\n`) |

## Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `format` | string\|function | `'{bar}'` | Output format with placeholders |
| `fps` | float | 10 | Max update rate |
| `stream` | stream | `process.stderr` | Output stream |
| `stopOnComplete` | boolean | false | Auto-stop on total |
| `clearOnComplete` | boolean | false | Clear bar on stop |
| `barsize` | int | 40 | Bar length in chars |
| `align` | char | `'left'` | Bar position: left, right, center |
| `barCompleteChar` | char | `'='` | Complete indicator |
| `barIncompleteChar` | char | `'-'` | Incomplete indicator |
| `hideCursor` | boolean | false | Hide cursor during operation |
| `linewrap` | boolean | false | Disable line wrapping |
| `gracefulExit` | boolean | false | Stop bars on SIGINT/SIGTERM |
| `etaBuffer` | int | 10 | Updates for ETA calculation (higher = more stable) |
| `etaAsynchronousUpdate` | boolean | false | Trigger ETA update in async render |
| `progressCalculationRelative` | boolean | false | Use startValue as zero-offset |
| `synchronousUpdate` | boolean | true | Trigger redraw during update() |
| `noTTYOutput` | boolean | false | Enable output to non-tty streams |
| `notTTYSchedule` | int | 2000 | Output interval for non-tty (ms) |
| `emptyOnZero` | boolean | false | Show bar as empty when total=0 |
| `forceRedraw` | boolean | false | Redraw every frame even without change |
| `barGlue` | string | `''` | String between complete/incomplete bar |
| `autopadding` | boolean | false | Fixed-width padding for time/percentage |
| `autopaddingChar` | string | `' '` | Padding char (need 3 identical chars) |
| `formatBar` | function | default | Custom bar renderer |
| `formatTime` | function | default | Custom time formatter |
| `formatValue` | function | default | Custom value formatter |

## Format Placeholders

| Placeholder | Description |
|-------------|-------------|
| `{bar}` | The progress bar |
| `{percentage}` | Progress percent (0-100) |
| `{total}` | End value |
| `{value}` | Current value |
| `{eta}` | ETA in seconds |
| `{duration}` | Elapsed seconds |
| `{eta_formatted}` | ETA in human units |
| `{duration_formatted}` | Duration in human units |
| `{<key>}` | Custom payload token |

## Presets

- **legacy** — v1.3.0 style
- **shades-classic** — Unicode background shades
- **shades-grey** — Unicode background shades with grey bar
- **rect** — Unicode Rectangles

```js
const bar = new cliProgress.SingleBar({}, cliProgress.Presets.shades_grey);
```

## Custom Payload

```js
const bar = new cliProgress.SingleBar({
  format: 'progress [{bar}] {percentage}% | Speed: {speed} kbit'
});
bar.start(200, 0, { speed: "N/A" });
bar.update(50, { speed: '125' });
bar.stop();
```

## Custom Formatter

```js
function formatter(options, params, payload) {
  const bar = options.barCompleteString.substr(0, Math.round(params.progress * options.barsize));
  return `# ${payload.task} ${params.value}/${params.total} --[${bar}]--`;
}
new cliProgress.SingleBar({ format: formatter });
```

## Custom Preset

```js
const colors = require('ansi-colors');
module.exports = {
  format: colors.red('{bar}') + ' {percentage}% | {value}/{total}',
  barCompleteChar: '\u2588',
  barIncompleteChar: '\u2591'
};
```

## Events

Extends EventEmitter. See [docs/events.md](https://github.com/npkgz/cli-progress/blob/master/docs/events.md).

## Compatibility

Linux/macOS/container TTY terminals and Windows 10 PowerShell. Non-TTY output for logging daemons.
