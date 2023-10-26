package middlewares

import (
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	"ecommerce/web/helpers/responce"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserAuthentication struct {
	DB *gorm.DB
}

func NewUserAuthentication(db *gorm.DB) *UserAuthentication {
	return &UserAuthentication{DB: db}
}

func (user *UserAuthentication) UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		config, err := config.LoadConfig()

		if err != nil {
			panic("Cannot Load env files...")
		}

		secret := config.SECRET

		cookie, err := c.Request.Cookie("jwtToken")
		if err != nil {
			fmt.Println("Cookie cannot be retrieved ...")
			c.JSON(http.StatusUnauthorized, "Cookie cannot be retrieved ...Please login")
			c.Abort()
			return
		}

		cookieString := cookie.Value

		values, err := jwt.ValidateToken(cookieString, []byte(secret))

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		var userr responce.UserData
		var admin responce.AdminData

		user.DB.Raw(`SELECT * FROM users WHERE email = $1`, values["email"].(string)).Scan(&userr)

		if userr.Email == "" {

			user.DB.Raw(`SELECT * FROM admins WHERE email = $1`, values["email"].(string)).Scan(&admin)

			if admin.Email != "" && admin.Isblocked {

				c.JSON(http.StatusServiceUnavailable, fmt.Errorf("you have been blocked"))

			}

		} else if userr.Isblocked {

			c.JSON(http.StatusServiceUnavailable, fmt.Errorf("you have been blocked"))

		}

		c.Set("values", values)
	}
}
