# SubstatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The channel through which the event occurred. | [optional] 
**EffectiveDate** | Pointer to **time.Time** | Effective date of the deactivation, in UTC. | [optional] 
**Reason** | Pointer to **string** | Reason for the event. | [optional] 
**State** | **string** | The initial state of the substatus. ACTIVE - Required for the substatuses of HARDSHIP,FRAUD,MLA,SCRA,DECEASED. BANKRUPTCY_FILED - Required for the substatus of BANKRUPTCY. | 

## Methods

### NewSubstatusEvent

`func NewSubstatusEvent(state string, ) *SubstatusEvent`

NewSubstatusEvent instantiates a new SubstatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubstatusEventWithDefaults

`func NewSubstatusEventWithDefaults() *SubstatusEvent`

NewSubstatusEventWithDefaults instantiates a new SubstatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SubstatusEvent) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SubstatusEvent) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SubstatusEvent) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SubstatusEvent) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *SubstatusEvent) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SubstatusEvent) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SubstatusEvent) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *SubstatusEvent) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetReason

`func (o *SubstatusEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubstatusEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubstatusEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SubstatusEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *SubstatusEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubstatusEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubstatusEvent) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


