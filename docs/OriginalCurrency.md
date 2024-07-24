# OriginalCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **decimal.Decimal** | original amount | [optional] 
**Code** | Pointer to **string** | Currency code, such as EUR or USD. | [optional] 

## Methods

### NewOriginalCurrency

`func NewOriginalCurrency() *OriginalCurrency`

NewOriginalCurrency instantiates a new OriginalCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalCurrencyWithDefaults

`func NewOriginalCurrencyWithDefaults() *OriginalCurrency`

NewOriginalCurrencyWithDefaults instantiates a new OriginalCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OriginalCurrency) GetAmount() decimal.Decimal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OriginalCurrency) GetAmountOk() (*decimal.Decimal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OriginalCurrency) SetAmount(v decimal.Decimal)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OriginalCurrency) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCode

`func (o *OriginalCurrency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OriginalCurrency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OriginalCurrency) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OriginalCurrency) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


