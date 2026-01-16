package main

import (
	"backend/internal/adapter"
	"backend/internal/service"
	"backend/internal/web"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib" // Postgres Driver
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("database/.env")
	if err != nil {
		log.Println("Info: Keine .env Datei in 'database/.env' gefunden oder Fehler beim Laden. Nutze System-Umgebungsvariablen.")
	} else {
		log.Println("Konfiguration aus database/.env geladen.")
	}

	// Variablen auslesen (egal ob aus .env oder System-Env)
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST") // Lokal: "localhost", Docker: "db"

	if dbHost == "" {
		dbHost = "localhost"
	}

	awsBucket := os.Getenv("AWS_BUCKET_NAME")
	awsRegion := os.Getenv("AWS_REGION")

	// 2. DB Verbindung
	// DSN bauen
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=require", dbUser, dbPass, dbHost, dbName)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Konnte DB nicht verbinden. DSN: postgres://%s:***@%s:5432/%s Error: %v", dbUser, dbHost, dbName, err)
	}
	defer db.Close()
	log.Println("Datenbank verbunden auf Host:", dbHost)

	// 3. S3 Setup
	s3Storage, err := adapter.NewS3Storage(context.Background(), awsBucket, awsRegion)
	if err != nil {
		log.Fatalf("Konnte S3 nicht initialisieren: %v", err)
	}

	// 4. Wiring (Dependency Injection)
	repo := adapter.NewPostgresRepo(db, s3Storage)

	userService := service.NewUserService(repo, s3Storage)
	projectService := service.NewProjectService(repo, s3Storage)
	projectStepService := service.NewProjectStepService(repo, s3Storage)
	historyService := service.NewHistoryService(repo, s3Storage)

	userHandler := web.NewUserHandler(userService)
	projectHandler := web.NewProjectHandler(projectService)
	projectStepHandler := web.NewProjectStepHandler(projectStepService)
	historyHandler := web.NewHistoryHandler(historyService)

	// 5. Router Setup
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "3600")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	//Routes
	r.Route("/api/v1", func(r chi.Router) {

		r.Post("/users/create", userHandler.HandleCreateUser)
		r.Post("/users/login", userHandler.HandleUserLogin)
		r.Get("/users/getAll", userHandler.HandleGetAllUsers)

		r.Post("/projects/create", projectHandler.HandleCreateProject)
		r.Get("/projects/export/{projectID}", projectHandler.HandleExportProjectZip)
		r.Put("/projects/editProject/{projectID}", projectHandler.HandleUpdateProject)
		r.Get("/projects/getAll", projectHandler.HandleGetAllProjects)
		r.Get("/projects/getByID/{id}", projectHandler.HandleGetProjectByID)
		r.Get("/projects/getAllCustomerLastnames", projectHandler.HandleGetAllCustomerLastnames)
		r.Get("/projects/getAllAddresses", projectHandler.HandleGetAllAddresses)
		r.Get("/projects/getByManagerID/{managerID}", projectHandler.HandleGetByManagerID)
		r.Get("/projects/getByCustomerLastname/{lastname}", projectHandler.HandleGetProjectByCustomerLastname)
		r.Get("/projects/getByAddress/{address}", projectHandler.HandleGetProjectByAddress)

		r.Get("/project-steps/getAllByProjectID/{projectID}", projectStepHandler.HandleGetProjectSteps)
		r.Get("/project-steps/getByID/{projectID}/{stepID}", projectStepHandler.HandleGetProjectStepByID)
		r.Post("/project-steps/create", projectStepHandler.HandleCreateProjectStep)
		r.Post("/project-steps/updateProgress/{stepID}", projectStepHandler.HandleUpdateStepProgress)

		r.Get("/history/getHistory/{projectID}", historyHandler.HandleGetHistory)

	})

	frontendRouter := chi.NewRouter()
	frontendPath := "./frontend/dist"
	fileServer := http.FileServer(http.Dir(frontendPath))

	frontendRouter.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		fullPath := frontendPath + r.URL.Path
		info, err := os.Stat(fullPath)

		if os.IsNotExist(err) || info.IsDir() {
			http.ServeFile(w, r, frontendPath+"/index.html")
			return
		}
		fileServer.ServeHTTP(w, r)
	})

	go func() {
		log.Println("Frontend startet auf Port 8040")
		if err := http.ListenAndServe(":8040", frontendRouter); err != nil {
			log.Fatalf("Frontend-Server Fehler: %v", err)
		}
	}()

	log.Println("Server startet auf :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Backend-Server Fehler: %v", err)
	}
}
