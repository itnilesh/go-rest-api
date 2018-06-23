package models


// CreateEmployeeReq - model for create employee
type CreateEmployeeReq struct {
	Fname string `json:"first_name,omitempty"`
	Lname string `json:"last_name,omitempty"`
	Address *EmpAdress `json:"address,omitempty"`
}
// EmployeeResp api model for employee
type EmployeeResp struct {
	ID    string `json:"id"`
	Fname string `json:"first_name,omitempty"`
	Lname string `json:"last_name,omitempty"`
	Address *EmpAdress `json:"address,omitempty"`
}


// UpdateEmployeeReq - api model to udpate
type UpdateEmployeeReq struct {
	Address *EmpAdress `json:"address,omitempty"`
}


// EmpAdress - api model for address
type EmpAdress struct {
	HouseNumber string `json:"house_number,omitempty"`
	Street string  `json:"street,omitempty"`
	ZipCode string  `json:"zip_code,omitempty"`
	City string  `json:"city,omitempty"`
	District string  `json:"district,omitempty"`
	State string  `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
}


// APIErrorResp - Api error response
type APIErrorResp struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Target string  `target:"message,omitempty"`
}

