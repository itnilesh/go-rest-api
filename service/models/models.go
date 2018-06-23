//Ideally API models and service models should be different and evolve as per need 
// I see people keep same and create problems when using pagation or gathering data from multiple models

package models


// Employee api model for employee
type Employee struct {
	ID    string 
	Emp *NewEmployee
}

// Address - Address of employee
type Address struct {
	HouseNumber string
	Street string 
	ZipCode string 
	City string 
	District string 
	State string 
	Country  string
}

// UpdateEmployee contains updatable fields of employee
type UpdateEmployee struct {
	Address *Address
}

// NewEmployee creates new employee
type NewEmployee struct{
	Fname string 
	Lname string 
	Address *Address

}
