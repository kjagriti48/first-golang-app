package handlers

import (
	"encoding/json"
	"first-golang-app/models"
	"first-golang-app/utils"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Students)
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

	models.Students = append(models.Students, s)
	utils.SaveStudentsToFile()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func deleteStudentByName(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/students/")
	if name == "" {
		writeError(w, http.StatusBadRequest, "Student name required in path")
		return
	}
	found := false
	var updated []models.Student
	for _, s := range models.Students {
		if s.Name != name {
			updated = append(updated, s)
		} else {
			found = true
		}
	}
	if !found {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	models.Students = updated
	utils.SaveStudentsToFile()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
}

func TopStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}

	if len(models.Students) == 0 {
		writeError(w, http.StatusNotFound, "No students Found")
		return
	}
	top := models.Students[0]
	topAvg := utils.CalculateAverage(top.Marks)
	for _, s := range models.Students[1:] {
		avg := utils.CalculateAverage(s.Marks)
		if avg > topAvg {
			top = s
			topAvg = avg
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(top)
}
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
