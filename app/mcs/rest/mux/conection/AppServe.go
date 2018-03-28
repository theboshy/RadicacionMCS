package main

import (
	//"encoding/json"
	//"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	end "../endpoints"
	//"encoding/json"
	//"fmt"
	//"os"
	//"fmt"
)


func main() {

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})


	router := mux.NewRouter()

	router.HandleFunc("/findAllRadicacion", end.GetAllRadicacionData).Methods("GET")
	router.HandleFunc("/getCurrentRadicacion", end.GetCurrentRadicacionData).Methods("GET")
	router.HandleFunc("/setNewRadicacion", end.SetNewRadicacionData).Methods("POST")

	/*router.HandleFunc("/.../{id}", end....).Methods("GET")*/
	//log.Fatal(http.ListenAndServe(":8000", router))
	http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
}
