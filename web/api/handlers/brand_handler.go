package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	BrandUsecase usecasesinterface.BrandUsecaseInterface
}

func NewBrandHandler(usecase usecasesinterface.BrandUsecaseInterface) *BrandHandler {

	return &BrandHandler{BrandUsecase: usecase}

}

func (brand *BrandHandler) CreateBrand(c *gin.Context) {

	var req helperstructs.BrandReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	branddata, err := brand.BrandUsecase.CreateBrand(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while creating brand",
			Data:       branddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "brand created Successfully",
		Data:       branddata,
		Errors:     nil,
	})

}

func (brand *BrandHandler) DeleteBrand(c *gin.Context) {

	id := c.Param("id")

	u, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "the id is not of type uint",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := brand.BrandUsecase.DeleteBrand(c, uint(u)); err != nil {
		c.JSON(http.StatusNotModified, responce.Response{
			StatusCode: 304,
			Message:    "Could not delete the brand",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Deleted the brand",
		Data:       nil,
		Errors:     nil,
	})

}

func (brand *BrandHandler) UpdateBrand(c *gin.Context) {

	var req helperstructs.BrandReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	branddata, err := brand.BrandUsecase.UpdateBrand(c, req)

	if err != nil {
		c.JSON(http.StatusNotModified, responce.Response{
			StatusCode: 304,
			Message:    "Coouldnt update the brandegory",
			Data:       branddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "brandegory updated successfully",
		Data:       branddata,
		Errors:     nil,
	})

}

func (brand *BrandHandler) GetAllbrandegories(c *gin.Context) {

	branddata, err := brand.BrandUsecase.GetallBrand(c)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get all the brand",
			Data:       branddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Loaded all the brands",
		Data:       branddata,
		Errors:     nil,
	})

}
