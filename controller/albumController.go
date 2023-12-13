package controller

import (
	"gogin/service"

	"github.com/gin-gonic/gin"
)

type AlbumController interface {
	GetAlbums(c *gin.Context)
	GetAlbumById(c *gin.Context)
	PostAlbums(c *gin.Context)
	DeleteAlbum(c *gin.Context)
}

type AlbumControllerImpl struct {
	svc service.AlbumService
}

func (a AlbumControllerImpl) GetAlbums(c *gin.Context) {
	a.svc.GetAlbums(c)
}

func (a AlbumControllerImpl) GetAlbumById(c *gin.Context) {
	a.svc.GetAlbumById(c)
}

func (a AlbumControllerImpl) PostAlbums(c *gin.Context) {
	a.svc.PostAlbums(c)
}

func (a AlbumControllerImpl) DeleteAlbum(c *gin.Context) {
	a.svc.DeleteAlbum(c)
}

func AlbumControllerInit(albumService service.AlbumService) *AlbumControllerImpl {
	return &AlbumControllerImpl{
		svc: albumService,
	}
}
