package echokit

import (
	"net/http"

	"github.com/aldisaputra17/go-micro/toolkit/log"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type HTTPError struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
}

type defaultErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func loggerHTTPErrorHandler(w echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		logger := log.FromCtx(c.Request().Context())
		msg := "request completed with error"

		prevCommitted := c.Response().Committed

		if !c.Response().Committed {
			// writer may commit response
			w(err, c)
		}

		if c.Response().Committed {
			if !prevCommitted {
				logErrorAndResponse(logger, msg, err, c)
			}

			return
		}

		// found error & response not yet written

		// check for echo.NewHTTPError returned from handler / controller
		var errEchoHTTP *echo.HTTPError
		if ok := errors.As(err, &errEchoHTTP); ok {
			if errEchoHTTP.Internal != nil {
				err = errEchoHTTP.Internal
			}

			errWriteResp := c.JSON(errEchoHTTP.Code, errEchoHTTP)

			if errWriteResp != nil {
				logger.Error(errWriteResp, "error writing JSON response", "path", c.Request().URL.Path)
			}

			logErrorAndResponse(logger, msg, err, c)

			return
		}

		// check for web.Validation error
		var httpErr *HTTPError
		if ok := errors.As(err, &httpErr); ok {
			errWriteResp := c.JSON(httpErr.Code, httpErr)

			if errWriteResp != nil {
				logger.Error(errWriteResp, "error writing JSON response", "path", c.Request().URL.Path)
			}

			logErrorAndResponse(logger, msg, err, c)

			return
		}

		// unhandled errors returned types from controller / handler
		resp := defaultErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}

		errWriteResp := c.JSON(resp.Code, resp)
		if errWriteResp != nil {
			logger.Error(errWriteResp, "error writing JSON response", "path", c.Request().URL.Path)
		}

		logErrorAndResponse(logger, "request completed with unhandled error. add error type inspection in your echo.HTTPErrorHandler", err, c)
	}
}

func logErrorAndResponse(logger *log.Logger, msg string, err error, c echo.Context) {
	if c.Response().Status >= http.StatusInternalServerError {
		logger.Error(err, msg,
			"path", c.Request().URL.Path,
			"status_code", c.Response().Status,
		)
	} else {
		logger.Info(msg,
			"error", err,
			"path", c.Request().URL.Path,
			"status_code", c.Response().Status,
		)
	}
}
