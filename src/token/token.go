package token

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aldisaputra17/go-micro/constant"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type JWTPayload struct {
	Token     string
	ExpiredAt time.Time
}

func ParseHeaderToken(h http.Header) (reqToken string, err error) {
	headerDataToken := h.Get(constant.DefaultMdwHeaderToken)
	if !strings.Contains(headerDataToken, "Bearer") {
		err = echo.NewHTTPError(http.StatusUnauthorized, constant.MsgHeaderTokenNotFound).SetInternal(
			errors.Wrap(constant.ErrHeaderTokenNotFound, constant.MsgHeaderTokenNotFound),
		)

		return
	}

	splitToken := strings.Split(headerDataToken, fmt.Sprintf("%s ", constant.DefaultMdwHeaderBearer))
	if len(splitToken) > 1 {
		reqToken = splitToken[1]
	} else {
		err = echo.NewHTTPError(http.StatusUnauthorized, constant.MsgHeaderTokenInvalid).SetInternal(
			errors.Wrap(constant.ErrHeaderTokenInvalid, constant.MsgHeaderTokenInvalid),
		)

		return
	}

	return
}
