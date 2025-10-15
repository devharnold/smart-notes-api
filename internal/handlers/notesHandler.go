package handlers

import (
	"net/http"

	"smart-notes-api/internal/services"

	"github.com/gin-gonic/gin"
)

type NotesHandler struct {
	Service *services.NotesService
}

func NewNotesHandler(s *services.NotesService) *NotesHandler {
	return &NotesHandler{Service: s}
}

type UploadRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type RetrieveRequest struct {
	Title string `json:"title"`
}

type UpdateRequest struct {
	Title      string `json:"title"`
	NewContent string `json:"new_content" binding:"required"`
}

type DeleteRequest struct {
	Title string `json:"title"`
}

// helper to get UserID --> For security --> authorization
func getUserID(c *gin.Context) int {
	return 1
}

// upload handler
func (h *NotesHandler) Upload(c *gin.Context) {
	var req UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	s3Path, err := h.Service.SaveNote(c.Request.Context(), userID, req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Note uploaded successfully",
		"s3_path": s3Path,
	})
}

// retrieve note handler
func (h *NotesHandler) Retrieve(c *gin.Context) {
	var req RetrieveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	content, err := h.Service.RetrieveNote(c.Request.Context(), userID, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   req.Title,
		"content": content,
	})
}

// update handler
func (h *NotesHandler) Update(c *gin.Context) {
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	if err := h.Service.UpdateNote(c.Request.Context(), userID, req.Title, req.NewContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note updated successfully",
	})
}

// delete handler
func (h *NotesHandler) Delete(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	if err := h.Service.DeleteNote(c.Request.Context(), userID, req.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}
