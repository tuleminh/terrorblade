package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"terrorblade/cmd/terrorbladed/internal/dtos"
	"terrorblade/cmd/terrorbladed/internal/services"
)

// NewUserHandler returns a new instance of UserHandler.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// UserHandler provides access to User HTTP handler.
type UserHandler struct {
	userService *services.UserService
}

// CreateUser inserts a new User.
func (_this *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request dtos.CreateUserRequest
		if err := c.Bind(&request); err != nil {
			return err
		}

		resp, err := _this.userService.CreateUser(&request)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, resp)
	}
}

// GetUser returns a User by ID.
func (_this *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		resp, err := _this.userService.GetUser(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, resp)
	}
}
