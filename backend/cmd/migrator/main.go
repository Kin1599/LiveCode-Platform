package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var dbUser, dbName, dbHost, dbPass, migrationsPath, direction string

	flag.StringVar(&dbUser, "db-user", "", "database user")
	flag.StringVar(&dbPass, "db-pass", "", "database password")
	flag.StringVar(&dbHost, "db-host", "", "database host")
	flag.StringVar(&dbName, "db-name", "", "database name")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&direction, "direction", "up", "migration direction: up or down")
	flag.Parse()

	if dbUser == "" {
		panic("db-user is required")
	}
	if dbName == "" {
		panic("db-name is required")
	}

	if dbHost == "" {
		panic("db-host is required")
	}

	if dbPass == "" {
		panic("db-pass is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName),
	)
	if err != nil {
		panic(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")
				return
			}

			panic(err)
		}

		fmt.Println("migrations applied")
	case "down":
		if err := m.Force(1); err != nil {
			panic(err)
		}

		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to revert")
				return
			}
		}

		fmt.Println("migrations reverted")
	default:
		panic("invalid direction. Use 'up' or 'down'.")
	}

}

// go run cmd/migrator/main.go --db-user="yourname" --db-pass="pass" --db-name="example" --db-host=localhost --migrations-path="./migrations"
