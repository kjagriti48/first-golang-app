package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connStr := "postgres://localhost/studentdb?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to PostgreSQL: %v", err))
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

	fmt.Println("Connected to PostgreSQL and ready ")

}

func PrintAllStudents() {
	rows, err := DB.Query("SELECT name, age, status FROM students")
	if err != nil {
		fmt.Println("Failed to read from DB:", err)
		return
	}

	defer rows.Close()

	fmt.Println("Students in Database:")
	for rows.Next() {
		var name string
		var age int
		var status string

		rows.Scan(&name, &age, &status)
		fmt.Printf("- %s (Age: %d, Status: %s)\n", name, age, status)

	}
}
