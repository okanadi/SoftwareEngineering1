package port

import (
	"backend/internal/domain"
	"context"
	"io"
)

type ProjectRepository interface {
	//UserHandler
	CreateUser(ctx context.Context, user *domain.CreateUserDTO) (string, error)

	//ProjectHandler
	CreateProject(ctx context.Context, project *domain.CreateProjectDTO) (string, error)

	GetAllProjects(ctx context.Context) ([]domain.ProjectDB, error)

	//ProjectStepHandler
	CreateProjectStep(ctx context.Context, step *domain.CreateProjectStepDTO) (string, error)
	GetStepsProjectByProjectID(ctx context.Context, projectID string) ([]domain.ProjectStepDB, error)
}

type FileStorage interface {
	UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
}
