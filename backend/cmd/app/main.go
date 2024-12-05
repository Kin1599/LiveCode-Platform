package main

import (
	"fmt"

	"livecode/internal/app"
	"livecode/internal/config"
	"livecode/internal/handlers"
	"livecode/internal/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title LiveCode API
// @version 1.0
// @description LiveCode API
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.MustLoad()
	storagePath := config.ConStringFromCfg(cfg.StoragePath)
	authService := app.New(storagePath)

	handlers.InitAuthService(authService)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	router.Use(cors.New(config))

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.POST("/api/session", handlers.CreateSession)
	router.GET("/api/session", handlers.GetSession)
	router.GET("/ws", func(ctx *gin.Context) {
		websocket.WsHandler(ctx.Writer, ctx.Request)
	})
	go websocket.HandleMessages()

	fmt.Println("Starting server on port 8080")
	err_server := router.Run(":8080")
	if err_server != nil {
		fmt.Println("Error starting server:", err_server)
	}
}
