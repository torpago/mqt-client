/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TransactionDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDevice{}

// TransactionDevice Contains information about the device used in the transaction to enhance the risk decisioning process. Use this data to improve fraud prevention and dispute resolution.
type TransactionDevice struct {
	// Unique identifier of the data component bound to the credential.
	BindingId *string `json:"binding_id,omitempty"`
	// IP address of the device. The presence of the IP address helps determine if the transaction was initiated from an unusual network, helping establish a pattern of safe device usage that further confirms the authenticity of the consumer who initiated the transaction.
	IpAddress *string `json:"ip_address,omitempty"`
	// Geographic coordinates of the device. Contains the latitude and longitude of the device used when the cardholder was authenticated during checkout. This field helps to determine if the transaction was initiated from an unexpected location.
	Location *string `json:"location,omitempty"`
	// Telephone number of the device. Contains the phone number that was used to authenticate the consumer during checkout, or the consumer's preferred phone number. The presence of the phone number helps establish the consumer's authenticity when matching the phone number provided during checkout to a list of known phone numbers for the consumer.
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// NewTransactionDevice instantiates a new TransactionDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDevice() *TransactionDevice {
	this := TransactionDevice{}
	return &this
}

// NewTransactionDeviceWithDefaults instantiates a new TransactionDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDeviceWithDefaults() *TransactionDevice {
	this := TransactionDevice{}
	return &this
}

// GetBindingId returns the BindingId field value if set, zero value otherwise.
func (o *TransactionDevice) GetBindingId() string {
	if o == nil || IsNil(o.BindingId) {
		var ret string
		return ret
	}
	return *o.BindingId
}

// GetBindingIdOk returns a tuple with the BindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDevice) GetBindingIdOk() (*string, bool) {
	if o == nil || IsNil(o.BindingId) {
		return nil, false
	}
	return o.BindingId, true
}

// HasBindingId returns a boolean if a field has been set.
func (o *TransactionDevice) HasBindingId() bool {
	if o != nil && !IsNil(o.BindingId) {
		return true
	}

	return false
}

// SetBindingId gets a reference to the given string and assigns it to the BindingId field.
func (o *TransactionDevice) SetBindingId(v string) {
	o.BindingId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *TransactionDevice) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDevice) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *TransactionDevice) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *TransactionDevice) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *TransactionDevice) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDevice) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *TransactionDevice) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *TransactionDevice) SetLocation(v string) {
	o.Location = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *TransactionDevice) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDevice) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *TransactionDevice) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *TransactionDevice) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o TransactionDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BindingId) {
		toSerialize["binding_id"] = o.BindingId
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableTransactionDevice struct {
	value *TransactionDevice
	isSet bool
}

func (v NullableTransactionDevice) Get() *TransactionDevice {
	return v.value
}

func (v *NullableTransactionDevice) Set(val *TransactionDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDevice(val *TransactionDevice) *NullableTransactionDevice {
	return &NullableTransactionDevice{value: val, isSet: true}
}

func (v NullableTransactionDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


