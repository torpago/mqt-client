# PolicyAprTierResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apr** | Pointer to **decimal.Decimal** | Value of the APR. | [optional] [default to 0]
**MarginRate** | Pointer to **decimal.Decimal** | Margin rate for the risk tier for a pricing strategy. | [optional] 

## Methods

### NewPolicyAprTierResponse

`func NewPolicyAprTierResponse() *PolicyAprTierResponse`

NewPolicyAprTierResponse instantiates a new PolicyAprTierResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprTierResponseWithDefaults

`func NewPolicyAprTierResponseWithDefaults() *PolicyAprTierResponse`

NewPolicyAprTierResponseWithDefaults instantiates a new PolicyAprTierResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApr

`func (o *PolicyAprTierResponse) GetApr() decimal.Decimal`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *PolicyAprTierResponse) GetAprOk() (*decimal.Decimal, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *PolicyAprTierResponse) SetApr(v decimal.Decimal)`

SetApr sets Apr field to given value.

### HasApr

`func (o *PolicyAprTierResponse) HasApr() bool`

HasApr returns a boolean if a field has been set.

### GetMarginRate

`func (o *PolicyAprTierResponse) GetMarginRate() decimal.Decimal`

GetMarginRate returns the MarginRate field if non-nil, zero value otherwise.

### GetMarginRateOk

`func (o *PolicyAprTierResponse) GetMarginRateOk() (*decimal.Decimal, bool)`

GetMarginRateOk returns a tuple with the MarginRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRate

`func (o *PolicyAprTierResponse) SetMarginRate(v decimal.Decimal)`

SetMarginRate sets MarginRate field to given value.

### HasMarginRate

`func (o *PolicyAprTierResponse) HasMarginRate() bool`

HasMarginRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


