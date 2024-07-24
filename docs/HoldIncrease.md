# HoldIncrease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Controls whether the &#x60;value&#x60; field represents a fixed amount or a percentage of the authorization amount. | [default to "AMOUNT"]
**Value** | **decimal.Decimal** | Specifies the amount of the automatic increase to the authorization amount.  The &#x60;type&#x60; field controls whether this amount is a fixed amount or a percentage. | 

## Methods

### NewHoldIncrease

`func NewHoldIncrease(type_ string, value decimal.Decimal, ) *HoldIncrease`

NewHoldIncrease instantiates a new HoldIncrease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldIncreaseWithDefaults

`func NewHoldIncreaseWithDefaults() *HoldIncrease`

NewHoldIncreaseWithDefaults instantiates a new HoldIncrease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HoldIncrease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HoldIncrease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HoldIncrease) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *HoldIncrease) GetValue() decimal.Decimal`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HoldIncrease) GetValueOk() (*decimal.Decimal, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HoldIncrease) SetValue(v decimal.Decimal)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


