package utils

import "github.com/google/uuid"

func GenerateSessionID() (uuid.UUID, error) {
	return uuid.NewUUID()
}
