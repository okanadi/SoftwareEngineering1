package service

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"
)

type HistoryService struct {
	repo    port.ProjectRepository
	storage port.FileStorage
}

func NewHistoryService(repo port.ProjectRepository, storage port.FileStorage) *HistoryService {
	return &HistoryService{repo: repo, storage: storage}
}

func (s *HistoryService) GetHistory(ctx context.Context, projectID string) ([]domain.ProjectStepHistoryDTO, error) {
	if projectID == "" {
		return nil, fmt.Errorf("ProjektID sind Pflichtfeld")
	}

	return s.repo.GetHistory(ctx, projectID)
}
