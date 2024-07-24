# AccountNameVerificationSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First or given name of the cardholder. | [optional] 
**LastName** | Pointer to **string** | Last or family name of the cardholder. | [optional] 
**MiddleName** | Pointer to **string** | Middle name of the cardholder. | [optional] 
**OnFile** | Pointer to [**AniInformation**](AniInformation.md) |  | [optional] 
**Response** | Pointer to [**Response**](Response.md) |  | [optional] 

## Methods

### NewAccountNameVerificationSource

`func NewAccountNameVerificationSource() *AccountNameVerificationSource`

NewAccountNameVerificationSource instantiates a new AccountNameVerificationSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountNameVerificationSourceWithDefaults

`func NewAccountNameVerificationSourceWithDefaults() *AccountNameVerificationSource`

NewAccountNameVerificationSourceWithDefaults instantiates a new AccountNameVerificationSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *AccountNameVerificationSource) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountNameVerificationSource) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountNameVerificationSource) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountNameVerificationSource) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AccountNameVerificationSource) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountNameVerificationSource) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountNameVerificationSource) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountNameVerificationSource) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *AccountNameVerificationSource) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *AccountNameVerificationSource) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *AccountNameVerificationSource) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *AccountNameVerificationSource) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetOnFile

`func (o *AccountNameVerificationSource) GetOnFile() AniInformation`

GetOnFile returns the OnFile field if non-nil, zero value otherwise.

### GetOnFileOk

`func (o *AccountNameVerificationSource) GetOnFileOk() (*AniInformation, bool)`

GetOnFileOk returns a tuple with the OnFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFile

`func (o *AccountNameVerificationSource) SetOnFile(v AniInformation)`

SetOnFile sets OnFile field to given value.

### HasOnFile

`func (o *AccountNameVerificationSource) HasOnFile() bool`

HasOnFile returns a boolean if a field has been set.

### GetResponse

`func (o *AccountNameVerificationSource) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AccountNameVerificationSource) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AccountNameVerificationSource) SetResponse(v Response)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AccountNameVerificationSource) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


