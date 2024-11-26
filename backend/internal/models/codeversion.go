package models

import (
	"time"

	"github.com/google/uuid"
)

type CodeVersion struct {
	ID        uuid.UUID
	IdSnippet uuid.UUID
	Code      string
	CreatedAt time.Time
	IdAuthor  uuid.UUID
}
