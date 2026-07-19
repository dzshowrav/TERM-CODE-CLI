package eventbus

import (
	"fmt"
	"sync"
	"time"
)

type EventType string

const (
	// Deprecated: unused
	// EventAppStarted       EventType = "app.started"
	EventModelChanged EventType = "model.changed"
	// Deprecated: unused
	// EventProviderChanged  EventType = "provider.changed"
	// Deprecated: unused
	// EventSessionCreated   EventType = "session.created"
	// Deprecated: unused
	// EventSessionDeleted   EventType = "session.deleted"
	// Deprecated: unused
	// EventMessageSent      EventType = "message.sent"
	// Deprecated: unused
	// EventMessageReceived  EventType = "message.received"
	EventStreamStarted EventType = "stream.started"
	// Deprecated: unused
	// EventStreamChunk      EventType = "stream.chunk"
	EventStreamComplete EventType = "stream.complete"
	// Deprecated: unused
	// EventTokenUpdate      EventType = "token.update"
	EventToolStarted  EventType = "tool.started"
	EventToolComplete EventType = "tool.complete"
	EventToolFailed   EventType = "tool.failed"
	// Deprecated: unused
	// EventWorkspaceChanged EventType = "workspace.changed"
	// Deprecated: unused
	// EventThemeChanged     EventType = "theme.changed"
	// Deprecated: unused
	// EventConfigChanged    EventType = "config.changed"
	EventAttention    EventType = "attention.required"
	EventNotification EventType = "notification.show"
	// Deprecated: unused
	// EventScreenChanged    EventType = "screen.changed"
	EventError EventType = "error"
	// Deprecated: unused
	// EventProgress         EventType = "progress"
	// Deprecated: unused
	// EventToolOutput       EventType = "tool.output"
	// Deprecated: unused
	// EventDialogOpened     EventType = "dialog.opened"
	// Deprecated: unused
	// EventDialogClosed     EventType = "dialog.closed"
)

type Event struct {
	Type      EventType
	Data      any
	Timestamp time.Time
	Source    string
}

type HandlerFunc func(Event)

type Subscription struct {
	id     uint64
	etype  EventType
	handle HandlerFunc
	async  bool
}

type Bus struct {
	mu           sync.RWMutex
	subs         map[EventType][]*Subscription
	nextID       uint64
	history      []Event
	historyLimit int
}

func New() *Bus {
	return &Bus{
		subs:         make(map[EventType][]*Subscription),
		historyLimit: 100,
	}
}

func (b *Bus) Subscribe(etype EventType, handler HandlerFunc) uint64 {
	return b.subscribe(etype, handler, false)
}

func (b *Bus) SubscribeAsync(etype EventType, handler HandlerFunc) uint64 {
	return b.subscribe(etype, handler, true)
}

func (b *Bus) subscribe(etype EventType, handler HandlerFunc, async bool) uint64 {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.nextID++
	sub := &Subscription{
		id:     b.nextID,
		etype:  etype,
		handle: handler,
		async:  async,
	}
	b.subs[etype] = append(b.subs[etype], sub)
	return sub.id
}

func (b *Bus) Unsubscribe(id uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for etype, subs := range b.subs {
		for i, sub := range subs {
			if sub.id == id {
				b.subs[etype] = append(subs[:i], subs[i+1:]...)
				return
			}
		}
	}
}

func (b *Bus) Emit(etype EventType, data any) {
	b.EmitWithSource(etype, data, "")
}

func (b *Bus) EmitWithSource(etype EventType, data any, source string) {
	event := Event{
		Type:      etype,
		Data:      data,
		Timestamp: time.Now(),
		Source:    source,
	}

	b.mu.RLock()
	subs := b.subs[etype]
	allSubs := b.subs[""]
	b.mu.RUnlock()

	b.addHistory(event)

	for _, sub := range subs {
		b.dispatch(sub, event)
	}
	for _, sub := range allSubs {
		b.dispatch(sub, event)
	}
}

func (b *Bus) dispatch(sub *Subscription, event Event) {
	if sub.async {
		go sub.handle(event)
	} else {
		sub.handle(event)
	}
}

func (b *Bus) addHistory(event Event) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.history = append(b.history, event)
	if len(b.history) > b.historyLimit {
		b.history = b.history[1:]
	}
}

func (b *Bus) History(limit int) []Event {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if limit <= 0 || limit > len(b.history) {
		limit = len(b.history)
	}
	result := make([]Event, limit)
	copy(result, b.history[len(b.history)-limit:])
	return result
}

func (b *Bus) ClearHistory() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.history = nil
}

func (b *Bus) SubscriberCount(etype EventType) int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.subs[etype])
}

type MultiBus struct {
	buses map[string]*Bus
	mu    sync.RWMutex
}

func NewMulti() *MultiBus {
	return &MultiBus{
		buses: make(map[string]*Bus),
	}
}

func (m *MultiBus) Get(name string) *Bus {
	m.mu.RLock()
	bus, ok := m.buses[name]
	m.mu.RUnlock()
	if ok {
		return bus
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	bus = New()
	m.buses[name] = bus
	return bus
}

func (m *MultiBus) Names() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	names := make([]string, 0, len(m.buses))
	for n := range m.buses {
		names = append(names, n)
	}
	return names
}

var (
	ErrSubNotFound = fmt.Errorf("subscription not found")
	ErrBusNotFound = fmt.Errorf("bus not found")
)
