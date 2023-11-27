package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUsecase usecasesinterface.OrderUsecaseInterface
}

func NewOrderHandler(usecase usecasesinterface.OrderUsecaseInterface) *OrderHandler {

	return &OrderHandler{OrderUsecase: usecase}

}

func (order *OrderHandler) CheckoutCart(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	var orderreq helperstructs.OrderReq

	if err := c.BindJSON(&orderreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	orderreq.Email = email

	res, err := order.OrderUsecase.AddOrder(c, orderreq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot place an order now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully placed the order",
		Data:       res,
		Errors:     nil,
	})

}

func (order *OrderHandler) CancelOrder(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	var orderid helperstructs.OrderCancel

	if err := c.BindJSON(&orderid); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if valueMap["isadmin"].(bool) {
		var err error
		email, err = order.OrderUsecase.GetEmailByID(c, orderid.OrderID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responce.Response{
				StatusCode: 500,
				Message:    "can't get the email of the user",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if err := order.OrderUsecase.CancelOrder(c, orderid.OrderID, email); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot cancel the order now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully canceled the order",
		Data:       nil,
		Errors:     nil,
	})

}

func (order *OrderHandler) ReturnOrder(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	var orderid helperstructs.OrderCancel

	if err := c.BindJSON(&orderid); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := order.OrderUsecase.ReturnOrder(c, orderid.OrderID, email); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot return the order now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully returned the order",
		Data:       nil,
		Errors:     nil,
	})

}

func (order *OrderHandler) GetAllOrders(c *gin.Context) {

	count := c.DefaultQuery("count", "4")
	page := c.DefaultQuery("page", "1")

	orderdata, err := order.OrderUsecase.GetAllOrders(c,count,page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot return the order now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully fetched all the orders",
		Data:       orderdata,
		Errors:     nil,
	})

}

func (order *OrderHandler) GetAllOrdersByEmail(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	orderdata, err := order.OrderUsecase.GetAllOrdersByEmail(c, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot return the order now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully fetched all the orders",
		Data:       orderdata,
		Errors:     nil,
	})

}

func (order *OrderHandler) ChangeStatus(c *gin.Context) {

	var status helperstructs.OrderStatus

	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	exists, statuss := helpers.StatusCheck(status.StatusCode)

	if !exists {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "invalid status code",
			Data:       nil,
			Errors:     errors.New("the given status code doesnt exists"),
		})
		return
	}

	if err := order.OrderUsecase.ChangeStatus(c, status.OrderID, statuss); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Cannot change the status now",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "successfully changed order status",
		Data:       nil,
		Errors:     nil,
	})

}
