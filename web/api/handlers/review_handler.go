package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	ReviewUsecase usecasesinterface.ReviewUsecaseInterface
}

func NewReviewHandler(usecase usecasesinterface.ReviewUsecaseInterface) *ReviewHandler {
	return &ReviewHandler{ReviewUsecase: usecase}
}

func (review *ReviewHandler) CreateReview(c *gin.Context) {

	var req helperstructs.ReviewReq

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	req.ReviewedBy = valueMap["email"].(string)

	if err := review.ReviewUsecase.CreateReview(c, req); err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt add review right now please try again later",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "successfully added the review",
		Data:       nil,
		Errors:     nil,
	})

}

func (review *ReviewHandler) UpdateReview(c *gin.Context) {

	var req helperstructs.ReviewReq

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	req.ReviewedBy = valueMap["email"].(string)

	if err := review.ReviewUsecase.UpdateReview(c, req); err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt update review right now please try again later",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "successfully updated the review",
		Data:       nil,
		Errors:     nil,
	})

}

func (review *ReviewHandler) GetReviewsByID(c *gin.Context) {

	strid := c.Param("id")

	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind the given id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	reviews,err := review.ReviewUsecase.GetReviewsByID(c,uint(id))

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get the reviews for the product",
			Data:       reviews,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "successfully retrieved the reviews for the order",
		Data:       reviews,
		Errors:     nil,
	})

}

func (review *ReviewHandler) GetReviwByEmail(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	reviews,err := review.ReviewUsecase.GetReviwByEmail(c,email)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get the reviews of the user",
			Data:       reviews,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "successfully retrieved the reviews of the user",
		Data:       reviews,
		Errors:     nil,
	})

}

func (review *ReviewHandler) DeleteReview(c *gin.Context) {

	idstring := c.Param("id")

	id,err :=strconv.ParseUint(idstring,10,64)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "an error occcured while deleting the review",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := review.ReviewUsecase.DeleteReview(c,uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while deleting the review",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "review deleted Successfully",
		Data:       nil,
		Errors:     nil,
	})

}
