package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/itnilesh/go-rest-api/api"

	cfg "github.com/itnilesh/go-rest-api/cfg"
	env "github.com/kelseyhightower/envconfig"
	logger "github.com/sirupsen/logrus"
)

// LOG for the project
var LOG = logger.New()
var empSrvCfg cfg.EmpServiceConfig

func init() {
	var LOG = logger.New()
	logger.SetFormatter(&logger.JSONFormatter{})
	LOG.Out = os.Stdout

	err := env.Process("emp_service", &empSrvCfg)
	if err != nil {
		LOG.Fatal(err.Error())
	}
	level,err :=logger.ParseLevel(empSrvCfg.LogLevel);
	if err !=nil {
		LOG.Warn("Could not initialize log level to ", empSrvCfg.LogLevel)
	}
	logger.SetLevel(level)
}

func main() {

	// Load config from environment
	LOG.Info("Loading environment vars")
	


	LOG.WithFields(logger.Fields{
		"config": empSrvCfg.ToString(),
	}).Info("Employee service congfigs")

	LOG.Info("Starting employee service....")
	router := mux.NewRouter()
	empCtr, _ := api.NewEmployeeController()

	LOG.Info("Registerted rest end-point", "/employees")

	router.HandleFunc("/employees", empCtr.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", empCtr.AddEmployee).Methods("POST")

	LOG.Info("Registerted rest end-point", "/employees/{id}")
	router.HandleFunc("/employees/{id}", empCtr.GetEmployee).Methods("GET")
	router.HandleFunc("/employees/{id}", empCtr.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", empCtr.DeleteEmployee).Methods("DELETE")

	LOG.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", empSrvCfg.Port), router))
	LOG.Info("Employee service sarted successfully ..!")
}
