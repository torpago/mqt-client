# PolicyProductPaymentConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationOrder** | [**[]PaymentAllocationOrderEnum**](PaymentAllocationOrderEnum.md) | Ordered list of balance types to which payments are allocated, from first to last. | [default to ["INTEREST","FEES","PRINCIPAL"]]
**BillingCycleDay** | **int32** | Day of month the billing cycle starts. | 
**BillingCycleDayStrategy** | Pointer to **string** | Determines if the billing cycle day is manually set or determined dynamically during account creation based on cycling logic. | [optional] [default to "MANUAL"]
**BillingCycleFrequency** | Pointer to **string** | Frequency at which your account is billed. | [optional] [default to "MONTHLY"]
**DueDay** | Pointer to **int32** | Day of month the payment for the previous billing cycle is due.  This field is being deprecated and replaced by &#x60;payment_due_interval&#x60; of a product policy. To retrieve &#x60;payment_due_interval&#x60;, see &lt;&lt;/core-api/policies#retrieveProductPolicy, Retrieve credit product policy, payments.payment_due_interval&gt;&gt;. | [optional] 
**MinPaymentCalculation** | [**PolicyProductMinPaymentCalculation**](PolicyProductMinPaymentCalculation.md) |  | 
**PaymentDueInterval** | Pointer to **int32** | Specifies the payment due interval that is used to determine the payment due date for a billing cycle. The accepted values are either -1 or a value between 1 and 26. A value of -1 indicates one day prior to the next billing cycle date | [optional] [default to -1]

## Methods

### NewPolicyProductPaymentConfiguration

`func NewPolicyProductPaymentConfiguration(allocationOrder []PaymentAllocationOrderEnum, billingCycleDay int32, minPaymentCalculation PolicyProductMinPaymentCalculation, ) *PolicyProductPaymentConfiguration`

NewPolicyProductPaymentConfiguration instantiates a new PolicyProductPaymentConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyProductPaymentConfigurationWithDefaults

`func NewPolicyProductPaymentConfigurationWithDefaults() *PolicyProductPaymentConfiguration`

NewPolicyProductPaymentConfigurationWithDefaults instantiates a new PolicyProductPaymentConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationOrder

`func (o *PolicyProductPaymentConfiguration) GetAllocationOrder() []PaymentAllocationOrderEnum`

GetAllocationOrder returns the AllocationOrder field if non-nil, zero value otherwise.

### GetAllocationOrderOk

`func (o *PolicyProductPaymentConfiguration) GetAllocationOrderOk() (*[]PaymentAllocationOrderEnum, bool)`

GetAllocationOrderOk returns a tuple with the AllocationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationOrder

`func (o *PolicyProductPaymentConfiguration) SetAllocationOrder(v []PaymentAllocationOrderEnum)`

SetAllocationOrder sets AllocationOrder field to given value.


### GetBillingCycleDay

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleDay() int32`

GetBillingCycleDay returns the BillingCycleDay field if non-nil, zero value otherwise.

### GetBillingCycleDayOk

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleDayOk() (*int32, bool)`

GetBillingCycleDayOk returns a tuple with the BillingCycleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDay

`func (o *PolicyProductPaymentConfiguration) SetBillingCycleDay(v int32)`

SetBillingCycleDay sets BillingCycleDay field to given value.


### GetBillingCycleDayStrategy

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleDayStrategy() string`

GetBillingCycleDayStrategy returns the BillingCycleDayStrategy field if non-nil, zero value otherwise.

### GetBillingCycleDayStrategyOk

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleDayStrategyOk() (*string, bool)`

GetBillingCycleDayStrategyOk returns a tuple with the BillingCycleDayStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleDayStrategy

`func (o *PolicyProductPaymentConfiguration) SetBillingCycleDayStrategy(v string)`

SetBillingCycleDayStrategy sets BillingCycleDayStrategy field to given value.

### HasBillingCycleDayStrategy

`func (o *PolicyProductPaymentConfiguration) HasBillingCycleDayStrategy() bool`

HasBillingCycleDayStrategy returns a boolean if a field has been set.

### GetBillingCycleFrequency

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleFrequency() string`

GetBillingCycleFrequency returns the BillingCycleFrequency field if non-nil, zero value otherwise.

### GetBillingCycleFrequencyOk

`func (o *PolicyProductPaymentConfiguration) GetBillingCycleFrequencyOk() (*string, bool)`

GetBillingCycleFrequencyOk returns a tuple with the BillingCycleFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleFrequency

`func (o *PolicyProductPaymentConfiguration) SetBillingCycleFrequency(v string)`

SetBillingCycleFrequency sets BillingCycleFrequency field to given value.

### HasBillingCycleFrequency

`func (o *PolicyProductPaymentConfiguration) HasBillingCycleFrequency() bool`

HasBillingCycleFrequency returns a boolean if a field has been set.

### GetDueDay

`func (o *PolicyProductPaymentConfiguration) GetDueDay() int32`

GetDueDay returns the DueDay field if non-nil, zero value otherwise.

### GetDueDayOk

`func (o *PolicyProductPaymentConfiguration) GetDueDayOk() (*int32, bool)`

GetDueDayOk returns a tuple with the DueDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDay

`func (o *PolicyProductPaymentConfiguration) SetDueDay(v int32)`

SetDueDay sets DueDay field to given value.

### HasDueDay

`func (o *PolicyProductPaymentConfiguration) HasDueDay() bool`

HasDueDay returns a boolean if a field has been set.

### GetMinPaymentCalculation

`func (o *PolicyProductPaymentConfiguration) GetMinPaymentCalculation() PolicyProductMinPaymentCalculation`

GetMinPaymentCalculation returns the MinPaymentCalculation field if non-nil, zero value otherwise.

### GetMinPaymentCalculationOk

`func (o *PolicyProductPaymentConfiguration) GetMinPaymentCalculationOk() (*PolicyProductMinPaymentCalculation, bool)`

GetMinPaymentCalculationOk returns a tuple with the MinPaymentCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPaymentCalculation

`func (o *PolicyProductPaymentConfiguration) SetMinPaymentCalculation(v PolicyProductMinPaymentCalculation)`

SetMinPaymentCalculation sets MinPaymentCalculation field to given value.


### GetPaymentDueInterval

`func (o *PolicyProductPaymentConfiguration) GetPaymentDueInterval() int32`

GetPaymentDueInterval returns the PaymentDueInterval field if non-nil, zero value otherwise.

### GetPaymentDueIntervalOk

`func (o *PolicyProductPaymentConfiguration) GetPaymentDueIntervalOk() (*int32, bool)`

GetPaymentDueIntervalOk returns a tuple with the PaymentDueInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueInterval

`func (o *PolicyProductPaymentConfiguration) SetPaymentDueInterval(v int32)`

SetPaymentDueInterval sets PaymentDueInterval field to given value.

### HasPaymentDueInterval

`func (o *PolicyProductPaymentConfiguration) HasPaymentDueInterval() bool`

HasPaymentDueInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


