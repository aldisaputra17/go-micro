package application

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := Login()

		return c.JSON(http.StatusOK, url)
	}
}
