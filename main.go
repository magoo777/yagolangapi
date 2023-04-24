package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func playerJson(c *gin.Context) {
	c.JSON(http.StatusOK, players)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/api/mats", playerJson)
	router.Run(":8080")
}
