# MsaBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | **decimal.Decimal** |  | 
**Balances** | [**map[string]MsaBalances**](MsaBalances.md) |  | 
**CachedBalance** | **decimal.Decimal** |  | 
**CreditBalance** | **decimal.Decimal** |  | 
**CurrencyCode** | **string** |  | 
**ImpactedAmount** | Pointer to **decimal.Decimal** |  | [optional] 
**LastUpdatedTime** | **time.Time** |  | 
**LedgerBalance** | **decimal.Decimal** |  | 
**PendingCredits** | **decimal.Decimal** |  | 

## Methods

### NewMsaBalances

`func NewMsaBalances(availableBalance decimal.Decimal, balances map[string]MsaBalances, cachedBalance decimal.Decimal, creditBalance decimal.Decimal, currencyCode string, lastUpdatedTime time.Time, ledgerBalance decimal.Decimal, pendingCredits decimal.Decimal, ) *MsaBalances`

NewMsaBalances instantiates a new MsaBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsaBalancesWithDefaults

`func NewMsaBalancesWithDefaults() *MsaBalances`

NewMsaBalancesWithDefaults instantiates a new MsaBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *MsaBalances) GetAvailableBalance() decimal.Decimal`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *MsaBalances) GetAvailableBalanceOk() (*decimal.Decimal, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *MsaBalances) SetAvailableBalance(v decimal.Decimal)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalances

`func (o *MsaBalances) GetBalances() map[string]MsaBalances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *MsaBalances) GetBalancesOk() (*map[string]MsaBalances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *MsaBalances) SetBalances(v map[string]MsaBalances)`

SetBalances sets Balances field to given value.


### GetCachedBalance

`func (o *MsaBalances) GetCachedBalance() decimal.Decimal`

GetCachedBalance returns the CachedBalance field if non-nil, zero value otherwise.

### GetCachedBalanceOk

`func (o *MsaBalances) GetCachedBalanceOk() (*decimal.Decimal, bool)`

GetCachedBalanceOk returns a tuple with the CachedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedBalance

`func (o *MsaBalances) SetCachedBalance(v decimal.Decimal)`

SetCachedBalance sets CachedBalance field to given value.


### GetCreditBalance

`func (o *MsaBalances) GetCreditBalance() decimal.Decimal`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *MsaBalances) GetCreditBalanceOk() (*decimal.Decimal, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *MsaBalances) SetCreditBalance(v decimal.Decimal)`

SetCreditBalance sets CreditBalance field to given value.


### GetCurrencyCode

`func (o *MsaBalances) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *MsaBalances) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *MsaBalances) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetImpactedAmount

`func (o *MsaBalances) GetImpactedAmount() decimal.Decimal`

GetImpactedAmount returns the ImpactedAmount field if non-nil, zero value otherwise.

### GetImpactedAmountOk

`func (o *MsaBalances) GetImpactedAmountOk() (*decimal.Decimal, bool)`

GetImpactedAmountOk returns a tuple with the ImpactedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedAmount

`func (o *MsaBalances) SetImpactedAmount(v decimal.Decimal)`

SetImpactedAmount sets ImpactedAmount field to given value.

### HasImpactedAmount

`func (o *MsaBalances) HasImpactedAmount() bool`

HasImpactedAmount returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *MsaBalances) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *MsaBalances) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *MsaBalances) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLedgerBalance

`func (o *MsaBalances) GetLedgerBalance() decimal.Decimal`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *MsaBalances) GetLedgerBalanceOk() (*decimal.Decimal, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *MsaBalances) SetLedgerBalance(v decimal.Decimal)`

SetLedgerBalance sets LedgerBalance field to given value.


### GetPendingCredits

`func (o *MsaBalances) GetPendingCredits() decimal.Decimal`

GetPendingCredits returns the PendingCredits field if non-nil, zero value otherwise.

### GetPendingCreditsOk

`func (o *MsaBalances) GetPendingCreditsOk() (*decimal.Decimal, bool)`

GetPendingCreditsOk returns a tuple with the PendingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCredits

`func (o *MsaBalances) SetPendingCredits(v decimal.Decimal)`

SetPendingCredits sets PendingCredits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


