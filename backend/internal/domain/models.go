package domain

import (
	"io"
)

// --- DB Models (entsprechen den Tabellen) ---

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
