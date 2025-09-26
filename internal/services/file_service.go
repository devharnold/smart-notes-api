package main

import (
	"context"
	"smart-notes-api/internal/repositories"
	"smart-notes-api/internal/storage"
)

type FileService struct {
	Repo *repositories.FileMeta
	S3	*storage.S3Client
}

func (s *FileService) UploadNote(ctx context.Context, Title, content []byte) error {
	// upload the notes file to S3
	key := "uploads/" + filename
	_, err := s.S3.UploadFile(ctx, key, content)
	if err != nil {
		return err
	}

	// save the metadata in postgres
	meta := repositories.FileMeta{
		Tile:	title,
	}
	return s.Repo.Save(ctx, meta)
}