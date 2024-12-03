package main

import (
	"context"
	"fmt"
	"livecode/internal/app"
	"livecode/internal/config"
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
}
