package models

type Student struct {
    Name   string         `json:"name"`
    Age    int            `json:"age"`
    Marks  map[string]int `json:"marks"`
    Status string         `json:"status"`
}

var Students []Student