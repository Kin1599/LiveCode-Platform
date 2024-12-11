package routes

import (
	"livecode/internal/handlers"
	"livecode/internal/websocket"
	"livecode/internal/websocket/chat"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// gin.SetMode(gin.ReleaseMode)

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
	router.DELETE("/api/session", handlers.DeleteSession)
	router.GET("/ws", func(ctx *gin.Context) {
		websocket.WsHandler(ctx.Writer, ctx.Request)
	})

	router.GET("/chat", func(ctx *gin.Context) {
		chat.ServeWs(ctx.Writer, ctx.Request)
	})

	router.POST("/api/uploadProject", handlers.UploadProject)
	router.GET("/api/downloadProject", handlers.DownloadProject)

	router.GET("/api/user", handlers.GetUserInfo)

	router.POST("/api/block", handlers.BlockIP)
	router.POST("/api/unblock", handlers.UnblockIP)
	router.GET("/api/blocked", handlers.GetBlockedIPs)

	return router
}
