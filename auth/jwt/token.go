package jwt_auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Object any
	Params map[string]any
	jwt.StandardClaims
}

func NewClaims() *Claims {
	return new(Claims)
}

func (claims *Claims) InvalidateToken(secretKey string) (string, error) {
	claims.Object = nil
	claims.Params = make(map[string]any)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(-time.Hour * 24).Unix(),
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, er := claim.SignedString([]byte(secretKey))
	if er != nil {
		return "", er
	}
	return token, nil
}

func (claims *Claims) GenerateJwtToken(secretKey string, duration ...int64) (token string, err error) {

	if len(duration) == 0 {
		duration = append(duration, time.Now().Add(time.Hour*24).Unix())
	}

	timeDuration := duration[0]

	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: timeDuration}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = claim.SignedString([]byte(secretKey))

	return
}

func (claims *Claims) ValidateJwtToken(secretKey string, tokenString string) (err error) {
	token, er := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if er != nil || token.Claims.Valid() != nil {
		err = ErrJwtTokenNotValid
	}
	return
}

func (claims *Claims) GetJwtTokenClaims(secretKey string, tokenString string) (err error) {
	token, er := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if er != nil || token.Claims.Valid() != nil {
		err = ErrJwtTokenNotValid
		return
	}
	claims = token.Claims.(*Claims)
	return
}
