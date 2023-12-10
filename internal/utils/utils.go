package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

// Doing base64 Encoding For Base64
func EncodeingId(id string) string {
	hashInstance := sha256.New()
	hashInstance.Write([]byte(id))
	hashedBytes := hashInstance.Sum(nil)
	encodedString := base64.URLEncoding.EncodeToString(hashedBytes)

	return encodedString[:8]
}
