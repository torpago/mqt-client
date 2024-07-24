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

// checks if the MccDynamicFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MccDynamicFilter{}

// MccDynamicFilter Contains information on the dynamic merchant category code (MCC) for a reward.
type MccDynamicFilter struct {
	// One or more dynamic MCCs.
	Includes []DynamicMccType `json:"includes"`
}

type _MccDynamicFilter MccDynamicFilter

// NewMccDynamicFilter instantiates a new MccDynamicFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMccDynamicFilter(includes []DynamicMccType) *MccDynamicFilter {
	this := MccDynamicFilter{}
	this.Includes = includes
	return &this
}

// NewMccDynamicFilterWithDefaults instantiates a new MccDynamicFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMccDynamicFilterWithDefaults() *MccDynamicFilter {
	this := MccDynamicFilter{}
	return &this
}

// GetIncludes returns the Includes field value
func (o *MccDynamicFilter) GetIncludes() []DynamicMccType {
	if o == nil {
		var ret []DynamicMccType
		return ret
	}

	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value
// and a boolean to check if the value has been set.
func (o *MccDynamicFilter) GetIncludesOk() ([]DynamicMccType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Includes, true
}

// SetIncludes sets field value
func (o *MccDynamicFilter) SetIncludes(v []DynamicMccType) {
	o.Includes = v
}

func (o MccDynamicFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MccDynamicFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["includes"] = o.Includes
	return toSerialize, nil
}

func (o *MccDynamicFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"includes",
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

	varMccDynamicFilter := _MccDynamicFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMccDynamicFilter)

	if err != nil {
		return err
	}

	*o = MccDynamicFilter(varMccDynamicFilter)

	return err
}

type NullableMccDynamicFilter struct {
	value *MccDynamicFilter
	isSet bool
}

func (v NullableMccDynamicFilter) Get() *MccDynamicFilter {
	return v.value
}

func (v *NullableMccDynamicFilter) Set(val *MccDynamicFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMccDynamicFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMccDynamicFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMccDynamicFilter(val *MccDynamicFilter) *NullableMccDynamicFilter {
	return &NullableMccDynamicFilter{value: val, isSet: true}
}

func (v NullableMccDynamicFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMccDynamicFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


