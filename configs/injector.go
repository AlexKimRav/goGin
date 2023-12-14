// go:build wireinject
//go:build wireinject
// +build wireinject

//It would be a trouble if we inject each
// component manually in Go, especially when our program becomes bigger
// and complex. In this case we’re using
// google wire library to handle dependency injection in our components.

package configs

//TODO change user to album
import (
	"gogin/controller"
	"gogin/repository"
	"gogin/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet)
	return nil
}
