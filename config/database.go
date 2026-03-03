package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = os.Getenv("SUPABASE_DB_URL")
	}

	if dbURL == "" {
		// Gunakan log.Println alih-alih panic agar server tidak langsung mati
		log.Println("WARNING: DATABASE_URL is not set")
		return 
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("ERROR: failed to open database: %v\n", err)
		return
	}

	// Ubah bagian ini: Jangan gunakan panic jika Ping gagal
	if err := db.Ping(); err != nil {
		log.Printf("WARNING: database ping failed on startup (koneksi mungkin butuh waktu): %v\n", err)
		// Kita tetap menyimpan instance 'db' ke 'DB' karena sql.Open otomatis akan melakukan 
		// reconnecting saat ada request ke database nantinya.
	} else {
		fmt.Println("Database connected successfully (supabase)")
	}

	DB = db
}