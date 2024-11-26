package models

import (
	"time"

	"github.com/google/uuid"
)

type CodeSnippet struct {
	ID          uuid.UUID
	IdSession   uuid.UUID
	Language    string
	Code        string
	CreatedAt   time.Time
	UpdatedAtAt time.Time
}
