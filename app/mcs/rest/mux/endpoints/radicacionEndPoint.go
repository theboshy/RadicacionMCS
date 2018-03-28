package endpoints

import (
	//"github.com/gorilla/mux"
	"net/http"
	"../../../../../dao/factory"
	"../../../../../utilities"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)



func GetAllRadicacionData(w http.ResponseWriter, r *http.Request) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	radicacionDao := factory.FactoryDao(config.Engine)
	radicaResult ,_ := radicacionDao.GetAll()
	json.NewEncoder(w).Encode(radicaResult)
}

func GetCurrentRadicacionData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(params["param1"])
	/*config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	radicacionDao := factory.FactoryDao(config.Engine)
	radicaResult ,_ := radicacionDao.GetAll()*/
	json.NewEncoder(w).Encode(params)
}

func SetNewRadicacionData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println(params["param1"])
	/*config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}
	radicacionDao := factory.FactoryDao(config.Engine)
	radicaResult ,_ := radicacionDao.GetAll()*/
	json.NewEncoder(w).Encode(params)
}
