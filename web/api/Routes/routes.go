package routes

import (
	"ecommerce/web/api/handlers"

	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	engine *gin.Engine
}

func NewGinEngine(userhandler *handlers.UserHandler) *GinEngine {

	engine := gin.New()

	engine.Use(gin.Logger())

	user := engine.Group("/user")
	{
		user.POST("/signup", userhandler.UserSignUp)
	}

	return &GinEngine{engine: engine}

}

func (sh *GinEngine) Start() {
	sh.engine.Run(":3000")
}
