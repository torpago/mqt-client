# MaxAPRSchedulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** | Date and time when the override APR ends, in UTC. | 
**Reason** | **string** | Reason for the override APR. | 
**StartDate** | **time.Time** | Date and time when the override APR goes into effect, in UTC. | 
**Value** | **decimal.Decimal** | The APR percentage value. This is the value of the fixed rate during the override period. The APR value must adhere to the constraints of the main schedule, such as maximum allowable values. | 

## Methods

### NewMaxAPRSchedulesResponse

`func NewMaxAPRSchedulesResponse(endDate time.Time, reason string, startDate time.Time, value decimal.Decimal, ) *MaxAPRSchedulesResponse`

NewMaxAPRSchedulesResponse instantiates a new MaxAPRSchedulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxAPRSchedulesResponseWithDefaults

`func NewMaxAPRSchedulesResponseWithDefaults() *MaxAPRSchedulesResponse`

NewMaxAPRSchedulesResponseWithDefaults instantiates a new MaxAPRSchedulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *MaxAPRSchedulesResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MaxAPRSchedulesResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MaxAPRSchedulesResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetReason

`func (o *MaxAPRSchedulesResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MaxAPRSchedulesResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MaxAPRSchedulesResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStartDate

`func (o *MaxAPRSchedulesResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MaxAPRSchedulesResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MaxAPRSchedulesResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetValue

`func (o *MaxAPRSchedulesResponse) GetValue() decimal.Decimal`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MaxAPRSchedulesResponse) GetValueOk() (*decimal.Decimal, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MaxAPRSchedulesResponse) SetValue(v decimal.Decimal)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


