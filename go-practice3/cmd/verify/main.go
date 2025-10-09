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

	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	if err != nil {
		log.Fatalf(" failed to query tables: %v", err)
	}
	defer rows.Close()

	fmt.Println(" Connected successfully!")
	fmt.Println("Tables in database:")
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		fmt.Println(" -", tableName)
	}
}
