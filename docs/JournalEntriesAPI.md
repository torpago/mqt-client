# \JournalEntriesAPI

All URIs are relative to *https://raw.githubusercontent.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountJournalEntry**](JournalEntriesAPI.md#GetAccountJournalEntry) | **Get** /accounts/{account_token}/journalentries/{journal_entry_token} | Retrieve account journal entry
[**ListAccountJournalEntries**](JournalEntriesAPI.md#ListAccountJournalEntries) | **Get** /accounts/{account_token}/journalentries | List account journal entries
[**ResendWebhookEvent**](JournalEntriesAPI.md#ResendWebhookEvent) | **Post** /webhooks/{event_type}/{resource_token} | Resend credit event notification



## GetAccountJournalEntry

> JournalEntry GetAccountJournalEntry(ctx, accountToken, journalEntryToken).Execute()

Retrieve account journal entry



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve journal entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	journalEntryToken := "journalEntryToken_example" // string | Unique identifier of the journal entry you want to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalEntriesAPI.GetAccountJournalEntry(context.Background(), accountToken, journalEntryToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalEntriesAPI.GetAccountJournalEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountJournalEntry`: JournalEntry
	fmt.Fprintf(os.Stdout, "Response from `JournalEntriesAPI.GetAccountJournalEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve journal entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 
**journalEntryToken** | **string** | Unique identifier of the journal entry you want to retrieve.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/journalentries&#x60; to retrieve existing journal entry tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountJournalEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JournalEntry**](JournalEntry.md)

### Authorization

[mqAppAndAccessToken](../README.md#mqAppAndAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountJournalEntries

> JournalEntriesPage ListAccountJournalEntries(ctx, accountToken).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).StartImpactTime(startImpactTime).EndImpactTime(endImpactTime).StartCreatedTime(startCreatedTime).EndCreatedTime(endCreatedTime).Statuses(statuses).Groups(groups).Expand(expand).SortBy(sortBy).CardTokens(cardTokens).UserTokens(userTokens).Types(types).Execute()

List account journal entries



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
	accountToken := "accountToken_example" // string | Unique identifier of the credit account for which you want to retrieve journal entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
	count := int32(56) // int32 | Number of journal entry resources to retrieve. (optional) (default to 5)
	startIndex := int32(56) // int32 | Sort order index of the first resource in the returned array. (optional) (default to 0)
	startDate := "startDate_example" // string | Starting date of the date range from which to return journal entries. (optional)
	endDate := "endDate_example" // string | Ending date of the date range from which to return journal entries. (optional)
	startImpactTime := "startImpactTime_example" // string | Starting `impact_time` of the date range from which to return journal entries. (optional)
	endImpactTime := "endImpactTime_example" // string | Ending `impact_time` of the date range from which to return journal entries. (optional)
	startCreatedTime := "startCreatedTime_example" // string | Starting `created_date` of the date range from which to return journal entries. (optional)
	endCreatedTime := "endCreatedTime_example" // string | Ending `created_date` of the date range from which to return journal entries. (optional)
	statuses := []string{"Statuses_example"} // []string | Array of statuses by which to filter journal entries. (optional)
	groups := []string{"Groups_example"} // []string | Array of groups by which to filter journal entries.  To return all journal entry groups, do not include this query parameter. (optional)
	expand := []string{"Expand_example"} // []string | Embeds the specified object into the response. (optional)
	sortBy := "sortBy_example" // string | Field on which to sort. Prefix the field name with a hyphen (`-`) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as `createdTime`, and not by the field names appearing in response bodies such as `created_time`. (optional) (default to "-createdTime")
	cardTokens := []string{"Inner_example"} // []string | Array of card tokens by which to filter journal entries. Returns journal entries associated with the specified card tokens.  Send a `GET` request to `/credit/accounts/{account_token}/cards/` to retrieve existing card tokens. (optional)
	userTokens := []string{"Inner_example"} // []string | Array of user tokens by which to filter journal entries. Returns journal entries associated with the specified user tokens.  Send a `GET` request to `/users` to retrieve existing user tokens. (optional)
	types := []string{"Types_example"} // []string | Array of <</core-api/event-types#_credit_journal_entry_events, event types>> by which to filter journal entries.  To return all event types, do not include this query parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalEntriesAPI.ListAccountJournalEntries(context.Background(), accountToken).Count(count).StartIndex(startIndex).StartDate(startDate).EndDate(endDate).StartImpactTime(startImpactTime).EndImpactTime(endImpactTime).StartCreatedTime(startCreatedTime).EndCreatedTime(endCreatedTime).Statuses(statuses).Groups(groups).Expand(expand).SortBy(sortBy).CardTokens(cardTokens).UserTokens(userTokens).Types(types).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalEntriesAPI.ListAccountJournalEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountJournalEntries`: JournalEntriesPage
	fmt.Fprintf(os.Stdout, "Response from `JournalEntriesAPI.ListAccountJournalEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountToken** | **string** | Unique identifier of the credit account for which you want to retrieve journal entries.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts&#x60; to retrieve existing credit account tokens. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountJournalEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | Number of journal entry resources to retrieve. | [default to 5]
 **startIndex** | **int32** | Sort order index of the first resource in the returned array. | [default to 0]
 **startDate** | **string** | Starting date of the date range from which to return journal entries. | 
 **endDate** | **string** | Ending date of the date range from which to return journal entries. | 
 **startImpactTime** | **string** | Starting &#x60;impact_time&#x60; of the date range from which to return journal entries. | 
 **endImpactTime** | **string** | Ending &#x60;impact_time&#x60; of the date range from which to return journal entries. | 
 **startCreatedTime** | **string** | Starting &#x60;created_date&#x60; of the date range from which to return journal entries. | 
 **endCreatedTime** | **string** | Ending &#x60;created_date&#x60; of the date range from which to return journal entries. | 
 **statuses** | **[]string** | Array of statuses by which to filter journal entries. | 
 **groups** | **[]string** | Array of groups by which to filter journal entries.  To return all journal entry groups, do not include this query parameter. | 
 **expand** | **[]string** | Embeds the specified object into the response. | 
 **sortBy** | **string** | Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;. | [default to &quot;-createdTime&quot;]
 **cardTokens** | **[]string** | Array of card tokens by which to filter journal entries. Returns journal entries associated with the specified card tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/cards/&#x60; to retrieve existing card tokens. | 
 **userTokens** | **[]string** | Array of user tokens by which to filter journal entries. Returns journal entries associated with the specified user tokens.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve existing user tokens. | 
 **types** | **[]string** | Array of &lt;&lt;/core-api/event-types#_credit_journal_entry_events, event types&gt;&gt; by which to filter journal entries.  To return all event types, do not include this query parameter. | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	eventType := "eventType_example" // string | Specifies the type of event you want to resend.
	resourceToken := "resourceToken_example" // string | Unique identifier of the resource for which you want to resend a notification.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.  Send a `GET` request to `/credit/accounts/{account_token}/ledgerentries` to retrieve existing ledger entry tokens.  Send a `GET` request to `/accounts/{account_token}/accounttransitions` to retrieve existing account transition tokens.  Send a `GET` request to `/credit/accounts/{account_token}/payments/{payment_token}` to retrieve existing payment transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements` to retrieve existing statement summary tokens.  Send a `GET` request to `/accounts/{account_token}/delinquencystate/transitions` to retrieve existing delinquency state transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token}` to retrieve existing payment reminder tokens.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalEntriesAPI.ResendWebhookEvent(context.Background(), eventType, resourceToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalEntriesAPI.ResendWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendWebhookEvent`: WebhookEventResendContainerResponse
	fmt.Fprintf(os.Stdout, "Response from `JournalEntriesAPI.ResendWebhookEvent`: %v\n", resp)
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

