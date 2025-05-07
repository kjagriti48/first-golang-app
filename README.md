# 🎓 Student Management API (Golang)

A modular, file-persistent REST API built in Go to manage student records, calculate average scores, determine top performers, and return clean JSON responses with proper validation and logging.

---

## 📦 Project Structure
---
first-golang-app/
├── main.go
├── go.mod
├── routes/
├── handlers/
├── models/
├── utils/
└── students.json

## 🔧 Features

- ✅ Modular folder structure: `handlers`, `routes`, `models`, `utils`
- ✅ REST endpoints for student creation, listing, deletion
- ✅ Top student identification
- ✅ File-based JSON persistence
- ✅ Input validation & structured error responses
- ✅ Logging of every request and action (Day 12)

---

## 🔗 API Endpoints

### `GET /students`
Returns all students in the system.

### `POST /students`
Adds a new student.  
Example JSON body:

```json
{
  "name": "Arjun",
  "age": 20,
  "marks": {
    "math": 90,
    "english": 85,
    "science": 95
  }
}

### `git aDELETE /students/{name}`
Deletes a student by their name.

GET /top-student
Returns the student with the highest average score.

⚠️ Input Validation
name is required

age must be greater than 0

At least one subject mark must be provided

Error responses are returned in JSON:

json
Copy
Edit
{
  "error": "Name is required"
}