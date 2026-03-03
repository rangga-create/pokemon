package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	// Supabase project URL is NOT a Postgres DSN; use DATABASE_URL (or SUPABASE_DB_URL)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = os.Getenv("SUPABASE_DB_URL")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY") // optional for client libs

	if dbURL == "" {
		panic("DATABASE_URL (or SUPABASE_DB_URL) is required for Postgres connection")
	}
	if supabaseURL == "" {
		fmt.Println("warning: SUPABASE_URL not set; only needed for client libraries")
	}
	if supabaseKey == "" {
		fmt.Println("warning: SUPABASE_API_KEY not set; only needed for client libraries")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %v", err))
	}

	// simple health check
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}

	DB = db
	fmt.Println("Database connected successfully (supabase)")
}
