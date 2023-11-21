package core

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

func checkAuthHash(values url.Values, merchantCcId string, merchantPw string) bool {

	if values == nil || !values.Has("authParams") || !values.Has("vAuthInfo") {
		return false
	}

	authParams := values.Get("authParams")
	vAuthInfo := values.Get("vAuthInfo")

	if authParams == "" || vAuthInfo == "" {
		return false
	}

	hash, err := createAuthHashInfo(values, authParams, merchantCcId, merchantPw)
	if err != nil {
		return false
	}

	return vAuthInfo == hash

}

func createAuthHashInfo(values url.Values, authParams string, merchantCcId string, merchantPw string) (string, error) {

	buffer, err := base64.StdEncoding.DecodeString(authParams)

	if err != nil {
		return "", err
	}

	decoded := string(buffer)

	params := strings.Split(decoded, ",")

	str := &strings.Builder{}

	str.WriteString(merchantCcId)
	for _, param := range params {
		if values.Has(param) {
			str.WriteString(values.Get(param))
		}
	}
	str.WriteString(merchantPw)

	hash := sha256.Sum256([]byte(str.String()))

	return fmt.Sprintf("%x", hash), nil

}
