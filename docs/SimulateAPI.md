# \SimulateAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSimulateAuthorization**](SimulateAPI.md#PostSimulateAuthorization) | **Post** /simulate/authorization | Simulates an authorization
[**PostSimulateAuthorizationAdvice**](SimulateAPI.md#PostSimulateAuthorizationAdvice) | **Post** /simulate/authorization/advice | Simulates an authorization advice transaction
[**PostSimulateClearing**](SimulateAPI.md#PostSimulateClearing) | **Post** /simulate/clearing | Simulates a clearing/settlement transaction
[**PostSimulateDirectdeposits**](SimulateAPI.md#PostSimulateDirectdeposits) | **Post** /simulate/directdeposits | Simulates the creation of direct deposit
[**PostSimulateFinancial**](SimulateAPI.md#PostSimulateFinancial) | **Post** /simulate/financial | Simulates a financial request (PIN debit) transaction with optional cash back
[**PostSimulateFinancialAdvice**](SimulateAPI.md#PostSimulateFinancialAdvice) | **Post** /simulate/financial/advice | Simulates a financial advice transaction
[**PostSimulateFinancialBalanceinquiry**](SimulateAPI.md#PostSimulateFinancialBalanceinquiry) | **Post** /simulate/financial/balanceinquiry | Simulates a balance inquiry
[**PostSimulateFinancialOriginalcredit**](SimulateAPI.md#PostSimulateFinancialOriginalcredit) | **Post** /simulate/financial/originalcredit | Simulates an orignal credit transaction
[**PostSimulateFinancialWithdrawal**](SimulateAPI.md#PostSimulateFinancialWithdrawal) | **Post** /simulate/financial/withdrawal | Simulates an ATM withdrawal transaction
[**PostSimulateReversal**](SimulateAPI.md#PostSimulateReversal) | **Post** /simulate/reversal | Simulates a reversal transaction



## PostSimulateAuthorization

> SimulationResponseModel PostSimulateAuthorization(ctx).AuthRequestModel(authRequestModel).Execute()

Simulates an authorization

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
	authRequestModel := *openapiclient.NewAuthRequestModel(decimal.Decimal(123), "CardToken_example", "Mid_example") // AuthRequestModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateAuthorization(context.Background()).AuthRequestModel(authRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateAuthorization`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRequestModel** | [**AuthRequestModel**](AuthRequestModel.md) |  | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateAuthorizationAdvice

> SimulationResponseModel PostSimulateAuthorizationAdvice(ctx).AuthorizationAdviceModel(authorizationAdviceModel).Execute()

Simulates an authorization advice transaction

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
	authorizationAdviceModel := *openapiclient.NewAuthorizationAdviceModel(decimal.Decimal(123), "OriginalTransactionToken_example") // AuthorizationAdviceModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateAuthorizationAdvice(context.Background()).AuthorizationAdviceModel(authorizationAdviceModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateAuthorizationAdvice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateAuthorizationAdvice`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateAuthorizationAdvice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateAuthorizationAdviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationAdviceModel** | [**AuthorizationAdviceModel**](AuthorizationAdviceModel.md) |  | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateClearing

> SimulationResponseModel PostSimulateClearing(ctx).ClearingModel(clearingModel).Execute()

Simulates a clearing/settlement transaction

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
	clearingModel := *openapiclient.NewClearingModel(decimal.Decimal(123), "OriginalTransactionToken_example") // ClearingModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateClearing(context.Background()).ClearingModel(clearingModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateClearing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateClearing`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateClearing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateClearingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clearingModel** | [**ClearingModel**](ClearingModel.md) |  | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateDirectdeposits

> DepositDepositResponse PostSimulateDirectdeposits(ctx).DirectDepositRequest(directDepositRequest).Execute()

Simulates the creation of direct deposit

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/torpago/mqt-client"
)

func main() {
	directDepositRequest := *openapiclient.NewDirectDepositRequest("AccountNumber_example", decimal.Decimal(123), time.Now(), "Type_example") // DirectDepositRequest | Direct deposit simulate request model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateDirectdeposits(context.Background()).DirectDepositRequest(directDepositRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateDirectdeposits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateDirectdeposits`: DepositDepositResponse
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateDirectdeposits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateDirectdepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directDepositRequest** | [**DirectDepositRequest**](DirectDepositRequest.md) | Direct deposit simulate request model | 

### Return type

[**DepositDepositResponse**](DepositDepositResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateFinancial

> SimulationResponseModel PostSimulateFinancial(ctx).FinancialRequestModel(financialRequestModel).Execute()

Simulates a financial request (PIN debit) transaction with optional cash back

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
	financialRequestModel := *openapiclient.NewFinancialRequestModel(decimal.Decimal(123), *openapiclient.NewCardAcceptorModel(), "CardToken_example", "Mid_example") // FinancialRequestModel | Financial request model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateFinancial(context.Background()).FinancialRequestModel(financialRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateFinancial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateFinancial`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateFinancial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateFinancialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **financialRequestModel** | [**FinancialRequestModel**](FinancialRequestModel.md) | Financial request model | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateFinancialAdvice

> SimulationResponseModel PostSimulateFinancialAdvice(ctx).AuthorizationAdviceModel(authorizationAdviceModel).Execute()

Simulates a financial advice transaction

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
	authorizationAdviceModel := *openapiclient.NewAuthorizationAdviceModel(decimal.Decimal(123), "OriginalTransactionToken_example") // AuthorizationAdviceModel | Financial advice request model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateFinancialAdvice(context.Background()).AuthorizationAdviceModel(authorizationAdviceModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateFinancialAdvice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateFinancialAdvice`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateFinancialAdvice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateFinancialAdviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationAdviceModel** | [**AuthorizationAdviceModel**](AuthorizationAdviceModel.md) | Financial advice request model | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateFinancialBalanceinquiry

> SimulationResponseModel PostSimulateFinancialBalanceinquiry(ctx).BalanceInquiryRequestModel(balanceInquiryRequestModel).Execute()

Simulates a balance inquiry

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
	balanceInquiryRequestModel := *openapiclient.NewBalanceInquiryRequestModel("AccountType_example", *openapiclient.NewCardAcceptorModel(), "CardToken_example", "Mid_example") // BalanceInquiryRequestModel | Balance inquiry request model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateFinancialBalanceinquiry(context.Background()).BalanceInquiryRequestModel(balanceInquiryRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateFinancialBalanceinquiry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateFinancialBalanceinquiry`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateFinancialBalanceinquiry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateFinancialBalanceinquiryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **balanceInquiryRequestModel** | [**BalanceInquiryRequestModel**](BalanceInquiryRequestModel.md) | Balance inquiry request model | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateFinancialOriginalcredit

> SimulationResponseModel PostSimulateFinancialOriginalcredit(ctx).OrignalcreditRequestModel(orignalcreditRequestModel).Execute()

Simulates an orignal credit transaction

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
	orignalcreditRequestModel := *openapiclient.NewOrignalcreditRequestModel(decimal.Decimal(123), "CardToken_example", "Mid_example", "Type_example") // OrignalcreditRequestModel | Orignal Credit request model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateFinancialOriginalcredit(context.Background()).OrignalcreditRequestModel(orignalcreditRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateFinancialOriginalcredit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateFinancialOriginalcredit`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateFinancialOriginalcredit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateFinancialOriginalcreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orignalcreditRequestModel** | [**OrignalcreditRequestModel**](OrignalcreditRequestModel.md) | Orignal Credit request model | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateFinancialWithdrawal

> SimulationResponseModel PostSimulateFinancialWithdrawal(ctx).WithdrawalRequestModel(withdrawalRequestModel).Execute()

Simulates an ATM withdrawal transaction

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
	withdrawalRequestModel := *openapiclient.NewWithdrawalRequestModel(decimal.Decimal(123), "CardToken_example", "Mid_example") // WithdrawalRequestModel | ATM withdrawal request model

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateFinancialWithdrawal(context.Background()).WithdrawalRequestModel(withdrawalRequestModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateFinancialWithdrawal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateFinancialWithdrawal`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateFinancialWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateFinancialWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withdrawalRequestModel** | [**WithdrawalRequestModel**](WithdrawalRequestModel.md) | ATM withdrawal request model | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimulateReversal

> SimulationResponseModel PostSimulateReversal(ctx).ReversalModel(reversalModel).Execute()

Simulates a reversal transaction

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
	reversalModel := *openapiclient.NewReversalModel(decimal.Decimal(123), "OriginalTransactionToken_example") // ReversalModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimulateAPI.PostSimulateReversal(context.Background()).ReversalModel(reversalModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimulateAPI.PostSimulateReversal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimulateReversal`: SimulationResponseModel
	fmt.Fprintf(os.Stdout, "Response from `SimulateAPI.PostSimulateReversal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSimulateReversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reversalModel** | [**ReversalModel**](ReversalModel.md) |  | 

### Return type

[**SimulationResponseModel**](SimulationResponseModel.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

