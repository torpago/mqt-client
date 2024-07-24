# \ProgramGatewaysAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProgramGateway**](ProgramGatewaysAPI.md#CreateProgramGateway) | **Post** /programgateways | Create Program Gateway
[**ListProgramGateways**](ProgramGatewaysAPI.md#ListProgramGateways) | **Get** /programgateways | List Program Gateways
[**RetrieveProgramGateway**](ProgramGatewaysAPI.md#RetrieveProgramGateway) | **Get** /programgateways/{token} | Retrieve Program Gateway
[**UpdateProgramGateway**](ProgramGatewaysAPI.md#UpdateProgramGateway) | **Put** /programgateways/{token} | Update Program Gateway



## CreateProgramGateway

> ProgramGatewayResponse CreateProgramGateway(ctx).ProgramGatewayCreateReq(programGatewayCreateReq).Execute()

Create Program Gateway



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
	programGatewayCreateReq := *openapiclient.NewProgramGatewayCreateReq("BasicAuthPassword_example", "BasicAuthUsername_example", "Name_example", "Url_example") // ProgramGatewayCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramGatewaysAPI.CreateProgramGateway(context.Background()).ProgramGatewayCreateReq(programGatewayCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewaysAPI.CreateProgramGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProgramGateway`: ProgramGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgramGatewaysAPI.CreateProgramGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProgramGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programGatewayCreateReq** | [**ProgramGatewayCreateReq**](ProgramGatewayCreateReq.md) |  | 

### Return type

[**ProgramGatewayResponse**](ProgramGatewayResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProgramGateways

> ProgramGatewayPage ListProgramGateways(ctx).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List Program Gateways



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
	count := int32(56) // int32 | Number of program gateway resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "-lastModifiedTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramGatewaysAPI.ListProgramGateways(context.Background()).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewaysAPI.ListProgramGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProgramGateways`: ProgramGatewayPage
	fmt.Fprintf(os.Stdout, "Response from `ProgramGatewaysAPI.ListProgramGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProgramGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of program gateway resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;-lastModifiedTime&quot;]

### Return type

[**ProgramGatewayPage**](ProgramGatewayPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveProgramGateway

> ProgramGatewayResponse RetrieveProgramGateway(ctx, token).Execute()

Retrieve Program Gateway



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
	token := "token_example" // string | Unique identifier of the Program Gateway to retrieve.  Send a `GET` request to `/credit/programgateways` to retrieve existing Program Gateway tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramGatewaysAPI.RetrieveProgramGateway(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewaysAPI.RetrieveProgramGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveProgramGateway`: ProgramGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgramGatewaysAPI.RetrieveProgramGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the Program Gateway to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/programgateways&#x60; to retrieve existing Program Gateway tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveProgramGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramGatewayResponse**](ProgramGatewayResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProgramGateway

> ProgramGatewayResponse UpdateProgramGateway(ctx, token).ProgramGatewayUpdateReq(programGatewayUpdateReq).Execute()

Update Program Gateway



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
	token := "token_example" // string | Unique identifier of the Program Gateway to update.
	programGatewayUpdateReq := *openapiclient.NewProgramGatewayUpdateReq() // ProgramGatewayUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramGatewaysAPI.UpdateProgramGateway(context.Background(), token).ProgramGatewayUpdateReq(programGatewayUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramGatewaysAPI.UpdateProgramGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProgramGateway`: ProgramGatewayResponse
	fmt.Fprintf(os.Stdout, "Response from `ProgramGatewaysAPI.UpdateProgramGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the Program Gateway to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProgramGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **programGatewayUpdateReq** | [**ProgramGatewayUpdateReq**](ProgramGatewayUpdateReq.md) |  | 

### Return type

[**ProgramGatewayResponse**](ProgramGatewayResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

