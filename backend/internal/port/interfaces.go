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

	//UserHandler
	CreateUser(ctx context.Context, user *domain.CreateUserDTO) (string, error)
}

type FileStorage interface {
	UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
}
