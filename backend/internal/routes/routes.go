package routes

import (
	"livecode/internal/handlers"
	"livecode/internal/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/api/ping", handlers.Ping)

	router.POST("/api/register", handlers.Register)
	router.POST("/api/login", handlers.Login)

	router.POST("/api/session", handlers.CreateSession)
	router.GET("/api/session", handlers.GetSession)
	router.GET("/ws", func(ctx *gin.Context) {
		websocket.WsHandler(ctx.Writer, ctx.Request)
	})

	return router
}
