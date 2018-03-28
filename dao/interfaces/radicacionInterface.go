package interfaces

import "../../models"

type RadicacionDao interface {
	Create(u models.Radicacion) (int64,error)
	/*Update(u *models.Radicacion) error
	Delete(i int) error
	GetById(i int) (models.Radicacion, error)*/
	Find(id string)(models.Radicacion, error)
	GetAll() ([]models.Radicacion, error)
}
