package routers



import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"../../../../../dao/factory"
	"../../../../../utilities"
	"log"

)

func FindBySede(c *gin.Context) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}


	radicacionDao := factory.FactoryDaoSedes(config.Engine)
	radicaResult ,_ := radicacionDao.FindBySedes(1)
	c.JSON(200, radicaResult)
}