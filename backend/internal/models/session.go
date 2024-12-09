package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID              uuid.UUID
	IdOwner        uuid.UUID
	Title           string
	Language        string
	AccessType 		string
	ExpirationTime time.Time
	MaxUsers       int64
	IsEditable     bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsActive       bool
}
