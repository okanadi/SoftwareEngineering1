package adapter

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"backend/internal/port"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	client     *s3.Client
	bucketName string
}

// NewS3Storage initialisiert die Verbindung.
// Es liest automatisch die Env-Vars (AWS_ACCESS_KEY_ID etc.)
func NewS3Storage(ctx context.Context, bucketName string, region string) (port.FileStorage, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Storage{
		client:     client,
		bucketName: bucketName,
	}, nil
}

func (s *S3Storage) UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error) {
	// Wir bauen eine Ordnerstruktur: uploads/JAHR-MONAT-TAG/TIMESTAMP_Dateiname
	// Das verhindert Namenskonflikte.
	datePath := time.Now().Format("2006-01-02")
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("uploads/%s/%d_%s", datePath, timestamp, filename)

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		log.Printf("Fehler beim Upload zu Bucket %s: %v", s.bucketName, err)
		return "", err
	}

	return key, nil
}

func (s *S3Storage) GetPresignedURL(ctx context.Context, key string) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	// URL ist 15 Minuten gültig
	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *S3Storage) DownloadFile(ctx context.Context, key string) ([]byte, error) {
	resp, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
