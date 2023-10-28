package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
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

	cookie, _ := c.Request.Cookie("jwtToken")

	if cookie != nil {
		c.AbortWithError(http.StatusConflict, fmt.Errorf("the super admin is already logged in"))
		return
	}

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

func (su *SuAdminHandler) BlockUser(c *gin.Context) {

	var blockreq helperstructs.BlockReq

	if err := c.BindJSON(&blockreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := su.SuAdminUsecase.BlockUser(c.Request.Context(), blockreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't block the user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Blocked User Successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (su *SuAdminHandler) ListUsers(c *gin.Context) {

	usersdata, err := su.SuAdminUsecase.GetAllUsers(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't Retrive all the users",
			Data:       usersdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrived all the Users Successfully",
		Data:       usersdata,
		Errors:     nil,
	})

}

func (su *SuAdminHandler) ListAdmins(c *gin.Context) {

	admindata, err := su.SuAdminUsecase.GetAllAdmins(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't Retrive all the admins",
			Data:       admindata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrived all the Admins Successfully",
		Data:       admindata,
		Errors:     nil,
	})

}

func (su *SuAdminHandler) ListReports(c *gin.Context) {

	reports, err := su.SuAdminUsecase.GetReportes(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't Retrive all the reports",
			Data:       reports,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrived all the Reports Successfully",
		Data:       reports,
		Errors:     nil,
	})

}

func (su *SuAdminHandler) DetailedReport(c *gin.Context) {

	email := c.Param("email")

	report, err := su.SuAdminUsecase.GetDetailedReport(c, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't Retrive the report",
			Data:       report,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrived the Report Successfully",
		Data:       report,
		Errors:     nil,
	})

}

func (su *SuAdminHandler) Logout(c *gin.Context) {

	c.SetCookie("jwtToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Logged out successfully",
		Data:       nil,
		Errors:     nil,
	})

}
