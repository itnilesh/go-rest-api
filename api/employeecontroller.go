package api

import (
	"encoding/json"
	"github.com/itnilesh/go-rest-api/service"
	"net/http"
)

// GetEmployees returns list of employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, _ := service.GetEmployees()
	json.NewEncoder(w).Encode(employees)
}
