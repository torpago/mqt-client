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

// checks if the AvsInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvsInformation{}

// AvsInformation Contains address verification information.
type AvsInformation struct {
	// Postal code of the address.
	PostalCode *string `json:"postal_code,omitempty"`
	// Street name and number of the address.
	StreetAddress *string `json:"street_address,omitempty"`
	// United States ZIP code of the address.
	Zip *string `json:"zip,omitempty"`
}

// NewAvsInformation instantiates a new AvsInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvsInformation() *AvsInformation {
	this := AvsInformation{}
	return &this
}

// NewAvsInformationWithDefaults instantiates a new AvsInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvsInformationWithDefaults() *AvsInformation {
	this := AvsInformation{}
	return &this
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AvsInformation) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvsInformation) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AvsInformation) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AvsInformation) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *AvsInformation) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvsInformation) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *AvsInformation) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *AvsInformation) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *AvsInformation) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvsInformation) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *AvsInformation) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *AvsInformation) SetZip(v string) {
	o.Zip = &v
}

func (o AvsInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvsInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableAvsInformation struct {
	value *AvsInformation
	isSet bool
}

func (v NullableAvsInformation) Get() *AvsInformation {
	return v.value
}

func (v *NullableAvsInformation) Set(val *AvsInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAvsInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAvsInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvsInformation(val *AvsInformation) *NullableAvsInformation {
	return &NullableAvsInformation{value: val, isSet: true}
}

func (v NullableAvsInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvsInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


