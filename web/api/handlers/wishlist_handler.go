package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WishListHandler struct {
	WishListUseCase usecasesinterface.WishListUseCaseInterface
	CartUsecase     usecasesinterface.CartUseCaseInterface
}

func NewWishListHandler(wishlistusecase usecasesinterface.WishListUseCaseInterface, cartusecase usecasesinterface.CartUseCaseInterface) *WishListHandler {

	return &WishListHandler{
		WishListUseCase: wishlistusecase,
		CartUsecase:     cartusecase,
	}

}

func (wishlist *WishListHandler) AddToWishList(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	var wishlistreq helperstructs.WishListItemsReq

	if err := c.BindJSON(&wishlistreq); err != nil {

		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}

	wishlistreq.Email = valueMap["email"].(string)

	proddata, err := wishlist.WishListUseCase.AddToWishList(c, wishlistreq)

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

func (wishlist *WishListHandler) GetWishListItems(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	proddata, err := wishlist.WishListUseCase.GetWishListitems(c, email)

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

func (wishlist *WishListHandler) DeleteWishListItem(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	var wishreq helperstructs.WishListItemsReq

	if err := c.BindJSON(&wishreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	wishreq.Email = email

	if err := wishlist.WishListUseCase.DeleteWishListItem(c, wishreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot remove the wishlist item right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Removed the item from wishlist",
		Data:       nil,
		Errors:     nil,
	})

}

func (wishlist *WishListHandler) AddItemtoCart(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	var cartreq helperstructs.CartItemReq
	var wishreq helperstructs.WishListItemsReq

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
	cartreq.Quantity = 1
	wishreq.Email = email
	wishreq.ProductId = cartreq.ProductId

	proddata, err := wishlist.CartUsecase.AddToCart(c, cartreq)

	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "Coouldnt add to the cart",
			Data:       proddata,
			Errors:     err.Error(),
		})
		return
	}

	if err := wishlist.WishListUseCase.DeleteWishListItem(c, wishreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot remove the wishlist item right now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Successfully added to cart",
		Data:       proddata,
		Errors:     nil,
	})

}

func (wishlist *WishListHandler) TransferAlltoCart(c *gin.Context) {

}
