package routers

import (
	"github.com/gin-gonic/gin"
	//"fmt"
)

func GetInstructions2(c *gin.Context) {
	c.JSON(200, gin.H{"part1/2": "is the finaalal countdownnn"})
	return
}

func GetDelivery2(c *gin.Context) {
	c.XML(200, gin.H{"part2/2": "tuurururur tutu tututu rtuuu"})
}
