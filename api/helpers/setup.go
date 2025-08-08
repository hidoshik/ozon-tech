package helpers

import (
	"errors"
	"os"
)

// получить токен https://superheroapi.com/index.html

func GetAPIToken() (string, error) {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		return "", errors.New("API_TOKEN environment variable is not set")
	}
	return apiToken, nil
}