package main

import (
	"net/http"
	"github.com/gin-gonic/gin"


	"context"
	"log"
	"smart-notes-api/internal/db"
	"smart-notes-api/internal/repositories"
	"smart-notes-api/internal/services"
	"smart-notes-api/internal/storage"
)

func main() {
	r := gin.Default()
	db.Init()
	defer db.Close()

	fileRepo	:= &repositories.FileRepo{}
	s3Client	:= storage.NewS3Client()
	FileService	:= &services.FileService{Repo: fileRepo, S3: s3Client}


	err := fileService.UploadUserFile(context.Background(), 1, "", []byte())
	if err != nil {

	}
}

func UploadFile


// TODO: Clean up the main.go file, also handle the http routes