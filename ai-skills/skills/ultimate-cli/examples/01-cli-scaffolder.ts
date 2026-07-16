#!/usr/bin/env node
/**
 * ============================================================
 *  CLI SCAFFOLDER — Complete CLI App with Commander + Inquirer
 * ============================================================
 * Libraries: commander, @inquirer/prompts, chalk, fast-glob
 * 
 * A production-ready CLI that scaffolds new projects.
 * Demonstrates: subcommands, options, arguments, interactive prompts,
 * progress spinners, file generation with templates.
 * 
 * Run: npx ts-node examples/01-cli-scaffolder.ts init my-app
 * ============================================================
 */

import { Command } from 'commander';
import { input, select, checkbox, confirm, password } from '@inquirer/prompts';
import chalk from 'chalk';
import fg from 'fast-glob';
import fs from 'fs/promises';
import path from 'path';

// ─── Types ───────────────────────────────────────────────
interface ScaffoldOptions {
  name: string;
  template: 'react' | 'vue' | 'svelte' | 'vanilla';
  packageManager: 'npm' | 'yarn' | 'pnpm' | 'bun';
  features: string[];
  typescript: boolean;
  git: boolean;
  apiKey?: string;
}

interface TemplateFile {
  path: string;
  content: string;
}

// ─── Templates ───────────────────────────────────────────
const TEMPLATES: Record<string, Record<string, string>> = {
  react: {
    'package.json': JSON.stringify({
      name: '{{name}}',
      version: '1.0.0',
      private: true,
      scripts: {
        dev: 'vite',
        build: 'tsc && vite build',
        preview: 'vite preview',
      },
    }, null, 2),
    'index.html': `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>{{name}}</title>
</head>
<body>
  <div id="root"></div>
  <script type="module" src="/src/main.tsx"></script>
</body>
</html>`,
    'src/main.tsx': `import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);`,
    'src/App.tsx': `import React from 'react';

const App: React.FC = () => {
  return (
    <div>
      <h1>{{name}}</h1>
      <p>Scaffolded with ultimate-cli</p>
    </div>
  );
};

export default App;`,
  },
  vanilla: {
    'index.html': `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>{{name}}</title>
  <link rel="stylesheet" href="/src/style.css" />
</head>
<body>
  <h1>{{name}}</h1>
  <script type="module" src="/src/main.js"></script>
</body>
</html>`,
    'src/main.js': `console.log('{{name}} initialized!');`,
    'src/style.css': `* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: system-ui, sans-serif; padding: 2rem; }`,
  },
};

// ─── Helpers ─────────────────────────────────────────────
function renderTemplate(content: string, vars: Record<string, string>): string {
  return content.replace(/\{\{(\w+)\}\}/g, (_, key) => vars[key] || `{{${key}}}`);
}

async function writeFiles(baseDir: string, files: TemplateFile[]): Promise<void> {
  for (const file of files) {
    const fullPath = path.join(baseDir, file.path);
    await fs.mkdir(path.dirname(fullPath), { recursive: true });
    await fs.writeFile(fullPath, file.content, 'utf-8');
    console.log(chalk.green('  ✓'), chalk.dim(file.path));
  }
}

async function initGit(dir: string): Promise<void> {
  const { execSync } = await import('child_process');
  execSync('git init', { cwd: dir, stdio: 'ignore' });
  await fs.writeFile(path.join(dir, '.gitignore'), `node_modules\ndist\n.env\n`);
}

// ─── Main Scaffold Function ─────────────────────────────
async function scaffold(options: ScaffoldOptions): Promise<void> {
  const projectDir = path.resolve(process.cwd(), options.name);
  const vars = { name: options.name };

  console.log(chalk.bold.cyan('\n  ⚡ Creating project:'), chalk.bold(options.name), '\n');

  // 1. Create directory
  await fs.mkdir(projectDir, { recursive: true });

  // 2. Generate template files
  const template = TEMPLATES[options.template] || TEMPLATES.vanilla;
  const files: TemplateFile[] = Object.entries(template).map(([filePath, content]) => ({
    path: filePath,
    content: renderTemplate(content, vars),
  }));

  // 3. Write files
  console.log(chalk.dim('  Creating files...'));
  await writeFiles(projectDir, files);

  // 4. Initialize git
  if (options.git) {
    console.log(chalk.dim('\n  Initializing git...'));
    await initGit(projectDir);
    console.log(chalk.green('  ✓'), chalk.dim('.git/ initialized'));
  }

  // 5. Success message
  console.log(chalk.bold.green('\n  ✅ Project created successfully!\n'));
  console.log(chalk.dim(`  cd ${options.name}`));
  console.log(chalk.dim(`  ${options.packageManager} install\n`));
}

// ─── CLI Definition ──────────────────────────────────────
const program = new Command()
  .name('scaffolder')
  .description('Scaffold new projects from templates')
  .version('1.0.0');

// === Main command: init ===
program
  .command('init')
  .description('Create a new project')
  .argument('[name]', 'Project name')
  .option('-t, --template <type>', 'Template type (react|vue|svelte|vanilla)')
  .option('--ts, --typescript', 'Use TypeScript', true)
  .option('--no-git', 'Skip git init')
  .option('-p, --pm <manager>', 'Package manager')
  .action(async (name?: string, options?: Record<string, any>) => {
    try {
      // Interactive prompts for missing args
      const projectName = name || await input({
        message: 'Project name:',
        default: 'my-app',
        validate: (v) => /^[a-z0-9-]+$/.test(v) || 'Use lowercase, numbers, hyphens',
      });

      const template = (options?.template || await select({
        message: 'Select template:',
        choices: [
          { name: 'React + Vite', value: 'react', description: 'React with TypeScript and Vite' },
          { name: 'Vue 3', value: 'vue', description: 'Vue 3 Composition API' },
          { name: 'Svelte', value: 'svelte', description: 'Svelte + SvelteKit' },
          { name: 'Vanilla', value: 'vanilla', description: 'Plain HTML/CSS/JS' },
        ],
      })) as ScaffoldOptions['template'];

      const pm = (options?.pm || await select({
        message: 'Package manager:',
        choices: [
          { name: 'npm', value: 'npm' },
          { name: 'yarn', value: 'yarn' },
          { name: 'pnpm', value: 'pnpm' },
          { name: 'bun', value: 'bun' },
        ],
      })) as ScaffoldOptions['packageManager'];

      const features = await checkbox({
        message: 'Additional features:',
        choices: [
          { name: 'ESLint', value: 'eslint', checked: true },
          { name: 'Prettier', value: 'prettier', checked: true },
          { name: 'Husky + lint-staged', value: 'husky' },
          { name: 'Tailwind CSS', value: 'tailwind' },
          { name: 'Vitest', value: 'vitest' },
        ],
      });

      const initGit = options?.git !== false ? await confirm({
        message: 'Initialize git repository?',
        default: true,
      }) : false;

      // Optional API key for private packages
      let apiKey: string | undefined;
      if (features.includes('private-registry')) {
        apiKey = await password({ message: 'Registry API key:' });
      }

      // Scaffold!
      await scaffold({
        name: projectName,
        template,
        packageManager: pm,
        features,
        typescript: options?.typescript ?? true,
        git: initGit,
        apiKey,
      });

    } catch (error: any) {
      if (error.name === 'AbortError') {
        console.log(chalk.yellow('\n  ⏹ Cancelled'));
        process.exit(0);
      }
      console.error(chalk.red('\n  ✖ Error:'), error.message);
      process.exit(1);
    }
  });

// === Utility: list-templates ===
program
  .command('list-templates')
  .description('List available templates')
  .action(() => {
    console.log(chalk.bold('\n  Available templates:\n'));
    console.log(chalk.cyan('  react   '), chalk.dim('React + Vite + TypeScript'));
    console.log(chalk.cyan('  vue     '), chalk.dim('Vue 3 Composition API'));
    console.log(chalk.cyan('  svelte  '), chalk.dim('Svelte + SvelteKit'));
    console.log(chalk.cyan('  vanilla '), chalk.dim('Plain HTML/CSS/JS'));
    console.log();
  });

// === Utility: info ===
program
  .command('info')
  .description('Show CLI info')
  .action(() => {
    console.log(chalk.bold('\n  Scaffolder CLI\n'));
    console.log(chalk.dim('  Version:'), '1.0.0');
    console.log(chalk.dim('  Author:'), 'ultimate-cli');
    console.log(chalk.dim('  Templates:'), 'react, vue, svelte, vanilla');
    console.log();
  });

// === Parse ===
program.parse(process.argv);

// Show help if no args
if (!process.argv.slice(2).length) {
  program.help();
}
