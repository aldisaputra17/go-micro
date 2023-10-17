package token

import (
	"os"
	"time"

	"github.com/aldisaputra17/go-micro/helper"
	"github.com/golang-jwt/jwt"
)

type RefreshTokenPayload struct {
	GUID     string `json:"guid"`
	UserGUID string `json:"user_guid"`
}

func CreateRefreshToken(request RefreshTokenPayload) (response JWTPayload, err error) {
	expiredDuration, err := time.ParseDuration(os.Getenv("AUTH_REFRESH_TOKEN_EXPIRES"))
	if err != nil {
		return
	}

	expiredAt := time.Now().Add(expiredDuration)

	claims := &jwt.MapClaims{
		"jti": request.GUID,
		"usi": request.UserGUID,
		"exp": expiredAt.Unix(),
	}

	token, err := helper.CreateJWT(claims, os.Getenv("AUTH_REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		return
	}

	response = JWTPayload{
		Token:     token,
		ExpiredAt: expiredAt,
	}

	return
}

// Parse JWT, and return claims object.
func ClaimsRefreshToken(token string) (response RefreshTokenPayload, err error) {
	claims, err := helper.ClaimsJWT(token, os.Getenv("AUTH_REFRESH_TOKEN_SECRET_KEY"))
	if err != nil {
		return
	}

	response = RefreshTokenPayload{
		GUID:     claims["jti"].(string),
		UserGUID: claims["usi"].(string),
	}

	return
}
