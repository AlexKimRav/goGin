package main

import (
	"gogin/configs"
	"gogin/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var albums = []model.Album{
	{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	configs.Init()
	router := gin.Default()
	router.GET("/hello", getHello)
	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbumById)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err.Error())
	}

	for _, album := range albums {
		if album.Id == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
}

func postAlbums(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
