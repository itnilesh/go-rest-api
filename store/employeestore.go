package store

import (
	"errors"
	"github.com/itnilesh/go-rest-api/service/models"
	"github.com/google/uuid"

)

// Store persistemnce inrterface
type Store interface {
	AddEmployee(emp *models.NewEmployee) error
	UpdateEmployee(id string , emp *models.UpdateEmployee) error
	DeleteEmployee(id string) error
	GetEmployee(id string) (*models.Employee, error)
	GetEmployees() (*[]models.Employee, error)
}


//InMemoryStore in memory store may be used for testing 
type inMemoryStore struct{
	 index map[string] *models.Employee
	 ids []string


}

//NewInMemoryStore Creates new in memory store 
func NewInMemoryStore() Store {
	 return &inMemoryStore{index:make(map[string]* models.Employee ) ,ids : make([]string,0,0)}


}
//GetEmployees get employee list
func ( s *inMemoryStore) GetEmployees() (*[]models.Employee, error) {
	list :=make([]models.Employee,0,0)
	for _ ,id := range  s.ids {
		emp, ok := s.index[id]
		if ok {
			list = append(list, *emp)
		}
	}
	return &list,nil
}


//GetEmployees get employee list
func ( s *inMemoryStore) GetEmployee(id string ) (*models.Employee, error) {
	employee, ok := s.index[id]
	if !ok {
		return nil, errors.New("Employee not found with id")
	}
	return employee,nil
}


func ( s *inMemoryStore) AddEmployee(emp *models.NewEmployee) error  {
	  uuid,err := uuid.NewUUID()
	  if err !=nil {
		 return errors.New("Failed to generate UUID")
	  }
	  id := uuid.String()
	  // Ideally this should be cloned
	  employee := &models.Employee{ID:id, Emp:emp }
	  s.ids = append(s.ids, id)
	  s.index[id] = employee
	  return nil
}

func ( s *inMemoryStore) UpdateEmployee(id string , emp *models.UpdateEmployee) error  {

	employee, ok := s.index[id]
	if  ok {
		employee.Emp.Address = emp.Address;
		return nil
	}
	return errors.New("Employee not found with id")
	
}


func ( s *inMemoryStore) DeleteEmployee(id string) error  {
	// I am not sure this is automic operation , seemns to be some race condition 
	delete(s.index, id)
	return nil
}


