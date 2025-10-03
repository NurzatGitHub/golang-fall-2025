# Practice 2 — Tiny HTTP Service in Go

##  Overview
This project is a simple HTTP service built with Go `net/http`.  
It implements:
- `GET /user?id=123` → returns a user id in JSON
- `POST /user` → creates a user with JSON request body
- Middleware that checks API key (`X-API-Key: secret123`) and logs requests

## 📂 Project Structure

practice-2/
├── cmd/
│ └── api/
│ └── main.go # entrypoint
├── internal/
│ ├── handlers/
│ │ └── user.go # GET /user and POST /user
│ └── middleware/
│ └── auth.go # API key middleware
├── go.mod
├── .gitignore
└── README.md


---

##  Running the server
```bash
go run ./cmd/api

http://localhost:8080