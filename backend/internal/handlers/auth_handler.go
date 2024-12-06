package handlers

import (
	"livecode/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

var authService *auth.Auth

func InitAuthService(service *auth.Auth) {
	authService = service
}

// Register godoc
// @summary Регистрация нового пользователя
// @Description Регистрация нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/register [post]
func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	userUUID, err := authService.RegisterNewUser(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userUUID})
}

// Register godoc
// @summary Авторизация пользователя
// @Description Вход пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/login [post]
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	token, err := authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}