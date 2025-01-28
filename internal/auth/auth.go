package auth

import (
	"errors"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
//
//	Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}

func GetAPIKeyFromCookies(cookies []*http.Cookie) (string, error) {
	for _, cookie := range cookies {
		if cookie.Name == "ApiKey" {
			return cookie.Value, nil
		}
	}
	return "", errors.New("no ApiKey Found")
}

func HashPassword(passwordString string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
