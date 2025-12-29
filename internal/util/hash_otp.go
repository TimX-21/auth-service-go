package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashWithHMAC(input, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
