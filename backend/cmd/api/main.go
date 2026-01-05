package main

import (
	"backend/internal/adapter"
	"backend/internal/service"
	"backend/internal/web"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib" // Postgres Driver
	"github.com/jmoiron/sqlx"
)

func main() {
	// 1. Umgebungsvariablen laden
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := "db" // Im Docker Container heißt der Host "db"
	if os.Getenv("GO_ENV") == "local" {
		dbHost = "localhost" // Für lokales Testen ohne Docker Network
	}

	bucketName := os.Getenv("AWS_BUCKET_NAME")
	awsRegion := os.Getenv("AWS_REGION")

	// 2. DB Verbindung
	dsn := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":5432/" + dbName + "?sslmode=disable"
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("Konnte DB nicht verbinden (%s): %v", dsn, err)
	}
	defer db.Close()
	log.Println("Datenbank verbunden.")

	// 3. S3 Setup
	s3Storage, err := adapter.NewS3Storage(context.Background(), bucketName, awsRegion)
	if err != nil {
		log.Fatalf("Konnte S3 nicht initialisieren: %v", err)
	}

	// 4. Wiring (Dependency Injection)
	repo := adapter.NewPostgresRepo(db)
	svc := service.NewProjectService(repo, s3Storage)
	handler := web.NewProjectHandler(svc)

	// 5. Router Setup
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS (wichtig für Frontend Zugriff)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// API Routen
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Smart Builders API is running!"))
		})

		r.Post("/customer/login", handler.HandleCustomerLogin)
		r.Get("/projects/{id}/steps", handler.HandleGetSteps)
		r.Post("/steps/{id}/update", handler.HandleUpdateStep)
	})

	log.Println("Server startet auf :8080")
	http.ListenAndServe(":8080", r)
}
