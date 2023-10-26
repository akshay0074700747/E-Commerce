// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireeeeeee

import (
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	"ecommerce/web/api/Routes"
	"ecommerce/web/api/handlers"
	"ecommerce/web/api/middlewares"
	"ecommerce/web/config"
	"ecommerce/web/database"
)

// Injectors from wire.go:

func InitializeAPI(config2 config.Config) (*routes.GinEngine, error) {
	db := database.Connect_to(config2)
	userRepo := adapters.NewUserRepository(db)
	userUsecaseInterface := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(config2, userUsecaseInterface)
	adminRepo := adapters.NewAdminRepository(db)
	adminUsecaseInterface := usecases.NewAdminUsecase(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminUsecaseInterface, config2)
	suAdminRepo := adapters.NewSuAdminRepository(db)
	suAdminUsecaseInterface := usecases.NewSuAdminUsecase(suAdminRepo)
	suAdminHandler := handlers.NewSuAdminHandler(suAdminUsecaseInterface, config2)
	categoryRepo := adapters.NewCategoryRepository(db)
	categoryUsecaseInterface := usecases.NewCategoryUsecase(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecaseInterface)
	productsRepo := adapters.NewProductDataBase(db)
	productUsecaseInterface := usecases.NewProductUsecases(productsRepo)
	productHandler := handlers.NewProductHandler(productUsecaseInterface)
	brandRepo := adapters.NewBrandRepository(db)
	brandUsecaseInterface := usecases.NewBrandUsecase(brandRepo)
	brandHandler := handlers.NewBrandHandler(brandUsecaseInterface)
	discountRepo := adapters.NewDiscountAdapter(db)
	discountUsecaseInterface := usecases.NewDiscountUsecase(discountRepo)
	discountHandler := handlers.NewDiscountHandler(discountUsecaseInterface)
	userAuthentication := middlewares.NewUserAuthentication(db)
	ginEngine := routes.NewGinEngine(userHandler, adminHandler, suAdminHandler, categoryHandler, productHandler, brandHandler, discountHandler, userAuthentication)
	return ginEngine, nil
}
