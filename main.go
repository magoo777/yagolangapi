package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageView struct {
	Title  string
	Rubrik string
}

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "Jag testar", Rubrik: "Hej Golang, please work"})
}

// HTML
// JSON

var player = struct {
	Name string `json:"name"`
	City string `json:"city"`
}{
	Name: "Wayne Rooney",
	City: "Manchester",
}

func playerJson(c *gin.Context) {

	c.JSON(http.StatusOK, player)
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)
	router.GET("/api/mittnamn", playerJson)
	router.Run(":8080")

}
