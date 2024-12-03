package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          int64
	IdSession   uuid.UUID
	IdUser      uuid.UUID
	MessageText string
	CreatedAt   time.Time
	UpdatedAtAt time.Time
}
