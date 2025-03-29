package handlers

import (
	"context"
	"errors"
	"fmt"
	"livecode/internal/models"
	"livecode/internal/services/session"
	"livecode/internal/utils"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessions = make(map[uuid.UUID]models.Session)
var mutex = &sync.Mutex{}
var sessionService *session.SessionService

func InitSessionService(service *session.SessionService) {
	sessionService = service
}

// CreateSession godoc
// @Summary Создание новой сессии
// @Description Создание новой сессии
// @Tags session
// @Accept json
// @Produce json
// @Param owner_id formData string true "ID пользователя"
// @Param editable formData boolean true "Редактируемая ли сессия"
// @Param title formData string true "Название сессии"
// @Param language formData string true "Язык программирования"
// @Param max_users formData integer true "Максимальное количество пользователей"
// @Success 200 {object} CreateSessionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/session [post]
func CreateSession(c *gin.Context) {
	ownerId := c.PostForm("owner_id")
	editable := c.PostForm("editable") == "true"
	title := c.PostForm("title")
	language := c.PostForm("language")
	maxUsersStr := c.PostForm("max_users")

	maxUsers, err := strconv.ParseInt(maxUsersStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid max_users value"})
		return
	}

	ownerUUID, err := uuid.Parse(ownerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid owner_id value"})
		return
	}

	sessionID, err := utils.GenerateSessionID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to generate session ID"})
		return
	}

	session := models.Session{
		ID:             sessionID,
		IdOwner:        ownerUUID,
		Title:          title,
		Language:       language,
		ExpirationTime: time.Now().Add(time.Hour * 24),
		MaxUsers:       maxUsers,
		IsEditable:     editable,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsActive:       true,
	}

	_, err = sessionService.CreateNewSession(context.Background(), sessionID, ownerUUID, title, language, "Public", maxUsers, editable)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	mutex.Lock()
	sessions[sessionID] = session
	mutex.Unlock()

	c.JSON(http.StatusOK, CreateSessionResponse{
		SessionID: sessionID.String(),
		URL:       "/code-input/" + sessionID.String(),
	})
}

// GetSession godoc
// @Summary Получение сессии по ID
// @Description Получение сессии по ID
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Success 200 {object} GetSessionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/session [get]
func GetSession(c *gin.Context) {
	sessionIDStr := c.Query("session_id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid session_id value"})
		return
	}

	// mutex.Lock()
	// // _, _ = sessions[sessionID]
	// mutex.Unlock()

	sessionModel, err := sessionService.GetSession(context.Background(), sessionID)
	if err != nil {
		if errors.Is(err, session.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "Session not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetSessionResponse{Session: sessionModel})
}

// DeleteSession godoc
// @Summary Удаление сессии по ID
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Success 200 {object} string
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/session [delete]
func DeleteSession(c *gin.Context) {
	sessionIDStr := c.Query("session_id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid session_id value"})
		return
	}

	// mutex.Lock()
	// // _, _ = sessions[sessionID]
	// mutex.Unlock()

	err = sessionService.DeleteSession(context.Background(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "DELETED")
}

// GetTemplate godoc
// @Summary Получение шаблона по ID
// @Tags template
// @Accept json
// @Produce json
// @Param id query string true "ID шаблона"
// @Success 200 {object} GetTemplateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/template [get]
func GetTemplate(c *gin.Context) {
	templateIDstr := c.Query("id")
	templateID, err := uuid.Parse(templateIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid session_id value"})
		return
	}

	template, err := sessionService.TemplateByID(context.Background(), templateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetTemplateResponse{Template: template})
}

// GetAllTemplates godoc
// @Summary Получение всех шаблонов
// @Tags template
// @Accept json
// @Produce json
// @Success 200 {object} GetAllTemplatesResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/templates [get]
func GetAllTemplates(c *gin.Context) {
	templates, err := sessionService.AllTemplates(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetAllTemplatesResponse{Templates: templates})
}

// CreateTemplate godoc
// @Summary Создание нового шаблона
// @Tags template
// @Accept json
// @Produce json
// @Param template_name formData string true "Название шаблона"
// @Param language formData string true "Язык программирования"
// @Param template_code formData string true "Код шаблона"
// @Param creator_id formData string true "ID создателя"
// @Success 200 {object} CreateTemplateResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/template [post]
func CreateTemplate(c *gin.Context) {
	templateName := c.PostForm("template_name")
	language := c.PostForm("language")
	templateCode := c.PostForm("template_code")
	creatorUUIDStr := c.PostForm("creator_id")

	creatorUUID, err := uuid.Parse(creatorUUIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid owner_id value"})
		return
	}

	templatesUUID, err := sessionService.CreateTemplate(context.Background(), templateName,
		language, templateCode, creatorUUID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateTemplateResponse{
		TemplateID: templatesUUID,
	})
}
