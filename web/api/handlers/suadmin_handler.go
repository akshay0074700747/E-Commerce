package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuAdminHandler struct {
	SuAdminUsecase usecasesinterface.SuAdminUsecaseInterface
	Config         config.Config
}

func NewSuAdminHandler(usecase usecasesinterface.SuAdminUsecaseInterface, config config.Config) *SuAdminHandler {

	return &SuAdminHandler{
		SuAdminUsecase: usecase,
		Config:         config,
	}

}

func (su *SuAdminHandler) Login(c *gin.Context) {

	var req helperstructs.SuAdminReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	suadmindta, err := su.SuAdminUsecase.SuAdminLogin(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responce.Response{
			StatusCode: 401,
			Message:    "couldn't login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	jwt, err := jwt.GenerateJwt(suadmindta.Email, false, true, []byte(su.Config.SECRET))

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't generate jwt",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetCookie("jwtToken", jwt, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusAccepted, responce.Response{
		StatusCode: 202,
		Message:    "super admin logged in successfully",
		Data:       suadmindta,
		Errors:     nil,
	})

}

func (su *SuAdminHandler) CreateAdmin(c *gin.Context) {

	cookie, err := c.Cookie("jwtToken")

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "the cookie cannot be accessed",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	var req helperstructs.AdminReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	admindta, err := su.SuAdminUsecase.CreateAdmin(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't create admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "admin created Successfully",
		Data:       admindta,
		Errors:     nil,
	})

}
