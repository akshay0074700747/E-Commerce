package handlers

import (
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	"ecommerce/web/database"
	"ecommerce/web/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase usecases.UserUsecase
}

func NewUserHandler(userusecase usecases.UserUsecase) *UserHandler {
	return &UserHandler{UserUseCase: userusecase}
}

func (userhandler *UserHandler) RegisterUser(c echo.Context, dbs database.ItsaDatabase) error {
	var jsondta map[string]interface{}

	if err := c.Bind(&jsondta); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid json data..."})
	}

	useradapter := adapters.NewUserAdapter(dbs.Returnconn())

	email, mobile, name := jsondta["email"], jsondta["mobile"], jsondta["name"]

	password := helpers.Hash_pass(jsondta["password"].(string))

	userusecase := usecases.NewUserUsecase(useradapter)

	// Assuming RegisterUser returns an error
	if err := userusecase.RegisterUser(email.(string), password, mobile.(string), name.(string)); err != nil {
		// Handle the error, e.g., return an HTTP response with the error message.
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Registration was successful; you can return a success response.
	return c.JSON(http.StatusOK, map[string]string{"message": "User registered successfully"})
}
