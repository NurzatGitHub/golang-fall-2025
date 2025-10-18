package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// User represents a user row
type User struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Email   string  `db:"email"`
	Balance float64 `db:"balance"`
}

// InsertUser inserts new user using NamedExec
func InsertUser(db *sqlx.DB, user User) error {
	query := `INSERT INTO users (name, email, balance) VALUES (:name, :email, :balance)`
	_, err := db.NamedExec(query, user)
	return err
}

// GetAllUsers returns all users using Select
func GetAllUsers(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT id, name, email, balance FROM users ORDER BY id")
	return users, err
}

// GetUserByID returns a single user by id using Get
func GetUserByID(db *sqlx.DB, id int) (User, error) {
	var u User
	err := db.Get(&u, "SELECT id, name, email, balance FROM users WHERE id=$1", id)
	return u, err
}

// TransferBalance transfers amount from fromID to toID within a transaction.
// It performs SELECT ... FOR UPDATE (row lock) to avoid races.
// If any check fails, it rollbacks.
func TransferBalance(db *sqlx.DB, fromID int, toID int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	// ensure rollback on any failure
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	// Lock sender row
	var from User
	if err := tx.Get(&from, "SELECT id, balance FROM users WHERE id=$1 FOR UPDATE", fromID); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("sender not found: %w", err)
	}

	// Lock receiver row
	var to User
	if err := tx.Get(&to, "SELECT id, balance FROM users WHERE id=$1 FOR UPDATE", toID); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("receiver not found: %w", err)
	}

	// Check balance
	if from.Balance < amount {
		_ = tx.Rollback()
		return fmt.Errorf("insufficient funds: have %.2f need %.2f", from.Balance, amount)
	}

	// Update balances
	if _, err := tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, fromID); err != nil {
		_ = tx.Rollback()
		return err
	}
	if _, err := tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, toID); err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
