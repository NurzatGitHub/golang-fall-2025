package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "practice3.db")
	if err != nil {
		log.Fatalf(" failed to open database: %v", err)
	}
	defer db.Close()

	fmt.Println(" Users:")
	rows, err := db.Query("SELECT id, username, email, created_at FROM users")
	if err != nil {
		log.Fatalf(" failed to select users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, email, createdAt string
		rows.Scan(&id, &username, &email, &createdAt)
		fmt.Printf(" - [%d] %s (%s) — created at %s\n", id, username, email, createdAt)
	}

	fmt.Println("\n Expenses:")
	expRows, err := db.Query(`
		SELECT u.username, c.name, e.amount, e.created_at
		FROM expenses e
		JOIN users u ON e.user_id = u.id
		JOIN categories c ON e.category_id = c.id
	`)
	if err != nil {
		log.Fatalf(" failed to select expenses: %v", err)
	}
	defer expRows.Close()

	for expRows.Next() {
		var username, category string
		var amount float64
		var createdAt string
		expRows.Scan(&username, &category, &amount, &createdAt)
		fmt.Printf(" - %s spent %.2f on %s at %s\n", username, amount, category, createdAt)
	}
}
