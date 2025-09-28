package storage

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Uploader *manager.Uploader
	Bucket   string
}

func NewS3Client() *S3Client {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Unable to log aws sdk config: %v", err)
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		log.Fatal("Unable to find bucket. Consider creating one")
	}

	return &S3Client{
		Uploader: uploader,
		Bucket:   bucket,
	}
}

func (s *S3Client) UploadFile(ctx context.Context, key string, body []byte) (string, error) {
	_, err := s.Uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return "", err
	}

	return key, nil
}
