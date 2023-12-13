package router

import (
	"gogin/configs"

	"github.com/gin-gonic/gin"
)

// TODO add Logger in router.Use(log)
func Init(init *configs.Initialization) *gin.Engine {
	router := gin.Default()
	// router.Use()

	api := router.Group("/api")
	{
		album := api.Group("/albums")
		album.GET("", init.AlbumCtrl.GetAlbums)
		album.GET("/:albumId", init.AlbumCtrl.GetAlbumById)
		album.POST("", init.AlbumCtrl.PostAlbums)
		album.DELETE("/:albumId", init.AlbumCtrl.DeleteAlbum)
	}
	return router
}
