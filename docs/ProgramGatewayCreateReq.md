# ProgramGatewayCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the Program Gateway is active. | [optional] [default to true]
**BasicAuthPassword** | **string** | Basic Authentication password for authenticating your environment. | 
**BasicAuthUsername** | **string** | Basic Authentication username for authenticating your environment. | 
**CustomHeader** | Pointer to **map[string]string** | Additional custom information included in the HTTP header. For example, this might contain security information, along with Basic Authentication, when making a Program Gateway request. Custom headers also appear in the associated webhook&#39;s notifications. | [optional] 
**Mtls** | Pointer to **bool** | Indicates whether the Program Gateway uses mutual Transport Layer Security (mTLS). | [optional] [default to false]
**Name** | **string** | Name of the Program Gateway. | 
**TimeoutMillis** | Pointer to **int32** | Total timeout for Program Gateway calls, in milliseconds. | [optional] [default to 2000]
**Token** | Pointer to **string** | Unique identifier of the Program Gateway.  If you do not include a token, the system generates one automatically. As this token is necessary for use in other calls, it is recommended that you define a simple and easy to remember string rather than letting the system generate a token for you. This value cannot be updated. | [optional] 
**Url** | **string** | URL of the Program Gateway endpoint hosted in your environment and configured to receive authorization requests made by the Marqeta platform. Must be HTTPS. | 

## Methods

### NewProgramGatewayCreateReq

`func NewProgramGatewayCreateReq(basicAuthPassword string, basicAuthUsername string, name string, url string, ) *ProgramGatewayCreateReq`

NewProgramGatewayCreateReq instantiates a new ProgramGatewayCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramGatewayCreateReqWithDefaults

`func NewProgramGatewayCreateReqWithDefaults() *ProgramGatewayCreateReq`

NewProgramGatewayCreateReqWithDefaults instantiates a new ProgramGatewayCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProgramGatewayCreateReq) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProgramGatewayCreateReq) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProgramGatewayCreateReq) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProgramGatewayCreateReq) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *ProgramGatewayCreateReq) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *ProgramGatewayCreateReq) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *ProgramGatewayCreateReq) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### GetBasicAuthUsername

`func (o *ProgramGatewayCreateReq) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *ProgramGatewayCreateReq) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *ProgramGatewayCreateReq) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.


### GetCustomHeader

`func (o *ProgramGatewayCreateReq) GetCustomHeader() map[string]string`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *ProgramGatewayCreateReq) GetCustomHeaderOk() (*map[string]string, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *ProgramGatewayCreateReq) SetCustomHeader(v map[string]string)`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *ProgramGatewayCreateReq) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.

### GetMtls

`func (o *ProgramGatewayCreateReq) GetMtls() bool`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *ProgramGatewayCreateReq) GetMtlsOk() (*bool, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *ProgramGatewayCreateReq) SetMtls(v bool)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *ProgramGatewayCreateReq) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetName

`func (o *ProgramGatewayCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramGatewayCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramGatewayCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutMillis

`func (o *ProgramGatewayCreateReq) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *ProgramGatewayCreateReq) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *ProgramGatewayCreateReq) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *ProgramGatewayCreateReq) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetToken

`func (o *ProgramGatewayCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramGatewayCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramGatewayCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramGatewayCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUrl

`func (o *ProgramGatewayCreateReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProgramGatewayCreateReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProgramGatewayCreateReq) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


