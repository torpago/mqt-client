# DelinquencyBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketNumber** | **int32** | Delinquency bucket number in the returned array. Delinquency buckets are returned from most recent to least; the most recent delinquency bucket is &#x60;1&#x60;. | 
**CurrentDue** | **decimal.Decimal** | Current amount that is due for this delinquency bucket. | 
**DaysPastDue** | **int32** | Total number of days that the payment is past due for this delinquency bucket. | 
**PastDueCarriedForward** | **decimal.Decimal** | Amount that is past due and carried forward from previous delinquency buckets. | 
**PaymentDueDate** | **time.Time** | Date that the payment was due for this delinquency bucket. | 
**TotalDue** | **decimal.Decimal** | Total amount that is due for this delinquency bucket; the sum of &#x60;past_due_carried_forward&#x60; and &#x60;current_due&#x60;. | 

## Methods

### NewDelinquencyBucketResponse

`func NewDelinquencyBucketResponse(bucketNumber int32, currentDue decimal.Decimal, daysPastDue int32, pastDueCarriedForward decimal.Decimal, paymentDueDate time.Time, totalDue decimal.Decimal, ) *DelinquencyBucketResponse`

NewDelinquencyBucketResponse instantiates a new DelinquencyBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelinquencyBucketResponseWithDefaults

`func NewDelinquencyBucketResponseWithDefaults() *DelinquencyBucketResponse`

NewDelinquencyBucketResponseWithDefaults instantiates a new DelinquencyBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketNumber

`func (o *DelinquencyBucketResponse) GetBucketNumber() int32`

GetBucketNumber returns the BucketNumber field if non-nil, zero value otherwise.

### GetBucketNumberOk

`func (o *DelinquencyBucketResponse) GetBucketNumberOk() (*int32, bool)`

GetBucketNumberOk returns a tuple with the BucketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketNumber

`func (o *DelinquencyBucketResponse) SetBucketNumber(v int32)`

SetBucketNumber sets BucketNumber field to given value.


### GetCurrentDue

`func (o *DelinquencyBucketResponse) GetCurrentDue() decimal.Decimal`

GetCurrentDue returns the CurrentDue field if non-nil, zero value otherwise.

### GetCurrentDueOk

`func (o *DelinquencyBucketResponse) GetCurrentDueOk() (*decimal.Decimal, bool)`

GetCurrentDueOk returns a tuple with the CurrentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDue

`func (o *DelinquencyBucketResponse) SetCurrentDue(v decimal.Decimal)`

SetCurrentDue sets CurrentDue field to given value.


### GetDaysPastDue

`func (o *DelinquencyBucketResponse) GetDaysPastDue() int32`

GetDaysPastDue returns the DaysPastDue field if non-nil, zero value otherwise.

### GetDaysPastDueOk

`func (o *DelinquencyBucketResponse) GetDaysPastDueOk() (*int32, bool)`

GetDaysPastDueOk returns a tuple with the DaysPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPastDue

`func (o *DelinquencyBucketResponse) SetDaysPastDue(v int32)`

SetDaysPastDue sets DaysPastDue field to given value.


### GetPastDueCarriedForward

`func (o *DelinquencyBucketResponse) GetPastDueCarriedForward() decimal.Decimal`

GetPastDueCarriedForward returns the PastDueCarriedForward field if non-nil, zero value otherwise.

### GetPastDueCarriedForwardOk

`func (o *DelinquencyBucketResponse) GetPastDueCarriedForwardOk() (*decimal.Decimal, bool)`

GetPastDueCarriedForwardOk returns a tuple with the PastDueCarriedForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastDueCarriedForward

`func (o *DelinquencyBucketResponse) SetPastDueCarriedForward(v decimal.Decimal)`

SetPastDueCarriedForward sets PastDueCarriedForward field to given value.


### GetPaymentDueDate

`func (o *DelinquencyBucketResponse) GetPaymentDueDate() time.Time`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *DelinquencyBucketResponse) GetPaymentDueDateOk() (*time.Time, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *DelinquencyBucketResponse) SetPaymentDueDate(v time.Time)`

SetPaymentDueDate sets PaymentDueDate field to given value.


### GetTotalDue

`func (o *DelinquencyBucketResponse) GetTotalDue() decimal.Decimal`

GetTotalDue returns the TotalDue field if non-nil, zero value otherwise.

### GetTotalDueOk

`func (o *DelinquencyBucketResponse) GetTotalDueOk() (*decimal.Decimal, bool)`

GetTotalDueOk returns a tuple with the TotalDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDue

`func (o *DelinquencyBucketResponse) SetTotalDue(v decimal.Decimal)`

SetTotalDue sets TotalDue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


