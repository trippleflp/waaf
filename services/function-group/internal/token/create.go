package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token"
	"time"
)

type MyCustomClaims struct {
	TempTokens []string `json:"tempTokens"`
	jwt.RegisteredClaims
}

func CreateToken(userId string, hashes []string) (string, string, error) {
	jti := uuid.NewString()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		hashes,
		jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "waaf",
			ID:        jti,
		},
	})
	ss, err := t.SignedString(token.SigningKey)
	return ss, jti, err
}
