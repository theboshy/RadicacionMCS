package factory

import (
	"../interfaces"
	"../mysql"
	"log"
)

func FactoryDao(e string) interfaces.RadicacionDao {
	var i interfaces.RadicacionDao
	switch e {
	case "mysql":
		i = mysql.RadicacionImplMysql{}
	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}
