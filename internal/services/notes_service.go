package services

import (
	"context"
	"fmt"
	"io"
	"smart-notes-api/internal/repositories"
	"smart-notes-api/internal/storage"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type NotesService struct {
	repo *repositories.NotesMeta
	S3   *storage.S3Client
}

func NewNotesService(s3 *storage.S3Client, repo *repositories.NotesMeta) *NotesService {
	return &NotesService{
		repo: repo,
		S3:   s3,
	}
}

func (s *NotesService) SaveNote(ctx context.Context, userID int, title string, content string) (string, error) {
	// first insert into DB, to help us get the noteID
	meta := repositories.NotesMeta{
		UserID: userID,
		Title:  title,
	}

	// insert without S3key first because we dont know it yet
	insertQuery := "INSERT INTO notes(user_id, title) VALUES ($1, $2) RETURNING id"
	err := storage.Pool.QueryRow(ctx, insertQuery, meta.UserID, meta.Title).Scan(&meta.ID)
	if err != nil {
		return "", err
	}

	// Form a deterministic S3Key using UserID and NoteID
	key := fmt.Sprintf("notes/%d/%d.txt", userID, meta.ID)

	// upload note content to S3
	_, err = s.S3.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.S3.Bucket),
		Key:         aws.String(key),
		Body:        strings.NewReader(content),
		ContentType: aws.String("text/plain"),
	})
	if err != nil {
		return "", err
	}

	// update record with s3 key
	updateQuery := "UPDATE notes SET s3_key = $1 WHERE id = $2"
	_, err = storage.Pool.Exec(ctx, updateQuery, key, meta.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("s3://%s/%s", s.S3.Bucket, key), nil
}

func (s *NotesService) RetrieveNote(ctx context.Context, UserID int, title string) (string, error) {
	// step1: Fetch s3_key for the note
	var s3key string
	SelectQuery := "SELECT s3_key FROM notes WHERE user_id = $1 AND title = $2"
	err := storage.Pool.QueryRow(ctx, SelectQuery, UserID, title).Scan(&s3key)
	if err != nil {
		return "", err
	}

	output, err := s.S3.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.S3.Bucket),
		Key:    aws.String(s3key),
	})
	if err != nil {
		return "", err
	}
	defer output.Body.Close()

	body, err := io.ReadAll(output.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Update the existing text in S3, but we keep the same key
func (s *NotesService) UpdateNote(ctx context.Context, userID int, title string, newContent string) error {
	var s3Key string
	query := "SELECT s3_key FROM notes WHERE user_id = $1 AND title = $2"
	if err := storage.Pool.QueryRow(ctx, query, userID, title).Scan(&s3Key); err != nil {
		return err
	}

	_, err := s.S3.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.S3.Bucket),
		Key:         aws.String(s3Key),
		Body:        strings.NewReader(newContent),
		ContentType: aws.String("text/plain"),
	})
	return err
}

func (s *NotesService) DeleteNote(ctx context.Context, userID int, title string) error {
	var s3key string
	selectQuery := "SELECT s3_key FROM notes WHERE user_id = $1 AND title = $2"
	if err := storage.Pool.QueryRow(ctx, selectQuery, userID, title).Scan(&s3key); err != nil {
		return err
	}

	_, err := s.S3.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.S3.Bucket),
		Key:    aws.String(s3key),
	})
	if err != nil {
		return err
	}

	// also delete record from postgres
	delQuery := "DELETE FROM notes WHERE user_id = $1 AND title = $2"
	_, err = storage.Pool.Exec(ctx, delQuery, userID, title)
	if err != nil {
		return err
	}
	return nil
}
