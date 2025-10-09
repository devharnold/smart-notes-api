package main

import (
	"context"
	"errors"
	"fmt"

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

	userID, err := s.repo.UploadUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload user: %w", err)
	}

	user.UserID = fmt.Sprintf("%s", userID)
	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("User not found", err)
	}

	if !auth.VerifyPassword(user.Password, password) {
		return "", fmt.Errorf("Invalid password")
	}

	// token, err :=
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user by email: %w", err)
	}

	return user, nil
}
