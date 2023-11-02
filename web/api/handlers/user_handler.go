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
	UserUseCase     usecasesinterface.UserUsecaseInterface
	Config          config.Config
	CartUseCase     usecasesinterface.CartUseCaseInterface
	WishListUseCase usecasesinterface.WishListUseCaseInterface
	Addressusecase  usecasesinterface.AddessUsecaseInterface
	OrderUsecase    usecasesinterface.OrderUsecaseInterface
}

func NewUserHandler(config config.Config, usecase usecasesinterface.UserUsecaseInterface, cartusecase usecasesinterface.CartUseCaseInterface, wishlistusecase usecasesinterface.WishListUseCaseInterface, addressusecse usecasesinterface.AddessUsecaseInterface, orderusecse usecasesinterface.OrderUsecaseInterface) *UserHandler {
	return &UserHandler{
		UserUseCase:     usecase,
		Config:          config,
		CartUseCase:     cartusecase,
		WishListUseCase: wishlistusecase,
		Addressusecase:  addressusecse,
		OrderUsecase:    orderusecse,
	}
}

func (cr *UserHandler) UserSignUp(c *gin.Context) {

	cookie, _ := c.Request.Cookie("jwtToken")

	if cookie != nil {
		c.AbortWithError(http.StatusConflict, fmt.Errorf("the user is already logged in"))
		return
	}

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

	userData, err := cr.UserUseCase.UserSignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, responce.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       responce.UserData{},
			Errors:     err.Error(),
		})
		return
	}

	jwt, err := jwt.GenerateJwt(user.Email, false, false, []byte(cr.Config.SECRET))

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "cannot generate token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := cr.CartUseCase.CreateCart(c, user.Email); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "cannot create cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := cr.WishListUseCase.CreateWishList(c, user.Email); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "cannot create wishlist",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetCookie("jwtToken", jwt, 3600, "/", "localhost", false, false)

	fmt.Println(jwt)

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})

}

func (cr *UserHandler) UserLogin(c *gin.Context) {

	cookie, _ := c.Request.Cookie("jwtToken")

	if cookie != nil {
		c.AbortWithError(http.StatusConflict, fmt.Errorf("the user is already logged in"))
		return
	}

	var req helperstructs.UserReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userdta, err := cr.UserUseCase.UserLogin(c.Request.Context(), req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responce.Response{
			StatusCode: 401,
			Message:    "couldn't login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	jwt, err := jwt.GenerateJwt(userdta.Email, false, false, []byte(cr.Config.SECRET))

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
		Message:    "userlogged in successfully",
		Data:       userdta,
		Errors:     nil,
	})
}

func (cr *UserHandler) Logout(c *gin.Context) {

	c.SetCookie("jwtToken", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Logged out successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *UserHandler) ReportAdmin(c *gin.Context) {

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

	if err := cr.UserUseCase.ReportAdmin(c.Request.Context(), req); err != nil {
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
		Message:    "Reported Admin Successfully",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *UserHandler) GetUserDetails(c *gin.Context) {

	var userresponce responce.UserResponce

	var user helperstructs.UserReq

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	user.Email = email

	userdata, err := cr.UserUseCase.GetByEmail(c, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't get user details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	addressdata, err := cr.Addressusecase.GetallAddress(c, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't get user details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	orderdata, err := cr.OrderUsecase.GetAllOrdersByEmail(c, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't get user details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userresponce.Userdata = userdata
	userresponce.Address = addressdata
	userresponce.Orders = orderdata

	c.JSON(http.StatusOK, responce.Response{
		StatusCode: 200,
		Message:    "Successfully fetched user details",
		Data:       userresponce,
		Errors:     nil,
	})

}

func (cr *UserHandler) UpdateUserDetails(c *gin.Context) {

	var userreq helperstructs.UserReq

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	if err := c.BindJSON(&userreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userreq.Email = email

	if err := cr.UserUseCase.UpdateUserData(c, userreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't update user details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Successfully modified user details",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *UserHandler) ChangePassword(c *gin.Context) {

	var userreq helperstructs.UserReq

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	if err := c.BindJSON(&userreq); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userreq.Email = email

	if err := cr.UserUseCase.ChangePassword(c, userreq); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "couldn't update user details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Successfully changed user password",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *UserHandler) ForgotPassword(c *gin.Context) {

	value, _ := c.Get("values")

	valueMap, _ := value.(map[string]interface{})

	email := valueMap["email"].(string)

	var user helperstructs.UserReq

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, responce.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	user.Email = email

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

	if err := cr.UserUseCase.ForgotPassword(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, responce.Response{
			StatusCode: 500,
			Message:    "error while changing password",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, responce.Response{
		StatusCode: 201,
		Message:    "Successfully changed user password",
		Data:       nil,
		Errors:     nil,
	})

}
