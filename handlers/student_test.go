package handlers

import (
	"bytes"
	"first-golang-app/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddStudentCases(t *testing.T) {
	utils.InitDB()

	//Clean table before testing
	_, err := utils.DB.Exec("DELETE FROM students")
	if err != nil {
		t.Fatalf("Failed to clear table, %v", err)
	}

	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Empty JSON body",
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Name is required",
		},
		{
			name: "Missing age",
			body: `{
				"name": "MissingAge",
				"marks": {"math": 90}
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Age should be greater than 0",
		},
		{
			name: "Missing name",
			body: `{
				"age": 20,
				"marks": {"math": 90}
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Name is required",
		},
		{
			name: "Empty marks",
			body: `{
				"name": "NoMarks",
				"age": 20,
				"marks": {}
			}`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "At least one subject with marks is required",
		},
		{
			name: "Valid input",
			body: `{
				"name": "UnitTestStudent",
				"age": 18,
				"marks": {"math": 90, "science": 85}
			}`,
			expectedStatus: http.StatusCreated,
			expectedError:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer([]byte(tc.body)))
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()

			addStudent(res, req)

			if res.Code != tc.expectedStatus {
				t.Errorf("[%s] expected status %d, got %d", tc.name, tc.expectedStatus, res.Code)
			}

			// Optional: check if error message is in response
			if tc.expectedError != "" && !strings.Contains(res.Body.String(), tc.expectedError) {
				t.Errorf("[%s] expected error %q, got %q", tc.name, tc.expectedError, res.Body.String())
			}
		})
	}
}
