package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/NurzatGitHub/practice4-sqlx/internal/store"
)

func main() {
	// DSN: postgres://user:pass@host:port/dbname?sslmode=disable
	dsn := "postgres://pguser:pgpass@localhost:5432/practisedb?sslmode=disable"

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	// Connection pooling
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("connected to db")

	// Пример использования методов

	// 1) Insert a new user
	u := store.User{Name: "Charlie", Email: "charlie@example.com", Balance: 20.0}
	if err := store.InsertUser(db, u); err != nil {
		log.Printf("InsertUser error (maybe duplicate): %v", err)
	} else {
		fmt.Println("Inserted user Charlie")
	}

	// 2) Get all users
	users, err := store.GetAllUsers(db)
	if err != nil {
		log.Fatalf("GetAllUsers error: %v", err)
	}
	fmt.Println("All users:")
	for _, uu := range users {
		fmt.Printf(" - [%d] %s (%s) balance=%.2f\n", uu.ID, uu.Name, uu.Email, uu.Balance)
	}

	// 3) Transfer money example: from id 1 to id 2 amount 10.0
	fmt.Println("Transferring 10.0 from user 1 to user 2...")
	if err := store.TransferBalance(db, 1, 2, 10.0); err != nil {
		log.Printf("TransferBalance failed: %v", err)
	} else {
		fmt.Println("Transfer succeeded")
	}

	// Check balances after transfer
	u1, _ := store.GetUserByID(db, 1)
	u2, _ := store.GetUserByID(db, 2)
	fmt.Printf("User1 balance=%.2f, User2 balance=%.2f\n", u1.Balance, u2.Balance)
}
