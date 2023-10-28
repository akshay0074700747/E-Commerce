package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CartAndWishListAdder() gin.HandlerFunc {

	return func(c *gin.Context) {

		value, _ := c.Get("values")

		valueMap, _ := value.(map[string]interface{})

		email := valueMap["email"].(string)

		c.Set("userhandler", email)

	}

}
