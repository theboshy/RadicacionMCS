package interfaces

import "../../models"

type UsuarioDao interface {
	CreateUsuario(u models.Usuarios) (int64,error)
	FindUsuario(id string,pass string)(models.Usuarios, error)
	GetAllUsuario() ([]models.Usuarios, error)
	GetAllUsuarioExtend() ([]models.UsuariosSedes, error)
}
