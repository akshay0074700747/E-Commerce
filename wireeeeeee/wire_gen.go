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

func InitializeAPI(config2 config.Config, togglecrone chan bool, listencrone chan int) (*routes.GinEngine, error) {
	db := database.Connect_to(config2)
	userRepo := adapters.NewUserRepository(db)
	coupons := adapters.NewCouponAdapter(db)
	userUsecaseInterface := usecases.NewUserUseCase(userRepo, coupons)
	cartRepo := adapters.NewCartAdapter(db)
	cartUseCaseInterface := usecases.NewCartUseCase(cartRepo)
	wishListRepo := adapters.NewWishListAdapter(db)
	wishListUseCaseInterface := usecases.NewWishListUseCase(wishListRepo)
	addressRepo := adapters.NewAddressAdapter(db)
	addessUsecaseInterface := usecases.NewAddressUsecase(addressRepo)
	orderRepo := adapters.NewOrderAdapter(db)
	productsRepo := adapters.NewProductDataBase(db)
	discountRepo := adapters.NewDiscountAdapter(db)
	orderUsecaseInterface := usecases.NewOrderUsecase(orderRepo, cartRepo, productsRepo, userRepo, discountRepo, coupons)
	userHandler := handlers.NewUserHandler(config2, userUsecaseInterface, cartUseCaseInterface, wishListUseCaseInterface, addessUsecaseInterface, orderUsecaseInterface)
	adminRepo := adapters.NewAdminRepository(db)
	adminUsecaseInterface := usecases.NewAdminUsecase(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminUsecaseInterface, config2, togglecrone, listencrone)
	suAdminRepo := adapters.NewSuAdminRepository(db)
	suAdminUsecaseInterface := usecases.NewSuAdminUsecase(suAdminRepo)
	suAdminHandler := handlers.NewSuAdminHandler(suAdminUsecaseInterface, config2)
	categoryRepo := adapters.NewCategoryRepository(db)
	categoryUsecaseInterface := usecases.NewCategoryUsecase(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecaseInterface)
	productUsecaseInterface := usecases.NewProductUsecases(productsRepo, cartRepo, wishListRepo)
	productHandler := handlers.NewProductHandler(productUsecaseInterface)
	brandRepo := adapters.NewBrandRepository(db)
	brandUsecaseInterface := usecases.NewBrandUsecase(brandRepo)
	brandHandler := handlers.NewBrandHandler(brandUsecaseInterface)
	discountUsecaseInterface := usecases.NewDiscountUsecase(discountRepo)
	discountHandler := handlers.NewDiscountHandler(discountUsecaseInterface)
	userAuthentication := middlewares.NewUserAuthentication(db)
	cartHandler := handlers.NewCartHandler(cartUseCaseInterface)
	wishListHandler := handlers.NewWishListHandler(wishListUseCaseInterface, cartUseCaseInterface)
	addressHandler := handlers.NewAddressHandler(addessUsecaseInterface)
	orderHandler := handlers.NewOrderHandler(orderUsecaseInterface)
	paymentRepo := adapters.NewPaymentAdapter(db)
	paymentUsecaseInterface := usecases.NewPaymentUsecase(paymentRepo)
	couponUsecaseInterface := usecases.NewCouponUsecase(coupons)
	paymentHandler := handlers.NewPaymentHandler(paymentUsecaseInterface, orderUsecaseInterface, couponUsecaseInterface, config2)
	reviewRepo := adapters.NewReviewAdapter(db)
	reviewUsecaseInterface := usecases.NewReviewUsecase(reviewRepo)
	reviewHandler := handlers.NewReviewHandler(reviewUsecaseInterface)
	couponHandler := handlers.NewCouponHandler(couponUsecaseInterface)
	ginEngine := routes.NewGinEngine(userHandler, adminHandler, suAdminHandler, categoryHandler, productHandler, brandHandler, discountHandler, userAuthentication, cartHandler, wishListHandler, addressHandler, orderHandler, paymentHandler, reviewHandler, couponHandler)
	return ginEngine, nil
}
