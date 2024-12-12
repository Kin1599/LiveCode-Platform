package handlers

import (
	"fmt"
	"livecode/internal/services/auth"
	"livecode/internal/services/jwt"
	"net/http"
	"strings"

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
	token, err := authService.Login(ctx, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetUserInfo(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	userModel, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	ctx := context.Background()
	userInfo, err := authService.GetUserInfo(ctx, userModel.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get info about user"})
		return
	}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("Failed to change user: %w", err)})
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
