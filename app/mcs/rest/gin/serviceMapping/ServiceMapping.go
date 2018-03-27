package serviceMapping

import (
	"../routers"
	"github.com/gin-gonic/gin"
)

func MapRouterGroup(router *gin.Engine) {
	//primer grupo de empaquetado (mappings para mcs/principal)
	mappingsUrl := router.Group("mcs/principal")
	{
		mappingsUrl.GET("/instructions", routers.GetInstructions)
		mappingsUrl.GET("/delivery", routers.GetDelivery)
		mappingsUrl.POST("/captureRecords", routers.PostConsoleParams)
	}

	//segundo grupo de empaquetado (mappings para mcs/secondary)
	mappingsUrl2 := router.Group("mcs/secondary")
	{
		mappingsUrl2.GET("/instructions2", routers.GetInstructions2)
		mappingsUrl2.GET("/delivery2", routers.GetDelivery2)
	}
}
