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
	authentication *middlewares.UserAuthentication) *GinEngine {

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
		user.POST("/logout", userhandler.Logout)
		user.GET("/view/product/:id", prodhandler.GetProductByID).Use(middlewares.CartAndWishListAdder())
		user.POST("/report", userhandler.ReportAdmin)
	}

	engine.POST("/admin/login", adminhandler.Login)

	admin := engine.Group("/admin")
	admin.Use(authentication.UserAuth(), middlewares.AdminAuth())
	{
		admin.GET("/users", adminhandler.GetAllUsers)
		admin.POST("/report", adminhandler.ReportUser)
		admin.GET("/categories", cathandler.GetAllCategories)
		admin.POST("/categories/add", cathandler.CreateCategory)
		admin.PATCH("/categories/update", cathandler.UpdateCategory)
		admin.DELETE("/categories/delete/:id", cathandler.DeleteCategory)
		admin.GET("/brands", brandhandler.GetAllbrans)
		admin.POST("/brands/add", brandhandler.CreateBrand)
		admin.PATCH("/brands/update", brandhandler.UpdateBrand)
		admin.DELETE("/brands/delete/:id", brandhandler.DeleteBrand)
		admin.GET("/products", prodhandler.GetProducts)
		admin.POST("/products/add", prodhandler.AddProduct)
		admin.PATCH("/products/update", prodhandler.UpdateProducts)
		admin.DELETE("/products/delete/:id", prodhandler.DeleteProduct)
		admin.PATCH("/products/stockupdate", prodhandler.UpdateStocks)
		admin.GET("/discounts", dischandler.GetAllDiscounts)
		admin.POST("discounts/add", dischandler.AddDiscount)
		admin.DELETE("/discounts/delete/:id", dischandler.DeleteDiscount)
		admin.PATCH("/discounts/update", dischandler.UpdateDiscount)
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
