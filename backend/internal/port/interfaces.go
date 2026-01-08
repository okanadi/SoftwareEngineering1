package port

import (
	"backend/internal/domain"
	"context"
	"io"
)

type ProjectRepository interface {

	//ProjectHandler
	CreateProject(ctx context.Context, project *domain.CreateProjectDTO) (string, error)

	GetAllProjects(ctx context.Context) ([]domain.ProjectDB, error)
	GetProjectByID(ctx context.Context, projectID string) (*domain.ProjectDB, error)
	GetProjectByCustomerLastname(ctx context.Context, lastname string) ([]domain.ProjectDB, error)
	GetProjectByAddress(ctx context.Context, address string) ([]domain.ProjectDB, error)
	GetAllCustomerLastnames(ctx context.Context) ([]string, error)
	GetAllAddresses(ctx context.Context) ([]string, error)

	//UserHandler
	CreateUser(ctx context.Context, user *domain.CreateUserDTO) (string, error)
}

type FileStorage interface {
	UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
}
