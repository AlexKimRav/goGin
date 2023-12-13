package service

import (
	"fmt"
	"gogin/model"
	"gogin/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumService interface {
	GetAlbums(c *gin.Context)
	GetAlbumById(c *gin.Context)
	PostAlbums(c *gin.Context)
	DeleteAlbum(c *gin.Context)
}

type AlbumServiceImpl struct {
	albumRepository repository.AlbumRepository
}

func (a AlbumServiceImpl) GetAlbums(c *gin.Context) {
	data, err := a.albumRepository.FindAllAlbums()
	if err != nil {
		fmt.Println("Happened Error when find all album data. Error: ", err)
	}

	c.JSON(http.StatusOK, data)
}

func (a AlbumServiceImpl) GetAlbumById(c *gin.Context) {
	albumID, _ := strconv.Atoi(c.Param("albumId"))

	data, err := a.albumRepository.FindAlbumById(albumID)
	if err != nil {
		fmt.Println("Happened error when get data from database. Error", err)
	}

	c.JSON(http.StatusOK, data)
}

func (a AlbumServiceImpl) PostAlbums(c *gin.Context) {
	var request model.Album
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("Happened error when mapping request from FE. Error", err)
	}

	data, err := a.albumRepository.Save(&request)
	if err != nil {
		fmt.Println("Happened error when saving data to database. Error", err)
	}

	c.JSON(http.StatusOK, data)
}

func (u AlbumServiceImpl) DeleteAlbum(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Param("albumId"))

	err := u.albumRepository.DeleteAlbumById(albumId)
	if err != nil {
		fmt.Println("Happened Error when try delete data album from DB. Error:", err)
	}

	c.JSON(http.StatusOK, nil)
}

func AlbumServiceInit(albumRepository repository.AlbumRepository) *AlbumServiceImpl {
	return &AlbumServiceImpl{
		albumRepository: albumRepository,
	}
}
