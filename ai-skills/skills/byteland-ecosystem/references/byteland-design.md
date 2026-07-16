# byteland-design — Reference

> **GitHub**: https://github.com/ByteLandTechnology/byteland-design
> **Stars**: 0 (new) | **License**: MIT
> **Status**: New (1 commit)

ByteLand's design-language Skill suite — 5 Skills defining the brand's visual, web, icon, motion, and video direction. Each Skill is a plain Markdown `SKILL.md` with structured frontmatter, designed for Skill-aware agents.

---

## Skill Hierarchy

```
byteland-visual-style (foundation)
├── byteland-web-style
├── byteland-icon-style
├── byteland-motion-style
└── byteland-video-style
```

| Skill | Focus |
|-------|-------|
| `byteland-visual-style` | Foundation: tokens, color logic, surface, anti-patterns |
| `byteland-web-style` | Layout, components, responsive, quality checks |
| `byteland-icon-style` | Brand mark, app/function/status icons, exports |
| `byteland-motion-style` | Timing, patterns, reduced-motion |
| `byteland-video-style` | Remotion defaults, segment structure |

---

## Installation

```bash
git clone https://github.com/ByteLandTechnology/byteland-design.git
cd byteland-design
npm test
```

There is nothing to install — tests use Node's built-in runner. Point a Skill-aware agent at `skills/` to use the design language.

---

## Testing (Node built-in test runner)

```bash
npm test
```

Runs 3 suites:
1. `tests/style-rules.test.js` — Token-level invariants (hex colors, viewBox, reduced-motion, etc.)
2. `tests/contrast.test.js` — WCAG AA contrast checks (>= 4.5:1 on white for text-safe colors)
3. `tests/eval-suite.test.js` — Alignment between eval prompts and iteration workspace

Requires Node 18+ (`node --test`).
