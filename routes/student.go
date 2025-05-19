package routes

import (
	"first-golang-app/handlers"
	"first-golang-app/utils"
	"log"
	"net/http"
)

func RegisterRoutes() {
	//http.HandleFunc("/students", utils.LogRequest(utils.JSONOnly(handlers.StudentHandler)))
	http.HandleFunc("/students/", utils.LogRequest(handlers.StudentHandler))
	http.HandleFunc("/top-student", utils.LogRequest(handlers.TopStudentHandler))
	http.HandleFunc("/signup", handlers.SignUpHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/students", utils.LogRequest(utils.JWTMiddleware(handlers.StudentHandler)))
	//http.HandleFunc("/students", utils.JWTMiddleware(handlers.StudentHandler))
}

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}
