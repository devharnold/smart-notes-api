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

type MockNotesService struct {
	mock.Mock
}

func (m *MockNotesService) SaveNote(ctx context.Context, userID int, title, content string) (string, error) {
	args := m.Called(ctx, userID, title, content)
	return args.String(0), args.Error(1)
}

func (m *MockNotesService) RetrieveNote(ctx context.Context, title string) (string, error) {
	args := m.Called(ctx, title)
	return args.String(0), args.Error(1)
}

func (m *MockNotesService) UpdateNote(ctx context.Context, title, content string) (string, error) {
	args := m.Called(ctx, title, content)
	return args.String(0), args.Error(1)
}

func (m *MockNotesService) DeleteNote(ctx context.Context, userID, noteID int) (string, error) {
	args := m.Called(ctx, userID, noteID)
	return args.String(0), args.Error(1)
}

func SetUpRouter(handler *handlers.NotesHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/upload", handler.Upload)
	r.GET("/retrieve", handler.Retrieve)
	r.PUT("/update", handler.Update)
	r.DELETE("/delete", handler.Delete)
	return r
}

func TestUploadNoteService(t *testing.T) {
	mockService := new(MockNotesService)
	handler := handlers.NewNotesHandler((*services.NotesService)(mockService))
	router := SetUpRouter(handler)

	mockService.On("SaveNote", mock.Anything, 1, "Test Title", "Test Content").Return("s3://bucket/test.txt", nil)

	body, _ := json.Marshal(handlers.UploadRequest{
		Title:   "Test Title",
		Content: "Test Content",
	})

	req, _ := http.NewRequest(http.MethodPost, "/upload", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Note uploaded successfully")
	mockService.AssertExpectations(t)
}

func TestUploadNote_BadRequest(t *testing.T) {
	mockService := new(MockNotesService)
	handler := handlers.NewNotesHandler((*services.NotesService)(mockService))
	router := SetUpRouter(handler)

	req, _ := http.NewRequest(http.MethodPost, "/upload", bytes.NewBuffer([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRetrieveNote_Success(t *testing.T) {
	mockService := new(MockNotesService)
	handler := handlers.NewNotesHandler((*services.NotesService)(mockService))
	router := SetUpRouter(handler)

	mockService.On("RetrieveNote", mock.Anything, 1, "Test Title").Return("Found Test Note Content", nil)

	body, _ := json.Marshal(handlers.RetrieveRequest{Title: "Test Title"})
	req, _ := http.NewRequest(http.MethodPost, "/retrieve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Found Test Note Content")
	mockService.AssertExpectations(t)
}

func TestUpdateNote_Error(t *testing.T) {
	mockService := new(MockNotesService)
	handler := handlers.NewNotesHandler((*services.NotesService)(mockService))
	router := SetUpRouter(handler)

	mockService.On("UpdateNote", mock.Anything, 1, "Test Title", "New Content").Return(errors.New("failed to update"))

	body, _ := json.Marshal(handlers.UpdateRequest{Title: "Test Title", NewContent: "New Content"})
	req, _ := http.NewRequest(http.MethodPut, "/update", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to update")
	mockService.AssertExpectations(t)
}

func TestDeleteNote_Success(t *testing.T) {
	mockService := new(MockNotesService)
	handler := handlers.NewNotesHandler((*services.NotesService)(mockService))
	router := SetUpRouter(handler)

	mockService.On("DeleteNote", mock.Anything, 1, "Old Note").Return(nil)

	body, _ := json.Marshal(DeleteRequest{Title: "Old Note"})
	req, _ := http.NewRequest(http.MethodDelete, "/delete", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Note deleted successfully")
	mockService.AssertExpectations(t)
}
