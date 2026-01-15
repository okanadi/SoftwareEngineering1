package port

import (
	"backend/internal/domain"
	"context"
	"io"
)

type ProjectRepository interface {
	//UserHandler
	CreateUser(ctx context.Context, user *domain.CreateUserDTO) (string, error)
	UserLogin(ctx context.Context, input *domain.UserLoginDTO) (*domain.UserDB, error)
	GetAllUsers(ctx context.Context) ([]domain.AllUsersDTO, error)

	//ProjectHandler
	CreateProject(ctx context.Context, project *domain.CreateProjectDTO) (string, error)
	GetByManagerID(ctx context.Context, managerId string) ([]domain.ProjectDB, error)
	GetAllProjects(ctx context.Context) ([]domain.ProjectDB, error)
	GetProjectByID(ctx context.Context, projectID string) (*domain.ProjectDB, error)
	GetProjectByCustomerLastname(ctx context.Context, lastname string) ([]domain.ProjectDB, error)
	GetProjectByAddress(ctx context.Context, address string) ([]domain.ProjectDB, error)
	GetAllCustomerLastnames(ctx context.Context) ([]string, error)
	GetAllAddresses(ctx context.Context) ([]string, error)

	//ProjectStepHandler
	CreateProjectStep(ctx context.Context, step *domain.CreateProjectStepDTO) (string, error)
	GetStepsProjectByProjectID(ctx context.Context, projectID string) ([]domain.ProjectStepDB, error)
	UpdateStepWithHistoryAndMedia(ctx context.Context, stepID, userID, status, note, s3Key, fileType string) error
}

type FileStorage interface {
	UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error)
	GetPresignedURL(ctx context.Context, key string) (string, error)
	DownloadFile(ctx context.Context, key string) ([]byte, error)
}
