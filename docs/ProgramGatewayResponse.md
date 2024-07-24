# ProgramGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the Program Gateway is active. | [optional] [default to true]
**BasicAuthPassword** | Pointer to **string** | Basic Authentication password for authenticating your environment. | [optional] 
**BasicAuthUsername** | Pointer to **string** | Basic Authentication username for authenticating your environment. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the Program Gateway was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**CustomHeader** | Pointer to **map[string]interface{}** | Additional custom information included in the HTTP header. | [optional] 
**Mtls** | Pointer to **bool** | Indicates whether the Program Gateway uses mutual Transport Layer Security (mTLS). | [optional] 
**Name** | Pointer to **string** | Name of the Program Gateway. | [optional] 
**TimeoutMillis** | Pointer to **int32** | Total timeout for Program Gateway calls, in milliseconds. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the Program Gateway. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the Program Gateway was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Url** | Pointer to **string** | URL of the Program Gateway endpoint hosted in your environment and configured to receive authorization requests made by the Marqeta platform. | [optional] 

## Methods

### NewProgramGatewayResponse

`func NewProgramGatewayResponse() *ProgramGatewayResponse`

NewProgramGatewayResponse instantiates a new ProgramGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramGatewayResponseWithDefaults

`func NewProgramGatewayResponseWithDefaults() *ProgramGatewayResponse`

NewProgramGatewayResponseWithDefaults instantiates a new ProgramGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProgramGatewayResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProgramGatewayResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProgramGatewayResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ProgramGatewayResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *ProgramGatewayResponse) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *ProgramGatewayResponse) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *ProgramGatewayResponse) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *ProgramGatewayResponse) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.

### GetBasicAuthUsername

`func (o *ProgramGatewayResponse) GetBasicAuthUsername() string`

GetBasicAuthUsername returns the BasicAuthUsername field if non-nil, zero value otherwise.

### GetBasicAuthUsernameOk

`func (o *ProgramGatewayResponse) GetBasicAuthUsernameOk() (*string, bool)`

GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUsername

`func (o *ProgramGatewayResponse) SetBasicAuthUsername(v string)`

SetBasicAuthUsername sets BasicAuthUsername field to given value.

### HasBasicAuthUsername

`func (o *ProgramGatewayResponse) HasBasicAuthUsername() bool`

HasBasicAuthUsername returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ProgramGatewayResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProgramGatewayResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProgramGatewayResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ProgramGatewayResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomHeader

`func (o *ProgramGatewayResponse) GetCustomHeader() map[string]interface{}`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *ProgramGatewayResponse) GetCustomHeaderOk() (*map[string]interface{}, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *ProgramGatewayResponse) SetCustomHeader(v map[string]interface{})`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *ProgramGatewayResponse) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.

### GetMtls

`func (o *ProgramGatewayResponse) GetMtls() bool`

GetMtls returns the Mtls field if non-nil, zero value otherwise.

### GetMtlsOk

`func (o *ProgramGatewayResponse) GetMtlsOk() (*bool, bool)`

GetMtlsOk returns a tuple with the Mtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtls

`func (o *ProgramGatewayResponse) SetMtls(v bool)`

SetMtls sets Mtls field to given value.

### HasMtls

`func (o *ProgramGatewayResponse) HasMtls() bool`

HasMtls returns a boolean if a field has been set.

### GetName

`func (o *ProgramGatewayResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProgramGatewayResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProgramGatewayResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProgramGatewayResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *ProgramGatewayResponse) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *ProgramGatewayResponse) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *ProgramGatewayResponse) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *ProgramGatewayResponse) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetToken

`func (o *ProgramGatewayResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramGatewayResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramGatewayResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProgramGatewayResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ProgramGatewayResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ProgramGatewayResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ProgramGatewayResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ProgramGatewayResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUrl

`func (o *ProgramGatewayResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProgramGatewayResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProgramGatewayResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProgramGatewayResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


