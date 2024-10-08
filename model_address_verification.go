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

// checks if the AddressVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressVerification{}

// AddressVerification Contains address verification information.
type AddressVerification struct {
	// Name of the cardholder.
	Name *string `json:"name,omitempty"`
	// Postal code.
	PostalCode *string `json:"postal_code,omitempty"`
	// Street address provided by the cardholder.
	StreetAddress *string `json:"street_address,omitempty"`
	// United States ZIP code.
	Zip *string `json:"zip,omitempty"`
}

// NewAddressVerification instantiates a new AddressVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressVerification() *AddressVerification {
	this := AddressVerification{}
	return &this
}

// NewAddressVerificationWithDefaults instantiates a new AddressVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressVerificationWithDefaults() *AddressVerification {
	this := AddressVerification{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AddressVerification) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerification) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AddressVerification) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AddressVerification) SetName(v string) {
	o.Name = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressVerification) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerification) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressVerification) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressVerification) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *AddressVerification) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerification) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *AddressVerification) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *AddressVerification) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *AddressVerification) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressVerification) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *AddressVerification) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *AddressVerification) SetZip(v string) {
	o.Zip = &v
}

func (o AddressVerification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableAddressVerification struct {
	value *AddressVerification
	isSet bool
}

func (v NullableAddressVerification) Get() *AddressVerification {
	return v.value
}

func (v *NullableAddressVerification) Set(val *AddressVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressVerification(val *AddressVerification) *NullableAddressVerification {
	return &NullableAddressVerification{value: val, isSet: true}
}

func (v NullableAddressVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


