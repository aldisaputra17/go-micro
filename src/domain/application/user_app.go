package application

import (
	"github.com/aldisaputra17/go-micro/src/domain/service"
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/labstack/echo/v4"
)

func AddUserRoute(e *echo.Echo, m *module.Module) {
	svc := service.NewUserService(m.GetDB())

	usrRoute := e.Group("/user")

	usrRoute.POST("", createUser(svc))
}
