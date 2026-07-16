#!/usr/bin/env node
/**
 * ============================================================
 *  PROMPT COMPARISON — Inquirer × Enquirer × Prompts
 * ============================================================
 * Libraries: @inquirer/prompts, enquirer, prompts, chalk
 *
 * A side-by-side comparison and demo of all 3 prompt libraries.
 * Run each prompt type with all libraries and compare behavior.
 *
 * Run: npx ts-node examples/08-prompt-comparison.ts
 *      npx ts-node examples/08-prompt-comparison.ts --lib inquirer
 *      npx ts-node examples/08-prompt-comparison.ts input confirm select
 * ============================================================
 */

import { Command } from 'commander';
import chalk from 'chalk';

// ─── Inquirer Demos ─────────────────────────────────────
async function demoInquirer() {
  console.log(chalk.bold.cyan('\n  📦 @inquirer/prompts Demos\n'));

  const { input, confirm, select, checkbox, password, search } = await import('@inquirer/prompts');

  // Input
  const name = await input({
    message: 'What is your name?',
    default: 'Guest',
    validate: (v) => v.length >= 2 || 'At least 2 characters',
  });
  console.log(chalk.green(`  → ${name}\n`));

  // Password
  const secret = await password({
    message: 'Enter a secret:',
    mask: true,
  });
  console.log(chalk.green(`  → (${secret.length} chars hidden)\n`));

  // Confirm
  const agreed = await confirm({
    message: 'Do you agree to the terms?',
    default: false,
  });
  console.log(chalk.green(`  → ${agreed ? 'Yes' : 'No'}\n`));

  // Select
  const color = await select({
    message: 'Pick a color:',
    choices: [
      { name: 'Red', value: 'red', description: 'Like fire' },
      { name: 'Blue', value: 'blue', description: 'Like sky' },
      { name: 'Green', value: 'green', description: 'Like nature' },
      { name: 'Purple', value: 'purple', description: 'Like royalty' },
    ],
  });
  console.log(chalk.green(`  → ${color}\n`));

  // Checkbox
  const toppings = await checkbox({
    message: 'Select pizza toppings:',
    choices: [
      { name: 'Pepperoni', value: 'pepperoni', checked: true },
      { name: 'Mushrooms', value: 'mushrooms' },
      { name: 'Olives', value: 'olives', disabled: 'Out of stock' },
      { name: 'Extra Cheese', value: 'cheese', checked: true },
    ],
    required: true,
    validate: (selected) => selected.length >= 1 || 'Select at least 1',
  });
  console.log(chalk.green(`  → ${toppings.join(', ')}\n`));

  // Search
  try {
    const country = await search({
      message: 'Search for a country:',
      source: async (input) => {
        if (!input) return [];
        const countries = ['Bangladesh', 'India', 'Japan', 'Canada', 'Brazil', 'Australia',
          'Germany', 'France', 'Italy', 'Spain', 'UK', 'USA', 'China'];
        return countries
          .filter(c => c.toLowerCase().includes(input.toLowerCase()))
          .map(c => ({ name: c, value: c }));
      },
    });
    console.log(chalk.green(`  → ${country}\n`));
  } catch {}

  return { name, color, toppings };
}

// ─── Enquirer Demos ─────────────────────────────────────
async function demoEnquirer() {
  console.log(chalk.bold.cyan('\n  📦 Enquirer Demos\n'));

  const Enquirer = (await import('enquirer')).default;

  // Input
  const { name } = await Enquirer.prompt({
    type: 'input',
    name: 'name',
    message: 'What is your name?',
    initial: 'Guest',
    validate: (v: string) => v.length >= 2 || 'At least 2 characters',
  });
  console.log(chalk.green(`  → ${name}\n`));

  // Password
  const { secret } = await Enquirer.prompt({
    type: 'password',
    name: 'secret',
    message: 'Enter a secret:',
    mask: '*',
  });
  console.log(chalk.green(`  → (${secret.length} chars)\n`));

  // Toggle (unique to Enquirer)
  const { active } = await Enquirer.prompt({
    type: 'toggle',
    name: 'active',
    message: 'Enable dark mode?',
    enabled: '🌙 Dark',
    disabled: '☀️ Light',
    initial: true,
  });
  console.log(chalk.green(`  → ${active ? 'Dark mode' : 'Light mode'}\n`));

  // Select
  const { color } = await Enquirer.prompt({
    type: 'select',
    name: 'color',
    message: 'Pick a color:',
    choices: ['Red', 'Blue', 'Green', 'Purple'].map(c => ({
      name: c,
      hint: c === 'Blue' ? '(Recommended!)' : undefined,
    })),
  });
  console.log(chalk.green(`  → ${color}\n`));

  // MultiSelect
  const { items } = await Enquirer.prompt({
    type: 'multiselect',
    name: 'items',
    message: 'Select features:',
    choices: [
      { name: 'Autosave', value: 'autosave' },
      { name: 'Notifications', value: 'notifications' },
      { name: 'Dark Mode', value: 'darkmode', disabled: true },
    ],
  });
  console.log(chalk.green(`  → ${items.join(', ')}\n`));

  // AutoComplete (unique to Enquirer)
  const { framework } = await Enquirer.prompt({
    type: 'autocomplete',
    name: 'framework',
    message: 'Search framework:',
    choices: ['React', 'Vue', 'Svelte', 'Angular', 'Next.js', 'Nuxt', 'SvelteKit'],
    limit: 3,
  });
  console.log(chalk.green(`  → ${framework}\n`));

  // Form (unique to Enquirer)
  const { user } = await Enquirer.prompt({
    type: 'form',
    name: 'user',
    message: 'Enter user details:',
    choices: [
      { name: 'username', message: 'Username', initial: 'john' },
      { name: 'email', message: 'Email' },
      { name: 'role', message: 'Role', initial: 'developer' },
    ],
    validate: (values: any) => {
      if (!values.email?.includes('@')) return 'Valid email required';
      return true;
    },
  });
  console.log(chalk.green(`  → ${user.username} <${user.email}> (${user.role})\n`));

  // Sort (unique to Enquirer)
  const { order } = await Enquirer.prompt({
    type: 'sort',
    name: 'order',
    message: 'Sort your priorities:',
    choices: ['Feature A', 'Feature B', 'Feature C'],
    numbered: true,
  });
  console.log(chalk.green(`  → ${order.join(' > ')}\n`));
}

// ─── Prompts Demos ──────────────────────────────────────
async function demoPrompts() {
  console.log(chalk.bold.cyan('\n  📦 prompts Demos\n'));

  const prompts = (await import('prompts')).default;

  const response = await prompts([
    {
      type: 'text',
      name: 'name',
      message: 'What is your name?',
      initial: 'Guest',
      validate: (v: string) => v.length >= 2 ? true : 'At least 2 chars',
    },
    {
      type: 'password',
      name: 'secret',
      message: 'Enter a secret:',
    },
    {
      type: 'number',
      name: 'age',
      message: 'How old are you?',
      initial: 25,
      min: 0,
      max: 150,
      validate: (v: number) => v >= 18 ? true : 'Must be 18+',
    },
    {
      type: 'confirm',
      name: 'confirmed',
      message: 'Can you confirm?',
      initial: true,
    },
    {
      type: 'toggle',
      name: 'enabled',
      message: 'Toggle this?',
      initial: true,
      active: 'yes',
      inactive: 'no',
    },
    {
      type: 'select',
      name: 'color',
      message: 'Pick a color',
      choices: [
        { title: 'Red', value: '#ff0000', description: 'Fire' },
        { title: 'Green', value: '#00ff00', description: 'Nature' },
        { title: 'Blue', value: '#0000ff', description: 'Sky', disabled: true },
      ],
    },
    {
      type: 'multiselect',
      name: 'tags',
      message: 'Select tags',
      choices: [
        { title: 'JavaScript', value: 'js', selected: true },
        { title: 'TypeScript', value: 'ts', selected: true },
        { title: 'Python', value: 'py' },
        { title: 'Rust', value: 'rs' },
      ],
      min: 1,
      max: 3,
      hint: 'Select 1-3 options',
    },
    {
      type: 'list',
      name: 'keywords',
      message: 'Enter keywords (comma-separated):',
      initial: 'cli, terminal, ui',
      separator: ',',
    },
    {
      type: 'date',
      name: 'birthday',
      message: 'Pick your birthday',
      initial: new Date(1990, 0, 1),
    },
    {
      type: 'autocomplete',
      name: 'language',
      message: 'Search language:',
      choices: [
        { title: 'TypeScript', value: 'ts' },
        { title: 'JavaScript', value: 'js' },
        { title: 'Python', value: 'py' },
        { title: 'Rust', value: 'rs' },
        { title: 'Go', value: 'go' },
      ].map(c => ({ ...c, title: `${c.title} (${c.value})` })),
    },
  ]);

  console.log(chalk.green(`\n  → All answers: ${JSON.stringify(response, null, 2)}\n`));
}

// ─── Comparison Table ───────────────────────────────────
function showComparison() {
  console.log(chalk.bold.cyan('\n  📊 Prompt Library Comparison\n'));

  const data = [
    { feature: 'Text Input', inquirer: '✅', enquirer: '✅', prompts: '✅' },
    { feature: 'Password', inquirer: '✅', enquirer: '✅', prompts: '✅' },
    { feature: 'Confirm (Yes/No)', inquirer: '✅', enquirer: '✅ (Toggle)', prompts: '✅' },
    { feature: 'Select (Single)', inquirer: '✅', enquirer: '✅', prompts: '✅' },
    { feature: 'Multi-Select', inquirer: '✅', enquirer: '✅', prompts: '✅' },
    { feature: 'AutoComplete', inquirer: '✅ (Search)', enquirer: '✅', prompts: '✅' },
    { feature: 'Search/Filter', inquirer: '✅', enquirer: '✅', prompts: '✅' },
    { feature: 'Number Input', inquirer: '❌', enquirer: '✅', prompts: '✅' },
    { feature: 'Form (multi-field)', inquirer: '❌', enquirer: '✅', prompts: '❌' },
    { feature: 'Toggle Switch', inquirer: '❌', enquirer: '✅', prompts: '✅' },
    { feature: 'Sort / Reorder', inquirer: '❌', enquirer: '✅', prompts: '❌' },
    { feature: 'Date Picker', inquirer: '❌', enquirer: '❌', prompts: '✅' },
    { feature: 'List (CSV)', inquirer: '❌', enquirer: '✅', prompts: '✅' },
    { feature: 'Quiz', inquirer: '❌', enquirer: '✅', prompts: '❌' },
    { feature: 'Survey', inquirer: '❌', enquirer: '✅', prompts: '❌' },
    { feature: 'Scale/Rating', inquirer: '❌', enquirer: '✅', prompts: '❌' },
    { feature: 'Raw List', inquirer: '✅', enquirer: '❌', prompts: '❌' },
    { feature: 'Expand', inquirer: '✅', enquirer: '❌', prompts: '❌' },
    { feature: 'Editor', inquirer: '✅', enquirer: '❌', prompts: '❌' },
    { feature: 'Invisible Input', inquirer: '❌', enquirer: '✅', prompts: '✅' },
    { feature: 'Snippet/Code', inquirer: '❌', enquirer: '✅', prompts: '❌' },
    { feature: 'npm Weekly Downloads', inquirer: '~28.5M', enquirer: '~26.7M', prompts: '~45.2M' },
    { feature: 'Bundle Size', inquirer: '~40KB', enquirer: '~78KB', prompts: '~30KB' },
    { feature: 'Dependencies', inquirer: '3', enquirer: '7', prompts: '2' },
    { feature: 'TypeScript', inquirer: '✅ Native', enquirer: '✅', prompts: '✅' },
    { feature: 'Custom Prompts', inquirer: '✅ (createPrompt)', enquirer: '✅ (register)', prompts: '❌' },
  ];

  const headers = ['Feature', 'Inquirer', 'Enquirer', 'Prompts'];
  const rows = data.map(d => [d.feature, d.inquirer, d.enquirer, d.prompts]);

  // Simple table
  console.log(`  ${chalk.bold(headers[0].padEnd(25))} ${chalk.bold(headers[1].padEnd(14))} ${chalk.bold(headers[2].padEnd(14))} ${chalk.bold(headers[3])}`);
  console.log(chalk.dim(`  ${'-'.repeat(70)}`));

  for (const row of rows) {
    console.log(`  ${row[0].padEnd(25)} ${String(row[1]).padEnd(14)} ${String(row[2]).padEnd(14)} ${row[3]}`);
  }
  console.log();

  // Recommendation
  console.log(chalk.bold('\n  🏆 Recommendations:\n'));
  console.log(`  ${chalk.cyan('Inquirer')}  — Best for: Production CLIs, TypeScript, custom prompts`);
  console.log(`  ${chalk.cyan('Enquirer')}  — Best for: Feature-rich prompts (forms, sort, survey, quiz)`);
  console.log(`  ${chalk.cyan('Prompts')}   — Best for: Small bundles, date picker, minimal deps`);
  console.log();
}

// ─── CLI ────────────────────────────────────────────────
const program = new Command()
  .name('prompt-compare')
  .description('Compare Inquirer, Enquirer, and Prompts')
  .version('1.0.0');

program
  .command('all')
  .description('Run all prompt library demos')
  .action(async () => {
    await demoInquirer();
    await demoEnquirer();
    await demoPrompts();
    showComparison();
  });

program
  .command('inquirer')
  .description('Run Inquirer demos')
  .action(demoInquirer);

program
  .command('enquirer')
  .description('Run Enquirer demos')
  .action(demoEnquirer);

program
  .command('prompts')
  .description('Run Prompts demos')
  .action(demoPrompts);

program
  .command('compare')
  .description('Show comparison table')
  .action(showComparison);

program.parse(process.argv);
