package service

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"
)

type ProjectStepService struct {
	repo    port.ProjectRepository
	storage port.FileStorage
}

func NewProjectStepService(repo port.ProjectRepository, storage port.FileStorage) *ProjectStepService {
	return &ProjectStepService{repo: repo, storage: storage}
}

func (s *ProjectStepService) CreateProjectStep(ctx context.Context, input domain.CreateProjectStepDTO) (string, error) {
	if input.ProjectID == "" || input.Title == "" {
		return "", fmt.Errorf("Projekt-ID und Titel sind Pflichtfelder")
	}
	return s.repo.CreateProjectStep(ctx, &input)
}

func (s *ProjectStepService) GetProjectSteps(ctx context.Context, input string) ([]domain.ProjectStepDB, error) {
	if input == "" {
		return nil, fmt.Errorf("Projekt-ID ist ein Pflichtfeld")
	}

	return s.repo.GetStepsProjectByProjectID(ctx, input)
}

func (s *ProjectStepService) GetProjectStepByID(ctx context.Context, projectID string, stepID string) (*domain.ProjectStepDB, error) {
	if projectID == "" || stepID == "" {
		return nil, fmt.Errorf("Projekt-ID und Schritt-ID sind Pflichtfelder")
	}
	steps, err := s.repo.GetStepsProjectByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	for _, step := range steps {
		if step.ID.String() == stepID {
			return &step, nil
		}
	}
	return nil, fmt.Errorf("Schritt mit der ID %s nicht gefunden", stepID)
}

func (s *ProjectStepService) UpdateStepProgress(ctx context.Context, input domain.UpdateProjectStepDTO) error {
	var s3Key string
	var err error

	// 1. Datei zu S3 (außerhalb der DB-Transaktion)
	if input.File != nil {
		s3Key, err = s.storage.UploadFile(ctx, input.File, input.FileName, input.FileContentType)
		if err != nil {
			return fmt.Errorf("S3 upload failed: %w", err)
		}
	}

	// 2. Transaktionale DB-Operation
	err = s.repo.UpdateStepWithHistoryAndMedia(ctx,
		input.StepId,
		input.UserId,
		input.NewStatus,
		input.Note,
		s3Key,
		input.FileContentType,
	)

	if err != nil {
		// Optional: Hier könnte man den S3-Key wieder löschen (Cleanup),
		// falls die DB fehlschlägt.
		return fmt.Errorf("database transaction failed: %w", err)
	}

	return nil
}
