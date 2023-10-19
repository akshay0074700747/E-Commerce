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

type AdminHandler struct {
	AdminUsecase usecasesinterface.AdminUsecaseInterface
	Config       config.Config
}

func NewAdminHandler(usecase usecasesinterface.AdminUsecaseInterface, config config.Config) *AdminHandler {

	return &AdminHandler{
		AdminUsecase: usecase,
		Config:       config,
	}

}

func (ad *AdminHandler) Login(c *gin.Context) {

	var req helperstructs.AdminReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	admindta, err := ad.AdminUsecase.AdminLogin(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responce.Response{
			StatusCode: 401,
			Message:    "couldn't login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	jwt, err := jwt.GenerateJwt(admindta.Email, true, false, []byte(ad.Config.SECRET))

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
		Data:       admindta,
		Errors:     nil,
	})

}
