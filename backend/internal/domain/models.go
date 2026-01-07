package domain

import (
	"io"
	"time"

	"github.com/google/uuid"
)

// --- DB Models (entsprechen den Tabellen) ---

type ProjectStep struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	ProjectID   uuid.UUID  `db:"project_id" json:"project_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	StartDate   *time.Time `db:"start_date" json:"start_date"`
	EndDate     *time.Time `db:"end_date" json:"end_date"`
	Progress    string     `db:"progress" json:"progress"`
}

// --- DTOs (Data Transfer Objects) für API Requests ---

// UpdateStepInput: Daten, die der Handwerker sendet
type UpdateStepInput struct {
	StepID      string
	UserID      string // Kommt später aus dem Token, aktuell Dummy
	NewStatus   string // "in_arbeit", "fertiggestellt", "problem"
	Note        string
	File        io.Reader // Der Datei-Stream
	FileName    string
	ContentType string
}

// LoginRequest: Für den Kunden-Login
type CustomerLoginRequest struct {
	OrderID  string `json:"order_id"` // Die UUID des Projekts
	Lastname string `json:"lastname"`
}
