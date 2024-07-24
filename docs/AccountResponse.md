# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationTime** | Pointer to **time.Time** | Date and time when the credit account was activated on Marqeta&#39;s credit platform, in UTC. | [optional] 
**AvailableCredit** | **decimal.Decimal** | Amount of credit available for use on the credit account. | 
**BundleToken** | Pointer to **string** | Unique identifier of the associated bundle product. | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the parent business program.  Either a &#x60;user_token&#x60; or &#x60;business_token&#x60; is returned, not both. | [optional] 
**Config** | [**AccountConfigResponse**](AccountConfigResponse.md) |  | 
**CreatedTime** | **time.Time** | Date and time when the credit account was created on Marqeta&#39;s credit platform, in UTC. | 
**CreditLimit** | **decimal.Decimal** | Maximum balance the credit account can carry. | 
**CreditProductToken** | Pointer to **string** | Unique identifier of the associated credit product. | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [default to CURRENCYCODE_USD]
**CurrentBalance** | **decimal.Decimal** | Current purchase balance on the credit account. | 
**Description** | Pointer to **string** | Description for the credit account. | [optional] 
**ExternalOfferId** | Pointer to **string** | Unique identifier you provide of the associated external credit offer. | [optional] 
**LatestStatementCycleType** | Pointer to [**CycleType**](CycleType.md) |  | [optional] 
**MaxAprSchedules** | Pointer to [**[]MaxAPRSchedulesResponse**](MaxAPRSchedulesResponse.md) | Contains &#x60;max_apr_schedule&#x60; objects, which provide information about any temporary overrides of the APRs on the credit account. This could include special APR rates due to account/user sub status changes. | [optional] 
**Name** | Pointer to **string** | Name of the credit account. | [optional] 
**RemainingMinPaymentDue** | **decimal.Decimal** | Amount remaining on the latest statement&#39;s minimum payment, after it&#39;s adjusted for payments, returned payments, and applicable credits that occurred after the latest statement&#39;s closing date. | 
**RemainingStatementBalance** | **decimal.Decimal** | Amount remaining on the latest statement&#39;s balance, after it&#39;s adjusted for payments, returned payments, and applicable credits that occurred after the latest statement&#39;s closing date. | 
**Status** | [**AccountStatusEnum**](AccountStatusEnum.md) |  | 
**Substatuses** | Pointer to **[]string** | Substatuses of the credit account. | [optional] 
**Token** | **string** | Unique identifier of the credit account. | 
**Type** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**UpdatedTime** | **time.Time** | Date and time when the credit account was last updated on Marqeta&#39;s credit platform, in UTC. | 
**Usages** | [**[]AccountUsageResponse**](AccountUsageResponse.md) | Contains one or more &#x60;usages&#x60; objects that contain information on how a credit account is used and what types of balances are permitted on the account. | 
**UserSubstatuses** | Pointer to **[]string** | Substatuses of the users under the credit account. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the primary account holder.  Either a &#x60;user_token&#x60; or &#x60;business_token&#x60; is returned, not both. | [optional] 

## Methods

### NewAccountResponse

`func NewAccountResponse(availableCredit decimal.Decimal, config AccountConfigResponse, createdTime time.Time, creditLimit decimal.Decimal, currencyCode CurrencyCode, currentBalance decimal.Decimal, remainingMinPaymentDue decimal.Decimal, remainingStatementBalance decimal.Decimal, status AccountStatusEnum, token string, updatedTime time.Time, usages []AccountUsageResponse, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationTime

`func (o *AccountResponse) GetActivationTime() time.Time`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *AccountResponse) GetActivationTimeOk() (*time.Time, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *AccountResponse) SetActivationTime(v time.Time)`

SetActivationTime sets ActivationTime field to given value.

### HasActivationTime

`func (o *AccountResponse) HasActivationTime() bool`

HasActivationTime returns a boolean if a field has been set.

### GetAvailableCredit

`func (o *AccountResponse) GetAvailableCredit() decimal.Decimal`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *AccountResponse) GetAvailableCreditOk() (*decimal.Decimal, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *AccountResponse) SetAvailableCredit(v decimal.Decimal)`

SetAvailableCredit sets AvailableCredit field to given value.


### GetBundleToken

`func (o *AccountResponse) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *AccountResponse) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *AccountResponse) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.

### HasBundleToken

`func (o *AccountResponse) HasBundleToken() bool`

HasBundleToken returns a boolean if a field has been set.

### GetBusinessToken

`func (o *AccountResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *AccountResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *AccountResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *AccountResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetConfig

`func (o *AccountResponse) GetConfig() AccountConfigResponse`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountResponse) GetConfigOk() (*AccountConfigResponse, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountResponse) SetConfig(v AccountConfigResponse)`

SetConfig sets Config field to given value.


### GetCreatedTime

`func (o *AccountResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccountResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccountResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCreditLimit

`func (o *AccountResponse) GetCreditLimit() decimal.Decimal`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountResponse) GetCreditLimitOk() (*decimal.Decimal, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountResponse) SetCreditLimit(v decimal.Decimal)`

SetCreditLimit sets CreditLimit field to given value.


### GetCreditProductToken

`func (o *AccountResponse) GetCreditProductToken() string`

GetCreditProductToken returns the CreditProductToken field if non-nil, zero value otherwise.

### GetCreditProductTokenOk

`func (o *AccountResponse) GetCreditProductTokenOk() (*string, bool)`

GetCreditProductTokenOk returns a tuple with the CreditProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditProductToken

`func (o *AccountResponse) SetCreditProductToken(v string)`

SetCreditProductToken sets CreditProductToken field to given value.

### HasCreditProductToken

`func (o *AccountResponse) HasCreditProductToken() bool`

HasCreditProductToken returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AccountResponse) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountResponse) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountResponse) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetCurrentBalance

`func (o *AccountResponse) GetCurrentBalance() decimal.Decimal`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *AccountResponse) GetCurrentBalanceOk() (*decimal.Decimal, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *AccountResponse) SetCurrentBalance(v decimal.Decimal)`

SetCurrentBalance sets CurrentBalance field to given value.


### GetDescription

`func (o *AccountResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalOfferId

`func (o *AccountResponse) GetExternalOfferId() string`

GetExternalOfferId returns the ExternalOfferId field if non-nil, zero value otherwise.

### GetExternalOfferIdOk

`func (o *AccountResponse) GetExternalOfferIdOk() (*string, bool)`

GetExternalOfferIdOk returns a tuple with the ExternalOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOfferId

`func (o *AccountResponse) SetExternalOfferId(v string)`

SetExternalOfferId sets ExternalOfferId field to given value.

### HasExternalOfferId

`func (o *AccountResponse) HasExternalOfferId() bool`

HasExternalOfferId returns a boolean if a field has been set.

### GetLatestStatementCycleType

`func (o *AccountResponse) GetLatestStatementCycleType() CycleType`

GetLatestStatementCycleType returns the LatestStatementCycleType field if non-nil, zero value otherwise.

### GetLatestStatementCycleTypeOk

`func (o *AccountResponse) GetLatestStatementCycleTypeOk() (*CycleType, bool)`

GetLatestStatementCycleTypeOk returns a tuple with the LatestStatementCycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStatementCycleType

`func (o *AccountResponse) SetLatestStatementCycleType(v CycleType)`

SetLatestStatementCycleType sets LatestStatementCycleType field to given value.

### HasLatestStatementCycleType

`func (o *AccountResponse) HasLatestStatementCycleType() bool`

HasLatestStatementCycleType returns a boolean if a field has been set.

### GetMaxAprSchedules

`func (o *AccountResponse) GetMaxAprSchedules() []MaxAPRSchedulesResponse`

GetMaxAprSchedules returns the MaxAprSchedules field if non-nil, zero value otherwise.

### GetMaxAprSchedulesOk

`func (o *AccountResponse) GetMaxAprSchedulesOk() (*[]MaxAPRSchedulesResponse, bool)`

GetMaxAprSchedulesOk returns a tuple with the MaxAprSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAprSchedules

`func (o *AccountResponse) SetMaxAprSchedules(v []MaxAPRSchedulesResponse)`

SetMaxAprSchedules sets MaxAprSchedules field to given value.

### HasMaxAprSchedules

`func (o *AccountResponse) HasMaxAprSchedules() bool`

HasMaxAprSchedules returns a boolean if a field has been set.

### GetName

`func (o *AccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemainingMinPaymentDue

`func (o *AccountResponse) GetRemainingMinPaymentDue() decimal.Decimal`

GetRemainingMinPaymentDue returns the RemainingMinPaymentDue field if non-nil, zero value otherwise.

### GetRemainingMinPaymentDueOk

`func (o *AccountResponse) GetRemainingMinPaymentDueOk() (*decimal.Decimal, bool)`

GetRemainingMinPaymentDueOk returns a tuple with the RemainingMinPaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingMinPaymentDue

`func (o *AccountResponse) SetRemainingMinPaymentDue(v decimal.Decimal)`

SetRemainingMinPaymentDue sets RemainingMinPaymentDue field to given value.


### GetRemainingStatementBalance

`func (o *AccountResponse) GetRemainingStatementBalance() decimal.Decimal`

GetRemainingStatementBalance returns the RemainingStatementBalance field if non-nil, zero value otherwise.

### GetRemainingStatementBalanceOk

`func (o *AccountResponse) GetRemainingStatementBalanceOk() (*decimal.Decimal, bool)`

GetRemainingStatementBalanceOk returns a tuple with the RemainingStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingStatementBalance

`func (o *AccountResponse) SetRemainingStatementBalance(v decimal.Decimal)`

SetRemainingStatementBalance sets RemainingStatementBalance field to given value.


### GetStatus

`func (o *AccountResponse) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountResponse) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountResponse) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.


### GetSubstatuses

`func (o *AccountResponse) GetSubstatuses() []string`

GetSubstatuses returns the Substatuses field if non-nil, zero value otherwise.

### GetSubstatusesOk

`func (o *AccountResponse) GetSubstatusesOk() (*[]string, bool)`

GetSubstatusesOk returns a tuple with the Substatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatuses

`func (o *AccountResponse) SetSubstatuses(v []string)`

SetSubstatuses sets Substatuses field to given value.

### HasSubstatuses

`func (o *AccountResponse) HasSubstatuses() bool`

HasSubstatuses returns a boolean if a field has been set.

### GetToken

`func (o *AccountResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccountResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccountResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *AccountResponse) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountResponse) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountResponse) SetType(v AccountType)`

SetType sets Type field to given value.

### HasType

`func (o *AccountResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AccountResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AccountResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AccountResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.


### GetUsages

`func (o *AccountResponse) GetUsages() []AccountUsageResponse`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *AccountResponse) GetUsagesOk() (*[]AccountUsageResponse, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *AccountResponse) SetUsages(v []AccountUsageResponse)`

SetUsages sets Usages field to given value.


### GetUserSubstatuses

`func (o *AccountResponse) GetUserSubstatuses() []string`

GetUserSubstatuses returns the UserSubstatuses field if non-nil, zero value otherwise.

### GetUserSubstatusesOk

`func (o *AccountResponse) GetUserSubstatusesOk() (*[]string, bool)`

GetUserSubstatusesOk returns a tuple with the UserSubstatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubstatuses

`func (o *AccountResponse) SetUserSubstatuses(v []string)`

SetUserSubstatuses sets UserSubstatuses field to given value.

### HasUserSubstatuses

`func (o *AccountResponse) HasUserSubstatuses() bool`

HasUserSubstatuses returns a boolean if a field has been set.

### GetUserToken

`func (o *AccountResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *AccountResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *AccountResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *AccountResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


