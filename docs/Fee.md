# Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **decimal.Decimal** | Amount of the fee. | 
**CreatedTime** | **time.Time** | Date and time when the &#x60;fees&#x60; object was created, in UTC. | 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**LastModifiedTime** | **time.Time** | Date and time when the &#x60;fees&#x60; object was last modified, in UTC. | 
**Name** | **string** | Name of the fee. | 
**Tags** | Pointer to **string** | Descriptive metadata about the fee. | [optional] 
**Token** | **string** | Unique identifier of the &#x60;fees&#x60; object. | 

## Methods

### NewFee

`func NewFee(amount decimal.Decimal, createdTime time.Time, currencyCode string, lastModifiedTime time.Time, name string, token string, ) *Fee`

NewFee instantiates a new Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeWithDefaults

`func NewFeeWithDefaults() *Fee`

NewFeeWithDefaults instantiates a new Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Fee) GetAmount() decimal.Decimal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Fee) GetAmountOk() (*decimal.Decimal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Fee) SetAmount(v decimal.Decimal)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *Fee) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Fee) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Fee) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *Fee) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Fee) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Fee) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetLastModifiedTime

`func (o *Fee) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *Fee) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *Fee) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *Fee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Fee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Fee) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *Fee) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Fee) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Fee) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Fee) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *Fee) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Fee) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Fee) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


