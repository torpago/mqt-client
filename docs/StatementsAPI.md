# \StatementsAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentReminder**](StatementsAPI.md#GetPaymentReminder) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token} | Get payment reminder
[**GetPaymentRemindersByStatementSummary**](StatementsAPI.md#GetPaymentRemindersByStatementSummary) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/ | List payment reminders by statement summary
[**GetStatementFilesByAccount**](StatementsAPI.md#GetStatementFilesByAccount) | **Get** /accounts/{account_token}/statements/files | List files for an account
[**GetStatementSummariesByAccount**](StatementsAPI.md#GetStatementSummariesByAccount) | **Get** /accounts/{account_token}/statements | List account statement summaries
[**ListStatementJournalEntries**](StatementsAPI.md#ListStatementJournalEntries) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/journalentries | List account statement journal entries
[**ListStatementLedgerEntries**](StatementsAPI.md#ListStatementLedgerEntries) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/ledgerentries | List account statement ledger entries
[**ResendWebhookEvent**](StatementsAPI.md#ResendWebhookEvent) | **Post** /webhooks/{event_type}/{resource_token} | Resend credit event notification
[**RetrieveStatementFiles**](StatementsAPI.md#RetrieveStatementFiles) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/files | List files for a statement summary
[**RetrieveStatementInterestCharges**](StatementsAPI.md#RetrieveStatementInterestCharges) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/interestcharges | Retrieve account statement interest charges
[**RetrieveStatementPaymentInfo**](StatementsAPI.md#RetrieveStatementPaymentInfo) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/paymentinfo | Retrieve account statement payment information
[**RetrieveStatementReward**](StatementsAPI.md#RetrieveStatementReward) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/rewards | Retrieve account statement rewards
[**RetrieveStatementSummary**](StatementsAPI.md#RetrieveStatementSummary) | **Get** /accounts/{account_token}/statements/{statement_summary_token} | Retrieve account statement summary
[**RetrieveYearToDateForStatementSummary**](StatementsAPI.md#RetrieveYearToDateForStatementSummary) | **Get** /accounts/{account_token}/statements/{statement_summary_token}/yeartodate | Retrieve account statement year-to-date totals



## GetPaymentReminder

> PaymentReminderResponse GetPaymentReminder(ctx, accountToken, statementSummaryToken, token).Execute()

Get payment reminder



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve the statement payment reminder. Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary for which you want to retrieve the statement payment reminder. Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.
	token := "token_example" // string | Unique identifier of the payment reminder you want to retrieve. Send a `GET` request to `/credit/accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token}` to retrieve existing payment reminder tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetPaymentReminder(context.Background(), accountToken, statementSummaryToken, token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.GetPaymentReminder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentReminder`: PaymentReminderResponse
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.GetPaymentReminder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve the statement payment reminder. Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary for which you want to retrieve the statement payment reminder. Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 
**token** | **string** | Unique identifier of the payment reminder you want to retrieve. Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token}&#x60; to retrieve existing payment reminder tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentReminderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaymentReminderResponse**](PaymentReminderResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentRemindersByStatementSummary

> PaymentReminderPage GetPaymentRemindersByStatementSummary(ctx, accountToken, statementSummaryToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List payment reminders by statement summary



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve the statement payment information. Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve the payment information. Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.
	count := int32(56) // int32 | Number of payment reminder resources to retrieve. (optional) (default to 10)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order. *NOTE:* You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetPaymentRemindersByStatementSummary(context.Background(), accountToken, statementSummaryToken).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.GetPaymentRemindersByStatementSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentRemindersByStatementSummary`: PaymentReminderPage
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.GetPaymentRemindersByStatementSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve the statement payment information. Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve the payment information. Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRemindersByStatementSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Number of payment reminder resources to retrieve. | [default to 10]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order. *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]

### Return type

[**PaymentReminderPage**](PaymentReminderPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatementFilesByAccount

> StatementFilePage GetStatementFilesByAccount(ctx, accountToken).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List files for an account



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which to retrieve statement files.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	startDate := time.Now() // time.Time | Start date of the date range for which to return statement files. (optional)
	endDate := time.Now() // time.Time | End date of the date range for which to return statement files. (optional)
	count := int32(56) // int32 | Number of statement file resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetStatementFilesByAccount(context.Background(), accountToken).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.GetStatementFilesByAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatementFilesByAccount`: StatementFilePage
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.GetStatementFilesByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which to retrieve statement files.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementFilesByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Start date of the date range for which to return statement files. | 
 **endDate** | **time.Time** | End date of the date range for which to return statement files. | 
 **count** | **int32** | Number of statement file resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]

### Return type

[**StatementFilePage**](StatementFilePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatementSummariesByAccount

> StatementSummaryPage GetStatementSummariesByAccount(ctx, accountToken).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()

List account statement summaries



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve statement summaries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	startDate := time.Now() // time.Time | Returns statements with a matching opening date.  If both `start_date` and `end_date` are specified, statements whose closing date falls between the start and end dates are returned. (optional)
	endDate := time.Now() // time.Time | Returns statements with a matching closing date.  If both `start_date` and `end_date` are specified, statements whose closing date falls between the start and end dates are returned. (optional)
	count := int32(56) // int32 | Number of account statement resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.GetStatementSummariesByAccount(context.Background(), accountToken).StartDate(startDate).EndDate(endDate).Count(count).StartIndex(startIndex).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.GetStatementSummariesByAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatementSummariesByAccount`: StatementSummaryPage
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.GetStatementSummariesByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve statement summaries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementSummariesByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Returns statements with a matching opening date.  If both &#x60;start_date&#x60; and &#x60;end_date&#x60; are specified, statements whose closing date falls between the start and end dates are returned. | 
 **endDate** | **time.Time** | Returns statements with a matching closing date.  If both &#x60;start_date&#x60; and &#x60;end_date&#x60; are specified, statements whose closing date falls between the start and end dates are returned. | 
 **count** | **int32** | Number of account statement resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]

### Return type

[**StatementSummaryPage**](StatementSummaryPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStatementJournalEntries

> JournalEntriesPage ListStatementJournalEntries(ctx, accountToken, statementSummaryToken).Count(count).StartIndex(startIndex).Expand(expand).SortBy(sortBy).Execute()

List account statement journal entries



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which to retrieve the statement journal entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve journal entries.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.
	count := int32(56) // int32 | Number of journal entry resources to return. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	expand := []string{"Expand_example"} // []string | Embeds the specified object into the response. (optional)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "lastModifiedTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.ListStatementJournalEntries(context.Background(), accountToken, statementSummaryToken).Count(count).StartIndex(startIndex).Expand(expand).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.ListStatementJournalEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStatementJournalEntries`: JournalEntriesPage
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.ListStatementJournalEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which to retrieve the statement journal entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve journal entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStatementJournalEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Number of journal entry resources to return. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **expand** | **[]string** | Embeds the specified object into the response. | 
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;lastModifiedTime&quot;]

### Return type

[**JournalEntriesPage**](JournalEntriesPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStatementLedgerEntries

> []LedgerEntry ListStatementLedgerEntries(ctx, accountToken, statementSummaryToken).Expand(expand).SortBy(sortBy).Execute()

List account statement ledger entries



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which to retrieve the statement ledger entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve ledger entries.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.
	expand := []string{"Expand_example"} // []string | Embeds the specified object into the response. (optional)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `lastModifiedTime`, and not by the field names appearing in response bodies such as `last_modified_time`. (optional) (default to "lastModifiedTime")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.ListStatementLedgerEntries(context.Background(), accountToken, statementSummaryToken).Expand(expand).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.ListStatementLedgerEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStatementLedgerEntries`: []LedgerEntry
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.ListStatementLedgerEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which to retrieve the statement ledger entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve ledger entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStatementLedgerEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **[]string** | Embeds the specified object into the response. | 
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;lastModifiedTime&#x60;, and not by the field names appearing in response bodies such as &#x60;last_modified_time&#x60;. | [default to &quot;lastModifiedTime&quot;]

### Return type

[**[]LedgerEntry**](LedgerEntry.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendWebhookEvent

> WebhookEventResendContainerResponse ResendWebhookEvent(ctx, eventType, resourceToken).Execute()

Resend credit event notification



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
	eventType := "eventType_example" // string | Specifies the type of event you want to resend.
	resourceToken := "resourceToken_example" // string | Unique identifier of the resource for which you want to resend a notification.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.  Send a `GET` request to `/credit/accounts/{account_token}/ledgerentries` to retrieve existing ledger entry tokens.  Send a `GET` request to `/accounts/{account_token}/accounttransitions` to retrieve existing account transition tokens.  Send a `GET` request to `/credit/accounts/{account_token}/payments/{payment_token}` to retrieve existing payment transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements` to retrieve existing statement summary tokens.  Send a `GET` request to `/accounts/{account_token}/delinquencystate/transitions` to retrieve existing delinquency state transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token}` to retrieve existing payment reminder tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.ResendWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendWebhookEvent`: WebhookEventResendContainerResponse
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.ResendWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventType** | **string** | Specifies the type of event you want to resend. | 
**resourceToken** | **string** | Unique identifier of the resource for which you want to resend a notification.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/journalentries&#x60; to retrieve existing journal entry tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/ledgerentries&#x60; to retrieve existing ledger entry tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/accounttransitions&#x60; to retrieve existing account transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/payments/{payment_token}&#x60; to retrieve existing payment transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/statements&#x60; to retrieve existing statement summary tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/delinquencystate/transitions&#x60; to retrieve existing delinquency state transition tokens.  Send a &#x60;GET&#x60; request to &#x60;/accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token}&#x60; to retrieve existing payment reminder tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebhookEventResendContainerResponse**](WebhookEventResendContainerResponse.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatementFiles

> StatementFilePage RetrieveStatementFiles(ctx, accountToken, statementSummaryToken).Count(count).StartIndex(startIndex).Execute()

List files for a statement summary



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which to retrieve statement files for a statement summary.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary whose statement files you want to retrieve.  Send a `GET` request to `/credit/accounts/{token}/statements` to retrieve existing statement summary tokens.
	count := int32(56) // int32 | Number of statement files to return. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index from which to begin returning files. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.RetrieveStatementFiles(context.Background(), accountToken, statementSummaryToken).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.RetrieveStatementFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveStatementFiles`: StatementFilePage
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.RetrieveStatementFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which to retrieve statement files for a statement summary.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary whose statement files you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatementFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | Number of statement files to return. | [default to 5]
 **startIndex** | **int32** | Sort order index from which to begin returning files. | [default to 0]

### Return type

[**StatementFilePage**](StatementFilePage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatementInterestCharges

> StatementInterestChargesPage RetrieveStatementInterestCharges(ctx, accountToken, statementSummaryToken).Execute()

Retrieve account statement interest charges



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve the statement interest charges.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve the interest charges.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.RetrieveStatementInterestCharges(context.Background(), accountToken, statementSummaryToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.RetrieveStatementInterestCharges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveStatementInterestCharges`: StatementInterestChargesPage
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.RetrieveStatementInterestCharges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve the statement interest charges.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve the interest charges.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatementInterestChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatementInterestChargesPage**](StatementInterestChargesPage.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatementPaymentInfo

> StatementPaymentInfo RetrieveStatementPaymentInfo(ctx, accountToken, statementSummaryToken).Execute()

Retrieve account statement payment information



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve the statement payment information.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve the payment information.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.RetrieveStatementPaymentInfo(context.Background(), accountToken, statementSummaryToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.RetrieveStatementPaymentInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveStatementPaymentInfo`: StatementPaymentInfo
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.RetrieveStatementPaymentInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve the statement payment information.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve the payment information.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatementPaymentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatementPaymentInfo**](StatementPaymentInfo.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatementReward

> StatementReward RetrieveStatementReward(ctx, accountToken, statementSummaryToken).Execute()

Retrieve account statement rewards



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account from which to retrieve statement rewards.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve rewards.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.RetrieveStatementReward(context.Background(), accountToken, statementSummaryToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.RetrieveStatementReward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveStatementReward`: StatementReward
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.RetrieveStatementReward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account from which to retrieve statement rewards.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve rewards.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatementRewardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatementReward**](StatementReward.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatementSummary

> StatementSummary RetrieveStatementSummary(ctx, accountToken, statementSummaryToken).Execute()

Retrieve account statement summary



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve a statement summary.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary to retrieve.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.RetrieveStatementSummary(context.Background(), accountToken, statementSummaryToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.RetrieveStatementSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveStatementSummary`: StatementSummary
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.RetrieveStatementSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve a statement summary.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatementSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatementSummary**](StatementSummary.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveYearToDateForStatementSummary

> YearToDate RetrieveYearToDateForStatementSummary(ctx, accountToken, statementSummaryToken).Execute()

Retrieve account statement year-to-date totals



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account from which to retrieve statement year-to-date totals.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	statementSummaryToken := "statementSummaryToken_example" // string | Unique identifier of the statement summary from which to retrieve year-to-date totals.  Send a `GET` request to `/credit/accounts/{token}/statements/` to retrieve existing statement summary tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatementsAPI.RetrieveYearToDateForStatementSummary(context.Background(), accountToken, statementSummaryToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatementsAPI.RetrieveYearToDateForStatementSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveYearToDateForStatementSummary`: YearToDate
	fmt.Fprintf(os.Stdout, "Response from `StatementsAPI.RetrieveYearToDateForStatementSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account from which to retrieve statement year-to-date totals.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**statementSummaryToken** | **string** | Unique identifier of the statement summary from which to retrieve year-to-date totals.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{token}/statements/&#x60; to retrieve existing statement summary tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveYearToDateForStatementSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**YearToDate**](YearToDate.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

