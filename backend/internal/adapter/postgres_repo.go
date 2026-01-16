package adapter

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type PostgresRepo struct {
	db      *sqlx.DB
	storage port.FileStorage
}

func NewPostgresRepo(db *sqlx.DB, storage port.FileStorage) port.ProjectRepository {
	return &PostgresRepo{db: db, storage: storage}
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

func (r *PostgresRepo) GetAllUsers(ctx context.Context) ([]domain.AllUsersDTO, error) {
	query := `
        SELECT id, email, name, role
        FROM users
    `
	var users []domain.AllUsersDTO
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("Get all users failed: %w", err)
	}
	return users, nil
}

func (r *PostgresRepo) GetUserByID(ctx context.Context, userID string) (domain.UserDB, error) {
	query := `
        SELECT id, email, name, role
        FROM users
    	WHERE id = $1
		LIMIT 1
	`

	var user domain.UserDB
	err := r.db.GetContext(ctx, &user, query, userID)

	if err != nil {
		return user, fmt.Errorf("Get user failed: %w", err)
	}

	return user, nil

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

func (r *PostgresRepo) GetByManagerID(ctx context.Context, managerId string) ([]domain.ProjectDB, error) {
	query := `
		SELECT *
		FROM projects
		WHERE manager_id = $1
	`

	var projects []domain.ProjectDB
	err := r.db.SelectContext(ctx, &projects, query, managerId)

	if err != nil {
		return nil, fmt.Errorf("Get projects by manager id failed: %w", err)
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

func (r *PostgresRepo) CreateProjectHistoryEntry(ctx context.Context, history *domain.HistoryDB) (string, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("Failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	var historyID string
	query := `
		INSERT INTO project_history (step_id, user_id, new_status, note)
		VALUES ($1, NULLIF($2, '')::uuid, $3, $4)
		RETURNING id`

	err = tx.QueryRowContext(ctx, query,
		history.StepId,
		history.UserName,
		history.Status,
		history.Note,
	).Scan(&historyID)
	if err != nil {
		return "", fmt.Errorf("Create project history entry failed: %w", err)
	}

	//Hier weitermachen
	return "", nil
}

func (r *PostgresRepo) CreateHistoryEntry(ctx context.Context, stepID string, userID string, status string, note string) (string, error) {
	query := `
        INSERT INTO project_history (step_id, user_id, new_status, note)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var historyID string
	err := r.db.QueryRowContext(ctx, query, stepID, userID, status, note).Scan(&historyID)
	return historyID, err
}

func (r *PostgresRepo) CreateMedia(ctx context.Context, historyID string, s3Key string, fileType string) error {
	query := `INSERT INTO media (history_id, s3_key, file_type) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, historyID, s3Key, fileType)
	return err
}

func (r *PostgresRepo) UpdateStepWithHistoryAndMedia(ctx context.Context, stepID, userID, status, note, s3Key, fileType string) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	// Sicherstellen, dass bei einem Fehler zurückgerollt wird
	defer tx.Rollback()

	// 1. Projektschritt aktualisieren
	_, err = tx.ExecContext(ctx, `
		UPDATE project_steps 
		SET progress = $1 
		WHERE id = $2`,
		status, stepID)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}

	// 2. History-Eintrag schreiben
	var historyID string
	err = tx.QueryRowContext(ctx, `
		INSERT INTO project_history (step_id, user_id, new_status, note)
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		stepID, userID, status, note).Scan(&historyID)
	if err != nil {
		return fmt.Errorf("failed to insert history: %w", err)
	}

	// 3. Media-Eintrag schreiben (nur wenn ein Bild hochgeladen wurde)
	if s3Key != "" {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO media (history_id, s3_key, file_type)
			VALUES ($1, $2, $3)`,
			historyID, s3Key, fileType)
		if err != nil {
			return fmt.Errorf("failed to insert media: %w", err)
		}
	}

	// 4. Transaktion abschließen
	return tx.Commit()
}

func (r *PostgresRepo) GetHistory(ctx context.Context, projectID string) ([]domain.ProjectStepHistoryDTO, error) {
	query := `
		WITH media_agg AS (
			SELECT 
				history_id, 
				jsonb_agg(jsonb_build_object(
					'id', id,
					's3_key', s3_key,
					'file_type', file_type
				)) AS photos
			FROM media
			GROUP BY history_id
		),
		history_agg AS (
			SELECT 
				h.step_id,
				jsonb_agg(jsonb_build_object(
					'id', h.id,
					'status', h.new_status,
					'note', h.note,
					'user_name', COALESCE(u.name, 'Unbekannt'),
					'timestamp', h.timestamp,
					'photos', COALESCE(m.photos, '[]'::jsonb)
				) ORDER BY h.timestamp DESC) AS entries
			FROM project_history h
			LEFT JOIN users u ON h.user_id = u.id
			LEFT JOIN media_agg m ON h.id = m.history_id
			GROUP BY h.step_id
		)
		SELECT 
			ps.id, ps.project_id, ps.title, ps.description, 
			ps.start_date, ps.end_date, ps.progress, ps.created_at,
			COALESCE(ha.entries, '[]'::jsonb) as history_json
		FROM project_steps ps
		LEFT JOIN history_agg ha ON ps.id = ha.step_id
		WHERE ps.project_id = $1::uuid
		ORDER BY ps.created_at ASC
	`

	var results []struct {
		domain.ProjectStepDB
		HistoryJSON []byte `db:"history_json"`
	}

	err := r.db.SelectContext(ctx, &results, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch project steps with history: %w", err)
	}

	finalSteps := make([]domain.ProjectStepHistoryDTO, len(results))

	for i, res := range results {
		finalSteps[i].ProjectStepDB = res.ProjectStepDB

		if err := json.Unmarshal(res.HistoryJSON, &finalSteps[i].History); err != nil {
			return nil, fmt.Errorf("failed to unmarshal history for step %s: %w", res.ID, err)
		}

		for j := range finalSteps[i].History {
			for k := range finalSteps[i].History[j].Photos {
				photo := &finalSteps[i].History[j].Photos[k]

				url, err := r.storage.GetPresignedURL(ctx, photo.S3Key)
				if err != nil {
					fmt.Printf("Warning: could not generate URL for key %s: %v\n", photo.S3Key, err)
					continue
				}

				photo.Url = url
			}
		}
	}

	return finalSteps, nil
}

func (r *PostgresRepo) UpdateProject(ctx context.Context, p *domain.UpdateProjectDTO) error {
	query := `
		UPDATE projects 
		SET manager_id = $1, 
		    customer_lastname = $2, 
		    address = $3, 
		    description = $4, 
		    start_date = NULLIF($5, '')::date, 
		    end_date = NULLIF($6, '')::date, 
		    progress = $7
		WHERE id = $8
	`
	_, err := r.db.ExecContext(ctx, query,
		p.ManagerID,
		p.CustomerLastname,
		p.Address,
		p.Description,
		p.StartDate,
		p.EndDate,
		p.Progress,
		p.ID,
	)
	if err != nil {
		return fmt.Errorf("update project failed: %w", err)
	}
	return nil
}
