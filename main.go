package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "AdaDeSion",
		"status":  http.StatusOK,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", GetMessage)

	r.GET("/ada/:id/*action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")
		message := id + " is " + action

		c.String(http.StatusOK, message)
	})

	r.Run()
}
