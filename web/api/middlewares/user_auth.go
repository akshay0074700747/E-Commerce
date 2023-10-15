package middlewares

import (
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		envs,err := config.LoadEnv("SECRET")

		if err != nil {
			panic("Cannot Load env files...")
		}

		secret := envs["SECRET"]

		cookie, err := c.Cookie("jwtToken")
		if err != nil {
			fmt.Println("Cookie cannot be retrieved ...")
			return err
		}

		cookieString := cookie.Value

		values, err := jwt.ValidateToken(cookieString, []byte(secret)) // Replace with your validation logic

		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		// Store the 'values' in context
		c.Set("values", values)

		return next(c)
	}
}
