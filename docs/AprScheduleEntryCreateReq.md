# AprScheduleEntryCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Margin** | Pointer to **decimal.Decimal** | Number of percentage points added to the prime rate, used to calculate a variable value.  Used for variable values only. | [optional] 
**Type** | Pointer to **string** | Indicates whether the APR value is fixed or variable. | [optional] [default to "FIXED"]
**Value** | **decimal.Decimal** | Percentage value of the APR.  If the APR type is &#x60;FIXED&#x60;, this is the value of the fixed rate. If the APR type is &#x60;VARIABLE&#x60;, the value is calculated by adding the margin to the prime rate that was stored on Marqeta&#39;s credit platform when your credit program was created.  When backdating an APR, this value cannot be greater than the value of the effective APR on the backdated date. | 

## Methods

### NewAprScheduleEntryCreateReq

`func NewAprScheduleEntryCreateReq(value decimal.Decimal, ) *AprScheduleEntryCreateReq`

NewAprScheduleEntryCreateReq instantiates a new AprScheduleEntryCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAprScheduleEntryCreateReqWithDefaults

`func NewAprScheduleEntryCreateReqWithDefaults() *AprScheduleEntryCreateReq`

NewAprScheduleEntryCreateReqWithDefaults instantiates a new AprScheduleEntryCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMargin

`func (o *AprScheduleEntryCreateReq) GetMargin() decimal.Decimal`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *AprScheduleEntryCreateReq) GetMarginOk() (*decimal.Decimal, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *AprScheduleEntryCreateReq) SetMargin(v decimal.Decimal)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *AprScheduleEntryCreateReq) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetType

`func (o *AprScheduleEntryCreateReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AprScheduleEntryCreateReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AprScheduleEntryCreateReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AprScheduleEntryCreateReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *AprScheduleEntryCreateReq) GetValue() decimal.Decimal`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AprScheduleEntryCreateReq) GetValueOk() (*decimal.Decimal, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AprScheduleEntryCreateReq) SetValue(v decimal.Decimal)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


