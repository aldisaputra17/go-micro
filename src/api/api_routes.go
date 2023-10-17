package api

import (
	"net/http"

	userRoute "github.com/aldisaputra17/go-micro/src/domain/application"
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, mdl *module.Module) {
	e.HTTPErrorHandler = errorHandler()

	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "success login")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	userRoute.AddUserRoute(e, mdl)
}
