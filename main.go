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

func main() {
	LOG.Out = os.Stdout
	LOG.Info("Starting server..")
	router := mux.NewRouter()
	router.HandleFunc("/employees", api.GetEmployees).Methods("GET")
	LOG.Fatal(http.ListenAndServe(":8000", router))
}
