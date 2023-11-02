package wireeeeeee

import (
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	routes "ecommerce/web/api/Routes"
	"ecommerce/web/api/handlers"
	"ecommerce/web/api/middlewares"
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
		adapters.NewBrandRepository,
		adapters.NewDiscountAdapter,
		middlewares.NewUserAuthentication,
		adapters.NewCartAdapter,
		adapters.NewWishListAdapter,
		adapters.NewAddressAdapter,
		adapters.NewOrderAdapter,
		usecases.NewUserUseCase,
		usecases.NewAdminUsecase,
		usecases.NewSuAdminUsecase,
		usecases.NewCategoryUsecase,
		usecases.NewProductUsecases,
		usecases.NewBrandUsecase,
		usecases.NewDiscountUsecase,
		usecases.NewCartUseCase,
		usecases.NewWishListUseCase,
		usecases.NewAddressUsecase,
		usecases.NewOrderUsecase,
		handlers.NewUserHandler,
		handlers.NewAdminHandler,
		handlers.NewSuAdminHandler,
		handlers.NewCategoryHandler,
		handlers.NewProductHandler,
		handlers.NewBrandHandler,
		handlers.NewDiscountHandler,
		handlers.NewCartHandler,
		handlers.NewWishListHandler,
		handlers.NewAddressHandler,
		handlers.NewOrderHandler,
		routes.NewGinEngine)

	return &routes.GinEngine{}, nil

}
