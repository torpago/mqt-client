# PaymentAllocationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Total amount of the payment allocation. | 
**Bucket** | **string** | category a portion of the payment is allocated to. | 

## Methods

### NewPaymentAllocationResponse

`func NewPaymentAllocationResponse(amount float32, bucket string, ) *PaymentAllocationResponse`

NewPaymentAllocationResponse instantiates a new PaymentAllocationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAllocationResponseWithDefaults

`func NewPaymentAllocationResponseWithDefaults() *PaymentAllocationResponse`

NewPaymentAllocationResponseWithDefaults instantiates a new PaymentAllocationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentAllocationResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAllocationResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAllocationResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBucket

`func (o *PaymentAllocationResponse) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *PaymentAllocationResponse) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *PaymentAllocationResponse) SetBucket(v string)`

SetBucket sets Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


