package routes

import (
	"livecode/internal/handlers"
	"livecode/internal/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	authRoutes := router.Group("/api")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
		authRoutes.GET("/user", handlers.GetUserInfo)
		authRoutes.POST("/refresh-token", handlers.RefreshToken)
	}

	sessionRoutes := router.Group("/api/session")
	{
		sessionRoutes.POST("", handlers.CreateSession)
		sessionRoutes.GET("", handlers.GetSession)
		sessionRoutes.DELETE("", handlers.DeleteSession)
	}

	projectRoutes := router.Group("/api")
	{
		projectRoutes.POST("/api/uploadProject", handlers.UploadProject)
		projectRoutes.GET("/api/downloadProject", handlers.DownloadProject)
	}

	wsRoutes := router.Group("/")
	{
		wsRoutes.GET("/ws", func(ctx *gin.Context) {
			websocket.WsHandler(ctx.Writer, ctx.Request)
		})

		wsRoutes.GET("/chat", func(ctx *gin.Context) {
			websocket.ServeWs(ctx.Writer, ctx.Request)
		})
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
