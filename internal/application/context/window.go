package context

import (
	"sort"
	"strings"
	"time"
)

type Message struct {
	Role      string
	Content   string
	Tokens    int
	Timestamp time.Time
}

type Window struct {
	messages   []Message
	maxTokens  int
	usedTokens int
	reserveTop int
}

func NewWindow(maxTokens int) *Window {
	return &Window{
		maxTokens:  maxTokens,
		reserveTop: 1,
	}
}

func (w *Window) SetMaxTokens(n int) {
	w.maxTokens = n
}

func (w *Window) SetReserveTop(n int) {
	w.reserveTop = n
}

func (w *Window) Add(msg Message) {
	w.messages = append(w.messages, msg)
	w.usedTokens += msg.Tokens
}

func (w *Window) Messages() []Message {
	return w.messages
}

func (w *Window) UsedTokens() int {
	return w.usedTokens
}

func (w *Window) AvailableTokens() int {
	avail := w.maxTokens - w.usedTokens
	if avail < 0 {
		return 0
	}
	return avail
}

func (w *Window) Percentage() float64 {
	if w.maxTokens <= 0 {
		return 0
	}
	return float64(w.usedTokens) / float64(w.maxTokens) * 100
}

func (w *Window) Trim() int {
	if w.usedTokens <= w.maxTokens {
		return 0
	}
	removed := 0

	preserved := w.reserveTop
	if preserved > len(w.messages) {
		preserved = len(w.messages)
	}

	var keep []Message
	keep = append(keep, w.messages[:preserved]...)

	middle := w.messages[preserved:]
	sort.Slice(middle, func(i, j int) bool {
		return middle[i].Timestamp.After(middle[j].Timestamp)
	})

	tokens := 0
	for _, m := range keep {
		tokens += m.Tokens
	}

	for _, m := range middle {
		if tokens+m.Tokens > w.maxTokens {
			removed++
			continue
		}
		tokens += m.Tokens
		keep = append(keep, m)
	}

	sort.Slice(keep, func(i, j int) bool {
		return keep[i].Timestamp.Before(keep[j].Timestamp)
	})

	w.messages = keep
	w.usedTokens = tokens
	return removed
}

func (w *Window) Clear() {
	w.messages = nil
	w.usedTokens = 0
}

func (w *Window) Build() []Message {
	w.Trim()
	result := make([]Message, len(w.messages))
	copy(result, w.messages)
	return result
}

func EstimateTokens(text string) int {
	return len(strings.Fields(text)) + len(text)/4
}
