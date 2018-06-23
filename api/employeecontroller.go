package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	apimodels "github.com/itnilesh/go-rest-api/api/models"
	"github.com/itnilesh/go-rest-api/service"
	srvmodels "github.com/itnilesh/go-rest-api/service/models"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// LOG - Lgger
var LOG = log.New()

func init() {
	LOG.Out = os.Stdout
}

// EmployeeController - class
type EmployeeController struct {
	EmployeeService *service.EmployeeService
}

// NewEmployeeController - creates new rest controller
func NewEmployeeController() (*EmployeeController, error) {
	service, _ := service.NewEmployeeService()
	return &EmployeeController{EmployeeService: service}, nil
}

// AddEmployee - Add new employee
func (e *EmployeeController) AddEmployee(w http.ResponseWriter, r *http.Request) {
	createEmpReq := &apimodels.CreateEmployeeReq{}
	json.NewDecoder(r.Body).Decode(createEmpReq)
	e.EmployeeService.AddEmployee(convEmployeeSrvModel(createEmpReq))
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func convEmployeeSrvModel(cer *apimodels.CreateEmployeeReq) *srvmodels.NewEmployee {
	newEmp := &srvmodels.NewEmployee{
		Fname:   cer.Fname,
		Lname:   cer.Lname,
		Address: convEmpAddrSrvModel(cer.Address)}
	return newEmp
}

func convEmpAddrSrvModel(edr *apimodels.EmpAdress) *srvmodels.Address {
	adr := &srvmodels.Address{HouseNumber: edr.HouseNumber,
		Street:   edr.Street,
		ZipCode:  edr.ZipCode,
		District: edr.District,
		City:     edr.City,
		State:    edr.State,
		Country:  edr.Country}
	return adr

}

// UpdateEmployee - update employee
func (e *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	updateEmpReq := &apimodels.UpdateEmployeeReq{}
	json.NewDecoder(r.Body).Decode(updateEmpReq)
	updateEmp := convUpdateEmpSrvModel(updateEmpReq)
	e.EmployeeService.UpdateEmployee(id, updateEmp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
}

func convUpdateEmpSrvModel(uer *apimodels.UpdateEmployeeReq) *srvmodels.UpdateEmployee {
	updateEmp := &srvmodels.UpdateEmployee{Address: convEmpAddrSrvModel(uer.Address)}
	return updateEmp
}

// DeleteEmployee - delete employee
func (e *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	e.EmployeeService.DeleteEmployee(id)
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("content-type", "application/json")
}

// GetEmployee - get single employee by id
func (e *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	emp, _ := e.EmployeeService.GetEmployee(id)
	empResp, _ := convEmployeeResp(emp)
	empRespJSON, _ := json.Marshal(empResp)
	LOG.Info(emp.Emp.Address.HouseNumber)
	LOG.Info(string(empRespJSON))
	w.Write(empRespJSON)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// GetEmployees returns list of employees
func (e *EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, _ := e.EmployeeService.GetEmployees()
	employeesResp, _ := convEmployeesToResp(employees)
	// convert to view
	json.NewEncoder(w).Encode(employeesResp)
}

func convEmployeesToResp(srvmodels *[]srvmodels.Employee) (*[]apimodels.EmployeeResp, error) {
	respList := make([]apimodels.EmployeeResp, 0, 0)

	for _, emp := range *srvmodels {
		empResp, _ := convEmployeeResp(&emp)
		respList = append(respList, *empResp)
	}

	return &respList, nil
}

func convEmployeeResp(srvm *srvmodels.Employee) (*apimodels.EmployeeResp, error) {
	addressResp, _ := convEmpAddress(srvm.Emp.Address)
	resp := &apimodels.EmployeeResp{ID: srvm.ID, 
		Fname: srvm.Emp.Fname, 
		Lname: srvm.Emp.Lname, 
		Address: addressResp}
	return resp, nil
}

func convEmpAddress(srva *srvmodels.Address) (*apimodels.EmpAdress, error) {
	resp := &apimodels.EmpAdress{
	  HouseNumber:srva.HouseNumber,
	  Street:srva.State,
	  ZipCode:srva.ZipCode,
	  District:srva.District,
	  State:srva.State,
	  Country:srva.Country}

	return resp, nil
}
