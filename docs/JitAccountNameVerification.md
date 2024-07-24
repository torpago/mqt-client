# JitAccountNameVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to [**AccountNameVerificationSource**](AccountNameVerificationSource.md) |  | [optional] 
**Issuer** | Pointer to [**AccountNameVerificationSource**](AccountNameVerificationSource.md) |  | [optional] 
**Request** | Pointer to [**AniInformation1**](AniInformation1.md) |  | [optional] 

## Methods

### NewJitAccountNameVerification

`func NewJitAccountNameVerification() *JitAccountNameVerification`

NewJitAccountNameVerification instantiates a new JitAccountNameVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitAccountNameVerificationWithDefaults

`func NewJitAccountNameVerificationWithDefaults() *JitAccountNameVerification`

NewJitAccountNameVerificationWithDefaults instantiates a new JitAccountNameVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *JitAccountNameVerification) GetGateway() AccountNameVerificationSource`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *JitAccountNameVerification) GetGatewayOk() (*AccountNameVerificationSource, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *JitAccountNameVerification) SetGateway(v AccountNameVerificationSource)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *JitAccountNameVerification) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIssuer

`func (o *JitAccountNameVerification) GetIssuer() AccountNameVerificationSource`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *JitAccountNameVerification) GetIssuerOk() (*AccountNameVerificationSource, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *JitAccountNameVerification) SetIssuer(v AccountNameVerificationSource)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *JitAccountNameVerification) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetRequest

`func (o *JitAccountNameVerification) GetRequest() AniInformation1`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *JitAccountNameVerification) GetRequestOk() (*AniInformation1, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *JitAccountNameVerification) SetRequest(v AniInformation1)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *JitAccountNameVerification) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


