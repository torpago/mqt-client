# VelocityControlBalanceAllOfAvailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of money remaining for the user. This value is returned only if the user has a limit on the amount of money. | [optional] 
**DaysRemaining** | Pointer to **int32** | Number of days remaining for the user. This value is returned only if the user has a limit on the number of days. | [optional] 
**Uses** | Pointer to **int32** | Number of uses remaining for the user. This value is returned only if the user has a limit on the number of uses. | [optional] 

## Methods

### NewVelocityControlBalanceAllOfAvailable

`func NewVelocityControlBalanceAllOfAvailable() *VelocityControlBalanceAllOfAvailable`

NewVelocityControlBalanceAllOfAvailable instantiates a new VelocityControlBalanceAllOfAvailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityControlBalanceAllOfAvailableWithDefaults

`func NewVelocityControlBalanceAllOfAvailableWithDefaults() *VelocityControlBalanceAllOfAvailable`

NewVelocityControlBalanceAllOfAvailableWithDefaults instantiates a new VelocityControlBalanceAllOfAvailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VelocityControlBalanceAllOfAvailable) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VelocityControlBalanceAllOfAvailable) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VelocityControlBalanceAllOfAvailable) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VelocityControlBalanceAllOfAvailable) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDaysRemaining

`func (o *VelocityControlBalanceAllOfAvailable) GetDaysRemaining() int32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *VelocityControlBalanceAllOfAvailable) GetDaysRemainingOk() (*int32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *VelocityControlBalanceAllOfAvailable) SetDaysRemaining(v int32)`

SetDaysRemaining sets DaysRemaining field to given value.

### HasDaysRemaining

`func (o *VelocityControlBalanceAllOfAvailable) HasDaysRemaining() bool`

HasDaysRemaining returns a boolean if a field has been set.

### GetUses

`func (o *VelocityControlBalanceAllOfAvailable) GetUses() int32`

GetUses returns the Uses field if non-nil, zero value otherwise.

### GetUsesOk

`func (o *VelocityControlBalanceAllOfAvailable) GetUsesOk() (*int32, bool)`

GetUsesOk returns a tuple with the Uses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUses

`func (o *VelocityControlBalanceAllOfAvailable) SetUses(v int32)`

SetUses sets Uses field to given value.

### HasUses

`func (o *VelocityControlBalanceAllOfAvailable) HasUses() bool`

HasUses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


