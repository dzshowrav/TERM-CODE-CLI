package command

import (
	"strings"
)

type ParsedCommand struct {
	Raw      string
	Name     string
	Args     []string
	Flags    map[string]string
	HasFlags bool
}

func Parse(input string) *ParsedCommand {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "/") {
		return nil
	}

	pc := &ParsedCommand{
		Raw:   input,
		Flags: make(map[string]string),
	}

	input = input[1:]
	parts := splitArgs(input)
	if len(parts) == 0 {
		return pc
	}

	pc.Name = parts[0]

	i := 1
	for i < len(parts) {
		part := parts[i]
		if strings.HasPrefix(part, "--") {
			flagName := part[2:]
			flagVal := "true"
			if i+1 < len(parts) && !strings.HasPrefix(parts[i+1], "-") {
				i++
				flagVal = parts[i]
			}
			pc.Flags[flagName] = flagVal
			pc.HasFlags = true
		} else if strings.HasPrefix(part, "-") && len(part) == 2 {
			flagName := part[1:]
			flagVal := "true"
			if i+1 < len(parts) && !strings.HasPrefix(parts[i+1], "-") {
				i++
				flagVal = parts[i]
			}
			pc.Flags[flagName] = flagVal
			pc.HasFlags = true
		} else {
			pc.Args = append(pc.Args, part)
		}
		i++
	}

	return pc
}

func splitArgs(input string) []string {
	var args []string
	var cur strings.Builder
	inQuote := false
	quoteChar := byte(0)

	for i := 0; i < len(input); i++ {
		c := input[i]

		if inQuote {
			if c == quoteChar {
				inQuote = false
				continue
			}
			cur.WriteByte(c)
			continue
		}

		if c == '"' || c == '\'' {
			inQuote = true
			quoteChar = c
			continue
		}

		if c == ' ' {
			if cur.Len() > 0 {
				args = append(args, cur.String())
				cur.Reset()
			}
			continue
		}

		cur.WriteByte(c)
	}

	if cur.Len() > 0 {
		args = append(args, cur.String())
	}

	return args
}

func (pc *ParsedCommand) Arg(n int) string {
	if n >= 0 && n < len(pc.Args) {
		return pc.Args[n]
	}
	return ""
}

func (pc *ParsedCommand) Flag(name string) (string, bool) {
	v, ok := pc.Flags[name]
	return v, ok
}

func (pc *ParsedCommand) String() string {
	if pc == nil {
		return ""
	}
	var b strings.Builder
	b.WriteString("/")
	b.WriteString(pc.Name)
	for _, a := range pc.Args {
		b.WriteString(" ")
		b.WriteString(a)
	}
	for k, v := range pc.Flags {
		b.WriteString(" --")
		b.WriteString(k)
		if v != "true" {
			b.WriteString(" ")
			b.WriteString(v)
		}
	}
	return b.String()
}
