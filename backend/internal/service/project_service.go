package service

import (
	"archive/zip"
	"backend/internal/domain"
	"backend/internal/port"
	"context"
	"fmt"
	"net/http"
)

type ProjectService struct {
	repo    port.ProjectRepository
	storage port.FileStorage
}

func NewProjectService(repo port.ProjectRepository, storage port.FileStorage) *ProjectService {
	return &ProjectService{repo: repo, storage: storage}
}

func (s *ProjectService) CreateProject(ctx context.Context, input domain.CreateProjectDTO) (string, error) {
	if input.CustomerLastname == "" || input.Address == "" {
		return "", fmt.Errorf("Nachname und Adresse sind Pflichtfelder")
	}

	return s.repo.CreateProject(ctx, &input)
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]domain.ProjectDB, error) {
	return s.repo.GetAllProjects(ctx)
}

func (s *ProjectService) GetProjectByID(ctx context.Context, projectID string) (*domain.ProjectDB, error) {
	if projectID == "" {
		return nil, fmt.Errorf("Projekt ID darf nicht leer sein")
	}

	return s.repo.GetProjectByID(ctx, projectID)
}

func (s *ProjectService) GetProjectByCustomerLastname(ctx context.Context, lastname string) ([]domain.ProjectDB, error) {
	if lastname == "" {
		return nil, fmt.Errorf("Nachname darf nicht leer sein")
	}

	return s.repo.GetProjectByCustomerLastname(ctx, lastname)
}

func (s *ProjectService) GetProjectByAddress(ctx context.Context, address string) ([]domain.ProjectDB, error) {
	if address == "" {
		return nil, fmt.Errorf("Adresse darf nicht leer sein")
	}

	return s.repo.GetProjectByAddress(ctx, address)
}

func (s *ProjectService) GetAllCustomerLastnames(ctx context.Context) ([]string, error) {
	return s.repo.GetAllCustomerLastnames(ctx)
}

func (s *ProjectService) GetAllAddresses(ctx context.Context) ([]string, error) {
	return s.repo.GetAllAddresses(ctx)
}

func (s *ProjectService) GetByManagerID(ctx context.Context, managerId string) ([]domain.ProjectDB, error) {
	return s.repo.GetByManagerID(ctx, managerId)
}

func (s *ProjectService) UpdateProject(ctx context.Context, input domain.UpdateProjectDTO) error {
	// Hier könnten Validierungen stehen (z.B. existiert der Manager?)
	return s.repo.UpdateProject(ctx, &input)
}

func (s *ProjectService) ExportProjectAsZip(ctx context.Context, projectID string, w http.ResponseWriter) error {
	// 1. Daten aus der DB holen (unsere vorhandene SQL-Funktion)
	steps, err := s.repo.GetHistory(ctx, projectID)
	if err != nil {
		return err
	}

	// 2. ZIP-Writer auf den ResponseWriter legen
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// 3. Einen Text-Bericht generieren (Zusammenfassung der Historie)
	report := "PROJEKT EXPORT PROTOKOLL\n========================\n"

	for _, step := range steps {
		report += fmt.Sprintf("\nSchritt: %s (Status: %s)\n", step.Title, step.Progress)

		for _, entry := range step.History {
			report += fmt.Sprintf("- [%s] %s: %s\n",
				entry.Timestamp.Format("02.01.2006 15:04"),
				entry.UserName,
				entry.Note)

			// 4. Bilder für diesen History-Eintrag hinzufügen
			for i, photo := range entry.Photos {
				// Bild von S3 streamen
				imgBytes, err := s.storage.DownloadFile(ctx, photo.S3Key)
				if err != nil {
					continue // Falls ein Bild fehlt, trotzdem weitermachen
				}

				// Dateiname im ZIP: "Bilder/Schrittname_Datum_Index.jpg"
				fileName := fmt.Sprintf("Bilder/%s_%s_%d.jpg",
					step.Title,
					entry.Timestamp.Format("2006-01-02"),
					i)

				f, _ := zipWriter.Create(fileName)
				f.Write(imgBytes)
			}
		}
	}

	// 5. Das Protokoll als Text-Datei ins ZIP legen
	f, _ := zipWriter.Create("protokoll.txt")
	f.Write([]byte(report))

	return nil
}
