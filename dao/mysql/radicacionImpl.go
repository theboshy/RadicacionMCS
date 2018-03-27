package mysql

import (
	"../../models"
	//"fmt"
)

func (dao MysqlImplDb) Create(radicaci *models.Radicacion) error {

	return nil
}

func (dao MysqlImplDb) GetAll() ([]models.Radicacion, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := GetInstance()
	defer db.Close()
	var radci []models.Radicacion
	error :=db.Find(&radci)
	if error != nil {
		return radci,error
	}
		return radci,nil
}

