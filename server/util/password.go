package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashPassword(password, salt1, salt2 string) string {
	password = fmt.Sprintf("%v%v%v", salt1, password, salt2)
	hashedPassword := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashedPassword[:])
}
