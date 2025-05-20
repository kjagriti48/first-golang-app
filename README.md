# 🚀 Student Management API (Go + PostgreSQL + JWT)

A production-ready backend API built with Go, PostgreSQL, and JWT authentication. Deployed on Railway.

---

## 🌐 Live API

Base URL:  
`https://first-golang-app-production.up.railway.app`

---

## 🧠 Features

- JWT authentication (signup/login)
- Middleware-based route protection
- PostgreSQL persistence
- RESTful structure
- Modular project layout
- Cloud-deployed via Railway

---

## 📦 Endpoints

### 🔐 `POST /signup`

Registers a new user.

**Request body:**

```json
{
  "username": "admin",
  "password": "secure123"
}

🔐 POST /login
Logs in and returns a JWT token.

Request body:

{
  "username": "admin",
  "password": "secure123"
}

Response:

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}

🔐 GET /students (Protected)
Returns the list of students.
Requires Authorization header:

Authorization: Bearer <your_token>

🔐 POST /students (Protected)
Adds a student.

{
  "name": "Arjun",
  "age": 20,
  "marks": {
    "math": 90,
    "science": 85
  }
}

🔐 DELETE /students/{name} (Protected)
Deletes a student by name.


🔐 GET /top-student (Protected)
Returns the student with the highest average score.

🧪 How to Test with Postman
POST /signup → create a user

POST /login → receive token

Use the token in the Authorization header for all other requests
Authorization: Bearer <token>

🛠️ Tech Stack
Go 1.21
PostgreSQL (via Railway)
golang-jwt for auth
bcrypt for hashing
Hosted on Railway


👤 Author
Created by Jagriti.