# FraudFeedbackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | **string** | This is the party making a call. | 
**Amount** | **string** |  | 
**IsFraud** | **bool** |  | 
**Status** | **string** | This is the value of the status of the fraud. | 
**TransactionToken** | **string** |  | 

## Methods

### NewFraudFeedbackRequest

`func NewFraudFeedbackRequest(actor string, amount string, isFraud bool, status string, transactionToken string, ) *FraudFeedbackRequest`

NewFraudFeedbackRequest instantiates a new FraudFeedbackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFraudFeedbackRequestWithDefaults

`func NewFraudFeedbackRequestWithDefaults() *FraudFeedbackRequest`

NewFraudFeedbackRequestWithDefaults instantiates a new FraudFeedbackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *FraudFeedbackRequest) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *FraudFeedbackRequest) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *FraudFeedbackRequest) SetActor(v string)`

SetActor sets Actor field to given value.


### GetAmount

`func (o *FraudFeedbackRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FraudFeedbackRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FraudFeedbackRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsFraud

`func (o *FraudFeedbackRequest) GetIsFraud() bool`

GetIsFraud returns the IsFraud field if non-nil, zero value otherwise.

### GetIsFraudOk

`func (o *FraudFeedbackRequest) GetIsFraudOk() (*bool, bool)`

GetIsFraudOk returns a tuple with the IsFraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFraud

`func (o *FraudFeedbackRequest) SetIsFraud(v bool)`

SetIsFraud sets IsFraud field to given value.


### GetStatus

`func (o *FraudFeedbackRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FraudFeedbackRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FraudFeedbackRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactionToken

`func (o *FraudFeedbackRequest) GetTransactionToken() string`

GetTransactionToken returns the TransactionToken field if non-nil, zero value otherwise.

### GetTransactionTokenOk

`func (o *FraudFeedbackRequest) GetTransactionTokenOk() (*string, bool)`

GetTransactionTokenOk returns a tuple with the TransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionToken

`func (o *FraudFeedbackRequest) SetTransactionToken(v string)`

SetTransactionToken sets TransactionToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


