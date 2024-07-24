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
	"time"
)


// RedemptionsAPIService RedemptionsAPI service
type RedemptionsAPIService service

type RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest struct {
	ctx context.Context
	ApiService *RedemptionsAPIService
	settlementStartDate *time.Time
	settlementEndDate *time.Time
	destination *DestinationType
	count *int32
	startIndex *int64
}

// Settlement start date to filter by.
func (r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) SettlementStartDate(settlementStartDate time.Time) RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest {
	r.settlementStartDate = &settlementStartDate
	return r
}

// Settlement end date to filter by.
func (r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) SettlementEndDate(settlementEndDate time.Time) RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest {
	r.settlementEndDate = &settlementEndDate
	return r
}

// Specifies the destination for external redemptions to filter for.
func (r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) Destination(destination DestinationType) RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest {
	r.destination = &destination
	return r
}

// Number of resources to retrieve.
func (r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) Count(count int32) RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest {
	r.count = &count
	return r
}

// Sort order index of the first resource in the returned array.
func (r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) StartIndex(startIndex int64) RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest {
	r.startIndex = &startIndex
	return r
}

func (r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) Execute() (*RedemptionsBySettlementDatePage, *http.Response, error) {
	return r.ApiService.RetrieveRedemptionsBySettlementDateExecute(r)
}

/*
RetrieveRedemptionsBySettlementDate retrieves all completed redemptions by settlement date

retrieves all completed redemptions by settlement date

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest
*/
func (a *RedemptionsAPIService) RetrieveRedemptionsBySettlementDate(ctx context.Context) RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest {
	return RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RedemptionsBySettlementDatePage
func (a *RedemptionsAPIService) RetrieveRedemptionsBySettlementDateExecute(r RedemptionsAPIRetrieveRedemptionsBySettlementDateRequest) (*RedemptionsBySettlementDatePage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RedemptionsBySettlementDatePage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RedemptionsAPIService.RetrieveRedemptionsBySettlementDate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rewardprograms/redemptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.settlementStartDate == nil {
		return localVarReturnValue, nil, reportError("settlementStartDate is required and must be specified")
	}
	if r.settlementEndDate == nil {
		return localVarReturnValue, nil, reportError("settlementEndDate is required and must be specified")
	}

	if r.destination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "destination", r.destination, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "settlement_start_date", r.settlementStartDate, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "settlement_end_date", r.settlementEndDate, "")
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 5
		r.count = &defaultValue
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_index", r.startIndex, "")
	} else {
		var defaultValue int64 = 0
		r.startIndex = &defaultValue
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
