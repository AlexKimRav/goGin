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
