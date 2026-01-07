package service

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"
)

type UserService struct {
	repo    port.ProjectRepository
	storage port.FileStorage
}

func NewUserService(repo port.ProjectRepository, storage port.FileStorage) *UserService {
	return &UserService{repo: repo, storage: storage}
}

func (s *UserService) CreateUser(ctx context.Context, input domain.CreateUserDTO) (string, error) {
	if input.Name == "" || input.Email == "" {
		return "", fmt.Errorf("Name und Email sind Pflichtfelder")
	}
	return s.repo.CreateUser(ctx, &input)
}
