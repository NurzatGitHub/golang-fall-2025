package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type ProductQuery struct {
	Category string `form:"category"`
	MinPrice int    `form:"min_price"`
	MaxPrice int    `form:"max_price"`
	Sort     string `form:"sort"`
	Limit    int    `form:"limit" binding:"min=1"`
	Offset   int    `form:"offset" binding:"min=0"`
}

func main() {
	// Connect to PostgreSQL
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=products sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	router := gin.Default()

	// Products endpoint
	router.GET("/products", func(c *gin.Context) {
		var query ProductQuery

		// Parse query parameters
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}

		// Install default values
		if query.Limit == 0 {
			query.Limit = 10
		}
		if query.Sort == "" {
			query.Sort = "price_asc"
		}

		// Start timer
		startTime := time.Now()

		// Build SQL query
		products, err := getProducts(db, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}

		// Calculate query execution time
		queryTime := time.Since(startTime).Milliseconds()

		// Set custom header with query time
		c.Header("X-Query-Time", fmt.Sprintf("%dms", queryTime))

		c.JSON(http.StatusOK, products)
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	log.Println("Server starting on :8080")
	router.Run(":8080")
}

func getProducts(db *sql.DB, query ProductQuery) ([]Product, error) {
	// Build dynamic SQL query
	sqlQuery := `
		SELECT p.id, p.name, c.name as category, p.price 
		FROM products p 
		JOIN categories c ON p.category_id = c.id 
		WHERE 1=1
	`
	
	var args []interface{}
	argIndex := 1

	if query.Category != "" {
		sqlQuery += fmt.Sprintf(" AND c.name = $%d", argIndex)
		args = append(args, query.Category)
		argIndex++
	}

	if query.MinPrice > 0 {
		sqlQuery += fmt.Sprintf(" AND p.price >= $%d", argIndex)
		args = append(args, query.MinPrice)
		argIndex++
	}

	if query.MaxPrice > 0 {
		sqlQuery += fmt.Sprintf(" AND p.price <= $%d", argIndex)
		args = append(args, query.MaxPrice)
		argIndex++
	}

	switch query.Sort {
	case "price_desc":
		sqlQuery += " ORDER BY p.price DESC"
	case "price_asc":
		sqlQuery += " ORDER BY p.price ASC"
	default:
		sqlQuery += " ORDER BY p.price ASC"
	}

	// adding pagination
	sqlQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, query.Limit, query.Offset)

	// Execute query
	rows, err := db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}