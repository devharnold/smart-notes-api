package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"smart-notes-api/internal/handlers"
	"smart-notes-api/internal/services"
)

type MockUsersService struct {
	mock.Mock
}

func (m *MockUsersService) Create(ctx context.Context, firstName, lastName, email string) (string, error) {
	args := m.Called(ctx, firstName, lastName, email)
	return args.String(0), args.Error(1)
}

func (m *MockUsersService) Login(ctx context.Context, email, password string) (string, error) {
	args := m.Called(ctx, email, password)
	return args.String(0), args.Error(1)
}

func setUpRouter(handler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/create", handler.Register)
	r.POST("/login", handler.Login)
	return r
}

func TestRegisterUserService(t *testing.T) {
	mockService := new(MockUsersService)
	handler := handlers.NewUserHandler((*services.NewUserService)(mockService))
	router := setUpRouter(handler)

	mockService.On("RegisterUser", mock.Anything, 1, "Test firstName", "Test lastName", "Test Email").Return("UserId", nil)

	body, _ := json.Marshal(handlers.SignUpRequest{
		FirstName: "TestFirstName",
		LastName:  "TestLastName",
		Email:     "TestEmail",
	})

	req, _ := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "UserId")
	mockService.AssertNumberOfCalls(t, "RegisterUser", 1)
	mockService.AssertExpectations(t)
}

func TestLoginUserService(t *testing.T) {
	mockService := new(MockUsersService)
	handler := handlers.NewUserHandler((*services.UserService)(mockService))
	router := setUpRouter(handler)

	mockService.On("LoginUser", mock.Anything, 1, "Test Email", "Test Password").Return("UserId", nil)

	body, _ := json.Marshal(handlers.LoginRequest{
		Email:    "TestEmail",
		Password: "TestPassword",
	})

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "UserId")
	mockService.AssertNumberOfCalls(t, "LoginUser", 1)
	mockService.AssertExpectations(t)
}
