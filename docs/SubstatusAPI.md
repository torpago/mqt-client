# \SubstatusAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubStatus**](SubstatusAPI.md#CreateSubStatus) | **Post** /substatuses | Create substatus
[**ListSubStatuses**](SubstatusAPI.md#ListSubStatuses) | **Get** /substatuses | List substatuses
[**RetrieveSubStatus**](SubstatusAPI.md#RetrieveSubStatus) | **Get** /substatuses/{substatus_token} | Retrieve Substatus
[**UpdateSubStatus**](SubstatusAPI.md#UpdateSubStatus) | **Put** /substatuses/{substatus_token} | Update substatus



## CreateSubStatus

> SubstatusResponse CreateSubStatus(ctx).SubstatusCreateReq(substatusCreateReq).Execute()

Create substatus



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
	substatusCreateReq := *openapiclient.NewSubstatusCreateReq([]openapiclient.SubstatusEvent{*openapiclient.NewSubstatusEvent("State_example")}, "ResourceToken_example", "ResourceType_example", "Substatus_example") // SubstatusCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubstatusAPI.CreateSubStatus(context.Background()).SubstatusCreateReq(substatusCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubstatusAPI.CreateSubStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubStatus`: SubstatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SubstatusAPI.CreateSubStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **substatusCreateReq** | [**SubstatusCreateReq**](SubstatusCreateReq.md) |  | 

### Return type

[**SubstatusResponse**](SubstatusResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubStatuses

> SubstatusPage ListSubStatuses(ctx).AccountToken(accountToken).UserToken(userToken).IsActive(isActive).Substatuses(substatuses).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List substatuses



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
	accountToken := "accountToken_example" // string | The account token to filter by. (optional)
	userToken := "userToken_example" // string | The user token to filter by. (optional)
	isActive := true // bool |  (optional)
	substatuses := []string{"Substatuses_example"} // []string | comma-deliminited list of substatuses types to include Allowable values: HARDSHIP,FRAUD,MLA,SCRA,DECEASED,BANKRUPTCY (optional)
	count := int32(56) // int32 | The number of resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Show n-th paginated page (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. (optional) (default to "-createdTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubstatusAPI.ListSubStatuses(context.Background()).AccountToken(accountToken).UserToken(userToken).IsActive(isActive).Substatuses(substatuses).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubstatusAPI.ListSubStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubStatuses`: SubstatusPage
	fmt.Fprintf(os.Stdout, "Response from `SubstatusAPI.ListSubStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountToken** | **string** | The account token to filter by. | 
 **userToken** | **string** | The user token to filter by. | 
 **isActive** | **bool** |  | 
 **substatuses** | **[]string** | comma-deliminited list of substatuses types to include Allowable values: HARDSHIP,FRAUD,MLA,SCRA,DECEASED,BANKRUPTCY | 
 **count** | **int32** | The number of resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Show n-th paginated page | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. | [default to &quot;-createdTime&quot;]

### Return type

[**SubstatusPage**](SubstatusPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSubStatus

> SubstatusResponse RetrieveSubStatus(ctx, substatusToken).Execute()

Retrieve Substatus



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
	substatusToken := "substatusToken_example" // string | The unique identifier of the substatus to retrieve.  Send a `GET` request to `/credit/substatuses` to retrieve existing substatus tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubstatusAPI.RetrieveSubStatus(context.Background(), substatusToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubstatusAPI.RetrieveSubStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveSubStatus`: SubstatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SubstatusAPI.RetrieveSubStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**substatusToken** | **string** | The unique identifier of the substatus to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/substatuses&#x60; to retrieve existing substatus tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSubStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubstatusResponse**](SubstatusResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubStatus

> SubstatusResponse UpdateSubStatus(ctx, substatusToken).SubstatusUpdateReq(substatusUpdateReq).Execute()

Update substatus



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
	substatusToken := "substatusToken_example" // string | The unique identifier of the substatus to update.
	substatusUpdateReq := *openapiclient.NewSubstatusUpdateReq("State_example") // SubstatusUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubstatusAPI.UpdateSubStatus(context.Background(), substatusToken).SubstatusUpdateReq(substatusUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubstatusAPI.UpdateSubStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubStatus`: SubstatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SubstatusAPI.UpdateSubStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**substatusToken** | **string** | The unique identifier of the substatus to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **substatusUpdateReq** | [**SubstatusUpdateReq**](SubstatusUpdateReq.md) |  | 

### Return type

[**SubstatusResponse**](SubstatusResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

