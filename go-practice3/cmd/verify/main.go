package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir := "internal/db/migrations"
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read %s: %v", dir, err)
	}
	fmt.Printf("Found %d migration files in %s:\n", len(entries), dir)
	for _, e := range entries {
		fmt.Println(" -", e.Name())
	}
	fmt.Println("\nTo apply migrations run:")
	fmt.Println(`migrate -path internal/db/migrations -database "sqlite3://./expense.db" up`)
}
