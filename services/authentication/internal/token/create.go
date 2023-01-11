package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var signingKey = []byte("testKey")

func CreateToken(userId string) (string, string, error) {
	jti := uuid.NewString()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "waaf",
		Subject:   userId,
		ID:        jti,
	})
	ss, err := token.SignedString(signingKey)
	return ss, jti, err
}
