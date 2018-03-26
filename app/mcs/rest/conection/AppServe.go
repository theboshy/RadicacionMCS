package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	end "../endpoints"
	//"encoding/json"
	//"fmt"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/findAllRadicacion", end.GetAllRadicacionData).Methods("GET")
	/*router.HandleFunc("/.../{id}", end....).Methods("GET")*/
	log.Fatal(http.ListenAndServe(":8000", router))
}