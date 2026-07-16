#!/usr/bin/env node
/**
 * ============================================================
 *  INTERACTIVE TUI APP — Ink Ecosystem Demo
 * ============================================================
 * Libraries: ink, ink-text-input, ink-spinner, ink-select-input,
 *            ink-multi-select, ink-progress-bar, ink-table,
 *            ink-gradient, ink-divider, ink-big-text, chalk
 *
 * A fully interactive terminal UI demonstrating all Ink features:
 *   - Big text header with gradient
 *   - Interactive text input with live preview
 *   - Select/multi-select menus
 *   - Real-time progress bars
 *   - Data tables
 *   - Spinners for async operations
 *   - Color themes
 *
 * Run: npx ts-node examples/07-ink-tui.tsx
 * Note: This is a TSX file — requires ts-node JSX support
 * ============================================================
 */

import React, { useState, useEffect, useCallback } from 'react';
import { render, Text, Box, Newline, useInput, useApp, useStdout } from 'ink';
import TextInput from 'ink-text-input';
import Spinner from 'ink-spinner';
import SelectInput from 'ink-select-input';
import MultiSelect from 'ink-multi-select';
import Gradient from 'ink-gradient';
import BigText from 'ink-big-text';
import Divider from 'ink-divider';
import Table from 'ink-table';
import ProgressBar from 'ink-progress-bar';

// ─── Types ───────────────────────────────────────────────
type View = 'home' | 'input' | 'select' | 'multiselect' | 'progress' | 'table' | 'spinner' | 'about';

interface Todo {
  id: number;
  text: string;
  done: boolean;
}

interface AppState {
  view: View;
  name: string;
  color: string;
  hobbies: string[];
  progress: number;
  todos: Todo[];
  loading: boolean;
  logs: string[];
}

// ─── Components ─────────────────────────────────────────

// Home View
const HomeView: React.FC<{ onSelect: (view: View) => void }> = ({ onSelect }) => {
  const items = [
    { label: '📝  Text Input Demo', value: 'input' as View },
    { label: '🎨  Select Menu Demo', value: 'select' as View },
    { label: '✅  Multi-Select Demo', value: 'multiselect' as View },
    { label: '⏳  Progress Bar Demo', value: 'progress' as View },
    { label: '📊  Table Demo', value: 'table' as View },
    { label: '🔄  Spinner Demo', value: 'spinner' as View },
    { label: 'ℹ️  About', value: 'about' as View },
    { label: '🚪  Exit', value: 'home' as View },
  ];

  return (
    <Box flexDirection="column" padding={1}>
      <Gradient name="fruit">
        <BigText text="Ink TUI" font="simple" />
      </Gradient>
      <Divider title="Interactive Demo App" />
      <Box marginTop={1}>
        <Text dimColor>Select a demo:</Text>
      </Box>
      <Box marginTop={1}>
        <SelectInput
          items={items}
          onSelect={(item) => {
            if (item.value === 'home') process.exit(0);
            onSelect(item.value);
          }}
        />
      </Box>
    </Box>
  );
};

// Input Demo
const InputDemo: React.FC<{ onBack: () => void; state: AppState; setState: any }> = ({ onBack, state, setState }) => {
  const { stdin } = useStdout();

  useInput((input, key) => {
    if (key.escape) onBack();
  });

  return (
    <Box flexDirection="column" padding={1}>
      <Divider title="Text Input Demo" />
      <Box marginTop={1} flexDirection="column">
        <Text>Type your name (ESC to go back):</Text>
        <Box marginTop={1}>
          <Text bold>Name: </Text>
          <TextInput
            value={state.name}
            onChange={(v) => setState((s: AppState) => ({ ...s, name: v }))}
            placeholder="Enter your name..."
          />
        </Box>
        {state.name && (
          <Box marginTop={1}>
            <Gradient name="rainbow">
              <Text>Hello, {state.name}! 👋</Text>
            </Gradient>
          </Box>
        )}
        <Box marginTop={1}>
          <Text dimColor>Preview: </Text>
          <Text>{state.name || '...'}</Text>
        </Box>
      </Box>
    </Box>
  );
};

// Select Demo
const SelectDemo: React.FC<{ onBack: () => void; state: AppState; setState: any }> = ({ onBack, state, setState }) => {
  useInput((input, key) => {
    if (key.escape) onBack();
  });

  const colors = [
    { label: '🔴  Red', value: 'red' },
    { label: '🟢  Green', value: 'green' },
    { label: '🔵  Blue', value: 'blue' },
    { label: '🟡  Yellow', value: 'yellow' },
    { label: '🟣  Purple', value: 'purple' },
    { label: '⚪  Gray', value: 'gray' },
  ];

  return (
    <Box flexDirection="column" padding={1}>
      <Divider title="Select Menu Demo" />
      <Box marginTop={1}>
        <Text>Favorite color (ESC to go back):</Text>
      </Box>
      <Box marginTop={1}>
        <SelectInput
          items={colors}
          onSelect={(item) => setState((s: AppState) => ({ ...s, color: item.value }))}
        />
      </Box>
      {state.color && (
        <Box marginTop={1}>
          <Text>
            Selected:{' '}
            <Text color={state.color as any} bold>
              {state.color}
            </Text>
          </Text>
        </Box>
      )}
    </Box>
  );
};

// Multi-Select Demo
const MultiSelectDemo: React.FC<{ onBack: () => void; state: AppState; setState: any }> = ({ onBack, state, setState }) => {
  useInput((input, key) => {
    if (key.escape) onBack();
  });

  const hobbies = [
    { label: '📚  Reading', value: 'reading' },
    { label: '🎮  Gaming', value: 'gaming' },
    { label: '🎵  Music', value: 'music' },
    { label: '🏃  Running', value: 'running' },
    { label: '🍳  Cooking', value: 'cooking' },
    { label: '✈️  Travel', value: 'travel' },
  ];

  return (
    <Box flexDirection="column" padding={1}>
      <Divider title="Multi-Select Demo" />
      <Box marginTop={1}>
        <Text>Select hobbies (Space=Toggle, Enter=Done, ESC=Back):</Text>
      </Box>
      <Box marginTop={1}>
        <MultiSelect
          items={hobbies}
          selected={state.hobbies}
          onSubmit={(items) => {
            setState((s: AppState) => ({ ...s, hobbies: items.map((i: any) => i.value) }));
          }}
        />
      </Box>
      {state.hobbies.length > 0 && (
        <Box marginTop={1} flexDirection="column">
          <Text bold>Selected hobbies:</Text>
          {state.hobbies.map((h) => (
            <Text key={h} color="green">  ✓ {h}</Text>
          ))}
        </Box>
      )}
    </Box>
  );
};

// Progress Demo
const ProgressDemo: React.FC<{ onBack: () => void; state: AppState; setState: any }> = ({ onBack, state, setState }) => {
  useInput((input, key) => {
    if (key.escape) {
      setState((s: AppState) => ({ ...s, progress: 0 }));
      onBack();
    }
  });

  useEffect(() => {
    if (state.progress < 100) {
      const timer = setInterval(() => {
        setState((s: AppState) => ({
          ...s,
          progress: Math.min(s.progress + Math.random() * 15, 100),
        }));
      }, 500);
      return () => clearInterval(timer);
    }
  }, [state.progress < 100]);

  return (
    <Box flexDirection="column" padding={1}>
      <Divider title="Progress Bar Demo" />
      <Box marginTop={1} flexDirection="column">
        <Text bold>Task Progress (ESC to go back):</Text>
        <Box marginTop={1} width={60}>
          <ProgressBar
            percent={state.progress / 100}
            character="█"
            rightLabel={`${Math.round(state.progress)}%`}
          />
        </Box>
        {state.progress >= 100 && (
          <Box marginTop={1}>
            <Gradient name="rainbow">
              <Text bold>✅ Task Complete!</Text>
            </Gradient>
          </Box>
        )}
        <Box marginTop={1}>
          <Text dimColor>Press ESC to reset and go back</Text>
        </Box>
      </Box>
    </Box>
  );
};

// Table Demo
const TableDemo: React.FC<{ onBack: () => void; state: AppState; setState: any }> = ({ onBack, state, setState }) => {
  useInput((input, key) => {
    if (key.escape) onBack();
  });

  const data = [
    { id: '1', task: 'Learn Ink', status: '✅ Done', priority: 'High', assignee: 'You' },
    { id: '2', task: 'Build TUI', status: '✅ Done', priority: 'High', assignee: 'You' },
    { id: '3', task: 'Add animations', status: '🔄 Progress', priority: 'Medium', assignee: 'You' },
    { id: '4', task: 'Release v1.0', status: '⏳ Pending', priority: 'Low', assignee: 'TBD' },
    { id: '5', task: 'Write docs', status: '⏳ Pending', priority: 'Medium', assignee: 'You' },
  ];

  return (
    <Box flexDirection="column" padding={1}>
      <Divider title="Data Table Demo" />
      <Box marginTop={1}>
        <Text>Task List (ESC to go back):</Text>
      </Box>
      <Box marginTop={1}>
        <Table data={data} />
      </Box>
    </Box>
  );
};

// Spinner Demo
const SpinnerDemo: React.FC<{ onBack: () => void; state: AppState; setState: any }> = ({ onBack, state, setState }) => {
  useInput((input, key) => {
    if (key.escape) onBack();
  });

  useEffect(() => {
    const tasks = [
      'Loading configuration...',
      'Connecting to server...',
      'Fetching data...',
      'Processing results...',
      'Almost done...',
    ];

    let i = 0;
    const timer = setInterval(() => {
      if (i < tasks.length) {
        setState((s: AppState) => ({
          ...s,
          logs: [...s.logs, tasks[i]],
        }));
        i++;
      } else {
        clearInterval(timer);
      }
    }, 1500);

    return () => clearInterval(timer);
  }, []);

  return (
    <Box flexDirection="column" padding={1}>
      <Divider title="Spinner Demo" />
      <Box marginTop={1} flexDirection="column">
        {state.logs.length < 5 ? (
          <Box>
            <Spinner type="dots" />
            <Text>  Processing...</Text>
          </Box>
        ) : (
          <Gradient name="passion">
            <Text bold>✅ All tasks complete!</Text>
          </Gradient>
        )}

        <Box marginTop={1} flexDirection="column">
          {state.logs.map((log, i) => (
            <Box key={i}>
              <Text color={i === state.logs.length - 1 && state.logs.length < 5 ? 'yellow' : 'green'}>
                {i === state.logs.length - 1 && state.logs.length < 5 ? <Spinner type="simpleDotsScrolling" /> : '✓'}  {log}
              </Text>
            </Box>
          ))}
        </Box>

        <Box marginTop={1}>
          <Text dimColor>Press ESC to go back</Text>
        </Box>
      </Box>
    </Box>
  );
};

// About View
const AboutView: React.FC<{ onBack: () => void }> = ({ onBack }) => {
  useInput((input, key) => {
    if (key.escape) onBack();
  });

  return (
    <Box flexDirection="column" padding={1}>
      <Gradient name="rainbow">
        <BigText text="Ink TUI" font="tiny" />
      </Gradient>
      <Divider title="About" />
      <Box marginTop={1} flexDirection="column">
        <Text bold>Interactive Terminal UI Demo</Text>
        <Newline />
        <Text>Built with:</Text>
        <Text>  • Ink — React for CLIs</Text>
        <Text>  • Ink Text Input — Interactive input</Text>
        <Text>  • Ink Select Input — Menu selection</Text>
        <Text>  • Ink Multi Select — Checkbox lists</Text>
        <Text>  • Ink Progress Bar — Progress indicators</Text>
        <Text>  • Ink Table — Data tables</Text>
        <Text>  • Ink Spinner — Loading animations</Text>
        <Text>  • Ink Gradient — Color gradients</Text>
        <Text>  • Ink Big Text — ASCII art text</Text>
        <Text>  • Ink Divider — Section dividers</Text>
        <Newline />
        <Text dimColor>Use ↑↓ arrows to navigate, Enter to select, ESC to go back</Text>
        <Newline />
        <Text>Version 1.0.0 — {new Date().getFullYear()}</Text>
      </Box>
    </Box>
  );
};

// ─── Main App ───────────────────────────────────────────
const App: React.FC = () => {
  const [state, setState] = useState<AppState>({
    view: 'home',
    name: '',
    color: '',
    hobbies: [],
    progress: 0,
    todos: [],
    loading: false,
    logs: [],
  });

  const handleSelect = useCallback((view: View) => {
    setState((s) => ({ ...s, view, progress: 0 }));
  }, []);

  const handleBack = useCallback(() => {
    setState((s) => ({ ...s, view: 'home', logs: [], progress: 0 }));
  }, []);

  switch (state.view) {
    case 'home':
      return <HomeView onSelect={handleSelect} />;
    case 'input':
      return <InputDemo onBack={handleBack} state={state} setState={setState} />;
    case 'select':
      return <SelectDemo onBack={handleBack} state={state} setState={setState} />;
    case 'multiselect':
      return <MultiSelectDemo onBack={handleBack} state={state} setState={setState} />;
    case 'progress':
      return <ProgressDemo onBack={handleBack} state={state} setState={setState} />;
    case 'table':
      return <TableDemo onBack={handleBack} state={state} setState={setState} />;
    case 'spinner':
      return <SpinnerDemo onBack={handleBack} state={state} setState={setState} />;
    case 'about':
      return <AboutView onBack={handleBack} />;
    default:
      return <HomeView onSelect={handleSelect} />;
  }
};

// ─── Render ─────────────────────────────────────────────
render(<App />);
