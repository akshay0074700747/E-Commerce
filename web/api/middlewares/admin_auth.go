package middlewares

import (
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("values")

		if !exists {
			c.JSON(http.StatusInternalServerError, responce.Response{
				StatusCode: 500,
				Message:    "the cookie cannot be accessed",
				Data:       nil,
				Errors:     "the cookie couldnt be accessed",
			})
			return
		}

		if valueMap, ok := value.(map[string]interface{}); ok {

			if !valueMap["isadmin"].(bool) {
				c.JSON(http.StatusUnauthorized, responce.Response{
					StatusCode: 401,
					Message:    "You have no admin privilages to do this operation",
					Data:       nil,
					Errors:     "Only admins are authorized to this route",
				})
				c.Abort()
				return
			}

		}
	}
}
