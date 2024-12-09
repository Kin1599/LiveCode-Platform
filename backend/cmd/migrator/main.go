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
	var dbUser, dbName, dbHost, dbPass, migrationsPath string

	flag.StringVar(&dbUser, "db-user", "", "database user")
	flag.StringVar(&dbPass, "db-pass", "", "database password")
	flag.StringVar(&dbHost, "db-host", "", "database host")
	flag.StringVar(&dbName, "db-name", "", "database name")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
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
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}

		panic(err)
	}

	fmt.Println("migrations applied")
}

// go run cmd/migrator/main.go --db-user="yourname" --db-pass="pass" --db-name="example" --migrations-path="./migrations"
