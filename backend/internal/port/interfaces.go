package port

import (
	"backend/internal/domain"
	"context"
	"io"
)

type ProjectRepository interface {
	GetProjectByLogin(ctx context.Context, id string, lastname string) (*domain.Project, error)
	GetStepsByProjectID(ctx context.Context, projectID string) ([]domain.ProjectStep, error)
	// AddHistoryEntry speichert Update + Media + aktualisiert Step-Status in EINER Transaktion
	AddHistoryEntry(ctx context.Context, stepID string, userID string, status string, note string, s3Key string, fileType string) error
}

type FileStorage interface {
	UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
}
