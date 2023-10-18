package wireeeeeee

import (
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	routes "ecommerce/web/api/Routes"
	"ecommerce/web/api/handlers"
	"ecommerce/web/config"
	"ecommerce/web/database"

	"github.com/google/wire"
)

func InitializeAPI(config config.Config) (*routes.GinEngine, error) {

	wire.Build(database.Connect_to, adapters.NewUserRepository, usecases.NewUserUseCase, handlers.NewUserHandler, routes.NewGinEngine)

	return &routes.GinEngine{}, nil

}
