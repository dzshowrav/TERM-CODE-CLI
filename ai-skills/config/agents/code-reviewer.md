---
description: Reviews code for quality, security, performance, and best practices
mode: subagent
temperature: 0.1
color: "#10b981"
prompt: "{file:../prompts/code-reviewer.txt}"
permission:
  edit: deny
  bash:
    "*": deny
    "git diff*": allow
    "git log*": allow
    "git status": allow
    "rg *": allow
---

