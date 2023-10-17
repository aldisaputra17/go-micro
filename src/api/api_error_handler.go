package api

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func errorHandler() echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		var echoError *echo.HTTPError

		// if *echo.HTTPError, let echokit middleware handles it
		if errors.As(err, &echoError) {
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			err = constant.ErrDataNotFound
		}

		appDebug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))

		statusCode := http.StatusInternalServerError
		message := err.Error()

		switch {
		case
			errors.Is(err, constant.ErrFailedParseRequest) ||
				errors.Is(err, constant.ErrValidationFailed) ||
				errors.Is(err, constant.ErrPasswordIncorrect) ||
				errors.Is(err, constant.ErrEmailAlreadyExists):
			statusCode = http.StatusBadRequest
		case
			errors.Is(err, constant.ErrHeaderTokenNotFound) ||
				errors.Is(err, constant.ErrHeaderTokenInvalid) ||
				errors.Is(err, constant.ErrTokenUnauthorized) ||
				errors.Is(err, constant.ErrTokenInvalid) ||
				errors.Is(err, constant.ErrTokenExpired):
			statusCode = http.StatusUnauthorized
		case
			errors.Is(err, constant.ErrForbiddenRole) ||
				errors.Is(err, constant.ErrForbiddenPermission):
			statusCode = http.StatusForbidden
		case
			errors.Is(err, constant.ErrAccountNotFound) ||
				errors.Is(err, constant.ErrDataNotFound):
			statusCode = http.StatusNotFound
		default:
			if !appDebug {
				message = constant.ErrUnknownSource.Error()
			}
		}

		_ = ctx.JSON(statusCode, echo.NewHTTPError(statusCode, message))
	}
}
