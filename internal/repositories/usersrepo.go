package repositories

import (
	"context"
	"time"

	"smart-notes-api/internal/models"
	"smart-notes-api/internal/storage"
)

type UsersRepository struct{}

func (r *UsersRepository) UploadUser(ctx context.Context, users models.Users) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO users(user_email, first_name, last_name, password) VALUES ($1, $2, $3, $4)"

	// execute query
	_, err := storage.Pool.Exec(ctx, insertQuery, users.UserEmail, users.FirstName, users.LastName, users.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var user models.User

	selectQuery := "SELECT * FROM users WHERE user_email = $1"

	err := storage.Pool.QueryRow(ctx, selectQuery, email).Scan(
		&user.UserEmail,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
