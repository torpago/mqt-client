# StatementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the statement summary is generated. | 
**AvailableCredit** | **decimal.Decimal** | Amount available to spend on the credit account, as of the statement closing date. | 
**ClosingBalance** | **decimal.Decimal** | Balance of the credit account when the statement period ended. | 
**ClosingDate** | **time.Time** | Date and time when the statement period ended. | 
**CreatedTime** | **time.Time** | Date and time when the statement summary was created on Marqeta&#39;s credit platform, in UTC. | 
**CreditLimit** | Pointer to **decimal.Decimal** | Maximum balance the credit account can carry, as of the statement closing date. | [optional] 
**Credits** | **decimal.Decimal** | Total amount of credits received during the statement period. | 
**CycleType** | [**CycleType**](CycleType.md) |  | 
**DaysInBillingCycle** | **int32** | Number of days in the billing cycle, also known as the statement period. | 
**Fees** | **decimal.Decimal** | Total amount of fees charged during the statement period. Does not include periodic fees. | 
**Interest** | **decimal.Decimal** | Total amount of interest charged during the statement period. | 
**OpeningBalance** | **decimal.Decimal** | Balance of the credit account when the statement period began. | 
**OpeningDate** | **time.Time** | Date and time when the statement period began. | 
**PastDueAmount** | **decimal.Decimal** | Total payment amount, required to make the account current. | 
**Payments** | **decimal.Decimal** | Total amount of payments made during the statement period. | 
**Purchases** | **decimal.Decimal** | Total amount of purchases made during the statement period. | 
**Token** | **string** | Unique identifier of the statement summary. | 

## Methods

### NewStatementSummary

`func NewStatementSummary(accountToken string, availableCredit decimal.Decimal, closingBalance decimal.Decimal, closingDate time.Time, createdTime time.Time, credits decimal.Decimal, cycleType CycleType, daysInBillingCycle int32, fees decimal.Decimal, interest decimal.Decimal, openingBalance decimal.Decimal, openingDate time.Time, pastDueAmount decimal.Decimal, payments decimal.Decimal, purchases decimal.Decimal, token string, ) *StatementSummary`

NewStatementSummary instantiates a new StatementSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementSummaryWithDefaults

`func NewStatementSummaryWithDefaults() *StatementSummary`

NewStatementSummaryWithDefaults instantiates a new StatementSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *StatementSummary) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *StatementSummary) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *StatementSummary) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetAvailableCredit

`func (o *StatementSummary) GetAvailableCredit() decimal.Decimal`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *StatementSummary) GetAvailableCreditOk() (*decimal.Decimal, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *StatementSummary) SetAvailableCredit(v decimal.Decimal)`

SetAvailableCredit sets AvailableCredit field to given value.


### GetClosingBalance

`func (o *StatementSummary) GetClosingBalance() decimal.Decimal`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *StatementSummary) GetClosingBalanceOk() (*decimal.Decimal, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *StatementSummary) SetClosingBalance(v decimal.Decimal)`

SetClosingBalance sets ClosingBalance field to given value.


### GetClosingDate

`func (o *StatementSummary) GetClosingDate() time.Time`

GetClosingDate returns the ClosingDate field if non-nil, zero value otherwise.

### GetClosingDateOk

`func (o *StatementSummary) GetClosingDateOk() (*time.Time, bool)`

GetClosingDateOk returns a tuple with the ClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDate

`func (o *StatementSummary) SetClosingDate(v time.Time)`

SetClosingDate sets ClosingDate field to given value.


### GetCreatedTime

`func (o *StatementSummary) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *StatementSummary) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *StatementSummary) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreditLimit

`func (o *StatementSummary) GetCreditLimit() decimal.Decimal`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *StatementSummary) GetCreditLimitOk() (*decimal.Decimal, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *StatementSummary) SetCreditLimit(v decimal.Decimal)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *StatementSummary) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCredits

`func (o *StatementSummary) GetCredits() decimal.Decimal`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *StatementSummary) GetCreditsOk() (*decimal.Decimal, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *StatementSummary) SetCredits(v decimal.Decimal)`

SetCredits sets Credits field to given value.


### GetCycleType

`func (o *StatementSummary) GetCycleType() CycleType`

GetCycleType returns the CycleType field if non-nil, zero value otherwise.

### GetCycleTypeOk

`func (o *StatementSummary) GetCycleTypeOk() (*CycleType, bool)`

GetCycleTypeOk returns a tuple with the CycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleType

`func (o *StatementSummary) SetCycleType(v CycleType)`

SetCycleType sets CycleType field to given value.


### GetDaysInBillingCycle

`func (o *StatementSummary) GetDaysInBillingCycle() int32`

GetDaysInBillingCycle returns the DaysInBillingCycle field if non-nil, zero value otherwise.

### GetDaysInBillingCycleOk

`func (o *StatementSummary) GetDaysInBillingCycleOk() (*int32, bool)`

GetDaysInBillingCycleOk returns a tuple with the DaysInBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInBillingCycle

`func (o *StatementSummary) SetDaysInBillingCycle(v int32)`

SetDaysInBillingCycle sets DaysInBillingCycle field to given value.


### GetFees

`func (o *StatementSummary) GetFees() decimal.Decimal`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *StatementSummary) GetFeesOk() (*decimal.Decimal, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *StatementSummary) SetFees(v decimal.Decimal)`

SetFees sets Fees field to given value.


### GetInterest

`func (o *StatementSummary) GetInterest() decimal.Decimal`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *StatementSummary) GetInterestOk() (*decimal.Decimal, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *StatementSummary) SetInterest(v decimal.Decimal)`

SetInterest sets Interest field to given value.


### GetOpeningBalance

`func (o *StatementSummary) GetOpeningBalance() decimal.Decimal`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *StatementSummary) GetOpeningBalanceOk() (*decimal.Decimal, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *StatementSummary) SetOpeningBalance(v decimal.Decimal)`

SetOpeningBalance sets OpeningBalance field to given value.


### GetOpeningDate

`func (o *StatementSummary) GetOpeningDate() time.Time`

GetOpeningDate returns the OpeningDate field if non-nil, zero value otherwise.

### GetOpeningDateOk

`func (o *StatementSummary) GetOpeningDateOk() (*time.Time, bool)`

GetOpeningDateOk returns a tuple with the OpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDate

`func (o *StatementSummary) SetOpeningDate(v time.Time)`

SetOpeningDate sets OpeningDate field to given value.


### GetPastDueAmount

`func (o *StatementSummary) GetPastDueAmount() decimal.Decimal`

GetPastDueAmount returns the PastDueAmount field if non-nil, zero value otherwise.

### GetPastDueAmountOk

`func (o *StatementSummary) GetPastDueAmountOk() (*decimal.Decimal, bool)`

GetPastDueAmountOk returns a tuple with the PastDueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastDueAmount

`func (o *StatementSummary) SetPastDueAmount(v decimal.Decimal)`

SetPastDueAmount sets PastDueAmount field to given value.


### GetPayments

`func (o *StatementSummary) GetPayments() decimal.Decimal`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *StatementSummary) GetPaymentsOk() (*decimal.Decimal, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *StatementSummary) SetPayments(v decimal.Decimal)`

SetPayments sets Payments field to given value.


### GetPurchases

`func (o *StatementSummary) GetPurchases() decimal.Decimal`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *StatementSummary) GetPurchasesOk() (*decimal.Decimal, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *StatementSummary) SetPurchases(v decimal.Decimal)`

SetPurchases sets Purchases field to given value.


### GetToken

`func (o *StatementSummary) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StatementSummary) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StatementSummary) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


