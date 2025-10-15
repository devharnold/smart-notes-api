package main

import (
	"smart-notes-api/internal/handlers"

	"github.com/gin-gonic/gin"

	"log"
	"smart-notes-api/internal/repositories"
	"smart-notes-api/internal/services"
)

func main() {
	r := gin.Default()

	// dependencies
	userRepo := &repositories.UsersRepository{}
	notesRepo := &repositories.NotesMeta{}

	// services
	userService := &services.UserService{Repo: userRepo}
	noteService := &services.NotesService{Repo: notesRepo}

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	notesHandler := handlers.NewNotesHandler(noteService)

	// Routes
	api := r.Group("/api")
	{
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/", (*userHandler).Register)
			userRoutes.GET("/:email", (*userHandler).GetUserByEmail)
		}
		noteRoutes := api.Group("/notes")
		{
			noteRoutes.POST("/", (*notesHandler).Upload)
			noteRoutes.GET("/:title", (*notesHandler).Retrieve)
			noteRoutes.PUT("/:title", (*notesHandler).Update)
			noteRoutes.DELETE("/:title", (*notesHandler).Delete)
		}
	}

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
