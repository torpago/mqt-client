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
)


// FeeRefundsAPIService FeeRefundsAPI service
type FeeRefundsAPIService service

type FeeRefundsAPIPostFeeRefundsRequest struct {
	ctx context.Context
	ApiService *FeeRefundsAPIService
	feeRefundRequest *FeeRefundRequest
}

func (r FeeRefundsAPIPostFeeRefundsRequest) FeeRefundRequest(feeRefundRequest FeeRefundRequest) FeeRefundsAPIPostFeeRefundsRequest {
	r.feeRefundRequest = &feeRefundRequest
	return r
}

func (r FeeRefundsAPIPostFeeRefundsRequest) Execute() (*FeeTransferResponse, *http.Response, error) {
	return r.ApiService.PostFeeRefundsExecute(r)
}

/*
PostFeeRefunds Create fee refund

Use this endpoint to create a fee refund.
Include the fee charge `token` path parameter to specify the fee to return.

If there are insufficient funds in your fee account to refund the fee, funds are drawn from your program reserve account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FeeRefundsAPIPostFeeRefundsRequest
*/
func (a *FeeRefundsAPIService) PostFeeRefunds(ctx context.Context) FeeRefundsAPIPostFeeRefundsRequest {
	return FeeRefundsAPIPostFeeRefundsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FeeTransferResponse
func (a *FeeRefundsAPIService) PostFeeRefundsExecute(r FeeRefundsAPIPostFeeRefundsRequest) (*FeeTransferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FeeTransferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeeRefundsAPIService.PostFeeRefunds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feerefunds"

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
	localVarPostBody = r.feeRefundRequest
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
