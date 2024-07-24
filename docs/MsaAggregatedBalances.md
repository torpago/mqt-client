# MsaAggregatedBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | **decimal.Decimal** |  | 
**Balances** | [**map[string]MsaAggregatedBalances**](MsaAggregatedBalances.md) |  | 
**CachedBalance** | **decimal.Decimal** |  | 
**CreditBalance** | **decimal.Decimal** |  | 
**CurrencyCode** | **string** |  | 
**ImpactedAmount** | Pointer to **decimal.Decimal** |  | [optional] 
**LastUpdatedTime** | **time.Time** |  | 
**LedgerBalance** | **decimal.Decimal** |  | 
**PendingCredits** | **decimal.Decimal** |  | 

## Methods

### NewMsaAggregatedBalances

`func NewMsaAggregatedBalances(availableBalance decimal.Decimal, balances map[string]MsaAggregatedBalances, cachedBalance decimal.Decimal, creditBalance decimal.Decimal, currencyCode string, lastUpdatedTime time.Time, ledgerBalance decimal.Decimal, pendingCredits decimal.Decimal, ) *MsaAggregatedBalances`

NewMsaAggregatedBalances instantiates a new MsaAggregatedBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaAggregatedBalancesWithDefaults

`func NewMsaAggregatedBalancesWithDefaults() *MsaAggregatedBalances`

NewMsaAggregatedBalancesWithDefaults instantiates a new MsaAggregatedBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *MsaAggregatedBalances) GetAvailableBalance() decimal.Decimal`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *MsaAggregatedBalances) GetAvailableBalanceOk() (*decimal.Decimal, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *MsaAggregatedBalances) SetAvailableBalance(v decimal.Decimal)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalances

`func (o *MsaAggregatedBalances) GetBalances() map[string]MsaAggregatedBalances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *MsaAggregatedBalances) GetBalancesOk() (*map[string]MsaAggregatedBalances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *MsaAggregatedBalances) SetBalances(v map[string]MsaAggregatedBalances)`

SetBalances sets Balances field to given value.


### GetCachedBalance

`func (o *MsaAggregatedBalances) GetCachedBalance() decimal.Decimal`

GetCachedBalance returns the CachedBalance field if non-nil, zero value otherwise.

### GetCachedBalanceOk

`func (o *MsaAggregatedBalances) GetCachedBalanceOk() (*decimal.Decimal, bool)`

GetCachedBalanceOk returns a tuple with the CachedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedBalance

`func (o *MsaAggregatedBalances) SetCachedBalance(v decimal.Decimal)`

SetCachedBalance sets CachedBalance field to given value.


### GetCreditBalance

`func (o *MsaAggregatedBalances) GetCreditBalance() decimal.Decimal`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *MsaAggregatedBalances) GetCreditBalanceOk() (*decimal.Decimal, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *MsaAggregatedBalances) SetCreditBalance(v decimal.Decimal)`

SetCreditBalance sets CreditBalance field to given value.


### GetCurrencyCode

`func (o *MsaAggregatedBalances) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MsaAggregatedBalances) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MsaAggregatedBalances) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetImpactedAmount

`func (o *MsaAggregatedBalances) GetImpactedAmount() decimal.Decimal`

GetImpactedAmount returns the ImpactedAmount field if non-nil, zero value otherwise.

### GetImpactedAmountOk

`func (o *MsaAggregatedBalances) GetImpactedAmountOk() (*decimal.Decimal, bool)`

GetImpactedAmountOk returns a tuple with the ImpactedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedAmount

`func (o *MsaAggregatedBalances) SetImpactedAmount(v decimal.Decimal)`

SetImpactedAmount sets ImpactedAmount field to given value.

### HasImpactedAmount

`func (o *MsaAggregatedBalances) HasImpactedAmount() bool`

HasImpactedAmount returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *MsaAggregatedBalances) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *MsaAggregatedBalances) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *MsaAggregatedBalances) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLedgerBalance

`func (o *MsaAggregatedBalances) GetLedgerBalance() decimal.Decimal`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *MsaAggregatedBalances) GetLedgerBalanceOk() (*decimal.Decimal, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *MsaAggregatedBalances) SetLedgerBalance(v decimal.Decimal)`

SetLedgerBalance sets LedgerBalance field to given value.


### GetPendingCredits

`func (o *MsaAggregatedBalances) GetPendingCredits() decimal.Decimal`

GetPendingCredits returns the PendingCredits field if non-nil, zero value otherwise.

### GetPendingCreditsOk

`func (o *MsaAggregatedBalances) GetPendingCreditsOk() (*decimal.Decimal, bool)`

GetPendingCreditsOk returns a tuple with the PendingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCredits

`func (o *MsaAggregatedBalances) SetPendingCredits(v decimal.Decimal)`

SetPendingCredits sets PendingCredits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


