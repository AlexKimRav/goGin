package configs

import (
	"gogin/controller"
	"gogin/repository"
	"gogin/service"
)

type Initialization struct {
	albumRepo repository.AlbumRepository
	albumSvc  service.AlbumService
	AlbumCtrl controller.AlbumController
}

// Change user to album
func NewInitialization(albumRepo repository.AlbumRepository,
	albumService service.AlbumService,
	albumController controller.AlbumController) *Initialization {
	return &Initialization{
		albumRepo: albumRepo,
		albumSvc:  albumService,
		AlbumCtrl: albumController,
	}
}
