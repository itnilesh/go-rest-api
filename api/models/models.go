package models

// Employee api model for employee
type Employee struct {
	ID    string `json:"id"`
	Fname string `json:"fname,omitempty"`
	Lname string `json:"lname,omitempty"`
}
