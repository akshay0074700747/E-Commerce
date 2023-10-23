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

func InitializeAPI1(config config.Config) (*routes.GinEngine, error) {

	wire.Build(database.Connect_to,
		adapters.NewUserRepository,
		adapters.NewAdminRepository,
		adapters.NewSuAdminRepository,
		adapters.NewCategoryRepository,
		adapters.NewProductDataBase,
		usecases.NewUserUseCase,
		usecases.NewAdminUsecase,
		usecases.NewSuAdminUsecase,
		usecases.NewCategoryUsecase,
		usecases.NewProductUsecases,
		handlers.NewUserHandler,
		handlers.NewAdminHandler,
		handlers.NewSuAdminHandler,
		handlers.NewCategoryHandler,
		handlers.NewProductHandler,
		routes.NewGinEngine)

	return &routes.GinEngine{}, nil

}
