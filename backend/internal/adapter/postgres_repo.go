package adapter

import (
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) port.ProjectRepository {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetProjectByLogin(ctx context.Context, id string, lastname string) (*domain.Project, error) {
	var p domain.Project
	// Login via ID (Auftragsnummer) + Nachname
	query := `SELECT * FROM projects WHERE id = $1 AND customer_lastname = $2`
	err := r.db.GetContext(ctx, &p, query, id, lastname)
	return &p, err
}

func (r *PostgresRepo) GetStepsByProjectID(ctx context.Context, projectID string) ([]domain.ProjectStep, error) {
	var steps []domain.ProjectStep
	query := `SELECT * FROM project_steps WHERE project_id = $1 ORDER BY start_date ASC`
	err := r.db.SelectContext(ctx, &steps, query, projectID)
	return steps, err
}

func (r *PostgresRepo) AddHistoryEntry(ctx context.Context, stepID string, userID string, status string, note string, s3Key string, fileType string) error {
	// Transaktion starten
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback bei Fehler

	// 1. History Eintrag
	var historyID string
	queryHistory := `
		INSERT INTO project_history (step_id, user_id, new_status, note)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	// Falls userID leer ist (z.B. Test), setzen wir NULL
	var uid *string
	if userID != "" {
		uid = &userID
	}

	err = tx.QueryRowContext(ctx, queryHistory, stepID, uid, status, note).Scan(&historyID)
	if err != nil {
		return fmt.Errorf("insert history failed: %w", err)
	}

	// 2. Media Eintrag (falls Bild vorhanden)
	if s3Key != "" {
		queryMedia := `INSERT INTO media (history_id, s3_key, file_type) VALUES ($1, $2, $3)`
		_, err = tx.ExecContext(ctx, queryMedia, historyID, s3Key, fileType)
		if err != nil {
			return fmt.Errorf("insert media failed: %w", err)
		}
	}

	// 3. Step Status aktualisieren
	queryStepUpdate := `UPDATE project_steps SET progress = $1 WHERE id = $2`
	_, err = tx.ExecContext(ctx, queryStepUpdate, status, stepID)
	if err != nil {
		return fmt.Errorf("update step failed: %w", err)
	}

	return tx.Commit()
}
