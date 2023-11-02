package routes

import (
	"ecommerce/web/api/handlers"
	"ecommerce/web/api/middlewares"

	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	engine *gin.Engine
}

func NewGinEngine(userhandler *handlers.UserHandler,
	adminhandler *handlers.AdminHandler,
	suadminhandler *handlers.SuAdminHandler,
	cathandler *handlers.CategoryHandler,
	prodhandler *handlers.ProductHandler,
	brandhandler *handlers.BrandHandler,
	dischandler *handlers.DiscountHandler,
	authentication *middlewares.UserAuthentication,
	carthandler *handlers.CartHandler,
	wishlisthandler *handlers.WishListHandler,
	addresshandler *handlers.AddressHandler,
	orderhandler *handlers.OrderHandler) *GinEngine {

	engine := gin.New()

	engine.Use(gin.Logger())

	engine.POST("/user/login", userhandler.UserLogin)
	engine.POST("/user/signup", userhandler.UserSignUp)

	user := engine.Group("/user")
	user.Use(authentication.UserAuth())
	{

		product := user.Group("/products")
		product.Use(middlewares.CartAndWishListAdder())
		{

			product.GET("", prodhandler.GetProducts)
			product.GET("/:category", prodhandler.FilterByCategory)
			product.GET("/:category/:sub", prodhandler.FilterByCategoryAndSub)

		}

		cart := user.Group("/cart")
		{

			cart.GET("", carthandler.GetCartItems)
			cart.POST("/add", carthandler.AddToCart)
			cart.PATCH("/update/quantity", carthandler.UpdateCartItemQuantity)
			cart.DELETE("/delete", carthandler.DeleteCartItem)
			cart.POST("/checkout/cod", orderhandler.AddOrderCOD)

		}

		order := user.Group("/orders")
		{
			order.GET("", orderhandler.GetAllOrdersByEmail)
			order.POST("/cancel", orderhandler.CancelOrder)
			order.POST("/return", orderhandler.ReturnOrder)
		}

		wishlist := user.Group("/wishlist")
		{

			wishlist.GET("", wishlisthandler.GetWishListItems)
			wishlist.POST("/add", wishlisthandler.AddToWishList)
			wishlist.DELETE("/delete", wishlisthandler.DeleteWishListItem)
			wishlist.POST("/addtocart", wishlisthandler.AddItemtoCart)

		}

		profile := user.Group("/profile")
		{
			profile.GET("", userhandler.GetUserDetails)
			profile.PATCH("/updatedetails", userhandler.UpdateUserDetails)
			profile.PATCH("/password/change", userhandler.ChangePassword)
			profile.PATCH("/password/forgot", userhandler.ForgotPassword)
		}

		address := profile.Group("/address")
		{
			address.GET("", addresshandler.GetAlladdress)
			address.POST("/add", addresshandler.AddAddress)
			address.PATCH("/update", addresshandler.UpdateAddress)
			address.DELETE("/remove", addresshandler.DeleteAddress)
		}

		user.POST("/logout", userhandler.Logout)
		user.GET("/view/product/:id", middlewares.CartAndWishListAdder(), prodhandler.GetProductByID)
		user.POST("/report", userhandler.ReportAdmin)

	}

	engine.POST("/admin/login", adminhandler.Login)

	admin := engine.Group("/admin")
	admin.Use(authentication.UserAuth(), middlewares.AdminAuth())
	{

		admin.GET("/users", adminhandler.GetAllUsers)
		admin.POST("/report", adminhandler.ReportUser)

		categories := admin.Group("/categories")
		{

			categories.GET("", cathandler.GetAllCategories)
			categories.POST("/add", cathandler.CreateCategory)
			categories.PATCH("/update", cathandler.UpdateCategory)
			categories.DELETE("/delete/:id", cathandler.DeleteCategory)

		}
		brands := admin.Group("/brands")
		{

			brands.GET("", brandhandler.GetAllbrans)
			brands.POST("/add", brandhandler.CreateBrand)
			brands.PATCH("/update", brandhandler.UpdateBrand)
			brands.DELETE("/delete/:id", brandhandler.DeleteBrand)

		}
		products := admin.Group("/products")
		{

			products.GET("", prodhandler.GetProducts)
			products.POST("/add", prodhandler.AddProduct)
			products.PATCH("/update", prodhandler.UpdateProducts)
			products.DELETE("/delete/:id", prodhandler.DeleteProduct)
			products.PATCH("/stockupdate", prodhandler.UpdateStocks)

		}
		discounts := admin.Group("/discounts")
		{

			discounts.GET("", dischandler.GetAllDiscounts)
			discounts.POST("/add", dischandler.AddDiscount)
			discounts.PATCH("/update", dischandler.UpdateDiscount)
			discounts.DELETE("/delete/:id", dischandler.DeleteDiscount)

		}
		order := admin.Group("/orders")
		{
			order.GET("", orderhandler.GetAllOrders)
			order.POST("/cancel", orderhandler.CancelOrder)
			order.PATCH("/change/status", orderhandler.ChangeStatus)
		}
		admin.POST("/logout", adminhandler.Logout)
	}

	engine.POST("/suadmin/login", suadminhandler.Login)

	suadmin := engine.Group("/suadmin")
	suadmin.Use(authentication.UserAuth(), middlewares.SuAdminAuth())
	{

		suadmin.POST("/createadmin", suadminhandler.CreateAdmin)
		suadmin.POST("/block", suadminhandler.BlockUser)
		suadmin.GET("/users", suadminhandler.ListUsers)
		suadmin.GET("/admins", suadminhandler.ListAdmins)
		suadmin.GET("/reports", suadminhandler.ListReports)
		suadmin.GET("/reports/:email", suadminhandler.DetailedReport)
		suadmin.POST("/logout", suadminhandler.Logout)

	}

	return &GinEngine{engine: engine}

}

func (sh *GinEngine) Start() {
	sh.engine.Run(":3000")
}
