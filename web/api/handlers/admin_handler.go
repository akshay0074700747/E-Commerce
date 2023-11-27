package handlers

import (
	usecasesinterface "ecommerce/internal/interfaces/usecases_interface"
	"ecommerce/web/api/middlewares/jwt"
	"ecommerce/web/config"
	"ecommerce/web/helpers"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUsecase usecasesinterface.AdminUsecaseInterface
	Config       config.Config
	ToggleCrone  chan bool
	ListenCrone  chan int
}

func NewAdminHandler(usecase usecasesinterface.AdminUsecaseInterface, config config.Config, togglecrone chan bool, listencrone chan int) *AdminHandler {

	return &AdminHandler{
		AdminUsecase: usecase,
		Config:       config,
		ToggleCrone:  togglecrone,
		ListenCrone:  listencrone,
	}

}

func (ad *AdminHandler) Login(c *gin.Context) {

	cookie, _ := c.Request.Cookie("jwtToken")

	if cookie != nil {
		c.AbortWithError(http.StatusConflict, fmt.Errorf("the admin is already logged in"))
		return
	}

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
		Message:    "admin logged in successfully",
		Data:       admindta,
		Errors:     nil,
	})

}

func (ad *AdminHandler) GetAllUsers(c *gin.Context) {

	count := c.DefaultQuery("count", "4")
	page := c.DefaultQuery("page", "1")

	usersdata, err := ad.AdminUsecase.GetUsers(c, count, page)

	if err != nil {
		c.JSON(http.StatusNoContent, responce.Response{
			StatusCode: 204,
			Message:    "couldn't get all the usesers...",
			Data:       usersdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Fetched all the users successfully",
		Data:       usersdata,
		Errors:     nil,
	})

}

func (ad *AdminHandler) ReportUser(c *gin.Context) {

	var req helperstructs.ReportReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := ad.AdminUsecase.Reportuser(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "couldn't report the user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Reported User Successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (ad *AdminHandler) GetUser(c *gin.Context) {

	email := c.Param("email")

	userdata, err := ad.AdminUsecase.GetUser(c, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't get the user",
			Data:       userdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrieved User Successfully",
		Data:       userdata,
		Errors:     nil,
	})

}

func (ad *AdminHandler) Logout(c *gin.Context) {

	c.SetCookie("jwtToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Logged out successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (ad *AdminHandler) GetAdminDashBoard(c *gin.Context) {

	dashboard, err := ad.AdminUsecase.GetAdminDashBoard(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't get the dashboard",
			Data:       dashboard,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrieved Admin DashBoard Successfully",
		Data:       dashboard,
		Errors:     nil,
	})

}

func (ad *AdminHandler) StartOrStopCron(c *gin.Context) {

	status := c.Param("status")

	intstat, err := strconv.Atoi(status)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind url param",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err, boolstaat := helpers.ToggleCroneHelper(intstat)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "already in the same state or status code doesnt exist",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	fmt.Println("dfcuyyguytgfuyuy", boolstaat)

	if boolstaat {
		ad.ListenCrone <- 1
		ad.ToggleCrone <- true
	} else {
		ad.ToggleCrone <- false
	}

	fmt.Println("okkkkkkkkkkkkkkkkkkk")

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "changed the crone job status",
		Data:       nil,
		Errors:     nil,
	})

}

func (ad *AdminHandler) SalesReport(c *gin.Context) {

	var req helperstructs.SalesReportTime

	code := c.Param("code")
	var err error

	if code == "" || code == " " {

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, responce.Response{
				StatusCode: 422,
				Message:    "can't bind",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

	} else {

		intcode, err := strconv.Atoi(code)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, responce.Response{
				StatusCode: 422,
				Message:    "can't bind url param",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

		req.Starttime, err = helpers.SalesReportHelper(intcode)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, responce.Response{
				StatusCode: 422,
				Message:    "the given code doesnt exist",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

		req.EndTime = time.Now()

	}

	salesdata, err := ad.AdminUsecase.GetSalesReport(c, req.Starttime, req.EndTime)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldnt get sales report",
			Data:       salesdata,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrieved sales report successfully",
		Data:       salesdata,
		Errors:     nil,
	})

}

func (ad *AdminHandler) UserWalletDetails(c *gin.Context) {

	count := c.DefaultQuery("count", "4")
	page := c.DefaultQuery("page", "1")

	res, err := ad.AdminUsecase.GetUsersWalletDetails(c, count, page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldnt get users wallet details",
			Data:       res,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Retrieved wallet details successfully",
		Data:       res,
		Errors:     nil,
	})

}
