# ProductConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleDay** | **int32** | Day of month the billing cycle starts. | 
**BillingCycleDayStrategy** | Pointer to **string** | Determines if the billing cycle day is manually set or determined dynamically during account creation based on cycling logic. | [optional] [default to "MANUAL"]
**BillingCycleFrequency** | Pointer to **string** | Frequency at which your account is billed. | [optional] [default to "MONTHLY"]
**Fees** | Pointer to [**[]ProductFeeType**](ProductFeeType.md) | One or more fee types. | [optional] 
**PaymentDueDay** | Pointer to **int32** | Day of month the payment for the previous billing cycle is due.  This field is deprecated. Use the &#x60;product.payment_due_interval&#x60; field instead. To retrieve &#x60;payment_due_interval&#x60;, see &lt;&lt;/core-api/credit-products#retrieveProduct, Retrieve credit product, config.payment_due_interval&gt;&gt;. | [optional] 
**PaymentDueInterval** | Pointer to **int32** | Specifies the payment due interval that is used to determine the payment due date for a billing cycle. The accepted values are either -1 or a value between 1 and 26. A value of -1 indicates one day prior to the next billing cycle date. | [optional] [default to -1]
**PeriodicFees** | Pointer to [**[]ProductConfigPeriodicFeesInner**](ProductConfigPeriodicFeesInner.md) | Contains one or more periodic fees. | [optional] 

## Methods

### NewProductConfig

`func NewProductConfig(billingCycleDay int32, ) *ProductConfig`

NewProductConfig instantiates a new ProductConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductConfigWithDefaults

`func NewProductConfigWithDefaults() *ProductConfig`

NewProductConfigWithDefaults instantiates a new ProductConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleDay

`func (o *ProductConfig) GetBillingCycleDay() int32`

GetBillingCycleDay returns the BillingCycleDay field if non-nil, zero value otherwise.

### GetBillingCycleDayOk

`func (o *ProductConfig) GetBillingCycleDayOk() (*int32, bool)`

GetBillingCycleDayOk returns a tuple with the BillingCycleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDay

`func (o *ProductConfig) SetBillingCycleDay(v int32)`

SetBillingCycleDay sets BillingCycleDay field to given value.


### GetBillingCycleDayStrategy

`func (o *ProductConfig) GetBillingCycleDayStrategy() string`

GetBillingCycleDayStrategy returns the BillingCycleDayStrategy field if non-nil, zero value otherwise.

### GetBillingCycleDayStrategyOk

`func (o *ProductConfig) GetBillingCycleDayStrategyOk() (*string, bool)`

GetBillingCycleDayStrategyOk returns a tuple with the BillingCycleDayStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDayStrategy

`func (o *ProductConfig) SetBillingCycleDayStrategy(v string)`

SetBillingCycleDayStrategy sets BillingCycleDayStrategy field to given value.

### HasBillingCycleDayStrategy

`func (o *ProductConfig) HasBillingCycleDayStrategy() bool`

HasBillingCycleDayStrategy returns a boolean if a field has been set.

### GetBillingCycleFrequency

`func (o *ProductConfig) GetBillingCycleFrequency() string`

GetBillingCycleFrequency returns the BillingCycleFrequency field if non-nil, zero value otherwise.

### GetBillingCycleFrequencyOk

`func (o *ProductConfig) GetBillingCycleFrequencyOk() (*string, bool)`

GetBillingCycleFrequencyOk returns a tuple with the BillingCycleFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleFrequency

`func (o *ProductConfig) SetBillingCycleFrequency(v string)`

SetBillingCycleFrequency sets BillingCycleFrequency field to given value.

### HasBillingCycleFrequency

`func (o *ProductConfig) HasBillingCycleFrequency() bool`

HasBillingCycleFrequency returns a boolean if a field has been set.

### GetFees

`func (o *ProductConfig) GetFees() []ProductFeeType`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ProductConfig) GetFeesOk() (*[]ProductFeeType, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ProductConfig) SetFees(v []ProductFeeType)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ProductConfig) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetPaymentDueDay

`func (o *ProductConfig) GetPaymentDueDay() int32`

GetPaymentDueDay returns the PaymentDueDay field if non-nil, zero value otherwise.

### GetPaymentDueDayOk

`func (o *ProductConfig) GetPaymentDueDayOk() (*int32, bool)`

GetPaymentDueDayOk returns a tuple with the PaymentDueDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDay

`func (o *ProductConfig) SetPaymentDueDay(v int32)`

SetPaymentDueDay sets PaymentDueDay field to given value.

### HasPaymentDueDay

`func (o *ProductConfig) HasPaymentDueDay() bool`

HasPaymentDueDay returns a boolean if a field has been set.

### GetPaymentDueInterval

`func (o *ProductConfig) GetPaymentDueInterval() int32`

GetPaymentDueInterval returns the PaymentDueInterval field if non-nil, zero value otherwise.

### GetPaymentDueIntervalOk

`func (o *ProductConfig) GetPaymentDueIntervalOk() (*int32, bool)`

GetPaymentDueIntervalOk returns a tuple with the PaymentDueInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueInterval

`func (o *ProductConfig) SetPaymentDueInterval(v int32)`

SetPaymentDueInterval sets PaymentDueInterval field to given value.

### HasPaymentDueInterval

`func (o *ProductConfig) HasPaymentDueInterval() bool`

HasPaymentDueInterval returns a boolean if a field has been set.

### GetPeriodicFees

`func (o *ProductConfig) GetPeriodicFees() []ProductConfigPeriodicFeesInner`

GetPeriodicFees returns the PeriodicFees field if non-nil, zero value otherwise.

### GetPeriodicFeesOk

`func (o *ProductConfig) GetPeriodicFeesOk() (*[]ProductConfigPeriodicFeesInner, bool)`

GetPeriodicFeesOk returns a tuple with the PeriodicFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicFees

`func (o *ProductConfig) SetPeriodicFees(v []ProductConfigPeriodicFeesInner)`

SetPeriodicFees sets PeriodicFees field to given value.

### HasPeriodicFees

`func (o *ProductConfig) HasPeriodicFees() bool`

HasPeriodicFees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


