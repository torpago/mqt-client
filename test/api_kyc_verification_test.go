/*
Core API

Testing KYCVerificationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/torpago/mqt-client"
)

func Test_openapi_KYCVerificationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KYCVerificationAPIService GetKycBusinessBusinesstoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var businessToken string

		resp, httpRes, err := apiClient.KYCVerificationAPI.GetKycBusinessBusinesstoken(context.Background(), businessToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KYCVerificationAPIService GetKycToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.KYCVerificationAPI.GetKycToken(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KYCVerificationAPIService GetKycUserUsertoken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userToken string

		resp, httpRes, err := apiClient.KYCVerificationAPI.GetKycUserUsertoken(context.Background(), userToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KYCVerificationAPIService PostKyc", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KYCVerificationAPI.PostKyc(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
