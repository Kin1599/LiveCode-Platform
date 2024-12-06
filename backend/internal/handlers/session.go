package handlers

import (
	"livecode/internal/models"
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

// CreateSession godoc
// @summary Создание новой сессии
// @Description Создание новой сессии
// @Tags session
// @Accept json
// @Produce json
// @Param owner_id query string true "ID пользователя"
// @Param editable query boolean true "Редактируемая ли сессия"
// @Param title query string true "Название сессии"
// @Param language query string true "Язык программирования"
// @Param max_users query integer true "Максимальное количество пользователей"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/session [post]
func CreateSession(c *gin.Context) {
	ownerId := c.PostForm("owner_id")
	editable := c.PostForm("editable") == "true"
	title := c.PostForm("title")
	language := c.PostForm("language")
	maxUsersStr := c.PostForm("max_users")

	maxUsers, err := strconv.ParseInt(maxUsersStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid max_users value"})
		return
	}

	ownerUUID, err := uuid.Parse(ownerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner_id value"})
		return
	}

	sessionID, err := utils.GenerateSessionID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session ID"})
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

	mutex.Lock()
	sessions[sessionID] = session
	mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"session_id": sessionID.String(),
		"url":        "/session/" + sessionID.String(),
	})
}

// GetSession godoc
// @summary Получение сессии по ID
// @Description Получение сессии по ID
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/session [get]
func GetSession(c *gin.Context) {
	sessionIDStr := c.Query("session_id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id value"})
		return
	}

	mutex.Lock()
	session, exists := sessions[sessionID]
	mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	c.JSON(http.StatusOK, session)
}
