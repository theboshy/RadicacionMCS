package interfaces

import "../../models"

type RadicacionDao interface {
	Create(u *models.Radicacion) error
	/*Update(u *models.Radicacion) error
	Delete(i int) error
	GetById(i int) (models.Radicacion, error)*/
	GetAll() ([]models.Radicacion, error)
}
