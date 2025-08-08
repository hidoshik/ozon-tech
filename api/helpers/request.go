package helpers

import "net/http"

func MakeRequest(name, apiToken string) (*http.Response, error) {
	return http.Get("https://superheroapi.com/api/" + apiToken + "/search/" + name)
}