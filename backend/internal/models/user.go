package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Nickname     string
	Avatar       string
	Email        string
	PasswordHash string
	OathProvider string
	Id_oath      int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
