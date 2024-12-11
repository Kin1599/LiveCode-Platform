// package main

// import (
// 	"fmt"

// 	"livecode/internal/app"
// 	"livecode/internal/config"
// 	"livecode/internal/handlers"
// 	"livecode/internal/routes"
// 	"livecode/internal/websocket"
// 	"livecode/internal/websocket/chat"
// )

// // @title LiveCode API
// // @version 1.0
// // @description LiveCode API
// // @host localhost:8080
// // @BasePath /api

// func main() {
// 	cfg := config.MustLoad()
// 	storagePath := config.ConStringFromCfg(cfg.StoragePath)
// 	authService := app.New(storagePath)

// 	handlers.InitAuthService(authService)

// 	s3Serve := app.NewS3Storage(cfg.StoragePath.BucketName)
// 	handlers.InitS3Service(s3Serve)

// 	sessionClient := app.NewSessionService(storagePath)
// 	handlers.InitSessionService(sessionClient)

// 	router := routes.SetupRouter()

// 	go websocket.HandleMessages()
// 	go chat.Run()

// 	fmt.Println("Starting server on port 80")
// 	err_server := router.Run(":8080")
// 	if err_server != nil {
// 		fmt.Println("Error starting server:", err_server)
// 	}
// }

// export CONFIG_PATH="./configs/local.yaml"

package main

import (
	"context"
	"fmt"

	"livecode/internal/app"
	"livecode/internal/config"

	"github.com/google/uuid"
)

// @title LiveCode API
// @version 1.0
// @description LiveCode API
// @host localhost:8080
// @BasePath /api

func main() {
	cfg := config.MustLoad()
	storagePath := config.ConStringFromCfg(cfg.StoragePath)
	sessionClient := app.NewSessionService(storagePath)

	for i := 0; i < 3; i++ {
		res, err := sessionClient.CreateNewSession(context.Background(), uuid.New(),
			uuid.New(), "123", "123", "Public", 123, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
	for {

	}
}
