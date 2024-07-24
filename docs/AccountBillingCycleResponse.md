# AccountBillingCycleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | Token of the associated account. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the Billing Cycle was created on Marqeta&#39;s credit platform | [optional] 
**CurrentEndTime** | Pointer to **time.Time** | End time of the current billing cycle. | [optional] 
**CurrentPaymentDueDate** | Pointer to **time.Time** | Payment due date for the current billing cycle. | [optional] 
**CurrentStartTime** | Pointer to **time.Time** | Start time of the current billing cycle. | [optional] 
**NextEndTime** | Pointer to **time.Time** | End time of the next billing cycle. | [optional] 
**NextPaymentDueDate** | Pointer to **time.Time** | Payment due date for the next billing cycle. | [optional] 
**NextStartTime** | Pointer to **time.Time** | Start time of the next billing cycle. | [optional] 
**ShortCode** | Pointer to **string** | Unique identifier of the billing cycle&#39;s short code. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the BillingCycle was last updated on Marqeta&#39;s credit platform | [optional] 

## Methods

### NewAccountBillingCycleResponse

`func NewAccountBillingCycleResponse() *AccountBillingCycleResponse`

NewAccountBillingCycleResponse instantiates a new AccountBillingCycleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBillingCycleResponseWithDefaults

`func NewAccountBillingCycleResponseWithDefaults() *AccountBillingCycleResponse`

NewAccountBillingCycleResponseWithDefaults instantiates a new AccountBillingCycleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *AccountBillingCycleResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *AccountBillingCycleResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *AccountBillingCycleResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *AccountBillingCycleResponse) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AccountBillingCycleResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountBillingCycleResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountBillingCycleResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AccountBillingCycleResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrentEndTime

`func (o *AccountBillingCycleResponse) GetCurrentEndTime() time.Time`

GetCurrentEndTime returns the CurrentEndTime field if non-nil, zero value otherwise.

### GetCurrentEndTimeOk

`func (o *AccountBillingCycleResponse) GetCurrentEndTimeOk() (*time.Time, bool)`

GetCurrentEndTimeOk returns a tuple with the CurrentEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEndTime

`func (o *AccountBillingCycleResponse) SetCurrentEndTime(v time.Time)`

SetCurrentEndTime sets CurrentEndTime field to given value.

### HasCurrentEndTime

`func (o *AccountBillingCycleResponse) HasCurrentEndTime() bool`

HasCurrentEndTime returns a boolean if a field has been set.

### GetCurrentPaymentDueDate

`func (o *AccountBillingCycleResponse) GetCurrentPaymentDueDate() time.Time`

GetCurrentPaymentDueDate returns the CurrentPaymentDueDate field if non-nil, zero value otherwise.

### GetCurrentPaymentDueDateOk

`func (o *AccountBillingCycleResponse) GetCurrentPaymentDueDateOk() (*time.Time, bool)`

GetCurrentPaymentDueDateOk returns a tuple with the CurrentPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPaymentDueDate

`func (o *AccountBillingCycleResponse) SetCurrentPaymentDueDate(v time.Time)`

SetCurrentPaymentDueDate sets CurrentPaymentDueDate field to given value.

### HasCurrentPaymentDueDate

`func (o *AccountBillingCycleResponse) HasCurrentPaymentDueDate() bool`

HasCurrentPaymentDueDate returns a boolean if a field has been set.

### GetCurrentStartTime

`func (o *AccountBillingCycleResponse) GetCurrentStartTime() time.Time`

GetCurrentStartTime returns the CurrentStartTime field if non-nil, zero value otherwise.

### GetCurrentStartTimeOk

`func (o *AccountBillingCycleResponse) GetCurrentStartTimeOk() (*time.Time, bool)`

GetCurrentStartTimeOk returns a tuple with the CurrentStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStartTime

`func (o *AccountBillingCycleResponse) SetCurrentStartTime(v time.Time)`

SetCurrentStartTime sets CurrentStartTime field to given value.

### HasCurrentStartTime

`func (o *AccountBillingCycleResponse) HasCurrentStartTime() bool`

HasCurrentStartTime returns a boolean if a field has been set.

### GetNextEndTime

`func (o *AccountBillingCycleResponse) GetNextEndTime() time.Time`

GetNextEndTime returns the NextEndTime field if non-nil, zero value otherwise.

### GetNextEndTimeOk

`func (o *AccountBillingCycleResponse) GetNextEndTimeOk() (*time.Time, bool)`

GetNextEndTimeOk returns a tuple with the NextEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEndTime

`func (o *AccountBillingCycleResponse) SetNextEndTime(v time.Time)`

SetNextEndTime sets NextEndTime field to given value.

### HasNextEndTime

`func (o *AccountBillingCycleResponse) HasNextEndTime() bool`

HasNextEndTime returns a boolean if a field has been set.

### GetNextPaymentDueDate

`func (o *AccountBillingCycleResponse) GetNextPaymentDueDate() time.Time`

GetNextPaymentDueDate returns the NextPaymentDueDate field if non-nil, zero value otherwise.

### GetNextPaymentDueDateOk

`func (o *AccountBillingCycleResponse) GetNextPaymentDueDateOk() (*time.Time, bool)`

GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentDueDate

`func (o *AccountBillingCycleResponse) SetNextPaymentDueDate(v time.Time)`

SetNextPaymentDueDate sets NextPaymentDueDate field to given value.

### HasNextPaymentDueDate

`func (o *AccountBillingCycleResponse) HasNextPaymentDueDate() bool`

HasNextPaymentDueDate returns a boolean if a field has been set.

### GetNextStartTime

`func (o *AccountBillingCycleResponse) GetNextStartTime() time.Time`

GetNextStartTime returns the NextStartTime field if non-nil, zero value otherwise.

### GetNextStartTimeOk

`func (o *AccountBillingCycleResponse) GetNextStartTimeOk() (*time.Time, bool)`

GetNextStartTimeOk returns a tuple with the NextStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartTime

`func (o *AccountBillingCycleResponse) SetNextStartTime(v time.Time)`

SetNextStartTime sets NextStartTime field to given value.

### HasNextStartTime

`func (o *AccountBillingCycleResponse) HasNextStartTime() bool`

HasNextStartTime returns a boolean if a field has been set.

### GetShortCode

`func (o *AccountBillingCycleResponse) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *AccountBillingCycleResponse) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *AccountBillingCycleResponse) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *AccountBillingCycleResponse) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AccountBillingCycleResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AccountBillingCycleResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AccountBillingCycleResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AccountBillingCycleResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


