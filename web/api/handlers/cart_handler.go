package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	CartUseCase usecasesinterface.CartUseCaseInterface
}

func NewCartHandler(usecase usecasesinterface.CartUseCaseInterface) *CartHandler {

	return &CartHandler{CartUseCase: usecase}

}

func (cart *CartHandler) AddToCart(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	var cartreq helperstructs.CartItemReq

	if err := c.BindJSON(&cartreq); err != nil {

		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}

	cartreq.Email = valueMap["email"].(string)

	proddata, err := cart.CartUseCase.AddToCart(c, cartreq)

	if err != nil {
		c.JSON(http.StatusNotModified, responce.Response{
			StatusCode: 304,
			Message:    "Coouldnt add to the cart",
			Data:       proddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Added to the cart successfully",
		Data:       proddata,
		Errors:     nil,
	})

}

func (cart *CartHandler) GetCartItems(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	proddata, err := cart.CartUseCase.GetCartitems(c, email)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Cannot retrive your cart items right now",
			Data:       proddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrieved all the cart iems successfully",
		Data:       proddata,
		Errors:     nil,
	})

}
