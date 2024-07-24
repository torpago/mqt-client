# ProductCreditLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | **decimal.Decimal** | Maximum credit limit. | 
**Min** | **decimal.Decimal** | Minimum credit limit. | 

## Methods

### NewProductCreditLine

`func NewProductCreditLine(max decimal.Decimal, min decimal.Decimal, ) *ProductCreditLine`

NewProductCreditLine instantiates a new ProductCreditLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCreditLineWithDefaults

`func NewProductCreditLineWithDefaults() *ProductCreditLine`

NewProductCreditLineWithDefaults instantiates a new ProductCreditLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *ProductCreditLine) GetMax() decimal.Decimal`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ProductCreditLine) GetMaxOk() (*decimal.Decimal, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ProductCreditLine) SetMax(v decimal.Decimal)`

SetMax sets Max field to given value.


### GetMin

`func (o *ProductCreditLine) GetMin() decimal.Decimal`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ProductCreditLine) GetMinOk() (*decimal.Decimal, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ProductCreditLine) SetMin(v decimal.Decimal)`

SetMin sets Min field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


