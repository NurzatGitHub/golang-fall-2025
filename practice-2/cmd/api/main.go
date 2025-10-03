package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NurzatGitHub/golang-fall-2025/practice-2/internal/handlers"
	"github.com/NurzatGitHub/golang-fall-2025/practice-2/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/user", middleware.WithAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetUserHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			handlers.CreateUserHandler(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})))

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
