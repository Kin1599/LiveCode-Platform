package main

import (
	"fmt"

	"livecode/internal/app"
	"livecode/internal/config"
	"livecode/internal/handlers"
	"livecode/internal/routes"
	"livecode/internal/websocket"
)

import (
	
)

// @title LiveCode API
// @version 1.0
// @description LiveCode API
// @host localhost:8080
// @BasePath /api

type Project struct {
	Root map[string]interface{} `json:"root"`
}

func main() {
	cfg := config.MustLoad()
	storagePath := config.ConStringFromCfg(cfg.StoragePath)
	authService := app.New(storagePath)

	handlers.InitAuthService(authService)

	router := routes.SetupRouter()

	go websocket.HandleMessages()

	fmt.Println("Starting server on port 8080")
	err_server := router.Run(":8080")
	if err_server != nil {
		fmt.Println("Error starting server:", err_server)
	}
}
