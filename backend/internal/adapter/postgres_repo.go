package adapter

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) port.ProjectRepository {
	return &PostgresRepo{db: db}
}

// User
func (r *PostgresRepo) CreateUser(ctx context.Context, user *domain.CreateUserDTO) (string, error) {
	query := `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var newID string

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Fehler beim hashen des Passworts: %w", err)
	}

	// Konvertiere das byte-slice zurück in einen String für die DB
	hashedPassword := string(hashedBytes)

	err = r.db.QueryRowContext(ctx, query,
		user.Name,
		user.Email,
		hashedPassword,
		user.Role,
	).Scan(&newID)

	if err != nil {
		return "", fmt.Errorf("create user failed: %w", err)
	}

	return newID, nil
}

func (r *PostgresRepo) UserLogin(ctx context.Context, input *domain.UserLoginDTO) (*domain.UserDB, error) {
	query := `
        SELECT *
        FROM users
        WHERE email = $1
    `
	var user domain.UserDB

	err := r.db.GetContext(ctx, &user, query, input.Email)
	if err != nil {
		return nil, fmt.Errorf("Benutzer mit dieser E-Mail nicht gefunden: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return nil, fmt.Errorf("Falsches Passwort")
	}

	return &user, nil
}

// Project
func (r *PostgresRepo) CreateProject(ctx context.Context, project *domain.CreateProjectDTO) (string, error) {
	query := `
		INSERT INTO projects (manager_id, customer_lastname, address, description, start_date, end_date, progress)
		VALUES ($1, $2, $3, $4, $5, $6, 'geplant')
		RETURNING id
	`

	var start, end *string
	if project.StartDate != "" {
		start = &project.StartDate
	}
	if project.EndDate != "" {
		end = &project.EndDate
	}

	var newID string
	err := r.db.QueryRowContext(ctx, query,
		project.ManagerID,
		project.CustomerLastname,
		project.Address,
		project.Description,
		start,
		end,
	).Scan(&newID)

	if err != nil {
		return "", fmt.Errorf("create project failed: %w", err)
	}

	return newID, nil
}

func (r *PostgresRepo) GetAllProjects(ctx context.Context) ([]domain.ProjectDB, error) {
	query := `
		SELECT *
		FROM projects
		ORDER BY created_at DESC
	`
	var projects []domain.ProjectDB
	err := r.db.SelectContext(ctx, &projects, query)
	if err != nil {
		return nil, fmt.Errorf("Get all projects failed: %w", err)
	}
	return projects, nil
}

// Project Step
func (r *PostgresRepo) CreateProjectStep(ctx context.Context, step *domain.CreateProjectStepDTO) (string, error) {
	query := `
		INSERT INTO project_steps (project_id, title, description, start_date, end_date, progress)
		VALUES ($1, $2, $3, $4, $5, 'geplant')
		RETURNING id
	`
	var start, end *string
	if step.StartDate != "" {
		start = &step.StartDate
	}
	if step.EndDate != "" {
		end = &step.EndDate
	}
	var newID string
	err := r.db.QueryRowContext(ctx, query,
		step.ProjectID,
		step.Title,
		step.Description,
		start,
		end,
	).Scan(&newID)
	if err != nil {
		return "", fmt.Errorf("create project step failed: %w", err)
	}

	return newID, nil
}

func (r *PostgresRepo) GetStepsProjectByProjectID(ctx context.Context, projectID string) ([]domain.ProjectStepDB, error) {
	query := `
		SELECT *
		FROM project_steps
		WHERE project_id = $1
		ORDER BY created_at DESC
	`
	var steps []domain.ProjectStepDB
	err := r.db.SelectContext(ctx, &steps, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("Get steps by project ID failed: %w", err)
	}
	return steps, nil
}
func (r *PostgresRepo) GetProjectByID(ctx context.Context, projectID string) (*domain.ProjectDB, error) {
	query := `
		SELECT *
		FROM projects
		WHERE id = $1
	`

	var project domain.ProjectDB
	err := r.db.GetContext(ctx, &project, query, projectID)

	if err != nil {
		return nil, fmt.Errorf("Get project by ID failed: %w", err)
	}

	return &project, nil
}

func (r *PostgresRepo) GetProjectByCustomerLastname(ctx context.Context, lastname string) ([]domain.ProjectDB, error) {
	query := `
		SELECT *
		FROM projects
		WHERE customer_lastname ILIKE $1
		ORDER BY created_at DESC
	`

	var projects []domain.ProjectDB
	err := r.db.SelectContext(ctx, &projects, query, "%"+lastname+"%")

	if err != nil {
		return nil, fmt.Errorf("Get project by customer lastname failed: %w", err)
	}

	return projects, nil
}

func (r *PostgresRepo) GetProjectByAddress(ctx context.Context, address string) ([]domain.ProjectDB, error) {
	query := `
		SELECT *
		FROM projects
		WHERE address ILIKE $1
		ORDER BY created_at DESC
	`

	var projects []domain.ProjectDB
	err := r.db.SelectContext(ctx, &projects, query, "%"+address+"%")

	if err != nil {
		return nil, fmt.Errorf("Get project by address failed: %w", err)
	}

	return projects, nil
}

func (r *PostgresRepo) GetAllCustomerLastnames(ctx context.Context) ([]string, error) {
	query := `
		SELECT DISTINCT customer_lastname
		FROM projects
		ORDER BY customer_lastname
	`

	var lastnames []string
	err := r.db.SelectContext(ctx, &lastnames, query)
	if err != nil {
		return nil, fmt.Errorf("Get all customer lastnames failed: %w", err)
	}
	return lastnames, nil
}

func (r *PostgresRepo) GetAllAddresses(ctx context.Context) ([]string, error) {
	query := `
		SELECT DISTINCT address
		FROM projects
		ORDER BY address
	`

	var addresses []string
	err := r.db.SelectContext(ctx, &addresses, query)
	if err != nil {
		return nil, fmt.Errorf("Get all addresses failed: %w", err)
	}
	return addresses, nil
}
