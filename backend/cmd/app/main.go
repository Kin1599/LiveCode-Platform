package main

import (
	"context"
	"fmt"
	"net/http"

	"livecode/internal/app"
	"livecode/internal/config"
	"livecode/internal/websocket"
)

func main() {
	cfg := config.MustLoad()
	storagePath := config.ConStringFromCfg(cfg.StoragePath)
	authService := app.New(storagePath)

	ctx := context.Background()

	userUUID, err := authService.RegisterNewUser(ctx, "TestUserEmail@gmail.com", "TestUserPass")
	if err != nil {
		panic(err)
	}
	fmt.Println(userUUID)

	tkn, err := authService.Login(ctx, "TestUserEmail@gmail.com", "TestUserPass")
	if err != nil {
		panic(err)
	}

	fmt.Println(tkn)

	// tkn, err = authService.Login(ctx, "TestUserEmail@gmail.com", "WrongPassWord")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(tkn)

	http.HandleFunc("/ws", websocket.WsHandler)
	go websocket.HandleMessages()

	fmt.Println("Starting server on port 8080")
	err_server := http.ListenAndServe(":8080", nil)
	if err_server != nil {
		fmt.Println("Error starting server:", err_server)
	}
}
