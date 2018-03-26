package mysql

import (
	"../../models"
	//"fmt"
)

type RadicacionImplMysql struct {
}

func (dao RadicacionImplMysql) Create(radicaci *models.Radicacion) error {

	return nil
}

func (dao RadicacionImplMysql) GetAll() ([]models.Radicacion, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := get()
	defer db.Close()
	var radci []models.Radicacion
	error :=db.Find(&radci)
	if error != nil {
		return radci,error
	}
		return radci,nil
}

