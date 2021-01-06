package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacSha256(secret string, message string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sha), nil
}
