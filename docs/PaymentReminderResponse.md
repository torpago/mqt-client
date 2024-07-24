# PaymentReminderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | Token of the associated account. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the Billing Cycle was created on Marqeta&#39;s credit platform | [optional] 
**DaysUntilDue** | Pointer to **int32** | Days until payment cutoff date | [optional] 
**PaymentCutoffDate** | Pointer to **time.Time** | Last day a payment can be made before interest and fees are charged to the account. | [optional] 
**PaymentDueDate** | Pointer to **time.Time** | Payment due date, based on the credit account settings. | [optional] 
**RemainingMinimumPaymentDue** | Pointer to **decimal.Decimal** | Amount remaining on the latest statement&#39;s minimum payment, after it&#39;s adjusted for payments, returned payments, and applicable credits that occurred after the latest statement&#39;s closing date. | [optional] 
**StatementSummaryToken** | Pointer to **string** | Token of the associated statement summary | [optional] 
**Token** | Pointer to **string** | Token of the payment reminder | [optional] 

## Methods

### NewPaymentReminderResponse

`func NewPaymentReminderResponse() *PaymentReminderResponse`

NewPaymentReminderResponse instantiates a new PaymentReminderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentReminderResponseWithDefaults

`func NewPaymentReminderResponseWithDefaults() *PaymentReminderResponse`

NewPaymentReminderResponseWithDefaults instantiates a new PaymentReminderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *PaymentReminderResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentReminderResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentReminderResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *PaymentReminderResponse) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PaymentReminderResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentReminderResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentReminderResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PaymentReminderResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDaysUntilDue

`func (o *PaymentReminderResponse) GetDaysUntilDue() int32`

GetDaysUntilDue returns the DaysUntilDue field if non-nil, zero value otherwise.

### GetDaysUntilDueOk

`func (o *PaymentReminderResponse) GetDaysUntilDueOk() (*int32, bool)`

GetDaysUntilDueOk returns a tuple with the DaysUntilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilDue

`func (o *PaymentReminderResponse) SetDaysUntilDue(v int32)`

SetDaysUntilDue sets DaysUntilDue field to given value.

### HasDaysUntilDue

`func (o *PaymentReminderResponse) HasDaysUntilDue() bool`

HasDaysUntilDue returns a boolean if a field has been set.

### GetPaymentCutoffDate

`func (o *PaymentReminderResponse) GetPaymentCutoffDate() time.Time`

GetPaymentCutoffDate returns the PaymentCutoffDate field if non-nil, zero value otherwise.

### GetPaymentCutoffDateOk

`func (o *PaymentReminderResponse) GetPaymentCutoffDateOk() (*time.Time, bool)`

GetPaymentCutoffDateOk returns a tuple with the PaymentCutoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCutoffDate

`func (o *PaymentReminderResponse) SetPaymentCutoffDate(v time.Time)`

SetPaymentCutoffDate sets PaymentCutoffDate field to given value.

### HasPaymentCutoffDate

`func (o *PaymentReminderResponse) HasPaymentCutoffDate() bool`

HasPaymentCutoffDate returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *PaymentReminderResponse) GetPaymentDueDate() time.Time`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *PaymentReminderResponse) GetPaymentDueDateOk() (*time.Time, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *PaymentReminderResponse) SetPaymentDueDate(v time.Time)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *PaymentReminderResponse) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetRemainingMinimumPaymentDue

`func (o *PaymentReminderResponse) GetRemainingMinimumPaymentDue() decimal.Decimal`

GetRemainingMinimumPaymentDue returns the RemainingMinimumPaymentDue field if non-nil, zero value otherwise.

### GetRemainingMinimumPaymentDueOk

`func (o *PaymentReminderResponse) GetRemainingMinimumPaymentDueOk() (*decimal.Decimal, bool)`

GetRemainingMinimumPaymentDueOk returns a tuple with the RemainingMinimumPaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMinimumPaymentDue

`func (o *PaymentReminderResponse) SetRemainingMinimumPaymentDue(v decimal.Decimal)`

SetRemainingMinimumPaymentDue sets RemainingMinimumPaymentDue field to given value.

### HasRemainingMinimumPaymentDue

`func (o *PaymentReminderResponse) HasRemainingMinimumPaymentDue() bool`

HasRemainingMinimumPaymentDue returns a boolean if a field has been set.

### GetStatementSummaryToken

`func (o *PaymentReminderResponse) GetStatementSummaryToken() string`

GetStatementSummaryToken returns the StatementSummaryToken field if non-nil, zero value otherwise.

### GetStatementSummaryTokenOk

`func (o *PaymentReminderResponse) GetStatementSummaryTokenOk() (*string, bool)`

GetStatementSummaryTokenOk returns a tuple with the StatementSummaryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementSummaryToken

`func (o *PaymentReminderResponse) SetStatementSummaryToken(v string)`

SetStatementSummaryToken sets StatementSummaryToken field to given value.

### HasStatementSummaryToken

`func (o *PaymentReminderResponse) HasStatementSummaryToken() bool`

HasStatementSummaryToken returns a boolean if a field has been set.

### GetToken

`func (o *PaymentReminderResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentReminderResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentReminderResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentReminderResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


