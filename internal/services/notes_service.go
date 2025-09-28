package main

import (
	"context"
	"smart-notes-api/internal/repositories"
	"smart-notes-api/internal/storage"
)

type NotesService struct {
	Repo *repositories.NotesMeta
	S3   *storage.S3Client
}

func (s *NotesService) UploadNote(ctx context.Context, Title, content []byte) error {
	// upload the notes file to S3
	key := "uploads/" + filename
	_, err := s.S3.UploadFile(ctx, key, content)
	if err != nil {
		return err
	}

	// save the metadata in postgres
	meta := repositories.NotesMeta{
		Tile: title,
	}
	return s.Repo.Save(ctx, meta)
}
