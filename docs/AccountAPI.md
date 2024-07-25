# \AccountAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeriodicFeeSchedules**](AccountAPI.md#GetPeriodicFeeSchedules) | **Get** /accounts/{account_token}/periodicfeeschedules | Get all active and upcoming periodic fee schedules of an account
[**RetrieveBillingCycleForAccount**](AccountAPI.md#RetrieveBillingCycleForAccount) | **Get** /accounts/{account_token}/billingcycle | Get billing cycle by account token



## GetPeriodicFeeSchedules

> PeriodicFeeSchedulePage GetPeriodicFeeSchedules(ctx, accountToken).Count(count).StartIndex(startIndex).Execute()

Get all active and upcoming periodic fee schedules of an account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/torpago/mqt-client"
)

func main() {
	accountToken := "accountToken_example" // string | account token
	count := int32(56) // int32 | Number of periodic fee schedule resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetPeriodicFeeSchedules(context.Background(), accountToken).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetPeriodicFeeSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPeriodicFeeSchedules`: PeriodicFeeSchedulePage
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetPeriodicFeeSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | account token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeriodicFeeSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of periodic fee schedule resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]

### Return type

[**PeriodicFeeSchedulePage**](PeriodicFeeSchedulePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveBillingCycleForAccount

> AccountBillingCycleResponse RetrieveBillingCycleForAccount(ctx, accountToken).Execute()

Get billing cycle by account token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/torpago/mqt-client"
)

func main() {
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve delinquency state details.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.RetrieveBillingCycleForAccount(context.Background(), accountToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RetrieveBillingCycleForAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveBillingCycleForAccount`: AccountBillingCycleResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RetrieveBillingCycleForAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve delinquency state details.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveBillingCycleForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountBillingCycleResponse**](AccountBillingCycleResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

