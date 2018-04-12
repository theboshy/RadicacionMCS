package mysql

import (
	"../../models"
	//"fmt"
)

func (dao MysqlImplDb) CreateRadicacion(model models.Radicacion) (int64, error) {
	db := getConection()
	defer db.Close()
	radci := model
	affec,error :=db.Insert(&radci)
	if error != nil {
		return affec,error
	}
	return affec,nil
}

func (dao MysqlImplDb) FindRadicacion(id string) (models.Radicacion, error) {
	db := getConection()
	defer db.Close()
	radii := models.Radicacion{}
	affect,error :=db.Where("numero_radicacion = ?", id).Get(&radii)
	if error != nil || affect == false{
		return radii,error
	}
	return radii,nil
}

func (dao MysqlImplDb) GetAllRadicacion() ([]models.Radicacion, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := getConection()
	defer db.Close()
	var radii []models.Radicacion
	error :=db.Find(&radii)
	if error != nil {
		return radii,error
	}
		return radii,nil
}


