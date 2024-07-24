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

// checks if the AmountFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmountFilter{}

// AmountFilter Contains information on the minimum and maximum amounts that the balance for a billing cycle can be to earn the reward.
type AmountFilter struct {
	// Minimum amount that a balance for a billing cycle can be to earn the reward.
	GreaterThan *float32 `json:"greater_than,omitempty"`
	// Maximum amount that a balance for a billing cycle can be to earn the reward.
	LessThan *float32 `json:"less_than,omitempty"`
}

// NewAmountFilter instantiates a new AmountFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmountFilter() *AmountFilter {
	this := AmountFilter{}
	return &this
}

// NewAmountFilterWithDefaults instantiates a new AmountFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmountFilterWithDefaults() *AmountFilter {
	this := AmountFilter{}
	return &this
}

// GetGreaterThan returns the GreaterThan field value if set, zero value otherwise.
func (o *AmountFilter) GetGreaterThan() float32 {
	if o == nil || IsNil(o.GreaterThan) {
		var ret float32
		return ret
	}
	return *o.GreaterThan
}

// GetGreaterThanOk returns a tuple with the GreaterThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountFilter) GetGreaterThanOk() (*float32, bool) {
	if o == nil || IsNil(o.GreaterThan) {
		return nil, false
	}
	return o.GreaterThan, true
}

// HasGreaterThan returns a boolean if a field has been set.
func (o *AmountFilter) HasGreaterThan() bool {
	if o != nil && !IsNil(o.GreaterThan) {
		return true
	}

	return false
}

// SetGreaterThan gets a reference to the given float32 and assigns it to the GreaterThan field.
func (o *AmountFilter) SetGreaterThan(v float32) {
	o.GreaterThan = &v
}

// GetLessThan returns the LessThan field value if set, zero value otherwise.
func (o *AmountFilter) GetLessThan() float32 {
	if o == nil || IsNil(o.LessThan) {
		var ret float32
		return ret
	}
	return *o.LessThan
}

// GetLessThanOk returns a tuple with the LessThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmountFilter) GetLessThanOk() (*float32, bool) {
	if o == nil || IsNil(o.LessThan) {
		return nil, false
	}
	return o.LessThan, true
}

// HasLessThan returns a boolean if a field has been set.
func (o *AmountFilter) HasLessThan() bool {
	if o != nil && !IsNil(o.LessThan) {
		return true
	}

	return false
}

// SetLessThan gets a reference to the given float32 and assigns it to the LessThan field.
func (o *AmountFilter) SetLessThan(v float32) {
	o.LessThan = &v
}

func (o AmountFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmountFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GreaterThan) {
		toSerialize["greater_than"] = o.GreaterThan
	}
	if !IsNil(o.LessThan) {
		toSerialize["less_than"] = o.LessThan
	}
	return toSerialize, nil
}

type NullableAmountFilter struct {
	value *AmountFilter
	isSet bool
}

func (v NullableAmountFilter) Get() *AmountFilter {
	return v.value
}

func (v *NullableAmountFilter) Set(val *AmountFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAmountFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAmountFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmountFilter(val *AmountFilter) *NullableAmountFilter {
	return &NullableAmountFilter{value: val, isSet: true}
}

func (v NullableAmountFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmountFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


