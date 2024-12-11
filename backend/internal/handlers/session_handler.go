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
		fmt.Println(err)
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

	_, err = sessionService.CreateNewSession(context.Background(), sessionID, ownerUUID, title, language, "Public", maxUsers, editable)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
// @Summary Получение сессии по ID
// @Description Получение сессии по ID
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Success 200 {object} models.Session
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/session [get]
func GetSession(c *gin.Context) {
	sessionIDStr := c.Query("session_id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id value"})
		return
	}

	// mutex.Lock()
	// // _, _ = sessions[sessionID]
	// mutex.Unlock()

	sessionModel, err := sessionService.GetSession(context.Background(), sessionID)
	if err != nil {
		if errors.Is(err, session.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessionModel)
}

func DeleteSession(c *gin.Context) {
	sessionIDStr := c.Query("session_id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id value"})
		return
	}

	// mutex.Lock()
	// // _, _ = sessions[sessionID]
	// mutex.Unlock()

	err = sessionService.DeleteSession(context.Background(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "DELETED")
}

// BlockIP godoc
// @Summary Блокировка IP
// @Description Блокировка IP для сессии
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Param ip query string true "IP для блокировки"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/block [post]
func BlockIP(c *gin.Context) {
	sessionIDStr := c.DefaultQuery("session_id", "")
	ip := c.DefaultQuery("ip", "")

	if sessionIDStr == "" || ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id and ip are required"})
		return
	}

	sessionUUID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
		return
	}

	_, err = sessionService.BlockUser(context.Background(), ip, sessionUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "IP blocked successfully"})
}

// UnblockIP godoc
// @Summary Разблокировка IP
// @Description Разблокировка IP для сессии
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Param ip query string true "IP для разблокировки"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/unblock [post]
func UnblockIP(c *gin.Context) {
	sessionIDStr := c.DefaultQuery("session_id", "")
	ip := c.DefaultQuery("ip", "")

	if sessionIDStr == "" || ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id and ip are required"})
		return
	}

	sessionUUID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
		return
	}

	err = sessionService.DeleteBlockByIP(context.Background(), ip, sessionUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "IP unblocked successfully"})
}

// GetBlockedIPs godoc
// @Summary Получить заблокированные IP для сессии
// @Description Получить список заблокированных IP для сессии
// @Tags session
// @Accept json
// @Produce json
// @Param session_id query string true "ID сессии"
// @Success 200 {object} []string
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/blocked [get]
func GetBlockedIPs(c *gin.Context) {
	sessionIDStr := c.DefaultQuery("session_id", "")
	if sessionIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	sessionUUID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session_id"})
		return
	}

	blockedIPs, err := sessionService.GetBlockedIPsBySession(context.Background(), sessionUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blockedIPs)
}
