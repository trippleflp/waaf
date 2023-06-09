package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token"
	"time"
)

func CreateToken(userId string) (string, string, error) {
	jti := uuid.NewString()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "waaf",
		Subject:   userId,
		ID:        jti,
	})
	ss, err := t.SignedString(token.SigningKey)
	return ss, jti, err
}
