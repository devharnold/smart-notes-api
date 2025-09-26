package repositories


import (
	"context"
	"smart-notes-api/internal/storage/db"
)

type FileMeta struct {
	ID			int
	UserID  	int
	FileName	string
	S3Key		string
}

type FileRepo struct {}

func (r *FileRepo) Save(ctx context.Context, f FileMeta) error {
	_, err := db.Pool.Exec(
		ctx,
		"INSERT INTO files (user_id, filename, s3_key) VALUES ($1, $2, $3)",
		f.UserID, f.FileName, f.S3Key,
	)
	return err
}