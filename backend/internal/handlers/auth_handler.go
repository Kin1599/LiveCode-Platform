package handlers

import (
	"fmt"
	"livecode/internal/services/auth"
	"livecode/internal/services/jwt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

var authService *auth.Auth

func InitAuthService(service *auth.Auth) {
	authService = service
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Регистрация нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body struct{Email string; Password string} true "User data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/register [post]
func Register(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	userUUID, err := authService.RegisterNewUser(ctx, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userUUID})
}

// Login godoc
// @Summary Авторизация пользователя
// @Description Вход пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body struct{Email string; Password string} true "User data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/login [post]
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	accessToken, refreshToken, err := authService.Login(ctx, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")
	userModel, err := jwt.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token"})
		return
	}

	accessToken, err := jwt.NewToken(userModel, time.Hour*1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

func GetUserInfo(c *gin.Context) {
	// Извлечение токена из заголовка Authorization
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Разделение заголовка на тип и токен
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	token := parts[1]

	// Валидация токена
	userModel, err := jwt.ValidateToken(token)
	if err != nil {
		log.Printf("Failed to validate token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Получение информации о пользователе
	ctx := context.Background()
	userInfo, err := authService.GetUserInfo(ctx, userModel.Email)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get info about user"})
		return
	}

	// Формирование ответа
	var info struct {
		ID       uuid.UUID
		Nickname string
		Avatar   string
		Email    string
	}

	info.ID = userInfo.ID
	info.Nickname = userInfo.Nickname
	info.Avatar = userInfo.Avatar
	info.Email = userInfo.Email

	c.JSON(http.StatusOK, gin.H{"UserInfo": info})
}

func ChangeUserInfo(c *gin.Context) {
	newEmail := c.PostForm("email")
	newNickname := c.PostForm("nickname")
	newAvatar := c.PostForm("avatar")
	password := c.PostForm("password")

	ctx := context.Background()
	err := authService.ChangeUser(ctx, newEmail, newNickname, newAvatar, password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to change user: %w", err)})
		return
	}

	c.JSON(http.StatusOK, "Changed")
}

// Ping godoc
// @Summary Проверка работы сервера
// @Description Эндпоинт для проверки работы сервера
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /api/ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
