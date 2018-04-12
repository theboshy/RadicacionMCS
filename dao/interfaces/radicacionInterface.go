package interfaces

import "../../models"

type RadicacionDao interface {
	CreateRadicacion(u models.Radicacion) (int64,error)
	/*Update(u *models.Radicacion) error
	Delete(i int) error
	GetById(i int) (models.Radicacion, error)*/
	FindRadicacion(id string)(models.Radicacion, error)
	GetAllRadicacion() ([]models.Radicacion, error)
}
