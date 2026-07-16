package session_test

import (
	"context"
	"log/slog"
	"sync"
	"testing"

	appsession "termcode/internal/application/session"
	domainsession "termcode/internal/domain/session"
)

type fakeRepo struct {
	mu       sync.Mutex
	sessions map[string]*domainsession.Session
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{sessions: make(map[string]*domainsession.Session)}
}

func (f *fakeRepo) Create(_ context.Context, s *domainsession.Session) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sessions[s.ID] = s
	return nil
}

func (f *fakeRepo) GetByID(_ context.Context, id string) (*domainsession.Session, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	s, ok := f.sessions[id]
	if !ok {
		return nil, nil
	}
	return s, nil
}

func (f *fakeRepo) List(_ context.Context) ([]*domainsession.Session, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	result := make([]*domainsession.Session, 0, len(f.sessions))
	for _, s := range f.sessions {
		result = append(result, s)
	}
	return result, nil
}

func (f *fakeRepo) ListByStatus(_ context.Context, status domainsession.Status) ([]*domainsession.Session, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var result []*domainsession.Session
	for _, s := range f.sessions {
		if s.Status == status {
			result = append(result, s)
		}
	}
	return result, nil
}

func (f *fakeRepo) Update(_ context.Context, s *domainsession.Session) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sessions[s.ID] = s
	return nil
}

func (f *fakeRepo) Delete(_ context.Context, id string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	delete(f.sessions, id)
	return nil
}

type fakeMsgRepo struct {
	mu       sync.Mutex
	messages map[string]*domainsession.Message
}

func newFakeMsgRepo() *fakeMsgRepo {
	return &fakeMsgRepo{messages: make(map[string]*domainsession.Message)}
}

func (f *fakeMsgRepo) Create(_ context.Context, m *domainsession.Message) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.messages[m.ID] = m
	return nil
}

func (f *fakeMsgRepo) ListBySession(_ context.Context, sessionID string) ([]*domainsession.Message, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var result []*domainsession.Message
	for _, m := range f.messages {
		if m.SessionID == sessionID {
			result = append(result, m)
		}
	}
	return result, nil
}

func (f *fakeMsgRepo) DeleteBySession(_ context.Context, sessionID string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	for id, m := range f.messages {
		if m.SessionID == sessionID {
			delete(f.messages, id)
		}
	}
	return nil
}

func (f *fakeMsgRepo) CountBySession(_ context.Context, sessionID string) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	count := 0
	for _, m := range f.messages {
		if m.SessionID == sessionID {
			count++
		}
	}
	return count, nil
}

func TestService_Create(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	ses, err := svc.Create(ctx, "test-session", "prov-1", "model-1")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if ses.ID == "" {
		t.Error("expected non-empty ID")
	}
	if ses.Name != "test-session" {
		t.Errorf("expected Name=test-session, got %q", ses.Name)
	}
	if ses.ProviderID != "prov-1" {
		t.Errorf("expected ProviderID=prov-1, got %q", ses.ProviderID)
	}
	if ses.ModelID != "model-1" {
		t.Errorf("expected ModelID=model-1, got %q", ses.ModelID)
	}
	if ses.Status != domainsession.StatusActive {
		t.Errorf("expected Status=active, got %q", ses.Status)
	}
}

func TestService_GetByID(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	created, _ := svc.Create(ctx, "test", "p1", "m1")
	got, err := svc.GetByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("expected ID=%q, got %q", created.ID, got.ID)
	}
}

func TestService_GetByID_NotFound(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	got, err := svc.GetByID(context.Background(), "nonexistent")
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestService_List(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	svc.Create(ctx, "s1", "p1", "m1")
	svc.Create(ctx, "s2", "p2", "m2")

	sessions, err := svc.List(ctx)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(sessions) != 2 {
		t.Errorf("expected 2 sessions, got %d", len(sessions))
	}
}

func TestService_ListActive(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	svc.Create(ctx, "s1", "p1", "m1")
	svc.Create(ctx, "s2", "p2", "m2")

	active, err := svc.ListActive(ctx)
	if err != nil {
		t.Fatalf("ListActive() error = %v", err)
	}
	if len(active) != 2 {
		t.Errorf("expected 2 active sessions, got %d", len(active))
	}
}

func TestService_Update(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	ses, _ := svc.Create(ctx, "test", "p1", "m1")
	ses.Name = "updated"
	err := svc.Update(ctx, ses)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, _ := svc.GetByID(ctx, ses.ID)
	if got.Name != "updated" {
		t.Errorf("expected Name=updated, got %q", got.Name)
	}
}

func TestService_Delete(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	ses, _ := svc.Create(ctx, "test", "p1", "m1")

	msg, _ := svc.AddMessage(ctx, ses.ID, domainsession.RoleUser, "hello")
	if msg == nil {
		t.Fatal("expected message to be created")
	}

	err := svc.Delete(ctx, ses.ID)
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	got, _ := svc.GetByID(ctx, ses.ID)
	if got != nil {
		t.Error("expected session to be deleted")
	}

	count, _ := svc.MessageCount(ctx, ses.ID)
	if count != 0 {
		t.Errorf("expected 0 messages after delete, got %d", count)
	}
}

func TestService_AddMessage(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	ses, _ := svc.Create(ctx, "test", "p1", "m1")

	msg, err := svc.AddMessage(ctx, ses.ID, domainsession.RoleUser, "hello")
	if err != nil {
		t.Fatalf("AddMessage() error = %v", err)
	}
	if msg.Role != domainsession.RoleUser {
		t.Errorf("expected Role=user, got %q", msg.Role)
	}
	if msg.Content != "hello" {
		t.Errorf("expected Content=hello, got %q", msg.Content)
	}
}

func TestService_Messages(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	ses, _ := svc.Create(ctx, "test", "p1", "m1")
	svc.AddMessage(ctx, ses.ID, domainsession.RoleUser, "q1")
	svc.AddMessage(ctx, ses.ID, domainsession.RoleAssistant, "a1")

	msgs, err := svc.Messages(ctx, ses.ID)
	if err != nil {
		t.Fatalf("Messages() error = %v", err)
	}
	if len(msgs) != 2 {
		t.Errorf("expected 2 messages, got %d", len(msgs))
	}
}

func TestService_MessageCount(t *testing.T) {
	svc := appsession.NewService(newFakeRepo(), newFakeMsgRepo(), slog.Default())
	ctx := context.Background()

	ses, _ := svc.Create(ctx, "test", "p1", "m1")
	svc.AddMessage(ctx, ses.ID, domainsession.RoleUser, "hello")

	count, err := svc.MessageCount(ctx, ses.ID)
	if err != nil {
		t.Fatalf("MessageCount() error = %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 message, got %d", count)
	}
}
