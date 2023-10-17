package helper

import (
	"github.com/aldisaputra17/go-micro/constant"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

// Create a JWT with signing method HS256.
func CreateJWT(claims jwt.Claims, secretKey string) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secretKey))

	return
}

// Parse JWT, and return claims object.
func ClaimsJWT(tokenStr string, secretKey string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, errors.Wrapf(err, "Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return
	}

	if jwtToken == nil || !jwtToken.Valid {
		err = errors.WithStack(constant.ErrTokenInvalid)
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.WithStack(constant.ErrTokenInvalid)
		return
	}

	return
}
