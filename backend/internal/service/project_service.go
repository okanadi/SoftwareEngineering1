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

func (s *ProjectService) CreateProject(ctx context.Context, input domain.CreateProjectDTO) (string, error) {
	if input.CustomerLastname == "" || input.Address == "" {
		return "", fmt.Errorf("Nachname und Adresse sind Pflichtfelder")
	}

	return s.repo.CreateProject(ctx, &input)
}
