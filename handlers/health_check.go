package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"terrorblade/dtos"
)

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

type HealthCheckHandler struct {
}

func (_this *HealthCheckHandler) Check() echo.HandlerFunc {
	return func(c echo.Context) error {
		statusCode := http.StatusOK
		return c.JSON(statusCode, dtos.Metadata{
			Code:    statusCode,
			Message: http.StatusText(statusCode),
		})
	}
}
