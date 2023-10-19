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
	suadminhandler *handlers.SuAdminHandler) *GinEngine {

	engine := gin.New()

	engine.Use(gin.Logger())

	user := engine.Group("/user")
	{
		user.POST("/signup", userhandler.UserSignUp)
	}

	admin := engine.Group("/admin")
	{
		admin.POST("/login", adminhandler.Login)
	}

	suadmin := engine.Group("/suadmin")
	{
		suadmin.POST("/login", suadminhandler.Login)
		suadmin.POST("/createadmin", middlewares.UserAuth(), suadminhandler.CreateAdmin)
	}

	return &GinEngine{engine: engine}

}

func (sh *GinEngine) Start() {
	sh.engine.Run(":3000")
}
