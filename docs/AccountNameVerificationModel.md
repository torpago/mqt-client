# AccountNameVerificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnFile** | Pointer to [**AniInformation**](AniInformation.md) |  | [optional] 
**Request** | Pointer to [**AniInformation**](AniInformation.md) |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**Response** | Pointer to [**Response**](Response.md) |  | [optional] 

## Methods

### NewAccountNameVerificationModel

`func NewAccountNameVerificationModel() *AccountNameVerificationModel`

NewAccountNameVerificationModel instantiates a new AccountNameVerificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountNameVerificationModelWithDefaults

`func NewAccountNameVerificationModelWithDefaults() *AccountNameVerificationModel`

NewAccountNameVerificationModelWithDefaults instantiates a new AccountNameVerificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnFile

`func (o *AccountNameVerificationModel) GetOnFile() AniInformation`

GetOnFile returns the OnFile field if non-nil, zero value otherwise.

### GetOnFileOk

`func (o *AccountNameVerificationModel) GetOnFileOk() (*AniInformation, bool)`

GetOnFileOk returns a tuple with the OnFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFile

`func (o *AccountNameVerificationModel) SetOnFile(v AniInformation)`

SetOnFile sets OnFile field to given value.

### HasOnFile

`func (o *AccountNameVerificationModel) HasOnFile() bool`

HasOnFile returns a boolean if a field has been set.

### GetRequest

`func (o *AccountNameVerificationModel) GetRequest() AniInformation`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AccountNameVerificationModel) GetRequestOk() (*AniInformation, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AccountNameVerificationModel) SetRequest(v AniInformation)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AccountNameVerificationModel) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetRequestType

`func (o *AccountNameVerificationModel) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *AccountNameVerificationModel) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *AccountNameVerificationModel) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *AccountNameVerificationModel) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetResponse

`func (o *AccountNameVerificationModel) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AccountNameVerificationModel) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AccountNameVerificationModel) SetResponse(v Response)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AccountNameVerificationModel) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


