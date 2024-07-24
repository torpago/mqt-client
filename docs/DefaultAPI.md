# \DefaultAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeedbackFraudPost**](DefaultAPI.md#FeedbackFraudPost) | **Post** /feedback/fraud | Creates a fraud feedback



## FeedbackFraudPost

> FraudFeedbackResponse FeedbackFraudPost(ctx).FraudFeedbackRequest(fraudFeedbackRequest).Execute()

Creates a fraud feedback



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
	fraudFeedbackRequest := *openapiclient.NewFraudFeedbackRequest("Actor_example", "Amount_example", false, "Status_example", "TransactionToken_example") // FraudFeedbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FeedbackFraudPost(context.Background()).FraudFeedbackRequest(fraudFeedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FeedbackFraudPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedbackFraudPost`: FraudFeedbackResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FeedbackFraudPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeedbackFraudPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fraudFeedbackRequest** | [**FraudFeedbackRequest**](FraudFeedbackRequest.md) |  | 

### Return type

[**FraudFeedbackResponse**](FraudFeedbackResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

