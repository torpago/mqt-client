# RequestForApplePayWppJWT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardToken** | **string** | Unique identifier of the card resource. | 
**ReqSysId** | Pointer to **string** | Random pseudo-unique value used for troubleshooting between multiple parties. | [optional] 

## Methods

### NewRequestForApplePayWppJWT

`func NewRequestForApplePayWppJWT(cardToken string, ) *RequestForApplePayWppJWT`

NewRequestForApplePayWppJWT instantiates a new RequestForApplePayWppJWT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestForApplePayWppJWTWithDefaults

`func NewRequestForApplePayWppJWTWithDefaults() *RequestForApplePayWppJWT`

NewRequestForApplePayWppJWTWithDefaults instantiates a new RequestForApplePayWppJWT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardToken

`func (o *RequestForApplePayWppJWT) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *RequestForApplePayWppJWT) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *RequestForApplePayWppJWT) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetReqSysId

`func (o *RequestForApplePayWppJWT) GetReqSysId() string`

GetReqSysId returns the ReqSysId field if non-nil, zero value otherwise.

### GetReqSysIdOk

`func (o *RequestForApplePayWppJWT) GetReqSysIdOk() (*string, bool)`

GetReqSysIdOk returns a tuple with the ReqSysId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSysId

`func (o *RequestForApplePayWppJWT) SetReqSysId(v string)`

SetReqSysId sets ReqSysId field to given value.

### HasReqSysId

`func (o *RequestForApplePayWppJWT) HasReqSysId() bool`

HasReqSysId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


