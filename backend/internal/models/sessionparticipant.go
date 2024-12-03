package models

import (
	"time"

	"github.com/google/uuid"
)

type SessionParticipant struct {
	ID             uuid.UUID
	IdSession      uuid.UUID
	IdUser         uuid.UUID
	Nickname       string
	ExpirationTime time.Time
	Avatar         string
	IsCreator      bool
	CanEdit        bool
	JoinedAt       time.Time
}
