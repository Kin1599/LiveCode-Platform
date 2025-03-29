package handlers

import (
	"livecode/internal/models"

	"github.com/google/uuid"
)

type RegisterResponse struct {
	UserID string `json:"user_id"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UploadProjectResponse struct {
	Message string `json:"message"`
}

type DownloadProjectResponse struct {
	ProjectStructure string `json:"project_structure"`
}

type CreateSessionResponse struct {
	SessionID string `json:"session_id"`
	URL       string `json:"url"`
}

type GetSessionResponse struct {
	Session models.Session `json:"session"`
}

type CreateTemplateResponse struct {
	TemplateID uuid.UUID `json:"template_id"`
}

type GetAllTemplatesResponse struct {
	Templates []models.Template `json:"templates"`
}

type GetTemplateResponse struct {
	Template models.Template `json:"template"`
}
