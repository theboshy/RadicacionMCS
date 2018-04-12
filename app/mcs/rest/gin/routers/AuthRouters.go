package routers


import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
)

func HapinnesHandlerRouter(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "sup| nigg**",
	})
}