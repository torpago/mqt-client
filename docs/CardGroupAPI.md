# \CardGroupAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCardGroup**](CardGroupAPI.md#CreateCardGroup) | **Post** /cardgroups | Create Card Group
[**ListCardGroups**](CardGroupAPI.md#ListCardGroups) | **Get** /cardgroups | List Card Groups
[**RetrieveCardGroup**](CardGroupAPI.md#RetrieveCardGroup) | **Get** /cardgroups/{token} | Retrieve Card Group



## CreateCardGroup

> CardGroup CreateCardGroup(ctx).CardGroupReq(cardGroupReq).Execute()

Create Card Group



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
	cardGroupReq := *openapiclient.NewCardGroupReq("CardToken_example") // CardGroupReq | Card group to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardGroupAPI.CreateCardGroup(context.Background()).CardGroupReq(cardGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardGroupAPI.CreateCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCardGroup`: CardGroup
	fmt.Fprintf(os.Stdout, "Response from `CardGroupAPI.CreateCardGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardGroupReq** | [**CardGroupReq**](CardGroupReq.md) | Card group to create. | 

### Return type

[**CardGroup**](CardGroup.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCardGroups

> CardGroupPage ListCardGroups(ctx).CardTokens(cardTokens).Count(count).StartIndex(startIndex).Execute()

List Card Groups



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
	cardTokens := []string{"Inner_example"} // []string | list of unique card identifiers to retrieve. (optional)
	count := int32(56) // int32 | Number of card group resources to retrieve. (optional)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardGroupAPI.ListCardGroups(context.Background()).CardTokens(cardTokens).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardGroupAPI.ListCardGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCardGroups`: CardGroupPage
	fmt.Fprintf(os.Stdout, "Response from `CardGroupAPI.ListCardGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCardGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardTokens** | **[]string** | list of unique card identifiers to retrieve. | 
 **count** | **int32** | Number of card group resources to retrieve. | 
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | 

### Return type

[**CardGroupPage**](CardGroupPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCardGroup

> CardGroup RetrieveCardGroup(ctx, token).Execute()

Retrieve Card Group



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
	token := "token_example" // string | Unique identifier of the card group to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CardGroupAPI.RetrieveCardGroup(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardGroupAPI.RetrieveCardGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCardGroup`: CardGroup
	fmt.Fprintf(os.Stdout, "Response from `CardGroupAPI.RetrieveCardGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Unique identifier of the card group to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCardGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardGroup**](CardGroup.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

