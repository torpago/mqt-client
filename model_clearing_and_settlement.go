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

// checks if the ClearingAndSettlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClearingAndSettlement{}

// ClearingAndSettlement Specifies the destination for overdraft funds.
type ClearingAndSettlement struct {
	// Specifies the destination for overdraft funds.  This field does not apply if JIT Funding is enabled.
	OverdraftDestination *string `json:"overdraft_destination,omitempty"`
}

// NewClearingAndSettlement instantiates a new ClearingAndSettlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClearingAndSettlement() *ClearingAndSettlement {
	this := ClearingAndSettlement{}
	var overdraftDestination string = "GPA"
	this.OverdraftDestination = &overdraftDestination
	return &this
}

// NewClearingAndSettlementWithDefaults instantiates a new ClearingAndSettlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClearingAndSettlementWithDefaults() *ClearingAndSettlement {
	this := ClearingAndSettlement{}
	var overdraftDestination string = "GPA"
	this.OverdraftDestination = &overdraftDestination
	return &this
}

// GetOverdraftDestination returns the OverdraftDestination field value if set, zero value otherwise.
func (o *ClearingAndSettlement) GetOverdraftDestination() string {
	if o == nil || IsNil(o.OverdraftDestination) {
		var ret string
		return ret
	}
	return *o.OverdraftDestination
}

// GetOverdraftDestinationOk returns a tuple with the OverdraftDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingAndSettlement) GetOverdraftDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.OverdraftDestination) {
		return nil, false
	}
	return o.OverdraftDestination, true
}

// HasOverdraftDestination returns a boolean if a field has been set.
func (o *ClearingAndSettlement) HasOverdraftDestination() bool {
	if o != nil && !IsNil(o.OverdraftDestination) {
		return true
	}

	return false
}

// SetOverdraftDestination gets a reference to the given string and assigns it to the OverdraftDestination field.
func (o *ClearingAndSettlement) SetOverdraftDestination(v string) {
	o.OverdraftDestination = &v
}

func (o ClearingAndSettlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClearingAndSettlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OverdraftDestination) {
		toSerialize["overdraft_destination"] = o.OverdraftDestination
	}
	return toSerialize, nil
}

type NullableClearingAndSettlement struct {
	value *ClearingAndSettlement
	isSet bool
}

func (v NullableClearingAndSettlement) Get() *ClearingAndSettlement {
	return v.value
}

func (v *NullableClearingAndSettlement) Set(val *ClearingAndSettlement) {
	v.value = val
	v.isSet = true
}

func (v NullableClearingAndSettlement) IsSet() bool {
	return v.isSet
}

func (v *NullableClearingAndSettlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearingAndSettlement(val *ClearingAndSettlement) *NullableClearingAndSettlement {
	return &NullableClearingAndSettlement{value: val, isSet: true}
}

func (v NullableClearingAndSettlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearingAndSettlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


