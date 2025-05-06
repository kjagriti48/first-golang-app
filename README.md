# 📘 Student API (Golang)

A modular Go backend that manages student records, calculates average marks, determines top performers, and saves data to a local file.

---

## 🚀 Features

•⁠  ⁠Modular folder structure (⁠ handlers ⁠, ⁠ routes ⁠, ⁠ models ⁠, ⁠ utils ⁠)
•⁠  ⁠⁠ GET ⁠, ⁠ POST ⁠, ⁠ DELETE ⁠ endpoints
•⁠  ⁠⁠ GET /top-student ⁠ to return the highest scorer
•⁠  ⁠File persistence using ⁠ students.json ⁠
•⁠  ⁠Status logic: ⁠ pass ⁠ or ⁠ fail ⁠ based on average
•⁠  ⁠Input Validation for POST requests
•⁠  Error response structure using 'writeError()'
•⁠  ⁠HTTP Status codes: 400, 404, 201
•⁠  ⁠Cleaner JSON formatting for all endpoints

---

## 🧪 API Endpoints

### ⁠ GET /students ⁠
Returns a list of all students.

### ⁠ POST /students ⁠
Add a student.

#### Example Request Body:
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


