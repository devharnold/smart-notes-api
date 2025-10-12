package repositories

import (
	"context"
	"smart-notes-api/internal/storage"
)

type NotesMeta struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	S3Key  string `json:"s3_key"`
}

//type NotesRepo struct{}

func (r *NotesMeta) Save(ctx context.Context, f NotesMeta) error {
	insertQuery := "INSERT INTO notes(user_id, title, s3_key) VALUES ($1, $2, $3) RETURNING id"

	err := storage.Pool.QueryRow(ctx, insertQuery, f.UserID, f.Title, f.S3Key).Scan(&f.ID)
	if err != nil {
		return err
	}
	return nil
}
