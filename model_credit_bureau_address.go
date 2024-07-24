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

// checks if the CreditBureauAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditBureauAddress{}

// CreditBureauAddress Contains information on the address of the credit bureau.
type CreditBureauAddress struct {
	// Street address.
	Address1 string `json:"address1"`
	// Additional address information.
	Address2 *string `json:"address2,omitempty"`
	// Address city.
	City string `json:"city"`
	// Address state.
	State string `json:"state"`
	// Address ZIP code.
	Zip string `json:"zip"`
}

type _CreditBureauAddress CreditBureauAddress

// NewCreditBureauAddress instantiates a new CreditBureauAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBureauAddress(address1 string, city string, state string, zip string) *CreditBureauAddress {
	this := CreditBureauAddress{}
	this.Address1 = address1
	this.City = city
	this.State = state
	this.Zip = zip
	return &this
}

// NewCreditBureauAddressWithDefaults instantiates a new CreditBureauAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBureauAddressWithDefaults() *CreditBureauAddress {
	this := CreditBureauAddress{}
	return &this
}

// GetAddress1 returns the Address1 field value
func (o *CreditBureauAddress) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *CreditBureauAddress) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *CreditBureauAddress) SetAddress1(v string) {
	o.Address1 = v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *CreditBureauAddress) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBureauAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *CreditBureauAddress) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *CreditBureauAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value
func (o *CreditBureauAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CreditBureauAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CreditBureauAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *CreditBureauAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreditBureauAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreditBureauAddress) SetState(v string) {
	o.State = v
}

// GetZip returns the Zip field value
func (o *CreditBureauAddress) GetZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zip
}

// GetZipOk returns a tuple with the Zip field value
// and a boolean to check if the value has been set.
func (o *CreditBureauAddress) GetZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zip, true
}

// SetZip sets field value
func (o *CreditBureauAddress) SetZip(v string) {
	o.Zip = v
}

func (o CreditBureauAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditBureauAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address1"] = o.Address1
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["zip"] = o.Zip
	return toSerialize, nil
}

func (o *CreditBureauAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address1",
		"city",
		"state",
		"zip",
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

	varCreditBureauAddress := _CreditBureauAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditBureauAddress)

	if err != nil {
		return err
	}

	*o = CreditBureauAddress(varCreditBureauAddress)

	return err
}

type NullableCreditBureauAddress struct {
	value *CreditBureauAddress
	isSet bool
}

func (v NullableCreditBureauAddress) Get() *CreditBureauAddress {
	return v.value
}

func (v *NullableCreditBureauAddress) Set(val *CreditBureauAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBureauAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBureauAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBureauAddress(val *CreditBureauAddress) *NullableCreditBureauAddress {
	return &NullableCreditBureauAddress{value: val, isSet: true}
}

func (v NullableCreditBureauAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBureauAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


