package app

import (
	"livecode/internal/database"
	"livecode/internal/services/auth"
	"livecode/internal/services/filestorage"
	"livecode/internal/services/session"
)

func New(storagePath string) *auth.Auth {
	storage, err := database.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(storage, storage)

	return authService
}

func NewS3Storage(bucketName string) *filestorage.S3Service {
	s3Client, err := database.NewS3Client(bucketName)
	if err != nil {
		panic(err)
	}

	s3Serve := filestorage.New(s3Client, s3Client)

	return s3Serve
}

func NewSessionService(storagePath string) *session.SessionService {
	storage, err := database.New(storagePath)
	if err != nil {
		panic(err)
	}

	sessionService := session.New(storage, storage, storage)

	return sessionService
}
