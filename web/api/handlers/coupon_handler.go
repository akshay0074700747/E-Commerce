package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	CouponUsecase usecasesinterface.CouponUsecaseInterface
}

func NewCouponHandler(usecase usecasesinterface.CouponUsecaseInterface) *CouponHandler {
	
	return &CouponHandler{CouponUsecase: usecase}

}

func (coupon *CouponHandler) AddCoupon(c *gin.Context) {
	
	var req helperstructs.CouponReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := coupon.CouponUsecase.AddCoupon(c,req); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "something went wrong with adding a new coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "coupon added successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (coupon *CouponHandler) GetAllCouponsByEmail(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	res,err := coupon.CouponUsecase.GetAllCouponsByEmail(c,email)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "something went wrong with retrieving the information",
			Data:       res,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "retrieved all the coupons available for you",
		Data:       res,
		Errors:     nil,
	})

}

func (coupon *CouponHandler) GetAllCoupons(c *gin.Context) {
	
	res,err := coupon.CouponUsecase.GetAllCoupons(c)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "something went wrong with retrieving the information",
			Data:       res,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "retrieved all the coupons successfully",
		Data:       res,
		Errors:     nil,
	})

}

func (coupon *CouponHandler) UpdateCoupon(c *gin.Context) {
	
	var req helperstructs.CouponReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := coupon.CouponUsecase.UpdateCoupon(c,req); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "something went wrong with updating the coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "coupon updated successfully",
		Data:       nil,
		Errors:     nil,
	})
	
}
func (coupon *CouponHandler) DeleteCoupon(c *gin.Context) {
	
	idstring := c.Param("id")

	id,err :=strconv.ParseUint(idstring,10,64)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "an error occcured while removing the coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := coupon.CouponUsecase.DeleteCoupon(c,uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while removing the coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "coupon removed Successfully",
		Data:       nil,
		Errors:     nil,
	})

}

