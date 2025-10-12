package handlers

import (
	"net/http"

	"smart-notes-api/internal/models"
	"smart-notes-api/internal/services"

	"github.com/gin-gonic/gin"
)

type NotesHandler struct {
	notesService services.NotesService
}

func NewNotesHandler(notesService services.NotesService) *NotesHandler {
	return &NotesHandler{notesService: notesService}
}

type UploadRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type RetrieveRequest struct {
	NoteID string `json:"noteId"`
	Title  string `json:"title"`
	s3Key  *string
}

type UpdateRequest struct {
	NoteID string `json:"noteId"`
	Title  string `json:"title"`
	s3Key  *string
}

type DeleteRequest struct {
	NoteID string `json:"noteId"`
	Title  string `json:"title"`
	s3Key  *string
}

func (h *NotesHandler) Upload(c *gin.Context) {
	var req UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notes := &models.Notes{
		Title: req.Title,
		Body:  req.Body,
	}

	uploadedNote, err := h.notesService.SaveNote(c.Request.Context(), notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Note uploaded successfully",
		"user": gin.H{
			"title": uploadedNote.Title,
		},
	})
}

func (h *NotesHandler) Retrieve(c *gin.Context) {
	var req RetrieveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	retrievedNote, err := h.notesService.RetrieveNote(c.Request.Context(), req.NoteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note": gin.H{
			"noteID": retrievedNote.NoteID,
			"title":  retrievedNote.Title,
			"body":   retrievedNote.Body,
		},
	})
}
