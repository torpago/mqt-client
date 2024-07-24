# PolicyAprTierReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarginRate** | **decimal.Decimal** | Number of percentage points added to the prime rate, used to calculate a variable APR value. | 

## Methods

### NewPolicyAprTierReq

`func NewPolicyAprTierReq(marginRate decimal.Decimal, ) *PolicyAprTierReq`

NewPolicyAprTierReq instantiates a new PolicyAprTierReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAprTierReqWithDefaults

`func NewPolicyAprTierReqWithDefaults() *PolicyAprTierReq`

NewPolicyAprTierReqWithDefaults instantiates a new PolicyAprTierReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarginRate

`func (o *PolicyAprTierReq) GetMarginRate() decimal.Decimal`

GetMarginRate returns the MarginRate field if non-nil, zero value otherwise.

### GetMarginRateOk

`func (o *PolicyAprTierReq) GetMarginRateOk() (*decimal.Decimal, bool)`

GetMarginRateOk returns a tuple with the MarginRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRate

`func (o *PolicyAprTierReq) SetMarginRate(v decimal.Decimal)`

SetMarginRate sets MarginRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


