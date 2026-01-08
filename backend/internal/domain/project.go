package domain

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectDTO struct {
	ManagerID        string `json:"manager_id"`
	CustomerLastname string `json:"customer_lastname"`
	Address          string `json:"address"`
	Description      string `json:"description"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	//progress wird initial im Repo auf 'geplant' gesetzt
	//CreatedAt wird automatisch gesetzt
}

type ProjectDB struct {
	ID               uuid.UUID  `db:"id" json:"id"`
	ManagerID        uuid.UUID  `db:"manager_id" json:"manager_id"`
	CustomerLastname string     `db:"customer_lastname" json:"customer_lastname"`
	Address          string     `db:"address" json:"address"`
	Description      string     `db:"description" json:"description"`
	StartDate        *time.Time `db:"start_date" json:"start_date"`
	EndDate          *time.Time `db:"end_date" json:"end_date"`
	Progress         string     `db:"progress" json:"progress"`
	CreatedAt        *time.Time `db:"created_at" json:"created_at"`
}
