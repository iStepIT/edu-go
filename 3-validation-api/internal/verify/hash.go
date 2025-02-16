package verify

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenHash(email string) string {
	hash := sha256.New()
	hash.Write([]byte(email))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)[:32]
}
