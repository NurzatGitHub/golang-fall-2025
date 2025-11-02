package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Простая строка подключения
	connStr := "postgres://postgres:password@localhost:5432/products?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Ждем подключения с ретраями
	var dbErr error
	for i := 0; i < 5; i++ {
		dbErr = db.Ping()
		if dbErr == nil {
			break
		}
		log.Printf("Attempt %d: Failed to ping database: %v", i+1, dbErr)
		time.Sleep(2 * time.Second)
	}
	
	if dbErr != nil {
		log.Fatal("Failed to connect after retries:", dbErr)
	}

	log.Println("✅ Successfully connected to PostgreSQL!")

	router := gin.Default()

	// Простой health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "database": "connected"})
	})

	// Простой products endpoint
	router.GET("/products", func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT p.id, p.name, c.name as category, p.price 
			FROM products p 
			JOIN categories c ON p.category_id = c.id 
			LIMIT 10
		`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var products []map[string]interface{}
		for rows.Next() {
			var id int
			var name, category string
			var price int
			
			err := rows.Scan(&id, &name, &category, &price)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			
			products = append(products, map[string]interface{}{
				"id":       id,
				"name":     name,
				"category": category,
				"price":    price,
			})
		}

		c.JSON(http.StatusOK, products)
	})

	log.Println("🚀 Server starting on :8080")
	router.Run(":8080")
}
