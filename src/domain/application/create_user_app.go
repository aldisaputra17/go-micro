package application

import (
	"net/http"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/aldisaputra17/go-micro/src/domain/payload"
	"github.com/aldisaputra17/go-micro/src/domain/service"
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/aldisaputra17/go-micro/validator"
	"github.com/labstack/echo/v4"
)

func createUser(svc *service.UserService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var req payload.CreateUserRequest

		if err = c.Bind(&req); err != nil {
			err = log.FromCtx(c.Request().Context()).NewError(err, constant.ErrFailedParseRequest.Error())
			return
		}

		if err := c.Validate(req); err != nil {
			log.FromCtx(c.Request().Context()).Error(constant.ErrValidationFailed, "error validation create user request")
			return module.ResponseData(c, http.StatusBadRequest, nil, validator.ValidationErrors(err), constant.ErrValidationFailed.Error())
		}

		data, err := svc.CreateUser(c.Request().Context(), req)
		if err != nil {
			return module.ResponseData(c, http.StatusBadRequest, nil, constant.ErrDataFailedCreated, "failed create user")
		}

		res := payload.CreateUserResponses(data)

		return module.ResponseData(c, http.StatusOK, res, nil, "success create user")
	}
}
