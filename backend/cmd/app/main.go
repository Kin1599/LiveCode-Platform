package main

import (
	"context"
	"fmt"
	"livecode/internal/config"
	"livecode/internal/database"
)

func main() {

	cfg := config.MustLoad()
	db, err := database.New(cfg.StoragePath)
	if err != nil {
		fmt.Print("1", err)
	}

	defer db.Stop()

	ctx := context.Background()

	usr, err := db.User(ctx, "gemail.com")
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
	}
	fmt.Println(usr)
}
