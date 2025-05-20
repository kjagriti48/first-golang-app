package utils

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	//raw := os.Getenv("DATABASE_URL")
	//fmt.Printf("🚨 DATABASE_URL = %s\n", raw) // <== ADD THIS
	// Replace "postgresql" with "postgres"
	//connStr := strings.Replace(raw, "postgresql://", "postgres://", 1)

	raw := os.Getenv("DATABASE_URL")
	if raw == "" {
		panic("🚨 DATABASE_URL is EMPTY at runtime")
	}
	fmt.Printf("🚨 DATABASE_URL found: %s\n", raw)

	connStr := strings.Replace(raw, "postgresql://", "postgres://", 1)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect: %v", err))
	}

	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Database unreachable: %v", err))
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS students (
		name TEXT PRIMARY KEY,
		age INTEGER,
		marks TEXT,
		status TEXT
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		panic(fmt.Sprintf("Failed to create table: %v", err))
	}

	fmt.Println("✅ Connected to PostgreSQL and table ready.")
}
