package app

import (
	"livecode/internal/database"
	"livecode/internal/services/auth"
	"livecode/internal/services/filestorage"
	"livecode/internal/services/session"
)

func NewAuth(storagePath string) (*auth.Auth, error) {
	storage, err := database.New(storagePath)
	if err != nil {
		return nil, err
	}

	authService := auth.New(storage, storage)

	return authService, nil
}

func NewS3Storage(bucketName string) (*filestorage.S3Client, error) {
	s3Client, err := filestorage.New(bucketName)
	if err != nil {
		return nil, err
	}

	return s3Client, nil
}

func NewSessionService(storagePath string) (*session.SessionService, error) {
	storage, err := database.New(storagePath)
	if err != nil {
		return nil, err
	}

	sessionService := session.New(storage, storage, storage)

	return sessionService, nil
}
