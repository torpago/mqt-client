# SubstatusUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The channel through which deactivation is occurring: - **ADMIN**: Added through the Marqeta Dashboard. - **API**: Initiated through the Core API. - **FRAUD**: Determined by Marqeta or the card network. - **SYSTEM**: Initiated by Marqeta  | [optional] 
**EffectiveDate** | Pointer to **time.Time** | Effective date of the deactivation, in UTC. | [optional] 
**Reason** | Pointer to **string** | Reason for deactivating the substatus. | [optional] 
**State** | **string** | The state of the substatus. | 

## Methods

### NewSubstatusUpdateReq

`func NewSubstatusUpdateReq(state string, ) *SubstatusUpdateReq`

NewSubstatusUpdateReq instantiates a new SubstatusUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubstatusUpdateReqWithDefaults

`func NewSubstatusUpdateReqWithDefaults() *SubstatusUpdateReq`

NewSubstatusUpdateReqWithDefaults instantiates a new SubstatusUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SubstatusUpdateReq) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SubstatusUpdateReq) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SubstatusUpdateReq) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SubstatusUpdateReq) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *SubstatusUpdateReq) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SubstatusUpdateReq) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SubstatusUpdateReq) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *SubstatusUpdateReq) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetReason

`func (o *SubstatusUpdateReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubstatusUpdateReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubstatusUpdateReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SubstatusUpdateReq) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetState

`func (o *SubstatusUpdateReq) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubstatusUpdateReq) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubstatusUpdateReq) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


