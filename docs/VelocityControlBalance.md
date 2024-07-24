# VelocityControlBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the velocity control is active. | [optional] [default to false]
**AmountLimit** | **float32** | Maximum monetary sum that can be cleared within the time period defined by velocity period. | 
**Association** | Pointer to [**Association**](Association.md) |  | [optional] 
**CurrencyCode** | **string** | Three-character ISO 4217 currency code. | 
**MerchantScope** | Pointer to [**MerchantScope**](MerchantScope.md) |  | [optional] 
**Name** | Pointer to **string** | Description of how the velocity control restricts spending. For example, \&quot;Max spend of $500 per day\&quot; or \&quot;Max spend of $5000 per month for non-exempt employees\&quot;. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the velocity control. | [optional] 
**UsageLimit** | Pointer to **int32** | Maximum number of times a card can be used within the time period defined by the &#x60;velocity_window&#x60; field.  Leave &#x60;null&#x60; to indicate that the card can be used an unlimited number of times. | [optional] 
**VelocityWindow** | Pointer to [**VelocityWindow**](VelocityWindow.md) |  | [optional] 
**VelocityWindowStartDay** | Pointer to **int32** | Start day of the velocity window defined by the &#x60;velocity_window&#x60; field. Default value is &#x60;1&#x60; | [optional] 
**Available** | Pointer to [**VelocityControlBalanceAllOfAvailable**](VelocityControlBalanceAllOfAvailable.md) |  | [optional] 

## Methods

### NewVelocityControlBalance

`func NewVelocityControlBalance(amountLimit float32, currencyCode string, ) *VelocityControlBalance`

NewVelocityControlBalance instantiates a new VelocityControlBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityControlBalanceWithDefaults

`func NewVelocityControlBalanceWithDefaults() *VelocityControlBalance`

NewVelocityControlBalanceWithDefaults instantiates a new VelocityControlBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *VelocityControlBalance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *VelocityControlBalance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *VelocityControlBalance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *VelocityControlBalance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAmountLimit

`func (o *VelocityControlBalance) GetAmountLimit() float32`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *VelocityControlBalance) GetAmountLimitOk() (*float32, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *VelocityControlBalance) SetAmountLimit(v float32)`

SetAmountLimit sets AmountLimit field to given value.


### GetAssociation

`func (o *VelocityControlBalance) GetAssociation() Association`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *VelocityControlBalance) GetAssociationOk() (*Association, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *VelocityControlBalance) SetAssociation(v Association)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *VelocityControlBalance) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *VelocityControlBalance) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *VelocityControlBalance) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *VelocityControlBalance) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetMerchantScope

`func (o *VelocityControlBalance) GetMerchantScope() MerchantScope`

GetMerchantScope returns the MerchantScope field if non-nil, zero value otherwise.

### GetMerchantScopeOk

`func (o *VelocityControlBalance) GetMerchantScopeOk() (*MerchantScope, bool)`

GetMerchantScopeOk returns a tuple with the MerchantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantScope

`func (o *VelocityControlBalance) SetMerchantScope(v MerchantScope)`

SetMerchantScope sets MerchantScope field to given value.

### HasMerchantScope

`func (o *VelocityControlBalance) HasMerchantScope() bool`

HasMerchantScope returns a boolean if a field has been set.

### GetName

`func (o *VelocityControlBalance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VelocityControlBalance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VelocityControlBalance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VelocityControlBalance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *VelocityControlBalance) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VelocityControlBalance) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VelocityControlBalance) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VelocityControlBalance) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsageLimit

`func (o *VelocityControlBalance) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *VelocityControlBalance) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *VelocityControlBalance) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *VelocityControlBalance) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetVelocityWindow

`func (o *VelocityControlBalance) GetVelocityWindow() VelocityWindow`

GetVelocityWindow returns the VelocityWindow field if non-nil, zero value otherwise.

### GetVelocityWindowOk

`func (o *VelocityControlBalance) GetVelocityWindowOk() (*VelocityWindow, bool)`

GetVelocityWindowOk returns a tuple with the VelocityWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityWindow

`func (o *VelocityControlBalance) SetVelocityWindow(v VelocityWindow)`

SetVelocityWindow sets VelocityWindow field to given value.

### HasVelocityWindow

`func (o *VelocityControlBalance) HasVelocityWindow() bool`

HasVelocityWindow returns a boolean if a field has been set.

### GetVelocityWindowStartDay

`func (o *VelocityControlBalance) GetVelocityWindowStartDay() int32`

GetVelocityWindowStartDay returns the VelocityWindowStartDay field if non-nil, zero value otherwise.

### GetVelocityWindowStartDayOk

`func (o *VelocityControlBalance) GetVelocityWindowStartDayOk() (*int32, bool)`

GetVelocityWindowStartDayOk returns a tuple with the VelocityWindowStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityWindowStartDay

`func (o *VelocityControlBalance) SetVelocityWindowStartDay(v int32)`

SetVelocityWindowStartDay sets VelocityWindowStartDay field to given value.

### HasVelocityWindowStartDay

`func (o *VelocityControlBalance) HasVelocityWindowStartDay() bool`

HasVelocityWindowStartDay returns a boolean if a field has been set.

### GetAvailable

`func (o *VelocityControlBalance) GetAvailable() VelocityControlBalanceAllOfAvailable`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *VelocityControlBalance) GetAvailableOk() (*VelocityControlBalanceAllOfAvailable, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *VelocityControlBalance) SetAvailable(v VelocityControlBalanceAllOfAvailable)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *VelocityControlBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


