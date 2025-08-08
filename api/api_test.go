package api

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"api-test/helpers"
)

type SuperHeroAPISuite struct {
	suite.Suite
}

func (s *SuperHeroAPISuite) BeforeAll(t provider.T) {
	_, err := helpers.GetAPIToken();

	if err != nil {
		t.Fatalf("Failed to get API token: %v", err)
	}
	
	t.Log("API token has been set successfully")
}

func (s *SuperHeroAPISuite) TestShouldHaveStatusCode200(t provider.T) {
	t.WithNewStep("Make and validate request", func(step provider.StepCtx) {
		apiToken, _ := helpers.GetAPIToken();

		resp, err := helpers.MakeRequest(NAME, apiToken)
		t.Require().NoError(err, "Request failed")
		defer resp.Body.Close()

		//вложенный шаг
		step.WithNewStep("Check status code", func(sCtx provider.StepCtx) {
			got := resp.StatusCode
			want := http.StatusOK
			t.Assert().Equal(want, got, "Unexpected status code")
		})
	})
}

func (s *SuperHeroAPISuite) TestShouldBeSuccessful(t provider.T) {
	t.WithNewStep("Make and validate request", func(step provider.StepCtx) {
		apiToken, _ := helpers.GetAPIToken();

		resp, err := helpers.MakeRequest(NAME, apiToken)
		t.Require().NoError(err, "Request failed")
		defer resp.Body.Close()

		// вложенный шаг
		step.WithNewStep("Parse and validate response", func(sCtx provider.StepCtx) {
			body, err := io.ReadAll(resp.Body)
			t.Require().NoError(err, "Failed to read response body")
			
			var data APIResponse
			err = json.Unmarshal(body, &data)
			t.Assert().NoError(err, "JSON decode error")

			t.Assert().Equal(NAME, data.ResultsFor, "Unexpected superhero name")
			t.Assert().Equal("success", data.Response, "Unexpected response value")
		})
	})
}

func (s *SuperHeroAPISuite) TestShouldGiveErrorForInvalidName(t provider.T) {
	t.WithNewStep("Make request with invalid name", func(step provider.StepCtx) {
		apiToken, _ := helpers.GetAPIToken();

		resp, err := helpers.MakeRequest(INVALID_NAME, apiToken)
		t.Require().NoError(err, "Request failed")
		defer resp.Body.Close()

		// вложенный шаг
		step.WithNewStep("Parse and validate error response", func(sCtx provider.StepCtx) {
			body, err := io.ReadAll(resp.Body)
			t.Require().NoError(err, "Failed to read response body")
			
			var data APIResponseError
			err = json.Unmarshal(body, &data)
			t.Assert().NoError(err, "JSON decode error")

			t.Assert().Equal("character with given name not found", data.Error, "Unexpected error message")
			t.Assert().Equal("error", data.Response, "Unexpected response value")
		})
	})
}

func TestSuperHeroAPI(t *testing.T) {
	suite.RunSuite(t, new(SuperHeroAPISuite))
}
