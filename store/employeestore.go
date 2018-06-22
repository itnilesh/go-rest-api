package store

import (
	"strconv"
	"github.com/itnilesh/go-rest-api/api/models"
)

//GetEmployees get employee list
func GetEmployees() (*[]models.Employee, error) {
	employees := make([]models.Employee, 0)

	for i := 0; i < 5; i++ {
		index := strconv.Itoa(i)
		fname := "Nilesh_" + index
		lname := "Salpe_" + index
		emp := models.Employee{ID: index, Fname: fname, Lname: lname}
		employees = append(employees, emp)
	}
	return &employees, nil
}
