# \VelocityControlsCardGroupBalanceAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVelocityControlsCardGroupBalances**](VelocityControlsCardGroupBalanceAPI.md#ListVelocityControlsCardGroupBalances) | **Get** /velocitycontrols/cardgroup/{card_group_token}/available | List Velocity Controls Card Group Balances



## ListVelocityControlsCardGroupBalances

> VelocityControlBalancePage ListVelocityControlsCardGroupBalances(ctx, cardGroupToken).Execute()

List Velocity Controls Card Group Balances



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
	cardGroupToken := "cardGroupToken_example" // string | Unique identifier of the card group for which to retrieve balances.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VelocityControlsCardGroupBalanceAPI.ListVelocityControlsCardGroupBalances(context.Background(), cardGroupToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VelocityControlsCardGroupBalanceAPI.ListVelocityControlsCardGroupBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVelocityControlsCardGroupBalances`: VelocityControlBalancePage
	fmt.Fprintf(os.Stdout, "Response from `VelocityControlsCardGroupBalanceAPI.ListVelocityControlsCardGroupBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardGroupToken** | **string** | Unique identifier of the card group for which to retrieve balances. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVelocityControlsCardGroupBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VelocityControlBalancePage**](VelocityControlBalancePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

