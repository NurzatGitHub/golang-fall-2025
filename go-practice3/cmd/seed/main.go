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
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Users
	_, err = db.Exec(`INSERT INTO users (username, email) VALUES 
		('Alice', 'alice@example.com'),
		('Bob', 'bob@example.com');`)
	if err != nil {
		log.Fatalf("failed to insert users: %v", err)
	}

	// Categories
	_, err = db.Exec(`INSERT INTO categories (name) VALUES 
		('Food'),
		('Transport'),
		('Entertainment');`)
	if err != nil {
		log.Fatalf("failed to insert categories: %v", err)
	}

	// Expenses
	_, err = db.Exec(`INSERT INTO expenses (user_id, category_id, amount) VALUES
	(1, 1, 15.50),
	(2, 2, 5.00),
	(1, 3, 12.00);`)
	if err != nil {
		log.Fatalf("failed to insert expenses: %v", err)
	}

	fmt.Println(" Dummy data inserted successfully!")
}
