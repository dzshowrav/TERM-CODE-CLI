package session

import (
	"context"
	"encoding/json"
	"fmt"
)

type ExportData struct {
	Session  *SessionExport   `json:"session"`
	Messages []*MessageExport `json:"messages"`
}

type SessionExport struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProviderID string `json:"provider_id"`
	ModelID    string `json:"model_id"`
	CreatedAt  string `json:"created_at"`
}

type MessageExport struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (s *Service) Export(ctx context.Context, sessionID string) ([]byte, error) {
	ses, err := s.repo.GetByID(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("get session: %w", err)
	}

	msgs, err := s.msgRepo.ListBySession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("get messages: %w", err)
	}

	export := ExportData{
		Session: &SessionExport{
			ID:         ses.ID,
			Name:       ses.Name,
			ProviderID: ses.ProviderID,
			ModelID:    ses.ModelID,
			CreatedAt:  ses.CreatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}

	for _, m := range msgs {
		export.Messages = append(export.Messages, &MessageExport{
			Role:    string(m.Role),
			Content: m.Content,
		})
	}

	data, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	return data, nil
}
