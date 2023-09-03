package token

import (
	"time"

	"github.com/anotherandrey/token-rest-api/internal/app/model"
	"github.com/golang-jwt/jwt"
)

var TokenSignedKey = []byte("tokenSignedKey")

type CustomClaims struct {
	*jwt.StandardClaims
	Id       int
	Username string
	Password string
}

func CreateToken(u *model.UserModel) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	expiresAt := time.Now().Add(time.Second * 60).Unix()

	token.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		u.Id,
		u.Username,
		u.Password,
	}

	return token.SignedString(TokenSignedKey)
}

func ValidToken(tokenStr string) error {
	claims, err := ParseToken(tokenStr)

	if err != nil {
		return err
	}

	return claims.Valid()
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSignedKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*CustomClaims), nil
}
