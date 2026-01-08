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
