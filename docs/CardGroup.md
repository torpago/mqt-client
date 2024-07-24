# CardGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardTokens** | Pointer to **[]string** | Array of card tokens associated with group. There will be at least one card token in the array. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time the card group was created in the system. The date and time is provided in ISO 8601 format. | [optional] 
**LastIssuedCardToken** | Pointer to **string** | Unique identifier of the last reissued card token associated with group. It may be empty if there is no reissued card. | [optional] 
**SourceCardToken** | **string** | Unique identifier of the card token associated with group. This is the card that will be used to create the card group. The Card Group Service will send a request to JCard to verify that this card is not a reissue or replacement. | 
**Token** | **string** | Unique identifier of the card group.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | 
**UpdatedTime** | Pointer to **time.Time** | Date and time the card group was last updated in the system. The date and time is provided in ISO 8601 format. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user this card group belongs to. | [optional] 

## Methods

### NewCardGroup

`func NewCardGroup(sourceCardToken string, token string, ) *CardGroup`

NewCardGroup instantiates a new CardGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardGroupWithDefaults

`func NewCardGroupWithDefaults() *CardGroup`

NewCardGroupWithDefaults instantiates a new CardGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardTokens

`func (o *CardGroup) GetCardTokens() []string`

GetCardTokens returns the CardTokens field if non-nil, zero value otherwise.

### GetCardTokensOk

`func (o *CardGroup) GetCardTokensOk() (*[]string, bool)`

GetCardTokensOk returns a tuple with the CardTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTokens

`func (o *CardGroup) SetCardTokens(v []string)`

SetCardTokens sets CardTokens field to given value.

### HasCardTokens

`func (o *CardGroup) HasCardTokens() bool`

HasCardTokens returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CardGroup) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CardGroup) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CardGroup) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CardGroup) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastIssuedCardToken

`func (o *CardGroup) GetLastIssuedCardToken() string`

GetLastIssuedCardToken returns the LastIssuedCardToken field if non-nil, zero value otherwise.

### GetLastIssuedCardTokenOk

`func (o *CardGroup) GetLastIssuedCardTokenOk() (*string, bool)`

GetLastIssuedCardTokenOk returns a tuple with the LastIssuedCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIssuedCardToken

`func (o *CardGroup) SetLastIssuedCardToken(v string)`

SetLastIssuedCardToken sets LastIssuedCardToken field to given value.

### HasLastIssuedCardToken

`func (o *CardGroup) HasLastIssuedCardToken() bool`

HasLastIssuedCardToken returns a boolean if a field has been set.

### GetSourceCardToken

`func (o *CardGroup) GetSourceCardToken() string`

GetSourceCardToken returns the SourceCardToken field if non-nil, zero value otherwise.

### GetSourceCardTokenOk

`func (o *CardGroup) GetSourceCardTokenOk() (*string, bool)`

GetSourceCardTokenOk returns a tuple with the SourceCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCardToken

`func (o *CardGroup) SetSourceCardToken(v string)`

SetSourceCardToken sets SourceCardToken field to given value.


### GetToken

`func (o *CardGroup) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardGroup) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardGroup) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedTime

`func (o *CardGroup) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *CardGroup) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *CardGroup) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *CardGroup) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUserToken

`func (o *CardGroup) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardGroup) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardGroup) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *CardGroup) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


