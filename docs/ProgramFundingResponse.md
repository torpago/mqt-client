# ProgramFundingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the funding. | 
**CreatedTime** | **time.Time** | Timestamp when the funding entry was created. | 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**Memo** | **string** | Additional notes for the funding entry. | 
**PostTime** | **time.Time** | Timestamp when the funding entry was posted. | 
**ShortCode** | **string** | Unique identifier of the program. | 
**Token** | **string** | Unique identifier of the funding entry. | 
**UpdatedTime** | **time.Time** | Timestamp when the funding entry was last updated. | 

## Methods

### NewProgramFundingResponse

`func NewProgramFundingResponse(amount float32, createdTime time.Time, currencyCode CurrencyCode, memo string, postTime time.Time, shortCode string, token string, updatedTime time.Time, ) *ProgramFundingResponse`

NewProgramFundingResponse instantiates a new ProgramFundingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramFundingResponseWithDefaults

`func NewProgramFundingResponseWithDefaults() *ProgramFundingResponse`

NewProgramFundingResponseWithDefaults instantiates a new ProgramFundingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ProgramFundingResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProgramFundingResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProgramFundingResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCreatedTime

`func (o *ProgramFundingResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ProgramFundingResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ProgramFundingResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *ProgramFundingResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProgramFundingResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProgramFundingResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetMemo

`func (o *ProgramFundingResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ProgramFundingResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ProgramFundingResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPostTime

`func (o *ProgramFundingResponse) GetPostTime() time.Time`

GetPostTime returns the PostTime field if non-nil, zero value otherwise.

### GetPostTimeOk

`func (o *ProgramFundingResponse) GetPostTimeOk() (*time.Time, bool)`

GetPostTimeOk returns a tuple with the PostTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTime

`func (o *ProgramFundingResponse) SetPostTime(v time.Time)`

SetPostTime sets PostTime field to given value.


### GetShortCode

`func (o *ProgramFundingResponse) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *ProgramFundingResponse) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *ProgramFundingResponse) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.


### GetToken

`func (o *ProgramFundingResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProgramFundingResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProgramFundingResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *ProgramFundingResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ProgramFundingResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ProgramFundingResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


