package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	fmt.Println("SQLite database initialized.")

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
