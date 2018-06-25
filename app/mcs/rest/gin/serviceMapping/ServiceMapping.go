package serviceMapping

import (
	"../routers"
	"github.com/gin-gonic/gin"
	//"github.com/appleboy/gin-jwt"
)

func MapRouterGroup(router gin.Engine /*,authmidleware *jwt.GinJWTMiddleware*/) {
	/* oauth interno - router.POST("/login",authmidleware.LoginHandler)
	router.Use(authmidleware.MiddlewareFunc())*/

	//primer grupo de empaquetado (mappings para mcs/principal)
	mappingsUrl := router.Group("mcs/principal")
	{
		mappingsUrl.GET("/instructions", routers.GetInstructions)
		mappingsUrl.GET("/findALLUsuarioExtend", routers.GetAlllExtends)
		mappingsUrl.GET("/findALLUsuario", routers.FinAllUsuario)
		mappingsUrl.POST("/captureFile", routers.CaptureFile)
		//mappingsUrl.POST("/findByUsuario", routers.FindUsuario)
		mappingsUrl.POST("/findBySede", routers.FindBySede)
		mappingsUrl.GET("/delivery", routers.GetDelivery)
		mappingsUrl.POST("/captureRecords", routers.PostConsoleParams)
		mappingsUrl.GET("/findAllRadicacion", routers.FinAllRadicacion)
		mappingsUrl.POST("/setNewRadicacion", routers.SetNewRadicacion)
		mappingsUrl.GET("/findByIdRadicacion", routers.FindByIdRadicacion)
	}

	//segundo grupo de empaquetado (mappings para mcs/secondary)
	mappingsUrl2 := router.Group("mcs/secondary")
	{
		mappingsUrl2.GET("/instructions2", routers.GetInstructions2)
		mappingsUrl2.GET("/delivery2", routers.GetDelivery2)
	}

	auth := router.Group("/auth")

	{
		auth.GET("/makeMeHappy", routers.HapinnesHandlerRouter)
		//auth.GET("/refresh_token", authmidleware.RefreshHandler)
	}
}
