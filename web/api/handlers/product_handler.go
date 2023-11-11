package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase usecasesinterface.ProductUsecaseInterface
}

func NewProductHandler(usecase usecasesinterface.ProductUsecaseInterface) *ProductHandler {
	return &ProductHandler{ProductUsecase: usecase}
}

func (product *ProductHandler) AddProduct(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	var productreq helperstructs.ProductReq

	if err := c.BindJSON(&productreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	productreq.UpdatedBy = valueMap["email"].(string)

	proddata, err := product.ProductUsecase.AddProduct(c, productreq)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while Adding Product",
			Data:       proddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Product created Successfully",
		Data:       proddata,
		Errors:     nil,
	})

}

func (product *ProductHandler) GetProducts(c *gin.Context) {

	count := c.DefaultQuery("count", "4")
	page := c.DefaultQuery("page", "1")

	email, exists := c.Get("userhandler")

	if !exists {
		email = ""
	}

	productdata, err := product.ProductUsecase.GetProducts(c, email.(string), count, page)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get all the Products",
			Data:       productdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Loaded all the Products",
		Data:       productdata,
		Errors:     nil,
	})

}

func (product *ProductHandler) UpdateProducts(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	var productreq helperstructs.ProductReq

	if err := c.BindJSON(&productreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	productreq.UpdatedBy = valueMap["email"].(string)

	proddata, err := product.ProductUsecase.UpdateProduct(c, productreq)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get all the Products",
			Data:       proddata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Updated the Product",
		Data:       proddata,
		Errors:     nil,
	})

}

func (product *ProductHandler) DeleteProduct(c *gin.Context) {

	del := c.Param("id")

	del_id, err := strconv.ParseUint(del, 10, 0)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't convert to uint",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err = product.ProductUsecase.DeleteProduct(c, uint(del_id)); err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt delete the product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Deleted the Product",
		Data:       nil,
		Errors:     nil,
	})

}

func (product *ProductHandler) FilterByCategory(c *gin.Context) {

	count := c.DefaultQuery("count", "4")
	page := c.DefaultQuery("page", "1")

	email, exists := c.Get("userhandler")

	if !exists {
		email = ""
	}

	category := c.Param("category")

	productdata, err := product.ProductUsecase.GetProducts(c, email.(string), count, page)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get the Products",
			Data:       productdata,
			Errors:     err.Error(),
		})
		return
	}

	for i := range productdata {

		if productdata[i].Category != category {
			productdata = append(productdata[:i], productdata[i+1:]...)
		}

	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Loaded all the Products",
		Data:       productdata,
		Errors:     nil,
	})

}

func (product *ProductHandler) FilterByCategoryAndSub(c *gin.Context) {

	count := c.DefaultQuery("count", "4")
	page := c.DefaultQuery("page", "1")

	email, exists := c.Get("userhandler")

	if !exists {
		email = ""
	}

	category := c.Param("category")
	sub := c.Param("sub")

	productdata, err := product.ProductUsecase.GetProducts(c, email.(string), count, page)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get the Products",
			Data:       productdata,
			Errors:     err.Error(),
		})
		return
	}

	var newproducts []responce.ProuctData

	for i := range productdata {

		if productdata[i].Category == category && productdata[i].SubCategory == sub {
			newproducts = append(newproducts, productdata[i])
		}

	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Loaded all the Products",
		Data:       newproducts,
		Errors:     nil,
	})

}

func (product *ProductHandler) UpdateStocks(c *gin.Context) {

	var stock helperstructs.StockReq

	if err := c.BindJSON(&stock); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't convert to uint",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := product.ProductUsecase.UpdateStock(c, stock.ID, stock.Stock); err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get the Products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Updated the stock",
		Data:       nil,
		Errors:     nil,
	})

}

func (product *ProductHandler) GetProductByID(c *gin.Context) {

	email, exists := c.Get("userhandler")

	if !exists {
		email = ""
	}

	id := c.Param("id")

	productdata, err := product.ProductUsecase.GetProductByID(c, id, email.(string))

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get the Product",
			Data:       productdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrieved the product details",
		Data:       productdata,
		Errors:     nil,
	})

}
