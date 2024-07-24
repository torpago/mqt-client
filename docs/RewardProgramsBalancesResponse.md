# RewardProgramsBalancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleClosingDate** | **time.Time** | Closing date of the billing cycle for which rewards were accrued, in UTC. | 
**BillingCycleOpeningDate** | **time.Time** | Opening date of the billing cycle for which rewards were accrued, in UTC. | 
**NetBalance** | **decimal.Decimal** | The net balance for a billing cycle, which is total amount spent during a billing cycle, minus any refunds or reversals. Used to determine reward accrual. | 
**PendingRewardBalance** | **decimal.Decimal** | The pending balance of the rewards accrued for the current billing cycle. Pending rewards cannot be redeemed. | 
**Percentage** | **decimal.Decimal** | The reward percentage applied to the balance for the current billing cycle. Determined by the reward rules config. | 
**RewardProgramToken** | **string** | Unique identifier of reward program for which to return balances. | 
**TotalRewardBalance** | **decimal.Decimal** | The total balance of the rewards accrued to date minus the rewards redeemed to date. | 

## Methods

### NewRewardProgramsBalancesResponse

`func NewRewardProgramsBalancesResponse(billingCycleClosingDate time.Time, billingCycleOpeningDate time.Time, netBalance decimal.Decimal, pendingRewardBalance decimal.Decimal, percentage decimal.Decimal, rewardProgramToken string, totalRewardBalance decimal.Decimal, ) *RewardProgramsBalancesResponse`

NewRewardProgramsBalancesResponse instantiates a new RewardProgramsBalancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardProgramsBalancesResponseWithDefaults

`func NewRewardProgramsBalancesResponseWithDefaults() *RewardProgramsBalancesResponse`

NewRewardProgramsBalancesResponseWithDefaults instantiates a new RewardProgramsBalancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleClosingDate

`func (o *RewardProgramsBalancesResponse) GetBillingCycleClosingDate() time.Time`

GetBillingCycleClosingDate returns the BillingCycleClosingDate field if non-nil, zero value otherwise.

### GetBillingCycleClosingDateOk

`func (o *RewardProgramsBalancesResponse) GetBillingCycleClosingDateOk() (*time.Time, bool)`

GetBillingCycleClosingDateOk returns a tuple with the BillingCycleClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleClosingDate

`func (o *RewardProgramsBalancesResponse) SetBillingCycleClosingDate(v time.Time)`

SetBillingCycleClosingDate sets BillingCycleClosingDate field to given value.


### GetBillingCycleOpeningDate

`func (o *RewardProgramsBalancesResponse) GetBillingCycleOpeningDate() time.Time`

GetBillingCycleOpeningDate returns the BillingCycleOpeningDate field if non-nil, zero value otherwise.

### GetBillingCycleOpeningDateOk

`func (o *RewardProgramsBalancesResponse) GetBillingCycleOpeningDateOk() (*time.Time, bool)`

GetBillingCycleOpeningDateOk returns a tuple with the BillingCycleOpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleOpeningDate

`func (o *RewardProgramsBalancesResponse) SetBillingCycleOpeningDate(v time.Time)`

SetBillingCycleOpeningDate sets BillingCycleOpeningDate field to given value.


### GetNetBalance

`func (o *RewardProgramsBalancesResponse) GetNetBalance() decimal.Decimal`

GetNetBalance returns the NetBalance field if non-nil, zero value otherwise.

### GetNetBalanceOk

`func (o *RewardProgramsBalancesResponse) GetNetBalanceOk() (*decimal.Decimal, bool)`

GetNetBalanceOk returns a tuple with the NetBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBalance

`func (o *RewardProgramsBalancesResponse) SetNetBalance(v decimal.Decimal)`

SetNetBalance sets NetBalance field to given value.


### GetPendingRewardBalance

`func (o *RewardProgramsBalancesResponse) GetPendingRewardBalance() decimal.Decimal`

GetPendingRewardBalance returns the PendingRewardBalance field if non-nil, zero value otherwise.

### GetPendingRewardBalanceOk

`func (o *RewardProgramsBalancesResponse) GetPendingRewardBalanceOk() (*decimal.Decimal, bool)`

GetPendingRewardBalanceOk returns a tuple with the PendingRewardBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRewardBalance

`func (o *RewardProgramsBalancesResponse) SetPendingRewardBalance(v decimal.Decimal)`

SetPendingRewardBalance sets PendingRewardBalance field to given value.


### GetPercentage

`func (o *RewardProgramsBalancesResponse) GetPercentage() decimal.Decimal`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *RewardProgramsBalancesResponse) GetPercentageOk() (*decimal.Decimal, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *RewardProgramsBalancesResponse) SetPercentage(v decimal.Decimal)`

SetPercentage sets Percentage field to given value.


### GetRewardProgramToken

`func (o *RewardProgramsBalancesResponse) GetRewardProgramToken() string`

GetRewardProgramToken returns the RewardProgramToken field if non-nil, zero value otherwise.

### GetRewardProgramTokenOk

`func (o *RewardProgramsBalancesResponse) GetRewardProgramTokenOk() (*string, bool)`

GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardProgramToken

`func (o *RewardProgramsBalancesResponse) SetRewardProgramToken(v string)`

SetRewardProgramToken sets RewardProgramToken field to given value.


### GetTotalRewardBalance

`func (o *RewardProgramsBalancesResponse) GetTotalRewardBalance() decimal.Decimal`

GetTotalRewardBalance returns the TotalRewardBalance field if non-nil, zero value otherwise.

### GetTotalRewardBalanceOk

`func (o *RewardProgramsBalancesResponse) GetTotalRewardBalanceOk() (*decimal.Decimal, bool)`

GetTotalRewardBalanceOk returns a tuple with the TotalRewardBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRewardBalance

`func (o *RewardProgramsBalancesResponse) SetTotalRewardBalance(v decimal.Decimal)`

SetTotalRewardBalance sets TotalRewardBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


