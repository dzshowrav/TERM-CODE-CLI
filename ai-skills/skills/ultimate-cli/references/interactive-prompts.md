# Interactive Prompt Libraries — Complete A-to-Z Reference

---

## 1. Inquirer (@inquirer/prompts)
**GitHub**: https://github.com/SBoudrias/Inquirer.js | **Stars**: 20K+
**npm**: `@inquirer/prompts` | **Weekly**: ~28.5M | **License**: MIT

### 1.1 Installation
```bash
npm install @inquirer/prompts
```

### 1.2 Complete API — All Prompt Types

#### input()
```typescript
import { input } from '@inquirer/prompts';

const answer: string = await input({
  message: 'Enter your name',
  default?: 'User',
  required?: boolean,           // default: false
  validate?: (value: string) => boolean | string | Promise<boolean | string>,
  transformer?: (value: string, {isFinal: boolean}) => string,
  filter?: (value: string) => string,
  theme?: { prefix?: string; spinner?: object; style?: object },
});
```

#### password()
```typescript
import { password } from '@inquirer/prompts';

const answer: string = await password({
  message: 'Enter password',
  mask?: boolean | string,      // true = '*', or any character
  validate?: (value: string) => boolean | string,
  theme?: { prefix?: string; style?: object },
});
```

#### select()
```typescript
import { select } from '@inquirer/prompts';

const answer: string = await select({
  message: 'Pick an option',
  choices: [
    { name: 'Option A', value: 'a', description?: 'Description shown on hover' },
    { name: 'Option B', value: 'b', disabled?: boolean | string },
    new inquirer.Separator(),                               // visual divider
    { name: 'Option C', value: 'c', key?: 'c' },           // keyboard shortcut
  ],
  default?: string,                  // default value
  loop?: boolean,                     // loop list (default: true)
  pageSize?: number,                  // visible items (default: 7)
  theme?: object,
});
```

#### checkbox() (multi-select)
```typescript
import { checkbox } from '@inquirer/prompts';

const answer: string[] = await checkbox({
  message: 'Select options',
  choices: [
    { name: 'Option A', value: 'a', checked?: boolean },
    { name: 'Option B', value: 'b', disabled?: boolean | string },
  ],
  required?: boolean,         // require at least 1
  loop?: boolean,
  pageSize?: number,
  validate?: (choices: string[]) => boolean | string,
  theme?: object,
});
```

#### confirm()
```typescript
import { confirm } from '@inquirer/prompts';

const answer: boolean = await confirm({
  message: 'Continue?',
  default?: boolean,          // default: true
  theme?: object,
});
```

#### search()
```typescript
import { search } from '@inquirer/prompts';

const answer: string = await search({
  message: 'Search items',
  source: async (input: string | undefined, { signal }: { signal: AbortSignal }) => {
    // Return choices based on search input
    if (!input) return [];
    const results = await searchAPI(input, { signal });
    return results.map(r => ({ name: r.name, value: r.value }));
  },
  pageSize?: number,
  validate?: (value: string) => boolean | string,
  theme?: object,
});
```

#### rawlist()
```typescript
import { rawlist } from '@inquirer/prompts';

const answer: string = await rawlist({
  message: 'Pick an option',
  choices: [
    { name: 'Option A', value: 'a' },
    { name: 'Option B', value: 'b' },
  ],
  theme?: object,
});
```

#### editor()
```typescript
import { editor } from '@inquirer/prompts';

const answer: string = await editor({
  message: 'Edit content',
  default?: string,
  validate?: (text: string) => boolean | string,
  postfix?: string,           // file extension for editor (default: '.txt')
  waitForUseInput?: boolean,  // wait for user input before opening editor
  theme?: object,
});
```

#### expand()
```typescript
import { expand } from '@inquirer/prompts';

const answer: string = await expand({
  message: 'Choose an option',
  choices: [
    { key: 'y', name: 'Yes', value: 'yes' },
    { key: 'n', name: 'No', value: 'no' },
  ],
  theme?: object,
});
```

### 1.3 Custom Theme
```typescript
const answer = await input({
  message: 'Name',
  theme: {
    prefix: '?',              // prefix symbol
    spinner: { ... },         // spinner animation
    style: {
      answer: (text: string) => text.green,
      message: (text: string) => text.bold.cyan,
      error: (text: string) => text.red,
      defaultAnswer: (text: string) => text.dim,
      highlight: (text: string) => text.cyan,
    },
  },
});
```

### 1.4 AbortSignal Support
```typescript
const controller = new AbortController();
setTimeout(() => controller.abort(), 10000);

try {
  const answer = await input({
    message: 'Enter quickly...',
  }, { signal: controller.signal });
} catch (err: any) {
  if (err.name === 'AbortError') console.log('Timed out');
}
```

### 1.5 Custom Inquirer Prompt
```bash
npm install @inquirer/core @inquirer/type
```

```typescript
import { createPrompt, useState, useKeypress } from '@inquirer/core';

type MyConfig = {
  message: string;
  choices: string[];
};

export default createPrompt<string, MyConfig>((config, done) => {
  const [status, setStatus] = useState<string>('pending');
  const [value, setValue] = useState<string>('');

  useKeypress((key, rl) => {
    if (key.name === 'enter') {
      setStatus('done');
      done(value);
    } else {
      setValue(rl.line);
    }
  });

  return `? ${config.message} ${value}`;
});
```

---

## 2. Enquirer (enquirer/enquirer)
**GitHub**: https://github.com/enquirer/enquirer | **Stars**: 7.5K+
**npm**: `enquirer` | **Weekly**: ~26.7M | **License**: MIT

### 2.1 Installation
```bash
npm install enquirer
```

### 2.2 Complete API — All 17 Prompt Types

#### AutoComplete — Filterable Select
```typescript
import Enquirer from 'enquirer';

const response = await Enquirer.prompt({
  type: 'autocomplete',
  name: 'framework',
  message: 'Select framework',
  choices: ['React', 'Vue', 'Svelte', 'Angular'],
  limit: 5,                    // visible items
  suggest?: (input, choices) => choices.filter(c => c.includes(input)),
  highlight?: (haystack, needle) => string,
  initial?: number,
  fallback?: string,           // no results message
  sort?: boolean | ((a, b) => number),
});
```

#### BasicAuth — Username/Password
```typescript
const response = await Enquirer.prompt({
  type: 'basicauth',
  name: 'credentials',
  message: 'Enter credentials:',
  username: 'admin',
  password: '',
  showPassword: false,
});
```

#### Confirm — Yes/No
```typescript
const response = await Enquirer.prompt({
  type: 'confirm',
  name: 'confirmed',
  message: 'Continue?',
  initial: true,        // default selection
});
```

#### Form — Multi-Field Form
```typescript
const response = await Enquirer.prompt({
  type: 'form',
  name: 'user',
  message: 'Enter user details:',
  choices: [
    { name: 'name', message: 'Name', initial: 'John' },
    { name: 'email', message: 'Email', validate: (val) => val.includes('@') },
    { name: 'phone', message: 'Phone' },
  ],
  validate: (values) => {
    if (!values.name) return 'Name is required';
    return true;
  },
});
```

#### Input — Text Input
```typescript
const response = await Enquirer.prompt({
  type: 'input',
  name: 'name',
  message: 'What is your name?',
  initial: 'Guest',
  required: true,
  validate: (value) => value.length >= 2 || 'Too short',
  transformer: (value) => value.trim(),
  result: (value) => value.toLowerCase(),
  format: (value) => value.toUpperCase(),
  filter: (value) => value.replace(/[^a-zA-Z]/g, ''),
});
```

#### Invisible — Hidden Input
```typescript
const response = await Enquirer.prompt({
  type: 'invisible',
  name: 'token',
  message: 'Enter token:',
});
```

#### List — Comma-Separated List
```typescript
const response = await Enquirer.prompt({
  type: 'list',
  name: 'tags',
  message: 'Enter tags (comma-separated):',
  initial: 'a, b, c',
  separator: ',',          // default: /,\s*/
  format: (list) => list.join(', '),
});
```

#### MultiSelect — Checkbox
```typescript
const response = await Enquirer.prompt({
  type: 'multiselect',
  name: 'toppings',
  message: 'Choose toppings',
  choices: [
    { name: 'pepperoni', value: 'pep' },
    { name: 'mushrooms', value: 'mush' },
    { name: 'onions', value: 'onion', disabled: true },
  ],
  initial: ['pepperoni'],   // pre-selected
  required: true,
  validate: (selected) => selected.length > 0 || 'Select at least 1',
  limit: 5,
});
```

#### Numeral — Number Input
```typescript
const response = await Enquirer.prompt({
  type: 'numeral',
  name: 'age',
  message: 'How old are you?',
  min: 0,
  max: 150,
  float: false,
  round: false,
  default: 25,
});
```

#### Password — Masked Input
```typescript
const response = await Enquirer.prompt({
  type: 'password',
  name: 'password',
  message: 'Enter password:',
  mask: '*',
  showMask: true,
  maxLength: 20,
});
```

#### Quiz
```typescript
const response = await Enquirer.prompt({
  type: 'quiz',
  name: 'answer',
  message: 'What is 2+2?',
  choices: ['3', '4', '5'],
  correctChoice: '4',
  points: 10,
});
```

#### Scale — Rating
```typescript
const response = await Enquirer.prompt({
  type: 'scale',
  name: 'rating',
  message: 'Rate the experience',
  scale: [
    { name: 'Poor', value: 1 },
    { name: 'Average', value: 2 },
    { name: 'Good', value: 3 },
    { name: 'Excellent', value: 4 },
  ],
  choices: [
    { name: 'Ease of use', initial: 3 },
    { name: 'Performance', initial: 4 },
    { name: 'Support', initial: 2 },
  ],
});
```

#### Select — Single Select
```typescript
const response = await Enquirer.prompt({
  type: 'select',
  name: 'color',
  message: 'Pick a color:',
  choices: [
    { name: 'Red' },
    { name: 'Blue', hint: 'Recommended!' },
    { name: 'Green', disabled: 'coming soon' },
  ],
  initial: 1,           // index of default
  sort: true,
  scroll: true,
  result: (name) => name.toLowerCase(),
});
```

#### Sort — Reorderable List
```typescript
const response = await Enquirer.prompt({
  type: 'sort',
  name: 'priorities',
  message: 'Sort by priority:',
  choices: ['Feature A', 'Feature B', 'Feature C'],
  numbered: true,       // show numbers
});
```

#### Survey — Multiple Questions
```typescript
const response = await Enquirer.prompt({
  type: 'survey',
  name: 'survey',
  message: 'Please rate:',
  scale: [{ name: '1', value: 1 }, { name: '5', value: 5 }],
  choices: [
    { name: 'Satisfaction', message: 'How satisfied are you? (1-5)' },
    { name: 'Recommend', message: 'Would you recommend? (1-5)' },
  ],
});
```

#### Snippet — Code Template
```typescript
const response = await Enquirer.prompt({
  type: 'snippet',
  name: 'config',
  message: 'Fill in the config:',
  required: true,
  fields: [
    { name: 'host', message: 'Host', initial: 'localhost' },
    { name: 'port', message: 'Port', initial: '3000' },
  ],
  template: 'Host: ${host}:${port}',
});
```

#### Toggle — On/Off Switch
```typescript
const response = await Enquirer.prompt({
  type: 'toggle',
  name: 'enabled',
  message: 'Enable feature?',
  enabled: 'Yes',
  disabled: 'No',
  initial: true,
});
```

### 2.3 Enquirer Class API
```typescript
import Enquirer from 'enquirer';
const enquirer = new Enquirer();

// Register custom prompt
enquirer.register('custom', CustomPrompt);

// Use registered prompts
const response = await enquirer.prompt({
  type: 'custom',
  name: 'test',
  message: 'Custom prompt',
});

// Multiple prompts at once
const answers = await enquirer.prompt([
  { type: 'input', name: 'name', message: 'Name' },
  { type: 'password', name: 'pass', message: 'Password' },
]);
```

### 2.4 Enquirer Types (TypeScript)
```typescript
import Enquirer from 'enquirer';
import type { Prompts, Choice, ChoiceOptions } from 'enquirer';

// Typed prompts
type MyPromptParams = Prompts.Input | Prompts.Select;

const response = await Enquirer.prompt<{name: string}>({
  type: 'input',
  name: 'name',
  message: 'Name',
});
```

---

## 3. Prompts (terkelg/prompts)
**GitHub**: https://github.com/terkelg/prompts | **Stars**: 9K+
**npm**: `prompts` | **Weekly**: ~45.2M | **License**: MIT | **Dependencies**: 2

### 3.1 Installation
```bash
npm install prompts
```

### 3.2 Complete API

#### Basic Usage
```typescript
import prompts from 'prompts';

const response = await prompts({
  type: 'text',
  name: 'value',
  message: 'Enter value:',
  initial: 'default',
  validate: (val: string) => val.length > 0 || 'Required',
});
```

#### All Prompt Types

**text**
```typescript
{
  type: 'text',
  name: string,
  message: string,
  initial?: string,
  validate?: (val: string, values: object, prompt: PromptObject) => boolean | string,
  format?: (val: string, values: object) => any,
  onState?: (state: {value: string, aborted: boolean, exited: boolean}) => void,
  style?: 'default' | 'password' | 'invisible',
}
```

**password**
```typescript
{
  type: 'password',
  name: string,
  message: string,
  initial?: string,
  validate?: (val: string) => boolean | string,
}
```

**invisible**
```typescript
{
  type: 'invisible',
  name: string,
  message: string,
  initial?: string,
}
```

**number**
```typescript
{
  type: 'number',
  name: string,
  message: string,
  initial?: number,
  validate?: (val: number) => boolean | string,
  min?: number,
  max?: number,
  float?: boolean,
  round?: number,
  increment?: number,
}
```

**confirm**
```typescript
{
  type: 'confirm',
  name: string,
  message: string,
  initial?: boolean,
  active?: string,       // text for 'yes' (default: 'yes')
  inactive?: string,     // text for 'no' (default: 'no')
}
```

**list**
```typescript
{
  type: 'list',
  name: string,
  message: string,
  initial?: string,
  separator?: string | RegExp,
}
```

**toggle**
```typescript
{
  type: 'toggle',
  name: string,
  message: string,
  initial?: boolean,
  active?: string,
  inactive?: string,
}
```

**select**
```typescript
{
  type: 'select',
  name: string,
  message: string,
  choices: Array<{
    title: string,
    value: any,
    description?: string,
    disabled?: boolean | string,
  }>,
  initial?: number,
  hint?: string,
  warn?: string,
}
```

**multiselect**
```typescript
{
  type: 'multiselect',
  name: string,
  message: string,
  choices: Array<{
    title: string,
    value: any,
    selected?: boolean,
    disabled?: boolean | string,
  }>,
  initial?: number | number[],
  max?: number,           // max selections
  hint?: string,
  warn?: string,
  min?: number,           // min selections
  optionsPerPage?: number,
}
```

**autocomplete**
```typescript
{
  type: 'autocomplete',
  name: string,
  message: string,
  choices: Array<{ title: string, value: any }>,
  initial?: number,
  suggest?: (input: string, choices: Choice[]) => Promise<Choice[]>,
  limit?: number,
  style?: 'default' | 'password' | 'invisible',
}
```

**date**
```typescript
{
  type: 'date',
  name: string,
  message: string,
  initial?: Date,
  mask?: string,         // date format mask
  validate?: (val: Date) => boolean | string,
  locale?: object,
  min?: Date,
  max?: Date,
}
```

### 3.3 Multi-Prompt
```typescript
const responses = await prompts([
  { type: 'text', name: 'username', message: 'Username:' },
  { type: 'password', name: 'password', message: 'Password:' },
  { type: 'select', name: 'role', message: 'Role:',
    choices: [
      { title: 'Admin', value: 'admin' },
      { title: 'User', value: 'user' },
    ],
  },
]);

// responses = { username: '...', password: '...', role: '...' }
```

### 3.4 Override Answers (Testing)
```typescript
// Programmatic answering
prompts.override({
  name: 'Alice',
  age: 30,
});
```

### 3.5 Injection (Testing)
```typescript
const prompts = require('prompts');
prompts.inject(['Alice', 30, true]);
```

### 3.6 Options
```typescript
const response = await prompts(prompts, {
  onSubmit: (prompt, answer, answers) => {
    // called before each prompt submits
    return true;     // continue (or false to skip remaining)
  },
  onCancel: (prompt, answers) => {
    // called on cancel
    return true;     // continue (or false to abort)
  },
});
```
