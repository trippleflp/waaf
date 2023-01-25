package token

import (
	"github.com/golang-jwt/jwt/v4"
)

var SigningKey = []byte("signing_secret")

func Validate(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return SigningKey, nil })
	if err != nil || !token.Valid {
		return err
	}

	return nil
}
