package routes

import (
	"first-golang-app/handlers"
	"first-golang-app/utils"
	"log"
	"net/http"
)

func RegisterRoutes() {
	// Public endpoints
	http.HandleFunc("/signup", utils.EnableCORS(handlers.SignUpHandler))
	http.HandleFunc("/login", utils.EnableCORS(handlers.LoginHandler))

	// Protected endpoints
	http.HandleFunc("/students", utils.EnableCORS(utils.JWTMiddleware(handlers.StudentHandler)))
	http.HandleFunc("/students/", utils.EnableCORS(utils.JWTMiddleware(handlers.StudentHandler)))
	http.HandleFunc("/top-student", utils.EnableCORS(utils.JWTMiddleware(handlers.TopStudentHandler)))
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}
