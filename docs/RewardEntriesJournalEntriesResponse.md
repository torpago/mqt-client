# RewardEntriesJournalEntriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelatedJournalEntryToken** | Pointer to **string** | Unique identifier of the related journal entry to which the reward rule was applied to trigger the reward entry. | [optional] 
**TransactionAmount** | Pointer to **float32** | The transaction amount to which the reward rule was applied. Used to determine the value of the reward entry. | [optional] 
**Value** | Pointer to **float32** | Value of the reward entry. | [optional] 

## Methods

### NewRewardEntriesJournalEntriesResponse

`func NewRewardEntriesJournalEntriesResponse() *RewardEntriesJournalEntriesResponse`

NewRewardEntriesJournalEntriesResponse instantiates a new RewardEntriesJournalEntriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardEntriesJournalEntriesResponseWithDefaults

`func NewRewardEntriesJournalEntriesResponseWithDefaults() *RewardEntriesJournalEntriesResponse`

NewRewardEntriesJournalEntriesResponseWithDefaults instantiates a new RewardEntriesJournalEntriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelatedJournalEntryToken

`func (o *RewardEntriesJournalEntriesResponse) GetRelatedJournalEntryToken() string`

GetRelatedJournalEntryToken returns the RelatedJournalEntryToken field if non-nil, zero value otherwise.

### GetRelatedJournalEntryTokenOk

`func (o *RewardEntriesJournalEntriesResponse) GetRelatedJournalEntryTokenOk() (*string, bool)`

GetRelatedJournalEntryTokenOk returns a tuple with the RelatedJournalEntryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedJournalEntryToken

`func (o *RewardEntriesJournalEntriesResponse) SetRelatedJournalEntryToken(v string)`

SetRelatedJournalEntryToken sets RelatedJournalEntryToken field to given value.

### HasRelatedJournalEntryToken

`func (o *RewardEntriesJournalEntriesResponse) HasRelatedJournalEntryToken() bool`

HasRelatedJournalEntryToken returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *RewardEntriesJournalEntriesResponse) GetTransactionAmount() float32`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *RewardEntriesJournalEntriesResponse) GetTransactionAmountOk() (*float32, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *RewardEntriesJournalEntriesResponse) SetTransactionAmount(v float32)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *RewardEntriesJournalEntriesResponse) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetValue

`func (o *RewardEntriesJournalEntriesResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RewardEntriesJournalEntriesResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RewardEntriesJournalEntriesResponse) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RewardEntriesJournalEntriesResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


