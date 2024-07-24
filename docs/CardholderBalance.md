# CardholderBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | **float32** | Ledger balance minus any authorized transactions that have not yet cleared. Also known as the cardholder&#39;s purchasing power. When using JIT Funding, this balance is usually equal to $0.00. | 
**Balances** | [**map[string]CardholderBalance**](CardholderBalance.md) | Contains GPA balance information, organized by currency code. | 
**CachedBalance** | **float32** | Not currently in use. | 
**CreditBalance** | **float32** | Not currently in use. | 
**CurrencyCode** | **string** | Three-digit ISO 4217 currency code. | 
**ImpactedAmount** | Pointer to **float32** | Balance change based on the amount of the transaction. | [optional] 
**LastUpdatedTime** | **time.Time** | Date and time when the resource was last updated, in UTC. | 
**LedgerBalance** | **float32** | When using standard funding: The funds that are available to spend immediately, including funds from any authorized transactions that have not yet cleared. When using Just-in-Time (JIT) Funding: Authorized funds that are currently on hold, but not yet cleared. | 
**PendingCredits** | **float32** | ACH loads that have been accepted, but for which the funding time has not yet elapsed. | 

## Methods

### NewCardholderBalance

`func NewCardholderBalance(availableBalance float32, balances map[string]CardholderBalance, cachedBalance float32, creditBalance float32, currencyCode string, lastUpdatedTime time.Time, ledgerBalance float32, pendingCredits float32, ) *CardholderBalance`

NewCardholderBalance instantiates a new CardholderBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderBalanceWithDefaults

`func NewCardholderBalanceWithDefaults() *CardholderBalance`

NewCardholderBalanceWithDefaults instantiates a new CardholderBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *CardholderBalance) GetAvailableBalance() float32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *CardholderBalance) GetAvailableBalanceOk() (*float32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *CardholderBalance) SetAvailableBalance(v float32)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalances

`func (o *CardholderBalance) GetBalances() map[string]CardholderBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *CardholderBalance) GetBalancesOk() (*map[string]CardholderBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *CardholderBalance) SetBalances(v map[string]CardholderBalance)`

SetBalances sets Balances field to given value.


### GetCachedBalance

`func (o *CardholderBalance) GetCachedBalance() float32`

GetCachedBalance returns the CachedBalance field if non-nil, zero value otherwise.

### GetCachedBalanceOk

`func (o *CardholderBalance) GetCachedBalanceOk() (*float32, bool)`

GetCachedBalanceOk returns a tuple with the CachedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedBalance

`func (o *CardholderBalance) SetCachedBalance(v float32)`

SetCachedBalance sets CachedBalance field to given value.


### GetCreditBalance

`func (o *CardholderBalance) GetCreditBalance() float32`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *CardholderBalance) GetCreditBalanceOk() (*float32, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *CardholderBalance) SetCreditBalance(v float32)`

SetCreditBalance sets CreditBalance field to given value.


### GetCurrencyCode

`func (o *CardholderBalance) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CardholderBalance) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CardholderBalance) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetImpactedAmount

`func (o *CardholderBalance) GetImpactedAmount() float32`

GetImpactedAmount returns the ImpactedAmount field if non-nil, zero value otherwise.

### GetImpactedAmountOk

`func (o *CardholderBalance) GetImpactedAmountOk() (*float32, bool)`

GetImpactedAmountOk returns a tuple with the ImpactedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedAmount

`func (o *CardholderBalance) SetImpactedAmount(v float32)`

SetImpactedAmount sets ImpactedAmount field to given value.

### HasImpactedAmount

`func (o *CardholderBalance) HasImpactedAmount() bool`

HasImpactedAmount returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CardholderBalance) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CardholderBalance) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CardholderBalance) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLedgerBalance

`func (o *CardholderBalance) GetLedgerBalance() float32`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *CardholderBalance) GetLedgerBalanceOk() (*float32, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *CardholderBalance) SetLedgerBalance(v float32)`

SetLedgerBalance sets LedgerBalance field to given value.


### GetPendingCredits

`func (o *CardholderBalance) GetPendingCredits() float32`

GetPendingCredits returns the PendingCredits field if non-nil, zero value otherwise.

### GetPendingCreditsOk

`func (o *CardholderBalance) GetPendingCreditsOk() (*float32, bool)`

GetPendingCreditsOk returns a tuple with the PendingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCredits

`func (o *CardholderBalance) SetPendingCredits(v float32)`

SetPendingCredits sets PendingCredits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


