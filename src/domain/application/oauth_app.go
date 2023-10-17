package application

import (
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/labstack/echo/v4"
)

func AddRouteOauth(e *echo.Echo, mdl *module.Module) {
	e.GET("/login", Login())
}
