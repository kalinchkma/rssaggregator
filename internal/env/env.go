package env

import (
	"fmt"
	"os"
)

func GetString(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("%v not found", key)
	}
	return value, nil
}
