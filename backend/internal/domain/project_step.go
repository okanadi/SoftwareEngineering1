package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectStepDTO struct {
	ProjectID   string `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	//progress wird initial im Repo auf 'geplant' gesetzt
}

type ProjectStepDB struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ProjectID   uuid.UUID  `db:"project_id" json:"project_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	StartDate   *time.Time `db:"start_date" json:"start_date"`
	EndDate     *time.Time `db:"end_date" json:"end_date"`
	Progress    string     `db:"progress" json:"progress"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
}
