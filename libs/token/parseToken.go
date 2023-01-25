package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func ParseToken(tokenString string) (*string, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		subject := claims["sub"].(string)
		return &subject, nil
	}
	return nil, fmt.Errorf("could not parse token claims")
}
