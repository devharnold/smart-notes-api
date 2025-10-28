package auth_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"smart-notes-api/internal/auth"
)

func TestAndValidateJWT(t *testing.T) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		t.Fatal("SECRET_KEY environment variable not set")
	}

	// New JWT service
	jwtService := auth.NewJWTService()

	token, err := jwtService.GenerateToken("test-user_id")
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}

	// validate token
	claims, err := jwtService.ValidateToken(token)
	if err != nil {
		t.Fatalf("Error validating token: %v", err)
	}

	// assertions
	assert.NotNil(t, claims)
	assert.Equal(t, "test-user-id", claims.UserID)
	assert.NotEmpty(t, token)
}
