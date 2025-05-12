package main

import (
	"fmt"
	"log"
	"net/http"

	"first-golang-app/routes"
	"first-golang-app/utils"
)

func main() {
	utils.InitDB("students.db")
	utils.PrintAllStudents()
	//utils.LoadStudentsFromFile()
	routes.RegisterRoutes()

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
