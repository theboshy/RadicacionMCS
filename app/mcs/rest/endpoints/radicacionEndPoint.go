package endpoints

import (
	//"github.com/gorilla/mux"
	"net/http"
	"../../../../dao/factory"
	"../../../../utilities"
	"log"
	"encoding/json"
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
