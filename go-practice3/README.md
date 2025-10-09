Go Practice 3 — Database Migrations and Verification

## 📘 Overview
This project implements **database migrations**, **dummy data seeding**, and **verification** commands using **Go** and **SQLite**.  
All migrations are managed using the [`golang-migrate`](https://github.com/golang-migrate/migrate) package, with database access handled by the pure Go driver [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite).

---

## ⚙️ Commands

### 1️⃣ Apply Migrations
Creates the database and applies all `.up.sql` migrations.

```bash
go run ./cmd/apply_migrations

2️⃣ Insert Dummy Data

Populates the users, categories, and expenses tables.

go run ./cmd/seed

3️⃣ Verify Database Tables

Checks if tables were created successfully.

go run ./cmd/verify

4️⃣ Query Data

Displays users and their expenses (JOIN across 3 tables).

go run ./cmd/query

🛠️ Technologies Used

Go 1.23

SQLite (modernc.org/sqlite)

golang-migrate v4

modular command-based structure