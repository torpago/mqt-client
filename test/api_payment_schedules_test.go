/*
Core API

Testing PaymentSchedulesAPIService

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

func Test_openapi_PaymentSchedulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PaymentSchedulesAPIService CreatePaymentSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.PaymentSchedulesAPI.CreatePaymentSchedule(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentSchedulesAPIService CreatePaymentScheduleTransition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string
		var paymentScheduleToken string

		resp, httpRes, err := apiClient.PaymentSchedulesAPI.CreatePaymentScheduleTransition(context.Background(), accountToken, paymentScheduleToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentSchedulesAPIService RetrievePaymentSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string
		var paymentScheduleToken string

		resp, httpRes, err := apiClient.PaymentSchedulesAPI.RetrievePaymentSchedule(context.Background(), accountToken, paymentScheduleToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentSchedulesAPIService RetrievePaymentScheduleTransition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string
		var paymentScheduleToken string
		var token string

		resp, httpRes, err := apiClient.PaymentSchedulesAPI.RetrievePaymentScheduleTransition(context.Background(), accountToken, paymentScheduleToken, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentSchedulesAPIService RetrievePaymentScheduleTransitions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string
		var paymentScheduleToken string

		resp, httpRes, err := apiClient.PaymentSchedulesAPI.RetrievePaymentScheduleTransitions(context.Background(), accountToken, paymentScheduleToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PaymentSchedulesAPIService RetrievePaymentSchedules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountToken string

		resp, httpRes, err := apiClient.PaymentSchedulesAPI.RetrievePaymentSchedules(context.Background(), accountToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
