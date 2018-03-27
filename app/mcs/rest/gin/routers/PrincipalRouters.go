package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetInstructions(c *gin.Context) {
	c.JSON(200, gin.H{"part1": "welcome to the jungle"})
}

func GetDelivery(c *gin.Context) {
	c.XML(200, gin.H{"part2": "If you want it you're gonna bleed but it's the price to pay"})
}

func PostConsoleParams(c *gin.Context) {
	id := c.Query("id")
	name := c.DefaultQuery("name", "0")
	valor1 := c.PostForm("valor1")
	message := c.PostForm("message")
	fmt.Printf("id: %s; valor1: %s; name: %s; message: %s", id, valor1, name, message)
}
