package service

import ()
import (
	 "github.com/itnilesh/go-rest-api/service/models"
	 errorz "github.com/itnilesh/go-rest-api/service/errors"
	"github.com/itnilesh/go-rest-api/store"
)


// EmployeeService - Service class
type EmployeeService struct {
	Store store.Store

}

// NewEmployeeService  - creates new employee service
func NewEmployeeService() (*EmployeeService, errorz.EmployeeError){
	store:= store.NewInMemoryStore()
	return &EmployeeService{Store:store}, nil
}


//AddEmployee Add employee to registory
func (e *EmployeeService) AddEmployee(newEmp *models.NewEmployee) errorz.EmployeeError {
	 err := e.Store.AddEmployee(newEmp)
	return &errorz.GenericEmployeeError{Errorz:err}
}

//GetEmployee get employee from registory
func (e *EmployeeService) GetEmployee(id string)( *models.Employee,  errorz.EmployeeError) {
	employee, err := e.Store.GetEmployee(id)
   return employee, &errorz.GenericEmployeeError{Errorz:err}
}

//UpdateEmployee - update employee details
func (e *EmployeeService) UpdateEmployee(id string, updateEmp *models.UpdateEmployee)  errorz.EmployeeError {
	 err := e.Store.UpdateEmployee(id, updateEmp)
	return &errorz.GenericEmployeeError{Errorz:err}
}


// DeleteEmployee - delete employee from recoirds 
func (e *EmployeeService) DeleteEmployee(id string) errorz.EmployeeError {
	err := e.Store.DeleteEmployee(id)
	return &errorz.GenericEmployeeError{Errorz:err}
}

// GetEmployees - get list of employees
func (e *EmployeeService) GetEmployees() (*[]models.Employee, errorz.EmployeeError) {
	employees, err := e.Store.GetEmployees()
	return employees, &errorz.GenericEmployeeError{Errorz:err}
}

