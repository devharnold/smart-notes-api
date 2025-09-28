package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"smart-notes-api/internal/auth"
	"smart-notes-api/internal/models"
	"smart-notes-api/internal/repositories"
)

type UserService struct {
	repo *repositories.UsersRepository
}

func NewUserService(repo *repositories.UsersRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if user.UserEmail == "" {
		return nil, errors.New("An email address is required")
	}
	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: Correct the return statement
	return s.repo.UploadUser(ctx)
}
