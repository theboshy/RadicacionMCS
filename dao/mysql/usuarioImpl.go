package mysql

import (
	"../../models"
	//"fmt"
	"fmt"
)

func (dao MysqlImplDb) CreateUsuario(model models.Usuarios) (int64, error) {
	db := getConection()
	defer db.Close()
	radci := model
	affec,error :=db.Insert(&radci)
	if error != nil {
		return affec,error
	}
	return affec,nil
}

func (dao MysqlImplDb) FindUsuario(id string,pass string) (models.Usuarios, error) {
	db := getConection()
	defer db.Close()
	radii := models.Usuarios{}
	var affect bool
	var error error
	if pass !="" {
		affect, error = db.Where("nom_id_usuario = ?", id).Where("contrasena = ?", pass).Get(&radii)
	}else{
		affect, error = db.Where("nom_id_usuario = ?", id).Get(&radii)
	}
	if error != nil || affect == false{
		return radii,error
	}
	return radii,nil
}

func (dao MysqlImplDb) GetAllUsuario() ([]models.Usuarios, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := getConection()
	defer db.Close()
	var radii []models.Usuarios
	error :=db.Find(&radii)
	if error != nil {
		return radii,error
	}
	return radii,nil
}

func (dao MysqlImplDb) GetAllUsuarioExtend() ([]models.UsuariosSedes, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := getConection()
	defer db.Close()

	var users []models.UsuariosSedes
	err := engine.Join("INNER", "sedes", "sedes.idsedes = usuarios.sedes_idsedes").Find(&users)
	fmt.Println(err)
	return users,err
}



