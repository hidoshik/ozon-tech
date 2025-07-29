package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// получить токен https://superheroapi.com/index.html
var (
    apiToken = os.Getenv("API_TOKEN")
)

func TestSuperHeroAPI_ByName(t *testing.T) {
	makeRequest := func(name string) (*http.Response, error) {
		return http.Get("https://superheroapi.com/api/" + apiToken + "/search/" + name)
	}

	// синтаксис используемый в learn Go with tests
	t.Run("should have status code of 200", func(t *testing.T) {
		resp, err := makeRequest(NAME)

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		defer resp.Body.Close()

		got := resp.StatusCode
		want := http.StatusOK

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	// далее использую assert для читаемости
	t.Run("should be successful", func(t *testing.T) {
		resp, err := makeRequest(NAME)
		require.NoError(t, err, "Request failed")
		defer resp.Body.Close()
		
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err, "Failed to read response body")
		
		var data APIResponse
		err = json.Unmarshal(body, &data)
		assert.NoError(t, err, "JSON decode error")

		assert.Equal(t, "ironman", data.ResultsFor)
		assert.Equal(t, "success", data.Response)
	})

	t.Run("should give error when using invalid superhero name", func(t *testing.T) {
		resp, err := makeRequest(INVALID_NAME)
		require.NoError(t, err, "Request failed")

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err, "Failed to read response body")
		
		var data APIResponseError
		err = json.Unmarshal(body, &data)
		assert.NoError(t, err, "JSON decode error")

		assert.Equal(t, "character with given name not found", data.Error)
		assert.Equal(t, "error", data.Response)
	})
}
