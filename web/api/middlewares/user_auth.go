package middlewares

import (
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, err := config.LoadConfig()

		if err != nil {
			panic("Cannot Load env files...")
		}

		secret := config.SECRET

		cookie, err := c.Request.Cookie("jwtToken")
		if err != nil {
			fmt.Println("Cookie cannot be retrieved ...")
			c.JSON(http.StatusUnauthorized, "Cookie cannot be retrieved ...")
			c.Abort()
			return
		}

		cookieString := cookie.Value

		values, err := jwt.ValidateToken(cookieString, []byte(secret)) // Replace with your validation logic

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("values", values)
	}
}
