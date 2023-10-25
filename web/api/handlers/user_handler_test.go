package handlers_test

import (
	"bytes"
	"ecommerce/mocks"
	"ecommerce/web/api/handlers"
	"ecommerce/web/config"
	helperstructs "ecommerce/web/helpers/helper_structs"
	"ecommerce/web/helpers/responce"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserLogin(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {

		mockUserUsecase := new(mocks.UserUsecaseInterface)
		mockUserUsecase.On("UserLogin", mock.Anything, mock.Anything).Return(responce.UserData{}, nil)

		userHandler := handlers.UserHandler{
			UserUseCase: mockUserUsecase,
			Config:      config.Config{SECRET: "itsmysecetkeyyyyyyy"},
		}

		router := gin.Default()
		router.POST("/login", userHandler.UserLogin)

		body, _ := json.Marshal(helperstructs.UserReq{})

		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))

		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusAccepted, resp.Code)
		mockUserUsecase.AssertExpectations(t)

	})

	t.Run("error", func(t *testing.T) {

		mockUserUsecase := new(mocks.UserUsecaseInterface)
		mockUserUsecase.On("UserLogin", mock.Anything, mock.Anything).Return(responce.UserData{}, errors.New("error"))

		userHandler := handlers.UserHandler{
			UserUseCase: mockUserUsecase,
			Config:      config.Config{SECRET: "itsmysecetkeyyyyyyy"},
		}

		router := gin.Default()
		router.POST("/login", userHandler.UserLogin)

		body, _ := json.Marshal(helperstructs.UserReq{})

		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))

		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		mockUserUsecase.AssertExpectations(t)
		
	})
}

func TestUserSignup(t *testing.T) {

	gin.SetMode(gin.TestMode)

}
