# PolicyFeePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMethod** | Pointer to **string** | Method used to calculate the fee value. | [optional] 
**DefaultValue** | Pointer to **decimal.Decimal** | Amount of the fee. | [optional] 

## Methods

### NewPolicyFeePayment

`func NewPolicyFeePayment() *PolicyFeePayment`

NewPolicyFeePayment instantiates a new PolicyFeePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFeePaymentWithDefaults

`func NewPolicyFeePaymentWithDefaults() *PolicyFeePayment`

NewPolicyFeePaymentWithDefaults instantiates a new PolicyFeePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMethod

`func (o *PolicyFeePayment) GetDefaultMethod() string`

GetDefaultMethod returns the DefaultMethod field if non-nil, zero value otherwise.

### GetDefaultMethodOk

`func (o *PolicyFeePayment) GetDefaultMethodOk() (*string, bool)`

GetDefaultMethodOk returns a tuple with the DefaultMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMethod

`func (o *PolicyFeePayment) SetDefaultMethod(v string)`

SetDefaultMethod sets DefaultMethod field to given value.

### HasDefaultMethod

`func (o *PolicyFeePayment) HasDefaultMethod() bool`

HasDefaultMethod returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PolicyFeePayment) GetDefaultValue() decimal.Decimal`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PolicyFeePayment) GetDefaultValueOk() (*decimal.Decimal, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PolicyFeePayment) SetDefaultValue(v decimal.Decimal)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PolicyFeePayment) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


