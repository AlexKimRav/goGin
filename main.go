package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"price"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/hello", getHello)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
