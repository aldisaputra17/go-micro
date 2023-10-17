package application

import (
	"net/http"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/aldisaputra17/go-micro/helper"
	"github.com/aldisaputra17/go-micro/src/domain/payload"
	"github.com/aldisaputra17/go-micro/src/domain/repository/query/dto"
	"github.com/aldisaputra17/go-micro/src/domain/service"
	"github.com/aldisaputra17/go-micro/src/module"
	"github.com/aldisaputra17/go-micro/src/token"
	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/aldisaputra17/go-micro/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func LoginApp(svc *service.UserService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var req payload.LoginRequest

		if err = c.Bind(&req); err != nil {
			err = log.FromCtx(c.Request().Context()).NewError(err, constant.ErrFailedParseRequest.Error())
			return
		}

		if err := c.Validate(req); err != nil {
			log.FromCtx(c.Request().Context()).Error(constant.ErrValidationFailed, "error validation login user request")
			return module.ResponseData(c, http.StatusBadRequest, nil, validator.ValidationErrors(err), constant.ErrValidationFailed.Error())
		}

		user, err := svc.GetUser(c.Request().Context(), req.Email)
		if err != nil {
			return module.ResponseData(c, http.StatusNotFound, nil, constant.ErrDataAlreadyExists, "email not found")
		}

		if err = helper.CompareHashPassword(user.Password, req.Password); err != nil {
			return module.ResponseData(c, http.StatusBadRequest, nil, constant.ErrPasswordIncorrect, "please check your password again")
		}

		tknPayload := token.AccessTokenPayload{
			GUID:     uuid.NewString(),
			UserGUID: user.GUID,
		}

		accessToken, err := token.CreateAccessToken(tknPayload)
		if err != nil {
			return module.ResponseData(c, http.StatusBadRequest, nil, constant.ErrFailedParseRequest, "failed create access token")
		}

		refreshToken, err := token.CreateRefreshToken(token.RefreshTokenPayload(tknPayload))
		if err != nil {
			return module.ResponseData(c, http.StatusBadRequest, nil, constant.ErrFailedParseRequest, "failed create refresh token")
		}

		session := dto.CreateSessionParam{
			GUID:         uuid.NewString(),
			UserGUID:     user.GUID,
			RefreshToken: refreshToken.Token,
			UserAgent:    c.Request().UserAgent(),
			ClientIP:     c.Request().RemoteAddr,
			ExpiredAt:    refreshToken.ExpiredAt,
		}

		data, err := svc.CreateSession(c.Request().Context(), session)
		if err != nil {
			return module.ResponseData(c, http.StatusBadRequest, nil, constant.ErrFailedLogin, "failed session login")
		}

		userResponse := payload.ToReadUserResponse(user)

		sessionResponse := payload.LoginResponse{
			GUID:                  data.GUID,
			AccessToken:           accessToken.Token,
			AccessTokenExpiredAt:  accessToken.ExpiredAt,
			RefreshToken:          refreshToken.Token,
			RefreshTokenExpiresAt: refreshToken.ExpiredAt,
			CreatedAt:             data.CreatedAt,
			User:                  userResponse,
		}

		return module.ResponseData(c, http.StatusOK, sessionResponse, nil, "successfuly login")
	}
}
