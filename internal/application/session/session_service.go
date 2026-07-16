package session

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"termcode/internal/domain/session"
)

type Repository interface {
	Create(ctx context.Context, s *session.Session) error
	GetByID(ctx context.Context, id string) (*session.Session, error)
	List(ctx context.Context) ([]*session.Session, error)
	ListByStatus(ctx context.Context, status session.Status) ([]*session.Session, error)
	Update(ctx context.Context, s *session.Session) error
	Delete(ctx context.Context, id string) error
}

type MessageRepository interface {
	Create(ctx context.Context, m *session.Message) error
	ListBySession(ctx context.Context, sessionID string) ([]*session.Message, error)
	DeleteBySession(ctx context.Context, sessionID string) error
	CountBySession(ctx context.Context, sessionID string) (int, error)
}

type Service struct {
	repo    Repository
	msgRepo MessageRepository
	logger  *slog.Logger
}

func NewService(repo Repository, msgRepo MessageRepository, logger *slog.Logger) *Service {
	return &Service{
		repo:    repo,
		msgRepo: msgRepo,
		logger:  logger.With("svc", "session"),
	}
}

func (s *Service) Create(ctx context.Context, name, providerID, modelID string) (*session.Session, error) {
	ses := session.New(name, providerID, modelID)
	if err := s.repo.Create(ctx, ses); err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	s.logger.Info("session created", "id", ses.ID, "name", name)
	return ses, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*session.Session, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]*session.Session, error) {
	return s.repo.List(ctx)
}

func (s *Service) ListActive(ctx context.Context) ([]*session.Session, error) {
	return s.repo.ListByStatus(ctx, session.StatusActive)
}

func (s *Service) Update(ctx context.Context, ses *session.Session) error {
	ses.UpdatedAt = time.Now()
	return s.repo.Update(ctx, ses)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.msgRepo.DeleteBySession(ctx, id); err != nil {
		return fmt.Errorf("delete messages: %w", err)
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	s.logger.Info("session deleted", "id", id)
	return nil
}

func (s *Service) AddMessage(ctx context.Context, sesID string, role session.Role, content string) (*session.Message, error) {
	msg := session.NewMessage(sesID, role, content)
	if err := s.msgRepo.Create(ctx, msg); err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}
	return msg, nil
}

func (s *Service) Messages(ctx context.Context, sessionID string) ([]*session.Message, error) {
	return s.msgRepo.ListBySession(ctx, sessionID)
}

func (s *Service) MessageCount(ctx context.Context, sessionID string) (int, error) {
	return s.msgRepo.CountBySession(ctx, sessionID)
}
