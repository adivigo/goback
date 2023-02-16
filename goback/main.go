package main

import (
	"log"
	"net/http"

	"github.com/adivigo/goback/controllers/vehiclecontroller"
	"github.com/adivigo/goback/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/vehicle", vehiclecontroller.Index).Methods("GET")
	// r.HandleFunc("/vehicle/{id}", vehiclecontroller.Show).Methods("GET")
	// r.HandleFunc("/vehicle", vehiclecontroller.Create).Methods("POST")
	// r.HandleFunc("/vehicle/{id}", vehiclecontroller.Update).Methods("PUT")
	// r.HandleFunc("/vehicle", vehiclecontroller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
