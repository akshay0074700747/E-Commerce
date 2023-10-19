package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/api/middlewares"
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecasesinterface.UserUsecaseInterface
	config      config.Config
}

func NewUserHandler(config config.Config, usecase usecasesinterface.UserUsecaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		config:      config,
	}
}

func (cr *UserHandler) UserSignUp(c *gin.Context) {

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

	if user.Otp == "" {

		if err := middlewares.Otp_Gen(user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, responce.Response{
				StatusCode: 500,
				Message:    "error while sending otp",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

		return

	} else {

		errr, otpstat := middlewares.Otp_Verify(user.Email, user.Otp)
		if errr != nil {
			c.JSON(http.StatusInternalServerError, responce.Response{
				StatusCode: 500,
				Message:    "otp not verified",
				Data:       nil,
				Errors:     errr.Error(),
			})
			return
		}

		if !otpstat {
			c.JSON(http.StatusConflict, responce.Response{
				StatusCode: 409,
				Message:    "entered otp is not correct",
				Data:       nil,
				Errors:     "otp is not correct...",
			})
			return
		}
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

	jwt, err := jwt.GenerateJwt(user.Email, false, false, []byte(cr.config.SECRET))

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "cannot generate token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetCookie("JwtToken", jwt, 3600, "/", "localhost", false, false)

	fmt.Println(jwt)

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})

}
