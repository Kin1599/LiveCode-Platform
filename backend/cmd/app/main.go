package main

import (
	"fmt"
	"log"

	"livecode/internal/app"
	"livecode/internal/config"
	"livecode/internal/handlers"
	"livecode/internal/routes"
	"livecode/internal/services/filestorage"
	"livecode/internal/websocket"
	"livecode/internal/websocket/chat"
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

	s3Client, err := filestorage.New(cfg.StoragePath.BucketName)
	if err != nil {
		log.Fatalf("Error creating S3 client: %v", err)
	}

	handlers.InitS3Client(s3Client)

	sessionClient := app.NewSessionService(storagePath)
	handlers.InitSessionService(sessionClient)

	websocket.Init(sessionClient)
	router := routes.SetupRouter()

	go websocket.HandleMessages()
	go chat.Run()

	fmt.Println("Starting server on port 80")
	err_server := router.Run(":8080")
	if err_server != nil {
		fmt.Println("Error starting server:", err_server)
	}
}
