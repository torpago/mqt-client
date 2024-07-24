# TransactionDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingId** | Pointer to **string** | Unique identifier of the data component bound to the credential. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the device. The presence of the IP address helps determine if the transaction was initiated from an unusual network, helping establish a pattern of safe device usage that further confirms the authenticity of the consumer who initiated the transaction. | [optional] 
**Location** | Pointer to **string** | Geographic coordinates of the device. Contains the latitude and longitude of the device used when the cardholder was authenticated during checkout. This field helps to determine if the transaction was initiated from an unexpected location. | [optional] 
**PhoneNumber** | Pointer to **string** | Telephone number of the device. Contains the phone number that was used to authenticate the consumer during checkout, or the consumer&#39;s preferred phone number. The presence of the phone number helps establish the consumer&#39;s authenticity when matching the phone number provided during checkout to a list of known phone numbers for the consumer. | [optional] 

## Methods

### NewTransactionDevice

`func NewTransactionDevice() *TransactionDevice`

NewTransactionDevice instantiates a new TransactionDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDeviceWithDefaults

`func NewTransactionDeviceWithDefaults() *TransactionDevice`

NewTransactionDeviceWithDefaults instantiates a new TransactionDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingId

`func (o *TransactionDevice) GetBindingId() string`

GetBindingId returns the BindingId field if non-nil, zero value otherwise.

### GetBindingIdOk

`func (o *TransactionDevice) GetBindingIdOk() (*string, bool)`

GetBindingIdOk returns a tuple with the BindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingId

`func (o *TransactionDevice) SetBindingId(v string)`

SetBindingId sets BindingId field to given value.

### HasBindingId

`func (o *TransactionDevice) HasBindingId() bool`

HasBindingId returns a boolean if a field has been set.

### GetIpAddress

`func (o *TransactionDevice) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *TransactionDevice) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *TransactionDevice) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *TransactionDevice) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLocation

`func (o *TransactionDevice) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TransactionDevice) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TransactionDevice) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TransactionDevice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *TransactionDevice) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TransactionDevice) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TransactionDevice) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TransactionDevice) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


