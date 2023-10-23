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
	prodhandler *handlers.ProductHandler) *GinEngine {

	engine := gin.New()

	engine.Use(gin.Logger())

	user := engine.Group("/user")
	{
		user.POST("/signup", userhandler.UserSignUp)
		user.POST("/login", userhandler.UserLogin)
	}

	engine.POST("/admin/login", adminhandler.Login)

	admin := engine.Group("/admin")
	admin.Use(middlewares.UserAuth(), middlewares.AdminAuth())
	{
		admin.GET("/users", adminhandler.GetAllUsers)
		admin.POST("/report/:email", adminhandler.ReportUser)
		admin.GET("/categories", cathandler.GetAllCategories)
		admin.POST("/categories/add", cathandler.CreateCategory)
		admin.PATCH("/categories/update", cathandler.UpdateCategory)
		admin.DELETE("/categories/delete/:id", cathandler.DeleteCategory)
		admin.GET("/products", prodhandler.GetProducts)
		admin.POST("/products/add", prodhandler.AddProduct)
		admin.PATCH("/products/update", prodhandler.UpdateProducts)
		admin.DELETE("/products/delete/:id", prodhandler.DeleteProduct)
	}

	engine.POST("/suadmin/login", suadminhandler.Login)

	suadmin := engine.Group("/suadmin")
	suadmin.Use(middlewares.UserAuth(), middlewares.SuAdminAuth())
	{
		suadmin.POST("/createadmin", suadminhandler.CreateAdmin)
		suadmin.POST("/block/:email", suadminhandler.BlockUser)
	}

	return &GinEngine{engine: engine}

}

func (sh *GinEngine) Start() {
	sh.engine.Run(":3000")
}
