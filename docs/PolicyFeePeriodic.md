# PolicyFeePeriodic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFromInterestCalc** | Pointer to **bool** | Whether the periodic fee is excluded from interest calculation. | [optional] 
**FeeAmount** | Pointer to **float32** | Amount of the fee. | [optional] 
**NumberOfDaysPostActivation** | Pointer to **float32** | Number of days after an account is activated that the initial fee is charged. For example, if the value in this field is &#x60;30&#x60;, then the initial fee is charged 30 days after an account is activated. | [optional] 

## Methods

### NewPolicyFeePeriodic

`func NewPolicyFeePeriodic() *PolicyFeePeriodic`

NewPolicyFeePeriodic instantiates a new PolicyFeePeriodic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFeePeriodicWithDefaults

`func NewPolicyFeePeriodicWithDefaults() *PolicyFeePeriodic`

NewPolicyFeePeriodicWithDefaults instantiates a new PolicyFeePeriodic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFromInterestCalc

`func (o *PolicyFeePeriodic) GetExcludeFromInterestCalc() bool`

GetExcludeFromInterestCalc returns the ExcludeFromInterestCalc field if non-nil, zero value otherwise.

### GetExcludeFromInterestCalcOk

`func (o *PolicyFeePeriodic) GetExcludeFromInterestCalcOk() (*bool, bool)`

GetExcludeFromInterestCalcOk returns a tuple with the ExcludeFromInterestCalc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromInterestCalc

`func (o *PolicyFeePeriodic) SetExcludeFromInterestCalc(v bool)`

SetExcludeFromInterestCalc sets ExcludeFromInterestCalc field to given value.

### HasExcludeFromInterestCalc

`func (o *PolicyFeePeriodic) HasExcludeFromInterestCalc() bool`

HasExcludeFromInterestCalc returns a boolean if a field has been set.

### GetFeeAmount

`func (o *PolicyFeePeriodic) GetFeeAmount() float32`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *PolicyFeePeriodic) GetFeeAmountOk() (*float32, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *PolicyFeePeriodic) SetFeeAmount(v float32)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *PolicyFeePeriodic) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetNumberOfDaysPostActivation

`func (o *PolicyFeePeriodic) GetNumberOfDaysPostActivation() float32`

GetNumberOfDaysPostActivation returns the NumberOfDaysPostActivation field if non-nil, zero value otherwise.

### GetNumberOfDaysPostActivationOk

`func (o *PolicyFeePeriodic) GetNumberOfDaysPostActivationOk() (*float32, bool)`

GetNumberOfDaysPostActivationOk returns a tuple with the NumberOfDaysPostActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDaysPostActivation

`func (o *PolicyFeePeriodic) SetNumberOfDaysPostActivation(v float32)`

SetNumberOfDaysPostActivation sets NumberOfDaysPostActivation field to given value.

### HasNumberOfDaysPostActivation

`func (o *PolicyFeePeriodic) HasNumberOfDaysPostActivation() bool`

HasNumberOfDaysPostActivation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


