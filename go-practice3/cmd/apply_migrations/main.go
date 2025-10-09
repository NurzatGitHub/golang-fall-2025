package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

func main() {
	m, err := migrate.New(
		"file://internal/db/migrations",
		"sqlite://practice3.db",
	)
	if err != nil {
		log.Fatalf("migrate init error: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration error: %v", err)
	}

	fmt.Println("Migrations applied successfully!")
}
