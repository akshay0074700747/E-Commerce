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
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
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

func (cart *CartHandler) DeleteCartItem(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

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

	cartreq.Email = email

	if err := cart.CartUseCase.DeleteCartItem(c, cartreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot remove the cart item right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Removed the item from cart",
		Data:       nil,
		Errors:     nil,
	})

}

func (cart *CartHandler) UpdateCartItemQuantity(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

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

	cartreq.Email = email

	if err := cart.CartUseCase.UpdateQuantity(c, cartreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot update the quantity of cart item right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully updated the quantity of cart item",
		Data:       nil,
		Errors:     nil,
	})

}
