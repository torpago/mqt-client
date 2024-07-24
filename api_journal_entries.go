/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// JournalEntriesAPIService JournalEntriesAPI service
type JournalEntriesAPIService service

type JournalEntriesAPIGetAccountJournalEntryRequest struct {
	ctx context.Context
	ApiService *JournalEntriesAPIService
	accountToken string
	journalEntryToken string
}

func (r JournalEntriesAPIGetAccountJournalEntryRequest) Execute() (*JournalEntry, *http.Response, error) {
	return r.ApiService.GetAccountJournalEntryExecute(r)
}

/*
GetAccountJournalEntry Retrieve account journal entry

Retrieve a journal entry for a credit account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken Unique identifier of the credit account for which you want to retrieve journal entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @param journalEntryToken Unique identifier of the journal entry you want to retrieve.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.
 @return JournalEntriesAPIGetAccountJournalEntryRequest
*/
func (a *JournalEntriesAPIService) GetAccountJournalEntry(ctx context.Context, accountToken string, journalEntryToken string) JournalEntriesAPIGetAccountJournalEntryRequest {
	return JournalEntriesAPIGetAccountJournalEntryRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
		journalEntryToken: journalEntryToken,
	}
}

// Execute executes the request
//  @return JournalEntry
func (a *JournalEntriesAPIService) GetAccountJournalEntryExecute(r JournalEntriesAPIGetAccountJournalEntryRequest) (*JournalEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JournalEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JournalEntriesAPIService.GetAccountJournalEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_token}/journalentries/{journal_entry_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"journal_entry_token"+"}", url.PathEscape(parameterValueToString(r.journalEntryToken, "journalEntryToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JournalEntriesAPIListAccountJournalEntriesRequest struct {
	ctx context.Context
	ApiService *JournalEntriesAPIService
	accountToken string
	count *int32
	startIndex *int32
	startDate *string
	endDate *string
	startImpactTime *string
	endImpactTime *string
	startCreatedTime *string
	endCreatedTime *string
	statuses *[]string
	groups *[]string
	expand *[]string
	sortBy *string
	cardTokens *[]string
	userTokens *[]string
	types *[]string
}

// Number of journal entry resources to retrieve.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) Count(count int32) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) StartIndex(startIndex int32) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.startIndex = &startIndex
	return r
}

// Starting date of the date range from which to return journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) StartDate(startDate string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.startDate = &startDate
	return r
}

// Ending date of the date range from which to return journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) EndDate(endDate string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.endDate = &endDate
	return r
}

// Starting &#x60;impact_time&#x60; of the date range from which to return journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) StartImpactTime(startImpactTime string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.startImpactTime = &startImpactTime
	return r
}

// Ending &#x60;impact_time&#x60; of the date range from which to return journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) EndImpactTime(endImpactTime string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.endImpactTime = &endImpactTime
	return r
}

// Starting &#x60;created_date&#x60; of the date range from which to return journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) StartCreatedTime(startCreatedTime string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.startCreatedTime = &startCreatedTime
	return r
}

// Ending &#x60;created_date&#x60; of the date range from which to return journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) EndCreatedTime(endCreatedTime string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.endCreatedTime = &endCreatedTime
	return r
}

// Array of statuses by which to filter journal entries.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) Statuses(statuses []string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.statuses = &statuses
	return r
}

// Array of groups by which to filter journal entries.  To return all journal entry groups, do not include this query parameter.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) Groups(groups []string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.groups = &groups
	return r
}

// Embeds the specified object into the response.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) Expand(expand []string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.expand = &expand
	return r
}

// Field on which to sort. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.  *NOTE:* You must sort using system field names such as &#x60;createdTime&#x60;, and not by the field names appearing in response bodies such as &#x60;created_time&#x60;.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) SortBy(sortBy string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.sortBy = &sortBy
	return r
}

// Array of card tokens by which to filter journal entries. Returns journal entries associated with the specified card tokens.  Send a &#x60;GET&#x60; request to &#x60;/credit/accounts/{account_token}/cards/&#x60; to retrieve existing card tokens.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) CardTokens(cardTokens []string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.cardTokens = &cardTokens
	return r
}

// Array of user tokens by which to filter journal entries. Returns journal entries associated with the specified user tokens.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve existing user tokens.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) UserTokens(userTokens []string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.userTokens = &userTokens
	return r
}

// Array of &lt;&lt;/core-api/event-types#_credit_journal_entry_events, event types&gt;&gt; by which to filter journal entries.  To return all event types, do not include this query parameter.
func (r JournalEntriesAPIListAccountJournalEntriesRequest) Types(types []string) JournalEntriesAPIListAccountJournalEntriesRequest {
	r.types = &types
	return r
}

func (r JournalEntriesAPIListAccountJournalEntriesRequest) Execute() (*JournalEntriesPage, *http.Response, error) {
	return r.ApiService.ListAccountJournalEntriesExecute(r)
}

/*
ListAccountJournalEntries List account journal entries

Retrieve an array of journal entries on a credit account.

This endpoint supports <</core-api/sorting-and-pagination, sorting and pagination>> and <</core-api/object-expansion, object expansion>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken Unique identifier of the credit account for which you want to retrieve journal entries.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @return JournalEntriesAPIListAccountJournalEntriesRequest
*/
func (a *JournalEntriesAPIService) ListAccountJournalEntries(ctx context.Context, accountToken string) JournalEntriesAPIListAccountJournalEntriesRequest {
	return JournalEntriesAPIListAccountJournalEntriesRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
	}
}

// Execute executes the request
//  @return JournalEntriesPage
func (a *JournalEntriesAPIService) ListAccountJournalEntriesExecute(r JournalEntriesAPIListAccountJournalEntriesRequest) (*JournalEntriesPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JournalEntriesPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JournalEntriesAPIService.ListAccountJournalEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_token}/journalentries"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int32 = 0
		r.startIndex = &defaultValue
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	if r.startImpactTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_impact_time", r.startImpactTime, "")
	}
	if r.endImpactTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_impact_time", r.endImpactTime, "")
	}
	if r.startCreatedTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_created_time", r.startCreatedTime, "")
	}
	if r.endCreatedTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_created_time", r.endCreatedTime, "")
	}
	if r.statuses != nil {
		t := *r.statuses
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "statuses", t, "multi")
		}
	}
	if r.groups != nil {
		t := *r.groups
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "groups", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "groups", t, "multi")
		}
	}
	if r.expand != nil {
		t := *r.expand
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "expand", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "expand", t, "multi")
		}
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-createdTime"
		r.sortBy = &defaultValue
	}
	if r.cardTokens != nil {
		t := *r.cardTokens
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "card_tokens", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "card_tokens", t, "multi")
		}
	}
	if r.userTokens != nil {
		t := *r.userTokens
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "user_tokens", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "user_tokens", t, "multi")
		}
	}
	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type JournalEntriesAPIResendWebhookEventRequest struct {
	ctx context.Context
	ApiService *JournalEntriesAPIService
	eventType string
	resourceToken string
}

func (r JournalEntriesAPIResendWebhookEventRequest) Execute() (*WebhookEventResendContainerResponse, *http.Response, error) {
	return r.ApiService.ResendWebhookEventExecute(r)
}

/*
ResendWebhookEvent Resend credit event notification

Resends a credit event notification to your webhook endpoint.

Although you send this request as a `POST`, all parameters are passed in the URL and the body is empty.
The event notification is resent to your webhook endpoint and also returned in the response to this request.

For details on how to configure your webhook endpoint, see the About Webhooks <</developer-guides/about-webhooks#_tutorial, tutorial>>.
For the complete `/webhooks` endpoint reference, see <</core-api/webhooks, Webhooks>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventType Specifies the type of event you want to resend.
 @param resourceToken Unique identifier of the resource for which you want to resend a notification.  Send a `GET` request to `/credit/accounts/{account_token}/journalentries` to retrieve existing journal entry tokens.  Send a `GET` request to `/credit/accounts/{account_token}/ledgerentries` to retrieve existing ledger entry tokens.  Send a `GET` request to `/accounts/{account_token}/accounttransitions` to retrieve existing account transition tokens.  Send a `GET` request to `/credit/accounts/{account_token}/payments/{payment_token}` to retrieve existing payment transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements` to retrieve existing statement summary tokens.  Send a `GET` request to `/accounts/{account_token}/delinquencystate/transitions` to retrieve existing delinquency state transition tokens.  Send a `GET` request to `/accounts/{account_token}/statements/{statement_summary_token}/paymentreminders/{token}` to retrieve existing payment reminder tokens.
 @return JournalEntriesAPIResendWebhookEventRequest
*/
func (a *JournalEntriesAPIService) ResendWebhookEvent(ctx context.Context, eventType string, resourceToken string) JournalEntriesAPIResendWebhookEventRequest {
	return JournalEntriesAPIResendWebhookEventRequest{
		ApiService: a,
		ctx: ctx,
		eventType: eventType,
		resourceToken: resourceToken,
	}
}

// Execute executes the request
//  @return WebhookEventResendContainerResponse
func (a *JournalEntriesAPIService) ResendWebhookEventExecute(r JournalEntriesAPIResendWebhookEventRequest) (*WebhookEventResendContainerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEventResendContainerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JournalEntriesAPIService.ResendWebhookEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{event_type}/{resource_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"event_type"+"}", url.PathEscape(parameterValueToString(r.eventType, "eventType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resource_token"+"}", url.PathEscape(parameterValueToString(r.resourceToken, "resourceToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
