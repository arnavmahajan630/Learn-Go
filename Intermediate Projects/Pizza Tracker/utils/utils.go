package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortID() (string, error) {
	b := make([]byte, 11)
	if _, err := rand.Read(b); err != nil {
		return "", err 
	}
	return base64.RawURLEncoding.EncodeToString(b)[:14], nil
}