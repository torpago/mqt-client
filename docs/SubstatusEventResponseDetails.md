# SubstatusEventResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel through which the event occurred. | 
**CreatedTime** | **time.Time** | Creation time of the event. | 
**EffectiveDate** | **time.Time** | Effective date of the event, in UTC. | 
**Reason** | Pointer to **string** | Reason for the event. | [optional] 
**State** | **string** | The state of the event | 

## Methods

### NewSubstatusEventResponseDetails

`func NewSubstatusEventResponseDetails(channel string, createdTime time.Time, effectiveDate time.Time, state string, ) *SubstatusEventResponseDetails`

NewSubstatusEventResponseDetails instantiates a new SubstatusEventResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubstatusEventResponseDetailsWithDefaults

`func NewSubstatusEventResponseDetailsWithDefaults() *SubstatusEventResponseDetails`

NewSubstatusEventResponseDetailsWithDefaults instantiates a new SubstatusEventResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SubstatusEventResponseDetails) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SubstatusEventResponseDetails) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SubstatusEventResponseDetails) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetCreatedTime

`func (o *SubstatusEventResponseDetails) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SubstatusEventResponseDetails) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SubstatusEventResponseDetails) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEffectiveDate

`func (o *SubstatusEventResponseDetails) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SubstatusEventResponseDetails) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SubstatusEventResponseDetails) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetReason

`func (o *SubstatusEventResponseDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubstatusEventResponseDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubstatusEventResponseDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SubstatusEventResponseDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *SubstatusEventResponseDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubstatusEventResponseDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubstatusEventResponseDetails) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


