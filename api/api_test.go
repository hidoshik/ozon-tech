package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// получить токен https://superheroapi.com/index.html

type SuperHeroAPISuite struct {
	suite.Suite
	apiToken string
}

func (s *SuperHeroAPISuite) BeforeAll(t provider.T) {
	s.apiToken = os.Getenv("API_TOKEN")
	require.NotEmpty(t, s.apiToken, "API_TOKEN must be set")
}

func (s *SuperHeroAPISuite) makeRequest(name string) (*http.Response, error) {
	return http.Get("https://superheroapi.com/api/" + s.apiToken + "/search/" + name)
}

func (s *SuperHeroAPISuite) TestShouldHaveStatusCode200(t provider.T) {
	t.WithNewStep("Make and validate request", func(step provider.StepCtx) {
		resp, err := s.makeRequest(NAME)
		require.NoError(t, err, "Request failed")
		defer resp.Body.Close()

		step.WithNewStep("Check status code", func(sCtx provider.StepCtx) {
			got := resp.StatusCode
			want := http.StatusOK
			assert.Equal(t, want, got, "Unexpected status code")
		})
	})
}

func (s *SuperHeroAPISuite) TestShouldBeSuccessful(t provider.T) {
	t.WithNewStep("Make and validate request", func(step provider.StepCtx) {
		resp, err := s.makeRequest(NAME)
		require.NoError(t, err, "Request failed")
		defer resp.Body.Close()

		step.WithNewStep("Parse and validate response", func(sCtx provider.StepCtx) {
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err, "Failed to read response body")
			
			var data APIResponse
			err = json.Unmarshal(body, &data)
			assert.NoError(t, err, "JSON decode error")

			assert.Equal(t, NAME, data.ResultsFor)
			assert.Equal(t, "success", data.Response)
		})
	})
}

func (s *SuperHeroAPISuite) TestShouldGiveErrorForInvalidName(t provider.T) {
	t.WithNewStep("Make request with invalid name", func(step provider.StepCtx) {
		resp, err := s.makeRequest(INVALID_NAME)
		require.NoError(t, err, "Request failed")
		defer resp.Body.Close()

		step.WithNewStep("Parse and validate error response", func(sCtx provider.StepCtx) {
			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err, "Failed to read response body")
			
			var data APIResponseError
			err = json.Unmarshal(body, &data)
			assert.NoError(t, err, "JSON decode error")

			assert.Equal(t, "character with given name not found", data.Error)
			assert.Equal(t, "error", data.Response)
		})
	})
}

func TestSuperHeroAPI(t *testing.T) {
	suite.RunSuite(t, new(SuperHeroAPISuite))
}
