package handlers

import (
	"encoding/json"
	"first-golang-app/models"
	"first-golang-app/utils"
	"log"
	"net/http"
	"strings"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/students":
		getStudents(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/students":
		addStudent(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/students/"):
		deleteStudentByName(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.DB.Query("SELECT name, age, marks, status FROM students")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch students from the database")
		return
	}

	defer rows.Close()

	students := make([]models.Student, 0)

	for rows.Next() {
		var s models.Student
		var marksText string

		err := rows.Scan(&s.Name, &s.Age, &marksText, &s.Status)
		if err != nil {
			log.Printf("Scan failed: %v", err)
			writeError(w, http.StatusInternalServerError, "Error reading student record")
			return
		}

		json.Unmarshal([]byte(marksText), &s.Marks)
		students = append(students, s)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	//Validation
	if s.Name == "" {
		writeError(w, http.StatusBadRequest, "Name is required")
		return
	}

	if s.Age <= 0 {
		writeError(w, http.StatusBadRequest, "Age should be greater than 0")
		return
	}

	if len(s.Marks) == 0 {
		writeError(w, http.StatusBadRequest, "At least one subject with marks is required")
		return
	}
	total := 0
	for _, mark := range s.Marks {
		total += mark
	}
	avg := float64(total) / float64(len(s.Marks))
	if avg >= 70 {
		s.Status = "pass"
	} else {
		s.Status = "fail"
	}

	//Convert marks map to JSON string
	marksJSON, err := json.Marshal(s.Marks)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to encode marks")
		return
	}

	//Insert into DB
	_, err = utils.DB.Exec(
		`INSERT INTO students (name, age, marks, status) VALUES ($1, $2, $3, $4)`,
		s.Name, s.Age, string(marksJSON), s.Status,
	)

	if err != nil {
		log.Printf("Insert failed: %v", err)
		writeError(w, http.StatusInternalServerError, "Could not insert into database")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
	log.Printf("Added Student: %s", s.Name)
}

func deleteStudentByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/students/")
	if name == "" {
		writeError(w, http.StatusBadRequest, "Student name required in path")
		return
	}

	result, err := utils.DB.Exec("DELETE FROM students WHERE name = ?", name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		writeError(w, http.StatusNotFound, "Student not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
	log.Printf("Deleted Student: %s", name)
}

func TopStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	rows, err := utils.DB.Query("SELECT name, age, marks, status FROM students")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch students")
		return
	}

	defer rows.Close()

	var top models.Student
	var topAvg float64
	found := false

	for rows.Next() {
		var s models.Student
		var marksJSON string

		err := rows.Scan(&s.Name, &s.Age, &marksJSON, &s.Status)
		if err != nil {
			log.Printf("Scan failed: %v", err)
			writeError(w, http.StatusInternalServerError, "Error reading student")
			return
		}
		err = json.Unmarshal([]byte(marksJSON), &s.Marks)
		if err != nil {
			log.Printf("Failed to unmarshal marks for %s: %v", s.Name, err)
			continue
		}
		avg := utils.CalculateAverage(s.Marks)
		if !found || avg > topAvg {
			top = s
			topAvg = avg
			found = true
		}

	}
	if !found {
		writeError(w, http.StatusNotFound, "No students found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(top)
	log.Printf("Top Student Requested")
}
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
