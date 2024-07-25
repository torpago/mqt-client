# \ProgramFundingAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProgramFundings**](ProgramFundingAPI.md#GetProgramFundings) | **Get** /admin/programs/funding | List program fundings
[**GetProgramFundingsByShortCode**](ProgramFundingAPI.md#GetProgramFundingsByShortCode) | **Get** /programs/funding | List program fundings



## GetProgramFundings

> ProgramFundingPage GetProgramFundings(ctx).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).ShortCode(shortCode).Execute()

List program fundings



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
	count := int32(56) // int32 | Number of program funding resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	startDate := "2014-01-01" // string | Start date for filtering program funding entries. (optional)
	endDate := "2014-04-01" // string | End date for filtering program funding entries. (optional)
	shortCode := "shortCode_example" // string | Short code for filtering program funding entries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramFundingAPI.GetProgramFundings(context.Background()).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).ShortCode(shortCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingAPI.GetProgramFundings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramFundings`: ProgramFundingPage
	fmt.Fprintf(os.Stdout, "Response from `ProgramFundingAPI.GetProgramFundings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramFundingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of program funding resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **startDate** | **string** | Start date for filtering program funding entries. | 
 **endDate** | **string** | End date for filtering program funding entries. | 
 **shortCode** | **string** | Short code for filtering program funding entries. | 

### Return type

[**ProgramFundingPage**](ProgramFundingPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProgramFundingsByShortCode

> ProgramFundingPage GetProgramFundingsByShortCode(ctx).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).Execute()

List program fundings



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
	count := int32(56) // int32 | Number of program funding resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	startDate := "2014-01-01" // string | Start date for filtering program funding entries. (optional)
	endDate := "2014-04-01" // string | End date for filtering program funding entries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProgramFundingAPI.GetProgramFundingsByShortCode(context.Background()).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProgramFundingAPI.GetProgramFundingsByShortCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProgramFundingsByShortCode`: ProgramFundingPage
	fmt.Fprintf(os.Stdout, "Response from `ProgramFundingAPI.GetProgramFundingsByShortCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProgramFundingsByShortCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Number of program funding resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **startDate** | **string** | Start date for filtering program funding entries. | 
 **endDate** | **string** | End date for filtering program funding entries. | 

### Return type

[**ProgramFundingPage**](ProgramFundingPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

