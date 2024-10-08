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
	"bytes"
	"fmt"
)

// checks if the PTCSoftDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PTCSoftDescriptor{}

// PTCSoftDescriptor struct for PTCSoftDescriptor
type PTCSoftDescriptor struct {
	Address PTCAddress `json:"address"`
	Email *string `json:"email,omitempty"`
	Name string `json:"name"`
	Phone *PTCPhone `json:"phone,omitempty"`
}

type _PTCSoftDescriptor PTCSoftDescriptor

// NewPTCSoftDescriptor instantiates a new PTCSoftDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPTCSoftDescriptor(address PTCAddress, name string) *PTCSoftDescriptor {
	this := PTCSoftDescriptor{}
	this.Address = address
	this.Name = name
	return &this
}

// NewPTCSoftDescriptorWithDefaults instantiates a new PTCSoftDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPTCSoftDescriptorWithDefaults() *PTCSoftDescriptor {
	this := PTCSoftDescriptor{}
	return &this
}

// GetAddress returns the Address field value
func (o *PTCSoftDescriptor) GetAddress() PTCAddress {
	if o == nil {
		var ret PTCAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PTCSoftDescriptor) GetAddressOk() (*PTCAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PTCSoftDescriptor) SetAddress(v PTCAddress) {
	o.Address = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PTCSoftDescriptor) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTCSoftDescriptor) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PTCSoftDescriptor) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PTCSoftDescriptor) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value
func (o *PTCSoftDescriptor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PTCSoftDescriptor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PTCSoftDescriptor) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PTCSoftDescriptor) GetPhone() PTCPhone {
	if o == nil || IsNil(o.Phone) {
		var ret PTCPhone
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTCSoftDescriptor) GetPhoneOk() (*PTCPhone, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PTCSoftDescriptor) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given PTCPhone and assigns it to the Phone field.
func (o *PTCSoftDescriptor) SetPhone(v PTCPhone) {
	o.Phone = &v
}

func (o PTCSoftDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PTCSoftDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

func (o *PTCSoftDescriptor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPTCSoftDescriptor := _PTCSoftDescriptor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPTCSoftDescriptor)

	if err != nil {
		return err
	}

	*o = PTCSoftDescriptor(varPTCSoftDescriptor)

	return err
}

type NullablePTCSoftDescriptor struct {
	value *PTCSoftDescriptor
	isSet bool
}

func (v NullablePTCSoftDescriptor) Get() *PTCSoftDescriptor {
	return v.value
}

func (v *NullablePTCSoftDescriptor) Set(val *PTCSoftDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullablePTCSoftDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullablePTCSoftDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePTCSoftDescriptor(val *PTCSoftDescriptor) *NullablePTCSoftDescriptor {
	return &NullablePTCSoftDescriptor{value: val, isSet: true}
}

func (v NullablePTCSoftDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePTCSoftDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


