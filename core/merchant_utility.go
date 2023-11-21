package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func CheckMessage(secretKey string, msgBody string, sContentHmac string) bool {

	const delimiter = ";v="
	hmacPos := strings.Index(sContentHmac, delimiter)
	if hmacPos < 0 {
		return false
	}

	pos := hmacPos + len(delimiter)
	paramHmac := sContentHmac[pos:]

	key, err := hex.DecodeString(secretKey)
	if err != nil {
		return false
	}

	hmacSHA256 := hmac.New(sha256.New, key)
	hmacSHA256.Write([]byte(msgBody))

	hmacStr := hex.EncodeToString(hmacSHA256.Sum(nil))

	return hmacStr == paramHmac
}
