# Permanent Agent Rules

## Language Rule (CRITICAL)
- NO Bengali script (বাংলা) characters allowed ever in any output.
- Use Banglish (Bengali written in English/Latin chars) only.
- E.g., "ami banglay kotha bolchi" ✓ | "আমি বাংলায় কথা বলছি" ✗
- Applies to: all comments, docs, commit msgs, agent prompts, filenames, user communication.

## Execution Rules
- CI=true, non-interactive, no TTY — always use `-y`, `--yes`, `-f` flags
- No editors/pagers (vim, nano, less, man)
- Never stop after tool output unless task is complete — drive process forward
