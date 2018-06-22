package service

import ()
import ("github.com/itnilesh/go-rest-api/store"
 "github.com/itnilesh/go-rest-api/api/models")


//GetEmployees list employees
func GetEmployees() (*[]models.Employee, error) {

	return store.GetEmployees()

}
