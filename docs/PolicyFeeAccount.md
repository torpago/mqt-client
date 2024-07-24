# PolicyFeeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualFee** | Pointer to [**PolicyFeePeriodic**](PolicyFeePeriodic.md) |  | [optional] 
**ForeignTransactionFee** | Pointer to [**PolicyFeeForeignTransaction**](PolicyFeeForeignTransaction.md) |  | [optional] 
**LatePayment** | Pointer to [**PolicyFeePayment**](PolicyFeePayment.md) |  | [optional] 
**MonthlyFee** | Pointer to [**PolicyFeePeriodic**](PolicyFeePeriodic.md) |  | [optional] 
**ReturnedPayment** | Pointer to [**PolicyFeePayment**](PolicyFeePayment.md) |  | [optional] 

## Methods

### NewPolicyFeeAccount

`func NewPolicyFeeAccount() *PolicyFeeAccount`

NewPolicyFeeAccount instantiates a new PolicyFeeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFeeAccountWithDefaults

`func NewPolicyFeeAccountWithDefaults() *PolicyFeeAccount`

NewPolicyFeeAccountWithDefaults instantiates a new PolicyFeeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualFee

`func (o *PolicyFeeAccount) GetAnnualFee() PolicyFeePeriodic`

GetAnnualFee returns the AnnualFee field if non-nil, zero value otherwise.

### GetAnnualFeeOk

`func (o *PolicyFeeAccount) GetAnnualFeeOk() (*PolicyFeePeriodic, bool)`

GetAnnualFeeOk returns a tuple with the AnnualFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualFee

`func (o *PolicyFeeAccount) SetAnnualFee(v PolicyFeePeriodic)`

SetAnnualFee sets AnnualFee field to given value.

### HasAnnualFee

`func (o *PolicyFeeAccount) HasAnnualFee() bool`

HasAnnualFee returns a boolean if a field has been set.

### GetForeignTransactionFee

`func (o *PolicyFeeAccount) GetForeignTransactionFee() PolicyFeeForeignTransaction`

GetForeignTransactionFee returns the ForeignTransactionFee field if non-nil, zero value otherwise.

### GetForeignTransactionFeeOk

`func (o *PolicyFeeAccount) GetForeignTransactionFeeOk() (*PolicyFeeForeignTransaction, bool)`

GetForeignTransactionFeeOk returns a tuple with the ForeignTransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTransactionFee

`func (o *PolicyFeeAccount) SetForeignTransactionFee(v PolicyFeeForeignTransaction)`

SetForeignTransactionFee sets ForeignTransactionFee field to given value.

### HasForeignTransactionFee

`func (o *PolicyFeeAccount) HasForeignTransactionFee() bool`

HasForeignTransactionFee returns a boolean if a field has been set.

### GetLatePayment

`func (o *PolicyFeeAccount) GetLatePayment() PolicyFeePayment`

GetLatePayment returns the LatePayment field if non-nil, zero value otherwise.

### GetLatePaymentOk

`func (o *PolicyFeeAccount) GetLatePaymentOk() (*PolicyFeePayment, bool)`

GetLatePaymentOk returns a tuple with the LatePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatePayment

`func (o *PolicyFeeAccount) SetLatePayment(v PolicyFeePayment)`

SetLatePayment sets LatePayment field to given value.

### HasLatePayment

`func (o *PolicyFeeAccount) HasLatePayment() bool`

HasLatePayment returns a boolean if a field has been set.

### GetMonthlyFee

`func (o *PolicyFeeAccount) GetMonthlyFee() PolicyFeePeriodic`

GetMonthlyFee returns the MonthlyFee field if non-nil, zero value otherwise.

### GetMonthlyFeeOk

`func (o *PolicyFeeAccount) GetMonthlyFeeOk() (*PolicyFeePeriodic, bool)`

GetMonthlyFeeOk returns a tuple with the MonthlyFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyFee

`func (o *PolicyFeeAccount) SetMonthlyFee(v PolicyFeePeriodic)`

SetMonthlyFee sets MonthlyFee field to given value.

### HasMonthlyFee

`func (o *PolicyFeeAccount) HasMonthlyFee() bool`

HasMonthlyFee returns a boolean if a field has been set.

### GetReturnedPayment

`func (o *PolicyFeeAccount) GetReturnedPayment() PolicyFeePayment`

GetReturnedPayment returns the ReturnedPayment field if non-nil, zero value otherwise.

### GetReturnedPaymentOk

`func (o *PolicyFeeAccount) GetReturnedPaymentOk() (*PolicyFeePayment, bool)`

GetReturnedPaymentOk returns a tuple with the ReturnedPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedPayment

`func (o *PolicyFeeAccount) SetReturnedPayment(v PolicyFeePayment)`

SetReturnedPayment sets ReturnedPayment field to given value.

### HasReturnedPayment

`func (o *PolicyFeeAccount) HasReturnedPayment() bool`

HasReturnedPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


