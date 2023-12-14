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
func NewInitialization(userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,
	}
}
