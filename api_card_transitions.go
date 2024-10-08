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
)


// CardTransitionsAPIService CardTransitionsAPI service
type CardTransitionsAPIService service

type CardTransitionsAPIGetCardtransitionsCardTokenRequest struct {
	ctx context.Context
	ApiService *CardTransitionsAPIService
	token string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
}

// Number of card transitions to retrieve.
func (r CardTransitionsAPIGetCardtransitionsCardTokenRequest) Count(count int32) CardTransitionsAPIGetCardtransitionsCardTokenRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r CardTransitionsAPIGetCardtransitionsCardTokenRequest) StartIndex(startIndex int32) CardTransitionsAPIGetCardtransitionsCardTokenRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r CardTransitionsAPIGetCardtransitionsCardTokenRequest) Fields(fields string) CardTransitionsAPIGetCardtransitionsCardTokenRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r CardTransitionsAPIGetCardtransitionsCardTokenRequest) SortBy(sortBy string) CardTransitionsAPIGetCardtransitionsCardTokenRequest {
	r.sortBy = &sortBy
	return r
}

func (r CardTransitionsAPIGetCardtransitionsCardTokenRequest) Execute() (*CardTransitionListResponse, *http.Response, error) {
	return r.ApiService.GetCardtransitionsCardTokenExecute(r)
}

/*
GetCardtransitionsCardToken List transitions for card

Retrieves a list of the transitions for a specific card.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the card.
 @return CardTransitionsAPIGetCardtransitionsCardTokenRequest
*/
func (a *CardTransitionsAPIService) GetCardtransitionsCardToken(ctx context.Context, token string) CardTransitionsAPIGetCardtransitionsCardTokenRequest {
	return CardTransitionsAPIGetCardtransitionsCardTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return CardTransitionListResponse
func (a *CardTransitionsAPIService) GetCardtransitionsCardTokenExecute(r CardTransitionsAPIGetCardtransitionsCardTokenRequest) (*CardTransitionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardTransitionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardTransitionsAPIService.GetCardtransitionsCardToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cardtransitions/card/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

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
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-createdTime"
		r.sortBy = &defaultValue
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

type CardTransitionsAPIGetCardtransitionsTokenRequest struct {
	ctx context.Context
	ApiService *CardTransitionsAPIService
	token string
	fields *string
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r CardTransitionsAPIGetCardtransitionsTokenRequest) Fields(fields string) CardTransitionsAPIGetCardtransitionsTokenRequest {
	r.fields = &fields
	return r
}

func (r CardTransitionsAPIGetCardtransitionsTokenRequest) Execute() (*CardTransitionResponse, *http.Response, error) {
	return r.ApiService.GetCardtransitionsTokenExecute(r)
}

/*
GetCardtransitionsToken Retrieve card transition

Retrieves a specific card transition.
This endpoint supports <</core-api/field-filtering, field filtering>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the card transition. Send a `GET` request to `/cardtransitions/card/{token}` to retrieve card transition tokens for a specific card.
 @return CardTransitionsAPIGetCardtransitionsTokenRequest
*/
func (a *CardTransitionsAPIService) GetCardtransitionsToken(ctx context.Context, token string) CardTransitionsAPIGetCardtransitionsTokenRequest {
	return CardTransitionsAPIGetCardtransitionsTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return CardTransitionResponse
func (a *CardTransitionsAPIService) GetCardtransitionsTokenExecute(r CardTransitionsAPIGetCardtransitionsTokenRequest) (*CardTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardTransitionsAPIService.GetCardtransitionsToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cardtransitions/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
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

type CardTransitionsAPIPostCardtransitionsRequest struct {
	ctx context.Context
	ApiService *CardTransitionsAPIService
	cardTransitionRequest *CardTransitionRequest
}

func (r CardTransitionsAPIPostCardtransitionsRequest) CardTransitionRequest(cardTransitionRequest CardTransitionRequest) CardTransitionsAPIPostCardtransitionsRequest {
	r.cardTransitionRequest = &cardTransitionRequest
	return r
}

func (r CardTransitionsAPIPostCardtransitionsRequest) Execute() (*CardTransitionResponse, *http.Response, error) {
	return r.ApiService.PostCardtransitionsExecute(r)
}

/*
PostCardtransitions Create card transition

Creates a card state transition to set the state of an existing card.

If your system uses IVR, you can send a `POST` request to `/cards/getbypan` to retrieve a card token, which you can then use in your `POST` request to `/cardtransitions`.

It may not be possible to move a card from one user to another once the card has been activated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CardTransitionsAPIPostCardtransitionsRequest
*/
func (a *CardTransitionsAPIService) PostCardtransitions(ctx context.Context) CardTransitionsAPIPostCardtransitionsRequest {
	return CardTransitionsAPIPostCardtransitionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CardTransitionResponse
func (a *CardTransitionsAPIService) PostCardtransitionsExecute(r CardTransitionsAPIPostCardtransitionsRequest) (*CardTransitionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CardTransitionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CardTransitionsAPIService.PostCardtransitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cardtransitions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.cardTransitionRequest
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
