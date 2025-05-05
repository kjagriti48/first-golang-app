package utils

import (
    "encoding/json"
    "fmt"
    "os"
    "first-golang-app/models"
)

const filePath = "students.json"

func LoadStudentsFromFile() {
    data, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Starting with empty student list.")
        models.Students = []models.Student{}
        return
    }
    json.Unmarshal(data, &models.Students)
    fmt.Println("Loaded students from file.")
}

func SaveStudentsToFile() {
    data, err := json.MarshalIndent(models.Students, "", "  ")
    if err != nil {
        fmt.Println("Failed to save:", err)
        return
    }
    os.WriteFile(filePath, data, 0644)
}

func CalculateAverage(marks map[string]int) float64 {
    total := 0
    for _, score := range marks {
        total += score
    }
    return float64(total) / float64(len(marks))
}