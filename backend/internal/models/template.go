package models

import (
	"time"

	"github.com/google/uuid"
)

type Template struct {
	ID           int64
	Name         string
	Language     string
	Nickname     string
	TemplateCode string
	CreatedBy    uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
