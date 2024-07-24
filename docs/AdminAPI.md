# \AdminAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplayFailedStatements**](AdminAPI.md#ReplayFailedStatements) | **Post** /admin/replayfailedstatements | Replays all failed statement from statement error processing table
[**ReplayFailedStatementsByShortCode**](AdminAPI.md#ReplayFailedStatementsByShortCode) | **Post** /admin/replayfailedstatements/{short_code} | Replays all failed statements by short code from statement error processing table
[**ReplaySingleFailedStatement**](AdminAPI.md#ReplaySingleFailedStatement) | **Post** /admin/{short_code}/replayfailedstatement/{account_token} | Replays single failed statement by short code  and account token from statement error processing table
[**RetryAchPayment**](AdminAPI.md#RetryAchPayment) | **Post** /admin/{short_code}/retryachpayments | Create a new ACHO ACH transfer
[**ScheduleStatements**](AdminAPI.md#ScheduleStatements) | **Post** /admin/scheduleStatements | Schedules statements for applicable accounts



## ReplayFailedStatements

> Success ReplayFailedStatements(ctx).Execute()

Replays all failed statement from statement error processing table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ReplayFailedStatements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ReplayFailedStatements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayFailedStatements`: Success
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ReplayFailedStatements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReplayFailedStatementsRequest struct via the builder pattern


### Return type

[**Success**](Success.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayFailedStatementsByShortCode

> Success ReplayFailedStatementsByShortCode(ctx, shortCode).Execute()

Replays all failed statements by short code from statement error processing table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	shortCode := "shortCode_example" // string | Short code of the program

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ReplayFailedStatementsByShortCode(context.Background(), shortCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ReplayFailedStatementsByShortCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayFailedStatementsByShortCode`: Success
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ReplayFailedStatementsByShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shortCode** | **string** | Short code of the program | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayFailedStatementsByShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success**](Success.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaySingleFailedStatement

> Success ReplaySingleFailedStatement(ctx, shortCode, accountToken).Execute()

Replays single failed statement by short code  and account token from statement error processing table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	shortCode := "shortCode_example" // string | Short code of the program
	accountToken := "accountToken_example" // string | account token of the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ReplaySingleFailedStatement(context.Background(), shortCode, accountToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ReplaySingleFailedStatement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaySingleFailedStatement`: Success
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ReplaySingleFailedStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shortCode** | **string** | Short code of the program | 
**accountToken** | **string** | account token of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaySingleFailedStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Success**](Success.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryAchPayment

> RetryAchPaymentReq RetryAchPayment(ctx, shortCode).RetryAchPaymentReq(retryAchPaymentReq).Execute()

Create a new ACHO ACH transfer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	shortCode := "shortCode_example" // string | Short code of the program
	retryAchPaymentReq := *openapiclient.NewRetryAchPaymentReq("NewAchoAchTransferToken_example", "PaymentToken_example") // RetryAchPaymentReq | Create a new ACHO ACH transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.RetryAchPayment(context.Background(), shortCode).RetryAchPaymentReq(retryAchPaymentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RetryAchPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryAchPayment`: RetryAchPaymentReq
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.RetryAchPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shortCode** | **string** | Short code of the program | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryAchPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retryAchPaymentReq** | [**RetryAchPaymentReq**](RetryAchPaymentReq.md) | Create a new ACHO ACH transfer | 

### Return type

[**RetryAchPaymentReq**](RetryAchPaymentReq.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleStatements

> Success ScheduleStatements(ctx).Execute()

Schedules statements for applicable accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ScheduleStatements(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ScheduleStatements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleStatements`: Success
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ScheduleStatements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleStatementsRequest struct via the builder pattern


### Return type

[**Success**](Success.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

