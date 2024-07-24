# SubstatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]SubstatusCreateReqAttributesInner**](SubstatusCreateReqAttributesInner.md) | Additional dynamic attributes related to the substatus. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the substatus was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Events** | Pointer to [**[]SubstatusEventResponseDetails**](SubstatusEventResponseDetails.md) | List of events related to the substatus. | [optional] 
**IsActive** | **bool** | Flag indicating whether the substatus is active (false when deactivated). | 
**ResourceToken** | **string** | substatus resource token | 
**ResourceType** | **string** | substatus resource type | 
**State** | Pointer to **string** | state of the substatus | [optional] 
**Substatus** | **string** | substatus | 
**Token** | **string** | substatus token | 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the substatus was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewSubstatusResponse

`func NewSubstatusResponse(isActive bool, resourceToken string, resourceType string, substatus string, token string, ) *SubstatusResponse`

NewSubstatusResponse instantiates a new SubstatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubstatusResponseWithDefaults

`func NewSubstatusResponseWithDefaults() *SubstatusResponse`

NewSubstatusResponseWithDefaults instantiates a new SubstatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SubstatusResponse) GetAttributes() []SubstatusCreateReqAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubstatusResponse) GetAttributesOk() (*[]SubstatusCreateReqAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubstatusResponse) SetAttributes(v []SubstatusCreateReqAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubstatusResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreatedTime

`func (o *SubstatusResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SubstatusResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SubstatusResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *SubstatusResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetEvents

`func (o *SubstatusResponse) GetEvents() []SubstatusEventResponseDetails`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SubstatusResponse) GetEventsOk() (*[]SubstatusEventResponseDetails, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SubstatusResponse) SetEvents(v []SubstatusEventResponseDetails)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SubstatusResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIsActive

`func (o *SubstatusResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SubstatusResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SubstatusResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetResourceToken

`func (o *SubstatusResponse) GetResourceToken() string`

GetResourceToken returns the ResourceToken field if non-nil, zero value otherwise.

### GetResourceTokenOk

`func (o *SubstatusResponse) GetResourceTokenOk() (*string, bool)`

GetResourceTokenOk returns a tuple with the ResourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceToken

`func (o *SubstatusResponse) SetResourceToken(v string)`

SetResourceToken sets ResourceToken field to given value.


### GetResourceType

`func (o *SubstatusResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *SubstatusResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *SubstatusResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetState

`func (o *SubstatusResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SubstatusResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SubstatusResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SubstatusResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubstatus

`func (o *SubstatusResponse) GetSubstatus() string`

GetSubstatus returns the Substatus field if non-nil, zero value otherwise.

### GetSubstatusOk

`func (o *SubstatusResponse) GetSubstatusOk() (*string, bool)`

GetSubstatusOk returns a tuple with the Substatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatus

`func (o *SubstatusResponse) SetSubstatus(v string)`

SetSubstatus sets Substatus field to given value.


### GetToken

`func (o *SubstatusResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubstatusResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubstatusResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *SubstatusResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *SubstatusResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *SubstatusResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *SubstatusResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


