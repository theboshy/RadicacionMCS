package routers



import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"../../../../../dao/factory"
	"../../../../../utilities"
	"log"

)

func FinAllUsuario(c *gin.Context) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}


	radicacionDao := factory.FactoryDaoUsuario(config.Engine)
	radicaResult ,_ := radicacionDao.GetAllUsuario()
	c.JSON(200, radicaResult)
}

func GetAlllExtends(c *gin.Context) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}


	radicacionDao := factory.FactoryDaoUsuario(config.Engine)
	radicaResult ,_ := radicacionDao.GetAllUsuarioExtend()
	c.JSON(200, radicaResult)
}
