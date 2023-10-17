package token

import (
	"os"
	"time"

	"github.com/aldisaputra17/go-micro/helper"
	"github.com/golang-jwt/jwt"
)

type AccessTokenPayload struct {
	GUID     string `json:"guid"`
	UserGUID string `json:"user_guid"`
}

func CreateAccessToken(req AccessTokenPayload) (res JWTPayload, err error) {
	expiredDuration, err := time.ParseDuration(os.Getenv("AUTH_ACCESS_TOKEN_EXPIRES"))
	if err != nil {
		return
	}

	expiredAt := time.Now().Add(expiredDuration)

	claims := &jwt.MapClaims{
		"usi": req.UserGUID,
		"jti": req.GUID,
		"exp": expiredAt.Unix(),
	}

	token, err := helper.CreateJWT(claims, os.Getenv("AUTH_ACCESS_TOKEN_SECRET_KEY"))
	if err != nil {
		return
	}

	res = JWTPayload{
		Token:     token,
		ExpiredAt: expiredAt,
	}

	return

}
