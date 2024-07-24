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

// checks if the CardholderBalances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardholderBalances{}

// CardholderBalances Returns general purpose account (GPA) balances for a user or business. This object includes a link to balances of related user GPAs.
type CardholderBalances struct {
	Gpa CardholderBalance `json:"gpa"`
	// Array of links to balances of related user GPAs.
	Links []Link `json:"links"`
}

type _CardholderBalances CardholderBalances

// NewCardholderBalances instantiates a new CardholderBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardholderBalances(gpa CardholderBalance, links []Link) *CardholderBalances {
	this := CardholderBalances{}
	this.Gpa = gpa
	this.Links = links
	return &this
}

// NewCardholderBalancesWithDefaults instantiates a new CardholderBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardholderBalancesWithDefaults() *CardholderBalances {
	this := CardholderBalances{}
	return &this
}

// GetGpa returns the Gpa field value
func (o *CardholderBalances) GetGpa() CardholderBalance {
	if o == nil {
		var ret CardholderBalance
		return ret
	}

	return o.Gpa
}

// GetGpaOk returns a tuple with the Gpa field value
// and a boolean to check if the value has been set.
func (o *CardholderBalances) GetGpaOk() (*CardholderBalance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpa, true
}

// SetGpa sets field value
func (o *CardholderBalances) SetGpa(v CardholderBalance) {
	o.Gpa = v
}

// GetLinks returns the Links field value
func (o *CardholderBalances) GetLinks() []Link {
	if o == nil {
		var ret []Link
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CardholderBalances) GetLinksOk() ([]Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *CardholderBalances) SetLinks(v []Link) {
	o.Links = v
}

func (o CardholderBalances) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardholderBalances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gpa"] = o.Gpa
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CardholderBalances) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gpa",
		"links",
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

	varCardholderBalances := _CardholderBalances{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardholderBalances)

	if err != nil {
		return err
	}

	*o = CardholderBalances(varCardholderBalances)

	return err
}

type NullableCardholderBalances struct {
	value *CardholderBalances
	isSet bool
}

func (v NullableCardholderBalances) Get() *CardholderBalances {
	return v.value
}

func (v *NullableCardholderBalances) Set(val *CardholderBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableCardholderBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableCardholderBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardholderBalances(val *CardholderBalances) *NullableCardholderBalances {
	return &NullableCardholderBalances{value: val, isSet: true}
}

func (v NullableCardholderBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardholderBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


