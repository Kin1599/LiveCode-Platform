package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID          uuid.UUID
	IdSession   uuid.UUID
	IdAuthor    uuid.UUID
	Content     string
	LineStart   int64
	LineEnd     int64
	CreatedAt   time.Time
	UpdatedAtAt time.Time
}
