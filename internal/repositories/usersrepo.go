package repositories

import (
	"context"
	"time"

	"smart-notes-api/internal/models"
	"smart-notes-api/internal/storage"
)

type UsersRepository struct{}

func (r *UsersRepository) UploadUser(ctx context.Context, user *models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	insertQuery := "INSERT INTO users(user_email, first_name, last_name, password) VALUES ($1, $2, $3, $4) RETURNING id"

	// execute query
	var userID int64
	err := storage.Pool.QueryRow(ctx, insertQuery,
		user.UserEmail,
		user.FirstName,
		user.LastName,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
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

func (r *UsersRepository) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var user models.User

	selectQuery := "SELECT * FROM users WHERE id = $1"

	err := storage.Pool.QueryRow(ctx, selectQuery, id).Scan(
		&user.UserEmail,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
