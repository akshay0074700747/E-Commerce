package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUsecase usecasesinterface.CategoryUsecaseInterface
}

func NewCategoryHandler(usecase usecasesinterface.CategoryUsecaseInterface) *CategoryHandler {
	
	return &CategoryHandler{CategoryUsecase: usecase}

}

func (cat *CategoryHandler) CreateCategory(c *gin.Context)  {
	
	var req helperstructs.CategoryReq

	if err:= c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity,responce.Response{
			StatusCode: 422,
			Message: "can't bind",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}

	catdata, err := cat.CategoryUsecase.CreateCategory(c,req)

	if  err != nil {
		c.JSON(http.StatusInternalServerError,responce.Response{
			StatusCode: 500,
			Message: "an error occcured while creating category",
			Data: catdata,
			Errors: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Category created Successfully",
		Data:       catdata,
		Errors:     nil,
	})

}


func (cat *CategoryHandler) DeleteCategory(c *gin.Context)  {
	
	id := c.Param("productid")

	u,err := strconv.ParseUint(id,10,0)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,responce.Response{
			StatusCode: 422,
			Message: "the id is not of type uint",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}

	if err := cat.CategoryUsecase.DeleteCategory(c,uint(u)); err != nil {
		c.JSON(http.StatusNotModified,responce.Response{
			StatusCode: 304,
			Message: "Could not delete the category",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,responce.Response{
		StatusCode: 200,
		Message: "Deleted the category",
		Data: nil,
		Errors: nil,
	})

}

func (cat *CategoryHandler) UpdateCategory(c *gin.Context)  {

	var req helperstructs.CategoryReq

	if err:= c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity,responce.Response{
			StatusCode: 422,
			Message: "can't bind",
			Data: nil,
			Errors: err.Error(),
		})
		return
	}

	catdata,err := cat.CategoryUsecase.UpdateCategory(c,req)

	if err != nil {
		c.JSON(http.StatusNotModified,responce.Response{
			StatusCode: 304,
			Message: "Coouldnt update the category",
			Data: catdata,
			Errors: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,responce.Response{
		StatusCode: 201,
		Message: "Category updated successfully",
		Data: catdata,
		Errors: nil,
	})

}

func (cat *CategoryHandler) GetAllCategories(c *gin.Context)  {
	
	catdata,err := cat.CategoryUsecase.GetAllCategories(c)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable,responce.Response{
			StatusCode: 503,
			Message: "Coouldnt get all the categories",
			Data: catdata,
			Errors: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,responce.Response{
		StatusCode: 200,
		Message: "Loaded all the categories",
		Data: catdata,
		Errors: nil,
	})

}