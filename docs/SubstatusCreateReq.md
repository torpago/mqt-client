# SubstatusCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]SubstatusCreateReqAttributesInner**](SubstatusCreateReqAttributesInner.md) | Additional dynamic attributes related to the substatus. | [optional] 
**Events** | [**[]SubstatusEvent**](SubstatusEvent.md) | List of events related to the substatus. | 
**ResourceToken** | **string** | The unique identifier of the user or account for which you want to create a substatus. | 
**ResourceType** | **string** | Possible values: USER, ACCOUNT.  | 
**Substatus** | **string** | Possible values: HARDSHIP, FRAUD, MLA, SCRA, DECEASED, BANKRUPTCY.  | 
**Token** | Pointer to **string** | Unique identifier of the substatus. | [optional] 

## Methods

### NewSubstatusCreateReq

`func NewSubstatusCreateReq(events []SubstatusEvent, resourceToken string, resourceType string, substatus string, ) *SubstatusCreateReq`

NewSubstatusCreateReq instantiates a new SubstatusCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubstatusCreateReqWithDefaults

`func NewSubstatusCreateReqWithDefaults() *SubstatusCreateReq`

NewSubstatusCreateReqWithDefaults instantiates a new SubstatusCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SubstatusCreateReq) GetAttributes() []SubstatusCreateReqAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubstatusCreateReq) GetAttributesOk() (*[]SubstatusCreateReqAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubstatusCreateReq) SetAttributes(v []SubstatusCreateReqAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubstatusCreateReq) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEvents

`func (o *SubstatusCreateReq) GetEvents() []SubstatusEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SubstatusCreateReq) GetEventsOk() (*[]SubstatusEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SubstatusCreateReq) SetEvents(v []SubstatusEvent)`

SetEvents sets Events field to given value.


### GetResourceToken

`func (o *SubstatusCreateReq) GetResourceToken() string`

GetResourceToken returns the ResourceToken field if non-nil, zero value otherwise.

### GetResourceTokenOk

`func (o *SubstatusCreateReq) GetResourceTokenOk() (*string, bool)`

GetResourceTokenOk returns a tuple with the ResourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceToken

`func (o *SubstatusCreateReq) SetResourceToken(v string)`

SetResourceToken sets ResourceToken field to given value.


### GetResourceType

`func (o *SubstatusCreateReq) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *SubstatusCreateReq) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *SubstatusCreateReq) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetSubstatus

`func (o *SubstatusCreateReq) GetSubstatus() string`

GetSubstatus returns the Substatus field if non-nil, zero value otherwise.

### GetSubstatusOk

`func (o *SubstatusCreateReq) GetSubstatusOk() (*string, bool)`

GetSubstatusOk returns a tuple with the Substatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatus

`func (o *SubstatusCreateReq) SetSubstatus(v string)`

SetSubstatus sets Substatus field to given value.


### GetToken

`func (o *SubstatusCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SubstatusCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SubstatusCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SubstatusCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


