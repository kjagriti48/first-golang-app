package routes

import (
	"first-golang-app/handlers"
	"log"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/students", handlers.StudentHandler)
	http.HandleFunc("/students/", handlers.StudentHandler)
	http.HandleFunc("/top-student", handlers.TopStudentHandler)
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}
