# CLI Argument Parsers — Complete A-to-Z Reference

---

## 1. Commander.js (tj/commander.js)
**GitHub**: https://github.com/tj/commander.js | **Stars**: 26K+
**npm**: `commander` | **Weekly**: ~50M | **License**: MIT | **Dependencies**: 0
**Website**: https://tj.github.io/commander.js/

### 1.1 Installation
```bash
npm install commander
```

### 1.2 Complete API

#### Command Class
```typescript
import { Command } from 'commander';

const program = new Command(name?: string);
```

#### Configuration Methods
```typescript
program
  .name('my-cli')                                    // program name (shown in help)
  .description('CLI description')                     // description
  .summary('Short summary')                            // brief summary (v8+)
  .version('1.0.0', '-v, --version')                   // version option
  .usage('[options] <command>')                        // custom usage
  .helpOption('-h, --help', 'Show help')               // help flag
  .addHelpCommand(true | false | string)                // enable/disable 'help [cmd]'
  .helpCommand(name, description)                       // custom help command
  .configureHelp({ optionFlags: '--custom' })           // configure help format
  .configureOutput({ writeOut, writeErr, outputError }) // custom output
  .showHelpAfterError(true | string)                   // show help on error
  .showSuggestionAfterError(true)                       // suggest similar commands
  .allowUnknownOption(true | false)                     // allow unknown opts
  .allowExcessArguments(true | false)                   // allow extra args
  .exitOverride((err) => {})                            // override exit
  .passThroughOptions(true)                             // pass options through
  .storeOptionsAsProperties(true)                       // store as props
  .enablePositionalOptions(true)                        // positional opts
;

// Parser callback
program.parse(process.argv);                           // parse argv
program.parseAsync(process.argv);                      // async parse
program.parseOptions(process.argv);                    // parse without actions
```

#### Options
```typescript
// Syntax:
// --flag                          boolean
// -c, --command                   boolean with short flag
// -p, --port <number>            required value (angle brackets)
// -o, --output [file]             optional value (square brackets)
// -d, --drink <drink>             required with description
// --no-color                      negated boolean
// --yes                           positive alias for --no-yes

program
  .option('-d, --debug', 'Enable debug')                          // boolean
  .option('-p, --port <number>', 'Port', parseInt, 3000)          // with parser + default
  .option('--no-color', 'Disable color')                          // negated
  .option('-n, --number <n>', 'Number', /^[0-9]+$/)              // regex validation
  .option('--env <env>', 'Environment', /^(dev|prod)$/i)         // choices via regex
  .requiredOption('-k, --api-key <key>', 'API key')              // required option
  .option('--list <items...>', 'List of items')                  // variadic
  .option('-c, --choice <type>', 'Choice', ['a', 'b', 'c'])      // choices array
;

// Access options in action handler:
program.action((opts) => {
  opts.debug;    // boolean
  opts.port;     // number (parsed)
  opts.color;    // boolean
});
// Or anywhere:
program.opts();          // all options
program.optsWithGlobals();  // options + global options
```

#### Commands
```typescript
// === Standalone command ===
program
  .command('serve')
  .description('Start server')
  .argument('[port]', 'Port number', parseInt, 3000)
  .option('-v, --verbose', 'Verbose')
  .action((port, options) => {
    console.log(`Serving on port ${port}`);
  });

// === Subcommand with alias ===
const deploy = program
  .command('deploy')
  .alias('d')
  .description('Deploy the application');

deploy
  .command('staging')
  .description('Deploy to staging')
  .action(() => {});

deploy
  .command('production')
  .description('Deploy to production')
  .action(() => {});

// === Command with argument definition ===
program
  .command('make:component <name>')
  .description('Generate a component')
  .argument('[dir]', 'Target directory', './components')
  .option('-t, --type <type>', 'Component type', 'functional')
  .action((name, dir, opts) => {});
```

#### Arguments
```typescript
program
  .argument('<required>', 'Required argument')
  .argument('[optional]', 'Optional argument')
  .argument('[variadic...]', 'Variadic argument')     // rest args
  .argument('<file>', 'Input file', /\.\w+$/)         // regex validation
  .argument('[num]', 'Number', parseInt, 10)           // with parser + default
;

// Add arguments to command
const cmd = program.command('process');
cmd.argument('<input>', 'Input file');
cmd.argument('<output>', 'Output file');
cmd.argument('[options...]', 'Additional options');
cmd.action((input, output, options) => {});
```

#### Hooks & Lifecycle
```typescript
// Pre/post action hooks (v7+)
program.hook('preAction', (thisCommand, actionCommand) => {
  // runs before action
});

program.hook('postAction', (thisCommand, actionCommand) => {
  // runs after action
});
```

#### Help Customization
```typescript
program.helpInformation();              // get help text
program.outputHelp();                   // output help
program.outputHelp(callback);           // output and callback

// Custom help
program.configureHelp({
  sortSubcommands: true,
  sortOptions: true,
  showGlobalOptions: true,
  // Custom format functions:
  formatHelp: (cmd, helper) => '...',
  subcommandTerm: (cmd) => cmd.name(),
  optionTerm: (option) => option.flags,
  visibleCommands: (cmd) => cmd.commands,
  visibleOptions: (cmd) => cmd.options,
  visibleGlobalOptions: (cmd) => cmd.options,
});
```

#### Events
```typescript
program.on('--help', () => { console.log('Extra help text'); });
program.on('command:*', (cmd) => { console.log('Unknown command:', cmd); });
```

#### TypeScript Usage
```typescript
import { Command, type OptionValues } from 'commander';

interface DeployOptions extends OptionValues {
  force?: boolean;
  timeout?: number;
}

const program = new Command();
program
  .option('-f, --force', 'Force deploy')
  .option('-t, --timeout <seconds>', 'Timeout', parseInt, 60);

program.parse();
const options = program.opts<DeployOptions>();

// v8+ — action handler with auto-typed args
program
  .command('build')
  .argument('<target>')
  .option('--optimize', 'Optimize build')
  .action((target: string, options: { optimize?: boolean }) => {
    // options are auto-typed in v8+
  });
```

---

## 2. Yargs (yargs/yargs)
**GitHub**: https://github.com/yargs/yargs | **Stars**: 11K+
**npm**: `yargs` | **Weekly**: ~30M | **License**: MIT | **Dependencies**: 7 (small)
**Website**: https://yargs.js.org/

### 2.1 Installation
```bash
npm install yargs
```

### 2.2 Complete API

#### Setup
```typescript
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

yargs(hideBin(process.argv))
  .scriptName('my-cli')
  .usage('$0 <command> [options]')
  .parse();
```

#### Options
```typescript
yargs
  // Simple option
  .option('name', {
    alias: 'n',
    type: 'string',
    description: 'Your name',
    default: 'Guest',
    demandOption: false,        // required
    requiresArg: true,          // requires value
    nargs: 1,                   // num args consumed
    group: 'User Options:',     // group in help
    hidden: false,              // hide from help
    choices: ['a', 'b', 'c'],   // allowed values
    coerce: (val) => val.toUpperCase(),  // transform
  })

  // Shorthand
  .option('verbose', { alias: 'v', type: 'boolean' })
  .option('config', { alias: 'c', type: 'string' })
  .option('port', { type: 'number', default: 3000 })

  // Alternative API
  .string('name')
  .boolean('verbose')
  .number('port')
  .array('items')
  .count('verbose')       // -vvv = 3

  // Defaults
  .default('port', 3000)
  .default('config', './config.json')

  // Required
  .demandOption('api-key')
  .demandOption(['key', 'secret'])  // multiple required

  // Choices
  .choices('env', ['dev', 'staging', 'prod'])

  // Conflicts
  .conflicts('optionA', 'optionB')     // A and B can't coexist
  .implies('key', 'secret')             // key requires secret
;
```

#### Positional Arguments
```typescript
yargs.command(
  'build <source> [dest]',
  'Build the project',
  (yargs) => {
    yargs
      .positional('source', {
        describe: 'Source directory',
        type: 'string',
        demandOption: true,
      })
      .positional('dest', {
        describe: 'Destination directory',
        type: 'string',
        default: './dist',
      });
  },
  (argv) => {
    console.log(`Building from ${argv.source} to ${argv.dest}`);
  }
);
```

#### Subcommands
```typescript
// Command with builder and handler
yargs.command({
  command: 'deploy <env>',
  aliases: ['d'],
  describe: 'Deploy to environment',
  builder: (yargs) => {
    return yargs
      .positional('env', {
        type: 'string',
        choices: ['staging', 'production'],
        describe: 'Target environment',
      })
      .option('force', { alias: 'f', type: 'boolean' });
  },
  handler: (argv) => {
    console.log(`Deploying to ${argv.env}`);
  },
});

// Command directory (auto-load modules)
yargs.commandDir('./commands', {
  extensions: ['js', 'ts'],
  exclude: /\.test\./,
  recurse: true,        // recursive search
  visit: (cmd) => cmd,  // transform loaded command
});
```

#### Middleware
```typescript
yargs.middleware([
  (argv) => {
    // pre-processing all commands
    argv.startTime = Date.now();
  },
  {
    applyBeforeValidation: true,  // run before validation
    global: true,                   // apply to all commands
  },
]);
```

#### Help System
```typescript
yargs
  .help()                              // enable --help
  .help('help', 'Show this help')      // custom help flag
  .alias('h', 'help')                  // alias for help
  .showHelpOnFail(true)                // show help on failure
  .wrap(80)                            // help line width
  .epilog('Copyright 2026')            // footer text
  .epilogue('See more at...')
  .example('$0 deploy staging', 'Deploy to staging')
  .example('$0 deploy --force', 'Force deploy');
```

#### Validation
```typescript
yargs
  .check((argv, options) => {
    if (argv.port < 0 || argv.port > 65535) {
      throw new Error('Invalid port');
    }
    return true;      // validation passed
  })

  // Strict mode
  .strict()                         // error on unknown options
  .strictCommands()                 // error on unknown commands
  .strictOptions()                  // error on unknown options only

  // Demand
  .demandCommand(1, 'Need at least 1 command')
  .demandCommand(1, 3, 'Too many commands')  // min, max
  .demandOption(['key'])
;
```

#### Shell Completion
```typescript
yargs.completion('completion', (current, argv, completed) => {
  // Custom completion function
  completed(['option1', 'option2']);
});

// Or generate completion scripts:
// my-cli completion >> ~/.bashrc
```

#### Parser Configuration
```typescript
yargs.parserConfiguration({
  'short-option-groups': true,        // -abc = -a -b -c
  'camel-case-expansion': true,       // --my-opt → argv.myOpt
  'dot-notation': true,               // foo.bar → argv.foo.bar
  'parse-numbers': true,              // auto-parse numbers
  'boolean-negation': true,           // --no-flag
  'combine-arrays': false,            // -a 1 -a 2 = [1,2] or [1,-a,2]
  'duplicate-arguments-array': true,  // --x a --x b = [a,b]
  'flatten-duplicate-arrays': true,   // flatten nested arrays
  'halt-at-non-option': false,        // stop at --
  'negation-prefix': 'no-',           // prefix for negation
  'populate--': false,                // populate argv['--']
  'set-placeholder-key': false,       // set key for placeholder
  'unknown-options-as-args': false,   // treat unknown opts as args
});
```

#### TypeScript
```typescript
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

interface DeployArgs {
  environment: 'staging' | 'production';
  force?: boolean;
  timeout: number;
}

yargs(hideBin(process.argv))
  .command<DeployArgs>(
    'deploy <environment>',
    'Deploy',
    (yargs) => yargs
      .positional('environment', {
        type: 'string',
        choices: ['staging', 'production'] as const,
      })
      .option('force', { type: 'boolean' })
      .option('timeout', { type: 'number', default: 60 }),
    (argv) => {
      // argv is typed as DeployArgs
      console.log(argv.environment);
    }
  )
  .parse();

// Utility type for inferred options
import type { InferredOptionTypes } from 'yargs';
const options = {
  port: { type: 'number' as const, default: 3000 },
  verbose: { type: 'boolean' as const },
} as const;
type Options = InferredOptionTypes<typeof options>;
// { port: number; verbose: boolean }
```
