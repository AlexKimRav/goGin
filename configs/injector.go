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

var db = wire.NewSet(GetDb())

var albumServiceSet = wire.NewSet(service.AlbumServiceInit,
	wire.Bind(new(service.AlbumService), new(*service.AlbumServiceImpl)),
)

var albumRepoSet = wire.NewSet(repository.AlbumRepositoryInit,
	wire.Bind(new(repository.AlbumRepository), new(*repository.AlbumRepositoryImpl)),
)

var albumCtrlSet = wire.NewSet(controller.AlbumControllerInit,
	wire.Bind(new(controller.AlbumController), new(*controller.AlbumControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, albumCtrlSet, albumServiceSet, albumRepoSet)
	return nil
}
