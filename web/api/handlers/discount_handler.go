package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	DiscountUsecase usecasesinterface.DiscountUsecaseInterface
}

func NewDiscountHandler(usecase usecasesinterface.DiscountUsecaseInterface) *DiscountHandler {

	return &DiscountHandler{DiscountUsecase: usecase}

}

func (discount *DiscountHandler) AddDiscount(c *gin.Context) {

	var req helperstructs.DiscountReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	discountdata,err := discount.DiscountUsecase.AddDiscount(c,req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while Adding a discount",
			Data:       discountdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Discount Added Successfully",
		Data:       discountdata,
		Errors:     nil,
	})

}

func (discount *DiscountHandler) UpdateDiscount(c *gin.Context) {

	var req helperstructs.DiscountReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	discountdata,err := discount.DiscountUsecase.UpdateDiscount(c,req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while Updating the discount",
			Data:       discountdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Discount Updated Successfully",
		Data:       discountdata,
		Errors:     nil,
	})

}

func (discount *DiscountHandler) DeleteDiscount(c *gin.Context) {

	idstring := c.Param("id")

	if err := discount.DiscountUsecase.DeleteDiscount(c,idstring); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while deleting the discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Discount deleted Successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (discount *DiscountHandler) GetAllDiscounts(c *gin.Context) {

	discountdata,err := discount.DiscountUsecase.GetAllDiscounts(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while Retrieving all the discounts",
			Data:       discountdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Discounts Retrieved Successfully",
		Data:       discountdata,
		Errors:     nil,
	})

}

func (discount *DiscountHandler) GetDiscountByID(c *gin.Context) {

	idstring := c.Param("id")

	discountdata,err := discount.DiscountUsecase.GetByID(c,idstring)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while Getting the discount",
			Data:       discountdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Discount retrieved Successfully",
		Data:       discountdata,
		Errors:     nil,
	})

}
