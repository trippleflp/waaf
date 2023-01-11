package handler

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func hash(data string) string {
	passwordHash := sha3.Sum256([]byte(data))
	return hex.EncodeToString(passwordHash[:])
}
