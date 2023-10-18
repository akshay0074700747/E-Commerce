package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/api/middlewares"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecasesinterface.UserUsecaseInterface
}

func NewUserHandler(usecase usecasesinterface.UserUsecaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) UserSignUp(c *gin.Context) {

	buffer := c.Copy()

	var user helperstructs.UserReq
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	errr, otpstat := middlewares.Otp_Gen(buffer)

	if errr != nil {
		c.JSON(http.StatusConflict, errr.Error())
		return
	}

	if !otpstat {
		c.JSON(http.StatusNotFound, gin.H{"error": "otp is not verified"})
		return
	}

	userData, err := cr.userUseCase.UserSignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       responce.UserData{},
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})

}
