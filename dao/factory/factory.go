package factory

import (
	"../interfaces"
	"../mysql"
	"log"
)

func FactoryDaoRadicacion(e string) interfaces.RadicacionDao {
	var i interfaces.RadicacionDao

	switch e {
	case "mysql":
		i = mysql.MysqlImplDb{}

	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}

func FactoryDaoUsuario(e string) interfaces.UsuarioDao {
	var i interfaces.UsuarioDao

	switch e {
	case "mysql":
		i = mysql.MysqlImplDb{}

	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}


func FactoryDaoSedes(e string) interfaces.SedesDao {
	var i interfaces.SedesDao

	switch e {
	case "mysql":
		i = mysql.MysqlImplDb{}

	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}