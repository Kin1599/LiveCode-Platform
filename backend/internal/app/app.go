package app

import (
	"livecode/internal/database"
	"livecode/internal/services/auth"
)

func New(storagePath string) *auth.Auth {
	storage, err := database.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(storage, storage)

	return authService
}
