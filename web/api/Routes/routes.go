package routes

import (
	"ecommerce/web/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoEngine struct {
	engine *echo.Echo
}

func NewEchoEngine(handlers.UserHandler) *EchoEngine {

	e := echo.New()

	e.Use(middleware.Logger())

	user := e.Group("/user")
	{
		user.POST("/signup",)
	}

	return &EchoEngine{engine: e}

}

func (en *EchoEngine) Start() {
	en.engine.Start(":3000")
}
