# ascii-ansi-colorizer: gradient example

## Input

```
PARTME AI CLI
```

## Script usage

```bash
python3 scripts/colorize.py --mode gradient --direction lr < ./input.txt
```

## Output (schematic, with ANSI)

```
\x1b[38;5;33mP\x1b[0m\x1b[38;5;39mA\x1b[0m\x1b[38;5;45mR\x1b[0m...
```

## Notes
- In a real terminal this displays as gradient-colored text
- When redirecting to a file, disable color and output the plain-text version
