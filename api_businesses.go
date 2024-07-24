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


// BusinessesAPIService BusinessesAPI service
type BusinessesAPIService service

type BusinessesAPIGetBusinessesRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	count *int32
	startIndex *int32
	businessNameDba *string
	businessNameLegal *string
	searchType *string
	fields *string
	sortBy *string
}

// Number of business resources to retrieve.
func (r BusinessesAPIGetBusinessesRequest) Count(count int32) BusinessesAPIGetBusinessesRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r BusinessesAPIGetBusinessesRequest) StartIndex(startIndex int32) BusinessesAPIGetBusinessesRequest {
	r.startIndex = &startIndex
	return r
}

// Fictitious or \&quot;doing business as (DBA)\&quot; name of the business.
func (r BusinessesAPIGetBusinessesRequest) BusinessNameDba(businessNameDba string) BusinessesAPIGetBusinessesRequest {
	r.businessNameDba = &businessNameDba
	return r
}

// Legal name of the business.
func (r BusinessesAPIGetBusinessesRequest) BusinessNameLegal(businessNameLegal string) BusinessesAPIGetBusinessesRequest {
	r.businessNameLegal = &businessNameLegal
	return r
}

// Specifies the search type for the query.
func (r BusinessesAPIGetBusinessesRequest) SearchType(searchType string) BusinessesAPIGetBusinessesRequest {
	r.searchType = &searchType
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r BusinessesAPIGetBusinessesRequest) Fields(fields string) BusinessesAPIGetBusinessesRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r BusinessesAPIGetBusinessesRequest) SortBy(sortBy string) BusinessesAPIGetBusinessesRequest {
	r.sortBy = &sortBy
	return r
}

func (r BusinessesAPIGetBusinessesRequest) Execute() (*BusinessCardHolderListResponse, *http.Response, error) {
	return r.ApiService.GetBusinessesExecute(r)
}

/*
GetBusinesses List businesses

To return an array of all businesses, send a `GET` request to the `/businesses` endpoint.

To narrow your result set to businesses that match a particular legal or fictitious name, include the appropriate parameters from the following query parameters table.
This endpoint also supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, sorting and pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BusinessesAPIGetBusinessesRequest
*/
func (a *BusinessesAPIService) GetBusinesses(ctx context.Context) BusinessesAPIGetBusinessesRequest {
	return BusinessesAPIGetBusinessesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BusinessCardHolderListResponse
func (a *BusinessesAPIService) GetBusinessesExecute(r BusinessesAPIGetBusinessesRequest) (*BusinessCardHolderListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessCardHolderListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.GetBusinesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses"

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
	if r.businessNameDba != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "business_name_dba", r.businessNameDba, "")
	}
	if r.businessNameLegal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "business_name_legal", r.businessNameLegal, "")
	}
	if r.searchType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_type", r.searchType, "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_by", r.sortBy, "")
	} else {
		var defaultValue string = "-lastModifiedTime"
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

type BusinessesAPIGetBusinessesParenttokenChildrenRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	parentToken string
	count *int32
	startIndex *int32
	fields *string
	sortBy *string
}

// Number of child cardholders to retrieve.
func (r BusinessesAPIGetBusinessesParenttokenChildrenRequest) Count(count int32) BusinessesAPIGetBusinessesParenttokenChildrenRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r BusinessesAPIGetBusinessesParenttokenChildrenRequest) StartIndex(startIndex int32) BusinessesAPIGetBusinessesParenttokenChildrenRequest {
	r.startIndex = &startIndex
	return r
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r BusinessesAPIGetBusinessesParenttokenChildrenRequest) Fields(fields string) BusinessesAPIGetBusinessesParenttokenChildrenRequest {
	r.fields = &fields
	return r
}

// Field on which to sort. Use any field in the resource model, or one of the system fields &#x60;lastModifiedTime&#x60; or &#x60;createdTime&#x60;. Prefix the field name with a hyphen (&#x60;-&#x60;) to sort in descending order. Omit the hyphen to sort in ascending order.
func (r BusinessesAPIGetBusinessesParenttokenChildrenRequest) SortBy(sortBy string) BusinessesAPIGetBusinessesParenttokenChildrenRequest {
	r.sortBy = &sortBy
	return r
}

func (r BusinessesAPIGetBusinessesParenttokenChildrenRequest) Execute() (*BusinessUserCardHolderListResponse, *http.Response, error) {
	return r.ApiService.GetBusinessesParenttokenChildrenExecute(r)
}

/*
GetBusinessesParenttokenChildren List business children

To return an array of all child cardholders of a particular business, send a `GET` request to the `/businesses/{parent_token}/children` endpoint.
Include the `parent_token` as a URL path parameter.

This endpoint supports <</core-api/field-filtering, field filtering>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param parentToken Unique identifier of the parent business.
 @return BusinessesAPIGetBusinessesParenttokenChildrenRequest
*/
func (a *BusinessesAPIService) GetBusinessesParenttokenChildren(ctx context.Context, parentToken string) BusinessesAPIGetBusinessesParenttokenChildrenRequest {
	return BusinessesAPIGetBusinessesParenttokenChildrenRequest{
		ApiService: a,
		ctx: ctx,
		parentToken: parentToken,
	}
}

// Execute executes the request
//  @return BusinessUserCardHolderListResponse
func (a *BusinessesAPIService) GetBusinessesParenttokenChildrenExecute(r BusinessesAPIGetBusinessesParenttokenChildrenRequest) (*BusinessUserCardHolderListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessUserCardHolderListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.GetBusinessesParenttokenChildren")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses/{parent_token}/children"
	localVarPath = strings.Replace(localVarPath, "{"+"parent_token"+"}", url.PathEscape(parameterValueToString(r.parentToken, "parentToken")), -1)

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
		var defaultValue string = "-lastModifiedTime"
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

type BusinessesAPIGetBusinessesTokenRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	token string
	fields *string
}

// Comma-delimited list of fields to return (&#x60;field_1,field_2&#x60;, and so on). Leave blank to return all fields.
func (r BusinessesAPIGetBusinessesTokenRequest) Fields(fields string) BusinessesAPIGetBusinessesTokenRequest {
	r.fields = &fields
	return r
}

func (r BusinessesAPIGetBusinessesTokenRequest) Execute() (*BusinessCardHolderResponse, *http.Response, error) {
	return r.ApiService.GetBusinessesTokenExecute(r)
}

/*
GetBusinessesToken Retrieve business

To retrieve a specific business, send a `GET` request to the `/businesses/{token}` endpoint.
Include the business `token` path parameter to specify the business to return.

This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, sorting and pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the business resource.
 @return BusinessesAPIGetBusinessesTokenRequest
*/
func (a *BusinessesAPIService) GetBusinessesToken(ctx context.Context, token string) BusinessesAPIGetBusinessesTokenRequest {
	return BusinessesAPIGetBusinessesTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BusinessCardHolderResponse
func (a *BusinessesAPIService) GetBusinessesTokenExecute(r BusinessesAPIGetBusinessesTokenRequest) (*BusinessCardHolderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessCardHolderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.GetBusinessesToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses/{token}"
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

type BusinessesAPIGetBusinessesTokenSsnRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	token string
	fullSsn *bool
}

// To return the full identification number, set to &#x60;true&#x60;. To return only the last four digits, set to &#x60;false&#x60;. If the &#x60;proprietor_or_officer.identifications&#x60; array contains only the last four digits of the identification number, the &#x60;/businesses/{token}/ssn&#x60; endpoint will return only the last four digits regardless of the &#x60;full_ssn&#x60; parameter.
func (r BusinessesAPIGetBusinessesTokenSsnRequest) FullSsn(fullSsn bool) BusinessesAPIGetBusinessesTokenSsnRequest {
	r.fullSsn = &fullSsn
	return r
}

func (r BusinessesAPIGetBusinessesTokenSsnRequest) Execute() (*SsnResponseModel, *http.Response, error) {
	return r.ApiService.GetBusinessesTokenSsnExecute(r)
}

/*
GetBusinessesTokenSsn Retrieve business identification number

To retrieve the government-issued identification number of a business' proprietor, send a `GET` request to the `/businesses/{token}/ssn` endpoint.
Include the `token` path parameter to specify the business whose identification number (SSN, TIN, NIN, SIN) you want to return.
You can indicate whether to return the full number or the last four digits only.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the business resource.
 @return BusinessesAPIGetBusinessesTokenSsnRequest
*/
func (a *BusinessesAPIService) GetBusinessesTokenSsn(ctx context.Context, token string) BusinessesAPIGetBusinessesTokenSsnRequest {
	return BusinessesAPIGetBusinessesTokenSsnRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return SsnResponseModel
func (a *BusinessesAPIService) GetBusinessesTokenSsnExecute(r BusinessesAPIGetBusinessesTokenSsnRequest) (*SsnResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SsnResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.GetBusinessesTokenSsn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses/{token}/ssn"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fullSsn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "full_ssn", r.fullSsn, "")
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

type BusinessesAPIPostBusinessesRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	businessCardholder *BusinessCardholder
}

func (r BusinessesAPIPostBusinessesRequest) BusinessCardholder(businessCardholder BusinessCardholder) BusinessesAPIPostBusinessesRequest {
	r.businessCardholder = &businessCardholder
	return r
}

func (r BusinessesAPIPostBusinessesRequest) Execute() (*BusinessCardHolderResponse, *http.Response, error) {
	return r.ApiService.PostBusinessesExecute(r)
}

/*
PostBusinesses Create business

Create a business.
The initial status of a newly created business depends on the <</core-api/kyc-verification, Know Your Customer (KYC) requirements>> of the program or associated <</core-api/account-holder-groups, account holder group>>.

[cols="1,1,1,2"]
|===
| KYC Required | Initial Business State | Business Active on Creation | Business Limitations

| Always
| `UNVERIFIED`
| No
| Cannot load funds.

| Conditionally
| `LIMITED`
| No
| Restricted by rules in `accountholdergroups.pre_kyc_controls`.

| Never
| `ACTIVE`
| Required
| None.
|===

To change or track the history of a business' status, use the `/businesstransitions` endpoint.
For more information on status changes, see <<create_business_transition, Create Business Transition>>.

For information about configuring the required fields for KYC verification, see <</core-api/kyc-verification#_perform_kyc, Perform KYC>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BusinessesAPIPostBusinessesRequest
*/
func (a *BusinessesAPIService) PostBusinesses(ctx context.Context) BusinessesAPIPostBusinessesRequest {
	return BusinessesAPIPostBusinessesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BusinessCardHolderResponse
func (a *BusinessesAPIService) PostBusinessesExecute(r BusinessesAPIPostBusinessesRequest) (*BusinessCardHolderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessCardHolderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.PostBusinesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses"

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
	localVarPostBody = r.businessCardholder
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

type BusinessesAPIPostBusinessesLookupRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	dDARequest *DDARequest
}

func (r BusinessesAPIPostBusinessesLookupRequest) DDARequest(dDARequest DDARequest) BusinessesAPIPostBusinessesLookupRequest {
	r.dDARequest = &dDARequest
	return r
}

func (r BusinessesAPIPostBusinessesLookupRequest) Execute() (*BusinessCardholder, *http.Response, error) {
	return r.ApiService.PostBusinessesLookupExecute(r)
}

/*
PostBusinessesLookup Search businesses

To search for one or more businesses, send a `POST` request to the `/businesses/lookup` endpoint.
Include in the message body any parameters by which you want to query.
This endpoint supports <</core-api/field-filtering, field filtering>> and <</core-api/sorting-and-pagination, pagination>>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BusinessesAPIPostBusinessesLookupRequest
*/
func (a *BusinessesAPIService) PostBusinessesLookup(ctx context.Context) BusinessesAPIPostBusinessesLookupRequest {
	return BusinessesAPIPostBusinessesLookupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BusinessCardholder
func (a *BusinessesAPIService) PostBusinessesLookupExecute(r BusinessesAPIPostBusinessesLookupRequest) (*BusinessCardholder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessCardholder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.PostBusinessesLookup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses/lookup"

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
	localVarPostBody = r.dDARequest
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

type BusinessesAPIPutBusinessesTokenRequest struct {
	ctx context.Context
	ApiService *BusinessesAPIService
	token string
	businessCardHolderUpdate *BusinessCardHolderUpdate
}

// Business object
func (r BusinessesAPIPutBusinessesTokenRequest) BusinessCardHolderUpdate(businessCardHolderUpdate BusinessCardHolderUpdate) BusinessesAPIPutBusinessesTokenRequest {
	r.businessCardHolderUpdate = &businessCardHolderUpdate
	return r
}

func (r BusinessesAPIPutBusinessesTokenRequest) Execute() (*BusinessCardholder, *http.Response, error) {
	return r.ApiService.PutBusinessesTokenExecute(r)
}

/*
PutBusinessesToken Update business

To update a business, send a `PUT` request to `/businesses/{token}`.
Use the `token` path parameter to specify the business to update.
Include the business details to update in link:http://www.json.org/[JSON, window="_blank"] format in the body of the request.
Only values of parameters in the request are modified; all others are left unchanged.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Unique identifier of the business resource
 @return BusinessesAPIPutBusinessesTokenRequest
*/
func (a *BusinessesAPIService) PutBusinessesToken(ctx context.Context, token string) BusinessesAPIPutBusinessesTokenRequest {
	return BusinessesAPIPutBusinessesTokenRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return BusinessCardholder
func (a *BusinessesAPIService) PutBusinessesTokenExecute(r BusinessesAPIPutBusinessesTokenRequest) (*BusinessCardholder, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BusinessCardholder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BusinessesAPIService.PutBusinessesToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/businesses/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.businessCardHolderUpdate == nil {
		return localVarReturnValue, nil, reportError("businessCardHolderUpdate is required and must be specified")
	}

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
	localVarPostBody = r.businessCardHolderUpdate
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
