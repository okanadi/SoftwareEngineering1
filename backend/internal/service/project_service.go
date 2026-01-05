package service

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"
)

type ProjectService struct {
	repo    port.ProjectRepository
	storage port.FileStorage
}

func NewProjectService(repo port.ProjectRepository, storage port.FileStorage) *ProjectService {
	return &ProjectService{repo: repo, storage: storage}
}

func (s *ProjectService) LoginCustomer(ctx context.Context, req domain.CustomerLoginRequest) (*domain.Project, error) {
	return s.repo.GetProjectByLogin(ctx, req.OrderID, req.Lastname)
}

func (s *ProjectService) GetProjectSteps(ctx context.Context, projectID string) ([]domain.ProjectStep, error) {
	return s.repo.GetStepsByProjectID(ctx, projectID)
}

func (s *ProjectService) UpdateStepProgress(ctx context.Context, input domain.UpdateStepInput) error {
	var s3Key string
	var err error

	// Upload zu S3 wenn Datei da ist
	if input.File != nil {
		s3Key, err = s.storage.UploadFile(ctx, input.File, input.FileName, input.ContentType)
		if err != nil {
			return fmt.Errorf("upload failed: %w", err)
		}
	}

	// In DB speichern
	return s.repo.AddHistoryEntry(ctx, input.StepID, input.UserID, input.NewStatus, input.Note, s3Key, input.ContentType)
}
