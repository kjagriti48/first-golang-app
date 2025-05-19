package handlers

import (
	"encoding/json"
	"first-golang-app/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var userStore = map[string]string{}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	if u.Username == "" || u.Password == "" {
		writeError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	//Check if it already exists
	if _, exists := userStore[u.Username]; exists {
		writeError(w, http.StatusBadRequest, "User already exists")
		return
	}

	//Hash Password
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	userStore[u.Username] = string(hashed)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Signup successful",
	})
}

var jwtKey = []byte("my-secret-key")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	hashedPassword, ok := userStore[u.Username]
	if !ok {
		writeError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		writeError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	//Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), //1 hour expiry
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Could not sign token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}
