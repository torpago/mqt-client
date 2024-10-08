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


// BalanceRefundsAPIService BalanceRefundsAPI service
type BalanceRefundsAPIService service

type BalanceRefundsAPICreateBalanceRefundRequest struct {
	ctx context.Context
	ApiService *BalanceRefundsAPIService
	accountToken string
	accountCreditBalanceRefundReq *AccountCreditBalanceRefundReq
}

func (r BalanceRefundsAPICreateBalanceRefundRequest) AccountCreditBalanceRefundReq(accountCreditBalanceRefundReq AccountCreditBalanceRefundReq) BalanceRefundsAPICreateBalanceRefundRequest {
	r.accountCreditBalanceRefundReq = &accountCreditBalanceRefundReq
	return r
}

func (r BalanceRefundsAPICreateBalanceRefundRequest) Execute() (*AccountCreditBalanceRefundResponse, *http.Response, error) {
	return r.ApiService.CreateBalanceRefundExecute(r)
}

/*
CreateBalanceRefund Create balance refund

Create a new balance refund, which can be issued to the account holder if their credit account balance is negative.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountToken Unique identifier of the credit account for which you want to create a balance refund.  Send a `GET` request to `/credit/accounts` to retrieve existing credit account tokens.
 @return BalanceRefundsAPICreateBalanceRefundRequest
*/
func (a *BalanceRefundsAPIService) CreateBalanceRefund(ctx context.Context, accountToken string) BalanceRefundsAPICreateBalanceRefundRequest {
	return BalanceRefundsAPICreateBalanceRefundRequest{
		ApiService: a,
		ctx: ctx,
		accountToken: accountToken,
	}
}

// Execute executes the request
//  @return AccountCreditBalanceRefundResponse
func (a *BalanceRefundsAPIService) CreateBalanceRefundExecute(r BalanceRefundsAPICreateBalanceRefundRequest) (*AccountCreditBalanceRefundResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountCreditBalanceRefundResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalanceRefundsAPIService.CreateBalanceRefund")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_token}/creditbalancerefunds"
	localVarPath = strings.Replace(localVarPath, "{"+"account_token"+"}", url.PathEscape(parameterValueToString(r.accountToken, "accountToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountCreditBalanceRefundReq == nil {
		return localVarReturnValue, nil, reportError("accountCreditBalanceRefundReq is required and must be specified")
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
	localVarPostBody = r.accountCreditBalanceRefundReq
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
