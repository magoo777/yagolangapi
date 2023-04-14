package main

import (
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
)

type PageView struct {
	Title  string
	Rubrik string
}

var theRandom *rand.Rand

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "Testar", Rubrik: "Hej Golang, please work"})
}

// HTML
// JSON

func playerJson(c *gin.Context) {
	var player = struct {
		Name string `json:"name"`
		City string `json:"city"`
	}{
		Name: "Wayne Rooney",
		City: "Manchester",
	}

	c.JSON(http.StatusOK, player)
}

var config Config

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/matsgudmundsson", playerJson)
	router.Run(":8080")

}
