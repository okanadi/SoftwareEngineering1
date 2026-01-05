package adapter

import (
	"backend/internal/port"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	client     *s3.Client
	bucketName string
}

func NewS3Storage(ctx context.Context, bucketName string, region string) (port.FileStorage, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	return &S3Storage{
		client:     s3.NewFromConfig(cfg),
		bucketName: bucketName,
	}, nil
}

func (s *S3Storage) UploadFile(ctx context.Context, file io.Reader, filename string, contentType string) (string, error) {
	// Eindeutiger Name: uploads/2024-01-01/123456_bild.jpg
	key := fmt.Sprintf("uploads/%s/%d_%s", time.Now().Format("2006-01-02"), time.Now().Unix(), filename)

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("S3 upload failed: %w", err)
	}
	return key, nil
}
