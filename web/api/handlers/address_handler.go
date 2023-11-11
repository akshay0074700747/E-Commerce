package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	AddressUsecase usecasesinterface.AddessUsecaseInterface
}

func NewAddressHandler(usecase usecasesinterface.AddessUsecaseInterface) *AddressHandler {

	return &AddressHandler{AddressUsecase: usecase}

}

func (address *AddressHandler) AddAddress(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	var req helperstructs.AddressReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	req.Email = valueMap["email"].(string)

	addressdata, err := address.AddressUsecase.AddAddress(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "an error occcured while adding address",
			Data:       addressdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "addressegory created Successfully",
		Data:       addressdata,
		Errors:     nil,
	})

}

func (address *AddressHandler) DeleteAddress(c *gin.Context) {

	var req helperstructs.AddrID

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "the id is not of type uint",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := address.AddressUsecase.DeleteAddress(c, req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "Could not delete the address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Deleted the address",
		Data:       nil,
		Errors:     nil,
	})

}

func (address *AddressHandler) UpdateAddress(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	var req helperstructs.AddressReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	req.Email = valueMap["email"].(string)

	addressdata, err := address.AddressUsecase.UpdateAddress(c, req)

	if err != nil {
		c.JSON(http.StatusNotModified, responce.Response{
			StatusCode: 304,
			Message:    "Coouldnt update the address",
			Data:       addressdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "address updated successfully",
		Data:       addressdata,
		Errors:     nil,
	})

}

func (address *AddressHandler) GetAlladdress(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	addressdata, err := address.AddressUsecase.GetallAddress(c, email)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responce.Response{
			StatusCode: 503,
			Message:    "Coouldnt get all the address",
			Data:       addressdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Loaded all the address",
		Data:       addressdata,
		Errors:     nil,
	})

}
