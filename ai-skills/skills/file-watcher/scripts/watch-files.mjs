#!/usr/bin/env node
import chokidar from 'chokidar';
import { execSync } from 'child_process';
import path from 'path';

const [target, promptText] = process.argv.slice(2);
const globIdx = process.argv.indexOf('--glob');
const debounceIdx = process.argv.indexOf('--debounce');
const glob = globIdx !== -1 ? process.argv[globIdx + 1] : undefined;
const debounce = debounceIdx !== -1 ? parseInt(process.argv[debounceIdx + 1], 10) : 500;

if (!target || !promptText) {
  console.error('Usage: node watch-files.mjs <path> <prompt> [--glob <pattern>] [--debounce <ms>]');
  process.exit(1);
}

let timer;
const watcher = chokidar.watch(target, { ignored: glob ? (p) => !p.match(glob) : undefined });

watcher.on('add', (file) => handleChange(file));
watcher.on('change', (file) => handleChange(file));

function handleChange(file) {
  clearTimeout(timer);
  timer = setTimeout(() => {
    const prompt = promptText.replace(/\{\{file\}\}/g, file);
    console.log(`\n[watcher] ${file} changed — running prompt...`);
    try {
      const output = execSync(`claude -p "${prompt.replace(/"/g, '\\"')}" --print`, { encoding: 'utf-8' });
      console.log(output);
    } catch (err) {
      console.error(`[watcher] Error: ${err.message}`);
    }
  }, debounce);
}

console.log(`[watcher] Watching ${target} (debounce: ${debounce}ms)` + (glob ? `, glob: ${glob}` : ''));
