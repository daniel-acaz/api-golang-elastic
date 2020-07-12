package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/daniel-acaz/api-golang-elastic/route"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetAllProperty).Methods("GET")
	r.HandleFunc("/{id}", controller.GetPropertyById).Methods("GET")
	r.HandleFunc("/", controller.CreateProperty).Methods("POST")
	r.HandleFunc("/{id}", controller.UpdateProperty).Methods("PUT")
	r.HandleFunc("/{id}", controller.DeleteProperty).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
