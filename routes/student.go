package routes

import (
    "net/http"
    "first-golang-app/handlers"
)

func RegisterRoutes() {
    http.HandleFunc("/students", handlers.StudentHandler)
    http.HandleFunc("/students/", handlers.StudentHandler)
    http.HandleFunc("/top-student", handlers.TopStudentHandler)
}