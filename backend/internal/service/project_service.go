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

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]domain.ProjectDB, error) {
	return s.repo.GetAllProjects(ctx)
}

func (s *ProjectService) GetProjectByID(ctx context.Context, projectID string) (*domain.ProjectDB, error) {
	if projectID == "" {
		return nil, fmt.Errorf("Projekt ID darf nicht leer sein")
	}

	return s.repo.GetProjectByID(ctx, projectID)
}

func (s *ProjectService) GetProjectByCustomerLastname(ctx context.Context, lastname string) ([]domain.ProjectDB, error) {
	if lastname == "" {
		return nil, fmt.Errorf("Nachname darf nicht leer sein")
	}

	return s.repo.GetProjectByCustomerLastname(ctx, lastname)
}

func (s *ProjectService) GetProjectByAddress(ctx context.Context, address string) ([]domain.ProjectDB, error) {
	if address == "" {
		return nil, fmt.Errorf("Adresse darf nicht leer sein")
	}

	return s.repo.GetProjectByAddress(ctx, address)
}

func (s *ProjectService) GetAllCustomerLastnames(ctx context.Context) ([]string, error) {
	return s.repo.GetAllCustomerLastnames(ctx)
}

func (s *ProjectService) GetAllAddresses(ctx context.Context) ([]string, error) {
	return s.repo.GetAllAddresses(ctx)
}

func (s *ProjectService) GetByManagerID(ctx context.Context, managerId string) ([]domain.ProjectDB, error) {
	return s.repo.GetByManagerID(ctx, managerId)
}

func (s *ProjectService) UpdateProject(ctx context.Context, input domain.UpdateProjectDTO) error {
	// Hier könnten Validierungen stehen (z.B. existiert der Manager?)
	return s.repo.UpdateProject(ctx, &input)
}
