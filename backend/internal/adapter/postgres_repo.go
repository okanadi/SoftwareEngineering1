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
