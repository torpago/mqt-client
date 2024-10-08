/*
Core API

Testing SubstatusAPIService

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

func Test_openapi_SubstatusAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubstatusAPIService CreateSubStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SubstatusAPI.CreateSubStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubstatusAPIService ListSubStatuses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SubstatusAPI.ListSubStatuses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubstatusAPIService RetrieveSubStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var substatusToken string

		resp, httpRes, err := apiClient.SubstatusAPI.RetrieveSubStatus(context.Background(), substatusToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubstatusAPIService UpdateSubStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var substatusToken string

		resp, httpRes, err := apiClient.SubstatusAPI.UpdateSubStatus(context.Background(), substatusToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
