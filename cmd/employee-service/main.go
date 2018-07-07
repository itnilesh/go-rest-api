package main

import (
	"github.com/gorilla/mux"
	"github.com/itnilesh/go-rest-api/api"
	"net/http"

	logger "github.com/sirupsen/logrus"
	"os"
)

// LOG for the project
var LOG = logger.New()

func init() {
	LOG.Out = os.Stdout
}

func main() {

	LOG.Info("Starting server..")
	router := mux.NewRouter()
	empCtr, _ := api.NewEmployeeController()

	router.HandleFunc("/employees", empCtr.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", empCtr.AddEmployee).Methods("POST")

	router.HandleFunc("/employees/{id}", empCtr.GetEmployee).Methods("GET")
	router.HandleFunc("/employees/{id}", empCtr.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", empCtr.DeleteEmployee).Methods("DELETE")

	LOG.Fatal(http.ListenAndServe(":8000", router))
}
